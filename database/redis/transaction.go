package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// Transaction executes commands in a Redis transaction with optimistic locking
// Transaction은 낙관적 잠금을 사용하여 Redis 트랜잭션에서 명령을 실행합니다
func (c *Client) Transaction(ctx context.Context, fn func(tx *Tx) error, keys ...string) error {
	return c.executeWithRetry(ctx, func() error {
		return c.rdb.Watch(ctx, func(tx *redis.Tx) error {
			txClient := &Tx{
				tx:     tx,
				client: c,
			}
			return fn(txClient)
		}, keys...)
	})
}

// Exec executes a transaction
// Exec은 트랜잭션을 실행합니다
func (tx *Tx) Exec(ctx context.Context, fn func(pipe redis.Pipeliner) error) error {
	pipe := tx.tx.TxPipeline()
	if err := fn(pipe); err != nil {
		return err
	}
	_, err := pipe.Exec(ctx)
	return err
}

// Get gets a value within a transaction
// Get은 트랜잭션 내에서 값을 가져옵니다
func (tx *Tx) Get(ctx context.Context, key string) (string, error) {
	val, err := tx.tx.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", ErrNil
		}
		return "", err
	}
	return val, nil
}

// Set sets a value within a transaction pipeline
// Set은 트랜잭션 파이프라인 내에서 값을 설정합니다
func (tx *Tx) Set(ctx context.Context, pipe redis.Pipeliner, key string, value interface{}) error {
	return pipe.Set(ctx, key, value, 0).Err()
}

// Del deletes keys within a transaction pipeline
// Del은 트랜잭션 파이프라인 내에서 키를 삭제합니다
func (tx *Tx) Del(ctx context.Context, pipe redis.Pipeliner, keys ...string) error {
	return pipe.Del(ctx, keys...).Err()
}
