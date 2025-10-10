# Logging Package - Developer Guide / 로깅 패키지 - 개발자 가이드

**Version / 버전**: v1.2.004
**Last Updated / 최종 업데이트**: 2025-10-10

---

## Table of Contents / 목차

1. [Architecture Overview / 아키텍처 개요](#architecture-overview--아키텍처-개요)
2. [Package Structure / 패키지 구조](#package-structure--패키지-구조)
3. [Core Components / 핵심 컴포넌트](#core-components--핵심-컴포넌트)
4. [Internal Implementation / 내부 구현](#internal-implementation--내부-구현)
5. [Options Pattern / 옵션 패턴](#options-pattern--옵션-패턴)
6. [Log Levels and Colors / 로그 레벨 및 색상](#log-levels-and-colors--로그-레벨-및-색상)
7. [Banner System / 배너 시스템](#banner-system--배너-시스템)
8. [File Rotation with Lumberjack / Lumberjack을 사용한 파일 로테이션](#file-rotation-with-lumberjack--lumberjack을-사용한-파일-로테이션)
9. [Version Management / 버전 관리](#version-management--버전-관리)
10. [Testing Guide / 테스트 가이드](#testing-guide--테스트-가이드)
11. [Contributing Guidelines / 기여 가이드라인](#contributing-guidelines--기여-가이드라인)
12. [Code Style / 코드 스타일](#code-style--코드-스타일)

---

## Architecture Overview / 아키텍처 개요

The Logging package is designed with simplicity, performance, and flexibility in mind. It follows the **Options pattern** for configuration and uses **Lumberjack** for automatic log rotation.

Logging 패키지는 단순성, 성능, 유연성을 염두에 두고 설계되었습니다. 설정에 **Options 패턴**을 따르며 자동 로그 로테이션에 **Lumberjack**을 사용합니다.

### Design Principles / 설계 원칙

1. **Simplicity / 단순성**: Zero-configuration defaults, easy to use / 제로 설정 기본값, 사용 용이
2. **Performance / 성능**: Minimal allocations, mutex-protected concurrent access / 최소 할당, mutex 보호 동시 접근
3. **Flexibility / 유연성**: Options pattern allows extensive customization / Options 패턴으로 광범위한 커스터마이징 가능
4. **Production-Ready / 프로덕션 준비**: Built-in rotation, compression, and version management / 내장 로테이션, 압축, 버전 관리

### High-Level Architecture / 상위 수준 아키텍처

```
┌─────────────────────────────────────────────────────────┐
│                    User Application                      │
└─────────────────────┬───────────────────────────────────┘
                      │
                      ▼
┌─────────────────────────────────────────────────────────┐
│                 Logger (Main Interface)                  │
│  ┌───────────────────────────────────────────────────┐  │
│  │  • Info(), Debug(), Warn(), Error(), Fatal()     │  │
│  │  • Infof(), Debugf(), Warnf(), Errorf(), Fatalf()│  │
│  │  • Banner(), SimpleBanner(), CustomBanner()      │  │
│  │  • SetLevel(), GetLevel(), Rotate(), Close()     │  │
│  └───────────────────────────────────────────────────┘  │
└─────────────────────┬───────────────────────────────────┘
                      │
        ┌─────────────┴─────────────┐
        │                           │
        ▼                           ▼
┌───────────────┐          ┌────────────────┐
│  File Writer  │          │ Stdout Writer  │
│  (Lumberjack) │          │   (os.Stdout)  │
└───────┬───────┘          └────────────────┘
        │
        │ Automatic Rotation
        │ • Size-based
        │ • Time-based
        │ • Compression
        ▼
┌───────────────┐
│  Log Files    │
│  • app.log    │
│  • app-*.log  │
│  • *.log.gz   │
└───────────────┘
```

### Component Interaction Flow / 컴포넌트 상호작용 흐름

```
1. User calls logger.Info("message", "key", "value")
   사용자가 logger.Info("message", "key", "value") 호출

2. Logger.log() processes the message
   Logger.log()가 메시지 처리
   • Checks log level / 로그 레벨 확인
   • Formats timestamp / 타임스탬프 형식화
   • Builds log message / 로그 메시지 작성

3. Mutex locks for thread safety
   스레드 안전성을 위한 Mutex 잠금

4. Writes to enabled outputs
   활성화된 출력에 작성
   • Stdout (with color if enabled) / Stdout (색상 활성화 시)
   • File (via Lumberjack) / 파일 (Lumberjack 사용)

5. Lumberjack handles rotation if needed
   필요 시 Lumberjack이 로테이션 처리
   • Checks file size / 파일 크기 확인
   • Rotates if necessary / 필요 시 로테이션
   • Compresses old files / 이전 파일 압축
```

---

## Package Structure / 패키지 구조

### File Organization / 파일 구성

```
logging/
├── logger.go       # Main logger implementation / 메인 로거 구현
├── level.go        # Log level definitions / 로그 레벨 정의
├── options.go      # Configuration options / 설정 옵션
├── banner.go       # Banner printing functions / 배너 출력 함수
├── appconfig.go    # app.yaml loading / app.yaml 로딩
├── logger_test.go  # Comprehensive tests / 종합 테스트
└── README.md       # Package documentation / 패키지 문서
```

### File Responsibilities / 파일 책임

| File / 파일 | Purpose / 목적 | Key Components / 주요 컴포넌트 |
|------------|---------------|------------------------------|
| `logger.go` | Core logging functionality / 핵심 로깅 기능 | `Logger`, `New()`, `Default()`, logging methods |
| `level.go` | Log level management / 로그 레벨 관리 | `Level`, `DEBUG`, `INFO`, `WARN`, `ERROR`, `FATAL` |
| `options.go` | Configuration options / 설정 옵션 | `Option`, `config`, `WithXXX()` functions |
| `banner.go` | Banner printing / 배너 출력 | `Banner()`, `SimpleBanner()`, `CustomBanner()` |
| `appconfig.go` | app.yaml integration / app.yaml 통합 | `LoadAppConfig()`, `TryLoadAppVersion()` |
| `logger_test.go` | Unit and integration tests / 단위 및 통합 테스트 | Test functions, benchmarks |

---

## Core Components / 핵심 컴포넌트

### 1. Logger Struct / Logger 구조체

The main logger type:

메인 로거 타입:

```go
type Logger struct {
    config       *config                // Configuration / 설정
    fileWriter   *lumberjack.Logger     // File rotation writer / 파일 로테이션 writer
    stdoutWriter io.Writer              // Console writer / 콘솔 writer
    mu           sync.Mutex             // Thread safety / 스레드 안전성
}
```

**Fields / 필드:**

- `config`: Holds all configuration options / 모든 설정 옵션 보유
- `fileWriter`: Lumberjack logger for automatic rotation / 자동 로테이션을 위한 Lumberjack 로거
- `stdoutWriter`: Standard output writer (os.Stdout) / 표준 출력 writer (os.Stdout)
- `mu`: Mutex for thread-safe operations / 스레드 안전 작업을 위한 Mutex

### 2. Config Struct / Config 구조체

Configuration holder:

설정 보유자:

```go
type config struct {
    // File rotation settings / 파일 로테이션 설정
    filename   string
    maxSize    int
    maxBackups int
    maxAge     int
    compress   bool

    // Logger settings / 로거 설정
    level        Level
    prefix       string
    enableColor  bool
    enableStdout bool
    enableFile   bool
    timeFormat   string

    // Banner settings / 배너 설정
    autoBanner bool
    appName    string
    appVersion string
}
```

### 3. Level Type / Level 타입

Log level enumeration:

로그 레벨 열거형:

```go
type Level int

const (
    DEBUG Level = iota  // 0
    INFO                // 1
    WARN                // 2
    ERROR               // 3
    FATAL               // 4
)
```

**Methods / 메서드:**

- `String() string`: Returns level name (e.g., "INFO") / 레벨 이름 반환 (예: "INFO")
- `ColorCode() string`: Returns ANSI color code / ANSI 색상 코드 반환
- `ParseLevel(s string) Level`: Parses string to Level / 문자열을 Level로 파싱

---

## Internal Implementation / 내부 구현

### Logger Creation Flow / 로거 생성 흐름

```go
// Step 1: Create default configuration
// 1단계: 기본 설정 생성
cfg := defaultConfig()

// Step 2: Apply user options
// 2단계: 사용자 옵션 적용
for _, opt := range opts {
    opt(cfg)  // Each option modifies config / 각 옵션이 config 수정
}

// Step 3: Create logger instance
// 3단계: 로거 인스턴스 생성
logger := &Logger{
    config:       cfg,
    stdoutWriter: os.Stdout,
}

// Step 4: Initialize file writer (if enabled)
// 4단계: 파일 writer 초기화 (활성화된 경우)
if cfg.enableFile {
    logger.fileWriter = &lumberjack.Logger{...}
}

// Step 5: Print auto banner (if enabled)
// 5단계: 자동 배너 출력 (활성화된 경우)
if cfg.autoBanner {
    logger.Banner(cfg.appName, cfg.appVersion)
}
```

### Logging Flow / 로깅 흐름

```go
func (l *Logger) log(level Level, msg string, keysAndValues ...interface{}) {
    // Step 1: Check log level
    // 1단계: 로그 레벨 확인
    if level < l.config.level {
        return  // Skip if below minimum / 최소값보다 낮으면 건너뜀
    }

    // Step 2: Acquire mutex lock
    // 2단계: Mutex 잠금 획득
    l.mu.Lock()
    defer l.mu.Unlock()

    // Step 3: Format timestamp
    // 3단계: 타임스탬프 형식화
    timestamp := time.Now().Format(l.config.timeFormat)

    // Step 4: Build log message
    // 4단계: 로그 메시지 작성
    logMsg := fmt.Sprintf("%s [%s] %s%s",
        timestamp,
        level.String(),
        l.config.prefix,
        msg,
    )

    // Step 5: Add key-value pairs
    // 5단계: 키-값 쌍 추가
    for i := 0; i < len(keysAndValues); i += 2 {
        logMsg += fmt.Sprintf("%v=%v ", keysAndValues[i], keysAndValues[i+1])
    }

    // Step 6: Write to stdout (with color)
    // 6단계: stdout에 작성 (색상 포함)
    if l.config.enableStdout {
        colorMsg := level.ColorCode() + logMsg + ResetColor()
        l.stdoutWriter.Write([]byte(colorMsg))
    }

    // Step 7: Write to file (without color)
    // 7단계: 파일에 작성 (색상 없음)
    if l.config.enableFile && l.fileWriter != nil {
        l.fileWriter.Write([]byte(logMsg))
    }
}
```

### Printf-Style Logging / Printf 스타일 로깅

```go
func (l *Logger) logf(level Level, format string, args ...interface{}) {
    // Step 1: Check log level
    // 1단계: 로그 레벨 확인
    if level < l.config.level {
        return
    }

    // Step 2: Format the message
    // 2단계: 메시지 형식화
    msg := fmt.Sprintf(format, args...)

    // Step 3: Delegate to structured logging
    // 3단계: 구조화된 로깅에 위임
    l.log(level, msg)
}
```

---

## Options Pattern / 옵션 패턴

The package uses the **Functional Options Pattern** for flexible configuration.

패키지는 유연한 설정을 위해 **함수형 옵션 패턴**을 사용합니다.

### Pattern Definition / 패턴 정의

```go
// Option is a function that modifies config
// Option은 config를 수정하는 함수
type Option func(*config) error
```

### Creating New Options / 새 옵션 생성

**Template / 템플릿:**

```go
// WithXXX sets the XXX configuration
// WithXXX는 XXX 설정을 지정합니다
//
// Parameters / 매개변수:
//   - value: the value to set / 설정할 값
//
// Example / 예제:
//
//  logger, _ := logging.New(logging.WithXXX(value))
func WithXXX(value Type) Option {
    return func(c *config) error {
        // Validation (optional) / 검증 (선택 사항)
        if value < 0 {
            return fmt.Errorf("invalid value: %v", value)
        }

        // Set the config field / config 필드 설정
        c.fieldName = value
        return nil
    }
}
```

**Example: Adding a New Option / 예제: 새 옵션 추가**

Let's add a `WithBufferSize` option:

`WithBufferSize` 옵션을 추가해봅시다:

```go
// Step 1: Add field to config / 1단계: config에 필드 추가
type config struct {
    // ... existing fields
    bufferSize int  // New field / 새 필드
}

// Step 2: Update defaultConfig / 2단계: defaultConfig 업데이트
func defaultConfig() *config {
    return &config{
        // ... existing defaults
        bufferSize: 4096,  // Default 4KB buffer / 기본 4KB 버퍼
    }
}

// Step 3: Create the option function / 3단계: 옵션 함수 생성
func WithBufferSize(size int) Option {
    return func(c *config) error {
        if size <= 0 {
            return fmt.Errorf("buffer size must be positive")
        }
        c.bufferSize = size
        return nil
    }
}

// Step 4: Use the buffer in logger / 4단계: 로거에서 버퍼 사용
// Implement buffering logic in logger.go
// logger.go에 버퍼링 로직 구현
```

### Option Validation / 옵션 검증

Always validate option values:

항상 옵션 값을 검증하세요:

```go
func WithMaxSize(mb int) Option {
    return func(c *config) error {
        if mb <= 0 {
            return fmt.Errorf("maxSize must be positive, got %d", mb)
        }
        if mb > 1000 {
            return fmt.Errorf("maxSize too large (max 1000 MB), got %d", mb)
        }
        c.maxSize = mb
        return nil
    }
}
```

---

## Log Levels and Colors / 로그 레벨 및 색상

### Level Implementation / 레벨 구현

```go
type Level int

const (
    DEBUG Level = iota
    INFO
    WARN
    ERROR
    FATAL
)
```

**Why iota? / 왜 iota를 사용하나요?**

- Automatic sequential numbering / 자동 순차 번호 매기기
- Easy to compare (DEBUG < INFO < WARN...) / 비교 용이 (DEBUG < INFO < WARN...)
- Type-safe / 타입 안전

### Color System / 색상 시스템

ANSI color codes for terminal output:

터미널 출력을 위한 ANSI 색상 코드:

```go
func (l Level) ColorCode() string {
    switch l {
    case DEBUG:
        return "\033[36m"  // Cyan / 청록색
    case INFO:
        return "\033[32m"  // Green / 녹색
    case WARN:
        return "\033[33m"  // Yellow / 노란색
    case ERROR:
        return "\033[31m"  // Red / 빨간색
    case FATAL:
        return "\033[35m"  // Magenta / 자홍색
    default:
        return "\033[0m"   // Reset / 재설정
    }
}
```

**ANSI Code Format / ANSI 코드 형식:**

- `\033[`: Escape sequence start / 이스케이프 시퀀스 시작
- `XXm`: Color code / 색상 코드
  - `31` = Red / 빨간색
  - `32` = Green / 녹색
  - `33` = Yellow / 노란색
  - `36` = Cyan / 청록색
  - `35` = Magenta / 자홍색
  - `0` = Reset / 재설정

**Color Application / 색상 적용:**

```go
// Only applied to console output / 콘솔 출력에만 적용
if l.config.enableStdout && l.config.enableColor {
    colorMsg := fmt.Sprintf("%s%s%s",
        level.ColorCode(),  // Start color
        logMsg,             // Message
        ResetColor(),       // Reset color
    )
    l.stdoutWriter.Write([]byte(colorMsg))
}

// File output has no color / 파일 출력에는 색상 없음
if l.config.enableFile {
    l.fileWriter.Write([]byte(logMsg))  // Plain text
}
```

### Adding New Log Levels / 새 로그 레벨 추가

If you need to add a new level:

새 레벨을 추가해야 하는 경우:

```go
// Step 1: Add to Level constants / 1단계: Level 상수에 추가
const (
    DEBUG Level = iota
    INFO
    WARN
    ERROR
    FATAL
    CRITICAL  // New level / 새 레벨
)

// Step 2: Update String() method / 2단계: String() 메서드 업데이트
func (l Level) String() string {
    switch l {
    // ... existing cases
    case CRITICAL:
        return "CRITICAL"
    default:
        return "UNKNOWN"
    }
}

// Step 3: Update ColorCode() method / 3단계: ColorCode() 메서드 업데이트
func (l Level) ColorCode() string {
    switch l {
    // ... existing cases
    case CRITICAL:
        return "\033[1;31m"  // Bold red / 굵은 빨간색
    default:
        return "\033[0m"
    }
}

// Step 4: Add logging method / 4단계: 로깅 메서드 추가
func (l *Logger) Critical(msg string, keysAndValues ...interface{}) {
    l.log(CRITICAL, msg, keysAndValues...)
    os.Exit(2)  // Exit with code 2 / 코드 2로 종료
}
```

---

## Banner System / 배너 시스템

### Banner Architecture / 배너 아키텍처

```go
// Banner prints a formatted banner
// Banner는 형식화된 배너를 출력합니다
func (l *Logger) Banner(appName, version string) {
    // 1. Build banner string / 배너 문자열 작성
    var banner strings.Builder

    // 2. Calculate dimensions / 크기 계산
    text := fmt.Sprintf("%s %s", appName, version)
    width := max(len(text) + 12, 60)

    // 3. Build box / 박스 작성
    // Top border / 상단 경계
    banner.WriteString("╔" + strings.Repeat("═", width) + "╗\n")

    // Empty line / 빈 줄
    banner.WriteString("║" + strings.Repeat(" ", width) + "║\n")

    // Centered text / 중앙 정렬 텍스트
    padding := (width - len(text)) / 2
    banner.WriteString("║" +
        strings.Repeat(" ", padding) +
        text +
        strings.Repeat(" ", width-padding-len(text)) +
        "║\n")

    // Empty line / 빈 줄
    banner.WriteString("║" + strings.Repeat(" ", width) + "║\n")

    // Bottom border / 하단 경계
    banner.WriteString("╚" + strings.Repeat("═", width) + "╝\n")

    // 4. Print using printRaw / printRaw를 사용하여 출력
    l.printRaw(banner.String())
}
```

### printRaw Implementation / printRaw 구현

```go
// printRaw prints without timestamp/level formatting
// printRaw는 타임스탬프/레벨 형식 없이 출력합니다
func (l *Logger) printRaw(text string) {
    l.mu.Lock()
    defer l.mu.Unlock()

    // Write to stdout / stdout에 작성
    if l.config.enableStdout {
        l.stdoutWriter.Write([]byte(text))
    }

    // Write to file / 파일에 작성
    if l.config.enableFile && l.fileWriter != nil {
        l.fileWriter.Write([]byte(text))
    }
}
```

### Smart App Name Extraction / 스마트 앱 이름 추출

```go
// Auto-extract app name from filename
// 파일명에서 앱 이름 자동 추출
if bannerName == "Application" && cfg.filename != "" {
    base := filepath.Base(cfg.filename)     // "database.log"
    ext := filepath.Ext(base)               // ".log"
    if ext != "" {
        bannerName = base[:len(base)-len(ext)]  // "database"
    }
}
```

**Example Flow / 예제 흐름:**

```
Input:  "./logs/api-server.log"
        "./logs/api-server.log"

Step 1: filepath.Base() → "api-server.log"
1단계:  filepath.Base() → "api-server.log"

Step 2: filepath.Ext() → ".log"
2단계:  filepath.Ext() → ".log"

Step 3: Remove extension → "api-server"
3단계:  확장자 제거 → "api-server"

Output: "api-server"
출력:   "api-server"
```

---

## File Rotation with Lumberjack / Lumberjack을 사용한 파일 로테이션

### Lumberjack Integration / Lumberjack 통합

```go
import "gopkg.in/natefinch/lumberjack.v2"

// Create Lumberjack logger
// Lumberjack 로거 생성
logger.fileWriter = &lumberjack.Logger{
    Filename:   cfg.filename,     // "./logs/app.log"
    MaxSize:    cfg.maxSize,      // 100 MB
    MaxBackups: cfg.maxBackups,   // 3 files
    MaxAge:     cfg.maxAge,       // 28 days
    Compress:   cfg.compress,     // true = gzip
}
```

### How Lumberjack Works / Lumberjack 작동 방식

```
1. Write log message to file
   로그 메시지를 파일에 작성

2. Lumberjack checks file size
   Lumberjack이 파일 크기 확인

3. If size >= MaxSize:
   크기 >= MaxSize인 경우:
   a. Rename current file to app-TIMESTAMP.log
      현재 파일을 app-TIMESTAMP.log로 이름 변경
   b. Create new app.log
      새 app.log 생성
   c. If Compress: gzip old file
      Compress가 true면: 이전 파일을 gzip 압축

4. If old files > MaxBackups:
   이전 파일 > MaxBackups인 경우:
   Delete oldest file
   가장 오래된 파일 삭제

5. If file age > MaxAge:
   파일 수명 > MaxAge인 경우:
   Delete old files
   오래된 파일 삭제
```

### Rotation Example / 로테이션 예제

```
Initial state / 초기 상태:
logs/app.log (95 MB)

After writing 10 MB / 10 MB 작성 후:
logs/app.log (5 MB)              ← New file / 새 파일
logs/app-2025-10-10-143045.log (100 MB)  ← Rotated / 로테이션됨

After compression / 압축 후:
logs/app.log (5 MB)
logs/app-2025-10-10-143045.log.gz (10 MB)  ← Compressed / 압축됨
```

### Manual Rotation / 수동 로테이션

```go
func (l *Logger) Rotate() error {
    l.mu.Lock()
    defer l.mu.Unlock()

    if l.fileWriter != nil {
        return l.fileWriter.Rotate()
    }
    return nil
}
```

---

## Version Management / 버전 관리

### app.yaml Loading / app.yaml 로딩

```go
func LoadAppConfig() (*AppConfig, error) {
    // Search paths / 검색 경로
    searchPaths := []string{
        "cfg/app.yaml",
        "apps/app.yaml",
        "app.yaml",
        "../cfg/app.yaml",
        // ... more paths
    }

    // Try each path / 각 경로 시도
    for _, path := range searchPaths {
        config, err := loadFromPath(path)
        if err == nil {
            return config, nil  // Found / 찾음
        }
    }

    return &AppConfig{}, fmt.Errorf("app.yaml not found")
}
```

### Integration with Logger / 로거와의 통합

```go
func defaultConfig() *config {
    // Try to load from app.yaml
    // app.yaml에서 로드 시도
    appVersion := TryLoadAppVersion()
    if appVersion == "" {
        appVersion = "v1.0.0"  // Fallback / 대체
    }

    appName := TryLoadAppName()
    if appName == "" {
        appName = "Application"  // Fallback / 대체
    }

    return &config{
        // ... other settings
        appName:    appName,
        appVersion: appVersion,
    }
}
```

---

## Testing Guide / 테스트 가이드

### Test Structure / 테스트 구조

```
logger_test.go
├── Unit Tests / 단위 테스트
│   ├── TestNew
│   ├── TestLogLevels
│   ├── TestStructuredLogging
│   ├── TestPrintfStyleLogging
│   └── ...
├── Integration Tests / 통합 테스트
│   ├── TestFileRotation
│   ├── TestMultipleLoggers
│   └── TestAppYamlIntegration
└── Benchmarks / 벤치마크
    ├── BenchmarkInfo
    └── BenchmarkInfof
```

### Running Tests / 테스트 실행

```bash
# Run all tests / 모든 테스트 실행
go test -v

# Run specific test / 특정 테스트 실행
go test -v -run TestLogLevels

# Run with coverage / 커버리지와 함께 실행
go test -cover

# Generate coverage report / 커버리지 리포트 생성
go test -coverprofile=coverage.out
go tool cover -html=coverage.out

# Run benchmarks / 벤치마크 실행
go test -bench=.

# Run benchmarks with memory stats / 메모리 통계와 함께 벤치마크 실행
go test -bench=. -benchmem
```

### Writing Tests / 테스트 작성

**Example: Testing a New Feature / 예제: 새 기능 테스트**

```go
func TestNewFeature(t *testing.T) {
    // Setup / 설정
    tempDir := t.TempDir()  // Auto-cleanup / 자동 정리
    logFile := filepath.Join(tempDir, "test.log")

    // Create logger / 로거 생성
    logger, err := logging.New(
        logging.WithFilePath(logFile),
        logging.WithStdoutOnly(),  // No file for this test
    )
    if err != nil {
        t.Fatalf("Failed to create logger: %v", err)
    }
    defer logger.Close()

    // Test the feature / 기능 테스트
    logger.Info("Test message")

    // Verify / 검증
    // ... assertions
}
```

### Test Helpers / 테스트 헬퍼

```go
// Helper: Create test logger / 헬퍼: 테스트 로거 생성
func createTestLogger(t *testing.T) (*logging.Logger, string) {
    tempDir := t.TempDir()
    logFile := filepath.Join(tempDir, "test.log")

    logger, err := logging.New(
        logging.WithFilePath(logFile),
        logging.WithAutoBanner(false),
    )
    if err != nil {
        t.Fatalf("Failed to create logger: %v", err)
    }

    return logger, logFile
}

// Helper: Read log file / 헬퍼: 로그 파일 읽기
func readLogFile(t *testing.T, path string) string {
    data, err := os.ReadFile(path)
    if err != nil {
        t.Fatalf("Failed to read log file: %v", err)
    }
    return string(data)
}
```

---

## Contributing Guidelines / 기여 가이드라인

### Contribution Process / 기여 프로세스

1. **Fork the repository / 저장소 포크**
2. **Create a feature branch / 기능 브랜치 생성**
   ```bash
   git checkout -b feature/my-new-feature
   ```
3. **Make changes / 변경 사항 작성**
4. **Write tests / 테스트 작성**
   - All new features must have tests / 모든 새 기능에는 테스트 필요
   - Maintain >80% coverage / >80% 커버리지 유지
5. **Run tests / 테스트 실행**
   ```bash
   go test -v
   go test -cover
   ```
6. **Commit with descriptive message / 설명적 메시지로 커밋**
   ```bash
   git commit -m "feat: add new logging feature"
   ```
7. **Push and create PR / 푸시 및 PR 생성**

### Contribution Checklist / 기여 체크리스트

- [ ] Code compiles without errors / 코드가 오류 없이 컴파일됨
- [ ] All tests pass / 모든 테스트 통과
- [ ] New tests added for new features / 새 기능에 대한 새 테스트 추가
- [ ] Code coverage >80% / 코드 커버리지 >80%
- [ ] Documentation updated / 문서 업데이트
- [ ] Bilingual comments (English/Korean) / 이중 언어 주석 (영문/한글)
- [ ] Follows code style guidelines / 코드 스타일 가이드라인 준수
- [ ] No breaking changes (or clearly documented) / 주요 변경 사항 없음 (또는 명확히 문서화)

### What to Contribute / 기여 가능한 내용

**Welcome contributions / 환영하는 기여:**

- Bug fixes / 버그 수정
- New features / 새 기능
- Performance improvements / 성능 개선
- Documentation improvements / 문서 개선
- Test coverage improvements / 테스트 커버리지 개선
- Examples / 예제

**Please discuss first / 먼저 논의해주세요:**

- Major architectural changes / 주요 아키텍처 변경
- Breaking changes / 주요 변경 사항
- New dependencies / 새 의존성

---

## Code Style / 코드 스타일

### Naming Conventions / 명명 규칙

**Packages / 패키지:**
```go
package logging  // Lowercase, singular / 소문자, 단수형
```

**Functions / 함수:**
```go
// Public: PascalCase / 공개: 파스칼 케이스
func New() *Logger {}
func (l *Logger) Info() {}

// Private: camelCase / 비공개: 카멜 케이스
func defaultConfig() *config {}
func (l *Logger) log() {}
```

**Variables / 변수:**
```go
// Public: PascalCase / 공개: 파스칼 케이스
const MaxSize = 100

// Private: camelCase / 비공개: 카멜 케이스
var enableColor = true
```

**Types / 타입:**
```go
// Public: PascalCase / 공개: 파스칼 케이스
type Logger struct {}
type Level int

// Private: camelCase / 비공개: 카멜 케이스
type config struct {}
```

### Comments / 주석

**Function comments / 함수 주석:**

```go
// Info logs a message at INFO level
// Info는 INFO 레벨로 메시지를 로깅합니다
//
// Parameters / 매개변수:
//   - msg: log message / 로그 메시지
//   - keysAndValues: optional key-value pairs / 선택적 키-값 쌍
//
// Example / 예제:
//
//  logger.Info("Server started", "port", 8080)
func (l *Logger) Info(msg string, keysAndValues ...interface{}) {
    l.log(INFO, msg, keysAndValues...)
}
```

**Inline comments / 인라인 주석:**

```go
// Acquire mutex lock / Mutex 잠금 획득
l.mu.Lock()
defer l.mu.Unlock()

// Format timestamp / 타임스탬프 형식화
timestamp := time.Now().Format(l.config.timeFormat)
```

### Error Handling / 에러 처리

**Always check errors / 항상 에러 확인:**

```go
// ✅ Good / 좋음
logger, err := logging.New(opts...)
if err != nil {
    return fmt.Errorf("failed to create logger: %w", err)
}

// ❌ Bad / 나쁨
logger, _ := logging.New(opts...)
```

**Wrap errors with context / 컨텍스트와 함께 에러 래핑:**

```go
if err := os.MkdirAll(logDir, 0755); err != nil {
    return fmt.Errorf("failed to create log directory: %w", err)
}
```

### Formatting / 형식화

**Use gofmt / gofmt 사용:**

```bash
gofmt -w .
```

**Use goimports / goimports 사용:**

```bash
goimports -w .
```

### Best Practices / 모범 사례

1. **Keep functions small / 함수를 작게 유지**
   - Max 50 lines per function / 함수당 최대 50줄
   - Single responsibility / 단일 책임

2. **Avoid deep nesting / 깊은 중첩 피하기**
   ```go
   // ✅ Good - Early return / 좋음 - 조기 반환
   if err != nil {
       return err
   }
   // Continue...

   // ❌ Bad - Nested / 나쁨 - 중첩
   if err == nil {
       // Long code...
   }
   ```

3. **Use meaningful variable names / 의미 있는 변수 이름 사용**
   ```go
   // ✅ Good / 좋음
   timestamp := time.Now()

   // ❌ Bad / 나쁨
   t := time.Now()
   ```

4. **Document all public APIs / 모든 공개 API 문서화**
5. **Write tests for all features / 모든 기능에 대한 테스트 작성**
6. **Use bilingual comments / 이중 언어 주석 사용**

---

**End of Developer Guide / 개발자 가이드 끝**

For user documentation, see [USER_MANUAL.md](./USER_MANUAL.md).

사용자 문서는 [USER_MANUAL.md](./USER_MANUAL.md)를 참조하세요.
