package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

// QueryBuilder provides a fluent API for building complex SQL queries
// QueryBuilder는 복잡한 SQL 쿼리를 빌드하기 위한 Fluent API를 제공합니다
type QueryBuilder struct {
	client         *Client        // Client reference / 클라이언트 참조
	tx             *Tx            // Transaction reference (optional) / 트랜잭션 참조 (선택)
	columns        []string       // SELECT columns / SELECT 컬럼
	table          string         // FROM table / FROM 테이블
	joins          []joinClause   // JOIN clauses / JOIN 절
	whereClauses   []whereClause  // WHERE conditions / WHERE 조건
	groupByCols    []string       // GROUP BY columns / GROUP BY 컬럼
	havingClauses  []whereClause  // HAVING conditions / HAVING 조건
	orderBy        string         // ORDER BY clause / ORDER BY 절
	limitNum       *int           // LIMIT value / LIMIT 값
	offsetNum      *int           // OFFSET value / OFFSET 값
	args           []interface{}  // Query arguments / 쿼리 인자
}

// joinClause represents a JOIN clause
// joinClause는 JOIN 절을 나타냅니다
type joinClause struct {
	joinType  string // INNER, LEFT, RIGHT / 조인 타입
	table     string // Table to join / 조인할 테이블
	condition string // ON condition / ON 조건
}

// whereClause represents a WHERE or HAVING clause
// whereClause는 WHERE 또는 HAVING 절을 나타냅니다
type whereClause struct {
	condition string        // SQL condition / SQL 조건
	args      []interface{} // Arguments for placeholders / 플레이스홀더 인자
}

// Select initiates a query builder with specified columns
// Select는 지정된 컬럼으로 쿼리 빌더를 시작합니다
func (c *Client) Select(cols ...string) *QueryBuilder {
	if len(cols) == 0 {
		cols = []string{"*"}
	}
	return &QueryBuilder{
		client:  c,
		columns: cols,
		args:    make([]interface{}, 0),
	}
}

// Select initiates a query builder within a transaction
// Select는 트랜잭션 내에서 쿼리 빌더를 시작합니다
func (tx *Tx) Select(cols ...string) *QueryBuilder {
	if len(cols) == 0 {
		cols = []string{"*"}
	}
	return &QueryBuilder{
		tx:      tx,
		columns: cols,
		args:    make([]interface{}, 0),
	}
}

// From specifies the table to query from
// From은 쿼리할 테이블을 지정합니다
func (qb *QueryBuilder) From(table string) *QueryBuilder {
	qb.table = table
	return qb
}

// Join adds an INNER JOIN clause
// Join은 INNER JOIN 절을 추가합니다
func (qb *QueryBuilder) Join(table, condition string) *QueryBuilder {
	qb.joins = append(qb.joins, joinClause{
		joinType:  "INNER JOIN",
		table:     table,
		condition: condition,
	})
	return qb
}

// InnerJoin adds an INNER JOIN clause (alias for Join)
// InnerJoin은 INNER JOIN 절을 추가합니다 (Join의 별칭)
func (qb *QueryBuilder) InnerJoin(table, condition string) *QueryBuilder {
	return qb.Join(table, condition)
}

// LeftJoin adds a LEFT JOIN clause
// LeftJoin은 LEFT JOIN 절을 추가합니다
func (qb *QueryBuilder) LeftJoin(table, condition string) *QueryBuilder {
	qb.joins = append(qb.joins, joinClause{
		joinType:  "LEFT JOIN",
		table:     table,
		condition: condition,
	})
	return qb
}

// RightJoin adds a RIGHT JOIN clause
// RightJoin은 RIGHT JOIN 절을 추가합니다
func (qb *QueryBuilder) RightJoin(table, condition string) *QueryBuilder {
	qb.joins = append(qb.joins, joinClause{
		joinType:  "RIGHT JOIN",
		table:     table,
		condition: condition,
	})
	return qb
}

// Where adds a WHERE condition
// Where는 WHERE 조건을 추가합니다
func (qb *QueryBuilder) Where(condition string, args ...interface{}) *QueryBuilder {
	qb.whereClauses = append(qb.whereClauses, whereClause{
		condition: condition,
		args:      args,
	})
	return qb
}

// GroupBy adds GROUP BY columns
// GroupBy는 GROUP BY 컬럼을 추가합니다
func (qb *QueryBuilder) GroupBy(cols ...string) *QueryBuilder {
	qb.groupByCols = append(qb.groupByCols, cols...)
	return qb
}

// Having adds a HAVING condition (used with GROUP BY)
// Having은 HAVING 조건을 추가합니다 (GROUP BY와 함께 사용)
func (qb *QueryBuilder) Having(condition string, args ...interface{}) *QueryBuilder {
	qb.havingClauses = append(qb.havingClauses, whereClause{
		condition: condition,
		args:      args,
	})
	return qb
}

// OrderBy adds an ORDER BY clause
// OrderBy는 ORDER BY 절을 추가합니다
func (qb *QueryBuilder) OrderBy(order string) *QueryBuilder {
	qb.orderBy = order
	return qb
}

// Limit sets the LIMIT clause
// Limit은 LIMIT 절을 설정합니다
func (qb *QueryBuilder) Limit(n int) *QueryBuilder {
	qb.limitNum = &n
	return qb
}

// Offset sets the OFFSET clause
// Offset은 OFFSET 절을 설정합니다
func (qb *QueryBuilder) Offset(n int) *QueryBuilder {
	qb.offsetNum = &n
	return qb
}

// buildQuery constructs the final SQL query and argument list
// buildQuery는 최종 SQL 쿼리와 인자 목록을 생성합니다
func (qb *QueryBuilder) buildQuery() (string, []interface{}) {
	var parts []string
	var args []interface{}

	// SELECT clause
	// SELECT 절
	parts = append(parts, "SELECT "+strings.Join(qb.columns, ", "))

	// FROM clause
	// FROM 절
	if qb.table != "" {
		parts = append(parts, "FROM "+qb.table)
	}

	// JOIN clauses
	// JOIN 절들
	for _, join := range qb.joins {
		parts = append(parts, fmt.Sprintf("%s %s ON %s", join.joinType, join.table, join.condition))
	}

	// WHERE clause
	// WHERE 절
	if len(qb.whereClauses) > 0 {
		conditions := make([]string, 0, len(qb.whereClauses))
		for _, wc := range qb.whereClauses {
			conditions = append(conditions, "("+wc.condition+")")
			args = append(args, wc.args...)
		}
		parts = append(parts, "WHERE "+strings.Join(conditions, " AND "))
	}

	// GROUP BY clause
	// GROUP BY 절
	if len(qb.groupByCols) > 0 {
		parts = append(parts, "GROUP BY "+strings.Join(qb.groupByCols, ", "))
	}

	// HAVING clause
	// HAVING 절
	if len(qb.havingClauses) > 0 {
		conditions := make([]string, 0, len(qb.havingClauses))
		for _, hc := range qb.havingClauses {
			conditions = append(conditions, "("+hc.condition+")")
			args = append(args, hc.args...)
		}
		parts = append(parts, "HAVING "+strings.Join(conditions, " AND "))
	}

	// ORDER BY clause
	// ORDER BY 절
	if qb.orderBy != "" {
		parts = append(parts, "ORDER BY "+qb.orderBy)
	}

	// LIMIT clause
	// LIMIT 절
	if qb.limitNum != nil {
		parts = append(parts, fmt.Sprintf("LIMIT %d", *qb.limitNum))
	}

	// OFFSET clause
	// OFFSET 절
	if qb.offsetNum != nil {
		parts = append(parts, fmt.Sprintf("OFFSET %d", *qb.offsetNum))
	}

	return strings.Join(parts, " "), args
}

// All executes the query and returns all matching rows
// All은 쿼리를 실행하고 일치하는 모든 행을 반환합니다
func (qb *QueryBuilder) All(ctx context.Context) ([]map[string]interface{}, error) {
	query, args := qb.buildQuery()

	// Use transaction if available, otherwise use client
	// 트랜잭션이 있으면 트랜잭션을 사용하고 없으면 클라이언트를 사용합니다
	if qb.tx != nil {
		return qb.executeQueryTx(ctx, query, args)
	}
	return qb.executeQueryClient(ctx, query, args)
}

// One executes the query and returns a single row
// One은 쿼리를 실행하고 단일 행을 반환합니다
func (qb *QueryBuilder) One(ctx context.Context) (map[string]interface{}, error) {
	// Add LIMIT 1 if not already set
	// 이미 설정되지 않았으면 LIMIT 1 추가
	if qb.limitNum == nil {
		limit := 1
		qb.limitNum = &limit
	}

	results, err := qb.All(ctx)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no rows found")
	}

	return results[0], nil
}

// executeQueryClient executes the query using the client
// executeQueryClient는 클라이언트를 사용하여 쿼리를 실행합니다
func (qb *QueryBuilder) executeQueryClient(ctx context.Context, query string, args []interface{}) ([]map[string]interface{}, error) {
	var rows *sql.Rows
	err := qb.client.executeWithRetry(ctx, func() error {
		db := qb.client.getCurrentConnection()
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

// executeQueryTx executes the query using the transaction
// executeQueryTx는 트랜잭션을 사용하여 쿼리를 실행합니다
func (qb *QueryBuilder) executeQueryTx(ctx context.Context, query string, args []interface{}) ([]map[string]interface{}, error) {
	rows, err := qb.tx.tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanRows(rows)
}
