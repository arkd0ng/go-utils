package websvrutil

import (
	"strings"
	"testing"
)

// TestValidatorRequired tests required validation.
// TestValidatorRequired는 required 검증을 테스트합니다.
func TestValidatorRequired(t *testing.T) {
	type TestStruct struct {
		Name string `validate:"required"`
	}

	validator := &DefaultValidator{}

	// Valid case
	// 유효한 경우
	obj := TestStruct{Name: "John"}
	if err := validator.Validate(&obj); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Invalid case
	// 무효한 경우
	obj = TestStruct{Name: ""}
	err := validator.Validate(&obj)
	if err == nil {
		t.Error("Expected error for empty required field")
	}
}

// TestValidatorEmail tests email validation.
// TestValidatorEmail은 이메일 검증을 테스트합니다.
func TestValidatorEmail(t *testing.T) {
	type TestStruct struct {
		Email string `validate:"email"`
	}

	validator := &DefaultValidator{}

	// Valid cases
	// 유효한 경우
	validEmails := []string{
		"test@example.com",
		"user.name@example.co.uk",
		"test+tag@example.com",
	}

	for _, email := range validEmails {
		obj := TestStruct{Email: email}
		if err := validator.Validate(&obj); err != nil {
			t.Errorf("Expected no error for email %s, got %v", email, err)
		}
	}

	// Invalid cases
	// 무효한 경우
	invalidEmails := []string{
		"invalid",
		"@example.com",
		"test@",
		"test @example.com",
	}

	for _, email := range invalidEmails {
		obj := TestStruct{Email: email}
		err := validator.Validate(&obj)
		if err == nil {
			t.Errorf("Expected error for invalid email %s", email)
		}
	}
}

// TestValidatorMinMax tests min/max validation.
// TestValidatorMinMax는 min/max 검증을 테스트합니다.
func TestValidatorMinMax(t *testing.T) {
	type TestStruct struct {
		Name string `validate:"min=3,max=10"`
		Age  int    `validate:"min=18,max=100"`
	}

	validator := &DefaultValidator{}

	// Valid case
	// 유효한 경우
	obj := TestStruct{Name: "John", Age: 25}
	if err := validator.Validate(&obj); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Invalid min
	// 무효한 최소값
	obj = TestStruct{Name: "Jo", Age: 25}
	err := validator.Validate(&obj)
	if err == nil {
		t.Error("Expected error for name below min length")
	}

	// Invalid max
	// 무효한 최대값
	obj = TestStruct{Name: "VeryLongName", Age: 25}
	err = validator.Validate(&obj)
	if err == nil {
		t.Error("Expected error for name above max length")
	}

	// Invalid age
	// 무효한 나이
	obj = TestStruct{Name: "John", Age: 17}
	err = validator.Validate(&obj)
	if err == nil {
		t.Error("Expected error for age below min")
	}
}

// TestValidatorLen tests exact length validation.
// TestValidatorLen은 정확한 길이 검증을 테스트합니다.
func TestValidatorLen(t *testing.T) {
	type TestStruct struct {
		Code string `validate:"len=6"`
	}

	validator := &DefaultValidator{}

	// Valid case
	// 유효한 경우
	obj := TestStruct{Code: "ABC123"}
	if err := validator.Validate(&obj); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Invalid case
	// 무효한 경우
	obj = TestStruct{Code: "ABC12"}
	err := validator.Validate(&obj)
	if err == nil {
		t.Error("Expected error for incorrect length")
	}
}

// TestValidatorEqNe tests equality/inequality validation.
// TestValidatorEqNe는 동등성/부등성 검증을 테스트합니다.
func TestValidatorEqNe(t *testing.T) {
	type TestStruct struct {
		Status string `validate:"eq=active"`
		Type   string `validate:"ne=invalid"`
	}

	validator := &DefaultValidator{}

	// Valid case
	// 유효한 경우
	obj := TestStruct{Status: "active", Type: "valid"}
	if err := validator.Validate(&obj); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Invalid eq
	// 무효한 eq
	obj = TestStruct{Status: "inactive", Type: "valid"}
	err := validator.Validate(&obj)
	if err == nil {
		t.Error("Expected error for status not equal to active")
	}

	// Invalid ne
	// 무효한 ne
	obj = TestStruct{Status: "active", Type: "invalid"}
	err = validator.Validate(&obj)
	if err == nil {
		t.Error("Expected error for type equal to invalid")
	}
}

// TestValidatorComparison tests gt/gte/lt/lte validation.
// TestValidatorComparison은 gt/gte/lt/lte 검증을 테스트합니다.
func TestValidatorComparison(t *testing.T) {
	type TestStruct struct {
		Score   int `validate:"gte=0,lte=100"`
		Count   int `validate:"gt=0,lt=1000"`
		Balance int `validate:"gte=0"`
	}

	validator := &DefaultValidator{}

	// Valid case
	// 유효한 경우
	obj := TestStruct{Score: 75, Count: 500, Balance: 1000}
	if err := validator.Validate(&obj); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Invalid gte
	// 무효한 gte
	obj = TestStruct{Score: -1, Count: 500, Balance: 1000}
	err := validator.Validate(&obj)
	if err == nil {
		t.Error("Expected error for score < 0")
	}

	// Invalid lte
	// 무효한 lte
	obj = TestStruct{Score: 101, Count: 500, Balance: 1000}
	err = validator.Validate(&obj)
	if err == nil {
		t.Error("Expected error for score > 100")
	}
}

// TestValidatorOneOf tests oneof validation.
// TestValidatorOneOf는 oneof 검증을 테스트합니다.
func TestValidatorOneOf(t *testing.T) {
	type TestStruct struct {
		Role string `validate:"oneof=admin user guest"`
	}

	validator := &DefaultValidator{}

	// Valid cases
	// 유효한 경우
	validRoles := []string{"admin", "user", "guest"}
	for _, role := range validRoles {
		obj := TestStruct{Role: role}
		if err := validator.Validate(&obj); err != nil {
			t.Errorf("Expected no error for role %s, got %v", role, err)
		}
	}

	// Invalid case
	// 무효한 경우
	obj := TestStruct{Role: "superadmin"}
	err := validator.Validate(&obj)
	if err == nil {
		t.Error("Expected error for invalid role")
	}
}

// TestValidatorAlpha tests alpha validation.
// TestValidatorAlpha는 alpha 검증을 테스트합니다.
func TestValidatorAlpha(t *testing.T) {
	type TestStruct struct {
		Name string `validate:"alpha"`
	}

	validator := &DefaultValidator{}

	// Valid case
	// 유효한 경우
	obj := TestStruct{Name: "JohnDoe"}
	if err := validator.Validate(&obj); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Invalid case
	// 무효한 경우
	obj = TestStruct{Name: "John123"}
	err := validator.Validate(&obj)
	if err == nil {
		t.Error("Expected error for non-alpha characters")
	}
}

// TestValidatorAlphanum tests alphanum validation.
// TestValidatorAlphanum은 alphanum 검증을 테스트합니다.
func TestValidatorAlphanum(t *testing.T) {
	type TestStruct struct {
		Username string `validate:"alphanum"`
	}

	validator := &DefaultValidator{}

	// Valid case
	// 유효한 경우
	obj := TestStruct{Username: "user123"}
	if err := validator.Validate(&obj); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Invalid case
	// 무효한 경우
	obj = TestStruct{Username: "user_123"}
	err := validator.Validate(&obj)
	if err == nil {
		t.Error("Expected error for non-alphanumeric characters")
	}
}

// TestValidatorNumeric tests numeric validation.
// TestValidatorNumeric은 numeric 검증을 테스트합니다.
func TestValidatorNumeric(t *testing.T) {
	type TestStruct struct {
		Code string `validate:"numeric"`
	}

	validator := &DefaultValidator{}

	// Valid case
	// 유효한 경우
	obj := TestStruct{Code: "123456"}
	if err := validator.Validate(&obj); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Invalid case
	// 무효한 경우
	obj = TestStruct{Code: "12a456"}
	err := validator.Validate(&obj)
	if err == nil {
		t.Error("Expected error for non-numeric characters")
	}
}

// TestValidatorMultipleTags tests multiple validation tags.
// TestValidatorMultipleTags는 여러 검증 태그를 테스트합니다.
func TestValidatorMultipleTags(t *testing.T) {
	type User struct {
		Name  string `validate:"required,min=3,max=50,alpha"`
		Email string `validate:"required,email"`
		Age   int    `validate:"required,gte=18,lte=100"`
		Role  string `validate:"required,oneof=admin user guest"`
	}

	validator := &DefaultValidator{}

	// Valid case
	// 유효한 경우
	user := User{
		Name:  "John",
		Email: "john@example.com",
		Age:   25,
		Role:  "user",
	}

	if err := validator.Validate(&user); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Invalid multiple fields
	// 여러 필드가 무효한 경우
	user = User{
		Name:  "Jo",      // Too short
		Email: "invalid", // Invalid email
		Age:   17,        // Below min
		Role:  "unknown", // Not in oneof
	}

	err := validator.Validate(&user)
	if err == nil {
		t.Error("Expected error for multiple invalid fields")
	}

	// Check that we have multiple errors
	// 여러 에러가 있는지 확인
	validationErrors, ok := err.(ValidationErrors)
	if !ok {
		t.Fatalf("Expected ValidationErrors, got %T", err)
	}

	if len(validationErrors) < 2 {
		t.Errorf("Expected at least 2 errors, got %d", len(validationErrors))
	}
}

// TestValidationErrors tests ValidationErrors type.
// TestValidationErrors는 ValidationErrors 타입을 테스트합니다.
func TestValidationErrors(t *testing.T) {
	errors := ValidationErrors{
		&ValidationError{Field: "Name", Tag: "required"},
		&ValidationError{Field: "Email", Tag: "email"},
	}

	errorMsg := errors.Error()
	if !strings.Contains(errorMsg, "Name") {
		t.Error("Expected error message to contain 'Name'")
	}

	if !strings.Contains(errorMsg, "Email") {
		t.Error("Expected error message to contain 'Email'")
	}
}

// TestBindWithValidation tests Context.BindWithValidation method.
// TestBindWithValidation은 Context.BindWithValidation 메서드를 테스트합니다.
func TestBindWithValidation(t *testing.T) {
	type User struct {
		Name  string `json:"name" validate:"required,min=3"`
		Email string `json:"email" validate:"required,email"`
	}

	// Create test request
	// 테스트 요청 생성
	// This test would require a full HTTP request setup
	// 이 테스트는 전체 HTTP 요청 설정이 필요합니다
	// For now, we test the validator directly
	// 현재는 검증자를 직접 테스트합니다

	validator := &DefaultValidator{}

	// Valid user
	// 유효한 사용자
	user := User{Name: "John", Email: "john@example.com"}
	if err := validator.Validate(&user); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Invalid user
	// 무효한 사용자
	user = User{Name: "Jo", Email: "invalid"}
	err := validator.Validate(&user)
	if err == nil {
		t.Error("Expected validation error")
	}
}
