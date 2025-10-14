package redis

import (
	"context"
	"testing"
)

func TestPipelineExecutesCommands(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	key1 := testKey("pipeline")
	key2 := testKey("pipeline")

	err := client.Pipeline(ctx, func(pipe Pipeliner) error {
		if err := pipe.Set(ctx, key1, "value1", 0).Err(); err != nil {
			return err
		}
		if err := pipe.Set(ctx, key2, "value2", 0).Err(); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		t.Fatalf("Pipeline failed: %v", err)
	}

	val1, err := client.Get(ctx, key1)
	if err != nil || val1 != "value1" {
		t.Fatalf("unexpected pipeline result: %v, %s", err, val1)
	}
	val2, err := client.Get(ctx, key2)
	if err != nil || val2 != "value2" {
		t.Fatalf("unexpected pipeline result: %v, %s", err, val2)
	}
}

func TestTxPipelineExecutesAtomically(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	counter := testKey("txpipeline")
	if err := client.Set(ctx, counter, "0"); err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	err := client.TxPipeline(ctx, func(pipe Pipeliner) error {
		if err := pipe.Incr(ctx, counter).Err(); err != nil {
			return err
		}
		if err := pipe.IncrBy(ctx, counter, 4).Err(); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		t.Fatalf("TxPipeline failed: %v", err)
	}

	val, err := client.Get(ctx, counter)
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if val != "5" {
		t.Fatalf("expected counter to be 5, got %s", val)
	}
}
