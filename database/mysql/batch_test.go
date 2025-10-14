package mysql

import "testing"

func TestBatchInsertUpdateDelete(t *testing.T) {
	client := newTestClient(t)
	resetTable(t, "users")

	ctx := testContext()
	data := []map[string]interface{}{
		{"name": "Batch User 1", "email": uniqueEmail("batch"), "age": 21, "city": "Seoul"},
		{"name": "Batch User 2", "email": uniqueEmail("batch"), "age": 22, "city": "Busan"},
	}

	if _, err := client.BatchInsert(ctx, "users", data); err != nil {
		t.Fatalf("BatchInsert failed: %v", err)
	}

	rows, err := client.SelectWhere(ctx, "users", "email LIKE ?", "%batch%")
	if err != nil {
		t.Fatalf("SelectWhere failed: %v", err)
	}
	if len(rows) != len(data) {
		t.Fatalf("expected %d rows, got %d", len(data), len(rows))
	}

	var ids []interface{}
	for _, row := range rows {
		ids = append(ids, toInt64(row["id"]))
	}

	updates := []BatchUpdateItem{
		{Data: map[string]interface{}{"age": 30}, ConditionAndArgs: []interface{}{"id = ?", ids[0]}},
		{Data: map[string]interface{}{"city": "Daegu"}, ConditionAndArgs: []interface{}{"id = ?", ids[1]}},
	}

	if err := client.BatchUpdate(ctx, "users", updates); err != nil {
		t.Fatalf("BatchUpdate failed: %v", err)
	}

	updated, err := client.BatchSelectByIDs(ctx, "users", "id", ids)
	if err != nil {
		t.Fatalf("BatchSelectByIDs failed: %v", err)
	}
	if len(updated) != len(ids) {
		t.Fatalf("expected %d rows, got %d", len(ids), len(updated))
	}

	for _, row := range updated {
		if row["id"].(int64) == ids[0].(int64) && toInt(row["age"]) != 30 {
			t.Fatal("expected age update to 30")
		}
		if row["id"].(int64) == ids[1].(int64) && row["city"] != "Daegu" {
			t.Fatal("expected city update to Daegu")
		}
	}

	if _, err := client.BatchDelete(ctx, "users", "id", ids); err != nil {
		t.Fatalf("BatchDelete failed: %v", err)
	}

	remaining, err := client.SelectWhere(ctx, "users", "email LIKE ?", "%batch%")
	if err != nil {
		t.Fatalf("SelectWhere after delete failed: %v", err)
	}
	if len(remaining) != 0 {
		t.Fatalf("expected no rows, got %d", len(remaining))
	}
}

func TestBatchInsertValidations(t *testing.T) {
	client := newTestClient(t)

	ctx := testContext()
	if _, err := client.BatchInsert(ctx, "users", nil); err == nil {
		t.Fatal("expected error for nil data")
	}

	if _, err := client.BatchInsert(ctx, "users", []map[string]interface{}{}); err == nil {
		t.Fatal("expected error for empty data slice")
	}

	bad := []map[string]interface{}{{}}
	if _, err := client.BatchInsert(ctx, "users", bad); err == nil {
		t.Fatal("expected error for empty column set")
	}
}

func TestBatchDeleteValidations(t *testing.T) {
	client := newTestClient(t)
	ctx := testContext()

	if _, err := client.BatchDelete(ctx, "users", "id", nil); err == nil {
		t.Fatal("expected error for nil ids")
	}

	if _, err := client.BatchDelete(ctx, "users", "id", []interface{}{}); err == nil {
		t.Fatal("expected error for empty ids slice")
	}
}
