package entities

import "go_rabbitmq/internal/adapters/rabbitmq/factory/args"

type Consumer struct {
	Exchange struct {
		Name       string    `yaml:"name"`
		Kind       string    `yaml:"kind"`
		RoutingKey string    `yaml:"routing-key"`
		Durable    bool      `yaml:"durable"`
		AutoDelete bool      `yaml:"auto-delete"`
		Internal   bool      `yaml:"internal"`
		NoWait     bool      `yaml:"no-wait"`
		Args       args.Args `yaml:"args"`
		Bind       string    `yaml:"bind"`
	} `yaml:"exchange"`
	Queue struct {
		Name       string    `yaml:"name"`
		Durable    bool      `yaml:"durable"`
		Exclusive  bool      `yaml:"exclusive"`
		AutoDelete bool      `yaml:"auto-delete"`
		NoWait     bool      `yaml:"no-wait"`
		Args       args.Args `yaml:"args"`
		Bind       string    `yaml:"bind"`
	} `yaml:"queue"`
	Consumer struct {
		Name      string    `yaml:"name"`
		AutoAck   bool      `yaml:"auto-ack"`
		Exclusive bool      `yaml:"exclusive"`
		NoLocal   bool      `yaml:"no-local"`
		NoWait    bool      `yaml:"no-wait"`
		Args      args.Args `yaml:"args"`
	} `yaml:"consumer"`
}
