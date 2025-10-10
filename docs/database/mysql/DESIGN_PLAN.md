# Database/MySQL Package - Design Plan / 설계 계획서
# database/mysql 패키지 - 설계 계획서

**Version / 버전**: v1.3.x
**Author / 작성자**: arkd0ng
**Created / 작성일**: 2025-10-10
**Status / 상태**: Final Design - Extreme Simplicity / 최종 설계 - 극도의 간결함

---

## Table of Contents / 목차

1. [Why This Package Exists / 왜 이 패키지가 존재하는가](#why-this-package-exists--왜-이-패키지가-존재하는가)
2. [Design Philosophy / 설계 철학](#design-philosophy--설계-철학)
3. [What Users Get / 사용자가 얻는 것](#what-users-get--사용자가-얻는-것)
4. [API Design / API 설계](#api-design--api-설계)
5. [Implementation Architecture / 구현 아키텍처](#implementation-architecture--구현-아키텍처)
6. [File Structure / 파일 구조](#file-structure--파일-구조)
7. [Detailed Features / 상세 기능](#detailed-features--상세-기능)

---

## Why This Package Exists / 왜 이 패키지가 존재하는가

### The Problem / 문제점

Using the standard `database/sql` package with MySQL requires developers to:

표준 `database/sql` 패키지를 MySQL과 함께 사용하려면 개발자가 다음을 해야 합니다:

1. **Manually manage connections / 수동으로 연결 관리**:
   ```go
   // 매번 연결 상태 확인
   if err := db.Ping(); err != nil {
       // 재접속 로직 작성
       db, err = sql.Open(...)
       // ...
   }
   ```

2. **Manually manage resources / 수동으로 리소스 관리**:
   ```go
   rows, err := db.Query("SELECT * FROM users WHERE id = ?", 123)
   if err != nil {
       return err
   }
   defer rows.Close() // ← 개발자가 매번 기억해야 함!

   for rows.Next() {
       // scan logic...
   }
   if err := rows.Err(); err != nil {
       return err
   }
   ```

3. **Write complex SQL for simple operations / 간단한 작업에 복잡한 SQL 작성**:
   ```go
   // INSERT
   _, err := db.Exec("INSERT INTO users (name, email, age) VALUES (?, ?, ?)",
       "John", "john@example.com", 30)

   // UPDATE
   _, err := db.Exec("UPDATE users SET name = ?, age = ? WHERE id = ?",
       "Jane", 31, 123)

   // 매번 SQL 문법 작성, 컬럼 순서 기억, ? 개수 맞추기...
   ```

4. **Handle connection loss manually / 연결 손실 수동 처리**:
   ```go
   // "MySQL server has gone away" 에러 발생 시
   // 개발자가 직접 재시도 로직 작성
   ```

### The Solution / 해결책

**이 패키지는 위의 모든 번거로움을 제거합니다**:

```go
// 1. 한 번 연결하면 끝 - 자동으로 계속 유지됨
db, _ := mysql.New(mysql.WithDSN("..."))

// 2. SQL 문법에 가깝게 - 간단하게
users, _ := db.SelectAll("users", "id = ?", 123)

// 3. INSERT/UPDATE/DELETE - map으로 간단하게
db.Insert("users", map[string]interface{}{
    "name":  "John",
    "email": "john@example.com",
    "age":   30,
})

// 4. 연결 끊김? 자동 재접속됨
// 5. defer rows.Close()? 필요 없음, 자동 처리됨
// 6. 에러 처리? 재시도 가능한 에러는 자동 재시도됨
```

### If It's Not This Simple, Don't Build It / 이 정도로 간단하지 않으면 만들지 마세요

**핵심 원칙**:
- 기존 방법보다 10배 간단하지 않으면 의미 없음
- 개발자가 DB/연결/SQL 문법을 신경 쓰지 않아야 함
- 비즈니스 로직만 작성하면 됨

---

## Design Philosophy / 설계 철학

### 1. Zero Mental Overhead / 제로 정신적 부담

**개발자는 다음을 신경 쓰지 않습니다 / 개발자가 신경 쓰지 않는 것**:
- ❌ 연결이 살아있는지
- ❌ 언제 재접속해야 하는지
- ❌ defer rows.Close()를 해야 하는지
- ❌ SQL 문법의 정확한 순서
- ❌ ? placeholder 개수
- ❌ 에러가 재시도 가능한지

**개발자가 신경 쓰는 것**:
- ✅ 어떤 데이터를 가져올 것인가
- ✅ 어떤 데이터를 저장할 것인가
- ✅ 비즈니스 로직

### 2. SQL-Like, Not ORM / SQL 같지만 ORM 아님

```go
// ✅ SQL 친화적 - SQL 문법에 가까움
db.SelectAll("users", "age > ? AND active = ?", 18, true)
// SELECT * FROM users WHERE age > ? AND active = ?

db.Insert("users", map[string]interface{}{
    "name": "John",
    "age":  30,
})
// INSERT INTO users (name, age) VALUES (?, ?)

// ❌ ORM 같은 복잡한 것은 하지 않음
// User{}.Where(...).Find(...) ← 이런 거 안 함
```

### 3. Auto Everything / 모든 것이 자동

```go
// 사용자 코드
db, _ := mysql.New(mysql.WithDSN("..."))
users, _ := db.SelectAll("users", "id = ?", 123)

// 내부적으로 자동으로 발생:
// 1. 연결 상태 확인
// 2. 필요시 재접속
// 3. 쿼리 실행
// 4. 에러 발생 시 재시도 (transient errors)
// 5. rows.Close() 자동 처리
// 6. 느린 쿼리 자동 로깅
// 7. 결과를 사용하기 쉬운 형태로 반환
```

### 4. Progressive Disclosure / 점진적 노출

**90% 사용 사례 - 초간단**:
```go
users, _ := db.SelectAll("users", "age > ?", 18)
```

**복잡한 경우 - 여전히 간단**:
```go
users, _ := db.Select("id", "name").
    From("users").
    Where("age > ?", 18).
    Where("active = ?", true).
    OrderBy("name").
    Limit(10).
    All()
```

**매우 복잡한 경우 - 기존 방법 사용**:
```go
rows, _ := db.Query("복잡한 SQL...")
// 여전히 자동 재접속, 자동 재시도는 됨
```

---

## What Users Get / 사용자가 얻는 것

### Before (기존 방법) vs After (이 패키지)

#### Example 1: Simple SELECT / 간단한 조회

**Before / 기존**:
```go
// 30줄
db, err := sql.Open("mysql", dsn)
if err != nil {
    return nil, err
}
defer db.Close()

if err := db.Ping(); err != nil {
    return nil, err
}

rows, err := db.Query("SELECT * FROM users WHERE age > ?", 18)
if err != nil {
    return nil, err
}
defer rows.Close()

var users []User
for rows.Next() {
    var u User
    if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age); err != nil {
        return nil, err
    }
    users = append(users, u)
}

if err := rows.Err(); err != nil {
    return nil, err
}

return users, nil
```

**After / 이 패키지**:
```go
// 2줄
db, _ := mysql.New(mysql.WithDSN(dsn))
users, _ := db.SelectAll("users", "age > ?", 18)
```

#### Example 2: INSERT / 삽입

**Before / 기존**:
```go
// 10줄
result, err := db.Exec(
    "INSERT INTO users (name, email, age, created_at) VALUES (?, ?, ?, ?)",
    "John", "john@example.com", 30, time.Now(),
)
if err != nil {
    return err
}
userID, err := result.LastInsertId()
// ...
```

**After / 이 패키지**:
```go
// 4줄
result, _ := db.Insert("users", map[string]interface{}{
    "name":  "John",
    "email": "john@example.com",
    "age":   30,
})
```

#### Example 3: Transaction / 트랜잭션

**Before / 기존**:
```go
// 20줄+
tx, err := db.Begin()
if err != nil {
    return err
}
defer func() {
    if p := recover(); p != nil {
        tx.Rollback()
        panic(p)
    } else if err != nil {
        tx.Rollback()
    } else {
        err = tx.Commit()
    }
}()

_, err = tx.Exec("INSERT INTO users (name) VALUES (?)", "John")
if err != nil {
    return err
}
_, err = tx.Exec("UPDATE accounts SET balance = ? WHERE user_id = ?", 100, userID)
if err != nil {
    return err
}
// ...
```

**After / 이 패키지**:
```go
// 6줄
err := db.Transaction(func(tx *mysql.Tx) error {
    tx.Insert("users", map[string]interface{}{"name": "John"})
    tx.Update("accounts", map[string]interface{}{"balance": 100}, "user_id = ?", userID)
    return nil // 자동 커밋, 에러 시 자동 롤백
})
```

### Code Reduction / 코드 감소

- **일반 쿼리**: ~90% 코드 감소
- **INSERT/UPDATE**: ~70% 코드 감소
- **트랜잭션**: ~70% 코드 감소
- **에러 처리**: ~100% 감소 (자동 처리)
- **리소스 관리**: ~100% 감소 (자동 처리)

---

## API Design / API 설계

### Layer 1: Simple API (90% 사용 사례)

#### SELECT Operations / 조회 작업

```go
// 모든 행 가져오기 / Get all rows
users, err := db.SelectAll("users", "age > ?", 18)
// SELECT * FROM users WHERE age > ?

// 한 행 가져오기 / Get one row
user, err := db.SelectOne("users", "id = ?", 123)
// SELECT * FROM users WHERE id = ? LIMIT 1

// 특정 컬럼만 / Specific columns
users, err := db.SelectColumns([]string{"id", "name"}, "users", "active = ?", true)
// SELECT id, name FROM users WHERE active = ?

// 개수 세기 / Count
count, err := db.Count("users", "age > ?", 18)
// SELECT COUNT(*) FROM users WHERE age > ?

// 존재 확인 / Check existence
exists, err := db.Exists("users", "email = ?", "test@example.com")
// SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)

// WHERE 없이 전체 조회 / All rows without WHERE
users, err := db.SelectAll("users", "")
// SELECT * FROM users
```

#### INSERT Operations / 삽입 작업

```go
// 단일 행 삽입 / Insert single row
result, err := db.Insert("users", map[string]interface{}{
    "name":  "John Doe",
    "email": "john@example.com",
    "age":   30,
})
// INSERT INTO users (name, email, age) VALUES (?, ?, ?)

// 마지막 삽입 ID / Last insert ID
lastID, _ := result.LastInsertId()

// 여러 컬럼 / Multiple columns
result, err := db.Insert("users", map[string]interface{}{
    "name":       "John",
    "email":      "john@example.com",
    "age":        30,
    "created_at": time.Now(),
    "active":     true,
})
```

#### UPDATE Operations / 업데이트 작업

```go
// 행 업데이트 / Update rows
result, err := db.Update("users",
    map[string]interface{}{
        "name": "Jane Doe",
        "age":  31,
    },
    "id = ?", 123,
)
// UPDATE users SET name = ?, age = ? WHERE id = ?

// 영향받은 행 수 / Affected rows
affected, _ := result.RowsAffected()

// 여러 행 업데이트 / Update multiple rows
result, err := db.Update("users",
    map[string]interface{}{"active": false},
    "last_login < ?", time.Now().AddDate(0, -6, 0),
)
// UPDATE users SET active = ? WHERE last_login < ?
```

#### DELETE Operations / 삭제 작업

```go
// 행 삭제 / Delete rows
result, err := db.Delete("users", "age < ?", 18)
// DELETE FROM users WHERE age < ?

// 단일 행 삭제 / Delete single row
result, err := db.Delete("users", "id = ?", 123)
// DELETE FROM users WHERE id = ?

// ⚠️ 전체 삭제 (조심!) / Delete all (careful!)
result, err := db.DeleteAll("users")
// DELETE FROM users
```

### Layer 2: Query Builder (복잡한 쿼리용)

```go
// Fluent API
users, err := db.Select("id", "name", "email").
    From("users").
    Where("age > ?", 18).
    Where("active = ?", true).
    OrderBy("name ASC").
    Limit(10).
    Offset(20).
    All()

// Single row
user, err := db.Select("*").
    From("users").
    Where("id = ?", 123).
    One()

// JOIN
results, err := db.Select("u.name", "o.total").
    From("users u").
    Join("orders o", "u.id = o.user_id").
    Where("o.status = ?", "completed").
    OrderBy("o.created_at DESC").
    All()

// GROUP BY, HAVING
results, err := db.Select("category", "COUNT(*) as count").
    From("products").
    GroupBy("category").
    Having("COUNT(*) > ?", 10).
    All()
```

### Layer 3: Raw SQL (고급 사용자용)

```go
// 완전한 제어가 필요할 때
rows, err := db.Query("SELECT * FROM users WHERE MATCH(name) AGAINST(?)", "John")
// 여전히 자동 재접속, 자동 재시도됨

result, err := db.Exec("UPDATE users SET last_login = NOW() WHERE id = ?", 123)

// 복잡한 쿼리
rows, err := db.Query(`
    WITH ranked_users AS (
        SELECT *, ROW_NUMBER() OVER (PARTITION BY country ORDER BY score DESC) as rank
        FROM users
    )
    SELECT * FROM ranked_users WHERE rank <= 10
`)
```

### Transaction API / 트랜잭션 API

```go
// 간단한 트랜잭션 / Simple transaction
err := db.Transaction(func(tx *mysql.Tx) error {
    // 모든 간단한 API 사용 가능 / All simple APIs available
    _, err := tx.Insert("users", map[string]interface{}{
        "name": "John",
    })
    if err != nil {
        return err // 자동 롤백 / Auto rollback
    }

    _, err = tx.Update("accounts",
        map[string]interface{}{"balance": 1000},
        "user_id = ?", 123,
    )
    if err != nil {
        return err // 자동 롤백 / Auto rollback
    }

    return nil // 자동 커밋 / Auto commit
})

// Query Builder도 사용 가능 / Query Builder also available
err := db.Transaction(func(tx *mysql.Tx) error {
    users, err := tx.Select("*").
        From("users").
        Where("id = ?", 123).
        All()
    // ...
    return nil
})
```

---

## Implementation Architecture / 구현 아키텍처

### High-Level Flow / 상위 수준 흐름

```
┌──────────────────────────────────────────────────────────┐
│                    User Code                              │
│                    사용자 코드                             │
│                                                           │
│   db.SelectAll("users", "id = ?", 123)                   │
└──────────────────────┬───────────────────────────────────┘
                       │
                       ▼
┌──────────────────────────────────────────────────────────┐
│                 ensureConnected()                         │
│                 연결 확인                                  │
│                                                           │
│  ┌─────────────────────────────────────────────┐        │
│  │ Is connection healthy?                       │        │
│  │ YES → Continue                               │        │
│  │ NO  → Ping database                          │        │
│  │       Failed? → Reconnect                    │        │
│  └─────────────────────────────────────────────┘        │
└──────────────────────┬───────────────────────────────────┘
                       │
                       ▼
┌──────────────────────────────────────────────────────────┐
│              executeWithRetry()                           │
│              재시도 로직                                    │
│                                                           │
│  ┌─────────────────────────────────────────────┐        │
│  │ Attempt 1: Execute query                    │        │
│  │   Success? → Return result                  │        │
│  │   Transient error? → Retry with backoff     │        │
│  │   Permanent error? → Return error           │        │
│  │ Attempt 2, 3...                              │        │
│  └─────────────────────────────────────────────┘        │
└──────────────────────┬───────────────────────────────────┘
                       │
                       ▼
┌──────────────────────────────────────────────────────────┐
│                Execute Query                              │
│                쿼리 실행                                    │
│                                                           │
│  1. Build SQL from parameters                            │
│  2. Execute with prepared statement                      │
│  3. Scan results (auto defer rows.Close())               │
│  4. Log query (if slow or enabled)                       │
│  5. Return results                                       │
└──────────────────────────────────────────────────────────┘
```

### Core Components / 핵심 컴포넌트

```go
// Client - Main database client / 메인 데이터베이스 클라이언트
type Client struct {
    // Connection pool / 연결 풀
    connections      []*sql.DB         // 여러 개의 연결 (credential rotation용)
    currentIdx       int               // 현재 사용 중인 연결 인덱스 (round-robin)
    rotationIdx      int               // 순환 교체 인덱스
    connectionsMu    sync.RWMutex      // 연결 배열 동기화

    // Configuration / 설정
    config           *config           // 설정

    // State / 상태
    logger           *logging.Logger   // 로거 (선택)
    healthy          bool              // 연결 상태

    // Background tasks / 백그라운드 작업
    stopChan         chan struct{}     // 종료 신호
    healthCheckStop  chan struct{}     // 헬스 체크 중지
    rotationStop     chan struct{}     // 순환 중지

    // Synchronization / 동기화
    mu               sync.RWMutex      // 일반 동기화
}

// CredentialRefreshFunc - User-provided function to get new DSN
// CredentialRefreshFunc - 사용자가 제공하는 새 DSN 가져오기 함수
type CredentialRefreshFunc func() (dsn string, error)

// config - 설정 구조체
type config struct {
    // Basic connection / 기본 연결
    dsn             string

    // Connection pool / 연결 풀
    maxOpenConns    int
    maxIdleConns    int
    connMaxLifetime time.Duration
    connMaxIdleTime time.Duration

    // Credential rotation (optional) / 자격 증명 순환 (선택)
    credRefreshFunc    CredentialRefreshFunc  // 사용자 제공 함수
    poolCount          int                     // 연결 풀 개수
    rotationInterval   time.Duration           // 교체 주기

    // Timeout / 타임아웃
    connectTimeout  time.Duration
    queryTimeout    time.Duration

    // Retry / 재시도
    maxRetries      int
    retryDelay      time.Duration

    // Logging / 로깅
    logger          *logging.Logger
}

// 핵심 메서드들
func (c *Client) ensureConnected() error
func (c *Client) reconnect() error
func (c *Client) executeWithRetry(fn func() error) error
func (c *Client) scanRows(rows *sql.Rows) ([]map[string]interface{}, error)
func (c *Client) startHealthCheck()
func (c *Client) startConnectionRotation()  // 새로운 기능!
func (c *Client) getCurrentConnection() *sql.DB
func (c *Client) rotateConnection() error
```

### Auto-Management Flow / 자동 관리 흐름

```go
// 모든 쿼리 실행 시
func (c *Client) SelectAll(table, condition string, args ...interface{}) ([]map[string]interface{}, error) {
    // 1️⃣ 연결 확인 및 필요시 재접속
    if err := c.ensureConnected(); err != nil {
        return nil, err
    }

    // 2️⃣ 재시도 로직으로 실행
    var results []map[string]interface{}
    err := c.executeWithRetry(func() error {
        // 3️⃣ SQL 빌드
        query := fmt.Sprintf("SELECT * FROM %s", table)
        if condition != "" {
            query += " WHERE " + condition
        }

        // 4️⃣ 쿼리 실행
        start := time.Now()
        rows, err := c.db.Query(query, args...)
        if err != nil {
            return err
        }
        defer rows.Close() // 5️⃣ 자동 정리

        // 6️⃣ 결과 스캔
        results, err = c.scanRows(rows)

        // 7️⃣ 로깅
        c.logQuery(query, args, time.Since(start), err)

        return err
    })

    return results, err
}
```

---

## File Structure / 파일 구조

```
database/mysql/
├── client.go          # Client struct, New(), Close()
├── connection.go      # ensureConnected(), reconnect(), startHealthCheck()
├── rotation.go        # startConnectionRotation(), rotateOneConnection() (credential rotation)
├── simple.go          # SelectAll, SelectOne, Insert, Update, Delete
├── builder.go         # Query builder (Select().From().Where()...)
├── transaction.go     # Transaction support
├── retry.go           # executeWithRetry(), isRetryableError()
├── scan.go            # scanRows(), type conversion
├── config.go          # config struct, defaults
├── options.go         # WithDSN, WithCredentialRefresh, WithLogger, etc.
├── errors.go          # Custom error types
├── types.go           # Common types (CredentialRefreshFunc, Tx struct)
├── client_test.go     # Unit tests
├── rotation_test.go   # Credential rotation tests
└── README.md          # Package documentation
```

### Key Files Breakdown / 주요 파일 설명

**client.go** - 메인 클라이언트
```go
type Client struct { ... }
func New(opts ...Option) (*Client, error)
func (c *Client) Close() error
```

**connection.go** - 연결 관리 (자동)
```go
func (c *Client) ensureConnected() error
func (c *Client) reconnect() error
func (c *Client) startHealthCheck()
```

**rotation.go** - 연결 순환 (credential rotation 기능)
```go
func (c *Client) startConnectionRotation()       // rotation goroutine 시작
func (c *Client) rotateOneConnection() error     // 하나의 연결 교체
func (c *Client) getCurrentConnection() *sql.DB  // round-robin으로 연결 선택
```

**simple.go** - 간단한 API
```go
func (c *Client) SelectAll(table, condition string, args ...interface{}) ([]map[string]interface{}, error)
func (c *Client) SelectOne(table, condition string, args ...interface{}) (map[string]interface{}, error)
func (c *Client) Insert(table string, data map[string]interface{}) (sql.Result, error)
func (c *Client) Update(table string, data map[string]interface{}, condition string, args ...interface{}) (sql.Result, error)
func (c *Client) Delete(table, condition string, args ...interface{}) (sql.Result, error)
func (c *Client) DeleteAll(table string) (sql.Result, error)
func (c *Client) Count(table, condition string, args ...interface{}) (int64, error)
func (c *Client) Exists(table, condition string, args ...interface{}) (bool, error)
```

**builder.go** - 쿼리 빌더
```go
type QueryBuilder struct { ... }
func (c *Client) Select(cols ...string) *QueryBuilder
func (qb *QueryBuilder) From(table string) *QueryBuilder
func (qb *QueryBuilder) Where(condition string, args ...interface{}) *QueryBuilder
func (qb *QueryBuilder) Join(table, condition string) *QueryBuilder
func (qb *QueryBuilder) OrderBy(order string) *QueryBuilder
func (qb *QueryBuilder) GroupBy(cols ...string) *QueryBuilder
func (qb *QueryBuilder) Having(condition string, args ...interface{}) *QueryBuilder
func (qb *QueryBuilder) Limit(n int) *QueryBuilder
func (qb *QueryBuilder) Offset(n int) *QueryBuilder
func (qb *QueryBuilder) All() ([]map[string]interface{}, error)
func (qb *QueryBuilder) One() (map[string]interface{}, error)
```

**transaction.go** - 트랜잭션
```go
type Tx struct { ... }
func (c *Client) Transaction(fn func(*Tx) error) error
// Tx에서 모든 simple.go 메서드 사용 가능
```

**retry.go** - 재시도 로직
```go
func (c *Client) executeWithRetry(fn func() error) error
func isRetryableError(err error) bool
```

**scan.go** - 결과 스캔
```go
func (c *Client) scanRows(rows *sql.Rows) ([]map[string]interface{}, error)
func convertType(val interface{}) interface{} // []byte → string, etc.
```

**config.go** - 설정 구조체
```go
type config struct { ... }
func defaultConfig() *config
```

**options.go** - 함수형 옵션
```go
type Option func(*config) error
func WithDSN(dsn string) Option
func WithCredentialRefresh(fn CredentialRefreshFunc, poolCount int, interval time.Duration) Option
func WithLogger(logger *logging.Logger) Option
// ... 기타 옵션들
```

**types.go** - 공통 타입
```go
type CredentialRefreshFunc func() (dsn string, error)
type Tx struct { ... }  // Transaction wrapper
```

**errors.go** - 에러 타입
```go
var ErrConnectionFailed = errors.New("...")
var ErrQueryFailed = errors.New("...")
// ... 기타 에러들
```

---

## Detailed Features / 상세 기능

### 1. Auto Connection Management / 자동 연결 관리

```go
func (c *Client) ensureConnected() error {
    // Read lock으로 빠른 확인
    c.mu.RLock()
    if c.healthy {
        c.mu.RUnlock()
        return nil // 연결 정상, 바로 리턴
    }
    c.mu.RUnlock()

    // Write lock으로 재연결 시도
    c.mu.Lock()
    defer c.mu.Unlock()

    // Ping으로 실제 연결 확인
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := c.db.PingContext(ctx); err != nil {
        // Ping 실패 → 재연결 시도
        if c.logger != nil {
            c.logger.Warn("Connection unhealthy, attempting reconnect", "error", err)
        }
        return c.reconnect()
    }

    c.healthy = true
    return nil
}

func (c *Client) reconnect() error {
    // 기존 연결 종료
    if c.db != nil {
        c.db.Close()
    }

    // 새 연결 생성
    db, err := sql.Open("mysql", c.config.dsn)
    if err != nil {
        return fmt.Errorf("reconnect failed: %w", err)
    }

    // 연결 풀 설정
    db.SetMaxOpenConns(c.config.maxOpenConns)
    db.SetMaxIdleConns(c.config.maxIdleConns)
    db.SetConnMaxLifetime(c.config.connMaxLifetime)
    db.SetConnMaxIdleTime(c.config.connMaxIdleTime)

    // Ping으로 확인
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := db.PingContext(ctx); err != nil {
        db.Close()
        return fmt.Errorf("reconnect ping failed: %w", err)
    }

    c.db = db
    c.healthy = true

    if c.logger != nil {
        c.logger.Info("Successfully reconnected to database")
    }

    return nil
}
```

### 2. Auto Retry Logic / 자동 재시도 로직

```go
func (c *Client) executeWithRetry(fn func() error) error {
    var lastErr error
    maxRetries := c.config.maxRetries

    for attempt := 0; attempt <= maxRetries; attempt++ {
        // 재시도 전 대기 (지수 백오프)
        if attempt > 0 {
            delay := c.config.retryDelay * time.Duration(1<<uint(attempt-1))
            if c.logger != nil {
                c.logger.Debug("Retrying query",
                    "attempt", attempt,
                    "delay", delay,
                )
            }
            time.Sleep(delay)
        }

        // 실행
        err := fn()
        if err == nil {
            return nil // 성공
        }

        // 재시도 가능한 에러인지 확인
        if !isRetryableError(err) {
            return err // 재시도 불가능, 바로 리턴
        }

        lastErr = err
    }

    return fmt.Errorf("failed after %d retries: %w", maxRetries, lastErr)
}

func isRetryableError(err error) bool {
    if err == nil {
        return false
    }

    errMsg := err.Error()

    // MySQL 에러 코드/메시지로 판단
    retryableErrors := []string{
        "driver: bad connection",
        "invalid connection",
        "connection refused",
        "MySQL server has gone away",
        "Error 1040",  // Too many connections
        "Error 1205",  // Lock wait timeout
        "Error 1213",  // Deadlock
        "Error 2006",  // MySQL server has gone away
        "Error 2013",  // Lost connection
    }

    for _, retryable := range retryableErrors {
        if strings.Contains(errMsg, retryable) {
            return true
        }
    }

    return false
}
```

### 3. Auto Resource Cleanup / 자동 리소스 정리

```go
func (c *Client) scanRows(rows *sql.Rows) ([]map[string]interface{}, error) {
    defer rows.Close() // ← 여기서 자동으로 처리!

    // 컬럼 이름 가져오기
    columns, err := rows.Columns()
    if err != nil {
        return nil, err
    }

    // 컬럼 타입 가져오기 (타입 변환용)
    columnTypes, err := rows.ColumnTypes()
    if err != nil {
        return nil, err
    }

    results := []map[string]interface{}{}

    for rows.Next() {
        // 각 행을 위한 values 준비
        values := make([]interface{}, len(columns))
        valuePtrs := make([]interface{}, len(columns))
        for i := range columns {
            valuePtrs[i] = &values[i]
        }

        // Scan
        if err := rows.Scan(valuePtrs...); err != nil {
            return nil, err
        }

        // map으로 변환
        row := make(map[string]interface{})
        for i, col := range columns {
            // 타입 변환 ([]byte → string 등)
            row[col] = convertType(values[i], columnTypes[i])
        }

        results = append(results, row)
    }

    // rows.Err() 확인
    if err := rows.Err(); err != nil {
        return nil, err
    }

    return results, nil
}

func convertType(val interface{}, colType *sql.ColumnType) interface{} {
    if val == nil {
        return nil
    }

    // []byte → string 변환
    if b, ok := val.([]byte); ok {
        return string(b)
    }

    // 필요시 다른 타입 변환 추가

    return val
}
```

### 4. Dynamic Credentials & Session Rotation / 동적 자격 증명 및 세션 순환

#### 왜 필요한가? / Why Needed?

**엔터프라이즈 환경에서**:
1. **보안 정책**: Vault 등에서 자격 증명이 주기적으로 변경됨 (예: 2시간마다)
2. **세션 타임아웃**: DB가 일정 시간 후 세션을 자동 종료
3. **Zero Downtime**: 서비스 중단 없이 자격 증명 갱신 필요

#### 해결 방법 / Solution

**여러 세션을 동시에 유지하고 주기적으로 하나씩 교체**:
- 예: **3개 세션 유지**, **1시간마다 하나씩 교체**
- Credential이 **2시간마다 변경**되더라도 항상 유효한 세션 유지!

```
시간 0:00 - [세션1, 세션2, 세션3] (모두 Credential A)
시간 1:00 - [세션1, 세션2, 세션3-NEW] (세션3 교체 → Credential B)
시간 2:00 - [세션1, 세션2-NEW, 세션3-NEW] (세션2 교체 → Credential B)
            ↑ Credential A가 만료되어도 세션2, 세션3은 Credential B로 정상 작동
시간 3:00 - [세션1-NEW, 세션2-NEW, 세션3-NEW] (세션1 교체 → Credential C)
```

**Zero Downtime 보장**: 항상 최소 1개 이상의 유효한 세션 유지!

#### How It Works / 작동 방식

**1. Multiple Connection Pools / 여러 연결 풀 유지**
```
┌─────────────────────────────────────────────────────┐
│  Client                                              │
│                                                      │
│  connections: []*sql.DB                              │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐          │
│  │ Pool 0   │  │ Pool 1   │  │ Pool 2   │          │
│  │ (Active) │  │ (Active) │  │ (Active) │          │
│  └──────────┘  └──────────┘  └──────────┘          │
│                                                      │
│  currentIdx: 0 (Round-robin 방식으로 사용)           │
└─────────────────────────────────────────────────────┘
```

**2. Rolling Rotation / 순환 교체**
```
시간 0분: [Pool0, Pool1, Pool2]  (모두 Credential A)
         ↓ 1시간 후

시간 60분: [Pool0, Pool1, Pool2-NEW]  (Pool2 교체 → Credential B)
          ↓ 1시간 후

시간 120분: [Pool0, Pool1-NEW, Pool2-NEW]  (Pool1 교체 → Credential B)
           ↑ Credential A 만료! 하지만 Pool1, Pool2가 정상 작동
           ↓ 1시간 후

시간 180분: [Pool0-NEW, Pool1-NEW, Pool2-NEW]  (Pool0 교체 → Credential C)
```

**Zero Downtime**: 항상 최소 2개의 정상 연결 유지!

#### Credential Refresh Function / 자격 증명 갱신 함수

**간단한 함수 타입으로 제공** - 사용자가 Vault든 다른 방법이든 자유롭게 구현:

```go
// CredentialRefreshFunc returns a new DSN (user can fetch from Vault, etc.)
// CredentialRefreshFunc는 새 DSN을 반환합니다 (사용자가 Vault 등에서 가져올 수 있음)
type CredentialRefreshFunc func() (dsn string, error)
```

#### Implementation 1: Static Credentials / 정적 자격 증명 (기본값)

```go
// 고정 DSN 사용 - credential rotation 없음
db, err := mysql.New(
    mysql.WithDSN("user:password@tcp(localhost:3306)/dbname"),
)
// 내부적으로 단일 연결 풀만 사용
```

#### Implementation 2: Dynamic Credentials (Vault, etc.) / 동적 자격 증명 (Vault 등)

**사용자가 credential을 가져오는 함수를 직접 구현**:

```go
// 사용자가 Vault에서 credential 가져오기 (사용자 구현)
func getCredentialsFromVault() (string, error) {
    // 사용자가 직접 Vault 클라이언트 사용
    // 예시 (의사 코드):
    // secret, err := vaultClient.Logical().Read("database/creds/my-role")
    // username := secret.Data["username"]
    // password := secret.Data["password"]

    username := "dynamic_user_123"
    password := "dynamic_pass_456"

    dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/mydb", username, password)
    return dsn, nil
}

// MySQL 클라이언트 생성 - credential refresh 설정
db, err := mysql.New(
    mysql.WithCredentialRefresh(
        getCredentialsFromVault,  // 사용자 구현 함수
        3,                         // 3개 세션 유지
        1*time.Hour,               // 1시간마다 하나씩 교체
    ),
    mysql.WithLogger(logger),
)

// 결과:
// - 처음에 3개 세션 생성 (모두 같은 credential)
// - 1시간 후: 세션 1개를 새 credential로 교체
// - 2시간 후: 또 다른 세션 1개 교체
// - 3시간 후: 마지막 세션 교체
// → Credential이 2시간마다 바뀌어도 항상 유효한 세션 유지!
```

**더 간단한 예시 - 환경 변수나 파일에서 읽기**:

```go
func getCredentialsFromFile() (string, error) {
    // 파일이나 환경 변수에서 읽기
    username := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASS")

    dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/mydb", username, password)
    return dsn, nil
}

db, err := mysql.New(
    mysql.WithCredentialRefresh(
        getCredentialsFromFile,
        3,              // 3개 세션
        30*time.Minute, // 30분마다 교체
    ),
)
```

#### Connection Rotation Logic / 연결 순환 로직

```go
func (c *Client) startConnectionRotation() {
    // credRefreshFunc이 없으면 rotation 불필요 (정적 credential)
    if c.config.credRefreshFunc == nil {
        return
    }

    go func() {
        ticker := time.NewTicker(c.config.rotationInterval)
        defer ticker.Stop()

        for {
            select {
            case <-ticker.C:
                if err := c.rotateOneConnection(); err != nil {
                    if c.logger != nil {
                        c.logger.Error("Connection rotation failed", "error", err)
                    }
                }
            case <-c.rotationStop:
                return
            }
        }
    }()
}

func (c *Client) rotateOneConnection() error {
    // 1. 사용자 함수로 새 DSN 가져오기
    newDSN, err := c.config.credRefreshFunc()
    if err != nil {
        return fmt.Errorf("credential refresh function failed: %w", err)
    }

    // 2. 새 연결 풀 생성
    newDB, err := sql.Open("mysql", newDSN)
    if err != nil {
        return fmt.Errorf("failed to create new connection: %w", err)
    }

    // 3. 연결 풀 설정
    newDB.SetMaxOpenConns(c.config.maxOpenConns)
    newDB.SetMaxIdleConns(c.config.maxIdleConns)
    newDB.SetConnMaxLifetime(c.config.connMaxLifetime)
    newDB.SetConnMaxIdleTime(c.config.connMaxIdleTime)

    // 4. Ping으로 연결 확인
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := newDB.PingContext(ctx); err != nil {
        newDB.Close()
        return fmt.Errorf("new connection ping failed: %w", err)
    }

    // 5. 가장 오래된 연결 찾기 (순환 교체)
    c.connectionsMu.Lock()
    oldestIdx := c.rotationIdx % len(c.connections)
    c.rotationIdx++
    oldDB := c.connections[oldestIdx]

    // 6. 교체
    c.connections[oldestIdx] = newDB
    c.connectionsMu.Unlock()

    // 7. 오래된 연결 종료 (graceful - 30초 후)
    go func() {
        time.Sleep(30 * time.Second) // 활성 쿼리 완료 대기
        oldDB.Close()
    }()

    if c.logger != nil {
        c.logger.Info("Connection rotated successfully",
            "pool_index", oldestIdx,
            "total_pools", len(c.connections),
        )
    }

    return nil
}

func (c *Client) getCurrentConnection() *sql.DB {
    c.connectionsMu.RLock()
    defer c.connectionsMu.RUnlock()

    // Round-robin으로 연결 선택
    conn := c.connections[c.currentIdx]
    c.currentIdx = (c.currentIdx + 1) % len(c.connections)

    return conn
}
```

#### Configuration / 설정

**1. 기본 (정적 자격 증명)**:
```go
// 단일 연결 풀 - rotation 없음
db, err := mysql.New(
    mysql.WithDSN("user:password@tcp(localhost:3306)/dbname"),
)
```

**2. 동적 자격 증명 (Vault, 파일, 환경변수 등)**:
```go
// 사용자가 credential 가져오는 함수 구현
func getDSN() (string, error) {
    // Vault, 파일, 환경변수 등에서 가져오기
    return "user:password@tcp(localhost:3306)/dbname", nil
}

// 3개 세션, 1시간마다 하나씩 교체
db, err := mysql.New(
    mysql.WithCredentialRefresh(
        getDSN,        // 사용자 함수
        3,             // 세션 수
        1*time.Hour,   // 교체 주기
    ),
    mysql.WithLogger(logger),
)
```

#### Benefits / 이점

1. **Zero Downtime / 무중단**:
   - 항상 최소 N-1개의 정상 연결 유지
   - Credential이 바뀌어도 서비스 중단 없음

2. **Security / 보안**:
   - 주기적인 자격 증명 순환
   - Vault, AWS Secrets Manager 등 모든 시스템과 호환

3. **Simple / 간단함**:
   - 사용자는 DSN 가져오는 함수만 구현
   - 나머지는 패키지가 자동 처리

4. **Flexible / 유연함**:
   - Vault, 파일, 환경변수, API 등 어떤 방법이든 지원
   - 사용자가 원하는 방식으로 구현 가능

---

## Summary / 요약

### Core Principles / 핵심 원칙

1. **극도의 간결함** - 기존 방법보다 10배 간단
2. **자동 관리** - 연결, 재시도, 리소스 정리 모두 자동
3. **SQL 친화적** - SQL 문법에 가깝게
4. **제로 정신적 부담** - DB 상태를 신경 쓰지 않음

### What Makes This Different / 차별점

| 기존 database/sql | 이 패키지 |
|---|---|
| 연결 상태 수동 관리 | ✅ 자동 관리 |
| 재접속 로직 직접 작성 | ✅ 자동 재접속 |
| defer rows.Close() 필수 | ✅ 자동 처리 |
| SQL 문법 직접 작성 | ✅ 간단한 메서드 |
| 에러 재시도 직접 구현 | ✅ 자동 재시도 |
| 30줄 코드 | ✅ 2줄 코드 |

### Success Criteria / 성공 기준

이 패키지가 성공하려면:

1. ✅ 기존 방법보다 **최소 5배** 간단해야 함
2. ✅ 개발자가 DB 연결 상태를 **전혀 신경 쓰지 않아야** 함
3. ✅ SQL 문법을 **거의 그대로** 사용할 수 있어야 함
4. ✅ **자동으로** 모든 번거로운 일이 처리되어야 함

**이 정도로 간단하지 않으면 만들지 않는 것이 맞습니다.**

---

**Document Version / 문서 버전**: 3.0.0 (Final - Extreme Simplicity)
**Last Updated / 최종 업데이트**: 2025-10-10
**Status / 상태**: Ready for Implementation / 구현 준비 완료
