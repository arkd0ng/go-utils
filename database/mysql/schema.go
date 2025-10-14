package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

// ColumnInfo represents information about a table column
// ColumnInfo는 테이블 컬럼에 대한 정보를 나타냅니다
type ColumnInfo struct {
	Name     string         // Column name / 컬럼 이름
	Type     string         // Column type / 컬럼 타입
	Nullable bool           // Whether NULL is allowed / NULL 허용 여부
	Default  sql.NullString // Default value / 기본값
	Key      string         // Key type (PRI, UNI, MUL) / 키 타입
	Extra    string         // Extra information (auto_increment, etc.) / 추가 정보
}

// IndexInfo represents information about a table index
// IndexInfo는 테이블 인덱스에 대한 정보를 나타냅니다
type IndexInfo struct {
	Name      string   // Index name / 인덱스 이름
	Columns   []string // Indexed columns / 인덱싱된 컬럼
	Unique    bool     // Whether index is unique / 유니크 인덱스 여부
	IndexType string   // Index type (BTREE, HASH, etc.) / 인덱스 타입
}

// TableInfo represents information about a table
// TableInfo는 테이블에 대한 정보를 나타냅니다
type TableInfo struct {
	Name    string // Table name / 테이블 이름
	Engine  string // Storage engine / 스토리지 엔진
	Rows    int64  // Approximate row count / 대략적인 행 수
	Comment string // Table comment / 테이블 주석
}

// GetTables returns a list of all tables in the current database
// GetTables는 현재 데이터베이스의 모든 테이블 목록을 반환합니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	tables, err := client.GetTables(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, table := range tables {
//	    fmt.Printf("Table: %s (Engine: %s, Rows: %d)\n",
//	        table.Name, table.Engine, table.Rows)
//	}
func (c *Client) GetTables(ctx context.Context) ([]TableInfo, error) {
	query := `
		SELECT
			TABLE_NAME,
			ENGINE,
			TABLE_ROWS,
			TABLE_COMMENT
		FROM information_schema.TABLES
		WHERE TABLE_SCHEMA = DATABASE()
		ORDER BY TABLE_NAME
	`

	rows, err := c.Query(ctx, query)
	if err != nil {
		return nil, c.wrapError("GetTables", query, nil, err, 0)
	}
	defer rows.Close()

	var tables []TableInfo
	for rows.Next() {
		var table TableInfo
		var engine, comment sql.NullString
		var rowCount sql.NullInt64

		if err := rows.Scan(&table.Name, &engine, &rowCount, &comment); err != nil {
			return nil, fmt.Errorf("failed to scan table info: %w", err)
		}

		table.Engine = engine.String
		table.Rows = rowCount.Int64
		table.Comment = comment.String

		tables = append(tables, table)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating table rows: %w", err)
	}

	return tables, nil
}

// GetColumns returns information about all columns in a table
// GetColumns는 테이블의 모든 컬럼에 대한 정보를 반환합니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	columns, err := client.GetColumns(ctx, "users")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, col := range columns {
//	    nullable := "NOT NULL"
//	    if col.Nullable {
//	        nullable = "NULL"
//	    }
//	    fmt.Printf("%s %s %s\n", col.Name, col.Type, nullable)
//	    if col.Default.Valid {
//	        fmt.Printf("  DEFAULT: %s\n", col.Default.String)
//	    }
//	}
func (c *Client) GetColumns(ctx context.Context, table string) ([]ColumnInfo, error) {
	query := fmt.Sprintf("DESCRIBE %s", table)

	rows, err := c.Query(ctx, query)
	if err != nil {
		return nil, c.wrapError("GetColumns", query, []interface{}{table}, err, 0)
	}
	defer rows.Close()

	var columns []ColumnInfo
	for rows.Next() {
		var col ColumnInfo
		var nullStr string

		if err := rows.Scan(&col.Name, &col.Type, &nullStr, &col.Key, &col.Default, &col.Extra); err != nil {
			return nil, fmt.Errorf("failed to scan column info: %w", err)
		}

		col.Nullable = (nullStr == "YES")
		columns = append(columns, col)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating column rows: %w", err)
	}

	return columns, nil
}

// GetIndexes returns information about all indexes on a table
// GetIndexes는 테이블의 모든 인덱스에 대한 정보를 반환합니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	indexes, err := client.GetIndexes(ctx, "users")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, idx := range indexes {
//	    uniqueStr := ""
//	    if idx.Unique {
//	        uniqueStr = "UNIQUE"
//	    }
//	    fmt.Printf("%s %s INDEX %s (%s)\n",
//	        uniqueStr, idx.IndexType, idx.Name, strings.Join(idx.Columns, ", "))
//	}
func (c *Client) GetIndexes(ctx context.Context, table string) ([]IndexInfo, error) {
	query := fmt.Sprintf("SHOW INDEX FROM %s", table)

	rows, err := c.Query(ctx, query)
	if err != nil {
		return nil, c.wrapError("GetIndexes", query, []interface{}{table}, err, 0)
	}
	defer rows.Close()

	// Map to group columns by index name / 인덱스 이름별로 컬럼을 그룹화하는 맵
	indexMap := make(map[string]*IndexInfo)

	for rows.Next() {
		var (
			tableName    string
			nonUnique    int
			keyName      string
			seqInIndex   int
			columnName   string
			collation    sql.NullString
			cardinality  sql.NullInt64
			subPart      sql.NullInt64
			packed       sql.NullString
			null         string
			indexType    string
			comment      string
			indexComment string
			visible      string
		)

		err := rows.Scan(
			&tableName,
			&nonUnique,
			&keyName,
			&seqInIndex,
			&columnName,
			&collation,
			&cardinality,
			&subPart,
			&packed,
			&null,
			&indexType,
			&comment,
			&indexComment,
			&visible,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan index info: %w", err)
		}

		// Get or create index info / 인덱스 정보 가져오기 또는 생성
		idx, exists := indexMap[keyName]
		if !exists {
			idx = &IndexInfo{
				Name:      keyName,
				Columns:   make([]string, 0),
				Unique:    nonUnique == 0,
				IndexType: indexType,
			}
			indexMap[keyName] = idx
		}

		idx.Columns = append(idx.Columns, columnName)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating index rows: %w", err)
	}

	// Convert map to slice / 맵을 슬라이스로 변환
	indexes := make([]IndexInfo, 0, len(indexMap))
	for _, idx := range indexMap {
		indexes = append(indexes, *idx)
	}

	return indexes, nil
}

// TableExists checks if a table exists in the database
// TableExists는 데이터베이스에 테이블이 존재하는지 확인합니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	exists, err := client.TableExists(ctx, "users")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	if exists {
//	    fmt.Println("Table 'users' exists")
//	} else {
//	    fmt.Println("Table 'users' does not exist")
//	}
func (c *Client) TableExists(ctx context.Context, table string) (bool, error) {
	query := `
		SELECT COUNT(*)
		FROM information_schema.TABLES
		WHERE TABLE_SCHEMA = DATABASE()
		AND TABLE_NAME = ?
	`

	var count int
	err := c.QueryRow(ctx, query, table).Scan(&count)
	if err != nil {
		return false, c.wrapError("TableExists", query, []interface{}{table}, err, 0)
	}

	return count > 0, nil
}

// GetTableSchema returns the CREATE TABLE statement for a table
// GetTableSchema는 테이블의 CREATE TABLE 문을 반환합니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	schema, err := client.GetTableSchema(ctx, "users")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(schema)
func (c *Client) GetTableSchema(ctx context.Context, table string) (string, error) {
	query := fmt.Sprintf("SHOW CREATE TABLE %s", table)

	var tableName, createStmt string
	err := c.QueryRow(ctx, query).Scan(&tableName, &createStmt)
	if err != nil {
		return "", c.wrapError("GetTableSchema", query, []interface{}{table}, err, 0)
	}

	return createStmt, nil
}

// GetPrimaryKey returns the primary key columns for a table
// GetPrimaryKey는 테이블의 기본 키 컬럼을 반환합니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	pkCols, err := client.GetPrimaryKey(ctx, "users")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	if len(pkCols) > 0 {
//	    fmt.Printf("Primary key: %s\n", strings.Join(pkCols, ", "))
//	} else {
//	    fmt.Println("No primary key defined")
//	}
func (c *Client) GetPrimaryKey(ctx context.Context, table string) ([]string, error) {
	columns, err := c.GetColumns(ctx, table)
	if err != nil {
		return nil, err
	}

	var pkCols []string
	for _, col := range columns {
		if col.Key == "PRI" {
			pkCols = append(pkCols, col.Name)
		}
	}

	return pkCols, nil
}

// GetForeignKeys returns information about foreign keys for a table
// GetForeignKeys는 테이블의 외래 키에 대한 정보를 반환합니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	fks, err := client.GetForeignKeys(ctx, "orders")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, fk := range fks {
//	    fmt.Printf("FK: %s.%s -> %s.%s\n",
//	        fk.TableName, fk.ColumnName,
//	        fk.ReferencedTable, fk.ReferencedColumn)
//	}
func (c *Client) GetForeignKeys(ctx context.Context, table string) ([]ForeignKeyInfo, error) {
	query := `
		SELECT
			CONSTRAINT_NAME,
			COLUMN_NAME,
			REFERENCED_TABLE_NAME,
			REFERENCED_COLUMN_NAME
		FROM information_schema.KEY_COLUMN_USAGE
		WHERE TABLE_SCHEMA = DATABASE()
		AND TABLE_NAME = ?
		AND REFERENCED_TABLE_NAME IS NOT NULL
		ORDER BY CONSTRAINT_NAME, ORDINAL_POSITION
	`

	rows, err := c.Query(ctx, query, table)
	if err != nil {
		return nil, c.wrapError("GetForeignKeys", query, []interface{}{table}, err, 0)
	}
	defer rows.Close()

	var fks []ForeignKeyInfo
	for rows.Next() {
		var fk ForeignKeyInfo
		fk.TableName = table

		if err := rows.Scan(&fk.ConstraintName, &fk.ColumnName, &fk.ReferencedTable, &fk.ReferencedColumn); err != nil {
			return nil, fmt.Errorf("failed to scan foreign key info: %w", err)
		}

		fks = append(fks, fk)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating foreign key rows: %w", err)
	}

	return fks, nil
}

// ForeignKeyInfo represents information about a foreign key
// ForeignKeyInfo는 외래 키에 대한 정보를 나타냅니다
type ForeignKeyInfo struct {
	ConstraintName    string // Foreign key constraint name / 외래 키 제약 조건 이름
	TableName         string // Table name / 테이블 이름
	ColumnName        string // Column name / 컬럼 이름
	ReferencedTable   string // Referenced table name / 참조된 테이블 이름
	ReferencedColumn  string // Referenced column name / 참조된 컬럼 이름
}

// GetTableSize returns the size of a table in bytes
// GetTableSize는 테이블의 크기를 바이트 단위로 반환합니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	size, err := client.GetTableSize(ctx, "users")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Table size: %.2f MB\n", float64(size)/(1024*1024))
func (c *Client) GetTableSize(ctx context.Context, table string) (int64, error) {
	query := `
		SELECT
			DATA_LENGTH + INDEX_LENGTH
		FROM information_schema.TABLES
		WHERE TABLE_SCHEMA = DATABASE()
		AND TABLE_NAME = ?
	`

	var size sql.NullInt64
	err := c.QueryRow(ctx, query, table).Scan(&size)
	if err != nil {
		return 0, c.wrapError("GetTableSize", query, []interface{}{table}, err, 0)
	}

	return size.Int64, nil
}

// GetDatabaseSize returns the total size of the current database in bytes
// GetDatabaseSize는 현재 데이터베이스의 총 크기를 바이트 단위로 반환합니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	size, err := client.GetDatabaseSize(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Database size: %.2f MB\n", float64(size)/(1024*1024))
func (c *Client) GetDatabaseSize(ctx context.Context) (int64, error) {
	query := `
		SELECT
			SUM(DATA_LENGTH + INDEX_LENGTH)
		FROM information_schema.TABLES
		WHERE TABLE_SCHEMA = DATABASE()
	`

	var size sql.NullInt64
	err := c.QueryRow(ctx, query).Scan(&size)
	if err != nil {
		return 0, c.wrapError("GetDatabaseSize", query, nil, err, 0)
	}

	return size.Int64, nil
}

// InspectTable returns comprehensive information about a table
// InspectTable은 테이블에 대한 종합적인 정보를 반환합니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	inspection, err := client.InspectTable(ctx, "users")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	fmt.Printf("Table: %s\n", inspection.Info.Name)
//	fmt.Printf("Columns: %d\n", len(inspection.Columns))
//	fmt.Printf("Indexes: %d\n", len(inspection.Indexes))
//	fmt.Printf("Size: %.2f MB\n", float64(inspection.Size)/(1024*1024))
func (c *Client) InspectTable(ctx context.Context, table string) (*TableInspection, error) {
	inspection := &TableInspection{}

	// Get basic table info / 기본 테이블 정보 가져오기
	tables, err := c.GetTables(ctx)
	if err != nil {
		return nil, err
	}

	for _, t := range tables {
		if t.Name == table {
			inspection.Info = t
			break
		}
	}

	// Get columns / 컬럼 가져오기
	columns, err := c.GetColumns(ctx, table)
	if err != nil {
		return nil, err
	}
	inspection.Columns = columns

	// Get indexes / 인덱스 가져오기
	indexes, err := c.GetIndexes(ctx, table)
	if err != nil {
		return nil, err
	}
	inspection.Indexes = indexes

	// Get primary key / 기본 키 가져오기
	pk, err := c.GetPrimaryKey(ctx, table)
	if err != nil {
		return nil, err
	}
	inspection.PrimaryKey = pk

	// Get foreign keys / 외래 키 가져오기
	fks, err := c.GetForeignKeys(ctx, table)
	if err != nil {
		return nil, err
	}
	inspection.ForeignKeys = fks

	// Get table size / 테이블 크기 가져오기
	size, err := c.GetTableSize(ctx, table)
	if err != nil {
		return nil, err
	}
	inspection.Size = size

	return inspection, nil
}

// TableInspection represents comprehensive table information
// TableInspection은 종합적인 테이블 정보를 나타냅니다
type TableInspection struct {
	Info        TableInfo        // Basic table info / 기본 테이블 정보
	Columns     []ColumnInfo     // Column information / 컬럼 정보
	Indexes     []IndexInfo      // Index information / 인덱스 정보
	PrimaryKey  []string         // Primary key columns / 기본 키 컬럼
	ForeignKeys []ForeignKeyInfo // Foreign key information / 외래 키 정보
	Size        int64            // Table size in bytes / 테이블 크기 (바이트)
}

// String returns a formatted string representation of the inspection
// String은 검사의 포맷된 문자열 표현을 반환합니다
func (ti *TableInspection) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Table: %s\n", ti.Info.Name))
	sb.WriteString(fmt.Sprintf("Engine: %s\n", ti.Info.Engine))
	sb.WriteString(fmt.Sprintf("Rows: %d\n", ti.Info.Rows))
	sb.WriteString(fmt.Sprintf("Size: %.2f MB\n\n", float64(ti.Size)/(1024*1024)))

	sb.WriteString("Columns:\n")
	for _, col := range ti.Columns {
		nullable := "NOT NULL"
		if col.Nullable {
			nullable = "NULL"
		}
		sb.WriteString(fmt.Sprintf("  - %s %s %s", col.Name, col.Type, nullable))
		if col.Key != "" {
			sb.WriteString(fmt.Sprintf(" [%s]", col.Key))
		}
		if col.Extra != "" {
			sb.WriteString(fmt.Sprintf(" %s", col.Extra))
		}
		sb.WriteString("\n")
	}

	if len(ti.Indexes) > 0 {
		sb.WriteString("\nIndexes:\n")
		for _, idx := range ti.Indexes {
			uniqueStr := ""
			if idx.Unique {
				uniqueStr = "UNIQUE "
			}
			sb.WriteString(fmt.Sprintf("  - %s%s (%s) [%s]\n",
				uniqueStr, idx.Name, strings.Join(idx.Columns, ", "), idx.IndexType))
		}
	}

	if len(ti.ForeignKeys) > 0 {
		sb.WriteString("\nForeign Keys:\n")
		for _, fk := range ti.ForeignKeys {
			sb.WriteString(fmt.Sprintf("  - %s.%s -> %s.%s\n",
				fk.TableName, fk.ColumnName, fk.ReferencedTable, fk.ReferencedColumn))
		}
	}

	return sb.String()
}
