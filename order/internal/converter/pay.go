package converter

import (
	"strings"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
	paymentv1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/proto/payment/v1"
)

func ToProtoPaymentRequest(info model.PaymentInfo) *paymentv1.PayOrderRequest {
	return &paymentv1.PayOrderRequest{
		Info: &paymentv1.PaymentInfo{
			OrderUuid:     info.OrderUuid,
			UserUuid:      info.UserUuid,
			PaymentMethod: toProtoPaymentMethod(info.PaymentMethod),
		},
	}
}

func toProtoPaymentMethod(pm string) paymentv1.PaymentMethod {
	switch strings.ToUpper(pm) {
	case "PAYMENT_METHOD_CARD":
		return paymentv1.PaymentMethod_PAYMENT_METHOD_CARD
	case "PAYMENT_METHOD_SBP":
		return paymentv1.PaymentMethod_PAYMENT_METHOD_SBP
	case "PAYMENT_METHOD_CREDIT_CARD":
		return paymentv1.PaymentMethod_PAYMENT_METHOD_CREDIT_CARD
	case "PAYMENT_METHOD_INVESTOR_MONEY":
		return paymentv1.PaymentMethod_PAYMENT_METHOD_INVESTOR_MONEY
	default:
		return paymentv1.PaymentMethod_PAYMENT_METHOD_UNSPECIFIED
	}
}
