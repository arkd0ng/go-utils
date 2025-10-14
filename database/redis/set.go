package redis

import (
	"context"
)

// SAdd adds members to a set
// SAdd는 집합에 멤버를 추가합니다
func (c *Client) SAdd(ctx context.Context, key string, members ...interface{}) error {
	if len(members) == 0 {
		return nil
	}

	return c.executeWithRetry(ctx, func() error {
		return c.rdb.SAdd(ctx, key, members...).Err()
	})
}

// SRem removes members from a set
// SRem은 집합에서 멤버를 제거합니다
func (c *Client) SRem(ctx context.Context, key string, members ...interface{}) error {
	if len(members) == 0 {
		return nil
	}

	return c.executeWithRetry(ctx, func() error {
		return c.rdb.SRem(ctx, key, members...).Err()
	})
}

// SMembers gets all members of a set
// SMembers는 집합의 모든 멤버를 가져옵니다
func (c *Client) SMembers(ctx context.Context, key string) ([]string, error) {
	var result []string
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.SMembers(ctx, key).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// SIsMember checks if a member exists in a set
// SIsMember는 집합에 멤버가 존재하는지 확인합니다
func (c *Client) SIsMember(ctx context.Context, key string, member interface{}) (bool, error) {
	var result bool
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.SIsMember(ctx, key, member).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// SCard gets the number of members in a set
// SCard는 집합의 멤버 수를 가져옵니다
func (c *Client) SCard(ctx context.Context, key string) (int64, error) {
	var result int64
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.SCard(ctx, key).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// SUnion gets the union of multiple sets
// SUnion은 여러 집합의 합집합을 가져옵니다
func (c *Client) SUnion(ctx context.Context, keys ...string) ([]string, error) {
	if len(keys) == 0 {
		return nil, nil
	}

	var result []string
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.SUnion(ctx, keys...).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// SInter gets the intersection of multiple sets
// SInter는 여러 집합의 교집합을 가져옵니다
func (c *Client) SInter(ctx context.Context, keys ...string) ([]string, error) {
	if len(keys) == 0 {
		return nil, nil
	}

	var result []string
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.SInter(ctx, keys...).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// SDiff gets the difference of multiple sets
// SDiff는 여러 집합의 차집합을 가져옵니다
func (c *Client) SDiff(ctx context.Context, keys ...string) ([]string, error) {
	if len(keys) == 0 {
		return nil, nil
	}

	var result []string
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.SDiff(ctx, keys...).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// SPop removes and returns a random member from a set
// SPop은 집합에서 랜덤 멤버를 제거하고 반환합니다
func (c *Client) SPop(ctx context.Context, key string) (string, error) {
	var result string
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.SPop(ctx, key).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}

// SRandMember gets random members from a set
// SRandMember는 집합에서 랜덤 멤버를 가져옵니다
func (c *Client) SRandMember(ctx context.Context, key string, count int64) ([]string, error) {
	var result []string
	err := c.executeWithRetry(ctx, func() error {
		val, err := c.rdb.SRandMemberN(ctx, key, count).Result()
		if err != nil {
			return err
		}
		result = val
		return nil
	})
	return result, err
}
