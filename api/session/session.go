package session

import (
	"db/operation"
	"github.com/satori/go.uuid"
	"log"
	"strconv"
	"sync"
	"time"
)

var sessionMap *sync.Map = &sync.Map{}

type Session struct {
	UserName string
	TTL string
}

func LoadSessionFromDB() {
    r, err := operation.GetAllSession()

    if err != nil {
    	log.Println(err)
    	return
	}

    storeFunc := func (k interface{}, v interface{}) bool {
    	sessionMap.Store(k, v)
    	return true
	}

    r.Range( storeFunc )
}

func GenerateNewSessionId(username string) (string) {

    newUUID, _ := uuid.NewV4()

    sessionid := newUUID.String()

    ct := time.Now().UnixNano() / 100000
    ttl := strconv.FormatInt(ct + 30 * 60 * 1000,10)

    newSession := &Session{
    	UserName: username,
    	TTL: ttl,
	}

    sessionMap.Store(sessionid, newSession)

    operation.AddSession(sessionid, ttl, username)

    return sessionid
}

func deleteExpiredSession(sessionId string) {
	sessionMap.Delete(sessionId)
	operation.DeleteSession(sessionId)
}

func IsSessionExpired(sessionId string) bool {

    se, ok := sessionMap.Load(sessionId)

    if ok {
    	ct := time.Now().UnixNano() / 100000

    	ttl, _ := strconv.ParseInt(se.(Session).TTL, 10, 64)

    	if ttl < ct {
    		deleteExpiredSession(sessionId)
    		return true
		}

    	return false
	}

    return true
}