/*
Copyright © 2023 JiAn Lim <limjian1990@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"slices"
	"time"

	"github.com/LimJiAn/go-rabbitmq-exam/utils"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/cobra"
)

var ExchangeTypes = []string{"direct", "fanout", "topic", "headers"}

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Send message to queue.",
	Long:  "publish.go, consume.go together make a simple example of using RabbitMQ",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		utils.CheckError(err)
		defer conn.Close()

		ch, err := conn.Channel()
		utils.CheckError(err)
		defer ch.Close()

		exchageName, _ := cmd.Flags().GetString("exchange")
		count, _ := cmd.Flags().GetInt("count")
		if exchageName != "" {
			routingKey, _ := cmd.Flags().GetString("routingkey")
			exchangeType, _ := cmd.Flags().GetString("type")
			if !slices.Contains(ExchangeTypes, exchangeType) {
				log.Fatalf(" 🚫 exchange type must be one of %s", ExchangeTypes)
			}

			log.Printf(" 📮 exchangeName: %s, routingKey: %s, exchangeType: %s, count: %d",
				exchageName, routingKey, exchangeType, count)

			err = ch.ExchangeDeclare(
				exchageName,  // name
				exchangeType, // type
				true,         // durable
				false,        // auto-deleted
				false,        // internal
				false,        // no-wait
				nil,          // arguments
			)
			utils.CheckError(err)

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			for i := 1; i < count+1; i++ {
				body := fmt.Sprintf("Hello World!! [%d]", i)
				err = ch.PublishWithContext(ctx,
					exchageName, // exchange
					routingKey,  // routing key
					false,       // mandatory
					false,       // immediate
					amqp.Publishing{
						ContentType: "text/plain",
						Body:        []byte(body),
					},
				)
				utils.CheckError(err)
				log.Printf(" 📧 Sent %s", body)
			}
			return
		}

		queueName, _ := cmd.Flags().GetString("queue")
		queue, err := ch.QueueDeclare(
			queueName, // name
			false,     // durable
			false,     // delete when unused
			false,     // exclusive
			false,     // no-wait
			nil,       // arguments
		)
		utils.CheckError(err)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

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
			utils.CheckError(err)
			log.Printf(" 📧 Sent %s", body)
		}
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)

	publishCmd.Flags().StringP("exchange", "e", "", "exchange name")
	publishCmd.Flags().StringP("routingkey", "r", "info", "routing key")
	publishCmd.Flags().StringP("type", "t", "direct", "exchange type [direct, fanout, topic, headers]")
	publishCmd.Flags().StringP("queue", "q", "hello", "queue name")
	publishCmd.Flags().IntP("count", "c", 1, "count of message to send")
}
