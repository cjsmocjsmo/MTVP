	
package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strings"
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
		log.Printf("[staticFileHandler] Initializing for prefix: '%s', dir: '%s'", prefix, dir)
		fs := http.FileServer(http.Dir(dir))
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("[staticFileHandler] Incoming request: Method=%s, URL=%s, RemoteAddr=%s, UserAgent=%s", r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
			relPath := strings.TrimPrefix(r.URL.Path, prefix)
			filePath := dir
			if !strings.HasSuffix(dir, "/") && relPath != "" {
				filePath += "/"
			}
			filePath += relPath
			log.Printf("[staticFileHandler] Mapped URL '%s' to file path '%s'", r.URL.Path, filePath)
			fileInfo, err := os.Stat(filePath)
			if err != nil {
				log.Printf("[staticFileHandler] File not found or error: %v", err)
			} else {
				log.Printf("[staticFileHandler] File exists: %s (size: %d bytes, mode: %s)", filePath, fileInfo.Size(), fileInfo.Mode())
			}
			fs.ServeHTTP(w, r)
			log.Printf("[staticFileHandler] Finished serving: %s", r.URL.Path)
		})
}

func StartServer() {

	dbPath := os.Getenv("MTVGO_DB_PATH")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Failed to open DB:", err)
	}
	defer db.Close()

		// Serve static files from templates at root
		//http.Handle("/", staticFileHandler("/", "templates"))
		// Serve /static/ from static for CSS, JS, etc.
	log.Println("[StartServer] Creating static file handler for /static/ route (using http.StripPrefix)")
	staticDir := "./static/"
	absStaticDir, err := os.Getwd()
	if err == nil {
		absStaticDir = absStaticDir + "/static"
	} else {
		absStaticDir = staticDir
	}
	log.Printf("[StartServer] Using static directory: %s", absStaticDir)
	handler := http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir)))
	http.Handle("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[staticFileHandler] Incoming request: Method=%s, URL=%s, RemoteAddr=%s, UserAgent=%s", r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
		relPath := strings.TrimPrefix(r.URL.Path, "/static/")
		filePath := staticDir
		if !strings.HasSuffix(staticDir, "/") && relPath != "" {
			filePath += "/"
		}
		filePath += relPath
		log.Printf("[staticFileHandler] Mapped URL '%s' to file path '%s'", r.URL.Path, filePath)
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			log.Printf("[staticFileHandler] File not found or error: %v", err)
		} else {
			log.Printf("[staticFileHandler] File exists: %s (size: %d bytes, mode: %s)", filePath, fileInfo.Size(), fileInfo.Mode())
		}
		handler.ServeHTTP(w, r)
		log.Printf("[staticFileHandler] Finished serving: %s", r.URL.Path)
	}))
	log.Println("[StartServer] Static file handler registered for /static/")

	// Serve movie thumbnails from /usr/share/MTV/gothumbnails/ at /thumbnails
	movThumbPath := os.Getenv("MTVGO_THUMBNAIL_PATH")
	http.Handle("/thumbnails/", http.StripPrefix("/thumbnails/", http.FileServer(http.Dir(movThumbPath))))
	log.Println("[StartServer] File server registered for /thumbnails ->", movThumbPath)

	// Serve TV thumbnails from /usr/share/MTV/gotvthumbnails/ at /tvthumbnails
	tvThumbPath := os.Getenv("MTVGO_TV_THUMBNAIL_PATH")
	http.Handle("/tvthumbnails/", http.StripPrefix("/tvthumbnails/", http.FileServer(http.Dir(tvThumbPath))))
	log.Println("[StartServer] File server registered for /tvthumbnails ->", tvThumbPath)

	http.HandleFunc("/", HomePageHandler())
	http.HandleFunc("/mainmoviepage", MainMoviePageHandler())
	http.HandleFunc("/maintvpage", MainTVPageHandler(db))
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
	http.HandleFunc("/ghostbusters", GhostBustersPageHandler(db))
	http.HandleFunc("/godzilla", GodzillaPageHandler(db))
	http.HandleFunc("/harrisonford", HarrisonFordPageHandler(db))
	http.HandleFunc("/harrypotter", HarryPotterPageHandler(db))
	http.HandleFunc("/hellboy", HellboyPageHandler(db))
	http.HandleFunc("/homevids", HomeVidsPageHandler(db))
	http.HandleFunc("/indianajones", IndianaJonesPageHandler(db))
	http.HandleFunc("/jamesbond", JamesBondPageHandler(db))
	http.HandleFunc("/johnwayne", JohnWaynePageHandler(db))
	http.HandleFunc("/johnwick", JohnWickPageHandler(db))
	http.HandleFunc("/jurassicpark", JurassicParkPageHandler(db))
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
	http.HandleFunc("/superheros", SuperHerosPageHandler(db))
	http.HandleFunc("/superman", SuperManPageHandler(db))
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
	http.HandleFunc("/tvactionpage", TVActionPageHandler())
	http.HandleFunc("/tvactiondarkwinds", TVDarkWindsPageHandler(db))
	http.HandleFunc("/tvactionmobland", TVMoblandPageHandler(db))
	http.HandleFunc("/tvactionshogun", TVShogunPageHandler(db))
	http.HandleFunc("/tvactionthecontinental", TVTheContinentalPageHandler(db))
	
	http.HandleFunc("/tvcartoonspage", TVCartoonsPageHandler())
	http.HandleFunc("/tvcartoonsmastersoftheuniverse", TVMastersOfTheUniversePageHandler(db))
	http.HandleFunc("/tvcartoonsflintstones", TVFlintstonesPageHandler(db))
	http.HandleFunc("/tvcartoonsjetsons", TVJetsonsPageHandler(db))
	http.HandleFunc("/tvcartoonsjonnyquest", TVJonnyQuestPageHandler(db))

	http.HandleFunc("/tvcomedypage", TVComedyPageHandler())
	http.HandleFunc("/tvcomedydmvpage", TVDMVPageHandler(db))
	http.HandleFunc("/tvcomedyfubarpage", TVFubarPageHandler(db))

	http.HandleFunc("/tvfantasypage", TVFantasyPageHandler())
	http.HandleFunc("/tvfantasylordoftherings", TVLordOfTheRingsPageHandler(db))
	http.HandleFunc("/tvfantasyhouseofthedragon", TVHouseOfTheDragonPageHandler(db))
	http.HandleFunc("/tvfantasypercyjacksonandtheolympians", TVPercyJacksonAndTheOlympiansPageHandler(db))
	http.HandleFunc("/tvfantasywednesday", TVWednesdayPageHandler(db))
	http.HandleFunc("/tvfantasywheeloftime", TVWheelOfTimePageHandler(db))



	wsAddr := os.Getenv("MTVGO_RAW_ADDR")
	go func() {
	    var port = os.Getenv("MTVGO_SERVER_PORT")
	    println(wsAddr+":"+port)
	    if err := http.ListenAndServe(wsAddr+":"+port, nil); err != nil {
	       log.Fatal("WebSocket server error:", err)
	    }
	}()

	// Block forever
	select {}
}
