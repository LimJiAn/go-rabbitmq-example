package amqp

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

var Conn *amqp.Connection

type Amqp struct {
	Conn *amqp.Connection
}

func AmqpConnect() {
	var err error
	Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	/*	production
		Conn, err = amqp.Dial(fmt.Sprintf(
			"amqp://%s:%s@%s:5672/",
			os.Getenv("GO_RABBITMQ_USER"),
			os.Getenv("GO_RABBITMQ_PASSWORD"),
			os.Getenv("GO_RABBITMQ_HOST"),
		))
	*/
	fmt.Printf("Co1nn: %v\n", Conn)
	if err != nil {
		panic(err)
	}
	defer Conn.Close()

	mq := Amqp{Conn: Conn}
	mq.Publish()

	mq.Consume()
}
