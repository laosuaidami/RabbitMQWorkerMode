package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

// 简单模式step1，创建简单RabbitMQ
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	return NewRabbitMQ(queueName, "", "")
}

// 简单模式step2，生成者代码
func (r *RabbitMQ) PublishSimple(message string) {
	// 1. 申请队列，如果队列不存在会自动创建，如果存在则跳过创建
	// 保证队列存在，消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false,  // 是否持久化
		false, // 是否自动删除
		false,  // 是否具有排他性
		false,  // 是否阻塞
		nil,  // 额外属性
	)
	if err != nil {
		fmt.Println(err)
	}
	// 2. 发送消息到队列中
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false, // 如果为true，根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返回给发送者
		false, // 如果为true，当Exchange发送消息到队列后发现队列上没有绑定消费者，则会把消息发还给发送者
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(message),
		})
}

// 简单模式step3，消费代码
func (r *RabbitMQ) ConsumeSimple() {
	// 1. 申请队列，如果队列不存在会自动创建，如果存在则跳过创建
	// 保证队列存在，消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false, // 是否持久化
		false, // 是否自动删除
		false,  // 是否具有排他性
		false, // 是否阻塞
		nil, // 额外属性
	)
	if err != nil {
		fmt.Println(err)
	}

	// 2.接收消息
	msgs, err := r.channel.Consume(
		r.QueueName,
		"", // 用来区分多个消费者
		true, // 是否自动应答
		false, // 是否具有排他性
		false, // 如果设置为true, 表示不能将同一个connection中发送的消息传递给这个connection中消费
		false, // 队列消费是否阻塞 false 表示阻塞
		nil)
	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	// 启用协程处理消息
	go func() {
		for d := range msgs {
			// 实现我们要处理的逻辑函数
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	log.Printf("[*] waiting for messages, [退出请按]To exit press CTRL+C")
	<-forever
}