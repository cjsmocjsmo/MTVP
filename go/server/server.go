package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	_ "github.com/mattn/go-sqlite3"
	"github.com/gorilla/websocket"
)

func wsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool { return true },
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("WebSocket upgrade error:", err)
			return
		}
		HandleWS(conn, db)
	}
}

func staticFileHandler(prefix, dir string) http.Handler {
	return http.StripPrefix(prefix, http.FileServer(http.Dir(dir)))
}

func StartServer() {

	dbPath := os.Getenv("MTVGO_DB_PATH")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Failed to open DB:", err)
	}
	defer db.Close()

		// Serve static files from go/templates at root
		http.Handle("/", staticFileHandler("/", "go/templates"))

	http.HandleFunc("/action", ActionPageHandler(db))
	http.HandleFunc("/arnold", ArnoldPageHandler(db))
	http.HandleFunc("/avatar", AvatarPageHandler(db))
	http.HandleFunc("/brucelee", BruceLeePageHandler(db))
	http.HandleFunc("/brucewillis", BruceWillisPageHandler(db))
	http.HandleFunc("/buzz", BuzzPageHandler(db))
	http.HandleFunc("/cartoons", CartoonsPageHandler(db))
	http.HandleFunc("/charliebrown", CharlieBrownPageHandler(db))
	http.HandleFunc("/cheechandchong", CheechAndChongPageHandler(db))
	http.HandleFunc("/chucknorris", ChuckNorrisPageHandler(db))
	http.HandleFunc("/clinteastwood", ClintEastwoodPageHandler(db))
	http.HandleFunc("/comedy", ComedyPageHandler(db))
	http.HandleFunc("/documentary", DocumentaryPageHandler(db))
	http.HandleFunc("/drama", DramaPageHandler(db))
	http.HandleFunc("/fantasy", FantasyPageHandler(db))
	http.HandleFunc("/ghostbusters", GhostbustersPageHandler(db))
	http.HandleFunc("/godzilla", GodzillaPageHandler(db))
	http.HandleFunc("/harrisonford", HarrisonFordPageHandler(db))
	http.HandleFunc("/harrypotter", HarryPotterPageHandler(db))
	http.HandleFunc("/hellboy", HellboyPageHandler(db))
	http.HandleFunc("/homevids", HomeVidsPageHandler(db))
	http.HandleFunc("/indianajones", IndianaJonesPageHandler(db))
	http.HandleFunc("/jamesbond", JamesBondPageHandler(db))
	http.HandleFunc("/johnwayne", JohnWaynePageHandler(db))
	http.HandleFunc("/johnwick", JohnWickPageHandler(db))
	http.HandleFunc("/jurrasicpark", JurrasicParkPageHandler(db))
	http.HandleFunc("/kevincostner", KevinCostnerPageHandler(db))
	http.HandleFunc("/kingsman", KingsmanPageHandler(db))
	http.HandleFunc("/lego", LegoPageHandler(db))
	http.HandleFunc("/meninblack", MenInBlackPageHandler(db))
	http.HandleFunc("/minions", MinionsPageHandler(db))
	http.HandleFunc("/misc", MiscPageHandler(db))
	http.HandleFunc("/mummy", MummyPageHandler(db))
	http.HandleFunc("/musicvids", MusicVidsPageHandler(db))
	http.HandleFunc("/nature", NaturePageHandler(db))
	http.HandleFunc("/nicolascage", NicolasCagePageHandler(db))
	http.HandleFunc("/oldies", OldiesPageHandler(db))
	http.HandleFunc("/pandas", PandasPageHandler(db))
	http.HandleFunc("/pirates", PiratesPageHandler(db))
	http.HandleFunc("/predator", PredatorPageHandler(db))
	http.HandleFunc("/riddick", RiddickPageHandler(db))
	http.HandleFunc("/scifi", SciFiPageHandler(db))
	http.HandleFunc("/science", SciencePageHandler(db))
	http.HandleFunc("/starwars", StarWarsPageHandler(db))
	http.HandleFunc("/stalone", StalonePageHandler(db))
	http.HandleFunc("/startrek", StarTrekPageHandler(db))
	http.HandleFunc("/stooges", StoogesPageHandler(db))
	http.HandleFunc("/superheroes", SuperheroesPageHandler(db))
	http.HandleFunc("/superman", SupermanPageHandler(db))
	http.HandleFunc("/therock", TheRockPageHandler(db))
	http.HandleFunc("/tinkerbell", TinkerbellPageHandler(db))
	http.HandleFunc("/tomcruize", TomCruizePageHandler(db))
	http.HandleFunc("/tremors", TremorsPageHandler(db))
	http.HandleFunc("/transformers", TransformersPageHandler(db))
	http.HandleFunc("/trolls", TrollsPageHandler(db))
	http.HandleFunc("/vandam", VandamPageHandler(db))
	http.HandleFunc("/ws", wsHandler(db))
	http.HandleFunc("/xmen", XmenPageHandler(db))
		// Register TV show pages
	http.HandleFunc("/tvactiondarkwinds", DarkWindsPageHandler(db))
	http.HandleFunc("/tvactionmobland", MoblandPageHandler(db))
	http.HandleFunc("/tvactionshogun", ShogunPageHandler(db))
	http.HandleFunc("/tvactionthecontinental", TheContinentalPageHandler(db))

	// wsAddr := os.Getenv("MTVGO_RAW_ADDR")
	// if wsAddr == "" {
	// 	wsAddr = "0.0.0.0"
	// }
	// // Remove protocol if present
	// cleanAddr := wsAddr
	//     //    if len(cleanAddr) > 7 && cleanAddr[:7] == "http://" {
	// 	//        cleanAddr = cleanAddr[7:]
	//     //    } else if len(cleanAddr) > 8 && cleanAddr[:8] == "https://" {
	// 	//        cleanAddr = cleanAddr[8:]
	//     //    }
	// log.Println("Starting WebSocket server on ws://" + cleanAddr + ":8765/ws")
	// log.Println("Starting static file server on http://" + cleanAddr + ":8080/thumbnails/ and /tvthumbnails/")
	
	wsAddr := os.Getenv("MTVGO_RAW_ADDR")
	go func() {
		var port = os.Getenv("MTVGO_SERVER_PORT")
		println(wsAddr+":"+port)
		if err := http.ListenAndServe(wsAddr+":"+port, nil); err != nil {
			log.Fatal("WebSocket server error:", err)
		}
	}()

	var staticPort = os.Getenv("MTVGO_STATIC_SERVER_PORT")
	println(wsAddr+":"+staticPort)
	if err := http.ListenAndServe(wsAddr+":"+staticPort, nil); err != nil {
		log.Fatal("Static file server error:", err)
	}
}
