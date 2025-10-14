package mysql

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestClientCreation(t *testing.T) {
	client := newTestClient(t)
	if client == nil {
		t.Fatal("expected client instance")
	}
}

func TestSimpleCRUD(t *testing.T) {
	client := newTestClient(t)
	resetTable(t, "users")

	email := uniqueEmail("simple")

	data := map[string]interface{}{
		"name":  "Alice",
		"email": email,
		"age":   30,
		"city":  "Seoul",
	}

	result, err := client.Insert("users", data)
	if err != nil {
		t.Fatalf("Insert failed: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		t.Fatalf("LastInsertId failed: %v", err)
	}

	user, err := client.SelectOne("users", "id = ?", id)
	if err != nil {
		t.Fatalf("SelectOne failed: %v", err)
	}

	assertEqual(t, user["name"], "Alice")
	assertEqual(t, user["email"], email)
	assertEqual(t, toInt(user["age"]), 30)

	_, err = client.Update("users", map[string]interface{}{
		"age":  31,
		"city": "Busan",
	}, "id = ?", id)
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}

	updated, err := client.SelectOne("users", "id = ?", id)
	if err != nil {
		t.Fatalf("SelectOne after update failed: %v", err)
	}

	assertEqual(t, toInt(updated["age"]), 31)
	assertEqual(t, updated["city"], "Busan")

	if _, err := client.Delete("users", "id = ?", id); err != nil {
		t.Fatalf("Delete failed: %v", err)
	}

	_, err = client.SelectOne("users", "id = ?", id)
	if err == nil {
		t.Fatal("expected error when selecting deleted user")
	}
}

func TestBatchOperations(t *testing.T) {
	client := newTestClient(t)
	resetTable(t, "users")

	ctx := context.Background()
	users := []map[string]interface{}{
		{"name": "User 1", "email": uniqueEmail("batch"), "age": 20, "city": "Seoul"},
		{"name": "User 2", "email": uniqueEmail("batch"), "age": 25, "city": "Busan"},
		{"name": "User 3", "email": uniqueEmail("batch"), "age": 30, "city": "Incheon"},
	}

	if _, err := client.BatchInsert(ctx, "users", users); err != nil {
		t.Fatalf("BatchInsert failed: %v", err)
	}

	rows, err := client.SelectWhere(ctx, "users", "", WithOrderBy("id ASC"))
	if err != nil {
		t.Fatalf("SelectWhere failed: %v", err)
	}
	if len(rows) != len(users) {
		t.Fatalf("expected %d rows, got %d", len(users), len(rows))
	}

	var ids []interface{}
	for _, row := range rows {
		id := toInt64(row["id"])
		ids = append(ids, id)
	}

	updates := []BatchUpdateItem{
		{
			Data:             map[string]interface{}{"age": 21},
			ConditionAndArgs: []interface{}{"id = ?", ids[0]},
		},
		{
			Data:             map[string]interface{}{"city": "Jeju"},
			ConditionAndArgs: []interface{}{"id = ?", ids[1]},
		},
	}

	if err := client.BatchUpdate(ctx, "users", updates); err != nil {
		t.Fatalf("BatchUpdate failed: %v", err)
	}

	updatedRows, err := client.BatchSelectByIDs(ctx, "users", "id", ids)
	if err != nil {
		t.Fatalf("BatchSelectByIDs failed: %v", err)
	}

	for _, row := range updatedRows {
		switch row["id"].(int64) {
		case toInt64(ids[0]):
			assertEqual(t, toInt(row["age"]), 21)
		case toInt64(ids[1]):
			assertEqual(t, row["city"], "Jeju")
		}
	}

	delIDs := []interface{}{ids[2]}
	if _, err := client.BatchDelete(ctx, "users", "id", delIDs); err != nil {
		t.Fatalf("BatchDelete failed: %v", err)
	}

	remaining, err := client.SelectAll("users")
	if err != nil {
		t.Fatalf("SelectAll after delete failed: %v", err)
	}
	if len(remaining) != 2 {
		t.Fatalf("expected 2 rows after delete, got %d", len(remaining))
	}
}

func TestUpsertOperations(t *testing.T) {
	client := newTestClient(t)
	resetTable(t, "users")

	ctx := context.Background()
	email := uniqueEmail("upsert")

	_, err := client.Upsert(ctx, "users", map[string]interface{}{
		"email": email,
		"name":  "Initial",
		"age":   30,
	}, []string{"name", "age"})
	if err != nil {
		t.Fatalf("Upsert insert failed: %v", err)
	}

	_, err = client.Upsert(ctx, "users", map[string]interface{}{
		"email": email,
		"name":  "Updated",
		"age":   31,
	}, []string{"name", "age"})
	if err != nil {
		t.Fatalf("Upsert update failed: %v", err)
	}

	row, err := client.SelectOne("users", "email = ?", email)
	if err != nil {
		t.Fatalf("SelectOne after upsert failed: %v", err)
	}
	assertEqual(t, row["name"], "Updated")
	assertEqual(t, toInt(row["age"]), 31)

	batchData := []map[string]interface{}{
		{"email": uniqueEmail("upsert_batch"), "name": "Batch 1", "age": 22},
		{"email": uniqueEmail("upsert_batch"), "name": "Batch 2", "age": 24},
	}
	if _, err := client.UpsertBatch(ctx, "users", batchData, []string{"name", "age"}); err != nil {
		t.Fatalf("UpsertBatch insert failed: %v", err)
	}

	batchData[0]["name"] = "Batch 1 Updated"
	if _, err := client.UpsertBatch(ctx, "users", batchData, []string{"name"}); err != nil {
		t.Fatalf("UpsertBatch update failed: %v", err)
	}

	list, err := client.SelectWhere(ctx, "users", "email LIKE ?", "%upsert_batch%", WithOrderBy("email ASC"))
	if err != nil {
		t.Fatalf("SelectWhere batch failed: %v", err)
	}
	if len(list) != 2 {
		t.Fatalf("expected 2 rows, got %d", len(list))
	}
}

func TestPagination(t *testing.T) {
	client := newTestClient(t)
	resetTable(t, "users")

	ctx := context.Background()
	var seed []map[string]interface{}
	for i := 0; i < 25; i++ {
		seed = append(seed, map[string]interface{}{
			"name":  fmt.Sprintf("User %02d", i),
			"email": uniqueEmail("page"),
			"age":   20 + i%5,
			"city":  "Seoul",
		})
	}

	if _, err := client.BatchInsert(ctx, "users", seed); err != nil {
		t.Fatalf("BatchInsert seed failed: %v", err)
	}

	result, err := client.Paginate(ctx, "users", 2, 10, "", WithOrderBy("id ASC"))
	if err != nil {
		t.Fatalf("Paginate failed: %v", err)
	}

	if result.TotalRows != 25 {
		t.Fatalf("TotalRows = %d, want 25", result.TotalRows)
	}
	if result.TotalPages != 3 {
		t.Fatalf("TotalPages = %d, want 3", result.TotalPages)
	}
	if !result.HasNext || !result.HasPrev {
		t.Fatal("expected both HasNext and HasPrev to be true on page 2")
	}
	if len(result.Data) != 10 {
		t.Fatalf("expected 10 rows on page 2, got %d", len(result.Data))
	}
}

func TestSoftDeleteWorkflowIntegration(t *testing.T) {
	client := newTestClient(t)
	resetTable(t, "users")

	ctx := context.Background()
	email := uniqueEmail("soft")
	result, err := client.Insert("users", map[string]interface{}{
		"name":  "Soft User",
		"email": email,
		"age":   28,
		"city":  "Daegu",
	})
	if err != nil {
		t.Fatalf("Insert failed: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		t.Fatalf("LastInsertId failed: %v", err)
	}

	if _, err := client.SoftDelete(ctx, "users", "id = ?", id); err != nil {
		t.Fatalf("SoftDelete failed: %v", err)
	}

	trashed, err := client.SelectAllOnlyTrashed(ctx, "users", "id = ?", id)
	if err != nil {
		t.Fatalf("SelectAllOnlyTrashed failed: %v", err)
	}
	if len(trashed) != 1 {
		t.Fatalf("expected 1 trashed row, got %d", len(trashed))
	}

	if _, err := client.Restore(ctx, "users", "id = ?", id); err != nil {
		t.Fatalf("Restore failed: %v", err)
	}

	active, err := client.SelectAll("users", "id = ?", id)
	if err != nil {
		t.Fatalf("SelectAll after restore failed: %v", err)
	}
	if len(active) != 1 {
		t.Fatalf("expected 1 active row, got %d", len(active))
	}
}

func TestExportToCSV(t *testing.T) {
	client := newTestClient(t)
	resetTable(t, "users")

	ctx := context.Background()
	users := []map[string]interface{}{
		{"name": "Export 1", "email": uniqueEmail("export"), "age": 27, "city": "Seoul"},
		{"name": "Export 2", "email": uniqueEmail("export"), "age": 29, "city": "Busan"},
	}

	if _, err := client.BatchInsert(ctx, "users", users); err != nil {
		t.Fatalf("BatchInsert seed failed: %v", err)
	}

	tmpDir := t.TempDir()
	target := filepath.Join(tmpDir, "users.csv")

	opts := DefaultCSVExportOptions()
	opts.Columns = []string{"id", "name", "email", "age", "city"}
	opts.OrderBy = "id ASC"

	if err := client.ExportTableToCSV(ctx, "users", target, opts); err != nil {
		t.Fatalf("ExportTableToCSV failed: %v", err)
	}

	info, err := os.Stat(target)
	if err != nil {
		t.Fatalf("expected export file, got error: %v", err)
	}
	if info.Size() == 0 {
		t.Fatal("expected non-empty export file")
	}
}
