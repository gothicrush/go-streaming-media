package operation

import (
	"database/sql"
	"log"
	"db/connection"
)

var Conn = connection.Conn

func AddUser(username string, password string) error {
	
     stmt, err := Conn.Prepare(`INSERT INTO users (username, password) VALUES (?,?)`)

     if err != nil {
     	log.Panicln(err)
     	return err
	 }

     if _, err = stmt.Exec(username, password); err != nil {
     	return err
	 }

     defer stmt.Close()

     return nil
}

func GetUser(username string) (string, error) {

	stmt, err := Conn.Prepare(`SELECT password FROM users WHERE username = ?`)

	if err != nil {
		log.Println(err)
		return "", err
	}

	var password string

	if err = stmt.QueryRow(username).Scan(&password); err != nil && err != sql.ErrNoRows {
		return "", err
	}

	defer stmt.Close()

	return username, nil
}

func DeleteUser(username string, password string) error {
	stmt, err := Conn.Prepare("DELETE FROM users WHERE username = ? AND password = ?")

	if err != nil {
		log.Println(err)
		return err
	}

	if _, err = stmt.Exec(username, password); err != nil {
		return err
	}
	
	defer stmt.Close()

	return nil
}
