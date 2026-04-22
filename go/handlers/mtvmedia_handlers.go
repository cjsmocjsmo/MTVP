package handlers

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"database/sql"
	"mtvp/mtvmedia"
)

// HandleSpecialCategoryCommand handles all TV/season and special category commands
func HandleSpecialCategoryCommand(conn *websocket.Conn, db *sql.DB, command string) {
	mtvMedia := mtvmedia.NewMTVMedia(db)
	var results []map[string]interface{}
	var err error
	switch command {
	// Media action commands (map to CategoryMovies)
	case "action", "arnold", "avatar", "brucelee", "brucewillis", "buzz", "cartoons", "charliebrown", "cheechandchong", "chucknorris", "clinteastwood", "comedy", "documentary", "drama", "fantasy", "ghostbusters", "godzilla", "harrisonford", "harrypotter", "hellboy", "indianajones", "jamesbond", "johnwayne", "johnwick", "jurassicpark", "kevincostner", "kingsman", "lego", "meninblack", "minions", "misc", "musicvids", "nature", "nicolascage", "oldies", "panda", "pirates", "predator", "riddick", "science", "scifi", "stalone", "startrek", "starwars", "stooges", "superheros", "superman", "therock", "tinkerbell", "tomcruize", "transformers", "tremors", "trolls", "vandam", "xmen":
		results, err = mtvMedia.CategoryMovies(command)
	// TV/season and special category commands
	case "alteredcarbons1":
		results, err = mtvMedia.CategoryTVShows("AlteredCarbon", "01")
	case "alteredcarbons2":
		results, err = mtvMedia.CategoryTVShows("AlteredCarbon", "02")
	case "columbias1":
		results, err = mtvMedia.CategoryTVShows("Columbia", "01")
	case "cowboybebops1":
		results, err = mtvMedia.CategoryTVShows("Cowboy Bebop", "01")
	case "fallouts1":
		results, err = mtvMedia.CategoryTVShows("Fallout", "01")
	case "fallouts2":
		results, err = mtvMedia.CategoryTVShows("Fallout", "02")
	case "forallmankinds1":
		results, err = mtvMedia.CategoryTVShows("For All Mankind", "01")
	case "forallmankinds2":
		results, err = mtvMedia.CategoryTVShows("For All Mankind", "02")
	case "forallmankinds3":
		results, err = mtvMedia.CategoryTVShows("For All Mankind", "03")
	case "forallmankinds4":
		results, err = mtvMedia.CategoryTVShows("For All Mankind", "04")
	case "foundations1":
		results, err = mtvMedia.CategoryTVShows("Foundation", "01")
	case "foundations2":
		results, err = mtvMedia.CategoryTVShows("Foundation", "02")
	case "foundations3":
		results, err = mtvMedia.CategoryTVShows("Foundation", "03")
	case "fuubars1":
		results, err = mtvMedia.CategoryTVShows("Fubar", "01")
	case "fuubars2":
		results, err = mtvMedia.CategoryTVShows("Fubar", "02")
	case "hford1923s1":
		results, err = mtvMedia.CategoryTVShows("1923", "01")
	case "hford1923s2":
		results, err = mtvMedia.CategoryTVShows("1923", "02")
	case "halos1":
		results, err = mtvMedia.CategoryTVShows("Halo", "01")
	case "halos2":
		results, err = mtvMedia.CategoryTVShows("Halo", "02")
	case "houseofthedragons1":
		results, err = mtvMedia.CategoryTVShows("House of the Dragon", "01")
	case "houseofthedragons2":
		results, err = mtvMedia.CategoryTVShows("House of the Dragon", "02")
	case "lostinspaces1":
		results, err = mtvMedia.CategoryTVShows("Lost in Space", "01")
	case "lostinspaces2":
		results, err = mtvMedia.CategoryTVShows("Lost in Space", "02")
	case "lostinspaces3":
		results, err = mtvMedia.CategoryTVShows("Lost in Space", "03")
	case "mastersoftheuniverse":
		results, err = mtvMedia.CategoryTVShows("Masters of the Universe", "")
	case "monarchlegacyofmonsterss1":
		results, err = mtvMedia.CategoryTVShows("Monarch: Legacy of Monsters", "01")
	case "monarchlegacyofmonsterss2":
		results, err = mtvMedia.CategoryTVShows("Monarch: Legacy of Monsters", "02")
	case "nightskys1":
		results, err = mtvMedia.CategoryTVShows("Night Sky", "01")
	case "orvilles1":
		results, err = mtvMedia.CategoryTVShows("The Orville", "01")
	case "orvilles2":
		results, err = mtvMedia.CategoryTVShows("The Orville", "02")
	case "orvilles3":
		results, err = mtvMedia.CategoryTVShows("The Orville", "03")
	case "percyjacksonandtheolympianss1":
		results, err = mtvMedia.CategoryTVShows("Percy Jackson and the Olympians", "01")
	case "percyjacksonandtheolympianss2":
		results, err = mtvMedia.CategoryTVShows("Percy Jackson and the Olympians", "02")
	case "prehistoricplanets1":
		results, err = mtvMedia.CategoryTVShows("Prehistoric Planet", "01")
	case "prehistoricplanets2":
		results, err = mtvMedia.CategoryTVShows("Prehistoric Planet", "02")
	case "prehistoricplanets3":
		results, err = mtvMedia.CategoryTVShows("Prehistoric Planet", "03")
	case "raisedbywolvess1":
		results, err = mtvMedia.CategoryTVShows("Raised by Wolves", "01")
	case "raisedbywolvess2":
		results, err = mtvMedia.CategoryTVShows("Raised by Wolves", "02")
	case "shoguns1":
		results, err = mtvMedia.CategoryTVShows("Shogun", "01")
	case "silos1":
		results, err = mtvMedia.CategoryTVShows("Silo", "01")
	case "silos2":
		results, err = mtvMedia.CategoryTVShows("Silo", "02")
	case "thecontinentals1":
		results, err = mtvMedia.CategoryTVShows("The Continental", "01")
	case "thelastofus1":
		results, err = mtvMedia.CategoryTVShows("The Last of Us", "01")
	case "thelastofus2":
		results, err = mtvMedia.CategoryTVShows("The Last of Us", "02")
	case "thelordoftheringstheringsofpowers1":
		results, err = mtvMedia.CategoryTVShows("The Lord of the Rings: The Rings of Power", "01")
	case "wheeloftimes1":
		results, err = mtvMedia.CategoryTVShows("The Wheel of Time", "01")
	case "wheeloftimes2":
		results, err = mtvMedia.CategoryTVShows("The Wheel of Time", "02")
	case "wheeloftimes3":
		results, err = mtvMedia.CategoryTVShows("The Wheel of Time", "03")
	// Star Trek series (examples, expand as needed)
	case "discoverys1":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Discovery", "01")
	case "discoverys2":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Discovery", "02")
	case "discoverys3":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Discovery", "03")
	case "discoverys4":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Discovery", "04")
	case "discoverys5":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Discovery", "05")
	case "enterprises1":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Enterprise", "01")
	case "enterprises2":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Enterprise", "02")
	case "enterprises3":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Enterprise", "03")
	case "enterprises4":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Enterprise", "04")
	case "enterprises5":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Enterprise", "05")
	case "lowerdeckss1":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Lower Decks", "01")
	case "lowerdeckss2":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Lower Decks", "02")
	case "lowerdeckss3":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Lower Decks", "03")
	case "lowerdeckss4":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Lower Decks", "04")
	case "lowerdeckss5":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Lower Decks", "05")
	case "picards1":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Picard", "01")
	case "picards2":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Picard", "02")
	case "prodigys1":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Prodigy", "01")
	case "prodigys2":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Prodigy", "02")
	case "sttvs1":
		results, err = mtvMedia.CategoryTVShows("Star Trek: The Animated Series", "01")
	case "sttvs2":
		results, err = mtvMedia.CategoryTVShows("Star Trek: The Animated Series", "02")
	case "sttvs3":
		results, err = mtvMedia.CategoryTVShows("Star Trek: The Animated Series", "03")
	case "strangenewworldss1":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Strange New Worlds", "01")
	case "strangenewworldss2":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Strange New Worlds", "02")
	case "strangenewworldss3":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Strange New Worlds", "03")
	case "tngs1":
		results, err = mtvMedia.CategoryTVShows("Star Trek: The Next Generation", "01")
	case "tngs2":
		results, err = mtvMedia.CategoryTVShows("Star Trek: The Next Generation", "02")
	case "tngs3":
		results, err = mtvMedia.CategoryTVShows("Star Trek: The Next Generation", "03")
	case "tngs4":
		results, err = mtvMedia.CategoryTVShows("Star Trek: The Next Generation", "04")
	case "tngs5":
		results, err = mtvMedia.CategoryTVShows("Star Trek: The Next Generation", "05")
	case "tngs6":
		results, err = mtvMedia.CategoryTVShows("Star Trek: The Next Generation", "06")
	case "tngs7":
		results, err = mtvMedia.CategoryTVShows("Star Trek: The Next Generation", "07")
	case "voyagers1":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Voyager", "01")
	case "voyagers2":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Voyager", "02")
	case "voyagers3":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Voyager", "03")
	case "voyagers4":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Voyager", "04")
	case "voyagers5":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Voyager", "05")
	case "voyagers6":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Voyager", "06")
	case "voyagers7":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Voyager", "07")
	case "deepspacenines1":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Deep Space Nine", "01")
	case "deepspacenines2":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Deep Space Nine", "02")
	case "deepspacenines3":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Deep Space Nine", "03")
	case "deepspacenines4":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Deep Space Nine", "04")
	case "deepspacenines5":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Deep Space Nine", "05")
	case "deepspacenines6":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Deep Space Nine", "06")
	case "deepspacenines7":
		results, err = mtvMedia.CategoryTVShows("Star Trek: Deep Space Nine", "07")
	case "continues1":
		results, err = mtvMedia.CategoryTVShows("Star Trek Continues", "01")
	case "starfleetacademys1":
		results, err = mtvMedia.CategoryTVShows("Starfleet Academy", "01")
	case "jetsonss1":
		results, err = mtvMedia.CategoryTVShows("The Jetsons", "01")
	case "jetsonss2":
		results, err = mtvMedia.CategoryTVShows("The Jetsons", "02")
	case "jonnyquests1":
		results, err = mtvMedia.CategoryTVShows("Jonny Quest", "01")
	default:
		conn.WriteMessage(websocket.TextMessage, []byte(`{"status":"not implemented"}`))
		return
	}
	if err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte(`[]`))
		return
	}
	resp, _ := json.Marshal(results)
	conn.WriteMessage(websocket.TextMessage, resp)
}

// HandleCategoryCommand handles all MTVMEDIA category/movie commands
func HandleCategoryCommand(conn *websocket.Conn, db *sql.DB, command string) {
	mtvMedia := mtvmedia.NewMTVMedia(db)
	results, err := mtvMedia.CategoryMovies(command)
	if err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte(`[]`))
		return
	}
	resp, _ := json.Marshal(results)
	conn.WriteMessage(websocket.TextMessage, resp)
}

// HandleTVShowCommand handles all MTVMEDIA TV/season commands
func HandleTVShowCommand(conn *websocket.Conn, db *sql.DB, category, season string) {
	mtvMedia := mtvmedia.NewMTVMedia(db)
	results, err := mtvMedia.CategoryTVShows(category, season)
	if err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte(`[]`))
		return
	}
	resp, _ := json.Marshal(results)
	conn.WriteMessage(websocket.TextMessage, resp)
}
