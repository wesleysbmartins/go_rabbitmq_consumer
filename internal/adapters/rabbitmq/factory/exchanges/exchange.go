package exchanges

import (
	"fmt"
	"go_rabbitmq/internal/adapters/rabbitmq"
	"go_rabbitmq/internal/adapters/rabbitmq/factory/args"
)

type Exchange struct{}

type IExchange interface {
	Create(name string, kind string, durable bool, autoDelete bool, internal bool, noWait bool, args args.Args)
	Bind(exchangeDestination string, routingKey string, originExchange string, noWait bool, args args.Args)
}

func (e *Exchange) Create(name string, kind string, durable bool, autoDelete bool, internal bool, noWait bool, args args.Args) {
	rabbitmq := rabbitmq.RabbitMQ{}
	channel := rabbitmq.Channel()
	err := channel.ExchangeDeclare(name, kind, durable, autoDelete, internal, noWait, args.Handle())
	if err != nil {
		panic(fmt.Errorf("Exchange config ERROR\n%s", err.Error()))
	}
}

func (e *Exchange) Bind(exchangeDestination string, routingKey string, originExchange string, noWait bool, args args.Args) {
	rabbitmq := rabbitmq.RabbitMQ{}
	channel := rabbitmq.Channel()
	err := channel.ExchangeBind(exchangeDestination, routingKey, originExchange, noWait, args.Handle())
	if err != nil {
		panic(fmt.Errorf("Exchange binding ERROR\n%s", err.Error()))
	}
}
