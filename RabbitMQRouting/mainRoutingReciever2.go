package main

import "imooc-rabbitmq/RabbitMQ"

func main() {
	RabbitMQTwo :=  RabbitMQ.NewRabbitMQRouting( "routingExchange", "routingKeyTwo")
	RabbitMQTwo.ConsumeRouting()
}