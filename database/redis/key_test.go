package redis

import (
	"context"
	"strings"
	"testing"
	"time"
)

func TestKeyOperations(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	key := testKey("key")
	if err := client.Set(ctx, key, "value"); err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	count, err := client.Exists(ctx, key)
	if err != nil {
		t.Fatalf("Exists failed: %v", err)
	}
	if count != 1 {
		t.Fatalf("expected Exists to return 1, got %d", count)
	}

	if err := client.Expire(ctx, key, time.Second); err != nil {
		t.Fatalf("Expire failed: %v", err)
	}

	ttl, err := client.TTL(ctx, key)
	if err != nil {
		t.Fatalf("TTL failed: %v", err)
	}
	if ttl <= 0 {
		t.Fatalf("expected positive TTL, got %v", ttl)
	}

	if err := client.Persist(ctx, key); err != nil {
		t.Fatalf("Persist failed: %v", err)
	}

	newKey := testKey("key:renamed")
	if err := client.Rename(ctx, key, newKey); err != nil {
		t.Fatalf("Rename failed: %v", err)
	}

	if keyType, err := client.Type(ctx, newKey); err != nil {
		t.Fatalf("Type failed: %v", err)
	} else if keyType != "string" {
		t.Fatalf("expected type string, got %s", keyType)
	}

	otherKey := testKey("key:other")
	if err := client.Set(ctx, otherKey, "other"); err != nil {
		t.Fatalf("Set otherKey failed: %v", err)
	}

	if renamed, err := client.RenameNX(ctx, newKey, otherKey); err != nil {
		t.Fatalf("RenameNX failed: %v", err)
	} else if renamed {
		t.Fatal("expected RenameNX to fail because destination exists")
	}

	prefix := strings.Split(otherKey, ":")[0]
	matching, err := client.Keys(ctx, prefix+"*")
	if err != nil {
		t.Fatalf("Keys failed: %v", err)
	}
	if len(matching) == 0 {
		t.Fatal("expected Keys to return results")
	}

	var collected []string
	var cursor uint64
	for {
		keys, nextCursor, err := client.Scan(ctx, cursor, prefix+"*", 10)
		if err != nil {
			t.Fatalf("Scan failed: %v", err)
		}
		collected = append(collected, keys...)
		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}
	if len(collected) == 0 {
		t.Fatal("expected Scan to return keys")
	}

	expireAtKey := testKey("key:expireat")
	if err := client.Set(ctx, expireAtKey, "temp"); err != nil {
		t.Fatalf("Set expireAtKey failed: %v", err)
	}
	if err := client.ExpireAt(ctx, expireAtKey, time.Now().Add(2*time.Second)); err != nil {
		t.Fatalf("ExpireAt failed: %v", err)
	}

	if err := client.Del(ctx, newKey, otherKey, expireAtKey); err != nil {
		t.Fatalf("Del failed: %v", err)
	}

	if count, err := client.Exists(ctx, newKey, otherKey, expireAtKey); err != nil {
		t.Fatalf("Exists after delete failed: %v", err)
	} else if count != 0 {
		t.Fatalf("expected deleted keys to not exist, got %d", count)
	}
}
