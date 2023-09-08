/*
Copyright © 2023 JiAn Lim <limjian1990@gmail.com>
*/
package cmd

import (
	"log"

	"github.com/LimJiAn/go-rabbitmq-exam/utils"
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

		q, err := ch.QueueDeclare(
			"hello", // name
			false,   // durable
			false,   // delete when unused
			false,   // exclusive
			false,   // no-wait
			nil,     // arguments
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// consumeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// consumeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
