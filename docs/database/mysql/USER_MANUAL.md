# MySQL Package - User Manual / 사용자 매뉴얼

**Package**: `github.com/arkd0ng/go-utils/database/mysql`
**Version**: v1.3.006
**Author**: arkd0ng
**Last Updated**: 2025-10-10

---

## Table of Contents / 목차

1. [Introduction / 소개](#introduction--소개)
2. [Installation / 설치](#installation--설치)
3. [Quick Start / 빠른 시작](#quick-start--빠른-시작)
4. [Configuration Reference / 설정 참조](#configuration-reference--설정-참조)
5. [API Reference / API 참조](#api-reference--api-참조)
6. [Usage Patterns / 사용 패턴](#usage-patterns--사용-패턴)
7. [Common Use Cases / 일반적인 사용 사례](#common-use-cases--일반적인-사용-사례)
8. [Best Practices / 모범 사례](#best-practices--모범-사례)
9. [Troubleshooting / 문제 해결](#troubleshooting--문제-해결)
10. [FAQ](#faq)

---

## Introduction / 소개

### What is this package? / 이 패키지는 무엇인가?

The `database/mysql` package is an **extremely simplified** MySQL client for Go that reduces typical database operations from **30+ lines to 1-2 lines** of code.

`database/mysql` 패키지는 일반적인 데이터베이스 작업을 **30줄 이상에서 1-2줄**로 줄이는 **매우 간단한** MySQL 클라이언트입니다.

### Key Features / 주요 기능

- **✅ Extreme Simplicity**: `30 lines → 2 lines` code reduction / 극도의 간결함: `30줄 → 2줄` 코드 감소
- **✅ Auto Everything**: Connection management, retry, resource cleanup / 모든 것이 자동: 연결 관리, 재시도, 리소스 정리
- **✅ Three-Layer API**: Simple, Query Builder, Raw SQL / 3계층 API: 간단, 쿼리 빌더, Raw SQL
- **✅ No `defer rows.Close()`**: Automatic resource management / 자동 리소스 관리
- **✅ Auto Reconnect**: Handles connection loss automatically / 자동 재연결: 연결 손실 자동 처리
- **✅ Auto Retry**: Retries transient errors automatically / 자동 재시도: 일시적 에러 자동 재시도
- **✅ Credential Rotation**: Zero-downtime credential updates (optional) / 무중단 자격 증명 순환 (선택)
- **✅ Transaction Support**: Auto commit/rollback / 트랜잭션 지원: 자동 커밋/롤백

### Design Philosophy / 설계 철학

**"If it's not 10x simpler, don't build it"**
**"10배 간단하지 않으면 만들지 마세요"**

This package exists to eliminate all the boilerplate code you write with `database/sql`:
- No manual connection management
- No `defer rows.Close()`
- No complex scanning logic
- No manual retry logic
- No SQL string building for simple operations

이 패키지는 `database/sql`에서 작성하는 모든 보일러플레이트 코드를 제거합니다:
- 수동 연결 관리 불필요
- `defer rows.Close()` 불필요
- 복잡한 스캔 로직 불필요
- 수동 재시도 로직 불필요
- 간단한 작업을 위한 SQL 문자열 빌드 불필요

---

## Installation / 설치

### Prerequisites / 전제 조건

- **Go version**: 1.18 or higher / Go 1.18 이상
- **MySQL server**: 5.7 or higher / MySQL 서버 5.7 이상
- **Network access**: To your MySQL server / MySQL 서버에 대한 네트워크 액세스

### Install Package / 패키지 설치

```bash
go get github.com/arkd0ng/go-utils/database/mysql
```

### Import / 임포트

```go
import "github.com/arkd0ng/go-utils/database/mysql"
```

---

## Quick Start / 빠른 시작

### Example 1: Basic Usage / 기본 사용법

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/arkd0ng/go-utils/database/mysql"
)

func main() {
    // Connect to MySQL / MySQL 연결
    db, err := mysql.New(
        mysql.WithDSN("user:password@tcp(localhost:3306)/database"),
    )
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    ctx := context.Background()

    // Select all users / 모든 사용자 선택
    users, err := db.SelectAll(ctx, "users")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Found %d users\n", len(users))
}
```

### Example 2: Insert, Update, Delete / 삽입, 업데이트, 삭제

```go
// Insert / 삽입
result, err := db.Insert(ctx, "users", map[string]any{
    "name":  "John Doe",
    "email": "john@example.com",
    "age":   30,
})
id, _ := result.LastInsertId()

// Update / 업데이트
db.Update(ctx, "users",
    map[string]any{"age": 31},
    "id = ?", id)

// Delete / 삭제
db.Delete(ctx, "users", "id = ?", id)
```

### Example 3: Query with Options / 옵션을 사용한 쿼리

```go
// One-liner with multiple options / 여러 옵션을 사용한 한 줄 쿼리
users, _ := db.SelectWhere(ctx, "users", "age > ?", 25,
    mysql.WithColumns("name", "email", "age"),
    mysql.WithOrderBy("age DESC"),
    mysql.WithLimit(10))
```

### Example 4: Transaction / 트랜잭션

```go
err := db.Transaction(ctx, func(tx *mysql.Tx) error {
    tx.Insert(ctx, "users", map[string]any{"name": "Alice"})
    tx.Insert(ctx, "users", map[string]any{"name": "Bob"})
    return nil // Auto commit / 자동 커밋
})
```

---

## Configuration Reference / 설정 참조

### Connection Options / 연결 옵션

All options use the **functional options pattern** for flexible configuration.

모든 옵션은 유연한 설정을 위해 **함수형 옵션 패턴**을 사용합니다.

| Option / 옵션 | Type / 타입 | Default / 기본값 | Description / 설명 |
|--------------|------------|----------------|-------------------|
| `WithDSN(string)` | string | required / 필수 | MySQL DSN connection string / MySQL DSN 연결 문자열 |
| `WithMaxOpenConns(int)` | int | 25 | Maximum open connections / 최대 열린 연결 수 |
| `WithMaxIdleConns(int)` | int | 10 | Maximum idle connections / 최대 유휴 연결 수 |
| `WithConnMaxLifetime(duration)` | duration | 5m | Maximum connection lifetime / 최대 연결 수명 |
| `WithConnMaxIdleTime(duration)` | duration | 5m | Maximum idle time / 최대 유휴 시간 |
| `WithRetryMaxAttempts(int)` | int | 3 | Maximum retry attempts / 최대 재시도 횟수 |
| `WithRetryInitialInterval(duration)` | duration | 100ms | Initial retry interval / 초기 재시도 간격 |
| `WithRetryMaxInterval(duration)` | duration | 1s | Maximum retry interval / 최대 재시도 간격 |
| `WithRetryMultiplier(float64)` | float64 | 2.0 | Retry backoff multiplier / 재시도 백오프 배수 |

### DSN Format / DSN 형식

```
[username[:password]@][protocol[(address)]]/dbname[?param1=value1&paramN=valueN]
```

**Examples / 예제**:
```go
// Basic / 기본
"user:password@tcp(localhost:3306)/database"

// With parameters / 파라미터 포함
"user:password@tcp(localhost:3306)/database?parseTime=true&charset=utf8mb4"

// Unix socket / 유닉스 소켓
"user:password@unix(/tmp/mysql.sock)/database"
```

### Connection Pool Configuration / 연결 풀 설정

```go
db, err := mysql.New(
    mysql.WithDSN("user:password@tcp(localhost:3306)/database"),
    mysql.WithMaxOpenConns(25),          // Max connections / 최대 연결 수
    mysql.WithMaxIdleConns(10),          // Idle connections / 유휴 연결 수
    mysql.WithConnMaxLifetime(5*time.Minute),  // Connection lifetime / 연결 수명
)
```

**Recommendations / 권장사항**:
- `MaxOpenConns`: Set based on your MySQL `max_connections` setting / MySQL `max_connections` 설정에 따라 설정
- `MaxIdleConns`: Typically 50% of `MaxOpenConns` / 일반적으로 `MaxOpenConns`의 50%
- `ConnMaxLifetime`: 5-10 minutes to handle firewall timeouts / 방화벽 타임아웃 처리를 위해 5-10분

### Retry Configuration / 재시도 설정

```go
db, err := mysql.New(
    mysql.WithDSN("..."),
    mysql.WithRetryMaxAttempts(5),                    // Retry up to 5 times / 최대 5회 재시도
    mysql.WithRetryInitialInterval(100*time.Millisecond),  // Start with 100ms / 100ms로 시작
    mysql.WithRetryMaxInterval(2*time.Second),        // Cap at 2s / 최대 2초
    mysql.WithRetryMultiplier(2.0),                   // Double each time / 매번 2배
)
```

**Exponential Backoff / 지수 백오프**:
- Attempt 1: 100ms
- Attempt 2: 200ms
- Attempt 3: 400ms
- Attempt 4: 800ms
- Attempt 5: 1600ms (capped at 2s)

---

## API Reference / API 참조

### Layer 1: Simple API / 간단한 API

The simplest API for common CRUD operations.

일반적인 CRUD 작업을 위한 가장 간단한 API입니다.

#### SelectAll

Select all rows from a table with optional WHERE condition.

선택적 WHERE 조건으로 테이블의 모든 행을 선택합니다.

```go
func (c *Client) SelectAll(ctx context.Context, table string, conditionAndArgs ...interface{}) ([]map[string]interface{}, error)
```

**Examples / 예제**:
```go
// Select all / 전체 선택
users, _ := db.SelectAll(ctx, "users")

// With condition / 조건 포함
users, _ := db.SelectAll(ctx, "users", "age > ?", 18)
users, _ := db.SelectAll(ctx, "users", "age > ? AND city = ?", 18, "Seoul")
```

#### SelectOne

Select a single row from a table.

테이블에서 단일 행을 선택합니다.

```go
func (c *Client) SelectOne(ctx context.Context, table string, conditionAndArgs ...interface{}) (map[string]interface{}, error)
```

**Examples / 예제**:
```go
user, _ := db.SelectOne(ctx, "users", "id = ?", 123)
user, _ := db.SelectOne(ctx, "users", "email = ?", "john@example.com")
```

#### SelectColumn

Select all rows with a single column from a table.

테이블에서 단일 컬럼으로 모든 행을 선택합니다.

```go
func (c *Client) SelectColumn(ctx context.Context, table string, column string, conditionAndArgs ...interface{}) ([]map[string]interface{}, error)
```

**Examples / 예제**:
```go
// SELECT email FROM users
emails, _ := db.SelectColumn(ctx, "users", "email")

// SELECT name FROM users WHERE age > 25
names, _ := db.SelectColumn(ctx, "users", "name", "age > ?", 25)

// Process results / 결과 처리
for _, row := range emails {
    fmt.Println(row["email"])
}
```

#### SelectColumns

Select all rows with multiple columns from a table.

테이블에서 여러 컬럼으로 모든 행을 선택합니다.

```go
func (c *Client) SelectColumns(ctx context.Context, table string, columns []string, conditionAndArgs ...interface{}) ([]map[string]interface{}, error)
```

**Examples / 예제**:
```go
// SELECT name, email FROM users
users, _ := db.SelectColumns(ctx, "users", []string{"name", "email"})

// SELECT name, age, city FROM users WHERE age > 25
users, _ := db.SelectColumns(ctx, "users", []string{"name", "age", "city"}, "age > ?", 25)

// Process results / 결과 처리
for _, user := range users {
    fmt.Printf("%s <%s>\n", user["name"], user["email"])
}
```

#### Insert

Insert a new row into a table.

테이블에 새 행을 삽입합니다.

```go
func (c *Client) Insert(ctx context.Context, table string, data map[string]interface{}) (sql.Result, error)
```

**Examples / 예제**:
```go
result, _ := db.Insert(ctx, "users", map[string]any{
    "name":  "John Doe",
    "email": "john@example.com",
    "age":   30,
    "city":  "Seoul",
})

id, _ := result.LastInsertId()
fmt.Printf("Inserted ID: %d\n", id)
```

#### Update

Update rows in a table.

테이블의 행을 업데이트합니다.

```go
func (c *Client) Update(ctx context.Context, table string, data map[string]interface{}, conditionAndArgs ...interface{}) (sql.Result, error)
```

**Examples / 예제**:
```go
// Update specific row / 특정 행 업데이트
result, _ := db.Update(ctx, "users",
    map[string]any{"age": 31, "city": "Busan"},
    "id = ?", 123)

// Update all rows / 모든 행 업데이트 (조심!)
result, _ := db.Update(ctx, "users",
    map[string]any{"status": "active"})

rows, _ := result.RowsAffected()
fmt.Printf("Updated %d rows\n", rows)
```

#### Delete

Delete rows from a table.

테이블에서 행을 삭제합니다.

```go
func (c *Client) Delete(ctx context.Context, table string, conditionAndArgs ...interface{}) (sql.Result, error)
```

**Examples / 예제**:
```go
// Delete specific row / 특정 행 삭제
result, _ := db.Delete(ctx, "users", "id = ?", 123)

// Delete multiple rows / 여러 행 삭제
result, _ := db.Delete(ctx, "users", "age < ? AND status = ?", 18, "inactive")

rows, _ := result.RowsAffected()
fmt.Printf("Deleted %d rows\n", rows)
```

#### Count

Count rows in a table.

테이블의 행 수를 계산합니다.

```go
func (c *Client) Count(ctx context.Context, table string, conditionAndArgs ...interface{}) (int64, error)
```

**Examples / 예제**:
```go
// Count all / 전체 수
total, _ := db.Count(ctx, "users")

// Count with condition / 조건으로 계산
adults, _ := db.Count(ctx, "users", "age >= ?", 18)
```

#### Exists

Check if rows exist in a table.

테이블에 행이 존재하는지 확인합니다.

```go
func (c *Client) Exists(ctx context.Context, table string, conditionAndArgs ...interface{}) (bool, error)
```

**Examples / 예제**:
```go
// Check existence / 존재 확인
exists, _ := db.Exists(ctx, "users", "email = ?", "john@example.com")
if exists {
    fmt.Println("User already exists")
}
```

### Layer 2: SelectWhere API (Functional Options) / SelectWhere API (함수형 옵션)

One-liner queries with functional options for simple to moderate complexity.

간단~중간 복잡도를 위한 함수형 옵션을 사용한 한 줄 쿼리입니다.

#### SelectWhere

Select rows with flexible options.

유연한 옵션으로 행을 선택합니다.

```go
func (c *Client) SelectWhere(ctx context.Context, table string, conditionAndArgs ...interface{}) ([]map[string]interface{}, error)
```

**Available Options / 사용 가능한 옵션**:

| Option / 옵션 | Description / 설명 |
|--------------|-------------------|
| `WithColumns(...string)` | SELECT specific columns / 특정 컬럼 선택 |
| `WithOrderBy(string)` | ORDER BY clause / ORDER BY 절 |
| `WithLimit(int)` | LIMIT clause / LIMIT 절 |
| `WithOffset(int)` | OFFSET clause / OFFSET 절 |
| `WithGroupBy(...string)` | GROUP BY clause / GROUP BY 절 |
| `WithHaving(string, ...interface{})` | HAVING clause / HAVING 절 |
| `WithJoin(table, condition)` | INNER JOIN / INNER JOIN |
| `WithLeftJoin(table, condition)` | LEFT JOIN / LEFT JOIN |
| `WithRightJoin(table, condition)` | RIGHT JOIN / RIGHT JOIN |
| `WithDistinct()` | DISTINCT keyword / DISTINCT 키워드 |

**Examples / 예제**:

```go
// Simple query with columns and ordering / 컬럼과 정렬을 사용한 간단한 쿼리
users, _ := db.SelectWhere(ctx, "users", "age > ?", 25,
    mysql.WithColumns("name", "email", "age"),
    mysql.WithOrderBy("age DESC"),
    mysql.WithLimit(10))

// GROUP BY with HAVING / HAVING을 사용한 GROUP BY
results, _ := db.SelectWhere(ctx, "users", "",
    mysql.WithColumns("city", "COUNT(*) as count"),
    mysql.WithGroupBy("city"),
    mysql.WithHaving("COUNT(*) > ?", 5),
    mysql.WithOrderBy("count DESC"))

// DISTINCT query / DISTINCT 쿼리
cities, _ := db.SelectWhere(ctx, "users", "age > ?", 25,
    mysql.WithColumns("city"),
    mysql.WithDistinct(),
    mysql.WithOrderBy("city ASC"))

// Pagination / 페이징
users, _ := db.SelectWhere(ctx, "users", "status = ?", "active",
    mysql.WithOrderBy("created_at DESC"),
    mysql.WithLimit(20),
    mysql.WithOffset(40))  // Page 3 (20 per page)
```

#### SelectOneWhere

Select a single row with options.

옵션으로 단일 행을 선택합니다.

```go
func (c *Client) SelectOneWhere(ctx context.Context, table string, conditionAndArgs ...interface{}) (map[string]interface{}, error)
```

**Examples / 예제**:
```go
// Select specific columns / 특정 컬럼 선택
user, _ := db.SelectOneWhere(ctx, "users", "id = ?", 123,
    mysql.WithColumns("name", "email"))
```

### Layer 3: Query Builder API / 쿼리 빌더 API

Fluent API for complex queries with multiple JOINs.

여러 JOIN이 있는 복잡한 쿼리를 위한 Fluent API입니다.

#### Select

Start a query builder chain.

쿼리 빌더 체인을 시작합니다.

```go
func (c *Client) Select(cols ...string) *QueryBuilder
```

**Methods / 메서드**:
- `From(table string)` - FROM clause / FROM 절
- `Join(table, condition)` - INNER JOIN
- `LeftJoin(table, condition)` - LEFT JOIN
- `RightJoin(table, condition)` - RIGHT JOIN
- `Where(condition, args...)` - WHERE clause / WHERE 절
- `GroupBy(cols...)` - GROUP BY
- `Having(condition, args...)` - HAVING
- `OrderBy(order)` - ORDER BY
- `Limit(n)` - LIMIT
- `Offset(n)` - OFFSET
- `All(ctx)` - Execute and return all rows / 실행 및 모든 행 반환
- `One(ctx)` - Execute and return one row / 실행 및 단일 행 반환

**Examples / 예제**:

```go
// Simple query / 간단한 쿼리
users, _ := db.Select("name", "email", "age").
    From("users").
    Where("age > ?", 25).
    OrderBy("age DESC").
    Limit(10).
    All(ctx)

// Complex JOIN query / 복잡한 JOIN 쿼리
results, _ := db.Select("u.name", "o.id as order_id", "o.total").
    From("users u").
    Join("orders o", "u.id = o.user_id").
    Where("o.status = ?", "completed").
    Where("o.total > ?", 100).
    OrderBy("o.total DESC").
    All(ctx)

// LEFT JOIN with aggregation / 집계를 사용한 LEFT JOIN
results, _ := db.Select("u.name", "COUNT(o.id) as order_count", "SUM(o.total) as total_spent").
    From("users u").
    LeftJoin("orders o", "u.id = o.user_id").
    GroupBy("u.id", "u.name").
    Having("COUNT(o.id) > ?", 5).
    OrderBy("total_spent DESC").
    All(ctx)

// Single result / 단일 결과
user, _ := db.Select("*").
    From("users").
    Where("email = ?", "john@example.com").
    One(ctx)
```

### Layer 4: Raw SQL / Raw SQL

Direct SQL execution for maximum control.

최대 제어를 위한 직접 SQL 실행입니다.

#### Query

Execute a SELECT query and return rows.

SELECT 쿼리를 실행하고 행을 반환합니다.

```go
func (c *Client) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
```

**Examples / 예제**:
```go
// Complex query / 복잡한 쿼리
rows, _ := db.Query(ctx, `
    WITH ranked_users AS (
        SELECT *, ROW_NUMBER() OVER (PARTITION BY country ORDER BY score DESC) as rank
        FROM users
    )
    SELECT * FROM ranked_users WHERE rank <= 10
`)
defer rows.Close()

// Manual scanning / 수동 스캔
for rows.Next() {
    var id int
    var name string
    rows.Scan(&id, &name)
    fmt.Printf("%d: %s\n", id, name)
}
```

#### Exec

Execute an INSERT, UPDATE, or DELETE statement.

INSERT, UPDATE 또는 DELETE 문을 실행합니다.

```go
func (c *Client) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
```

**Examples / 예제**:
```go
// Complex update / 복잡한 업데이트
result, _ := db.Exec(ctx, `
    UPDATE users
    SET last_login = NOW(), login_count = login_count + 1
    WHERE id = ?
`, userID)

// Batch insert / 배치 삽입
result, _ := db.Exec(ctx, `
    INSERT INTO logs (user_id, action, created_at)
    VALUES (?, ?, NOW()), (?, ?, NOW()), (?, ?, NOW())
`, user1, "login", user2, "logout", user3, "purchase")
```

### Transaction API / 트랜잭션 API

#### Transaction

Execute a function within a transaction with automatic commit/rollback.

자동 커밋/롤백으로 트랜잭션 내에서 함수를 실행합니다.

```go
func (c *Client) Transaction(ctx context.Context, fn func(tx *Tx) error) error
```

**All APIs available within transaction / 트랜잭션 내에서 사용 가능한 모든 API**:
- Simple API: `SelectAll`, `SelectOne`, `Insert`, `Update`, `Delete`, `Count`, `Exists`
- SelectWhere API: `SelectWhere`, `SelectOneWhere`
- Query Builder: `Select().From()...`
- Raw SQL: `Query`, `Exec`

**Examples / 예제**:

```go
// Simple transaction / 간단한 트랜잭션
err := db.Transaction(ctx, func(tx *mysql.Tx) error {
    // All operations within this function are in a transaction
    // 이 함수 내의 모든 작업은 트랜잭션 내에서 실행됩니다

    tx.Insert(ctx, "users", map[string]any{"name": "Alice"})
    tx.Insert(ctx, "users", map[string]any{"name": "Bob"})

    return nil // Commits / 커밋
})
if err != nil {
    // Transaction was rolled back / 트랜잭션이 롤백되었습니다
}

// Complex transaction with error handling / 에러 처리를 포함한 복잡한 트랜잭션
err := db.Transaction(ctx, func(tx *mysql.Tx) error {
    // Deduct from account / 계정에서 차감
    result, err := tx.Update(ctx, "accounts",
        map[string]any{"balance": db.Raw("balance - ?", amount)},
        "user_id = ?", fromUserID)
    if err != nil {
        return err // Auto rollback / 자동 롤백
    }

    rows, _ := result.RowsAffected()
    if rows == 0 {
        return fmt.Errorf("account not found") // Auto rollback / 자동 롤백
    }

    // Add to account / 계정에 추가
    _, err = tx.Update(ctx, "accounts",
        map[string]any{"balance": db.Raw("balance + ?", amount)},
        "user_id = ?", toUserID)
    if err != nil {
        return err // Auto rollback / 자동 롤백
    }

    // Log transaction / 트랜잭션 로그
    tx.Insert(ctx, "transaction_logs", map[string]any{
        "from_user": fromUserID,
        "to_user":   toUserID,
        "amount":    amount,
    })

    return nil // Commit / 커밋
})
```

**Auto Rollback on Panic / 패닉 시 자동 롤백**:
```go
err := db.Transaction(ctx, func(tx *mysql.Tx) error {
    tx.Insert(ctx, "users", map[string]any{"name": "Alice"})

    panic("Something went wrong!") // Auto rollback / 자동 롤백

    return nil
})
// Transaction is automatically rolled back / 트랜잭션이 자동으로 롤백됩니다
```

---

## Usage Patterns / 사용 패턴

### Pattern 1: Basic CRUD Operations / 기본 CRUD 작업

```go
// Create / 생성
result, _ := db.Insert(ctx, "products", map[string]any{
    "name":  "iPhone 14",
    "price": 999.99,
    "stock": 100,
})
productID, _ := result.LastInsertId()

// Read / 읽기
product, _ := db.SelectOne(ctx, "products", "id = ?", productID)
fmt.Printf("Product: %v\n", product)

// Update / 업데이트
db.Update(ctx, "products",
    map[string]any{"stock": 95},
    "id = ?", productID)

// Delete / 삭제
db.Delete(ctx, "products", "id = ?", productID)
```

### Pattern 2: Pagination / 페이징

```go
page := 1
pageSize := 20
offset := (page - 1) * pageSize

// Get total count / 전체 수 가져오기
total, _ := db.Count(ctx, "users", "status = ?", "active")

// Get page data / 페이지 데이터 가져오기
users, _ := db.SelectWhere(ctx, "users", "status = ?", "active",
    mysql.WithOrderBy("created_at DESC"),
    mysql.WithLimit(pageSize),
    mysql.WithOffset(offset))

fmt.Printf("Page %d of %d\n", page, (total+int64(pageSize)-1)/int64(pageSize))
```

### Pattern 3: Bulk Insert / 대량 삽입

```go
// Using transaction for bulk insert / 대량 삽입을 위한 트랜잭션 사용
err := db.Transaction(ctx, func(tx *mysql.Tx) error {
    for _, user := range users {
        _, err := tx.Insert(ctx, "users", map[string]any{
            "name":  user.Name,
            "email": user.Email,
        })
        if err != nil {
            return err // Rollback all / 모두 롤백
        }
    }
    return nil // Commit all / 모두 커밋
})
```

### Pattern 4: Search with Multiple Conditions / 여러 조건으로 검색

```go
// Dynamic search / 동적 검색
conditions := []string{}
args := []interface{}{}

if nameFilter != "" {
    conditions = append(conditions, "name LIKE ?")
    args = append(args, "%"+nameFilter+"%")
}

if minAge > 0 {
    conditions = append(conditions, "age >= ?")
    args = append(args, minAge)
}

if city != "" {
    conditions = append(conditions, "city = ?")
    args = append(args, city)
}

// Build WHERE clause / WHERE 절 빌드
whereClause := strings.Join(conditions, " AND ")
allArgs := append([]interface{}{whereClause}, args...)

// Execute / 실행
users, _ := db.SelectAll(ctx, "users", allArgs...)
```

### Pattern 5: Aggregation with GROUP BY / GROUP BY를 사용한 집계

```go
// Count users by city / 도시별 사용자 수 계산
results, _ := db.SelectWhere(ctx, "users", "",
    mysql.WithColumns("city", "COUNT(*) as user_count", "AVG(age) as avg_age"),
    mysql.WithGroupBy("city"),
    mysql.WithHaving("COUNT(*) >= ?", 10),
    mysql.WithOrderBy("user_count DESC"))

for _, row := range results {
    fmt.Printf("City: %s, Users: %v, Avg Age: %v\n",
        row["city"], row["user_count"], row["avg_age"])
}
```

### Pattern 6: JOIN Operations / JOIN 작업

```go
// Query Builder for complex JOIN / 복잡한 JOIN을 위한 쿼리 빌더
results, _ := db.Select("u.name", "u.email", "o.id", "o.total", "o.status").
    From("users u").
    Join("orders o", "u.id = o.user_id").
    Where("o.created_at >= ?", startDate).
    Where("o.status IN (?, ?)", "pending", "completed").
    OrderBy("o.created_at DESC").
    Limit(100).
    All(ctx)

// Process results / 결과 처리
for _, row := range results {
    fmt.Printf("%s ordered $%.2f (Status: %s)\n",
        row["name"], row["total"], row["status"])
}
```

### Pattern 7: Conditional Updates / 조건부 업데이트

```go
// Update only if condition met / 조건이 충족되는 경우에만 업데이트
result, _ := db.Update(ctx, "inventory",
    map[string]any{"stock": db.Raw("stock - ?", quantity)},
    "product_id = ? AND stock >= ?", productID, quantity)

affected, _ := result.RowsAffected()
if affected == 0 {
    return errors.New("insufficient stock")
}
```

### Pattern 8: Soft Delete / 소프트 삭제

```go
// Soft delete (mark as deleted) / 소프트 삭제 (삭제로 표시)
db.Update(ctx, "users",
    map[string]any{
        "deleted_at": time.Now(),
        "status":     "deleted",
    },
    "id = ?", userID)

// Query excluding soft deleted / 소프트 삭제된 항목 제외하고 쿼리
activeUsers, _ := db.SelectAll(ctx, "users", "deleted_at IS NULL")
```

### Pattern 9: Upsert (Insert or Update) / Upsert (삽입 또는 업데이트)

```go
// Check if exists / 존재하는지 확인
exists, _ := db.Exists(ctx, "users", "email = ?", email)

if exists {
    // Update / 업데이트
    db.Update(ctx, "users",
        map[string]any{"last_seen": time.Now()},
        "email = ?", email)
} else {
    // Insert / 삽입
    db.Insert(ctx, "users", map[string]any{
        "email":     email,
        "last_seen": time.Now(),
    })
}
```

### Pattern 10: Hierarchical Data / 계층적 데이터

```go
// Get parent with children / 부모와 자식 가져오기
parent, _ := db.SelectOne(ctx, "categories", "id = ?", parentID)

children, _ := db.SelectAll(ctx, "categories", "parent_id = ?", parentID)

result := map[string]interface{}{
    "parent":   parent,
    "children": children,
}
```

---

## Common Use Cases / 일반적인 사용 사례

### Use Case 1: User Authentication / 사용자 인증

```go
package auth

import (
    "context"
    "errors"
    "golang.org/x/crypto/bcrypt"

    "github.com/arkd0ng/go-utils/database/mysql"
)

type AuthService struct {
    db *mysql.Client
}

func (s *AuthService) Register(ctx context.Context, email, password string) error {
    // Check if user exists / 사용자 존재 확인
    exists, _ := s.db.Exists(ctx, "users", "email = ?", email)
    if exists {
        return errors.New("email already registered")
    }

    // Hash password / 패스워드 해시
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

    // Create user / 사용자 생성
    _, err := s.db.Insert(ctx, "users", map[string]any{
        "email":    email,
        "password": string(hashedPassword),
        "status":   "active",
    })
    return err
}

func (s *AuthService) Login(ctx context.Context, email, password string) (map[string]interface{}, error) {
    // Get user / 사용자 가져오기
    user, err := s.db.SelectOne(ctx, "users", "email = ? AND status = ?", email, "active")
    if err != nil {
        return nil, errors.New("invalid credentials")
    }

    // Verify password / 패스워드 확인
    err = bcrypt.CompareHashAndPassword(
        []byte(user["password"].(string)),
        []byte(password),
    )
    if err != nil {
        return nil, errors.New("invalid credentials")
    }

    // Update last login / 마지막 로그인 업데이트
    s.db.Update(ctx, "users",
        map[string]any{"last_login": time.Now()},
        "id = ?", user["id"])

    // Remove sensitive data / 민감한 데이터 제거
    delete(user, "password")

    return user, nil
}
```

### Use Case 2: E-commerce Order System / 전자상거래 주문 시스템

```go
package orders

import (
    "context"
    "errors"
    "time"

    "github.com/arkd0ng/go-utils/database/mysql"
)

type OrderService struct {
    db *mysql.Client
}

func (s *OrderService) CreateOrder(ctx context.Context, userID int64, items []OrderItem) (int64, error) {
    var orderID int64

    err := s.db.Transaction(ctx, func(tx *mysql.Tx) error {
        // Create order / 주문 생성
        result, err := tx.Insert(ctx, "orders", map[string]any{
            "user_id":    userID,
            "status":     "pending",
            "created_at": time.Now(),
        })
        if err != nil {
            return err
        }

        orderID, _ = result.LastInsertId()

        // Add order items and update inventory / 주문 항목 추가 및 재고 업데이트
        for _, item := range items {
            // Check stock / 재고 확인
            product, err := tx.SelectOne(ctx, "products", "id = ?", item.ProductID)
            if err != nil {
                return err
            }

            stock := product["stock"].(int64)
            if stock < item.Quantity {
                return errors.New("insufficient stock")
            }

            // Insert order item / 주문 항목 삽입
            _, err = tx.Insert(ctx, "order_items", map[string]any{
                "order_id":   orderID,
                "product_id": item.ProductID,
                "quantity":   item.Quantity,
                "price":      product["price"],
            })
            if err != nil {
                return err
            }

            // Update stock / 재고 업데이트
            _, err = tx.Update(ctx, "products",
                map[string]any{"stock": stock - item.Quantity},
                "id = ?", item.ProductID)
            if err != nil {
                return err
            }
        }

        return nil // Commit / 커밋
    })

    return orderID, err
}

func (s *OrderService) GetOrderWithItems(ctx context.Context, orderID int64) (map[string]interface{}, error) {
    // Get order details / 주문 세부사항 가져오기
    order, err := s.db.SelectOne(ctx, "orders", "id = ?", orderID)
    if err != nil {
        return nil, err
    }

    // Get order items with product details / 제품 세부사항과 함께 주문 항목 가져오기
    items, _ := s.db.Select("oi.*, p.name", "p.description").
        From("order_items oi").
        Join("products p", "oi.product_id = p.id").
        Where("oi.order_id = ?", orderID).
        All(ctx)

    order["items"] = items
    return order, nil
}
```

### Use Case 3: Blog System / 블로그 시스템

```go
package blog

import (
    "context"
    "time"

    "github.com/arkd0ng/go-utils/database/mysql"
)

type BlogService struct {
    db *mysql.Client
}

func (s *BlogService) CreatePost(ctx context.Context, authorID int64, title, content string, tags []string) (int64, error) {
    var postID int64

    err := s.db.Transaction(ctx, func(tx *mysql.Tx) error {
        // Create post / 게시물 생성
        result, err := tx.Insert(ctx, "posts", map[string]any{
            "author_id":  authorID,
            "title":      title,
            "content":    content,
            "status":     "published",
            "created_at": time.Now(),
        })
        if err != nil {
            return err
        }

        postID, _ = result.LastInsertId()

        // Add tags / 태그 추가
        for _, tag := range tags {
            // Get or create tag / 태그 가져오기 또는 생성
            existingTag, err := tx.SelectOne(ctx, "tags", "name = ?", tag)
            var tagID int64

            if err != nil {
                // Create new tag / 새 태그 생성
                result, _ := tx.Insert(ctx, "tags", map[string]any{"name": tag})
                tagID, _ = result.LastInsertId()
            } else {
                tagID = existingTag["id"].(int64)
            }

            // Link post to tag / 게시물을 태그에 연결
            tx.Insert(ctx, "post_tags", map[string]any{
                "post_id": postID,
                "tag_id":  tagID,
            })
        }

        return nil
    })

    return postID, err
}

func (s *BlogService) GetPostsWithTags(ctx context.Context, page, pageSize int) ([]map[string]interface{}, error) {
    offset := (page - 1) * pageSize

    // Get posts / 게시물 가져오기
    posts, _ := s.db.SelectWhere(ctx, "posts", "status = ?", "published",
        mysql.WithOrderBy("created_at DESC"),
        mysql.WithLimit(pageSize),
        mysql.WithOffset(offset))

    // Get tags for each post / 각 게시물의 태그 가져오기
    for i, post := range posts {
        postID := post["id"].(int64)

        tags, _ := s.db.Select("t.name").
            From("tags t").
            Join("post_tags pt", "t.id = pt.tag_id").
            Where("pt.post_id = ?", postID).
            All(ctx)

        posts[i]["tags"] = tags
    }

    return posts, nil
}

func (s *BlogService) SearchPosts(ctx context.Context, keyword string) ([]map[string]interface{}, error) {
    // Full-text search / 전체 텍스트 검색
    return s.db.Select("id", "title", "content", "created_at").
        From("posts").
        Where("status = ?", "published").
        Where("(title LIKE ? OR content LIKE ?)", "%"+keyword+"%", "%"+keyword+"%").
        OrderBy("created_at DESC").
        Limit(50).
        All(ctx)
}
```

### Use Case 4: Analytics and Reporting / 분석 및 보고

```go
package analytics

import (
    "context"
    "time"

    "github.com/arkd0ng/go-utils/database/mysql"
)

type AnalyticsService struct {
    db *mysql.Client
}

func (s *AnalyticsService) GetDailyActiveUsers(ctx context.Context, days int) ([]map[string]interface{}, error) {
    startDate := time.Now().AddDate(0, 0, -days)

    return s.db.SelectWhere(ctx, "user_activities", "created_at >= ?", startDate,
        mysql.WithColumns("DATE(created_at) as date", "COUNT(DISTINCT user_id) as active_users"),
        mysql.WithGroupBy("DATE(created_at)"),
        mysql.WithOrderBy("date ASC"))
}

func (s *AnalyticsService) GetTopProducts(ctx context.Context, limit int) ([]map[string]interface{}, error) {
    return s.db.Select("p.id", "p.name", "COUNT(oi.id) as order_count", "SUM(oi.quantity) as total_sold").
        From("products p").
        Join("order_items oi", "p.id = oi.product_id").
        Join("orders o", "oi.order_id = o.id").
        Where("o.status = ?", "completed").
        GroupBy("p.id", "p.name").
        OrderBy("order_count DESC").
        Limit(limit).
        All(ctx)
}

func (s *AnalyticsService) GetRevenueByMonth(ctx context.Context, year int) ([]map[string]interface{}, error) {
    return s.db.SelectWhere(ctx, "orders", "YEAR(created_at) = ? AND status = ?", year, "completed",
        mysql.WithColumns("MONTH(created_at) as month", "SUM(total) as revenue"),
        mysql.WithGroupBy("MONTH(created_at)"),
        mysql.WithOrderBy("month ASC"))
}
```

### Use Case 5: Notification System / 알림 시스템

```go
package notifications

import (
    "context"
    "time"

    "github.com/arkd0ng/go-utils/database/mysql"
)

type NotificationService struct {
    db *mysql.Client
}

func (s *NotificationService) SendNotification(ctx context.Context, userID int64, notifType, title, message string) error {
    _, err := s.db.Insert(ctx, "notifications", map[string]any{
        "user_id":    userID,
        "type":       notifType,
        "title":      title,
        "message":    message,
        "read":       false,
        "created_at": time.Now(),
    })
    return err
}

func (s *NotificationService) GetUnreadNotifications(ctx context.Context, userID int64) ([]map[string]interface{}, error) {
    return s.db.SelectWhere(ctx, "notifications", "user_id = ? AND read = ?", userID, false,
        mysql.WithOrderBy("created_at DESC"),
        mysql.WithLimit(50))
}

func (s *NotificationService) MarkAsRead(ctx context.Context, notificationIDs []int64) error {
    // Bulk update / 대량 업데이트
    placeholders := strings.Repeat("?,", len(notificationIDs)-1) + "?"
    query := fmt.Sprintf("UPDATE notifications SET read = true WHERE id IN (%s)", placeholders)

    args := make([]interface{}, len(notificationIDs))
    for i, id := range notificationIDs {
        args[i] = id
    }

    _, err := s.db.Exec(ctx, query, args...)
    return err
}
```

---

## Best Practices / 모범 사례

### 1. Always Use Context / 항상 Context 사용

```go
// ✅ Good / 좋은 예
ctx := context.Background()
users, err := db.SelectAll(ctx, "users")

// ✅ Better - with timeout / 더 좋은 예 - 타임아웃 포함
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
users, err := db.SelectAll(ctx, "users")

// ❌ Bad - no context / 나쁜 예 - context 없음
// Not possible with this package / 이 패키지에서는 불가능
```

### 2. Close the Client When Done / 완료 시 클라이언트 닫기

```go
// ✅ Good / 좋은 예
db, err := mysql.New(mysql.WithDSN("..."))
if err != nil {
    log.Fatal(err)
}
defer db.Close()

// Application code...
```

### 3. Use Transactions for Related Operations / 관련 작업에 트랜잭션 사용

```go
// ✅ Good - atomic operations / 좋은 예 - 원자적 작업
err := db.Transaction(ctx, func(tx *mysql.Tx) error {
    tx.Insert(ctx, "orders", orderData)
    tx.Update(ctx, "inventory", stockUpdate)
    return nil
})

// ❌ Bad - non-atomic / 나쁜 예 - 비원자적
db.Insert(ctx, "orders", orderData)
db.Update(ctx, "inventory", stockUpdate) // May fail after insert / 삽입 후 실패할 수 있음
```

### 4. Handle Errors Properly / 에러 적절히 처리

```go
// ✅ Good / 좋은 예
user, err := db.SelectOne(ctx, "users", "id = ?", userID)
if err != nil {
    log.Printf("Failed to get user: %v", err)
    return err
}

// ❌ Bad - ignore errors / 나쁜 예 - 에러 무시
user, _ := db.SelectOne(ctx, "users", "id = ?", userID)
```

### 5. Use Prepared Statements (Automatic) / Prepared Statement 사용 (자동)

```go
// ✅ Good - parameterized query / 좋은 예 - 파라미터화된 쿼리
users, _ := db.SelectAll(ctx, "users", "email = ?", email)

// ❌ Bad - SQL injection risk / 나쁜 예 - SQL 인젝션 위험
users, _ := db.SelectAll(ctx, "users", fmt.Sprintf("email = '%s'", email))
```

### 6. Connection Pool Configuration / 연결 풀 설정

```go
// ✅ Good - proper pool settings / 좋은 예 - 적절한 풀 설정
db, _ := mysql.New(
    mysql.WithDSN("..."),
    mysql.WithMaxOpenConns(25),      // Based on MySQL max_connections / MySQL max_connections 기반
    mysql.WithMaxIdleConns(10),      // ~40% of MaxOpenConns
    mysql.WithConnMaxLifetime(5*time.Minute),
)

// ❌ Bad - default settings may not be optimal / 나쁜 예 - 기본 설정은 최적이 아닐 수 있음
db, _ := mysql.New(mysql.WithDSN("..."))
```

### 7. Choose the Right API / 적절한 API 선택

```go
// For simple queries / 간단한 쿼리용
users, _ := db.SelectAll(ctx, "users", "age > ?", 18)

// For queries with options / 옵션이 있는 쿼리용
users, _ := db.SelectWhere(ctx, "users", "age > ?", 18,
    mysql.WithOrderBy("name"),
    mysql.WithLimit(10))

// For complex JOINs / 복잡한 JOIN용
results, _ := db.Select("u.name", "o.total").
    From("users u").
    Join("orders o", "u.id = o.user_id").
    Where("o.status = ?", "completed").
    All(ctx)

// For complete control / 완전한 제어용
rows, _ := db.Query(ctx, "SELECT * FROM users WHERE ...")
```

### 8. Pagination Best Practices / 페이징 모범 사례

```go
// ✅ Good - efficient pagination / 좋은 예 - 효율적인 페이징
users, _ := db.SelectWhere(ctx, "users", "status = ?", "active",
    mysql.WithOrderBy("id DESC"),
    mysql.WithLimit(pageSize),
    mysql.WithOffset((page-1)*pageSize))

// ✅ Better - cursor-based pagination / 더 좋은 예 - 커서 기반 페이징
users, _ := db.SelectWhere(ctx, "users", "id < ? AND status = ?", lastID, "active",
    mysql.WithOrderBy("id DESC"),
    mysql.WithLimit(pageSize))
```

### 9. Avoid N+1 Query Problem / N+1 쿼리 문제 방지

```go
// ❌ Bad - N+1 queries / 나쁜 예 - N+1 쿼리
users, _ := db.SelectAll(ctx, "users")
for _, user := range users {
    orders, _ := db.SelectAll(ctx, "orders", "user_id = ?", user["id"])
    // Process orders...
}

// ✅ Good - single JOIN query / 좋은 예 - 단일 JOIN 쿼리
results, _ := db.Select("u.*, o.id as order_id", "o.total").
    From("users u").
    LeftJoin("orders o", "u.id = o.user_id").
    All(ctx)
```

### 10. Logging and Monitoring / 로깅 및 모니터링

```go
// Add logging wrapper / 로깅 래퍼 추가
func (s *Service) GetUser(ctx context.Context, id int64) (map[string]interface{}, error) {
    start := time.Now()
    defer func() {
        log.Printf("GetUser(%d) took %v", id, time.Since(start))
    }()

    return s.db.SelectOne(ctx, "users", "id = ?", id)
}
```

### 11. Soft Deletes / 소프트 삭제

```go
// ✅ Good - implement soft delete / 좋은 예 - 소프트 삭제 구현
db.Update(ctx, "users",
    map[string]any{"deleted_at": time.Now()},
    "id = ?", userID)

// Always filter out soft deleted / 항상 소프트 삭제된 항목 필터링
users, _ := db.SelectAll(ctx, "users", "deleted_at IS NULL")
```

### 12. Use Indexes / 인덱스 사용

```sql
-- Create indexes for frequently queried columns
-- 자주 쿼리되는 컬럼에 인덱스 생성
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_status_created ON users(status, created_at);
CREATE INDEX idx_orders_user_status ON orders(user_id, status);
```

### 13. Batch Operations / 배치 작업

```go
// For bulk inserts, use transaction / 대량 삽입은 트랜잭션 사용
err := db.Transaction(ctx, func(tx *mysql.Tx) error {
    for _, item := range items {
        tx.Insert(ctx, "items", item)
    }
    return nil
})

// Or use raw SQL for very large batches / 또는 매우 큰 배치는 raw SQL 사용
values := []string{}
args := []interface{}{}
for _, item := range items {
    values = append(values, "(?, ?, ?)")
    args = append(args, item.Name, item.Price, item.Stock)
}
query := "INSERT INTO items (name, price, stock) VALUES " + strings.Join(values, ",")
db.Exec(ctx, query, args...)
```

### 14. Connection Timeout / 연결 타임아웃

```go
// Set timeout in DSN / DSN에 타임아웃 설정
dsn := "user:password@tcp(localhost:3306)/database?timeout=5s&readTimeout=10s&writeTimeout=10s"

db, _ := mysql.New(mysql.WithDSN(dsn))
```

### 15. Testing / 테스트

```go
// Use test database / 테스트 데이터베이스 사용
func TestUserService(t *testing.T) {
    db, _ := mysql.New(
        mysql.WithDSN("user:password@tcp(localhost:3306)/testdb"),
    )
    defer db.Close()

    // Clean up test data / 테스트 데이터 정리
    defer db.Delete(context.Background(), "users", "email LIKE ?", "test_%")

    // Run tests...
}
```

---

## Troubleshooting / 문제 해결

### Problem 1: Connection Refused / 연결 거부

**Error / 에러**:
```
dial tcp [::1]:3306: connect: connection refused
```

**Solutions / 해결책**:
1. Check if MySQL is running / MySQL이 실행 중인지 확인:
   ```bash
   # macOS
   brew services list
   brew services start mysql

   # Linux
   sudo systemctl status mysql
   sudo systemctl start mysql
   ```

2. Check host and port / 호스트와 포트 확인:
   ```go
   // Try different addresses / 다른 주소 시도
   "user:password@tcp(localhost:3306)/database"  // localhost
   "user:password@tcp(127.0.0.1:3306)/database"  // IP address
   ```

3. Check firewall / 방화벽 확인

### Problem 2: Access Denied / 액세스 거부

**Error / 에러**:
```
Error 1045: Access denied for user 'user'@'localhost'
```

**Solutions / 해결책**:
1. Check credentials / 자격 증명 확인:
   ```bash
   mysql -u user -p
   ```

2. Grant permissions / 권한 부여:
   ```sql
   GRANT ALL PRIVILEGES ON database.* TO 'user'@'localhost';
   FLUSH PRIVILEGES;
   ```

3. Check MySQL user host / MySQL 사용자 호스트 확인:
   ```sql
   SELECT user, host FROM mysql.user;
   ```

### Problem 3: Too Many Connections / 연결이 너무 많음

**Error / 에러**:
```
Error 1040: Too many connections
```

**Solutions / 해결책**:
1. Increase MySQL max_connections / MySQL max_connections 증가:
   ```sql
   SET GLOBAL max_connections = 200;
   ```

2. Reduce application connection pool / 애플리케이션 연결 풀 감소:
   ```go
   db, _ := mysql.New(
       mysql.WithDSN("..."),
       mysql.WithMaxOpenConns(10),  // Reduce / 감소
   )
   ```

3. Close unused connections / 미사용 연결 닫기:
   ```go
   defer db.Close()
   ```

### Problem 4: Deadlock / 교착상태

**Error / 에러**:
```
Error 1213: Deadlock found when trying to get lock
```

**Solutions / 해결책**:
1. Retry transaction / 트랜잭션 재시도 (automatic in this package / 이 패키지에서 자동)

2. Order operations consistently / 작업 순서 일관성 유지:
   ```go
   // Always update in same order / 항상 같은 순서로 업데이트
   tx.Update(ctx, "table1", ...)
   tx.Update(ctx, "table2", ...)
   ```

3. Keep transactions short / 트랜잭션을 짧게 유지

### Problem 5: Slow Queries / 느린 쿼리

**Solutions / 해결책**:
1. Add indexes / 인덱스 추가:
   ```sql
   CREATE INDEX idx_users_email ON users(email);
   ```

2. Use EXPLAIN / EXPLAIN 사용:
   ```sql
   EXPLAIN SELECT * FROM users WHERE email = 'john@example.com';
   ```

3. Enable slow query log / 느린 쿼리 로그 활성화:
   ```sql
   SET GLOBAL slow_query_log = 'ON';
   SET GLOBAL long_query_time = 2;
   ```

4. Optimize queries / 쿼리 최적화:
   ```go
   // Select only needed columns / 필요한 컬럼만 선택
   users, _ := db.SelectWhere(ctx, "users", "age > ?", 18,
       mysql.WithColumns("id", "name"))  // Not SELECT *
   ```

### Problem 6: Connection Lost / 연결 손실

**Error / 에러**:
```
invalid connection
```

**Solution / 해결책**:
Auto reconnect is built-in / 자동 재연결이 내장되어 있습니다:
```go
// Package handles this automatically / 패키지가 자동으로 처리
users, _ := db.SelectAll(ctx, "users")
// If connection is lost, will retry / 연결이 손실되면 재시도
```

### Problem 7: Context Deadline Exceeded / Context 데드라인 초과

**Error / 에러**:
```
context deadline exceeded
```

**Solutions / 해결책**:
1. Increase timeout / 타임아웃 증가:
   ```go
   ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
   defer cancel()
   ```

2. Optimize query / 쿼리 최적화

3. Check network latency / 네트워크 지연 확인

### Problem 8: Memory Issues with Large Result Sets / 큰 결과 세트의 메모리 문제

**Solution / 해결책**:
Use pagination or streaming / 페이징 또는 스트리밍 사용:
```go
// Pagination / 페이징
pageSize := 1000
for page := 1; ; page++ {
    users, _ := db.SelectWhere(ctx, "users", "",
        mysql.WithLimit(pageSize),
        mysql.WithOffset((page-1)*pageSize))

    if len(users) == 0 {
        break
    }

    // Process batch / 배치 처리
    processBatch(users)
}
```

### Problem 9: Character Encoding Issues / 문자 인코딩 문제

**Solution / 해결책**:
Set charset in DSN / DSN에 charset 설정:
```go
dsn := "user:password@tcp(localhost:3306)/database?charset=utf8mb4&parseTime=true"
db, _ := mysql.New(mysql.WithDSN(dsn))
```

### Problem 10: Time Zone Issues / 시간대 문제

**Solution / 해결책**:
Set time zone in DSN / DSN에 시간대 설정:
```go
dsn := "user:password@tcp(localhost:3306)/database?parseTime=true&loc=Local"
db, _ := mysql.New(mysql.WithDSN(dsn))
```

---

## FAQ

### Q1: Do I need to call `defer rows.Close()`?
### Q1: `defer rows.Close()`를 호출해야 하나요?

**A**: No! This package handles resource cleanup automatically. You only need `defer rows.Close()` when using the raw `Query()` method.

**A**: 아니요! 이 패키지는 리소스 정리를 자동으로 처리합니다. raw `Query()` 메서드를 사용할 때만 `defer rows.Close()`가 필요합니다.

```go
// ✅ No defer needed / defer 불필요
users, _ := db.SelectAll(ctx, "users")

// ⚠️ Defer needed for raw queries / raw 쿼리는 defer 필요
rows, _ := db.Query(ctx, "SELECT * FROM users")
defer rows.Close()
```

---

### Q2: How do I handle NULL values?
### Q2: NULL 값은 어떻게 처리하나요?

**A**: NULL values are returned as `nil` in the result map.

**A**: NULL 값은 결과 맵에서 `nil`로 반환됩니다.

```go
user, _ := db.SelectOne(ctx, "users", "id = ?", 123)

if user["middle_name"] == nil {
    fmt.Println("No middle name")
} else {
    fmt.Printf("Middle name: %s\n", user["middle_name"])
}
```

---

### Q3: Can I use this with existing `database/sql` code?
### Q3: 기존 `database/sql` 코드와 함께 사용할 수 있나요?

**A**: Yes! You can mix this package with standard `database/sql`:

**A**: 네! 이 패키지를 표준 `database/sql`과 혼합하여 사용할 수 있습니다:

```go
// Use simple API for simple queries / 간단한 쿼리는 simple API 사용
users, _ := db.SelectAll(ctx, "users")

// Use raw SQL for complex queries / 복잡한 쿼리는 raw SQL 사용
rows, _ := db.Query(ctx, "WITH ...")
```

---

### Q4: How do I handle transactions that span multiple tables?
### Q4: 여러 테이블에 걸친 트랜잭션은 어떻게 처리하나요?

**A**: Use the `Transaction()` method with all operations inside:

**A**: 모든 작업을 내부에 포함하는 `Transaction()` 메서드를 사용하세요:

```go
err := db.Transaction(ctx, func(tx *mysql.Tx) error {
    tx.Insert(ctx, "table1", data1)
    tx.Update(ctx, "table2", data2, "id = ?", id)
    tx.Delete(ctx, "table3", "id = ?", id)
    return nil // All commit together / 모두 함께 커밋
})
```

---

### Q5: What happens if a query fails due to network issues?
### Q5: 네트워크 문제로 쿼리가 실패하면 어떻게 되나요?

**A**: The package automatically retries transient errors with exponential backoff. You can configure retry behavior:

**A**: 패키지는 지수 백오프로 일시적 에러를 자동으로 재시도합니다. 재시도 동작을 설정할 수 있습니다:

```go
db, _ := mysql.New(
    mysql.WithDSN("..."),
    mysql.WithRetryMaxAttempts(5),
    mysql.WithRetryMaxInterval(2*time.Second),
)
```

---

### Q6: Can I use this package in production?
### Q6: 프로덕션에서 이 패키지를 사용할 수 있나요?

**A**: Yes! The package is production-ready and includes:
- Auto reconnection
- Auto retry
- Connection pooling
- Transaction support
- Comprehensive error handling

**A**: 네! 패키지는 프로덕션 준비가 완료되었으며 다음을 포함합니다:
- 자동 재연결
- 자동 재시도
- 연결 풀링
- 트랜잭션 지원
- 포괄적인 에러 처리

---

### Q7: How do I debug SQL queries?
### Q7: SQL 쿼리를 디버그하려면 어떻게 하나요?

**A**: Enable MySQL general log or use a wrapper with logging:

**A**: MySQL 일반 로그를 활성화하거나 로깅이 있는 래퍼를 사용하세요:

```sql
-- Enable general log / 일반 로그 활성화
SET GLOBAL general_log = 'ON';
SET GLOBAL log_output = 'TABLE';

-- View queries / 쿼리 보기
SELECT * FROM mysql.general_log ORDER BY event_time DESC LIMIT 10;
```

---

### Q8: What's the difference between Query Builder and SelectWhere?
### Q8: Query Builder와 SelectWhere의 차이점은 무엇인가요?

**A**:
- **Query Builder**: Fluent API, better for complex JOINs, IDE autocomplete
- **SelectWhere**: Functional options, one-liner queries, simpler for moderate complexity

**A**:
- **Query Builder**: Fluent API, 복잡한 JOIN에 더 좋음, IDE 자동완성
- **SelectWhere**: 함수형 옵션, 한 줄 쿼리, 중간 복잡도에 더 간단

```go
// Query Builder / 쿼리 빌더
users, _ := db.Select("name").
    From("users").
    Where("age > ?", 18).
    All(ctx)

// SelectWhere
users, _ := db.SelectWhere(ctx, "users", "age > ?", 18,
    mysql.WithColumns("name"))
```

---

### Q9: How do I handle connection pooling?
### Q9: 연결 풀링은 어떻게 처리하나요?

**A**: Connection pooling is automatic. Configure it during initialization:

**A**: 연결 풀링은 자동입니다. 초기화 중에 설정하세요:

```go
db, _ := mysql.New(
    mysql.WithDSN("..."),
    mysql.WithMaxOpenConns(25),
    mysql.WithMaxIdleConns(10),
    mysql.WithConnMaxLifetime(5*time.Minute),
)
```

---

### Q10: Can I use this with Docker containers?
### Q10: Docker 컨테이너와 함께 사용할 수 있나요?

**A**: Yes! Just point to the correct host:

**A**: 네! 올바른 호스트를 가리키기만 하면 됩니다:

```go
// Docker Compose / Docker Compose
dsn := "user:password@tcp(mysql:3306)/database"

// Docker with port mapping / 포트 매핑이 있는 Docker
dsn := "user:password@tcp(localhost:3307)/database"

db, _ := mysql.New(mysql.WithDSN(dsn))
```

---

### Q11: How do I migrate from `database/sql`?
### Q11: `database/sql`에서 마이그레이션하려면 어떻게 하나요?

**A**: Replace common patterns:

**A**: 일반적인 패턴을 교체하세요:

```go
// Before (database/sql) / 이전 (database/sql)
rows, _ := db.Query("SELECT * FROM users WHERE age > ?", 18)
defer rows.Close()
// ... scanning logic ...

// After (this package) / 이후 (이 패키지)
users, _ := db.SelectAll(ctx, "users", "age > ?", 18)
```

---

### Q12: How do I handle large file uploads (BLOB)?
### Q12: 큰 파일 업로드(BLOB)는 어떻게 처리하나요?

**A**: Use raw SQL for BLOB operations:

**A**: BLOB 작업에는 raw SQL을 사용하세요:

```go
// Insert BLOB / BLOB 삽입
fileData, _ := ioutil.ReadFile("image.jpg")
db.Exec(ctx, "INSERT INTO files (name, data) VALUES (?, ?)", "image.jpg", fileData)

// Read BLOB / BLOB 읽기
rows, _ := db.Query(ctx, "SELECT data FROM files WHERE id = ?", fileID)
defer rows.Close()
rows.Next()
var data []byte
rows.Scan(&data)
```

---

### Q13: Is this package thread-safe?
### Q13: 이 패키지는 스레드 안전한가요?

**A**: Yes! The package is safe for concurrent use across multiple goroutines.

**A**: 네! 패키지는 여러 고루틴에서 동시 사용에 안전합니다.

```go
// Safe to use concurrently / 동시 사용에 안전
db, _ := mysql.New(mysql.WithDSN("..."))

go func() {
    db.SelectAll(ctx, "users")
}()

go func() {
    db.Insert(ctx, "users", data)
}()
```

---

### Q14: How do I handle database migrations?
### Q14: 데이터베이스 마이그레이션은 어떻게 처리하나요?

**A**: Use a migration tool like `golang-migrate` or run SQL directly:

**A**: `golang-migrate` 같은 마이그레이션 도구를 사용하거나 SQL을 직접 실행하세요:

```go
// Run migration SQL / 마이그레이션 SQL 실행
db.Exec(ctx, `
    CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        email VARCHAR(255) UNIQUE NOT NULL
    )
`)
```

---

### Q15: Where can I get help?
### Q15: 어디서 도움을 받을 수 있나요?

**A**:
- GitHub Issues: https://github.com/arkd0ng/go-utils/issues
- Documentation: Check the DEVELOPER_GUIDE.md for advanced topics
- Examples: See `examples/mysql/main.go` for working examples

**A**:
- GitHub Issues: https://github.com/arkd0ng/go-utils/issues
- 문서: 고급 주제는 DEVELOPER_GUIDE.md 확인
- 예제: 작동 예제는 `examples/mysql/main.go` 참조

---

**End of User Manual / 사용자 매뉴얼 끝**
