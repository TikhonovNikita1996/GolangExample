package v1

import (
	"github.com/TikhonovNikita1996/Go-microservises-project/payment/internal/service"
	payment "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/proto/payment/v1"
)

type Api struct {
	payment.UnimplementedPaymentServiceServer

	paymentService service.PaymentService
}

func NewAPI(paymentService service.PaymentService) *Api {
	return &Api{
		paymentService: paymentService,
	}
}
