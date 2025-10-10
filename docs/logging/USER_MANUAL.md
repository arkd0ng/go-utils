# Logging Package - User Manual / 로깅 패키지 - 사용자 매뉴얼

**Version / 버전**: v1.2.004
**Last Updated / 최종 업데이트**: 2025-10-10

---

## Table of Contents / 목차

1. [Introduction / 소개](#introduction--소개)
2. [Installation / 설치](#installation--설치)
3. [Quick Start / 빠른 시작](#quick-start--빠른-시작)
4. [Configuration Reference / 설정 참조](#configuration-reference--설정-참조)
5. [Logging Methods / 로깅 메서드](#logging-methods--로깅-메서드)
6. [Banner Features / 배너 기능](#banner-features--배너-기능)
7. [Version Management / 버전 관리](#version-management--버전-관리)
8. [Usage Patterns / 사용 패턴](#usage-patterns--사용-패턴)
9. [Common Use Cases / 일반적인 사용 사례](#common-use-cases--일반적인-사용-사례)
10. [Best Practices / 모범 사례](#best-practices--모범-사례)
11. [Troubleshooting / 문제 해결](#troubleshooting--문제-해결)
12. [FAQ](#faq)

---

## Introduction / 소개

The Logging package is an enterprise-grade logging solution for Go applications with automatic file rotation, structured logging, and customizable banners.

Logging 패키지는 자동 파일 로테이션, 구조화된 로깅, 커스터마이징 가능한 배너를 갖춘 Go 애플리케이션용 엔터프라이즈급 로깅 솔루션입니다.

### Key Features / 주요 기능

- **Zero Configuration / 제로 설정**: Works out of the box with sensible defaults / 합리적인 기본값으로 즉시 사용 가능
- **Automatic File Rotation / 자동 파일 로테이션**: Built on lumberjack for automatic log file management / lumberjack 기반 자동 로그 파일 관리
- **Dual Logging Styles / 이중 로깅 스타일**: Supports both structured (key-value) and Printf-style logging / 구조화 및 Printf 스타일 모두 지원
- **Multiple Log Levels / 다중 로그 레벨**: DEBUG, INFO, WARN, ERROR, FATAL / 5가지 로그 레벨
- **Colored Console Output / 색상 콘솔 출력**: ANSI color-coded output for easy visual parsing / ANSI 색상 코딩으로 시각적 파싱 용이
- **Flexible Output / 유연한 출력**: File-only, console-only, or both / 파일만, 콘솔만, 또는 둘 다
- **Smart Banners / 스마트 배너**: Auto-extracts app name from log filename / 로그 파일명에서 앱 이름 자동 추출
- **Version Management / 버전 관리**: Auto-loads app name and version from app.yaml / app.yaml에서 앱 이름과 버전 자동 로드
- **Thread-Safe / 스레드 안전**: Safe for concurrent use / 동시 사용 안전
- **Production-Ready / 프로덕션 준비**: Default settings optimized for production / 프로덕션에 최적화된 기본 설정

### Use Cases / 사용 사례

This package is ideal for:

이 패키지는 다음과 같은 경우에 이상적입니다:

- Web applications / 웹 애플리케이션
- Microservices / 마이크로서비스
- CLI tools / CLI 도구
- Background services / 백그라운드 서비스
- Any Go application requiring robust logging / 강력한 로깅이 필요한 모든 Go 애플리케이션

---

## Installation / 설치

### Prerequisites / 전제 조건

- Go 1.16 or higher / Go 1.16 이상
- No other dependencies required (lumberjack and yaml.v3 are automatically managed) / 다른 의존성 불필요 (lumberjack 및 yaml.v3는 자동 관리)

### Install Package / 패키지 설치

```bash
go get github.com/arkd0ng/go-utils/logging
```

### Import in Your Code / 코드에 임포트

```go
import "github.com/arkd0ng/go-utils/logging"
```

---

## Quick Start / 빠른 시작

### Example 1: Default Logger / 기본 로거

The simplest way to get started:

가장 간단한 시작 방법:

```go
package main

import "github.com/arkd0ng/go-utils/logging"

func main() {
    // Create default logger
    // 기본 로거 생성
    logger := logging.Default()
    defer logger.Close()

    // Log messages
    // 메시지 로깅
    logger.Info("Application started")
    logger.Warn("This is a warning")
    logger.Error("An error occurred")
}
```

**Output / 출력**: Logs to `./logs/app.log` / `./logs/app.log`에 로그 기록

### Example 2: Custom Configuration / 커스텀 설정

Configure the logger for your needs:

필요에 맞게 로거 설정:

```go
logger, err := logging.New(
    logging.WithFilePath("./logs/myapp.log"),
    logging.WithMaxSize(50),        // 50 MB
    logging.WithMaxBackups(3),      // Keep 3 backups / 3개 백업 유지
    logging.WithLevel(logging.DEBUG),
    logging.WithStdout(true),       // Enable console output / 콘솔 출력 활성화
)
if err != nil {
    panic(err)
}
defer logger.Close()

logger.Info("Server started", "port", 8080)
```

### Example 3: Printf-Style Logging / Printf 스타일 로깅

Use familiar fmt.Printf syntax:

친숙한 fmt.Printf 문법 사용:

```go
logger := logging.Default()
defer logger.Close()

port := 8080
logger.Infof("Server listening on port %d", port)
logger.Errorf("Failed to connect to %s: %v", "database", err)
```

---

## Configuration Reference / 설정 참조

### File Rotation Options / 파일 로테이션 옵션

These options control how log files are rotated and managed:

이 옵션들은 로그 파일의 로테이션 및 관리 방법을 제어합니다:

| Option / 옵션 | Type / 타입 | Default / 기본값 | Description / 설명 |
|--------------|------------|-----------------|-------------------|
| `WithFilePath(path)` | string | `./logs/app.log` | Log file path / 로그 파일 경로 |
| `WithMaxSize(mb)` | int | `100` | Max file size (MB) before rotation / 로테이션 전 최대 파일 크기 (MB) |
| `WithMaxBackups(n)` | int | `3` | Number of old log files to keep / 보관할 이전 로그 파일 개수 |
| `WithMaxAge(days)` | int | `28` | Max days to keep old log files / 이전 로그 파일 보관 최대 일수 |
| `WithCompress(bool)` | bool | `true` | Compress rotated files (gzip) / 로테이션된 파일 압축 (gzip) |

**Example / 예제:**

```go
logger, _ := logging.New(
    logging.WithFilePath("./logs/app.log"),
    logging.WithMaxSize(100),      // 100 MB max file size
    logging.WithMaxBackups(5),     // Keep 5 old files
    logging.WithMaxAge(30),        // Keep for 30 days
    logging.WithCompress(true),    // Compress old files
)
```

### Logger Options / 로거 옵션

| Option / 옵션 | Type / 타입 | Default / 기본값 | Description / 설명 |
|--------------|------------|-----------------|-------------------|
| `WithLevel(level)` | Level | `INFO` | Minimum log level / 최소 로그 레벨 |
| `WithPrefix(string)` | string | `""` | Log message prefix / 로그 메시지 프리픽스 |
| `WithColor(bool)` | bool | `true` | Enable colored console output / 색상 콘솔 출력 활성화 |
| `WithStdout(bool)` | bool | `false` | Enable console output / 콘솔 출력 활성화 |
| `WithFile(bool)` | bool | `true` | Enable file output / 파일 출력 활성화 |
| `WithStdoutOnly()` | - | - | Console only (no file) / 콘솔만 (파일 없음) |
| `WithFileOnly()` | - | - | File only (no console) / 파일만 (콘솔 없음) |
| `WithTimeFormat(format)` | string | `2006-01-02 15:04:05` | Timestamp format / 타임스탬프 형식 |

**Example / 예제:**

```go
logger, _ := logging.New(
    logging.WithLevel(logging.DEBUG),
    logging.WithPrefix("[APP]"),
    logging.WithStdout(true),
    logging.WithColor(true),
)
```

### Banner Options / 배너 옵션

| Option / 옵션 | Type / 타입 | Default / 기본값 | Description / 설명 |
|--------------|------------|-----------------|-------------------|
| `WithAutoBanner(bool)` | bool | `true` | Auto-print banner on creation / 생성 시 자동 배너 출력 |
| `WithAppName(string)` | string | From app.yaml or `"Application"` | Application name / 애플리케이션 이름 |
| `WithAppVersion(string)` | string | From app.yaml or `"v1.0.0"` | Application version / 애플리케이션 버전 |
| `WithBanner(name, ver)` | strings | - | Set name, version & enable banner / 이름, 버전 설정 및 배너 활성화 |

**Example / 예제:**

```go
logger, _ := logging.New(
    logging.WithBanner("MyApp", "v2.0.0"),
)
// Automatically prints: MyApp v2.0.0
// 자동으로 출력: MyApp v2.0.0
```

### Log Levels / 로그 레벨

| Level / 레벨 | Constant / 상수 | Color / 색상 | Use Case / 사용 사례 |
|-------------|----------------|-------------|---------------------|
| DEBUG | `logging.DEBUG` | Cyan / 청록색 | Detailed debugging info / 상세 디버깅 정보 |
| INFO | `logging.INFO` | Green / 녹색 | General information / 일반 정보 |
| WARN | `logging.WARN` | Yellow / 노란색 | Warning messages / 경고 메시지 |
| ERROR | `logging.ERROR` | Red / 빨간색 | Error messages / 에러 메시지 |
| FATAL | `logging.FATAL` | Magenta / 자홍색 | Critical errors (exits program) / 치명적 에러 (프로그램 종료) |

---

## Logging Methods / 로깅 메서드

The logger provides two styles of logging: **Structured** and **Printf-style**.

로거는 두 가지 로깅 스타일을 제공합니다: **구조화** 및 **Printf 스타일**.

### Structured Logging (Recommended for Production) / 구조화된 로깅 (프로덕션 권장)

Structured logging uses key-value pairs for easy parsing and filtering:

구조화된 로깅은 쉬운 파싱 및 필터링을 위해 키-값 쌍을 사용합니다:

#### Methods / 메서드

```go
logger.Debug(msg string, keysAndValues ...interface{})
logger.Info(msg string, keysAndValues ...interface{})
logger.Warn(msg string, keysAndValues ...interface{})
logger.Error(msg string, keysAndValues ...interface{})
logger.Fatal(msg string, keysAndValues ...interface{})
```

#### Examples / 예제

```go
// Simple message / 간단한 메시지
logger.Info("Application started")

// With key-value pairs / 키-값 쌍 포함
logger.Info("User login",
    "username", "john.doe",
    "user_id", 12345,
    "ip", "192.168.1.100",
)

// Multiple key-value pairs / 여러 키-값 쌍
logger.Error("Database connection failed",
    "host", "db.example.com",
    "port", 5432,
    "error", err,
    "retry_count", 3,
)
```

**Output Format / 출력 형식:**

```
2025-10-10 15:30:45 [INFO] User login username=john.doe user_id=12345 ip=192.168.1.100
```

### Printf-Style Logging (Convenient for Development) / Printf 스타일 로깅 (개발 편의성)

Printf-style logging uses format strings, just like fmt.Printf:

Printf 스타일 로깅은 fmt.Printf와 같이 형식 문자열을 사용합니다:

#### Methods / 메서드

```go
logger.Debugf(format string, args ...interface{})
logger.Infof(format string, args ...interface{})
logger.Warnf(format string, args ...interface{})
logger.Errorf(format string, args ...interface{})
logger.Fatalf(format string, args ...interface{})
```

#### Examples / 예제

```go
// Simple formatted message / 간단한 형식화 메시지
logger.Infof("Server started on port %d", 8080)

// Multiple arguments / 여러 인자
logger.Infof("User %s (ID: %d) logged in from %s",
    username, userID, ipAddress)

// With error / 에러 포함
logger.Errorf("Failed to connect to %s: %v", host, err)
```

**Output Format / 출력 형식:**

```
2025-10-10 15:30:45 [INFO] Server started on port 8080
```

### Comparison / 비교

| Aspect / 측면 | Structured / 구조화 | Printf-Style / Printf 스타일 |
|--------------|-------------------|----------------------------|
| **Parsing / 파싱** | Easy to parse / 파싱 용이 | Harder to parse / 파싱 어려움 |
| **Searchability / 검색성** | Field-based search / 필드 기반 검색 | Text-based search / 텍스트 기반 검색 |
| **Readability / 가독성** | Less human-readable / 사람이 읽기 어려움 | More human-readable / 사람이 읽기 쉬움 |
| **Production Use / 프로덕션 사용** | ✅ Recommended / 권장 | ⚠️ Use with caution / 주의하여 사용 |
| **Development Use / 개발 사용** | ⚠️ More verbose / 더 장황함 | ✅ Quick and easy / 빠르고 쉬움 |

### Additional Methods / 추가 메서드

#### SetLevel / 레벨 설정

Change the minimum log level at runtime:

런타임에 최소 로그 레벨 변경:

```go
logger.SetLevel(logging.DEBUG)  // Show all logs / 모든 로그 표시
logger.SetLevel(logging.WARN)   // Only WARN and above / WARN 이상만
```

#### GetLevel / 레벨 가져오기

Get the current log level:

현재 로그 레벨 가져오기:

```go
currentLevel := logger.GetLevel()
fmt.Println(currentLevel) // Output: INFO
```

#### Rotate / 로테이션

Manually trigger log rotation:

수동으로 로그 로테이션 트리거:

```go
if err := logger.Rotate(); err != nil {
    log.Printf("Failed to rotate log: %v", err)
}
```

#### Close / 닫기

Close the logger and flush buffered data:

로거를 닫고 버퍼링된 데이터 플러시:

```go
defer logger.Close()
```

---

## Banner Features / 배너 기능

Banners provide a visual indicator of application startup and version information.

배너는 애플리케이션 시작 및 버전 정보의 시각적 표시를 제공합니다.

### Auto Banner / 자동 배너

By default, a banner is automatically printed when creating a logger:

기본적으로 로거 생성 시 배너가 자동으로 출력됩니다:

```go
logger := logging.Default()
// Automatically prints banner with app name from app.yaml
// app.yaml의 앱 이름으로 배너 자동 출력
```

### Smart App Name Extraction / 스마트 앱 이름 추출

The logger automatically extracts the app name from the log filename:

로거는 로그 파일명에서 앱 이름을 자동으로 추출합니다:

```go
// Creates banner: "database v1.0.0"
// 배너 생성: "database v1.0.0"
logger, _ := logging.New(
    logging.WithFilePath("./logs/database.log"),
)

// Creates banner: "api-server v1.0.0"
// 배너 생성: "api-server v1.0.0"
logger, _ := logging.New(
    logging.WithFilePath("./logs/api-server.log"),
)
```

### Banner Styles / 배너 스타일

#### 1. Standard Banner / 표준 배너

```go
logger.Banner("My Application", "v1.0.0")
```

**Output / 출력:**

```
╔════════════════════════════════════════════════════════════╗
║                                                            ║
║                 My Application v1.0.0                      ║
║                                                            ║
╚════════════════════════════════════════════════════════════╝
```

#### 2. Simple Banner / 간단한 배너

```go
logger.SimpleBanner("My App", "v1.0.0")
```

**Output / 출력:**

```
============================================================
My App v1.0.0
============================================================
```

#### 3. Double Banner / 이중 배너

With description:

설명 포함:

```go
logger.DoubleBanner("Production Server", "v1.0.0", "North America")
```

**Output / 출력:**

```
╔════════════════════════════════════════════════════════════╗
║                 Production Server v1.0.0                   ║
║                      North America                         ║
╚════════════════════════════════════════════════════════════╝
```

#### 4. Custom ASCII Art Banner / 커스텀 ASCII 아트 배너

```go
logger.CustomBanner([]string{
    "  __  __            _             ",
    " |  \\/  |_   _     / \\   _ __  _ __",
    " | |\\/| | | | |   / _ \\ | '_ \\| '_ \\",
    " | |  | | |_| |  / ___ \\| |_) | |_) |",
    " |_|  |_|\\__, | /_/   \\_\\ .__/| .__/",
    "         |___/          |_|   |_|",
})
```

#### 5. Separator Line / 구분선

```go
logger.SeparatorLine("=", 60)
logger.SeparatorLine("-", 60)
```

### Disable Auto Banner / 자동 배너 비활성화

```go
logger, _ := logging.New(
    logging.WithAutoBanner(false),  // No automatic banner
)
```

---

## Version Management / 버전 관리

The logger automatically loads app name and version from `app.yaml` if it exists.

로거는 `app.yaml` 파일이 존재하면 자동으로 앱 이름과 버전을 로드합니다.

### Create app.yaml / app.yaml 생성

Place the file in one of these locations:

다음 위치 중 하나에 파일을 배치하세요:

- `cfg/app.yaml` (recommended / 권장)
- `apps/app.yaml`
- `app.yaml` (root directory / 루트 디렉토리)

**File Format / 파일 형식:**

```yaml
# cfg/app.yaml
app:
  name: go-utils
  version: v1.2.004
  description: A collection of frequently used utility functions
```

### Automatic Loading / 자동 로딩

The logger searches for app.yaml in the following order:

로거는 다음 순서로 app.yaml을 검색합니다:

1. `cfg/app.yaml` (current directory / 현재 디렉토리)
2. `apps/app.yaml`
3. `app.yaml`
4. `../cfg/app.yaml` (parent directory / 상위 디렉토리)
5. `../apps/app.yaml`
6. `../app.yaml`
7. And so on... / 계속...

**Example / 예제:**

```go
// Automatically loads from cfg/app.yaml
// cfg/app.yaml에서 자동으로 로드
logger := logging.Default()
// Banner displays: "go-utils v1.2.004"
// 배너 표시: "go-utils v1.2.004"
```

### Override with Options / 옵션으로 재정의

You can override app.yaml values:

app.yaml 값을 재정의할 수 있습니다:

```go
logger, _ := logging.New(
    logging.WithAppName("CustomName"),     // Overrides app.yaml
    logging.WithAppVersion("v2.0.0"),      // Overrides app.yaml
)
```

---

## Usage Patterns / 사용 패턴

### Pattern 1: Default Logger (Quick Start) / 기본 로거 (빠른 시작)

Fastest way to get logging:

가장 빠른 로깅 방법:

```go
logger := logging.Default()
defer logger.Close()

logger.Info("Application started")
```

**When to use / 사용 시기:**
- Quick prototyping / 빠른 프로토타이핑
- Small applications / 소형 애플리케이션
- Getting started / 시작 단계

### Pattern 2: Production Logger / 프로덕션 로거

Optimized for production:

프로덕션에 최적화:

```go
logger, err := logging.New(
    logging.WithFilePath("./logs/production.log"),
    logging.WithLevel(logging.INFO),       // INFO and above
    logging.WithMaxSize(100),              // 100 MB
    logging.WithMaxBackups(10),            // Keep 10 backups
    logging.WithMaxAge(90),                // 90 days retention
    logging.WithCompress(true),            // Compress old logs
    logging.WithFileOnly(),                // No console output
)
if err != nil {
    log.Fatalf("Failed to create logger: %v", err)
}
defer logger.Close()
```

**When to use / 사용 시기:**
- Production environments / 프로덕션 환경
- High-volume logging / 대용량 로깅
- Long-term log retention / 장기 로그 보관

### Pattern 3: Development Logger / 개발 로거

Optimized for development:

개발에 최적화:

```go
logger, err := logging.New(
    logging.WithLevel(logging.DEBUG),      // All logs
    logging.WithStdout(true),              // Console output
    logging.WithColor(true),               // Colored output
    logging.WithFile(true),                // Also save to file
)
if err != nil {
    log.Fatalf("Failed to create logger: %v", err)
}
defer logger.Close()
```

**When to use / 사용 시기:**
- Local development / 로컬 개발
- Debugging / 디버깅
- Testing / 테스트

### Pattern 4: Multiple Loggers / 여러 로거

Separate concerns with multiple loggers:

여러 로거로 관심사 분리:

```go
// Application logger
// 애플리케이션 로거
appLogger, _ := logging.New(
    logging.WithFilePath("./logs/app.log"),
    logging.WithPrefix("[APP]"),
)

// Database logger
// 데이터베이스 로거
dbLogger, _ := logging.New(
    logging.WithFilePath("./logs/database.log"),
    logging.WithPrefix("[DB]"),
)

// API logger
// API 로거
apiLogger, _ := logging.New(
    logging.WithFilePath("./logs/api.log"),
    logging.WithPrefix("[API]"),
)

appLogger.Info("Application started")
dbLogger.Info("Database connected")
apiLogger.Info("API server listening on :8080")
```

**When to use / 사용 시기:**
- Microservices / 마이크로서비스
- Complex applications / 복잡한 애플리케이션
- Component separation / 컴포넌트 분리

### Pattern 5: Console-Only Logger / 콘솔 전용 로거

No file output:

파일 출력 없음:

```go
logger, _ := logging.New(
    logging.WithStdoutOnly(),
    logging.WithColor(true),
)

logger.Info("This only goes to console")
```

**When to use / 사용 시기:**
- CLI tools / CLI 도구
- Docker containers (with log forwarding) / Docker 컨테이너 (로그 포워딩 사용)
- Temporary debugging / 임시 디버깅

---

## Common Use Cases / 일반적인 사용 사례

### Use Case 1: Web Application Logging / 웹 애플리케이션 로깅

Separate logs for different components:

컴포넌트별로 로그 분리:

```go
package main

import "github.com/arkd0ng/go-utils/logging"

type Server struct {
    appLogger    *logging.Logger
    accessLogger *logging.Logger
    errorLogger  *logging.Logger
}

func NewServer() *Server {
    appLogger, _ := logging.New(
        logging.WithFilePath("./logs/app.log"),
        logging.WithBanner("WebApp", "v1.0.0"),
    )

    accessLogger, _ := logging.New(
        logging.WithFilePath("./logs/access.log"),
        logging.WithAutoBanner(false),  // No banner for access log
    )

    errorLogger, _ := logging.New(
        logging.WithFilePath("./logs/error.log"),
        logging.WithLevel(logging.ERROR),  // Only errors
        logging.WithAutoBanner(false),
    )

    return &Server{
        appLogger:    appLogger,
        accessLogger: accessLogger,
        errorLogger:  errorLogger,
    }
}

func (s *Server) Start() {
    s.appLogger.Info("Server starting")
    s.accessLogger.Info("GET /api/users", "status", 200, "duration_ms", 45)
    s.errorLogger.Error("Database connection failed", "error", err)
}
```

### Use Case 2: Microservice Logging / 마이크로서비스 로깅

Service-specific logging with structured data:

구조화된 데이터를 사용한 서비스별 로깅:

```go
package main

import "github.com/arkd0ng/go-utils/logging"

type AuthService struct {
    logger *logging.Logger
}

func NewAuthService() *AuthService {
    logger, _ := logging.New(
        logging.WithFilePath("./logs/auth-service.log"),
        logging.WithPrefix("[AUTH]"),
        logging.WithBanner("AuthService", "v2.1.0"),
    )

    return &AuthService{logger: logger}
}

func (s *AuthService) Login(username, ip string) error {
    s.logger.Info("User login attempt",
        "username", username,
        "ip", ip,
        "timestamp", time.Now(),
    )

    // Login logic...

    s.logger.Info("User login successful",
        "username", username,
        "session_id", sessionID,
    )

    return nil
}
```

### Use Case 3: CLI Tool Logging / CLI 도구 로깅

Console-only colored output:

콘솔 전용 색상 출력:

```go
package main

import "github.com/arkd0ng/go-utils/logging"

func main() {
    logger, _ := logging.New(
        logging.WithStdoutOnly(),
        logging.WithColor(true),
        logging.WithBanner("MyTool", "v1.0.0"),
    )
    defer logger.Close()

    logger.Info("Processing files...")
    logger.Infof("Processed %d files", count)
    logger.Warn("Some files were skipped")
}
```

### Use Case 4: Background Service / 백그라운드 서비스

Long-running service with log rotation:

로그 로테이션을 사용하는 장기 실행 서비스:

```go
package main

import "github.com/arkd0ng/go-utils/logging"

func main() {
    logger, _ := logging.New(
        logging.WithFilePath("./logs/service.log"),
        logging.WithMaxSize(50),       // 50 MB
        logging.WithMaxBackups(30),    // Keep 30 files
        logging.WithMaxAge(365),       // Keep for 1 year
        logging.WithCompress(true),
        logging.WithFileOnly(),
    )
    defer logger.Close()

    logger.Info("Service started")

    // Long-running process
    for {
        logger.Debug("Processing batch", "batch_id", batchID)
        // Process...
        time.Sleep(time.Minute)
    }
}
```

### Use Case 5: Environment-Specific Logging / 환경별 로깅

Different settings for dev/prod:

개발/프로덕션별 다른 설정:

```go
package main

import (
    "os"
    "github.com/arkd0ng/go-utils/logging"
)

func createLogger() *logging.Logger {
    env := os.Getenv("ENV")

    if env == "production" {
        logger, _ := logging.New(
            logging.WithLevel(logging.INFO),
            logging.WithFileOnly(),
            logging.WithCompress(true),
        )
        return logger
    } else {
        logger, _ := logging.New(
            logging.WithLevel(logging.DEBUG),
            logging.WithStdout(true),
            logging.WithColor(true),
        )
        return logger
    }
}

func main() {
    logger := createLogger()
    defer logger.Close()

    logger.Info("Application started")
}
```

---

## Best Practices / 모범 사례

### 1. Always Close Loggers / 항상 로거 닫기

Use `defer` to ensure the logger is properly closed:

`defer`를 사용하여 로거가 올바르게 닫히도록 보장:

```go
logger := logging.Default()
defer logger.Close()  // ✅ Good / 좋음
```

### 2. Use Structured Logging in Production / 프로덕션에서 구조화된 로깅 사용

Structured logging is easier to parse and analyze:

구조화된 로깅은 파싱 및 분석이 쉽습니다:

```go
// ✅ Good - Structured / 좋음 - 구조화
logger.Info("User login",
    "username", username,
    "user_id", userID,
)

// ❌ Avoid in production - Printf style / 프로덕션에서 피할 것 - Printf 스타일
logger.Infof("User %s (ID: %d) logged in", username, userID)
```

### 3. Separate Loggers for Different Concerns / 관심사별로 로거 분리

Create separate loggers for different components:

컴포넌트별로 로거를 분리 생성:

```go
appLogger := logging.New(logging.WithFilePath("./logs/app.log"))
dbLogger := logging.New(logging.WithFilePath("./logs/db.log"))
apiLogger := logging.New(logging.WithFilePath("./logs/api.log"))
```

### 4. Set Appropriate Log Levels / 적절한 로그 레벨 설정

Choose log levels based on environment:

환경에 따라 로그 레벨 선택:

- **Development / 개발**: `DEBUG` - See everything / 모든 것 확인
- **Staging / 스테이징**: `INFO` - General information / 일반 정보
- **Production / 프로덕션**: `WARN` or `INFO` - Only important events / 중요한 이벤트만

```go
// Development / 개발
logger, _ := logging.New(logging.WithLevel(logging.DEBUG))

// Production / 프로덕션
logger, _ := logging.New(logging.WithLevel(logging.INFO))
```

### 5. Configure Log Rotation / 로그 로테이션 설정

Set appropriate rotation settings for your application:

애플리케이션에 맞는 로테이션 설정:

```go
logger, _ := logging.New(
    logging.WithMaxSize(100),      // Adjust based on write volume
    logging.WithMaxBackups(10),    // Balance disk space vs history
    logging.WithMaxAge(30),        // Comply with retention policies
    logging.WithCompress(true),    // Save disk space
)
```

### 6. Use Prefixes for Clarity / 명확성을 위해 프리픽스 사용

Add prefixes to distinguish log sources:

로그 소스를 구분하기 위해 프리픽스 추가:

```go
logger, _ := logging.New(logging.WithPrefix("[API]"))
logger.Info("Request received")
// Output: 2025-10-10 15:30:45 [INFO] [API] Request received
```

### 7. Monitor Disk Usage / 디스크 사용량 모니터링

Set reasonable `MaxSize`, `MaxBackups`, and `MaxAge` to prevent disk exhaustion:

디스크 소진 방지를 위해 합리적인 `MaxSize`, `MaxBackups`, `MaxAge` 설정:

```go
// Example: 100 MB × 10 backups = max 1 GB disk usage
// 예: 100 MB × 10 백업 = 최대 1 GB 디스크 사용량
logger, _ := logging.New(
    logging.WithMaxSize(100),
    logging.WithMaxBackups(10),
)
```

### 8. Avoid Logging Sensitive Data / 민감한 데이터 로깅 피하기

Never log passwords, API keys, or personal information:

비밀번호, API 키, 개인정보를 로깅하지 마세요:

```go
// ❌ Bad - Logs password / 나쁨 - 비밀번호 로깅
logger.Info("User login", "username", user, "password", pass)

// ✅ Good - No sensitive data / 좋음 - 민감한 데이터 없음
logger.Info("User login", "username", user)
```

### 9. Use app.yaml for Version Management / 버전 관리에 app.yaml 사용

Centralize version management:

버전 관리 중앙화:

```yaml
# cfg/app.yaml
app:
  name: myapp
  version: v1.2.004
```

```go
// Automatically loads version from app.yaml
// app.yaml에서 버전 자동 로드
logger := logging.Default()
```

### 10. Test Your Logging Configuration / 로깅 설정 테스트

Verify log rotation and output:

로그 로테이션 및 출력 확인:

```go
logger, _ := logging.New(
    logging.WithFilePath("./logs/test.log"),
    logging.WithStdout(true),  // See output during testing
)

logger.Info("Test message")

// Check that file was created
// 파일이 생성되었는지 확인
if _, err := os.Stat("./logs/test.log"); err != nil {
    panic("Log file not created")
}
```

---

## Troubleshooting / 문제 해결

### Problem 1: Logs Not Appearing / 로그가 표시되지 않음

**Symptom / 증상**: No log output

**Possible Causes / 가능한 원인:**

1. **Log level too high / 로그 레벨이 너무 높음**
   ```go
   // If level is WARN, DEBUG and INFO are hidden
   // 레벨이 WARN이면 DEBUG와 INFO가 숨겨짐
   logger, _ := logging.New(logging.WithLevel(logging.WARN))
   logger.Info("This won't appear")  // ❌
   ```

   **Solution / 해결책:**
   ```go
   logger.SetLevel(logging.DEBUG)  // Show all logs
   ```

2. **File output only (no console) / 파일 출력만 (콘솔 없음)**
   ```go
   // Default is file-only
   // 기본값은 파일 전용
   logger := logging.Default()
   ```

   **Solution / 해결책:**
   ```go
   logger, _ := logging.New(logging.WithStdout(true))
   ```

### Problem 2: File Permission Errors / 파일 권한 오류

**Symptom / 증상**: `permission denied` error

**Solution / 해결책:**

1. Check directory permissions:
   ```bash
   chmod 755 ./logs/
   ```

2. Ensure user has write access:
   ```bash
   ls -la ./logs/
   ```

### Problem 3: Disk Space Full / 디스크 공간 가득 참

**Symptom / 증상**: Application crashes or logs stop

**Solution / 해결책:**

1. Reduce `MaxSize`:
   ```go
   logger, _ := logging.New(logging.WithMaxSize(50))  // 50 MB instead of 100
   ```

2. Reduce `MaxBackups`:
   ```go
   logger, _ := logging.New(logging.WithMaxBackups(5))  // Keep fewer backups
   ```

3. Enable compression:
   ```go
   logger, _ := logging.New(logging.WithCompress(true))
   ```

### Problem 4: Logs Not Rotating / 로그가 로테이션되지 않음

**Symptom / 증상**: Single log file grows indefinitely

**Solution / 해결책:**

Verify rotation settings are configured:
```go
logger, _ := logging.New(
    logging.WithMaxSize(100),      // ✅ Set max size
    logging.WithMaxBackups(3),     // ✅ Set backups
    logging.WithMaxAge(28),        // ✅ Set age
)
```

### Problem 5: Missing Banner / 배너 누락

**Symptom / 증상**: No banner on startup

**Possible Causes / 가능한 원인:**

1. **Auto banner disabled / 자동 배너 비활성화**
   ```go
   logger, _ := logging.New(logging.WithAutoBanner(false))
   ```

2. **app.yaml not found / app.yaml을 찾을 수 없음**

   **Solution / 해결책:**
   ```go
   // Manually specify app name and version
   // 수동으로 앱 이름과 버전 지정
   logger, _ := logging.New(
       logging.WithBanner("MyApp", "v1.0.0"),
   )
   ```

### Problem 6: Color Codes in Log Files / 로그 파일에 색상 코드

**Symptom / 증상**: ANSI color codes (`\033[32m`) appear in log files

**Cause / 원인**: Color is enabled but shouldn't be in files

**Solution / 해결책:**

Color codes are only applied to console output, not files. If you see them in files, check your configuration:

```go
// This is correct - color only applies to console
// 올바름 - 색상은 콘솔에만 적용됨
logger, _ := logging.New(
    logging.WithColor(true),
    logging.WithStdout(true),
)
```

---

## FAQ

### Q1: Can I use this logger with other logging libraries? / 다른 로깅 라이브러리와 함께 사용할 수 있나요?

**A**: Yes, this logger is independent and can coexist with other logging solutions. However, using multiple loggers for the same purpose may cause confusion.

**답**: 네, 이 로거는 독립적이며 다른 로깅 솔루션과 공존할 수 있습니다. 하지만 동일한 목적으로 여러 로거를 사용하면 혼란을 야기할 수 있습니다.

### Q2: How do I log to both console and file? / 콘솔과 파일 모두에 로그하려면?

**A**: Enable both outputs:

**답**: 두 출력을 모두 활성화하세요:

```go
logger, _ := logging.New(
    logging.WithStdout(true),  // Console
    logging.WithFile(true),    // File
)
```

### Q3: Can I change log level at runtime? / 런타임에 로그 레벨을 변경할 수 있나요?

**A**: Yes, use `SetLevel()`:

**답**: 네, `SetLevel()`을 사용하세요:

```go
logger.SetLevel(logging.DEBUG)
```

### Q4: How do I disable banners? / 배너를 비활성화하려면?

**A**: Use `WithAutoBanner(false)`:

**답**: `WithAutoBanner(false)`를 사용하세요:

```go
logger, _ := logging.New(logging.WithAutoBanner(false))
```

### Q5: What's the difference between structured and Printf-style logging? / 구조화 로깅과 Printf 스타일 로깅의 차이점은?

**A**:
- **Structured**: Uses key-value pairs, easier to parse and search. Recommended for production.
- **Printf**: Uses format strings, more human-readable. Good for development.

**답**:
- **구조화**: 키-값 쌍 사용, 파싱 및 검색 용이. 프로덕션 권장.
- **Printf**: 형식 문자열 사용, 사람이 읽기 쉬움. 개발에 적합.

### Q6: How do I create multiple loggers? / 여러 로거를 생성하려면?

**A**: Call `New()` multiple times with different configurations:

**답**: 다른 설정으로 `New()`를 여러 번 호출하세요:

```go
appLogger, _ := logging.New(logging.WithFilePath("./logs/app.log"))
dbLogger, _ := logging.New(logging.WithFilePath("./logs/db.log"))
```

### Q7: Is this logger thread-safe? / 이 로거는 스레드 안전한가요?

**A**: Yes, all operations are protected by mutex locks for safe concurrent use.

**답**: 네, 모든 작업이 mutex 잠금으로 보호되어 안전한 동시 사용이 가능합니다.

### Q8: How do I rotate logs manually? / 수동으로 로그를 로테이션하려면?

**A**: Use the `Rotate()` method:

**답**: `Rotate()` 메서드를 사용하세요:

```go
if err := logger.Rotate(); err != nil {
    log.Printf("Rotation failed: %v", err)
}
```

### Q9: Can I use custom time formats? / 커스텀 시간 형식을 사용할 수 있나요?

**A**: Yes, use `WithTimeFormat()`:

**답**: 네, `WithTimeFormat()`을 사용하세요:

```go
logger, _ := logging.New(
    logging.WithTimeFormat("2006/01/02 15:04:05.000"),
)
```

### Q10: Where should I place app.yaml? / app.yaml을 어디에 배치해야 하나요?

**A**: Recommended location is `cfg/app.yaml`. The logger searches multiple paths automatically.

**답**: 권장 위치는 `cfg/app.yaml`입니다. 로거가 자동으로 여러 경로를 검색합니다.

---

**End of User Manual / 사용자 매뉴얼 끝**

For developer documentation, see [DEVELOPER_GUIDE.md](./DEVELOPER_GUIDE.md).

개발자 문서는 [DEVELOPER_GUIDE.md](./DEVELOPER_GUIDE.md)를 참조하세요.
