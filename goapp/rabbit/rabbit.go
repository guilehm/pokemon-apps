package rabbit

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func RabbitConnection() (*amqp.Connection, error) {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	FailOnError(err, "Failed to connect to RabbitMQ")
	return conn, err
}
