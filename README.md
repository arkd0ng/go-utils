# go-utils

A collection of frequently used utility functions for Golang development.

Golang 개발에 자주 사용되는 유틸리티 함수 모음입니다.

[![go-utils version](https://img.shields.io/badge/dynamic/yaml?url=https%3A%2F%2Fraw.githubusercontent.com%2Farkd0ng%2Fgo-utils%2Fmain%2Fcfg%2Fapp.yaml&query=$.app.version&label=go-utils&color=blue&cacheSeconds=300)](https://github.com/arkd0ng/go-utils)
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
├── random/          # Random generation utilities (14 methods) / 랜덤 생성 유틸리티 (14개 메서드)
├── logging/         # Logging with file rotation / 파일 로테이션 로깅
├── database/
│   ├── mysql/       # Extreme simplicity MySQL client / 극도로 간단한 MySQL 클라이언트
│   └── redis/       # Extreme simplicity Redis client / 극도로 간단한 Redis 클라이언트
├── stringutil/      # String manipulation utilities (53 functions) / 문자열 처리 유틸리티 (53개 함수)
├── timeutil/        # Time and date utilities (114+ functions) / 시간 및 날짜 유틸리티 (114+ 함수)
├── sliceutil/       # Slice utilities (95 functions) / 슬라이스 유틸리티 (95개 함수)
├── maputil/         # Map utilities (99 functions) / 맵 유틸리티 (99개 함수)
├── fileutil/        # File and path utilities (~91 functions) / 파일 및 경로 유틸리티 (약 91개 함수)
├── httputil/        # HTTP client utilities / HTTP 클라이언트 유틸리티
├── websvrutil/      # Web server utilities / 웹 서버 유틸리티
├── errorutil/       # Error handling utilities (21 functions) / 에러 처리 유틸리티 (21개 함수)
├── validation/      # Fluent validation library (135+ validators) / Fluent 검증 라이브러리 (135개 이상 검증기)
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

### ✅ [timeutil](./timeutil/) - Time and Date Utilities

Extreme simplicity time utilities - reduce 20 lines of time manipulation code to just 1 line with KST (GMT+9) as default timezone.

극도로 간단한 시간 유틸리티 - 20줄의 시간 처리 코드를 단 1줄로 줄이며, KST (GMT+9)를 기본 타임존으로 사용합니다.

**Core Features**: 80+ functions, KST default timezone, custom format tokens, business day support / 80개 이상 함수, KST 기본 타임존, 커스텀 포맷 토큰, 영업일 지원

**Categories / 카테고리**:
- **Time Difference (8)**: SubTime, DiffInSeconds, DiffInMinutes, DiffInHours, DiffInDays / 시간 차이
- **Timezone Operations (10)**: ConvertTimezone, ToKST, NowKST, SetDefaultTimezone / 타임존 작업
- **Date Arithmetic (16)**: AddDays, AddWeeks, StartOfDay, EndOfMonth, StartOfYear / 날짜 연산
- **Date Formatting (8)**: FormatISO8601, FormatKorean, Format (YYYY-MM-DD) / 날짜 포맷팅
- **Time Parsing (6)**: ParseDate, ParseDateTime, Parse (auto-detect format) / 시간 파싱
- **Time Comparisons (18)**: IsToday, IsWeekend, IsBetween, IsThisMonth / 시간 비교
- **Age Calculations (4)**: AgeInYears, Age (years/months/days) / 나이 계산
- **Relative Time (3)**: RelativeTime ("2 hours ago"), TimeAgo / 상대 시간
- **Unix Timestamp (12)**: Now, NowMilli, FromUnix, ToUnix / Unix 타임스탬프
- **Business Days (7)**: IsBusinessDay, AddBusinessDays, AddKoreanHolidays / 영업일

```go
import "github.com/arkd0ng/go-utils/timeutil"

// Time difference with human-readable output / 사람이 읽기 쉬운 시간 차이
start := time.Date(2025, 1, 1, 9, 0, 0, 0, time.UTC)
end := time.Date(2025, 1, 3, 15, 30, 0, 0, time.UTC)
diff := timeutil.SubTime(start, end)
fmt.Println(diff.String()) // "2 days 6 hours 30 minutes"

// Timezone operations with KST default / KST 기본 타임존 작업
kstNow := timeutil.NowKST()
nyTime, _ := timeutil.ConvertTimezone(time.Now(), "America/New_York")

// Custom format tokens (YYYY-MM-DD instead of 2006-01-02) / 커스텀 포맷 토큰
formatted := timeutil.Format(time.Now(), "YYYY-MM-DD HH:mm:ss") // "2025-10-14 15:04:05"
korean := timeutil.FormatKorean(time.Now()) // "2025년 10월 14일 15시 04분 05초"

// Business days with Korean holidays / 한국 공휴일을 포함한 영업일
timeutil.AddKoreanHolidays(2025)
nextBizDay := timeutil.AddBusinessDays(time.Now(), 5)
isHoliday := timeutil.IsHoliday(time.Date(2025, 1, 1, 0, 0, 0, 0, timeutil.KST))

// Relative time / 상대 시간
past := time.Now().Add(-2 * time.Hour)
fmt.Println(timeutil.RelativeTime(past)) // "2 hours ago"

// Age calculation / 나이 계산
birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
age := timeutil.Age(birthDate)
fmt.Println(age.String()) // "35 years 4 months 29 days"
```

**Before vs After**:
```go
// ❌ Before: 20+ lines with standard time package
start := time.Date(2025, 1, 1, 9, 0, 0, 0, time.UTC)
end := time.Date(2025, 1, 3, 15, 30, 0, 0, time.UTC)
duration := end.Sub(start)
hours := duration.Hours()
days := hours / 24
if days > 0 {
    fmt.Printf("%d days %d hours", int(days), int(hours)%24)
} else if hours > 0 {
    fmt.Printf("%d hours %d minutes", int(hours), int(duration.Minutes())%60)
}
// ... 더 많은 코드

// ✅ After: 1-2 lines with this package
diff := timeutil.SubTime(start, end)
fmt.Println(diff.String()) // "2 days 6 hours 30 minutes"
```

**[→ View full documentation / 전체 문서 보기](./timeutil/README.md)**

---

### ✅ [sliceutil](./sliceutil/) - Slice Utilities

Extreme simplicity slice utilities - reduce 20 lines of repetitive slice manipulation code to just 1 line with **95 type-safe functions**.

극도로 간단한 슬라이스 유틸리티 - 20줄의 반복적인 슬라이스 조작 코드를 단 1줄로 줄이며, **95개의 타입 안전 함수**를 제공합니다.

**Core Features**: 95 functions across 14 categories, Go 1.18+ generics, functional programming style, immutable operations, zero dependencies, 100% test coverage / 14개 카테고리에 걸쳐 95개 함수, Go 1.18+ 제네릭, 함수형 프로그래밍 스타일, 불변 작업, 제로 의존성, 100% 테스트 커버리지

**Categories / 카테고리**:
- **Basic Operations (11)**: Contains, IndexOf, Find, Count, FindLast, Equal / 기본 작업
- **Transformation (8)**: Map, Filter, Unique, Reverse, Flatten, FlatMap / 변환
- **Aggregation (11)**: Reduce, ReduceRight, Sum, Min, Max, MinBy, MaxBy, Average, GroupBy / 집계
- **Slicing (11)**: Chunk, Take, TakeLast, TakeWhile, Drop, DropWhile, Sample, Window, Interleave / 슬라이싱
- **Set Operations (6)**: Union, Intersection, Difference, SymmetricDifference, IsSubset, IsSuperset / 집합 작업
- **Sorting (6)**: Sort, SortBy, SortByMulti, SortDesc, IsSorted, IsSortedDesc / 정렬
- **Predicates (6)**: All, Any, None, AllEqual, IsSortedBy, ContainsAll / 조건자
- **Utilities (12)**: ForEach, ForEachIndexed, Join, Clone, Shuffle, Zip, Unzip, Tap / 유틸리티
- **Combinatorial (2)**: Permutations, Combinations / 조합 작업
- **Statistics (8)**: Median, Mode, Frequencies, Percentile, StandardDeviation, Variance, MostCommon, LeastCommon / 통계
- **Diff/Comparison (4)**: Diff, DiffBy, EqualUnordered, HasDuplicates / 차이/비교
- **Index-based (3)**: FindIndices, AtIndices, RemoveIndices / 인덱스 기반
- **Conditional (3)**: ReplaceIf, ReplaceAll, UpdateWhere / 조건부
- **Advanced (4)**: Scan, ZipWith, RotateLeft, RotateRight / 고급

```go
import "github.com/arkd0ng/go-utils/sliceutil"

// Filter and Map pipeline / 필터 및 Map 파이프라인
numbers := []int{1, 2, 3, 4, 5, 6}
evens := sliceutil.Filter(numbers, func(n int) bool { return n%2 == 0 })
doubled := sliceutil.Map(evens, func(n int) int { return n * 2 })
// Result: [4 8 12]

// GroupBy and aggregate / 그룹화 및 집계
users := []User{
    {Name: "Alice", City: "Seoul", Age: 28},
    {Name: "Bob", City: "Busan", Age: 35},
    {Name: "Charlie", City: "Seoul", Age: 42},
}
byCity := sliceutil.GroupBy(users, func(u User) string { return u.City })
// Map[Seoul: [{Alice Seoul 28} {Charlie Seoul 42}], Busan: [{Bob Busan 35}]]

// Set operations / 집합 작업
set1 := []int{1, 2, 3, 4, 5}
set2 := []int{4, 5, 6, 7, 8}
union := sliceutil.Union(set1, set2)        // [1 2 3 4 5 6 7 8]
intersection := sliceutil.Intersection(set1, set2) // [4 5]

// Functional operations / 함수형 작업
result := sliceutil.Reduce(numbers, 0, func(acc, n int) int {
    return acc + n
}) // 21

// Batch processing / 배치 처리
data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
chunks := sliceutil.Chunk(data, 3)
// Result: [[1 2 3] [4 5 6] [7 8 9] [10]]
```

**Before vs After**:
```go
// ❌ Before: 20+ lines with standard Go
numbers := []int{1, 2, 3, 4, 5, 6}
var evens []int
for _, n := range numbers {
    if n%2 == 0 {
        evens = append(evens, n)
    }
}
var doubled []int
for _, n := range evens {
    doubled = append(doubled, n*2)
}
// ... 더 많은 코드

// ✅ After: 2 lines with this package
evens := sliceutil.Filter(numbers, func(n int) bool { return n%2 == 0 })
doubled := sliceutil.Map(evens, func(n int) int { return n * 2 })
```

**Documentation / 문서**:
- [Package README](./sliceutil/README.md) - Quick start and examples / 빠른 시작 및 예제
- [User Manual](./docs/sliceutil/USER_MANUAL.md) - Comprehensive user guide (3,887 lines) / 포괄적인 사용자 가이드 (3,887줄)
- [Developer Guide](./docs/sliceutil/DEVELOPER_GUIDE.md) - Technical documentation (2,205 lines) / 기술 문서 (2,205줄)
- [Performance Benchmarks](./docs/sliceutil/PERFORMANCE_BENCHMARKS.md) - Real benchmark data / 실제 벤치마크 데이터

**[→ View full documentation / 전체 문서 보기](./sliceutil/README.md)**

---

### ✅ [maputil](./maputil/) - Map Utilities

Extreme simplicity map utilities - reduce 20 lines of repetitive map manipulation code to just 1-2 lines with **99 type-safe functions**.

극도로 간단한 맵 유틸리티 - 20줄의 반복적인 맵 조작 코드를 단 1-2줄로 줄이며, **99개의 타입 안전 함수**를 제공합니다.

**Core Features**: 99 functions across 14 categories, Go 1.18+ generics, functional programming style, immutable operations, zero dependencies, 92.8% test coverage / 14개 카테고리에 걸쳐 99개 함수, Go 1.18+ 제네릭, 함수형 프로그래밍 스타일, 불변 작업, 제로 의존성, 92.8% 테스트 커버리지

**Categories / 카테고리**:
- **Basic Operations (11)**: Get, Set, Delete, Has, Clone, Equal, IsEmpty / 기본 작업
- **Transformation (10)**: Map, MapKeys, Invert, Flatten, Partition / 변환
- **Aggregation (9)**: Reduce, Sum, Min, Max, Average, GroupBy, CountBy, Median, Frequencies / 집계 및 통계
- **Merge Operations (8)**: Merge, Union, Intersection, Difference / 병합 작업
- **Filter Operations (7)**: Filter, Pick, Omit, Partition / 필터 작업
- **Conversion (10)**: Keys, Values, Entries, ToJSON, FromJSON, ToYAML, FromYAML / 변환 (YAML 지원)
- **Predicate Checks (7)**: Every, Some, None, HasValue, IsSubset / 조건 검사
- **Key Operations (8)**: KeysSorted, RenameKey, SwapKeys, FindKey / 키 작업
- **Value Operations (7)**: ValuesSorted, UniqueValues, ReplaceValue / 값 작업
- **Comparison (6)**: Diff, Compare, CommonKeys, AllKeys / 비교
- **Utility Functions (6)**: ForEach, GetMany, SetMany, Tap, ContainsAllKeys, Apply / 유틸리티
- **Default Functions (3)**: GetOrSet, SetDefault, Defaults / 기본값 관리
- **Nested Map Functions (5)**: GetNested, SetNested, HasNested, DeleteNested, SafeGet / 중첩 맵 작업
- **Statistics Functions (2)**: Median, Frequencies / 통계 함수

```go
import "github.com/arkd0ng/go-utils/maputil"

// Filter map by value / 값으로 맵 필터링
data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
result := maputil.Filter(data, func(k string, v int) bool {
    return v > 2
}) // map[string]int{"c": 3, "d": 4}

// Transform values / 값 변환
doubled := maputil.MapValues(data, func(v int) int {
    return v * 2
}) // map[string]int{"a": 2, "b": 4, "c": 6, "d": 8}

// Merge maps / 맵 병합
map1 := map[string]int{"a": 1, "b": 2}
map2 := map[string]int{"b": 3, "c": 4}
merged := maputil.Merge(map1, map2) // map[string]int{"a": 1, "b": 3, "c": 4}

// Group slice by key / 키로 슬라이스 그룹화
users := []User{
    {Name: "Alice", City: "Seoul"},
    {Name: "Bob", City: "Seoul"},
    {Name: "Charlie", City: "Busan"},
}
byCity := maputil.GroupBy[string, User, string](users, func(u User) string {
    return u.City
})
// Map[Seoul: [{Alice Seoul} {Bob Seoul}], Busan: [{Charlie Busan}]]

// Set operations / 집합 작업
m1 := map[string]int{"a": 1, "b": 2, "c": 3}
m2 := map[string]int{"b": 2, "c": 4, "d": 5}
intersection := maputil.Intersection(m1, m2) // map[string]int{"b": 2}
difference := maputil.Difference(m1, m2)     // map[string]int{"a": 1}

// Nested map operations (NEW) / 중첩 맵 작업 (신규)
config := map[string]interface{}{
    "server": map[string]interface{}{
        "host": "localhost",
        "port": 8080,
    },
}
host, ok := maputil.GetNested(config, "server", "host") // "localhost", true
maputil.SetNested(config, "api.example.com", "server", "host")

// Default value management (NEW) / 기본값 관리 (신규)
cache := map[string]int{"a": 1}
value := maputil.GetOrSet(cache, "b", 10) // Returns 10 and sets cache["b"] = 10

// Statistics (NEW) / 통계 (신규)
scores := map[string]int{"Alice": 85, "Bob": 90, "Charlie": 75}
median, _ := maputil.Median(scores) // 85.0
freq := maputil.Frequencies(scores) // Count occurrences of each score

// YAML conversion (NEW) / YAML 변환 (신규)
yamlStr, _ := maputil.ToYAML(config)
parsedConfig, _ := maputil.FromYAML(yamlStr)
```

**Before vs After**:
```go
// ❌ Before: 20+ lines with standard Go
data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
result := make(map[string]int)
for k, v := range data {
    if v > 2 {
        result[k] = v
    }
}
// ... 더 많은 코드

// ✅ After: 1 line with this package
result := maputil.Filter(data, func(k string, v int) bool { return v > 2 })
```

**Documentation / 문서**:
- [Package README](./maputil/README.md) - Quick start and examples / 빠른 시작 및 예제
- [User Manual](./docs/maputil/USER_MANUAL.md) - Comprehensive user guide (2,207 lines) / 포괄적인 사용자 가이드 (2,207줄)
- [Developer Guide](./docs/maputil/DEVELOPER_GUIDE.md) - Technical documentation (2,356 lines) / 기술 문서 (2,356줄)
- [Design Plan](./docs/maputil/DESIGN_PLAN.md) - Architecture and design decisions / 아키텍처 및 설계 결정

**[→ View full documentation / 전체 문서 보기](./maputil/README.md)**

---

### ✅ [fileutil](./fileutil/) - File and Path Utilities

Extreme simplicity file and path utilities - reduce 20+ lines of repetitive file manipulation code to just 1-2 lines with **~91 cross-platform functions**.

극도로 간단한 파일 및 경로 유틸리티 - 20줄의 반복적인 파일 조작 코드를 단 1-2줄로 줄이며, **약 91개의 크로스 플랫폼 함수**를 제공합니다.

**Core Features**: ~91 functions across 12 categories, automatic directory creation, cross-platform compatibility, buffered I/O, atomic operations, progress callbacks, multiple hash algorithms, zero external dependencies / 12개 카테고리에 걸쳐 약 91개 함수, 자동 디렉토리 생성, 크로스 플랫폼 호환성, 버퍼링된 I/O, 원자적 작업, 진행 상황 콜백, 여러 해시 알고리즘, 외부 의존성 없음

**Categories / 카테고리**:
- **File Reading (8)**: ReadFile, ReadString, ReadLines, ReadJSON, ReadYAML, ReadCSV, ReadBytes, ReadChunk / 파일 읽기
- **File Writing (11)**: WriteFile, WriteString, WriteLines, WriteJSON, WriteYAML, WriteCSV, WriteAtomic, Append* / 파일 쓰기
- **File Information (15)**: Exists, IsFile, IsDir, Size, SizeHuman, Chmod, Chown, ModTime, Touch / 파일 정보
- **Path Operations (18)**: Join, Split, Base, Dir, Ext, Abs, CleanPath, Normalize, IsAbs, IsValid, IsSafe, Match, Glob / 경로 작업
- **File Copying (4)**: CopyFile, CopyDir, CopyRecursive, SyncDirs (with progress callbacks) / 파일 복사
- **File Moving (5)**: MoveFile, MoveDir, Rename, RenameExt, SafeMove / 파일 이동
- **File Deleting (7)**: DeleteFile, DeleteDir, DeleteRecursive, DeletePattern, DeleteFiles, Clean, RemoveEmpty / 파일 삭제
- **Directory Operations (13)**: MkdirAll, CreateTemp, IsEmpty, DirSize, ListFiles, Walk, FindFiles / 디렉토리 작업
- **File Hashing (10)**: MD5, SHA1, SHA256, SHA512, Hash, CompareFiles, CompareHash, Checksum, VerifyChecksum / 파일 해싱

```go
import "github.com/arkd0ng/go-utils/fileutil"

// Write file with auto directory creation / 자동 디렉토리 생성과 함께 파일 쓰기
err := fileutil.WriteString("path/to/file.txt", "Hello, World!")

// Read file / 파일 읽기
content, err := fileutil.ReadString("path/to/file.txt")

// Copy with progress / 진행 상황과 함께 복사
err = fileutil.CopyFile("large.dat", "backup.dat",
    fileutil.WithProgress(func(written, total int64) {
        percent := float64(written) / float64(total) * 100
        fmt.Printf("\rProgress: %.1f%%", percent)
    }))

// Calculate file hash / 파일 해시 계산
hash, err := fileutil.SHA256("file.dat")

// Find all .txt files / 모든 .txt 파일 찾기
txtFiles, err := fileutil.FindFiles(".", func(path string, info os.FileInfo) bool {
    return fileutil.Ext(path) == ".txt"
})

// Atomic write (safe update) / 원자적 쓰기 (안전한 업데이트)
err = fileutil.WriteAtomic("important.json", data)

// JSON/YAML support / JSON/YAML 지원
var config Config
err = fileutil.ReadJSON("config.json", &config)
err = fileutil.WriteYAML("config.yaml", config)
```

**Before vs After**:
```go
// ❌ Before: 20+ lines with standard Go
dir := filepath.Dir(path)
if err := os.MkdirAll(dir, 0755); err != nil {
    return err
}
file, err := os.Create(path)
if err != nil {
    return err
}
defer file.Close()
if _, err := file.WriteString(content); err != nil {
    return err
}
// ... 더 많은 코드

// ✅ After: 1 line with this package
err := fileutil.WriteString(path, content)
```

**Documentation / 문서**:
- [Package README](./fileutil/README.md) - Quick start and examples / 빠른 시작 및 예제

**[→ View full documentation / 전체 문서 보기](./fileutil/README.md)**

---

### ✅ [httputil](./httputil/) - HTTP Client Utilities

Extremely simple HTTP client that reduces 30+ lines of boilerplate code to just 2-3 lines with **automatic retry logic**, **JSON handling**, **rich error types**, and **advanced features**.

극도로 간단한 HTTP 클라이언트로 30줄 이상의 보일러플레이트 코드를 단 2-3줄로 줄이며, **자동 재시도 로직**, **JSON 처리**, **풍부한 에러 타입**, **고급 기능**을 제공합니다.

**Core Features**: RESTful methods (GET/POST/PUT/PATCH/DELETE), automatic JSON encoding/decoding, smart retry with exponential backoff, 14 configuration options, rich error types, zero external dependencies / RESTful 메서드, 자동 JSON 인코딩/디코딩, 지수 백오프를 통한 스마트 재시도, 14개 설정 옵션, 풍부한 에러 타입, 외부 의존성 없음

**API Levels / API 레벨**:
- **Simple API (26+ functions)**: Package-level convenience functions / 패키지 레벨 편의 함수
- **Client API**: Configured HTTP client for multiple requests / 여러 요청을 위한 설정된 HTTP 클라이언트
- **Response Helpers (20+ methods)**: Status checks, body access, headers / 상태 확인, 본문 접근, 헤더
- **File Operations**: Upload/download with progress tracking / 진행 상황 추적이 있는 업로드/다운로드
- **URL Builder**: Fluent API for building URLs / URL 구축을 위한 Fluent API
- **Form Builder**: Fluent API for building forms / 폼 구축을 위한 Fluent API
- **Cookie Management**: In-memory and persistent cookie jars / 메모리 내 및 지속성 쿠키 저장소
- **Options Pattern**: 14 built-in options (timeout, auth, retry, cookies, etc.) / 14개 내장 옵션
- **Error Types**: HTTPError, RetryError, TimeoutError / 에러 타입

```go
import "github.com/arkd0ng/go-utils/httputil"

// Simple GET request / 간단한 GET 요청
var users []User
err := httputil.Get("https://api.example.com/users", &users,
    httputil.WithBearerToken("your-token"))

// POST request with automatic JSON handling / 자동 JSON 처리를 가진 POST 요청
payload := CreateUserRequest{Name: "John", Email: "john@example.com"}
var response CreateUserResponse
err := httputil.Post("https://api.example.com/users", payload, &response,
    httputil.WithTimeout(30*time.Second),
    httputil.WithRetry(3))

// Client with cookies and base URL / 쿠키와 베이스 URL을 가진 클라이언트
client := httputil.NewClient(
    httputil.WithBaseURL("https://api.example.com/v1"),
    httputil.WithBearerToken("your-token"),
    httputil.WithRetry(5),
    httputil.WithPersistentCookies("cookies.json"))

client.Get("/users", &users)
client.Post("/users", newUser, &created)
client.Delete("/users/123", nil)

// File download with progress / 진행 상황과 함께 파일 다운로드
err = httputil.DownloadFile(
    "https://example.com/large-file.zip",
    "./downloads/file.zip",
    httputil.WithProgress(func(bytesRead, totalBytes int64) {
        progress := float64(bytesRead) / float64(totalBytes) * 100
        fmt.Printf("\rDownloading: %.2f%%", progress)
    }))

// Response helpers / 응답 헬퍼
resp, _ := httputil.DoRaw("GET", "https://api.example.com/users", nil)
if resp.IsSuccess() {
    bodyString := resp.String()
    var users []User
    resp.JSON(&users)
}

// URL and Form builders / URL 및 Form 빌더
url := httputil.NewURL("https://api.example.com").
    Path("users", "search").
    Param("q", "golang").
    Build()

form := httputil.NewForm().
    Set("username", "john").
    Set("email", "john@example.com").
    AddIf(hasPromo, "promo_code", "SAVE20")
```

**Before vs After**:
```go
// ❌ Before: 30+ lines with standard Go
client := &http.Client{Timeout: 30 * time.Second}
req, _ := http.NewRequest("GET", url, nil)
req.Header.Set("Authorization", "Bearer token")
req.Header.Set("Content-Type", "application/json")
resp, _ := client.Do(req)
defer resp.Body.Close()
if resp.StatusCode >= 400 {
    body, _ := io.ReadAll(resp.Body)
    return fmt.Errorf("HTTP %d: %s", resp.StatusCode, body)
}
var users []User
json.NewDecoder(resp.Body).Decode(&users)
// Plus retry logic, error handling... 20+ more lines

// ✅ After: 2 lines with httputil
var users []User
err := httputil.Get(url, &users, httputil.WithBearerToken("token"))
```

**Documentation / 문서**:
- [Package README](./httputil/README.md) - Quick start and API reference / 빠른 시작 및 API 참조
- [User Manual](./docs/httputil/USER_MANUAL.md) - Comprehensive usage guide / 종합 사용 가이드
- [Developer Guide](./docs/httputil/DEVELOPER_GUIDE.md) - Architecture and internals / 아키텍처 및 내부 구조
- [Work Plan](./docs/httputil/WORK_PLAN.md) - Development roadmap / 개발 로드맵

**[→ View full documentation / 전체 문서 보기](./httputil/README.md)**

---

### ✅ [websvrutil](./websvrutil/) - Web Server Utilities

Extreme simplicity web server utilities - reduce 50+ lines of server setup code to just 5 lines.

극도로 간단한 웹 서버 유틸리티 - 50줄 이상의 서버 설정 코드를 단 5줄로 줄입니다.

**Core Features**: Simple Router, RESTful routing with path parameters, middleware chaining, handler helpers / 간단한 라우터, 경로 매개변수를 가진 RESTful 라우팅, 미들웨어 체이닝, 핸들러 헬퍼

**Middleware (10+ built-in)**: CORS, logging, recovery, auth, rate limiting, timeout, compression / 미들웨어 (10개 이상 내장): CORS, 로깅, 복구, 인증, 속도 제한, 타임아웃, 압축

**Handler Helpers**: JSON response, error response, file serving, redirect / 핸들러 헬퍼: JSON 응답, 에러 응답, 파일 서빙, 리다이렉트

**Request/Response Utilities**: Body binding, cookie management, headers, query/path params / 요청/응답 유틸리티: 본문 바인딩, 쿠키 관리, 헤더, 쿼리/경로 매개변수

**Server Management**: Graceful shutdown, signal handling, health check / 서버 관리: 정상 종료, 시그널 처리, 헬스 체크

**Template Engine**: Auto-discovery, hot reload, layout support / 템플릿 엔진: 자동 발견, 핫 리로드, 레이아웃 지원

```go
import "github.com/arkd0ng/go-utils/websvrutil"

// Create app with middleware / 미들웨어와 함께 앱 생성
app := websvrutil.New()
app.Use(websvrutil.Logger(), websvrutil.Recovery())

// Define routes / 라우트 정의
app.GET("/", func(c *websvrutil.Context) error {
    return c.JSON(200, map[string]string{"message": "Hello, World!"})
})

app.GET("/users/:id", func(c *websvrutil.Context) error {
    id := c.Param("id")
    return c.JSON(200, map[string]string{"user_id": id})
})

// Run with graceful shutdown / 정상 종료와 함께 실행
app.RunWithGracefulShutdown(":8080", 10*time.Second)
```

**Before vs After**:
```go
// ❌ Before: 50+ lines with standard net/http
mux := http.NewServeMux()
mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "Hello"})
})
// ... CORS setup, logging, recovery, graceful shutdown - 40+ more lines

// ✅ After: 5 lines with websvrutil
app := websvrutil.New()
app.Use(websvrutil.Logger(), websvrutil.Recovery(), websvrutil.CORS())
app.GET("/", func(c *websvrutil.Context) error {
    return c.JSON(200, map[string]string{"message": "Hello"})
})
app.RunWithGracefulShutdown(":8080", 10*time.Second)
```

**[→ View full documentation / 전체 문서 보기](./websvrutil/README.md)**

---

### ✅ [errorutil](./errorutil/) - Error Handling Utilities

Comprehensive error handling utilities with error codes, context, and stack traces.

에러 코드, 컨텍스트, 스택 트레이스를 갖춘 포괄적인 에러 처리 유틸리티.

**Core Features**: Error creation/wrapping with codes, error inspection, standard library compatible / 핵심 기능: 코드를 가진 에러 생성/래핑, 에러 검사, 표준 라이브러리 호환

**Error Creation (6 functions)**: New, Newf, WithCode, WithCodef, WithNumericCode, WithNumericCodef / 에러 생성 (6개 함수)

**Error Wrapping (6 functions)**: Wrap, Wrapf, WrapWithCode, WrapWithCodef, WrapWithNumericCode, WrapWithNumericCodef / 에러 래핑 (6개 함수)

**Error Inspection (9 functions)**: HasCode, HasNumericCode, GetCode, GetNumericCode, GetStackTrace, GetContext, Root, UnwrapAll, Contains / 에러 검사 (9개 함수)

**Standard Compatible**: Works with errors.Is() and errors.As() / 표준 호환: errors.Is() 및 errors.As()와 동작

**Zero Dependencies**: Standard library only / 외부 의존성 없음: 표준 라이브러리만

**Test Coverage**: 99.2% / 테스트 커버리지: 99.2%

```go
import "github.com/arkd0ng/go-utils/errorutil"

// Create error with code / 코드를 가진 에러 생성
err := errorutil.WithCode("DB_ERROR", "failed to connect to database")

// Wrap error with additional code / 추가 코드로 에러 래핑
err = errorutil.WrapWithCode("API_ERROR", err, "failed to fetch user data")

// Check for specific error code / 특정 에러 코드 확인
if errorutil.HasCode(err, "DB_ERROR") {
    log.Println("Database error occurred")
}

// Get error code / 에러 코드 가져오기
code := errorutil.GetCode(err) // Returns "API_ERROR"

// Numeric error codes / 숫자 에러 코드
httpErr := errorutil.WithNumericCode(500, "internal server error")
statusCode := errorutil.GetNumericCode(httpErr) // Returns 500

// Get root cause / 근본 원인 가져오기
root := errorutil.Root(err)

// Check if error chain contains specific error / 에러 체인이 특정 에러를 포함하는지 확인
if errorutil.Contains(err, io.EOF) {
    log.Println("EOF found in error chain")
}
```

**[→ View full documentation / 전체 문서 보기](./errorutil/README.md)**

---

### ✅ [validation](./validation/) - Fluent Validation Library

Fluent, type-safe validation library with 135+ built-in validators. Reduce 20-30 lines of validation code to just 1-2 lines.

135개 이상의 내장 검증기를 가진 플루언트하고 타입 안전한 검증 라이브러리입니다. 20-30줄의 검증 코드를 단 1-2줄로 줄입니다.

**Features**: Fluent API, 135+ validators, bilingual errors (EN/KR), 100% test coverage, zero dependencies / Fluent API, 135개 이상 검증기, 이중 언어 에러 (영어/한글), 100% 테스트 커버리지, 제로 의존성

**Validator Categories**: String (19), Numeric (10), Collection (10), Comparison (11), Type (7), Network (5), DateTime (4), File (6), Security (6), Credit Card (3), Business Codes (3), Color (4), Data Format (4), Format (3), Geographic (3), Logical (4), Range (3)

**검증기 카테고리**: 문자열 (19), 숫자 (10), 컬렉션 (10), 비교 (11), 타입 (7), 네트워크 (5), 날짜/시간 (4), 파일 (6), 보안 (6), 신용카드 (3), 비즈니스 코드 (3), 색상 (4), 데이터 형식 (4), 포맷 (3), 지리 (3), 논리 (4), 범위 (3)

```go
import "github.com/arkd0ng/go-utils/validation"

// Single field validation / 단일 필드 검증
email := "user@example.com"
v := validation.New(email, "email")
v.Required().Email().MaxLength(100)

if err := v.Validate(); err != nil {
    log.Fatal(err)
}

// Multi-field validation / 다중 필드 검증
type User struct {
    Username string
    Email    string
    Age      int
}

user := User{
    Username: "johndoe",
    Email:    "john@example.com",
    Age:      25,
}

mv := validation.NewValidator()
mv.Field(user.Username, "username").Required().MinLength(3).MaxLength(20).Alphanumeric()
mv.Field(user.Email, "email").Required().Email()
mv.Field(user.Age, "age").Positive().Between(13, 120)

if err := mv.Validate(); err != nil {
    // Get all validation errors / 모든 검증 에러 가져오기
    verrs := err.(validation.ValidationErrors)
    for _, e := range verrs {
        log.Printf("Field '%s': %s", e.Field, e.Message)
    }
}

// Custom validators / 사용자 정의 검증기
password := "MyPassword123!"
v := validation.New(password, "password")
v.MinLength(8).Custom(func(val interface{}) bool {
    s := val.(string)
    return strings.ContainsAny(s, "!@#$%^&*()")
}, "Password must contain at least one special character")
```

**[→ View full documentation / 전체 문서 보기](./validation/README.md)**

---

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

### v1.8.x (Current / 현재)

- **NEW**: `maputil` package - Map utilities / 맵 유틸리티
  - 20 lines → 1-2 lines code reduction / 20줄 → 1-2줄 코드 감소
  - 81 functions across 10 categories / 10개 카테고리에 걸쳐 81개 함수
  - Go 1.18+ generics for type safety / Go 1.18+ 제네릭으로 타입 안전성
  - Functional programming style (Map, Filter, Reduce) / 함수형 프로그래밍 스타일
  - Immutable operations (original maps unchanged) / 불변 작업 (원본 맵 변경 없음)
  - Zero dependencies / 제로 의존성
  - Entry type for key-value pairs / 키-값 쌍을 위한 Entry 타입
  - Type constraints (Number, Ordered) / 타입 제약조건

### v1.7.x

- **NEW**: `sliceutil` package - Slice utilities / 슬라이스 유틸리티
  - 20 lines → 1 line code reduction / 20줄 → 1줄 코드 감소
  - 95 functions across 14 categories / 14개 카테고리에 걸쳐 95개 함수
  - Go 1.18+ generics for type safety / Go 1.18+ 제네릭으로 타입 안전성
  - Functional programming style / 함수형 프로그래밍 스타일
  - Immutable operations / 불변 작업
  - Zero dependencies / 제로 의존성
  - 100% test coverage / 100% 테스트 커버리지
  - Comprehensive documentation (USER_MANUAL, DEVELOPER_GUIDE, PERFORMANCE_BENCHMARKS) / 포괄적인 문서화

### v1.6.x

- **NEW**: `timeutil` package - Time and date utilities / 시간 및 날짜 유틸리티
  - 20 lines → 1 line code reduction / 20줄 → 1줄 코드 감소
  - 114 functions across 12 categories / 12개 카테고리에 걸쳐 114개 함수
  - KST (GMT+9) as default timezone / KST (GMT+9)를 기본 타임존으로 설정
  - Custom format tokens (YYYY-MM-DD) / 커스텀 포맷 토큰
  - Business day support with Korean holidays / 한국 공휴일을 포함한 영업일 지원
  - Thread-safe timezone caching / 스레드 안전 타임존 캐싱

### v1.5.x

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
