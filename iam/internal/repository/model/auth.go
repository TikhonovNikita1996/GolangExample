package model

import "time"

type Session struct {
	Uuid      string
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiresAt time.Time
}

// SessionRedisView - модель для хранения в Redis hash map
type SessionRedisView struct {
	UUID        string `redis:"uuid"`
	CreatedAtNs *int64 `redis:"created_at,omitempty"`
	UpdatedAtNs *int64 `redis:"updated_at,omitempty"`
	ExpiresAt   *int64 `redis:"deleted_at,omitempty"`
}
