package main

import (
	"time"

	"github.com/arkd0ng/go-utils/logging"
)

func main() {
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

	// Example 6: Banner examples (배너 예제)
	// Example 6: Banner examples
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
		logging.WithFilePath("./logs/custom.log"),
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
		logging.WithFilePath("./logs/app.log"),
		logging.WithPrefix("[APP]"),
	)
	defer appLogger.Close()

	// Database logger / 데이터베이스 로거
	dbLogger, _ := logging.New(
		logging.WithFilePath("./logs/database.log"),
		logging.WithPrefix("[DB]"),
		logging.WithLevel(logging.DEBUG),
	)
	defer dbLogger.Close()

	// API logger / API 로거
	apiLogger, _ := logging.New(
		logging.WithFilePath("./logs/api.log"),
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
		logging.WithFilePath("./logs/levels.log"),
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
		logging.WithFilePath("./logs/structured.log"),
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

// bannerExample demonstrates various banner styles
// bannerExample은 다양한 배너 스타일을 보여줍니다
func bannerExample() {
	logger, _ := logging.New(
		logging.WithFilePath("./logs/banners.log"),
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
