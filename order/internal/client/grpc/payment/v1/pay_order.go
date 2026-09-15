package v1

import (
	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/converter"
	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
)

func (c *paymentServiceClient) PayOrder(ctx context.Context, info model.PaymentInfo) (string, error) {
	request := converter.ToProtoPaymentRequest(info)
	paymentResponse, err := c.generatedClient.PayOrder(ctx, request)
	if err != nil {
		return "", model.ErrPayment
	}
	return paymentResponse.TransactionUuid, nil
}
