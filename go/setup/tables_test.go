package setup

import (
	"database/sql"
	"testing"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestCreateTables(t *testing.T) {
               // Check videos table columns
               foundFields = map[string]bool{}
               rows4, err := db.Query("PRAGMA table_info(videos)")
               assert.NoError(t, err)
               defer rows4.Close()
               for rows4.Next() {
                       err := rows4.Scan(&cid, &name, &ctype, &notnull, &dfltValue, &pk)
                       assert.NoError(t, err)
                       foundFields[name] = true
               }
               expectedFields = []string{"VidId", "VidPath", "Size", "Name", "Idx"}
               for _, field := range expectedFields {
                       assert.True(t, foundFields[field], "videos table missing field: %s", field)
               }
	db, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(t, err)
	defer db.Close()

	err = createTables(db)
	assert.NoError(t, err)

	       // Check movies table columns
       var cid int
       var name, ctype string
       var notnull, pk int
       var dfltValue interface{}
       foundFields := map[string]bool{}
       rows, err := db.Query("PRAGMA table_info(movies)")
       assert.NoError(t, err)
       defer rows.Close()
       for rows.Next() {
               err := rows.Scan(&cid, &name, &ctype, &notnull, &dfltValue, &pk)
               assert.NoError(t, err)
               foundFields[name] = true
       }
       expectedFields := []string{"Name", "Year", "PosterAddr", "Size", "Path", "Idx", "MovId", "Catagory", "HttpThumbPath"}
       for _, field := range expectedFields {
               assert.True(t, foundFields[field], "movies table missing field: %s", field)
       }

       // Check tvshows table columns
       foundFields = map[string]bool{}
       rows2, err := db.Query("PRAGMA table_info(tvshows)")
       assert.NoError(t, err)
       defer rows2.Close()
       for rows2.Next() {
               err := rows2.Scan(&cid, &name, &ctype, &notnull, &dfltValue, &pk)
               assert.NoError(t, err)
               foundFields[name] = true
       }
       expectedFields = []string{"TvId", "Size", "Catagory", "Name", "Season", "Episode", "Path", "Idx"}
       for _, field := range expectedFields {
               assert.True(t, foundFields[field], "tvshows table missing field: %s", field)
       }

       // Check images table columns
       foundFields = map[string]bool{}
       rows3, err := db.Query("PRAGMA table_info(images)")
       assert.NoError(t, err)
       defer rows3.Close()
       for rows3.Next() {
               err := rows3.Scan(&cid, &name, &ctype, &notnull, &dfltValue, &pk)
               assert.NoError(t, err)
               foundFields[name] = true
       }
       expectedFields = []string{"ImgId", "Path", "ImgPath", "Size", "Name", "ThumbPath", "Idx", "HttpThumbPath"}
       for _, field := range expectedFields {
               assert.True(t, foundFields[field], "images table missing field: %s", field)
       }
}
