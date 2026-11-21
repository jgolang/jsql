package jsql

import (
	"database/sql"

	"github.com/jgolang/log"
)

// Exec doc ...
func Exec(query string, args ...interface{}) (sql.Result, error) {

	query, err := GetQueryString(query, args...)
	if err != nil {
		return nil, err
	}

	result, err := db.Exec(query)
	if err != nil {
		return nil, err
	}

	response := result

	go func() {
		liID, _ := response.LastInsertId()
		log.Debug(liID)
	}()

	return result, err

}
