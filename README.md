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

### v0.2.0 (Current / 현재)

- **BREAKING CHANGE**: Refactored to subpackage structure / 서브패키지 구조로 리팩토링
  - Moved `GenRandomString` to `random.GenString` / `GenRandomString`을 `random.GenString`으로 이동
  - Import path changed / import 경로 변경
- Added bilingual documentation (English/Korean) / 이중 언어 문서 추가 (영문/한글)
- Improved package organization / 패키지 구조 개선

### v0.1.0 (Initial Release / 첫 릴리스)

- Added `random` package with string generation utilities / 문자열 생성 유틸리티가 포함된 `random` 패키지 추가
