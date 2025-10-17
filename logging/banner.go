package logging

import (
	"fmt"
	"strings"
)

// Banner prints a formatted banner with application name and version
// Banner는 애플리케이션 이름과 버전이 포함된 형식화된 배너를 출력합니다
//
// Parameters
// 매개변수:
// - appName: application name
// 애플리케이션 이름
// - version: version string
// 버전 문자열
//
// Example
// 예제:
//
//	logger.Banner("My Application", "v1.0.0")
//
// Output
// 출력:
//
//	╔════════════════════════════════════════════════════════════╗
//	║                                                            ║
//	║                 My Application v1.0.0                      ║
//	║                                                            ║
//	╚════════════════════════════════════════════════════════════╝
func (l *Logger) Banner(appName, version string) {
	// Combine app name and version
	// 앱 이름과 버전 결합
	text := fmt.Sprintf("%s %s", appName, version)
	width := len(text) + 12 // Padding / 패딩

	if width < 60 {
		width = 60
	}

	// Build banner
	// 배너 작성
	var banner strings.Builder

	// Top border
	// 상단 경계
	banner.WriteString("╔")
	banner.WriteString(strings.Repeat("═", width))
	banner.WriteString("╗\n")

	// Empty line
	// 빈 줄
	banner.WriteString("║")
	banner.WriteString(strings.Repeat(" ", width))
	banner.WriteString("║\n")

	// Center text line
	// 중앙 정렬 텍스트 줄
	padding := (width - len(text)) / 2
	banner.WriteString("║")
	banner.WriteString(strings.Repeat(" ", padding))
	banner.WriteString(text)
	banner.WriteString(strings.Repeat(" ", width-padding-len(text)))
	banner.WriteString("║\n")

	// Empty line
	// 빈 줄
	banner.WriteString("║")
	banner.WriteString(strings.Repeat(" ", width))
	banner.WriteString("║\n")

	// Bottom border
	// 하단 경계
	banner.WriteString("╚")
	banner.WriteString(strings.Repeat("═", width))
	banner.WriteString("╝\n")

	// Print banner using Info level
	// INFO 레벨로 배너 출력
	l.printRaw(banner.String())
}

// SimpleBanner prints a simple banner with just a separator line
// SimpleBanner는 구분선만 있는 간단한 배너를 출력합니다
//
// Parameters
// 매개변수:
// - appName: application name
// 애플리케이션 이름
// - version: version string
// 버전 문자열
//
// Example
// 예제:
//
//	logger.SimpleBanner("My Application", "v1.0.0")
//
// Output
// 출력:
//
//	========================================
//	My Application v1.0.0
//	========================================
func (l *Logger) SimpleBanner(appName, version string) {
	text := fmt.Sprintf("%s %s", appName, version)
	width := len(text)
	if width < 40 {
		width = 40
	}

	var banner strings.Builder
	separator := strings.Repeat("=", width)

	banner.WriteString(separator + "\n")
	banner.WriteString(text + "\n")
	banner.WriteString(separator + "\n")

	l.printRaw(banner.String())
}

// CustomBanner prints a custom ASCII art banner
// CustomBanner는 사용자 정의 ASCII 아트 배너를 출력합니다
//
// Parameters
// 매개변수:
// - lines: array of banner lines
// 배너 줄의 배열
//
// Example
// 예제:
//
//	logger.CustomBanner([]string{
//	    "  __  __            _             ",
//	    " |  \\/  |_   _     / \\   _ __  _ __",
//	    " | |\\/| | | | |   / _ \\ | '_ \\| '_ \\",
//	    " | |  | | |_| |  / ___ \\| |_) | |_) |",
//	    " |_|  |_|\\__, | /_/   \\_\\ .__/| .__/",
//	    "         |___/          |_|   |_|",
//	})
func (l *Logger) CustomBanner(lines []string) {
	var banner strings.Builder
	for _, line := range lines {
		banner.WriteString(line + "\n")
	}
	l.printRaw(banner.String())
}

// DoubleBanner prints a banner with double-line borders
// DoubleBanner는 이중선 경계가 있는 배너를 출력합니다
//
// Parameters
// 매개변수:
// - appName: application name
// 애플리케이션 이름
// - version: version string
// 버전 문자열
// - description: optional description
// 선택적 설명
//
// Example
// 예제:
//
//	logger.DoubleBanner("My Application", "v1.0.0", "Production Server")
//
// Output
// 출력:
//
//	╔════════════════════════════════════════════════════════════╗
//	║                 My Application v1.0.0                      ║
//	║                   Production Server                        ║
//	╚════════════════════════════════════════════════════════════╝
func (l *Logger) DoubleBanner(appName, version, description string) {
	text1 := fmt.Sprintf("%s %s", appName, version)
	text2 := description

	width := len(text1)
	if len(text2) > width {
		width = len(text2)
	}
	width += 12 // Padding / 패딩

	if width < 60 {
		width = 60
	}

	var banner strings.Builder

	// Top border
	// 상단 경계
	banner.WriteString("╔")
	banner.WriteString(strings.Repeat("═", width))
	banner.WriteString("╗\n")

	// First text line
	// 첫 번째 텍스트 줄
	padding1 := (width - len(text1)) / 2
	banner.WriteString("║")
	banner.WriteString(strings.Repeat(" ", padding1))
	banner.WriteString(text1)
	banner.WriteString(strings.Repeat(" ", width-padding1-len(text1)))
	banner.WriteString("║\n")

	// Second text line if description provided
	// 설명이 제공된 경우 두 번째 텍스트 줄
	if description != "" {
		padding2 := (width - len(text2)) / 2
		banner.WriteString("║")
		banner.WriteString(strings.Repeat(" ", padding2))
		banner.WriteString(text2)
		banner.WriteString(strings.Repeat(" ", width-padding2-len(text2)))
		banner.WriteString("║\n")
	}

	// Bottom border
	// 하단 경계
	banner.WriteString("╚")
	banner.WriteString(strings.Repeat("═", width))
	banner.WriteString("╝\n")

	l.printRaw(banner.String())
}

// SeparatorLine prints a separator line
// SeparatorLine은 구분선을 출력합니다
//
// Parameters
// 매개변수:
// - char: character to use for the line
// 줄에 사용할 문자
// - width: width of the line
// 줄의 너비
//
// Example
// 예제:
//
//	logger.SeparatorLine("=", 50)
//	logger.SeparatorLine("-", 50)
func (l *Logger) SeparatorLine(char string, width int) {
	if width <= 0 {
		width = 50
	}
	if char == "" {
		char = "="
	}
	l.printRaw(strings.Repeat(char, width) + "\n")
}

// printRaw prints raw text without timestamp or level formatting
// printRaw는 타임스탬프나 레벨 형식 없이 원시 텍스트를 출력합니다
func (l *Logger) printRaw(text string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Write to stdout
	// stdout에 작성
	if l.config.enableStdout {
		l.stdoutWriter.Write([]byte(text))
	}

	// Write to file
	// 파일에 작성
	if l.config.enableFile && l.fileWriter != nil {
		l.fileWriter.Write([]byte(text))
	}
}
