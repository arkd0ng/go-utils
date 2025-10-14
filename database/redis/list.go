package redis

import (
	"context"
)

// LPush inserts values at the head of the list
// LPush는 리스트의 헤드에 값을 삽입합니다
func (c *Client) LPush(ctx context.Context, key string, values ...interface{}) error {
	if len(values) == 0 {
		return nil
	}

	return c.executeWithRetry(ctx, func() error {
		return c.rdb.LPush(ctx, key, values...).Err()
	})
}

// RPush inserts values at the tail of the list
// RPush는 리스트의 테일에 값을 삽입합니다
func (c *Client) RPush(ctx context.Context, key string, values ...interface{}) error {
	if len(values) == 0 {
		return nil
	}

	return c.executeWithRetry(ctx, func() error {
		return c.rdb.RPush(ctx, key, values...).Err()
	})
}

// LPop removes and returns the first element of the list
// LPop은 리스트의 첫 번째 요소를 제거하고 반환합니다
func (c *Client) LPop(ctx context.Context, key string) (string, error) {
	var result string
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.LPop(ctx, key).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// RPop removes and returns the last element of the list
// RPop은 리스트의 마지막 요소를 제거하고 반환합니다
func (c *Client) RPop(ctx context.Context, key string) (string, error) {
	var result string
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.RPop(ctx, key).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// LRange gets a range of elements from the list
// LRange는 리스트에서 요소 범위를 가져옵니다
func (c *Client) LRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	var result []string
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.LRange(ctx, key, start, stop).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// LLen gets the length of the list
// LLen은 리스트의 길이를 가져옵니다
func (c *Client) LLen(ctx context.Context, key string) (int64, error) {
	var result int64
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.LLen(ctx, key).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// LIndex gets an element from the list by its index
// LIndex는 인덱스로 리스트의 요소를 가져옵니다
func (c *Client) LIndex(ctx context.Context, key string, index int64) (string, error) {
	var result string
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.LIndex(ctx, key, index).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// LSet sets the value of an element in the list by its index
// LSet은 인덱스로 리스트의 요소 값을 설정합니다
func (c *Client) LSet(ctx context.Context, key string, index int64, value interface{}) error {
	return c.executeWithRetry(ctx, func() error {
		return c.rdb.LSet(ctx, key, index, value).Err()
	})
}

// LRem removes elements from the list
// LRem은 리스트에서 요소를 제거합니다
func (c *Client) LRem(ctx context.Context, key string, count int64, value interface{}) error {
	return c.executeWithRetry(ctx, func() error {
		return c.rdb.LRem(ctx, key, count, value).Err()
	})
}

// LTrim trims the list to the specified range
// LTrim은 리스트를 지정된 범위로 자릅니다
func (c *Client) LTrim(ctx context.Context, key string, start, stop int64) error {
	return c.executeWithRetry(ctx, func() error {
		return c.rdb.LTrim(ctx, key, start, stop).Err()
	})
}
