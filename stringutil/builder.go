package stringutil

import "strings"

// =============================================================================
// File: builder.go
// Purpose: Fluent String Builder with Method Chaining
// 파일: builder.go
// 목적: 메서드 체이닝을 사용한 유창한 문자열 빌더
// =============================================================================
//
// OVERVIEW
// 개요
// --------
// The builder.go file provides a fluent, chainable API for performing multiple
// string operations in sequence without intermediate variables. The StringBuilder
// type wraps a string value and exposes methods that transform the string and
// return the builder itself, enabling method chaining. This approach leads to
// more readable and maintainable code when performing multiple transformations
// on strings.
//
// builder.go 파일은 중간 변수 없이 여러 문자열 연산을 순차적으로 수행하기
// 위한 유창하고 체인 가능한 API를 제공합니다. StringBuilder 타입은 문자열
// 값을 래핑하고 문자열을 변환하여 빌더 자체를 반환하는 메서드를 노출하여
// 메서드 체이닝을 가능하게 합니다. 이 접근 방식은 문자열에 여러 변환을
// 수행할 때 더 읽기 쉽고 유지 관리가 쉬운 코드로 이어집니다.
//
// DESIGN PHILOSOPHY
// 설계 철학
// -----------------
// 1. **Fluent Interface**: All transformation methods return *StringBuilder
//    **유창한 인터페이스**: 모든 변환 메서드는 *StringBuilder 반환
//
// 2. **Immutable Operations**: Each transformation creates new state internally
//    **불변 연산**: 각 변환은 내부적으로 새로운 상태 생성
//
// 3. **Method Chaining**: Enable readable, declarative transformation pipelines
//    **메서드 체이닝**: 읽기 쉽고 선언적인 변환 파이프라인 가능
//
// 4. **Zero Boilerplate**: Eliminate intermediate variables for multi-step transformations
//    **보일러플레이트 제로**: 다단계 변환을 위한 중간 변수 제거
//
// 5. **Composability**: Combine operations in any order for custom workflows
//    **조합 가능성**: 사용자 정의 워크플로를 위해 모든 순서로 연산 결합
//
// CORE CONCEPT: METHOD CHAINING
// 핵심 개념: 메서드 체이닝
// -----------------------------
// Traditional approach (verbose):
// 전통적 접근 방식 (장황함):
//
//     s1 := "user profile data"
//     s2 := stringutil.Clean(s1)
//     s3 := stringutil.ToSnakeCase(s2)
//     s4 := stringutil.Truncate(s3, 20)
//     result := s4
//
// StringBuilder approach (fluent):
// StringBuilder 접근 방식 (유창함):
//
//     result := stringutil.NewBuilder().
//         Append("user profile data").
//         Clean().
//         ToSnakeCase().
//         Truncate(20).
//         Build()
//
// Benefits of method chaining:
// 메서드 체이닝의 이점:
// - No intermediate variables (s1, s2, s3)
//   중간 변수 없음 (s1, s2, s3)
// - Clear transformation pipeline
//   명확한 변환 파이프라인
// - Easy to reorder operations
//   연산 재정렬 용이
// - Self-documenting code flow
//   자체 문서화 코드 흐름
//
// FUNCTION CATEGORIES
// 함수 범주
// -------------------
//
// 1. CONSTRUCTION (생성)
//    - NewBuilder: Create empty builder
//      NewBuilder: 빈 빌더 생성
//    - NewBuilderWithString: Create builder with initial value
//      NewBuilderWithString: 초기 값으로 빌더 생성
//
// 2. APPENDING (추가)
//    - Append: Add string to current value
//      Append: 현재 값에 문자열 추가
//    - AppendLine: Add string with newline
//      AppendLine: 줄바꿈과 함께 문자열 추가
//
// 3. CASE CONVERSION (케이스 변환)
//    - ToSnakeCase, ToCamelCase, ToKebabCase, ToPascalCase
//    - ToTitle, ToUpper, ToLower
//
// 4. MANIPULATION (조작)
//    - Truncate, TruncateWithSuffix: Truncate strings
//      Truncate, TruncateWithSuffix: 문자열 자르기
//    - Reverse: Reverse string order
//      Reverse: 문자열 순서 뒤집기
//    - Capitalize, CapitalizeFirst: Capitalize letters
//      Capitalize, CapitalizeFirst: 문자 대문자화
//
// 5. CLEANING (정리)
//    - Clean: Normalize whitespace
//      Clean: 공백 정규화
//    - RemoveSpaces: Remove all spaces
//      RemoveSpaces: 모든 공백 제거
//    - RemoveSpecialChars: Remove non-alphanumeric
//      RemoveSpecialChars: 영숫자가 아닌 문자 제거
//    - Trim: Remove leading/trailing whitespace
//      Trim: 앞뒤 공백 제거
//
// 6. FORMATTING (포맷팅)
//    - PadLeft, PadRight: Add padding
//      PadLeft, PadRight: 패딩 추가
//    - Slugify: Convert to URL slug
//      Slugify: URL 슬러그로 변환
//    - Quote, Unquote: Handle quotes
//      Quote, Unquote: 따옴표 처리
//
// 7. TRANSFORMATION (변환)
//    - Repeat: Repeat string n times
//      Repeat: 문자열 n번 반복
//    - Replace: Replace substrings
//      Replace: 부분 문자열 치환
//
// 8. FINALIZATION (완료)
//    - Build: Get final string
//      Build: 최종 문자열 가져오기
//    - String: Get current value (fmt.Stringer)
//      String: 현재 값 가져오기 (fmt.Stringer)
//    - Len: Get length in runes
//      Len: rune 단위 길이 가져오기
//    - Reset: Clear to empty string
//      Reset: 빈 문자열로 초기화
//
// KEY OPERATIONS SUMMARY
// 주요 연산 요약
// ----------------------
//
// NewBuilder() *StringBuilder
// - Purpose: Create empty string builder
// - 목적: 빈 문자열 빌더 생성
// - Time Complexity: O(1)
// - 시간 복잡도: O(1)
// - Space Complexity: O(1)
// - 공간 복잡도: O(1)
// - Use Cases: Start of transformation pipeline
// - 사용 사례: 변환 파이프라인 시작
//
// NewBuilderWithString(s string) *StringBuilder
// - Purpose: Create builder with initial string
// - 목적: 초기 문자열로 빌더 생성
// - Time Complexity: O(1)
// - 시간 복잡도: O(1)
// - Space Complexity: O(n) for string copy
// - 공간 복잡도: O(n), 문자열 복사용
// - Use Cases: Transform existing string
// - 사용 사례: 기존 문자열 변환
//
// Append(s string) *StringBuilder
// - Purpose: Concatenate string to current value
// - 목적: 현재 값에 문자열 연결
// - Time Complexity: O(n + m) where n is current, m is appended
// - 시간 복잡도: O(n + m), n은 현재, m은 추가
// - Space Complexity: O(n + m)
// - 공간 복잡도: O(n + m)
// - Chainable: Yes
// - 체인 가능: 예
// - Use Cases: Building strings piece by piece
// - 사용 사례: 문자열을 조각별로 구성
//
// ToSnakeCase() *StringBuilder
// - Purpose: Convert to snake_case
// - 목적: snake_case로 변환
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Delegates to: case.go ToSnakeCase
// - 위임: case.go ToSnakeCase
// - Use Cases: API parameter formatting, database columns
// - 사용 사례: API 매개변수 포맷팅, 데이터베이스 컬럼
//
// Clean() *StringBuilder
// - Purpose: Normalize whitespace (trim + deduplicate)
// - 목적: 공백 정규화 (제거 + 중복 제거)
// - Time Complexity: O(n)
// - 시간 복잡도: O(n)
// - Space Complexity: O(n)
// - 공간 복잡도: O(n)
// - Delegates to: manipulation.go Clean
// - 위임: manipulation.go Clean
// - Use Cases: User input sanitization
// - 사용 사례: 사용자 입력 정제
//
// Build() string
// - Purpose: Extract final string value
// - 목적: 최종 문자열 값 추출
// - Time Complexity: O(1)
// - 시간 복잡도: O(1)
// - Space Complexity: O(1) (returns reference)
// - 공간 복잡도: O(1) (참조 반환)
// - Terminal Operation: Ends the chain
// - 종료 연산: 체인 종료
// - Use Cases: Get result after transformations
// - 사용 사례: 변환 후 결과 가져오기
//
// METHOD CHAINING PATTERNS
// 메서드 체이닝 패턴
// ------------------------
//
// All transformation methods follow this pattern:
// 모든 변환 메서드는 이 패턴을 따릅니다:
//
//     func (sb *StringBuilder) MethodName(...) *StringBuilder {
//         sb.value = transform(sb.value)
//         return sb  // Enable chaining
//     }
//
// This design allows:
// 이 설계는 다음을 허용합니다:
// 1. Chaining multiple operations: obj.A().B().C()
//    여러 연산 체이닝: obj.A().B().C()
// 2. Readable left-to-right flow
//    읽기 쉬운 왼쪽에서 오른쪽 흐름
// 3. Easy insertion/removal of steps
//    단계의 쉬운 삽입/제거
// 4. Self-documenting transformation pipeline
//    자체 문서화 변환 파이프라인
//
// PERFORMANCE CHARACTERISTICS
// 성능 특성
// ---------------------------
//
// Time Complexities:
// 시간 복잡도:
// - Construction: O(1) for NewBuilder, O(n) for NewBuilderWithString
//   생성: NewBuilder는 O(1), NewBuilderWithString은 O(n)
// - Append: O(n + m) due to string concatenation
//   Append: 문자열 연결로 인해 O(n + m)
// - Transformations: O(n) for most operations (delegates to underlying functions)
//   변환: 대부분의 연산은 O(n) (기본 함수에 위임)
// - Build: O(1) - returns current value
//   Build: O(1) - 현재 값 반환
//
// Space Complexities:
// 공간 복잡도:
// - StringBuilder itself: O(1) - just a string reference
//   StringBuilder 자체: O(1) - 문자열 참조만
// - Intermediate strings: O(n) for each transformation
//   중간 문자열: 각 변환마다 O(n)
// - Total: O(n * t) where t is number of transformations
//   총: O(n * t), t는 변환 수
//
// Optimization Tips:
// 최적화 팁:
// 1. Avoid excessive Append calls - use AppendLine or single Append
//    과도한 Append 호출 피하기 - AppendLine 또는 단일 Append 사용
// 2. Order operations efficiently (e.g., Clean before ToSnakeCase)
//    연산 순서를 효율적으로 (예: ToSnakeCase 전에 Clean)
// 3. Use Reset() to reuse builder for multiple strings
//    여러 문자열에 빌더를 재사용하려면 Reset() 사용
// 4. For large strings, consider breaking chain into stages
//    대용량 문자열의 경우 체인을 단계로 나누는 것 고려
// 5. Avoid deep chains for performance-critical code
//    성능이 중요한 코드에서는 깊은 체인 피하기
//
// Performance Consideration: String Immutability
// 성능 고려사항: 문자열 불변성
// -----------------------------------------------
// Go strings are immutable. Each transformation creates a new string.
// StringBuilder does NOT use strings.Builder internally - it's a semantic
// wrapper for chaining operations, not a performance optimization for
// concatenation.
//
// Go 문자열은 불변입니다. 각 변환은 새로운 문자열을 생성합니다.
// StringBuilder는 내부적으로 strings.Builder를 사용하지 않습니다 -
// 연결에 대한 성능 최적화가 아니라 연산 체이닝을 위한 의미론적 래퍼입니다.
//
// For heavy string concatenation (loops), use strings.Builder:
// 무거운 문자열 연결(루프)의 경우 strings.Builder 사용:
//
//     var builder strings.Builder
//     for _, s := range manyStrings {
//         builder.WriteString(s)
//     }
//     result := builder.String()
//
// For transformation pipelines, use StringBuilder:
// 변환 파이프라인의 경우 StringBuilder 사용:
//
//     result := stringutil.NewBuilder().
//         Append(s).
//         Clean().
//         ToSnakeCase().
//         Build()
//
// EDGE CASES AND SPECIAL BEHAVIORS
// 엣지 케이스 및 특수 동작
// ---------------------------------
//
// Empty Builder:
// 빈 빌더:
// - NewBuilder() creates builder with empty string
//   NewBuilder()는 빈 문자열로 빌더 생성
// - Build() on empty builder returns ""
//   빈 빌더에서 Build()는 "" 반환
// - Len() on empty builder returns 0
//   빈 빌더에서 Len()은 0 반환
//
// Nil Safety:
// nil 안전성:
// - All methods assume non-nil StringBuilder
//   모든 메서드는 nil이 아닌 StringBuilder 가정
// - Calling methods on nil will panic (Go default)
//   nil에서 메서드 호출은 패닉 발생 (Go 기본)
//
// Reset:
// 초기화:
// - Reset() clears value to empty string
//   Reset()은 값을 빈 문자열로 지움
// - Allows builder reuse without reallocation
//   재할당 없이 빌더 재사용 허용
//
// String() vs Build():
// String() 대 Build():
// - Both return current string value
//   둘 다 현재 문자열 값 반환
// - String() implements fmt.Stringer
//   String()은 fmt.Stringer 구현
// - Build() is semantic terminal operation
//   Build()는 의미론적 종료 연산
//
// Operation Order:
// 연산 순서:
// - Order matters! Clean().ToSnakeCase() != ToSnakeCase().Clean()
//   순서가 중요! Clean().ToSnakeCase() != ToSnakeCase().Clean()
// - Plan transformation pipeline carefully
//   변환 파이프라인을 신중하게 계획
//
// COMMON USAGE PATTERNS
// 일반 사용 패턴
// ---------------------
//
// 1. Converting User Input to Database Field Name
//    사용자 입력을 데이터베이스 필드명으로 변환:
//
//    userInput := "  User Profile Data  "
//    fieldName := stringutil.NewBuilder().
//        Append(userInput).
//        Clean().
//        ToSnakeCase().
//        Build()
//    // "user_profile_data"
//    // Clean input, convert to database-friendly format
//    // 입력 정리, 데이터베이스 친화적 형식으로 변환
//
// 2. Building URL Slugs
//    URL 슬러그 구성:
//
//    title := "Hello, World! - A Guide"
//    slug := stringutil.NewBuilder().
//        Append(title).
//        Slugify().
//        Truncate(30).
//        Build()
//    // "hello-world-a-guide"
//    // Create SEO-friendly URLs
//    // SEO 친화적 URL 생성
//
// 3. Formatting API Parameter Names
//    API 매개변수명 포맷팅:
//
//    goField := "UserProfileID"
//    apiParam := stringutil.NewBuilder().
//        Append(goField).
//        ToCamelCase().
//        Build()
//    // "userProfileId"
//    // Convert Go style to JavaScript style
//    // Go 스타일을 JavaScript 스타일로 변환
//
// 4. Cleaning and Formatting Text
//    텍스트 정리 및 포맷팅:
//
//    messyText := "  HELLO   WORLD  "
//    clean := stringutil.NewBuilder().
//        Append(messyText).
//        Clean().
//        ToTitle().
//        Build()
//    // "Hello World"
//    // Normalize and capitalize
//    // 정규화 및 대문자화
//
// 5. Building Multi-Line Output
//    여러 줄 출력 구성:
//
//    output := stringutil.NewBuilder().
//        AppendLine("Header").
//        AppendLine("=======").
//        Append("Content here").
//        Build()
//    // "Header\n=======\nContent here"
//    // Construct formatted output
//    // 포맷된 출력 구성
//
// 6. Conditional Transformations
//    조건부 변환:
//
//    sb := stringutil.NewBuilder().Append(input)
//    if needsCleaning {
//        sb.Clean()
//    }
//    if toSnake {
//        sb.ToSnakeCase()
//    }
//    result := sb.Build()
//    // Apply transformations conditionally
//    // 조건부로 변환 적용
//
// 7. Reusable Builder
//    재사용 가능한 빌더:
//
//    sb := stringutil.NewBuilder()
//    for _, input := range inputs {
//        result := sb.
//            Append(input).
//            Clean().
//            ToSnakeCase().
//            Build()
//        process(result)
//        sb.Reset()
//    }
//    // Reuse builder for efficiency
//    // 효율성을 위해 빌더 재사용
//
// 8. Padding and Formatting Numbers
//    숫자 패딩 및 포맷팅:
//
//    num := 42
//    formatted := stringutil.NewBuilder().
//        Append(fmt.Sprintf("%d", num)).
//        PadLeft(5, "0").
//        Build()
//    // "00042"
//    // Zero-pad numbers
//    // 숫자 제로 패딩
//
// 9. Quoting for Shell Commands
//    셸 명령을 위한 따옴표:
//
//    filename := "my file.txt"
//    quoted := stringutil.NewBuilder().
//        Append(filename).
//        Quote().
//        Build()
//    // "\"my file.txt\""
//    // Safe for shell execution
//    // 셸 실행에 안전
//
// 10. Complex Transformation Pipeline
//     복잡한 변환 파이프라인:
//
//     result := stringutil.NewBuilder().
//         Append("  User@Profile#Data  ").
//         Clean().
//         RemoveSpecialChars().
//         ToSnakeCase().
//         Truncate(20).
//         Build()
//     // "user_profile_data"
//     // Multi-step cleaning and formatting
//     // 다단계 정리 및 포맷팅
//
// COMPARISON WITH ALTERNATIVES
// 대안과의 비교
// ----------------------------
//
// StringBuilder vs Direct Function Calls
// - StringBuilder: Fluent, chainable, self-documenting
//   StringBuilder: 유창함, 체인 가능, 자체 문서화
// - Functions: More verbose, requires intermediate variables
//   함수: 더 장황함, 중간 변수 필요
// - Use StringBuilder for: Multi-step transformations
//   StringBuilder 사용: 다단계 변환
// - Use Functions for: Single operations
//   함수 사용: 단일 연산
//
// StringBuilder vs strings.Builder
// - stringutil.StringBuilder: Transformation pipeline wrapper
//   stringutil.StringBuilder: 변환 파이프라인 래퍼
// - strings.Builder: Efficient concatenation for loops
//   strings.Builder: 루프를 위한 효율적 연결
// - Different purposes, not interchangeable
//   다른 목적, 교환 불가능
//
// StringBuilder vs Method Chaining Libraries (e.g., lodash)
// - Similar philosophy: fluent interfaces
//   유사한 철학: 유창한 인터페이스
// - StringBuilder: String-specific operations
//   StringBuilder: 문자열 전용 연산
// - Use for: Readable transformation pipelines
//   사용: 읽기 쉬운 변환 파이프라인
//
// THREAD SAFETY
// 스레드 안전성
// -------------
// StringBuilder is NOT thread-safe. Do not share a single StringBuilder instance
// across multiple goroutines without synchronization.
//
// StringBuilder는 스레드 안전하지 않습니다. 동기화 없이 여러 고루틴에서
// 단일 StringBuilder 인스턴스를 공유하지 마세요.
//
// Unsafe:
// 안전하지 않음:
//
//     sb := stringutil.NewBuilder()
//     go func() { sb.Append("A") }()
//     go func() { sb.Append("B") }()  // Race condition
//
// Safe:
// 안전:
//
//     // Each goroutine gets its own builder
//     // 각 고루틴이 자체 빌더 가져오기
//     go func() {
//         sb := stringutil.NewBuilder().Append("A")
//     }()
//
// DESIGN RATIONALE
// 설계 근거
// ----------------
// Why not use strings.Builder directly?
// strings.Builder를 직접 사용하지 않는 이유는?
//
// 1. **Different Purpose**: strings.Builder optimizes concatenation,
//    StringBuilder provides transformation pipeline
//    **다른 목적**: strings.Builder는 연결을 최적화,
//    StringBuilder는 변환 파이프라인 제공
//
// 2. **Semantic Clarity**: NewBuilder().Append(s).ToSnakeCase() is clearer
//    than manually calling functions
//    **의미론적 명확성**: NewBuilder().Append(s).ToSnakeCase()가
//    수동으로 함수 호출하는 것보다 명확
//
// 3. **Composability**: Easy to add/remove transformation steps
//    **조합 가능성**: 변환 단계 추가/제거 용이
//
// 4. **Discoverability**: IDE autocomplete shows available operations
//    **발견 가능성**: IDE 자동 완성이 사용 가능한 연산 표시
//
// RELATED FILES
// 관련 파일
// -------------
// - manipulation.go: Provides Truncate, Reverse, Clean, etc.
//   manipulation.go: Truncate, Reverse, Clean 등 제공
// - case.go: Provides case conversion functions
//   case.go: 케이스 변환 함수 제공
// - formatting.go: Provides formatting utilities
//   formatting.go: 포맷팅 유틸리티 제공
// - utils.go: Provides PadLeft, PadRight, etc.
//   utils.go: PadLeft, PadRight 등 제공
//
// =============================================================================

// StringBuilder provides a fluent API for chaining string operations.
// StringBuilder는 문자열 작업을 체이닝하기 위한 fluent API를 제공합니다.
//
// Example
// 예제:
//
//	result := stringutil.NewBuilder().
//		Append("user profile data").
//		Clean().
//		ToSnakeCase().
//		Truncate(20).
//		Build()
type StringBuilder struct {
	value string
}

// NewBuilder creates a new StringBuilder.
// NewBuilder는 새 StringBuilder를 생성합니다.
//
// Example
// 예제:
//
//	sb := stringutil.NewBuilder()
func NewBuilder() *StringBuilder {
	return &StringBuilder{value: ""}
}

// NewBuilderWithString creates a new StringBuilder with an initial string.
// NewBuilderWithString은 초기 문자열로 새 StringBuilder를 생성합니다.
//
// Example
// 예제:
//
//	sb := stringutil.NewBuilderWithString("hello")
func NewBuilderWithString(s string) *StringBuilder {
	return &StringBuilder{value: s}
}

// Append appends a string to the builder.
// Append는 빌더에 문자열을 추가합니다.
//
// Example
// 예제:
//
//	sb.Append("hello").Append(" world")
func (sb *StringBuilder) Append(s string) *StringBuilder {
	sb.value += s
	return sb
}

// AppendLine appends a string followed by a newline.
// AppendLine은 문자열 뒤에 줄바꿈을 추가합니다.
//
// Example
// 예제:
//
//	sb.AppendLine("line1").AppendLine("line2")
func (sb *StringBuilder) AppendLine(s string) *StringBuilder {
	sb.value += s + "\n"
	return sb
}

// ToSnakeCase converts the current value to snake_case.
// ToSnakeCase는 현재 값을 snake_case로 변환합니다.
//
// Example
// 예제:
//
//	sb.Append("UserProfile").ToSnakeCase().Build()  // "user_profile"
func (sb *StringBuilder) ToSnakeCase() *StringBuilder {
	sb.value = ToSnakeCase(sb.value)
	return sb
}

// ToCamelCase converts the current value to camelCase.
// ToCamelCase는 현재 값을 camelCase로 변환합니다.
//
// Example
// 예제:
//
//	sb.Append("user_profile").ToCamelCase().Build()  // "userProfile"
func (sb *StringBuilder) ToCamelCase() *StringBuilder {
	sb.value = ToCamelCase(sb.value)
	return sb
}

// ToKebabCase converts the current value to kebab-case.
// ToKebabCase는 현재 값을 kebab-case로 변환합니다.
//
// Example
// 예제:
//
//	sb.Append("UserProfile").ToKebabCase().Build()  // "user-profile"
func (sb *StringBuilder) ToKebabCase() *StringBuilder {
	sb.value = ToKebabCase(sb.value)
	return sb
}

// ToPascalCase converts the current value to PascalCase.
// ToPascalCase는 현재 값을 PascalCase로 변환합니다.
//
// Example
// 예제:
//
//	sb.Append("user_profile").ToPascalCase().Build()  // "UserProfile"
func (sb *StringBuilder) ToPascalCase() *StringBuilder {
	sb.value = ToPascalCase(sb.value)
	return sb
}

// ToTitle converts the current value to Title Case.
// ToTitle은 현재 값을 Title Case로 변환합니다.
//
// Example
// 예제:
//
//	sb.Append("hello world").ToTitle().Build()  // "Hello World"
func (sb *StringBuilder) ToTitle() *StringBuilder {
	sb.value = ToTitle(sb.value)
	return sb
}

// ToUpper converts the current value to uppercase.
// ToUpper는 현재 값을 대문자로 변환합니다.
//
// Example
// 예제:
//
//	sb.Append("hello").ToUpper().Build()  // "HELLO"
func (sb *StringBuilder) ToUpper() *StringBuilder {
	sb.value = strings.ToUpper(sb.value)
	return sb
}

// ToLower converts the current value to lowercase.
// ToLower는 현재 값을 소문자로 변환합니다.
//
// Example
// 예제:
//
//	sb.Append("HELLO").ToLower().Build()  // "hello"
func (sb *StringBuilder) ToLower() *StringBuilder {
	sb.value = strings.ToLower(sb.value)
	return sb
}

// Truncate truncates the current value to the specified length and appends "...".
// Truncate는 현재 값을 지정된 길이로 자르고 "..."를 추가합니다.
//
// Example
// 예제:
//
//	sb.Append("Hello World").Truncate(8).Build()  // "Hello..."
func (sb *StringBuilder) Truncate(length int) *StringBuilder {
	sb.value = Truncate(sb.value, length)
	return sb
}

// TruncateWithSuffix truncates the current value with a custom suffix.
// TruncateWithSuffix는 현재 값을 사용자 정의 suffix로 자릅니다.
//
// Example
// 예제:
//
//	sb.Append("Hello World").TruncateWithSuffix(8, "…").Build()  // "Hello Wo…"
func (sb *StringBuilder) TruncateWithSuffix(length int, suffix string) *StringBuilder {
	sb.value = TruncateWithSuffix(sb.value, length, suffix)
	return sb
}

// Reverse reverses the current value (Unicode-safe).
// Reverse는 현재 값을 뒤집습니다 (유니코드 안전).
//
// Example
// 예제:
//
//	sb.Append("hello").Reverse().Build()  // "olleh"
func (sb *StringBuilder) Reverse() *StringBuilder {
	sb.value = Reverse(sb.value)
	return sb
}

// Capitalize capitalizes each word in the current value.
// Capitalize는 현재 값의 각 단어를 대문자로 시작하게 합니다.
//
// Example
// 예제:
//
//	sb.Append("hello world").Capitalize().Build()  // "Hello World"
func (sb *StringBuilder) Capitalize() *StringBuilder {
	sb.value = Capitalize(sb.value)
	return sb
}

// CapitalizeFirst capitalizes only the first letter.
// CapitalizeFirst는 첫 글자만 대문자로 변환합니다.
//
// Example
// 예제:
//
//	sb.Append("hello world").CapitalizeFirst().Build()  // "Hello world"
func (sb *StringBuilder) CapitalizeFirst() *StringBuilder {
	sb.value = CapitalizeFirst(sb.value)
	return sb
}

// Clean trims and deduplicates spaces.
// Clean은 공백을 제거하고 중복 공백을 정리합니다.
//
// Example
// 예제:
//
//	sb.Append("  hello   world  ").Clean().Build()  // "hello world"
func (sb *StringBuilder) Clean() *StringBuilder {
	sb.value = Clean(sb.value)
	return sb
}

// RemoveSpaces removes all whitespace.
// RemoveSpaces는 모든 공백을 제거합니다.
//
// Example
// 예제:
//
//	sb.Append("hello world").RemoveSpaces().Build()  // "helloworld"
func (sb *StringBuilder) RemoveSpaces() *StringBuilder {
	sb.value = RemoveSpaces(sb.value)
	return sb
}

// RemoveSpecialChars removes all special characters (keeps only alphanumeric).
// RemoveSpecialChars는 모든 특수 문자를 제거합니다 (영숫자만 유지).
//
// Example
// 예제:
//
//	sb.Append("hello@world!").RemoveSpecialChars().Build()  // "helloworld"
func (sb *StringBuilder) RemoveSpecialChars() *StringBuilder {
	sb.value = RemoveSpecialChars(sb.value)
	return sb
}

// Repeat repeats the current value count times.
// Repeat는 현재 값을 count번 반복합니다.
//
// Example
// 예제:
//
//	sb.Append("ab").Repeat(3).Build()  // "ababab"
func (sb *StringBuilder) Repeat(count int) *StringBuilder {
	sb.value = Repeat(sb.value, count)
	return sb
}

// Slugify converts the current value to a URL-friendly slug.
// Slugify는 현재 값을 URL 친화적인 slug로 변환합니다.
//
// Example
// 예제:
//
//	sb.Append("Hello World!").Slugify().Build()  // "hello-world"
func (sb *StringBuilder) Slugify() *StringBuilder {
	sb.value = Slugify(sb.value)
	return sb
}

// Quote wraps the current value in double quotes.
// Quote는 현재 값을 큰따옴표로 감쌉니다.
//
// Example
// 예제:
//
//	sb.Append("hello").Quote().Build()  // "\"hello\""
func (sb *StringBuilder) Quote() *StringBuilder {
	sb.value = Quote(sb.value)
	return sb
}

// Unquote removes surrounding quotes.
// Unquote는 주변 따옴표를 제거합니다.
//
// Example
// 예제:
//
//	sb.Append("\"hello\"").Unquote().Build()  // "hello"
func (sb *StringBuilder) Unquote() *StringBuilder {
	sb.value = Unquote(sb.value)
	return sb
}

// PadLeft pads the current value on the left side to reach the specified length.
// PadLeft는 지정된 길이에 도달하도록 현재 값의 왼쪽에 패딩을 추가합니다.
//
// Example
// 예제:
//
//	sb.Append("5").PadLeft(3, "0").Build()  // "005"
func (sb *StringBuilder) PadLeft(length int, pad string) *StringBuilder {
	sb.value = PadLeft(sb.value, length, pad)
	return sb
}

// PadRight pads the current value on the right side to reach the specified length.
// PadRight는 지정된 길이에 도달하도록 현재 값의 오른쪽에 패딩을 추가합니다.
//
// Example
// 예제:
//
//	sb.Append("5").PadRight(3, "0").Build()  // "500"
func (sb *StringBuilder) PadRight(length int, pad string) *StringBuilder {
	sb.value = PadRight(sb.value, length, pad)
	return sb
}

// Trim removes leading and trailing whitespace.
// Trim은 앞뒤 공백을 제거합니다.
//
// Example
// 예제:
//
//	sb.Append("  hello  ").Trim().Build()  // "hello"
func (sb *StringBuilder) Trim() *StringBuilder {
	sb.value = strings.TrimSpace(sb.value)
	return sb
}

// Replace replaces all occurrences of old with new.
// Replace는 old를 모두 new로 치환합니다.
//
// Example
// 예제:
//
//	sb.Append("hello world").Replace("world", "there").Build()  // "hello there"
func (sb *StringBuilder) Replace(old, new string) *StringBuilder {
	sb.value = strings.ReplaceAll(sb.value, old, new)
	return sb
}

// Build returns the final string value.
// Build는 최종 문자열 값을 반환합니다.
//
// Example
// 예제:
//
//	result := sb.Append("hello").ToUpper().Build()  // "HELLO"
func (sb *StringBuilder) Build() string {
	return sb.value
}

// String returns the current string value (implements fmt.Stringer).
// String은 현재 문자열 값을 반환합니다 (fmt.Stringer 인터페이스 구현).
func (sb *StringBuilder) String() string {
	return sb.value
}

// Len returns the length of the current value in runes (Unicode-safe).
// Len은 현재 값의 길이를 rune 단위로 반환합니다 (유니코드 안전).
func (sb *StringBuilder) Len() int {
	return len([]rune(sb.value))
}

// Reset resets the builder to an empty string.
// Reset은 빌더를 빈 문자열로 초기화합니다.
func (sb *StringBuilder) Reset() *StringBuilder {
	sb.value = ""
	return sb
}
