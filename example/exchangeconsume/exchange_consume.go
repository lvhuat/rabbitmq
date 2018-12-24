package main

import (
	"fmt"

	mq "github.com/lvhuat/rabbitmq"
	"github.com/streadway/amqp"
)

func exchangeHandler(delivery *amqp.Delivery) {
	fmt.Println("Delivery coming:", string(delivery.Body))
}

func main() {
	sess, err := mq.Dial("amqp://guest:guest@127.0.0.1:5672/test")
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	if err := sess.HandleExchange(
		"queue-to-bind",
		"exchange-name",
		"topic",
		exchangeHandler,
		mq.MakeupSettings(
			mq.NewExchangeSettings().Durable(),
			mq.NewQueueSettings().Durable().AutoDelete(),
			mq.NewConsumeSettings().AutoAck(),
		), "fruit.*", "vegetables.*"); err != nil {
		panic(err)
	}

	select {
	case <-sess.Closed:
	}
}
