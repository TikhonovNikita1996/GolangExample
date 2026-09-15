package config

import (
	"os"

	"github.com/joho/godotenv"

	"github.com/TikhonovNikita1996/Go-microservises-project/iam/internal/config/env"
)

var appConfig *config

type config struct {
	Logger   LoggerConfig
	Redis    RedisConfig
	Postgres PostgresConfig
}

func Load(path ...string) error {
	err := godotenv.Load(path...)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	loggerCfg, err := env.NewLoggerConfig()
	if err != nil {
		return err
	}

	redisCfg, err := env.NewRedisConfig()
	if err != nil {
		return err
	}

	postgresCfg, err := env.NewPostgresConfig()
	if err != nil {
		return err
	}

	appConfig = &config{
		Logger:   loggerCfg,
		Redis:    redisCfg,
		Postgres: postgresCfg,
	}

	return nil
}

func AppConfig() *config {
	return appConfig
}
