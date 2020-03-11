package main

import (
	"fmt"
	"imooc-rabbitmq/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("" + "newProduct")

	for i := 0; i <= 10; i++ {
		rabbitmq.PublishPub("订阅模式生成第" + strconv.Itoa(i) + "条数据")
		time.Sleep(time.Second)
		fmt.Println("订阅模式生成第" + strconv.Itoa(i) + "条数据")
	}
}
