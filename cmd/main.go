package main

import (
	"go_rabbitmq/internal/adapters/rabbitmq"
	sales_usecases "go_rabbitmq/internal/usecases/sales"
)

func init() {
	rabbitmq := rabbitmq.RabbitMQ{}
	rabbitmq.Connect()
}

func main() {
	saleUsecase := sales_usecases.SaleUsecase{}

	saleUsecase.CreateQueue()
	go saleUsecase.Consume()

	select {}
}
