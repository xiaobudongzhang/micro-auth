package main

import (
	"fmt"

	"github.com/micro-in-cn/tutorials/microservice-in-micro/part4/basic/common"
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

var (
	appName = "auth_service"
	cfg     = &authCfg{}
)

type authCfg struct {
	common.AppCfg
}

func main() {
	initCfg()
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
	err := config.C().App("etcd", etcdCfg)
	if err != nil {
		panic(err)
	}
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}

func initCfg() {
	source := grpc.NewSource(
		grpc.WithAddress("127.0.0.1:9600"),
		grpc.WithPath("micro"),
	)

	basic.Init(config.WithSource(source))

	err := config.C().App(appName, cfg)
	if err != nil {
		panic(err)
	}

	log.Logf("配置 cfg:%v", cfg)

	return
}
