package args

import "github.com/streadway/amqp"

type Args struct {
	XMessageTTL           int32  `yaml:"x-message-ttl"`
	XExpires              int32  `yaml:"x-expires"`
	XMaxLength            int32  `yaml:"x-max-length"`
	XMaxLengthBytes       int32  `yaml:"x-max-length-bytes"`
	XDeadLetterExchange   string `yaml:"x-dead-letter-exchange"`
	XDeadLetterRoutingKey string `yaml:"x-dead-letter-routing-key"`
	XMaxPriority          int32  `yaml:"x-max-priority"`
	XQueueMode            string `yaml:"x-queue-mode"`
	XQueueMasterLocator   string `yaml:"x-queue-master-locator"`
	AlternativeExchange   string `yaml:"alternative-exchange"`
	XMatch                string `yaml:"x-match"`
}

type IArgs interface {
	Handle() amqp.Table
}

func (a *Args) Handle() amqp.Table {
	table := amqp.Table{}

	if a.XMessageTTL != 0 {
		table["x-message-ttl"] = a.XMessageTTL
	}

	if a.XExpires != 0 {
		table["x-expires"] = a.XExpires
	}

	if a.XMaxLength != 0 {
		table["x-max-length"] = a.XMaxLength
	}

	if a.XMaxLengthBytes != 0 {
		table["x-max-length-bytes"] = a.XMaxLengthBytes
	}

	if a.XDeadLetterExchange != "" {
		table["x-dead-letter-exchange"] = a.XDeadLetterExchange
	}

	if a.XDeadLetterRoutingKey != "" {
		table["x-dead-letter-routing-key"] = a.XDeadLetterRoutingKey
	}

	if a.XMaxPriority != 0 {
		table["x-max-priority"] = a.XMaxPriority
	}

	if a.XQueueMode != "" {
		table["x-queue-mode"] = a.XQueueMode
	}

	if a.XQueueMasterLocator != "" {
		table["x-queue-master-locator"] = a.XQueueMasterLocator
	}

	if a.AlternativeExchange != "" {
		table["alternative-exchange"] = a.AlternativeExchange
	}

	if a.XMatch != "" {
		table["x-match"] = a.XMatch
	}

	return table
}
