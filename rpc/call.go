package rpc

import (
	"context"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
	"github.com/uncleyd/core/config"
	"log"
)

var addr string
var basePath string

func Init() {
	addr = config.Get().Rpcx.ConsulAddr
	basePath = config.Get().Rpcx.BasePath
}

// 请求单个服务
func RpcxCall(req, resp interface{}, servicepath, serviceMethod string) error {
	d, _ := client.NewConsulDiscovery(basePath, servicepath, []string{addr}, nil)

	option := client.DefaultOption
	option.SerializeType = protocol.JSON

	xclient := client.NewXClient(servicepath, client.Failtry, client.SelectByUser, d, option)
	defer xclient.Close()

	xclient.SetSelector(&consistentHashSelector{})

	err := xclient.Call(context.Background(), serviceMethod, req, resp)
	if err != nil {
		log.Println("failed to call: %v", err)
		return err
	}
	return nil
}

// 广播请求所有服务
func RpcxBroadcast(req, resp interface{}, servicepath, serviceMethod string) error {
	d, _ := client.NewConsulDiscovery(basePath, servicepath, []string{addr}, nil)

	option := client.DefaultOption
	option.SerializeType = protocol.JSON

	xclient := client.NewXClient(servicepath, client.Failtry, client.RoundRobin, d, option)
	defer xclient.Close()

	err := xclient.Broadcast(context.Background(), serviceMethod, req, resp)
	if err != nil {
		log.Println("failed to call: %v", err)
		return err
	}
	return nil
}
