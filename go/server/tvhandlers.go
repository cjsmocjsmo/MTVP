package server

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

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
			rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "TheFlintstones", seasonNum)
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
			rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "JonnyQuest", seasonNum)
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
			rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "TheLordOfTheRingsTheRingsOfPower", seasonNum)
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
			rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "FalconWinterSoldier", seasonNum)
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
			rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Supers", seasonNum)
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
		// Support up to 10 seasons, extendable
		seasons := map[string][]map[string]interface{}{}
		for i := 1; i <= 10; i++ {
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

func TVSpiderNoirPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Support up to 4 seasons, extendable
		seasons := map[string][]map[string]interface{}{}
		for i := 1; i <= 2; i++ {
			seasonNum := fmt.Sprintf("%02d", i)
			rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "SpiderNoir", seasonNum)
			if err != nil {
				log.Printf("TVSpiderNoirPageHandler: DB query failed season=%s err=%v", seasonNum, err)
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
				} else {
					log.Printf("TVSpiderNoirPageHandler: season=%s row scan failed err=%v", seasonNum, err)
				}
			}
			if err := rows.Err(); err != nil {
				log.Printf("TVSpiderNoirPageHandler: season=%s rows iteration error=%v", seasonNum, err)
			}
			if len(episodes) > 0 {
				seasons[seasonNum] = episodes
			}
		}
		tmpl, err := template.ParseFiles("templates/tv/cartoons/tvcartoonsspidernoirpage.html")
		if err != nil {
			log.Printf("TVSpiderNoirPageHandler: template parse failed template=%s err=%v", "templates/tv/cartoons/tvcartoonsspidernoirpage.html", err)
			http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			Seasons map[string][]map[string]interface{}
		}{Seasons: seasons}
		err = tmpl.Execute(w, data)
		if err != nil {
			log.Printf("TVSpiderNoirPageHandler: template execute failed template=%s err=%v", "templates/tv/cartoons/tvcartoonsspidernoirpage.html", err)
			http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func TVStarCityPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Support up to 4 seasons, extendable
		seasons := map[string][]map[string]interface{}{}
		for i := 1; i <= 2; i++ {
			seasonNum := fmt.Sprintf("%02d", i)
			rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "StarCity", seasonNum)
			if err != nil {
				log.Println("DB error (Star City S", seasonNum, "): ", err)
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
		tmpl, err := template.ParseFiles("templates/tv/scifi/tvscifistarcitypage.html")
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
			rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Continues", seasonNum)
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
		for i := 1; i <= 7; i++ {
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
		for i := 1; i <= 5; i++ {
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
		for i := 1; i <= 4; i++ {
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
		for i := 1; i <= 5; i++ {
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
		tmpl, err := template.ParseFiles("templates/tv/startrek/tvstartrekprodigypage.html")
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
			rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "StarfleetAcademy", seasonNum)
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
		tmpl, err := template.ParseFiles("templates/tv/startrek/tvstartrekstarfleetacademypage.html")
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
		for i := 1; i <= 3; i++ {
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
		tmpl, err := template.ParseFiles("templates/tv/startrek/tvstartrekstrangenewworldspage.html")
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
		for i := 1; i <= 3; i++ {
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
		tmpl, err := template.ParseFiles("templates/tv/startrek/tvstartreksttvpage.html")
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
		for i := 1; i <= 7; i++ {
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
		tmpl, err := template.ParseFiles("templates/tv/startrek/tvstartrektngpage.html")
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
		for i := 1; i <= 7; i++ {
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
		tmpl, err := template.ParseFiles("templates/tv/startrek/tvstartrekvoyagerpage.html")
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

func TVStarWarsPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/tv/starwars/tvstarwarspage.html")
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

func TVAcolytePageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Support up to 4 seasons, extendable
		seasons := map[string][]map[string]interface{}{}
		for i := 1; i <= 7; i++ {
			seasonNum := fmt.Sprintf("%02d", i)
			rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Acolyte", seasonNum)
			if err != nil {
				log.Println("DB error (Acolyte S", seasonNum, "): ", err)
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
		tmpl, err := template.ParseFiles("templates/tv/starwars/tvstarwarsacolytepage.html")
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

func TVAhsokaPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Support up to 4 seasons, extendable
		seasons := map[string][]map[string]interface{}{}
		for i := 1; i <= 7; i++ {
			seasonNum := fmt.Sprintf("%02d", i)
			rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Ahsoka", seasonNum)
			if err != nil {
				log.Println("DB error (Ahsoka S", seasonNum, "): ", err)
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
		tmpl, err := template.ParseFiles("templates/tv/starwars/tvstarwarsahsokapage.html")
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

func TVAndorPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Support up to 4 seasons, extendable
		seasons := map[string][]map[string]interface{}{}
		for i := 1; i <= 7; i++ {
			seasonNum := fmt.Sprintf("%02d", i)
			rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Andor", seasonNum)
			if err != nil {
				log.Println("DB error (Andor S", seasonNum, "): ", err)
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
		tmpl, err := template.ParseFiles("templates/tv/starwars/tvstarwarsandorpage.html")
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

func TVBookOfBobafettPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Support up to 4 seasons, extendable
		seasons := map[string][]map[string]interface{}{}
		for i := 1; i <= 7; i++ {
			seasonNum := fmt.Sprintf("%02d", i)
			rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Book of Boba Fett", seasonNum)
			if err != nil {
				log.Println("DB error (Book of Boba Fett S", seasonNum, "): ", err)
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
		tmpl, err := template.ParseFiles("templates/tv/starwars/tvstarwarsbookofbobafettpage.html")
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

func TVMandalorianPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Support up to 4 seasons, extendable
		seasons := map[string][]map[string]interface{}{}
		for i := 1; i <= 7; i++ {
			seasonNum := fmt.Sprintf("%02d", i)
			rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Mandalorian", seasonNum)
			if err != nil {
				log.Println("DB error (The Mandalorian S", seasonNum, "): ", err)
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
		tmpl, err := template.ParseFiles("templates/tv/starwars/tvstarwarsmandalorianpage.html")
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

func TVMaulShadowLordPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Support up to 4 seasons, extendable
		seasons := map[string][]map[string]interface{}{}
		for i := 1; i <= 7; i++ {
			seasonNum := fmt.Sprintf("%02d", i)
			rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "MaulShadowLord", seasonNum)
			if err != nil {
				log.Println("DB error (MaulShadowLord S", seasonNum, "): ", err)
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
		tmpl, err := template.ParseFiles("templates/tv/starwars/tvstarwarsmaulshadowlordpage.html")
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

func TVObiWanKenobiPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Support up to 4 seasons, extendable
		seasons := map[string][]map[string]interface{}{}
		for i := 1; i <= 7; i++ {
			seasonNum := fmt.Sprintf("%02d", i)
			rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "ObiWanKenobi", seasonNum)
			if err != nil {
				log.Println("DB error (ObiWanKenobi S", seasonNum, "): ", err)
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
		tmpl, err := template.ParseFiles("templates/tv/starwars/tvstarwarsobiwankenobipage.html")
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

func TVSkeletonCrewPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Support up to 4 seasons, extendable
		seasons := map[string][]map[string]interface{}{}
		for i := 1; i <= 7; i++ {
			seasonNum := fmt.Sprintf("%02d", i)
			rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "SkeletonCrew", seasonNum)
			if err != nil {
				log.Println("DB error (SkeletonCrew S", seasonNum, "): ", err)
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
		tmpl, err := template.ParseFiles("templates/tv/starwars/tvstarwarsskeletoncrewpage.html")
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

func TVTalesOfTheEmpirePageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Support up to 4 seasons, extendable
		seasons := map[string][]map[string]interface{}{}
		for i := 1; i <= 7; i++ {
			seasonNum := fmt.Sprintf("%02d", i)
			rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "TalesOfTheEmpire", seasonNum)
			if err != nil {
				log.Println("DB error (TalesOfTheEmpire S", seasonNum, "): ", err)
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
		tmpl, err := template.ParseFiles("templates/tv/starwars/tvstarwarstalesoftheempirepage.html")
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

func TVTalesOfTheJediPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Support up to 4 seasons, extendable
		seasons := map[string][]map[string]interface{}{}
		for i := 1; i <= 7; i++ {
			seasonNum := fmt.Sprintf("%02d", i)
			rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "TalesOfTheJedi", seasonNum)
			if err != nil {
				log.Println("DB error (TalesOfTheJedi S", seasonNum, "): ", err)
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
		tmpl, err := template.ParseFiles("templates/tv/starwars/tvstarwarstalesofthejedipage.html")
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

func TVTalesOfTheUnderworldPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Support up to 4 seasons, extendable
		seasons := map[string][]map[string]interface{}{}
		for i := 1; i <= 7; i++ {
			seasonNum := fmt.Sprintf("%02d", i)
			rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "TalesOfTheUnderworld", seasonNum)
			if err != nil {
				log.Println("DB error (TalesOfTheUnderWorld S", seasonNum, "): ", err)
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
		tmpl, err := template.ParseFiles("templates/tv/starwars/tvstarwarstalesoftheunderworldpage.html")
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

func TVTheBadBatchPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Support up to 4 seasons, extendable
		seasons := map[string][]map[string]interface{}{}
		for i := 1; i <= 7; i++ {
			seasonNum := fmt.Sprintf("%02d", i)
			rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "TheBadBatch", seasonNum)
			if err != nil {
				log.Println("DB error (TheBadBatch S", seasonNum, "): ", err)
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
		tmpl, err := template.ParseFiles("templates/tv/starwars/tvstarwarsthebadbatchpage.html")
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

func TVVisionsPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Support up to 4 seasons, extendable
		seasons := map[string][]map[string]interface{}{}
		for i := 1; i <= 7; i++ {
			seasonNum := fmt.Sprintf("%02d", i)
			rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Visions", seasonNum)
			if err != nil {
				log.Println("DB error (TVVisions S", seasonNum, "): ", err)
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
		tmpl, err := template.ParseFiles("templates/tv/starwars/tvstarwarsvisionspage.html")
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

func TVWesternsPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/tv/westerns/tvwesternspage.html")
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

func TV1923PageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Support up to 4 seasons, extendable
		seasons := map[string][]map[string]interface{}{}
		for i := 1; i <= 7; i++ {
			seasonNum := fmt.Sprintf("%02d", i)
			rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "HFord1923", seasonNum)
			if err != nil {
				log.Println("DB error (1923 S", seasonNum, "): ", err)
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
		tmpl, err := template.ParseFiles("templates/tv/westerns/tvhford1923page.html")
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

func TVDeadlochPageHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Support up to 2 seasons, extendable
		seasons := map[string][]map[string]interface{}{}
		for i := 1; i <= 2; i++ {
			seasonNum := fmt.Sprintf("%02d", i)
			rows, err := db.Query("SELECT * FROM tvshows WHERE catagory=? AND season=? ORDER BY Episode ASC", "Deadloch", seasonNum)
			if err != nil {
				log.Println("DB error (Deadloch S", seasonNum, "): ", err)
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
		tmpl, err := template.ParseFiles("templates/tv/comedy/tvcomedydeadlochpage.html")
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