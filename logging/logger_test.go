package logging

import (
	"os"
	"strings"
	"testing"
	"time"
)

// TestNew tests creating a new logger with various options
// TestNew는 다양한 옵션으로 새 로거를 생성하는 테스트입니다
func TestNew(t *testing.T) {
	tests := []struct {
		name string
		opts []Option
	}{
		{
			name: "Default logger",
			opts: []Option{},
		},
		{
			name: "Custom file path",
			opts: []Option{
				WithFilePath("./test_logs/custom.log"),
			},
		},
		{
			name: "Custom settings",
			opts: []Option{
				WithFilePath("./test_logs/test.log"),
				WithMaxSize(50),
				WithMaxBackups(5),
				WithMaxAge(14),
				WithLevel(DEBUG),
			},
		},
		{
			name: "Stdout only",
			opts: []Option{
				WithStdoutOnly(),
			},
		},
		{
			name: "File only",
			opts: []Option{
				WithFilePath("./test_logs/fileonly.log"),
				WithFileOnly(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger, err := New(tt.opts...)
			if err != nil {
				t.Fatalf("Failed to create logger: %v", err)
			}
			defer logger.Close()
			defer cleanupTestLogs()

			if logger == nil {
				t.Error("Logger should not be nil")
			}
		})
	}
}

// TestDefault tests the Default logger
// TestDefault는 기본 로거를 테스트합니다
func TestDefault(t *testing.T) {
	logger := Default()
	defer logger.Close()
	defer cleanupTestLogs()

	if logger == nil {
		t.Error("Default logger should not be nil")
	}

	if logger.config.level != INFO {
		t.Errorf("Default level should be INFO, got %v", logger.config.level)
	}
}

// TestLogLevels tests different log levels
// TestLogLevels는 다양한 로그 레벨을 테스트합니다
func TestLogLevels(t *testing.T) {
	logger, err := New(
		WithFilePath("./test_logs/levels.log"),
		WithLevel(DEBUG),
	)
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Close()
	defer cleanupTestLogs()

	// Test all log levels / 모든 로그 레벨 테스트
	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message")

	// Verify log file exists / 로그 파일이 존재하는지 확인
	if _, err := os.Stat("./test_logs/levels.log"); os.IsNotExist(err) {
		t.Error("Log file should exist")
	}
}

// TestSetLevel tests changing log level
// TestSetLevel은 로그 레벨 변경을 테스트합니다
func TestSetLevel(t *testing.T) {
	logger, err := New(
		WithFilePath("./test_logs/setlevel.log"),
		WithLevel(INFO),
	)
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Close()
	defer cleanupTestLogs()

	// Initial level / 초기 레벨
	if logger.GetLevel() != INFO {
		t.Errorf("Initial level should be INFO, got %v", logger.GetLevel())
	}

	// Change level / 레벨 변경
	logger.SetLevel(DEBUG)
	if logger.GetLevel() != DEBUG {
		t.Errorf("Level should be DEBUG after change, got %v", logger.GetLevel())
	}
}

// TestStructuredLogging tests structured logging with key-value pairs
// TestStructuredLogging은 키-값 쌍을 사용한 구조화된 로깅을 테스트합니다
func TestStructuredLogging(t *testing.T) {
	logger, err := New(
		WithFilePath("./test_logs/structured.log"),
	)
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Close()
	defer cleanupTestLogs()

	// Log with structured data / 구조화된 데이터로 로깅
	logger.Info("User login",
		"user_id", 12345,
		"ip", "192.168.1.1",
		"duration", time.Millisecond*120,
	)

	// Read log file / 로그 파일 읽기
	content, err := os.ReadFile("./test_logs/structured.log")
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	// Verify structured data is in the log / 구조화된 데이터가 로그에 있는지 확인
	logStr := string(content)
	if !strings.Contains(logStr, "user_id=12345") {
		t.Error("Log should contain user_id=12345")
	}
	if !strings.Contains(logStr, "ip=192.168.1.1") {
		t.Error("Log should contain ip=192.168.1.1")
	}
}

// TestPrefix tests log prefix functionality
// TestPrefix는 로그 프리픽스 기능을 테스트합니다
func TestPrefix(t *testing.T) {
	logger, err := New(
		WithFilePath("./test_logs/prefix.log"),
		WithPrefix("[TEST]"),
	)
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Close()
	defer cleanupTestLogs()

	logger.Info("Test message")

	// Read log file / 로그 파일 읽기
	content, err := os.ReadFile("./test_logs/prefix.log")
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	// Verify prefix is in the log / 프리픽스가 로그에 있는지 확인
	if !strings.Contains(string(content), "[TEST]") {
		t.Error("Log should contain [TEST] prefix")
	}
}

// TestRotate tests manual log rotation
// TestRotate는 수동 로그 로테이션을 테스트합니다
func TestRotate(t *testing.T) {
	logger, err := New(
		WithFilePath("./test_logs/rotate.log"),
	)
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Close()
	defer cleanupTestLogs()

	logger.Info("Before rotation")

	// Manually trigger rotation / 수동으로 로테이션 트리거
	err = logger.Rotate()
	if err != nil {
		t.Errorf("Rotate should not return error: %v", err)
	}

	logger.Info("After rotation")
}

// TestLevelParsing tests parsing level from string
// TestLevelParsing은 문자열에서 레벨 파싱을 테스트합니다
func TestLevelParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected Level
	}{
		{"DEBUG", DEBUG},
		{"debug", DEBUG},
		{"INFO", INFO},
		{"info", INFO},
		{"WARN", WARN},
		{"WARNING", WARN},
		{"ERROR", ERROR},
		{"FATAL", FATAL},
		{"unknown", INFO}, // Default / 기본값
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := ParseLevel(tt.input)
			if result != tt.expected {
				t.Errorf("ParseLevel(%s) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// TestLevelString tests Level.String() method
// TestLevelString은 Level.String() 메서드를 테스트합니다
func TestLevelString(t *testing.T) {
	tests := []struct {
		level    Level
		expected string
	}{
		{DEBUG, "DEBUG"},
		{INFO, "INFO"},
		{WARN, "WARN"},
		{ERROR, "ERROR"},
		{FATAL, "FATAL"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := tt.level.String()
			if result != tt.expected {
				t.Errorf("Level.String() = %s, want %s", result, tt.expected)
			}
		})
	}
}

// TestBanner tests banner functionality
// TestBanner는 배너 기능을 테스트합니다
func TestBanner(t *testing.T) {
	logger, err := New(
		WithFilePath("./test_logs/banner.log"),
		WithAutoBanner(false), // Disable auto banner for this test / 이 테스트에서는 자동 배너 비활성화
	)
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Close()
	defer cleanupTestLogs()

	// Test different banner types / 다양한 배너 타입 테스트
	logger.Banner("Test App", "v1.0.0")
	logger.SimpleBanner("Simple", "v2.0.0")
	logger.DoubleBanner("Double", "v3.0.0", "Description")
	logger.SeparatorLine("=", 50)
	logger.CustomBanner([]string{"Custom", "Banner"})

	// Verify log file exists / 로그 파일이 존재하는지 확인
	if _, err := os.Stat("./test_logs/banner.log"); os.IsNotExist(err) {
		t.Error("Log file should exist")
	}
}

// TestAutoBanner tests automatic banner printing on logger creation
// TestAutoBanner는 로거 생성 시 자동 배너 출력을 테스트합니다
func TestAutoBanner(t *testing.T) {
	tests := []struct {
		name           string
		opts           []Option
		shouldHaveBanner bool
	}{
		{
			name: "Auto banner enabled (default)",
			opts: []Option{
				WithFilePath("./test_logs/auto_banner_default.log"),
			},
			shouldHaveBanner: true,
		},
		{
			name: "Auto banner disabled",
			opts: []Option{
				WithFilePath("./test_logs/auto_banner_disabled.log"),
				WithAutoBanner(false),
			},
			shouldHaveBanner: false,
		},
		{
			name: "Custom app name and version",
			opts: []Option{
				WithFilePath("./test_logs/auto_banner_custom.log"),
				WithAppName("TestApp"),
				WithAppVersion("v2.0.0"),
			},
			shouldHaveBanner: true,
		},
		{
			name: "WithBanner convenience function",
			opts: []Option{
				WithFilePath("./test_logs/auto_banner_convenience.log"),
				WithBanner("MyApp", "v3.0.0"),
			},
			shouldHaveBanner: true,
		},
		{
			name: "Auto extract app name from filename",
			opts: []Option{
				WithFilePath("./test_logs/database.log"),
				// With app.yaml present, should use "go-utils" from cfg/app.yaml
				// Without app.yaml, would extract "database" from filename
			},
			shouldHaveBanner: true,
		},
		{
			name: "Auto extract from complex path",
			opts: []Option{
				WithFilePath("./test_logs/api-server.log"),
				// With app.yaml present, should use "go-utils" from cfg/app.yaml
			},
			shouldHaveBanner: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger, err := New(tt.opts...)
			if err != nil {
				t.Fatalf("Failed to create logger: %v", err)
			}
			defer logger.Close()
			defer cleanupTestLogs()

			// Write a test log to ensure file is created / 파일 생성을 위해 테스트 로그 작성
			logger.Info("Test log message")

			// Read log file / 로그 파일 읽기
			content, err := os.ReadFile(logger.config.filename)
			if err != nil {
				t.Fatalf("Failed to read log file: %v", err)
			}

			logStr := string(content)
			hasBanner := strings.Contains(logStr, "╔") || strings.Contains(logStr, "═")

			if tt.shouldHaveBanner && !hasBanner {
				t.Error("Log file should contain auto banner")
			}
			if !tt.shouldHaveBanner && hasBanner {
				t.Error("Log file should not contain auto banner")
			}

			// Verify app name in banner
			// 배너에 앱 이름 확인
			if logger.config.appName != "" && logger.config.appName != "Application" && tt.shouldHaveBanner {
				// Verify config appName if it's not the default "Application"
				// 기본값 "Application"이 아닌 경우 config appName 확인
				if !strings.Contains(logStr, logger.config.appName) {
					t.Errorf("Log file should contain app name: %s", logger.config.appName)
				}
			}
		})
	}
}

// TestFileOnly tests file-only output
// TestFileOnly는 파일 전용 출력을 테스트합니다
func TestFileOnly(t *testing.T) {
	logger, err := New(
		WithFilePath("./test_logs/fileonly.log"),
		WithFileOnly(),
	)
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Close()
	defer cleanupTestLogs()

	logger.Info("File only message")

	// Verify log file exists and contains the message
	// 로그 파일이 존재하고 메시지를 포함하는지 확인
	content, err := os.ReadFile("./test_logs/fileonly.log")
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	if !strings.Contains(string(content), "File only message") {
		t.Error("Log file should contain the message")
	}
}

// TestMultipleLoggers tests creating multiple independent loggers
// TestMultipleLoggers는 여러 독립적인 로거 생성을 테스트합니다
func TestMultipleLoggers(t *testing.T) {
	defer cleanupTestLogs()

	// Create multiple loggers with different configurations
	// 다양한 설정으로 여러 로거 생성
	logger1, err := New(
		WithFilePath("./test_logs/logger1.log"),
		WithPrefix("[APP1]"),
	)
	if err != nil {
		t.Fatalf("Failed to create logger1: %v", err)
	}
	defer logger1.Close()

	logger2, err := New(
		WithFilePath("./test_logs/logger2.log"),
		WithPrefix("[APP2]"),
	)
	if err != nil {
		t.Fatalf("Failed to create logger2: %v", err)
	}
	defer logger2.Close()

	logger3, err := New(
		WithFilePath("./test_logs/logger3.log"),
		WithPrefix("[DB]"),
	)
	if err != nil {
		t.Fatalf("Failed to create logger3: %v", err)
	}
	defer logger3.Close()

	// Log to different loggers / 다른 로거에 로깅
	logger1.Info("Message from logger1")
	logger2.Info("Message from logger2")
	logger3.Info("Message from logger3")

	// Verify each log file exists and contains correct message
	// 각 로그 파일이 존재하고 올바른 메시지를 포함하는지 확인
	verifyLogFile(t, "./test_logs/logger1.log", "[APP1]", "Message from logger1")
	verifyLogFile(t, "./test_logs/logger2.log", "[APP2]", "Message from logger2")
	verifyLogFile(t, "./test_logs/logger3.log", "[DB]", "Message from logger3")
}

// Helper function to verify log file content
// 로그 파일 내용을 확인하는 헬퍼 함수
func verifyLogFile(t *testing.T, filename, prefix, message string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to read %s: %v", filename, err)
	}

	logStr := string(content)
	if !strings.Contains(logStr, prefix) {
		t.Errorf("%s should contain prefix %s", filename, prefix)
	}
	if !strings.Contains(logStr, message) {
		t.Errorf("%s should contain message %s", filename, message)
	}
}

// cleanupTestLogs removes test log files
// cleanupTestLogs는 테스트 로그 파일을 제거합니다
func cleanupTestLogs() {
	os.RemoveAll("./test_logs")
	os.RemoveAll("./logs")
}

// BenchmarkInfo benchmarks Info logging
// BenchmarkInfo는 Info 로깅을 벤치마크합니다
func BenchmarkInfo(b *testing.B) {
	logger, _ := New(WithFilePath("./test_logs/bench.log"))
	defer logger.Close()
	defer cleanupTestLogs()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("Benchmark message", "iteration", i)
	}
}

// BenchmarkStructuredLogging benchmarks structured logging
// BenchmarkStructuredLogging은 구조화된 로깅을 벤치마크합니다
func BenchmarkStructuredLogging(b *testing.B) {
	logger, _ := New(WithFilePath("./test_logs/bench_structured.log"))
	defer logger.Close()
	defer cleanupTestLogs()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("Benchmark message",
			"iteration", i,
			"timestamp", time.Now(),
			"value", float64(i)*1.5,
		)
	}
}

// TestPrintfStyleLogging tests Printf-style formatted logging methods
// TestPrintfStyleLogging은 Printf 스타일 형식화된 로깅 메서드를 테스트합니다
func TestPrintfStyleLogging(t *testing.T) {
	logger, err := New(
		WithFilePath("./test_logs/printf_style.log"),
		WithLevel(DEBUG),
	)
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Close()
	defer cleanupTestLogs()

	// Test all Printf-style methods / 모든 Printf 스타일 메서드 테스트
	user := "john_doe"
	userID := 12345

	logger.Debugf("Debug: User %s with ID %d", user, userID)
	logger.Infof("Info: User %s with ID %d", user, userID)
	logger.Warnf("Warning: User %s with ID %d", user, userID)
	logger.Errorf("Error: User %s with ID %d", user, userID)

	// Read log file and verify content / 로그 파일 읽기 및 내용 확인
	content, err := os.ReadFile("./test_logs/printf_style.log")
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	logStr := string(content)

	// Verify formatted messages are in the log / 형식화된 메시지가 로그에 있는지 확인
	expectedMessages := []string{
		"Debug: User john_doe with ID 12345",
		"Info: User john_doe with ID 12345",
		"Warning: User john_doe with ID 12345",
		"Error: User john_doe with ID 12345",
	}

	for _, expected := range expectedMessages {
		if !strings.Contains(logStr, expected) {
			t.Errorf("Log should contain: %s", expected)
		}
	}

	// Verify log levels are present / 로그 레벨이 있는지 확인
	for _, level := range []string{"[DEBUG]", "[INFO]", "[WARN]", "[ERROR]"} {
		if !strings.Contains(logStr, level) {
			t.Errorf("Log should contain level: %s", level)
		}
	}
}

// TestPrintfVsStructured tests the difference between Printf and structured logging
// TestPrintfVsStructured는 Printf와 구조화된 로깅의 차이를 테스트합니다
func TestPrintfVsStructured(t *testing.T) {
	logger, err := New(
		WithFilePath("./test_logs/printf_vs_structured.log"),
	)
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Close()
	defer cleanupTestLogs()

	user := "alice"
	userID := 67890

	// Printf-style logging / Printf 스타일 로깅
	logger.Infof("User login: %s (ID: %d)", user, userID)

	// Structured logging / 구조화된 로깅
	logger.Info("User login", "username", user, "user_id", userID)

	// Read log file / 로그 파일 읽기
	content, err := os.ReadFile("./test_logs/printf_vs_structured.log")
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	logStr := string(content)

	// Both should contain user information / 둘 다 사용자 정보를 포함해야 함
	if !strings.Contains(logStr, "alice") {
		t.Error("Log should contain username 'alice'")
	}
	if !strings.Contains(logStr, "67890") {
		t.Error("Log should contain user ID '67890'")
	}

	// Structured logging should have key=value format / 구조화된 로깅은 키=값 형식이어야 함
	if !strings.Contains(logStr, "username=alice") {
		t.Error("Log should contain structured data 'username=alice'")
	}
	if !strings.Contains(logStr, "user_id=67890") {
		t.Error("Log should contain structured data 'user_id=67890'")
	}
}

// TestAppYamlIntegration tests that app.yaml is loaded and used in banner
// TestAppYamlIntegration은 app.yaml이 로드되고 배너에 사용되는지 테스트합니다
func TestAppYamlIntegration(t *testing.T) {
	logger, err := New(
		WithFilePath("./test_logs/app_yaml.log"),
	)
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Close()
	defer cleanupTestLogs()

	// Write a test log to ensure file is created / 파일 생성을 위해 테스트 로그 작성
	logger.Info("Test log message")

	// Read log file / 로그 파일 읽기
	content, err := os.ReadFile("./test_logs/app_yaml.log")
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	logStr := string(content)

	// Verify that app.yaml values are in the banner / app.yaml 값이 배너에 있는지 확인
	// Should contain "go-utils" (app name from cfg/app.yaml)
	if !strings.Contains(logStr, "go-utils") {
		t.Error("Log file should contain app name 'go-utils' from cfg/app.yaml")
	}

	// Should contain "v1.4.012" (version from cfg/app.yaml)
	if !strings.Contains(logStr, "v1.4.012") {
		t.Error("Log file should contain version 'v1.4.012' from cfg/app.yaml")
	}
}
