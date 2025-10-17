package websvrutil

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// ============================================================================
// Validation
// 검증
// ============================================================================

// Validator defines the interface for custom validators.
// Validator는 커스텀 검증자를 위한 인터페이스를 정의합니다.
type Validator interface {
	Validate(interface{}) error
}

// ValidationError represents a validation error.
// ValidationError는 검증 에러를 나타냅니다.
type ValidationError struct {
	Field string
	Tag   string
	Value interface{}
}

// Error returns the error message.
// Error는 에러 메시지를 반환합니다.
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed for field '%s' with tag '%s'", e.Field, e.Tag)
}

// ValidationErrors represents multiple validation errors.
// ValidationErrors는 여러 검증 에러를 나타냅니다.
type ValidationErrors []*ValidationError

// Error returns the concatenated error messages.
// Error는 연결된 에러 메시지를 반환합니다.
func (e ValidationErrors) Error() string {
	if len(e) == 0 {
		return ""
	}
	messages := make([]string, len(e))
	for i, err := range e {
		messages[i] = err.Error()
	}
	return strings.Join(messages, "; ")
}

// DefaultValidator is the built-in validator.
// DefaultValidator는 내장 검증자입니다.
type DefaultValidator struct{}

// Validate validates a struct using validation tags.
// Validate는 검증 태그를 사용하여 구조체를 검증합니다.
//
// Supported tags
// 지원되는 태그:
// - required: field must not be empty
// 필드가 비어 있지 않아야 함
// - email: field must be a valid email address
// 유효한 이메일 주소여야 함
// - min=<value>: minimum length/value
// 최소 길이/값
// - max=<value>: maximum length/value
// 최대 길이/값
// - len=<value>: exact length
// 정확한 길이
// - eq=<value>: equal to value
// 값과 같아야 함
// - ne=<value>: not equal to value
// 값과 같지 않아야 함
// - gt=<value>: greater than value
// 값보다 커야 함
// - gte=<value>: greater than or equal to value
// 값보다 크거나 같아야 함
// - lt=<value>: less than value
// 값보다 작아야 함
// - lte=<value>: less than or equal to value
// 값보다 작거나 같아야 함
// - oneof=<values>: one of the values (comma-separated)
// 값 중 하나 (쉼표로 구분)
// - alpha: alphabetic characters only
// 알파벳 문자만
// - alphanum: alphanumeric characters only
// 영숫자 문자만
// - numeric: numeric characters only
// 숫자 문자만
//
// Example
// 예제:
//
//	type User struct {
//	    Name  string `validate:"required,min=3,max=50"`
//	    Email string `validate:"required,email"`
//	    Age   int    `validate:"required,gte=18,lte=100"`
//	    Role  string `validate:"required,oneof=admin,user,guest"`
//	}
func (v *DefaultValidator) Validate(obj interface{}) error {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return fmt.Errorf("validate: expected struct, got %s", val.Kind())
	}

	var errors ValidationErrors

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		// Get validation tag
		// 검증 태그 가져오기
		tag := fieldType.Tag.Get("validate")
		if tag == "" {
			continue
		}

		// Parse and validate tags
		// 태그 파싱 및 검증
		if err := validateField(fieldType.Name, field, tag); err != nil {
			if ve, ok := err.(*ValidationError); ok {
				errors = append(errors, ve)
			} else if ves, ok := err.(ValidationErrors); ok {
				errors = append(errors, ves...)
			}
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// validateField validates a single field.
// validateField는 단일 필드를 검증합니다.
func validateField(fieldName string, field reflect.Value, tag string) error {
	// Split tag by comma
	// 쉼표로 태그 분할
	rules := strings.Split(tag, ",")

	for _, rule := range rules {
		rule = strings.TrimSpace(rule)
		if rule == "" {
			continue
		}

		// Parse rule (e.g., "min=3" -> "min", "3")
		// 규칙 파싱 (예: "min=3" -> "min", "3")
		parts := strings.SplitN(rule, "=", 2)
		ruleName := parts[0]
		var ruleValue string
		if len(parts) > 1 {
			ruleValue = parts[1]
		}

		// Validate based on rule
		// 규칙에 따라 검증
		if err := applyRule(fieldName, field, ruleName, ruleValue); err != nil {
			return err
		}
	}

	return nil
}

// applyRule applies a validation rule to a field.
// applyRule은 필드에 검증 규칙을 적용합니다.
func applyRule(fieldName string, field reflect.Value, ruleName, ruleValue string) error {
	switch ruleName {
	case "required":
		return validateRequired(fieldName, field)
	case "email":
		return validateEmail(fieldName, field)
	case "min":
		return validateMin(fieldName, field, ruleValue)
	case "max":
		return validateMax(fieldName, field, ruleValue)
	case "len":
		return validateLen(fieldName, field, ruleValue)
	case "eq":
		return validateEq(fieldName, field, ruleValue)
	case "ne":
		return validateNe(fieldName, field, ruleValue)
	case "gt":
		return validateGt(fieldName, field, ruleValue)
	case "gte":
		return validateGte(fieldName, field, ruleValue)
	case "lt":
		return validateLt(fieldName, field, ruleValue)
	case "lte":
		return validateLte(fieldName, field, ruleValue)
	case "oneof":
		return validateOneOf(fieldName, field, ruleValue)
	case "alpha":
		return validateAlpha(fieldName, field)
	case "alphanum":
		return validateAlphanum(fieldName, field)
	case "numeric":
		return validateNumeric(fieldName, field)
	default:
		// Unknown rule, skip
		// 알 수 없는 규칙, 건너뛰기
		return nil
	}
}

// validateRequired validates that a field is not empty.
// validateRequired는 필드가 비어 있지 않은지 검증합니다.
func validateRequired(fieldName string, field reflect.Value) error {
	if isZero(field) {
		return &ValidationError{Field: fieldName, Tag: "required", Value: field.Interface()}
	}
	return nil
}

// validateEmail validates that a field is a valid email.
// validateEmail은 필드가 유효한 이메일인지 검증합니다.
func validateEmail(fieldName string, field reflect.Value) error {
	if field.Kind() != reflect.String {
		return nil
	}

	email := field.String()
	if email == "" {
		// Empty is handled by required
		// 빈 값은 required에서 처리
		return nil
	}

	// Simple email regex
	// 간단한 이메일 정규식
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return &ValidationError{Field: fieldName, Tag: "email", Value: email}
	}

	return nil
}

// validateMin validates minimum length/value.
// validateMin은 최소 길이/값을 검증합니다.
func validateMin(fieldName string, field reflect.Value, min string) error {
	minVal, err := strconv.ParseInt(min, 10, 64)
	if err != nil {
		return nil
	}

	switch field.Kind() {
	case reflect.String:
		if int64(len(field.String())) < minVal {
			return &ValidationError{Field: fieldName, Tag: "min=" + min, Value: field.Interface()}
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if field.Int() < minVal {
			return &ValidationError{Field: fieldName, Tag: "min=" + min, Value: field.Interface()}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if field.Uint() < uint64(minVal) {
			return &ValidationError{Field: fieldName, Tag: "min=" + min, Value: field.Interface()}
		}
	case reflect.Float32, reflect.Float64:
		if field.Float() < float64(minVal) {
			return &ValidationError{Field: fieldName, Tag: "min=" + min, Value: field.Interface()}
		}
	case reflect.Slice, reflect.Array, reflect.Map:
		if int64(field.Len()) < minVal {
			return &ValidationError{Field: fieldName, Tag: "min=" + min, Value: field.Interface()}
		}
	}

	return nil
}

// validateMax validates maximum length/value.
// validateMax는 최대 길이/값을 검증합니다.
func validateMax(fieldName string, field reflect.Value, max string) error {
	maxVal, err := strconv.ParseInt(max, 10, 64)
	if err != nil {
		return nil
	}

	switch field.Kind() {
	case reflect.String:
		if int64(len(field.String())) > maxVal {
			return &ValidationError{Field: fieldName, Tag: "max=" + max, Value: field.Interface()}
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if field.Int() > maxVal {
			return &ValidationError{Field: fieldName, Tag: "max=" + max, Value: field.Interface()}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if field.Uint() > uint64(maxVal) {
			return &ValidationError{Field: fieldName, Tag: "max=" + max, Value: field.Interface()}
		}
	case reflect.Float32, reflect.Float64:
		if field.Float() > float64(maxVal) {
			return &ValidationError{Field: fieldName, Tag: "max=" + max, Value: field.Interface()}
		}
	case reflect.Slice, reflect.Array, reflect.Map:
		if int64(field.Len()) > maxVal {
			return &ValidationError{Field: fieldName, Tag: "max=" + max, Value: field.Interface()}
		}
	}

	return nil
}

// validateLen validates exact length.
// validateLen은 정확한 길이를 검증합니다.
func validateLen(fieldName string, field reflect.Value, length string) error {
	lenVal, err := strconv.ParseInt(length, 10, 64)
	if err != nil {
		return nil
	}

	switch field.Kind() {
	case reflect.String:
		if int64(len(field.String())) != lenVal {
			return &ValidationError{Field: fieldName, Tag: "len=" + length, Value: field.Interface()}
		}
	case reflect.Slice, reflect.Array, reflect.Map:
		if int64(field.Len()) != lenVal {
			return &ValidationError{Field: fieldName, Tag: "len=" + length, Value: field.Interface()}
		}
	}

	return nil
}

// validateEq validates equality.
// validateEq는 동등성을 검증합니다.
func validateEq(fieldName string, field reflect.Value, value string) error {
	switch field.Kind() {
	case reflect.String:
		if field.String() != value {
			return &ValidationError{Field: fieldName, Tag: "eq=" + value, Value: field.Interface()}
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intVal, _ := strconv.ParseInt(value, 10, 64)
		if field.Int() != intVal {
			return &ValidationError{Field: fieldName, Tag: "eq=" + value, Value: field.Interface()}
		}
	}

	return nil
}

// validateNe validates inequality.
// validateNe는 부등성을 검증합니다.
func validateNe(fieldName string, field reflect.Value, value string) error {
	switch field.Kind() {
	case reflect.String:
		if field.String() == value {
			return &ValidationError{Field: fieldName, Tag: "ne=" + value, Value: field.Interface()}
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intVal, _ := strconv.ParseInt(value, 10, 64)
		if field.Int() == intVal {
			return &ValidationError{Field: fieldName, Tag: "ne=" + value, Value: field.Interface()}
		}
	}

	return nil
}

// validateGt validates greater than.
// validateGt는 보다 큰 값을 검증합니다.
func validateGt(fieldName string, field reflect.Value, value string) error {
	intVal, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil
	}

	switch field.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if field.Int() <= intVal {
			return &ValidationError{Field: fieldName, Tag: "gt=" + value, Value: field.Interface()}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if field.Uint() <= uint64(intVal) {
			return &ValidationError{Field: fieldName, Tag: "gt=" + value, Value: field.Interface()}
		}
	case reflect.Float32, reflect.Float64:
		if field.Float() <= float64(intVal) {
			return &ValidationError{Field: fieldName, Tag: "gt=" + value, Value: field.Interface()}
		}
	}

	return nil
}

// validateGte validates greater than or equal.
// validateGte는 보다 크거나 같은 값을 검증합니다.
func validateGte(fieldName string, field reflect.Value, value string) error {
	intVal, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil
	}

	switch field.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if field.Int() < intVal {
			return &ValidationError{Field: fieldName, Tag: "gte=" + value, Value: field.Interface()}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if field.Uint() < uint64(intVal) {
			return &ValidationError{Field: fieldName, Tag: "gte=" + value, Value: field.Interface()}
		}
	case reflect.Float32, reflect.Float64:
		if field.Float() < float64(intVal) {
			return &ValidationError{Field: fieldName, Tag: "gte=" + value, Value: field.Interface()}
		}
	}

	return nil
}

// validateLt validates less than.
// validateLt는 보다 작은 값을 검증합니다.
func validateLt(fieldName string, field reflect.Value, value string) error {
	intVal, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil
	}

	switch field.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if field.Int() >= intVal {
			return &ValidationError{Field: fieldName, Tag: "lt=" + value, Value: field.Interface()}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if field.Uint() >= uint64(intVal) {
			return &ValidationError{Field: fieldName, Tag: "lt=" + value, Value: field.Interface()}
		}
	case reflect.Float32, reflect.Float64:
		if field.Float() >= float64(intVal) {
			return &ValidationError{Field: fieldName, Tag: "lt=" + value, Value: field.Interface()}
		}
	}

	return nil
}

// validateLte validates less than or equal.
// validateLte는 보다 작거나 같은 값을 검증합니다.
func validateLte(fieldName string, field reflect.Value, value string) error {
	intVal, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil
	}

	switch field.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if field.Int() > intVal {
			return &ValidationError{Field: fieldName, Tag: "lte=" + value, Value: field.Interface()}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if field.Uint() > uint64(intVal) {
			return &ValidationError{Field: fieldName, Tag: "lte=" + value, Value: field.Interface()}
		}
	case reflect.Float32, reflect.Float64:
		if field.Float() > float64(intVal) {
			return &ValidationError{Field: fieldName, Tag: "lte=" + value, Value: field.Interface()}
		}
	}

	return nil
}

// validateOneOf validates that value is one of the specified values.
// validateOneOf는 값이 지정된 값 중 하나인지 검증합니다.
func validateOneOf(fieldName string, field reflect.Value, values string) error {
	if field.Kind() != reflect.String {
		return nil
	}

	fieldValue := field.String()
	// Split by comma
	// 쉼표로 분할
	validValues := strings.Split(values, " ")

	for _, valid := range validValues {
		if fieldValue == strings.TrimSpace(valid) {
			return nil
		}
	}

	return &ValidationError{Field: fieldName, Tag: "oneof=" + values, Value: fieldValue}
}

// validateAlpha validates alphabetic characters only.
// validateAlpha는 알파벳 문자만 검증합니다.
func validateAlpha(fieldName string, field reflect.Value) error {
	if field.Kind() != reflect.String {
		return nil
	}

	str := field.String()
	if str == "" {
		return nil
	}

	alphaRegex := regexp.MustCompile(`^[a-zA-Z]+$`)
	if !alphaRegex.MatchString(str) {
		return &ValidationError{Field: fieldName, Tag: "alpha", Value: str}
	}

	return nil
}

// validateAlphanum validates alphanumeric characters only.
// validateAlphanum은 영숫자 문자만 검증합니다.
func validateAlphanum(fieldName string, field reflect.Value) error {
	if field.Kind() != reflect.String {
		return nil
	}

	str := field.String()
	if str == "" {
		return nil
	}

	alphanumRegex := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if !alphanumRegex.MatchString(str) {
		return &ValidationError{Field: fieldName, Tag: "alphanum", Value: str}
	}

	return nil
}

// validateNumeric validates numeric characters only.
// validateNumeric은 숫자 문자만 검증합니다.
func validateNumeric(fieldName string, field reflect.Value) error {
	if field.Kind() != reflect.String {
		return nil
	}

	str := field.String()
	if str == "" {
		return nil
	}

	numericRegex := regexp.MustCompile(`^[0-9]+$`)
	if !numericRegex.MatchString(str) {
		return &ValidationError{Field: fieldName, Tag: "numeric", Value: str}
	}

	return nil
}

// isZero checks if a value is zero.
// isZero는 값이 제로인지 확인합니다.
func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	case reflect.Slice, reflect.Map, reflect.Array:
		return v.Len() == 0
	}

	return false
}

// BindWithValidation binds and validates the request data.
// BindWithValidation은 요청 데이터를 바인딩하고 검증합니다.
//
// Example
// 예제:
//
//	type User struct {
//	    Name  string `json:"name" validate:"required,min=3"`
//	    Email string `json:"email" validate:"required,email"`
//	}
//	var user User
//	if err := ctx.BindWithValidation(&user); err != nil {
//	    return ctx.Error(400, err.Error())
//	}
func (c *Context) BindWithValidation(obj interface{}) error {
	// First bind the data
	// 먼저 데이터 바인딩
	if err := c.Bind(obj); err != nil {
		return err
	}

	// Then validate
	// 그 다음 검증
	validator := &DefaultValidator{}
	return validator.Validate(obj)
}
