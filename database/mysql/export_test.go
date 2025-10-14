package mysql

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"testing"
)

func TestExportAndImportCSV(t *testing.T) {
	client := newTestClient(t)
	resetTable(t, "users")

	ctx := testContext()
	seed := []map[string]interface{}{
		{"name": "Export 1", "email": uniqueEmail("export"), "age": 25, "city": "Seoul"},
		{"name": "Export 2", "email": uniqueEmail("export"), "age": 26, "city": "Busan"},
	}
	if _, err := client.BatchInsert(ctx, "users", seed); err != nil {
		t.Fatalf("BatchInsert failed: %v", err)
	}

	file := filepath.Join(t.TempDir(), "users.csv")
	opts := DefaultCSVExportOptions()
	opts.Columns = []string{"id", "name", "email", "age", "city"}
	if err := client.ExportTableToCSV(ctx, "users", file, opts); err != nil {
		t.Fatalf("ExportTableToCSV failed: %v", err)
	}

	info, err := os.Stat(file)
	if err != nil || info.Size() == 0 {
		t.Fatalf("expected CSV file to be created, err=%v size=%d", err, info.Size())
	}

	importFile := filepath.Join(t.TempDir(), "import.csv")
	writeCSV(t, importFile, [][]string{
		{"name", "email", "age", "city"},
		{"Imported", uniqueEmail("import"), "28", "Daegu"},
	})

	resetTable(t, "users")
	importOpts := DefaultCSVImportOptions()
	importOpts.Columns = []string{"name", "email", "age", "city"}
	importOpts.BatchSize = 1

	if err := client.ImportFromCSV(ctx, "users", importFile, importOpts); err != nil {
		t.Fatalf("ImportFromCSV failed: %v", err)
	}

	rows, err := client.SelectAll("users")
	if err != nil {
		t.Fatalf("SelectAll after import failed: %v", err)
	}
	if len(rows) != 1 {
		t.Fatalf("expected 1 imported row, got %d", len(rows))
	}
}

func TestExportQueryToCSV(t *testing.T) {
	client := newTestClient(t)
	resetTable(t, "users")

	ctx := testContext()
	if _, err := client.Insert("users", map[string]interface{}{
		"name":  "QueryExport",
		"email": uniqueEmail("export_query"),
		"age":   33,
		"city":  "Seoul",
	}); err != nil {
		t.Fatalf("Insert failed: %v", err)
	}

	file := filepath.Join(t.TempDir(), "query.csv")
	query := "SELECT name, email FROM users WHERE city = ?"
	opts := DefaultCSVExportOptions()
	if err := client.ExportQueryToCSV(ctx, query, []interface{}{"Seoul"}, file, opts); err != nil {
		t.Fatalf("ExportQueryToCSV failed: %v", err)
	}

	info, err := os.Stat(file)
	if err != nil || info.Size() == 0 {
		t.Fatalf("expected query CSV file, err=%v size=%d", err, info.Size())
	}
}

func TestImportValidations(t *testing.T) {
	client := newTestClient(t)
	ctx := testContext()

	if err := client.ImportFromCSV(ctx, "users", "missing.csv", DefaultCSVImportOptions()); err == nil {
		t.Fatal("expected error for missing file")
	}
}

func writeCSV(t *testing.T, path string, rows [][]string) {
	file, err := os.Create(path)
	if err != nil {
		t.Fatalf("failed to create csv: %v", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	if err := w.WriteAll(rows); err != nil {
		t.Fatalf("failed to write csv: %v", err)
	}
}
