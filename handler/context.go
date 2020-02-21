package handler

import "mychat/mychannel"

type HandlerContext struct {
	UserChan map[string]*mychannel.MyChannel
	ChanGrp  map[string][]*mychannel.MyChannel
}

func NewHandlerContext() *HandlerContext {
	return &HandlerContext{
		UserChan: make(map[string]*mychannel.MyChannel),
		ChanGrp:  make(map[string][]*mychannel.MyChannel),
	}
}
