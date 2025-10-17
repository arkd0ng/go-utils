package mysql

import (
	"database/sql"
	"fmt"
	"time"
)

// scanRows scans all rows from sql.Rows into a slice of maps
// scanRows는 sql.Rows의 모든 행을 map 슬라이스로 스캔합니다
func scanRows(rows *sql.Rows) ([]map[string]interface{}, error) {
	defer rows.Close()

	// Get column names
	// 컬럼 이름 가져오기
	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("failed to get columns: %w", err)
	}

	// Get column types
	// 컬럼 타입 가져오기
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, fmt.Errorf("failed to get column types: %w", err)
	}

	var results []map[string]interface{}

	for rows.Next() {
		// Create a slice of interface{} to hold each column value
		// 각 컬럼 값을 보유할 interface{} 슬라이스 생성
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		// Scan the row
		// 행 스캔
		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		// Create a map for this row
		// 이 행에 대한 map 생성
		row := make(map[string]interface{})
		for i, col := range columns {
			val := values[i]

			// Convert types
			// 타입 변환
			row[col] = convertValue(val, columnTypes[i])
		}

		results = append(results, row)
	}

	// Check for errors from iterating over rows
	// 행 반복 중 에러 확인
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return results, nil
}

// scanRow scans a single row from sql.Row into a map
// scanRow는 sql.Row의 단일 행을 map으로 스캔합니다
func scanRow(rows *sql.Rows) (map[string]interface{}, error) {
	defer rows.Close()

	// Get column names
	// 컬럼 이름 가져오기
	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("failed to get columns: %w", err)
	}

	// Get column types
	// 컬럼 타입 가져오기
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, fmt.Errorf("failed to get column types: %w", err)
	}

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("error getting row: %w", err)
		}
		return nil, ErrNoRows
	}

	// Create a slice of interface{} to hold each column value
	// 각 컬럼 값을 보유할 interface{} 슬라이스 생성
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for i := range values {
		valuePtrs[i] = &values[i]
	}

	// Scan the row
	// 행 스캔
	if err := rows.Scan(valuePtrs...); err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	// Create a map for this row
	// 이 행에 대한 map 생성
	row := make(map[string]interface{})
	for i, col := range columns {
		val := values[i]
		row[col] = convertValue(val, columnTypes[i])
	}

	return row, nil
}

// convertValue converts database values to appropriate Go types
// convertValue는 데이터베이스 값을 적절한 Go 타입으로 변환합니다
func convertValue(val interface{}, colType *sql.ColumnType) interface{} {
	if val == nil {
		return nil
	}

	// Convert []byte to string for text types
	// 텍스트 타입에 대해 []byte를 string으로 변환
	switch v := val.(type) {
	case []byte:
		// Check if it's a text/varchar type
		// text/varchar 타입인지 확인
		dbType := colType.DatabaseTypeName()
		if isTextType(dbType) {
			return string(v)
		}
		return v
	case time.Time:
		return v
	case int64:
		return v
	case float64:
		return v
	case bool:
		return v
	case string:
		return v
	default:
		return v
	}
}

// isTextType checks if a database type is a text type
// isTextType은 데이터베이스 타입이 텍스트 타입인지 확인합니다
func isTextType(dbType string) bool {
	textTypes := []string{
		"VARCHAR", "CHAR", "TEXT", "TINYTEXT", "MEDIUMTEXT", "LONGTEXT",
		"ENUM", "SET", "JSON",
	}

	for _, t := range textTypes {
		if dbType == t {
			return true
		}
	}
	return false
}

// scanCount scans a COUNT(*) result
// scanCount는 COUNT(*) 결과를 스캔합니다
func scanCount(rows *sql.Rows) (int64, error) {
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return 0, fmt.Errorf("error getting row: %w", err)
		}
		return 0, ErrNoRows
	}

	var count int64
	if err := rows.Scan(&count); err != nil {
		return 0, fmt.Errorf("failed to scan count: %w", err)
	}

	return count, nil
}
