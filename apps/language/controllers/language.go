package apps

import (
	"log"

	"github.com/streadway/amqp"
	languageService "github.com/tiket-dev/tiket-microservice-configuration/apps/language/services"
	config "github.com/tiket-dev/tiket-microservice-configuration/config"
	coreMQ "github.com/tiket-dev/tiket-microservice-configuration/core/amqp"
)

func ServeLanguage(ch *amqp.Channel) {
	queues := config.GetQueueConfig()
	queueName := queues.Language.QueueName

	q, err := coreMQ.DeclareQueueMQ(ch, queueName)
	msgs := coreMQ.ConsumeQueueMQ(ch, q)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			response := string(d.Body)
			result := languageService.GetLanguage(response)
			coreMQ.PublishQueueMQ(ch, err, d, result)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
