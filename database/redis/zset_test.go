package redis

import (
	"context"
	"reflect"
	"testing"
)

func TestSortedSetOperations(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	key := testKey("zset")

	members := map[string]float64{
		"alice": 10,
		"bob":   20,
		"carol": 15,
	}
	if err := client.ZAddMultiple(ctx, key, members); err != nil {
		t.Fatalf("ZAddMultiple failed: %v", err)
	}

	if err := client.ZAdd(ctx, key, 25, "dave"); err != nil {
		t.Fatalf("ZAdd failed: %v", err)
	}

	ordered, err := client.ZRange(ctx, key, 0, -1)
	if err != nil {
		t.Fatalf("ZRange failed: %v", err)
	}
	if !reflect.DeepEqual(ordered, []string{"alice", "carol", "bob", "dave"}) {
		t.Fatalf("unexpected ZRange order: %#v", ordered)
	}

	scoreRange, err := client.ZRangeByScore(ctx, key, 12, 22)
	if err != nil {
		t.Fatalf("ZRangeByScore failed: %v", err)
	}
	if !reflect.DeepEqual(scoreRange, []string{"carol", "bob"}) {
		t.Fatalf("unexpected ZRangeByScore: %#v", scoreRange)
	}

	score, err := client.ZScore(ctx, key, "bob")
	if err != nil {
		t.Fatalf("ZScore failed: %v", err)
	}
	if score != 20 {
		t.Fatalf("expected score 20, got %f", score)
	}

	rank, err := client.ZRank(ctx, key, "bob")
	if err != nil {
		t.Fatalf("ZRank failed: %v", err)
	}
	if rank != 2 {
		t.Fatalf("expected rank 2, got %d", rank)
	}

	reverseRank, err := client.ZRevRank(ctx, key, "alice")
	if err != nil {
		t.Fatalf("ZRevRank failed: %v", err)
	}
	if reverseRank != 3 {
		t.Fatalf("expected reverse rank 3, got %d", reverseRank)
	}

	newScore, err := client.ZIncrBy(ctx, key, 4, "carol")
	if err != nil {
		t.Fatalf("ZIncrBy failed: %v", err)
	}
	if newScore != 19 {
		t.Fatalf("expected new score 19, got %f", newScore)
	}

	card, err := client.ZCard(ctx, key)
	if err != nil {
		t.Fatalf("ZCard failed: %v", err)
	}
	if card != 4 {
		t.Fatalf("expected 4 members, got %d", card)
	}

	reverseRange, err := client.ZRevRange(ctx, key, 0, 1)
	if err != nil {
		t.Fatalf("ZRevRange failed: %v", err)
	}
	if !reflect.DeepEqual(reverseRange, []string{"dave", "bob"}) {
		t.Fatalf("unexpected reverse range: %#v", reverseRange)
	}

	if err := client.ZRem(ctx, key, "alice"); err != nil {
		t.Fatalf("ZRem failed: %v", err)
	}

	if _, err := client.ZRank(ctx, key, "alice"); err == nil {
		t.Fatal("expected error when querying removed member")
	}
}
