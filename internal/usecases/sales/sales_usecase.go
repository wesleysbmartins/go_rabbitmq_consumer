package sales_usecases

import (
	"fmt"
	configs "go_rabbitmq/config"
	"go_rabbitmq/internal/adapters/rabbitmq/factory/exchanges"
	"go_rabbitmq/internal/adapters/rabbitmq/factory/queues"
	"go_rabbitmq/internal/entities"
)

type SaleUsecase struct{}

type ISaleUsecase interface {
	CreateQueue()
	Produce()
}

var c *entities.Consumer

func (u *SaleUsecase) CreateQueue() {
	config := configs.Config{}
	config.Load("sale", &c)

	exchange := exchanges.Exchange{}

	exchange.Create(c.Exchange.Name, c.Exchange.Kind, c.Exchange.Durable, c.Exchange.AutoDelete, c.Exchange.Internal, c.Exchange.NoWait, c.Exchange.Args)

	if c.Exchange.Bind != "" {
		exchange.Bind(c.Exchange.Bind, c.Exchange.RoutingKey, c.Exchange.Name, c.Exchange.NoWait, c.Exchange.Args)
	}

	queue := queues.Queue{}

	queue.Create(c.Queue.Name, c.Queue.Durable, c.Queue.AutoDelete, c.Queue.Exclusive, c.Queue.NoWait, c.Queue.Args)

	queue.Bind(c.Queue.Name, c.Exchange.RoutingKey, c.Exchange.Name, c.Queue.NoWait, c.Queue.Args)

}

func (u *SaleUsecase) Consume() error {

	queue := queues.Queue{}

	err := queue.Consume(c.Queue.Name, c.Consumer.Name, c.Consumer.AutoAck, c.Consumer.Exclusive, c.Consumer.NoLocal, c.Consumer.NoWait, c.Consumer.Args, &SalesUsecaseConsumer{})
	if err != nil {
		panic(fmt.Sprintf("ERROR TO CONSUME QUEUE: %s", err.Error()))
	}

	return err
}
