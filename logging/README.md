# Logging Package

A simple and powerful logging package with automatic file rotation and structured logging support.

파일 로테이션과 구조화된 로깅을 지원하는 간단하고 강력한 로깅 패키지입니다.

## Installation / 설치

```bash
go get github.com/arkd0ng/go-utils/logging
```

## Features / 주요 기능

- **Zero Configuration** - Works out of the box with sensible defaults / 기본 설정으로 즉시 사용 가능
- **Automatic File Rotation** - Uses lumberjack for log file management / lumberjack을 사용한 로그 파일 자동 관리
- **Multiple Log Levels** - DEBUG, INFO, WARN, ERROR, FATAL / 5가지 로그 레벨 지원
- **Two Logging Styles** - Both structured (key-value) and Printf-style logging / 구조화 및 Printf 스타일 모두 지원
- **Structured Logging** - Key-value pairs for searchable logs / 검색 가능한 키-값 쌍 로깅
- **Printf-Style Logging** - Familiar `fmt.Printf` syntax / 친숙한 `fmt.Printf` 문법
- **Colored Output** - Color-coded console output / 색상으로 구분된 콘솔 출력
- **Multiple Loggers** - Create separate loggers for different purposes / 용도별 독립 로거 생성
- **Automatic Banner** - Prints banner on logger creation by default / 로거 생성 시 자동 배너 출력
- **Banner Support** - ASCII art banners for application startup / 애플리케이션 시작 배너 지원
- **Thread-Safe** - Safe for concurrent use / 동시성 안전

## Quick Start / 빠른 시작

### 1. Simple Usage / 간단한 사용

```go
package main

import "github.com/arkd0ng/go-utils/logging"

func main() {
    // Create default logger / 기본 로거 생성
    logger := logging.Default()
    defer logger.Close()

    // Print banner / 배너 출력
    logger.Banner("My Application", "v1.0.0")

    // Log messages / 메시지 로깅
    logger.Info("Application started")
    logger.Warn("This is a warning")
    logger.Error("An error occurred")

    // Printf-style logging / Printf 스타일 로깅
    port := 8080
    logger.Infof("Server listening on port %d", port)
}
```

### 2. Custom Configuration / 커스텀 설정

```go
logger, err := logging.New(
    logging.WithFilePath("./logs/myapp.log"),
    logging.WithMaxSize(50),       // 50 MB
    logging.WithMaxBackups(3),     // Keep 3 backups / 3개 백업 유지
    logging.WithMaxAge(28),         // Keep for 28 days / 28일 동안 보관
    logging.WithLevel(logging.DEBUG),
    logging.WithPrefix("[APP]"),
)
if err != nil {
    panic(err)
}
defer logger.Close()
```

### 3. Multiple Loggers / 여러 로거 사용

```go
// Application logger / 애플리케이션 로거
appLogger, _ := logging.New(
    logging.WithFilePath("./logs/app.log"),
    logging.WithPrefix("[APP]"),
)

// Database logger / 데이터베이스 로거
dbLogger, _ := logging.New(
    logging.WithFilePath("./logs/database.log"),
    logging.WithPrefix("[DB]"),
)

// API logger / API 로거
apiLogger, _ := logging.New(
    logging.WithFilePath("./logs/api.log"),
    logging.WithPrefix("[API]"),
)

appLogger.Info("Application started")
dbLogger.Info("Database connected")
apiLogger.Info("API server listening on :8080")
```

## Usage / 사용법

### Log Levels / 로그 레벨

```go
logger := logging.Default()

logger.Debug("Detailed debugging information")
logger.Info("General informational messages")
logger.Warn("Warning messages")
logger.Error("Error messages")
logger.Fatal("Fatal errors - exits program") // Calls os.Exit(1) / os.Exit(1) 호출
```

### Two Logging Styles: Structured vs Printf / 두 가지 로깅 스타일: 구조화 vs Printf

이 로깅 패키지는 두 가지 로깅 스타일을 모두 지원합니다:

This logging package supports two logging styles:

#### 1. Structured Logging (권장 / Recommended)

구조화된 로깅은 **키-값 쌍**을 사용하여 로그를 기록합니다. 로그 분석 도구에서 검색하고 필터링하기 쉽습니다.

Structured logging uses **key-value pairs**. This format is easy to search and filter in log analysis tools.

```go
// Structured logging with key-value pairs / 키-값 쌍을 사용한 구조화된 로깅
logger.Info("User login",
    "user_id", 12345,
    "username", "john.doe",
    "ip", "192.168.1.100",
)

// Output / 출력:
// 2025-10-10 15:30:45 [INFO] User login user_id=12345 username=john.doe ip=192.168.1.100
```

**장점 (Advantages):**
- 로그 분석 도구에서 쉽게 파싱 가능 / Easy to parse in log analysis tools
- 필드별 검색/필터링 가능 / Searchable and filterable by field
- 구조화된 데이터 형식 / Structured data format

**사용 사례 (Use Cases):**
- 프로덕션 환경 / Production environments
- 로그 분석이 필요한 경우 / When log analysis is needed
- 자동화된 모니터링 / Automated monitoring

#### 2. Printf-Style Logging (친숙함 / Familiar)

Printf 스타일은 **형식 문자열**을 사용하여 로그를 기록합니다. 표준 라이브러리의 `fmt.Printf`와 동일한 방식입니다.

Printf-style uses **format strings**, just like `fmt.Printf` from the standard library.

```go
// Printf-style logging with format string / 형식 문자열을 사용한 Printf 스타일 로깅
logger.Infof("User login: %s (ID: %d, IP: %s)", "john.doe", 12345, "192.168.1.100")

// Output / 출력:
// 2025-10-10 15:30:45 [INFO] User login: john.doe (ID: 12345, IP: 192.168.1.100)
```

**장점 (Advantages):**
- 친숙한 문법 (표준 라이브러리와 동일) / Familiar syntax (same as standard library)
- 읽기 쉬운 메시지 / Human-readable messages
- 빠른 작성 / Quick to write

**사용 사례 (Use Cases):**
- 개발/디버깅 환경 / Development/debugging
- 사람이 읽는 로그 메시지 / Human-readable log messages
- 간단한 로깅 / Simple logging

### 비교 예제 / Comparison Examples

```go
user := "alice"
userID := 67890
loginTime := time.Now()

// 1. Structured logging (키-값 쌍) / Structured logging (key-value pairs)
logger.Info("User login successful",
    "username", user,
    "user_id", userID,
    "timestamp", loginTime,
)
// Output / 출력:
// 2025-10-10 15:30:45 [INFO] User login successful username=alice user_id=67890 timestamp=2025-10-10 15:30:45.123456789 +0900 KST

// 2. Printf-style logging (형식 문자열) / Printf-style logging (format string)
logger.Infof("User login successful: %s (ID: %d) at %s", user, userID, loginTime.Format("15:04:05"))
// Output / 출력:
// 2025-10-10 15:30:45 [INFO] User login successful: alice (ID: 67890) at 15:30:45
```

### 모든 로그 레벨에서 두 스타일 모두 지원 / Both Styles Supported for All Levels

```go
// Structured logging / 구조화된 로깅
logger.Debug("Debug info", "key", "value")
logger.Info("Info message", "key", "value")
logger.Warn("Warning", "key", "value")
logger.Error("Error occurred", "key", "value")

// Printf-style logging / Printf 스타일 로깅
logger.Debugf("Debug: %s = %v", "key", "value")
logger.Infof("Info: %s = %v", "key", "value")
logger.Warnf("Warning: %s = %v", "key", "value")
logger.Errorf("Error: %s = %v", "key", "value")
```

### 어떤 스타일을 사용해야 할까요? / Which Style Should You Use?

| 상황 / Situation | 권장 스타일 / Recommended Style |
|-----------------|-------------------------------|
| 프로덕션 환경 / Production | Structured (`Info`, `Error`, etc.) |
| 로그 분석/모니터링 / Log analysis/monitoring | Structured (`Info`, `Error`, etc.) |
| 개발/디버깅 / Development/debugging | Printf (`Infof`, `Errorf`, etc.) |
| 사람이 읽는 메시지 / Human-readable messages | Printf (`Infof`, `Errorf`, etc.) |
| 빠른 프로토타이핑 / Quick prototyping | Printf (`Infof`, `Errorf`, etc.) |

**💡 권장사항 / Recommendation:**

프로덕션 환경에서는 **구조화된 로깅(Structured Logging)**을 사용하세요. 로그 분석 도구(예: ELK, Splunk, Datadog)에서 쉽게 검색하고 필터링할 수 있습니다.

For production environments, use **Structured Logging**. It's easier to search and filter in log analysis tools (e.g., ELK, Splunk, Datadog).

개발 중이거나 빠르게 로그를 확인하고 싶을 때는 **Printf 스타일**이 더 편리할 수 있습니다.

During development or when you want to quickly check logs, **Printf-style** may be more convenient.

### Setting Log Level / 로그 레벨 설정

```go
// Set minimum log level / 최소 로그 레벨 설정
logger.SetLevel(logging.WARN)

// These won't be logged / 이것들은 로깅되지 않음
logger.Debug("Debug message")
logger.Info("Info message")

// These will be logged / 이것들은 로깅됨
logger.Warn("Warning message")
logger.Error("Error message")
```

### Automatic Banner / 자동 배너

By default, a banner is automatically printed when a logger is created.

기본적으로 로거 생성 시 자동으로 배너가 출력됩니다.

```go
// Default auto banner (prints "Application v1.0.0")
// 기본 자동 배너 ("Application v1.0.0" 출력)
logger := logging.Default()
// Banner is automatically printed / 배너가 자동으로 출력됨

// Custom app name and version / 커스텀 앱 이름과 버전
logger, _ := logging.New(
    logging.WithAppName("MyApp"),
    logging.WithAppVersion("v2.0.0"),
)
// Prints "MyApp v2.0.0" banner automatically / "MyApp v2.0.0" 배너가 자동으로 출력됨

// Convenience function / 편의 함수
logger, _ := logging.New(
    logging.WithBanner("ProductionAPI", "v3.2.1"),
)
// Prints "ProductionAPI v3.2.1" banner automatically

// Disable auto banner / 자동 배너 비활성화
logger, _ := logging.New(
    logging.WithAutoBanner(false),
)
// No automatic banner / 자동 배너 없음
```

### Banner Styles / 배너 스타일

```go
// Standard banner / 표준 배너
logger.Banner("My Application", "v1.0.0")
/* Output / 출력:
╔════════════════════════════════════════╗
║                                        ║
║       My Application v1.0.0            ║
║                                        ║
╚════════════════════════════════════════╝
*/

// Simple banner / 간단한 배너
logger.SimpleBanner("My App", "v1.0.0")
/* Output / 출력:
========================================
My App v1.0.0
========================================
*/

// Double banner with description / 설명이 있는 이중 배너
logger.DoubleBanner("Production Server", "v1.0.0", "North America")
/* Output / 출력:
╔════════════════════════════════════════╗
║       Production Server v1.0.0         ║
║       North America                    ║
╚════════════════════════════════════════╝
*/

// Custom ASCII art / 커스텀 ASCII 아트
logger.CustomBanner([]string{
    "  __  __            _             ",
    " |  \\/  |_   _     / \\   _ __  _ __",
    " | |\\/| | | | |   / _ \\ | '_ \\| '_ \\",
})

// Separator line / 구분선
logger.SeparatorLine("=", 50)
```

## Configuration Options / 설정 옵션

### File Rotation Options (Lumberjack) / 파일 로테이션 옵션

| Option / 옵션 | Description / 설명 | Default / 기본값 |
|--------------|-------------------|-----------------|
| `WithFilePath(path)` | Log file path / 로그 파일 경로 | `./logs/app.log` |
| `WithMaxSize(mb)` | Max size before rotation (MB) / 로테이션 전 최대 크기 (MB) | `100` |
| `WithMaxBackups(n)` | Max number of old files / 보관할 이전 파일 최대 개수 | `3` |
| `WithMaxAge(days)` | Max days to keep old files / 이전 파일 보관 최대 일수 | `28` |
| `WithCompress(bool)` | Compress rotated files / 로테이션된 파일 압축 | `true` |

### Logger Options / 로거 옵션

| Option / 옵션 | Description / 설명 | Default / 기본값 |
|--------------|-------------------|-----------------|
| `WithLevel(level)` | Minimum log level / 최소 로그 레벨 | `INFO` |
| `WithPrefix(string)` | Log prefix / 로그 프리픽스 | `""` |
| `WithColor(bool)` | Enable colored output / 색상 출력 활성화 | `true` |
| `WithStdout(bool)` | Enable stdout output / 표준 출력 활성화 | `true` |
| `WithFile(bool)` | Enable file output / 파일 출력 활성화 | `true` |
| `WithStdoutOnly()` | Stdout only (no file) / 표준 출력만 (파일 없음) | - |
| `WithFileOnly()` | File only (no stdout) / 파일만 (표준 출력 없음) | - |
| `WithTimeFormat(format)` | Time format / 시간 형식 | `2006-01-02 15:04:05` |

### Banner Options / 배너 옵션

| Option / 옵션 | Description / 설명 | Default / 기본값 |
|--------------|-------------------|-----------------|
| `WithAutoBanner(bool)` | Auto-print banner on creation / 생성 시 자동 배너 출력 | `true` |
| `WithAppName(string)` | Application name for banner / 배너 애플리케이션 이름 | `"Application"` |
| `WithAppVersion(string)` | Application version for banner / 배너 애플리케이션 버전 | `"v1.0.0"` |
| `WithBanner(name, version)` | Convenience: set name, version & enable auto banner / 편의 함수 | - |

## Advanced Usage / 고급 사용법

### File-Only Logging / 파일 전용 로깅

```go
logger, _ := logging.New(
    logging.WithFilePath("./logs/production.log"),
    logging.WithFileOnly(), // No console output / 콘솔 출력 없음
)
```

### Stdout-Only Logging / 표준 출력 전용 로깅

```go
logger, _ := logging.New(
    logging.WithStdoutOnly(), // No file output / 파일 출력 없음
)
```

### Manual Rotation / 수동 로테이션

```go
// Manually trigger log rotation / 수동으로 로그 로테이션 트리거
if err := logger.Rotate(); err != nil {
    log.Printf("Failed to rotate log: %v", err)
}
```

### Custom Time Format / 커스텀 시간 형식

```go
logger, _ := logging.New(
    logging.WithTimeFormat("2006/01/02 15:04:05.000"),
)
```

## Use Cases / 사용 사례

### Web Application / 웹 애플리케이션

```go
// Separate logs for different components / 컴포넌트별 로그 분리
appLogger, _ := logging.New(logging.WithFilePath("./logs/app.log"))
accessLogger, _ := logging.New(logging.WithFilePath("./logs/access.log"))
errorLogger, _ := logging.New(
    logging.WithFilePath("./logs/error.log"),
    logging.WithLevel(logging.ERROR),
)
```

### Microservices / 마이크로서비스

```go
// Service-specific logging / 서비스별 로깅
authLogger, _ := logging.New(
    logging.WithFilePath("./logs/auth-service.log"),
    logging.WithPrefix("[AUTH]"),
)

paymentLogger, _ := logging.New(
    logging.WithFilePath("./logs/payment-service.log"),
    logging.WithPrefix("[PAYMENT]"),
)
```

### Development vs Production / 개발 vs 프로덕션

```go
var logger *logging.Logger

if os.Getenv("ENV") == "production" {
    logger, _ = logging.New(
        logging.WithLevel(logging.INFO),
        logging.WithFileOnly(), // Production: file only / 프로덕션: 파일만
    )
} else {
    logger, _ = logging.New(
        logging.WithLevel(logging.DEBUG),
        // Development: console + file / 개발: 콘솔 + 파일
    )
}
```

## Examples / 예제

See the [examples directory](../examples/logging/) for complete working examples.

완전한 실행 예제는 [examples 디렉토리](../examples/logging/)를 참조하세요.

```bash
# Run the example / 예제 실행
go run examples/logging/main.go
```

## Testing / 테스트

```bash
# Run tests / 테스트 실행
go test -v

# Run tests with coverage / 커버리지와 함께 테스트 실행
go test -cover

# Run benchmarks / 벤치마크 실행
go test -bench=.
```

## Dependencies / 의존성

This package uses [lumberjack](https://github.com/natefinch/lumberjack) for log file rotation.

이 패키지는 로그 파일 로테이션을 위해 [lumberjack](https://github.com/natefinch/lumberjack)을 사용합니다.

- **lumberjack** (v2.2.1) - MIT License - Copyright (c) 2014 Nate Finch
  - Provides automatic log file rotation and compression
  - 자동 로그 파일 로테이션 및 압축 제공

## Performance / 성능

The logger is optimized for performance:
- Mutex locks for thread-safety / 스레드 안전성을 위한 Mutex 잠금
- Minimal allocations / 최소 메모리 할당
- Buffered writes via lumberjack / lumberjack을 통한 버퍼링된 쓰기

로거는 성능을 위해 최적화되었습니다.

## Best Practices / 모범 사례

1. **Always close loggers** / 항상 로거 닫기
   ```go
   logger := logging.Default()
   defer logger.Close()
   ```

2. **Use structured logging for searchability** / 검색 가능성을 위해 구조화된 로깅 사용
   ```go
   logger.Info("Event", "key1", value1, "key2", value2)
   ```

3. **Separate loggers for different concerns** / 관심사별로 로거 분리
   ```go
   appLogger := logging.New(logging.WithFilePath("./logs/app.log"))
   dbLogger := logging.New(logging.WithFilePath("./logs/db.log"))
   ```

4. **Set appropriate log levels** / 적절한 로그 레벨 설정
   - Development: DEBUG / 개발: DEBUG
   - Production: INFO or WARN / 프로덕션: INFO 또는 WARN

5. **Monitor log file sizes** / 로그 파일 크기 모니터링
   - Configure MaxSize, MaxBackups, MaxAge appropriately
   - MaxSize, MaxBackups, MaxAge를 적절히 설정

## License / 라이선스

MIT License - see the [LICENSE](../LICENSE) file for details.

MIT 라이선스 - 자세한 내용은 [LICENSE](../LICENSE) 파일을 참조하세요.

## Credits / 크레딧

- Built with [lumberjack](https://github.com/natefinch/lumberjack) by Nate Finch
- Part of the [go-utils](https://github.com/arkd0ng/go-utils) collection
