package redis

import (
	"context"
	"testing"
	"time"
)

func TestPublishSubscribe(t *testing.T) {
	client := newTestClient(t)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	channel := testKey("pubsub")
	message := "hello world"

	sub, err := client.Subscribe(ctx, channel)
	if err != nil {
		t.Fatalf("Subscribe failed: %v", err)
	}
	defer sub.Close()

	if _, err := sub.Receive(ctx); err != nil {
		t.Fatalf("Receive ack failed: %v", err)
	}

	ch := sub.Channel()

	if err := client.Publish(ctx, channel, message); err != nil {
		t.Fatalf("Publish failed: %v", err)
	}

	select {
	case msg := <-ch:
		if msg.Payload != message {
			t.Fatalf("expected payload %s, got %s", message, msg.Payload)
		}
	case <-ctx.Done():
		t.Fatal("timed out waiting for message")
	}

	if err := sub.Unsubscribe(ctx, channel); err != nil {
		t.Fatalf("Unsubscribe failed: %v", err)
	}
}

func TestPatternSubscribe(t *testing.T) {
	client := newTestClient(t)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pattern := testKey("pattern") + "*"
	channel := pattern + "1"

	sub, err := client.PSubscribe(ctx, pattern)
	if err != nil {
		t.Fatalf("PSubscribe failed: %v", err)
	}
	defer sub.Close()

	if _, err := sub.Receive(ctx); err != nil {
		t.Fatalf("Receive ack failed: %v", err)
	}

	ch := sub.Channel()

	if err := client.Publish(ctx, channel, "payload"); err != nil {
		t.Fatalf("Publish failed: %v", err)
	}

	select {
	case msg := <-ch:
		if msg.Channel != channel {
			t.Fatalf("expected channel %s, got %s", channel, msg.Channel)
		}
	case <-ctx.Done():
		t.Fatal("timed out waiting for pattern message")
	}

	if err := sub.PUnsubscribe(ctx, pattern); err != nil {
		t.Fatalf("PUnsubscribe failed: %v", err)
	}
}
