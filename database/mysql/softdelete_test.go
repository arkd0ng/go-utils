package mysql

import "testing"

func TestSoftDeleteWorkflow(t *testing.T) {
	client := newTestClient(t)
	resetTable(t, "users")

	ctx := testContext()
	result, err := client.Insert("users", map[string]interface{}{
		"name":  "SoftDelete User",
		"email": uniqueEmail("soft"),
		"age":   27,
		"city":  "Seoul",
	})
	if err != nil {
		t.Fatalf("Insert failed: %v", err)
	}

	id, _ := result.LastInsertId()

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

	count, err := client.CountOnlyTrashed(ctx, "users", "id = ?", id)
	if err != nil {
		t.Fatalf("CountOnlyTrashed failed: %v", err)
	}
	if count != 1 {
		t.Fatalf("expected trashed count 1, got %d", count)
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

func TestSoftDeleteValidations(t *testing.T) {
	client := newTestClient(t)
	ctx := testContext()

	if _, err := client.SoftDelete(ctx, "users"); err == nil {
		t.Fatal("expected error when condition missing for soft delete")
	}
	if _, err := client.Restore(ctx, "users"); err == nil {
		t.Fatal("expected error when condition missing for restore")
	}
}
