package app

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	grpcGo "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	orderv1 "github.com/TikhonovNikita1996/Go-microservises-project/order/internal/api/order/v1"
	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/client/grpc"
	inventoryClient "github.com/TikhonovNikita1996/Go-microservises-project/order/internal/client/grpc/inventory/v1"
	paymentClient "github.com/TikhonovNikita1996/Go-microservises-project/order/internal/client/grpc/payment/v1"
	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/config"
	kafkaConverter "github.com/TikhonovNikita1996/Go-microservises-project/order/internal/converter/kafka"
	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/converter/kafka/decoder"
	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/migrator"
	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/repository"
	orderRepository "github.com/TikhonovNikita1996/Go-microservises-project/order/internal/repository/order"
	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/service"
	orderConsumer "github.com/TikhonovNikita1996/Go-microservises-project/order/internal/service/consumer"
	orderService "github.com/TikhonovNikita1996/Go-microservises-project/order/internal/service/order"
	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/service/producer"
	"github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/closer"
	wrappedKafka "github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/kafka"
	wrappedKafkaConsumer "github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/kafka/consumer"
	wrappedKafkaProducer "github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/kafka/producer"
	"github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/logger"
	kafkaMidlaware "github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/middleware/kafka"
	inventoryV1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/proto/inventory/v1"
	paymentv1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/proto/payment/v1"
)

type diContainer struct {
	orderV1API           *orderv1.OrderAPI
	orderService         service.OrderService
	orderProducerService service.OrderProducerService
	orderConsumerService service.ConsumerService

	orderRepository       repository.OrderRepository
	inventoryClient       grpc.InventoryClient
	paymentClient         grpc.PaymentClient
	postgresDbConn        *pgx.Conn
	consumerGroup         sarama.ConsumerGroup
	shipAssembledConsumer wrappedKafka.Consumer
	shipAssembledDecoder  kafkaConverter.ShipAssembledDecoder
	syncProducer          sarama.SyncProducer
	orderPaidProducer     wrappedKafka.Producer
}

func NewDiContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) OrderV1API(ctx context.Context) *orderv1.OrderAPI {
	if d.orderV1API == nil {
		d.orderV1API = orderv1.NewAPI(d.OrderService(ctx))
	}
	return d.orderV1API
}

func (d *diContainer) OrderService(ctx context.Context) service.OrderService {
	if d.orderService == nil {
		d.orderService = orderService.NewService(d.OrderRepository(ctx), d.InventoryClient(ctx), d.PaymentClient(ctx), d.OrderProducerService())
	}
	return d.orderService
}

func (d *diContainer) OrderProducerService() service.OrderProducerService {
	if d.orderProducerService == nil {
		d.orderProducerService = producer.NewService(d.OrderPaidProducer())
	}

	return d.orderProducerService
}

func (d *diContainer) OrderConsumerService() service.ConsumerService {
	if d.orderConsumerService == nil {
		d.orderConsumerService = orderConsumer.NewService(d.ShipAssembledConsumer(), d.ShipAssembledDecoder(), d.OrderService(context.Background()))
	}

	return d.orderConsumerService
}

func (d *diContainer) OrderRepository(ctx context.Context) repository.OrderRepository {
	if d.orderRepository == nil {
		d.orderRepository = orderRepository.NewRepository(d.PostgresDBConn(ctx))
	}
	return d.orderRepository
}

func (d *diContainer) InventoryClient(ctx context.Context) grpc.InventoryClient {
	if d.inventoryClient == nil {
		inventoryConn, err := createConnection(config.AppConfig().InventoryGRPCConfig.InventoryGRPCAddress())
		if err != nil {
			panic(fmt.Sprintf("failed to connect to inventory service: %s\n", err.Error()))
		}
		generatedInventoryServiceClient := inventoryV1.NewInventoryServiceClient(inventoryConn)

		d.inventoryClient = inventoryClient.NewClient(generatedInventoryServiceClient)

		closer.AddNamed("Inventory connection", func(ctx context.Context) error {
			return inventoryConn.Close()
		})
	}
	return d.inventoryClient
}

func (d *diContainer) PaymentClient(ctx context.Context) grpc.PaymentClient {
	if d.paymentClient == nil {
		paymentConn, err := createConnection(config.AppConfig().PaymentGRPCConfig.PaymentGRPCAddress())
		if err != nil {
			panic(fmt.Sprintf("failed to connect to inventory service: %s\n", err.Error()))
		}
		generatedPaymentServiceClient := paymentv1.NewPaymentServiceClient(paymentConn)

		d.paymentClient = paymentClient.NewClient(generatedPaymentServiceClient)

		closer.AddNamed("Payment connection", func(ctx context.Context) error {
			return paymentConn.Close()
		})
	}
	return d.paymentClient
}

func (d *diContainer) PostgresDBConn(ctx context.Context) *pgx.Conn {
	if d.postgresDbConn == nil {
		con, err := pgx.Connect(ctx, config.AppConfig().Postgres.URI())
		if err != nil {
			panic(fmt.Sprintf("failed to connect to PostgreSQL: %v\n", err))
		}

		err = con.Ping(ctx)
		if err != nil {
			panic(fmt.Sprintf("failed to ping PostgreSQL: %v\n", err))
		}

		migratorRunner := migrator.NewMigrator(stdlib.OpenDB(*con.Config().Copy()), config.AppConfig().Postgres.MigrationsPath())
		err = migratorRunner.Up()
		if err != nil {
			panic(fmt.Sprintf("failed to apply migrations: %v\n", err))
		}

		closer.AddNamed("PosgreSQL connection", func(ctx context.Context) error {
			return con.Close(ctx)
		})
		d.postgresDbConn = con
	}
	return d.postgresDbConn
}

func (d *diContainer) ConsumerGroup() sarama.ConsumerGroup {
	if d.consumerGroup == nil {
		consumerGroup, err := sarama.NewConsumerGroup(
			config.AppConfig().Kafka.Brokers(),
			config.AppConfig().ShipAssembledConsumer.GroupID(),
			config.AppConfig().ShipAssembledConsumer.Config(),
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

func (d *diContainer) ShipAssembledConsumer() wrappedKafka.Consumer {
	if d.shipAssembledConsumer == nil {
		d.shipAssembledConsumer = wrappedKafkaConsumer.NewConsumer(
			d.ConsumerGroup(),
			[]string{
				config.AppConfig().ShipAssembledConsumer.Topic(),
			},
			logger.Logger(),
			kafkaMidlaware.Logging(logger.Logger()),
		)
	}

	return d.shipAssembledConsumer
}

func (d *diContainer) ShipAssembledDecoder() kafkaConverter.ShipAssembledDecoder {
	if d.shipAssembledDecoder == nil {
		d.shipAssembledDecoder = decoder.NewShipAssembledDecoder()
	}

	return d.shipAssembledDecoder
}

func (d *diContainer) SyncProducer() sarama.SyncProducer {
	if d.syncProducer == nil {
		p, err := sarama.NewSyncProducer(
			config.AppConfig().Kafka.Brokers(),
			config.AppConfig().OrderPaidProducer.Config(),
		)
		if err != nil {
			panic(fmt.Sprintf("failed to create sync producer: %s\n", err.Error()))
		}
		closer.AddNamed("Kafka sync producer", func(ctx context.Context) error {
			return p.Close()
		})

		d.syncProducer = p
	}

	return d.syncProducer
}

func (d *diContainer) OrderPaidProducer() wrappedKafka.Producer {
	if d.orderPaidProducer == nil {
		d.orderPaidProducer = wrappedKafkaProducer.NewProducer(
			d.SyncProducer(),
			config.AppConfig().OrderPaidProducer.Topic(),
			logger.Logger(),
		)
	}

	return d.orderPaidProducer
}

func createConnection(address string) (*grpcGo.ClientConn, error) {
	conn, err := grpcGo.NewClient(
		address,
		grpcGo.WithTransportCredentials(insecure.NewCredentials()),
	)
	return conn, err
}
