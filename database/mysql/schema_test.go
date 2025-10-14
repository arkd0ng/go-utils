package mysql

import "testing"

func TestSchemaInspection(t *testing.T) {
	client := newTestClient(t)
	resetTable(t, "users")

	ctx := testContext()
	if _, err := client.Insert("users", map[string]interface{}{
		"name":  "Schema User",
		"email": uniqueEmail("schema"),
		"age":   29,
		"city":  "Seoul",
	}); err != nil {
		t.Fatalf("Insert failed: %v", err)
	}

	tables, err := client.GetTables(ctx)
	if err != nil {
		t.Fatalf("GetTables failed: %v", err)
	}
	if len(tables) == 0 {
		t.Fatal("expected at least one table")
	}

	columns, err := client.GetColumns(ctx, "users")
	if err != nil {
		t.Fatalf("GetColumns failed: %v", err)
	}
	if len(columns) == 0 {
		t.Fatal("expected columns metadata")
	}

	indexes, err := client.GetIndexes(ctx, "users")
	if err != nil {
		t.Fatalf("GetIndexes failed: %v", err)
	}
	if len(indexes) == 0 {
		t.Fatal("expected indexes metadata")
	}
}
