package config

import (
	"os"

	"github.com/joho/godotenv"

	"github.com/TikhonovNikita1996/Go-microservises-project/notification/internal/config/env"
)

var appConfig *config

type config struct {
	Logger                LoggerConfig
	Kafka                 KafkaConfig
	ShipAssembledConsumer ConsumerConfig
	OrderPaidConsumer     ConsumerConfig
	TelegramConfig        TelegramConfig
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

	shipAssembledConsumerCfg, err := env.NewShipAssembledConsumerConfig()
	if err != nil {
		return err
	}

	orderPaidConsumerCfg, err := env.NewOrderPaidConsumerConfig()
	if err != nil {
		return err
	}

	telegramCfg, err := env.NewTelegramConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		Logger:                loggerCfg,
		Kafka:                 kafkaCfg,
		ShipAssembledConsumer: shipAssembledConsumerCfg,
		OrderPaidConsumer:     orderPaidConsumerCfg,
		TelegramConfig:        telegramCfg,
	}

	return nil
}

func AppConfig() *config {
	return appConfig
}
