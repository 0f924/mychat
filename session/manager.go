package session

import "sync"

type SessionManager struct {
	Sessions map[uint64]*Session
	Count    uint32
	Lock     sync.RWMutex
}

func NewSessionM() *SessionManager {
	return &SessionManager{
		Sessions: make(map[uint64]*Session),
		Count:    0,
	}
}
