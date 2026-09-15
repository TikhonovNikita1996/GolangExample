package repository

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
)

func (r *OrderRepository) CreateOrder(ctx context.Context, req *model.CreateOrderRequest) (string, error) {
	query := `SELECT add_order($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.dbConnection.Exec(ctx, query,
		req.OrderUUID,
		req.UserUUID,
		req.PartUuids,
		req.TotalPrice,
		model.Status_WaitingForPayment,
		nil,
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("failed to execute add_order: %w", err)
	}

	return req.OrderUUID, nil
}
