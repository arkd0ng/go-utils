# MySQL Package / MySQL 패키지

Extreme simplicity MySQL/MariaDB client with zero-downtime credential rotation.

무중단 자격 증명 순환을 갖춘 극도로 간단한 MySQL/MariaDB 클라이언트.

## Features / 특징

### Core Features / 핵심 기능

✅ **Extreme Simplicity**: 30 lines → 2 lines of code / 극도의 간결함: 30줄 → 2줄 코드
✅ **Auto Everything**: Connection management, retry, cleanup / 모든 것 자동: 연결 관리, 재시도, 정리
✅ **Zero-Downtime Credential Rotation**: Multiple connection pools with rolling rotation / 무중단 자격 증명 순환
✅ **SQL-Like API**: Close to actual SQL syntax / SQL 문법에 가까운 API
✅ **No defer rows.Close()**: Automatic resource cleanup / 자동 리소스 정리
✅ **Auto Retry**: Transient errors are retried automatically / 일시적 에러 자동 재시도
✅ **Health Check**: Automatic connection monitoring / 자동 연결 모니터링

### Advanced Features / 고급 기능

✅ **Batch Operations**: BatchInsert, BatchUpdate, BatchDelete, BatchSelectByIDs / 배치 작업
✅ **Upsert Operations**: Upsert, UpsertBatch, Replace (ON DUPLICATE KEY) / Upsert 작업
✅ **Pagination**: Easy pagination with metadata (Paginate, PaginateQuery) / 페이지네이션
✅ **Soft Delete**: SoftDelete, Restore, SelectAllWithTrashed / 소프트 삭제
✅ **Query Statistics**: Performance monitoring, slow query logging / 쿼리 통계
✅ **Pool Metrics**: Connection pool health monitoring / 풀 메트릭
✅ **Schema Inspector**: Database schema introspection / 스키마 검사
✅ **Migration Helpers**: CreateTable, AddColumn, AddIndex, etc. / 마이그레이션 헬퍼
✅ **CSV Export/Import**: Export and import data in CSV format / CSV 내보내기/가져오기

## Installation / 설치

```bash
go get github.com/arkd0ng/go-utils/database/mysql
```

## Quick Start / 빠른 시작

### Basic Usage / 기본 사용법

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/arkd0ng/go-utils/database/mysql"
)

func main() {
    // Create client / 클라이언트 생성
    db, err := mysql.New(
        mysql.WithDSN("user:password@tcp(localhost:3306)/dbname?parseTime=true"),
    )
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    ctx := context.Background()

    // Insert / 삽입
    result, err := db.Insert(ctx, "users", map[string]interface{}{
        "name":  "John Doe",
        "email": "john@example.com",
        "age":   30,
    })
    if err != nil {
        log.Fatal(err)
    }

    userID, _ := result.LastInsertId()
    fmt.Printf("Inserted user with ID: %d\n", userID)

    // Select all / 모두 선택
    users, err := db.SelectAll(ctx, "users", "age > ?", 18)
    if err != nil {
        log.Fatal(err)
    }

    for _, user := range users {
        fmt.Printf("User: %+v\n", user)
    }

    // Update / 업데이트
    _, err = db.Update(ctx, "users",
        map[string]interface{}{"age": 31},
        "id = ?", userID)
    if err != nil {
        log.Fatal(err)
    }

    // Delete / 삭제
    _, err = db.Delete(ctx, "users", "id = ?", userID)
    if err != nil {
        log.Fatal(err)
    }
}
```

### Dynamic Credentials (Vault, AWS Secrets Manager, etc.) / 동적 자격 증명

```go
// User-provided function to get DSN / DSN을 가져오는 사용자 제공 함수
func getDSN() (string, error) {
    // Fetch from Vault, file, environment variable, etc.
    // Vault, 파일, 환경 변수 등에서 가져오기
    user := os.Getenv("DB_USER")
    pass := os.Getenv("DB_PASS")
    return fmt.Sprintf("%s:%s@tcp(localhost:3306)/mydb", user, pass), nil
}

// Create client with credential rotation / 자격 증명 순환으로 클라이언트 생성
db, err := mysql.New(
    mysql.WithCredentialRefresh(
        getDSN,        // User function / 사용자 함수
        3,             // 3 connection pools / 3개 연결 풀
        1*time.Hour,   // Rotate one per hour / 1시간마다 하나씩 교체
    ),
    mysql.WithLogger(logger),
)

// Result: Zero-downtime credential rotation!
// 결과: 무중단 자격 증명 순환!
// - Time 0:00: [Pool0, Pool1, Pool2] (Credential A)
// - Time 1:00: [Pool0, Pool1, Pool2-NEW] (Credential B)
// - Time 2:00: [Pool0, Pool1-NEW, Pool2-NEW] (Credential B)
//   → Credential A expires, but Pool1 & Pool2 still work!
```

## API Reference / API 참조

### Simple API / 간단한 API

#### SelectAll - Select all rows / 모든 행 선택

```go
// Select all / 모두 선택
users, err := db.SelectAll(ctx, "users")

// With condition / 조건 포함
users, err := db.SelectAll(ctx, "users", "age > ?", 18)
adults, err := db.SelectAll(ctx, "users", "age > ? AND city = ?", 18, "Seoul")
```

#### SelectOne - Select single row / 단일 행 선택

```go
user, err := db.SelectOne(ctx, "users", "id = ?", 123)
```

#### Insert - Insert new row / 새 행 삽입

```go
result, err := db.Insert(ctx, "users", map[string]interface{}{
    "name":  "John",
    "email": "john@example.com",
    "age":   30,
})
```

#### Update - Update rows / 행 업데이트

```go
result, err := db.Update(ctx, "users",
    map[string]interface{}{"name": "Jane", "age": 31},
    "id = ?", 123)
```

#### Delete - Delete rows / 행 삭제

```go
result, err := db.Delete(ctx, "users", "id = ?", 123)
```

#### Count - Count rows / 행 개수

```go
count, err := db.Count(ctx, "users")
count, err := db.Count(ctx, "users", "age > ?", 18)
```

#### Exists - Check existence / 존재 확인

```go
exists, err := db.Exists(ctx, "users", "email = ?", "john@example.com")
```

### Transaction API / 트랜잭션 API

```go
err := db.Transaction(ctx, func(tx *mysql.Tx) error {
    // Insert user / 사용자 삽입
    result, err := tx.Insert(ctx, "users", map[string]interface{}{
        "name": "John",
        "email": "john@example.com",
    })
    if err != nil {
        return err // Auto rollback / 자동 롤백
    }

    userID, _ := result.LastInsertId()

    // Insert profile / 프로필 삽입
    _, err = tx.Insert(ctx, "profiles", map[string]interface{}{
        "user_id": userID,
        "bio": "Hello world",
    })
    if err != nil {
        return err // Auto rollback / 자동 롤백
    }

    return nil // Auto commit / 자동 커밋
})
```

### Raw SQL API / Raw SQL API

```go
// Query / 쿼리
rows, err := db.Query(ctx, "SELECT * FROM users WHERE age > ?", 18)

// QueryRow / 단일 행 쿼리
row := db.QueryRow(ctx, "SELECT * FROM users WHERE id = ?", 123)

// Exec / 실행
result, err := db.Exec(ctx, "UPDATE users SET name = ? WHERE id = ?", "John", 123)
```

## Configuration Options / 설정 옵션

```go
db, err := mysql.New(
    // Connection / 연결
    mysql.WithDSN("user:pass@tcp(localhost:3306)/db"),
    mysql.WithMaxOpenConns(50),
    mysql.WithMaxIdleConns(10),
    mysql.WithConnMaxLifetime(5*time.Minute),

    // Credential Rotation / 자격 증명 순환
    mysql.WithCredentialRefresh(getDSN, 3, 1*time.Hour),

    // Timeout / 타임아웃
    mysql.WithConnectTimeout(10*time.Second),
    mysql.WithQueryTimeout(30*time.Second),

    // Retry / 재시도
    mysql.WithMaxRetries(3),
    mysql.WithRetryDelay(100*time.Millisecond),

    // Logging / 로깅
    mysql.WithLogger(logger),
    mysql.WithQueryLogging(true),
    mysql.WithSlowQueryLogging(true),
    mysql.WithSlowQueryThreshold(1*time.Second),

    // Health Check / 헬스 체크
    mysql.WithHealthCheck(true),
    mysql.WithHealthCheckInterval(30*time.Second),

    // Security / 보안
    mysql.WithTLS(tlsConfig),
)
```

## Why This Package? / 왜 이 패키지인가?

### Standard database/sql / 표준 database/sql

```go
// ❌ 30+ lines of boilerplate code / 30줄 이상의 보일러플레이트 코드
db, err := sql.Open("mysql", dsn)
if err != nil {
    return err
}
defer db.Close()

if err := db.Ping(); err != nil {
    // Manual reconnect logic / 수동 재연결 로직
}

rows, err := db.Query("SELECT * FROM users WHERE age > ?", 18)
if err != nil {
    return err
}
defer rows.Close() // ← Must remember! / 기억해야 함!

var users []User
for rows.Next() {
    var u User
    err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age)
    if err != nil {
        return err
    }
    users = append(users, u)
}
if err := rows.Err(); err != nil {
    return err
}
```

### This Package / 이 패키지

```go
// ✅ 2 lines / 2줄
db, _ := mysql.New(mysql.WithDSN(dsn))
users, _ := db.SelectAll(ctx, "users", "age > ?", 18)

// That's it! Everything else is automatic:
// 끝! 나머지는 모두 자동:
// ✓ Auto connection management / 자동 연결 관리
// ✓ Auto reconnect / 자동 재연결
// ✓ Auto retry / 자동 재시도
// ✓ Auto rows.Close() / 자동 rows.Close()
// ✓ Auto error handling / 자동 에러 처리
```

## Advanced Features / 고급 기능

### Batch Operations / 배치 작업

Perform bulk operations for better performance / 더 나은 성능을 위한 대량 작업:

```go
// BatchInsert - Insert multiple rows in one query / 한 쿼리로 여러 행 삽입
data := []map[string]interface{}{
    {"name": "John", "age": 30, "email": "john@example.com"},
    {"name": "Jane", "age": 25, "email": "jane@example.com"},
    {"name": "Bob", "age": 35, "email": "bob@example.com"},
}
result, err := db.BatchInsert(ctx, "users", data)

// BatchUpdate - Update multiple rows with different values / 다른 값으로 여러 행 업데이트
updates := []map[string]interface{}{
    {"id": 1, "name": "John Updated", "age": 31},
    {"id": 2, "name": "Jane Updated", "age": 26},
}
results, err := db.BatchUpdate(ctx, "users", updates, "id")

// BatchDelete - Delete multiple rows by IDs / ID로 여러 행 삭제
ids := []interface{}{1, 2, 3, 4, 5}
result, err := db.BatchDelete(ctx, "users", "id", ids)

// BatchSelectByIDs - Select multiple rows by IDs / ID로 여러 행 선택
ids := []interface{}{1, 2, 3}
users, err := db.BatchSelectByIDs(ctx, "users", "id", ids)
```

### Upsert Operations / Upsert 작업

Insert or update with ON DUPLICATE KEY UPDATE / ON DUPLICATE KEY UPDATE로 삽입 또는 업데이트:

```go
// Upsert - Insert or update on duplicate key / 중복 키 시 삽입 또는 업데이트
data := map[string]interface{}{
    "email": "john@example.com",  // Unique key / 고유 키
    "name": "John Doe",
    "age": 30,
}
updateColumns := []string{"name", "age"}  // Columns to update on duplicate / 중복 시 업데이트할 컬럼
result, err := db.Upsert(ctx, "users", data, updateColumns)

// UpsertBatch - Batch upsert / 배치 Upsert
data := []map[string]interface{}{
    {"email": "john@example.com", "name": "John", "age": 30},
    {"email": "jane@example.com", "name": "Jane", "age": 25},
}
result, err := db.UpsertBatch(ctx, "users", data, []string{"name", "age"})

// Replace - Replace row (DELETE + INSERT) / 행 교체 (DELETE + INSERT)
result, err := db.Replace(ctx, "users", map[string]interface{}{
    "id": 1,
    "name": "John Replaced",
    "email": "john@example.com",
})
```

### Pagination / 페이지네이션

Easy pagination with metadata / 메타데이터와 함께 쉬운 페이지네이션:

```go
// Paginate - Simple pagination / 간단한 페이지네이션
result, err := db.Paginate(ctx, "users", 1, 10)  // Page 1, 10 items per page
fmt.Printf("Page %d of %d (Total: %d rows)\n", result.Page, result.TotalPages, result.TotalRows)
fmt.Printf("HasNext: %v, HasPrev: %v\n", result.HasNext, result.HasPrev)

// With conditions and ordering / 조건 및 정렬 포함
result, err := db.Paginate(ctx, "users", 2, 20,
    "age > ?", 18,  // WHERE condition / WHERE 조건
    mysql.WithOrderBy("created_at DESC"))

// PaginateQuery - Paginate custom query / 커스텀 쿼리 페이지네이션
query := "SELECT * FROM users WHERE age > ?"
result, err := db.PaginateQuery(ctx, query, 1, 10, 18)
```

### Soft Delete / 소프트 삭제

Soft delete with automatic restoration / 자동 복원과 함께 소프트 삭제:

```go
// SoftDelete - Mark as deleted / 삭제로 표시
result, err := db.SoftDelete(ctx, "users", "id = ?", 1)

// Restore - Restore soft-deleted row / 소프트 삭제된 행 복원
result, err := db.Restore(ctx, "users", "id = ?", 1)

// SelectAllWithTrashed - Include soft-deleted rows / 소프트 삭제된 행 포함
users, err := db.SelectAllWithTrashed(ctx, "users")

// SelectAllOnlyTrashed - Only soft-deleted rows / 소프트 삭제된 행만
users, err := db.SelectAllOnlyTrashed(ctx, "users")

// PermanentDelete - Permanently delete / 영구 삭제
result, err := db.PermanentDelete(ctx, "users", "id = ?", 1)
```

**Note**: Soft delete requires a `deleted_at` column (TIMESTAMP NULL) in your table.

**참고**: 소프트 삭제는 테이블에 `deleted_at` 컬럼(TIMESTAMP NULL)이 필요합니다.

### Query Statistics / 쿼리 통계

Monitor query performance / 쿼리 성능 모니터링:

```go
// GetQueryStats - Get query statistics / 쿼리 통계 가져오기
stats := db.GetQueryStats()
fmt.Printf("Total Queries: %d\n", stats.TotalQueries)
fmt.Printf("Average Duration: %v\n", stats.AverageDuration)
fmt.Printf("Slowest Query: %v (%s)\n", stats.SlowestQuery, stats.SlowestQuerySQL)

// EnableSlowQueryLog - Enable slow query logging / 느린 쿼리 로깅 활성화
db.EnableSlowQueryLog(1 * time.Second)  // Log queries slower than 1s

// GetSlowQueries - Get slow query list / 느린 쿼리 목록 가져오기
slowQueries := db.GetSlowQueries()
for _, q := range slowQueries {
    fmt.Printf("Slow query: %s (took %v)\n", q.SQL, q.Duration)
}

// ResetQueryStats - Reset statistics / 통계 초기화
db.ResetQueryStats()
```

### Pool Metrics / 풀 메트릭

Monitor connection pool health / 연결 풀 상태 모니터링:

```go
// GetPoolMetrics - Get connection pool metrics / 연결 풀 메트릭 가져오기
metrics := db.GetPoolMetrics()
fmt.Printf("Open Connections: %d / %d\n", metrics.OpenConnections, metrics.MaxOpenConnections)
fmt.Printf("Idle Connections: %d / %d\n", metrics.IdleConnections, metrics.MaxIdleConnections)
fmt.Printf("Wait Count: %d (Duration: %v)\n", metrics.WaitCount, metrics.WaitDuration)

// GetPoolHealthInfo - Get pool health status / 풀 상태 정보 가져오기
health := db.GetPoolHealthInfo()
fmt.Printf("Status: %s\n", health.Status)
fmt.Printf("Healthy: %v\n", health.Healthy)
fmt.Printf("Message: %s\n", health.Message)

// GetConnectionUtilization - Get connection utilization percentage / 연결 사용률 가져오기
utilization := db.GetConnectionUtilization()
fmt.Printf("Utilization: %.2f%%\n", utilization)
```

### Schema Inspector / 스키마 검사

Inspect database schema / 데이터베이스 스키마 검사:

```go
// GetTables - List all tables / 모든 테이블 목록
tables, err := db.GetTables(ctx)
for _, table := range tables {
    fmt.Printf("Table: %s\n", table)
}

// GetColumns - Get table columns / 테이블 컬럼 가져오기
columns, err := db.GetColumns(ctx, "users")
for _, col := range columns {
    fmt.Printf("%s: %s (Nullable: %v)\n", col.Name, col.Type, col.Nullable)
}

// GetIndexes - Get table indexes / 테이블 인덱스 가져오기
indexes, err := db.GetIndexes(ctx, "users")
for _, idx := range indexes {
    fmt.Printf("Index: %s on %v (Unique: %v)\n", idx.Name, idx.Columns, idx.Unique)
}

// TableExists - Check if table exists / 테이블 존재 확인
exists, err := db.TableExists(ctx, "users")

// InspectTable - Get complete table info / 전체 테이블 정보 가져오기
info, err := db.InspectTable(ctx, "users")
fmt.Printf("Table: %s (%d columns, %d indexes)\n",
    info.Name, len(info.Columns), len(info.Indexes))
```

### Migration Helpers / 마이그레이션 헬퍼

Database migration utilities / 데이터베이스 마이그레이션 유틸리티:

```go
// CreateTable - Create table with schema / 스키마로 테이블 생성
schema := map[string]string{
    "id": "INT AUTO_INCREMENT PRIMARY KEY",
    "name": "VARCHAR(100) NOT NULL",
    "email": "VARCHAR(255) UNIQUE",
    "created_at": "TIMESTAMP DEFAULT CURRENT_TIMESTAMP",
}
err := db.CreateTable(ctx, "users", schema)

// DropTable - Drop table / 테이블 삭제
err := db.DropTable(ctx, "users")

// TruncateTable - Truncate table / 테이블 초기화
err := db.TruncateTable(ctx, "users")

// AddColumn - Add column to table / 테이블에 컬럼 추가
err := db.AddColumn(ctx, "users", "phone", "VARCHAR(20)")

// DropColumn - Drop column from table / 테이블에서 컬럼 삭제
err := db.DropColumn(ctx, "users", "phone")

// AddIndex - Add index to table / 테이블에 인덱스 추가
err := db.AddIndex(ctx, "users", "idx_email", []string{"email"}, false)

// AddForeignKey - Add foreign key constraint / 외래 키 제약 조건 추가
err := db.AddForeignKey(ctx, "profiles", "fk_user", "user_id", "users", "id",
    mysql.FKOnDeleteCascade, mysql.FKOnUpdateCascade)
```

### CSV Export/Import / CSV 내보내기/가져오기

Export and import data in CSV format / CSV 형식으로 데이터 내보내기 및 가져오기:

```go
// ExportTableToCSV - Export entire table to CSV / 전체 테이블을 CSV로 내보내기
err := db.ExportTableToCSV(ctx, "users", "users.csv")

// ExportQueryToCSV - Export query results to CSV / 쿼리 결과를 CSV로 내보내기
query := "SELECT id, name, email FROM users WHERE age > ?"
err := db.ExportQueryToCSV(ctx, query, "active_users.csv", 18)

// ImportFromCSV - Import CSV to table / CSV를 테이블로 가져오기
rowsImported, err := db.ImportFromCSV(ctx, "users.csv", "users")
fmt.Printf("Imported %d rows\n", rowsImported)

// Import with column mapping / 컬럼 매핑과 함께 가져오기
columnMap := map[string]string{
    "FullName": "name",
    "EmailAddress": "email",
}
rowsImported, err := db.ImportFromCSV(ctx, "users.csv", "users", mysql.WithColumnMapping(columnMap))
```

## Best Practices / 모범 사례

1. **Always use context**: Pass context for timeout control / 타임아웃 제어를 위해 항상 context 전달
2. **Use transactions for multi-step operations**: Ensure atomicity / 다단계 작업에는 트랜잭션 사용
3. **Enable query logging in development**: Debug slow queries / 개발 중 쿼리 로깅 활성화
4. **Use credential rotation for production**: Enhanced security / 프로덕션에는 자격 증명 순환 사용
5. **Monitor health checks**: Track connection status / 헬스 체크 모니터링
6. **Use batch operations**: For bulk inserts/updates/deletes / 대량 삽입/업데이트/삭제에는 배치 작업 사용
7. **Monitor query statistics**: Identify slow queries / 느린 쿼리 식별
8. **Use pagination for large datasets**: Avoid loading all rows / 모든 행 로드 방지
9. **Use soft delete when needed**: Preserve data history / 데이터 히스토리 보존
10. **Check pool metrics**: Monitor connection pool health / 연결 풀 상태 모니터링

## License / 라이선스

MIT

## Contributing / 기여

Contributions are welcome! Please follow the project's coding standards.

기여를 환영합니다! 프로젝트의 코딩 표준을 따라주세요.
