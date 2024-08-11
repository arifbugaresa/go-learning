package rabbitmq

import (
	"fmt"
	"github.com/spf13/viper"
	"go-learning/utils/constant"
)

func (r *RabbitMQ) Consume() (err error) {
	//forever := make(chan bool)
	r.ConsumeEmailQueue()

	//<-forever

	return
}

func (r *RabbitMQ) ConsumeEmailQueue() {
	var (
		queue        = constant.EmailQueue
		exchangeName = viper.GetString("name")
		routingKey   = viper.GetString("app.mode")
	)

	emailQueue := MqConfig{
		QueueName:    queue,
		RoutingKey:   routingKey,
		ExchangeName: exchangeName,
	}

	_ = r.DeclareExchange(emailQueue)
	_ = r.DeclareQueue(emailQueue)
	_ = r.Bind(emailQueue)

	// declaring consumer with its properties over channel opened
	msgEmailQueues, err := r.Channel.Consume(
		emailQueue.QueueName.String(), // queue name
		"",                            // consumer
		true,                          // auto ack
		false,                         // exclusive
		false,                         // no local
		false,                         // no wait
		nil,                           //args
	)
	if err != nil {
		panic(err)
	}

	// do your logic here
	go func() {
		for msg := range msgEmailQueues {
			fmt.Printf("Received Message: %s From ConsumeEmailQueue \n", string(msg.Body))
		}
	}()
}
