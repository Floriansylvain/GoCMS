package internal

import (
	"crypto"
	"fmt"
	"hash"
	"time"
)

type Session struct {
	Token               hash.Hash `json:"token"`
	ExpirationTimestamp int64     `json:"expirationTimestamp"`
}

var SESSIONS []Session

const SESSION_DURATION int64 = 3600

func generateSessionToken(user User) hash.Hash {
	token := crypto.SHA256.New()
	stringToHash := fmt.Sprintf(user.Email + user.Password)
	token.Write([]byte(stringToHash))
	return token
}

func isSessionExpired(session Session) bool {
	return session.ExpirationTimestamp < time.Now().Unix()
}

func addSession(user User) {
	SESSIONS = append(SESSIONS, Session{
		Token:               generateSessionToken(user),
		ExpirationTimestamp: time.Now().Unix() + SESSION_DURATION,
	})
}

func removeSession(user User) {
	index := 0
	sessLen := len(SESSIONS)
	for i := 0; i < sessLen; i++ {
		if SESSIONS[i].Token == generateSessionToken(user) {
			index = i
			break
		}
	}
	fmt.Println(SESSIONS)
	SESSIONS[index] = SESSIONS[sessLen-1]
	SESSIONS[sessLen-1] = Session{}
	SESSIONS = SESSIONS[:sessLen-1]
	fmt.Println(SESSIONS)
}
