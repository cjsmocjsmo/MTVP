package server

import (
    "os/exec"
    "sync"
    "net"
    "encoding/json"
    "time"
    "database/sql"
    "log"
    "github.com/gorilla/websocket"
    "html/template"
    "path/filepath"
    "net/http"
)

// PlayerManager manages the media player process and state
type PlayerManager struct {
    mu      sync.Mutex
    cmd     *exec.Cmd
    playing bool
    paused  bool
    ipcSock string
}

var player = &PlayerManager{ipcSock: "/tmp/mpvsocket"}

func (p *PlayerManager) StartMPV(path string) error {
    p.mu.Lock()
    defer p.mu.Unlock()
    if p.cmd != nil && p.playing {
        p.cmd.Process.Kill()
        time.Sleep(500 * time.Millisecond)
    }
    _ = exec.Command("rm", "-f", p.ipcSock).Run()
    p.cmd = exec.Command("mpv", "--fs", "--volume=100", "--input-ipc-server="+p.ipcSock, path)
    err := p.cmd.Start()
    if err == nil {
        p.playing = true
        p.paused = false
    }
    return err
}

func (p *PlayerManager) StopMPV() {
    p.mu.Lock()
    defer p.mu.Unlock()
    if p.cmd != nil && p.playing {
        p.cmd.Process.Kill()
        p.playing = false
        p.paused = false
    }
}

func (p *PlayerManager) sendMPVCommand(cmd interface{}) error {
    conn, err := net.Dial("unix", p.ipcSock)
    if err != nil {
        return err
    }
    defer conn.Close()
    b, _ := json.Marshal(cmd)
    _, err = conn.Write(append(b, '\n'))
    return err
}

func (p *PlayerManager) Pause() error {
    return p.sendMPVCommand(map[string]interface{}{
        "command": []interface{}{ "set_property", "pause", true },
    })
}

func (p *PlayerManager) Play() error {
    return p.sendMPVCommand(map[string]interface{}{
        "command": []interface{}{ "set_property", "pause", false },
    })
}

func (p *PlayerManager) Next() error {
    return p.sendMPVCommand(map[string]interface{}{
        "command": []interface{}{ "seek", 35, "relative" },
    })
}

func (p *PlayerManager) Previous() error {
    return p.sendMPVCommand(map[string]interface{}{
        "command": []interface{}{ "seek", -35, "relative" },
    })
}

func HandleWS(conn *websocket.Conn, db *sql.DB) {
    defer conn.Close()
    for {
        _, message, err := conn.ReadMessage()
        if err != nil {
            log.Println("WebSocket read error:", err)
            break
        }
        var data map[string]interface{}
        if err := json.Unmarshal(message, &data); err != nil {
            log.Println("JSON unmarshal error:", err)
            continue
        }
        command, _ := data["command"].(string)
        switch command {
        case "set_media":
            mediaID, _ := data["media_id"].(string)
            if mediaID != "" {
                var path string
                err := db.QueryRow("SELECT Path FROM movies WHERE MovId = ?", mediaID).Scan(&path)
                if err != nil {
                    sendJSON(conn, map[string]interface{}{ "status": "error", "message": "media not found" })
                } else {
                    if err := player.StartMPV(path); err != nil {
                        sendJSON(conn, map[string]interface{}{ "status": "error", "message": err.Error() })
                    } else {
                        sendJSON(conn, map[string]interface{}{ "status": "media_set" })
                    }
                }
            }
        case "stop":
            player.StopMPV()
            sendJSON(conn, map[string]interface{}{ "status": "stopped" })
        case "play":
            if err := player.Play(); err != nil {
                sendJSON(conn, map[string]interface{}{ "status": "error", "message": err.Error() })
            } else {
                sendJSON(conn, map[string]interface{}{ "status": "playing" })
            }
        case "pause":
            if err := player.Pause(); err != nil {
                sendJSON(conn, map[string]interface{}{ "status": "error", "message": err.Error() })
            } else {
                sendJSON(conn, map[string]interface{}{ "status": "paused" })
            }
        case "next":
            if err := player.Next(); err != nil {
                sendJSON(conn, map[string]interface{}{ "status": "error", "message": err.Error() })
            } else {
                sendJSON(conn, map[string]interface{}{ "status": "next" })
            }
        case "previous":
            if err := player.Previous(); err != nil {
                sendJSON(conn, map[string]interface{}{ "status": "error", "message": err.Error() })
            } else {
                sendJSON(conn, map[string]interface{}{ "status": "previous" })
            }
        case "search":
            phrase, _ := data["phrase"].(string)
            results := getCategoryMovieImages(db, phrase)
            resp := map[string]interface{}{ "results": results }
            sendJSON(conn, resp)
        case "movcount":
            count := getMovieCount(db)
            sendJSON(conn, map[string]interface{}{ "movcount": count })
        case "tvcount":
            count := getTVShowCount(db)
            sendJSON(conn, map[string]interface{}{ "tvcount": count })
        case "test":
            sendJSON(conn, map[string]interface{}{ "status": "It worked" })
        // Movie categories
        case "action", "arnold", "avatar", 
            "brucelee", "brucewillis", "buzz",
            "cartoons", "charliebrown", "cheechandchong", "chucknorris", "clinteastwood", "comedy",
            "drama", "documentary",
            "fantasy",
            "ghostbusters", "godzilla",
            "harrisonford", "harrypotter", "hellboy",
            "indianajones",
            "jamesbond", "johnwayne", "johnwick", "jurassicpark",
            "kevincostner", "kingsman",
            "lego",
            "meninblack", "minions", "misc", "musicvids", "mummy",
            "nature", "nicolascage",
            "oldies",
            "panda", "pirates", "predator",
            "riddick", 
            "science", "scifi", "stalone", "startrek", "starwars", "stooges", "superheros", "superman",
            "therock", "tinkerbell", "tomcruize", "transformers", "tremors", "trolls",
			"vandam",
			"xmen":
            results := getCategoryMovieImages(db, command)
            sendJSON(conn, map[string]interface{}{ "images": results })
        // TV show categories (full parity with mtvserverutils.py)
        case "alteredcarbons1", "alteredcarbons2",
			 "cowboybebop", "columbia",
			 "fallouts1", "fallouts2",
             "forallmankinds1", "forallmankinds2", "forallmankinds3", "forallmankinds4", "forallmankinds5",
             "foundations1", "foundations2", "foundations3",
			 "fuubar1", "fuubar2",
			 "hford19231", "hford19232",
             "halos1", "halos2",
		  	 "houseofthedragons1", "houseofthedragons2",
			 "lostinspaces1", "lostinspaces2", "lostinspaces3",
			 "mastersoftheuniverse",
			 "mobland",
			 "monarchlegacyofmonsterss1", "monarchlegacyofmonsterss2",
             "nightsky",
			 "orvilles1", "orvilles2", "orvilles3",
			 "percyjacksonandtheolympianss1", "percyjacksonandtheolympianss2",
			 "prehistoricplanets1", "prehistoricplanets2", "prehistoricplanets3",
             "raisedbywolvess1", "raisedbywolvess2",
			 "shogun",
			 "silo1", "silo2", 
			 "thecontinental",
			 "thelastofus1", "thelastofus2",
			 "thelordoftheringstheringsofpower",
			 "wheeloftime1", "wheeloftime2", "wheeloftime3",
			 "wednesdays1", "wednesdays2",

             "discoverys1", "discoverys2", "discoverys3", "discoverys4", "discoverys5",
             "enterprises1", "enterprises2", "enterprises3", "enterprises4",
             "lowerdeckss1", "lowerdeckss2", "lowerdeckss3", "lowerdeckss4", "lowerdeckss5",
             "picards1", "picards2",
             "prodigy1", "prodigy2",
             "sttvs1", "sttvs2", "sttvs3",
             "strangenewworldss1", "strangenewworldss2", "strangenewworldss3",
             "tngs1", "tngs2", "tngs3", "tngs4", "tngs5", "tngs6", "tngs7",
             "voyagers1", "voyagers2", "voyagers3", "voyagers4", "voyagers5", "voyagers6", "voyagers7",
             "deepspacenines1", "deepspacenines2", "deepspacenines3", "deepspacenines4", "deepspacenines5", "deepspacenines6", "deepspacenines7",
             "continues1", "starfleetacademys1",

             "acolyte", "ahsoka", "andor1", "andor2", "bookofbobafett",
             "mandalorians1", "mandalorians2", "mandalorians3", "obiwankenobi",
             "talesoftheempire", "talesofthejedi", "thebadbatchs1", "thebadbatchs2", "thebadbatchs3",
             "visionss1", "visionss2", "visionss3", "maulshadowlords1",

             "falconwintersoldier", "hawkeye", "iamgroots1", "iamgroots2",
             "lokis1", "lokis2", "moonknight", "secretinvasion", "shehulk", "wandavision",
             "skeletoncrew",  "talesoftheunderworld",  "ironheart", "wondermans1",
              
            "tonyandzivas1",
            "nciss1", "nciss2", "nciss3", "nciss4", "nciss5", "nciss6", "nciss7", "nciss8", "nciss9", "nciss10", "nciss11", "nciss12", "nciss13", "nciss14", "nciss15", "nciss16", "nciss17", "nciss18", "nciss19", "nciss20", "nciss21", "nciss22", "nciss23", "nciss24",
            "ncislas1", "ncislas2", "ncislas3", "ncislas4", "ncislas5", "ncislas6", "ncislas7", "ncislas8", "ncislas9", "ncislas10", "ncislas11", "ncislas12", "ncislas13", "ncislas14",
            "ncissydneys1", "ncissydneys2", "ncissydneys3",
            "ncisorigins1",
            "ncishawaiis1", "ncishawaiis2", "ncishawaiis3",
            "ncisneworleanss1", "ncisneworleanss2", "ncisneworleanss3", "ncisneworleanss4", 
            "ncisneworleanss5", "ncisneworleanss6", "ncisneworleanss7":
            results := getCategoryTVShowEpisodes(db, command)
            sendJSON(conn, map[string]interface{}{ "episodes": results })
        default:
            sendJSON(conn, map[string]interface{}{ "status": "unknown command" })
        }
    }
}

func getCategoryTVShowEpisodes(db *sql.DB, key string) []map[string]interface{} {
    var cat, season string
    switch key {
    case "nciss1":
        cat, season = "NCIS", "01"
    case "nciss2":
        cat, season = "NCIS", "02"
    case "nciss3":
        cat, season = "NCIS", "03"
    case "nciss4":
        cat, season = "NCIS", "04"
    case "nciss5":
        cat, season = "NCIS", "05"
    case "nciss6":
        cat, season = "NCIS", "06"
    case "nciss7":
        cat, season = "NCIS", "07"
    case "nciss8":
        cat, season = "NCIS", "08"
    case "nciss9":
        cat, season = "NCIS", "09"
    case "nciss10":
        cat, season = "NCIS", "10"
    case "nciss11":
        cat, season = "NCIS", "11"
    case "nciss12":
        cat, season = "NCIS", "12"
    case "nciss13":
        cat, season = "NCIS", "13"
    case "nciss14":
        cat, season = "NCIS", "14"
    case "nciss15":
        cat, season = "NCIS", "15"
    case "nciss16":
        cat, season = "NCIS", "16"
    case "nciss17":
        cat, season = "NCIS", "17"
    case "nciss18":
        cat, season = "NCIS", "18"
    case "nciss19":
        cat, season = "NCIS", "19"
    case "nciss20":
        cat, season = "NCIS", "20"
    case "nciss21":
        cat, season = "NCIS", "21"
    case "nciss22":
        cat, season = "NCIS", "22"
    case "nciss23":
        cat, season = "NCIS", "23"
    case "nciss24":
        cat, season = "NCIS", "24"
    case "ncislas1":
        cat, season = "NCISLA", "01"
    case "ncislas2":
        cat, season = "NCISLA", "02"
    case "ncislas3":
        cat, season = "NCISLA", "03"
    case "ncislas4":
        cat, season = "NCISLA", "04"
    case "ncislas5":
        cat, season = "NCISLA", "05"
    case "ncislas6":
        cat, season = "NCISLA", "06"
    case "ncislas7":
        cat, season = "NCISLA", "07"
    case "ncislas8":
        cat, season = "NCISLA", "08"
    case "ncislas9":
        cat, season = "NCISLA", "09"
    case "ncislas10":
        cat, season = "NCISLA", "10"
    case "ncislas11":
        cat, season = "NCISLA", "11"
    case "ncislas12":
        cat, season = "NCISLA", "12"
    case "ncislas13":
        cat, season = "NCISLA", "13"
    case "ncislas14":
        cat, season = "NCISLA", "14"
    case "ncissydneys1":
        cat, season = "NCISSydney", "01"
    case "ncissydneys2":
        cat, season = "NCISSydney", "02"
    case "ncissydneys3":
        cat, season = "NCISSydney", "03"
    case "ncisorigins1":
        cat, season = "NCISOrigins", "01"
    case "alteredcarbons1":
        cat, season = "AlteredCarbon", "01"
    case "alteredcarbons2":
        cat, season = "AlteredCarbon", "02"
    case "cowboybebop":
        cat = "CowboyBebop"
    case "fallouts1":
        cat, season = "Fallout", "01"
    case "fallouts2":
        cat, season = "Fallout", "02"
    case "forallmankinds1":
        cat, season = "ForAllManKind", "01"
    case "forallmankinds2":
        cat, season = "ForAllManKind", "02"
    case "forallmankinds3":
        cat, season = "ForAllManKind", "03"
    case "forallmankinds4":
        cat, season = "ForAllManKind", "04"
    case "forallmankinds5":
        cat, season = "ForAllManKind", "05"
    case "foundations1":
        cat, season = "Foundation", "01"
    case "foundations2":
        cat, season = "Foundation", "02"
    case "foundations3":
        cat, season = "Foundation", "03"
    case "fuubar1":
        cat, season = "FuuBar", "01"
    case "fuubar2":
        cat, season = "FuuBar", "02"
    case "hford19231":
        cat, season = "HFord1923", "01"
    case "hford19232":
        cat, season = "HFord1923", "02"
    case "halos1":
        cat, season = "Halo", "01"
    case "halos2":
        cat, season = "Halo", "02"
    case "houseofthedragons1":
        cat, season = "HouseOfTheDragon", "01"
    case "houseofthedragons2":
        cat, season = "HouseOfTheDragon", "02"
    case "lostinspaces1":
        cat, season = "LostInSpace", "01"
    case "lostinspaces2":
        cat, season = "LostInSpace", "02"
    case "lostinspaces3":
        cat, season = "LostInSpace", "03"
    case "mastersoftheuniverse":
        cat = "MastersOfTheUniverse"
    case "monarchlegacyofmonsterss1":
        cat = "MonarchLegacyOfMonsters"
    case "monarchlegacyofmonsterss2":
        cat, season = "MonarchLegacyOfMonsters", "02"
    case "nightsky":
        cat = "NightSky"
    case "orvilles1":
        cat, season = "Orville", "01"
    case "orvilles2":
        cat, season = "Orville", "02"
    case "orvilles3":
        cat, season = "Orville", "03"
    case "percyjacksonandtheolympianss1":
        cat, season = "PercyJacksonAndTheOlympians", "01"
    case "percyjacksonandtheolympianss2":
        cat, season = "PercyJacksonAndTheOlympians", "02"
    case "prehistoricplanets1":
        cat, season = "PrehistoricPlanet", "01"
    case "prehistoricplanets2":
        cat, season = "PrehistoricPlanet", "02"
    case "prehistoricplanets3":
        cat, season = "PrehistoricPlanet", "03"
    case "raisedbywolvess1":
        cat, season = "RaisedByWolves", "01"
    case "raisedbywolvess2":
        cat, season = "RaisedByWolves", "02"
    case "shogun":
        cat = "Shogun"
    case "silo1":
        cat, season = "Silo", "01"
    case "silo2":
        cat, season = "Silo", "02"
    case "thecontinental":
        cat = "TheContinental"
    case "thelastofus1":
        cat, season = "TheLastOfUs", "01"
    case "thelastofus2":
        cat, season = "TheLastOfUs", "02"
    case "thelordoftheringstheringsofpower":
        cat = "TheLordOfTheRingsTheRingsOfPower"
    case "wheeloftime1":
        cat, season = "WheelOfTime", "01"
    case "wheeloftime2":
        cat, season = "WheelOfTime", "02"
    case "wheeloftime3":
        cat, season = "WheelOfTime", "03"
    case "discoverys1":
        cat, season = "Discovery", "01"
    case "discoverys2":
        cat, season = "Discovery", "02"
    case "discoverys3":
        cat, season = "Discovery", "03"
    case "discoverys4":
        cat, season = "Discovery", "04"
    case "discoverys5":
        cat, season = "Discovery", "05"
    case "enterprises1":
        cat, season = "Enterprise", "01"
    case "enterprises2":
        cat, season = "Enterprise", "02"
    case "enterprises3":
        cat, season = "Enterprise", "03"
    case "enterprises4":
        cat, season = "Enterprise", "04"
    case "lowerdeckss1":
        cat, season = "LowerDecks", "01"
    case "lowerdeckss2":
        cat, season = "LowerDecks", "02"
    case "lowerdeckss3":
        cat, season = "LowerDecks", "03"
    case "lowerdeckss4":
        cat, season = "LowerDecks", "04"
    case "lowerdeckss5":
        cat, season = "LowerDecks", "05"
    }
    var rows *sql.Rows
    var err error
    if season != "" {
        rows, err = db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", cat, season)
    } else {
        rows, err = db.Query("SELECT * FROM tvshows WHERE catagory=? ORDER BY Episode ASC", cat)
    }
    if err != nil {
        log.Println("DB error (tvshow episodes):", err)
        return nil
    }
    defer rows.Close()
    cols, _ := rows.Columns()
    results := []map[string]interface{}{}
    for rows.Next() {
        vals := make([]interface{}, len(cols))
        valPtrs := make([]interface{}, len(cols))
        for i := range vals {
            valPtrs[i] = &vals[i]
        }
        if err := rows.Scan(valPtrs...); err == nil {
            row := make(map[string]interface{})
            for i, col := range cols {
                b, ok := vals[i].([]byte)
                if ok {
                    row[col] = string(b)
                } else {
                    row[col] = vals[i]
                }
            }
            results = append(results, row)
        }
    }
    return results
}



func sendJSON(conn *websocket.Conn, v interface{}) {
    msg, err := json.Marshal(v)
    if err != nil {
        log.Println("JSON marshal error:", err)
        return
    }
    if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
        log.Println("WebSocket write error:", err)
    }
}


type WSResponse map[string]interface{}

// getCategoryMovieImages queries the DB for movies in a given category and returns a list of image URLs
func getCategoryMovieImages(db *sql.DB, category string) []string {
    query := "SELECT Path FROM movies WHERE LOWER(Title) LIKE ?"
    likePattern := "%" + category + "%"
    rows, err := db.Query(query, likePattern)
    if err != nil {
        log.Println("DB error (category images):", err)
        return nil
    }
    defer rows.Close()
    var images []string
    for rows.Next() {
        var path string
        if err := rows.Scan(&path); err == nil {
            images = append(images, "/thumbnails/"+path)
        }
    }
    return images
}


// getActionMovieImages queries the DB for action movies and returns a list of image URLs
func getActionMovieImages(db *sql.DB) []string {
    rows, err := db.Query("SELECT Path FROM movies WHERE Title LIKE '%Action%'")
    if err != nil {
        log.Println("DB error (action images):", err)
        return nil
    }
    defer rows.Close()
    var images []string
    for rows.Next() {
        var path string
        if err := rows.Scan(&path); err == nil {
            images = append(images, "/thumbnails/"+path)
        }
    }
    return images
}


func getMovieCount(db *sql.DB) int {
    var count int
    err := db.QueryRow("SELECT COUNT(*) FROM movies").Scan(&count)
    if err != nil {
        log.Println("DB error (movcount):", err)
        return 0
    }
    return count
}


func getTVShowCount(db *sql.DB) int {
    var count int
    err := db.QueryRow("SELECT COUNT(*) FROM tvshows").Scan(&count)
    if err != nil {
        log.Println("DB error (tvcount):", err)
        return 0
    }
    return count
}

// indexHandler serves the main page using Go's html/template
func indexHandlerWithPath(tmplPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// Legacy handler for production use
func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexHandlerWithPath(filepath.Join("go", "templates", "index.html"))(w, r)
}
