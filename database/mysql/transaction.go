package mysql

import (
	"context"
	"database/sql"
	"fmt"
)

// Transaction executes a function within a database transaction
// Transaction은 데이터베이스 트랜잭션 내에서 함수를 실행합니다
//
// If the function returns an error, the transaction is rolled back.
// If the function succeeds, the transaction is committed.
//
// 함수가 에러를 반환하면 트랜잭션이 롤백됩니다.
// 함수가 성공하면 트랜잭션이 커밋됩니다.
//
// Example / 예제:
//
//	err := db.Transaction(ctx, func(tx *Tx) error {
//	    // Insert user / 사용자 삽입
//	    result, err := tx.Insert(ctx, "users", map[string]interface{}{
//	        "name": "John",
//	        "email": "john@example.com",
//	    })
//	    if err != nil {
//	        return err // Will rollback / 롤백됨
//	    }
//
//	    userID, _ := result.LastInsertId()
//
//	    // Insert profile / 프로필 삽입
//	    _, err = tx.Insert(ctx, "profiles", map[string]interface{}{
//	        "user_id": userID,
//	        "bio": "Hello world",
//	    })
//	    if err != nil {
//	        return err // Will rollback / 롤백됨
//	    }
//
//	    return nil // Will commit / 커밋됨
//	})
func (c *Client) Transaction(ctx context.Context, fn func(*Tx) error) error {
	tx, err := c.Begin(ctx)
	if err != nil {
		return err
	}

	// Ensure rollback if panic / 패닉 시 롤백 보장
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Re-throw panic / 패닉 다시 발생
		}
	}()

	// Execute function / 함수 실행
	if err := fn(tx); err != nil {
		// Rollback on error / 에러 시 롤백
		if rbErr := tx.Rollback(); rbErr != nil {
			if c.config.logger != nil {
				c.config.logger.Error("Failed to rollback transaction",
					"error", rbErr,
					"original_error", err)
			}
		}
		return err
	}

	// Commit on success / 성공 시 커밋
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("%w: %v", ErrTransactionFailed, err)
	}

	return nil
}

// Insert inserts a new row within the transaction
// Insert는 트랜잭션 내에서 새 행을 삽입합니다
func (t *Tx) Insert(table string, data map[string]interface{}) (sql.Result, error) {
	return t.InsertContext(context.Background(), table, data)
}

// InsertContext inserts a new row within the transaction
// InsertContext는 트랜잭션 내에서 새 행을 삽입합니다
func (t *Tx) InsertContext(ctx context.Context, table string, data map[string]interface{}) (sql.Result, error) {
	if t.finished {
		return nil, ErrTransactionFailed
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("%w: no data provided for insert", ErrQueryFailed)
	}

	// Build query (same as simple.go but use tx) / 쿼리 빌드 (simple.go와 동일하지만 tx 사용)
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
		joinStrings(columns, ", "),
		joinStrings(placeholders, ", "))

	return t.tx.ExecContext(ctx, query, values...)
}

// Update updates rows within the transaction
// Update는 트랜잭션 내에서 행을 업데이트합니다
func (t *Tx) Update(table string, data map[string]interface{}, conditionAndArgs ...interface{}) (sql.Result, error) {
	return t.UpdateContext(context.Background(), table, data, conditionAndArgs...)
}

// UpdateContext updates rows within the transaction
// UpdateContext는 트랜잭션 내에서 행을 업데이트합니다
func (t *Tx) UpdateContext(ctx context.Context, table string, data map[string]interface{}, conditionAndArgs ...interface{}) (sql.Result, error) {
	if t.finished {
		return nil, ErrTransactionFailed
	}

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
	query := fmt.Sprintf("UPDATE %s SET %s", table, joinStrings(setClauses, ", "))

	// Add WHERE clause if provided / 제공된 경우 WHERE 절 추가
	if len(conditionAndArgs) > 0 {
		condition := fmt.Sprintf("%v", conditionAndArgs[0])
		query += " WHERE " + condition
		if len(conditionAndArgs) > 1 {
			values = append(values, conditionAndArgs[1:]...)
		}
	}

	return t.tx.ExecContext(ctx, query, values...)
}

// Delete deletes rows within the transaction
// Delete는 트랜잭션 내에서 행을 삭제합니다
func (t *Tx) Delete(table string, conditionAndArgs ...interface{}) (sql.Result, error) {
	return t.DeleteContext(context.Background(), table, conditionAndArgs...)
}

// DeleteContext deletes rows within the transaction
// DeleteContext는 트랜잭션 내에서 행을 삭제합니다
func (t *Tx) DeleteContext(ctx context.Context, table string, conditionAndArgs ...interface{}) (sql.Result, error) {
	if t.finished {
		return nil, ErrTransactionFailed
	}

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

	return t.tx.ExecContext(ctx, query, args...)
}

// SelectAll selects all rows within the transaction
// SelectAll은 트랜잭션 내에서 모든 행을 선택합니다
func (t *Tx) SelectAll(table string, conditionAndArgs ...interface{}) ([]map[string]interface{}, error) {
	return t.SelectAllContext(context.Background(), table, conditionAndArgs...)
}

// SelectAllContext selects all rows within the transaction
// SelectAllContext는 트랜잭션 내에서 모든 행을 선택합니다
func (t *Tx) SelectAllContext(ctx context.Context, table string, conditionAndArgs ...interface{}) ([]map[string]interface{}, error) {
	if t.finished {
		return nil, ErrTransactionFailed
	}

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

	rows, err := t.tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return scanRows(rows)
}

// SelectColumn selects all rows with a single column within the transaction
// SelectColumn은 트랜잭션 내에서 단일 컬럼으로 모든 행을 선택합니다
func (t *Tx) SelectColumn(table string, column string, conditionAndArgs ...interface{}) ([]map[string]interface{}, error) {
	return t.SelectColumnContext(context.Background(), table, column, conditionAndArgs...)
}

// SelectColumnContext selects all rows with a single column within the transaction
// SelectColumnContext는 트랜잭션 내에서 단일 컬럼으로 모든 행을 선택합니다
func (t *Tx) SelectColumnContext(ctx context.Context, table string, column string, conditionAndArgs ...interface{}) ([]map[string]interface{}, error) {
	if t.finished {
		return nil, ErrTransactionFailed
	}

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

	rows, err := t.tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return scanRows(rows)
}

// SelectColumns selects all rows with multiple columns within the transaction
// SelectColumns는 트랜잭션 내에서 여러 컬럼으로 모든 행을 선택합니다
func (t *Tx) SelectColumns(table string, columns []string, conditionAndArgs ...interface{}) ([]map[string]interface{}, error) {
	return t.SelectColumnsContext(context.Background(), table, columns, conditionAndArgs...)
}

// SelectColumnsContext selects all rows with multiple columns within the transaction
// SelectColumnsContext는 트랜잭션 내에서 여러 컬럼으로 모든 행을 선택합니다
func (t *Tx) SelectColumnsContext(ctx context.Context, table string, columns []string, conditionAndArgs ...interface{}) ([]map[string]interface{}, error) {
	if t.finished {
		return nil, ErrTransactionFailed
	}

	if len(columns) == 0 {
		return nil, fmt.Errorf("%w: no columns provided", ErrQueryFailed)
	}

	// Build query / 쿼리 빌드
	columnList := joinStrings(columns, ", ")
	query := fmt.Sprintf("SELECT %s FROM %s", columnList, table)
	var args []interface{}

	if len(conditionAndArgs) > 0 {
		condition := fmt.Sprintf("%v", conditionAndArgs[0])
		query += " WHERE " + condition
		if len(conditionAndArgs) > 1 {
			args = conditionAndArgs[1:]
		}
	}

	rows, err := t.tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return scanRows(rows)
}

// SelectOne selects a single row within the transaction
// SelectOne은 트랜잭션 내에서 단일 행을 선택합니다
func (t *Tx) SelectOne(table string, conditionAndArgs ...interface{}) (map[string]interface{}, error) {
	return t.SelectOneContext(context.Background(), table, conditionAndArgs...)
}

// SelectOneContext selects a single row within the transaction
// SelectOneContext는 트랜잭션 내에서 단일 행을 선택합니다
func (t *Tx) SelectOneContext(ctx context.Context, table string, conditionAndArgs ...interface{}) (map[string]interface{}, error) {
	if t.finished {
		return nil, ErrTransactionFailed
	}

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

	rows, err := t.tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return scanRow(rows)
}

// joinStrings joins a slice of strings with a separator
// joinStrings는 구분자로 문자열 슬라이스를 결합합니다
func joinStrings(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}
	result := strs[0]
	for i := 1; i < len(strs); i++ {
		result += sep + strs[i]
	}
	return result
}
