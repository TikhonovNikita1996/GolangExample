package order

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
)

func (s *OrderService) PayForOrderByUuid(ctx context.Context, userUuid, paymentMethod, orderId string) (string, error) {
	order, err := s.orderRepository.GetOrderByUuid(ctx, orderId)
	if err != nil {
		return "", model.ErrOrderNotFound
	}
	switch order.Status {
	case model.Status_Paid, model.Status_Canceled:
		return "", model.ErrOrderCanNotBePayed
	default:
		order.Status = model.Status_Paid
	}
	paymentInfo := model.PaymentInfo{
		UserUuid:      userUuid,
		PaymentMethod: paymentMethod,
		OrderUuid:     orderId,
	}
	transactionUuid, err := s.paymentClient.PayOrder(ctx, paymentInfo)
	if err != nil {
		return "", fmt.Errorf("paying order: %w", err)
	}
	request := &model.PaymentRequest{
		OrderUuid:       orderId,
		PaymentMethod:   paymentMethod,
		TransactionUuid: transactionUuid,
		PaymentStatus:   model.Status_Paid,
	}
	ok := s.orderRepository.PayForOrderByUuid(ctx, request)
	if ok != nil {
		return "", fmt.Errorf("payment error: %w", err)
	}

	event := model.OrderPaid{
		EventUuid:       uuid.NewString(),
		OrderUuid:       orderId,
		UserUuid:        userUuid,
		PaymentMethod:   paymentMethod,
		TransactionUuid: transactionUuid,
	}

	s.orderPaidProducer.ProduceOrderPaid(ctx, event)

	return orderId, nil
}
