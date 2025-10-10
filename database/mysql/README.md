# MySQL Package / MySQL 패키지

Extreme simplicity MySQL/MariaDB client with zero-downtime credential rotation.

무중단 자격 증명 순환을 갖춘 극도로 간단한 MySQL/MariaDB 클라이언트.

## Features / 특징

✅ **Extreme Simplicity**: 30 lines → 2 lines of code / 극도의 간결함: 30줄 → 2줄 코드
✅ **Auto Everything**: Connection management, retry, cleanup / 모든 것 자동: 연결 관리, 재시도, 정리
✅ **Zero-Downtime Credential Rotation**: Multiple connection pools with rolling rotation / 무중단 자격 증명 순환
✅ **SQL-Like API**: Close to actual SQL syntax / SQL 문법에 가까운 API
✅ **No defer rows.Close()**: Automatic resource cleanup / 자동 리소스 정리
✅ **Auto Retry**: Transient errors are retried automatically / 일시적 에러 자동 재시도
✅ **Health Check**: Automatic connection monitoring / 자동 연결 모니터링

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

## Best Practices / 모범 사례

1. **Always use context**: Pass context for timeout control / 타임아웃 제어를 위해 항상 context 전달
2. **Use transactions for multi-step operations**: Ensure atomicity / 다단계 작업에는 트랜잭션 사용
3. **Enable query logging in development**: Debug slow queries / 개발 중 쿼리 로깅 활성화
4. **Use credential rotation for production**: Enhanced security / 프로덕션에는 자격 증명 순환 사용
5. **Monitor health checks**: Track connection status / 헬스 체크 모니터링

## License / 라이선스

MIT

## Contributing / 기여

Contributions are welcome! Please follow the project's coding standards.

기여를 환영합니다! 프로젝트의 코딩 표준을 따라주세요.
