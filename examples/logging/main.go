package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/arkd0ng/go-utils/fileutil"
	"github.com/arkd0ng/go-utils/logging"
)

// backupLogFile handles log file backup and cleanup for any log file
// backupLogFile은 모든 로그 파일의 백업 및 정리를 처리합니다
func backupLogFile(logFilePath string) {
	if !fileutil.Exists(logFilePath) {
		return
	}

	// Get modification time of existing log file / 기존 로그 파일의 수정 시간 가져오기
	modTime, err := fileutil.ModTime(logFilePath)
	if err != nil {
		return
	}

	// Extract directory and filename / 디렉토리와 파일명 추출
	dir := filepath.Dir(logFilePath)
	filename := filepath.Base(logFilePath)
	ext := filepath.Ext(filename)
	nameWithoutExt := filename[:len(filename)-len(ext)]

	// Create backup filename with timestamp / 타임스탬프와 함께 백업 파일명 생성
	backupName := fmt.Sprintf("%s/%s-%s%s", dir, nameWithoutExt, modTime.Format("20060102-150405"), ext)

	// Backup existing log file / 기존 로그 파일 백업
	if err := fileutil.CopyFile(logFilePath, backupName); err == nil {
		fmt.Printf("✅ Backed up previous log to: %s\n", backupName)
	}

	// Cleanup old backup files - keep only 5 most recent / 오래된 백업 파일 정리 - 최근 5개만 유지
	backupPattern := fmt.Sprintf("%s/%s-*%s", dir, nameWithoutExt, ext)
	backupFiles, err := filepath.Glob(backupPattern)
	if err != nil || len(backupFiles) <= 5 {
		return
	}

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

func main() {
	// Backup all log files that will be used in examples / 예제에서 사용할 모든 로그 파일 백업
	logFiles := []string{
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

	// Example 1: Default logger (가장 간단한 사용)
	// Example 1: Default logger (simplest usage)
	defaultExample()

	// Example 2: Custom logger with options (옵션을 사용한 커스텀 로거)
	// Example 2: Custom logger with options
	customExample()

	// Example 3: Multiple loggers (여러 로거 사용)
	// Example 3: Multiple loggers
	multipleLoggersExample()

	// Example 4: Different log levels (다양한 로그 레벨)
	// Example 4: Different log levels
	logLevelsExample()

	// Example 5: Structured logging (구조화된 로깅)
	// Example 5: Structured logging
	structuredLoggingExample()

	// Example 6: Auto banner examples (자동 배너 예제)
	// Example 6: Auto banner examples
	autoBannerExample()

	// Example 7: Manual banner examples (수동 배너 예제)
	// Example 7: Manual banner examples
	bannerExample()
}

// defaultExample demonstrates the simplest usage with default settings
// defaultExample은 기본 설정으로 가장 간단한 사용법을 보여줍니다
func defaultExample() {
	logger := logging.Default()
	defer logger.Close()

	logger.Banner("Default Logger Example", "v1.0.0")
	logger.Info("Using default logger")
	logger.Info("Logs to ./logs/app.log by default")
	logger.SeparatorLine("=", 50)
}

// customExample demonstrates creating a custom logger with various options
// customExample은 다양한 옵션으로 커스텀 로거를 생성하는 것을 보여줍니다
func customExample() {
	logger, err := logging.New(
		logging.WithFilePath("logs/logging-example-custom.log"),
		logging.WithMaxSize(50),      // 50 MB
		logging.WithMaxBackups(5),    // Keep 5 backups / 5개 백업 유지
		logging.WithMaxAge(7),         // Keep for 7 days / 7일 동안 보관
		logging.WithLevel(logging.DEBUG),
		logging.WithPrefix("[CUSTOM]"),
	)
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	logger.SimpleBanner("Custom Logger", "v1.0.0")
	logger.Debug("This is a debug message")
	logger.Info("Custom logger with specific settings")
	logger.Warn("Log files will rotate at 50MB")
	logger.SeparatorLine("-", 50)
}

// multipleLoggersExample demonstrates using multiple independent loggers
// multipleLoggersExample은 여러 독립적인 로거 사용을 보여줍니다
func multipleLoggersExample() {
	// Application logger / 애플리케이션 로거
	appLogger, _ := logging.New(
		logging.WithFilePath("logs/logging-example-app.log"),
		logging.WithPrefix("[APP]"),
	)
	defer appLogger.Close()

	// Database logger / 데이터베이스 로거
	dbLogger, _ := logging.New(
		logging.WithFilePath("logs/logging-example-database.log"),
		logging.WithPrefix("[DB]"),
		logging.WithLevel(logging.DEBUG),
	)
	defer dbLogger.Close()

	// API logger / API 로거
	apiLogger, _ := logging.New(
		logging.WithFilePath("logs/logging-example-api.log"),
		logging.WithPrefix("[API]"),
	)
	defer apiLogger.Close()

	appLogger.DoubleBanner("Multi-Logger Example", "v1.0.0", "Demonstrating multiple loggers")

	// Each logger writes to its own file / 각 로거는 자체 파일에 작성
	appLogger.Info("Application started")
	dbLogger.Debug("Connecting to database")
	dbLogger.Info("Database connection established")
	apiLogger.Info("API server listening on port 8080")

	appLogger.SeparatorLine("=", 50)
}

// logLevelsExample demonstrates different log levels
// logLevelsExample은 다양한 로그 레벨을 보여줍니다
func logLevelsExample() {
	logger, _ := logging.New(
		logging.WithFilePath("logs/logging-example-levels.log"),
		logging.WithLevel(logging.DEBUG), // Show all levels / 모든 레벨 표시
	)
	defer logger.Close()

	logger.Banner("Log Levels Example", "v1.0.0")

	// All log levels / 모든 로그 레벨
	logger.Debug("This is a DEBUG message - detailed information for debugging")
	logger.Info("This is an INFO message - general informational messages")
	logger.Warn("This is a WARN message - warning that doesn't prevent operation")
	logger.Error("This is an ERROR message - error that should be investigated")

	// Change log level dynamically / 동적으로 로그 레벨 변경
	logger.Info("Changing log level to WARN...")
	logger.SetLevel(logging.WARN)

	// These won't be logged / 이것들은 로깅되지 않음
	logger.Debug("This DEBUG won't be logged")
	logger.Info("This INFO won't be logged")

	// These will be logged / 이것들은 로깅됨
	logger.Warn("This WARN will be logged")
	logger.Error("This ERROR will be logged")

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

	logger.SimpleBanner("Structured Logging", "v1.0.0")

	// User login event / 사용자 로그인 이벤트
	logger.Info("User login successful",
		"user_id", 12345,
		"username", "john.doe",
		"ip", "192.168.1.100",
		"timestamp", time.Now().Unix(),
	)

	// API request / API 요청
	start := time.Now()
	time.Sleep(50 * time.Millisecond) // Simulate processing / 처리 시뮬레이션
	logger.Info("API request completed",
		"method", "GET",
		"path", "/api/users/12345",
		"status", 200,
		"duration_ms", time.Since(start).Milliseconds(),
	)

	// Database query / 데이터베이스 쿼리
	logger.Debug("Database query executed",
		"query", "SELECT * FROM users WHERE id = ?",
		"params", 12345,
		"rows_affected", 1,
		"duration_ms", 15,
	)

	// Error with context / 컨텍스트가 있는 에러
	logger.Error("Failed to process payment",
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
	// Example 6-1: Auto banner with default settings (기본 설정으로 자동 배너)
	// By default, a banner is automatically printed when logger is created
	// 기본적으로 로거 생성 시 자동으로 배너가 출력됩니다
	logger1, _ := logging.New(
		logging.WithFilePath("logs/logging-example-auto_banner_default.log"),
	)
	defer logger1.Close()

	logger1.Info("Logger created with auto banner (default app name and version)")
	logger1.Info("Banner: 'Application v1.0.0' was printed automatically")
	logger1.SeparatorLine("-", 50)

	// Example 6-2: Auto banner with custom app name and version (커스텀 앱 이름/버전)
	// You can customize the app name and version for the auto banner
	// 자동 배너의 앱 이름과 버전을 커스터마이즈할 수 있습니다
	logger2, _ := logging.New(
		logging.WithFilePath("logs/logging-example-auto_banner_custom.log"),
		logging.WithAppName("MyApp"),
		logging.WithAppVersion("v2.0.0"),
	)
	defer logger2.Close()

	logger2.Info("Logger created with custom app name and version")
	logger2.Info("Banner: 'MyApp v2.0.0' was printed automatically")
	logger2.SeparatorLine("-", 50)

	// Example 6-3: Convenience function WithBanner (편의 함수)
	// Use WithBanner() to set both name and version at once
	// WithBanner()를 사용하여 이름과 버전을 한 번에 설정할 수 있습니다
	logger3, _ := logging.New(
		logging.WithFilePath("logs/logging-example-auto_banner_convenience.log"),
		logging.WithBanner("ProductionAPI", "v3.2.1"),
	)
	defer logger3.Close()

	logger3.Info("Logger created with WithBanner convenience function")
	logger3.Info("Banner: 'ProductionAPI v3.2.1' was printed automatically")
	logger3.SeparatorLine("-", 50)

	// Example 6-4: Disable auto banner (자동 배너 비활성화)
	// If you don't want auto banner, disable it explicitly
	// 자동 배너를 원하지 않으면 명시적으로 비활성화할 수 있습니다
	logger4, _ := logging.New(
		logging.WithFilePath("logs/logging-example-auto_banner_disabled.log"),
		logging.WithAutoBanner(false),
	)
	defer logger4.Close()

	logger4.Info("Logger created with auto banner disabled")
	logger4.Info("No automatic banner was printed")
	logger4.SeparatorLine("-", 50)

	// Example 6-5: Disable auto banner but use manual banner (수동 배너 사용)
	// You can disable auto banner and call Banner() manually when needed
	// 자동 배너를 비활성화하고 필요할 때 수동으로 배너를 호출할 수 있습니다
	logger5, _ := logging.New(
		logging.WithFilePath("logs/logging-example-manual_banner_only.log"),
		logging.WithAutoBanner(false),
	)
	defer logger5.Close()

	logger5.Info("Starting application...")
	logger5.Banner("Manual Banner Example", "v1.5.0")
	logger5.Info("Manual banner called when needed")

	logger5.SeparatorLine("=", 50)
}

// bannerExample demonstrates various banner styles
// bannerExample은 다양한 배너 스타일을 보여줍니다
func bannerExample() {
	logger, _ := logging.New(
		logging.WithFilePath("logs/logging-example-banners.log"),
	)
	defer logger.Close()

	// Standard banner / 표준 배너
	logger.Banner("My Application", "v1.0.0")

	logger.Info("This is a standard banner with border")
	logger.SeparatorLine("-", 50)

	// Simple banner / 간단한 배너
	logger.SimpleBanner("Simple Style", "v2.0.0")

	logger.Info("This is a simple banner with lines")
	logger.SeparatorLine("-", 50)

	// Double banner with description / 설명이 있는 이중 배너
	logger.DoubleBanner("Production Server", "v3.0.0", "North America Region")

	logger.Info("This is a double banner with description")
	logger.SeparatorLine("-", 50)

	// Custom ASCII art banner / 커스텀 ASCII 아트 배너
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

	logger.Info("This is a custom ASCII art banner")

	// Various separator styles / 다양한 구분선 스타일
	logger.SeparatorLine("=", 60)
	logger.SeparatorLine("-", 60)
	logger.SeparatorLine("*", 60)
	logger.SeparatorLine("#", 60)
}
