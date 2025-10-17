package validation

import (
	"testing"
	"time"
)

// Benchmark tests for validation package
// validation 패키지의 벤치마크 테스트

// BenchmarkRequired benchmarks the Required validator
// BenchmarkRequired는 Required 검증기를 벤치마크합니다
func BenchmarkRequired(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("test value", "field")
		v.Required()
		_ = v.Validate()
	}
}

// BenchmarkMinLength benchmarks the MinLength validator
// BenchmarkMinLength는 MinLength 검증기를 벤치마크합니다
func BenchmarkMinLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("test string value", "field")
		v.MinLength(5)
		_ = v.Validate()
	}
}

// BenchmarkMaxLength benchmarks the MaxLength validator
// BenchmarkMaxLength는 MaxLength 검증기를 벤치마크합니다
func BenchmarkMaxLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("test", "field")
		v.MaxLength(10)
		_ = v.Validate()
	}
}

// BenchmarkEmail benchmarks the Email validator
// BenchmarkEmail는 Email 검증기를 벤치마크합니다
func BenchmarkEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("test@example.com", "email")
		v.Email()
		_ = v.Validate()
	}
}

// BenchmarkURL benchmarks the URL validator
// BenchmarkURL는 URL 검증기를 벤치마크합니다
func BenchmarkURL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("https://example.com", "url")
		v.URL()
		_ = v.Validate()
	}
}

// BenchmarkMin benchmarks the Min validator
// BenchmarkMin는 Min 검증기를 벤치마크합니다
func BenchmarkMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New(100, "value")
		v.Min(0)
		_ = v.Validate()
	}
}

// BenchmarkMax benchmarks the Max validator
// BenchmarkMax는 Max 검증기를 벤치마크합니다
func BenchmarkMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New(50, "value")
		v.Max(100)
		_ = v.Validate()
	}
}

// BenchmarkBetween benchmarks the Min and Max validators together
// BenchmarkBetween는 Min과 Max 검증기를 함께 벤치마크합니다
func BenchmarkBetween(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New(50, "value")
		v.Min(0).Max(100)
		_ = v.Validate()
	}
}

// BenchmarkIn benchmarks the In validator
// BenchmarkIn는 In 검증기를 벤치마크합니다
func BenchmarkIn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("apple", "fruit")
		v.In("apple", "banana", "orange", "grape", "melon")
		_ = v.Validate()
	}
}

// BenchmarkNotIn benchmarks the NotIn validator
// BenchmarkNotIn는 NotIn 검증기를 벤치마크합니다
func BenchmarkNotIn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("pear", "fruit")
		v.NotIn("apple", "banana", "orange")
		_ = v.Validate()
	}
}

// BenchmarkArrayLength benchmarks the ArrayLength validator
// BenchmarkArrayLength는 ArrayLength 검증기를 벤치마크합니다
func BenchmarkArrayLength(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		v := New(arr, "numbers")
		v.ArrayLength(5)
		_ = v.Validate()
	}
}

// BenchmarkArrayUnique benchmarks the ArrayUnique validator
// BenchmarkArrayUnique는 ArrayUnique 검증기를 벤치마크합니다
func BenchmarkArrayUnique(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		v := New(arr, "numbers")
		v.ArrayUnique()
		_ = v.Validate()
	}
}

// BenchmarkMapHasKey benchmarks the MapHasKey validator
// BenchmarkMapHasKey는 MapHasKey 검증기를 벤치마크합니다
func BenchmarkMapHasKey(b *testing.B) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for i := 0; i < b.N; i++ {
		v := New(m, "data")
		v.MapHasKey("a")
		_ = v.Validate()
	}
}

// BenchmarkMapHasKeys benchmarks the MapHasKeys validator
// BenchmarkMapHasKeys는 MapHasKeys 검증기를 벤치마크합니다
func BenchmarkMapHasKeys(b *testing.B) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	for i := 0; i < b.N; i++ {
		v := New(m, "data")
		v.MapHasKeys("a", "b", "c")
		_ = v.Validate()
	}
}

// BenchmarkEquals benchmarks the Equals validator
// BenchmarkEquals는 Equals 검증기를 벤치마크합니다
func BenchmarkEquals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New(100, "value")
		v.Equals(100)
		_ = v.Validate()
	}
}

// BenchmarkBefore benchmarks the Before validator
// BenchmarkBefore는 Before 검증기를 벤치마크합니다
func BenchmarkBefore(b *testing.B) {
	now := time.Now()
	yesterday := now.Add(-24 * time.Hour)
	for i := 0; i < b.N; i++ {
		v := New(yesterday, "date")
		v.Before(now)
		_ = v.Validate()
	}
}

// BenchmarkAfter benchmarks the After validator
// BenchmarkAfter는 After 검증기를 벤치마크합니다
func BenchmarkAfter(b *testing.B) {
	now := time.Now()
	tomorrow := now.Add(24 * time.Hour)
	for i := 0; i < b.N; i++ {
		v := New(tomorrow, "date")
		v.After(now)
		_ = v.Validate()
	}
}

// BenchmarkCustom benchmarks the Custom validator
// BenchmarkCustom는 Custom 검증기를 벤치마크합니다
func BenchmarkCustom(b *testing.B) {
	customFunc := func(val interface{}) bool {
		str, ok := val.(string)
		return ok && len(str) > 0
	}
	for i := 0; i < b.N; i++ {
		v := New("test", "field")
		v.Custom(customFunc, "custom error")
		_ = v.Validate()
	}
}

// BenchmarkMultipleValidators benchmarks chaining multiple validators
// BenchmarkMultipleValidators는 여러 검증기 체이닝을 벤치마크합니다
func BenchmarkMultipleValidators(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("test@example.com", "email")
		v.Required().MinLength(5).MaxLength(50).Email()
		_ = v.Validate()
	}
}

// BenchmarkStopOnError benchmarks validation with StopOnError
// BenchmarkStopOnError는 StopOnError를 사용한 검증을 벤치마크합니다
func BenchmarkStopOnError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("", "field")
		v.StopOnError()
		v.Required().MinLength(5).MaxLength(50)
		_ = v.Validate()
	}
}

// BenchmarkMultiValidator benchmarks MultiValidator with multiple fields
// BenchmarkMultiValidator는 여러 필드를 가진 MultiValidator를 벤치마크합니다
func BenchmarkMultiValidator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mv := NewValidator()
		mv.Field("john@example.com", "email").Required().Email()
		mv.Field("John Doe", "name").Required().MinLength(2)
		mv.Field(25, "age").Required().Min(18)
		_ = mv.Validate()
	}
}

// BenchmarkValidationErrors benchmarks error collection
// BenchmarkValidationErrors는 에러 수집을 벤치마크합니다
func BenchmarkValidationErrors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("", "field")
		v.Required()
		v.MinLength(10)
		v.Email()
		err := v.Validate()
		if err != nil {
			verrs := err.(ValidationErrors)
			_ = verrs.Error()
		}
	}
}
