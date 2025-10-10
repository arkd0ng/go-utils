package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

// SelectAll selects all rows from a table with optional conditions
// SelectAll은 선택적 조건으로 테이블의 모든 행을 선택합니다
//
// Example / 예제:
//
//	// Select all users / 모든 사용자 선택
//	users, err := db.SelectAll(ctx, "users")
//
//	// Select with condition / 조건과 함께 선택
//	users, err := db.SelectAll(ctx, "users", "age > ?", 18)
func (c *Client) SelectAll(ctx context.Context, table string, conditionAndArgs ...interface{}) ([]map[string]interface{}, error) {
	c.mu.RLock()
	if c.closed {
		c.mu.RUnlock()
		return nil, ErrClosed
	}
	c.mu.RUnlock()

	// Build query / 쿼리 빌드
	query := fmt.Sprintf("SELECT * FROM %s", table)
	var args []interface{}

	if len(conditionAndArgs) > 0 {
		condition := fmt.Sprintf("%v", conditionAndArgs[0])
		query += " WHERE " + condition
		if len(conditionAndArgs) > 1 {
			args = conditionAndArgs[1:]
		}
	}

	start := time.Now()

	// Execute with retry / 재시도로 실행
	var rows *sql.Rows
	err := c.executeWithRetry(ctx, func() error {
		db := c.getCurrentConnection()
		var execErr error
		rows, execErr = db.QueryContext(ctx, query, args...)
		return execErr
	})

	duration := time.Since(start)

	if err != nil {
		c.logQuery(query, args, duration, err)
		return nil, c.wrapError("SelectAll", query, args, err, duration)
	}

	// Scan rows / 행 스캔
	results, err := scanRows(rows)
	if err != nil {
		c.logQuery(query, args, duration, err)
		return nil, c.wrapError("SelectAll", query, args, err, duration)
	}

	c.logQuery(query, args, duration, nil)
	return results, nil
}

// SelectColumn selects all rows with a single column from a table
// SelectColumn은 테이블에서 단일 컬럼으로 모든 행을 선택합니다
//
// Example / 예제:
//
//	// Select all emails / 모든 이메일 선택
//	emails, err := db.SelectColumn(ctx, "users", "email")
//
//	// Select with condition / 조건과 함께 선택
//	emails, err := db.SelectColumn(ctx, "users", "email", "age > ?", 18)
func (c *Client) SelectColumn(ctx context.Context, table string, column string, conditionAndArgs ...interface{}) ([]map[string]interface{}, error) {
	c.mu.RLock()
	if c.closed {
		c.mu.RUnlock()
		return nil, ErrClosed
	}
	c.mu.RUnlock()

	// Build query / 쿼리 빌드
	query := fmt.Sprintf("SELECT %s FROM %s", column, table)
	var args []interface{}

	if len(conditionAndArgs) > 0 {
		condition := fmt.Sprintf("%v", conditionAndArgs[0])
		query += " WHERE " + condition
		if len(conditionAndArgs) > 1 {
			args = conditionAndArgs[1:]
		}
	}

	start := time.Now()

	// Execute with retry / 재시도로 실행
	var rows *sql.Rows
	err := c.executeWithRetry(ctx, func() error {
		db := c.getCurrentConnection()
		var execErr error
		rows, execErr = db.QueryContext(ctx, query, args...)
		return execErr
	})

	duration := time.Since(start)

	if err != nil {
		c.logQuery(query, args, duration, err)
		return nil, c.wrapError("SelectColumn", query, args, err, duration)
	}

	// Scan rows / 행 스캔
	results, err := scanRows(rows)
	if err != nil {
		c.logQuery(query, args, duration, err)
		return nil, c.wrapError("SelectColumn", query, args, err, duration)
	}

	c.logQuery(query, args, duration, nil)
	return results, nil
}

// SelectColumns selects all rows with multiple columns from a table
// SelectColumns는 테이블에서 여러 컬럼으로 모든 행을 선택합니다
//
// Example / 예제:
//
//	// Select multiple columns / 여러 컬럼 선택
//	users, err := db.SelectColumns(ctx, "users", []string{"name", "email", "age"})
//
//	// Select with condition / 조건과 함께 선택
//	users, err := db.SelectColumns(ctx, "users", []string{"name", "email"}, "age > ?", 18)
func (c *Client) SelectColumns(ctx context.Context, table string, columns []string, conditionAndArgs ...interface{}) ([]map[string]interface{}, error) {
	c.mu.RLock()
	if c.closed {
		c.mu.RUnlock()
		return nil, ErrClosed
	}
	c.mu.RUnlock()

	if len(columns) == 0 {
		return nil, fmt.Errorf("%w: no columns provided", ErrQueryFailed)
	}

	// Build query / 쿼리 빌드
	columnList := strings.Join(columns, ", ")
	query := fmt.Sprintf("SELECT %s FROM %s", columnList, table)
	var args []interface{}

	if len(conditionAndArgs) > 0 {
		condition := fmt.Sprintf("%v", conditionAndArgs[0])
		query += " WHERE " + condition
		if len(conditionAndArgs) > 1 {
			args = conditionAndArgs[1:]
		}
	}

	start := time.Now()

	// Execute with retry / 재시도로 실행
	var rows *sql.Rows
	err := c.executeWithRetry(ctx, func() error {
		db := c.getCurrentConnection()
		var execErr error
		rows, execErr = db.QueryContext(ctx, query, args...)
		return execErr
	})

	duration := time.Since(start)

	if err != nil {
		c.logQuery(query, args, duration, err)
		return nil, c.wrapError("SelectColumns", query, args, err, duration)
	}

	// Scan rows / 행 스캔
	results, err := scanRows(rows)
	if err != nil {
		c.logQuery(query, args, duration, err)
		return nil, c.wrapError("SelectColumns", query, args, err, duration)
	}

	c.logQuery(query, args, duration, nil)
	return results, nil
}

// SelectOne selects a single row from a table with conditions
// SelectOne은 조건과 함께 테이블에서 단일 행을 선택합니다
//
// Example / 예제:
//
//	user, err := db.SelectOne(ctx, "users", "id = ?", 123)
func (c *Client) SelectOne(ctx context.Context, table string, conditionAndArgs ...interface{}) (map[string]interface{}, error) {
	c.mu.RLock()
	if c.closed {
		c.mu.RUnlock()
		return nil, ErrClosed
	}
	c.mu.RUnlock()

	// Build query / 쿼리 빌드
	query := fmt.Sprintf("SELECT * FROM %s", table)
	var args []interface{}

	if len(conditionAndArgs) > 0 {
		condition := fmt.Sprintf("%v", conditionAndArgs[0])
		query += " WHERE " + condition
		if len(conditionAndArgs) > 1 {
			args = conditionAndArgs[1:]
		}
	}

	query += " LIMIT 1"

	start := time.Now()

	// Execute with retry / 재시도로 실행
	var rows *sql.Rows
	err := c.executeWithRetry(ctx, func() error {
		db := c.getCurrentConnection()
		var execErr error
		rows, execErr = db.QueryContext(ctx, query, args...)
		return execErr
	})

	duration := time.Since(start)

	if err != nil {
		c.logQuery(query, args, duration, err)
		return nil, c.wrapError("SelectOne", query, args, err, duration)
	}

	// Scan single row / 단일 행 스캔
	result, err := scanRow(rows)
	if err != nil {
		c.logQuery(query, args, duration, err)
		return nil, c.wrapError("SelectOne", query, args, err, duration)
	}

	c.logQuery(query, args, duration, nil)
	return result, nil
}

// Insert inserts a new row into a table
// Insert는 테이블에 새 행을 삽입합니다
//
// Example / 예제:
//
//	result, err := db.Insert(ctx, "users", map[string]interface{}{
//	    "name":  "John",
//	    "email": "john@example.com",
//	    "age":   30,
//	})
func (c *Client) Insert(ctx context.Context, table string, data map[string]interface{}) (sql.Result, error) {
	c.mu.RLock()
	if c.closed {
		c.mu.RUnlock()
		return nil, ErrClosed
	}
	c.mu.RUnlock()

	if len(data) == 0 {
		return nil, fmt.Errorf("%w: no data provided for insert", ErrQueryFailed)
	}

	// Build query / 쿼리 빌드
	columns := make([]string, 0, len(data))
	placeholders := make([]string, 0, len(data))
	values := make([]interface{}, 0, len(data))

	for col, val := range data {
		columns = append(columns, col)
		placeholders = append(placeholders, "?")
		values = append(values, val)
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		table,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "))

	start := time.Now()

	// Execute with retry / 재시도로 실행
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
		return nil, c.wrapError("Insert", query, values, err, duration)
	}

	c.logQuery(query, values, duration, nil)
	return result, nil
}

// Update updates rows in a table
// Update는 테이블의 행을 업데이트합니다
//
// Example / 예제:
//
//	result, err := db.Update(ctx, "users",
//	    map[string]interface{}{"name": "Jane", "age": 31},
//	    "id = ?", 123)
func (c *Client) Update(ctx context.Context, table string, data map[string]interface{}, conditionAndArgs ...interface{}) (sql.Result, error) {
	c.mu.RLock()
	if c.closed {
		c.mu.RUnlock()
		return nil, ErrClosed
	}
	c.mu.RUnlock()

	if len(data) == 0 {
		return nil, fmt.Errorf("%w: no data provided for update", ErrQueryFailed)
	}

	// Build SET clause / SET 절 빌드
	setClauses := make([]string, 0, len(data))
	values := make([]interface{}, 0, len(data))

	for col, val := range data {
		setClauses = append(setClauses, fmt.Sprintf("%s = ?", col))
		values = append(values, val)
	}

	// Build query / 쿼리 빌드
	query := fmt.Sprintf("UPDATE %s SET %s", table, strings.Join(setClauses, ", "))

	// Add WHERE clause if provided / 제공된 경우 WHERE 절 추가
	if len(conditionAndArgs) > 0 {
		condition := fmt.Sprintf("%v", conditionAndArgs[0])
		query += " WHERE " + condition
		if len(conditionAndArgs) > 1 {
			values = append(values, conditionAndArgs[1:]...)
		}
	}

	start := time.Now()

	// Execute with retry / 재시도로 실행
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
		return nil, c.wrapError("Update", query, values, err, duration)
	}

	c.logQuery(query, values, duration, nil)
	return result, nil
}

// Delete deletes rows from a table
// Delete는 테이블에서 행을 삭제합니다
//
// Example / 예제:
//
//	result, err := db.Delete(ctx, "users", "id = ?", 123)
func (c *Client) Delete(ctx context.Context, table string, conditionAndArgs ...interface{}) (sql.Result, error) {
	c.mu.RLock()
	if c.closed {
		c.mu.RUnlock()
		return nil, ErrClosed
	}
	c.mu.RUnlock()

	// Build query / 쿼리 빌드
	query := fmt.Sprintf("DELETE FROM %s", table)
	var args []interface{}

	if len(conditionAndArgs) > 0 {
		condition := fmt.Sprintf("%v", conditionAndArgs[0])
		query += " WHERE " + condition
		if len(conditionAndArgs) > 1 {
			args = conditionAndArgs[1:]
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
		return nil, c.wrapError("Delete", query, args, err, duration)
	}

	c.logQuery(query, args, duration, nil)
	return result, nil
}

// Count counts rows in a table with optional conditions
// Count는 선택적 조건으로 테이블의 행 수를 계산합니다
//
// Example / 예제:
//
//	count, err := db.Count(ctx, "users")
//	count, err := db.Count(ctx, "users", "age > ?", 18)
func (c *Client) Count(ctx context.Context, table string, conditionAndArgs ...interface{}) (int64, error) {
	c.mu.RLock()
	if c.closed {
		c.mu.RUnlock()
		return 0, ErrClosed
	}
	c.mu.RUnlock()

	// Build query / 쿼리 빌드
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", table)
	var args []interface{}

	if len(conditionAndArgs) > 0 {
		condition := fmt.Sprintf("%v", conditionAndArgs[0])
		query += " WHERE " + condition
		if len(conditionAndArgs) > 1 {
			args = conditionAndArgs[1:]
		}
	}

	start := time.Now()

	// Execute with retry / 재시도로 실행
	var rows *sql.Rows
	err := c.executeWithRetry(ctx, func() error {
		db := c.getCurrentConnection()
		var execErr error
		rows, execErr = db.QueryContext(ctx, query, args...)
		return execErr
	})

	duration := time.Since(start)

	if err != nil {
		c.logQuery(query, args, duration, err)
		return 0, c.wrapError("Count", query, args, err, duration)
	}

	// Scan count / 카운트 스캔
	count, err := scanCount(rows)
	if err != nil {
		c.logQuery(query, args, duration, err)
		return 0, c.wrapError("Count", query, args, err, duration)
	}

	c.logQuery(query, args, duration, nil)
	return count, nil
}

// Exists checks if at least one row exists with the given conditions
// Exists는 주어진 조건으로 최소한 하나의 행이 존재하는지 확인합니다
//
// Example / 예제:
//
//	exists, err := db.Exists(ctx, "users", "email = ?", "john@example.com")
func (c *Client) Exists(ctx context.Context, table string, conditionAndArgs ...interface{}) (bool, error) {
	count, err := c.Count(ctx, table, conditionAndArgs...)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
