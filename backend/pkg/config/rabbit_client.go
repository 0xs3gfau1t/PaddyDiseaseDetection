package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func init() {
	log.Println("Testing rabbit client")
	NewRabbit(true)
}

// Param: test: true to close connection, used by startup test
func NewRabbit(test bool) func(string) error {
	conn, err := amqp091.Dial(os.Getenv("RABBIT_CONN_URL"))
	if err != nil {
		log.Println("[x] Couldn't create a rabbit connection")
		log.Println(err)
		return nil
	}
	if test {
		defer conn.Close()
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Println("[x] Couldn't create a rabbit channel")
		log.Println(err)
		conn.Close()
		return nil
	}

	q, err := ch.QueueDeclare(
		os.Getenv("RABBIT_QUEUE"), // name
		false,                     // durable
		false,                     // delete when unused
		false,                     // exclusive
		false,                     // no-wait
		nil,                       // arguments
	)
	if err != nil {
		log.Println("[x] Couldn't declare a rabbit queue")
		log.Println(err)
		conn.Close()
		return nil
	}

	log.Println("[+] Created a new rabbit connection, channel and queue")
	return func(body string) error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return ch.PublishWithContext(ctx,
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp091.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
	}

}
