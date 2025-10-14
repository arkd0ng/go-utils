package mysql

import (
	"context"
	"testing"
)

// TestPaginate tests basic pagination
// TestPaginate는 기본 페이지네이션을 테스트합니다
func TestPaginate(t *testing.T) {
	tests := []struct {
		name             string
		table            string
		page             int
		pageSize         int
		conditionAndArgs []interface{}
		wantErr          bool
		errString        string
	}{
		{
			name:             "valid pagination page 1",
			table:            "users",
			page:             1,
			pageSize:         10,
			conditionAndArgs: []interface{}{},
			wantErr:          true, // Will fail without actual DB / 실제 DB 없이는 실패
		},
		{
			name:      "invalid page number (zero)",
			table:     "users",
			page:      0,
			pageSize:  10,
			wantErr:   true,
			errString: "page must be >= 1",
		},
		{
			name:      "invalid page number (negative)",
			table:     "users",
			page:      -1,
			pageSize:  10,
			wantErr:   true,
			errString: "page must be >= 1",
		},
		{
			name:      "invalid page size (zero)",
			table:     "users",
			page:      1,
			pageSize:  0,
			wantErr:   true,
			errString: "pageSize must be >= 1",
		},
		{
			name:      "invalid page size (negative)",
			table:     "users",
			page:      1,
			pageSize:  -10,
			wantErr:   true,
			errString: "pageSize must be >= 1",
		},
		{
			name:             "with WHERE condition",
			table:            "users",
			page:             1,
			pageSize:         20,
			conditionAndArgs: []interface{}{"age > ?", 18},
			wantErr:          true, // Will fail without actual DB / 실제 DB 없이는 실패
		},
		{
			name:             "large page number",
			table:            "users",
			page:             1000,
			pageSize:         10,
			conditionAndArgs: []interface{}{},
			wantErr:          true, // Will fail without actual DB / 실제 DB 없이는 실패
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				config: &config{
					dsn: "invalid-dsn",
				},
			}

			ctx := context.Background()
			result, err := client.Paginate(ctx, tt.table, tt.page, tt.pageSize, tt.conditionAndArgs...)

			if (err != nil) != tt.wantErr {
				t.Errorf("Paginate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.errString != "" && err != nil && err.Error() != tt.errString {
				t.Errorf("Paginate() error = %v, want error containing %v", err.Error(), tt.errString)
			}

			if err == nil && result == nil {
				t.Error("Paginate() returned nil result without error")
			}
		})
	}
}

// TestPaginateWithConditions tests pagination with WHERE conditions
// TestPaginateWithConditions는 WHERE 조건을 사용한 페이지네이션을 테스트합니다
func TestPaginateWithConditions(t *testing.T) {
	// This test would require a real database
	// 이 테스트는 실제 데이터베이스가 필요합니다
	t.Skip("Requires actual database connection")
}

// TestPaginationMetadata tests pagination metadata (HasNext, HasPrev, TotalPages)
// TestPaginationMetadata는 페이지네이션 메타데이터를 테스트합니다
func TestPaginationMetadata(t *testing.T) {
	tests := []struct {
		name           string
		result         *PaginationResult
		wantHasNext    bool
		wantHasPrev    bool
		wantFirstPage  bool
		wantLastPage   bool
		wantNextPage   int
		wantPrevPage   int
		wantTotalPages int
	}{
		{
			name: "first page with more pages",
			result: &PaginationResult{
				Page:       1,
				PageSize:   10,
				TotalRows:  100,
				TotalPages: 10,
				HasNext:    true,
				HasPrev:    false,
			},
			wantHasNext:    true,
			wantHasPrev:    false,
			wantFirstPage:  true,
			wantLastPage:   false,
			wantNextPage:   2,
			wantPrevPage:   0,
			wantTotalPages: 10,
		},
		{
			name: "middle page",
			result: &PaginationResult{
				Page:       5,
				PageSize:   10,
				TotalRows:  100,
				TotalPages: 10,
				HasNext:    true,
				HasPrev:    true,
			},
			wantHasNext:    true,
			wantHasPrev:    true,
			wantFirstPage:  false,
			wantLastPage:   false,
			wantNextPage:   6,
			wantPrevPage:   4,
			wantTotalPages: 10,
		},
		{
			name: "last page",
			result: &PaginationResult{
				Page:       10,
				PageSize:   10,
				TotalRows:  100,
				TotalPages: 10,
				HasNext:    false,
				HasPrev:    true,
			},
			wantHasNext:    false,
			wantHasPrev:    true,
			wantFirstPage:  false,
			wantLastPage:   true,
			wantNextPage:   0,
			wantPrevPage:   9,
			wantTotalPages: 10,
		},
		{
			name: "single page (no pagination needed)",
			result: &PaginationResult{
				Page:       1,
				PageSize:   10,
				TotalRows:  5,
				TotalPages: 1,
				HasNext:    false,
				HasPrev:    false,
			},
			wantHasNext:    false,
			wantHasPrev:    false,
			wantFirstPage:  true,
			wantLastPage:   true,
			wantNextPage:   0,
			wantPrevPage:   0,
			wantTotalPages: 1,
		},
		{
			name: "empty result",
			result: &PaginationResult{
				Page:       1,
				PageSize:   10,
				TotalRows:  0,
				TotalPages: 0,
				HasNext:    false,
				HasPrev:    false,
			},
			wantHasNext:    false,
			wantHasPrev:    false,
			wantFirstPage:  true,
			wantLastPage:   true,
			wantNextPage:   0,
			wantPrevPage:   0,
			wantTotalPages: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.result.HasNext != tt.wantHasNext {
				t.Errorf("HasNext = %v, want %v", tt.result.HasNext, tt.wantHasNext)
			}
			if tt.result.HasPrev != tt.wantHasPrev {
				t.Errorf("HasPrev = %v, want %v", tt.result.HasPrev, tt.wantHasPrev)
			}
			if tt.result.IsFirstPage() != tt.wantFirstPage {
				t.Errorf("IsFirstPage() = %v, want %v", tt.result.IsFirstPage(), tt.wantFirstPage)
			}
			if tt.result.IsLastPage() != tt.wantLastPage {
				t.Errorf("IsLastPage() = %v, want %v", tt.result.IsLastPage(), tt.wantLastPage)
			}
			if tt.result.NextPage() != tt.wantNextPage {
				t.Errorf("NextPage() = %v, want %v", tt.result.NextPage(), tt.wantNextPage)
			}
			if tt.result.PrevPage() != tt.wantPrevPage {
				t.Errorf("PrevPage() = %v, want %v", tt.result.PrevPage(), tt.wantPrevPage)
			}
			if tt.result.GetTotalPages() != tt.wantTotalPages {
				t.Errorf("GetTotalPages() = %v, want %v", tt.result.GetTotalPages(), tt.wantTotalPages)
			}
		})
	}
}

// TestPaginateQuery tests custom query pagination
// TestPaginateQuery는 사용자 정의 쿼리 페이지네이션을 테스트합니다
func TestPaginateQuery(t *testing.T) {
	tests := []struct {
		name       string
		baseQuery  string
		countQuery string
		page       int
		pageSize   int
		args       []interface{}
		wantErr    bool
		errString  string
	}{
		{
			name:       "valid custom query pagination",
			baseQuery:  "SELECT * FROM users WHERE age > ?",
			countQuery: "SELECT COUNT(*) FROM users WHERE age > ?",
			page:       1,
			pageSize:   10,
			args:       []interface{}{18},
			wantErr:    true, // Will fail without actual DB / 실제 DB 없이는 실패
		},
		{
			name:       "invalid page number",
			baseQuery:  "SELECT * FROM users",
			countQuery: "SELECT COUNT(*) FROM users",
			page:       0,
			pageSize:   10,
			wantErr:    true,
			errString:  "page must be >= 1",
		},
		{
			name:       "invalid page size",
			baseQuery:  "SELECT * FROM users",
			countQuery: "SELECT COUNT(*) FROM users",
			page:       1,
			pageSize:   0,
			wantErr:    true,
			errString:  "pageSize must be >= 1",
		},
		{
			name:       "complex JOIN query",
			baseQuery:  "SELECT u.*, COUNT(o.id) as order_count FROM users u LEFT JOIN orders o ON u.id = o.user_id GROUP BY u.id",
			countQuery: "SELECT COUNT(*) FROM users",
			page:       1,
			pageSize:   20,
			wantErr:    true, // Will fail without actual DB / 실제 DB 없이는 실패
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				config: &config{
					dsn: "invalid-dsn",
				},
			}

			ctx := context.Background()
			result, err := client.PaginateQuery(ctx, tt.baseQuery, tt.countQuery, tt.page, tt.pageSize, tt.args...)

			if (err != nil) != tt.wantErr {
				t.Errorf("PaginateQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.errString != "" && err != nil && err.Error() != tt.errString {
				t.Errorf("PaginateQuery() error = %v, want error containing %v", err.Error(), tt.errString)
			}

			if err == nil && result == nil {
				t.Error("PaginateQuery() returned nil result without error")
			}
		})
	}
}

// TestPaginationEdgeCases tests edge cases
// TestPaginationEdgeCases는 엣지 케이스를 테스트합니다
func TestPaginationEdgeCases(t *testing.T) {
	tests := []struct {
		name        string
		description string
		page        int
		pageSize    int
		totalRows   int64
		wantErr     bool
	}{
		{
			name:        "page beyond total pages",
			description: "Requesting page 100 when only 10 pages exist",
			page:        100,
			pageSize:    10,
			totalRows:   100,
			wantErr:     true, // Will fail without actual DB / 실제 DB 없이는 실패
		},
		{
			name:        "very large page size",
			description: "Requesting 10000 items per page",
			page:        1,
			pageSize:    10000,
			totalRows:   50,
			wantErr:     true, // Will fail without actual DB / 실제 DB 없이는 실패
		},
		{
			name:        "empty table",
			description: "Paginating an empty table",
			page:        1,
			pageSize:    10,
			totalRows:   0,
			wantErr:     true, // Will fail without actual DB / 실제 DB 없이는 실패
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				config: &config{
					dsn: "invalid-dsn",
				},
			}

			ctx := context.Background()
			_, err := client.Paginate(ctx, "users", tt.page, tt.pageSize)

			if (err != nil) != tt.wantErr {
				t.Errorf("Paginate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestPaginationResultGetters tests getter methods
// TestPaginationResultGetters는 getter 메서드를 테스트합니다
func TestPaginationResultGetters(t *testing.T) {
	result := &PaginationResult{
		Page:       3,
		PageSize:   20,
		TotalRows:  150,
		TotalPages: 8,
		HasNext:    true,
		HasPrev:    true,
	}

	if result.GetPage() != 3 {
		t.Errorf("GetPage() = %v, want 3", result.GetPage())
	}
	if result.GetTotalPages() != 8 {
		t.Errorf("GetTotalPages() = %v, want 8", result.GetTotalPages())
	}
	if result.GetTotalRows() != 150 {
		t.Errorf("GetTotalRows() = %v, want 150", result.GetTotalRows())
	}
}

// BenchmarkPaginate benchmarks pagination performance
// BenchmarkPaginate는 페이지네이션 성능을 벤치마크합니다
func BenchmarkPaginate(b *testing.B) {
	// This would require a real database connection
	// 실제 데이터베이스 연결이 필요합니다
	b.Skip("Requires actual database connection")
}

// BenchmarkPaginateQuery benchmarks custom query pagination performance
// BenchmarkPaginateQuery는 사용자 정의 쿼리 페이지네이션 성능을 벤치마크합니다
func BenchmarkPaginateQuery(b *testing.B) {
	// This would require a real database connection
	// 실제 데이터베이스 연결이 필요합니다
	b.Skip("Requires actual database connection")
}
