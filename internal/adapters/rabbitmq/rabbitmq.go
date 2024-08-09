package rabbitmq

import (
	"fmt"
	configs "go_rabbitmq/config"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type IRabbitMQ interface {
	Connect()
	Channel() *amqp.Channel
}

var channel *amqp.Channel

func (r *RabbitMQ) Connect() {
	credentials := RabbitMQ{}

	config := configs.Config{}
	config.Load("rabbitmq", &credentials)

	connStr := fmt.Sprintf("amqp://%s:%s@%s:%v/", credentials.User, credentials.Password, credentials.Host, credentials.Port)

	conn, err := amqp.Dial(connStr)
	if err != nil {
		panic(fmt.Sprintf("RABBITMQ CONNECTION ERROR\n%s", err.Error()))
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(fmt.Sprintf("RABBITMQ CHANNEL ERROR\n%s", err.Error()))
	}

	channel = ch

	fmt.Println("RABBITMQ CONNECTION SUCCESS!")
}

func (r *RabbitMQ) Channel() *amqp.Channel {
	if channel == nil {
		r.Connect()
	}

	return channel
}
