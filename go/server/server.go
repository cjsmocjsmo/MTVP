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

	wsAddr := os.Getenv("MTVGO_SERVER_ADDR")
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
	// Register /misc route for Misc movies page
	http.HandleFunc("/misc", MiscPageHandler(db))
	// Register /mummy route for Mummy movies page
	http.HandleFunc("/mummy", MummyPageHandler(db))
	// Register /musicvids route for Music Videos page
	http.HandleFunc("/musicvids", MusicVidsPageHandler(db))
	// Register /xmen route for X-Men movies page
	http.HandleFunc("/xmen", XmenPageHandler(db))
	// Register /vandam route for Van Damme movies page
	http.HandleFunc("/vandam", VandamPageHandler(db))
	// Register /trolls route for Trolls movies page
	http.HandleFunc("/trolls", TrollsPageHandler(db))
	// Register /tremors route for Tremors movies page
	http.HandleFunc("/tremors", TremorsPageHandler(db))
	// Register /transformers route for Transformers movies page
	http.HandleFunc("/transformers", TransformersPageHandler(db))
	// Register /scifi route for Sci-Fi movies page
	http.HandleFunc("/scifi", SciFiPageHandler(db))
	// Register /science route for Science movies page
	http.HandleFunc("/science", SciencePageHandler(db))
	// Register /nature route for Nature movies page
	http.HandleFunc("/nature", NaturePageHandler(db))
	// Register /nicolascage route for Nicolas Cage movies page
	http.HandleFunc("/nicolascage", NicolasCagePageHandler(db))
	// Register /oldies route for Oldies movies page
	http.HandleFunc("/oldies", OldiesPageHandler(db))
	// Register /pandas route for Pandas movies page
	http.HandleFunc("/pandas", PandasPageHandler(db))
	// Register /pirates route for Pirates movies page
	http.HandleFunc("/pirates", PiratesPageHandler(db))
	// Register /predator route for Predator movies page
	http.HandleFunc("/predator", PredatorPageHandler(db))
	// Register /riddick route for Riddick movies page
	http.HandleFunc("/riddick", RiddickPageHandler(db))
	// Register /stalone route for Stallone movies page
	http.HandleFunc("/stalone", StalonePageHandler(db))
	// Register /startrek route for Star Trek movies page
	http.HandleFunc("/startrek", StarTrekPageHandler(db))
	// Register /starwars route for Star Wars movies page
	http.HandleFunc("/starwars", StarWarsPageHandler(db))
	// Register /stooges route for Stooges movies page
	http.HandleFunc("/stooges", StoogesPageHandler(db))
	// Register /superheroes route for Superheroes movies page
	http.HandleFunc("/superheroes", SuperheroesPageHandler(db))
	// Register /superman route for Superman movies page
	http.HandleFunc("/superman", SupermanPageHandler(db))
	// Register /therock route for The Rock movies page
	http.HandleFunc("/therock", TheRockPageHandler(db))
	// Register /tinkerbell route for Tinkerbell movies page
	http.HandleFunc("/tinkerbell", TinkerbellPageHandler(db))
	// Register /tomcruize route for Tom Cruise movies page
	http.HandleFunc("/tomcruize", TomCruizePageHandler(db))

	go func() {
		var port = os.Getenv("MTVGO_SERVER_PORT")
		if err := http.ListenAndServe(wsAddr+":"+port, nil); err != nil {
			log.Fatal("WebSocket server error:", err)
		}
	}()

	var staticPort = os.Getenv("MTVGO_STATIC_SERVER_PORT")
	if err := http.ListenAndServe(wsAddr+":"+staticPort, nil); err != nil {
		log.Fatal("Static file server error:", err)
	}
}
