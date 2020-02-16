package session

import (
	"mychat/utils"
	"net"
	"sync"
	"time"
)

type Session struct {
	Id    uint64
	Conn  net.Conn
	Times int64
	Attr  map[string]interface{}
	Lock  sync.Mutex
}

func NewSession(conn net.Conn) *Session {
	session := &Session{}
	session.Id = utils.GetRandomId()
	session.Conn = conn
	session.Attr = make(map[string]interface{})
	session.Times = time.Now().Unix()
	return session
}
