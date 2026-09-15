package repository

import (
	"fmt"

	"github.com/go-faster/errors"
	"github.com/jackc/pgx/v5"
	"golang.org/x/net/context"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/model"
	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/repository/converter"
	repoModel "github.com/TikhonovNikita1996/Go-microservises-project/order/internal/repository/model"
)

func (r *OrderRepository) GetOrderByUuid(ctx context.Context, uuid string) (model.Order, error) {
	query := `SELECT * FROM get_order($1);`
	var order repoModel.Order
	row := r.dbConnection.QueryRow(ctx, query, uuid)
	err := row.Scan(
		&order.OrderUUID,
		&order.UserUUID,
		&order.PartsUUID,
		&order.TotalPrice,
		&order.TransactionUUID,
		&order.PaymentMethod,
		&order.Status,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Order{}, model.ErrOrderNotFound
		}
		return model.Order{}, fmt.Errorf("failed to scan order: %w", err)
	}
	return converter.ToModelOrder(order), nil
}
