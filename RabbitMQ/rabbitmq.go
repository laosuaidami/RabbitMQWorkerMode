package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

// url 格式 "amqp://账号:密码@RabbitMQ服务器地址:端口号/vhost"
const MQURL = "amqp://username:password@192.168.10.129:5672/testmq"
//const MQURL = "amqp://imooc_user:password@192.168.10.129:5672/imooc"

type RabbitMQ struct {
	conn * amqp.Connection
	channel *amqp.Channel
	// 队列名称
	QueueName string
	// 交换机
	Exchange string
	// 链接信息
	Mqurl string
	// key
	Key string
}


// 创建RabbitMQ 结构体实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	rabbitMQ := &RabbitMQ{QueueName:queueName,Exchange:exchange,Key:key,Mqurl:MQURL}
	var err error
	// 创建rabbitMQ 连接
	rabbitMQ.conn, err = amqp.Dial(rabbitMQ.Mqurl)
	rabbitMQ.failOnErr(err, "创建链接错误！")
	rabbitMQ.channel, err = rabbitMQ.conn.Channel()
	rabbitMQ.failOnErr(err, "获取channel失败！")
	return rabbitMQ
}

// 断开channel and connection链接
func (r *RabbitMQ) Destory(){
	r.channel.Close()
	r.conn.Close()
}

// 错误处理函数
func (r *RabbitMQ) failOnErr (err error, message string) {
	if err != nil {
		log.Printf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}

}


