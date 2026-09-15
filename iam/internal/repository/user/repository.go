package user

import (
	"sync"

	"github.com/jackc/pgx/v5"
)

type userRepository struct {
	mu           sync.RWMutex
	dbConnection *pgx.Conn
}

func NewRepository(connection *pgx.Conn) *userRepository {
	return &userRepository{
		dbConnection: connection,
	}
}
