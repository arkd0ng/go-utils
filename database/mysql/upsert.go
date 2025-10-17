package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

// Upsert inserts a row or updates it if it already exists (INSERT ... ON DUPLICATE KEY UPDATE)
// Upsert는 행을 삽입하거나 이미 존재하면 업데이트합니다 (INSERT ... ON DUPLICATE KEY UPDATE)
//
// Example
// 예제:
//
//	data := map[string]interface{}{
//	    "email": "john@example.com",  // Unique key
//	    "name": "John Doe",
//	    "age": 30,
//	}
//	updateColumns := []string{"name", "age"}  // Columns to update on duplicate
//	result, err := db.Upsert(ctx, "users", data, updateColumns)
//
// This generates: INSERT INTO users (email, name, age) VALUES (?,?,?)
//                 ON DUPLICATE KEY UPDATE name=VALUES(name), age=VALUES(age)
func (c *Client) Upsert(ctx context.Context, table string, data map[string]interface{}, updateColumns []string) (sql.Result, error) {
	c.mu.RLock()
	if c.closed {
		c.mu.RUnlock()
		return nil, ErrClosed
	}
	c.mu.RUnlock()

	if len(data) == 0 {
		return nil, fmt.Errorf("no data to upsert")
	}

	// Extract columns and values
	// 컬럼과 값 추출
	var columns []string
	var values []interface{}
	for col, val := range data {
		columns = append(columns, col)
		values = append(values, val)
	}

	// Build column list
	// 컬럼 목록 빌드
	columnList := strings.Join(columns, ", ")

	// Build value placeholders
	// 값 플레이스홀더 빌드
	placeholders := strings.Repeat("?,", len(columns))
	placeholders = placeholders[:len(placeholders)-1] // Remove last comma / 마지막 쉼표 제거

	// Build UPDATE clause
	// UPDATE 절 빌드
	var updateClauses []string
	if len(updateColumns) == 0 {
		// If no update columns specified, update all columns except those in data
		// 업데이트 컬럼이 지정되지 않은 경우 데이터의 모든 컬럼 업데이트
		for _, col := range columns {
			updateClauses = append(updateClauses, fmt.Sprintf("%s=VALUES(%s)", col, col))
		}
	} else {
		// Update only specified columns
		// 지정된 컬럼만 업데이트
		for _, col := range updateColumns {
			updateClauses = append(updateClauses, fmt.Sprintf("%s=VALUES(%s)", col, col))
		}
	}

	updateClause := strings.Join(updateClauses, ", ")

	// Build query
	// 쿼리 빌드
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) ON DUPLICATE KEY UPDATE %s",
		table, columnList, placeholders, updateClause)

	start := time.Now()

	// Execute with retry
	// 재시도로 실행
	var result sql.Result
	err := c.executeWithRetry(ctx, func() error {
		db := c.getCurrentConnection()
		var execErr error
		result, execErr = db.ExecContext(ctx, query, values...)
		return execErr
	})

	duration := time.Since(start)

	if err != nil {
		c.logQuery(query, values, duration, err)
		return nil, c.wrapError("Upsert", query, values, err, duration)
	}

	c.logQuery(query, values, duration, nil)
	return result, nil
}

// UpsertBatch performs multiple upsert operations in a single query
// UpsertBatch는 단일 쿼리로 여러 upsert 작업을 수행합니다
//
// Example
// 예제:
//
//	data := []map[string]interface{}{
//	    {"email": "john@example.com", "name": "John", "age": 30},
//	    {"email": "jane@example.com", "name": "Jane", "age": 25},
//	}
//	updateColumns := []string{"name", "age"}
//	err := db.UpsertBatch(ctx, "users", data, updateColumns)
//
// This generates: INSERT INTO users (email, name, age) VALUES (?,?,?),(?,?,?)
//                 ON DUPLICATE KEY UPDATE name=VALUES(name), age=VALUES(age)
func (c *Client) UpsertBatch(ctx context.Context, table string, data []map[string]interface{}, updateColumns []string) (sql.Result, error) {
	c.mu.RLock()
	if c.closed {
		c.mu.RUnlock()
		return nil, ErrClosed
	}
	c.mu.RUnlock()

	if len(data) == 0 {
		return nil, fmt.Errorf("no data to upsert")
	}

	// Get column names from first row
	// 첫 번째 행에서 컬럼 이름 가져오기
	var columns []string
	for col := range data[0] {
		columns = append(columns, col)
	}

	if len(columns) == 0 {
		return nil, fmt.Errorf("no columns to upsert")
	}

	// Build column list
	// 컬럼 목록 빌드
	columnList := strings.Join(columns, ", ")

	// Build value placeholders
	// 값 플레이스홀더 빌드
	valuePlaceholder := "(" + strings.Repeat("?,", len(columns))
	valuePlaceholder = valuePlaceholder[:len(valuePlaceholder)-1] + ")"

	// Build multiple value placeholders
	// 여러 값 플레이스홀더 빌드
	var valuePlaceholders []string
	for i := 0; i < len(data); i++ {
		valuePlaceholders = append(valuePlaceholders, valuePlaceholder)
	}

	// Build UPDATE clause
	// UPDATE 절 빌드
	var updateClauses []string
	if len(updateColumns) == 0 {
		// If no update columns specified, update all columns
		// 업데이트 컬럼이 지정되지 않은 경우 모든 컬럼 업데이트
		for _, col := range columns {
			updateClauses = append(updateClauses, fmt.Sprintf("%s=VALUES(%s)", col, col))
		}
	} else {
		// Update only specified columns
		// 지정된 컬럼만 업데이트
		for _, col := range updateColumns {
			updateClauses = append(updateClauses, fmt.Sprintf("%s=VALUES(%s)", col, col))
		}
	}

	updateClause := strings.Join(updateClauses, ", ")

	// Build query
	// 쿼리 빌드
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s ON DUPLICATE KEY UPDATE %s",
		table, columnList, strings.Join(valuePlaceholders, ","), updateClause)

	// Collect all values
	// 모든 값 수집
	var args []interface{}
	for _, row := range data {
		for _, col := range columns {
			args = append(args, row[col])
		}
	}

	start := time.Now()

	// Execute with retry
	// 재시도로 실행
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
		return nil, c.wrapError("UpsertBatch", query, args, err, duration)
	}

	c.logQuery(query, args, duration, nil)
	return result, nil
}

// Replace performs a REPLACE operation (delete + insert)
// Replace는 REPLACE 작업을 수행합니다 (삭제 + 삽입)
//
// Note: REPLACE deletes the old row and inserts a new one, which can have side effects
// with foreign keys and auto-increment values.
// 주의: REPLACE는 기존 행을 삭제하고 새 행을 삽입하므로 외래 키 및 auto_increment 값에
// 부작용이 있을 수 있습니다.
//
// Example
// 예제:
//
//	data := map[string]interface{}{
//	    "id": 1,
//	    "name": "John Doe",
//	    "age": 30,
//	}
//	result, err := db.Replace(ctx, "users", data)
//
// This generates: REPLACE INTO users (id, name, age) VALUES (?,?,?)
func (c *Client) Replace(ctx context.Context, table string, data map[string]interface{}) (sql.Result, error) {
	c.mu.RLock()
	if c.closed {
		c.mu.RUnlock()
		return nil, ErrClosed
	}
	c.mu.RUnlock()

	if len(data) == 0 {
		return nil, fmt.Errorf("no data to replace")
	}

	// Extract columns and values
	// 컬럼과 값 추출
	var columns []string
	var values []interface{}
	for col, val := range data {
		columns = append(columns, col)
		values = append(values, val)
	}

	// Build column list
	// 컬럼 목록 빌드
	columnList := strings.Join(columns, ", ")

	// Build value placeholders
	// 값 플레이스홀더 빌드
	placeholders := strings.Repeat("?,", len(columns))
	placeholders = placeholders[:len(placeholders)-1] // Remove last comma / 마지막 쉼표 제거

	// Build query
	// 쿼리 빌드
	query := fmt.Sprintf("REPLACE INTO %s (%s) VALUES (%s)",
		table, columnList, placeholders)

	start := time.Now()

	// Execute with retry
	// 재시도로 실행
	var result sql.Result
	err := c.executeWithRetry(ctx, func() error {
		db := c.getCurrentConnection()
		var execErr error
		result, execErr = db.ExecContext(ctx, query, values...)
		return execErr
	})

	duration := time.Since(start)

	if err != nil {
		c.logQuery(query, values, duration, err)
		return nil, c.wrapError("Replace", query, values, err, duration)
	}

	c.logQuery(query, values, duration, nil)
	return result, nil
}
