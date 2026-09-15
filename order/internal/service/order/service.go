package order

import (
	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/client/grpc"
	repository "github.com/TikhonovNikita1996/Go-microservises-project/order/internal/repository"
	def "github.com/TikhonovNikita1996/Go-microservises-project/order/internal/service"
)

type OrderService struct {
	orderRepository repository.OrderRepository

	inventoryClient grpc.InventoryClient
	paymentClient   grpc.PaymentClient

	orderPaidProducer def.OrderProducerService
}

func NewService(orderRepository repository.OrderRepository, inventoryClient grpc.InventoryClient,
	paymentClient grpc.PaymentClient, orderProducer def.OrderProducerService,
) *OrderService {
	return &OrderService{
		orderRepository:   orderRepository,
		inventoryClient:   inventoryClient,
		paymentClient:     paymentClient,
		orderPaidProducer: orderProducer,
	}
}
