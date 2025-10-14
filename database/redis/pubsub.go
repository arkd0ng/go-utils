package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// Publish publishes a message to a channel
// Publish는 채널에 메시지를 발행합니다
func (c *Client) Publish(ctx context.Context, channel string, message interface{}) error {
	return c.executeWithRetry(ctx, func() error {
		return c.rdb.Publish(ctx, channel, message).Err()
	})
}

// Subscribe subscribes to channels
// Subscribe는 채널을 구독합니다
func (c *Client) Subscribe(ctx context.Context, channels ...string) (*PubSub, error) {
	pubsub := c.rdb.Subscribe(ctx, channels...)
	return &PubSub{pubsub: pubsub}, nil
}

// PSubscribe subscribes to channels matching patterns
// PSubscribe는 패턴과 일치하는 채널을 구독합니다
func (c *Client) PSubscribe(ctx context.Context, patterns ...string) (*PubSub, error) {
	pubsub := c.rdb.PSubscribe(ctx, patterns...)
	return &PubSub{pubsub: pubsub}, nil
}

// Channel returns the channel for receiving messages
// Channel은 메시지를 받기 위한 채널을 반환합니다
func (ps *PubSub) Channel() <-chan *redis.Message {
	return ps.pubsub.Channel()
}

// Close closes the pub/sub connection
// Close는 pub/sub 연결을 닫습니다
func (ps *PubSub) Close() error {
	return ps.pubsub.Close()
}

// Unsubscribe unsubscribes from channels
// Unsubscribe는 채널 구독을 취소합니다
func (ps *PubSub) Unsubscribe(ctx context.Context, channels ...string) error {
	return ps.pubsub.Unsubscribe(ctx, channels...)
}

// PUnsubscribe unsubscribes from patterns
// PUnsubscribe는 패턴 구독을 취소합니다
func (ps *PubSub) PUnsubscribe(ctx context.Context, patterns ...string) error {
	return ps.pubsub.PUnsubscribe(ctx, patterns...)
}

// Receive receives a message
// Receive는 메시지를 받습니다
func (ps *PubSub) Receive(ctx context.Context) (interface{}, error) {
	return ps.pubsub.Receive(ctx)
}
