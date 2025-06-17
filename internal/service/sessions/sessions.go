package sessions

import (
	"errors"
	"sync"
)

var store = struct {
	sync.RWMutex
	sessions map[string]string
}{
	sessions: make(map[string]string),
}

func Create(sessionId, username string) {
	store.Lock()
	defer store.Unlock()
	store.sessions[sessionId] = username
}

func IsValid(sessionId string) bool {
	store.RLock()
	defer store.RUnlock()
	_, exists := store.sessions[sessionId]
	return exists
}

func GetUserName(sessionId string) (string, error) {
	store.RLock()
	defer store.RUnlock()
	u, exists := store.sessions[sessionId]
	if !exists {
		return "", errors.New("Invalid session")
	}
	return u, nil
}

func Delete(sessionId string) {
	store.Lock()
	defer store.Unlock()
	delete(store.sessions, sessionId)
}
