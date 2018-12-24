package main

import (
	"fmt"

	"time"

	mq "github.com/lvhuat/rabbitmq"
)

func main() {
	sess, err := mq.Dial("amqp://guest:guest@127.0.0.1:5672/test")
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	rpc := mq.NewRpcUtil(sess, time.Minute)
	if err := rpc.SetupReplyQueue(""); err != nil {
		panic(err)
	}

	delivery, err := rpc.Call(
		[]byte("hello"),
		"",
		"test-rpc-queue",
		mq.OptionReplyTo("test-rpc-reply-queue"),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("rpc replied:", string(delivery.Body))
}
