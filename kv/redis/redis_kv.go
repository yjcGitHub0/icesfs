package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"icesfs/command/vars"
	"icesfs/errors"
	"icesfs/kv"
	"icesfs/log"
)

func (store *KvStore) KvPut(ctx context.Context, key string, val []byte) error {
	_, err := store.client.Set(ctx, key, val, 0).Result()
	if err != nil {
		log.Errorw("redis kv put error", vars.UUIDKey, ctx.Value(vars.UUIDKey), vars.UserKey, ctx.Value(vars.UserKey), vars.ErrorKey, err.Error(), "key", key)
		return errors.GetAPIErr(errors.ErrKvSever)
	}
	return nil
}

func (store *KvStore) KvGet(ctx context.Context, key string) ([]byte, error) {
	val, err := store.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, kv.NotFound
		} else {
			log.Errorw("redis kv get error", vars.UUIDKey, ctx.Value(vars.UUIDKey), vars.UserKey, ctx.Value(vars.UserKey), vars.ErrorKey, err.Error(), "key", key)
			return nil, errors.GetAPIErr(errors.ErrKvSever)
		}
	}
	return []byte(val), nil
}

func (store *KvStore) KvDelete(ctx context.Context, key string) (bool, error) {
	ret, err := store.client.Del(ctx, key).Result()
	if err != nil {
		log.Errorw("redis kv delete error", vars.UUIDKey, ctx.Value(vars.UUIDKey), vars.UserKey, ctx.Value(vars.UserKey), vars.ErrorKey, err.Error(), "key", key)
		err = errors.GetAPIErr(errors.ErrKvSever)
	}
	return ret != 0, nil
}
