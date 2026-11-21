package jsql

import (
	"database/sql"
	"os"
	"reflect"
	"testing"

	"github.com/go-sql-driver/mysql" // SQL driver
)

func TestExec(t *testing.T) {

	mysqlConnect(t)

	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    sql.Result
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				query: "select now();",
				// args:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Exec(tt.args.query, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}

// mysqlConnect connects to the database named by the JSQL_TEST_DSN
// environment variable (a go-sql-driver/mysql DSN) and skips the test when it
// is not set, so no credentials live in the repository.
func mysqlConnect(t *testing.T) {
	dsn := os.Getenv("JSQL_TEST_DSN")
	if dsn == "" {
		t.Skip("JSQL_TEST_DSN is not set")
	}

	cfg, err := mysql.ParseDSN(dsn)
	if err != nil {
		t.Fatal(err)
	}

	c, err := mysql.NewConnector(cfg)
	if err != nil {
		t.Fatal(err)
	}

	NewJSQL(c)
}
