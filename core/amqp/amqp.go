package core

import (
	"log"

	"github.com/streadway/amqp"
	logger "github.com/tiket-dev/tiket-microservice-configuration/helpers/logger"
)

func ConnectToMQ(amqpHost string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(amqpHost)

	logger.Error(err, "Failed to connect to RabbitMQ")

	return conn, nil
}

func CreateChannelMQ(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()

	logger.Error(err, "Failed to open a channel")

	return ch
}

func DeclareQueueMQ(ch *amqp.Channel, queueName string) (*amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	logger.Error(err, "Failed to declare a queue")

	return &q, err
}

func ConsumeQueueMQ(ch *amqp.Channel, q *amqp.Queue) <-chan amqp.Delivery {
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	logger.Error(err, "Failed to register a consumer")

	return msgs
}

func PublishQueueMQ(ch *amqp.Channel, err error, q amqp.Delivery, result []byte) {
	err = ch.Publish(
		"",        // exchange
		q.ReplyTo, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        result,
		})

	log.Printf(" [x] Sent %s", string(result))

	logger.Error(err, "Failed to publish a message")
}
