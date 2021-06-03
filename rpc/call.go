package rpc

import (
	"context"
	"github.com/smallnest/rpcx/client"
	"github.com/uncleyd/core/config"
	"log"
)

var addr string
var basePath string

func Init() {
	addr = config.Get().Rpcx.ConsulAddr
	basePath = config.Get().Rpcx.BasePath
}

func RpcxCall(req, resp interface{}, servicepath, serviceMethod string) error {
	d, _ := client.NewConsulDiscovery(basePath, servicepath, []string{addr}, nil)
	xclient := client.NewXClient(servicepath, client.Failtry, client.ConsistentHash, d, client.DefaultOption)
	defer xclient.Close()

	err := xclient.Call(context.Background(), serviceMethod, req, resp)
	if err != nil {
		log.Println("failed to call: %v", err)
		return err
	}
	return nil
}
