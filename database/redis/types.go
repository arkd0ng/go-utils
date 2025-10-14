package redis

import (
	"github.com/redis/go-redis/v9"
)

// Client is a simplified Redis client wrapper
// Client는 간소화된 Redis 클라이언트 래퍼입니다
type Client struct {
	rdb    *redis.Client
	config *Config
	done   chan struct{}
}

// Pipeliner is a type alias for redis.Pipeliner
// This allows users to use pipeline operations without importing redis/go-redis directly
// Pipeliner는 redis.Pipeliner의 타입 별칭입니다
// 사용자가 redis/go-redis를 직접 import하지 않고 파이프라인 작업을 사용할 수 있게 합니다
type Pipeliner = redis.Pipeliner

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
