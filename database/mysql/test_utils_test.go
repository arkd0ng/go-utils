package mysql

import (
	"context"
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

var uniqueCounter uint64

func nextUnique() uint64 {
	return atomic.AddUint64(&uniqueCounter, 1)
}

func uniqueEmail(prefix string) string {
	return fmt.Sprintf("%s_%d_%d@test.example.com", prefix, time.Now().UnixNano(), nextUnique())
}

func uniqueTableName(prefix string) string {
	return fmt.Sprintf("%s_%d_%d", prefix, time.Now().UnixNano(), nextUnique())
}

func assertEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v (%T), want %v (%T)", got, got, want, want)
	}
}

func toInt(value interface{}) int {
	switch v := value.(type) {
	case int:
		return v
	case int32:
		return int(v)
	case int64:
		return int(v)
	default:
		panic(fmt.Sprintf("unexpected int type %T", value))
	}
}

func toInt64(value interface{}) int64 {
	switch v := value.(type) {
	case int:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	default:
		panic(fmt.Sprintf("unexpected int64 type %T", value))
	}
}

func testContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx
}
