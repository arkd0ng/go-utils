# Example Code Writing Guide / 예제 코드 작성 가이드

This guide provides comprehensive standards for writing example code in the go-utils project.

이 가이드는 go-utils 프로젝트에서 예제 코드를 작성하기 위한 포괄적인 표준을 제공합니다.

## Table of Contents / 목차

- [Overview / 개요](#overview--개요)
- [Directory Structure / 디렉토리 구조](#directory-structure--디렉토리-구조)
- [File Naming Conventions / 파일 명명 규칙](#file-naming-conventions--파일-명명-규칙)
- [Logging Standards / 로깅 표준](#logging-standards--로깅-표준)
- [Code Structure / 코드 구조](#code-structure--코드-구조)
- [Example Categories / 예제 카테고리](#example-categories--예제-카테고리)
- [Documentation Standards / 문서화 표준](#documentation-standards--문서화-표준)
- [Testing Examples / 예제 테스트](#testing-examples--예제-테스트)
- [Package-Specific Guidelines / 패키지별 가이드라인](#package-specific-guidelines--패키지별-가이드라인)

---

## Overview / 개요

### Purpose / 목적

Example code serves three main purposes:

예제 코드는 세 가지 주요 목적을 제공합니다:

1. **Learning Tool / 학습 도구**: Help users understand how to use the package
2. **Reference / 참조**: Provide copy-paste ready code snippets
3. **Testing / 테스트**: Verify that the package works as expected in real scenarios

### Core Principles / 핵심 원칙

1. **Completeness / 완전성**: Cover ALL functions in the package
   - ✅ **MUST include every public function** / 모든 공개 함수를 반드시 포함
   - ✅ **MUST demonstrate every function at least once** / 모든 함수를 최소 한 번 시연
   - ✅ **100% function coverage in examples** / 예제의 100% 함수 커버리지
   - ⚠️ **Examples are incomplete if ANY function is missing** / 함수가 하나라도 누락되면 예제는 불완전함

2. **Real-World Usage / 실제 사용**: Show practical, production-ready examples
   - Individual function examples / 개별 함수 예제
   - Combination examples (multiple functions together) / 복합 예제 (여러 함수 조합)
   - Real-world scenarios / 실제 사용 시나리오
   - Edge cases and error handling / 엣지 케이스 및 에러 처리

3. **Detailed Logging / 상세한 로깅**: Log everything so users don't need to read docs
   - Every function call logged / 모든 함수 호출 로깅
   - All parameters and results logged / 모든 매개변수 및 결과 로깅

4. **Bilingual / 이중 언어**: All comments and logs in English and Korean
   - All comments bilingual / 모든 주석 이중 언어
   - All log messages bilingual / 모든 로그 메시지 이중 언어

---

## Directory Structure / 디렉토리 구조

```
go-utils/
├── examples/
│   └── {package_name}/
│       ├── main.go              # Main example file / 메인 예제 파일
│       └── README.md            # Optional: Package-specific notes / 선택 사항: 패키지별 노트
└── logs/                        # Shared log directory / 공용 로그 디렉토리
    ├── {package}-example.log                 # Current log / 현재 로그
    └── {package}-example-YYYYMMDD-HHMMSS.log # Backup logs / 백업 로그
```

### Example Directory Names / 예제 디렉토리 이름

- `random_string/` - for `random` package
- `logging/` - for `logging` package
- `mysql/` - for `database/mysql` package
- `redis/` - for `database/redis` package
- `stringutil/` - for `stringutil` package
- `timeutil/` - for `timeutil` package
- `sliceutil/` - for `sliceutil` package
- `maputil/` - for `maputil` package
- `fileutil/` - for `fileutil` package
- `websvrutil/` - for `websvrutil` package

---

## File Naming Conventions / 파일 명명 규칙

### Log Files / 로그 파일

All example logs MUST be written to `go-utils/logs/` (shared across packages).  
모든 예제 로그는 패키지 공용 디렉토리인 `go-utils/logs/`에 기록되어야 합니다.

**Current Log File Format / 현재 로그 파일 형식**:
```
logs/{package}-example.log
```

Examples / 예제:
- `logs/mysql-example.log`
- `logs/redis-example.log`
- `logs/websvrutil-example.log`

**Backup Log File Format / 백업 로그 파일 형식**:
```
logs/{package}-example-YYYYMMDD-HHMMSS.log
```

Examples / 예제:
- `logs/mysql-example-20251016-143025.log`
- `logs/redis-example-20251016-143530.log`
- `logs/websvrutil-example-20251016-144012.log`

**Backup Retention Policy / 백업 보관 정책**:
- Keep only the 5 most recent backup files
- Delete older backup files automatically
- 최근 5개의 백업 파일만 유지
- 오래된 백업 파일 자동 삭제

---

## Logging Standards / 로깅 표준

### Log Setup Template / 로그 설정 템플릿

All examples MUST use this log setup pattern:

모든 예제는 이 로그 설정 패턴을 사용해야 합니다:

```go
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/arkd0ng/go-utils/fileutil"
	"github.com/arkd0ng/go-utils/logging"
)

func main() {
	// Setup log file with backup management / 백업 관리와 함께 로그 파일 설정
	logFilePath := "logs/{package}-example.log"

	// Check if previous log file exists / 이전 로그 파일 존재 여부 확인
	if fileutil.Exists(logFilePath) {
		// Get modification time of existing log file / 기존 로그 파일의 수정 시간 가져오기
		modTime, err := fileutil.ModTime(logFilePath)
		if err == nil {
			// Create backup filename with timestamp / 타임스탬프와 함께 백업 파일명 생성
			backupName := fmt.Sprintf("logs/{package}-example-%s.log", modTime.Format("20060102-150405"))

			// Backup existing log file / 기존 로그 파일 백업
			if err := fileutil.CopyFile(logFilePath, backupName); err == nil {
				fmt.Printf("✅ Backed up previous log to: %s\n", backupName)
				// Delete original log file to prevent content duplication / 내용 중복 방지를 위해 원본 로그 파일 삭제
				fileutil.DeleteFile(logFilePath)
			}
		}

		// Cleanup old backup files - keep only 5 most recent / 오래된 백업 파일 정리 - 최근 5개만 유지
		backupPattern := "logs/{package}-example-*.log"
		backupFiles, err := filepath.Glob(backupPattern)
		if err == nil && len(backupFiles) > 5 {
			// Sort by modification time / 수정 시간으로 정렬
			type fileInfo struct {
				path    string
				modTime time.Time
			}
			var files []fileInfo
			for _, f := range backupFiles {
				if mt, err := fileutil.ModTime(f); err == nil {
					files = append(files, fileInfo{path: f, modTime: mt})
				}
			}

			// Sort oldest first / 가장 오래된 것부터 정렬
			for i := 0; i < len(files)-1; i++ {
				for j := i + 1; j < len(files); j++ {
					if files[i].modTime.After(files[j].modTime) {
						files[i], files[j] = files[j], files[i]
					}
				}
			}

			// Delete oldest files to keep only 5 / 5개만 유지하도록 가장 오래된 파일 삭제
			for i := 0; i < len(files)-5; i++ {
				fileutil.DeleteFile(files[i].path)
				fmt.Printf("🗑️  Deleted old backup: %s\n", files[i].path)
			}
		}
	}

	// Initialize logger with fixed filename / 고정 파일명으로 로거 초기화
	logger, err := logging.New(
		logging.WithFilePath(logFilePath),
		logging.WithLevel(logging.DEBUG),
		logging.WithMaxSize(10),    // 10 MB
		logging.WithMaxBackups(5),
		logging.WithMaxAge(30),     // 30 days
		logging.WithCompress(true),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Close()

	// Log application start / 애플리케이션 시작 로그
	logger.Info("===========================================")
	logger.Info("Starting {Package} Package Examples")
	logger.Info("{패키지} 패키지 예제 시작")
	logger.Info("===========================================")

	// Your example code here...
}
```

### Logging Guidelines / 로깅 가이드라인

1. **Log Everything / 모든 것을 로그**:
   - Every function call
   - Every parameter value
   - Every result
   - Every error (even if handled)
   - Entry and exit of each example function

2. **Log Level Usage / 로그 레벨 사용**:
   - `DEBUG`: Detailed execution flow / 상세한 실행 흐름
   - `INFO`: Normal operation / 정상 작업
   - `WARN`: Potential issues / 잠재적 문제
   - `ERROR`: Errors that were handled / 처리된 에러
   - `FATAL`: Unrecoverable errors / 복구 불가능한 에러

3. **Bilingual Logging / 이중 언어 로깅**:
   ```go
   logger.Info("Starting Example 1: Basic Usage")
   logger.Info("예제 1 시작: 기본 사용법")
   ```

4. **Structured Logging / 구조화된 로깅**:
   ```go
   logger.Info("Function called", "function", "DoSomething", "param1", value1, "param2", value2)
   logger.Info("함수 호출됨", "function", "DoSomething", "param1", value1, "param2", value2)
   ```

5. **Section Separators / 섹션 구분자**:
   ```go
   logger.Info("===========================================")
   logger.Info("Example 1: Basic Usage / 예제 1: 기본 사용법")
   logger.Info("===========================================")
   ```

6. **Result Logging / 결과 로깅**:
   ```go
   logger.Info("Result", "value", result, "type", fmt.Sprintf("%T", result))
   logger.Info("결과", "value", result, "type", fmt.Sprintf("%T", result))
   ```

7. **Timing Information / 시간 정보**:
   ```go
   start := time.Now()
   // ... operation ...
   logger.Info("Operation completed", "duration", time.Since(start))
   logger.Info("작업 완료", "duration", time.Since(start))
   ```

---

## Code Structure / 코드 구조

### Main Function Template / 메인 함수 템플릿

```go
func main() {
	// 1. Setup logging (as shown above)
	// 1. 로깅 설정 (위에서 보여준 대로)

	// 2. Print header to console
	// 2. 콘솔에 헤더 출력
	fmt.Println("=== {Package} Package Examples ===")
	fmt.Println("=== {패키지} 패키지 예제 ===\n")

	// 3. Run each example with descriptive output
	// 3. 설명적인 출력과 함께 각 예제 실행

	// Example 1: Description
	fmt.Println("Example 1: Basic Usage / 기본 사용법")
	example1BasicUsage(logger)

	// Example 2: Description
	fmt.Println("\nExample 2: Advanced Usage / 고급 사용법")
	example2AdvancedUsage(logger)

	// ... more examples ...

	// 4. Print footer
	// 4. 푸터 출력
	fmt.Println("\n=== All Examples Completed ===")
	fmt.Println("=== 모든 예제 완료 ===")

	// 5. Final log entry
	// 5. 최종 로그 항목
	logger.Info("===========================================")
	logger.Info("All examples completed successfully")
	logger.Info("모든 예제가 성공적으로 완료되었습니다")
	logger.Info("===========================================")
}
```

### Example Function Template / 예제 함수 템플릿

```go
// example1BasicUsage demonstrates basic usage of the package.
// example1BasicUsage는 패키지의 기본 사용법을 시연합니다.
func example1BasicUsage(logger *logging.Logger) {
	logger.Info("===========================================")
	logger.Info("Example 1: Basic Usage / 예제 1: 기본 사용법")
	logger.Info("===========================================")

	// Step 1: Describe what we're doing
	// 단계 1: 무엇을 하는지 설명
	logger.Info("Step 1: Initialize the component")
	logger.Info("단계 1: 컴포넌트 초기화")

	// Code example
	// 코드 예제
	component := somepackage.New()
	logger.Info("Component created", "type", fmt.Sprintf("%T", component))
	logger.Info("컴포넌트 생성됨", "type", fmt.Sprintf("%T", component))

	// Step 2: Perform operation
	// 단계 2: 작업 수행
	logger.Info("Step 2: Perform basic operation")
	logger.Info("단계 2: 기본 작업 수행")

	result, err := component.DoSomething("input")
	if err != nil {
		logger.Error("Operation failed", "error", err)
		logger.Error("작업 실패", "error", err)
		fmt.Printf("❌ Error: %v\n", err)
		return
	}

logger.Info("Operation succeeded", "result", result)
logger.Info("작업 성공", "result", result)
fmt.Printf("✅ Result: %v\n", result)

logger.Info("Example 1 completed successfully")
logger.Info("예제 1 완료")
}
```

- **Structured bilingual logging / 구조화된 이중 언어 로그 예시**:

```
logger.Info("Example 12: Custom Middleware / 커스텀 미들웨어")
logger.Info("Request", "method", r.Method, "path", r.URL.Path)
logger.Info("요청", "method", r.Method, "path", r.URL.Path)
logger.Info("Response", "status", rec.Code, "duration", time.Since(start))
logger.Info("응답", "status", rec.Code, "duration", time.Since(start))
logger.Info("Artifacts", "headers", rec.Header(), "cookie_count", len(rec.Result().Cookies()))
logger.Info("산출물", "headers", rec.Header(), "cookie_count", len(rec.Result().Cookies()))
```

- Always include inputs, outputs, headers, status codes, and generated artifacts (file paths, tokens, session IDs).  
  입력·출력·헤더·상태 코드·생성된 산출물(파일 경로, 토큰, 세션 ID)을 반드시 기록하세요.

---

## Example Categories / 예제 카테고리

Every package should include examples in these categories:

모든 패키지는 다음 카테고리의 예제를 포함해야 합니다:

### 1. Basic Examples / 기본 예제

Cover every single function with basic usage:

모든 함수를 기본 사용법으로 다룹니다:

```go
// Example: Cover each function individually
// 예제: 각 함수를 개별적으로 다룹니다

func example1FunctionA(logger *logging.Logger) { /* ... */ }
func example2FunctionB(logger *logging.Logger) { /* ... */ }
func example3FunctionC(logger *logging.Logger) { /* ... */ }
```

**Requirements / 요구사항**:
- One example per public function
- Show typical use case
- Include error handling
- Log input parameters and output results

### 2. Combination Examples / 복합 예제

Show how multiple functions work together:

여러 함수가 함께 작동하는 방법을 보여줍니다:

```go
func exampleCombined(logger *logging.Logger) {
	logger.Info("Demonstrating combined usage of FunctionA and FunctionB")
	logger.Info("FunctionA와 FunctionB의 복합 사용 시연")

	// Use Function A
	resultA := pkg.FunctionA()
	logger.Info("FunctionA result", "value", resultA)

	// Pass result to Function B
	resultB := pkg.FunctionB(resultA)
	logger.Info("FunctionB result", "value", resultB)

	// Show the combined effect
	logger.Info("Combined result", "final", resultB)
}
```

### 3. Real-World Examples / 실제 사용 예제

Show common production scenarios:

일반적인 프로덕션 시나리오를 보여줍니다:

```go
// Example: RESTful API with all middleware
func exampleProductionRESTAPI(logger *logging.Logger) {
	logger.Info("=== Production REST API Example ===")
	logger.Info("=== 프로덕션 REST API 예제 ===")

	// Setup with all production settings
	app := websvrutil.New(
		websvrutil.WithReadTimeout(30*time.Second),
		websvrutil.WithWriteTimeout(30*time.Second),
		websvrutil.WithMaxHeaderBytes(1<<20), // 1 MB
	)

	// Add production middleware
	app.Use(websvrutil.Logger())
	app.Use(websvrutil.Recovery())
	app.Use(websvrutil.CORS())
	app.Use(websvrutil.CSRF())

	// Define routes
	app.GET("/health", healthCheckHandler)
	app.POST("/api/users", createUserHandler)
	app.GET("/api/users/:id", getUserHandler)

	logger.Info("Production REST API configured")
	logger.Info("프로덕션 REST API 설정 완료")
}
```

### 4. Edge Case Examples / 엣지 케이스 예제

Show how to handle edge cases and errors:

엣지 케이스와 에러를 처리하는 방법을 보여줍니다:

```go
func exampleEdgeCases(logger *logging.Logger) {
	logger.Info("=== Edge Case Examples ===")
	logger.Info("=== 엣지 케이스 예제 ===")

	// Test with empty input
	logger.Info("Testing with empty input")
	logger.Info("빈 입력으로 테스트")
	result1, err := pkg.Function("")
	if err != nil {
		logger.Warn("Expected error with empty input", "error", err)
		logger.Warn("빈 입력으로 예상된 에러", "error", err)
	}

	// Test with nil input
	logger.Info("Testing with nil input")
	logger.Info("nil 입력으로 테스트")
	result2, err := pkg.Function(nil)
	// ... handle error ...

	// Test with very large input
	logger.Info("Testing with large input")
	logger.Info("큰 입력으로 테스트")
	largeInput := strings.Repeat("x", 1000000) // 1 MB
	result3, err := pkg.Function(largeInput)
	// ... handle error ...
}
```

### 5. Performance Examples / 성능 예제

Show performance characteristics and benchmarks:

성능 특성과 벤치마크를 보여줍니다:

```go
func examplePerformance(logger *logging.Logger) {
	logger.Info("=== Performance Example ===")
	logger.Info("=== 성능 예제 ===")

	// Benchmark single operation
	start := time.Now()
	result := pkg.Function(input)
	duration := time.Since(start)

	logger.Info("Single operation",
		"duration", duration,
		"ns/op", duration.Nanoseconds(),
	)

	// Benchmark bulk operations
	count := 10000
	start = time.Now()
	for i := 0; i < count; i++ {
		pkg.Function(input)
	}
	duration = time.Since(start)

	logger.Info("Bulk operations",
		"count", count,
		"total_duration", duration,
		"avg_ns/op", duration.Nanoseconds()/int64(count),
	)
}
```

### 6. Integration Examples / 통합 예제

Show integration with other packages:

다른 패키지와의 통합을 보여줍니다:

```go
func exampleIntegration(logger *logging.Logger) {
	logger.Info("=== Integration Example ===")
	logger.Info("=== 통합 예제 ===")

	// Integrate with database
	db, _ := mysql.New(/* ... */)
	defer db.Close()

	// Integrate with web server
	app := websvrutil.New()
	app.POST("/api/data", func(w http.ResponseWriter, r *http.Request) {
		// Use database in handler
		data, err := db.Query("SELECT * FROM users")
		// ... handle response ...
	})

	logger.Info("Integration example configured")
	logger.Info("통합 예제 설정 완료")
}
```

---

## Documentation Standards / 문서화 표준

### Function Documentation / 함수 문서화

Every example function must have:

모든 예제 함수는 다음을 가져야 합니다:

```go
// example1BasicUsage demonstrates basic usage of the XYZ function.
// This example shows how to:
// - Initialize the component
// - Perform basic operations
// - Handle common errors
//
// example1BasicUsage는 XYZ 함수의 기본 사용법을 시연합니다.
// 이 예제는 다음을 보여줍니다:
// - 컴포넌트 초기화
// - 기본 작업 수행
// - 일반적인 에러 처리
func example1BasicUsage(logger *logging.Logger) {
	// Implementation...
}
```

### Code Comments / 코드 주석

1. **Before Each Section / 각 섹션 전**:
   ```go
   // Step 1: Initialize client with custom options
   // 단계 1: 커스텀 옵션으로 클라이언트 초기화
   ```

2. **Inline Comments for Complex Code / 복잡한 코드에 인라인 주석**:
   ```go
   // Convert to lowercase for case-insensitive comparison
   // 대소문자 구분 없는 비교를 위해 소문자로 변환
   normalized := strings.ToLower(input)
   ```

3. **Error Handling Comments / 에러 처리 주석**:
   ```go
   if err != nil {
		// This error is expected when the resource doesn't exist
		// 리소스가 존재하지 않을 때 이 에러가 예상됩니다
		logger.Warn("Resource not found", "error", err)
		return
	}
   ```

### Console Output Format / 콘솔 출력 형식

Use Unicode symbols for visual clarity:

시각적 명확성을 위해 유니코드 기호를 사용합니다:

```go
fmt.Println("✅ Success: Operation completed") // Success
fmt.Println("❌ Error: Operation failed")      // Error
fmt.Println("⚠️  Warning: Potential issue")    // Warning
fmt.Println("ℹ️  Info: Additional information") // Info
fmt.Println("🔍 Debug: Detailed information")  // Debug
fmt.Println("📊 Result: %v", result)           // Result
fmt.Println("🔧 Config: Settings applied")     // Configuration
fmt.Println("🚀 Starting: Operation begins")   // Start
fmt.Println("🏁 Finished: Operation complete") // Finish
fmt.Println("📝 Note: Important information")  // Note
```

---

## Testing Examples / 예제 테스트

### Running Examples / 예제 실행

All examples must be runnable:

모든 예제는 실행 가능해야 합니다:

```bash
# Run the example from repository root / 저장소 루트에서 예제 실행
go run ./examples/{package_name}

# Check logs in shared directory / 공용 디렉토리에서 로그 확인
cat logs/{package}-example.log

# Check backup logs / 백업 로그 확인
ls -l logs/
```

### Example Testing Checklist / 예제 테스트 체크리스트

Before committing example code, verify:

예제 코드를 커밋하기 전에 확인:

- [ ] All functions in the package are demonstrated
- [ ] Code compiles without errors
- [ ] All examples run successfully
- [ ] Logs are created in repository root `logs/` directory / 저장소 루트 `logs/` 디렉토리에 로그 생성
- [ ] Log backup system works correctly
- [ ] Old backups are cleaned up (only 5 kept)
- [ ] All comments are bilingual (English/Korean)
- [ ] Console output uses Unicode symbols
- [ ] Error cases are handled gracefully
- [ ] Performance examples include timing
- [ ] Integration examples work with dependencies

---

## Package-Specific Guidelines / 패키지별 가이드라인

### websvrutil Package / websvrutil 패키지

**Must Include / 포함해야 할 예제**:

1. **Basic Server Examples / 기본 서버 예제**:
   - Creating a server with default options
   - Custom server options (timeouts, max header size)
   - Graceful shutdown

2. **Routing Examples / 라우팅 예제**:
   - GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS
   - Path parameters (`:id`, `:name`)
   - Wildcard routes (`/files/*filepath`)
   - Route groups with prefix
   - Custom 404 handler

3. **Context Examples / Context 예제**:
   - Path parameters: `Param()`, `PathParam()`
   - Query parameters: `Query()`, `QueryParam()`, `QueryParams()`
   - Form data: `FormValue()`, `FormParams()`
   - Headers: `GetHeader()`, `AddHeader()`, `GetHeaders()`
   - Cookies: `GetCookie()`, `SetCookie()`, `DeleteCookie()`
   - Custom values: `Set()`, `Get()`, `MustGet()`
   - Client IP: `ClientIP()`
   - User Agent: `UserAgent()`
   - Content type: `ContentType()`
   - Referer: `Referer()`

4. **Request Binding Examples / 요청 바인딩 예제**:
   - `Bind()`: Auto-detect content type
   - `BindJSON()`: JSON body with size limit
   - `BindForm()`: Form data
   - `BindWithValidation()`: Bind + validate
   - `BindQuery()`: Query parameters

5. **Response Examples / 응답 예제**:
   - `JSON()`: JSON response
   - `String()`: Plain text
   - `HTML()`: HTML response
   - `Redirect()`: HTTP redirect
   - `File()`: Send file
   - `FileAttachment()`: Download file
   - `Status()`: Set status code
   - `NoContent()`: 204 response

6. **Middleware Examples / 미들웨어 예제**:
   - Custom middleware creation
   - Logger middleware
   - Recovery middleware
   - CORS middleware
   - CSRF protection
   - Body limit
   - Compression
   - Rate limiting
   - Authentication middleware
   - Multiple middleware chaining

7. **Session Examples / 세션 예제**:
   - Session creation
   - Setting values
   - Getting values
   - Deleting values
   - Session destruction
   - Custom session options
   - Session with database store

8. **Template Examples / 템플릿 예제**:
   - Loading templates
   - Rendering templates
   - Template with layout
   - Custom template functions
   - Auto-reload templates
   - Template data passing

9. **CSRF Examples / CSRF 예제**:
   - CSRF protection setup
   - Token generation
   - Token validation
   - Custom CSRF config
   - CSRF with forms

10. **Validator Examples / 검증자 예제**:
    - All validation tags: `required`, `email`, `min`, `max`, `len`, `eq`, `ne`, `gt`, `gte`, `lt`, `lte`, `oneof`, `alpha`, `alphanum`, `numeric`
    - Multiple tags per field
    - Custom validation messages
    - Validation with binding

11. **File Upload Examples / 파일 업로드 예제**:
    - Single file upload
    - Multiple file upload
    - File size validation
    - File type validation
    - Saving uploaded files

12. **Static Files Examples / 정적 파일 예제**:
    - Serving static directory
    - Single file serving
    - Custom static middleware

13. **Production Examples / 프로덕션 예제**:
    - Complete REST API
    - Microservice architecture
    - Database integration
    - Authentication system
    - API versioning
    - Health check endpoints
    - Metrics endpoints

14. **Testing Examples / 테스트 예제**:
    - Unit testing handlers
    - Integration testing
    - Mocking requests
    - Testing middleware

**Logging Expectations / 로깅 기대치**:
- Mirror every console message to `logs/websvrutil-example.log` (English followed by Korean).  
  모든 콘솔 메시지를 영어 후 한국어 순으로 `logs/websvrutil-example.log`에 기록합니다.
- Capture inputs, headers, payloads, status codes, and artifacts such as saved files or CSRF 토큰.  
  입력값, 헤더, 페이로드, 상태 코드, 저장된 파일·CSRF 토큰과 같은 산출물을 기록합니다.
- Rotate logs: timestamp backups (`logs/websvrutil-example-YYYYMMDD-HHMMSS.log`) and keep only the most recent five.  
  로그 회전: 타임스탬프 백업(`logs/websvrutil-example-YYYYMMDD-HHMMSS.log`)을 생성하고 최신 다섯 개만 유지합니다.

**Logging Requirements / 로깅 요구사항**:

```go
logger.Info("=== Web Server Example ===")
logger.Info("Server Configuration:")
logger.Info("  - Address: :8080")
logger.Info("  - Read Timeout: 30s")
logger.Info("  - Write Timeout: 30s")
logger.Info("")

logger.Info("Registering routes...")
logger.Info("  - GET    /")
logger.Info("  - GET    /health")
logger.Info("  - POST   /api/users")
logger.Info("  - GET    /api/users/:id")
logger.Info("  - PUT    /api/users/:id")
logger.Info("  - DELETE /api/users/:id")
logger.Info("")

logger.Info("Middleware stack:")
logger.Info("  1. Logger")
logger.Info("  2. Recovery")
logger.Info("  3. CORS")
logger.Info("  4. CSRF")
logger.Info("")

logger.Info("Server ready to start")
logger.Info("Press Ctrl+C to gracefully shutdown")
```

---

## Best Practices / 모범 사례

### 1. Progressive Examples / 점진적 예제

Start simple, then add complexity:

간단하게 시작하고 복잡성을 추가합니다:

```go
// Example 1: Minimal setup
func example1Minimal(logger *logging.Logger) {
	app := websvrutil.New()
	app.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
}

// Example 2: Add middleware
func example2WithMiddleware(logger *logging.Logger) {
	app := websvrutil.New()
	app.Use(websvrutil.Logger())
	app.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
}

// Example 3: Add error handling
func example3WithErrorHandling(logger *logging.Logger) {
	app := websvrutil.New()
	app.Use(websvrutil.Logger())
	app.Use(websvrutil.Recovery())
	// ... more complexity ...
}
```

### 2. Self-Documenting Code / 자체 문서화 코드

The logs should tell the complete story:

로그가 완전한 이야기를 전해야 합니다:

```go
logger.Info("===========================================")
logger.Info("Example: User Registration Flow")
logger.Info("예제: 사용자 등록 흐름")
logger.Info("===========================================")

logger.Info("Step 1: Validate user input")
logger.Info("단계 1: 사용자 입력 검증")
// ... validation code ...
logger.Info("✓ Validation passed", "email", user.Email)

logger.Info("Step 2: Check if user already exists")
logger.Info("단계 2: 사용자 존재 여부 확인")
// ... check code ...
logger.Info("✓ User does not exist, proceeding")

logger.Info("Step 3: Hash password")
logger.Info("단계 3: 비밀번호 해시")
// ... hashing code ...
logger.Info("✓ Password hashed", "algorithm", "bcrypt")

logger.Info("Step 4: Save to database")
logger.Info("단계 4: 데이터베이스에 저장")
// ... save code ...
logger.Info("✓ User saved", "id", userID)

logger.Info("Registration completed successfully")
logger.Info("등록 성공적으로 완료")
```

### 3. Error Demonstration / 에러 시연

Show both success and error paths:

성공과 에러 경로를 모두 보여줍니다:

```go
// Success case
logger.Info("Testing valid input...")
result, err := pkg.Function(validInput)
if err == nil {
	logger.Info("✓ Success", "result", result)
	fmt.Println("✅ Valid input succeeded")
}

// Error case (expected)
logger.Info("Testing invalid input (expecting error)...")
result, err = pkg.Function(invalidInput)
if err != nil {
	logger.Warn("✓ Got expected error", "error", err)
	fmt.Println("⚠️  Invalid input correctly rejected")
}
```

### 4. Resource Cleanup / 리소스 정리

Always show proper cleanup:

항상 적절한 정리를 보여줍니다:

```go
func exampleWithCleanup(logger *logging.Logger) {
	logger.Info("Creating resource...")
	resource, err := createResource()
	if err != nil {
		logger.Error("Failed to create resource", "error", err)
		return
	}

	// Ensure cleanup happens
	// 정리가 발생하도록 보장
	defer func() {
		logger.Info("Cleaning up resource...")
		if err := resource.Close(); err != nil {
			logger.Error("Cleanup failed", "error", err)
		} else {
			logger.Info("✓ Resource cleaned up successfully")
		}
	}()

	// Use resource...
	logger.Info("Using resource...")
}
```

---

## Checklist for New Examples / 새 예제 체크리스트

When creating examples for a new package:

새 패키지의 예제를 만들 때:

### Pre-Development / 개발 전

- [ ] Read the package README.md
- [ ] List all public functions/methods
- [ ] Identify common use cases
- [ ] Research production patterns
- [ ] Review existing examples in other packages

### During Development / 개발 중

- [ ] Create `examples/{package}/` directory
- [ ] Ensure shared `logs/` directory exists at repository root / 저장소 루트의 공용 `logs/` 디렉토리 확인
- [ ] Implement log backup system
- [ ] Create example for each function
- [ ] Add combination examples
- [ ] Add real-world scenarios
- [ ] Add edge case handling
- [ ] Add performance examples
- [ ] Add integration examples

### Documentation / 문서화

- [ ] Add bilingual function comments
- [ ] Add bilingual inline comments
- [ ] Add detailed logs at each step
- [ ] Use Unicode symbols in console output
- [ ] Document all parameters
- [ ] Document all return values
- [ ] Show error handling

### Testing / 테스트

- [ ] Run all examples successfully
- [ ] Verify log file creation
- [ ] Verify backup log creation
- [ ] Verify backup cleanup (keep 5)
- [ ] Test with empty input
- [ ] Test with invalid input
- [ ] Test with large input
- [ ] Test error scenarios

### Code Review / 코드 리뷰

- [ ] All comments are bilingual
- [ ] Logs are extremely detailed
- [ ] Console output uses symbols
- [ ] No hardcoded values (use constants)
- [ ] Proper error handling
- [ ] Resource cleanup with defer
- [ ] Performance timing included
- [ ] Example categories covered

### Final Checks / 최종 확인

- [ ] Code compiles without warnings
- [ ] All examples run successfully
- [ ] Logs directory structure correct (`go-utils/logs/`) / 로그 디렉토리 구조 확인 (`go-utils/logs/`)
- [ ] Old backups are deleted
- [ ] README.md updated (if needed)
- [ ] CHANGELOG updated
- [ ] Git commit with proper message

---

## Conclusion / 결론

This guide ensures that all example code in the go-utils project:

이 가이드는 go-utils 프로젝트의 모든 예제 코드가 다음을 보장합니다:

1. **Is comprehensive** - covers all functionality
2. **Is practical** - shows real-world usage
3. **Is well-documented** - logs tell the complete story
4. **Is bilingual** - accessible to English and Korean speakers
5. **Is maintainable** - follows consistent patterns

By following this guide, example code becomes a powerful learning tool that can stand alone as documentation.

이 가이드를 따르면 예제 코드는 문서로서 독립적으로 사용될 수 있는 강력한 학습 도구가 됩니다.

---

**Last Updated / 마지막 업데이트**: 2025-10-16
**Version / 버전**: v1.0
