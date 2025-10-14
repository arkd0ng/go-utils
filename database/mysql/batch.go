package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

// BatchInsert inserts multiple rows in a single query for better performance
// BatchInsert는 성능 향상을 위해 단일 쿼리로 여러 행을 삽입합니다
//
// Example / 예제:
//
//	data := []map[string]interface{}{
//	    {"name": "John", "age": 30, "email": "john@example.com"},
//	    {"name": "Jane", "age": 25, "email": "jane@example.com"},
//	    {"name": "Bob", "age": 35, "email": "bob@example.com"},
//	}
//	result, err := db.BatchInsert(ctx, "users", data)
//
// This generates: INSERT INTO users (name, age, email) VALUES (?,?,?),(?,?,?),(?,?,?)
func (c *Client) BatchInsert(ctx context.Context, table string, data []map[string]interface{}) (sql.Result, error) {
	c.mu.RLock()
	if c.closed {
		c.mu.RUnlock()
		return nil, ErrClosed
	}
	c.mu.RUnlock()

	if len(data) == 0 {
		return nil, fmt.Errorf("no data to insert")
	}

	// Get column names from first row / 첫 번째 행에서 컬럼 이름 가져오기
	var columns []string
	for col := range data[0] {
		columns = append(columns, col)
	}

	if len(columns) == 0 {
		return nil, fmt.Errorf("no columns to insert")
	}

	// Build column list / 컬럼 목록 빌드
	columnList := strings.Join(columns, ", ")

	// Build value placeholders / 값 플레이스홀더 빌드
	valuePlaceholder := "(" + strings.Repeat("?,", len(columns))
	valuePlaceholder = valuePlaceholder[:len(valuePlaceholder)-1] + ")"

	// Build multiple value placeholders / 여러 값 플레이스홀더 빌드
	var valuePlaceholders []string
	for i := 0; i < len(data); i++ {
		valuePlaceholders = append(valuePlaceholders, valuePlaceholder)
	}

	// Build query / 쿼리 빌드
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s",
		table, columnList, strings.Join(valuePlaceholders, ","))

	// Collect all values / 모든 값 수집
	var args []interface{}
	for _, row := range data {
		for _, col := range columns {
			args = append(args, row[col])
		}
	}

	start := time.Now()

	// Execute with retry / 재시도로 실행
	var result sql.Result
	err := c.executeWithRetry(ctx, func() error {
		db := c.getCurrentConnection()
		var execErr error
		result, execErr = db.ExecContext(ctx, query, args...)
		return execErr
	})

	duration := time.Since(start)

	if err != nil {
		c.logQuery(query, args, duration, err)
		return nil, c.wrapError("BatchInsert", query, args, err, duration)
	}

	c.logQuery(query, args, duration, nil)
	return result, nil
}

// BatchUpdateItem represents a single update operation in a batch
// BatchUpdateItem은 배치의 단일 업데이트 작업을 나타냅니다
type BatchUpdateItem struct {
	Data              map[string]interface{} // Columns to update / 업데이트할 컬럼
	ConditionAndArgs  []interface{}          // WHERE condition and arguments / WHERE 조건 및 인자
}

// BatchUpdate performs multiple update operations in a transaction
// BatchUpdate는 트랜잭션에서 여러 업데이트 작업을 수행합니다
//
// Example / 예제:
//
//	updates := []mysql.BatchUpdateItem{
//	    {
//	        Data: map[string]interface{}{"age": 31},
//	        ConditionAndArgs: []interface{}{"id = ?", 1},
//	    },
//	    {
//	        Data: map[string]interface{}{"age": 26},
//	        ConditionAndArgs: []interface{}{"id = ?", 2},
//	    },
//	}
//	err := db.BatchUpdate(ctx, "users", updates)
func (c *Client) BatchUpdate(ctx context.Context, table string, updates []BatchUpdateItem) error {
	if len(updates) == 0 {
		return fmt.Errorf("no updates to perform")
	}

	// Use transaction for atomicity / 원자성을 위해 트랜잭션 사용
	return c.Transaction(ctx, func(tx *Tx) error {
		for _, item := range updates {
			_, err := tx.Update(table, item.Data, item.ConditionAndArgs...)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// BatchDelete deletes multiple rows by IDs in a single query
// BatchDelete는 단일 쿼리로 ID로 여러 행을 삭제합니다
//
// Example / 예제:
//
//	ids := []interface{}{1, 2, 3, 4, 5}
//	result, err := db.BatchDelete(ctx, "users", "id", ids)
//
// This generates: DELETE FROM users WHERE id IN (?,?,?,?,?)
func (c *Client) BatchDelete(ctx context.Context, table string, idColumn string, ids []interface{}) (sql.Result, error) {
	c.mu.RLock()
	if c.closed {
		c.mu.RUnlock()
		return nil, ErrClosed
	}
	c.mu.RUnlock()

	if len(ids) == 0 {
		return nil, fmt.Errorf("no IDs to delete")
	}

	// Build placeholders / 플레이스홀더 빌드
	placeholders := strings.Repeat("?,", len(ids))
	placeholders = placeholders[:len(placeholders)-1] // Remove last comma / 마지막 쉼표 제거

	// Build query / 쿼리 빌드
	query := fmt.Sprintf("DELETE FROM %s WHERE %s IN (%s)", table, idColumn, placeholders)

	start := time.Now()

	// Execute with retry / 재시도로 실행
	var result sql.Result
	err := c.executeWithRetry(ctx, func() error {
		db := c.getCurrentConnection()
		var execErr error
		result, execErr = db.ExecContext(ctx, query, ids...)
		return execErr
	})

	duration := time.Since(start)

	if err != nil {
		c.logQuery(query, ids, duration, err)
		return nil, c.wrapError("BatchDelete", query, ids, err, duration)
	}

	c.logQuery(query, ids, duration, nil)
	return result, nil
}

// BatchSelectByIDs selects multiple rows by IDs in a single query
// BatchSelectByIDs는 단일 쿼리로 ID로 여러 행을 선택합니다
//
// Example / 예제:
//
//	ids := []interface{}{1, 2, 3, 4, 5}
//	users, err := db.BatchSelectByIDs(ctx, "users", "id", ids)
//
// This generates: SELECT * FROM users WHERE id IN (?,?,?,?,?)
func (c *Client) BatchSelectByIDs(ctx context.Context, table string, idColumn string, ids []interface{}) ([]map[string]interface{}, error) {
	c.mu.RLock()
	if c.closed {
		c.mu.RUnlock()
		return nil, ErrClosed
	}
	c.mu.RUnlock()

	if len(ids) == 0 {
		return []map[string]interface{}{}, nil
	}

	// Build placeholders / 플레이스홀더 빌드
	placeholders := strings.Repeat("?,", len(ids))
	placeholders = placeholders[:len(placeholders)-1] // Remove last comma / 마지막 쉼표 제거

	// Build query / 쿼리 빌드
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s IN (%s)", table, idColumn, placeholders)

	start := time.Now()

	// Execute with retry / 재시도로 실행
	var rows *sql.Rows
	err := c.executeWithRetry(ctx, func() error {
		db := c.getCurrentConnection()
		var execErr error
		rows, execErr = db.QueryContext(ctx, query, ids...)
		return execErr
	})

	duration := time.Since(start)

	if err != nil {
		c.logQuery(query, ids, duration, err)
		return nil, c.wrapError("BatchSelectByIDs", query, ids, err, duration)
	}

	// Scan rows / 행 스캔
	results, err := scanRows(rows)
	if err != nil {
		c.logQuery(query, ids, duration, err)
		return nil, c.wrapError("BatchSelectByIDs", query, ids, err, duration)
	}

	c.logQuery(query, ids, duration, nil)
	return results, nil
}
