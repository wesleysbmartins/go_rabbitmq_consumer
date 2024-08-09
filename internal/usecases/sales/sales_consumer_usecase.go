package sales_usecases

import (
	"encoding/json"
	"fmt"
	"go_rabbitmq/internal/entities"

	"github.com/streadway/amqp"
)

type SalesUsecaseConsumer struct{}

type ISalesUsecaseConsumer interface {
	Execute(message amqp.Delivery)
}

func (u *SalesUsecaseConsumer) Execute(message amqp.Delivery) {
	sale := entities.Sale{}
	messageByte := message.Body
	json.Unmarshal(messageByte, &sale)

	saleMsg := fmt.Sprintf("Sale: [Order: %v, Product: %s, Company: %s, Client: %s]", sale.Order, sale.Product, sale.SellingCompany, sale.ClientName)
	rabbitmqDelivery := fmt.Sprintf("Message: [RoutingKey: %s, Consumer: %s, Delivery: %v, Count: %v]", message.RoutingKey, message.ConsumerTag, message.DeliveryTag, message.MessageCount)
	fmt.Println(rabbitmqDelivery, "\n", saleMsg)
}
