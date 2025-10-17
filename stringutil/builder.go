package stringutil

import "strings"

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

// Truncate truncates the current value to the specified length and appends "...". / Truncate는 현재 값을 지정된 길이로 자르고 "..."를 추가합니다.
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
