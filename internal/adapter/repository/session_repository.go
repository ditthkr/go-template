package repository

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"go-template/internal/domain/session"
	"time"
)

type Session struct {
	rdb *redis.Client
	ttl time.Duration
}

func (r *Session) key(uid string) string {
	return "session:" + uid
}

func (r *Session) Set(uid, jti string) error {
	ctx := context.Background()
	return r.rdb.Set(ctx, r.key(uid), jti, r.ttl).Err()
}

func (r *Session) Get(uid string) (string, bool) {
	ctx := context.Background()
	val, err := r.rdb.Get(ctx, r.key(uid)).Result()
	if errors.Is(err, redis.Nil) || err != nil {
		return "", false
	}
	return val, true
}

func NewSession(rdb *redis.Client) *Session {
	return &Session{rdb: rdb, ttl: 24 * time.Hour}
}

var _ session.Store = (*Session)(nil)
