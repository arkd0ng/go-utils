package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/arkd0ng/go-utils/logging"
	"github.com/arkd0ng/go-utils/stringutil"
)

func main() {
	// Create results directories / 결과 디렉토리 생성
	if err := os.MkdirAll("./results/logs", 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create logs directory: %v\n", err)
		os.Exit(1)
	}

	// Initialize logger / 로거 초기화
	logger, err := logging.New(
		logging.WithFilePath(fmt.Sprintf("./results/logs/stringutil_example_%s.log", time.Now().Format("20060102_150405"))),
		logging.WithLevel(logging.DEBUG),
		logging.WithStdout(true),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Close()

	// Print banner / 배너 출력
	logger.Banner("Stringutil Package Examples", "go-utils/stringutil")

	// Case Conversion Examples / 케이스 변환 예제
	logger.Info("=== Case Conversion Examples ===")
	logger.Info("=== 케이스 변환 예제 ===")
	logger.Info("")

	input := "UserProfileData"
	logger.Info(fmt.Sprintf("Input: %s", input))
	logger.Info(fmt.Sprintf("ToSnakeCase: %s", stringutil.ToSnakeCase(input)))
	logger.Info(fmt.Sprintf("ToCamelCase: %s", stringutil.ToCamelCase(input)))
	logger.Info(fmt.Sprintf("ToKebabCase: %s", stringutil.ToKebabCase(input)))
	logger.Info(fmt.Sprintf("ToPascalCase: %s", stringutil.ToPascalCase(input)))
	logger.Info(fmt.Sprintf("ToScreamingSnakeCase: %s", stringutil.ToScreamingSnakeCase(input)))
	logger.Info("")

	// String Manipulation Examples / 문자열 조작 예제
	logger.Info("=== String Manipulation Examples ===")
	logger.Info("=== 문자열 조작 예제 ===")
	logger.Info("")

	logger.Info(fmt.Sprintf("Truncate('Hello World', 8): %s", stringutil.Truncate("Hello World", 8)))
	logger.Info(fmt.Sprintf("Reverse('hello'): %s", stringutil.Reverse("hello")))
	logger.Info(fmt.Sprintf("Capitalize('hello world'): %s", stringutil.Capitalize("hello world")))
	logger.Info(fmt.Sprintf("Clean('  hello   world  '): '%s'", stringutil.Clean("  hello   world  ")))
	logger.Info(fmt.Sprintf("RemoveSpaces('h e l l o'): %s", stringutil.RemoveSpaces("h e l l o")))
	logger.Info(fmt.Sprintf("Repeat('*', 5): %s", stringutil.Repeat("*", 5)))
	logger.Info(fmt.Sprintf("Repeat('안녕', 3): %s", stringutil.Repeat("안녕", 3)))
	logger.Info(fmt.Sprintf("Substring('hello world', 0, 5): %s", stringutil.Substring("hello world", 0, 5)))
	logger.Info(fmt.Sprintf("Substring('안녕하세요', 0, 2): %s", stringutil.Substring("안녕하세요", 0, 2)))
	logger.Info(fmt.Sprintf("Left('hello world', 5): %s", stringutil.Left("hello world", 5)))
	logger.Info(fmt.Sprintf("Right('hello world', 5): %s", stringutil.Right("hello world", 5)))
	logger.Info(fmt.Sprintf("Insert('hello world', 5, ','): %s", stringutil.Insert("hello world", 5, ",")))
	logger.Info(fmt.Sprintf("SwapCase('Hello World'): %s", stringutil.SwapCase("Hello World")))
	logger.Info("")

	// Validation Examples / 유효성 검사 예제
	logger.Info("=== Validation Examples ===")
	logger.Info("=== 유효성 검사 예제 ===")
	logger.Info("")

	logger.Info(fmt.Sprintf("IsEmail('user@example.com'): %v", stringutil.IsEmail("user@example.com")))
	logger.Info(fmt.Sprintf("IsEmail('invalid.email'): %v", stringutil.IsEmail("invalid.email")))
	logger.Info(fmt.Sprintf("IsURL('https://example.com'): %v", stringutil.IsURL("https://example.com")))
	logger.Info(fmt.Sprintf("IsURL('example.com'): %v", stringutil.IsURL("example.com")))
	logger.Info(fmt.Sprintf("IsAlphanumeric('abc123'): %v", stringutil.IsAlphanumeric("abc123")))
	logger.Info(fmt.Sprintf("IsNumeric('12345'): %v", stringutil.IsNumeric("12345")))
	logger.Info("")

	// Search & Replace Examples / 검색 및 치환 예제
	logger.Info("=== Search & Replace Examples ===")
	logger.Info("=== 검색 및 치환 예제 ===")
	logger.Info("")

	logger.Info(fmt.Sprintf("ContainsAny('hello world', ['foo', 'world']): %v",
		stringutil.ContainsAny("hello world", []string{"foo", "world"})))
	logger.Info(fmt.Sprintf("StartsWithAny('https://example.com', ['http://', 'https://']): %v",
		stringutil.StartsWithAny("https://example.com", []string{"http://", "https://"})))
	logger.Info(fmt.Sprintf("ReplaceAll('a b c', {'a': 'x', 'b': 'y'}): %s",
		stringutil.ReplaceAll("a b c", map[string]string{"a": "x", "b": "y"})))
	logger.Info("")

	// Utilities Examples / 유틸리티 예제
	logger.Info("=== Utilities Examples ===")
	logger.Info("=== 유틸리티 예제 ===")
	logger.Info("")

	logger.Info(fmt.Sprintf("CountWords('hello world foo'): %d", stringutil.CountWords("hello world foo")))
	logger.Info(fmt.Sprintf("PadLeft('5', 3, '0'): %s", stringutil.PadLeft("5", 3, "0")))
	logger.Info(fmt.Sprintf("PadRight('5', 3, '0'): %s", stringutil.PadRight("5", 3, "0")))
	logger.Info(fmt.Sprintf("Words('hello world foo'): %v", stringutil.Words("hello world foo")))

	// Map and Filter examples / Map과 Filter 예제
	strs := []string{"hello", "world", "foo"}
	upper := stringutil.Map(strs, strings.ToUpper)
	logger.Info(fmt.Sprintf("Map(['hello', 'world', 'foo'], ToUpper): %v", upper))

	filtered := stringutil.Filter(strs, func(s string) bool { return len(s) > 3 })
	logger.Info(fmt.Sprintf("Filter(['hello', 'world', 'foo'], len > 3): %v", filtered))
	logger.Info("")

	// Comparison Examples / 비교 예제
	logger.Info("=== Comparison Examples ===")
	logger.Info("=== 비교 예제 ===")
	logger.Info("")

	logger.Info(fmt.Sprintf("EqualFold('hello', 'HELLO'): %v", stringutil.EqualFold("hello", "HELLO")))
	logger.Info(fmt.Sprintf("HasPrefix('hello world', 'hello'): %v", stringutil.HasPrefix("hello world", "hello")))
	logger.Info(fmt.Sprintf("HasSuffix('hello world', 'world'): %v", stringutil.HasSuffix("hello world", "world")))
	logger.Info("")

	// Title and Slug Examples / 타이틀 및 슬러그 예제
	logger.Info("=== Title and Slug Examples ===")
	logger.Info("=== 타이틀 및 슬러그 예제 ===")
	logger.Info("")

	logger.Info(fmt.Sprintf("ToTitle('hello world'): %s", stringutil.ToTitle("hello world")))
	logger.Info(fmt.Sprintf("ToTitle('user_profile_data'): %s", stringutil.ToTitle("user_profile_data")))
	logger.Info(fmt.Sprintf("Slugify('Hello World!'): %s", stringutil.Slugify("Hello World!")))
	logger.Info(fmt.Sprintf("Slugify('Go Utils -- Package'): %s", stringutil.Slugify("Go Utils -- Package")))
	logger.Info("")

	// Quote Examples / 따옴표 예제
	logger.Info("=== Quote Examples ===")
	logger.Info("=== 따옴표 예제 ===")
	logger.Info("")

	logger.Info(fmt.Sprintf("Quote('hello'): %s", stringutil.Quote("hello")))
	logger.Info(fmt.Sprintf("Quote('say \"hi\"'): %s", stringutil.Quote("say \"hi\"")))
	logger.Info(fmt.Sprintf("Unquote('\"hello\"'): %s", stringutil.Unquote("\"hello\"")))
	logger.Info(fmt.Sprintf("Unquote(\"'world'\"): %s", stringutil.Unquote("'world'")))
	logger.Info("")

	// Unicode Examples / 유니코드 예제
	logger.Info("=== Unicode Examples ===")
	logger.Info("=== 유니코드 예제 ===")
	logger.Info("")

	logger.Info(fmt.Sprintf("RuneCount('hello'): %d", stringutil.RuneCount("hello")))
	logger.Info(fmt.Sprintf("RuneCount('안녕하세요'): %d", stringutil.RuneCount("안녕하세요")))
	logger.Info(fmt.Sprintf("RuneCount('🔥🔥'): %d", stringutil.RuneCount("🔥🔥")))
	logger.Info(fmt.Sprintf("Width('hello'): %d", stringutil.Width("hello")))
	logger.Info(fmt.Sprintf("Width('안녕'): %d", stringutil.Width("안녕")))
	logger.Info(fmt.Sprintf("Width('hello世界'): %d", stringutil.Width("hello世界")))
	logger.Info(fmt.Sprintf("Normalize('café', 'NFC'): %s", stringutil.Normalize("café", "NFC")))
	logger.Info(fmt.Sprintf("Normalize('①②③', 'NFKC'): %s", stringutil.Normalize("①②③", "NFKC")))
	logger.Info("")

	logger.Info("=== All examples completed successfully! ===")
	logger.Info("=== 모든 예제가 성공적으로 완료되었습니다! ===")
}
