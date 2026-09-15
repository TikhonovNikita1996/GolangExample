package v1

import (
	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/service"
)

type OrderAPI struct {
	OrderService service.OrderService
}

func NewAPI(orderService service.OrderService) *OrderAPI {
	return &OrderAPI{
		OrderService: orderService,
	}
}
