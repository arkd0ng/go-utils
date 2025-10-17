package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

// SelectOption is a functional option for customizing SELECT queries
// SelectOption은 SELECT 쿼리를 커스터마이징하기 위한 함수형 옵션입니다
type SelectOption func(*selectConfig)

// selectConfig holds configuration for SELECT queries
// selectConfig는 SELECT 쿼리를 위한 설정을 보관합니다
type selectConfig struct {
	// Columns to select
	// 선택할 컬럼
	columns []string
	// JOIN clauses
	// JOIN 절
	joins []joinClause
	// ORDER BY clause
	// ORDER BY 절
	orderBy string
	// GROUP BY columns
	// GROUP BY 컬럼
	groupBy []string
	// HAVING clause
	// HAVING 절
	having *whereClause
	// LIMIT value
	// LIMIT 값
	limit *int
	// OFFSET value
	// OFFSET 값
	offset *int
	// DISTINCT flag
	// DISTINCT 플래그
	distinct bool
}

// WithColumns specifies which columns to select
// WithColumns는 선택할 컬럼을 지정합니다
//
// Example
// 예제:
//
//	db.SelectWhere(ctx, "users", "age > ?", 18, WithColumns("name", "email"))
func WithColumns(cols ...string) SelectOption {
	return func(c *selectConfig) {
		c.columns = cols
	}
}

// WithOrderBy adds ORDER BY clause
// WithOrderBy는 ORDER BY 절을 추가합니다
//
// Example
// 예제:
//
//	db.SelectWhere(ctx, "users", "age > ?", 18, WithOrderBy("age DESC"))
func WithOrderBy(order string) SelectOption {
	return func(c *selectConfig) {
		c.orderBy = order
	}
}

// WithLimit adds LIMIT clause
// WithLimit은 LIMIT 절을 추가합니다
//
// Example
// 예제:
//
//	db.SelectWhere(ctx, "users", "age > ?", 18, WithLimit(10))
func WithLimit(n int) SelectOption {
	return func(c *selectConfig) {
		c.limit = &n
	}
}

// WithOffset adds OFFSET clause
// WithOffset은 OFFSET 절을 추가합니다
//
// Example
// 예제:
//
//	db.SelectWhere(ctx, "users", "age > ?", 18, WithLimit(10), WithOffset(20))
func WithOffset(n int) SelectOption {
	return func(c *selectConfig) {
		c.offset = &n
	}
}

// WithGroupBy adds GROUP BY clause
// WithGroupBy는 GROUP BY 절을 추가합니다
//
// Example
// 예제:
//
//	db.SelectWhere(ctx, "users", "", WithGroupBy("city"), WithColumns("city", "COUNT(*) as count"))
func WithGroupBy(cols ...string) SelectOption {
	return func(c *selectConfig) {
		c.groupBy = cols
	}
}

// WithHaving adds HAVING clause (used with GROUP BY)
// WithHaving은 HAVING 절을 추가합니다 (GROUP BY와 함께 사용)
//
// Example
// 예제:
//
//	db.SelectWhere(ctx, "users", "",
//	  WithColumns("city", "COUNT(*) as count"),
//	  WithGroupBy("city"),
//	  WithHaving("COUNT(*) > ?", 2))
func WithHaving(condition string, args ...interface{}) SelectOption {
	return func(c *selectConfig) {
		c.having = &whereClause{
			condition: condition,
			args:      args,
		}
	}
}

// WithJoin adds INNER JOIN clause
// WithJoin은 INNER JOIN 절을 추가합니다
//
// Example
// 예제:
//
//	db.SelectWhere(ctx, "users u", "u.age > ?", 18,
//	  WithJoin("orders o", "u.id = o.user_id"),
//	  WithColumns("u.name", "o.total"))
func WithJoin(table, condition string) SelectOption {
	return func(c *selectConfig) {
		c.joins = append(c.joins, joinClause{
			joinType:  "INNER JOIN",
			table:     table,
			condition: condition,
		})
	}
}

// WithLeftJoin adds LEFT JOIN clause
// WithLeftJoin은 LEFT JOIN 절을 추가합니다
//
// Example
// 예제:
//
//	db.SelectWhere(ctx, "users u", "",
//	  WithLeftJoin("orders o", "u.id = o.user_id"),
//	  WithColumns("u.name", "COUNT(o.id) as order_count"),
//	  WithGroupBy("u.id", "u.name"))
func WithLeftJoin(table, condition string) SelectOption {
	return func(c *selectConfig) {
		c.joins = append(c.joins, joinClause{
			joinType:  "LEFT JOIN",
			table:     table,
			condition: condition,
		})
	}
}

// WithRightJoin adds RIGHT JOIN clause
// WithRightJoin은 RIGHT JOIN 절을 추가합니다
func WithRightJoin(table, condition string) SelectOption {
	return func(c *selectConfig) {
		c.joins = append(c.joins, joinClause{
			joinType:  "RIGHT JOIN",
			table:     table,
			condition: condition,
		})
	}
}

// WithDistinct adds DISTINCT keyword
// WithDistinct는 DISTINCT 키워드를 추가합니다
//
// Example
// 예제:
//
//	db.SelectWhere(ctx, "users", "", WithColumns("city"), WithDistinct())
func WithDistinct() SelectOption {
	return func(c *selectConfig) {
		c.distinct = true
	}
}

// SelectWhere selects rows with optional WHERE condition and options
// SelectWhere는 선택적 WHERE 조건과 옵션으로 행을 선택합니다
//
// Example
// 예제:
//
// // Simple query with columns
// 컬럼 지정 간단 쿼리
//
//	users, _ := db.SelectWhere(ctx, "users", "age > ?", 18,
//	  WithColumns("name", "email"),
//	  WithOrderBy("age DESC"),
//	  WithLimit(10))
//
// // GROUP BY query
// GROUP BY 쿼리
//
//	results, _ := db.SelectWhere(ctx, "users", "",
//	  WithColumns("city", "COUNT(*) as count"),
//	  WithGroupBy("city"),
//	  WithHaving("COUNT(*) > ?", 2),
//	  WithOrderBy("count DESC"))
//
// // JOIN query
// JOIN 쿼리
//
//	results, _ := db.SelectWhere(ctx, "users u", "u.age > ?", 25,
//	  WithJoin("orders o", "u.id = o.user_id"),
//	  WithColumns("u.name", "o.total"),
//	  WithOrderBy("o.total DESC"))
func (c *Client) SelectWhere(ctx context.Context, table string, conditionAndArgs ...interface{}) ([]map[string]interface{}, error) {
	c.mu.RLock()
	if c.closed {
		c.mu.RUnlock()
		return nil, ErrClosed
	}
	c.mu.RUnlock()

	// Parse condition and options
	// 조건과 옵션 파싱
	var condition string
	var args []interface{}
	var opts []SelectOption

	if len(conditionAndArgs) > 0 {
		// First argument is condition string
		// 첫 번째 인자는 조건 문자열
		condition = fmt.Sprintf("%v", conditionAndArgs[0])

		// Extract args and options
		// 인자와 옵션 추출
		for i := 1; i < len(conditionAndArgs); i++ {
			if opt, ok := conditionAndArgs[i].(SelectOption); ok {
				opts = append(opts, opt)
			} else {
				args = append(args, conditionAndArgs[i])
			}
		}
	}

	// Apply options
	// 옵션 적용
	cfg := &selectConfig{
		columns: []string{"*"},
	}
	for _, opt := range opts {
		opt(cfg)
	}

	// Build query
	// 쿼리 빌드
	query := buildSelectQuery(table, condition, cfg)

	// Add args from HAVING clause
	// HAVING 절의 인자 추가
	if cfg.having != nil {
		args = append(args, cfg.having.args...)
	}

	// Execute with retry
	// 재시도로 실행
	var rows *sql.Rows
	err := c.executeWithRetry(ctx, func() error {
		db := c.getCurrentConnection()
		var execErr error
		rows, execErr = db.QueryContext(ctx, query, args...)
		return execErr
	})

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return scanRows(rows)
}

// SelectOneWhere selects a single row with optional WHERE condition and options
// SelectOneWhere는 선택적 WHERE 조건과 옵션으로 단일 행을 선택합니다
//
// Example
// 예제:
//
//	user, _ := db.SelectOneWhere(ctx, "users", "email = ?", "john@example.com",
//	  WithColumns("name", "age", "city"))
func (c *Client) SelectOneWhere(ctx context.Context, table string, conditionAndArgs ...interface{}) (map[string]interface{}, error) {
	// Add LIMIT 1
	// LIMIT 1 추가
	conditionAndArgs = append(conditionAndArgs, WithLimit(1))

	results, err := c.SelectWhere(ctx, table, conditionAndArgs...)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no rows found")
	}

	return results[0], nil
}

// buildSelectQuery builds the SQL SELECT query from config
// buildSelectQuery는 설정에서 SQL SELECT 쿼리를 빌드합니다
func buildSelectQuery(table string, condition string, cfg *selectConfig) string {
	var parts []string

	// SELECT
	selectClause := "SELECT"
	if cfg.distinct {
		selectClause += " DISTINCT"
	}
	selectClause += " " + strings.Join(cfg.columns, ", ")
	parts = append(parts, selectClause)

	// FROM
	parts = append(parts, "FROM "+table)

	// JOINs
	for _, join := range cfg.joins {
		parts = append(parts, fmt.Sprintf("%s %s ON %s", join.joinType, join.table, join.condition))
	}

	// WHERE
	if condition != "" {
		parts = append(parts, "WHERE "+condition)
	}

	// GROUP BY
	if len(cfg.groupBy) > 0 {
		parts = append(parts, "GROUP BY "+strings.Join(cfg.groupBy, ", "))
	}

	// HAVING
	if cfg.having != nil {
		parts = append(parts, "HAVING "+cfg.having.condition)
	}

	// ORDER BY
	if cfg.orderBy != "" {
		parts = append(parts, "ORDER BY "+cfg.orderBy)
	}

	// LIMIT
	if cfg.limit != nil {
		parts = append(parts, fmt.Sprintf("LIMIT %d", *cfg.limit))
	}

	// OFFSET
	if cfg.offset != nil {
		parts = append(parts, fmt.Sprintf("OFFSET %d", *cfg.offset))
	}

	return strings.Join(parts, " ")
}

// SelectWhere for transactions
// 트랜잭션용 SelectWhere
func (tx *Tx) SelectWhere(ctx context.Context, table string, conditionAndArgs ...interface{}) ([]map[string]interface{}, error) {
	// Parse condition and options
	// 조건과 옵션 파싱
	var condition string
	var args []interface{}
	var opts []SelectOption

	if len(conditionAndArgs) > 0 {
		condition = fmt.Sprintf("%v", conditionAndArgs[0])

		for i := 1; i < len(conditionAndArgs); i++ {
			if opt, ok := conditionAndArgs[i].(SelectOption); ok {
				opts = append(opts, opt)
			} else {
				args = append(args, conditionAndArgs[i])
			}
		}
	}

	// Apply options
	// 옵션 적용
	cfg := &selectConfig{
		columns: []string{"*"},
	}
	for _, opt := range opts {
		opt(cfg)
	}

	// Build query
	// 쿼리 빌드
	query := buildSelectQuery(table, condition, cfg)

	// Add args from HAVING clause
	// HAVING 절의 인자 추가
	if cfg.having != nil {
		args = append(args, cfg.having.args...)
	}

	// Execute
	// 실행
	rows, err := tx.tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanRows(rows)
}

// SelectOneWhere for transactions
// 트랜잭션용 SelectOneWhere
func (tx *Tx) SelectOneWhere(ctx context.Context, table string, conditionAndArgs ...interface{}) (map[string]interface{}, error) {
	conditionAndArgs = append(conditionAndArgs, WithLimit(1))

	results, err := tx.SelectWhere(ctx, table, conditionAndArgs...)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no rows found")
	}

	return results[0], nil
}
