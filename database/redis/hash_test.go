package redis

import (
	"context"
	"reflect"
	"sort"
	"testing"
)

func TestHashOperations(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	key := testKey("hash")
	fields := map[string]interface{}{
		"name":  "Alice",
		"email": "alice@example.com",
		"age":   "30",
	}

	if err := client.HSetMap(ctx, key, fields); err != nil {
		t.Fatalf("HSetMap failed: %v", err)
	}

	if err := client.HSet(ctx, key, "country", "KR"); err != nil {
		t.Fatalf("HSet failed: %v", err)
	}

	val, err := client.HGet(ctx, key, "name")
	if err != nil {
		t.Fatalf("HGet failed: %v", err)
	}
	if val != "Alice" {
		t.Fatalf("expected Alice, got %s", val)
	}

	all, err := client.HGetAll(ctx, key)
	if err != nil {
		t.Fatalf("HGetAll failed: %v", err)
	}
	if len(all) != 4 {
		t.Fatalf("expected 4 fields, got %d", len(all))
	}

	type user struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Age     string `json:"age"`
		Country string `json:"country"`
	}

	asStruct, err := HGetAllAs[user](client, ctx, key)
	if err != nil {
		t.Fatalf("HGetAllAs failed: %v", err)
	}
	if asStruct.Name != "Alice" || asStruct.Country != "KR" {
		t.Fatalf("unexpected struct: %+v", asStruct)
	}

	if exists, err := client.HExists(ctx, key, "email"); err != nil {
		t.Fatalf("HExists failed: %v", err)
	} else if !exists {
		t.Fatal("expected field email to exist")
	}

	if length, err := client.HLen(ctx, key); err != nil {
		t.Fatalf("HLen failed: %v", err)
	} else if length != 4 {
		t.Fatalf("expected length 4, got %d", length)
	}

	keys, err := client.HKeys(ctx, key)
	if err != nil {
		t.Fatalf("HKeys failed: %v", err)
	}
	sort.Strings(keys)
	expectedKeys := []string{"age", "country", "email", "name"}
	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Fatalf("unexpected HKeys: %#v", keys)
	}

	values, err := client.HVals(ctx, key)
	if err != nil {
		t.Fatalf("HVals failed: %v", err)
	}
	sort.Strings(values)
	if len(values) != 4 {
		t.Fatalf("unexpected HVals length: %d", len(values))
	}

	if result, err := client.HIncrBy(ctx, key, "logins", 2); err != nil {
		t.Fatalf("HIncrBy failed: %v", err)
	} else if result != 2 {
		t.Fatalf("expected logins=2, got %d", result)
	}

	if result, err := client.HIncrByFloat(ctx, key, "balance", 10.5); err != nil {
		t.Fatalf("HIncrByFloat failed: %v", err)
	} else if result != 10.5 {
		t.Fatalf("expected balance=10.5, got %f", result)
	}

	if err := client.HDel(ctx, key, "email", "age"); err != nil {
		t.Fatalf("HDel failed: %v", err)
	}

	if exists, err := client.HExists(ctx, key, "email"); err != nil {
		t.Fatalf("HExists failed: %v", err)
	} else if exists {
		t.Fatal("expected email field to be deleted")
	}
}
