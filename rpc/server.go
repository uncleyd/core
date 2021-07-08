package rpc

import (
	"github.com/smallnest/rpcx/server"
	"sync"
)

type serverList struct {
	List         map[string]interface{}
	sync.RWMutex // 锁，实现并发安全
}

var ServerList *serverList

func InitServerList() {
	ServerList = NewServerList()
}

func NewServerList() *serverList {
	return &serverList{
		List: map[string]interface{}{},
	}
}

func (this *serverList) RegisterName(s *server.Server) {
	this.RLock()
	defer this.RUnlock()

	for k, v := range this.List {
		if k != "" {
			s.RegisterName(k, v, "")
		}
	}
}
