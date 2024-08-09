package queues

import "github.com/streadway/amqp"

type IConsumerHandler interface {
	Execute(message amqp.Delivery)
}
