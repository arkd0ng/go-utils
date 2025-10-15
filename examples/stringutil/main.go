package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/arkd0ng/go-utils/fileutil"
	"github.com/arkd0ng/go-utils/logging"
	"github.com/arkd0ng/go-utils/stringutil"
)

func main() {
	// Setup log file with backup management / 백업 관리와 함께 로그 파일 설정
	logFilePath := "logs/stringutil-example.log"

	// Check if previous log file exists / 이전 로그 파일 존재 여부 확인
	if fileutil.Exists(logFilePath) {
		// Get modification time of existing log file / 기존 로그 파일의 수정 시간 가져오기
		modTime, err := fileutil.ModTime(logFilePath)
		if err == nil {
			// Create backup filename with timestamp / 타임스탬프와 함께 백업 파일명 생성
			backupName := fmt.Sprintf("logs/stringutil-example-%s.log", modTime.Format("20060102-150405"))

			// Backup existing log file / 기존 로그 파일 백업
			if err := fileutil.CopyFile(logFilePath, backupName); err == nil {
				fmt.Printf("✅ Backed up previous log to: %s\n", backupName)
			}
		}

		// Cleanup old backup files - keep only 5 most recent / 오래된 백업 파일 정리 - 최근 5개만 유지
		backupPattern := "logs/stringutil-example-*.log"
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
		logging.WithStdout(true),
	)
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Close()

	// Print banner / 배너 출력
	logger.Banner("Stringutil Package Examples - All 79 Functions", "go-utils/stringutil")
	logger.Info("")
	logger.Info("This example demonstrates ALL 79 functions in the stringutil package")
	logger.Info("본 예제는 stringutil 패키지의 모든 79개 함수를 시연합니다")
	logger.Info("")

	// ========================================
	// 1. Case Conversion (9 functions) / 케이스 변환 (9개 함수)
	// ========================================
	logger.Info(stringutil.Repeat("=", 80))
	logger.Info("=== 1. CASE CONVERSION (9 functions) ===")
	logger.Info("=== 1. 케이스 변환 (9개 함수) ===")
	logger.Info(stringutil.Repeat("=", 80))
	logger.Info("")

	input := "UserProfileData"
	logger.Info(fmt.Sprintf("Original input / 원본 입력: %s", input))
	logger.Info("")

	// 1.1 ToSnakeCase
	logger.Info("1.1 ToSnakeCase - Converts string to snake_case")
	logger.Info("    문자열을 snake_case로 변환합니다")
	result := stringutil.ToSnakeCase(input)
	logger.Info(fmt.Sprintf("    ToSnakeCase('%s') = '%s'", input, result))
	logger.Info("")

	// 1.2 ToCamelCase
	logger.Info("1.2 ToCamelCase - Converts string to camelCase")
	logger.Info("    문자열을 camelCase로 변환합니다")
	result = stringutil.ToCamelCase(input)
	logger.Info(fmt.Sprintf("    ToCamelCase('%s') = '%s'", input, result))
	logger.Info("")

	// 1.3 ToKebabCase
	logger.Info("1.3 ToKebabCase - Converts string to kebab-case")
	logger.Info("    문자열을 kebab-case로 변환합니다")
	result = stringutil.ToKebabCase(input)
	logger.Info(fmt.Sprintf("    ToKebabCase('%s') = '%s'", input, result))
	logger.Info("")

	// 1.4 ToPascalCase
	logger.Info("1.4 ToPascalCase - Converts string to PascalCase")
	logger.Info("    문자열을 PascalCase로 변환합니다")
	result = stringutil.ToPascalCase(input)
	logger.Info(fmt.Sprintf("    ToPascalCase('%s') = '%s'", input, result))
	logger.Info("")

	// 1.5 ToScreamingSnakeCase
	logger.Info("1.5 ToScreamingSnakeCase - Converts string to SCREAMING_SNAKE_CASE")
	logger.Info("    문자열을 SCREAMING_SNAKE_CASE로 변환합니다")
	result = stringutil.ToScreamingSnakeCase(input)
	logger.Info(fmt.Sprintf("    ToScreamingSnakeCase('%s') = '%s'", input, result))
	logger.Info("")

	// 1.6 ToTitle
	logger.Info("1.6 ToTitle - Converts string to Title Case (each word capitalized)")
	logger.Info("    문자열을 Title Case로 변환합니다 (각 단어의 첫 글자를 대문자로)")
	result = stringutil.ToTitle("hello world")
	logger.Info(fmt.Sprintf("    ToTitle('hello world') = '%s'", result))
	logger.Info("")

	// 1.7 Slugify
	logger.Info("1.7 Slugify - Converts string to URL-friendly slug")
	logger.Info("    문자열을 URL 친화적인 슬러그로 변환합니다")
	result = stringutil.Slugify("Hello World! This is a Test")
	logger.Info(fmt.Sprintf("    Slugify('Hello World! This is a Test') = '%s'", result))
	logger.Info("")

	// 1.8 Quote
	logger.Info("1.8 Quote - Wraps string in double quotes and escapes internal quotes")
	logger.Info("    문자열을 큰따옴표로 감싸고 내부 따옴표를 이스케이프합니다")
	result = stringutil.Quote("say \"hello\"")
	logger.Info(fmt.Sprintf("    Quote('say \"hello\"') = %s", result))
	logger.Info("")

	// 1.9 Unquote
	logger.Info("1.9 Unquote - Removes surrounding quotes and unescapes internal quotes")
	logger.Info("    주변 따옴표를 제거하고 내부 따옴표의 이스케이프를 해제합니다")
	result = stringutil.Unquote("\"hello world\"")
	logger.Info(fmt.Sprintf("    Unquote('\"hello world\"') = '%s'", result))
	logger.Info("")

	// ========================================
	// 2. String Manipulation (17 functions) / 문자열 조작 (17개 함수)
	// ========================================
	logger.Info(stringutil.Repeat("=", 80))
	logger.Info("=== 2. STRING MANIPULATION (17 functions) ===")
	logger.Info("=== 2. 문자열 조작 (17개 함수) ===")
	logger.Info(stringutil.Repeat("=", 80))
	logger.Info("")

	// 2.1 Truncate
	logger.Info("2.1 Truncate - Truncates string to specified length and appends '...'")
	logger.Info("    문자열을 지정된 길이로 자르고 '...'를 추가합니다")
	result = stringutil.Truncate("Hello World", 8)
	logger.Info(fmt.Sprintf("    Truncate('Hello World', 8) = '%s'", result))
	logger.Info("")

	// 2.2 TruncateWithSuffix
	logger.Info("2.2 TruncateWithSuffix - Truncates string with custom suffix")
	logger.Info("    사용자 정의 suffix로 문자열을 자릅니다")
	result = stringutil.TruncateWithSuffix("Hello World", 8, "…")
	logger.Info(fmt.Sprintf("    TruncateWithSuffix('Hello World', 8, '…') = '%s'", result))
	logger.Info("")

	// 2.3 Reverse
	logger.Info("2.3 Reverse - Reverses a string (Unicode-safe)")
	logger.Info("    문자열을 뒤집습니다 (유니코드 안전)")
	result = stringutil.Reverse("hello")
	logger.Info(fmt.Sprintf("    Reverse('hello') = '%s'", result))
	result = stringutil.Reverse("안녕하세요")
	logger.Info(fmt.Sprintf("    Reverse('안녕하세요') = '%s'", result))
	logger.Info("")

	// 2.4 Capitalize
	logger.Info("2.4 Capitalize - Capitalizes first letter of each word")
	logger.Info("    각 단어의 첫 글자를 대문자로 만듭니다")
	result = stringutil.Capitalize("hello world")
	logger.Info(fmt.Sprintf("    Capitalize('hello world') = '%s'", result))
	logger.Info("")

	// 2.5 CapitalizeFirst
	logger.Info("2.5 CapitalizeFirst - Capitalizes only the first letter of the string")
	logger.Info("    문자열의 첫 글자만 대문자로 만듭니다")
	result = stringutil.CapitalizeFirst("hello world")
	logger.Info(fmt.Sprintf("    CapitalizeFirst('hello world') = '%s'", result))
	logger.Info("")

	// 2.6 RemoveDuplicates
	logger.Info("2.6 RemoveDuplicates - Removes duplicate characters from string")
	logger.Info("    문자열에서 중복 문자를 제거합니다")
	result = stringutil.RemoveDuplicates("hello")
	logger.Info(fmt.Sprintf("    RemoveDuplicates('hello') = '%s'", result))
	logger.Info("")

	// 2.7 RemoveSpaces
	logger.Info("2.7 RemoveSpaces - Removes all whitespace from string")
	logger.Info("    문자열에서 모든 공백을 제거합니다")
	result = stringutil.RemoveSpaces("h e l l o")
	logger.Info(fmt.Sprintf("    RemoveSpaces('h e l l o') = '%s'", result))
	logger.Info("")

	// 2.8 RemoveSpecialChars
	logger.Info("2.8 RemoveSpecialChars - Removes special characters, keeping only alphanumeric and spaces")
	logger.Info("    특수 문자를 제거하고 영숫자와 공백만 유지합니다")
	result = stringutil.RemoveSpecialChars("hello@#$123")
	logger.Info(fmt.Sprintf("    RemoveSpecialChars('hello@#$123') = '%s'", result))
	logger.Info("")

	// 2.9 Clean
	logger.Info("2.9 Clean - Trims whitespace and deduplicates spaces")
	logger.Info("    공백을 제거하고 중복 공백을 정리합니다")
	result = stringutil.Clean("  hello   world  ")
	logger.Info(fmt.Sprintf("    Clean('  hello   world  ') = '%s'", result))
	logger.Info("")

	// 2.10 Repeat
	logger.Info("2.10 Repeat - Repeats a string n times")
	logger.Info("     문자열을 n번 반복합니다")
	result = stringutil.Repeat("*", 5)
	logger.Info(fmt.Sprintf("     Repeat('*', 5) = '%s'", result))
	result = stringutil.Repeat("안녕", 3)
	logger.Info(fmt.Sprintf("     Repeat('안녕', 3) = '%s'", result))
	logger.Info("")

	// 2.11 Substring
	logger.Info("2.11 Substring - Extracts substring from start to end index (Unicode-safe)")
	logger.Info("     start부터 end 인덱스까지 부분 문자열을 추출합니다 (유니코드 안전)")
	result = stringutil.Substring("hello world", 0, 5)
	logger.Info(fmt.Sprintf("     Substring('hello world', 0, 5) = '%s'", result))
	result = stringutil.Substring("안녕하세요", 0, 2)
	logger.Info(fmt.Sprintf("     Substring('안녕하세요', 0, 2) = '%s'", result))
	logger.Info("")

	// 2.12 Left
	logger.Info("2.12 Left - Returns leftmost n characters (Unicode-safe)")
	logger.Info("     가장 왼쪽 n개 문자를 반환합니다 (유니코드 안전)")
	result = stringutil.Left("hello world", 5)
	logger.Info(fmt.Sprintf("     Left('hello world', 5) = '%s'", result))
	logger.Info("")

	// 2.13 Right
	logger.Info("2.13 Right - Returns rightmost n characters (Unicode-safe)")
	logger.Info("     가장 오른쪽 n개 문자를 반환합니다 (유니코드 안전)")
	result = stringutil.Right("hello world", 5)
	logger.Info(fmt.Sprintf("     Right('hello world', 5) = '%s'", result))
	logger.Info("")

	// 2.14 Insert
	logger.Info("2.14 Insert - Inserts string at specified index (Unicode-safe)")
	logger.Info("     지정된 인덱스에 문자열을 삽입합니다 (유니코드 안전)")
	result = stringutil.Insert("hello world", 5, ",")
	logger.Info(fmt.Sprintf("     Insert('hello world', 5, ',') = '%s'", result))
	logger.Info("")

	// 2.15 SwapCase
	logger.Info("2.15 SwapCase - Swaps case of all letters")
	logger.Info("     모든 글자의 대소문자를 반전합니다")
	result = stringutil.SwapCase("Hello World")
	logger.Info(fmt.Sprintf("     SwapCase('Hello World') = '%s'", result))
	logger.Info("")

	// 2.16 PadLeft
	logger.Info("2.16 PadLeft - Pads string on left to reach specified length")
	logger.Info("     지정된 길이에 도달하도록 문자열의 왼쪽에 패딩을 추가합니다")
	result = stringutil.PadLeft("5", 3, "0")
	logger.Info(fmt.Sprintf("     PadLeft('5', 3, '0') = '%s'", result))
	logger.Info("")

	// 2.17 PadRight
	logger.Info("2.17 PadRight - Pads string on right to reach specified length")
	logger.Info("     지정된 길이에 도달하도록 문자열의 오른쪽에 패딩을 추가합니다")
	result = stringutil.PadRight("5", 3, "0")
	logger.Info(fmt.Sprintf("     PadRight('5', 3, '0') = '%s'", result))
	logger.Info("")

	// ========================================
	// 3. Validation (8 functions) / 유효성 검사 (8개 함수)
	// ========================================
	logger.Info(stringutil.Repeat("=", 80))
	logger.Info("=== 3. VALIDATION (8 functions) ===")
	logger.Info("=== 3. 유효성 검사 (8개 함수) ===")
	logger.Info(stringutil.Repeat("=", 80))
	logger.Info("")

	// 3.1 IsEmail
	logger.Info("3.1 IsEmail - Validates if string is an email address")
	logger.Info("    문자열이 이메일 주소인지 검증합니다")
	logger.Info(fmt.Sprintf("    IsEmail('user@example.com') = %v", stringutil.IsEmail("user@example.com")))
	logger.Info(fmt.Sprintf("    IsEmail('invalid.email') = %v", stringutil.IsEmail("invalid.email")))
	logger.Info("")

	// 3.2 IsURL
	logger.Info("3.2 IsURL - Validates if string is a URL")
	logger.Info("    문자열이 URL인지 검증합니다")
	logger.Info(fmt.Sprintf("    IsURL('https://example.com') = %v", stringutil.IsURL("https://example.com")))
	logger.Info(fmt.Sprintf("    IsURL('example.com') = %v", stringutil.IsURL("example.com")))
	logger.Info("")

	// 3.3 IsAlphanumeric
	logger.Info("3.3 IsAlphanumeric - Checks if string contains only alphanumeric characters")
	logger.Info("    문자열이 영숫자만 포함하는지 확인합니다")
	logger.Info(fmt.Sprintf("    IsAlphanumeric('abc123') = %v", stringutil.IsAlphanumeric("abc123")))
	logger.Info(fmt.Sprintf("    IsAlphanumeric('abc-123') = %v", stringutil.IsAlphanumeric("abc-123")))
	logger.Info("")

	// 3.4 IsNumeric
	logger.Info("3.4 IsNumeric - Checks if string contains only digits")
	logger.Info("    문자열이 숫자만 포함하는지 확인합니다")
	logger.Info(fmt.Sprintf("    IsNumeric('12345') = %v", stringutil.IsNumeric("12345")))
	logger.Info(fmt.Sprintf("    IsNumeric('123.45') = %v", stringutil.IsNumeric("123.45")))
	logger.Info("")

	// 3.5 IsAlpha
	logger.Info("3.5 IsAlpha - Checks if string contains only letters")
	logger.Info("    문자열이 알파벳만 포함하는지 확인합니다")
	logger.Info(fmt.Sprintf("    IsAlpha('abcABC') = %v", stringutil.IsAlpha("abcABC")))
	logger.Info(fmt.Sprintf("    IsAlpha('abc123') = %v", stringutil.IsAlpha("abc123")))
	logger.Info("")

	// 3.6 IsBlank
	logger.Info("3.6 IsBlank - Checks if string is empty or contains only whitespace")
	logger.Info("    문자열이 비어있거나 공백만 포함하는지 확인합니다")
	logger.Info(fmt.Sprintf("    IsBlank('') = %v", stringutil.IsBlank("")))
	logger.Info(fmt.Sprintf("    IsBlank('   ') = %v", stringutil.IsBlank("   ")))
	logger.Info(fmt.Sprintf("    IsBlank('hello') = %v", stringutil.IsBlank("hello")))
	logger.Info("")

	// 3.7 IsLower
	logger.Info("3.7 IsLower - Checks if all letters are lowercase")
	logger.Info("    모든 글자가 소문자인지 확인합니다")
	logger.Info(fmt.Sprintf("    IsLower('hello') = %v", stringutil.IsLower("hello")))
	logger.Info(fmt.Sprintf("    IsLower('Hello') = %v", stringutil.IsLower("Hello")))
	logger.Info("")

	// 3.8 IsUpper
	logger.Info("3.8 IsUpper - Checks if all letters are uppercase")
	logger.Info("    모든 글자가 대문자인지 확인합니다")
	logger.Info(fmt.Sprintf("    IsUpper('HELLO') = %v", stringutil.IsUpper("HELLO")))
	logger.Info(fmt.Sprintf("    IsUpper('Hello') = %v", stringutil.IsUpper("Hello")))
	logger.Info("")

	// ========================================
	// 4. Comparison (3 functions) / 비교 (3개 함수)
	// ========================================
	logger.Info(stringutil.Repeat("=", 80))
	logger.Info("=== 4. COMPARISON (3 functions) ===")
	logger.Info("=== 4. 비교 (3개 함수) ===")
	logger.Info(stringutil.Repeat("=", 80))
	logger.Info("")

	// 4.1 EqualFold
	logger.Info("4.1 EqualFold - Compares strings case-insensitively")
	logger.Info("    두 문자열을 대소문자 구분 없이 비교합니다")
	logger.Info(fmt.Sprintf("    EqualFold('hello', 'HELLO') = %v", stringutil.EqualFold("hello", "HELLO")))
	logger.Info("")

	// 4.2 HasPrefix
	logger.Info("4.2 HasPrefix - Checks if string starts with prefix")
	logger.Info("    문자열이 접두사로 시작하는지 확인합니다")
	logger.Info(fmt.Sprintf("    HasPrefix('hello world', 'hello') = %v", stringutil.HasPrefix("hello world", "hello")))
	logger.Info("")

	// 4.3 HasSuffix
	logger.Info("4.3 HasSuffix - Checks if string ends with suffix")
	logger.Info("    문자열이 접미사로 끝나는지 확인합니다")
	logger.Info(fmt.Sprintf("    HasSuffix('hello world', 'world') = %v", stringutil.HasSuffix("hello world", "world")))
	logger.Info("")

	// ========================================
	// 5. Search & Replace (6 functions) / 검색 및 치환 (6개 함수)
	// ========================================
	logger.Info(stringutil.Repeat("=", 80))
	logger.Info("=== 5. SEARCH & REPLACE (6 functions) ===")
	logger.Info("=== 5. 검색 및 치환 (6개 함수) ===")
	logger.Info(stringutil.Repeat("=", 80))
	logger.Info("")

	// 5.1 ContainsAny
	logger.Info("5.1 ContainsAny - Returns true if string contains any of the substrings")
	logger.Info("    문자열이 부분 문자열 중 하나라도 포함하면 true를 반환합니다")
	logger.Info(fmt.Sprintf("    ContainsAny('hello world', ['foo', 'world']) = %v",
		stringutil.ContainsAny("hello world", []string{"foo", "world"})))
	logger.Info("")

	// 5.2 ContainsAll
	logger.Info("5.2 ContainsAll - Returns true if string contains all of the substrings")
	logger.Info("    문자열이 모든 부분 문자열을 포함하면 true를 반환합니다")
	logger.Info(fmt.Sprintf("    ContainsAll('hello world', ['hello', 'world']) = %v",
		stringutil.ContainsAll("hello world", []string{"hello", "world"})))
	logger.Info("")

	// 5.3 StartsWithAny
	logger.Info("5.3 StartsWithAny - Returns true if string starts with any of the prefixes")
	logger.Info("    문자열이 접두사 중 하나로 시작하면 true를 반환합니다")
	logger.Info(fmt.Sprintf("    StartsWithAny('https://example.com', ['http://', 'https://']) = %v",
		stringutil.StartsWithAny("https://example.com", []string{"http://", "https://"})))
	logger.Info("")

	// 5.4 EndsWithAny
	logger.Info("5.4 EndsWithAny - Returns true if string ends with any of the suffixes")
	logger.Info("    문자열이 접미사 중 하나로 끝나면 true를 반환합니다")
	logger.Info(fmt.Sprintf("    EndsWithAny('file.txt', ['.txt', '.md']) = %v",
		stringutil.EndsWithAny("file.txt", []string{".txt", ".md"})))
	logger.Info("")

	// 5.5 ReplaceAll
	logger.Info("5.5 ReplaceAll - Replaces multiple strings at once using a replacement map")
	logger.Info("    치환 맵을 사용하여 여러 문자열을 한 번에 치환합니다")
	logger.Info(fmt.Sprintf("    ReplaceAll('a b c', {'a': 'x', 'b': 'y'}) = '%s'",
		stringutil.ReplaceAll("a b c", map[string]string{"a": "x", "b": "y"})))
	logger.Info("")

	// 5.6 ReplaceIgnoreCase
	logger.Info("5.6 ReplaceIgnoreCase - Replaces substring ignoring case")
	logger.Info("    대소문자를 무시하고 부분 문자열을 치환합니다")
	logger.Info(fmt.Sprintf("    ReplaceIgnoreCase('Hello World', 'hello', 'hi') = '%s'",
		stringutil.ReplaceIgnoreCase("Hello World", "hello", "hi")))
	logger.Info("")

	// ========================================
	// 6. Unicode Operations (3 functions) / 유니코드 작업 (3개 함수)
	// ========================================
	logger.Info(stringutil.Repeat("=", 80))
	logger.Info("=== 6. UNICODE OPERATIONS (3 functions) ===")
	logger.Info("=== 6. 유니코드 작업 (3개 함수) ===")
	logger.Info(stringutil.Repeat("=", 80))
	logger.Info("")

	// 6.1 RuneCount
	logger.Info("6.1 RuneCount - Counts Unicode characters (not bytes)")
	logger.Info("    유니코드 문자 개수를 셉니다 (바이트가 아님)")
	logger.Info(fmt.Sprintf("    RuneCount('hello') = %d", stringutil.RuneCount("hello")))
	logger.Info(fmt.Sprintf("    RuneCount('안녕하세요') = %d", stringutil.RuneCount("안녕하세요")))
	logger.Info(fmt.Sprintf("    RuneCount('🔥🔥') = %d", stringutil.RuneCount("🔥🔥")))
	logger.Info("")

	// 6.2 Width
	logger.Info("6.2 Width - Calculates East Asian width (CJK characters count as 2)")
	logger.Info("    동아시아 폭을 계산합니다 (CJK 문자는 2로 계산)")
	logger.Info(fmt.Sprintf("    Width('hello') = %d", stringutil.Width("hello")))
	logger.Info(fmt.Sprintf("    Width('안녕') = %d", stringutil.Width("안녕")))
	logger.Info(fmt.Sprintf("    Width('hello世界') = %d", stringutil.Width("hello世界")))
	logger.Info("")

	// 6.3 Normalize
	logger.Info("6.3 Normalize - Performs Unicode normalization (NFC, NFD, NFKC, NFKD)")
	logger.Info("    유니코드 정규화를 수행합니다 (NFC, NFD, NFKC, NFKD)")
	logger.Info(fmt.Sprintf("    Normalize('café', 'NFC') = '%s'", stringutil.Normalize("café", "NFC")))
	logger.Info(fmt.Sprintf("    Normalize('①②③', 'NFKC') = '%s'", stringutil.Normalize("①②③", "NFKC")))
	logger.Info("")

	// ========================================
	// 7. Collection Utilities (7 functions) / 컬렉션 유틸리티 (7개 함수)
	// ========================================
	logger.Info(stringutil.Repeat("=", 80))
	logger.Info("=== 7. COLLECTION UTILITIES (7 functions) ===")
	logger.Info("=== 7. 컬렉션 유틸리티 (7개 함수) ===")
	logger.Info(stringutil.Repeat("=", 80))
	logger.Info("")

	// 7.1 CountWords
	logger.Info("7.1 CountWords - Counts number of words (split by whitespace)")
	logger.Info("    단어 수를 셉니다 (공백으로 분리)")
	logger.Info(fmt.Sprintf("    CountWords('hello world foo') = %d", stringutil.CountWords("hello world foo")))
	logger.Info("")

	// 7.2 CountOccurrences
	logger.Info("7.2 CountOccurrences - Counts occurrences of substring")
	logger.Info("    부분 문자열이 나타나는 횟수를 셉니다")
	logger.Info(fmt.Sprintf("    CountOccurrences('hello hello', 'hello') = %d",
		stringutil.CountOccurrences("hello hello", "hello")))
	logger.Info("")

	// 7.3 Lines
	logger.Info("7.3 Lines - Splits string by newlines")
	logger.Info("    줄바꿈으로 문자열을 분리합니다")
	logger.Info(fmt.Sprintf("    Lines('line1\\nline2\\nline3') = %v",
		stringutil.Lines("line1\nline2\nline3")))
	logger.Info("")

	// 7.4 Words
	logger.Info("7.4 Words - Splits string by whitespace")
	logger.Info("    공백으로 문자열을 분리합니다")
	logger.Info(fmt.Sprintf("    Words('hello world foo') = %v", stringutil.Words("hello world foo")))
	logger.Info("")

	// 7.5 Map
	logger.Info("7.5 Map - Applies function to all strings in slice")
	logger.Info("    슬라이스의 모든 문자열에 함수를 적용합니다")
	strs := []string{"hello", "world", "foo"}
	upper := stringutil.Map(strs, strings.ToUpper)
	logger.Info(fmt.Sprintf("    Map(['hello', 'world', 'foo'], ToUpper) = %v", upper))
	logger.Info("")

	// 7.6 Filter
	logger.Info("7.6 Filter - Filters strings by predicate function")
	logger.Info("    조건 함수로 문자열을 필터링합니다")
	filtered := stringutil.Filter(strs, func(s string) bool { return len(s) > 3 })
	logger.Info(fmt.Sprintf("    Filter(['hello', 'world', 'foo'], len > 3) = %v", filtered))
	logger.Info("")

	// 7.7 Join
	logger.Info("7.7 Join - Joins slice of strings with separator")
	logger.Info("    구분자로 문자열 슬라이스를 연결합니다")
	logger.Info(fmt.Sprintf("    Join(['a', 'b', 'c'], '-') = '%s'",
		stringutil.Join([]string{"a", "b", "c"}, "-")))
	logger.Info("")

	// ========================================
	// 8. Encoding/Decoding (8 functions) / 인코딩/디코딩 (8개 함수)
	// ========================================
	logger.Info(stringutil.Repeat("=", 80))
	logger.Info("=== 8. ENCODING/DECODING (8 functions) ===")
	logger.Info("=== 8. 인코딩/디코딩 (8개 함수) ===")
	logger.Info(stringutil.Repeat("=", 80))
	logger.Info("")

	// 8.1 Base64Encode
	logger.Info("8.1 Base64Encode - Encodes string to Base64")
	logger.Info("    문자열을 Base64로 인코딩합니다")
	plainText := "Hello, 안녕하세요!"
	encoded := stringutil.Base64Encode(plainText)
	logger.Info(fmt.Sprintf("    Base64Encode('%s') = '%s'", plainText, encoded))
	logger.Info("")

	// 8.2 Base64Decode
	logger.Info("8.2 Base64Decode - Decodes Base64 string")
	logger.Info("    Base64 문자열을 디코딩합니다")
	decoded, _ := stringutil.Base64Decode(encoded)
	logger.Info(fmt.Sprintf("    Base64Decode(encoded) = '%s'", decoded))
	logger.Info("")

	// 8.3 Base64URLEncode
	logger.Info("8.3 Base64URLEncode - Encodes string to URL-safe Base64")
	logger.Info("    문자열을 URL 안전 Base64로 인코딩합니다")
	urlText := "hello?world=test&foo=bar"
	urlEncoded := stringutil.Base64URLEncode(urlText)
	logger.Info(fmt.Sprintf("    Base64URLEncode('%s') = '%s'", urlText, urlEncoded))
	logger.Info("")

	// 8.4 Base64URLDecode
	logger.Info("8.4 Base64URLDecode - Decodes URL-safe Base64 string")
	logger.Info("    URL 안전 Base64 문자열을 디코딩합니다")
	urlDecoded, _ := stringutil.Base64URLDecode(urlEncoded)
	logger.Info(fmt.Sprintf("    Base64URLDecode(encoded) = '%s'", urlDecoded))
	logger.Info("")

	// 8.5 URLEncode
	logger.Info("8.5 URLEncode - Encodes string for URL query parameters")
	logger.Info("    URL 쿼리 매개변수용으로 문자열을 인코딩합니다")
	urlParam := "hello world & foo=bar"
	paramEncoded := stringutil.URLEncode(urlParam)
	logger.Info(fmt.Sprintf("    URLEncode('%s') = '%s'", urlParam, paramEncoded))
	logger.Info("")

	// 8.6 URLDecode
	logger.Info("8.6 URLDecode - Decodes URL-encoded string")
	logger.Info("    URL 인코딩된 문자열을 디코딩합니다")
	paramDecoded, _ := stringutil.URLDecode(paramEncoded)
	logger.Info(fmt.Sprintf("    URLDecode(encoded) = '%s'", paramDecoded))
	logger.Info("")

	// 8.7 HTMLEscape
	logger.Info("8.7 HTMLEscape - Escapes HTML special characters")
	logger.Info("    HTML 특수 문자를 이스케이프합니다")
	htmlText := "<div>Hello & \"World\"</div>"
	htmlEscaped := stringutil.HTMLEscape(htmlText)
	logger.Info(fmt.Sprintf("    HTMLEscape('%s') = '%s'", htmlText, htmlEscaped))
	logger.Info("")

	// 8.8 HTMLUnescape
	logger.Info("8.8 HTMLUnescape - Unescapes HTML entities")
	logger.Info("    HTML 엔터티를 언이스케이프합니다")
	htmlUnescaped := stringutil.HTMLUnescape(htmlEscaped)
	logger.Info(fmt.Sprintf("    HTMLUnescape(escaped) = '%s'", htmlUnescaped))
	logger.Info("")

	// ========================================
	// 9. String Distance/Similarity (4 functions) / 문자열 거리/유사도 (4개 함수)
	// ========================================
	logger.Info(stringutil.Repeat("=", 80))
	logger.Info("=== 9. STRING DISTANCE/SIMILARITY (4 functions) ===")
	logger.Info("=== 9. 문자열 거리/유사도 (4개 함수) ===")
	logger.Info(stringutil.Repeat("=", 80))
	logger.Info("")

	// 9.1 LevenshteinDistance
	logger.Info("9.1 LevenshteinDistance - Calculates Levenshtein distance between two strings")
	logger.Info("    두 문자열 간의 Levenshtein 거리를 계산합니다")
	str1, str2 := "kitten", "sitting"
	levDist := stringutil.LevenshteinDistance(str1, str2)
	logger.Info(fmt.Sprintf("    LevenshteinDistance('%s', '%s') = %d", str1, str2, levDist))
	logger.Info("")

	// 9.2 Similarity
	logger.Info("9.2 Similarity - Calculates similarity ratio (0.0 to 1.0)")
	logger.Info("    유사도 비율을 계산합니다 (0.0에서 1.0)")
	sim := stringutil.Similarity(str1, str2)
	logger.Info(fmt.Sprintf("    Similarity('%s', '%s') = %.3f", str1, str2, sim))
	logger.Info("")

	// 9.3 HammingDistance
	logger.Info("9.3 HammingDistance - Calculates Hamming distance (equal-length strings only)")
	logger.Info("    Hamming 거리를 계산합니다 (동일 길이 문자열만)")
	str3, str4 := "karolin", "kathrin"
	hammingDist := stringutil.HammingDistance(str3, str4)
	logger.Info(fmt.Sprintf("    HammingDistance('%s', '%s') = %d", str3, str4, hammingDist))
	logger.Info("")

	// 9.4 JaroWinklerSimilarity
	logger.Info("9.4 JaroWinklerSimilarity - Calculates Jaro-Winkler similarity")
	logger.Info("    Jaro-Winkler 유사도를 계산합니다")
	str5, str6 := "martha", "marhta"
	jaroSim := stringutil.JaroWinklerSimilarity(str5, str6)
	logger.Info(fmt.Sprintf("    JaroWinklerSimilarity('%s', '%s') = %.3f", str5, str6, jaroSim))
	logger.Info("")

	// ========================================
	// 10. Formatting (10 functions) / 포맷팅 (10개 함수)
	// ========================================
	logger.Info(stringutil.Repeat("=", 80))
	logger.Info("=== 10. FORMATTING (10 functions) ===")
	logger.Info("=== 10. 포맷팅 (10개 함수) ===")
	logger.Info(stringutil.Repeat("=", 80))
	logger.Info("")

	// 10.1 FormatNumber
	logger.Info("10.1 FormatNumber - Formats number with thousand separators")
	logger.Info("     천 단위 구분자로 숫자를 포맷합니다")
	logger.Info(fmt.Sprintf("     FormatNumber(1000000, ',') = '%s'",
		stringutil.FormatNumber(1000000, ",")))
	logger.Info("")

	// 10.2 FormatBytes
	logger.Info("10.2 FormatBytes - Formats bytes to human-readable size")
	logger.Info("     바이트를 사람이 읽기 쉬운 크기로 포맷합니다")
	logger.Info(fmt.Sprintf("     FormatBytes(1024) = '%s'", stringutil.FormatBytes(1024)))
	logger.Info(fmt.Sprintf("     FormatBytes(1048576) = '%s'", stringutil.FormatBytes(1048576)))
	logger.Info("")

	// 10.3 FormatWithCount
	logger.Info("10.3 FormatWithCount - Formats string with count and plural form")
	logger.Info("     개수와 복수형으로 문자열을 포맷합니다")
	logger.Info(fmt.Sprintf("     FormatWithCount(1, 'item', 'items') = '%s'",
		stringutil.FormatWithCount(1, "item", "items")))
	logger.Info(fmt.Sprintf("     FormatWithCount(5, 'item', 'items') = '%s'",
		stringutil.FormatWithCount(5, "item", "items")))
	logger.Info("")

	// 10.4 Pluralize
	logger.Info("10.4 Pluralize - Returns plural form if count is not 1")
	logger.Info("     개수가 1이 아니면 복수형을 반환합니다")
	logger.Info(fmt.Sprintf("     Pluralize(1, 'item', 'items') = '%s'",
		stringutil.Pluralize(1, "item", "items")))
	logger.Info(fmt.Sprintf("     Pluralize(5, 'item', 'items') = '%s'",
		stringutil.Pluralize(5, "item", "items")))
	logger.Info("")

	// 10.5 Ellipsis
	logger.Info("10.5 Ellipsis - Truncates string with ellipsis in middle")
	logger.Info("     문자열을 중간에 ellipsis를 넣어 자릅니다")
	longFilename := "verylongfilename.txt"
	logger.Info(fmt.Sprintf("     Ellipsis('%s', 15) = '%s'",
		longFilename, stringutil.Ellipsis(longFilename, 15)))
	logger.Info("")

	// 10.6 Mask
	logger.Info("10.6 Mask - Masks characters except for first and last n characters")
	logger.Info("     처음과 마지막 n개 문자를 제외하고 마스킹합니다")
	logger.Info(fmt.Sprintf("     Mask('1234567890', 2, 2, '*') = '%s'",
		stringutil.Mask("1234567890", 2, 2, "*")))
	logger.Info("")

	// 10.7 MaskEmail
	logger.Info("10.7 MaskEmail - Masks email address")
	logger.Info("     이메일 주소를 마스킹합니다")
	logger.Info(fmt.Sprintf("     MaskEmail('john.doe@example.com') = '%s'",
		stringutil.MaskEmail("john.doe@example.com")))
	logger.Info("")

	// 10.8 MaskCreditCard
	logger.Info("10.8 MaskCreditCard - Masks credit card number")
	logger.Info("     신용카드 번호를 마스킹합니다")
	logger.Info(fmt.Sprintf("     MaskCreditCard('1234567890123456') = '%s'",
		stringutil.MaskCreditCard("1234567890123456")))
	logger.Info("")

	// 10.9 AddLineNumbers
	logger.Info("10.9 AddLineNumbers - Adds line numbers to multi-line text")
	logger.Info("     여러 줄 텍스트에 줄 번호를 추가합니다")
	multiLineText := "line 1\nline 2\nline 3"
	logger.Info(fmt.Sprintf("     AddLineNumbers:\n%s", stringutil.AddLineNumbers(multiLineText)))
	logger.Info("")

	// 10.10 Indent
	logger.Info("10.10 Indent - Adds indentation to each line")
	logger.Info("     각 줄에 들여쓰기를 추가합니다")
	codeSnippet := "func main() {\n  fmt.Println(\"hello\")\n}"
	logger.Info(fmt.Sprintf("     Indent (2 spaces):\n%s", stringutil.Indent(codeSnippet, "  ")))
	logger.Info("")

	// 10.11 Dedent
	logger.Info("10.11 Dedent - Removes common leading whitespace")
	logger.Info("      공통 앞 공백을 제거합니다")
	indentedCode := "    func main() {\n      fmt.Println(\"hello\")\n    }"
	logger.Info(fmt.Sprintf("      Dedent:\n%s", stringutil.Dedent(indentedCode)))
	logger.Info("")

	// 10.12 WrapText (bonus)
	logger.Info("10.12 WrapText - Wraps text to specified width")
	logger.Info("      텍스트를 지정된 너비로 줄바꿈합니다")
	longText := "The quick brown fox jumps over the lazy dog"
	logger.Info(fmt.Sprintf("      WrapText (width 20):\n%s", stringutil.WrapText(longText, 20)))
	logger.Info("")

	// ========================================
	// 11. Builder Pattern / 빌더 패턴
	// ========================================
	logger.Info(stringutil.Repeat("=", 80))
	logger.Info("=== 11. BUILDER PATTERN (Method Chaining) ===")
	logger.Info("=== 11. 빌더 패턴 (메서드 체이닝) ===")
	logger.Info(stringutil.Repeat("=", 80))
	logger.Info("")

	logger.Info("Builder Pattern allows method chaining for complex string transformations")
	logger.Info("빌더 패턴은 복잡한 문자열 변환을 위한 메서드 체이닝을 가능하게 합니다")
	logger.Info("")

	builderResult1 := stringutil.NewBuilder().
		Append("  user profile data  ").
		Clean().
		ToSnakeCase().
		ToUpper().
		Build()
	logger.Info(fmt.Sprintf("Example 1: NewBuilder().Append().Clean().ToSnakeCase().ToUpper().Build()"))
	logger.Info(fmt.Sprintf("Result: '%s'", builderResult1))
	logger.Info("")

	builderResult2 := stringutil.NewBuilder().
		Append("Hello World").
		ToKebabCase().
		Quote().
		Build()
	logger.Info(fmt.Sprintf("Example 2: NewBuilder().Append().ToKebabCase().Quote().Build()"))
	logger.Info(fmt.Sprintf("Result: %s", builderResult2))
	logger.Info("")

	builderResult3 := stringutil.NewBuilderWithString("the quick brown fox jumps over the lazy dog").
		Capitalize().
		Truncate(30).
		Build()
	logger.Info(fmt.Sprintf("Example 3: NewBuilderWithString().Capitalize().Truncate().Build()"))
	logger.Info(fmt.Sprintf("Result: '%s'", builderResult3))
	logger.Info("")

	complexBuilder := stringutil.NewBuilder().
		AppendLine("Line 1: User Profile").
		AppendLine("Line 2: Data Processing").
		ToLower().
		Replace(":", " →").
		Build()
	logger.Info(fmt.Sprintf("Example 4: Complex builder with multiple operations"))
	logger.Info(fmt.Sprintf("Result:\n%s", complexBuilder))
	logger.Info("")

	// ========================================
	// 12. Real-world Scenarios / 실제 사용 시나리오
	// ========================================
	logger.Info(stringutil.Repeat("=", 80))
	logger.Info("=== 12. REAL-WORLD SCENARIOS ===")
	logger.Info("=== 12. 실제 사용 시나리오 ===")
	logger.Info(stringutil.Repeat("=", 80))
	logger.Info("")

	// Scenario 1: Clean user input
	logger.Info("Scenario 1: Clean user input for database")
	logger.Info("시나리오 1: 데이터베이스용 사용자 입력 처리")
	userInput := "  John DOE  "
	processed := stringutil.NewBuilder().
		Append(userInput).
		Clean().
		ToTitle().
		Build()
	logger.Info(fmt.Sprintf("  Raw input: '%s'", userInput))
	logger.Info(fmt.Sprintf("  Processed: '%s'", processed))
	logger.Info("")

	// Scenario 2: URL-friendly slug
	logger.Info("Scenario 2: Generate URL-friendly slug from title")
	logger.Info("시나리오 2: URL 친화적 슬러그 생성")
	articleTitle := "How to Use Go Utils: A Complete Guide!"
	slug := stringutil.Slugify(articleTitle)
	logger.Info(fmt.Sprintf("  Title: '%s'", articleTitle))
	logger.Info(fmt.Sprintf("  Slug: '%s'", slug))
	logger.Info("")

	// Scenario 3: Format API response
	logger.Info("Scenario 3: Format API response with pluralization")
	logger.Info("시나리오 3: 복수형을 사용한 API 응답 포맷")
	filesFound := 42
	responseMsg := fmt.Sprintf("Found %s", stringutil.FormatWithCount(filesFound, "file", "files"))
	logger.Info(fmt.Sprintf("  Message: %s", responseMsg))
	logger.Info("")

	// Scenario 4: Mask sensitive data
	logger.Info("Scenario 4: Mask sensitive data in logs")
	logger.Info("시나리오 4: 로그에서 민감한 데이터 마스크")
	email := "sensitive.user@example.com"
	creditCard := "1234-5678-9012-3456"
	logger.Info(fmt.Sprintf("  Email: %s", stringutil.MaskEmail(email)))
	logger.Info(fmt.Sprintf("  Credit Card: %s", stringutil.MaskCreditCard(creditCard)))
	logger.Info("")

	// Scenario 5: Find similar strings
	logger.Info("Scenario 5: Find similar strings (typo correction)")
	logger.Info("시나리오 5: 유사한 문자열 찾기 (오타 수정)")
	userSearch := "golang"
	knownTerms := []string{"Go", "Golang", "Python", "Java", "JavaScript"}
	logger.Info(fmt.Sprintf("  User search: '%s'", userSearch))
	logger.Info("  Suggestions:")

	type suggestion struct {
		term  string
		score float64
	}
	var suggestions []suggestion
	for _, term := range knownTerms {
		score := stringutil.Similarity(strings.ToLower(userSearch), strings.ToLower(term))
		if score > 0.3 { // threshold
			suggestions = append(suggestions, suggestion{term, score})
		}
	}
	// Sort by score
	for i := 0; i < len(suggestions)-1; i++ {
		for j := 0; j < len(suggestions)-i-1; j++ {
			if suggestions[j].score < suggestions[j+1].score {
				suggestions[j], suggestions[j+1] = suggestions[j+1], suggestions[j]
			}
		}
	}
	for _, sug := range suggestions {
		logger.Info(fmt.Sprintf("    - %s (similarity: %.2f)", sug.term, sug.score))
	}
	logger.Info("")

	// ========================================
	// Summary / 요약
	// ========================================
	logger.Info(stringutil.Repeat("=", 80))
	logger.Info("=== SUMMARY / 요약 ===")
	logger.Info(stringutil.Repeat("=", 80))
	logger.Info("")
	logger.Info("This example demonstrated ALL 53 functions in the stringutil package:")
	logger.Info("본 예제는 stringutil 패키지의 모든 53개 함수를 시연했습니다:")
	logger.Info("")
	logger.Info("  1. Case Conversion (9 functions) - Case transformations")
	logger.Info("     케이스 변환 (9개 함수) - 케이스 변환")
	logger.Info("  2. String Manipulation (17 functions) - String operations")
	logger.Info("     문자열 조작 (17개 함수) - 문자열 작업")
	logger.Info("  3. Validation (8 functions) - String validation")
	logger.Info("     유효성 검사 (8개 함수) - 문자열 검증")
	logger.Info("  4. Comparison (3 functions) - String comparison")
	logger.Info("     비교 (3개 함수) - 문자열 비교")
	logger.Info("  5. Search & Replace (6 functions) - Finding and replacing")
	logger.Info("     검색 및 치환 (6개 함수) - 검색 및 치환")
	logger.Info("  6. Unicode Operations (3 functions) - Unicode handling")
	logger.Info("     유니코드 작업 (3개 함수) - 유니코드 처리")
	logger.Info("  7. Collection Utilities (7 functions) - Slice operations")
	logger.Info("     컬렉션 유틸리티 (7개 함수) - 슬라이스 작업")
	logger.Info("  8. Encoding/Decoding (8 functions) - Encode/decode strings")
	logger.Info("     인코딩/디코딩 (8개 함수) - 문자열 인코딩/디코딩")
	logger.Info("  9. String Distance/Similarity (4 functions) - Distance algorithms")
	logger.Info("     문자열 거리/유사도 (4개 함수) - 거리 알고리즘")
	logger.Info(" 10. Formatting (10+ functions) - String formatting")
	logger.Info("     포맷팅 (10개 이상 함수) - 문자열 포맷팅")
	logger.Info(" 11. Builder Pattern - Method chaining")
	logger.Info("     빌더 패턴 - 메서드 체이닝")
	logger.Info("")
	logger.Info("All examples completed successfully!")
	logger.Info("모든 예제가 성공적으로 완료되었습니다!")
	logger.Info("")
}
