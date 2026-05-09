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
	dbPath := os.Getenv("MTV_DB_PATH")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Failed to open DB:", err)
	}
	defer db.Close()

	wsAddr := os.Getenv("MTV_SERVER_ADDR")
	if wsAddr == "" {
		wsAddr = "0.0.0.0"
	}
	log.Println("Starting WebSocket server on ws://" + wsAddr + ":8765/ws")
	log.Println("Starting static file server on http://" + wsAddr + ":8080/thumbnails/ and /tvthumbnails/")

	http.HandleFunc("/ws", wsHandler(db))
	http.Handle("/thumbnails/", staticFileHandler("/thumbnails/", "/usr/share/MTV/thumbnails"))
	http.Handle("/tvthumbnails/", staticFileHandler("/tvthumbnails/", "/usr/share/MTV/tvthumbnails"))
	http.Handle("/templates/", staticFileHandler("/templates/", "go/templates/"))
	http.Handle("/static/css/", staticFileHandler("/static/css/", "go/static/css/"))
	http.Handle("/static/js/", staticFileHandler("/static/js/", "go/static/js/"))
	// Register /action route for action movies page
	http.HandleFunc("/action", ActionPageHandler(db))
	// Register /arnold route for Arnold movies page
	http.HandleFunc("/arnold", ArnoldPageHandler(db))
	// Register /avatar route for Avatar movies page
	http.HandleFunc("/avatar", AvatarPageHandler(db))
	// Register /brucelee route for Bruce Lee movies page
	http.HandleFunc("/brucelee", BruceLeePageHandler(db))
	// Register /brucewillis route for Bruce Willis movies page
	http.HandleFunc("/brucewillis", BruceWillisPageHandler(db))
	// Register /buzz route for Buzz movies page
	http.HandleFunc("/buzz", BuzzPageHandler(db))
	// Register /cartoons route for Cartoons movies page
	http.HandleFunc("/cartoons", CartoonsPageHandler(db))
	// Register /charliebrown route for Charlie Brown movies page
	http.HandleFunc("/charliebrown", CharlieBrownPageHandler(db))
	// Register /cheechandchong route for Cheech and Chong movies page
	http.HandleFunc("/cheechandchong", CheechAndChongPageHandler(db))
	// Register /chucknorris route for Chuck Norris movies page
	http.HandleFunc("/chucknorris", ChuckNorrisPageHandler(db))
	// Register /clinteastwood route for Clint Eastwood movies page
	http.HandleFunc("/clinteastwood", ClintEastwoodPageHandler(db))
	// Register /comedy route for Comedy movies page
	http.HandleFunc("/comedy", ComedyPageHandler(db))
	// Register /documentary route for Documentary movies page
	http.HandleFunc("/documentary", DocumentaryPageHandler(db))
	// Register /drama route for Drama movies page
	http.HandleFunc("/drama", DramaPageHandler(db))
	// Register /fantasy route for Fantasy movies page
	http.HandleFunc("/fantasy", FantasyPageHandler(db))
	// Register /ghostbusters route for Ghostbusters movies page
	http.HandleFunc("/ghostbusters", GhostbustersPageHandler(db))
	// Register /godzilla route for Godzilla movies page
	http.HandleFunc("/godzilla", GodzillaPageHandler(db))
	// Register /harrisonford route for Harrison Ford movies page
	http.HandleFunc("/harrisonford", HarrisonFordPageHandler(db))
	// Register /harrypotter route for Harry Potter movies page
	http.HandleFunc("/harrypotter", HarryPotterPageHandler(db))
	// Register /hellboy route for Hellboy movies page
	http.HandleFunc("/hellboy", HellboyPageHandler(db))
	// Register /homevids route for Home Vids movies page
	http.HandleFunc("/homevids", HomeVidsPageHandler(db))
	// Register /indianajones route for Indiana Jones movies page
	http.HandleFunc("/indianajones", IndianaJonesPageHandler(db))
	// Register /jamesbond route for James Bond movies page
	http.HandleFunc("/jamesbond", JamesBondPageHandler(db))
	// Register /johnwayne route for John Wayne movies page
	http.HandleFunc("/johnwayne", JohnWaynePageHandler(db))
	// Register /johnwick route for John Wick movies page
	http.HandleFunc("/johnwick", JohnWickPageHandler(db))
	// Register /jurrasicpark route for Jurassic Park movies page
	http.HandleFunc("/jurrasicpark", JurrasicParkPageHandler(db))
	// Register /kevincostner route for Kevin Costner movies page
	http.HandleFunc("/kevincostner", KevinCostnerPageHandler(db))
	// Register /kingsman route for Kingsman movies page
	http.HandleFunc("/kingsman", KingsmanPageHandler(db))
	// Register /lego route for Lego movies page
	http.HandleFunc("/lego", LegoPageHandler(db))
	// Register /meninblack route for Men in Black movies page
	http.HandleFunc("/meninblack", MenInBlackPageHandler(db))
	// Register /minions route for Minions movies page
	http.HandleFunc("/minions", MinionsPageHandler(db))

	go func() {
		if err := http.ListenAndServe(wsAddr+":8765", nil); err != nil {
			log.Fatal("WebSocket server error:", err)
		}
	}()

	if err := http.ListenAndServe(wsAddr+":8080", nil); err != nil {
		log.Fatal("Static file server error:", err)
	}
}
