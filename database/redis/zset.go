package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// ZAdd adds a member to a sorted set
// ZAdd는 정렬 집합에 멤버를 추가합니다
func (c *Client) ZAdd(ctx context.Context, key string, score float64, member interface{}) error {
	return c.executeWithRetry(ctx, func() error {
		return c.rdb.ZAdd(ctx, key, redis.Z{
			Score:  score,
			Member: member,
		}).Err()
	})
}

// ZAddMultiple adds multiple members to a sorted set
// ZAddMultiple은 정렬 집합에 여러 멤버를 추가합니다
func (c *Client) ZAddMultiple(ctx context.Context, key string, members map[string]float64) error {
	if len(members) == 0 {
		return nil
	}

	zs := make([]redis.Z, 0, len(members))
	for member, score := range members {
		zs = append(zs, redis.Z{
			Score:  score,
			Member: member,
		})
	}

	return c.executeWithRetry(ctx, func() error {
		return c.rdb.ZAdd(ctx, key, zs...).Err()
	})
}

// ZRange gets members in a sorted set by index range
// ZRange는 인덱스 범위로 정렬 집합의 멤버를 가져옵니다
func (c *Client) ZRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	var result []string
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.ZRange(ctx, key, start, stop).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// ZRangeByScore gets members in a sorted set by score range
// ZRangeByScore는 점수 범위로 정렬 집합의 멤버를 가져옵니다
func (c *Client) ZRangeByScore(ctx context.Context, key string, min, max float64) ([]string, error) {
	var result []string
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.ZRangeByScore(ctx, key, &redis.ZRangeBy{
			Min: fmt.Sprintf("%f", min),
			Max: fmt.Sprintf("%f", max),
		}).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// ZRem removes members from a sorted set
// ZRem은 정렬 집합에서 멤버를 제거합니다
func (c *Client) ZRem(ctx context.Context, key string, members ...interface{}) error {
	if len(members) == 0 {
		return nil
	}

	return c.executeWithRetry(ctx, func() error {
		return c.rdb.ZRem(ctx, key, members...).Err()
	})
}

// ZCard gets the number of members in a sorted set
// ZCard는 정렬 집합의 멤버 수를 가져옵니다
func (c *Client) ZCard(ctx context.Context, key string) (int64, error) {
	var result int64
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.ZCard(ctx, key).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// ZScore gets the score of a member in a sorted set
// ZScore는 정렬 집합에서 멤버의 점수를 가져옵니다
func (c *Client) ZScore(ctx context.Context, key string, member string) (float64, error) {
	var result float64
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.ZScore(ctx, key, member).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// ZIncrBy increments the score of a member in a sorted set
// ZIncrBy는 정렬 집합에서 멤버의 점수를 증가시킵니다
func (c *Client) ZIncrBy(ctx context.Context, key string, increment float64, member string) (float64, error) {
	var result float64
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.ZIncrBy(ctx, key, increment, member).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// ZRank gets the rank of a member in a sorted set (ascending)
// ZRank는 정렬 집합에서 멤버의 순위를 가져옵니다 (오름차순)
func (c *Client) ZRank(ctx context.Context, key, member string) (int64, error) {
	var result int64
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.ZRank(ctx, key, member).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// ZRevRank gets the rank of a member in a sorted set (descending)
// ZRevRank는 정렬 집합에서 멤버의 순위를 가져옵니다 (내림차순)
func (c *Client) ZRevRank(ctx context.Context, key, member string) (int64, error) {
	var result int64
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.ZRevRank(ctx, key, member).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// ZRevRange gets members in a sorted set by index range in reverse order
// ZRevRange는 역순으로 인덱스 범위로 정렬 집합의 멤버를 가져옵니다
func (c *Client) ZRevRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	var result []string
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.ZRevRange(ctx, key, start, stop).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}
