package validation

import (
	"runtime"
	"testing"
	"time"
)

// Performance tests for validation package
// validation 패키지의 성능 테스트

// TestPerformance_LargeDataset tests validation performance with large datasets
// TestPerformance_LargeDataset는 큰 데이터셋을 사용한 검증 성능을 테스트합니다
func TestPerformance_LargeDataset(t *testing.T) {
	t.Run("large string validation", func(t *testing.T) {
		// Create large string (10KB)
		// 큰 문자열 생성 (10KB)
		size := 10000
		data := make([]byte, size)
		for i := 0; i < size; i++ {
			data[i] = 'a'
		}
		str := string(data)

		start := time.Now()
		v := New(str, "large_string")
		v.Required().MinLength(100).MaxLength(20000)
		err := v.Validate()
		elapsed := time.Since(start)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		t.Logf("Validated large string (%d bytes) in %v", len(str), elapsed)

		// Performance assertion / 성능 단언
		// Should complete within 10ms
		// 10ms 이내에 완료되어야 함
		maxDuration := 10 * time.Millisecond
		if elapsed > maxDuration {
			t.Errorf("operation took %v, expected less than %v", elapsed, maxDuration)
		}
	})

	t.Run("large array validation", func(t *testing.T) {
		// Create large array (1000 elements)
		// 큰 배열 생성 (1000개 요소)
		size := 1000
		arr := make([]int, size)
		for i := 0; i < size; i++ {
			arr[i] = i
		}

		start := time.Now()
		v := New(arr, "large_array")
		v.ArrayMinLength(100).ArrayMaxLength(2000).ArrayUnique()
		err := v.Validate()
		elapsed := time.Since(start)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		t.Logf("Validated large array (%d elements) in %v", len(arr), elapsed)

		// Performance assertion / 성능 단언
		// Should complete within 50ms
		// 50ms 이내에 완료되어야 함
		maxDuration := 50 * time.Millisecond
		if elapsed > maxDuration {
			t.Errorf("operation took %v, expected less than %v", elapsed, maxDuration)
		}
	})

	t.Run("multiple validators chaining", func(t *testing.T) {
		email := "user@example.com"

		start := time.Now()
		v := New(email, "email")
		v.Required().
			MinLength(5).
			MaxLength(100).
			Email().
			Regex(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
		err := v.Validate()
		elapsed := time.Since(start)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		t.Logf("Validated with multiple validators in %v", elapsed)

		// Performance assertion / 성능 단언
		// Should complete within 5ms
		// 5ms 이내에 완료되어야 함
		maxDuration := 5 * time.Millisecond
		if elapsed > maxDuration {
			t.Errorf("operation took %v, expected less than %v", elapsed, maxDuration)
		}
	})
}

// TestPerformance_MemoryUsage tests memory consumption during validation
// TestPerformance_MemoryUsage는 검증 중 메모리 소비를 테스트합니다
func TestPerformance_MemoryUsage(t *testing.T) {
	var m runtime.MemStats

	t.Run("single validator memory", func(t *testing.T) {
		runtime.GC()
		runtime.ReadMemStats(&m)
		before := m.Alloc

		// Perform validation operations / 검증 작업 수행
		iterations := 1000
		for i := 0; i < iterations; i++ {
			v := New("test@example.com", "email")
			v.Required().Email()
			_ = v.Validate()
		}

		runtime.ReadMemStats(&m)
		after := m.Alloc
		used := after - before

		t.Logf("Memory used for %d validations: %d bytes (%.2f KB)",
			iterations, used, float64(used)/1024)

		// Memory limit assertion / 메모리 제한 단언
		// Should use less than 10MB for 1000 validations
		// 1000개 검증에 10MB 미만 사용해야 함
		maxMemory := uint64(10 * 1024 * 1024) // 10 MB
		if used > maxMemory {
			t.Errorf("used %d bytes, expected less than %d bytes", used, maxMemory)
		}
	})

	t.Run("multi validator memory", func(t *testing.T) {
		runtime.GC()
		runtime.ReadMemStats(&m)
		before := m.Alloc

		// Perform multi-field validation operations / 다중 필드 검증 작업 수행
		iterations := 1000
		for i := 0; i < iterations; i++ {
			mv := NewValidator()
			mv.Field("user@example.com", "email").Required().Email()
			mv.Field("John Doe", "name").Required().MinLength(2)
			mv.Field(25, "age").Required().Min(18).Max(100)
			_ = mv.Validate()
		}

		runtime.ReadMemStats(&m)
		after := m.Alloc
		used := after - before

		t.Logf("Memory used for %d multi-field validations: %d bytes (%.2f KB)",
			iterations, used, float64(used)/1024)

		// Memory limit assertion / 메모리 제한 단언
		// Should use less than 20MB for 1000 multi-field validations
		// 1000개 다중 필드 검증에 20MB 미만 사용해야 함
		maxMemory := uint64(20 * 1024 * 1024) // 20 MB
		if used > maxMemory {
			t.Errorf("used %d bytes, expected less than %d bytes", used, maxMemory)
		}
	})

	t.Run("error collection memory", func(t *testing.T) {
		runtime.GC()
		runtime.ReadMemStats(&m)
		before := m.Alloc

		// Generate validation errors / 검증 에러 생성
		iterations := 100
		for i := 0; i < iterations; i++ {
			v := New("", "field")
			v.Required()
			v.MinLength(10)
			v.MaxLength(5) // Will fail / 실패할 것임
			v.Email()
			err := v.Validate()
			if err != nil {
				verrs := err.(ValidationErrors)
				_ = verrs.Error()
				// Access ValidationErrors as a slice
				_ = len(verrs)
			}
		}

		runtime.ReadMemStats(&m)
		after := m.Alloc
		used := after - before

		t.Logf("Memory used for %d error collections: %d bytes (%.2f KB)",
			iterations, used, float64(used)/1024)

		// Memory limit assertion / 메모리 제한 단언
		// Should use less than 5MB for 100 error collections
		// 100개 에러 수집에 5MB 미만 사용해야 함
		maxMemory := uint64(5 * 1024 * 1024) // 5 MB
		if used > maxMemory {
			t.Errorf("used %d bytes, expected less than %d bytes", used, maxMemory)
		}
	})
}

// TestPerformance_ValidationSpeed tests validation speed for different validator types
// TestPerformance_ValidationSpeed는 다양한 검증기 유형의 검증 속도를 테스트합니다
func TestPerformance_ValidationSpeed(t *testing.T) {
	iterations := 10000

	tests := []struct {
		name     string
		setup    func() *Validator
		maxTime  time.Duration
	}{
		{
			name: "simple required",
			setup: func() *Validator {
				return New("test", "field").Required()
			},
			maxTime: 50 * time.Millisecond,
		},
		{
			name: "string length",
			setup: func() *Validator {
				return New("test string", "field").MinLength(5).MaxLength(20)
			},
			maxTime: 50 * time.Millisecond,
		},
		{
			name: "email validation",
			setup: func() *Validator {
				return New("user@example.com", "email").Email()
			},
			maxTime: 100 * time.Millisecond,
		},
		{
			name: "regex validation",
			setup: func() *Validator {
				return New("abc123", "code").Regex(`^[a-z0-9]+$`)
			},
			maxTime: 200 * time.Millisecond,
		},
		{
			name: "numeric range",
			setup: func() *Validator {
				return New(50, "age").Min(0).Max(100)
			},
			maxTime: 50 * time.Millisecond,
		},
		{
			name: "array operations",
			setup: func() *Validator {
				arr := []int{1, 2, 3, 4, 5}
				return New(arr, "numbers").ArrayMinLength(1).ArrayMaxLength(10)
			},
			maxTime: 50 * time.Millisecond,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			for i := 0; i < iterations; i++ {
				v := tt.setup()
				_ = v.Validate()
			}
			elapsed := time.Since(start)

			opsPerSec := float64(iterations) / elapsed.Seconds()
			t.Logf("%s: %d operations in %v (%.0f ops/sec)",
				tt.name, iterations, elapsed, opsPerSec)

			if elapsed > tt.maxTime {
				t.Errorf("%s took %v, expected less than %v",
					tt.name, elapsed, tt.maxTime)
			}

			// Minimum performance requirement / 최소 성능 요구사항
			// At least 1000 operations per second
			// 초당 최소 1000개 작업
			minOpsPerSec := 1000.0
			if opsPerSec < minOpsPerSec {
				t.Errorf("performance = %.0f ops/sec, want >= %.0f ops/sec",
					opsPerSec, minOpsPerSec)
			}
		})
	}
}

// TestPerformance_StopOnErrorImpact tests the performance impact of StopOnError
// TestPerformance_StopOnErrorImpact는 StopOnError의 성능 영향을 테스트합니다
func TestPerformance_StopOnErrorImpact(t *testing.T) {
	iterations := 10000

	t.Run("without StopOnError", func(t *testing.T) {
		start := time.Now()
		for i := 0; i < iterations; i++ {
			v := New("", "field")
			v.Required().MinLength(10).MaxLength(5).Email()
			_ = v.Validate()
		}
		elapsed := time.Since(start)
		t.Logf("Without StopOnError: %d operations in %v", iterations, elapsed)
	})

	t.Run("with StopOnError", func(t *testing.T) {
		start := time.Now()
		for i := 0; i < iterations; i++ {
			v := New("", "field")
			v.StopOnError()
			v.Required().MinLength(10).MaxLength(5).Email()
			_ = v.Validate()
		}
		elapsed := time.Since(start)
		t.Logf("With StopOnError: %d operations in %v", iterations, elapsed)

		// StopOnError should be faster (stops at first error)
		// StopOnError가 더 빨라야 함 (첫 에러에서 중단)
		maxDuration := 100 * time.Millisecond
		if elapsed > maxDuration {
			t.Errorf("operation took %v, expected less than %v", elapsed, maxDuration)
		}
	})
}

// TestPerformance_CustomMessageImpact tests the performance impact of custom messages
// TestPerformance_CustomMessageImpact는 커스텀 메시지의 성능 영향을 테스트합니다
func TestPerformance_CustomMessageImpact(t *testing.T) {
	iterations := 10000

	t.Run("default messages", func(t *testing.T) {
		start := time.Now()
		for i := 0; i < iterations; i++ {
			v := New("", "field")
			v.Required()
			_ = v.Validate()
		}
		elapsed := time.Since(start)
		t.Logf("Default messages: %d operations in %v", iterations, elapsed)
	})

	t.Run("custom messages", func(t *testing.T) {
		start := time.Now()
		for i := 0; i < iterations; i++ {
			v := New("", "field")
			v.WithCustomMessage("required", "Custom required message")
			v.Required()
			_ = v.Validate()
		}
		elapsed := time.Since(start)
		t.Logf("Custom messages: %d operations in %v", iterations, elapsed)

		// Custom messages should have minimal overhead
		// 커스텀 메시지는 최소한의 오버헤드를 가져야 함
		maxDuration := 150 * time.Millisecond
		if elapsed > maxDuration {
			t.Errorf("operation took %v, expected less than %v", elapsed, maxDuration)
		}
	})
}

// TestPerformance_GarbageCollection tests GC pressure during validation
// TestPerformance_GarbageCollection는 검증 중 GC 압력을 테스트합니다
func TestPerformance_GarbageCollection(t *testing.T) {
	var m runtime.MemStats

	// Get initial GC stats / 초기 GC 통계 가져오기
	runtime.GC()
	runtime.ReadMemStats(&m)
	startGC := m.NumGC

	// Perform many validation operations / 많은 검증 작업 수행
	iterations := 10000
	start := time.Now()
	for i := 0; i < iterations; i++ {
		v := New("test@example.com", "email")
		v.Required().MinLength(5).MaxLength(100).Email()
		_ = v.Validate()
	}
	elapsed := time.Since(start)

	// Get final GC stats / 최종 GC 통계 가져오기
	runtime.ReadMemStats(&m)
	endGC := m.NumGC
	gcRuns := endGC - startGC

	t.Logf("Performed %d validations in %v", iterations, elapsed)
	t.Logf("GC runs: %d", gcRuns)
	t.Logf("Total allocated: %.2f MB", float64(m.TotalAlloc)/1024/1024)
	t.Logf("Sys: %.2f MB", float64(m.Sys)/1024/1024)

	// GC should not run excessively / GC가 과도하게 실행되지 않아야 함
	// Maximum 30 GC runs for 10000 operations (allowing for system GC)
	// 10000개 작업에 최대 30회 GC 실행 (시스템 GC 허용)
	maxGCRuns := uint32(30)
	if gcRuns > maxGCRuns {
		t.Errorf("GC ran %d times, expected less than %d times", gcRuns, maxGCRuns)
	}
}
