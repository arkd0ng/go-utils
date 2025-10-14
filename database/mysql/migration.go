package mysql

import (
	"context"
	"fmt"
	"strings"
)

// CreateTable creates a new table with the given schema
// CreateTable은 주어진 스키마로 새 테이블을 생성합니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	schema := `
//	    id INT AUTO_INCREMENT PRIMARY KEY,
//	    name VARCHAR(255) NOT NULL,
//	    email VARCHAR(255) UNIQUE NOT NULL,
//	    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
//	`
//	err := client.CreateTable(ctx, "users", schema)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// Example with table options / 테이블 옵션 예제:
//
//	schema := `
//	    id BIGINT AUTO_INCREMENT PRIMARY KEY,
//	    data JSON,
//	    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//	    INDEX idx_created (created_at)
//	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci`
//	err := client.CreateTable(ctx, "events", schema)
func (c *Client) CreateTable(ctx context.Context, table string, schema string) error {
	// Check if schema ends with closing parenthesis / 스키마가 닫는 괄호로 끝나는지 확인
	schema = strings.TrimSpace(schema)
	if !strings.HasSuffix(schema, ")") && !strings.Contains(schema, "ENGINE=") {
		// Schema doesn't include table options, add default closing / 테이블 옵션이 포함되지 않은 경우 기본 닫기 추가
		schema = schema + "\n)"
	}

	query := fmt.Sprintf("CREATE TABLE %s (\n%s", table, schema)

	_, err := c.Exec(ctx, query)
	if err != nil {
		return c.wrapError("CreateTable", query, []interface{}{table}, err, 0)
	}

	if c.config.logger != nil {
		c.config.logger.Info("Table created",
			"table", table)
	}

	return nil
}

// CreateTableIfNotExists creates a table only if it doesn't already exist
// CreateTableIfNotExists는 테이블이 존재하지 않는 경우에만 생성합니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	schema := `
//	    id INT AUTO_INCREMENT PRIMARY KEY,
//	    name VARCHAR(255) NOT NULL
//	`
//	err := client.CreateTableIfNotExists(ctx, "users", schema)
//	if err != nil {
//	    log.Fatal(err)
//	}
func (c *Client) CreateTableIfNotExists(ctx context.Context, table string, schema string) error {
	exists, err := c.TableExists(ctx, table)
	if err != nil {
		return err
	}

	if exists {
		if c.config.logger != nil {
			c.config.logger.Debug("Table already exists, skipping creation",
				"table", table)
		}
		return nil
	}

	return c.CreateTable(ctx, table, schema)
}

// DropTable drops a table from the database
// DropTable은 데이터베이스에서 테이블을 삭제합니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	err := client.DropTable(ctx, "old_users", false)
//	if err != nil {
//	    log.Fatal(err)
//	}
func (c *Client) DropTable(ctx context.Context, table string, ifExists bool) error {
	query := fmt.Sprintf("DROP TABLE %s", table)
	if ifExists {
		query = fmt.Sprintf("DROP TABLE IF EXISTS %s", table)
	}

	_, err := c.Exec(ctx, query)
	if err != nil {
		return c.wrapError("DropTable", query, []interface{}{table}, err, 0)
	}

	if c.config.logger != nil {
		c.config.logger.Info("Table dropped",
			"table", table)
	}

	return nil
}

// TruncateTable removes all rows from a table
// TruncateTable은 테이블의 모든 행을 제거합니다
//
// Warning: This operation cannot be rolled back.
// 경고: 이 작업은 롤백할 수 없습니다.
//
// Example / 예제:
//
//	ctx := context.Background()
//	err := client.TruncateTable(ctx, "temp_data")
//	if err != nil {
//	    log.Fatal(err)
//	}
func (c *Client) TruncateTable(ctx context.Context, table string) error {
	query := fmt.Sprintf("TRUNCATE TABLE %s", table)

	_, err := c.Exec(ctx, query)
	if err != nil {
		return c.wrapError("TruncateTable", query, []interface{}{table}, err, 0)
	}

	if c.config.logger != nil {
		c.config.logger.Info("Table truncated",
			"table", table)
	}

	return nil
}

// AddColumn adds a new column to a table
// AddColumn은 테이블에 새 컬럼을 추가합니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	err := client.AddColumn(ctx, "users", "phone", "VARCHAR(20)")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// Example with position / 위치 지정 예제:
//
//	// Add column after 'email' / 'email' 뒤에 컬럼 추가
//	err := client.AddColumn(ctx, "users", "phone", "VARCHAR(20) AFTER email")
//
//	// Add column at the beginning / 처음에 컬럼 추가
//	err := client.AddColumn(ctx, "users", "status", "ENUM('active','inactive') FIRST")
func (c *Client) AddColumn(ctx context.Context, table string, column string, definition string) error {
	query := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s", table, column, definition)

	_, err := c.Exec(ctx, query)
	if err != nil {
		return c.wrapError("AddColumn", query, []interface{}{table, column, definition}, err, 0)
	}

	if c.config.logger != nil {
		c.config.logger.Info("Column added",
			"table", table,
			"column", column)
	}

	return nil
}

// DropColumn removes a column from a table
// DropColumn은 테이블에서 컬럼을 제거합니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	err := client.DropColumn(ctx, "users", "old_field")
//	if err != nil {
//	    log.Fatal(err)
//	}
func (c *Client) DropColumn(ctx context.Context, table string, column string) error {
	query := fmt.Sprintf("ALTER TABLE %s DROP COLUMN %s", table, column)

	_, err := c.Exec(ctx, query)
	if err != nil {
		return c.wrapError("DropColumn", query, []interface{}{table, column}, err, 0)
	}

	if c.config.logger != nil {
		c.config.logger.Info("Column dropped",
			"table", table,
			"column", column)
	}

	return nil
}

// ModifyColumn modifies the definition of an existing column
// ModifyColumn은 기존 컬럼의 정의를 수정합니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	// Change column type / 컬럼 타입 변경
//	err := client.ModifyColumn(ctx, "users", "age", "SMALLINT UNSIGNED")
//	if err != nil {
//	    log.Fatal(err)
//	}
func (c *Client) ModifyColumn(ctx context.Context, table string, column string, definition string) error {
	query := fmt.Sprintf("ALTER TABLE %s MODIFY COLUMN %s %s", table, column, definition)

	_, err := c.Exec(ctx, query)
	if err != nil {
		return c.wrapError("ModifyColumn", query, []interface{}{table, column, definition}, err, 0)
	}

	if c.config.logger != nil {
		c.config.logger.Info("Column modified",
			"table", table,
			"column", column)
	}

	return nil
}

// RenameColumn renames a column
// RenameColumn은 컬럼의 이름을 변경합니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	err := client.RenameColumn(ctx, "users", "old_name", "new_name", "VARCHAR(255)")
//	if err != nil {
//	    log.Fatal(err)
//	}
func (c *Client) RenameColumn(ctx context.Context, table string, oldName string, newName string, definition string) error {
	query := fmt.Sprintf("ALTER TABLE %s CHANGE COLUMN %s %s %s", table, oldName, newName, definition)

	_, err := c.Exec(ctx, query)
	if err != nil {
		return c.wrapError("RenameColumn", query, []interface{}{table, oldName, newName, definition}, err, 0)
	}

	if c.config.logger != nil {
		c.config.logger.Info("Column renamed",
			"table", table,
			"old_name", oldName,
			"new_name", newName)
	}

	return nil
}

// AddIndex adds an index to a table
// AddIndex는 테이블에 인덱스를 추가합니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	// Add simple index / 단순 인덱스 추가
//	err := client.AddIndex(ctx, "users", "idx_email", []string{"email"}, false)
//
//	// Add unique index / 유니크 인덱스 추가
//	err := client.AddIndex(ctx, "users", "idx_username", []string{"username"}, true)
//
//	// Add composite index / 복합 인덱스 추가
//	err := client.AddIndex(ctx, "orders", "idx_user_date",
//	    []string{"user_id", "created_at"}, false)
func (c *Client) AddIndex(ctx context.Context, table string, indexName string, columns []string, unique bool) error {
	indexType := "INDEX"
	if unique {
		indexType = "UNIQUE INDEX"
	}

	query := fmt.Sprintf("ALTER TABLE %s ADD %s %s (%s)",
		table, indexType, indexName, strings.Join(columns, ", "))

	_, err := c.Exec(ctx, query)
	if err != nil {
		return c.wrapError("AddIndex", query, []interface{}{table, indexName, columns}, err, 0)
	}

	if c.config.logger != nil {
		c.config.logger.Info("Index added",
			"table", table,
			"index", indexName,
			"columns", columns)
	}

	return nil
}

// DropIndex removes an index from a table
// DropIndex는 테이블에서 인덱스를 제거합니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	err := client.DropIndex(ctx, "users", "idx_email")
//	if err != nil {
//	    log.Fatal(err)
//	}
func (c *Client) DropIndex(ctx context.Context, table string, indexName string) error {
	query := fmt.Sprintf("ALTER TABLE %s DROP INDEX %s", table, indexName)

	_, err := c.Exec(ctx, query)
	if err != nil {
		return c.wrapError("DropIndex", query, []interface{}{table, indexName}, err, 0)
	}

	if c.config.logger != nil {
		c.config.logger.Info("Index dropped",
			"table", table,
			"index", indexName)
	}

	return nil
}

// RenameTable renames a table
// RenameTable은 테이블의 이름을 변경합니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	err := client.RenameTable(ctx, "old_users", "users")
//	if err != nil {
//	    log.Fatal(err)
//	}
func (c *Client) RenameTable(ctx context.Context, oldName string, newName string) error {
	query := fmt.Sprintf("RENAME TABLE %s TO %s", oldName, newName)

	_, err := c.Exec(ctx, query)
	if err != nil {
		return c.wrapError("RenameTable", query, []interface{}{oldName, newName}, err, 0)
	}

	if c.config.logger != nil {
		c.config.logger.Info("Table renamed",
			"old_name", oldName,
			"new_name", newName)
	}

	return nil
}

// AddForeignKey adds a foreign key constraint to a table
// AddForeignKey는 테이블에 외래 키 제약 조건을 추가합니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	err := client.AddForeignKey(ctx,
//	    "orders",           // table
//	    "fk_user",          // constraint name
//	    "user_id",          // column
//	    "users",            // referenced table
//	    "id",               // referenced column
//	    "CASCADE",          // on delete
//	    "CASCADE")          // on update
//	if err != nil {
//	    log.Fatal(err)
//	}
func (c *Client) AddForeignKey(ctx context.Context, table string, constraintName string,
	column string, refTable string, refColumn string, onDelete string, onUpdate string) error {

	query := fmt.Sprintf(
		"ALTER TABLE %s ADD CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s(%s) ON DELETE %s ON UPDATE %s",
		table, constraintName, column, refTable, refColumn, onDelete, onUpdate)

	_, err := c.Exec(ctx, query)
	if err != nil {
		return c.wrapError("AddForeignKey", query,
			[]interface{}{table, constraintName, column, refTable, refColumn}, err, 0)
	}

	if c.config.logger != nil {
		c.config.logger.Info("Foreign key added",
			"table", table,
			"constraint", constraintName,
			"column", column,
			"references", fmt.Sprintf("%s(%s)", refTable, refColumn))
	}

	return nil
}

// DropForeignKey removes a foreign key constraint from a table
// DropForeignKey는 테이블에서 외래 키 제약 조건을 제거합니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	err := client.DropForeignKey(ctx, "orders", "fk_user")
//	if err != nil {
//	    log.Fatal(err)
//	}
func (c *Client) DropForeignKey(ctx context.Context, table string, constraintName string) error {
	query := fmt.Sprintf("ALTER TABLE %s DROP FOREIGN KEY %s", table, constraintName)

	_, err := c.Exec(ctx, query)
	if err != nil {
		return c.wrapError("DropForeignKey", query, []interface{}{table, constraintName}, err, 0)
	}

	if c.config.logger != nil {
		c.config.logger.Info("Foreign key dropped",
			"table", table,
			"constraint", constraintName)
	}

	return nil
}

// CopyTable creates a copy of a table with a new name
// CopyTable은 새 이름으로 테이블의 복사본을 생성합니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	// Copy structure and data / 구조와 데이터 복사
//	err := client.CopyTable(ctx, "users", "users_backup", true)
//
//	// Copy only structure / 구조만 복사
//	err := client.CopyTable(ctx, "users", "users_template", false)
func (c *Client) CopyTable(ctx context.Context, sourceTable string, destTable string, withData bool) error {
	// First, create table structure / 먼저 테이블 구조 생성
	query := fmt.Sprintf("CREATE TABLE %s LIKE %s", destTable, sourceTable)

	_, err := c.Exec(ctx, query)
	if err != nil {
		return c.wrapError("CopyTable", query, []interface{}{sourceTable, destTable}, err, 0)
	}

	// If withData, copy data / withData가 true면 데이터 복사
	if withData {
		insertQuery := fmt.Sprintf("INSERT INTO %s SELECT * FROM %s", destTable, sourceTable)
		_, err := c.Exec(ctx, insertQuery)
		if err != nil {
			// Try to drop the newly created table / 새로 생성된 테이블 삭제 시도
			c.DropTable(ctx, destTable, true)
			return c.wrapError("CopyTable", insertQuery, []interface{}{sourceTable, destTable}, err, 0)
		}
	}

	if c.config.logger != nil {
		c.config.logger.Info("Table copied",
			"source", sourceTable,
			"destination", destTable,
			"with_data", withData)
	}

	return nil
}

// AlterTableEngine changes the storage engine of a table
// AlterTableEngine은 테이블의 스토리지 엔진을 변경합니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	err := client.AlterTableEngine(ctx, "users", "InnoDB")
//	if err != nil {
//	    log.Fatal(err)
//	}
func (c *Client) AlterTableEngine(ctx context.Context, table string, engine string) error {
	query := fmt.Sprintf("ALTER TABLE %s ENGINE=%s", table, engine)

	_, err := c.Exec(ctx, query)
	if err != nil {
		return c.wrapError("AlterTableEngine", query, []interface{}{table, engine}, err, 0)
	}

	if c.config.logger != nil {
		c.config.logger.Info("Table engine changed",
			"table", table,
			"engine", engine)
	}

	return nil
}

// AlterTableCharset changes the character set and collation of a table
// AlterTableCharset는 테이블의 문자 집합과 collation을 변경합니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	err := client.AlterTableCharset(ctx, "users", "utf8mb4", "utf8mb4_unicode_ci")
//	if err != nil {
//	    log.Fatal(err)
//	}
func (c *Client) AlterTableCharset(ctx context.Context, table string, charset string, collate string) error {
	query := fmt.Sprintf("ALTER TABLE %s CONVERT TO CHARACTER SET %s COLLATE %s",
		table, charset, collate)

	_, err := c.Exec(ctx, query)
	if err != nil {
		return c.wrapError("AlterTableCharset", query, []interface{}{table, charset, collate}, err, 0)
	}

	if c.config.logger != nil {
		c.config.logger.Info("Table charset changed",
			"table", table,
			"charset", charset,
			"collate", collate)
	}

	return nil
}
