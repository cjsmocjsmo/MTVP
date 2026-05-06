package server

import (
	"database/sql"
	"encoding/json"
	VideoID   string      `json:"video_id,omitempty"`
	Phrase    string      `json:"phrase,omitempty"`
}
	default:
		// Treat any other command as a movie category
		images := getCategoryMovieImages(db, cmd.Command)
		return WSResponse{"images": images}
	}
}

type WSResponse map[string]interface{}
func getCategoryMovieImages(db *sql.DB, category string) []string {
	// Use category in a LIKE query (case-insensitive)
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

	}
		// Example: select all movies with 'Arnold' in the title (customize as needed)
		rows, err := db.Query("SELECT Path FROM movies WHERE Title LIKE '%Arnold%'")
		if err != nil {
			log.Println("DB error (arnold images):", err)
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
	default:
		return WSResponse{"status": "unknown command"}
	}
}

// getActionMovieImages queries the DB for action movies and returns a list of image URLs
func getActionMovieImages(db *sql.DB) []string {
	// Example: select all movies with 'Action' in the title (customize as needed)
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
			// Assume images are served from /thumbnails/ and path is the filename
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
