package converter

import (
	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
	repoModel "github.com/TikhonovNikita1996/Go-microservises-project/order/internal/repository/model"
	orderv1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/openapi/order/v1"
)

func ToRepoOrder(order *model.Order) repoModel.Order {
	return repoModel.Order{
		OrderUUID:       order.OrderUUID,
		UserUUID:        order.UserUUID,
		PartsUUID:       order.PartsUUID,
		TotalPrice:      order.TotalPrice,
		TransactionUUID: order.TransactionUUID,
		PaymentMethod:   order.PaymentMethod,
		Status:          repoModel.Status(order.Status),
	}
}

func ToProtoOrder(order *model.Order) *orderv1.Order {
	var transactionUUID, PaymentMethod orderv1.OptNilString

	if order.TransactionUUID != nil {
		transactionUUID = orderv1.NewOptNilString(*order.TransactionUUID)
	} else {
		transactionUUID = orderv1.OptNilString{}
	}

	if order.PaymentMethod != nil {
		PaymentMethod = orderv1.NewOptNilString(*order.PaymentMethod)
	} else {
		PaymentMethod = orderv1.OptNilString{}
	}
	return &orderv1.Order{
		OrderUUID:       order.OrderUUID,
		UserUUID:        order.UserUUID,
		PartsUUID:       order.PartsUUID,
		TotalPrice:      order.TotalPrice,
		TransactionUUID: transactionUUID,
		PaymentMethod:   PaymentMethod,
		Status:          string(order.Status),
	}
}

func ToModelOrder(order repoModel.Order) model.Order {
	return model.Order{
		OrderUUID:       order.OrderUUID,
		UserUUID:        order.UserUUID,
		PartsUUID:       order.PartsUUID,
		TotalPrice:      order.TotalPrice,
		TransactionUUID: order.TransactionUUID,
		PaymentMethod:   order.PaymentMethod,
		Status:          model.Status(order.Status),
	}
}

func toRepoStatus(status model.Status) repoModel.Status {
	switch status {
	case model.Status_Paid:
		return repoModel.Status_Paid
	case model.Status_Canceled:
		return repoModel.Status_Canceled
	case model.Status_Assemdled:
		return repoModel.Status_Assemdled
	}
	return repoModel.Status_WaitingForPayment
}

func toModelStatus(status repoModel.Status) model.Status {
	switch status {
	case repoModel.Status_Paid:
		return model.Status_Paid
	case repoModel.Status_Canceled:
		return model.Status_Canceled
	case repoModel.Status_Assemdled:
		return model.Status_Assemdled
	}
	return model.Status_WaitingForPayment
}

func toProtoStatus(status model.Status) string {
	switch status {
	case model.Status_Paid:
		return "Paid"
	case model.Status_Canceled:
		return "Canceled"
	case model.Status_Assemdled:
		return "Assembled"
	}
	return "WaitingForPayment"
}
