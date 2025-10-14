package mysql

import "testing"

func TestPaginate(t *testing.T) {
	client := newTestClient(t)
	resetTable(t, "users")

	ctx := testContext()
	for i := 0; i < 35; i++ {
		if _, err := client.Insert("users", map[string]interface{}{
			"name":  "User",
			"email": uniqueEmail("paginate"),
			"age":   20 + i%5,
			"city":  "Seoul",
		}); err != nil {
			t.Fatalf("Insert seed failed: %v", err)
		}
	}

	result, err := client.Paginate(ctx, "users", 2, 10, "age >= ?", 21, WithOrderBy("id ASC"))
	if err != nil {
		t.Fatalf("Paginate failed: %v", err)
	}
	if result.Page != 2 || result.PageSize != 10 {
		t.Fatalf("unexpected pagination metadata: %+v", result)
	}
	if !result.HasNext || !result.HasPrev {
		t.Fatal("expected both HasNext and HasPrev to be true on page 2")
	}
	if len(result.Data) == 0 {
		t.Fatal("expected paginated data")
	}

	meta := result
	if meta.GetPage() != 2 {
		t.Fatalf("GetPage mismatch: %d", meta.GetPage())
	}
	if meta.GetTotalPages() == 0 {
		t.Fatal("expected total pages > 0")
	}
	if meta.NextPage() == 0 {
		t.Fatal("expected NextPage > 0")
	}
}

func TestPaginateQuery(t *testing.T) {
	client := newTestClient(t)
	resetTable(t, "users")

	ctx := testContext()
	for i := 0; i < 12; i++ {
		if _, err := client.Insert("users", map[string]interface{}{
			"name":  "QueryUser",
			"email": uniqueEmail("paginate_query"),
			"age":   25,
			"city":  "Busan",
		}); err != nil {
			t.Fatalf("Insert seed failed: %v", err)
		}
	}

	base := "SELECT id, name, email FROM users WHERE city = ? ORDER BY id"
	count := "SELECT COUNT(*) FROM users WHERE city = ?"

	result, err := client.PaginateQuery(ctx, base, count, 1, 5, "Busan")
	if err != nil {
		t.Fatalf("PaginateQuery failed: %v", err)
	}
	if len(result.Data) != 5 {
		t.Fatalf("expected 5 rows, got %d", len(result.Data))
	}
}

func TestPaginateValidations(t *testing.T) {
	client := newTestClient(t)
	ctx := testContext()

	if _, err := client.Paginate(ctx, "users", 0, 10); err == nil {
		t.Fatal("expected error for invalid page")
	}
	if _, err := client.Paginate(ctx, "users", 1, 0); err == nil {
		t.Fatal("expected error for invalid page size")
	}
}
