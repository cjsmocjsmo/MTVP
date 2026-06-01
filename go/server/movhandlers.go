package server

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

func MainMoviePageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/mov/movmainpage.html")
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

// ActionPageHandler serves the action movies page with images from the DB
func ActionPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "Action")
		fmt.Println("Fetched action movie images from DB:", len(images))
		tmpl, err := template.ParseFiles("templates/mov/movactionpage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// ArnoldPageHandler serves the Arnold movies page with images from the DB
func ArnoldPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "Arnold")
		tmpl, err := template.ParseFiles("templates/mov/movarnoldpage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// AvatarPageHandler serves the Avatar movies page with images from the DB
func AvatarPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "Avatar")
		tmpl, err := template.ParseFiles("templates/mov/movavatarpage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// BruceLeePageHandler serves the Bruce Lee movies page with images from the DB
func BruceLeePageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "BruceLee")
		tmpl, err := template.ParseFiles("templates/mov/movbruceleepage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// BruceWillisPageHandler serves the Bruce Willis movies page with images from the DB
func BruceWillisPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "BruceWillis")
		tmpl, err := template.ParseFiles("templates/mov/movbrucewillispage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// BuzzPageHandler serves the Buzz movies page with images from the DB
func BuzzPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "Buzz")
		tmpl, err := template.ParseFiles("templates/mov/movbuzzpage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// CartoonsPageHandler serves the Cartoons movies page with images from the DB
func CartoonsPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "Cartoons")
		tmpl, err := template.ParseFiles("templates/mov/movcartoonspage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// CharlieBrownPageHandler serves the Charlie Brown movies page with images from the DB
func CharlieBrownPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "CharlieBrown")
		tmpl, err := template.ParseFiles("templates/mov/movcharliebrownpage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// CheechAndChongPageHandler serves the Cheech and Chong movies page with images from the DB
func CheechAndChongPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "CheechAndChong")
		tmpl, err := template.ParseFiles("templates/mov/movcheechandchongpage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// ChuckNorrisPageHandler serves the Chuck Norris movies page with images from the DB
func ChuckNorrisPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "ChuckNorris")
		tmpl, err := template.ParseFiles("templates/mov/movchucknorrispage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// ClintEastwoodPageHandler serves the Clint Eastwood movies page with images from the DB
func ClintEastwoodPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "ClintEastwood")
		tmpl, err := template.ParseFiles("templates/mov/movclinteastwoodpage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// ComedyPageHandler serves the Comedy movies page with images from the DB
func ComedyPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "Comedy")
		tmpl, err := template.ParseFiles("templates/mov/movcomedypage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// DocumentaryPageHandler serves the Documentary movies page with images from the DB
func DocumentaryPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "Documentary")
		tmpl, err := template.ParseFiles("templates/mov/movdocumentarypage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// DramaPageHandler serves the Drama movies page with images from the DB
func DramaPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "Drama")
		tmpl, err := template.ParseFiles("templates/mov/movdramapage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// FantasyPageHandler serves the Fantasy movies page with images from the DB
func FantasyPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "Fantasy")
		tmpl, err := template.ParseFiles("templates/mov/movfantasypage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// GhostbustersPageHandler serves the Ghostbusters movies page with images from the DB
func GhostBustersPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "GhostBusters")
		tmpl, err := template.ParseFiles("templates/mov/movghostbusterspage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// GodzillaPageHandler serves the Godzilla movies page with images from the DB
func GodzillaPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "Godzilla")
		tmpl, err := template.ParseFiles("templates/mov/movgodzillapage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// HarrisonFordPageHandler serves the Harrison Ford movies page with images from the DB
func HarrisonFordPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "HarrisonFord")
		tmpl, err := template.ParseFiles("templates/mov/movharrisonfordpage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// HarryPotterPageHandler serves the Harry Potter movies page with images from the DB
func HarryPotterPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "HarryPotter")
		tmpl, err := template.ParseFiles("templates/mov/movharrypotterpage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// HellboyPageHandler serves the Hellboy movies page with images from the DB
func HellboyPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "HellBoy")
		tmpl, err := template.ParseFiles("templates/mov/movhellboypage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// HomeVidsPageHandler serves the Home Vids movies page with images from the DB
func HomeVidsPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryVideoImages(db)
		tmpl, err := template.ParseFiles("templates/mov/movhomevidspage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// getCategoryVideoImages fetches all videos from the videos table
func getCategoryVideoImages(db *sql.DB) []map[string]interface{} {
	query := "SELECT * FROM videos ORDER BY Idx ASC"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("DB error (video images):", err)
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

// IndianaJonesPageHandler serves the Indiana Jones movies page with images from the DB
func IndianaJonesPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "IndianaJones")
		tmpl, err := template.ParseFiles("templates/mov/movindianajonespage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// JamesBondPageHandler serves the James Bond movies page with images from the DB
func JamesBondPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "JamesBond")
		tmpl, err := template.ParseFiles("templates/mov/movjamesbondpage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// JohnWaynePageHandler serves the John Wayne movies page with images from the DB
func JohnWaynePageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "JohnWayne")
		tmpl, err := template.ParseFiles("templates/mov/movjohnwaynepage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// JohnWickPageHandler serves the John Wick movies page with images from the DB
func JohnWickPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "JohnWick")
		tmpl, err := template.ParseFiles("templates/mov/movjohnwickpage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// JurrasicParkPageHandler serves the Jurassic Park movies page with images from the DB
func JurassicParkPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "JurassicPark")
		tmpl, err := template.ParseFiles("templates/mov/movjurassicparkpage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// KevinCostnerPageHandler serves the Kevin Costner movies page with images from the DB
func KevinCostnerPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "KevinCostner")
		tmpl, err := template.ParseFiles("templates/mov/movkevincostnerpage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// KingsmanPageHandler serves the Kingsman movies page with images from the DB
func KingsmanPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "Kingsman")
		tmpl, err := template.ParseFiles("templates/mov/movkingsmanpage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// LegoPageHandler serves the Lego movies page with images from the DB
func LegoPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "Lego")
		tmpl, err := template.ParseFiles("templates/mov/movlegopage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// MenInBlackPageHandler serves the Men in Black movies page with images from the DB
func MenInBlackPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "MenInBlack")
		tmpl, err := template.ParseFiles("templates/mov/movmeninblackpage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// MinionsPageHandler serves the Minions movies page with images from the DB
func MinionsPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "Minions")
		tmpl, err := template.ParseFiles("templates/mov/movminionspage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// MiscPageHandler serves the Misc movies page with images from the DB
func MiscPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "Misc")
		tmpl, err := template.ParseFiles("templates/mov/movmiscpage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// MummyPageHandler serves the Mummy movies page with images from the DB
func MummyPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "Mummy")
		tmpl, err := template.ParseFiles("templates/mov/movmummypage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// MusicVidsPageHandler serves the Music Videos page with images from the DB
func MusicVidsPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "MusicVids")
		tmpl, err := template.ParseFiles("templates/mov/movmusicvidspage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// NaturePageHandler serves the Nature movies page with images from the DB
func NaturePageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "Nature")
		tmpl, err := template.ParseFiles("templates/mov/movnaturepage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// NicolasCagePageHandler serves the Nicolas Cage movies page with images from the DB
func NicolasCagePageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "NicolasCage")
		tmpl, err := template.ParseFiles("templates/mov/movnicolascagepage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// OldiesPageHandler serves the Oldies movies page with images from the DB
func OldiesPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "Oldies")
		tmpl, err := template.ParseFiles("templates/mov/movoldiespage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// PandasPageHandler serves the Pandas movies page with images from the DB
func PandasPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "Pandas")
		tmpl, err := template.ParseFiles("templates/mov/movpandaspage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// PiratesPageHandler serves the Pirates movies page with images from the DB
func PiratesPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "Pirates")
		tmpl, err := template.ParseFiles("templates/mov/movpiratespage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// PredatorPageHandler serves the Predator movies page with images from the DB
func PredatorPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "Predator")
		tmpl, err := template.ParseFiles("templates/mov/movpredatorpage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// RiddickPageHandler serves the Riddick movies page with images from the DB
func RiddickPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "Riddick")
		tmpl, err := template.ParseFiles("templates/mov/movriddickpage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// SciencePageHandler serves the Science movies page with images from the DB
func SciencePageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "Science")
		tmpl, err := template.ParseFiles("templates/mov/movsciencepage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// SciFiPageHandler serves the Sci-Fi movies page with images from the DB
func SciFiPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "SciFi")
		tmpl, err := template.ParseFiles("templates/mov/movscifipage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// StalonePageHandler serves the Stallone movies page with images from the DB
func StalonePageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "Stalone")
		tmpl, err := template.ParseFiles("templates/mov/movstalonepage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// StarTrekPageHandler serves the Star Trek movies page with images from the DB
func StarTrekPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "StarTrek")
		tmpl, err := template.ParseFiles("templates/mov/movstartrekpage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// StarWarsPageHandler serves the Star Wars movies page with images from the DB
func StarWarsPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "StarWars")
		tmpl, err := template.ParseFiles("templates/mov/movstarwarspage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// StoogesPageHandler serves the Stooges movies page with images from the DB
func StoogesPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "Stooges")
		tmpl, err := template.ParseFiles("templates/mov/movstoogespage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// SuperHerosPageHandler serves the SuperHeros movies page with images from the DB
func SuperHerosPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "SuperHeros")
		tmpl, err := template.ParseFiles("templates/mov/movsuperherospage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// SupermanPageHandler serves the Superman movies page with images from the DB
func SuperManPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "SuperMan")
		tmpl, err := template.ParseFiles("templates/mov/movsupermanpage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// TheRockPageHandler serves The Rock movies page with images from the DB
func TheRockPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "TheRock")
		tmpl, err := template.ParseFiles("templates/mov/movtherockpage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// TinkerbellPageHandler serves the Tinkerbell movies page with images from the DB
func TinkerbellPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "TinkerBell")
		tmpl, err := template.ParseFiles("templates/mov/movtinkerbellpage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// TomCruizePageHandler serves the Tom Cruise movies page with images from the DB
func TomCruizePageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "TomCruize")
		tmpl, err := template.ParseFiles("templates/mov/movtomcruizepage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// TransformersPageHandler serves the Transformers movies page with images from the DB
func TransformersPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "Transformers")
		tmpl, err := template.ParseFiles("templates/mov/movtransformerspage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// TremorsPageHandler serves the Tremors movies page with images from the DB
func TremorsPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "Tremors")
		tmpl, err := template.ParseFiles("templates/mov/movtremorspage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// TrollsPageHandler serves the Trolls movies page with images from the DB
func TrollsPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "Trolls")
		tmpl, err := template.ParseFiles("templates/mov/movtrollspage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// VandamPageHandler serves the Van Damme movies page with images from the DB
func VandamPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "VanDam")
		tmpl, err := template.ParseFiles("templates/mov/movvandampage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// XmenPageHandler serves the X-Men movies page with images from the DB
func XmenPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		images := getCategoryMovieImages(db, "XMen")
		tmpl, err := template.ParseFiles("templates/mov/movxmenpage.html")
		if err != nil {
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Images []map[string]interface{}
		}{Images: images}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}
