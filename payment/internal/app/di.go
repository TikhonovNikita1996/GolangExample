package app

import (
	api "github.com/TikhonovNikita1996/Go-microservises-project/payment/internal/api/payment/v1"
	"github.com/TikhonovNikita1996/Go-microservises-project/payment/internal/service"
	"github.com/TikhonovNikita1996/Go-microservises-project/payment/internal/service/payment"
	paymentv1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/proto/payment/v1"
)

type diContainer struct {
	paymentAPI     paymentv1.PaymentServiceServer
	paymentService service.PaymentService
}

func NewDiContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) UfoAPI() paymentv1.PaymentServiceServer {
	if d.paymentAPI == nil {
		d.paymentAPI = api.NewAPI(d.PaymentService())
	}
	return d.paymentAPI
}

func (d *diContainer) PaymentService() service.PaymentService {
	if d.paymentService == nil {
		d.paymentService = payment.NewService()
	}
	return d.paymentService
}
