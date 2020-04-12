package main

import (
	"fmt"

	"github.com/xiaobudongzhang/micro-auth/handler"

	basic "github.com/xiaobudongzhang/micro-basic"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/xiaobudongzhang/micro-auth/model"
	"github.com/xiaobudongzhang/micro-basic/config"

	auth "github.com/xiaobudongzhang/micro-auth/proto/auth"
)

func main() {
	basic.Init()
	micReg := etcd.NewRegistry(registryOptions)
	// New Service
	service := micro.NewService(
		micro.Name("mu.micro.book.service.auth"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(micro.Action(func(c *cli.Context) error {
		// 初始化handler
		model.Init()
		// 初始化handler
		handler.Init()
		return nil
	}))

	// Register Handler
	auth.RegisterServiceHandler(service.Server(), new(handler.Service))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("mu.micro.book.service.auth", service.Server(), new(subscriber.Auth))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
