package config

import (
	"time"
)

type LoggerConfig interface {
	Level() string
	AsJson() bool
}

type PostgresConfig interface {
	URI() string
	DatabaseName() string
	Port() string
	Host() string
	MigrationsPath() string
}

type RedisConfig interface {
	Address() string
	ConnectionTimeout() time.Duration
	MaxIdle() int
	IdleTimeout() time.Duration
	CacheTTL() time.Duration
}
