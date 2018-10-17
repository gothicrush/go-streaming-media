package operation

import (
	"api/session"
	"database/sql"
	"fmt"
	"log"
	"sync"
)

func AddSession(sessionId string, ttl string, username string) error {

    stmt, err := Conn.Prepare(
    	`INSERT INTO session (username, sessionid, ttl) VALUES(?,?,?);`)

    if err != nil {
    	log.Println(err)
        return err
    }

	defer stmt.Close()

    if _, err := stmt.Exec(username, sessionId, ttl); err != nil {
    	log.Println(err)
    	return err
	}

    return nil
}

func GetOneSession(sessionId string) (*session.Session, error) {

    stmt, err := Conn.Prepare(`SELECT username, ttl FROM session WHERE sessionid = ?`)

    if err != nil {
        fmt.Println(err)
        return nil, err
    }

	defer stmt.Close()

	var ttl string
	var username string

    err = stmt.QueryRow(sessionId).Scan(&ttl,&username)
    if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		return nil, err
    }

	return &session.Session{
	    UserName: username,
	    TTL: ttl,
	}, nil

}

func GetAllSession() (*sync.Map, error) {

    stmt, err := Conn.Prepare("SELECT * FROM session")

    if err != nil {
    	log.Println(err)
    	return nil, err
	}

    rows, err := stmt.Query()
    if err != nil {
    	log.Println(err)
    	return nil, err
	}

	m := &sync.Map{}

    for rows.Next() {
    	var username string
    	var sessionid string
    	var ttl string

    	if err := rows.Scan(&username,&sessionid,&ttl); err != nil {
    		log.Println(err)
    		break
		}

    	se := &session.Session{
    		UserName: username,
    		TTL: ttl,
		}

    	m.Store(sessionid, se)
	}

    return m, nil
}

func DeleteSession(sessionId string) (error) {

	stmt, err := Conn.Prepare(`DELETE FROM session WHERE sessionid=?`)

	if err != nil {
		log.Println(err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(sessionId)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

