package main

import (
	"fmt"
	"imooc-rabbitmq/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	rabbitMQOne := RabbitMQ.NewRabbitMQRouting( "routingExchange", "routingKeyOne")
	RabbitMQTwo :=  RabbitMQ.NewRabbitMQRouting( "routingExchange", "routingKeyTwo")
	for i := 0; i <= 10; i++ {
		rabbitMQOne.PublishRouting("RoutingOne模式生成第" + strconv.Itoa(i) + "条数据")
		RabbitMQTwo.PublishRouting("RoutingTwo模式生成第" + strconv.Itoa(i) + "条数据")
		time.Sleep(time.Second)
		fmt.Println("Routing模式生成第" + strconv.Itoa(i) + "条数据")
	}
}
