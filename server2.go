package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/hashicorp/consul/api"
    "github.com/micro/go-micro/broker"
    "github.com/micro/go-micro/registry/consul"
    "github.com/micro/go-micro/util/log"
    "github.com/micro/go-micro/web"
    "github.com/micro/go-plugins/broker/rabbitmq"
    "time"
)

func main() {
    reg := consul.NewRegistry(consul.Config(&api.Config{Address: "127.0.0.1:8500"})) // 服务发现和注册
    b := rabbitmq.NewBroker(broker.Addrs("localhost:4306"))
    if err := b.Connect(); err != nil {
        fmt.Println(err)
    }
    b.Subscribe("test1", func(event broker.Event) error {
        fmt.Println(string(event.Message().Body), "12313312")
        return nil
    }, func(options *broker.SubscribeOptions) {
        options.AutoAck = true
    })
    e := gin.Default()
    broker.NewBroker()
    service := web.NewService(
        web.Name("micro-demo"),
        web.RegisterInterval(time.Second*10),
        web.Handler(e),
        web.Registry(reg),
        web.Version("0.0.1"),
        web.Advertise("47.0.0.123"),
    )
    if err := service.Init(); err != nil {
        log.Fatal(err)
    }

    if err := service.Run(); err != nil {
        log.Fatal(err)
    }
}
