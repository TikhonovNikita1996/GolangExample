package config

import "github.com/IBM/sarama"

type LoggerConfig interface {
	Level() string
	AsJson() bool
}

type KafkaConfig interface {
	Brokers() []string
}

type ConsumerConfig interface {
	Topic() string
	GroupID() string
	Config() *sarama.Config
}

type TelegramConfig interface {
	Token() string
}
