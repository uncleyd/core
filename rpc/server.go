package rpc

import "github.com/smallnest/rpcx/server"

type serverList struct {
	List map[string]interface{}
}

var ServerList *serverList

func InitServerList() {
	ServerList = NewServerList()
}

func NewServerList() *serverList {
	return &serverList{
		map[string]interface{}{},
	}
}

func (this *serverList) RegisterName(s *server.Server) {
	for k, v := range this.List {
		if k != "" {
			s.RegisterName(k, v, "")
		}
	}
}
