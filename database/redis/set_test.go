package redis

import (
	"context"
	"reflect"
	"sort"
	"testing"
)

func TestSetOperations(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	setA := testKey("set:a")
	setB := testKey("set:b")

	if err := client.SAdd(ctx, setA, "red", "green", "blue"); err != nil {
		t.Fatalf("SAdd setA failed: %v", err)
	}
	if err := client.SAdd(ctx, setB, "green", "yellow"); err != nil {
		t.Fatalf("SAdd setB failed: %v", err)
	}

	card, err := client.SCard(ctx, setA)
	if err != nil {
		t.Fatalf("SCard failed: %v", err)
	}
	if card != 3 {
		t.Fatalf("expected cardinality 3, got %d", card)
	}

	members, err := client.SMembers(ctx, setA)
	if err != nil {
		t.Fatalf("SMembers failed: %v", err)
	}
	sort.Strings(members)
	if !reflect.DeepEqual(members, []string{"blue", "green", "red"}) {
		t.Fatalf("unexpected set members: %#v", members)
	}

	if exists, err := client.SIsMember(ctx, setA, "green"); err != nil {
		t.Fatalf("SIsMember failed: %v", err)
	} else if !exists {
		t.Fatal("expected green to be a member")
	}

	union, err := client.SUnion(ctx, setA, setB)
	if err != nil {
		t.Fatalf("SUnion failed: %v", err)
	}
	sort.Strings(union)
	if !reflect.DeepEqual(union, []string{"blue", "green", "red", "yellow"}) {
		t.Fatalf("unexpected union: %#v", union)
	}

	inter, err := client.SInter(ctx, setA, setB)
	if err != nil {
		t.Fatalf("SInter failed: %v", err)
	}
	if !reflect.DeepEqual(inter, []string{"green"}) {
		t.Fatalf("unexpected intersection: %#v", inter)
	}

	diff, err := client.SDiff(ctx, setA, setB)
	if err != nil {
		t.Fatalf("SDiff failed: %v", err)
	}
	sort.Strings(diff)
	if !reflect.DeepEqual(diff, []string{"blue", "red"}) {
		t.Fatalf("unexpected diff: %#v", diff)
	}

	pop, err := client.SPop(ctx, setA)
	if err != nil {
		t.Fatalf("SPop failed: %v", err)
	}
	if pop == "" {
		t.Fatal("expected SPop to return a member")
	}

	random, err := client.SRandMember(ctx, setB, 2)
	if err != nil {
		t.Fatalf("SRandMember failed: %v", err)
	}
	if len(random) == 0 {
		t.Fatal("expected SRandMember to return at least one member")
	}

	if err := client.SRem(ctx, setB, "green"); err != nil {
		t.Fatalf("SRem failed: %v", err)
	}

	if exists, err := client.SIsMember(ctx, setB, "green"); err != nil {
		t.Fatalf("SIsMember after remove failed: %v", err)
	} else if exists {
		t.Fatal("expected green to be removed from setB")
	}
}
