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

	DeclareExchange(config MqConfig) (err error)
	DeclareQueue(config MqConfig) (err error)
	Bind(config MqConfig) (err error)

	Publish(config MqConfig) (err error)
	Consume() (err error)

	ConsumeEmailQueue()
}

type MqConfig struct {
	ExchangeName string
	QueueName    constant.MqQueue
	RoutingKey   string
	Messsage     string
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

func (r *RabbitMQ) DeclareExchange(rabbitConfig MqConfig) (err error) {
	err = r.Channel.ExchangeDeclare(
		rabbitConfig.ExchangeName, //  exchange name
		"direct",                  //  exchange type
		true,                      // durable
		false,                     // auto delete
		false,                     // internal
		true,
		nil, // arguments
	)
	if err != nil {
		panic(err)
	}

	return
}

func (r *RabbitMQ) DeclareQueue(rabbitConfig MqConfig) (err error) {
	// declaring queue with its properties over the the channel opened
	_, err = r.Channel.QueueDeclare(
		rabbitConfig.QueueName.String(), // name
		true,                            // durable
		false,                           // auto delete
		false,                           // exclusive
		false,                           // no wait
		nil,                             // args
	)
	if err != nil {
		panic(err)
	}

	return
}

func (r *RabbitMQ) Bind(rabbitConfig MqConfig) (err error) {
	err = r.Channel.QueueBind(
		rabbitConfig.QueueName.String(), // queue name
		rabbitConfig.RoutingKey,         // routing key
		rabbitConfig.ExchangeName,       // exchange name
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	return
}
