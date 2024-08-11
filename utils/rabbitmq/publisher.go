package rabbitmq

import (
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"go-learning/utils/common"
	"go-learning/utils/constant"
)

func (r *RabbitMQ) Publish(rabbitConfig MqConfig) (err error) {
	var (
		routingKey   = viper.GetString("app.mode")
		exchangeName = viper.GetString("name")
	)

	rabbitConfig.ExchangeName = exchangeName
	rabbitConfig.RoutingKey = routingKey

	if common.IsEmptyField(rabbitConfig.RoutingKey) {
		rabbitConfig.QueueName = constant.DefaultQueue
	}

	_ = r.DeclareExchange(rabbitConfig)
	_ = r.DeclareQueue(rabbitConfig)
	_ = r.Bind(rabbitConfig)

	// publishing a message
	err = r.Channel.Publish(
		rabbitConfig.ExchangeName, // exchange name
		rabbitConfig.RoutingKey,   // key
		false,                     // mandatory
		false,                     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(rabbitConfig.Messsage),
		},
	)

	if err != nil {
		panic(err)
	}

	return
}
