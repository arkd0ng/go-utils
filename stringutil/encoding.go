package stringutil

import (
	"encoding/base64"
	"html"
	"net/url"
)

// =============================================================================
// File: encoding.go
// Purpose: String Encoding and Decoding Operations
// 파일: encoding.go
// 목적: 문자열 인코딩 및 디코딩 연산
// =============================================================================
//
// OVERVIEW
// 개요
// --------
// The encoding.go file provides essential encoding and decoding functions for
// common data interchange formats. These functions handle Base64 encoding (both
// standard and URL-safe variants), URL encoding/decoding, and HTML entity
// escaping/unescaping. These operations are crucial for data transmission,
// storage, and display in web applications, APIs, and data processing pipelines.
//
// encoding.go 파일은 일반적인 데이터 교환 형식을 위한 필수 인코딩 및 디코딩
// 함수를 제공합니다. 이러한 함수는 Base64 인코딩 (표준 및 URL 안전 변형 모두),
// URL 인코딩/디코딩, HTML 엔티티 이스케이프/언이스케이프를 처리합니다. 이러한
// 연산은 웹 애플리케이션, API 및 데이터 처리 파이프라인에서 데이터 전송, 저장
// 및 표시에 중요합니다.
//
// DESIGN PHILOSOPHY
// 설계 철학
// -----------------
// 1. **Standard Compliance**: Follow well-established encoding standards
//    **표준 준수**: 잘 확립된 인코딩 표준 따름
//
// 2. **Bidirectional Operations**: Provide both encode and decode for each format
//    **양방향 연산**: 각 형식에 대해 인코딩 및 디코딩 모두 제공
//
// 3. **Error Handling**: Return errors for invalid input during decoding
//    **오류 처리**: 디코딩 중 잘못된 입력에 대해 오류 반환
//
// 4. **Context-Specific Encoding**: Different Base64 variants for different use cases
//    **컨텍스트 특정 인코딩**: 다른 사용 사례를 위한 다른 Base64 변형
//
// 5. **Safety-First**: HTML escape prevents XSS, URL encode prevents injection
//    **안전 우선**: HTML 이스케이프는 XSS 방지, URL 인코딩은 주입 방지
//
// FUNCTION CATEGORIES
// 함수 범주
// -------------------
//
// 1. BASE64 ENCODING (STANDARD) (Base64 인코딩 (표준))
//    - Base64Encode: Encode string to standard Base64
//      Base64Encode: 문자열을 표준 Base64로 인코딩
//    - Base64Decode: Decode standard Base64 to string
//      Base64Decode: 표준 Base64를 문자열로 디코딩
//
// 2. BASE64 ENCODING (URL-SAFE) (Base64 인코딩 (URL 안전))
//    - Base64URLEncode: Encode string to URL-safe Base64
//      Base64URLEncode: 문자열을 URL 안전 Base64로 인코딩
//    - Base64URLDecode: Decode URL-safe Base64 to string
//      Base64URLDecode: URL 안전 Base64를 문자열로 디코딩
//
// 3. URL ENCODING (URL 인코딩)
//    - URLEncode: Encode string for URL query parameters
//      URLEncode: URL 쿼리 매개변수용 문자열 인코딩
//    - URLDecode: Decode URL-encoded string
//      URLDecode: URL 인코딩된 문자열 디코딩
//
// 4. HTML ENTITY ENCODING (HTML 엔티티 인코딩)
//    - HTMLEscape: Escape special HTML characters
//      HTMLEscape: 특수 HTML 문자 이스케이프
//    - HTMLUnescape: Unescape HTML entities
//      HTMLUnescape: HTML 엔티티 언이스케이프
//
// KEY OPERATIONS SUMMARY
// 주요 연산 요약
// ----------------------
//
// Base64Encode(s string) string
// - Purpose: Encode string to standard Base64 format
// - 목적: 문자열을 표준 Base64 형식으로 인코딩
// - Alphabet: A-Z, a-z, 0-9, +, /
// - 알파벳: A-Z, a-z, 0-9, +, /
// - Padding: Uses '=' for padding
// - 패딩: 패딩에 '=' 사용
// - Time Complexity: O(n) where n is input length
// - 시간 복잡도: O(n), n은 입력 길이
// - Space Complexity: O(n) - output is ~4/3 of input size
// - 공간 복잡도: O(n) - 출력은 입력 크기의 ~4/3
// - Use Cases: Email attachments, data storage, API tokens, binary data transmission
// - 사용 사례: 이메일 첨부, 데이터 저장, API 토큰, 바이너리 데이터 전송
//
// Base64Decode(s string) (string, error)
// - Purpose: Decode standard Base64 to original string
// - 목적: 표준 Base64를 원본 문자열로 디코딩
// - Error Handling: Returns error if input is not valid Base64
// - 오류 처리: 입력이 유효한 Base64가 아니면 오류 반환
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n) - output is ~3/4 of input size
// - 공간 복잡도: O(n) - 출력은 입력 크기의 ~3/4
// - Validation: Checks for valid Base64 characters and padding
// - 검증: 유효한 Base64 문자 및 패딩 확인
// - Use Cases: Decode Base64-encoded data, parse tokens, retrieve binary data
// - 사용 사례: Base64 인코딩 데이터 디코딩, 토큰 파싱, 바이너리 데이터 검색
//
// Base64URLEncode(s string) string
// - Purpose: Encode string to URL-safe Base64 format
// - 목적: 문자열을 URL 안전 Base64 형식으로 인코딩
// - Alphabet: A-Z, a-z, 0-9, -, _ (replaces +, /)
// - 알파벳: A-Z, a-z, 0-9, -, _ (+, / 대체)
// - URL-Safe: Can be used in URLs without additional encoding
// - URL 안전: 추가 인코딩 없이 URL에서 사용 가능
// - Padding: Uses '=' for padding (can be omitted in some contexts)
// - 패딩: 패딩에 '=' 사용 (일부 컨텍스트에서 생략 가능)
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Use Cases: JWT tokens, URL parameters, RESTful API identifiers
// - 사용 사례: JWT 토큰, URL 매개변수, RESTful API 식별자
//
// Base64URLDecode(s string) (string, error)
// - Purpose: Decode URL-safe Base64 to original string
// - 목적: URL 안전 Base64를 원본 문자열로 디코딩
// - Error Handling: Returns error if input is not valid URL-safe Base64
// - 오류 처리: 입력이 유효한 URL 안전 Base64가 아니면 오류 반환
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Use Cases: Parse JWT tokens, decode URL-encoded identifiers
// - 사용 사례: JWT 토큰 파싱, URL 인코딩된 식별자 디코딩
//
// URLEncode(s string) string
// - Purpose: Encode string for use in URL query parameters
// - 목적: URL 쿼리 매개변수에서 사용할 문자열 인코딩
// - Encoding: Percent-encoding (RFC 3986)
// - 인코딩: 퍼센트 인코딩 (RFC 3986)
// - Special Characters: Encodes spaces as '+', others as %XX
// - 특수 문자: 공백을 '+'로 인코딩, 다른 것은 %XX로
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n) - output can be up to 3x input size
// - 공간 복잡도: O(n) - 출력은 입력 크기의 최대 3배
// - Use Cases: URL query parameters, form submissions, HTTP GET requests
// - 사용 사례: URL 쿼리 매개변수, 폼 제출, HTTP GET 요청
//
// URLDecode(s string) (string, error)
// - Purpose: Decode URL-encoded string
// - 목적: URL 인코딩된 문자열 디코딩
// - Decoding: Converts %XX to characters, '+' to space
// - 디코딩: %XX를 문자로 변환, '+'를 공백으로
// - Error Handling: Returns error if invalid percent-encoding
// - 오류 처리: 잘못된 퍼센트 인코딩이면 오류 반환
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Use Cases: Parse URL query parameters, decode form submissions
// - 사용 사례: URL 쿼리 매개변수 파싱, 폼 제출 디코딩
//
// HTMLEscape(s string) string
// - Purpose: Escape special HTML characters to prevent XSS attacks
// - 목적: XSS 공격 방지를 위해 특수 HTML 문자 이스케이프
// - Escaped Characters: <, >, &, ", ' → &lt;, &gt;, &amp;, &#34;, &#39;
// - 이스케이프 문자: <, >, &, ", ' → &lt;, &gt;, &amp;, &#34;, &#39;
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n) - output can be larger than input
// - 공간 복잡도: O(n) - 출력이 입력보다 클 수 있음
// - Security: Critical for preventing Cross-Site Scripting (XSS)
// - 보안: 크로스 사이트 스크립팅 (XSS) 방지에 중요
// - Use Cases: Display user-generated content, HTML templates, web security
// - 사용 사례: 사용자 생성 콘텐츠 표시, HTML 템플릿, 웹 보안
//
// HTMLUnescape(s string) string
// - Purpose: Convert HTML entities back to original characters
// - 목적: HTML 엔티티를 원본 문자로 다시 변환
// - Unescaped Entities: &lt;, &gt;, &amp;, &#34;, &#39; → <, >, &, ", '
// - 언이스케이프 엔티티: &lt;, &gt;, &amp;, &#34;, &#39; → <, >, &, ", '
// - Named Entities: Supports common HTML entities (&nbsp;, &copy;, etc.)
// - 명명된 엔티티: 일반적인 HTML 엔티티 지원 (&nbsp;, &copy; 등)
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Use Cases: Parse HTML content, decode stored data, text processing
// - 사용 사례: HTML 콘텐츠 파싱, 저장된 데이터 디코딩, 텍스트 처리
//
// PERFORMANCE CHARACTERISTICS
// 성능 특성
// ---------------------------
//
// Time Complexities:
// 시간 복잡도:
// - All encode/decode functions: O(n) - single pass through input
//   모든 인코딩/디코딩 함수: O(n) - 입력 단일 패스
//
// Space Complexities:
// 공간 복잡도:
// - Base64Encode: O(n) - output ~133% of input (4/3 ratio)
//   Base64Encode: O(n) - 출력 입력의 ~133% (4/3 비율)
// - Base64Decode: O(n) - output ~75% of input (3/4 ratio)
//   Base64Decode: O(n) - 출력 입력의 ~75% (3/4 비율)
// - URLEncode: O(n) - output up to 300% of input (worst case)
//   URLEncode: O(n) - 출력 입력의 최대 300% (최악의 경우)
// - URLDecode: O(n) - output ~33%-100% of input
//   URLDecode: O(n) - 출력 입력의 ~33%-100%
// - HTMLEscape: O(n) - output up to 600% of input (worst case: all &amp;)
//   HTMLEscape: O(n) - 출력 입력의 최대 600% (최악의 경우: 모두 &amp;)
// - HTMLUnescape: O(n) - output smaller than input
//   HTMLUnescape: O(n) - 출력이 입력보다 작음
//
// Optimization Tips:
// 최적화 팁:
// 1. For large data, consider streaming encoding/decoding
//    큰 데이터의 경우 스트리밍 인코딩/디코딩 고려
// 2. Base64 adds ~33% overhead - consider alternatives for large files
//    Base64는 ~33% 오버헤드 추가 - 큰 파일의 경우 대안 고려
// 3. URLEncode worst case: all special characters (3x expansion)
//    URLEncode 최악의 경우: 모든 특수 문자 (3배 확장)
// 4. HTMLEscape only when displaying user content
//    사용자 콘텐츠 표시 시에만 HTMLEscape
// 5. Reuse buffers for repeated encoding/decoding operations
//    반복적인 인코딩/디코딩 연산에 버퍼 재사용
//
// ENCODING FORMAT DIFFERENCES
// 인코딩 형식 차이
// ----------------------------
//
// Standard Base64 vs URL-Safe Base64:
// 표준 Base64 vs URL 안전 Base64:
// - Standard: Uses + and / (62nd and 63rd characters)
//   표준: + 및 / 사용 (62번째 및 63번째 문자)
// - URL-Safe: Uses - and _ instead (safe in URLs without encoding)
//   URL 안전: 대신 - 및 _ 사용 (인코딩 없이 URL에서 안전)
// - Example: Standard "a+b/c==" → URL-Safe "a-b_c=="
//   예: 표준 "a+b/c==" → URL 안전 "a-b_c=="
// - Use Standard for: Email, file storage, general data
//   표준 사용: 이메일, 파일 저장, 일반 데이터
// - Use URL-Safe for: URL parameters, JWT tokens, filenames
//   URL 안전 사용: URL 매개변수, JWT 토큰, 파일 이름
//
// URL Encoding Specifics:
// URL 인코딩 세부사항:
// - Space: Encoded as '+' (or %20 in some contexts)
//   공백: '+'로 인코딩 (일부 컨텍스트에서 %20)
// - Reserved Characters: !, *, ', (, ), ;, :, @, &, =, +, $, ,, /, ?, #, [, ]
//   예약 문자: !, *, ', (, ), ;, :, @, &, =, +, $, ,, /, ?, #, [, ]
// - Safe Characters: A-Z, a-z, 0-9, -, _, ., ~
//   안전 문자: A-Z, a-z, 0-9, -, _, ., ~
// - Example: "hello world!" → "hello+world%21"
//   예: "hello world!" → "hello+world%21"
//
// HTML Escaping Rules:
// HTML 이스케이프 규칙:
// - Five XML entities: < > & " '
//   5개 XML 엔티티: < > & " '
// - &lt; and &gt; prevent tag injection
//   &lt; 및 &gt;는 태그 주입 방지
// - &amp; prevents entity confusion
//   &amp;는 엔티티 혼동 방지
// - &#34; and &#39; prevent attribute injection
//   &#34; 및 &#39;는 속성 주입 방지
//
// EDGE CASES AND SPECIAL BEHAVIORS
// 엣지 케이스 및 특수 동작
// ---------------------------------
//
// Empty Strings:
// 빈 문자열:
// - Base64Encode(""): "" (empty Base64)
//   Base64Encode(""): "" (빈 Base64)
// - URLEncode(""): ""
//   URLEncode(""): ""
// - HTMLEscape(""): ""
//   HTMLEscape(""): ""
// - All decode functions: ("", nil) or ("", error) depending on validity
//   모든 디코딩 함수: 유효성에 따라 ("", nil) 또는 ("", error)
//
// Invalid Input (Decoding):
// 잘못된 입력 (디코딩):
// - Base64Decode("invalid!"): ("", error) - invalid characters
//   Base64Decode("invalid!"): ("", error) - 잘못된 문자
// - URLDecode("%ZZ"): ("", error) - invalid percent encoding
//   URLDecode("%ZZ"): ("", error) - 잘못된 퍼센트 인코딩
// - HTMLUnescape handles invalid entities gracefully (returns as-is)
//   HTMLUnescape는 잘못된 엔티티를 우아하게 처리 (그대로 반환)
//
// Padding:
// 패딩:
// - Base64: Uses '=' for padding to make length multiple of 4
//   Base64: 길이를 4의 배수로 만들기 위해 '=' 패딩 사용
// - Example: "hello" → "aGVsbG8=" (1 padding char)
//   예: "hello" → "aGVsbG8=" (1개 패딩 문자)
// - URL-Safe Base64: Padding optional in some contexts (JWT uses no padding)
//   URL 안전 Base64: 일부 컨텍스트에서 패딩 선택 사항 (JWT는 패딩 없음)
//
// Unicode Handling:
// 유니코드 처리:
// - All functions handle Unicode correctly (UTF-8 encoding)
//   모든 함수는 유니코드를 올바르게 처리 (UTF-8 인코딩)
// - Base64Encode("你好"): Encodes UTF-8 bytes
//   Base64Encode("你好"): UTF-8 바이트 인코딩
// - URLEncode("你好"): Percent-encodes UTF-8 bytes
//   URLEncode("你好"): UTF-8 바이트 퍼센트 인코딩
//
// COMMON USAGE PATTERNS
// 일반 사용 패턴
// ---------------------
//
// 1. API Token Generation
//    API 토큰 생성:
//
//    token := generateRandomBytes(32)
//    encodedToken := stringutil.Base64URLEncode(string(token))
//    // Use in JWT or API authentication
//    // JWT 또는 API 인증에 사용
//
// 2. URL Query Parameter with Special Characters
//    특수 문자가 있는 URL 쿼리 매개변수:
//
//    searchTerm := "hello world & goodbye"
//    url := "https://example.com/search?q=" + stringutil.URLEncode(searchTerm)
//    // "https://example.com/search?q=hello+world+%26+goodbye"
//
// 3. Displaying User-Generated Content (XSS Prevention)
//    사용자 생성 콘텐츠 표시 (XSS 방지):
//
//    userInput := "<script>alert('XSS')</script>"
//    safeHTML := stringutil.HTMLEscape(userInput)
//    // "&lt;script&gt;alert(&#39;XSS&#39;)&lt;/script&gt;"
//    // Safe to display in HTML
//    // HTML에 표시하기에 안전
//
// 4. Email Attachment Encoding
//    이메일 첨부 인코딩:
//
//    fileContent := readFileContent("document.pdf")
//    base64Content := stringutil.Base64Encode(string(fileContent))
//    // Send in email as MIME attachment
//    // MIME 첨부로 이메일에 전송
//
// 5. Decoding URL Parameters
//    URL 매개변수 디코딩:
//
//    encodedParam := "hello+world%21"
//    decoded, err := stringutil.URLDecode(encodedParam)
//    if err != nil {
//        log.Fatal(err)
//    }
//    // "hello world!"
//
// 6. JWT Token Encoding
//    JWT 토큰 인코딩:
//
//    header := `{"alg":"HS256","typ":"JWT"}`
//    payload := `{"sub":"1234567890","name":"John Doe"}`
//    encodedHeader := stringutil.Base64URLEncode(header)
//    encodedPayload := stringutil.Base64URLEncode(payload)
//    token := encodedHeader + "." + encodedPayload + "." + signature
//    // Create JWT token
//    // JWT 토큰 생성
//
// 7. HTML Template Data Preparation
//    HTML 템플릿 데이터 준비:
//
//    username := `<admin>`
//    safeUsername := stringutil.HTMLEscape(username)
//    template := fmt.Sprintf("<div>Welcome, %s!</div>", safeUsername)
//    // "<div>Welcome, &lt;admin&gt;!</div>"
//    // Prevents tag injection
//    // 태그 주입 방지
//
// 8. Decoding Stored Base64 Data
//    저장된 Base64 데이터 디코딩:
//
//    storedData := "aGVsbG8gd29ybGQ="
//    original, err := stringutil.Base64Decode(storedData)
//    if err != nil {
//        log.Fatal(err)
//    }
//    // "hello world"
//
// 9. Building Search URLs
//    검색 URL 구축:
//
//    query := "golang tutorial"
//    category := "programming & development"
//    url := fmt.Sprintf("https://example.com/search?q=%s&cat=%s",
//        stringutil.URLEncode(query),
//        stringutil.URLEncode(category))
//    // Safe URL with encoded parameters
//    // 인코딩된 매개변수가 있는 안전한 URL
//
// 10. Parsing HTML Content
//     HTML 콘텐츠 파싱:
//
//     htmlContent := "&lt;p&gt;Hello &amp; goodbye&lt;/p&gt;"
//     plainText := stringutil.HTMLUnescape(htmlContent)
//     // "<p>Hello & goodbye</p>"
//     // Extract text from HTML entities
//     // HTML 엔티티에서 텍스트 추출
//
// SECURITY CONSIDERATIONS
// 보안 고려사항
// -----------------------
//
// XSS Prevention:
// XSS 방지:
// - ALWAYS use HTMLEscape when displaying user input in HTML
//   HTML에 사용자 입력 표시 시 항상 HTMLEscape 사용
// - Never trust user input
//   사용자 입력 절대 신뢰 안 함
// - Escape both content and attributes
//   콘텐츠 및 속성 모두 이스케이프
//
// URL Injection Prevention:
// URL 주입 방지:
// - Use URLEncode for all user-provided query parameters
//   모든 사용자 제공 쿼리 매개변수에 URLEncode 사용
// - Prevents SQL injection via URL parameters
//   URL 매개변수를 통한 SQL 주입 방지
//
// Base64 Is Not Encryption:
// Base64는 암호화가 아님:
// - Base64 is encoding, NOT encryption
//   Base64는 인코딩이지 암호화 아님
// - Anyone can decode Base64
//   누구나 Base64 디코딩 가능
// - Do NOT use for sensitive data without encryption
//   암호화 없이 민감한 데이터에 사용 안 함
//
// Validation After Decoding:
// 디코딩 후 검증:
// - Always check errors from decode functions
//   디코딩 함수의 오류 항상 확인
// - Validate decoded content before use
//   사용 전에 디코딩된 콘텐츠 검증
//
// COMPARISON WITH RELATED FUNCTIONS
// 관련 함수와의 비교
// ---------------------------------
//
// Base64Encode vs base64.StdEncoding.EncodeToString
// - Base64Encode: Convenience wrapper
//   Base64Encode: 편의 래퍼
// - base64.StdEncoding: Direct stdlib use
//   base64.StdEncoding: 직접 stdlib 사용
// - Use Base64Encode for: Cleaner code
//   Base64Encode 사용: 더 깔끔한 코드
//
// Base64URLEncode vs Base64Encode
// - Base64URLEncode: URL-safe (-, _)
//   Base64URLEncode: URL 안전 (-, _)
// - Base64Encode: Standard (+, /)
//   Base64Encode: 표준 (+, /)
// - Use Base64URLEncode for: URLs, JWT, filenames
//   Base64URLEncode 사용: URL, JWT, 파일 이름
//
// URLEncode vs url.QueryEscape
// - URLEncode: Convenience wrapper
//   URLEncode: 편의 래퍼
// - url.QueryEscape: Direct stdlib use
//   url.QueryEscape: 직접 stdlib 사용
// - Same functionality, different API
//   동일한 기능, 다른 API
//
// HTMLEscape vs template.HTMLEscapeString
// - HTMLEscape: Escape function
//   HTMLEscape: 이스케이프 함수
// - template.HTMLEscapeString: Template context
//   template.HTMLEscapeString: 템플릿 컨텍스트
// - Use HTMLEscape for: General text escaping
//   HTMLEscape 사용: 일반 텍스트 이스케이프
//
// THREAD SAFETY
// 스레드 안전성
// -------------
// All functions in this file are thread-safe as they operate on immutable strings
// and use thread-safe standard library functions.
//
// 이 파일의 모든 함수는 불변 문자열에서 작동하고 스레드 안전한 표준 라이브러리
// 함수를 사용하므로 스레드 안전합니다.
//
// Safe Concurrent Usage:
// 안전한 동시 사용:
//
//     go func() {
//         encoded := stringutil.Base64Encode(data)
//     }()
//
//     go func() {
//         escaped := stringutil.HTMLEscape(userInput)
//     }()
//
//     // All encoding functions safe for concurrent use
//     // 모든 인코딩 함수는 동시 사용에 안전
//
// RELATED FILES
// 관련 파일
// -------------
// - validation.go: String validation (IsEmail, IsURL, etc.)
//   validation.go: 문자열 검증 (IsEmail, IsURL 등)
// - unicode.go: Unicode handling operations
//   unicode.go: 유니코드 처리 연산
// - utils.go: General utilities
//   utils.go: 일반 유틸리티
//
// =============================================================================

// Base64Encode encodes a string to base64.
// Base64Encode는 문자열을 base64로 인코딩합니다.
//
// Example
// 예제:
//
//	Base64Encode("hello")  // "aGVsbG8="
func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Base64Decode decodes a base64 string.
// Base64Decode는 base64 문자열을 디코딩합니다.
//
// Returns an error if the input is not valid base64.
// 입력이 유효한 base64가 아니면 에러를 반환합니다.
//
// Example
// 예제:
//
//	Base64Decode("aGVsbG8=")  // "hello", nil
//	Base64Decode("invalid!")  // "", error
func Base64Decode(s string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

// Base64URLEncode encodes a string to URL-safe base64.
// Base64URLEncode는 문자열을 URL 안전 base64로 인코딩합니다.
//
// URL-safe base64 uses '-' and '_' instead of '+' and '/'.
// URL 안전 base64는 '+'와 '/' 대신 '-'와 '_'를 사용합니다.
//
// Example
// 예제:
//
//	Base64URLEncode("hello?world")  // "aGVsbG8_d29ybGQ="
func Base64URLEncode(s string) string {
	return base64.URLEncoding.EncodeToString([]byte(s))
}

// Base64URLDecode decodes a URL-safe base64 string.
// Base64URLDecode는 URL 안전 base64 문자열을 디코딩합니다.
//
// Returns an error if the input is not valid URL-safe base64.
// 입력이 유효한 URL 안전 base64가 아니면 에러를 반환합니다.
//
// Example
// 예제:
//
//	Base64URLDecode("aGVsbG8_d29ybGQ=")  // "hello?world", nil
func Base64URLDecode(s string) (string, error) {
	decoded, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

// URLEncode encodes a string for safe use in URLs.
// URLEncode는 URL에서 안전하게 사용하기 위해 문자열을 인코딩합니다.
//
// Example
// 예제:
//
//	URLEncode("hello world")  // "hello+world"
//	URLEncode("hello/world")  // "hello%2Fworld"
func URLEncode(s string) string {
	return url.QueryEscape(s)
}

// URLDecode decodes a URL-encoded string.
// URLDecode는 URL 인코딩된 문자열을 디코딩합니다.
//
// Returns an error if the input is not valid URL encoding.
// 입력이 유효한 URL 인코딩이 아니면 에러를 반환합니다.
//
// Example
// 예제:
//
//	URLDecode("hello+world")  // "hello world", nil
//	URLDecode("hello%2Fworld")  // "hello/world", nil
func URLDecode(s string) (string, error) {
	return url.QueryUnescape(s)
}

// HTMLEscape escapes special HTML characters.
// HTMLEscape는 특수 HTML 문자를 이스케이프합니다.
//
// Escapes: <, >, &, ", '
// 이스케이프: <, >, &, ", '
//
// Example
// 예제:
//
//	HTMLEscape("<div>hello</div>")  // "&lt;div&gt;hello&lt;/div&gt;"
//	HTMLEscape("'hello' & \"world\"")  // "&#39;hello&#39; &amp; &#34;world&#34;"
func HTMLEscape(s string) string {
	return html.EscapeString(s)
}

// HTMLUnescape unescapes HTML entities.
// HTMLUnescape는 HTML 엔티티를 언이스케이프합니다.
//
// Example
// 예제:
//
//	HTMLUnescape("&lt;div&gt;hello&lt;/div&gt;")  // "<div>hello</div>"
//	HTMLUnescape("&#39;hello&#39; &amp; &#34;world&#34;")  // "'hello' & \"world\""
func HTMLUnescape(s string) string {
	return html.UnescapeString(s)
}
