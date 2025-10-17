package validation

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

// TestMinMaxLengthProperties tests properties of MinLength and MaxLength validators
// TestMinMaxLengthProperties는 MinLength와 MaxLength 검증기의 속성을 테스트합니다
func TestMinMaxLengthProperties(t *testing.T) {
	// Property 1: MinLength(n) followed by MaxLength(m) where n <= m should work
	// 속성 1: MinLength(n) 다음 MaxLength(m)을 호출하면 n <= m일 때 작동해야 함
	t.Run("MinLength then MaxLength with valid range", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			minLen := rand.Intn(50)
			maxLen := minLen + rand.Intn(50)
			strLen := minLen + rand.Intn(maxLen-minLen+1)
			str := randomString(strLen)

			v := New(str, "field")
			v.MinLength(minLen).MaxLength(maxLen)

			if len(v.GetErrors()) > 0 {
				t.Errorf("valid string length %d (min=%d, max=%d) failed validation", strLen, minLen, maxLen)
			}
		}
	})

	// Property 2: String shorter than MinLength must fail
	// 속성 2: MinLength보다 짧은 문자열은 실패해야 함
	t.Run("MinLength property", func(t *testing.T) {
		for i := 0; i < 50; i++ {
			minLen := rand.Intn(20) + 1 // At least 1
			strLen := rand.Intn(minLen)  // Shorter than min
			str := randomString(strLen)

			v := New(str, "field")
			v.MinLength(minLen)

			if len(v.GetErrors()) == 0 {
				t.Errorf("string length %d should fail MinLength(%d)", strLen, minLen)
			}
		}
	})

	// Property 3: String longer than MaxLength must fail
	// 속성 3: MaxLength보다 긴 문자열은 실패해야 함
	t.Run("MaxLength property", func(t *testing.T) {
		for i := 0; i < 50; i++ {
			maxLen := rand.Intn(20) + 1
			strLen := maxLen + rand.Intn(10) + 1 // Longer than max
			str := randomString(strLen)

			v := New(str, "field")
			v.MaxLength(maxLen)

			if len(v.GetErrors()) == 0 {
				t.Errorf("string length %d should fail MaxLength(%d)", strLen, maxLen)
			}
		}
	})
}

// TestNumericRangeProperties tests properties of Min/Max validators
// TestNumericRangeProperties는 Min/Max 검증기의 속성을 테스트합니다
func TestNumericRangeProperties(t *testing.T) {
	// Property 1: Value between Min and Max should pass
	// 속성 1: Min과 Max 사이의 값은 통과해야 함
	t.Run("value in range", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			min := rand.Intn(100)
			max := min + rand.Intn(100) + 1
			val := min + rand.Intn(max-min+1)

			v := New(val, "field")
			v.Min(float64(min)).Max(float64(max))

			if len(v.GetErrors()) > 0 {
				t.Errorf("value %d (min=%d, max=%d) should pass", val, min, max)
			}
		}
	})

	// Property 2: Value < Min should fail
	// 속성 2: Min보다 작은 값은 실패해야 함
	t.Run("value below min", func(t *testing.T) {
		for i := 0; i < 50; i++ {
			min := rand.Intn(100) + 1
			val := rand.Intn(min)

			v := New(val, "field")
			v.Min(float64(min))

			if len(v.GetErrors()) == 0 {
				t.Errorf("value %d should fail Min(%d)", val, min)
			}
		}
	})

	// Property 3: Value > Max should fail
	// 속성 3: Max보다 큰 값은 실패해야 함
	t.Run("value above max", func(t *testing.T) {
		for i := 0; i < 50; i++ {
			max := rand.Intn(100)
			val := max + rand.Intn(100) + 1

			v := New(val, "field")
			v.Max(float64(max))

			if len(v.GetErrors()) == 0 {
				t.Errorf("value %d should fail Max(%d)", val, max)
			}
		}
	})
}

// TestBeforeAfterProperties tests temporal ordering properties
// TestBeforeAfterProperties는 시간 순서 속성을 테스트합니다
func TestBeforeAfterProperties(t *testing.T) {
	// Property 1: time1 < time2 implies time1.Before(time2) should pass
	// 속성 1: time1 < time2이면 time1.Before(time2)가 통과해야 함
	t.Run("Before property", func(t *testing.T) {
		now := time.Now()
		for i := 0; i < 50; i++ {
			offset1 := rand.Intn(1000)
			offset2 := offset1 + rand.Intn(1000) + 1
			time1 := now.Add(time.Duration(offset1) * time.Second)
			time2 := now.Add(time.Duration(offset2) * time.Second)

			v := New(time1, "field")
			v.Before(time2)

			if len(v.GetErrors()) > 0 {
				t.Errorf("time1 should be before time2")
			}
		}
	})

	// Property 2: time1 > time2 implies time1.After(time2) should pass
	// 속성 2: time1 > time2이면 time1.After(time2)가 통과해야 함
	t.Run("After property", func(t *testing.T) {
		now := time.Now()
		for i := 0; i < 50; i++ {
			offset1 := rand.Intn(1000) + 1
			offset2 := rand.Intn(offset1)
			time1 := now.Add(time.Duration(offset1) * time.Second)
			time2 := now.Add(time.Duration(offset2) * time.Second)

			v := New(time1, "field")
			v.After(time2)

			if len(v.GetErrors()) > 0 {
				t.Errorf("time1 should be after time2")
			}
		}
	})

	// Property 3: BetweenTime with time1 < time2 < time3
	// 속성 3: time1 < time2 < time3일 때 BetweenTime
	t.Run("BetweenTime property", func(t *testing.T) {
		now := time.Now()
		for i := 0; i < 50; i++ {
			offset1 := rand.Intn(1000)
			offset2 := offset1 + rand.Intn(500) + 1
			offset3 := offset2 + rand.Intn(500) + 1

			start := now.Add(time.Duration(offset1) * time.Second)
			middle := now.Add(time.Duration(offset2) * time.Second)
			end := now.Add(time.Duration(offset3) * time.Second)

			v := New(middle, "field")
			v.BetweenTime(start, end)

			if len(v.GetErrors()) > 0 {
				t.Errorf("middle time should be between start and end")
			}
		}
	})
}

// TestStopOnErrorProperty tests that StopOnError stops at first error
// TestStopOnErrorProperty는 StopOnError가 첫 에러에서 멈추는지 테스트합니다
func TestStopOnErrorProperty(t *testing.T) {
	// Property: With StopOnError, should only get one error even with multiple failures
	// 속성: StopOnError 사용 시 여러 실패가 있어도 하나의 에러만 받아야 함
	t.Run("stops at first error", func(t *testing.T) {
		for i := 0; i < 50; i++ {
			// Create conditions that will fail multiple validators
			v := New("", "field")
			v.StopOnError().
				Required().         // Will fail
				MinLength(10).      // Would fail
				Email().            // Would fail
				MaxLength(5)        // Would fail

			errors := v.GetErrors()
			if len(errors) != 1 {
				t.Errorf("expected 1 error with StopOnError, got %d", len(errors))
			}

			// Error should be from Required validator
			if errors[0].Rule != "required" {
				t.Errorf("first error should be 'required', got '%s'", errors[0].Rule)
			}
		}
	})
}

// TestValidatorIdempotence tests that validation is idempotent
// TestValidatorIdempotence는 검증이 멱등성을 가지는지 테스트합니다
func TestValidatorIdempotence(t *testing.T) {
	// Property: Calling Validate() multiple times should give same result
	// 속성: Validate()를 여러 번 호출해도 같은 결과를 얻어야 함
	t.Run("validate is idempotent", func(t *testing.T) {
		for i := 0; i < 50; i++ {
			str := randomString(rand.Intn(20))
			v := New(str, "field")
			v.Required().MinLength(5).MaxLength(15)

			// First validation
			err1 := v.Validate()
			errors1 := v.GetErrors()

			// Second validation
			err2 := v.Validate()
			errors2 := v.GetErrors()

			// Should get same results
			if (err1 == nil) != (err2 == nil) {
				t.Error("validate should be idempotent (error status changed)")
			}

			if len(errors1) != len(errors2) {
				t.Errorf("validate should be idempotent (error count changed: %d vs %d)", len(errors1), len(errors2))
			}
		}
	})
}

// TestCustomMessageProperty tests that custom messages are preserved
// TestCustomMessageProperty는 커스텀 메시지가 보존되는지 테스트합니다
func TestCustomMessageProperty(t *testing.T) {
	// Property: Custom messages should always be used when set
	// 속성: 커스텀 메시지가 설정되면 항상 사용되어야 함
	t.Run("custom message is preserved", func(t *testing.T) {
		for i := 0; i < 50; i++ {
			customMsg := fmt.Sprintf("Custom error %d", i)

			v := New("", "field")
			v.WithCustomMessage("required", customMsg)
			v.Required()

			errors := v.GetErrors()
			if len(errors) == 0 {
				t.Fatal("expected error for empty required field")
			}

			if errors[0].Message != customMsg {
				t.Errorf("expected custom message '%s', got '%s'", customMsg, errors[0].Message)
			}
		}
	})
}

// Helper function to generate random strings
// 랜덤 문자열을 생성하는 헬퍼 함수
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var sb strings.Builder
	sb.Grow(length)
	for i := 0; i < length; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}
