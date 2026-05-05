from mpv import MPV
import asyncio
import websockets
import urllib.parse
import mimetypes
import requests
import json
import logging
import mtvserverutils
from dotenv import load_dotenv
import sqlite3
import os
import utils as UTILS
import mtvmovies as MTVMOVIES
import mtvtvshows as MTVTVSHOWS


# Initialize MPV player
player = MPV()

load_dotenv(dotenv_path="./env/.env")

# Ensure log directory and file exist
log_file = os.getenv('MTV_SERVER_LOG')
log_dir = os.path.dirname(log_file)

if not os.path.exists(log_dir):
    os.makedirs(log_dir, exist_ok=True)

# Create log file if it doesn't exist
if not os.path.exists(log_file):
    open(log_file, 'a').close()

logging.basicConfig(
    level=logging.INFO,
    filename=log_file,
    filemode='a',  # Append mode
    format='%(asctime)s - %(levelname)s - %(message)s'
)

# import time
MTVMEDIA = mtvserverutils.Media()

async def get_media_path_from_media_id(media_id):
    """
    Retrieves the media_path from the db using the media_id.
    """
    try:
        with sqlite3.connect(os.getenv("MTV_DB_PATH")) as conn:
            cursor = conn.cursor()
            cursor.execute("SELECT Path FROM movies WHERE MovId = ?", (media_id,))
            result = cursor.fetchone()
            if result:
                media_path = result[0]
                return media_path
            else:
                logging.error(f"No media path found for media_id: {media_id}")
                return None
    except sqlite3.Error as e:
        logging.error(f"SQLite error: {e}")
        return None
    except Exception as e:
        logging.error(f"Error fetching media path: {e}")
        return None

async def get_media_path_from_media_tv_id(media_tv_id):
    """
    Retrieves the media_path from the db using the media_tv_id.
    """
    try:
        with sqlite3.connect(os.getenv("MTV_DB_PATH")) as conn:
            cursor = conn.cursor()
            cursor.execute("SELECT Path FROM tvshows WHERE TvId = ?", (media_tv_id,))
            result = cursor.fetchone()
            if result:
                media_path = result[0]
                return media_path
            else:
                logging.error(f"No media path found for media_tv_id: {media_tv_id}")
                return None
    except sqlite3.Error as e:
        logging.error(f"SQLite error: {e}")
        return None
    except Exception as e:
        logging.error(f"Error fetching media path: {e}")
        return None  

async def get_media_path_from_video_id(video_id):
    """
    Retrieves the media_path from the db using the video_id.
    """
    try:
        with sqlite3.connect(os.getenv("MTV_DB_PATH")) as conn:
            cursor = conn.cursor()
            cursor.execute("SELECT VidPath FROM videos WHERE VidId = ?", (video_id,))
            result = cursor.fetchone()
            if result:
                media_path = result[0]
                return media_path
            else:
                logging.error(f"No media path found for video_id: {video_id}")
                return None
    except sqlite3.Error as e:
        logging.error(f"SQLite error: {e}")
        return None
    except Exception as e:
        logging.error(f"Error fetching media path: {e}")
        return None    

async def get_weather_for_belfair_wa():
    """
    Retrieves and prints the current weather conditions for Belfair, WA
    from the National Weather Service.
    """
    try:
        latitude = 47.4281
        longitude = -122.8189
        point_url = f"https://api.weather.gov/points/{latitude},{longitude}"
        point_response = requests.get(point_url)
        point_response.raise_for_status()  # Raise an exception for bad status codes
        point_data = point_response.json()
        forecast_url = point_data['properties']['forecastHourly']
        weather_response = requests.get(forecast_url)
        weather_response.raise_for_status()
        weather_data = weather_response.json()
        current_forecast = weather_data['properties']['periods'][0]

        weather_data = {
            "location": "Belfair, WA",
            "temperature": current_forecast['temperature'],
            "temperature_unit": current_forecast['temperatureUnit'],
            "conditions": current_forecast['shortForecast'],
            "winddirection": current_forecast['windDirection'],
            "windspeed": current_forecast['windSpeed']
        }

        return weather_data

    except requests.exceptions.RequestException as e:
        # print(f"Error fetching weather data: {e}")
        logging.error(f"Error fetching weather data: {e}")
        return None
    except (KeyError, IndexError) as e:
        # print(f"Error parsing weather data: {e}")
        logging.error(f"Error parsing weather data: {e}")
        return None
    except Exception as e:
        logging.error(f"Error fetching media path: {e}")
        return None      


# async def handle_message(websocket, path):
async def handle_message(websocket):
    try:
        async for message in websocket:
            data = json.loads(message)
            command = data.get("command")

            if command == "set_media":
                try:
                    media_id = data.get("media_id")
                    if media_id:
                        media_path = await get_media_path_from_media_id(media_id)
                        player.play(media_path)
                        player.fullscreen = True
                        player.volume = 100
                        logging.info(f"Starting mpv mediaplayer with the path: {media_path}")
                        await websocket.send(json.dumps({"status": "media_set"}))
                except Exception as e:
                    logging.error(f"Error fetching media path: {e}")
                    return None
             
            elif command == "set_tv_media":
                try:
                    media_tv_id = data.get("media_tv_id")
                    if media_tv_id:
                        media_path = await get_media_path_from_media_tv_id(media_tv_id)
                        player.play(media_path)
                        player.fullscreen = True
                        player.volume = 100
                        logging.info(f"Starting TV mpv mediaplayer with the path: {media_path}")
                        await websocket.send(json.dumps({"status": "media_set"}))
                except Exception as e:
                    logging.error(f"Error setting player path with mediapath: {e}")
                    return None

            elif command == "set_video_media":
                try:
                    video_id = data.get("video_id")
                    if video_id:
                        media_path = await get_media_path_from_video_id(video_id)
                        player.play(media_path)
                        player.fullscreen = True
                        player.volume = 100
                        logging.info(f"Starting video mpv mediaplayer with the path: {media_path}")
                        await websocket.send(json.dumps({"status": "media_set"}))
                except Exception as e:
                    logging.error(f"Error setting player path with mediapath: {e}")
                    return None

            elif command == "search":
                phrase = data.get("phrase")
                if phrase:
                    search_results = MTVMEDIA.mtvsearch(phrase)
                    await websocket.send(json.dumps(search_results))

            elif command == "stop":
                player.stop()
                await websocket.send(json.dumps({"status": "stopped"}))

            elif command == "play":
                player.play()
                await websocket.send(json.dumps({"status": "playing"}))
            
            elif command == "pause":
                 # Toggle pause using mpv property
                 player.pause = not getattr(player, 'pause', False)
                 await websocket.send(json.dumps({"status": "paused" if player.pause else "playing"}))

            elif command == "next":
                    # Seek forward 35 seconds using mpv command
                    try:
                        player.command('seek', 35, 'relative')
                        await websocket.send(json.dumps({"status": "next"}))
                    except Exception as e:
                        await websocket.send(json.dumps({"status": "error", "message": str(e)}))

            elif command == "previous":
                    # Seek backward 35 seconds using mpv command
                    try:
                        player.command('seek', -35, 'relative')
                        await websocket.send(json.dumps({"status": "previous"}))
                    except Exception as e:
                        await websocket.send(json.dumps({"status": "error", "message": str(e)}))
            
            elif command == "weather":
                weather_data = await get_weather_for_belfair_wa()
                await websocket.send(json.dumps(weather_data))

            elif command == "test":
                await websocket.send(json.dumps({"status": "Fuck it worked"}))

            elif command == "movcount":
                mov_count = UTILS.movie_count()
                await websocket.send(json.dumps(mov_count))

            elif command == "tvcount":
                tv_count = UTILS.tvshow_count()
                await websocket.send(json.dumps(tv_count))

            elif command == "movsizeondisk":
                movsizeondisk = UTILS.movies_size_on_disk()
                await websocket.send(json.dumps(movsizeondisk))

            elif command == "tvsizeondisk":
                tvsizeondisk = UTILS.tvshows_size_on_disk()
                await websocket.send(json.dumps(tvsizeondisk))

            elif command == "checkformovupdates":
                conn = sqlite3.connect(os.getenv("MTV_DB_PATH"))
                cursor = conn.cursor()
                update_data = MTVMOVIES.UpdateMovies(conn, cursor).check_for_mov_updates()
                await websocket.send(json.dumps(update_data))
                conn.close()

            elif command == "updatemovs":
                update_data = MTVMOVIES.UpdateMovies().updateMovs()
                await websocket.send(json.dumps(update_data))
                conn.close()

            elif command == "checkfortvupdates":
                conn = sqlite3.connect(os.getenv("MTV_DB_PATH"))
                cursor = conn.cursor()
                update_data = MTVTVSHOWS.UpdateTVShows(conn, cursor).check_for_tv_updates()
                await websocket.send(json.dumps(update_data))
                conn.close()

            elif command == "updatetvs":
                update_data = MTVTVSHOWS.UpdateTVShows().updateTV()
                await websocket.send(json.dumps(update_data))
                conn.close()

            elif hasattr(MTVMEDIA, command):
                method = getattr(MTVMEDIA, command)
                result = method()
                await websocket.send(json.dumps(result))
            
    except Exception as e:
        logging.error(f"Exception in handle_message: {e}")
    finally:
        logging.info("WebSocket connection closed")

async def servermain():
    async def static_file_server(reader, writer):
        request = await reader.readline()
        if not request:
            writer.close()
            await writer.wait_closed()
            return

        try:
            method, path, *_ = request.decode().split()
        except Exception:
            writer.close()
            await writer.wait_closed()
            return

        path = urllib.parse.unquote(path)
        static_paths = {
            '/thumbnails': '/usr/share/MTV/thumbnails',
            '/tvthumbnails': '/usr/share/MTV/tvthumbnails',
        }
        for prefix, folder in static_paths.items():
            if path.startswith(prefix):
                rel_path = path[len(prefix):].lstrip('/')
                file_path = os.path.join(folder, rel_path)
                if os.path.isfile(file_path):
                    mime, _ = mimetypes.guess_type(file_path)
                    mime = mime or 'application/octet-stream'
                    try:
                        with open(file_path, 'rb') as f:
                            content = f.read()
                        response = (
                            "HTTP/1.1 200 OK\r\n"
                            f"Content-Type: {mime}\r\n"
                            f"Content-Length: {len(content)}\r\n"
                            "Connection: close\r\n"
                            "\r\n"
                        ).encode() + content
                    except Exception:
                        response = b"HTTP/1.1 500 Internal Server Error\r\n\r\n"
                else:
                    response = b"HTTP/1.1 404 Not Found\r\n\r\n"
                writer.write(response)
                await writer.drain()
                writer.close()
                await writer.wait_closed()
                return

        # Not a static file request
        response = b"HTTP/1.1 404 Not Found\r\n\r\n"
        writer.write(response)
        await writer.drain()
        writer.close()
        await writer.wait_closed()

    # Start both websocket and static file servers concurrently
    ws_server = websockets.serve(handle_message, "10.0.4.41", 8765)
    static_server = await asyncio.start_server(static_file_server, "10.0.4.41", 8080)
    print("WebSocket server running on ws://10.0.4.41:8765/")
    print("Static file server running on http://10.0.4.41:8080/thumbnails/... and /tvthumbnails/...")
    async with static_server:
        await asyncio.gather(ws_server, static_server.serve_forever())

if __name__ == "__main__":
    asyncio.run(servermain())