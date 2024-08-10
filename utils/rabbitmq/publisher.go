package rabbitmq

import (
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"go-learning/utils/constant"
)

func (r *RabbitMQ) Publish(queue constant.RabbitMqKey, msg string) (err error) {
	var (
		routingKey   = viper.GetString("app.mode")
		exchangeName = viper.GetString("name")
	)

	_ = r.DeclareExchange(exchangeName)
	_ = r.DeclareQueue(queue)
	_ = r.Bind(queue, routingKey, exchangeName)

	// publishing a message
	err = r.Channel.Publish(
		exchangeName, // exchange name
		routingKey,   // key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)

	if err != nil {
		panic(err)
	}

	return
}
