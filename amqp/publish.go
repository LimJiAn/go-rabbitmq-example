package amqp

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (mq *Amqp) Publish() {
	ch, err := mq.Conn.Channel()
	fmt.Printf("err: %v\n", err)
	fmt.Printf("ch: %v\n", ch)
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	if err != nil {
		panic(err)
	}

	body := "Hello World!"
	err = ch.PublishWithContext(
		context.Background(),
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

}
