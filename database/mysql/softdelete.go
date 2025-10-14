package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// SoftDelete marks rows as deleted by setting deleted_at timestamp
// SoftDelete는 deleted_at 타임스탬프를 설정하여 행을 삭제된 것으로 표시합니다
//
// Example / 예제:
//
//	result, err := db.SoftDelete(ctx, "users", "id = ?", 1)
func (c *Client) SoftDelete(ctx context.Context, table string, conditionAndArgs ...interface{}) (sql.Result, error) {
	if len(conditionAndArgs) == 0 {
		return nil, fmt.Errorf("condition is required for soft delete")
	}

	// Build UPDATE query to set deleted_at / deleted_at을 설정하는 UPDATE 쿼리 빌드
	data := map[string]interface{}{
		"deleted_at": time.Now(),
	}

	return c.UpdateContext(ctx, table, data, conditionAndArgs...)
}

// Restore restores soft-deleted rows by setting deleted_at to NULL
// Restore는 deleted_at을 NULL로 설정하여 소프트 삭제된 행을 복구합니다
//
// Example / 예제:
//
//	result, err := db.Restore(ctx, "users", "id = ?", 1)
func (c *Client) Restore(ctx context.Context, table string, conditionAndArgs ...interface{}) (sql.Result, error) {
	if len(conditionAndArgs) == 0 {
		return nil, fmt.Errorf("condition is required for restore")
	}

	// Build UPDATE query to set deleted_at to NULL / deleted_at을 NULL로 설정하는 UPDATE 쿼리 빌드
	data := map[string]interface{}{
		"deleted_at": nil,
	}

	return c.UpdateContext(ctx, table, data, conditionAndArgs...)
}

// SelectAllWithTrashed selects all rows including soft-deleted ones
// SelectAllWithTrashed는 소프트 삭제된 것을 포함하여 모든 행을 선택합니다
//
// Example / 예제:
//
//	users, err := db.SelectAllWithTrashed(ctx, "users")
//	users, err := db.SelectAllWithTrashed(ctx, "users", "age > ?", 18)
func (c *Client) SelectAllWithTrashed(ctx context.Context, table string, conditionAndArgs ...interface{}) ([]map[string]interface{}, error) {
	// Just use SelectAll - it doesn't filter by deleted_at
	// SelectAll을 사용 - deleted_at으로 필터링하지 않음
	return c.SelectAllContext(ctx, table, conditionAndArgs...)
}

// SelectAllOnlyTrashed selects only soft-deleted rows
// SelectAllOnlyTrashed는 소프트 삭제된 행만 선택합니다
//
// Example / 예제:
//
//	users, err := db.SelectAllOnlyTrashed(ctx, "users")
//	users, err := db.SelectAllOnlyTrashed(ctx, "users", "age > ?", 18)
func (c *Client) SelectAllOnlyTrashed(ctx context.Context, table string, conditionAndArgs ...interface{}) ([]map[string]interface{}, error) {
	// Add deleted_at IS NOT NULL condition / deleted_at IS NOT NULL 조건 추가
	var newCondition string
	var args []interface{}

	if len(conditionAndArgs) > 0 {
		condition := fmt.Sprintf("%v", conditionAndArgs[0])
		newCondition = fmt.Sprintf("(%s) AND deleted_at IS NOT NULL", condition)
		if len(conditionAndArgs) > 1 {
			args = conditionAndArgs[1:]
		}
	} else {
		newCondition = "deleted_at IS NOT NULL"
	}

	// Build new conditionAndArgs / 새 conditionAndArgs 빌드
	newConditionAndArgs := make([]interface{}, 0, 1+len(args))
	newConditionAndArgs = append(newConditionAndArgs, newCondition)
	newConditionAndArgs = append(newConditionAndArgs, args...)

	return c.SelectAllContext(ctx, table, newConditionAndArgs...)
}

// PermanentDelete performs actual deletion (physical delete) from database
// PermanentDelete는 데이터베이스에서 실제 삭제(물리적 삭제)를 수행합니다
//
// Example / 예제:
//
//	result, err := db.PermanentDelete(ctx, "users", "id = ?", 1)
func (c *Client) PermanentDelete(ctx context.Context, table string, conditionAndArgs ...interface{}) (sql.Result, error) {
	// Just use regular Delete / 일반 Delete 사용
	return c.DeleteContext(ctx, table, conditionAndArgs...)
}

// CountWithTrashed counts all rows including soft-deleted ones
// CountWithTrashed는 소프트 삭제된 것을 포함하여 모든 행을 계산합니다
func (c *Client) CountWithTrashed(ctx context.Context, table string, conditionAndArgs ...interface{}) (int64, error) {
	// Just use Count - it doesn't filter by deleted_at
	// Count를 사용 - deleted_at으로 필터링하지 않음
	return c.CountContext(ctx, table, conditionAndArgs...)
}

// CountOnlyTrashed counts only soft-deleted rows
// CountOnlyTrashed는 소프트 삭제된 행만 계산합니다
func (c *Client) CountOnlyTrashed(ctx context.Context, table string, conditionAndArgs ...interface{}) (int64, error) {
	// Add deleted_at IS NOT NULL condition / deleted_at IS NOT NULL 조건 추가
	var newCondition string
	var args []interface{}

	if len(conditionAndArgs) > 0 {
		condition := fmt.Sprintf("%v", conditionAndArgs[0])
		newCondition = fmt.Sprintf("(%s) AND deleted_at IS NOT NULL", condition)
		if len(conditionAndArgs) > 1 {
			args = conditionAndArgs[1:]
		}
	} else {
		newCondition = "deleted_at IS NOT NULL"
	}

	// Build new conditionAndArgs / 새 conditionAndArgs 빌드
	newConditionAndArgs := make([]interface{}, 0, 1+len(args))
	newConditionAndArgs = append(newConditionAndArgs, newCondition)
	newConditionAndArgs = append(newConditionAndArgs, args...)

	return c.CountContext(ctx, table, newConditionAndArgs...)
}
