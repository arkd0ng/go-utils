package mysql

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// CSVExportOptions represents options for CSV export
// CSVExportOptions는 CSV 내보내기 옵션을 나타냅니다
type CSVExportOptions struct {
	// Include column headers in the first row / 첫 번째 행에 컬럼 헤더 포함
	IncludeHeaders bool

	// Delimiter character (default: comma) / 구분 문자 (기본값: 쉼표)
	Delimiter rune

	// Columns to export (empty means all) / 내보낼 컬럼 (비어 있으면 전체)
	Columns []string

	// WHERE clause filter (without WHERE keyword) / WHERE 절 필터 (WHERE 키워드 제외)
	Where string

	// Arguments for WHERE clause / WHERE 절 인자
	WhereArgs []interface{}

	// ORDER BY clause / ORDER BY 절
	OrderBy string

	// LIMIT number of rows / 행 수 제한
	Limit int

	// NULL value representation / NULL 값 표현
	NullValue string
}

// CSVImportOptions represents options for CSV import
// CSVImportOptions는 CSV 가져오기 옵션을 나타냅니다
type CSVImportOptions struct {
	// First row contains headers / 첫 번째 행에 헤더 포함
	HasHeaders bool

	// Delimiter character (default: comma) / 구분 문자 (기본값: 쉼표)
	Delimiter rune

	// Columns to import (must match CSV column order if no headers) / 가져올 컬럼 (헤더가 없으면 CSV 컬럼 순서와 일치해야 함)
	Columns []string

	// Skip first N rows (after headers if present) / 첫 N개 행 건너뛰기 (헤더가 있는 경우 헤더 다음)
	SkipRows int

	// Batch size for bulk insert / 대량 삽입을 위한 배치 크기
	BatchSize int

	// If true, ignore duplicate key errors / true면 중복 키 에러 무시
	IgnoreDuplicates bool

	// If true, replace existing rows on duplicate keys / true면 중복 키에서 기존 행 교체
	ReplaceOnDuplicate bool

	// NULL value representation in CSV / CSV의 NULL 값 표현
	NullValue string
}

// DefaultCSVExportOptions returns default export options
// DefaultCSVExportOptions는 기본 내보내기 옵션을 반환합니다
func DefaultCSVExportOptions() CSVExportOptions {
	return CSVExportOptions{
		IncludeHeaders: true,
		Delimiter:      ',',
		NullValue:      "NULL",
	}
}

// DefaultCSVImportOptions returns default import options
// DefaultCSVImportOptions는 기본 가져오기 옵션을 반환합니다
func DefaultCSVImportOptions() CSVImportOptions {
	return CSVImportOptions{
		HasHeaders:         true,
		Delimiter:          ',',
		BatchSize:          1000,
		IgnoreDuplicates:   false,
		ReplaceOnDuplicate: false,
		NullValue:          "NULL",
	}
}

// ExportTableToCSV exports a table to a CSV file
// ExportTableToCSV는 테이블을 CSV 파일로 내보냅니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	opts := mysql.DefaultCSVExportOptions()
//	opts.Columns = []string{"id", "name", "email"}
//	opts.Where = "created_at >= ?"
//	opts.WhereArgs = []interface{}{"2024-01-01"}
//
//	err := client.ExportTableToCSV(ctx, "users", "/path/to/users.csv", opts)
//	if err != nil {
//	    log.Fatal(err)
//	}
func (c *Client) ExportTableToCSV(ctx context.Context, table string, filePath string, opts CSVExportOptions) error {
	// Build query / 쿼리 구성
	columns := "*"
	if len(opts.Columns) > 0 {
		columns = strings.Join(opts.Columns, ", ")
	}

	query := fmt.Sprintf("SELECT %s FROM %s", columns, table)

	if opts.Where != "" {
		query += " WHERE " + opts.Where
	}

	if opts.OrderBy != "" {
		query += " ORDER BY " + opts.OrderBy
	}

	if opts.Limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", opts.Limit)
	}

	// Execute query / 쿼리 실행
	rows, err := c.Query(ctx, query, opts.WhereArgs...)
	if err != nil {
		return fmt.Errorf("failed to query table: %w", err)
	}
	defer rows.Close()

	// Get column names / 컬럼 이름 가져오기
	columnNames, err := rows.Columns()
	if err != nil {
		return fmt.Errorf("failed to get column names: %w", err)
	}

	// Create CSV file / CSV 파일 생성
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %w", err)
	}
	defer file.Close()

	// Create CSV writer / CSV writer 생성
	writer := csv.NewWriter(file)
	if opts.Delimiter != 0 && opts.Delimiter != ',' {
		writer.Comma = opts.Delimiter
	}
	defer writer.Flush()

	// Write headers / 헤더 작성
	if opts.IncludeHeaders {
		if err := writer.Write(columnNames); err != nil {
			return fmt.Errorf("failed to write headers: %w", err)
		}
	}

	// Prepare scan destinations / 스캔 대상 준비
	values := make([]interface{}, len(columnNames))
	valuePtrs := make([]interface{}, len(columnNames))
	for i := range values {
		valuePtrs[i] = &values[i]
	}

	rowCount := 0

	// Write rows / 행 작성
	for rows.Next() {
		if err := rows.Scan(valuePtrs...); err != nil {
			return fmt.Errorf("failed to scan row: %w", err)
		}

		// Convert values to strings / 값을 문자열로 변환
		record := make([]string, len(columnNames))
		for i, val := range values {
			if val == nil {
				record[i] = opts.NullValue
			} else {
				record[i] = fmt.Sprintf("%v", val)
			}
		}

		if err := writer.Write(record); err != nil {
			return fmt.Errorf("failed to write row: %w", err)
		}

		rowCount++
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("error iterating rows: %w", err)
	}

	if c.config.logger != nil {
		c.config.logger.Info("Table exported to CSV",
			"table", table,
			"file", filePath,
			"rows", rowCount)
	}

	return nil
}

// ImportFromCSV imports data from a CSV file into a table
// ImportFromCSV는 CSV 파일에서 테이블로 데이터를 가져옵니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	opts := mysql.DefaultCSVImportOptions()
//	opts.Columns = []string{"id", "name", "email"}
//	opts.BatchSize = 500
//	opts.IgnoreDuplicates = true
//
//	err := client.ImportFromCSV(ctx, "users", "/path/to/users.csv", opts)
//	if err != nil {
//	    log.Fatal(err)
//	}
func (c *Client) ImportFromCSV(ctx context.Context, table string, filePath string, opts CSVImportOptions) error {
	// Open CSV file / CSV 파일 열기
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	// Create CSV reader / CSV reader 생성
	reader := csv.NewReader(file)
	if opts.Delimiter != 0 && opts.Delimiter != ',' {
		reader.Comma = opts.Delimiter
	}
	reader.FieldsPerRecord = -1 // Allow variable number of fields / 가변 필드 수 허용

	// Read all records / 모든 레코드 읽기
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV file: %w", err)
	}

	if len(records) == 0 {
		return fmt.Errorf("CSV file is empty")
	}

	startRow := 0
	columns := opts.Columns

	// Handle headers / 헤더 처리
	if opts.HasHeaders {
		if len(columns) == 0 {
			columns = records[0]
		}
		startRow = 1
	} else if len(columns) == 0 {
		return fmt.Errorf("columns must be specified when CSV has no headers")
	}

	// Skip rows if needed / 필요한 경우 행 건너뛰기
	startRow += opts.SkipRows

	if startRow >= len(records) {
		return fmt.Errorf("no data rows to import after skipping rows")
	}

	// Determine batch size / 배치 크기 결정
	batchSize := opts.BatchSize
	if batchSize <= 0 {
		batchSize = 1000
	}

	// Build insert query prefix / 삽입 쿼리 접두사 구성
	var insertPrefix string
	if opts.ReplaceOnDuplicate {
		insertPrefix = fmt.Sprintf("REPLACE INTO %s (%s) VALUES ",
			table, strings.Join(columns, ", "))
	} else if opts.IgnoreDuplicates {
		insertPrefix = fmt.Sprintf("INSERT IGNORE INTO %s (%s) VALUES ",
			table, strings.Join(columns, ", "))
	} else {
		insertPrefix = fmt.Sprintf("INSERT INTO %s (%s) VALUES ",
			table, strings.Join(columns, ", "))
	}

	totalRows := 0
	batch := make([][]string, 0, batchSize)

	// Process records in batches / 배치로 레코드 처리
	for i := startRow; i < len(records); i++ {
		record := records[i]

		// Skip empty rows / 빈 행 건너뛰기
		if len(record) == 0 || (len(record) == 1 && record[0] == "") {
			continue
		}

		// Validate record length / 레코드 길이 검증
		if len(record) != len(columns) {
			if c.config.logger != nil {
				c.config.logger.Warn("Skipping row with mismatched column count",
					"row", i+1,
					"expected", len(columns),
					"got", len(record))
			}
			continue
		}

		batch = append(batch, record)

		// Execute batch when full / 배치가 가득 차면 실행
		if len(batch) >= batchSize {
			if err := c.executeBatchInsert(ctx, insertPrefix, columns, batch, opts.NullValue); err != nil {
				return fmt.Errorf("failed to insert batch at row %d: %w", i+1, err)
			}
			totalRows += len(batch)
			batch = batch[:0] // Reset batch / 배치 초기화
		}
	}

	// Insert remaining records / 남은 레코드 삽입
	if len(batch) > 0 {
		if err := c.executeBatchInsert(ctx, insertPrefix, columns, batch, opts.NullValue); err != nil {
			return fmt.Errorf("failed to insert final batch: %w", err)
		}
		totalRows += len(batch)
	}

	if c.config.logger != nil {
		c.config.logger.Info("CSV imported",
			"table", table,
			"file", filePath,
			"rows", totalRows)
	}

	return nil
}

// executeBatchInsert executes a batch insert operation
// executeBatchInsert는 배치 삽입 작업을 실행합니다
func (c *Client) executeBatchInsert(ctx context.Context, insertPrefix string, columns []string,
	batch [][]string, nullValue string) error {

	// Build values part of query / 쿼리의 값 부분 구성
	valuePlaceholders := make([]string, len(batch))
	args := make([]interface{}, 0, len(batch)*len(columns))

	for i, record := range batch {
		placeholders := make([]string, len(columns))
		for j, value := range record {
			placeholders[j] = "?"
			// Convert NULL values / NULL 값 변환
			if value == nullValue {
				args = append(args, nil)
			} else {
				args = append(args, value)
			}
		}
		valuePlaceholders[i] = "(" + strings.Join(placeholders, ", ") + ")"
	}

	query := insertPrefix + strings.Join(valuePlaceholders, ", ")

	// Execute query / 쿼리 실행
	_, err := c.Exec(ctx, query, args...)
	return err
}

// ExportQueryToCSV exports the result of a query to a CSV file
// ExportQueryToCSV는 쿼리 결과를 CSV 파일로 내보냅니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	query := `
//	    SELECT u.id, u.name, COUNT(o.id) as order_count
//	    FROM users u
//	    LEFT JOIN orders o ON u.id = o.user_id
//	    GROUP BY u.id, u.name
//	    HAVING order_count > 10
//	`
//	opts := mysql.DefaultCSVExportOptions()
//	err := client.ExportQueryToCSV(ctx, query, nil, "/path/to/report.csv", opts)
//	if err != nil {
//	    log.Fatal(err)
//	}
func (c *Client) ExportQueryToCSV(ctx context.Context, query string, args []interface{},
	filePath string, opts CSVExportOptions) error {

	// Execute query / 쿼리 실행
	rows, err := c.Query(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	// Get column names / 컬럼 이름 가져오기
	columnNames, err := rows.Columns()
	if err != nil {
		return fmt.Errorf("failed to get column names: %w", err)
	}

	// Create CSV file / CSV 파일 생성
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %w", err)
	}
	defer file.Close()

	// Create CSV writer / CSV writer 생성
	writer := csv.NewWriter(file)
	if opts.Delimiter != 0 && opts.Delimiter != ',' {
		writer.Comma = opts.Delimiter
	}
	defer writer.Flush()

	// Write headers / 헤더 작성
	if opts.IncludeHeaders {
		if err := writer.Write(columnNames); err != nil {
			return fmt.Errorf("failed to write headers: %w", err)
		}
	}

	// Prepare scan destinations / 스캔 대상 준비
	values := make([]interface{}, len(columnNames))
	valuePtrs := make([]interface{}, len(columnNames))
	for i := range values {
		valuePtrs[i] = &values[i]
	}

	rowCount := 0

	// Write rows / 행 작성
	for rows.Next() {
		if err := rows.Scan(valuePtrs...); err != nil {
			return fmt.Errorf("failed to scan row: %w", err)
		}

		// Convert values to strings / 값을 문자열로 변환
		record := make([]string, len(columnNames))
		for i, val := range values {
			if val == nil {
				record[i] = opts.NullValue
			} else {
				// Handle different types / 다양한 타입 처리
				switch v := val.(type) {
				case []byte:
					record[i] = string(v)
				case int64:
					record[i] = strconv.FormatInt(v, 10)
				case float64:
					record[i] = strconv.FormatFloat(v, 'f', -1, 64)
				case bool:
					record[i] = strconv.FormatBool(v)
				default:
					record[i] = fmt.Sprintf("%v", val)
				}
			}
		}

		if err := writer.Write(record); err != nil {
			return fmt.Errorf("failed to write row: %w", err)
		}

		rowCount++
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("error iterating rows: %w", err)
	}

	if c.config.logger != nil {
		c.config.logger.Info("Query result exported to CSV",
			"file", filePath,
			"rows", rowCount)
	}

	return nil
}
