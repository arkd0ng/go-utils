package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// Client is a simplified Redis client wrapper
// Client는 간소화된 Redis 클라이언트 래퍼입니다
type Client struct {
	rdb    *redis.Client
	config *Config
	done   chan struct{}
}

// Pipeliner is an interface for pipeline operations
// Pipeliner는 파이프라인 작업을 위한 인터페이스입니다
type Pipeliner interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Exec(ctx context.Context) ([]redis.Cmder, error)
}

// Tx is a Redis transaction wrapper
// Tx는 Redis 트랜잭션 래퍼입니다
type Tx struct {
	tx     *redis.Tx
	client *Client
}

// PubSub is a Redis pub/sub wrapper
// PubSub은 Redis pub/sub 래퍼입니다
type PubSub struct {
	pubsub *redis.PubSub
}

// Message represents a pub/sub message
// Message는 pub/sub 메시지를 나타냅니다
type Message struct {
	Channel string
	Pattern string
	Payload string
}

// Z represents a sorted set member with score
// Z는 점수가 있는 정렬 집합 멤버를 나타냅니다
type Z struct {
	Score  float64
	Member interface{}
}
