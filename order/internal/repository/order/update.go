package repository

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
)

func (r *OrderRepository) UpdateOrder(ctx context.Context, updateModel model.OrderUpdate) error {
	query := `SELECT update_order($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.dbConnection.Exec(ctx, query,
		updateModel.OrderUUID,
		updateModel.UserUUID,
		updateModel.PartsUUID,
		updateModel.TotalPrice,
		updateModel.Status,
		updateModel.TransactionUUID,
		updateModel.PaymentMethod,
	)
	if err != nil {
		return fmt.Errorf("failed to execute update_order: %w", err)
	}
	return nil
}
