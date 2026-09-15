package config

import (
	"os"

	"github.com/joho/godotenv"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/config/env"
)

var appConfig *config

type config struct {
	Logger                LoggerConfig
	OrderHTTPConfig       orderHTTPConfig
	Postgres              PostgresConfig
	InventoryGRPCConfig   inventoryGRPCConfig
	PaymentGRPCConfig     paymentGRPCConfig
	Kafka                 KafkaConfig
	OrderPaidProducer     OrderPaidProducerConfig
	ShipAssembledConsumer ShipAssembledConsumerConfig
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

	orderHTTPCfg, err := env.NewOrderHTTPConfig()
	if err != nil {
		return err
	}

	inventoryGRPCCfg, err := env.NewInventoryGRPCConfig()
	if err != nil {
		return err
	}

	paymentGRPCCfg, err := env.NewPaymentGRPCConfig()
	if err != nil {
		return err
	}

	postgresConfig, err := env.NewPostgresConfig()
	if err != nil {
		return err
	}

	kafkaCfg, err := env.NewKafkaConfig()
	if err != nil {
		return err
	}

	orderPaidProducerCfg, err := env.NewOrderPaidProducerConfig()
	if err != nil {
		return err
	}

	shipAssembledConsumerCfg, err := env.NewShipAssembledConsumerConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		Logger:                loggerCfg,
		OrderHTTPConfig:       orderHTTPCfg,
		Postgres:              postgresConfig,
		InventoryGRPCConfig:   inventoryGRPCCfg,
		PaymentGRPCConfig:     paymentGRPCCfg,
		Kafka:                 kafkaCfg,
		ShipAssembledConsumer: shipAssembledConsumerCfg,
		OrderPaidProducer:     orderPaidProducerCfg,
	}

	return nil
}

func AppConfig() *config {
	return appConfig
}
