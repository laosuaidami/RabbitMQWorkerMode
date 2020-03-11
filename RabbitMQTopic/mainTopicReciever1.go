package main

import "imooc-rabbitmq/RabbitMQ"

func main() {
	rabbitMQOne := RabbitMQ.NewRabbitMQTopic( "TopicExchange", "#")
	rabbitMQOne.ConsumeTopic()
}