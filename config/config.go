package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Amqp struct {
	URI string `json:"uri"`
}

type Redis struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Mysql struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Config struct {
	Amqp  Amqp  `json:"amqp"`
	Redis Redis `json:"redis"`
	Mysql Mysql `json:"mysql"`
}

func GetConfig() Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	globalConfig := Config{
		Amqp: Amqp{
			URI: os.Getenv("AMQP_HOST"),
		},
		Redis: Redis{
			Host: os.Getenv("REDIS_HOST"),
			Port: os.Getenv("REDIS_PORT"),
		},
		Mysql: Mysql{
			Host:     os.Getenv("MYSQL_CONFIG_HOST"),
			Port:     os.Getenv("MYSQL_CONFIG_PORT"),
			Database: os.Getenv("MYSQL_CONFIG_DATABASE"),
			Username: os.Getenv("MYSQL_CONFIG_USERNAME"),
			Password: os.Getenv("MYSQL_CONFIG_PASSWORD"),
		},
	}

	return globalConfig
}

type Language struct {
	QueueName string `json:"queueName"`
}

type Country struct {
	QueueName string `json:"queueName"`
}

type AmqpQueueConfig struct {
	Language Language `json:"language"`
	Country  Country  `json:"country"`
}

func GetQueueConfig() AmqpQueueConfig {
	amqpQueueConfig := AmqpQueueConfig{
		Language: Language{
			QueueName: "queue_language",
		},
		Country: Country{
			QueueName: "queue_country",
		},
	}

	return amqpQueueConfig
}
