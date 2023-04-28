package imp

import (
	"context"
	"time"

	redis "github.com/go-redis/redis/v8"
)

// RedisCache - an implementation of the CacheInterface to connect to Redis
type RedisCache struct {
	rdb *redis.Client
	ctx context.Context
}

// NewRedisCache create a rediscache
func NewRedisCache(address string, password string, db int) *RedisCache {
	return &RedisCache{
		rdb: redis.NewClient(&redis.Options{
			Addr:     address,
			Password: password, // no password set
			DB:       db,       // use default DB
		}),
		ctx: context.Background(),
	}
}

// Set a value by key
func (rc *RedisCache) Set(key string, value []byte) error {
	return rc.rdb.Set(rc.ctx, key, value, 300*time.Second).Err()
}

// Get value by key
func (rc *RedisCache) Get(dst []byte, key string) []byte {
	val, err := rc.rdb.Get(rc.ctx, key).Result()
	if err == redis.Nil {
		return []byte{}
	}

	return []byte(val)
}

// Del removes the value by key
func (rc *RedisCache) Del(keyPattern string) error {
	keys := rc.rdb.Keys(rc.ctx, keyPattern)
	for _, v := range keys.Val() {
		if err := rc.rdb.Del(rc.ctx, v).Err(); err != nil {
			return err
		}
	}
	return nil
}

// Has returns true when the key exist
func (rc *RedisCache) Has(key string) bool {
	cnt := rc.rdb.Exists(rc.ctx, key)
	return cnt.Val() > 0
}

// Reset flushes all keys
func (rc *RedisCache) Reset() {
	rc.rdb.FlushAll(rc.ctx)
}

// ListKeys lists all keys
func (rc *RedisCache) ListKeys() []string {

	var (
		cursor  uint64
		allkeys []string
	)

	for {
		keys, cursor, err := rc.rdb.Scan(rc.ctx, cursor, "*", 10).Result()
		if err != nil {
			break
		}
		allkeys = append(allkeys, keys...)
		if cursor == 0 {
			break
		}
	}

	return allkeys
}
