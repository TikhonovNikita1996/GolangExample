package order

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	clients "github.com/TikhonovNikita1996/Go-microservises-project/order/internal/client/grpc/mocks"
	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/repository/mocks"
)

type ServiceSuite struct {
	suite.Suite

	ctx context.Context

	orderRepository *mocks.OrderRepository
	inventoryMock   *clients.InventoryClient
	paymentMock     *clients.PaymentClient

	service *OrderService
}

func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()

	s.orderRepository = mocks.NewOrderRepository(s.T())
	s.inventoryMock = clients.NewInventoryClient(s.T())
	s.paymentMock = clients.NewPaymentClient(s.T())

	s.service = NewService(
		s.orderRepository, s.inventoryMock, s.paymentMock,
	)
}

func (s *ServiceSuite) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
