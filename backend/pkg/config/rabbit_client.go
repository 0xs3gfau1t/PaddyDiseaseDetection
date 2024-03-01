package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

var QChannel *amqp091.Channel
var producerQueue amqp091.Queue
var consumerQueue amqp091.Queue

func init() {
	log.Println("Testing rabbit client")

	// Param: test: true to close connection, used by startup test
	conn, err := amqp091.Dial("amqp://" + os.Getenv("RABBIT_USER") + ":" + os.Getenv("RABBIT_PASS") + "@" + os.Getenv("RABBIT_HOST") + ":" + os.Getenv("RABBIT_PORT") + "/" + os.Getenv("RABBIT_VHOST"))
	if err != nil {
		log.Println("[x] Couldn't create a rabbit connection")
		log.Println(err)
		return
	}

	QChannel, err = conn.Channel()

	if err != nil {
		log.Println("[x] Couldn't create a rabbit channel")
		log.Println(err)
		conn.Close()
		return
	}

	producerQueue, err = QChannel.QueueDeclare(
		os.Getenv("RABBIT_QUEUE_PRODUCER"), // name
		true,                               // durable
		false,                              // delete when unused
		false,                              // exclusive
		false,                              // no-wait
		nil,                                // arguments
	)

	consumerQueue, err = QChannel.QueueDeclare(
		os.Getenv("RABBIT_QUEUE_CONSUMER"), // name
		true,                               // durable
		false,                              // delete when unused
		false,                              // exclusive
		false,                              // no-wait
		nil,                                // arguments
	)

	if err != nil {
		log.Println("[x] Couldn't declare a rabbit queue")
		log.Println(err)
		conn.Close()
		return
	}

	log.Println("[+] Created a new rabbit connection, channel and queue")
}

func Publisher(body string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return QChannel.PublishWithContext(ctx,
		"",                 // exchange
		producerQueue.Name, // routing key
		false,              // mandatory
		false,              // immediate
		amqp091.Publishing{
			DeliveryMode: amqp091.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
}
