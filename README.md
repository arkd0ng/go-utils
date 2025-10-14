# go-utils

A collection of frequently used utility functions for Golang development.

Golang 개발에 자주 사용되는 유틸리티 함수 모음입니다.

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.16-blue)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Overview / 개요

This library provides a growing collection of utility packages for common programming tasks in Go. Each package is designed to be independent, well-documented, and easy to use.

이 라이브러리는 Go의 일반적인 프로그래밍 작업을 위한 유틸리티 패키지 모음을 제공합니다. 각 패키지는 독립적이고 문서화가 잘 되어 있으며 사용하기 쉽게 설계되었습니다.

## Installation / 설치

```bash
go get github.com/arkd0ng/go-utils
```

Or import specific packages:

또는 특정 패키지만 import:

```bash
go get github.com/arkd0ng/go-utils/random
```

## Package Structure / 패키지 구조

This library is organized into subpackages for better modularity:

이 라이브러리는 모듈화를 위해 서브패키지로 구성되어 있습니다:

```
go-utils/
├── random/          # Random generation utilities / 랜덤 생성 유틸리티
├── logging/         # Logging with file rotation / 파일 로테이션 로깅
├── database/
│   ├── mysql/       # Extreme simplicity MySQL client / 극도로 간단한 MySQL 클라이언트
│   └── redis/       # Extreme simplicity Redis client / 극도로 간단한 Redis 클라이언트
├── stringutil/      # String manipulation utilities / 문자열 처리 유틸리티
├── sliceutil/       # Slice helpers (coming soon) / 슬라이스 헬퍼 (예정)
├── maputil/         # Map utilities (coming soon) / 맵 유틸리티 (예정)
└── ...
```

## Available Packages / 사용 가능한 패키지

### ✅ [random](./random/) - Random String Generation

Generate cryptographically secure random strings with various character sets.

다양한 문자 집합으로 암호학적으로 안전한 랜덤 문자열을 생성합니다.

**14 methods available** including: Letters, Alnum, Digits, Hex, AlphaUpper, AlphaLower, Base64URL, and more.

**14개 메서드 제공**: Letters, Alnum, Digits, Hex, AlphaUpper, AlphaLower, Base64URL 등.

**Flexible API**: Support both fixed length and range with variadic parameters and error handling.

**유연한 API**: 가변 인자와 에러 핸들링으로 고정 길이 및 범위 모두 지원.

```go
import (
    "log"
    "github.com/arkd0ng/go-utils/random"
)

// Fixed length: 32 characters / 고정 길이: 32자
str, err := random.GenString.Alnum(32)
if err != nil {
    log.Fatal(err)
}

// Range length: 32-128 characters / 범위 길이: 32-128자
str2, err := random.GenString.Alnum(32, 128)
if err != nil {
    log.Fatal(err)
}

// Generate PIN code (fixed 6 digits) / PIN 코드 생성 (고정 6자리)
pin, err := random.GenString.Digits(6)
if err != nil {
    log.Fatal(err)
}

// Generate hex color code / 16진수 색상 코드 생성
color, err := random.GenString.Hex(6)
if err != nil {
    log.Fatal(err)
}
```

**[→ View full documentation / 전체 문서 보기](./random/README.md)**

---

### ✅ [logging](./logging/) - Structured Logging with File Rotation

Simple and powerful logging with automatic file rotation (lumberjack), structured logging, and banner support.

자동 파일 로테이션(lumberjack), 구조화된 로깅, 배너 지원이 있는 간단하고 강력한 로깅 패키지입니다.

**Features**: Multiple log levels, key-value logging, colored output, thread-safe / 다중 로그 레벨, 키-값 로깅, 색상 출력, 스레드 안전

```go
import "github.com/arkd0ng/go-utils/logging"

// Default logger / 기본 로거
logger := logging.Default()
defer logger.Close()

logger.Banner("My Application", "v1.0.0")
logger.Info("Application started", "port", 8080)

// Multiple loggers for different purposes / 용도별 여러 로거
appLogger, _ := logging.New(logging.WithFilePath("./logs/app.log"))
dbLogger, _ := logging.New(logging.WithFilePath("./logs/db.log"))
```

**[→ View full documentation / 전체 문서 보기](./logging/README.md)**

---

### ✅ [database/mysql](./database/mysql/) - Extreme Simplicity MySQL/MariaDB Client

Reduce 30+ lines of boilerplate code to just 2 lines with auto-everything: connection management, retry, reconnection, and resource cleanup.

30줄 이상의 보일러플레이트 코드를 단 2줄로 줄이고, 연결 관리, 재시도, 재연결, 리소스 정리를 모두 자동화합니다.

**Core Features**: Zero-downtime credential rotation, SQL-like API, auto retry, transaction support, no defer rows.Close() / 무중단 자격 증명 순환, SQL 문법에 가까운 API, 자동 재시도, 트랜잭션 지원, defer rows.Close() 불필요

**Advanced Features** (v1.3.010+):
- **Batch Operations**: BatchInsert, BatchUpdate, BatchDelete, BatchSelectByIDs / 배치 작업
- **Upsert**: Upsert, UpsertBatch, Replace (ON DUPLICATE KEY UPDATE) / Upsert 작업
- **Pagination**: Paginate, PaginateQuery with metadata / 메타데이터를 포함한 페이지네이션
- **Soft Delete**: SoftDelete, Restore, SelectAllWithTrashed / 소프트 삭제
- **Query Statistics**: Performance monitoring, slow query logging / 쿼리 통계
- **Pool Metrics**: Connection pool health monitoring / 연결 풀 모니터링
- **Schema Inspector**: GetTables, GetColumns, GetIndexes, InspectTable / 스키마 검사
- **Migration Helpers**: CreateTable, AddColumn, AddIndex, AddForeignKey / 마이그레이션 헬퍼
- **CSV Export/Import**: ExportTableToCSV, ImportFromCSV / CSV 내보내기/가져오기

```go
import (
    "context"
    "github.com/arkd0ng/go-utils/database/mysql"
)

// Create client / 클라이언트 생성
db, _ := mysql.New(mysql.WithDSN("user:pass@tcp(localhost:3306)/dbname"))
defer db.Close()

// Simple API - 30 lines → 2 lines! / 간단한 API - 30줄 → 2줄!
users, _ := db.SelectAll(ctx, "users", "age > ?", 18)

// Insert with map / 맵으로 삽입
db.Insert(ctx, "users", map[string]interface{}{
    "name": "John", "email": "john@example.com", "age": 30,
})

// Transaction with auto commit/rollback / 자동 커밋/롤백 트랜잭션
db.Transaction(ctx, func(tx *mysql.Tx) error {
    tx.Insert(ctx, "users", map[string]interface{}{"name": "Jane"})
    tx.Insert(ctx, "profiles", map[string]interface{}{"user_id": 1})
    return nil // Auto commit / 자동 커밋
})

// Dynamic credentials (Vault, AWS Secrets Manager, etc.) / 동적 자격 증명
db, _ := mysql.New(
    mysql.WithCredentialRefresh(
        func() (string, error) {
            // User fetches credentials from Vault, file, etc.
            // Vault, 파일 등에서 자격 증명 가져오기
            return "user:pass@tcp(localhost:3306)/db", nil
        },
        3,            // 3 connection pools / 3개 연결 풀
        1*time.Hour,  // Rotate one per hour / 1시간마다 하나씩 교체
    ),
)
```

**Before vs After**:
```go
// ❌ Before: 30+ lines with standard database/sql
db, _ := sql.Open("mysql", dsn)
defer db.Close()
rows, _ := db.Query("SELECT * FROM users WHERE age > ?", 18)
defer rows.Close() // Must remember! / 기억해야 함!
var users []User
for rows.Next() {
    var u User
    rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age) // Manual scanning / 수동 스캔
    users = append(users, u)
}
// ... 20+ more lines

// ✅ After: 2 lines with this package
db, _ := mysql.New(mysql.WithDSN(dsn))
users, _ := db.SelectAll(ctx, "users", "age > ?", 18)
```

**[→ View full documentation / 전체 문서 보기](./database/mysql/README.md)**

---

### ✅ [database/redis](./database/redis/) - Extreme Simplicity Redis Client

Reduce 20+ lines of boilerplate code to just 2 lines with auto-everything: connection management, retry, reconnect, and resource cleanup.

20줄 이상의 보일러플레이트 코드를 단 2줄로 줄이고, 연결 관리, 재시도, 재연결, 리소스 정리를 모두 자동화합니다.

**Core Features**: Auto-retry with exponential backoff, connection pooling, health check, pipeline support, transaction support, Pub/Sub / 지수 백오프를 사용한 자동 재시도, 연결 풀링, 헬스 체크, 파이프라인 지원, 트랜잭션 지원, Pub/Sub

**Operations Supported / 지원되는 작업**:
- **String**: Set, Get, MGet, MSet, Incr, Decr, SetNX, SetEX / 문자열 작업
- **Hash**: HSet, HGet, HGetAll, HSetMap, HDel, HIncrBy / 해시 작업
- **List**: LPush, RPush, LPop, RPop, LRange, LLen / 리스트 작업
- **Set**: SAdd, SRem, SMembers, SUnion, SInter, SDiff / 집합 작업
- **Sorted Set**: ZAdd, ZRange, ZRangeByScore, ZScore / 정렬 집합 작업
- **Key**: Del, Exists, Expire, TTL, Keys, Scan / 키 작업
- **Advanced**: Pipeline, Transaction, Pub/Sub / 고급 기능

```go
import (
    "context"
    "github.com/arkd0ng/go-utils/database/redis"
)

// Create client / 클라이언트 생성
rdb, _ := redis.New(redis.WithAddr("localhost:6379"))
defer rdb.Close()

ctx := context.Background()

// String operations / 문자열 작업
rdb.Set(ctx, "key", "value")
val, _ := rdb.Get(ctx, "key")

// Hash operations / 해시 작업
rdb.HSetMap(ctx, "user:123", map[string]interface{}{
    "name":  "John",
    "email": "john@example.com",
    "age":   30,
})
fields, _ := rdb.HGetAll(ctx, "user:123")

// List operations (queue) / 리스트 작업 (큐)
rdb.RPush(ctx, "queue", "task1", "task2")
item, _ := rdb.LPop(ctx, "queue")

// Pipeline for batch operations / 배치 작업을 위한 파이프라인
rdb.Pipeline(ctx, func(pipe redis.Pipeliner) error {
    pipe.Set(ctx, "key1", "value1", 0)
    pipe.Set(ctx, "key2", "value2", 0)
    pipe.Incr(ctx, "counter")
    return nil
})
```

**Before vs After**:
```go
// ❌ Before: 20+ lines with standard go-redis
import "github.com/redis/go-redis/v9"

rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
err := rdb.Set(ctx, "key", "value", 0).Err()
if err != nil {
    return err
}
val, err := rdb.Get(ctx, "key").Result()
if err != nil {
    return err
}
// ... 매번 .Err() 또는 .Result() 호출...

// ✅ After: 2 lines with this package
rdb, _ := redis.New(redis.WithAddr("localhost:6379"))
rdb.Set(ctx, "key", "value")
val, _ := rdb.Get(ctx, "key")
```

**[→ View full documentation / 전체 문서 보기](./database/redis/README.md)**

---

### ✅ [stringutil](./stringutil/) - String Manipulation Utilities

Extreme simplicity string utilities - reduce 20 lines of string manipulation code to just 1 line.

극도로 간단한 문자열 유틸리티 - 20줄의 문자열 처리 코드를 단 1줄로 줄입니다.

**Core Features**: Unicode-safe operations, 53 functions across 9 categories / 유니코드 안전 작업, 9개 카테고리에 걸쳐 53개 함수

**Categories / 카테고리**:
- **Case Conversion (9)**: ToSnakeCase, ToCamelCase, ToTitle, Slugify, Quote, Unquote / 케이스 변환
- **String Manipulation (17)**: Truncate, Reverse, Substring, Insert, SwapCase, Repeat / 문자열 조작
- **Validation (8)**: IsEmail, IsURL, IsAlphanumeric, IsNumeric, IsBlank / 유효성 검사
- **Comparison (3)**: EqualFold, HasPrefix, HasSuffix / 비교
- **Search & Replace (6)**: ContainsAny, ContainsAll, ReplaceAll, ReplaceIgnoreCase / 검색 및 치환
- **Unicode Operations (3)**: RuneCount, Width, Normalize / 유니코드 작업
- **Collection Utilities (5)**: CountWords, Map, Filter, Join / 컬렉션 유틸리티
- **String Generation (2)**: PadLeft, PadRight / 문자열 생성
- **String Parsing (2)**: Lines, Words / 문자열 파싱

```go
import "github.com/arkd0ng/go-utils/stringutil"

// Case conversion / 케이스 변환
snake := stringutil.ToSnakeCase("HelloWorld")  // "hello_world"
camel := stringutil.ToCamelCase("hello_world") // "helloWorld"
title := stringutil.ToTitle("hello world")     // "Hello World"
slug := stringutil.Slugify("Hello World!")     // "hello-world"

// String manipulation / 문자열 조작
short := stringutil.Truncate("Long text here", 10)     // "Long text..."
sub := stringutil.Substring("hello world", 0, 5)       // "hello"
inserted := stringutil.Insert("hello world", 5, ",")   // "hello, world"
swapped := stringutil.SwapCase("Hello World")          // "hELLO wORLD"

// Validation & Comparison / 유효성 검사 및 비교
if stringutil.IsEmail("user@example.com") {
    // Valid email / 유효한 이메일
}
if stringutil.EqualFold("hello", "HELLO") {
    // Case-insensitive match / 대소문자 구분 없이 일치
}

// Unicode operations / 유니코드 작업
count := stringutil.RuneCount("안녕하세요")     // 5 (not 15 bytes)
width := stringutil.Width("hello世界")         // 9 (5 + 4)
normalized := stringutil.Normalize("café", "NFC") // "café"

// Functional programming (Map/Filter) / 함수형 프로그래밍
names := []string{"alice", "bob", "charlie"}
upper := stringutil.Map(names, func(s string) string {
    return strings.ToUpper(s)
}) // ["ALICE", "BOB", "CHARLIE"]

filtered := stringutil.Filter(names, func(s string) bool {
    return len(s) > 3
}) // ["alice", "charlie"]
```

**[→ View full documentation / 전체 문서 보기](./stringutil/README.md)**

---

### 🔜 Coming Soon / 개발 예정

- **sliceutil** - Slice/Array helpers / 슬라이스/배열 헬퍼
- **maputil** - Map utilities / 맵 유틸리티
- **fileutil** - File/Path utilities / 파일/경로 유틸리티
- **httputil** - HTTP helpers / HTTP 헬퍼
- **timeutil** - Time/Date utilities / 시간/날짜 유틸리티
- **validation** - Validation utilities / 검증 유틸리티
- **errorutil** - Error handling helpers / 에러 처리 헬퍼

## Quick Start / 빠른 시작

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/random"
)

func main() {
    // Generate a secure password / 안전한 비밀번호 생성
    password, err := random.GenString.Complex(16, 24)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Password:", password)

    // Generate an API key (fixed length) / API 키 생성 (고정 길이)
    apiKey, err := random.GenString.Alnum(40)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("API Key:", apiKey)
}
```

## Testing / 테스트

Run all tests:

모든 테스트 실행:

```bash
go test ./... -v
```

Run benchmarks:

벤치마크 실행:

```bash
go test ./... -bench=.
```

## Contributing / 기여하기

Contributions are welcome! This library will grow with frequently used utility functions.

기여를 환영합니다! 이 라이브러리는 자주 사용되는 유틸리티 함수들로 성장할 것입니다.

1. Fork the repository / 저장소 포크
2. Create your feature branch / 기능 브랜치 생성
3. Commit your changes / 변경사항 커밋
4. Push to the branch / 브랜치에 푸시
5. Create a Pull Request / Pull Request 생성

## License / 라이선스

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

이 프로젝트는 MIT 라이선스에 따라 배포됩니다 - 자세한 내용은 [LICENSE](LICENSE) 파일을 참조하세요.

## Author / 작성자

**arkd0ng**

- GitHub: [@arkd0ng](https://github.com/arkd0ng)

## Changelog / 변경 이력

For detailed version history, see:
- [CHANGELOG.md](./CHANGELOG.md) - Major/Minor version overview
- [docs/CHANGELOG/](./docs/CHANGELOG/) - Detailed patch-level changes

상세한 버전 히스토리는 다음을 참조하세요:
- [CHANGELOG.md](./CHANGELOG.md) - Major/Minor 버전 개요
- [docs/CHANGELOG/](./docs/CHANGELOG/) - 상세한 패치별 변경사항

### v1.5.x (Current / 현재)

- **NEW**: `stringutil` package - String manipulation utilities / 문자열 처리 유틸리티
  - 20 lines → 1 line code reduction / 20줄 → 1줄 코드 감소
  - 37 functions across 5 categories / 5개 카테고리에 걸쳐 37개 함수
  - Unicode-safe operations (rune-based) / 유니코드 안전 작업 (rune 기반)
  - Zero external dependencies / 외부 의존성 제로
  - Functional programming (Map/Filter) / 함수형 프로그래밍
  - Comprehensive documentation (USER_MANUAL, DEVELOPER_GUIDE) / 포괄적인 문서화

### v1.4.x

- **NEW**: `database/redis` package - Extreme simplicity Redis client / 극도로 간단한 Redis 클라이언트
  - 20 lines → 2 lines code reduction / 20줄 → 2줄 코드 감소
  - Auto-retry with exponential backoff / 지수 백오프를 사용한 자동 재시도
  - Connection pooling and health check / 연결 풀링 및 헬스 체크
  - 60+ methods: String, Hash, List, Set, Sorted Set, Key operations
  - Pipeline, Transaction, Pub/Sub support / 파이프라인, 트랜잭션, Pub/Sub 지원
  - Type-safe generic methods / 타입 안전 제네릭 메서드
- **DOCKER**: Docker Redis setup with automated scripts / 자동화된 스크립트를 사용한 Docker Redis 설정

### v1.3.x

- **NEW**: `database/mysql` package - Extreme simplicity MySQL/MariaDB client / 극도로 간단한 MySQL/MariaDB 클라이언트
  - 30 lines → 2 lines code reduction / 30줄 → 2줄 코드 감소
  - Zero-downtime credential rotation / 무중단 자격 증명 순환
  - Auto everything: connection, retry, cleanup / 모든 것 자동화
  - 7 Simple API methods + Advanced features / 7개 Simple API 메서드 + 고급 기능
- **DOCS**: Comprehensive documentation for Random and Logging packages / Random 및 Logging 패키지 종합 문서화
  - User manuals and developer guides / 사용자 매뉴얼 및 개발자 가이드
  - Bilingual documentation (English/Korean) / 이중 언어 문서 (영문/한글)

### v1.2.x

- Documentation improvements / 문서 개선
- CHANGELOG system restructured / CHANGELOG 시스템 재구성

### v1.1.x

- **NEW**: `logging` package with file rotation / 파일 로테이션 로깅 패키지
- Structured logging with lumberjack / lumberjack을 사용한 구조화된 로깅

### v1.0.x

- **NEW**: `random` package with 14 methods / 14개 메서드를 가진 랜덤 패키지
- Cryptographically secure random generation / 암호학적으로 안전한 랜덤 생성

### v0.2.0

- **BREAKING CHANGE**: Refactored to subpackage structure / 서브패키지 구조로 리팩토링
  - Moved `GenRandomString` to `random.GenString` / `GenRandomString`을 `random.GenString`으로 이동

### v0.1.0 (Initial Release / 첫 릴리스)

- Added `random` package with string generation utilities / 문자열 생성 유틸리티가 포함된 `random` 패키지 추가
