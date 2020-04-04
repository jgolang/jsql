package jsql

import (
	"database/sql"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// TimeLayout doc ...
var TimeLayout = "2006-01-02 15:04:05-06:00" // for Guatemala UTC

// GetQueryString Generate string sql query. Avoid SQL injection.
func GetQueryString(query string, params ...interface{}) (string, error) {

	for _, param := range params {

		sqlparam := param.(sql.NamedArg)

		if strings.Contains(sqlparam.Name, "@") {
			return "", fmt.Errorf("Error: The parameter name must not contain '@'. Please remamed %v param", sqlparam.Name)
		}

		value := getValue(sqlparam.Value)
		prm := regexp.MustCompile(`(\s|,|=|\()@` + sqlparam.Name + `($|,|\))`)

		for prm.MatchString(query) {
			query = strings.Replace(query, fmt.Sprintf("@%s", sqlparam.Name), fmt.Sprintf("%v", value), -1)
		}

	}

	paramLayout := regexp.MustCompile(`(\s|,|=|\()@(([a-z]|[A-Z]|[0-9])+)($|,|\))`)

	if paramLayout.MatchString(query) {
		return "", fmt.Errorf("Error: There should be no empty values in query. \n%v", query)
	}

	return query, nil
}

func getValue(v interface{}) string {

	if v == nil {
		return "NULL"
	}

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
	case byte:
		return fmt.Sprintf("%v", v.(byte))
	case []byte:
		return fmt.Sprintf("%v", v.([]byte))
	case time.Time:
		return fmt.Sprintf("'%v'", v.(time.Time).Format(TimeLayout))
	//... etc
	default:
		return fmt.Sprintf("%v", v)
	}
}
