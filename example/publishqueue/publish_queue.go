package main

import (
	mq "github.com/lvhuat/rabbitmq"
)

func main() {
	sess, err := mq.Dial("amqp://guest:guest@127.0.0.1:5672/test")
	if err != nil {
		panic(err)
	}
	defer sess.Close()
	if err := sess.PublishString(
		"good day today",
		"",
		"queue-to-publish",
		mq.OptionContentType("application/text"),
		mq.OptionAppId("my-service-id"),
		mq.OptionUserId("keto"),
	); err != nil {
		panic(err)
	}
}
