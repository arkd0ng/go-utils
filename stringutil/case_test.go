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
