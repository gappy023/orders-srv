package main

import (
	"fmt"

	"github.com/gappy023/orders-srv/handler"
	"github.com/gappy023/orders-srv/model"
	"github.com/gappy023/orders-srv/subscriber"

	proto "github.com/gappy023/orders-srv/proto/orders"

	"github.com/gappy023/basic"
	"github.com/gappy023/basic/common"
	"github.com/gappy023/basic/config"
	"github.com/micro/cli/v2"
	micro "github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	// 新建服务
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.orders"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) error {
			// 初始化模型层
			model.Init()
			// 初始化handler
			handler.Init()
			// 初始化sub
			subscriber.Init()

			return nil
		}),
	)

	// 侦听订单支付消息
	err := micro.RegisterSubscriber(common.TopicPaymentDone, service.Server(), subscriber.PayOrder)
	if err != nil {
		log.Fatal(err)
	}

	// 注册服务
	err = proto.RegisterOrdersHandler(service.Server(), new(handler.Orders))
	if err != nil {
		log.Fatal(err)
	}

	// 启动服务
	if err = service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
