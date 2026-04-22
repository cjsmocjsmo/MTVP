package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	_ "github.com/mattn/go-sqlite3"
	"github.com/gorilla/websocket"
	"mtvp/handlers"
	"mtvp/mpvipc"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}

		log.Printf("Received: %s", msg)

		// Parse JSON message
		var data map[string]interface{}
		if err := json.Unmarshal(msg, &data); err != nil {
			log.Println("JSON parse error:", err)
			conn.WriteMessage(websocket.TextMessage, []byte(`{"status":"error","message":"invalid json"}`))
			continue
		}

		command, _ := data["command"].(string)

				       // ...existing code...
				       var dbPath = os.Getenv("MTV_DB_PATH")
				       db, err := sql.Open("sqlite3", dbPath)
				       if err != nil {
					       log.Printf("DB open error: %v", err)
					       conn.WriteMessage(websocket.TextMessage, []byte(`{"status":"error","message":"db open error"}`))
					       break
				       }
				       defer db.Close()

						switch command {
						// Utility commands
						case "movcount", "tvcount", "movsizeondisk", "tvsizeondisk", "weather", "test", "checkformovupdates", "checkfortvupdates":
							handlers.HandleUtilityCommand(conn, db, command)
						// TV/season and special category commands
						case "action", "arnold", "avatar", "brucelee", "brucewillis", "buzz", "cartoons", "charliebrown", "cheechandchong", "chucknorris", "clinteastwood", "comedy", "documentary", "drama", "fantasy", "ghostbusters", "godzilla", "harrisonford", "harrypotter", "hellboy", "indianajones", "jamesbond", "johnwayne", "johnwick", "jurassicpark", "kevincostner", "kingsman", "lego", "meninblack", "minions", "misc", "musicvids", "nature", "nicolascage", "oldies", "panda", "pirates", "predator", "riddick", "science", "scifi", "stalone", "startrek", "starwars", "stooges", "superheros", "superman", "therock", "tinkerbell", "tomcruize", "transformers", "tremors", "trolls", "vandam", "xmen",
							"alteredcarbons1", "alteredcarbons2", "columbias1", "cowboybebops1", "fallouts1", "fallouts2", "forallmankinds1", "forallmankinds2", "forallmankinds3", "forallmankinds4", "foundations1", "foundations2", "foundations3", "fuubars1", "fuubars2", "hford1923s1", "hford1923s2", "halos1", "halos2", "houseofthedragons1", "houseofthedragons2", "lostinspaces1", "lostinspaces2", "lostinspaces3", "mastersoftheuniverse", "monarchlegacyofmonsterss1", "monarchlegacyofmonsterss2", "nightskys1", "orvilles1", "orvilles2", "orvilles3", "percyjacksonandtheolympianss1", "percyjacksonandtheolympianss2", "prehistoricplanets1", "prehistoricplanets2", "prehistoricplanets3", "raisedbywolvess1", "raisedbywolvess2", "shoguns1", "silos1", "silos2", "thecontinentals1", "thelastofus1", "thelastofus2", "thelordoftheringstheringsofpowers1", "wheeloftimes1", "wheeloftimes2", "wheeloftimes3", "discoverys1", "discoverys2", "discoverys3", "discoverys4", "discoverys5", "enterprises1", "enterprises2", "enterprises3", "enterprises4", "enterprises5", "lowerdeckss1", "lowerdeckss2", "lowerdeckss3", "lowerdeckss4", "lowerdeckss5", "picards1", "picards2", "prodigys1", "prodigys2", "sttvs1", "sttvs2", "sttvs3", "strangenewworldss1", "strangenewworldss2", "strangenewworldss3", "tngs1", "tngs2", "tngs3", "tngs4", "tngs5", "tngs6", "tngs7", "voyagers1", "voyagers2", "voyagers3", "voyagers4", "voyagers5", "voyagers6", "voyagers7", "deepspacenines1", "deepspacenines2", "deepspacenines3", "deepspacenines4", "deepspacenines5", "deepspacenines6", "deepspacenines7", "continues1", "starfleetacademys1", "jetsonss1", "jetsonss2", "jonnyquests1":
							handlers.HandleSpecialCategoryCommand(conn, db, command)
						// MPV media playback/controls
						case "set_media":
							mediaID, _ := data["media_id"].(string)
							var mediaPath string
							if mediaID != "" {
								err := db.QueryRow("SELECT Path FROM movies WHERE MovId = ?", mediaID).Scan(&mediaPath)
								if err != nil || mediaPath == "" {
									conn.WriteMessage(websocket.TextMessage, []byte(`{"status":"error","message":"media not found"}`))
									break
								}
							} else {
								mediaPath, _ = data["media_path"].(string)
							}
							mpv, err := mpvipc.New("/tmp/mpvsocket")
							if err != nil {
								conn.WriteMessage(websocket.TextMessage, []byte(`{"status":"error","message":"mpv not available"}`))
								break
							}
							err = mpv.LoadFile(mediaPath)
							if err != nil {
								conn.WriteMessage(websocket.TextMessage, []byte(`{"status":"error","message":"failed to play"}`))
								break
							}
							conn.WriteMessage(websocket.TextMessage, []byte(`{"status":"media_set"}`))
						case "play":
							mpv, err := mpvipc.New("/tmp/mpvsocket")
							if err == nil {
								mpv.Play()
							}
						case "pause":
							mpv, err := mpvipc.New("/tmp/mpvsocket")
							if err == nil {
								mpv.Pause()
							}
						case "stop":
							mpv, err := mpvipc.New("/tmp/mpvsocket")
							if err == nil {
								mpv.Stop()
							}
						case "next":
							mpv, err := mpvipc.New("/tmp/mpvsocket")
							if err == nil {
								mpv.Seek(35)
							}
						case "previous":
							mpv, err := mpvipc.New("/tmp/mpvsocket")
							if err == nil {
								mpv.Seek(-35)
							}
						case "set_tv_media":
							mediaTVID, _ := data["media_tv_id"].(string)
							var mediaPath string
							if mediaTVID != "" {
								err := db.QueryRow("SELECT Path FROM tvshows WHERE TvId = ?", mediaTVID).Scan(&mediaPath)
								if err != nil || mediaPath == "" {
									conn.WriteMessage(websocket.TextMessage, []byte(`{"status":"error","message":"media not found"}`))
									break
								}
							} else {
								mediaPath, _ = data["media_path"].(string)
							}
							mpv, err := mpvipc.New("/tmp/mpvsocket")
							if err != nil {
								conn.WriteMessage(websocket.TextMessage, []byte(`{"status":"error","message":"mpv not available"}`))
								break
							}
							err = mpv.LoadFile(mediaPath)
							if err != nil {
								conn.WriteMessage(websocket.TextMessage, []byte(`{"status":"error","message":"failed to play"}`))
								break
							}
							conn.WriteMessage(websocket.TextMessage, []byte(`{"status":"media_set"}`))
						case "set_video_media":
							videoID, _ := data["video_id"].(string)
							var mediaPath string
							err := db.QueryRow("SELECT VidPath FROM videos WHERE VidId = ?", videoID).Scan(&mediaPath)
							if err != nil || mediaPath == "" {
								conn.WriteMessage(websocket.TextMessage, []byte(`{"status":"error","message":"video not found"}`))
								break
							}
							mpv, err := mpvipc.New("/tmp/mpvsocket")
							if err != nil {
								conn.WriteMessage(websocket.TextMessage, []byte(`{"status":"error","message":"mpv not available"}`))
								break
							}
							err = mpv.LoadFile(mediaPath)
							if err != nil {
								conn.WriteMessage(websocket.TextMessage, []byte(`{"status":"error","message":"failed to play"}`))
								break
							}
							conn.WriteMessage(websocket.TextMessage, []byte(`{"status":"media_set"}`))
						case "search":
							phrase, _ := data["phrase"].(string)
							if phrase != "" {
								results := handlers.SearchMedia(db, phrase)
								resp, _ := json.Marshal(results)
								conn.WriteMessage(websocket.TextMessage, resp)
							} else {
								conn.WriteMessage(websocket.TextMessage, []byte(`{"status":"error","message":"no search phrase"}`))
							}
						}
	}
}

func staticFileHandler(prefix, dir string) http.Handler {
	return http.StripPrefix(prefix, http.FileServer(http.Dir(dir)))
}


