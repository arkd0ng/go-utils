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
│   └── mysql/       # Extreme simplicity MySQL client / 극도로 간단한 MySQL 클라이언트
├── stringutil/      # String manipulation (coming soon) / 문자열 처리 (예정)
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

### 🔜 Coming Soon / 개발 예정

- **stringutil** - String manipulation utilities / 문자열 처리 유틸리티
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

### v1.3.x (Current / 현재)

- **NEW**: `database/mysql` package - Extreme simplicity MySQL/MariaDB client / 극도로 간단한 MySQL/MariaDB 클라이언트
  - 30 lines → 2 lines code reduction / 30줄 → 2줄 코드 감소
  - Zero-downtime credential rotation / 무중단 자격 증명 순환
  - Auto everything: connection, retry, cleanup / 모든 것 자동화
  - 7 Simple API methods: SelectAll, SelectOne, Insert, Update, Delete, Count, Exists
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
