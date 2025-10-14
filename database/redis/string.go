package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// Set sets a string value with optional expiration
// Set은 선택적 만료 시간과 함께 문자열 값을 설정합니다
func (c *Client) Set(ctx context.Context, key string, value interface{}, expiration ...time.Duration) error {
	exp := time.Duration(0)
	if len(expiration) > 0 {
		exp = expiration[0]
	}

	return c.executeWithRetry(ctx, func() error {
		return c.rdb.Set(ctx, key, value, exp).Err()
	})
}

// Get gets a string value
// Get은 문자열 값을 가져옵니다
func (c *Client) Get(ctx context.Context, key string) (string, error) {
	var result string
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.Get(ctx, key).Result()
		if err != nil {
			if err == redis.Nil {
				return ErrNil
			}
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// GetAs gets a value and unmarshals it to the specified type
// GetAs는 값을 가져와 지정된 타입으로 언마샬합니다
func GetAs[T any](c *Client, ctx context.Context, key string) (T, error) {
	var result T

	val, err := c.Get(ctx, key)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal([]byte(val), &result); err != nil {
		return result, fmt.Errorf("failed to unmarshal value: %w", err)
	}

	return result, nil
}

// MGet gets multiple values
// MGet은 여러 값을 가져옵니다
func (c *Client) MGet(ctx context.Context, keys ...string) ([]string, error) {
	var result []string
	err := c.executeWithRetry(ctx, func() error {
		vals, err := c.rdb.MGet(ctx, keys...).Result()
		if err != nil {
			return err
		}

		result = make([]string, len(vals))
		for i, val := range vals {
			if val != nil {
				result[i] = val.(string)
			}
		}
		return nil
	})
	return result, err
}

// MSet sets multiple key-value pairs
// MSet은 여러 키-값 쌍을 설정합니다
func (c *Client) MSet(ctx context.Context, pairs map[string]interface{}) error {
	if len(pairs) == 0 {
		return nil
	}

	// Convert map to slice / map을 slice로 변환
	args := make([]interface{}, 0, len(pairs)*2)
	for k, v := range pairs {
		args = append(args, k, v)
	}

	return c.executeWithRetry(ctx, func() error {
		return c.rdb.MSet(ctx, args...).Err()
	})
}

// Incr increments a counter
// Incr은 카운터를 증가시킵니다
func (c *Client) Incr(ctx context.Context, key string) (int64, error) {
	var result int64
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.Incr(ctx, key).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// IncrBy increments a counter by the given amount
// IncrBy는 주어진 양만큼 카운터를 증가시킵니다
func (c *Client) IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	var result int64
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.IncrBy(ctx, key, value).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// Decr decrements a counter
// Decr은 카운터를 감소시킵니다
func (c *Client) Decr(ctx context.Context, key string) (int64, error) {
	var result int64
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.Decr(ctx, key).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// DecrBy decrements a counter by the given amount
// DecrBy는 주어진 양만큼 카운터를 감소시킵니다
func (c *Client) DecrBy(ctx context.Context, key string, value int64) (int64, error) {
	var result int64
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.DecrBy(ctx, key, value).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// Append appends a value to a key
// Append는 키에 값을 추가합니다
func (c *Client) Append(ctx context.Context, key, value string) (int64, error) {
	var result int64
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.Append(ctx, key, value).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// GetRange gets a substring of the string stored at a key
// GetRange는 키에 저장된 문자열의 부분 문자열을 가져옵니다
func (c *Client) GetRange(ctx context.Context, key string, start, end int64) (string, error) {
	var result string
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.GetRange(ctx, key, start, end).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// SetNX sets a key only if it doesn't exist
// SetNX는 키가 존재하지 않을 때만 설정합니다
func (c *Client) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	var result bool
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.SetNX(ctx, key, value, expiration).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// SetEX sets a key with expiration
// SetEX는 만료 시간과 함께 키를 설정합니다
func (c *Client) SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.executeWithRetry(ctx, func() error {
		return c.rdb.SetEx(ctx, key, value, expiration).Err()
	})
}
