package logging

import "strings"

// Level represents the severity level of a log message
// Level은 로그 메시지의 심각도 레벨을 나타냅니다
type Level int

// Log levels in ascending order of severity
// 심각도 오름차순으로 정렬된 로그 레벨
const (
	DEBUG Level = iota // Debug level for detailed information / 상세 정보를 위한 디버그 레벨
	INFO               // Info level for general information / 일반 정보를 위한 정보 레벨
	WARN               // Warning level for warning messages / 경고 메시지를 위한 경고 레벨
	ERROR              // Error level for error messages / 에러 메시지를 위한 에러 레벨
	FATAL              // Fatal level for critical errors / 치명적 에러를 위한 치명적 레벨
)

// String returns the string representation of the log level
// String은 로그 레벨의 문자열 표현을 반환합니다
func (l Level) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// ParseLevel parses a string into a Level
// ParseLevel은 문자열을 Level로 파싱합니다
//
// Parameters
// 매개변수:
// - s: level string (case-insensitive)
// 레벨 문자열 (대소문자 구분 안함)
//
// Returns
// 반환값:
// - Level: parsed level, defaults to INFO if invalid
// 파싱된 레벨, 유효하지 않으면 INFO 기본값
func ParseLevel(s string) Level {
	switch strings.ToUpper(s) {
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "WARN", "WARNING":
		return WARN
	case "ERROR":
		return ERROR
	case "FATAL":
		return FATAL
	default:
		return INFO
	}
}

// ColorCode returns ANSI color code for the log level
// ColorCode는 로그 레벨에 대한 ANSI 색상 코드를 반환합니다
func (l Level) ColorCode() string {
	switch l {
	case DEBUG:
		return "\033[36m" // Cyan / 청록색
	case INFO:
		return "\033[32m" // Green / 녹색
	case WARN:
		return "\033[33m" // Yellow / 노란색
	case ERROR:
		return "\033[31m" // Red / 빨간색
	case FATAL:
		return "\033[35m" // Magenta / 자홍색
	default:
		return "\033[0m" // Reset / 재설정
	}
}

// ResetColor returns ANSI reset code
// ResetColor는 ANSI 재설정 코드를 반환합니다
func ResetColor() string {
	return "\033[0m"
}
