package operation

import (
	"database/sql"
	"db/connection"
	"github.com/satori/go.uuid"
	"log"
)

var Conn = connection.Conn

func AddUser(username string, password string) error {

	newUUID, _ := uuid.NewV4()

	userid := newUUID.String()

	stmt, err := Conn.Prepare(`INSERT INTO users (userid, username, password) VALUES (?,?,?)`)

	if err != nil {
		log.Panicln(err)
		return err
	}

	defer stmt.Close()

	if _, err = stmt.Exec(userid, username, password); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func CheckUser(username string, password string) (bool, error) {

	stmt, err := Conn.Prepare(`SELECT password FROM users WHERE username = ?`)

	if err != nil {
		log.Println(err)
		return false, err
	}

	defer stmt.Close()

	var pw string

	if err = stmt.QueryRow(username).Scan(&pw); err != nil && err != sql.ErrNoRows {
		return false, err
	}

	if pw != password {
		return false, nil
	}

	return true, nil
}

func DeleteUser(username string, password string) error {
	stmt, err := Conn.Prepare("DELETE FROM users WHERE username = ? AND password = ?")

	if err != nil {
		log.Println(err)
		return err
	}

	defer stmt.Close()

	if _, err = stmt.Exec(username, password); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
