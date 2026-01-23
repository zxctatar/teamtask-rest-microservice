package myredis

import (
	"context"
	"errors"
	"time"
	"userservice/internal/repository/session"

	"github.com/redis/go-redis/v9"
)

var (
	invalidId uint32 = 0
)

type Redis struct {
	client *redis.Client
	ttl    *time.Duration
}

func NewRedis(client *redis.Client, ttl *time.Duration) *Redis {
	return &Redis{
		client: client,
		ttl:    ttl,
	}
}

func (r *Redis) Save(ctx context.Context, sessionId string, userId uint32) error {
	return r.client.Set(ctx, sessionId, userId, *r.ttl).Err()
}

func (r *Redis) Get(ctx context.Context, sessionId string) (uint32, error) {
	id, err := r.client.Get(ctx, sessionId).Int()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return invalidId, session.ErrKeyNotFound
		}
		return invalidId, err
	}
	return uint32(id), nil
}
