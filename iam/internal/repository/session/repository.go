package session

import (
	"context"
	"fmt"
	"time"

	"github.com/TikhonovNikita1996/Go-microservises-project/iam/internal/model"
	repoConverter "github.com/TikhonovNikita1996/Go-microservises-project/iam/internal/repository/converter"
	"github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/cache"
	redigo "github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"

	repoModel "github.com/TikhonovNikita1996/Go-microservises-project/iam/internal/repository/model"
)

const (
	cacheKeyPrefix = "user:session:"
)

type repository struct {
	cache cache.RedisClient
}

func NewRepository(cache cache.RedisClient) *repository {
	return &repository{
		cache: cache,
	}
}

func (r *repository) getCacheKey(uuid string) string {
	return fmt.Sprintf("%s%s", cacheKeyPrefix, uuid)
}

func (r *repository) CreateSession(ctx context.Context, uuid string) (model.Session, error) {
	cacheKey := r.getCacheKey(uuid)

	values, err := r.cache.HGetAll(ctx, cacheKey)
	if err != nil {
		if errors.Is(err, redigo.ErrNil) {
			return model.Session{}, model.ErrSessionNotFound
		}
		return model.Session{}, err
	}

	if len(values) == 0 {
		return model.Session{}, model.ErrSessionNotFound
	}

	var sessionRedisView repoModel.SessionRedisView
	err = redigo.ScanStruct(values, &sessionRedisView)
	if err != nil {
		return model.Session{}, err
	}

	return repoConverter.ToRedisViewSession()
}

func (r *repository) Set(ctx context.Context, uuid string, sighting model.Sighting, ttl time.Duration) error {
	cacheKey := r.getCacheKey(uuid)

	redisView := repoConverter.SightingToRedisView(ctx, sighting)

	err := r.cache.HashSet(ctx, cacheKey, redisView)
	if err != nil {
		return err
	}

	return r.cache.Expire(ctx, cacheKey, ttl)
}

func (r *repository) Delete(ctx context.Context, uuid string) error {
	cacheKey := r.getCacheKey(uuid)
	return r.cache.Del(ctx, cacheKey)
}
