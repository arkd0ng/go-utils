package redis

import (
	"context"
	"encoding/json"
	"fmt"
)

// HSet sets a hash field
// HSet은 해시 필드를 설정합니다
func (c *Client) HSet(ctx context.Context, key, field string, value interface{}) error {
	return c.executeWithRetry(ctx, func() error {
		return c.rdb.HSet(ctx, key, field, value).Err()
	})
}

// HSetMap sets multiple hash fields
// HSetMap은 여러 해시 필드를 설정합니다
func (c *Client) HSetMap(ctx context.Context, key string, fields map[string]interface{}) error {
	if len(fields) == 0 {
		return nil
	}

	return c.executeWithRetry(ctx, func() error {
		return c.rdb.HSet(ctx, key, fields).Err()
	})
}

// HGet gets a hash field
// HGet은 해시 필드를 가져옵니다
func (c *Client) HGet(ctx context.Context, key, field string) (string, error) {
	var result string
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.HGet(ctx, key, field).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// HGetAll gets all hash fields
// HGetAll은 모든 해시 필드를 가져옵니다
func (c *Client) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	var result map[string]string
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.HGetAll(ctx, key).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// HGetAllAs gets all hash fields and unmarshals to a struct
// HGetAllAs는 모든 해시 필드를 가져와 구조체로 언마샬합니다
func HGetAllAs[T any](c *Client, ctx context.Context, key string) (T, error) {
	var result T

	fields, err := c.HGetAll(ctx, key)
	if err != nil {
		return result, err
	}

	// Convert map to JSON and unmarshal to struct
	// map을 JSON으로 변환하고 구조체로 언마샬
	data, err := json.Marshal(fields)
	if err != nil {
		return result, fmt.Errorf("failed to marshal fields: %w", err)
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return result, fmt.Errorf("failed to unmarshal to struct: %w", err)
	}

	return result, nil
}

// HDel deletes hash fields
// HDel은 해시 필드를 삭제합니다
func (c *Client) HDel(ctx context.Context, key string, fields ...string) error {
	if len(fields) == 0 {
		return nil
	}

	return c.executeWithRetry(ctx, func() error {
		return c.rdb.HDel(ctx, key, fields...).Err()
	})
}

// HExists checks if a hash field exists
// HExists는 해시 필드가 존재하는지 확인합니다
func (c *Client) HExists(ctx context.Context, key, field string) (bool, error) {
	var result bool
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.HExists(ctx, key, field).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// HLen gets the number of fields in a hash
// HLen은 해시의 필드 수를 가져옵니다
func (c *Client) HLen(ctx context.Context, key string) (int64, error) {
	var result int64
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.HLen(ctx, key).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// HKeys gets all field names in a hash
// HKeys는 해시의 모든 필드 이름을 가져옵니다
func (c *Client) HKeys(ctx context.Context, key string) ([]string, error) {
	var result []string
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.HKeys(ctx, key).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// HVals gets all values in a hash
// HVals는 해시의 모든 값을 가져옵니다
func (c *Client) HVals(ctx context.Context, key string) ([]string, error) {
	var result []string
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.HVals(ctx, key).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// HIncrBy increments a hash field by the given amount
// HIncrBy는 주어진 양만큼 해시 필드를 증가시킵니다
func (c *Client) HIncrBy(ctx context.Context, key, field string, incr int64) (int64, error) {
	var result int64
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.HIncrBy(ctx, key, field, incr).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// HIncrByFloat increments a hash field by the given float amount
// HIncrByFloat는 주어진 float 양만큼 해시 필드를 증가시킵니다
func (c *Client) HIncrByFloat(ctx context.Context, key, field string, incr float64) (float64, error) {
	var result float64
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.HIncrByFloat(ctx, key, field, incr).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}
