package mysqltools

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"time"
)

// TimeLayout doc ...
var TimeLayout = "2006-01-02 15:04:05-06:00" // for Guatemala UTC

// GetQueryString Generate string sql query. Avoid SQL injection.
func GetQueryString(query string, params ...interface{}) (string, error) {
	valid := regexp.MustCompile(` @(([a-z]|[A-Z]|[0-9])+)( |$|)`)

	for _, param := range params {
		sqlparam := param.(sql.NamedArg)
		value := getValue(sqlparam.Value)
		prm := regexp.MustCompile(` @` + sqlparam.Name + `( |$|)`)
		for prm.MatchString(query) {
			query = prm.ReplaceAllLiteralString(query, fmt.Sprintf(" %s ", value))
		}
	}

	if valid.MatchString(query) {
		return query, errors.New("Existen parametros vacíos")
	}

	return query, nil
}

func getValue(v interface{}) string {
	switch v.(type) {
	case int:
		return fmt.Sprintf("%v", v.(int))
	case int64:
		return fmt.Sprintf("%v", v.(int64))
	case float64:
		return fmt.Sprintf("%v", v.(float64))
	case string:
		return fmt.Sprintf("'%v'", v.(string))
	case bool:
		return fmt.Sprintf("%v", v.(bool))
	case time.Time:
		return fmt.Sprintf("'%v'", v.(time.Time).Format(TimeLayout))
	//... etc
	default:
		return "unknown"
	}
}
