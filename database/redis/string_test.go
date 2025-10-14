package redis

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestStringBasicOperations(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	key := testKey("string")
	if err := client.Set(ctx, key, "hello"); err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	val, err := client.Get(ctx, key)
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if val != "hello" {
		t.Fatalf("expected hello, got %s", val)
	}

	if length, err := client.Append(ctx, key, " world"); err != nil {
		t.Fatalf("Append failed: %v", err)
	} else if length != int64(len("hello world")) {
		t.Fatalf("unexpected append length: %d", length)
	}

	rangeVal, err := client.GetRange(ctx, key, 0, 4)
	if err != nil {
		t.Fatalf("GetRange failed: %v", err)
	}
	if rangeVal != "hello" {
		t.Fatalf("expected prefix hello, got %s", rangeVal)
	}

	keys := []string{testKey("string"), testKey("string"), testKey("string")}
	pairs := map[string]interface{}{
		keys[0]: "v1",
		keys[1]: "v2",
		keys[2]: "v3",
	}

	if err := client.MSet(ctx, pairs); err != nil {
		t.Fatalf("MSet failed: %v", err)
	}

	values, err := client.MGet(ctx, keys...)
	if err != nil {
		t.Fatalf("MGet failed: %v", err)
	}
	if !reflect.DeepEqual(values, []string{"v1", "v2", "v3"}) {
		t.Fatalf("unexpected MGet values: %#v", values)
	}
}

func TestStringCounterOperations(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	counter := testKey("counter")

	if val, err := client.Incr(ctx, counter); err != nil {
		t.Fatalf("Incr failed: %v", err)
	} else if val != 1 {
		t.Fatalf("expected 1 after Incr, got %d", val)
	}

	if val, err := client.IncrBy(ctx, counter, 4); err != nil {
		t.Fatalf("IncrBy failed: %v", err)
	} else if val != 5 {
		t.Fatalf("expected 5 after IncrBy, got %d", val)
	}

	if val, err := client.Decr(ctx, counter); err != nil {
		t.Fatalf("Decr failed: %v", err)
	} else if val != 4 {
		t.Fatalf("expected 4 after Decr, got %d", val)
	}

	if val, err := client.DecrBy(ctx, counter, 2); err != nil {
		t.Fatalf("DecrBy failed: %v", err)
	} else if val != 2 {
		t.Fatalf("expected 2 after DecrBy, got %d", val)
	}
}

func TestStringConditionalSet(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	key := testKey("setnx")
	ok, err := client.SetNX(ctx, key, "value", time.Second)
	if err != nil {
		t.Fatalf("SetNX failed: %v", err)
	}
	if !ok {
		t.Fatal("expected SetNX to succeed for new key")
	}

	ok, err = client.SetNX(ctx, key, "other", time.Second)
	if err != nil {
		t.Fatalf("SetNX second call failed: %v", err)
	}
	if ok {
		t.Fatal("expected SetNX to fail for existing key")
	}
}

func TestStringGetAs(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	type user struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	key := testKey("getas")
	value := user{Name: "Alice", Age: 29}

	if err := client.Set(ctx, key, `{"name":"Alice","age":29}`); err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	got, err := GetAs[user](client, ctx, key)
	if err != nil {
		t.Fatalf("GetAs failed: %v", err)
	}
	if got != value {
		t.Fatalf("unexpected struct value: %+v", got)
	}
}
