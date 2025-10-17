package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

// BatchInsert collapses multiple rows into a single INSERT statement while preserving column order and retrying transient faults.
// BatchInsert는 여러 행을 하나의 INSERT 문으로 묶어 컬럼 순서를 유지하고, 일시적 장애가 발생하면 재시도를 수행합니다.
//
// Example
// 예제:
//
//	data := []map[string]interface{}{
//	    {"name": "John", "age": 30, "email": "john@example.com"},
//	    {"name": "Jane", "age": 25, "email": "jane@example.com"},
//	    {"name": "Bob", "age": 35, "email": "bob@example.com"},
//	}
//	result, err := db.BatchInsert(ctx, "users", data)
//
// The generated SQL looks like:
// 생성되는 SQL 예시는 다음과 같습니다:
//
//	INSERT INTO users (name, age, email) VALUES (?,?,?),(?,?,?),(?,?,?)
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

	// Capture the column order from the first row so every placeholder array aligns with the same ordering.
	// 첫 번째 행에서 컬럼 순서를 추출하여, 이후 플레이스홀더 배열이 항상 동일한 순서를 따르도록 보장합니다.
	var columns []string
	for col := range data[0] {
		columns = append(columns, col)
	}

	if len(columns) == 0 {
		return nil, fmt.Errorf("no columns to insert")
	}

	// Render the comma-separated column projection used in the INSERT clause.
	// INSERT 절에서 사용할 콤마로 구분된 컬럼 목록을 생성합니다.
	columnList := strings.Join(columns, ", ")

	// Prepare a single parenthesised placeholder sequence such as "(?,?,?)".
	// "(?,?,?)"와 같이 괄호로 감싼 플레이스홀더 시퀀스를 준비합니다.
	valuePlaceholder := "(" + strings.Repeat("?,", len(columns))
	valuePlaceholder = valuePlaceholder[:len(valuePlaceholder)-1] + ")"

	// Duplicate the placeholder block for every row so that all values can be supplied in a single Exec call.
	// 단일 Exec 호출로 모든 값을 전달할 수 있도록 행 수만큼 플레이스홀더 블록을 복제합니다.
	var valuePlaceholders []string
	for i := 0; i < len(data); i++ {
		valuePlaceholders = append(valuePlaceholders, valuePlaceholder)
	}

	// Combine the projection and placeholder list into the final INSERT statement.
	// 컬럼 목록과 플레이스홀더 목록을 결합하여 최종 INSERT 문을 구성합니다.
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s",
		table, columnList, strings.Join(valuePlaceholders, ","))

	// Flatten every map value into a contiguous args slice that matches the generated placeholders.
	// 각 행의 값을 생성된 플레이스홀더 순서와 일치하도록 하나의 연속된 슬라이스로 평탄화합니다.
	var args []interface{}
	for _, row := range data {
		for _, col := range columns {
			args = append(args, row[col])
		}
	}

	start := time.Now()

	// Execute the statement using the shared retry policy to absorb transient network or lock issues.
	// 네트워크 지연이나 잠금과 같은 일시적 문제를 흡수하기 위해 공유 재시도 정책을 사용해 구문을 실행합니다.
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

// BatchUpdateItem describes one UPDATE statement entry consisting of the SET map and the WHERE clause.
// BatchUpdateItem은 SET에 사용할 데이터 맵과 WHERE 절 조건으로 구성된 단일 UPDATE 항목을 표현합니다.
type BatchUpdateItem struct {
	// Data holds column/value pairs to set during the update.
	// Data에는 갱신 시 설정할 컬럼과 값의 쌍이 저장됩니다.
	Data map[string]interface{}
	// ConditionAndArgs keeps the WHERE clause string followed by placeholder arguments.
	// ConditionAndArgs에는 WHERE 절 문자열과 해당 플레이스홀더 인자가 순서대로 들어 있습니다.
	ConditionAndArgs []interface{}
}

// BatchUpdate applies every supplied BatchUpdateItem inside a single transaction to guarantee atomic writes.
// BatchUpdate는 전달된 모든 BatchUpdateItem을 하나의 트랜잭션 안에서 실행하여 원자적 쓰기를 보장합니다.
//
// Example
// 예제:
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

	// Run the per-row updates within a transaction so that either all succeed or none are applied.
	// 각 행에 대한 업데이트를 트랜잭션으로 감싸 모든 변경이 함께 성공하거나 전부 롤백되도록 합니다.
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

// BatchDelete removes every row whose identifier appears in ids by issuing a single DELETE ... IN (...) clause.
// BatchDelete는 IDs에 포함된 식별자를 가진 모든 행을 DELETE ... IN (...) 구문 한 번으로 제거합니다.
//
// Example
// 예제:
//
//	ids := []interface{}{1, 2, 3, 4, 5}
//	result, err := db.BatchDelete(ctx, "users", "id", ids)
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

	// Construct the "?,?,?" placeholder list matching the number of identifiers supplied.
	// 전달된 식별자 개수에 맞는 "?,?,?" 형태의 플레이스홀더 목록을 구성합니다.
	placeholders := strings.Repeat("?,", len(ids))
	// Trim the trailing comma that Repeat leaves at the end.
	// Repeat 호출이 남긴 마지막 쉼표를 제거합니다.
	placeholders = placeholders[:len(placeholders)-1]

	// Build the DELETE query using the generated placeholder list.
	// 생성된 플레이스홀더 목록을 사용해 DELETE 쿼리를 완성합니다.
	query := fmt.Sprintf("DELETE FROM %s WHERE %s IN (%s)", table, idColumn, placeholders)

	start := time.Now()

	// Execute the statement with the same retry semantics used for inserts.
	// INSERT와 동일한 재시도 규칙으로 구문을 실행합니다.
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

// BatchSelectByIDs fetches every row whose identifier is listed, returning the results as a slice of column maps.
// BatchSelectByIDs는 전달된 식별자 목록에 해당하는 모든 행을 조회하여 컬럼 맵 슬라이스로 반환합니다.
//
// Example
// 예제:
//
//	ids := []interface{}{1, 2, 3, 4, 5}
//	users, err := db.BatchSelectByIDs(ctx, "users", "id", ids)
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

	// Build the placeholder sequence that mirrors the number of IDs passed in.
	// 전달된 ID 개수와 동일한 플레이스홀더 시퀀스를 생성합니다.
	placeholders := strings.Repeat("?,", len(ids))
	// Remove the extra comma introduced by Repeat so the SQL remains valid.
	// Repeat로 인해 추가된 마지막 쉼표를 제거하여 올바른 SQL을 유지합니다.
	placeholders = placeholders[:len(placeholders)-1]

	// Assemble the SELECT statement that filters by the provided identifiers.
	// 전달받은 식별자에 따라 필터링하는 SELECT 문을 구성합니다.
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s IN (%s)", table, idColumn, placeholders)

	start := time.Now()

	// Issue the query using the shared retry wrapper to guard against transient errors.
	// 일시적 오류에 대비하기 위해 재시도 래퍼를 사용하여 쿼리를 실행합니다.
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

	// Convert the sql.Rows cursor into a slice of map[string]interface{} for ergonomic consumption.
	// sql.Rows 커서를 map[string]interface{} 슬라이스로 변환하여 사용성을 높입니다.
	results, err := scanRows(rows)
	if err != nil {
		c.logQuery(query, ids, duration, err)
		return nil, c.wrapError("BatchSelectByIDs", query, ids, err, duration)
	}

	c.logQuery(query, ids, duration, nil)
	return results, nil
}
