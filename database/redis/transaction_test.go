package redis

import (
	"context"
	"strconv"
	"testing"
)

func TestTransactionUpdatesValue(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	key := testKey("txn")
	if err := client.Set(ctx, key, "10"); err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	err := client.Transaction(ctx, func(tx *Tx) error {
		current, err := tx.Get(ctx, key)
		if err != nil {
			return err
		}

		val, err := strconv.Atoi(current)
		if err != nil {
			return err
		}

		return tx.Exec(ctx, func(pipe Pipeliner) error {
			if err := tx.Set(ctx, pipe, key, strconv.Itoa(val+5)); err != nil {
				return err
			}
			return nil
		})
	}, key)
	if err != nil {
		t.Fatalf("Transaction failed: %v", err)
	}

	final, err := client.Get(ctx, key)
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if final != "15" {
		t.Fatalf("expected value 15, got %s", final)
	}
}

func TestTransactionDeletesKeys(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	key1 := testKey("txn:del1")
	key2 := testKey("txn:del2")

	if err := client.Set(ctx, key1, "value1"); err != nil {
		t.Fatalf("Set failed: %v", err)
	}
	if err := client.Set(ctx, key2, "value2"); err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	err := client.Transaction(ctx, func(tx *Tx) error {
		return tx.Exec(ctx, func(pipe Pipeliner) error {
			if err := tx.Del(ctx, pipe, key1, key2); err != nil {
				return err
			}
			return nil
		})
	}, key1, key2)
	if err != nil {
		t.Fatalf("Transaction delete failed: %v", err)
	}

	exists, err := client.Exists(ctx, key1, key2)
	if err != nil {
		t.Fatalf("Exists failed: %v", err)
	}
	if exists != 0 {
		t.Fatalf("expected keys to be deleted, got exists=%d", exists)
	}
}
