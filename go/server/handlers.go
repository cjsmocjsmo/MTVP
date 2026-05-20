


package server

import (
    "os/exec"
    "sync"
    "net"
    "encoding/json"
    "time"
    "database/sql"
    "log"
    "fmt"
    "github.com/gorilla/websocket"
    "html/template"
    // "path/filepath"
    "net/http"
)

// HomePageHandler serves the index.html page for the root path
func HomePageHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tmpl, err := template.ParseFiles("templates/index.html")
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

// MovMainPageHandler serves the main movies page (movmainpage.html)
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
        images := getCategoryMovieImages(db, "HomeVids")
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



































func MainTVPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tmpl, err := template.ParseFiles("templates/tv/tvmainpage.html")
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

// TvMainPageHandler serves the main TV page (tvmainpage.html)
func TVActionPageHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tmpl, err := template.ParseFiles("templates/tv/action/tvactionpage.html")
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

// MoblandPageHandler serves all seasons of Mobland with episode info
func TVMoblandPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "MobLand", seasonNum)
            if err != nil {
                log.Println("DB error (Mobland S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/action/tvactionmoblandpage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

// DarkWindsPageHandler serves all seasons of Dark Winds with episode info
func TVDarkWindsPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // We'll support up to 4 seasons, but this is extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 4; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "DarkWinds", seasonNum)
            if err != nil {
                log.Println("DB error (DarkWinds S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/action/tvactiondarkwindspage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        fmt.Println(data) // Debug print to check data structure
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

// ShogunPageHandler serves all seasons of Shogun with episode info
func TVShogunPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Shogun", seasonNum)
            if err != nil {
                log.Println("DB error (Shogun S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/action/tvactionshogunpage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

// TheContinentalPageHandler serves all seasons of The Continental with episode info
func TVTheContinentalPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "TheContinental", seasonNum)
            if err != nil {
                log.Println("DB error (TheContinental S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/action/tvactionthecontinentalpage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVCartoonsPageHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tmpl, err := template.ParseFiles("templates/tv/cartoons/tvcartoonspage.html")
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

func TVMastersOfTheUniversePageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "MastersOfTheUniverse", seasonNum)
            if err != nil {
                log.Println("DB error (MastersOfTheUniverse S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/cartoons/tvcartoonsmastersoftheuniversepage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVFlintstonesPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Flintstones", seasonNum)
            if err != nil {
                log.Println("DB error (Flintstones S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/cartoons/tvcartoonsflintstonespage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVJetsonsPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 3; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Jetsons", seasonNum)
            if err != nil {
                log.Println("DB error (Jetsons S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/cartoons/tvcartoonsjetsonspage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVJonnyQuestPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Jonny Quest", seasonNum)
            if err != nil {
                log.Println("DB error (Jonny Quest S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/cartoons/tvcartoonsjonnyquestpage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVComedyPageHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tmpl, err := template.ParseFiles("templates/tv/comedy/tvcomedypage.html")
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

func TVDMVPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "DMV", seasonNum)
            if err != nil {
                log.Println("DB error (DMV S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/comedy/tvcomedydmvpage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVFubarPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "FuuBar", seasonNum)
            if err != nil {
                log.Println("DB error (FuuBar S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/comedy/tvcomedyfuubarpage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVFantasyPageHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tmpl, err := template.ParseFiles("templates/tv/fantasy/tvfantasypage.html")
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

func TVLordOfTheRingsPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Lord of the Rings", seasonNum)
            if err != nil {
                log.Println("DB error (Lord of the Rings S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/fantasy/tvfantasythelordoftheringstheringsofpowerpage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVHouseOfTheDragonPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "HouseOfTheDragon", seasonNum)
            if err != nil {
                log.Println("DB error (House of the Dragon S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/fantasy/tvfantasyhouseofthedragonpage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVPercyJacksonAndTheOlympiansPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "PercyJacksonAndTheOlympians", seasonNum)
            if err != nil {
                log.Println("DB error (Percy Jackson and the Olympians S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/fantasy/tvfantasypercyjacksonpage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVWednesdayPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Wednesday", seasonNum)
            if err != nil {
                log.Println("DB error (Wednesday S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/fantasy/tvfantasywednesdaypage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVWheelOfTimePageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "WheelOfTime", seasonNum)
            if err != nil {
                log.Println("DB error (Wheel of Time S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/fantasy/tvfantasywheeloftimepage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVMCUPageHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tmpl, err := template.ParseFiles("templates/tv/mcu/tvmcupage.html")
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

func TVFalconAndTheWinterSoldierPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "FalconAndTheWinterSoldier", seasonNum)
            if err != nil {
                log.Println("DB error (Falcon and the Winter Soldier S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/mcu/tvmcufalconandwintersoldierpage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVHawkeyePageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Hawkeye", seasonNum)
            if err != nil {
                log.Println("DB error (Hawkeye S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/mcu/tvmcuhawkeyepage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVIAmGrootPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "IAmGroot", seasonNum)
            if err != nil {
                log.Println("DB error (I Am Groot S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/mcu/tvmcuiamgrootpage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVIronHeartPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "IronHeart", seasonNum)
            if err != nil {
                log.Println("DB error (Iron Heart S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/mcu/tvmcuironheartpage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVLokiPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 3; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Loki", seasonNum)
            if err != nil {
                log.Println("DB error (Loki S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/mcu/tvmculokipage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVMoonKnightPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "MoonKnight", seasonNum)
            if err != nil {
                log.Println("DB error (Moon Knight S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/mcu/tvmcumoonknightpage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVSecretInvasionPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "SecretInvasion", seasonNum)
            if err != nil {
                log.Println("DB error (Secret Invasion S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/mcu/tvmcusecretinvasionpage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVSheHulkPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "SheHulk", seasonNum)
            if err != nil {
                log.Println("DB error (She-Hulk S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/mcu/tvmcushehulkpage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVWandaVisionPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "WandaVision", seasonNum)
            if err != nil {
                log.Println("DB error (WandaVision S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/mcu/tvmcuwandavisionpage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVWonderManPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Wonderman", seasonNum)
            if err != nil {
                log.Println("DB error (Wonderman S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/mcu/tvmcuwondermanpage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVNCISPageHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tmpl, err := template.ParseFiles("templates/tv/ncis/tvncispage.html")
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

func TVNCISHawaiiPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 3; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "NCISHawaii", seasonNum)
            if err != nil {
                log.Println("DB error (NCIS Hawaii S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/ncis/tvncisncishawaiipage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVNCISLAPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 13; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "NCISLA", seasonNum)
            if err != nil {
                log.Println("DB error (NCIS LA S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/ncis/tvncisncislapage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVNCISNewOrleansPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 7; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "NCISNewOrleans", seasonNum)
            if err != nil {
                log.Println("DB error (NCIS New Orleans S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/ncis/tvncisncisneworleanspage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVNCISOriginsPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "NCISOrigins", seasonNum)
            if err != nil {
                log.Println("DB error (NCIS Origins S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/ncis/tvncisncisoriginspage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVNCISNCISPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 23; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "NCISNCIS", seasonNum)
            if err != nil {
                log.Println("DB error (NCIS S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/ncis/tvncisncispage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVNCISSydneyPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 3; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "NCISSydney", seasonNum)
            if err != nil {
                log.Println("DB error (NCIS Sydney S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/ncis/tvncisncissydneypage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVNCISTonyAndZivaPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "TonyAndZiva", seasonNum)
            if err != nil {
                log.Println("DB error (NCIS Ziva and Tony S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/ncis/tvncisncistoniandzivapage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVSciencePageHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tmpl, err := template.ParseFiles("templates/tv/science/tvsciencepage.html")
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

func TVColumbiaPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Columbia", seasonNum)
            if err != nil {
                log.Println("DB error (Columbia S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/science/tvsciencecolumbiapage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVForgedInFirePageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 9; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "ForgedInFire", seasonNum)
            if err != nil {
                log.Println("DB error (Forged in Fire S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/science/tvscienceforgedinfirepage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVPersonOfInterestPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "PersonOfInterest", seasonNum)
            if err != nil {
                log.Println("DB error (Person of Interest S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/science/tvsciencepersonofinterestpage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVPreHistoricPlanetPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "PrehistoricPlanet", seasonNum)
            if err != nil {
                log.Println("DB error (PreHistoric Planet S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/science/tvscienceprehistoricplanetpage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVThePittPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "ThePitt", seasonNum)
            if err != nil {
                log.Println("DB error (The Pitt S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/science/tvsciencethepittpage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVSciFiPageHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tmpl, err := template.ParseFiles("templates/tv/scifi/tvscifipage.html")
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

func TVAlteredCarbonPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "AlteredCarbon", seasonNum)
            if err != nil {
                log.Println("DB error (Altered Carbon S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/scifi/tvscifialteredcarbonpage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVCowboyBebopPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "CowboyBebop", seasonNum)
            if err != nil {
                log.Println("DB error (Cowboy Bebop S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/scifi/tvscificowboybeboppage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVFalloutPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Fallout", seasonNum)
            if err != nil {
                log.Println("DB error (Fallout S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/scifi/tvscififalloutpage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVForAllMankindPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 5; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "ForAllManKind", seasonNum)
            if err != nil {
                log.Println("DB error (For All Mankind S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/scifi/tvscififorallmankindpage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVFoundationPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 3; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Foundation", seasonNum)
            if err != nil {
                log.Println("DB error (Foundation S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/scifi/tvscififoundationpage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVHaloPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Halo", seasonNum)
            if err != nil {
                log.Println("DB error (Halo S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/scifi/tvscifihalopage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVLostInSpacePageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 3; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "LostInSpace", seasonNum)
            if err != nil {
                log.Println("DB error (LostInSpace S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/scifi/tvscifilostinspacepage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVMonarchLegacyOfMonstersPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "MonarchLegacyOfMonsters", seasonNum)
            if err != nil {
                log.Println("DB error (MonarchLegacyOfMonsters S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/scifi/tvscifimonarchlegacyofmonsterspage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVNightSkyPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "NightSky", seasonNum)
            if err != nil {
                log.Println("DB error (NightSky S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/scifi/tvscifinightskypage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVOrvillePageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 3; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Orville", seasonNum)
            if err != nil {
                log.Println("DB error (Orville S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/scifi/tvscifiorvillepage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVRaisedByWolvesPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "RaisedByWolves", seasonNum)
            if err != nil {
                log.Println("DB error (RaisedByWolves S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/scifi/tvscifiraisedbywolvespage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVSiloPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Silo", seasonNum)
            if err != nil {
                log.Println("DB error (Silo S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/scifi/tvscifisilopage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVLastOfUsPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "TheLastOfUs", seasonNum)
            if err != nil {
                log.Println("DB error (TheLastOfUs S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/scifi/tvscifithelastofuspage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVStarTrekPageHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tmpl, err := template.ParseFiles("templates/tv/startrek/tvstartrekpage.html")
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

func TVStarTrekContinuesPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "StarTrekContinues", seasonNum)
            if err != nil {
                log.Println("DB error (StarTrekContinues S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/startrek/tvstartrekcontinuespage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVDeepSpaceNinePageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "DeepSpaceNine", seasonNum)
            if err != nil {
                log.Println("DB error (DeepSpaceNine S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/startrek/tvstartrekdeepspaceninepage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVDiscoveryPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Discovery", seasonNum)
            if err != nil {
                log.Println("DB error (Discovery S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/startrek/tvstartrekdiscoverypage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVEnterprisePageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Enterprise", seasonNum)
            if err != nil {
                log.Println("DB error (Enterprise S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/startrek/tvstartrekenterprisepage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVLowerDecksPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "LowerDecks", seasonNum)
            if err != nil {
                log.Println("DB error (LowerDecks S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/startrek/tvstartreklowerdeckspage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVPicardPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Picard", seasonNum)
            if err != nil {
                log.Println("DB error (Picard S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/startrek/tvstartrekpicardpage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVProdigyPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Prodigy", seasonNum)
            if err != nil {
                log.Println("DB error (Prodigy S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/starfleet/tvstarfleetprodigypage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVStarFleetAcademyPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "StarFleetAcademy", seasonNum)
            if err != nil {
                log.Println("DB error (StarFleetAcademy S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/starfleet/tvstarfleetacademypage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVStrangeNewWorldsPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "StrangeNewWorlds", seasonNum)
            if err != nil {
                log.Println("DB error (StrangeNewWorlds S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/scifi/tvscifistrangenewworldspage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVSTTVPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "STTV", seasonNum)
            if err != nil {
                log.Println("DB error (STTV S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/scifi/tvscifisttvpage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVTNGPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "TNG", seasonNum)
            if err != nil {
                log.Println("DB error (TNG S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/scifi/tvscifitngpage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}

func TVVoyagerPageHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Support up to 4 seasons, extendable
        seasons := map[string][]map[string]interface{}{}
        for i := 1; i <= 2; i++ {
            seasonNum := fmt.Sprintf("%02d", i)
            rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Voyager", seasonNum)
            if err != nil {
                log.Println("DB error (Voyager S", seasonNum, "): ", err)
                continue
            }
            defer rows.Close()
            cols, _ := rows.Columns()
            episodes := []map[string]interface{}{}
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
                    episodes = append(episodes, row)
                }
            }
            if len(episodes) > 0 {
                seasons[seasonNum] = episodes
            }
        }
        tmpl, err := template.ParseFiles("templates/tv/scifi/tvscifivoyagerpage.html")
        if err != nil {
            http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        data := struct {
            Seasons map[string][]map[string]interface{}
        }{Seasons: seasons}
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
        }
    }
}







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
func getCategoryMovieImages(db *sql.DB, category string) []map[string]interface{} {
    query := "SELECT * FROM movies WHERE Catagory=?"
    rows, err := db.Query(query, category)
    if err != nil {
        log.Println("DB error (category images):", err)
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

