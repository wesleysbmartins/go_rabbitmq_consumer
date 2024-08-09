package queues

import (
	"fmt"
	"go_rabbitmq/internal/adapters/rabbitmq"
	"go_rabbitmq/internal/adapters/rabbitmq/factory/args"
)

type Queue struct{}

type IQueue interface {
	Create(name string, durable bool, autoDelete bool, exclusive bool, noWait bool, args args.Args)
	Bind(routingKey string, exchangeName string, noWait bool, args args.Args)
	Consume(name string, consumer string, autoAck bool, exclusive bool, noLocal bool, noWait bool, args args.Args, handler IConsumerHandler) error
}

func (q *Queue) Create(name string, durable bool, autoDelete bool, exclusive bool, noWait bool, args args.Args) {
	if name == "" {
		panic("Param name is required!")
	}

	rabbitmq := rabbitmq.RabbitMQ{}
	channel := rabbitmq.Channel()
	_, err := channel.QueueDeclare(name, durable, autoDelete, exclusive, noWait, args.Handle())
	if err != nil {
		panic(fmt.Sprintf("RABBITMQ DECLARE QUEUE ERROR\n%s", err.Error()))
	}
}

func (q *Queue) Bind(queueName string, routingKey string, exchangeName string, noWait bool, args args.Args) {
	rabbitmq := rabbitmq.RabbitMQ{}
	channel := rabbitmq.Channel()
	err := channel.QueueBind(queueName, routingKey, exchangeName, noWait, args.Handle())
	if err != nil {
		panic(fmt.Sprintf("RABBITMQ BINDING QUEUE ERROR\n%s", err.Error()))
	}
}

func (q *Queue) Consume(name string, consumer string, autoAck bool, exclusive bool, noLocal bool, noWait bool, args args.Args, handler IConsumerHandler) error {
	rabbitmq := rabbitmq.RabbitMQ{}
	channel := rabbitmq.Channel()
	messages, err := channel.Consume(name, consumer, autoAck, exclusive, noLocal, noWait, args.Handle())
	if err != nil {
		panic(fmt.Sprintf("RABBITMQ CONSUME IN QUEUE ERROR\n%s", err.Error()))
	}

	listenChannel := make(chan bool)

	go func() {
		fmt.Println("Consumming: ", consumer)

		for message := range messages {
			go handler.Execute(message)
		}
	}()

	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")

	<-listenChannel

	return err
}
