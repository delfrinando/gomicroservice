package routes

import (
	"log"

	"github.com/jinzhu/gorm"
	country "github.com/tiket-dev/tiket-microservice-configuration/apps/country/controllers"
	config "github.com/tiket-dev/tiket-microservice-configuration/config"
	amqp "github.com/tiket-dev/tiket-microservice-configuration/core/amqp"
	mysql "github.com/tiket-dev/tiket-microservice-configuration/core/mysql"
)

func Routing() {

	mySqlConnect()
	amqpConnect()
}

func mySqlConnect() *gorm.DB {

	dbConnection := mysql.ConnectMySQL()

	return dbConnection
}

func amqpConnect() {

	configs := config.GetConfig()
	amqpHost := configs.Amqp.URI

	if conn, err := amqp.ConnectToMQ(amqpHost); err != nil {
		log.Println(err)
		log.Println("node will only be able to respond to local connections")
		log.Println("trying to reconnect in 5 seconds...")
	} else {
		log.Println("Connected to AMQP")
		ch := amqp.CreateChannelMQ(conn)
		// language.ServeLanguage(ch)
		country.ServeCountry(ch)
	}
}
