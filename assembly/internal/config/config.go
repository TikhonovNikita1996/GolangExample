package config

import (
	"os"

	"github.com/joho/godotenv"

	"github.com/TikhonovNikita1996/Go-microservises-project/assembly/internal/config/env"
)

var appConfig *config

type config struct {
	Logger            LoggerConfig
	Kafka             KafkaConfig
	OrderPaidProducer OrderPaidProducerConfig
	OrderPaidConsumer OrderPaidConsumerConfig
}

func Load(path ...string) error {
	err := godotenv.Load(path...)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	loggerCfg, err := env.NewLoggerConfig()
	if err != nil {
		return err
	}

	kafkaCfg, err := env.NewKafkaConfig()
	if err != nil {
		return err
	}

	OrderProducerCfg, err := env.NewShipAssembledProducerConfig()
	if err != nil {
		return err
	}

	OrderConsumerCfg, err := env.NewOrderPaidConsumerConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		Logger:            loggerCfg,
		Kafka:             kafkaCfg,
		OrderPaidProducer: OrderProducerCfg,
		OrderPaidConsumer: OrderConsumerCfg,
	}

	return nil
}

func AppConfig() *config {
	return appConfig
}
