package setup

import (
	"database/sql"
	"testing"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestCreateTables(t *testing.T) {
       db, err := sql.Open("sqlite3", ":memory:")
       assert.NoError(t, err)
       defer db.Close()

       err = createTables(db)
       assert.NoError(t, err)

       type tableCheck struct {
               name   string
               fields []string
       }
       checks := []tableCheck{
               {"movies", []string{"Name", "Year", "PosterAddr", "Size", "Path", "Idx", "MovId", "Catagory", "HttpThumbPath"}},
               {"tvshows", []string{"TvId", "Size", "Catagory", "Name", "Season", "Episode", "Path", "Idx"}},
               {"images", []string{"ImgId", "Path", "ImgPath", "Size", "Name", "ThumbPath", "Idx", "HttpThumbPath"}},
               {"videos", []string{"VidId", "VidPath", "Size", "Name", "Idx"}},
       }
       var cid int
       var name, ctype string
       var notnull, pk int
       var dfltValue interface{}
       for _, check := range checks {
               foundFields := map[string]bool{}
               rows, err := db.Query("PRAGMA table_info(" + check.name + ")")
               assert.NoError(t, err)
               defer rows.Close()
               for rows.Next() {
                       err := rows.Scan(&cid, &name, &ctype, &notnull, &dfltValue, &pk)
                       assert.NoError(t, err)
                       foundFields[name] = true
               }
               for _, field := range check.fields {
                       assert.True(t, foundFields[field], "%s table missing field: %s", check.name, field)
               }
       }
}
