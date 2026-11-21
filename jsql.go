package jsql

import (
	"database/sql"
	"database/sql/driver"
)

// JSQL doc ...
type JSQL struct {
	DB *sql.DB
}

// DB connect sql info...
var db *sql.DB

var connectionData string

// NewJSQL doc ...
func NewJSQL(c driver.Connector) *JSQL {

	// connectionData, err := c.Connect()

	db = sql.OpenDB(c)

	jsql := &JSQL{
		DB: db,
	}

	return jsql

}

// SetDBConnection doc ..
func SetDBConnection(inDB *sql.DB) {
	db = inDB
}
