package validation

import (
	"testing"
)

// TestWithCustomMessage tests the WithCustomMessage method
// TestWithCustomMessage는 WithCustomMessage 메서드를 테스트합니다
func TestWithCustomMessage(t *testing.T) {
	t.Run("Single custom message", func(t *testing.T) {
		v := New("", "email")
		v.WithCustomMessage("required", "Please enter your email address")
		v.Required()

		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Fatalf("Expected 1 error, got %d", len(errors))
		}

		if errors[0].Message != "Please enter your email address" {
			t.Errorf("Expected custom message 'Please enter your email address', got '%s'", errors[0].Message)
		}
	})

	t.Run("Multiple rules with custom messages", func(t *testing.T) {
		v := New("a", "password")
		v.WithCustomMessage("minlength", "Password must be at least 8 characters")
		v.WithCustomMessage("maxlength", "Password cannot exceed 20 characters")
		v.MinLength(8).MaxLength(20)

		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Fatalf("Expected 1 error, got %d", len(errors))
		}

		if errors[0].Message != "Password must be at least 8 characters" {
			t.Errorf("Expected custom minlength message, got '%s'", errors[0].Message)
		}
	})

	t.Run("Custom message not applied to different rule", func(t *testing.T) {
		v := New("test", "username")
		v.WithCustomMessage("required", "Username is required")
		v.MinLength(5) // This should use default message, not custom

		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Fatalf("Expected 1 error, got %d", len(errors))
		}

		// Should NOT have the custom "required" message
		if errors[0].Message == "Username is required" {
			t.Errorf("Custom message should not apply to different rule")
		}
	})
}

// TestWithCustomMessages tests the WithCustomMessages method
// TestWithCustomMessages는 WithCustomMessages 메서드를 테스트합니다
func TestWithCustomMessages(t *testing.T) {
	t.Run("Multiple custom messages at once", func(t *testing.T) {
		v := New("", "password")
		v.WithCustomMessages(map[string]string{
			"required":  "비밀번호를 입력해주세요",
			"minlength": "비밀번호는 8자 이상이어야 합니다",
			"maxlength": "비밀번호는 20자 이하여야 합니다",
		})
		v.StopOnError().Required().MinLength(8).MaxLength(20)

		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Fatalf("Expected 1 error (required), got %d", len(errors))
		}

		if errors[0].Message != "비밀번호를 입력해주세요" {
			t.Errorf("Expected Korean custom message, got '%s'", errors[0].Message)
		}
	})

	t.Run("Chaining with validation rules", func(t *testing.T) {
		v := New("abc", "username")
		v.WithCustomMessages(map[string]string{
			"minlength": "사용자명은 최소 5자 이상이어야 합니다",
			"maxlength": "사용자명은 최대 20자까지 가능합니다",
		}).MinLength(5).MaxLength(20)

		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Fatalf("Expected 1 error, got %d", len(errors))
		}

		if errors[0].Message != "사용자명은 최소 5자 이상이어야 합니다" {
			t.Errorf("Expected custom minlength message, got '%s'", errors[0].Message)
		}
	})

	t.Run("Overwriting custom messages", func(t *testing.T) {
		v := New("", "field")
		v.WithCustomMessage("required", "First message")
		v.WithCustomMessage("required", "Second message") // Should overwrite
		v.Required()

		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Fatalf("Expected 1 error, got %d", len(errors))
		}

		if errors[0].Message != "Second message" {
			t.Errorf("Expected overwritten message 'Second message', got '%s'", errors[0].Message)
		}
	})
}

// TestCustomMessageWithStopOnError tests custom messages with StopOnError
// TestCustomMessageWithStopOnError는 StopOnError와 함께 커스텀 메시지를 테스트합니다
func TestCustomMessageWithStopOnError(t *testing.T) {
	t.Run("StopOnError with custom message", func(t *testing.T) {
		v := New("invalid", "email")
		v.WithCustomMessages(map[string]string{
			"required": "이메일 주소를 입력해주세요",
			"email":    "올바른 이메일 형식이 아닙니다",
		})
		v.StopOnError().Email()

		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Fatalf("Expected 1 error (stopped after first), got %d", len(errors))
		}

		if errors[0].Message != "올바른 이메일 형식이 아닙니다" {
			t.Errorf("Expected custom email message, got '%s'", errors[0].Message)
		}
	})
}

// TestCustomMessageWithMultiValidator tests custom messages with MultiValidator
// TestCustomMessageWithMultiValidator는 MultiValidator와 함께 커스텀 메시지를 테스트합니다
func TestCustomMessageWithMultiValidator(t *testing.T) {
	t.Run("MultiValidator with custom messages", func(t *testing.T) {
		mv := NewValidator()

		mv.Field("", "email").WithCustomMessages(map[string]string{
			"required": "Email is required",
			"email":    "Invalid email format",
		}).Required()

		mv.Field("", "password").WithCustomMessages(map[string]string{
			"required":  "Password is required",
			"minlength": "Password must be at least 8 characters",
		}).Required()

		err := mv.Validate()
		if err == nil {
			t.Fatal("Expected validation errors")
		}

		errors := mv.GetErrors()
		if len(errors) != 2 {
			t.Fatalf("Expected 2 errors (email and password required), got %d", len(errors))
		}

		// Check email error message
		if errors[0].Message != "Email is required" {
			t.Errorf("Expected custom email message, got '%s'", errors[0].Message)
		}

		// Check password error message
		if errors[1].Message != "Password is required" {
			t.Errorf("Expected custom password message, got '%s'", errors[1].Message)
		}
	})
}

// TestCustomMessagePreservation tests that custom messages are preserved across validation chain
// TestCustomMessagePreservation는 검증 체인에서 커스텀 메시지가 보존되는지 테스트합니다
func TestCustomMessagePreservation(t *testing.T) {
	t.Run("Custom messages preserved in long chain", func(t *testing.T) {
		v := New("ab", "username")
		v.WithCustomMessages(map[string]string{
			"required":  "Username is required",
			"minlength": "Username too short",
			"maxlength": "Username too long",
			"alpha":     "Username must be alphabetic",
		})
		v.Required().MinLength(3).MaxLength(20).Alpha()

		errors := v.GetErrors()
		if len(errors) != 1 {
			t.Fatalf("Expected 1 error (minlength), got %d", len(errors))
		}

		if errors[0].Message != "Username too short" {
			t.Errorf("Expected 'Username too short', got '%s'", errors[0].Message)
		}
	})
}
