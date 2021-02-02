package main

import (
	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	pb "go-micro-etcd/service/general_recall_server"
	"go-micro-etcd/service/handler"
	"time"
)

func main() {
	serviceName := "example"
	reg := etcd.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"127.0.0.1:2379"} //这是服务发现【etcd集群】节点路径
	})

	// New Service
	service := micro.NewService(
		micro.Name(serviceName),
		micro.Registry(reg),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.Version("latest"),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(3000)), //服务端限流保护
	)

	service.Server().Init(
		server.Wait(nil),
	)

	service.Init()

	// Register Handler
	pb.RegisterGeneralRecallServiceHandler(service.Server(), new(handler.GeneralRecall))

	// Run service
	if err := service.Run(); err != nil {
		panic(err)
	}

}
