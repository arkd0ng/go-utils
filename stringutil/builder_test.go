package stringutil

import "testing"

// Test basic chaining
// 기본 체이닝 테스트
func TestBuilderBasicChaining(t *testing.T) {
	result := NewBuilder().
		Append("hello").
		Append(" ").
		Append("world").
		Build()

	expected := "hello world"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// Test case conversion chaining
// 케이스 변환 체이닝 테스트
func TestBuilderCaseConversion(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		chain    func(*StringBuilder) *StringBuilder
		expected string
	}{
		{
			name:  "ToSnakeCase",
			input: "UserProfile",
			chain: func(sb *StringBuilder) *StringBuilder {
				return sb.ToSnakeCase()
			},
			expected: "user_profile",
		},
		{
			name:  "ToCamelCase",
			input: "user_profile",
			chain: func(sb *StringBuilder) *StringBuilder {
				return sb.ToCamelCase()
			},
			expected: "userProfile",
		},
		{
			name:  "ToKebabCase",
			input: "UserProfile",
			chain: func(sb *StringBuilder) *StringBuilder {
				return sb.ToKebabCase()
			},
			expected: "user-profile",
		},
		{
			name:  "ToPascalCase",
			input: "user_profile",
			chain: func(sb *StringBuilder) *StringBuilder {
				return sb.ToPascalCase()
			},
			expected: "UserProfile",
		},
		{
			name:  "ToTitle",
			input: "hello world",
			chain: func(sb *StringBuilder) *StringBuilder {
				return sb.ToTitle()
			},
			expected: "Hello World",
		},
		{
			name:  "ToUpper",
			input: "hello",
			chain: func(sb *StringBuilder) *StringBuilder {
				return sb.ToUpper()
			},
			expected: "HELLO",
		},
		{
			name:  "ToLower",
			input: "HELLO",
			chain: func(sb *StringBuilder) *StringBuilder {
				return sb.ToLower()
			},
			expected: "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.chain(NewBuilder().Append(tt.input)).Build()
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

// Test complex chaining
// 복잡한 체이닝 테스트
func TestBuilderComplexChaining(t *testing.T) {
	result := NewBuilder().
		Append("  user profile data  ").
		Clean().
		ToSnakeCase().
		Truncate(15).
		Build()

	expected := "user_profile_da..."
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// Test manipulation chaining
// 조작 체이닝 테스트
func TestBuilderManipulation(t *testing.T) {
	result := NewBuilder().
		Append("hello").
		Capitalize().
		Reverse().
		Build()

	expected := "olleH"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// Test truncation chaining
// 잘라내기 체이닝 테스트
func TestBuilderTruncation(t *testing.T) {
	result := NewBuilder().
		Append("Hello World").
		Truncate(8).
		Build()

	expected := "Hello Wo..."
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// Test custom suffix truncation
// 사용자 정의 suffix 잘라내기 테스트
func TestBuilderTruncationWithSuffix(t *testing.T) {
	result := NewBuilder().
		Append("Hello World").
		TruncateWithSuffix(8, "…").
		Build()

	expected := "Hello Wo…"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// Test cleaning chaining
// 정리 체이닝 테스트
func TestBuilderCleaning(t *testing.T) {
	result := NewBuilder().
		Append("  hello   world  ").
		Clean().
		Build()

	expected := "hello world"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// Test remove spaces chaining
// 공백 제거 체이닝 테스트
func TestBuilderRemoveSpaces(t *testing.T) {
	result := NewBuilder().
		Append("hello world").
		RemoveSpaces().
		Build()

	expected := "helloworld"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// Test remove special chars chaining
// 특수 문자 제거 체이닝 테스트
func TestBuilderRemoveSpecialChars(t *testing.T) {
	result := NewBuilder().
		Append("hello@world!").
		RemoveSpecialChars().
		Build()

	expected := "helloworld"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// Test repeat chaining
// 반복 체이닝 테스트
func TestBuilderRepeat(t *testing.T) {
	result := NewBuilder().
		Append("ab").
		Repeat(3).
		Build()

	expected := "ababab"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// Test slugify chaining
// Slugify 체이닝 테스트
func TestBuilderSlugify(t *testing.T) {
	result := NewBuilder().
		Append("Hello World!").
		Slugify().
		Build()

	expected := "hello-world"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// Test quote/unquote chaining
// 따옴표 체이닝 테스트
func TestBuilderQuote(t *testing.T) {
	result := NewBuilder().
		Append("hello").
		Quote().
		Build()

	expected := `"hello"`
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestBuilderUnquote(t *testing.T) {
	result := NewBuilder().
		Append(`"hello"`).
		Unquote().
		Build()

	expected := "hello"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// Test padding chaining
// 패딩 체이닝 테스트
func TestBuilderPadLeft(t *testing.T) {
	result := NewBuilder().
		Append("5").
		PadLeft(3, "0").
		Build()

	expected := "005"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestBuilderPadRight(t *testing.T) {
	result := NewBuilder().
		Append("5").
		PadRight(3, "0").
		Build()

	expected := "500"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// Test trim chaining
// 공백 제거 체이닝 테스트
func TestBuilderTrim(t *testing.T) {
	result := NewBuilder().
		Append("  hello  ").
		Trim().
		Build()

	expected := "hello"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// Test replace chaining
// 치환 체이닝 테스트
func TestBuilderReplace(t *testing.T) {
	result := NewBuilder().
		Append("hello world").
		Replace("world", "there").
		Build()

	expected := "hello there"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// Test NewBuilderWithString
// NewBuilderWithString 테스트
func TestNewBuilderWithString(t *testing.T) {
	result := NewBuilderWithString("hello").
		ToUpper().
		Build()

	expected := "HELLO"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// Test AppendLine
// AppendLine 테스트
func TestBuilderAppendLine(t *testing.T) {
	result := NewBuilder().
		AppendLine("line1").
		AppendLine("line2").
		Build()

	expected := "line1\nline2\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// Test Len
// Len 테스트
func TestBuilderLen(t *testing.T) {
	sb := NewBuilder().Append("hello")
	if sb.Len() != 5 {
		t.Errorf("Expected length 5, got %d", sb.Len())
	}

	// Unicode test
	// 유니코드 테스트
	sb = NewBuilder().Append("안녕하세요")
	if sb.Len() != 5 {
		t.Errorf("Expected length 5, got %d", sb.Len())
	}
}

// Test Reset
// Reset 테스트
func TestBuilderReset(t *testing.T) {
	sb := NewBuilder().Append("hello")
	if sb.Build() != "hello" {
		t.Error("Initial value should be 'hello'")
	}

	sb.Reset()
	if sb.Build() != "" {
		t.Error("After reset, value should be empty")
	}

	sb.Append("world")
	if sb.Build() != "world" {
		t.Error("After reset and append, value should be 'world'")
	}
}

// Test String method
// String 메서드 테스트
func TestBuilderString(t *testing.T) {
	sb := NewBuilder().Append("hello")
	if sb.String() != "hello" {
		t.Errorf("Expected 'hello', got %q", sb.String())
	}
}

// Test real-world scenario
// 실제 시나리오 테스트
func TestBuilderRealWorld(t *testing.T) {
	// Scenario: User inputs a messy string, we want to clean it up and create a slug
	// 시나리오: 사용자가 지저분한 문자열을 입력하면, 정리하고 slug를 생성합니다
	input := "  Hello World! This is a TEST  "
	result := NewBuilder().
		Append(input).
		Clean().
		ToLower().
		Replace(" ", "-").
		Build()

	expected := "hello-world!-this-is-a-test"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// Benchmark builder vs manual
// Builder와 수동 비교 벤치마크
func BenchmarkBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewBuilder().
			Append("user profile data").
			Clean().
			ToSnakeCase().
			Truncate(15).
			Build()
	}
}

func BenchmarkManual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := "user profile data"
		s = Clean(s)
		s = ToSnakeCase(s)
		s = Truncate(s, 15)
		_ = s
	}
}
