package mysql

import "testing"

func TestUpsertAndUpsertBatch(t *testing.T) {
	client := newTestClient(t)
	resetTable(t, "users")

	ctx := testContext()
	email := uniqueEmail("upsert_single")

	if _, err := client.Upsert(ctx, "users", map[string]interface{}{
		"email": email,
		"name":  "Upsert One",
		"age":   29,
	}, []string{"name", "age"}); err != nil {
		t.Fatalf("Upsert insert failed: %v", err)
	}

	if _, err := client.Upsert(ctx, "users", map[string]interface{}{
		"email": email,
		"name":  "Upsert Updated",
		"age":   31,
	}, []string{"name", "age"}); err != nil {
		t.Fatalf("Upsert update failed: %v", err)
	}

	row, err := client.SelectOne("users", "email = ?", email)
	if err != nil {
		t.Fatalf("SelectOne failed: %v", err)
	}
	assertEqual(t, row["name"], "Upsert Updated")
	assertEqual(t, toInt(row["age"]), 31)

	batch := []map[string]interface{}{
		{"email": uniqueEmail("upsert_batch"), "name": "Batch 1", "age": 20},
		{"email": uniqueEmail("upsert_batch"), "name": "Batch 2", "age": 21},
	}

	if _, err := client.UpsertBatch(ctx, "users", batch, []string{"name", "age"}); err != nil {
		t.Fatalf("UpsertBatch insert failed: %v", err)
	}

	batch[0]["name"] = "Batch 1 Updated"
	if _, err := client.UpsertBatch(ctx, "users", batch, []string{"name"}); err != nil {
		t.Fatalf("UpsertBatch update failed: %v", err)
	}

	rows, err := client.SelectWhere(ctx, "users", "email LIKE ?", "%upsert_batch%", WithOrderBy("email ASC"))
	if err != nil {
		t.Fatalf("SelectWhere failed: %v", err)
	}
	if len(rows) != 2 {
		t.Fatalf("expected 2 rows, got %d", len(rows))
	}

	if rows[0]["name"] != "Batch 1 Updated" {
		t.Fatalf("expected updated batch name, got %v", rows[0]["name"])
	}
}

func TestUpsertValidations(t *testing.T) {
	client := newTestClient(t)
	ctx := testContext()

	if _, err := client.Upsert(ctx, "users", nil, []string{"name"}); err == nil {
		t.Fatal("expected error for nil data")
	}

	if _, err := client.UpsertBatch(ctx, "users", nil, []string{"name"}); err == nil {
		t.Fatal("expected error for nil batch data")
	}
}
