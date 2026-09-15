package config

import (
	"time"

	"github.com/IBM/sarama"
)

type LoggerConfig interface {
	Level() string
	AsJson() bool
}

type orderHTTPConfig interface {
	Port() string
	Host() string
	ShutdownTimeout() time.Duration
	ReadTimeout() time.Duration
}

type inventoryGRPCConfig interface {
	InventoryGRPCAddress() string
}

type paymentGRPCConfig interface {
	PaymentGRPCAddress() string
}

type PostgresConfig interface {
	URI() string
	DatabaseName() string
	Port() string
	Host() string
	MigrationsPath() string
}

type KafkaConfig interface {
	Brokers() []string
}

type OrderPaidProducerConfig interface {
	Topic() string
	Config() *sarama.Config
}

type ShipAssembledConsumerConfig interface {
	Topic() string
	GroupID() string
	Config() *sarama.Config
}
