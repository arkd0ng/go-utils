package validation

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// Stress tests for validation package
// validation 패키지의 스트레스 테스트

// TestStress_HighVolume tests validation with extremely high volume
// TestStress_HighVolume는 매우 많은 양으로 검증을 테스트합니다
func TestStress_HighVolume(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping stress test in short mode")
	}

	// Stress test with 1 million validations
	// 100만 검증으로 스트레스 테스트
	numValidations := 1000000
	var successCount atomic.Int64
	var errorCount atomic.Int64

	t.Logf("Starting stress test with %d validations", numValidations)

	start := time.Now()
	var wg sync.WaitGroup
	numWorkers := 100
	validationsPerWorker := numValidations / numWorkers

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			for j := 0; j < validationsPerWorker; j++ {
				v := New(fmt.Sprintf("user%d-%d@example.com", workerID, j), "email")
				v.Required().Email()
				err := v.Validate()

				if err != nil {
					errorCount.Add(1)
				} else {
					successCount.Add(1)
				}
			}
		}(i)
	}

	wg.Wait()
	elapsed := time.Since(start)

	successes := successCount.Load()
	errors := errorCount.Load()
	total := successes + errors

	t.Logf("Completed %d validations in %v", total, elapsed)
	t.Logf("Successes: %d (%.2f%%)", successes, float64(successes)/float64(total)*100)
	t.Logf("Errors: %d (%.2f%%)", errors, float64(errors)/float64(total)*100)
	t.Logf("Throughput: %.2f ops/sec", float64(total)/elapsed.Seconds())

	// All validations should succeed / 모든 검증이 성공해야 함
	if errors > 0 {
		t.Errorf("encountered %d errors during stress test", errors)
	}

	// Minimum success rate / 최소 성공률
	minSuccessRate := 99.9
	successRate := float64(successes) / float64(total) * 100
	if successRate < minSuccessRate {
		t.Errorf("success rate = %.2f%%, want >= %.2f%%", successRate, minSuccessRate)
	}
}

// TestStress_LargeDataValues tests validation with extremely large data values
// TestStress_LargeDataValues는 매우 큰 데이터 값으로 검증을 테스트합니다
func TestStress_LargeDataValues(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping stress test in short mode")
	}

	tests := []struct {
		name string
		size int
	}{
		{"10KB string", 10 * 1024},
		{"100KB string", 100 * 1024},
		{"1MB string", 1024 * 1024},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create large string / 큰 문자열 생성
			data := make([]byte, tt.size)
			for i := 0; i < tt.size; i++ {
				data[i] = 'a'
			}
			str := string(data)

			start := time.Now()
			v := New(str, "large_field")
			v.Required().MinLength(1).MaxLength(2 * 1024 * 1024) // Max 2MB
			err := v.Validate()
			elapsed := time.Since(start)

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			t.Logf("Validated %s in %v", tt.name, elapsed)

			// Should complete within reasonable time
			// 합리적인 시간 내에 완료되어야 함
			maxDuration := 100 * time.Millisecond
			if elapsed > maxDuration {
				t.Errorf("operation took %v, expected less than %v", elapsed, maxDuration)
			}
		})
	}
}

// TestStress_DeepNestedValidation tests deeply nested validation structures
// TestStress_DeepNestedValidation는 깊이 중첩된 검증 구조를 테스트합니다
func TestStress_DeepNestedValidation(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping stress test in short mode")
	}

	// Create deeply nested structure / 깊이 중첩된 구조 생성
	depth := 1000
	mv := NewValidator()

	for i := 0; i < depth; i++ {
		mv.Field(fmt.Sprintf("value%d", i), fmt.Sprintf("field%d", i)).Required()
	}

	start := time.Now()
	err := mv.Validate()
	elapsed := time.Since(start)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	t.Logf("Validated %d nested fields in %v", depth, elapsed)

	// Should complete within reasonable time
	// 합리적인 시간 내에 완료되어야 함
	maxDuration := 100 * time.Millisecond
	if elapsed > maxDuration {
		t.Errorf("operation took %v, expected less than %v", elapsed, maxDuration)
	}
}

// TestStress_MemoryPressure tests validation under memory pressure
// TestStress_MemoryPressure는 메모리 압박 하에서 검증을 테스트합니다
func TestStress_MemoryPressure(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping stress test in short mode")
	}

	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	startMem := m.Alloc

	// Allocate large amounts of data / 대량의 데이터 할당
	numValidations := 100000
	validations := make([]*Validator, numValidations)

	t.Logf("Creating %d validators under memory pressure", numValidations)

	start := time.Now()
	for i := 0; i < numValidations; i++ {
		v := New(fmt.Sprintf("user%d@example.com", i), "email")
		v.Required().Email().MinLength(5).MaxLength(100)
		validations[i] = v
	}
	elapsed := time.Since(start)

	runtime.ReadMemStats(&m)
	endMem := m.Alloc
	memUsed := endMem - startMem

	t.Logf("Created %d validators in %v", numValidations, elapsed)
	t.Logf("Memory used: %.2f MB", float64(memUsed)/1024/1024)
	t.Logf("GC runs: %d", m.NumGC)

	// Validate all / 모두 검증
	validationStart := time.Now()
	for i := 0; i < numValidations; i++ {
		_ = validations[i].Validate()

		// Force GC periodically / 주기적으로 GC 강제 실행
		if i%10000 == 0 {
			runtime.GC()
		}
	}
	validationElapsed := time.Since(validationStart)

	t.Logf("Validated %d validators in %v", numValidations, validationElapsed)

	// System should remain stable / 시스템이 안정적으로 유지되어야 함
	// No crashes or panics expected / 충돌이나 패닉 없어야 함
}

// TestStress_RapidCreationDestruction tests rapid creation and destruction of validators
// TestStress_RapidCreationDestruction는 검증기의 빠른 생성과 파괴를 테스트합니다
func TestStress_RapidCreationDestruction(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping stress test in short mode")
	}

	iterations := 100000
	var opsCompleted atomic.Int64

	t.Logf("Starting rapid creation/destruction test with %d iterations", iterations)

	start := time.Now()
	var wg sync.WaitGroup
	numWorkers := 50

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			for j := 0; j < iterations/numWorkers; j++ {
				// Create, use, and discard validator rapidly
				// 검증기를 빠르게 생성, 사용 및 폐기
				v := New("test@example.com", "email")
				v.Required().Email()
				_ = v.Validate()
				// Validator goes out of scope / 검증기가 범위를 벗어남
				opsCompleted.Add(1)
			}
		}(i)
	}

	wg.Wait()
	elapsed := time.Since(start)

	completed := opsCompleted.Load()
	opsPerSec := float64(completed) / elapsed.Seconds()

	t.Logf("Completed %d create/validate/destroy cycles in %v", completed, elapsed)
	t.Logf("Throughput: %.2f ops/sec", opsPerSec)

	// Check for memory leaks / 메모리 누수 확인
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)

	t.Logf("Final heap allocation: %.2f MB", float64(m.Alloc)/1024/1024)
	t.Logf("GC runs: %d", m.NumGC)

	// Minimum throughput / 최소 처리량
	minOpsPerSec := 10000.0
	if opsPerSec < minOpsPerSec {
		t.Errorf("throughput = %.2f ops/sec, want >= %.2f ops/sec",
			opsPerSec, minOpsPerSec)
	}
}

// TestStress_ConcurrentErrors tests error handling under stress
// TestStress_ConcurrentErrors는 스트레스 하에서 에러 처리를 테스트합니다
func TestStress_ConcurrentErrors(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping stress test in short mode")
	}

	numGoroutines := 100
	numOperations := 1000

	var wg sync.WaitGroup
	var errorCount atomic.Int64
	var successCount atomic.Int64

	t.Logf("Starting concurrent error test with %d goroutines x %d operations",
		numGoroutines, numOperations)

	start := time.Now()
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			for j := 0; j < numOperations; j++ {
				// Create intentional validation failures
				// 의도적인 검증 실패 생성
				v := New("", "field")
				v.Required().MinLength(10).Email().URL() // Multiple errors
				err := v.Validate()

				if err != nil {
					errorCount.Add(1)
					// Verify error structure / 에러 구조 검증
					if verrs, ok := err.(ValidationErrors); ok {
						if len(verrs) == 0 {
							t.Error("ValidationErrors has no errors")
						}
					}
				} else {
					successCount.Add(1)
				}
			}
		}(i)
	}

	wg.Wait()
	elapsed := time.Since(start)

	errors := errorCount.Load()
	successes := successCount.Load()
	total := errors + successes

	t.Logf("Completed %d operations in %v", total, elapsed)
	t.Logf("Errors: %d (%.2f%%)", errors, float64(errors)/float64(total)*100)
	t.Logf("Successes: %d (%.2f%%)", successes, float64(successes)/float64(total)*100)

	// All operations should generate errors (empty string with required, minlength, email, url)
	// 모든 작업이 에러를 생성해야 함 (빈 문자열 + required, minlength, email, url)
	if errors != total {
		t.Errorf("expected %d errors, got %d", total, errors)
	}
}

// TestStress_ExtremeConcurrency tests with extreme number of goroutines
// TestStress_ExtremeConcurrency는 극단적인 수의 고루틴으로 테스트합니다
func TestStress_ExtremeConcurrency(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping stress test in short mode")
	}

	// Extreme concurrency: 1000 goroutines
	// 극단적 동시성: 1000개 고루틴
	numGoroutines := 1000
	numOperations := 100

	var wg sync.WaitGroup
	var successCount atomic.Int64
	errors := make(chan error, numGoroutines)

	t.Logf("Starting extreme concurrency test with %d goroutines", numGoroutines)

	start := time.Now()
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			for j := 0; j < numOperations; j++ {
				v := New(fmt.Sprintf("user%d-%d@example.com", workerID, j), "email")
				v.Required().Email()
				err := v.Validate()

				if err != nil {
					select {
					case errors <- err:
					default:
						// Channel full, ignore / 채널 가득 참, 무시
					}
					return
				}
				successCount.Add(1)
			}
		}(i)
	}

	wg.Wait()
	close(errors)
	elapsed := time.Since(start)

	// Count errors / 에러 개수
	var errs []error
	for err := range errors {
		errs = append(errs, err)
	}

	successes := successCount.Load()
	expected := int64(numGoroutines * numOperations)

	t.Logf("Completed in %v", elapsed)
	t.Logf("Successes: %d/%d", successes, expected)
	t.Logf("Errors: %d", len(errs))

	// Minimum success rate / 최소 성공률
	minSuccessRate := 95.0
	successRate := float64(successes) / float64(expected) * 100
	if successRate < minSuccessRate {
		t.Errorf("success rate = %.2f%%, want >= %.2f%%", successRate, minSuccessRate)
	}
}

// TestStress_LongRunning tests validation under sustained stress
// TestStress_LongRunning는 지속적인 스트레스 하에서 검증을 테스트합니다
func TestStress_LongRunning(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping stress test in short mode")
	}

	// Run for 10 seconds continuously
	// 10초 동안 지속적으로 실행
	duration := 10 * time.Second
	var opsCompleted atomic.Int64

	t.Logf("Starting long-running stress test for %v", duration)

	stop := make(chan struct{})
	var wg sync.WaitGroup
	numWorkers := 20

	// Launch workers / 워커 실행
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			counter := 0
			for {
				select {
				case <-stop:
					return
				default:
					v := New(fmt.Sprintf("user%d-%d@test.com", workerID, counter), "email")
					v.Required().Email().MinLength(5).MaxLength(100)
					_ = v.Validate()
					opsCompleted.Add(1)
					counter++
				}
			}
		}(i)
	}

	// Run for specified duration / 지정된 시간 동안 실행
	time.Sleep(duration)
	close(stop)
	wg.Wait()

	completed := opsCompleted.Load()
	opsPerSec := float64(completed) / duration.Seconds()

	t.Logf("Completed %d operations in %v", completed, duration)
	t.Logf("Average throughput: %.2f ops/sec", opsPerSec)

	// Check system stability / 시스템 안정성 확인
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	t.Logf("Final heap allocation: %.2f MB", float64(m.Alloc)/1024/1024)
	t.Logf("GC runs during test: %d", m.NumGC)

	// Minimum sustained throughput / 최소 지속 처리량
	minOpsPerSec := 5000.0
	if opsPerSec < minOpsPerSec {
		t.Errorf("throughput = %.2f ops/sec, want >= %.2f ops/sec",
			opsPerSec, minOpsPerSec)
	}
}

// TestStress_ResourceExhaustion tests behavior when resources are constrained
// TestStress_ResourceExhaustion는 리소스가 제한될 때의 동작을 테스트합니다
func TestStress_ResourceExhaustion(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping stress test in short mode")
	}

	// Try to exhaust system resources / 시스템 리소스 고갈 시도
	numGoroutines := 500
	numOperations := 200

	var wg sync.WaitGroup
	var successCount atomic.Int64
	var failCount atomic.Int64

	t.Logf("Starting resource exhaustion test")

	start := time.Now()
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			for j := 0; j < numOperations; j++ {
				// Create validators with complex validation chains
				// 복잡한 검증 체인으로 검증기 생성
				mv := NewValidator()
				mv.Field("test@example.com", "email").Required().Email().MinLength(5).MaxLength(100)
				mv.Field("John Doe", "name").Required().MinLength(2).MaxLength(50)
				age := 25
				mv.Field(age, "age").Min(18).Max(100)
				mv.Field([]int{1, 2, 3}, "numbers").ArrayMinLength(1).ArrayMaxLength(10)

				err := mv.Validate()
				if err != nil {
					failCount.Add(1)
				} else {
					successCount.Add(1)
				}
			}
		}(i)
	}

	wg.Wait()
	elapsed := time.Since(start)

	successes := successCount.Load()
	failures := failCount.Load()
	total := successes + failures

	t.Logf("Completed %d operations in %v", total, elapsed)
	t.Logf("Successes: %d (%.2f%%)", successes, float64(successes)/float64(total)*100)
	t.Logf("Failures: %d (%.2f%%)", failures, float64(failures)/float64(total)*100)

	// System should remain stable with high success rate
	// 시스템이 높은 성공률로 안정적으로 유지되어야 함
	minSuccessRate := 99.0
	successRate := float64(successes) / float64(total) * 100
	if successRate < minSuccessRate {
		t.Errorf("success rate = %.2f%%, want >= %.2f%%", successRate, minSuccessRate)
	}
}
