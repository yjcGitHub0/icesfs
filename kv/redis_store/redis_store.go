package redis_store

import (
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

type RedisStore struct {
	client  redis.UniversalClient
	redSync *redsync.Redsync
}

func NewRedisStore(hostPort, password string, database int) *RedisStore {
	kvStore := &RedisStore{}
	kvStore.client = redis.NewClient(
		&redis.Options{
			Addr:     hostPort,
			Password: password,
			DB:       database,
		},
	)
	kvStore.redSync = redsync.New(goredis.NewPool(kvStore.client))
	return kvStore
}
