package rabbitmq

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"go-learning/utils/constant"
)

func Initiator() *RabbitMQ {
	var MQConnection = NewRabbitMQ()
	err := MQConnection.Connect()
	if err != nil {
		panic(err)
	}

	return MQConnection
}

type RabbitMQInterface interface {
	Connect() error
	DeclareExchange(exchangeName string) (err error)
	DeclareQueue(name constant.RabbitMqKey) (err error)
	Bind(queueName constant.RabbitMqKey, routingKey, exchangeName string) (err error)
	Publish(key constant.RabbitMqKey, msg string) (err error)
}

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewRabbitMQ() *RabbitMQ {
	return &RabbitMQ{}
}

func (r *RabbitMQ) Connect() (err error) {
	var (
		url = viper.GetString("connection.rabbit.url")
	)

	r.Conn, err = amqp.Dial(url)
	if err != nil {
		panic(err)
	}

	r.Channel, err = r.Conn.Channel()
	if err != nil {
		panic(err)
	}

	fmt.Println("successfully connected to rabbitmq")
	return
}

func (r *RabbitMQ) DeclareExchange(exchangeName string) (err error) {
	err = r.Channel.ExchangeDeclare(
		exchangeName, // nama exchange
		"direct",     // tipe exchange
		true,         // durable
		false,        // auto delete
		false,        // internal
		true,
		nil, // arguments
	)
	if err != nil {
		panic(err)
	}

	return
}

func (r *RabbitMQ) DeclareQueue(name constant.RabbitMqKey) (err error) {
	// declaring queue with its properties over the the channel opened
	_, err = r.Channel.QueueDeclare(
		name.String(), // name
		true,          // durable
		false,         // auto delete
		false,         // exclusive
		false,         // no wait
		nil,           // args
	)
	if err != nil {
		panic(err)
	}

	return
}

func (r *RabbitMQ) Bind(queueName constant.RabbitMqKey, routingKey, exchangeName string) (err error) {
	err = r.Channel.QueueBind(
		queueName.String(), // queue name
		routingKey,         // routing key
		exchangeName,       // exchange name
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	return
}
