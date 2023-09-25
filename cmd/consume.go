/*
Copyright © 2023 JiAn Lim <limjian1990@gmail.com>
*/
package cmd

import (
	"log"
	"slices"

	"github.com/LimJiAn/go-rabbitmq-example/utils"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/cobra"
)

// consumeCmd represents the consume command
var consumeCmd = &cobra.Command{
	Use:   "consume",
	Short: "receive message from queue",
	Long:  "publish.go, consume.go together make a simple example of using RabbitMQ",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		utils.CheckError(err)
		defer conn.Close()

		ch, err := conn.Channel()
		utils.CheckError(err)
		defer ch.Close()

		exchangeName, _ := cmd.Flags().GetString("exchange")
		queueName, _ := cmd.Flags().GetString("queue")
		if exchangeName != "" {
			routingKey, _ := cmd.Flags().GetString("routingkey")
			exchangeType, _ := cmd.Flags().GetString("type")
			if !slices.Contains(ExchangeTypes, exchangeType) {
				log.Fatalf(" 🚫 exchange type must be one of %s", ExchangeTypes)
			}

			err = ch.ExchangeDeclare(
				exchangeName, // name
				exchangeType, // type
				true,         // durable
				false,        // auto-deleted
				false,        // internal
				false,        // no-wait
				nil,          // arguments
			)
			utils.CheckError(err)

			q, err := ch.QueueDeclare(
				queueName, // name
				false,     // durable
				false,     // delete when unused
				false,     // exclusive
				false,     // no-wait
				nil,       // arguments
			)
			utils.CheckError(err)

			err = ch.QueueBind(
				q.Name,       // queue name
				routingKey,   // routing key
				exchangeName, // exchange
				false,        // no-wait
				nil,          // arguments
			)
			utils.CheckError(err)

			msgs, err := ch.Consume(
				q.Name, // queue
				"",     // consumer
				true,   // auto-ack
				false,  // exclusive
				false,  // no-local
				false,  // no-wait
				nil,    // args
			)

			forever := make(chan bool)
			count := 0

			go func() {
				for d := range msgs {
					count++
					log.Printf(" 🆗 Received a message: %s / Count: %d", d.Body, count)

				}
			}()

			log.Printf(" 📮 exchangeName: %s, routingKey: %s, exchangeType: %s, queue: %s",
				exchangeName, routingKey, exchangeType, queueName)
			log.Printf(" ✋ Waiting for messages. To exit press CTRL+C")
			<-forever
			return
		}

		q, err := ch.QueueDeclare(
			queueName, // name
			false,     // durable
			false,     // delete when unused
			false,     // exclusive
			false,     // no-wait
			nil,       // arguments
		)
		utils.CheckError(err)

		msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)

		forever := make(chan bool)
		count := 0
		go func() {
			for d := range msgs {
				count++
				log.Printf(" 🆗 Received a message: %s / Count: %d", d.Body, count)
			}
		}()
		log.Printf(" ✋ Waiting for messages. To exit press CTRL+C")
		<-forever
	},
}

func init() {
	rootCmd.AddCommand(consumeCmd)

	consumeCmd.Flags().StringP("exchange", "e", "", "exchange name")
	consumeCmd.Flags().StringP("type", "t", "direct", "exchange type")
	consumeCmd.Flags().StringP("routingkey", "r", "info", "routing key")
	consumeCmd.Flags().StringP("queue", "q", "hello", "queue name")
}
