package converter

import (
	serviceModel "github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
	"github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/openapi/order/v1"
)

func ToModelOrderCreateRequest(model *order_v1.CreateOrderRequest) *serviceModel.CreateOrderRequest {
	return &serviceModel.CreateOrderRequest{
		UserUUID:  model.GetUserUUID(),
		PartUuids: model.GetPartUuids(),
	}
}

func ToProtoOrderCreateResponse(model serviceModel.CreateOrderResponse) *order_v1.CreateOrderResponse {
	return &order_v1.CreateOrderResponse{
		OrderUUID:  model.OrderUUID,
		TotalPrice: model.TotalPrice,
	}
}

func ToProtoOrder(model serviceModel.Order) *order_v1.Order {
	return &order_v1.Order{
		OrderUUID:       model.OrderUUID,
		UserUUID:        model.UserUUID,
		PartsUUID:       model.PartsUUID,
		TotalPrice:      model.TotalPrice,
		TransactionUUID: order_v1.NewOptNilString(*model.TransactionUUID),
		PaymentMethod:   order_v1.NewOptNilString(*model.PaymentMethod),
		Status:          ToStringStatus(model.Status),
	}
}

func ToStringStatus(status serviceModel.Status) string {
	switch status {
	case serviceModel.Status_Paid:
		return "Paid"
	case serviceModel.Status_Canceled:
		return "Canceled"
	}
	return "WaitingForPayment"
}
