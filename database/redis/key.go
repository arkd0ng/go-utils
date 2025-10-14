package redis

import (
	"context"
	"time"
)

// Del deletes one or more keys
// Del은 하나 이상의 키를 삭제합니다
func (c *Client) Del(ctx context.Context, keys ...string) error {
	if len(keys) == 0 {
		return nil
	}

	return c.executeWithRetry(ctx, func() error {
		return c.rdb.Del(ctx, keys...).Err()
	})
}

// Exists checks if one or more keys exist
// Exists는 하나 이상의 키가 존재하는지 확인합니다
func (c *Client) Exists(ctx context.Context, keys ...string) (int64, error) {
	if len(keys) == 0 {
		return 0, nil
	}

	var result int64
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.Exists(ctx, keys...).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// Expire sets a timeout on a key
// Expire는 키에 타임아웃을 설정합니다
func (c *Client) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return c.executeWithRetry(ctx, func() error {
		return c.rdb.Expire(ctx, key, expiration).Err()
	})
}

// TTL gets the time to live for a key
// TTL은 키의 남은 시간을 가져옵니다
func (c *Client) TTL(ctx context.Context, key string) (time.Duration, error) {
	var result time.Duration
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.TTL(ctx, key).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// Keys finds all keys matching the given pattern
// Keys는 주어진 패턴과 일치하는 모든 키를 찾습니다
func (c *Client) Keys(ctx context.Context, pattern string) ([]string, error) {
	var result []string
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.Keys(ctx, pattern).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// Scan incrementally iterates over keys
// Scan은 키를 점진적으로 반복합니다
func (c *Client) Scan(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64, error) {
	var keys []string
	var newCursor uint64

	err := c.executeWithRetry(ctx, func() error {
		vals, cur, err := c.rdb.Scan(ctx, cursor, match, count).Result()
		if err != nil {
			return err
		}
		keys = vals
		newCursor = cur
		return nil
	})

	return keys, newCursor, err
}

// Rename renames a key
// Rename은 키의 이름을 변경합니다
func (c *Client) Rename(ctx context.Context, key, newKey string) error {
	return c.executeWithRetry(ctx, func() error {
		return c.rdb.Rename(ctx, key, newKey).Err()
	})
}

// RenameNX renames a key only if the new key doesn't exist
// RenameNX는 새 키가 존재하지 않을 때만 키의 이름을 변경합니다
func (c *Client) RenameNX(ctx context.Context, key, newKey string) (bool, error) {
	var result bool
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.RenameNX(ctx, key, newKey).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// Type gets the type of a key
// Type은 키의 타입을 가져옵니다
func (c *Client) Type(ctx context.Context, key string) (string, error) {
	var result string
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.Type(ctx, key).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// Persist removes the expiration from a key
// Persist는 키에서 만료를 제거합니다
func (c *Client) Persist(ctx context.Context, key string) error {
	return c.executeWithRetry(ctx, func() error {
		return c.rdb.Persist(ctx, key).Err()
	})
}

// ExpireAt sets a key to expire at a specific time
// ExpireAt은 특정 시간에 만료되도록 키를 설정합니다
func (c *Client) ExpireAt(ctx context.Context, key string, tm time.Time) error {
	return c.executeWithRetry(ctx, func() error {
		return c.rdb.ExpireAt(ctx, key, tm).Err()
	})
}
