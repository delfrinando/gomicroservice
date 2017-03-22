package apps

import (
	"log"

	"github.com/streadway/amqp"
	countryService "github.com/tiket-dev/tiket-microservice-configuration/apps/country/services"
	config "github.com/tiket-dev/tiket-microservice-configuration/config"
	coreMQ "github.com/tiket-dev/tiket-microservice-configuration/core/amqp"
)

func ServeCountry(ch *amqp.Channel) {
	queues := config.GetQueueConfig()
	queueName := queues.Country.QueueName

	q, err := coreMQ.DeclareQueueMQ(ch, queueName)
	msgs := coreMQ.ConsumeQueueMQ(ch, q)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			response := string(d.Body)
			result := countryService.GetCountry(response)
			coreMQ.PublishQueueMQ(ch, err, d, result)
		}
	}()

	log.Printf("Running.....")

	<-forever
}
