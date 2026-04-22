package handlers

import (
	"database/sql"
)

// SearchMedia searches movies and tvshows for the phrase in title/name
func SearchMedia(db *sql.DB, phrase string) []map[string]interface{} {
	results := []map[string]interface{}{}
	query := `SELECT * FROM movies WHERE Title LIKE ? UNION SELECT * FROM tvshows WHERE Name LIKE ?`
	likePhrase := "%" + phrase + "%"
	rows, err := db.Query(query, likePhrase, likePhrase)
	if err != nil {
		return results
	}
	defer rows.Close()
	cols, _ := rows.Columns()
	for rows.Next() {
		vals := make([]interface{}, len(cols))
		valPtrs := make([]interface{}, len(cols))
		for i := range vals {
			valPtrs[i] = &vals[i]
		}
		if err := rows.Scan(valPtrs...); err != nil {
			continue
		}
		rowMap := map[string]interface{}{}
		for i, col := range cols {
			if b, ok := vals[i].([]byte); ok {
				rowMap[col] = string(b)
			} else {
				rowMap[col] = vals[i]
			}
		}
		results = append(results, rowMap)
	}
	return results
}
