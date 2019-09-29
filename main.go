package main

import (
    "context"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/hashicorp/consul/api"
    "github.com/micro/go-micro/broker"
    "github.com/micro/go-micro/registry"
    "github.com/micro/go-micro/registry/consul"
    "github.com/micro/go-micro/util/log"
    "github.com/micro/go-micro/web"
    "github.com/micro/go-plugins/broker/rabbitmq"
    "time"
)

func main() {
    register := consul.NewRegistry(consul.Config(&api.Config{Address: "127.0.0.1:8500"}))
    b := rabbitmq.NewBroker(broker.Addrs("amqp://127.0.0.1:5672"))
    if err := b.Connect(); err != nil {
        fmt.Println(err)
    }
    ctx, _ := context.WithCancel(context.Background())
    watcher, err := register.Watch(func(options *registry.WatchOptions) {
        options.Context = ctx
    })
    fmt.Println(err)
    go func() {
        for {
            res, err := watcher.Next()
            fmt.Println(res.Service.Name, res.Action, err, 12313)
        }
    }()
    fmt.Println()
    e := gin.Default()
    service := web.NewService(
        web.Name("micro-demo"),
        web.RegisterInterval(time.Second*10),
        web.Handler(e),
        web.Registry(register),
        web.Version("0.0.1"),
        web.Advertise("47.0.0.124"), // ip 地址及端口号
    )
    go func() {
        for {
            time.Sleep(time.Second * 10)
            fmt.Println(b.Publish("test1", &broker.Message{Body: []byte(`Hello world!`)}), "send")
        }
    }()
    if err := service.Init(); err != nil {
        log.Fatal(err)
    }

    if err := service.Run(); err != nil {
        log.Fatal(err)
    }
}

//func RandomString(l int) string {
//    str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
//    bytes := []byte(str)
//    result := []byte{}
//    r := rand.New(rand.NewSource(time.Now().UnixNano()))
//    for i := 0; i < l; i++ {
//        result = append(result, bytes[r.Intn(len(bytes))])
//    }
//    return string(result)
//}
