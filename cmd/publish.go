/*
Copyright © 2023 JiAn Lim <limjian1990@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/cobra"
)

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:   "publish [send message count]",
	Short: "Send message to queue. Default count is 1",
	Long:  "publish.go, consume.go together make a simple example of using RabbitMQ",
	Run: func(cmd *cobra.Command, args []string) {
		sendCount := "1"
		if len(args) >= 1 {
			sendCount = args[0]
		}

		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		ch, err := conn.Channel()
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

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		count, err := strconv.Atoi(sendCount)
		if err != nil {
			panic(err)
		}

		for i := 1; i < count+1; i++ {
			body := fmt.Sprintf("Hello World!! [%d]", i)
			err = ch.PublishWithContext(ctx,
				"",         // exchange
				queue.Name, // routing key
				false,      // mandatory
				false,      // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(body),
				},
			)
			if err != nil {
				panic(err)
			}
			log.Printf(" 📧 Sent %s", body)
		}
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// publishCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// publishCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
