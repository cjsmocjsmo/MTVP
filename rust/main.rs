use serde::{Deserialize, Serialize};
use serde_json::json;
use tokio::net::TcpListener;
use tokio_tungstenite::accept_async;
use futures_util::{StreamExt, SinkExt};
use rusqlite::{params, Connection};
use std::env;
use std::path::Path;
use log::{info, error, LevelFilter};
use simple_logger::SimpleLogger;
use dotenv::dotenv;

// Replaces your MPV initialization
// Note: You will need libmpv headers installed on your Debian system
use mpv_client::{MpvClient, MpvClientBuilder}; 

#[derive(Deserialize)]
struct Command {
    command: String,
    media_id: Option<i32>,
    media_tv_id: Option<i32>,
    video_id: Option<i32>,
    phrase: Option<String>,
}

#[tokio::main]
async fn main() {
    dotenv().ok();
    
    // Logging setup
    SimpleLogger::new()
        .with_level(LevelFilter::Info)
        .init()
        .unwrap();

    let addr = "127.0.0.1:8080";
    let listener = TcpListener::bind(&addr).await.expect("Failed to bind");
    info!("Listening on: {}", addr);

    // Initialize MPV
    let mut mpv = MpvClientBuilder::new()
        .expect("Failed to create MPV builder")
        .build()
        .expect("Failed to build MPV client");

    while let Ok((stream, _)) = listener.accept().await {
        let mpv_ref = mpv.clone(); // In a real app, use a shared state handler
        tokio::spawn(handle_connection(stream, mpv_ref));
    }
}

async fn get_media_path(id: i32, table: &str, column: &str) -> Option<String> {
    let db_path = env::var("MTV_DB_PATH").ok()?;
    let conn = Connection::open(db_path).ok()?;
    
    let query = format!("SELECT {} FROM {} WHERE {} = ?", 
        if table == "videos" { "VidPath" } else { "Path" }, 
        table, 
        column
    );

    conn.query_row(&query, params![id], |row| row.get(0)).ok()
}

async fn handle_connection(stream: tokio::net::TcpStream, mpv: MpvClient) {
    let mut ws_stream = accept_async(stream).await.expect("Error during ws handshake");
    info!("New WebSocket connection");

    while let Some(msg) = ws_stream.next().await {
        let msg = msg.expect("Error reading message");
        if msg.is_text() {
            let text = msg.to_text().unwrap();
            if let Ok(cmd) = serde_json::from_str::<Command>(text) {
                match cmd.command.as_str() {
                    "set_media" => {
                        if let Some(id) = cmd.media_id {
                            if let Some(path) = get_media_path(id, "movies", "MovId").await {
                                let _ = mpv.command(&["loadfile", &path]);
                                let _ = ws_stream.send(serde_json::to_string(&json!({"status": "media_set"})).unwrap().into()).await;
                            }
                        }
                    },
                    "stop" => {
                        let _ = mpv.command(&["stop"]);
                        let _ = ws_stream.send(serde_json::to_string(&json!({"status": "stopped"})).unwrap().into()).await;
                    },
                    "pause" => {
                        // Toggle logic
                        let current_pause: bool = mpv.get_property("pause").unwrap_or(false);
                        let _ = mpv.set_property("pause", !current_pause);
                    },
                    "weather" => {
                        if let Ok(weather) = fetch_weather().await {
                             let _ = ws_stream.send(serde_json::to_string(&weather).unwrap().into()).await;
                        }
                    },
                    _ => info!("Unknown command received"),
                }
            }
        }
    }
}

async fn fetch_weather() -> Result<serde_json::Value, reqwest::Error> {
    let client = reqwest::Client::new();
    // National Weather Service requires a User-Agent header
    let res = client.get("https://api.weather.gov/points/47.4281,-122.8189")
        .header("User-Agent", "MTV-Server-Rust")
        .send()
        .await?
        .json::<serde_json::Value>()
        .await?;
        
    let forecast_url = res["properties"]["forecastHourly"].as_str().unwrap();
    let forecast = client.get(forecast_url)
        .header("User-Agent", "MTV-Server-Rust")
        .send()
        .await?
        .json::<serde_json::Value>()
        .await?;
        
    Ok(forecast["properties"]["periods"][0].clone())
}