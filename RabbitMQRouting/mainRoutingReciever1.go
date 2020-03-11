package main

import "imooc-rabbitmq/RabbitMQ"

func main() {
	rabbitMQOne := RabbitMQ.NewRabbitMQRouting( "routingExchange", "routingKeyOne")
	rabbitMQOne.ConsumeRouting()
}