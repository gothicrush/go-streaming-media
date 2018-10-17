package connection

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Conn *sql.DB
)

func init() {
	Conn = openConn("root","HIKARI27","stream_media")
}

func openConn(username string, password string, dbname string) *sql.DB {

	url := username + `:` + password + `@#@tcp(localhost:3306)/` + dbname

    Conn, err := sql.Open("mysql", url)

    if err != nil {
    	panic(err.Error())
	}

    err = Conn.Ping()
    if err != nil {
    	panic(err.Error())
	}

    return Conn
}