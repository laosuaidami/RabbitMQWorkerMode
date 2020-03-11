package main

import "imooc-rabbitmq/RabbitMQ"

func main() {
	RabbitMQTwo :=  RabbitMQ.NewRabbitMQTopic( "TopicExchange", "topic.*.two")
	RabbitMQTwo.ConsumeTopic()
}