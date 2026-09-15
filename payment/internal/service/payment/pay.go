package payment

import (
	"log"

	"github.com/google/uuid"
	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/payment/internal/model"
)

func (s *Service) PayOrder(ctx context.Context, orderId, userId string) (string, error) {
	if orderId == "" {
		return "", model.ErrArgumentEmptyString
	}
	if userId == "" {
		return "", model.ErrArgumentEmptyString
	}
	transactionId := uuid.NewString()
	log.Printf("Оплата прошла успешно. UUID транзакции: %s", transactionId)
	return transactionId, nil
}
