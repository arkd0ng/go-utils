package redis

import (
	"context"
)

// Pipeline executes multiple commands in a pipeline
// Pipeline은 파이프라인에서 여러 명령을 실행합니다
func (c *Client) Pipeline(ctx context.Context, fn func(pipe Pipeliner) error) error {
	return c.executeWithRetry(ctx, func() error {
		pipe := c.rdb.Pipeline()
		if err := fn(pipe); err != nil {
			return err
		}
		_, err := pipe.Exec(ctx)
		return err
	})
}

// TxPipeline executes multiple commands in a transaction pipeline
// TxPipeline은 트랜잭션 파이프라인에서 여러 명령을 실행합니다
func (c *Client) TxPipeline(ctx context.Context, fn func(pipe Pipeliner) error) error {
	return c.executeWithRetry(ctx, func() error {
		pipe := c.rdb.TxPipeline()
		if err := fn(pipe); err != nil {
			return err
		}
		_, err := pipe.Exec(ctx)
		return err
	})
}
