package main

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/registry"
	pb "go-micro-etcd/service/general_recall_server"
	"time"
)

type MicroRpcClient struct {
	GeneralRecall pb.GeneralRecallService
}

var MicroClient *MicroRpcClient

type MyClientWrapper struct {
	client.Client
}

func (c *MyClientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	return hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
		return c.Client.Call(ctx, req, rsp, opts...)
	}, func(e error) error {
		fmt.Println("这是一个备用的服务")
		return nil
	})
}

// NewClientWrapper returns a hystrix client Wrapper.
func NewMyClientWrapper() client.Wrapper {
	return func(c client.Client) client.Client {
		return &MyClientWrapper{c}
	}
}

func InitRpcClient() {

	reg := etcd.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:2379",
		}
	})
	service := micro.NewService(
		micro.Registry(reg),
		micro.WrapClient(NewMyClientWrapper()),
	)
	// Initialise the client and parse command line flags
	service.Init()

	MicroClient = &MicroRpcClient{
		GeneralRecall: pb.NewGeneralRecallService("example", service.Client()),
	}
}

func GetMicroClient() *MicroRpcClient {
	return MicroClient
}

// 使用方式：将下面代码插入到web，api项目中来调用user.QueryUserByName服务。可以参考web/gin.go。
func main() {
	InitRpcClient() //初始化

	//测试服务端RPC接口
	rsp, err := MicroClient.GeneralRecall.RecallResult(context.TODO(), &pb.RecallResultReq{
		Uid:                "1",
		Itemsize:           2,
		RsStrategyId:       "111",
		RsStrategyClientId: "1",
		RsStrategySceneId:  10,
	})
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
		return
	}
	fmt.Println("返回结果:", rsp.Code, ",return uid:", rsp.Uid)
	for _, v := range rsp.ItemList {
		fmt.Println(v)
	}

	time.Sleep(5 * time.Second)
}
