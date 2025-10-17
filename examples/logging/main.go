package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/arkd0ng/go-utils/fileutil"
	"github.com/arkd0ng/go-utils/logging"
)

// displayHeader shows comprehensive package information for the logging examples suite
// displayHeader는 로깅 예제 모음에 대한 종합 패키지 정보를 자세하게 안내합니다
func displayHeader() {
	fmt.Println("╔════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║            Logging Package - Comprehensive Examples                        ║")
	fmt.Println("║            Logging 패키지 - 종합 예제                                        ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════╝")
	fmt.Println()

	fmt.Println("📋 Package Information / 패키지 정보")
	fmt.Println("   Package: github.com/arkd0ng/go-utils/logging / 패키지: github.com/arkd0ng/go-utils/logging")
	fmt.Println("   Description: Structured logging with file rotation / 설명: 파일 로테이션을 지원하는 구조화된 로깅")
	fmt.Println("   Total Examples: 7 comprehensive examples / 총 예제 수: 7개의 종합 예제")
	fmt.Println("   Zero Dependencies: lumberjack.v2 for rotation only / 최소 의존성: 파일 로테이션을 위한 lumberjack.v2만 사용")
	fmt.Println()

	fmt.Println("🌟 Key Features / 주요 기능")
	fmt.Println("   • Structured Logging: Key-value pairs for context-rich logs / 구조화된 로깅: 문맥이 풍부한 로그를 위한 키-값 쌍")
	fmt.Println("   • File Rotation: Automatic rotation by size, age, and backup count / 파일 로테이션: 크기, 기간, 백업 개수에 따른 자동 회전")
	fmt.Println("   • Multiple Log Levels: DEBUG, INFO, WARN, ERROR, FATAL / 다양한 로그 레벨: DEBUG, INFO, WARN, ERROR, FATAL")
	fmt.Println("   • Color Output: Beautiful colored console output / 색상 출력: 보기 좋은 컬러 콘솔 출력")
	fmt.Println("   • Banner Support: Auto and manual banner generation / 배너 지원: 자동 및 수동 배너 생성")
	fmt.Println("   • Thread-Safe: Safe for concurrent use with sync.Mutex / 스레드 안전성: sync.Mutex로 동시 사용 시 안전 보장")
	fmt.Println("   • Options Pattern: Flexible configuration with functional options / 옵션 패턴: 함수형 옵션으로 유연한 구성 가능")
	fmt.Println()

	fmt.Println("📚 Examples Covered / 다루는 예제")
	fmt.Println("   1. Default Logger - Simplest usage / 기본 로거 - 가장 간단한 사용법")
	fmt.Println("   2. Custom Logger - With rotation options / 커스텀 로거 - 로테이션 옵션 포함")
	fmt.Println("   3. Multiple Loggers - Independent loggers / 여러 로거 - 독립적인 로거 구성")
	fmt.Println("   4. Log Levels - DEBUG, INFO, WARN, ERROR / 로그 레벨 - DEBUG, INFO, WARN, ERROR")
	fmt.Println("   5. Structured Logging - Key-value pairs / 구조화된 로깅 - 키-값 쌍")
	fmt.Println("   6. Auto Banner - Automatic banners / 자동 배너 - 자동 생성 배너")
	fmt.Println("   7. Manual Banner - Custom banners / 수동 배너 - 커스텀 배너 생성")
	fmt.Println()

	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()
}

// backupLogFile handles log file backup and cleanup for any log file before running examples
// backupLogFile은 예제를 실행하기 전에 모든 로그 파일의 백업과 정리를 세심하게 수행합니다
func backupLogFile(logFilePath string) {
	if !fileutil.Exists(logFilePath) {
		return
	}

	// Get modification time of existing log file
	// 기존 로그 파일의 수정 시간 가져오기
	modTime, err := fileutil.ModTime(logFilePath)
	if err != nil {
		return
	}

	// Extract directory and filename
	// 디렉토리와 파일명 추출
	dir := filepath.Dir(logFilePath)
	filename := filepath.Base(logFilePath)
	ext := filepath.Ext(filename)
	nameWithoutExt := filename[:len(filename)-len(ext)]

	// Create backup filename with timestamp
	// 타임스탬프와 함께 백업 파일명 생성
	backupName := fmt.Sprintf("%s/%s-%s%s", dir, nameWithoutExt, modTime.Format("20060102-150405"), ext)

	// Backup existing log file
	// 기존 로그 파일 백업
	if err := fileutil.CopyFile(logFilePath, backupName); err == nil {
		fmt.Printf("✅ Backed up previous log to: %s\n", backupName)
		// Delete original log file to prevent content duplication
		// 내용 중복 방지를 위해 원본 로그 파일 삭제
		fileutil.DeleteFile(logFilePath)
	}

	// Cleanup old backup files - keep only 5 most recent
	// 오래된 백업 파일 정리 - 최근 5개만 유지
	backupPattern := fmt.Sprintf("%s/%s-*%s", dir, nameWithoutExt, ext)
	backupFiles, err := filepath.Glob(backupPattern)
	if err != nil || len(backupFiles) <= 5 {
		return
	}

	// Sort by modification time
	// 수정 시간으로 정렬
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

	// Sort oldest first
	// 가장 오래된 것부터 정렬
	for i := 0; i < len(files)-1; i++ {
		for j := i + 1; j < len(files); j++ {
			if files[i].modTime.After(files[j].modTime) {
				files[i], files[j] = files[j], files[i]
			}
		}
	}

	// Delete oldest files to keep only 5
	// 5개만 유지하도록 가장 오래된 파일 삭제
	for i := 0; i < len(files)-5; i++ {
		fileutil.DeleteFile(files[i].path)
		fmt.Printf("🗑️  Deleted old backup: %s\n", files[i].path)
	}
}

func main() {
	// Backup all log files that will be used in examples
	// 예제에서 사용할 모든 로그 파일 백업
	logFiles := []string{
		"logs/logging-example-default.log",
		"logs/logging-example-app.log",
		"logs/logging-example-custom.log",
		"logs/logging-example-database.log",
		"logs/logging-example-api.log",
		"logs/logging-example-levels.log",
		"logs/logging-example-structured.log",
		"logs/logging-example-auto_banner_default.log",
		"logs/logging-example-auto_banner_custom.log",
		"logs/logging-example-auto_banner_convenience.log",
		"logs/logging-example-auto_banner_disabled.log",
		"logs/logging-example-manual_banner_only.log",
		"logs/logging-example-banners.log",
	}

	fmt.Println("🔄 Checking and backing up existing log files...")
	fmt.Println("🔄 기존 로그 파일 확인 및 백업 중...")
	for _, logFile := range logFiles {
		backupLogFile(logFile)
	}
	fmt.Println()

	// Display comprehensive header before running the examples
	// 예제를 실행하기 전에 종합 헤더를 출력합니다
	displayHeader()

	// Example 1: Default logger - Simplest usage
	// 예제 1: 기본 로거 - 가장 간단한 사용법
	defaultExample()

	// Example 2: Custom logger with options
	// 예제 2: 옵션을 사용한 커스텀 로거
	customExample()

	// Example 3: Multiple loggers
	// 예제 3: 여러 로거 구성
	multipleLoggersExample()

	// Example 4: Different log levels
	// 예제 4: 다양한 로그 레벨
	logLevelsExample()

	// Example 5: Structured logging
	// 예제 5: 구조화된 로깅
	structuredLoggingExample()

	// Example 6: Auto banner examples
	// 예제 6: 자동 배너 사용 예제
	autoBannerExample()

	// Example 7: Manual banner examples
	// 예제 7: 수동 배너 사용 예제
	bannerExample()
}

// defaultExample demonstrates the simplest usage with default settings, highlighting the zero-configuration experience
// defaultExample은 설정 없이도 바로 사용할 수 있는 기본 사용 예제를 자세히 보여주며 초기 경험을 강조합니다
func defaultExample() {
	logger, _ := logging.New(
		logging.WithFilePath("logs/logging-example-default.log"),
	)
	defer logger.Close()

	logger.Banner("Default Logger Example", "v1.0.0")
	logger.Info("Using logger with default settings / 기본 설정으로 로거 사용 중")
	logger.Info("Logs to ./logs/logging-example-default.log / 로그 파일 경로: ./logs/logging-example-default.log")
	logger.SeparatorLine("=", 50)
}

// customExample demonstrates creating a custom logger with rotation, retention, and level options
// customExample은 로테이션, 보존 기간, 로그 레벨 옵션을 포함한 커스텀 로거 구성을 자세히 설명합니다
func customExample() {
	logger, err := logging.New(
		logging.WithFilePath("logs/logging-example-custom.log"),
		logging.WithMaxSize(50), // 50 MB
		// Keep 5 backups
		// 5개 백업 유지
		logging.WithMaxBackups(5),
		// Keep for 7 days
		// 7일 동안 보관
		logging.WithMaxAge(7),
		logging.WithLevel(logging.DEBUG),
		logging.WithPrefix("[CUSTOM]"),
	)
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	logger.SimpleBanner("Custom Logger", "v1.0.0")
	logger.Debug("This is a debug message / 디버그 메시지 예시입니다")
	logger.Info("Custom logger with specific settings / 커스텀 설정이 적용된 로거입니다")
	logger.Warn("Log files will rotate at 50MB / 로그 파일이 50MB에 도달하면 로테이션됩니다")
	logger.SeparatorLine("-", 50)
}

// multipleLoggersExample demonstrates using multiple independent loggers for different subsystems
// multipleLoggersExample은 서로 다른 하위 시스템을 위한 독립적인 로거 구성을 자세히 보여줍니다
func multipleLoggersExample() {
	// Application logger
	// 애플리케이션 로거
	appLogger, _ := logging.New(
		logging.WithFilePath("logs/logging-example-app.log"),
		logging.WithPrefix("[APP]"),
	)
	defer appLogger.Close()

	// Database logger
	// 데이터베이스 로거
	dbLogger, _ := logging.New(
		logging.WithFilePath("logs/logging-example-database.log"),
		logging.WithPrefix("[DB]"),
		logging.WithLevel(logging.DEBUG),
	)
	defer dbLogger.Close()

	// API logger
	// API 로거
	apiLogger, _ := logging.New(
		logging.WithFilePath("logs/logging-example-api.log"),
		logging.WithPrefix("[API]"),
	)
	defer apiLogger.Close()

	appLogger.DoubleBanner("Multi-Logger Example / 다중 로거 예제", "v1.0.0", "Demonstrating multiple loggers / 여러 로거 사용을 시연합니다")

	// Each logger writes to its own file
	// 각 로거는 자체 파일에 작성
	appLogger.Info("Application started / 애플리케이션이 시작되었습니다")
	dbLogger.Debug("Connecting to database / 데이터베이스에 연결을 시도하는 중입니다")
	dbLogger.Info("Database connection established / 데이터베이스 연결이 완료되었습니다")
	apiLogger.Info("API server listening on port 8080 / API 서버가 포트 8080에서 대기 중입니다")

	appLogger.SeparatorLine("=", 50)
}

// logLevelsExample demonstrates different log levels
// logLevelsExample은 다양한 로그 레벨을 보여줍니다
func logLevelsExample() {
	logger, _ := logging.New(
		logging.WithFilePath("logs/logging-example-levels.log"),
		// Show all levels
		// 모든 레벨 표시
		logging.WithLevel(logging.DEBUG),
	)
	defer logger.Close()

	logger.Banner("Log Levels Example / 로그 레벨 예제", "v1.0.0")

	// All log levels
	// 모든 로그 레벨
	logger.Debug("This is a DEBUG message - detailed information for debugging / DEBUG 메시지 - 디버깅에 필요한 상세 정보")
	logger.Info("This is an INFO message - general informational messages / INFO 메시지 - 일반적인 안내 메시지")
	logger.Warn("This is a WARN message - warning that doesn't prevent operation / WARN 메시지 - 동작을 막지는 않지만 주의가 필요한 경고")
	logger.Error("This is an ERROR message - error that should be investigated / ERROR 메시지 - 조사해야 하는 오류")

	// Change log level dynamically
	// 동적으로 로그 레벨 변경
	logger.Info("Changing log level to WARN... / 로그 레벨을 WARN으로 변경합니다...")
	logger.SetLevel(logging.WARN)

	// These won't be logged
	// 이것들은 로깅되지 않음
	logger.Debug("This DEBUG won't be logged / 이 DEBUG 메시지는 기록되지 않습니다")
	logger.Info("This INFO won't be logged / 이 INFO 메시지는 기록되지 않습니다")

	// These will be logged
	// 이것들은 로깅됨
	logger.Warn("This WARN will be logged / 이 WARN 메시지는 기록됩니다")
	logger.Error("This ERROR will be logged / 이 ERROR 메시지는 기록됩니다")

	logger.SeparatorLine("-", 50)
}

// structuredLoggingExample demonstrates structured logging with key-value pairs
// structuredLoggingExample은 키-값 쌍을 사용한 구조화된 로깅을 보여줍니다
func structuredLoggingExample() {
	logger, _ := logging.New(
		logging.WithFilePath("logs/logging-example-structured.log"),
		logging.WithPrefix("[STRUCT]"),
	)
	defer logger.Close()

	logger.SimpleBanner("Structured Logging / 구조화된 로깅", "v1.0.0")

	// User login event
	// 사용자 로그인 이벤트
	logger.Info("User login successful / 사용자 로그인 성공",
		"user_id", 12345,
		"username", "john.doe",
		"ip", "192.168.1.100",
		"timestamp", time.Now().Unix(),
	)

	// API request
	// API 요청
	start := time.Now()
	// Simulate processing
	// 처리 시뮬레이션
	time.Sleep(50 * time.Millisecond)
	logger.Info("API request completed / API 요청이 완료되었습니다",
		"method", "GET",
		"path", "/api/users/12345",
		"status", 200,
		"duration_ms", time.Since(start).Milliseconds(),
	)

	// Database query
	// 데이터베이스 쿼리
	logger.Debug("Database query executed / 데이터베이스 쿼리가 실행되었습니다",
		"query", "SELECT * FROM users WHERE id = ?",
		"params", 12345,
		"rows_affected", 1,
		"duration_ms", 15,
	)

	// Error with context
	// 컨텍스트가 있는 에러
	logger.Error("Failed to process payment / 결제 처리에 실패했습니다",
		"order_id", "ORD-2024-001",
		"amount", 99.99,
		"currency", "USD",
		"error", "insufficient funds",
		"retry_count", 3,
	)

	logger.SeparatorLine("=", 50)
}

// autoBannerExample demonstrates automatic banner functionality
// autoBannerExample은 자동 배너 기능을 보여줍니다
func autoBannerExample() {
	// Example 6-1: Auto banner with default settings
	// 예제 6-1: 기본 설정으로 자동 배너 활용
	// By default, a banner is automatically printed when logger is created
	// 기본적으로 로거를 생성하면 배너가 자동으로 출력됩니다
	logger1, _ := logging.New(
		logging.WithFilePath("logs/logging-example-auto_banner_default.log"),
	)
	defer logger1.Close()

	logger1.Info("Logger created with auto banner (default app name and version) / 자동 배너가 기본 앱 이름과 버전으로 설정된 로거가 생성되었습니다")
	logger1.Info("Banner: 'Application v1.0.0' was printed automatically / 자동으로 'Application v1.0.0' 배너가 출력되었습니다")
	logger1.SeparatorLine("-", 50)

	// Example 6-2: Auto banner with custom app name and version
	// 예제 6-2: 사용자 지정 앱 이름과 버전으로 자동 배너 사용
	// You can customize the app name and version for the auto banner
	// 자동 배너에 표시될 앱 이름과 버전을 자유롭게 지정할 수 있습니다
	logger2, _ := logging.New(
		logging.WithFilePath("logs/logging-example-auto_banner_custom.log"),
		logging.WithAppName("MyApp"),
		logging.WithAppVersion("v2.0.0"),
	)
	defer logger2.Close()

	logger2.Info("Logger created with custom app name and version / 커스텀 앱 이름과 버전이 적용된 로거가 생성되었습니다")
	logger2.Info("Banner: 'MyApp v2.0.0' was printed automatically / 자동으로 'MyApp v2.0.0' 배너가 출력되었습니다")
	logger2.SeparatorLine("-", 50)

	// Example 6-3: Convenience function WithBanner
	// 예제 6-3: WithBanner 편의 함수 사용
	// Use WithBanner() to set both name and version at once
	// WithBanner()로 이름과 버전을 한 번에 설정할 수 있습니다
	logger3, _ := logging.New(
		logging.WithFilePath("logs/logging-example-auto_banner_convenience.log"),
		logging.WithBanner("ProductionAPI", "v3.2.1"),
	)
	defer logger3.Close()

	logger3.Info("Logger created with WithBanner convenience function / WithBanner 편의 함수로 생성된 로거입니다")
	logger3.Info("Banner: 'ProductionAPI v3.2.1' was printed automatically / 자동으로 'ProductionAPI v3.2.1' 배너가 출력되었습니다")
	logger3.SeparatorLine("-", 50)

	// Example 6-4: Disable auto banner (자동 배너 비활성화)
	// Example 6-4: Explicitly disable auto banner
	// 예제 6-4: 자동 배너를 명시적으로 비활성화
	// If you don't want auto banner, disable it explicitly
	// 자동 배너가 필요 없다면 옵션을 통해 명시적으로 끌 수 있습니다
	logger4, _ := logging.New(
		logging.WithFilePath("logs/logging-example-auto_banner_disabled.log"),
		logging.WithAutoBanner(false),
	)
	defer logger4.Close()

	logger4.Info("Logger created with auto banner disabled / 자동 배너가 비활성화된 상태로 로거가 생성되었습니다")
	logger4.Info("No automatic banner was printed / 자동으로 출력된 배너가 없습니다")
	logger4.SeparatorLine("-", 50)

	// Example 6-5: Disable auto banner but use manual banner
	// 예제 6-5: 자동 배너를 끄고 수동 배너 사용
	// You can disable auto banner and call Banner() manually when needed
	// 자동 배너를 끈 뒤 필요할 때 Banner()를 호출해 배너를 출력할 수 있습니다
	logger5, _ := logging.New(
		logging.WithFilePath("logs/logging-example-manual_banner_only.log"),
		logging.WithAutoBanner(false),
	)
	defer logger5.Close()

	logger5.Info("Starting application... / 애플리케이션을 시작합니다...")
	logger5.Banner("Manual Banner Example / 수동 배너 예제", "v1.5.0")
	logger5.Info("Manual banner called when needed / 필요한 시점에 수동 배너를 호출했습니다")

	logger5.SeparatorLine("=", 50)
}

// bannerExample demonstrates various banner styles to inspire customization
// bannerExample은 활용 가능한 다양한 배너 스타일을 소개하여 사용자 정의 아이디어를 제공합니다
func bannerExample() {
	logger, _ := logging.New(
		logging.WithFilePath("logs/logging-example-banners.log"),
	)
	defer logger.Close()

	// Standard banner
	// 표준 배너
	logger.Banner("My Application / 나의 애플리케이션", "v1.0.0")

	logger.Info("This is a standard banner with border / 테두리가 포함된 기본 배너 스타일입니다")
	logger.SeparatorLine("-", 50)

	// Simple banner
	// 간단한 배너
	logger.SimpleBanner("Simple Style / 심플 스타일", "v2.0.0")

	logger.Info("This is a simple banner with lines / 선으로 구성된 간단한 배너입니다")
	logger.SeparatorLine("-", 50)

	// Double banner with description
	// 설명이 있는 이중 배너
	logger.DoubleBanner("Production Server / 프로덕션 서버", "v3.0.0", "North America Region / 북미 리전")

	logger.Info("This is a double banner with description / 설명 문구가 포함된 이중 배너입니다")
	logger.SeparatorLine("-", 50)

	// Custom ASCII art banner
	// 커스텀 ASCII 아트 배너
	logger.CustomBanner([]string{
		"",
		"  ╔═╗╔═╗  ╦ ╦╔╦╗╦╦  ╔═╗",
		"  ║ ╦║ ║  ║ ║ ║ ║║  ╚═╗",
		"  ╚═╝╚═╝  ╚═╝ ╩ ╩╩═╝╚═╝",
		"",
		"  Logging Utility Package",
		"  Version 1.0.0",
		"",
	})

	logger.Info("This is a custom ASCII art banner / 커스텀 ASCII 아트 배너입니다")

	// Various separator styles
	// 다양한 구분선 스타일
	logger.SeparatorLine("=", 60)
	logger.SeparatorLine("-", 60)
	logger.SeparatorLine("*", 60)
	logger.SeparatorLine("#", 60)
}
