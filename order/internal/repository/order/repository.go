package repository

import (
	"sync"

	"github.com/jackc/pgx/v5"
)

type OrderRepository struct {
	mu           sync.RWMutex
	dbConnection *pgx.Conn
}

func NewRepository(connection *pgx.Conn) *OrderRepository {
	return &OrderRepository{
		dbConnection: connection,
	}
}
