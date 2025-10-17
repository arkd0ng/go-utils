package stringutil

import (
	"strings"
	"unicode"
)

// =============================================================================
// File: case.go
// Purpose: String Case Conversion Operations
// 파일: case.go
// 목적: 문자열 케이스 변환 연산
// =============================================================================
//
// OVERVIEW
// 개요
// --------
// The case.go file provides comprehensive case conversion operations for strings,
// handling various naming conventions commonly used in programming (snake_case,
// camelCase, PascalCase, kebab-case, etc.). These functions intelligently parse
// input strings regardless of their current format and convert them to the desired
// case style. Additionally, the file includes utilities for URL-friendly slugs,
// quoting/unquoting, and title case formatting.
//
// case.go 파일은 프로그래밍에서 일반적으로 사용되는 다양한 명명 규칙
// (snake_case, camelCase, PascalCase, kebab-case 등)을 처리하는 포괄적인
// 케이스 변환 연산을 제공합니다. 이러한 함수는 현재 형식에 관계없이 입력
// 문자열을 지능적으로 파싱하고 원하는 케이스 스타일로 변환합니다. 또한
// 이 파일에는 URL 친화적 슬러그, 따옴표 추가/제거, 제목 케이스 포맷팅을
// 위한 유틸리티가 포함되어 있습니다.
//
// DESIGN PHILOSOPHY
// 설계 철학
// -----------------
// 1. **Format-Agnostic Input**: Functions accept any common case format as input
//    **형식 독립적 입력**: 함수는 모든 일반적인 케이스 형식을 입력으로 받음
//
// 2. **Intelligent Parsing**: Automatically detect word boundaries from various indicators
//    **지능적 파싱**: 다양한 표시자에서 단어 경계 자동 감지
//
// 3. **Consistent Output**: Produce predictable, standard-compliant case formats
//    **일관된 출력**: 예측 가능하고 표준 준수 케이스 형식 생성
//
// 4. **Preserving Semantics**: Maintain word boundaries and meaning during conversion
//    **의미 보존**: 변환 중 단어 경계와 의미 유지
//
// 5. **Developer-Friendly**: Reduce boilerplate for common case conversion tasks
//    **개발자 친화적**: 일반적인 케이스 변환 작업의 보일러플레이트 감소
//
// FUNCTION CATEGORIES
// 함수 범주
// -------------------
//
// 1. PROGRAMMING CASE CONVENTIONS (프로그래밍 케이스 규칙)
//    - ToSnakeCase: Convert to snake_case (lowercase with underscores)
//      ToSnakeCase: snake_case로 변환 (밑줄이 있는 소문자)
//    - ToCamelCase: Convert to camelCase (first word lowercase)
//      ToCamelCase: camelCase로 변환 (첫 단어 소문자)
//    - ToPascalCase: Convert to PascalCase (all words capitalized)
//      ToPascalCase: PascalCase로 변환 (모든 단어 대문자로 시작)
//    - ToKebabCase: Convert to kebab-case (lowercase with hyphens)
//      ToKebabCase: kebab-case로 변환 (하이픈이 있는 소문자)
//    - ToScreamingSnakeCase: Convert to SCREAMING_SNAKE_CASE (uppercase with underscores)
//      ToScreamingSnakeCase: SCREAMING_SNAKE_CASE로 변환 (밑줄이 있는 대문자)
//
// 2. HUMAN-READABLE FORMATTING (사람이 읽을 수 있는 포맷팅)
//    - ToTitle: Convert to Title Case (each word capitalized, space-separated)
//      ToTitle: Title Case로 변환 (각 단어 대문자, 공백으로 구분)
//    - Slugify: Convert to URL-friendly slug
//      Slugify: URL 친화적 슬러그로 변환
//
// 3. QUOTING OPERATIONS (따옴표 연산)
//    - Quote: Wrap string in quotes with proper escaping
//      Quote: 적절한 이스케이프로 문자열을 따옴표로 감싸기
//    - Unquote: Remove surrounding quotes and unescape
//      Unquote: 주변 따옴표 제거 및 이스케이프 해제
//
// 4. INTERNAL UTILITIES (내부 유틸리티)
//    - splitIntoWords: Parse string into word tokens
//      splitIntoWords: 문자열을 단어 토큰으로 파싱
//
// KEY OPERATIONS SUMMARY
// 주요 연산 요약
// ----------------------
//
// ToSnakeCase(s string) string
// - Purpose: Convert any format to snake_case
// - 목적: 모든 형식을 snake_case로 변환
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n) for word array + result
// - 공간 복잡도: O(n), 단어 배열 + 결과용
// - Input Formats: PascalCase, camelCase, kebab-case, SCREAMING_SNAKE_CASE, space-separated
// - 입력 형식: PascalCase, camelCase, kebab-case, SCREAMING_SNAKE_CASE, 공백으로 구분
// - Use Cases: Database column names, API parameters, file names
// - 사용 사례: 데이터베이스 컬럼명, API 매개변수, 파일명
//
// ToCamelCase(s string) string
// - Purpose: Convert any format to camelCase
// - 목적: 모든 형식을 camelCase로 변환
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Behavior: First word lowercase, subsequent words capitalized
// - 동작: 첫 단어 소문자, 이후 단어 대문자로 시작
// - Use Cases: JavaScript variables, JSON keys, Java methods
// - 사용 사례: JavaScript 변수, JSON 키, Java 메서드
//
// ToPascalCase(s string) string
// - Purpose: Convert any format to PascalCase
// - 목적: 모든 형식을 PascalCase로 변환
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Behavior: All words capitalized, no separators
// - 동작: 모든 단어 대문자로 시작, 구분자 없음
// - Use Cases: Go struct names, C# classes, type names
// - 사용 사례: Go 구조체명, C# 클래스, 타입명
//
// ToKebabCase(s string) string
// - Purpose: Convert any format to kebab-case
// - 목적: 모든 형식을 kebab-case로 변환
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Behavior: Lowercase with hyphens
// - 동작: 하이픈이 있는 소문자
// - Use Cases: CSS class names, HTML attributes, URLs
// - 사용 사례: CSS 클래스명, HTML 속성, URL
//
// ToScreamingSnakeCase(s string) string
// - Purpose: Convert any format to SCREAMING_SNAKE_CASE
// - 목적: 모든 형식을 SCREAMING_SNAKE_CASE로 변환
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Behavior: Uppercase with underscores
// - 동작: 밑줄이 있는 대문자
// - Use Cases: Constants, environment variables, configuration keys
// - 사용 사례: 상수, 환경 변수, 설정 키
//
// ToTitle(s string) string
// - Purpose: Convert to Title Case with spaces
// - 목적: 공백이 있는 Title Case로 변환
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Behavior: Each word capitalized, space-separated
// - 동작: 각 단어 대문자로 시작, 공백으로 구분
// - Use Cases: Headings, titles, human-readable labels
// - 사용 사례: 제목, 타이틀, 사람이 읽을 수 있는 레이블
//
// Slugify(s string) string
// - Purpose: Convert to URL-friendly slug
// - 목적: URL 친화적 슬러그로 변환
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Behavior: Lowercase, hyphens for spaces/special chars, no consecutive hyphens
// - 동작: 소문자, 공백/특수 문자는 하이픈, 연속 하이픈 없음
// - Use Cases: URLs, SEO-friendly paths, blog post URLs
// - 사용 사례: URL, SEO 친화적 경로, 블로그 게시물 URL
//
// Quote(s string) string
// - Purpose: Wrap string in double quotes with escaping
// - 목적: 이스케이프와 함께 문자열을 큰따옴표로 감싸기
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Escaping: Handles internal quotes and backslashes
// - 이스케이프: 내부 따옴표와 백슬래시 처리
// - Use Cases: JSON encoding, shell escaping, string literals
// - 사용 사례: JSON 인코딩, 셸 이스케이프, 문자열 리터럴
//
// Unquote(s string) string
// - Purpose: Remove surrounding quotes and unescape
// - 목적: 주변 따옴표 제거 및 이스케이프 해제
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Supports: Double quotes (") and single quotes (')
// - 지원: 큰따옴표(") 및 작은따옴표(')
// - Use Cases: JSON parsing, configuration parsing, CLI argument processing
// - 사용 사례: JSON 파싱, 설정 파싱, CLI 인자 처리
//
// WORD BOUNDARY DETECTION
// 단어 경계 감지
// -----------------------
// The splitIntoWords helper intelligently detects word boundaries based on:
// splitIntoWords 헬퍼는 다음을 기반으로 단어 경계를 지능적으로 감지합니다:
//
// 1. **Delimiters**: Hyphens (-), underscores (_), spaces
//    **구분자**: 하이픈 (-), 밑줄 (_), 공백
//
// 2. **Case Transitions**: Lowercase to uppercase (e.g., "userId" → "user", "Id")
//    **케이스 전환**: 소문자에서 대문자로 (예: "userId" → "user", "Id")
//
// 3. **Consecutive Uppercase**: Multiple uppercase followed by lowercase
//    **연속 대문자**: 여러 대문자 다음에 소문자
//    (e.g., "HTTPServer" → "HTTP", "Server")
//    (예: "HTTPServer" → "HTTP", "Server")
//
// Examples of Word Splitting:
// 단어 분리 예제:
// - "UserProfileData" → ["User", "Profile", "Data"]
// - "user_profile_data" → ["user", "profile", "data"]
// - "user-profile-data" → ["user", "profile", "data"]
// - "HTTPServer" → ["HTTP", "Server"]
// - "XMLParser" → ["XML", "Parser"]
// - "parseHTMLString" → ["parse", "HTML", "String"]
//
// PERFORMANCE CHARACTERISTICS
// 성능 특성
// ---------------------------
//
// Time Complexities:
// 시간 복잡도:
// - All case conversion functions: O(n) - single pass through string
//   모든 케이스 변환 함수: O(n) - 문자열 단일 패스
// - splitIntoWords: O(n) - one scan for word boundaries
//   splitIntoWords: O(n) - 단어 경계를 위한 한 번의 스캔
// - Quote/Unquote: O(n) - character-by-character processing
//   Quote/Unquote: O(n) - 문자별 처리
// - Slugify: O(n) - single pass with rune processing
//   Slugify: O(n) - rune 처리를 통한 단일 패스
//
// Space Complexities:
// 공간 복잡도:
// - Word array: O(w) where w is number of words (typically w << n)
//   단어 배열: O(w), w는 단어 수 (일반적으로 w << n)
// - Result string: O(n) for output
//   결과 문자열: O(n), 출력용
// - Quote: O(n + 2e) where e is number of escaped characters
//   Quote: O(n + 2e), e는 이스케이프된 문자 수
//
// Optimization Tips:
// 최적화 팁:
// 1. Cache conversion results if converting same strings repeatedly
//    동일한 문자열을 반복적으로 변환하는 경우 변환 결과 캐시
// 2. For batch conversions, consider parallel processing
//    배치 변환의 경우 병렬 처리 고려
// 3. splitIntoWords is called by all case functions - minimize redundant calls
//    splitIntoWords는 모든 케이스 함수에서 호출됨 - 중복 호출 최소화
// 4. For ASCII-only strings, byte operations may be faster than rune operations
//    ASCII 전용 문자열의 경우 바이트 연산이 rune 연산보다 빠를 수 있음
// 5. Slugify allocates for rune slice - reuse builders for multiple slugs
//    Slugify는 rune 슬라이스를 할당 - 여러 슬러그에 빌더 재사용
//
// EDGE CASES AND SPECIAL BEHAVIORS
// 엣지 케이스 및 특수 동작
// ---------------------------------
//
// Empty Strings:
// 빈 문자열:
// - All functions return empty string for empty input
//   모든 함수는 빈 입력에 대해 빈 문자열 반환
// - ToCamelCase("") returns ""
//   ToCamelCase("")는 "" 반환
// - Slugify("") returns ""
//   Slugify("")는 "" 반환
//
// Single Word:
// 단일 단어:
// - ToSnakeCase("hello") returns "hello"
//   ToSnakeCase("hello")는 "hello" 반환
// - ToCamelCase("hello") returns "hello"
//   ToCamelCase("hello")는 "hello" 반환
// - ToPascalCase("hello") returns "Hello"
//   ToPascalCase("hello")는 "Hello" 반환
//
// All Uppercase:
// 모두 대문자:
// - ToSnakeCase("HTTP") returns "http"
//   ToSnakeCase("HTTP")는 "http" 반환
// - ToCamelCase("HTTP") returns "http"
//   ToCamelCase("HTTP")는 "http" 반환
// - Treated as single word
//   단일 단어로 처리
//
// Acronyms:
// 약어:
// - "HTTPServer" → snake_case: "http_server"
//   "HTTPServer" → snake_case: "http_server"
// - "XMLParser" → camelCase: "xmlParser"
//   "XMLParser" → camelCase: "xmlParser"
// - Consecutive uppercase letters kept together until lowercase appears
//   연속 대문자는 소문자가 나타날 때까지 함께 유지
//
// Multiple Delimiters:
// 여러 구분자:
// - "user__profile___data" → "user_profile_data" (delimiters deduplicated)
//   "user__profile___data" → "user_profile_data" (구분자 중복 제거)
// - "user---profile" → "user-profile" (for kebab-case)
//   "user---profile" → "user-profile" (kebab-case용)
//
// Special Characters in Slugify:
// Slugify의 특수 문자:
// - Non-alphanumeric characters converted to hyphens
//   영숫자가 아닌 문자는 하이픈으로 변환
// - Consecutive hyphens collapsed to single hyphen
//   연속 하이픈은 단일 하이픈으로 축소
// - Leading/trailing hyphens removed
//   앞뒤 하이픈 제거
// - "Hello, World!" → "hello-world"
//   "Hello, World!" → "hello-world"
//
// Unquote Without Quotes:
// 따옴표 없는 Unquote:
// - Unquote("hello") returns "hello" (no change)
//   Unquote("hello")는 "hello" 반환 (변경 없음)
// - Only removes quotes if both present
//   둘 다 있는 경우에만 따옴표 제거
//
// COMMON USAGE PATTERNS
// 일반 사용 패턴
// ---------------------
//
// 1. Converting Database Column Names
//    데이터베이스 컬럼명 변환:
//
//    structField := "UserProfileID"
//    columnName := stringutil.ToSnakeCase(structField)
//    // "user_profile_id"
//    // Suitable for PostgreSQL, MySQL columns
//    // PostgreSQL, MySQL 컬럼에 적합
//
// 2. Converting to JavaScript Variable Names
//    JavaScript 변수명으로 변환:
//
//    goStruct := "UserProfileData"
//    jsVariable := stringutil.ToCamelCase(goStruct)
//    // "userProfileData"
//    // API response fields, JSON keys
//    // API 응답 필드, JSON 키
//
// 3. Creating URL Slugs for Blog Posts
//    블로그 게시물용 URL 슬러그 생성:
//
//    title := "How to Use Go Utils Package"
//    slug := stringutil.Slugify(title)
//    // "how-to-use-go-utils-package"
//    // SEO-friendly URLs
//    // SEO 친화적 URL
//
// 4. Converting Environment Variable Names
//    환경 변수명 변환:
//
//    configKey := "databaseConnectionTimeout"
//    envVar := stringutil.ToScreamingSnakeCase(configKey)
//    // "DATABASE_CONNECTION_TIMEOUT"
//    // .env files, system environment
//    // .env 파일, 시스템 환경
//
// 5. Formatting Struct Names for Display
//    디스플레이용 구조체명 포맷팅:
//
//    structName := "UserProfileData"
//    displayName := stringutil.ToTitle(structName)
//    // "User Profile Data"
//    // Human-readable labels, UI text
//    // 사람이 읽을 수 있는 레이블, UI 텍스트
//
// 6. Converting API Parameter Names
//    API 매개변수명 변환:
//
//    goParam := "UserID"
//    apiParam := stringutil.ToKebabCase(goParam)
//    // "user-id"
//    // REST API paths, query parameters
//    // REST API 경로, 쿼리 매개변수
//
// 7. Handling JSON with Different Case Styles
//    다른 케이스 스타일의 JSON 처리:
//
//    // Go struct field
//    // Go 구조체 필드
//    field := "CreatedAt"
//
//    // Convert to various API formats
//    // 다양한 API 형식으로 변환
//    jsonSnake := stringutil.ToSnakeCase(field)      // "created_at" (Python API)
//    jsonCamel := stringutil.ToCamelCase(field)      // "createdAt" (JavaScript API)
//    jsonKebab := stringutil.ToKebabCase(field)      // "created-at" (Lisp-like API)
//
// 8. Creating CSS Class Names
//    CSS 클래스명 생성:
//
//    componentName := "UserProfileCard"
//    cssClass := stringutil.ToKebabCase(componentName)
//    // "user-profile-card"
//    // BEM methodology, CSS modules
//    // BEM 방법론, CSS 모듈
//
// 9. Quoting Strings for Shell Commands
//    셸 명령을 위한 문자열 따옴표:
//
//    filename := "my file.txt"
//    quoted := stringutil.Quote(filename)
//    // "\"my file.txt\""
//    // Safe for shell execution
//    // 셸 실행에 안전
//
// 10. Parsing Configuration Values
//     설정 값 파싱:
//
//     configValue := "\"localhost:8080\""
//     unquoted := stringutil.Unquote(configValue)
//     // "localhost:8080"
//     // Reading from config files
//     // 설정 파일에서 읽기
//
// COMPARISON WITH RELATED FUNCTIONS
// 관련 함수와의 비교
// ---------------------------------
//
// ToSnakeCase vs ToKebabCase
// - ToSnakeCase: Uses underscores (_), for programming identifiers
//   ToSnakeCase: 밑줄 (_) 사용, 프로그래밍 식별자용
// - ToKebabCase: Uses hyphens (-), for URLs and CSS
//   ToKebabCase: 하이픈 (-) 사용, URL 및 CSS용
// - Same word splitting logic
//   동일한 단어 분리 로직
// - Use ToSnakeCase for: Variables, database columns
//   ToSnakeCase 사용: 변수, 데이터베이스 컬럼
// - Use ToKebabCase for: URLs, HTML attributes
//   ToKebabCase 사용: URL, HTML 속성
//
// ToCamelCase vs ToPascalCase
// - ToCamelCase: First word lowercase (userProfile)
//   ToCamelCase: 첫 단어 소문자 (userProfile)
// - ToPascalCase: All words capitalized (UserProfile)
//   ToPascalCase: 모든 단어 대문자로 시작 (UserProfile)
// - Use ToCamelCase for: Methods, variables, JSON keys
//   ToCamelCase 사용: 메서드, 변수, JSON 키
// - Use ToPascalCase for: Types, classes, constructors
//   ToPascalCase 사용: 타입, 클래스, 생성자
//
// ToTitle vs ToPascalCase
// - ToTitle: Space-separated words (User Profile Data)
//   ToTitle: 공백으로 구분된 단어 (User Profile Data)
// - ToPascalCase: No separators (UserProfileData)
//   ToPascalCase: 구분자 없음 (UserProfileData)
// - Use ToTitle for: Human-readable display
//   ToTitle 사용: 사람이 읽을 수 있는 디스플레이
// - Use ToPascalCase for: Programming identifiers
//   ToPascalCase 사용: 프로그래밍 식별자
//
// Slugify vs ToKebabCase
// - Slugify: Removes special characters, only alphanumeric + hyphens
//   Slugify: 특수 문자 제거, 영숫자 + 하이픈만
// - ToKebabCase: Preserves all characters, just changes separators
//   ToKebabCase: 모든 문자 유지, 구분자만 변경
// - Slugify: More aggressive cleanup for URLs
//   Slugify: URL을 위한 더 적극적인 정리
// - Use Slugify for: Public-facing URLs
//   Slugify 사용: 공개 URL
// - Use ToKebabCase for: Internal identifiers
//   ToKebabCase 사용: 내부 식별자
//
// Quote vs strconv.Quote
// - Quote: Simple quote wrapping with basic escaping
//   Quote: 기본 이스케이프로 간단한 따옴표 감싸기
// - strconv.Quote: Full Go string literal escaping
//   strconv.Quote: 전체 Go 문자열 리터럴 이스케이프
// - Quote: Faster, less comprehensive
//   Quote: 더 빠르고 덜 포괄적
// - Use Quote for: Basic shell/JSON quoting
//   Quote 사용: 기본 셸/JSON 따옴표
// - Use strconv.Quote for: Go source code generation
//   strconv.Quote 사용: Go 소스 코드 생성
//
// THREAD SAFETY
// 스레드 안전성
// -------------
// All functions in this file are thread-safe as they operate on immutable strings
// and don't use shared mutable state. Multiple goroutines can safely call these
// functions concurrently.
//
// 이 파일의 모든 함수는 불변 문자열에서 작동하고 공유 가변 상태를 사용하지
// 않으므로 스레드 안전합니다. 여러 고루틴이 이러한 함수를 동시에 안전하게
// 호출할 수 있습니다.
//
// Safe Concurrent Usage:
// 안전한 동시 사용:
//
//     go func() {
//         snake := stringutil.ToSnakeCase("UserProfile")
//     }()
//
//     go func() {
//         camel := stringutil.ToCamelCase("user_profile")
//     }()
//
//     // Both goroutines safe, no synchronization needed
//     // 두 고루틴 모두 안전, 동기화 불필요
//
// RELATED FILES
// 관련 파일
// -------------
// - manipulation.go: String manipulation operations (Truncate, Reverse, etc.)
//   manipulation.go: 문자열 조작 연산 (Truncate, Reverse 등)
// - validation.go: String validation functions (IsEmail, IsURL, etc.)
//   validation.go: 문자열 검증 함수 (IsEmail, IsURL 등)
// - formatting.go: Advanced formatting operations
//   formatting.go: 고급 포맷팅 연산
// - comparison.go: String comparison utilities
//   comparison.go: 문자열 비교 유틸리티
//
// =============================================================================

// ToSnakeCase converts a string to snake_case.
// ToSnakeCase는 문자열을 snake_case로 변환합니다.
//
// Handles multiple input formats:
// 여러 입력 형식 처리:
//   - PascalCase: "UserProfileData" → "user_profile_data"
//   - camelCase: "userProfileData" → "user_profile_data"
//   - kebab-case: "user-profile-data" → "user_profile_data"
//   - SCREAMING_SNAKE_CASE: "USER_PROFILE_DATA" → "user_profile_data"
//
// Example:
//
//	ToSnakeCase("UserProfileData")  // "user_profile_data"
//	ToSnakeCase("userProfileData")  // "user_profile_data"
//	ToSnakeCase("user-profile-data") // "user_profile_data"
func ToSnakeCase(s string) string {
	words := splitIntoWords(s)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return strings.Join(words, "_")
}

// ToCamelCase converts a string to camelCase.
// ToCamelCase는 문자열을 camelCase로 변환합니다.
//
// Example:
//
//	ToCamelCase("user_profile_data")  // "userProfileData"
//	ToCamelCase("user-profile-data")  // "userProfileData"
//	ToCamelCase("UserProfileData")    // "userProfileData"
func ToCamelCase(s string) string {
	words := splitIntoWords(s)
	if len(words) == 0 {
		return ""
	}

	// First word lowercase, rest capitalized
	// 첫 단어는 소문자, 나머지는 대문자로 시작
	result := strings.ToLower(words[0])
	for i := 1; i < len(words); i++ {
		if len(words[i]) > 0 {
			result += strings.ToUpper(string(words[i][0])) + strings.ToLower(words[i][1:])
		}
	}
	return result
}

// ToKebabCase converts a string to kebab-case.
// ToKebabCase는 문자열을 kebab-case로 변환합니다.
//
// Example:
//
//	ToKebabCase("UserProfileData")   // "user-profile-data"
//	ToKebabCase("user_profile_data") // "user-profile-data"
//	ToKebabCase("userProfileData")   // "user-profile-data"
func ToKebabCase(s string) string {
	words := splitIntoWords(s)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return strings.Join(words, "-")
}

// ToPascalCase converts a string to PascalCase.
// ToPascalCase는 문자열을 PascalCase로 변환합니다.
//
// Example:
//
//	ToPascalCase("user_profile_data") // "UserProfileData"
//	ToPascalCase("user-profile-data") // "UserProfileData"
//	ToPascalCase("userProfileData")   // "UserProfileData"
func ToPascalCase(s string) string {
	words := splitIntoWords(s)
	result := ""
	for _, word := range words {
		if len(word) > 0 {
			result += strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	return result
}

// ToScreamingSnakeCase converts a string to SCREAMING_SNAKE_CASE.
// ToScreamingSnakeCase는 문자열을 SCREAMING_SNAKE_CASE로 변환합니다.
//
// Example:
//
//	ToScreamingSnakeCase("UserProfileData") // "USER_PROFILE_DATA"
//	ToScreamingSnakeCase("userProfileData") // "USER_PROFILE_DATA"
func ToScreamingSnakeCase(s string) string {
	words := splitIntoWords(s)
	for i, word := range words {
		words[i] = strings.ToUpper(word)
	}
	return strings.Join(words, "_")
}

// ToTitle converts a string to Title Case (each word capitalized).
// ToTitle은 문자열을 Title Case로 변환합니다 (각 단어의 첫 글자를 대문자로).
//
// Example:
//
//	ToTitle("hello world")       // "Hello World"
//	ToTitle("user_profile_data") // "User Profile Data"
//	ToTitle("hello-world")       // "Hello World"
func ToTitle(s string) string {
	words := splitIntoWords(s)
	result := make([]string, len(words))
	for i, word := range words {
		if len(word) > 0 {
			result[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(result, " ")
}

// Slugify converts a string to a URL-friendly slug.
// Slugify는 문자열을 URL 친화적인 슬러그로 변환합니다.
//
// Converts to lowercase, replaces spaces and special characters with hyphens,
// and removes consecutive hyphens.
// 소문자로 변환하고, 공백과 특수 문자를 하이픈으로 대체하며,
// 연속된 하이픈을 제거합니다.
//
// Example:
//
//	Slugify("Hello World!")           // "hello-world"
//	Slugify("User Profile Data")      // "user-profile-data"
//	Slugify("Go Utils -- Package")    // "go-utils-package"
func Slugify(s string) string {
	// Convert to lowercase
	// 소문자로 변환
	s = strings.ToLower(s)

	// Replace spaces and special characters with hyphens
	// 공백과 특수 문자를 하이픈으로 대체
	var result []rune
	lastWasHyphen := false

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result = append(result, r)
			lastWasHyphen = false
		} else if !lastWasHyphen {
			result = append(result, '-')
			lastWasHyphen = true
		}
	}

	// Trim leading and trailing hyphens
	// 앞뒤 하이픈 제거
	slug := strings.Trim(string(result), "-")
	return slug
}

// Quote wraps a string in double quotes and escapes internal quotes.
// Quote는 문자열을 큰따옴표로 감싸고 내부 따옴표를 이스케이프합니다.
//
// Example:
//
//	Quote("hello")       // "\"hello\""
//	Quote("say \"hi\"")  // "\"say \\\"hi\\\"\""
func Quote(s string) string {
	// Use strconv.Quote for proper escaping
	// 적절한 이스케이프를 위해 strconv.Quote 사용
	var result strings.Builder
	result.WriteRune('"')
	for _, r := range s {
		if r == '"' || r == '\\' {
			result.WriteRune('\\')
		}
		result.WriteRune(r)
	}
	result.WriteRune('"')
	return result.String()
}

// Unquote removes surrounding quotes from a string and unescapes internal quotes.
// Unquote는 문자열에서 주변 따옴표를 제거하고 내부 따옴표의 이스케이프를 해제합니다.
//
// Supports both double quotes (") and single quotes (').
// 큰따옴표(")와 작은따옴표(') 모두 지원합니다.
//
// Example:
//
//	Unquote("\"hello\"")       // "hello"
//	Unquote("'world'")         // "world"
//	Unquote("\"say \\\"hi\\\"\"") // "say \"hi\""
func Unquote(s string) string {
	if len(s) < 2 {
		return s
	}

	// Check if string is quoted
	// 문자열이 따옴표로 감싸져 있는지 확인
	if (s[0] == '"' && s[len(s)-1] == '"') || (s[0] == '\'' && s[len(s)-1] == '\'') {
		s = s[1 : len(s)-1]
	}

	// Unescape internal quotes
	// 내부 따옴표 이스케이프 해제
	s = strings.ReplaceAll(s, "\\\"", "\"")
	s = strings.ReplaceAll(s, "\\'", "'")
	s = strings.ReplaceAll(s, "\\\\", "\\")

	return s
}

// splitIntoWords splits a string into words based on delimiters and case changes.
// splitIntoWords는 구분자와 케이스 변경을 기반으로 문자열을 단어로 분리합니다.
//
// Handles:
// 처리:
// - Delimiters: -, _, space
// 구분자: -, _, 공백
//   - Case changes: "UserProfile" → ["User", "Profile"]
//   - Consecutive uppercase: "HTTPServer" → ["HTTP", "Server"]
func splitIntoWords(s string) []string {
	if s == "" {
		return []string{}
	}

	var words []string
	var currentWord []rune

	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		r := runes[i]

		// Skip delimiters
		// 구분자 건너뛰기
		if r == '-' || r == '_' || r == ' ' {
			if len(currentWord) > 0 {
				words = append(words, string(currentWord))
				currentWord = []rune{}
			}
			continue
		}

		// Handle case changes
		// 케이스 변경 처리
		if unicode.IsUpper(r) && len(currentWord) > 0 {
			// Check if previous character was lowercase
			// 이전 문자가 소문자였는지 확인
			if unicode.IsLower(currentWord[len(currentWord)-1]) {
				words = append(words, string(currentWord))
				currentWord = []rune{r}
				continue
			}

			// Check if next character is lowercase (e.g., "HTTPServer" → "HTTP" "Server")
			// 다음 문자가 소문자인지 확인 (예: "HTTPServer" → "HTTP" "Server")
			if i+1 < len(runes) && unicode.IsLower(runes[i+1]) && len(currentWord) > 0 {
				words = append(words, string(currentWord))
				currentWord = []rune{r}
				continue
			}
		}

		currentWord = append(currentWord, r)
	}

	if len(currentWord) > 0 {
		words = append(words, string(currentWord))
	}

	return words
}
