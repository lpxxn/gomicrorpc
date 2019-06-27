package main

import (
	"context"
	"fmt"
	"time"

	"github.com/lpxxn/gomicrorpc/grpcexample/common"
	model "github.com/lpxxn/gomicrorpc/grpcexample/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client/grpc"
	"github.com/micro/go-micro/registry"
)

func main() {
	// 初始化服务
	service := micro.NewService(
		micro.Client(grpc.NewClient()),
	)
	service.Init()
	service.Options().Registry.Init(func(options *registry.Options) {
		options.Timeout = time.Second * 2
	})

	sayClent := model.NewSayService(common.GrpcExampleName, service.Client())

	rsp, err := sayClent.Hello(context.Background(), &model.SayParam{Msg: "hello server"})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp)
}
