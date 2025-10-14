package mysql

import (
	"context"
	"fmt"
	"math"
)

// PaginationResult represents the result of a paginated query
// PaginationResult는 페이지네이션 쿼리의 결과를 나타냅니다
type PaginationResult struct {
	Data       []map[string]interface{} // Current page data / 현재 페이지 데이터
	TotalRows  int64                    // Total number of rows / 전체 행 수
	TotalPages int                      // Total number of pages / 전체 페이지 수
	Page       int                      // Current page number (1-indexed) / 현재 페이지 번호 (1부터 시작)
	PageSize   int                      // Number of rows per page / 페이지당 행 수
	HasNext    bool                     // Whether there is a next page / 다음 페이지 존재 여부
	HasPrev    bool                     // Whether there is a previous page / 이전 페이지 존재 여부
}

// Paginate performs paginated query on a table
// Paginate는 테이블에 대해 페이지네이션 쿼리를 수행합니다
//
// Example / 예제:
//
//	// Get page 1 with 10 items per page
//	// 페이지 1을 페이지당 10개 항목으로 가져오기
//	result, err := db.Paginate(ctx, "users", 1, 10)
//
//	// With WHERE condition and ORDER BY
//	// WHERE 조건 및 ORDER BY와 함께
//	result, err := db.Paginate(ctx, "users", 2, 20,
//	    mysql.WithColumns("id", "name", "email"),
//	    mysql.WithWhere("age > ?", 18),
//	    mysql.WithOrderBy("created_at DESC"))
func (c *Client) Paginate(ctx context.Context, table string, page, pageSize int, conditionAndArgs ...interface{}) (*PaginationResult, error) {
	if page < 1 {
		return nil, fmt.Errorf("page must be >= 1")
	}
	if pageSize < 1 {
		return nil, fmt.Errorf("pageSize must be >= 1")
	}

	// Parse condition and options / 조건 및 옵션 파싱
	var condition string
	var args []interface{}
	var opts []SelectOption

	// Extract condition and arguments / 조건 및 인자 추출
	if len(conditionAndArgs) > 0 {
		// Check if first argument is a string (condition) / 첫 번째 인자가 문자열(조건)인지 확인
		if cond, ok := conditionAndArgs[0].(string); ok {
			condition = cond
			// Find where SelectOptions start / SelectOption이 시작하는 위치 찾기
			i := 1
			for i < len(conditionAndArgs) {
				if opt, ok := conditionAndArgs[i].(SelectOption); ok {
					opts = append(opts, opt)
					i++
				} else {
					// This is an argument for the condition / 조건의 인자
					args = append(args, conditionAndArgs[i])
					i++
				}
			}
		} else {
			// No condition, only options / 조건 없이 옵션만
			for _, item := range conditionAndArgs {
				if opt, ok := item.(SelectOption); ok {
					opts = append(opts, opt)
				}
			}
		}
	}

	// Count total rows / 전체 행 수 계산
	var totalRows int64
	var err error
	if condition != "" {
		countArgs := make([]interface{}, 0, 1+len(args))
		countArgs = append(countArgs, condition)
		countArgs = append(countArgs, args...)
		totalRows, err = c.CountContext(ctx, table, countArgs...)
	} else {
		totalRows, err = c.CountContext(ctx, table)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to count rows: %w", err)
	}

	// Calculate pagination info / 페이지네이션 정보 계산
	totalPages := int(math.Ceil(float64(totalRows) / float64(pageSize)))
	offset := (page - 1) * pageSize

	// Add LIMIT and OFFSET to options / 옵션에 LIMIT 및 OFFSET 추가
	opts = append(opts, WithLimit(pageSize), WithOffset(offset))

	// Query data / 데이터 쿼리
	var data []map[string]interface{}
	if condition != "" {
		// Combine condition, args, and opts / 조건, 인자, 옵션 결합
		queryArgs := make([]interface{}, 0, 1+len(args)+len(opts))
		queryArgs = append(queryArgs, condition)
		queryArgs = append(queryArgs, args...)
		for _, opt := range opts {
			queryArgs = append(queryArgs, opt)
		}
		data, err = c.SelectWhere(ctx, table, queryArgs...)
	} else {
		// Only opts / 옵션만
		queryArgs := make([]interface{}, 0, 1+len(opts))
		queryArgs = append(queryArgs, "")
		for _, opt := range opts {
			queryArgs = append(queryArgs, opt)
		}
		data, err = c.SelectWhere(ctx, table, queryArgs...)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query data: %w", err)
	}

	// Build result / 결과 빌드
	result := &PaginationResult{
		Data:       data,
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Page:       page,
		PageSize:   pageSize,
		HasNext:    page < totalPages,
		HasPrev:    page > 1,
	}

	return result, nil
}

// PaginateQuery performs paginated query using a custom query
// PaginateQuery는 사용자 정의 쿼리를 사용하여 페이지네이션 쿼리를 수행합니다
//
// Example / 예제:
//
//	baseQuery := "SELECT u.*, COUNT(o.id) as order_count FROM users u LEFT JOIN orders o ON u.id = o.user_id GROUP BY u.id"
//	countQuery := "SELECT COUNT(*) FROM users"
//	result, err := db.PaginateQuery(ctx, baseQuery, countQuery, 1, 10)
func (c *Client) PaginateQuery(ctx context.Context, baseQuery, countQuery string, page, pageSize int, args ...interface{}) (*PaginationResult, error) {
	if page < 1 {
		return nil, fmt.Errorf("page must be >= 1")
	}
	if pageSize < 1 {
		return nil, fmt.Errorf("pageSize must be >= 1")
	}

	// Count total rows / 전체 행 수 계산
	row := c.QueryRow(ctx, countQuery, args...)
	var totalRows int64
	if err := row.Scan(&totalRows); err != nil {
		return nil, fmt.Errorf("failed to count rows: %w", err)
	}

	// Calculate pagination info / 페이지네이션 정보 계산
	totalPages := int(math.Ceil(float64(totalRows) / float64(pageSize)))
	offset := (page - 1) * pageSize

	// Add LIMIT and OFFSET to query / 쿼리에 LIMIT 및 OFFSET 추가
	paginatedQuery := fmt.Sprintf("%s LIMIT %d OFFSET %d", baseQuery, pageSize, offset)

	// Query data / 데이터 쿼리
	rows, err := c.Query(ctx, paginatedQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query data: %w", err)
	}

	// Scan results / 결과 스캔
	data, err := scanRows(rows)
	if err != nil {
		return nil, fmt.Errorf("failed to scan rows: %w", err)
	}

	// Build result / 결과 빌드
	result := &PaginationResult{
		Data:       data,
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Page:       page,
		PageSize:   pageSize,
		HasNext:    page < totalPages,
		HasPrev:    page > 1,
	}

	return result, nil
}

// GetPage returns a specific page number (1-indexed)
// GetPage는 특정 페이지 번호를 반환합니다 (1부터 시작)
func (pr *PaginationResult) GetPage() int {
	return pr.Page
}

// GetTotalPages returns the total number of pages
// GetTotalPages는 전체 페이지 수를 반환합니다
func (pr *PaginationResult) GetTotalPages() int {
	return pr.TotalPages
}

// GetTotalRows returns the total number of rows
// GetTotalRows는 전체 행 수를 반환합니다
func (pr *PaginationResult) GetTotalRows() int64 {
	return pr.TotalRows
}

// IsFirstPage returns true if this is the first page
// IsFirstPage는 첫 번째 페이지인 경우 true를 반환합니다
func (pr *PaginationResult) IsFirstPage() bool {
	return pr.Page == 1
}

// IsLastPage returns true if this is the last page
// IsLastPage는 마지막 페이지인 경우 true를 반환합니다
func (pr *PaginationResult) IsLastPage() bool {
	return pr.Page == pr.TotalPages || pr.TotalPages == 0
}

// NextPage returns the next page number, or 0 if there is no next page
// NextPage는 다음 페이지 번호를 반환하거나, 다음 페이지가 없으면 0을 반환합니다
func (pr *PaginationResult) NextPage() int {
	if pr.HasNext {
		return pr.Page + 1
	}
	return 0
}

// PrevPage returns the previous page number, or 0 if there is no previous page
// PrevPage는 이전 페이지 번호를 반환하거나, 이전 페이지가 없으면 0을 반환합니다
func (pr *PaginationResult) PrevPage() int {
	if pr.HasPrev {
		return pr.Page - 1
	}
	return 0
}
