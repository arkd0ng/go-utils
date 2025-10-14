package mysql

import (
	"context"
	"testing"
)

// TestBatchInsert tests inserting multiple rows in a single query
// TestBatchInsert는 단일 쿼리로 여러 행을 삽입하는 것을 테스트합니다
func TestBatchInsert(t *testing.T) {
	tests := []struct {
		name      string
		table     string
		data      []map[string]interface{}
		wantErr   bool
		errString string
	}{
		{
			name:  "valid batch insert with multiple rows",
			table: "users",
			data: []map[string]interface{}{
				{"name": "John", "age": 30, "email": "john@example.com"},
				{"name": "Jane", "age": 25, "email": "jane@example.com"},
				{"name": "Bob", "age": 35, "email": "bob@example.com"},
			},
			wantErr: true, // Will fail without actual DB / 실제 DB 없이는 실패
		},
		{
			name:      "empty data",
			table:     "users",
			data:      []map[string]interface{}{},
			wantErr:   true,
			errString: "no data to insert",
		},
		{
			name:      "nil data",
			table:     "users",
			data:      nil,
			wantErr:   true,
			errString: "no data to insert",
		},
		{
			name:  "single row",
			table: "users",
			data: []map[string]interface{}{
				{"name": "John", "age": 30},
			},
			wantErr: true, // Will fail without actual DB / 실제 DB 없이는 실패
		},
		{
			name:  "empty columns in first row",
			table: "users",
			data: []map[string]interface{}{
				{},
			},
			wantErr:   true,
			errString: "no columns to insert",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// For validation errors, we can test without DB connection
			// 검증 에러의 경우 DB 연결 없이 테스트 가능
			if tt.errString != "" {
				client := &Client{
					config: &config{
						dsn: "invalid-dsn",
					},
				}

				ctx := context.Background()
				_, err := client.BatchInsert(ctx, tt.table, tt.data)

				if err == nil {
					t.Error("BatchInsert() should return error for invalid input")
					return
				}

				if err.Error() != tt.errString {
					t.Errorf("BatchInsert() error = %v, want error %v", err.Error(), tt.errString)
				}
			} else {
				// For DB operations, skip test without actual DB
				// DB 작업의 경우 실제 DB 없이 테스트 건너뛰기
				t.Skip("Requires actual database connection")
			}
		})
	}
}

// TestBatchUpdate tests updating multiple rows in a transaction
// TestBatchUpdate는 트랜잭션에서 여러 행을 업데이트하는 것을 테스트합니다
func TestBatchUpdate(t *testing.T) {
	tests := []struct {
		name      string
		table     string
		updates   []BatchUpdateItem
		wantErr   bool
		errString string
	}{
		{
			name:  "valid batch update",
			table: "users",
			updates: []BatchUpdateItem{
				{
					Data:             map[string]interface{}{"age": 31},
					ConditionAndArgs: []interface{}{"id = ?", 1},
				},
				{
					Data:             map[string]interface{}{"age": 26},
					ConditionAndArgs: []interface{}{"id = ?", 2},
				},
			},
			wantErr: true, // Will fail without actual DB / 실제 DB 없이는 실패
		},
		{
			name:      "empty updates",
			table:     "users",
			updates:   []BatchUpdateItem{},
			wantErr:   true,
			errString: "no updates to perform",
		},
		{
			name:      "nil updates",
			table:     "users",
			updates:   nil,
			wantErr:   true,
			errString: "no updates to perform",
		},
		{
			name:  "single update",
			table: "users",
			updates: []BatchUpdateItem{
				{
					Data:             map[string]interface{}{"name": "Updated"},
					ConditionAndArgs: []interface{}{"id = ?", 1},
				},
			},
			wantErr: true, // Will fail without actual DB / 실제 DB 없이는 실패
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
			err := client.BatchUpdate(ctx, tt.table, tt.updates)

			if (err != nil) != tt.wantErr {
				t.Errorf("BatchUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.errString != "" && err != nil && err.Error() != tt.errString {
				t.Errorf("BatchUpdate() error = %v, want error containing %v", err.Error(), tt.errString)
			}
		})
	}
}

// TestBatchDelete tests deleting multiple rows by IDs
// TestBatchDelete는 ID로 여러 행을 삭제하는 것을 테스트합니다
func TestBatchDelete(t *testing.T) {
	tests := []struct {
		name      string
		table     string
		idColumn  string
		ids       []interface{}
		wantErr   bool
		errString string
	}{
		{
			name:     "valid batch delete",
			table:    "users",
			idColumn: "id",
			ids:      []interface{}{1, 2, 3, 4, 5},
			wantErr:  true, // Will fail without actual DB / 실제 DB 없이는 실패
		},
		{
			name:      "empty ids",
			table:     "users",
			idColumn:  "id",
			ids:       []interface{}{},
			wantErr:   true,
			errString: "no IDs to delete",
		},
		{
			name:      "nil ids",
			table:     "users",
			idColumn:  "id",
			ids:       nil,
			wantErr:   true,
			errString: "no IDs to delete",
		},
		{
			name:     "single id",
			table:    "users",
			idColumn: "id",
			ids:      []interface{}{1},
			wantErr:  true, // Will fail without actual DB / 실제 DB 없이는 실패
		},
		{
			name:     "mixed type ids",
			table:    "users",
			idColumn: "id",
			ids:      []interface{}{1, "2", 3},
			wantErr:  true, // Will fail without actual DB / 실제 DB 없이는 실패
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
			_, err := client.BatchDelete(ctx, tt.table, tt.idColumn, tt.ids)

			if (err != nil) != tt.wantErr {
				t.Errorf("BatchDelete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.errString != "" && err != nil && err.Error() != tt.errString {
				t.Errorf("BatchDelete() error = %v, want error containing %v", err.Error(), tt.errString)
			}
		})
	}
}

// TestBatchSelectByIDs tests selecting multiple rows by IDs
// TestBatchSelectByIDs는 ID로 여러 행을 선택하는 것을 테스트합니다
func TestBatchSelectByIDs(t *testing.T) {
	tests := []struct {
		name      string
		table     string
		idColumn  string
		ids       []interface{}
		wantErr   bool
		wantEmpty bool
	}{
		{
			name:     "valid batch select",
			table:    "users",
			idColumn: "id",
			ids:      []interface{}{1, 2, 3},
			wantErr:  true, // Will fail without actual DB / 실제 DB 없이는 실패
		},
		{
			name:      "empty ids returns empty result",
			table:     "users",
			idColumn:  "id",
			ids:       []interface{}{},
			wantErr:   false,
			wantEmpty: true,
		},
		{
			name:      "nil ids returns empty result",
			table:     "users",
			idColumn:  "id",
			ids:       nil,
			wantErr:   false,
			wantEmpty: true,
		},
		{
			name:     "single id",
			table:    "users",
			idColumn: "id",
			ids:      []interface{}{1},
			wantErr:  true, // Will fail without actual DB / 실제 DB 없이는 실패
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
			results, err := client.BatchSelectByIDs(ctx, tt.table, tt.idColumn, tt.ids)

			if (err != nil) != tt.wantErr {
				t.Errorf("BatchSelectByIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantEmpty && len(results) != 0 {
				t.Errorf("BatchSelectByIDs() expected empty results, got %d rows", len(results))
			}
		})
	}
}

// BenchmarkBatchInsert benchmarks batch insert performance
// BenchmarkBatchInsert는 배치 삽입 성능을 벤치마크합니다
func BenchmarkBatchInsert(b *testing.B) {
	// This would require a real database connection
	// 실제 데이터베이스 연결이 필요합니다
	b.Skip("Requires actual database connection")
}

// BenchmarkBatchDelete benchmarks batch delete performance
// BenchmarkBatchDelete는 배치 삭제 성능을 벤치마크합니다
func BenchmarkBatchDelete(b *testing.B) {
	// This would require a real database connection
	// 실제 데이터베이스 연결이 필요합니다
	b.Skip("Requires actual database connection")
}

// BenchmarkBatchSelectByIDs benchmarks batch select performance
// BenchmarkBatchSelectByIDs는 배치 선택 성능을 벤치마크합니다
func BenchmarkBatchSelectByIDs(b *testing.B) {
	// This would require a real database connection
	// 실제 데이터베이스 연결이 필요합니다
	b.Skip("Requires actual database connection")
}
