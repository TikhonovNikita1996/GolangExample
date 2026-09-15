package app

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/go-telegram/bot"

	httpClient "github.com/TikhonovNikita1996/Go-microservises-project/notification/internal/client/http"
	telegramClient "github.com/TikhonovNikita1996/Go-microservises-project/notification/internal/client/http/telegram"
	"github.com/TikhonovNikita1996/Go-microservises-project/notification/internal/config"
	kafkaConverter "github.com/TikhonovNikita1996/Go-microservises-project/notification/internal/converter/kafka"
	"github.com/TikhonovNikita1996/Go-microservises-project/notification/internal/converter/kafka/decoder"
	"github.com/TikhonovNikita1996/Go-microservises-project/notification/internal/service"
	orderConsumer "github.com/TikhonovNikita1996/Go-microservises-project/notification/internal/service/consumer/order_paid_consumer"
	shipConsumer "github.com/TikhonovNikita1996/Go-microservises-project/notification/internal/service/consumer/ship_assembled_consumer"
	telegramService "github.com/TikhonovNikita1996/Go-microservises-project/notification/internal/service/telegram"
	"github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/closer"
	wrappedKafka "github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/kafka"
	wrappedKafkaConsumer "github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/kafka/consumer"
	"github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/logger"
	kafkaMidlaware "github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/middleware/kafka"
)

type diContainer struct {
	assemblyConsumerService service.ConsumerService
	orderConsumerService    service.ConsumerService

	orderPaidConsumer     wrappedKafka.Consumer
	shipAssembledConsumer wrappedKafka.Consumer
	consumerGroup         sarama.ConsumerGroup

	orderPaidDecoder         kafkaConverter.OrderPaidDecoder
	shipAssembledPaidDecoder kafkaConverter.ShipAssembledDecoder

	telegramService service.TelegramService
	telegramClient  httpClient.TelegramClient
	telegramBot     *bot.Bot
}

func NewDiContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) OrderPaidConsumerService() service.ConsumerService {
	if d.orderConsumerService == nil {
		d.orderConsumerService = orderConsumer.NewService(d.OrderPaidConsumer(), d.OrderPaidDecoder(), d.TelegramService(context.Background()))
	}

	return d.orderConsumerService
}

func (d *diContainer) ShipAssembledConsumerService() service.ConsumerService {
	if d.assemblyConsumerService == nil {
		d.assemblyConsumerService = shipConsumer.NewService(d.ShipAssembledConsumer(), d.ShipAssembledDecoder(), d.TelegramService(context.Background()))
	}

	return d.assemblyConsumerService
}

func (d *diContainer) ConsumerGroup() sarama.ConsumerGroup {
	if d.consumerGroup == nil {
		consumerGroup, err := sarama.NewConsumerGroup(
			config.AppConfig().Kafka.Brokers(),
			config.AppConfig().OrderPaidConsumer.GroupID(),
			config.AppConfig().OrderPaidConsumer.Config(),
		)
		if err != nil {
			panic(fmt.Sprintf("failed to create consumer group: %s\n", err.Error()))
		}
		closer.AddNamed("Kafka consumer group", func(ctx context.Context) error {
			return d.consumerGroup.Close()
		})

		d.consumerGroup = consumerGroup
	}

	return d.consumerGroup
}

func (d *diContainer) OrderPaidConsumer() wrappedKafka.Consumer {
	if d.orderPaidConsumer == nil {
		d.orderPaidConsumer = wrappedKafkaConsumer.NewConsumer(
			d.ConsumerGroup(),
			[]string{
				config.AppConfig().OrderPaidConsumer.Topic(),
			},
			logger.Logger(),
			kafkaMidlaware.Logging(logger.Logger()),
		)
	}
	return d.orderPaidConsumer
}

func (d *diContainer) ShipAssembledConsumer() wrappedKafka.Consumer {
	if d.orderPaidConsumer == nil {
		d.orderPaidConsumer = wrappedKafkaConsumer.NewConsumer(
			d.ConsumerGroup(),
			[]string{
				config.AppConfig().ShipAssembledConsumer.Topic(),
			},
			logger.Logger(),
			kafkaMidlaware.Logging(logger.Logger()),
		)
	}

	return d.orderPaidConsumer
}

func (d *diContainer) TelegramService(ctx context.Context) service.TelegramService {
	if d.telegramService == nil {
		d.telegramService = telegramService.NewService(
			d.TelegramClient(ctx),
		)
	}

	return d.telegramService
}

func (d *diContainer) TelegramClient(ctx context.Context) httpClient.TelegramClient {
	if d.telegramClient == nil {
		d.telegramClient = telegramClient.NewClient(d.TelegramBot(ctx))
	}

	return d.telegramClient
}

func (d *diContainer) TelegramBot(ctx context.Context) *bot.Bot {
	if d.telegramBot == nil {
		token := config.AppConfig().TelegramConfig.Token()
		b, err := bot.New(token)
		if err != nil {
			panic(fmt.Sprintf("failed to create telegram bot: %s\n", err.Error()))
		}

		d.telegramBot = b
	}

	return d.telegramBot
}

func (d *diContainer) OrderPaidDecoder() kafkaConverter.OrderPaidDecoder {
	if d.orderPaidDecoder == nil {
		d.orderPaidDecoder = decoder.NewOrderPaidDecoder()
	}

	return d.orderPaidDecoder
}

func (d *diContainer) ShipAssembledDecoder() kafkaConverter.ShipAssembledDecoder {
	if d.shipAssembledPaidDecoder == nil {
		d.shipAssembledPaidDecoder = decoder.NewShipAssembledDecoder()
	}

	return d.shipAssembledPaidDecoder
}
