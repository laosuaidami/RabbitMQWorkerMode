package RabbitMQ

import (
	"github.com/streadway/amqp"
	"log"
)

// 订阅模式step1，创建简单RabbitMQ
func NewRabbitMQPubSub(exchangeName string) *RabbitMQ {
	return NewRabbitMQ("", exchangeName, "")
}

// 订阅模式step2，生成者代码
func (r *RabbitMQ) PublishPub(message string) {
	// 1. 尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,      // exchange Name
		"fanout",   // exchange mode
		true,     // 是否持久化
		false, // 是否自动删除
		false,    // true 表示这个Exchange不可以被client用来推送消息，仅用来Exchange 和 Exchange 之间绑定
		false,
		nil)
	r.failOnErr(err, "Failed to declare an exchange")

	// 2. 发送消息到队列中
	err = r.channel.Publish(
		r.Exchange,
		"",
		// 如果为true，根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返回给发送者
		false,
		// 如果为true，当Exchange发送消息到队列后发现队列上没有绑定消费者，则会把消息发还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(message),
		})
}

// 简单模式step3，消费代码
func (r *RabbitMQ) ConsumeSub() {
	// 1. 尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,      // exchange Name
		"fanout",   // exchange mode
		true,     // 是否持久化
		false, // 是否自动删除
		false,    // true 表示这个Exchange不可以被client用来推送消息，仅用来Exchange 和 Exchange 之间绑定
		false,
		nil)
	r.failOnErr(err, "Failed to declare an exchange")

	// 2. 试探申请队列，这里注意队列名称不要写
	q, err := r.channel.QueueDeclare(
		"",  // 随机生成队列名称
		false,  // 是否持久化
		false,  // 是否自动删除
		true, // 是否具有排他性
		false, // 是否阻塞
		nil, // 额外属性
	)
	r.failOnErr(err, "Failed to declare a queue")

	// 绑定队列到exchange 中
	err = r.channel.QueueBind(
		q.Name,  // 队列名称
		"",  // 在pub/sub 模式下，这里的key要为空
		r.Exchange, // 交换机名称
		false,
		nil,
		)

	// 2.接收消息
	msgs, err := r.channel.Consume(
		q.Name,   // 队列名称
		"", // 用来区分多个消费者
		true, // 是否自动应答
		false, // 是否具有排他性
		false, // 如果设置为true, 表示不能将同一个connection中发送的消息传递给这个connection中消费
		false, // 队列消费是否阻塞 false 表示阻塞
		nil)
	r.failOnErr(err, "Create consume Failed")

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
