package mtvmedia

import (
	"database/sql"
)

type MTVMedia struct {
	DB *sql.DB
}

func NewMTVMedia(db *sql.DB) *MTVMedia {
	return &MTVMedia{DB: db}
}

func (m *MTVMedia) CategoryMovies(category string) ([]map[string]interface{}, error) {
       rows, err := m.DB.Query("SELECT * FROM movies WHERE catagory=? ORDER BY year DESC", category)
       if err != nil {
	       return nil, err
       }
       defer rows.Close()
       return rowsToMaps(rows)
}

func (m *MTVMedia) CategoryTVShows(category, season string) ([]map[string]interface{}, error) {
       query := "SELECT * FROM tvshows WHERE catagory=?"
       args := []interface{}{category}
       if season != "" {
	       query += " AND season=?"
	       args = append(args, season)
       }
       query += " ORDER BY Episode ASC"
       rows, err := m.DB.Query(query, args...)
       if err != nil {
	       return nil, err
       }
       defer rows.Close()
       return rowsToMaps(rows)
}

// rowsToMaps converts SQL rows to a slice of maps
func rowsToMaps(rows *sql.Rows) ([]map[string]interface{}, error) {
       columns, err := rows.Columns()
       if err != nil {
	       return nil, err
       }
       var results []map[string]interface{}
       for rows.Next() {
	       vals := make([]interface{}, len(columns))
	       valPtrs := make([]interface{}, len(columns))
	       for i := range vals {
		       valPtrs[i] = &vals[i]
	       }
	       if err := rows.Scan(valPtrs...); err != nil {
		       return nil, err
	       }
	       rowMap := make(map[string]interface{})
	       for i, col := range columns {
		       b, ok := vals[i].([]byte)
		       if ok {
			       rowMap[col] = string(b)
		       } else {
			       rowMap[col] = vals[i]
		       }
	       }
	       results = append(results, rowMap)
       }
       return results, nil
}
