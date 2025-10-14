package stringutil

import "testing"

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"UserProfileData", "user_profile_data"},
		{"userProfileData", "user_profile_data"},
		{"user-profile-data", "user_profile_data"},
		{"USER_PROFILE_DATA", "user_profile_data"},
		{"", ""},
	}

	for _, tt := range tests {
		result := ToSnakeCase(tt.input)
		if result != tt.expected {
			t.Errorf("ToSnakeCase(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

func TestToCamelCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"user_profile_data", "userProfileData"},
		{"user-profile-data", "userProfileData"},
		{"UserProfileData", "userProfileData"},
		{"", ""},
	}

	for _, tt := range tests {
		result := ToCamelCase(tt.input)
		if result != tt.expected {
			t.Errorf("ToCamelCase(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

func TestToPascalCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"user_profile_data", "UserProfileData"},
		{"user-profile-data", "UserProfileData"},
		{"userProfileData", "UserProfileData"},
		{"", ""},
	}

	for _, tt := range tests {
		result := ToPascalCase(tt.input)
		if result != tt.expected {
			t.Errorf("ToPascalCase(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

func TestToTitle(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple lowercase", "hello world", "Hello World"},
		{"snake case", "user_profile_data", "User Profile Data"},
		{"kebab case", "hello-world", "Hello World"},
		{"camelCase", "helloWorld", "Hello World"},
		{"PascalCase", "HelloWorld", "Hello World"},
		{"mixed case", "hello_world-foo", "Hello World Foo"},
		{"empty string", "", ""},
		{"single word", "hello", "Hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToTitle(tt.input)
			if result != tt.expected {
				t.Errorf("ToTitle(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSlugify(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple text", "Hello World", "hello-world"},
		{"with punctuation", "Hello World!", "hello-world"},
		{"multiple spaces", "Hello   World", "hello-world"},
		{"special characters", "Go Utils -- Package", "go-utils-package"},
		{"mixed case", "User Profile Data", "user-profile-data"},
		{"numbers", "version 1.2.3", "version-1-2-3"},
		{"leading/trailing special", "!!hello world!!", "hello-world"},
		{"empty string", "", ""},
		{"only special chars", "!@#$%", ""},
		{"unicode", "hello 세계", "hello-세계"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Slugify(tt.input)
			if result != tt.expected {
				t.Errorf("Slugify(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestQuote(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple string", "hello", "\"hello\""},
		{"with quotes", "say \"hi\"", "\"say \\\"hi\\\"\""},
		{"with backslash", "path\\to\\file", "\"path\\\\to\\\\file\""},
		{"empty string", "", "\"\""},
		{"mixed", "he said \"hello\\world\"", "\"he said \\\"hello\\\\world\\\"\""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Quote(tt.input)
			if result != tt.expected {
				t.Errorf("Quote(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestUnquote(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"double quotes", "\"hello\"", "hello"},
		{"single quotes", "'world'", "world"},
		{"escaped quotes", "\"say \\\"hi\\\"\"", "say \"hi\""},
		{"escaped backslash", "\"path\\\\to\\\\file\"", "path\\to\\file"},
		{"no quotes", "hello", "hello"},
		{"empty string", "", ""},
		{"only quotes", "\"\"", ""},
		{"mixed escapes", "\"he said \\\"hello\\\\world\\\"\"", "he said \"hello\\world\""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Unquote(tt.input)
			if result != tt.expected {
				t.Errorf("Unquote(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestQuoteUnquoteRoundTrip(t *testing.T) {
	tests := []string{
		"hello",
		"say \"hi\"",
		"path\\to\\file",
		"mixed \"quotes\" and \\backslashes\\",
	}

	for _, original := range tests {
		quoted := Quote(original)
		unquoted := Unquote(quoted)
		if unquoted != original {
			t.Errorf("Quote/Unquote round trip failed for %q: got %q", original, unquoted)
		}
	}
}
