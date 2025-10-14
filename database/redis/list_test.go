package redis

import (
	"context"
	"reflect"
	"testing"
)

func TestListOperations(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	key := testKey("list")

	if err := client.RPush(ctx, key, "a", "b", "c"); err != nil {
		t.Fatalf("RPush failed: %v", err)
	}
	if err := client.LPush(ctx, key, "start"); err != nil {
		t.Fatalf("LPush failed: %v", err)
	}

	length, err := client.LLen(ctx, key)
	if err != nil {
		t.Fatalf("LLen failed: %v", err)
	}
	if length != 4 {
		t.Fatalf("expected length 4, got %d", length)
	}

	items, err := client.LRange(ctx, key, 0, -1)
	if err != nil {
		t.Fatalf("LRange failed: %v", err)
	}
	expected := []string{"start", "a", "b", "c"}
	if !reflect.DeepEqual(items, expected) {
		t.Fatalf("unexpected list contents: %#v", items)
	}

	item, err := client.LPop(ctx, key)
	if err != nil {
		t.Fatalf("LPop failed: %v", err)
	}
	if item != "start" {
		t.Fatalf("expected start, got %s", item)
	}

	item, err = client.RPop(ctx, key)
	if err != nil {
		t.Fatalf("RPop failed: %v", err)
	}
	if item != "c" {
		t.Fatalf("expected c, got %s", item)
	}

	if err := client.LSet(ctx, key, 0, "alpha"); err != nil {
		t.Fatalf("LSet failed: %v", err)
	}

	indexVal, err := client.LIndex(ctx, key, 0)
	if err != nil {
		t.Fatalf("LIndex failed: %v", err)
	}
	if indexVal != "alpha" {
		t.Fatalf("expected alpha at index 0, got %s", indexVal)
	}

	if err := client.RPush(ctx, key, "c", "d"); err != nil {
		t.Fatalf("RPush additional elements failed: %v", err)
	}

	if err := client.LRem(ctx, key, 1, "c"); err != nil {
		t.Fatalf("LRem failed: %v", err)
	}

	trimmed, err := client.LRange(ctx, key, 0, -1)
	if err != nil {
		t.Fatalf("LRange after LRem failed: %v", err)
	}
	expected = []string{"alpha", "b", "d"}
	if !reflect.DeepEqual(trimmed, expected) {
		t.Fatalf("unexpected list contents after remove: %#v", trimmed)
	}

	if err := client.LTrim(ctx, key, 1, -1); err != nil {
		t.Fatalf("LTrim failed: %v", err)
	}

	final, err := client.LRange(ctx, key, 0, -1)
	if err != nil {
		t.Fatalf("LRange after trim failed: %v", err)
	}
	if !reflect.DeepEqual(final, []string{"b", "d"}) {
		t.Fatalf("expected list [b d], got %#v", final)
	}
}
