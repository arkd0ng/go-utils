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

	// Builder Pattern Examples / 빌더 패턴 예제
	logger.Info("=== Builder Pattern Examples ===")
	logger.Info("=== 빌더 패턴 예제 ===")
	logger.Info("")

	builderResult1 := stringutil.NewBuilder().
		Append("  user profile data  ").
		Clean().
		ToSnakeCase().
		ToUpper().
		Build()
	logger.Info(fmt.Sprintf("Builder chain 1: %s", builderResult1))

	builderResult2 := stringutil.NewBuilder().
		Append("Hello World").
		ToKebabCase().
		Quote().
		Build()
	logger.Info(fmt.Sprintf("Builder chain 2: %s", builderResult2))

	builderResult3 := stringutil.NewBuilderWithString("the quick brown fox jumps over the lazy dog").
		Capitalize().
		Truncate(30).
		Build()
	logger.Info(fmt.Sprintf("Builder chain 3: %s", builderResult3))

	// Builder with multiple operations / 여러 작업을 연결한 빌더
	complexBuilder := stringutil.NewBuilder().
		AppendLine("Line 1: User Profile").
		AppendLine("Line 2: Data Processing").
		ToLower().
		Replace(":", " →").
		Build()
	logger.Info(fmt.Sprintf("Complex builder result:\n%s", complexBuilder))
	logger.Info("")

	// Encoding/Decoding Examples / 인코딩/디코딩 예제
	logger.Info("=== Encoding/Decoding Examples ===")
	logger.Info("=== 인코딩/디코딩 예제 ===")
	logger.Info("")

	// Base64
	plainText := "Hello, 안녕하세요!"
	base64Encoded := stringutil.Base64Encode(plainText)
	logger.Info(fmt.Sprintf("Base64Encode('%s'): %s", plainText, base64Encoded))
	base64Decoded, _ := stringutil.Base64Decode(base64Encoded)
	logger.Info(fmt.Sprintf("Base64Decode(encoded): %s", base64Decoded))

	// Base64URL
	urlText := "hello?world=test&foo=bar"
	base64URLEncoded := stringutil.Base64URLEncode(urlText)
	logger.Info(fmt.Sprintf("Base64URLEncode('%s'): %s", urlText, base64URLEncoded))
	base64URLDecoded, _ := stringutil.Base64URLDecode(base64URLEncoded)
	logger.Info(fmt.Sprintf("Base64URLDecode(encoded): %s", base64URLDecoded))

	// URL Encoding
	urlParam := "hello world & foo=bar"
	urlEncoded := stringutil.URLEncode(urlParam)
	logger.Info(fmt.Sprintf("URLEncode('%s'): %s", urlParam, urlEncoded))
	urlDecoded, _ := stringutil.URLDecode(urlEncoded)
	logger.Info(fmt.Sprintf("URLDecode(encoded): %s", urlDecoded))

	// HTML Escaping
	htmlText := "<div>Hello & \"World\"</div>"
	htmlEscaped := stringutil.HTMLEscape(htmlText)
	logger.Info(fmt.Sprintf("HTMLEscape('%s'): %s", htmlText, htmlEscaped))
	htmlUnescaped := stringutil.HTMLUnescape(htmlEscaped)
	logger.Info(fmt.Sprintf("HTMLUnescape(escaped): %s", htmlUnescaped))
	logger.Info("")

	// String Distance/Similarity Examples / 문자열 거리/유사도 예제
	logger.Info("=== String Distance/Similarity Examples ===")
	logger.Info("=== 문자열 거리/유사도 예제 ===")
	logger.Info("")

	str1, str2 := "kitten", "sitting"
	levDist := stringutil.LevenshteinDistance(str1, str2)
	logger.Info(fmt.Sprintf("LevenshteinDistance('%s', '%s'): %d", str1, str2, levDist))

	sim := stringutil.Similarity(str1, str2)
	logger.Info(fmt.Sprintf("Similarity('%s', '%s'): %.3f", str1, str2, sim))

	str3, str4 := "karolin", "kathrin"
	hammingDist := stringutil.HammingDistance(str3, str4)
	logger.Info(fmt.Sprintf("HammingDistance('%s', '%s'): %d", str3, str4, hammingDist))

	str5, str6 := "martha", "marhta"
	jaroSim := stringutil.JaroWinklerSimilarity(str5, str6)
	logger.Info(fmt.Sprintf("JaroWinklerSimilarity('%s', '%s'): %.3f", str5, str6, jaroSim))

	// Practical similarity use case / 실용적인 유사도 사용 예
	searchTerm := "golang"
	candidates := []string{"Go Language", "golang tutorial", "Python", "Java"}
	logger.Info(fmt.Sprintf("\nSearch term: %s", searchTerm))
	logger.Info("Similarity scores with candidates:")
	for _, candidate := range candidates {
		score := stringutil.Similarity(strings.ToLower(searchTerm), strings.ToLower(candidate))
		logger.Info(fmt.Sprintf("  '%s': %.3f", candidate, score))
	}
	logger.Info("")

	// Formatting Examples / 포맷팅 예제
	logger.Info("=== Formatting Examples ===")
	logger.Info("=== 포맷팅 예제 ===")
	logger.Info("")

	// Number formatting
	logger.Info(fmt.Sprintf("FormatNumber(1000000, ','): %s", stringutil.FormatNumber(1000000, ",")))
	logger.Info(fmt.Sprintf("FormatNumber(1234567, '.'): %s", stringutil.FormatNumber(1234567, ".")))
	logger.Info(fmt.Sprintf("FormatNumber(-1000000, ','): %s", stringutil.FormatNumber(-1000000, ",")))

	// Bytes formatting
	logger.Info(fmt.Sprintf("FormatBytes(1024): %s", stringutil.FormatBytes(1024)))
	logger.Info(fmt.Sprintf("FormatBytes(1536): %s", stringutil.FormatBytes(1536)))
	logger.Info(fmt.Sprintf("FormatBytes(1048576): %s", stringutil.FormatBytes(1048576)))
	logger.Info(fmt.Sprintf("FormatBytes(1073741824): %s", stringutil.FormatBytes(1073741824)))

	// Pluralization
	logger.Info(fmt.Sprintf("FormatWithCount(1, 'item', 'items'): %s", stringutil.FormatWithCount(1, "item", "items")))
	logger.Info(fmt.Sprintf("FormatWithCount(5, 'item', 'items'): %s", stringutil.FormatWithCount(5, "item", "items")))
	logger.Info(fmt.Sprintf("FormatWithCount(1, 'person', 'people'): %s", stringutil.FormatWithCount(1, "person", "people")))
	logger.Info(fmt.Sprintf("FormatWithCount(10, 'person', 'people'): %s", stringutil.FormatWithCount(10, "person", "people")))

	// Ellipsis
	longFilename := "verylongfilename.txt"
	logger.Info(fmt.Sprintf("Ellipsis('%s', 15): %s", longFilename, stringutil.Ellipsis(longFilename, 15)))

	// Masking
	logger.Info(fmt.Sprintf("Mask('1234567890', 2, 2, '*'): %s", stringutil.Mask("1234567890", 2, 2, "*")))
	logger.Info(fmt.Sprintf("MaskEmail('john.doe@example.com'): %s", stringutil.MaskEmail("john.doe@example.com")))
	logger.Info(fmt.Sprintf("MaskCreditCard('1234567890123456'): %s", stringutil.MaskCreditCard("1234567890123456")))
	logger.Info(fmt.Sprintf("MaskCreditCard('1234-5678-9012-3456'): %s", stringutil.MaskCreditCard("1234-5678-9012-3456")))

	// Line numbers
	multiLineText := "line 1\nline 2\nline 3"
	logger.Info(fmt.Sprintf("AddLineNumbers:\n%s", stringutil.AddLineNumbers(multiLineText)))

	// Indentation
	codeSnippet := "func main() {\n  fmt.Println(\"hello\")\n}"
	logger.Info(fmt.Sprintf("Indent (2 spaces):\n%s", stringutil.Indent(codeSnippet, "  ")))

	indentedCode := "    func main() {\n      fmt.Println(\"hello\")\n    }"
	logger.Info(fmt.Sprintf("Dedent:\n%s", stringutil.Dedent(indentedCode)))

	// Text wrapping
	longText := "The quick brown fox jumps over the lazy dog"
	logger.Info(fmt.Sprintf("WrapText (width 20):\n%s", stringutil.WrapText(longText, 20)))
	logger.Info("")

	// Real-world Scenarios / 실제 사용 시나리오
	logger.Info("=== Real-world Scenarios ===")
	logger.Info("=== 실제 사용 시나리오 ===")
	logger.Info("")

	// Scenario 1: Processing user input for database / 시나리오 1: 데이터베이스용 사용자 입력 처리
	logger.Info("Scenario 1: Clean user input for database")
	userInput := "  John DOE  "
	processed := stringutil.NewBuilder().
		Append(userInput).
		Clean().
		ToTitle().
		Build()
	logger.Info(fmt.Sprintf("  Raw input: '%s'", userInput))
	logger.Info(fmt.Sprintf("  Processed: '%s'", processed))

	// Scenario 2: Generate URL-friendly slug / 시나리오 2: URL 친화적 슬러그 생성
	logger.Info("\nScenario 2: Generate URL-friendly slug from title")
	articleTitle := "How to Use Go Utils: A Complete Guide!"
	slug := stringutil.Slugify(articleTitle)
	logger.Info(fmt.Sprintf("  Title: '%s'", articleTitle))
	logger.Info(fmt.Sprintf("  Slug: '%s'", slug))

	// Scenario 3: Format API response / 시나리오 3: API 응답 포맷
	logger.Info("\nScenario 3: Format API response")
	filesFound := 42
	responseMsg := fmt.Sprintf("Found %s", stringutil.FormatWithCount(filesFound, "file", "files"))
	logger.Info(fmt.Sprintf("  Message: %s", responseMsg))

	// Scenario 4: Mask sensitive data in logs / 시나리오 4: 로그에서 민감한 데이터 마스크
	logger.Info("\nScenario 4: Mask sensitive data in logs")
	email := "sensitive.user@example.com"
	creditCard := "1234-5678-9012-3456"
	logger.Info(fmt.Sprintf("  Email: %s", stringutil.MaskEmail(email)))
	logger.Info(fmt.Sprintf("  Credit Card: %s", stringutil.MaskCreditCard(creditCard)))

	// Scenario 5: Find similar strings (typo correction) / 시나리오 5: 유사한 문자열 찾기 (오타 수정)
	logger.Info("\nScenario 5: Find similar strings (typo correction)")
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
	// Sort by score (simple bubble sort for demo)
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

	logger.Info("=== All examples completed successfully! ===")
	logger.Info("=== 모든 예제가 성공적으로 완료되었습니다! ===")
}
