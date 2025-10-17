package logging

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger provides structured logging with file rotation
// Logger는 파일 로테이션을 지원하는 구조화된 로깅을 제공합니다
type Logger struct {
	config       *config
	fileWriter   *lumberjack.Logger
	stdoutWriter io.Writer
	mu           sync.Mutex
}

// New creates a new Logger with the given options
// New는 주어진 옵션으로 새로운 Logger를 생성합니다
//
// Parameters
// 매개변수:
// - opts: configuration options
// 설정 옵션
//
// Returns
// 반환값:
// - *Logger: new logger instance
// 새 로거 인스턴스
// - error: error if any
// 에러가 있으면 에러
//
// Example
// 예제:
//
//	logger, err := logging.New(
//	    logging.WithFilePath("./logs/app.log"),
//	    logging.WithLevel(logging.DEBUG),
//	)
func New(opts ...Option) (*Logger, error) {
	cfg := defaultConfig()
	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return nil, err
		}
	}

	logger := &Logger{
		config:       cfg,
		stdoutWriter: os.Stdout,
	}

	// Create file writer if file output is enabled
	// 파일 출력이 활성화된 경우 파일 writer 생성
	if cfg.enableFile {
		// Ensure log directory exists
		// 로그 디렉토리가 존재하는지 확인
		logDir := filepath.Dir(cfg.filename)
		if err := os.MkdirAll(logDir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create log directory: %w", err)
		}

		logger.fileWriter = &lumberjack.Logger{
			Filename:   cfg.filename,
			MaxSize:    cfg.maxSize,
			MaxBackups: cfg.maxBackups,
			MaxAge:     cfg.maxAge,
			Compress:   cfg.compress,
		}
	}

	// Print auto banner if enabled
	// 자동 배너가 활성화된 경우 배너 출력
	if cfg.autoBanner {
		bannerName := cfg.appName

		// If appName is default "Application", extract from filename / appName이 기본값 "Application"이면 파일명에서 추출
		if bannerName == "Application" && cfg.filename != "" {
			// Extract filename without path and extension
			// 경로와 확장자를 제외한 파일명 추출
			base := filepath.Base(cfg.filename) // "database.log"
			ext := filepath.Ext(base)           // ".log"
			if ext != "" {
				bannerName = base[:len(base)-len(ext)] // "database"
			} else {
				bannerName = base
			}
		}

		logger.Banner(bannerName, cfg.appVersion)
	}

	return logger, nil
}

// Default creates a Logger with default settings
// Default는 기본 설정으로 Logger를 생성합니다
//
// Returns
// 반환값:
// - *Logger: logger with default configuration
// 기본 설정의 로거
//
// Example
// 예제:
//
//	logger := logging.Default()
//	logger.Info("Application started")
func Default() *Logger {
	logger, _ := New()
	return logger
}

// log is the internal logging function
// log는 내부 로깅 함수입니다
func (l *Logger) log(level Level, msg string, keysAndValues ...interface{}) {
	// Skip if level is below configured minimum
	// 레벨이 설정된 최소값보다 낮으면 건너뜀
	if level < l.config.level {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	// Format timestamp
	// 타임스탬프 형식화
	timestamp := time.Now().Format(l.config.timeFormat)

	// Build log message
	// 로그 메시지 작성
	var logMsg string

	// Add prefix if configured
	// 설정된 경우 프리픽스 추가
	prefix := l.config.prefix
	if prefix != "" {
		prefix = prefix + " "
	}

	// Format: timestamp [LEVEL] prefix message key=value key=value...
	// 형식: 타임스탬프 [레벨] 프리픽스 메시지 키=값 키=값...
	logMsg = fmt.Sprintf("%s [%s] %s%s", timestamp, level.String(), prefix, msg)

	// Add structured key-value pairs
	// 구조화된 키-값 쌍 추가
	if len(keysAndValues) > 0 {
		logMsg += " "
		for i := 0; i < len(keysAndValues); i += 2 {
			if i+1 < len(keysAndValues) {
				logMsg += fmt.Sprintf("%v=%v ", keysAndValues[i], keysAndValues[i+1])
			}
		}
	}

	logMsg += "\n"

	// Write to stdout with color if enabled
	// 활성화된 경우 색상으로 stdout에 작성
	if l.config.enableStdout {
		colorMsg := logMsg
		if l.config.enableColor {
			colorMsg = fmt.Sprintf("%s%s%s", level.ColorCode(), logMsg, ResetColor())
		}
		l.stdoutWriter.Write([]byte(colorMsg))
	}

	// Write to file without color
	// 색상 없이 파일에 작성
	if l.config.enableFile && l.fileWriter != nil {
		l.fileWriter.Write([]byte(logMsg))
	}
}

// Debug logs a message at DEBUG level
// Debug는 DEBUG 레벨로 메시지를 로깅합니다
//
// Parameters
// 매개변수:
// - msg: log message
// 로그 메시지
// - keysAndValues: optional key-value pairs for structured logging
// 구조화된 로깅을 위한 선택적 키-값 쌍
//
// Example
// 예제:
//
//	logger.Debug("Processing request", "user_id", 12345, "ip", "192.168.1.1")
func (l *Logger) Debug(msg string, keysAndValues ...interface{}) {
	l.log(DEBUG, msg, keysAndValues...)
}

// Info logs a message at INFO level
// Info는 INFO 레벨로 메시지를 로깅합니다
//
// Parameters
// 매개변수:
// - msg: log message
// 로그 메시지
// - keysAndValues: optional key-value pairs for structured logging
// 구조화된 로깅을 위한 선택적 키-값 쌍
//
// Example
// 예제:
//
//	logger.Info("Server started", "port", 8080)
func (l *Logger) Info(msg string, keysAndValues ...interface{}) {
	l.log(INFO, msg, keysAndValues...)
}

// Warn logs a message at WARN level
// Warn은 WARN 레벨로 메시지를 로깅합니다
//
// Parameters
// 매개변수:
// - msg: log message
// 로그 메시지
// - keysAndValues: optional key-value pairs for structured logging
// 구조화된 로깅을 위한 선택적 키-값 쌍
//
// Example
// 예제:
//
//	logger.Warn("High memory usage", "usage", "85%")
func (l *Logger) Warn(msg string, keysAndValues ...interface{}) {
	l.log(WARN, msg, keysAndValues...)
}

// Error logs a message at ERROR level
// Error는 ERROR 레벨로 메시지를 로깅합니다
//
// Parameters
// 매개변수:
// - msg: log message
// 로그 메시지
// - keysAndValues: optional key-value pairs for structured logging
// 구조화된 로깅을 위한 선택적 키-값 쌍
//
// Example
// 예제:
//
//	logger.Error("Failed to connect", "error", err, "retry", 3)
func (l *Logger) Error(msg string, keysAndValues ...interface{}) {
	l.log(ERROR, msg, keysAndValues...)
}

// Fatal logs a message at FATAL level and exits the program
// Fatal은 FATAL 레벨로 메시지를 로깅하고 프로그램을 종료합니다
//
// Parameters
// 매개변수:
// - msg: log message
// 로그 메시지
// - keysAndValues: optional key-value pairs for structured logging
// 구조화된 로깅을 위한 선택적 키-값 쌍
//
// Example
// 예제:
//
//	logger.Fatal("Critical error", "error", err)
func (l *Logger) Fatal(msg string, keysAndValues ...interface{}) {
	l.log(FATAL, msg, keysAndValues...)
	os.Exit(1)
}

// logf is the internal logging function for Printf-style formatting
// logf는 Printf 스타일 형식화를 위한 내부 로깅 함수입니다
func (l *Logger) logf(level Level, format string, args ...interface{}) {
	// Skip if level is below configured minimum
	// 레벨이 설정된 최소값보다 낮으면 건너뜀
	if level < l.config.level {
		return
	}

	// Format the message
	// 메시지 형식화
	msg := fmt.Sprintf(format, args...)

	// Use the existing log function without key-value pairs
	// 키-값 쌍 없이 기존 log 함수 사용
	l.log(level, msg)
}

// Debugf logs a formatted message at DEBUG level
// Debugf는 DEBUG 레벨로 형식화된 메시지를 로깅합니다
//
// Parameters
// 매개변수:
// - format: format string
// 형식 문자열
// - args: arguments for formatting
// 형식화를 위한 인자
//
// Example
// 예제:
//
//	logger.Debugf("Processing request from %s (ID: %d)", username, userID)
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logf(DEBUG, format, args...)
}

// Infof logs a formatted message at INFO level
// Infof는 INFO 레벨로 형식화된 메시지를 로깅합니다
//
// Parameters
// 매개변수:
// - format: format string
// 형식 문자열
// - args: arguments for formatting
// 형식화를 위한 인자
//
// Example
// 예제:
//
//	logger.Infof("Server started on port %d", port)
func (l *Logger) Infof(format string, args ...interface{}) {
	l.logf(INFO, format, args...)
}

// Warnf logs a formatted message at WARN level
// Warnf는 WARN 레벨로 형식화된 메시지를 로깅합니다
//
// Parameters
// 매개변수:
// - format: format string
// 형식 문자열
// - args: arguments for formatting
// 형식화를 위한 인자
//
// Example
// 예제:
//
//	logger.Warnf("Memory usage at %d%%", memPercent)
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logf(WARN, format, args...)
}

// Errorf logs a formatted message at ERROR level
// Errorf는 ERROR 레벨로 형식화된 메시지를 로깅합니다
//
// Parameters
// 매개변수:
// - format: format string
// 형식 문자열
// - args: arguments for formatting
// 형식화를 위한 인자
//
// Example
// 예제:
//
//	logger.Errorf("Failed to connect to %s: %v", host, err)
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logf(ERROR, format, args...)
}

// Fatalf logs a formatted message at FATAL level and exits the program
// Fatalf는 FATAL 레벨로 형식화된 메시지를 로깅하고 프로그램을 종료합니다
//
// Parameters
// 매개변수:
// - format: format string
// 형식 문자열
// - args: arguments for formatting
// 형식화를 위한 인자
//
// Example
// 예제:
//
//	logger.Fatalf("Critical error: %v", err)
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logf(FATAL, format, args...)
	os.Exit(1)
}

// SetLevel changes the minimum log level
// SetLevel은 최소 로그 레벨을 변경합니다
//
// Parameters
// 매개변수:
// - level: new minimum log level
// 새 최소 로그 레벨
//
// Example
// 예제:
//
//	logger.SetLevel(logging.DEBUG)
func (l *Logger) SetLevel(level Level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.config.level = level
}

// GetLevel returns the current log level
// GetLevel은 현재 로그 레벨을 반환합니다
//
// Returns
// 반환값:
// - Level: current log level
// 현재 로그 레벨
func (l *Logger) GetLevel() Level {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.config.level
}

// Close closes the logger and flushes any buffered data
// Close는 로거를 닫고 버퍼링된 데이터를 플러시합니다
//
// Returns
// 반환값:
// - error: error if any
// 에러가 있으면 에러
//
// Example
// 예제:
//
//	defer logger.Close()
func (l *Logger) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.fileWriter != nil {
		return l.fileWriter.Close()
	}
	return nil
}

// Rotate manually triggers log file rotation
// Rotate는 로그 파일 로테이션을 수동으로 트리거합니다
//
// Returns
// 반환값:
// - error: error if any
// 에러가 있으면 에러
//
// Example
// 예제:
//
//	if err := logger.Rotate(); err != nil {
//	    log.Printf("Failed to rotate log: %v", err)
//	}
func (l *Logger) Rotate() error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.fileWriter != nil {
		return l.fileWriter.Rotate()
	}
	return nil
}
