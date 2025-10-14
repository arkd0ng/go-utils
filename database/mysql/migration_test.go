package mysql

import "testing"

func TestMigrationHelpers(t *testing.T) {
	client := newTestClient(t)
	ctx := testContext()
	name := uniqueTableName("migration")

	schema := `
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL
	`
	if err := client.CreateTable(ctx, name, schema); err != nil {
		t.Fatalf("CreateTable failed: %v", err)
	}
	defer client.DropTable(ctx, name, true)

	if exists, err := client.TableExists(ctx, name); err != nil || !exists {
		t.Fatalf("TableExists failed: exists=%v err=%v", exists, err)
	}

	if err := client.AddColumn(ctx, name, "email", "VARCHAR(100) NULL"); err != nil {
		t.Fatalf("AddColumn failed: %v", err)
	}

	if err := client.DropColumn(ctx, name, "email"); err != nil {
		t.Fatalf("DropColumn failed: %v", err)
	}

	if err := client.TruncateTable(ctx, name); err != nil {
		t.Fatalf("TruncateTable failed: %v", err)
	}
}
