package validation

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// Load tests for validation package
// validation 패키지의 부하 테스트

// TestLoad_ConcurrentValidations tests validation under concurrent load
// TestLoad_ConcurrentValidations는 동시 부하 하에서 검증을 테스트합니다
func TestLoad_ConcurrentValidations(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping load test in short mode")
	}

	numGoroutines := 100
	numOperations := 1000

	var wg sync.WaitGroup
	errors := make(chan error, numGoroutines)
	var successCount atomic.Int64

	// Launch concurrent goroutines / 동시 고루틴 실행
	start := time.Now()
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			for j := 0; j < numOperations; j++ {
				v := New(fmt.Sprintf("user%d-%d@example.com", workerID, j), "email")
				v.Required().Email().MinLength(5).MaxLength(100)
				err := v.Validate()

				if err != nil {
					errors <- fmt.Errorf("worker %d, op %d: %w", workerID, j, err)
					return
				}
				successCount.Add(1)
			}
		}(i)
	}

	// Wait for completion / 완료 대기
	wg.Wait()
	close(errors)
	elapsed := time.Since(start)

	// Check for errors / 에러 확인
	var errs []error
	for err := range errors {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		t.Errorf("encountered %d errors during concurrent operations", len(errs))
		for i, err := range errs {
			if i < 10 { // Log first 10 errors / 처음 10개 에러만 로그
				t.Logf("  %v", err)
			}
		}
	}

	// Verify total operations / 총 작업 수 검증
	expected := int64(numGoroutines * numOperations)
	actual := successCount.Load()
	if actual != expected {
		t.Errorf("success count = %d, want %d", actual, expected)
	}

	// Calculate throughput / 처리량 계산
	opsPerSec := float64(actual) / elapsed.Seconds()
	t.Logf("Completed %d operations in %v (%.2f ops/sec)",
		actual, elapsed, opsPerSec)

	// Minimum throughput requirement / 최소 처리량 요구사항
	// At least 10,000 operations per second
	// 초당 최소 10,000개 작업
	minOpsPerSec := 10000.0
	if opsPerSec < minOpsPerSec {
		t.Errorf("throughput = %.2f ops/sec, want >= %.2f ops/sec",
			opsPerSec, minOpsPerSec)
	}
}

// TestLoad_ConcurrentMultiValidator tests MultiValidator under concurrent load
// TestLoad_ConcurrentMultiValidator는 동시 부하 하에서 MultiValidator를 테스트합니다
func TestLoad_ConcurrentMultiValidator(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping load test in short mode")
	}

	numGoroutines := 50
	numOperations := 500

	var wg sync.WaitGroup
	errors := make(chan error, numGoroutines)
	var successCount atomic.Int64

	// Launch concurrent goroutines / 동시 고루틴 실행
	start := time.Now()
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			for j := 0; j < numOperations; j++ {
				age := 25 + j%50
				mv := NewValidator()
				mv.Field(fmt.Sprintf("user%d-%d@example.com", workerID, j), "email").Required().Email()
				mv.Field(fmt.Sprintf("User%d-%d", workerID, j), "name").Required().MinLength(5)
				mv.Field(age, "age").Min(18).Max(100)
				err := mv.Validate()

				if err != nil {
					errors <- fmt.Errorf("worker %d, op %d: %w", workerID, j, err)
					return
				}
				successCount.Add(1)
			}
		}(i)
	}

	// Wait for completion / 완료 대기
	wg.Wait()
	close(errors)
	elapsed := time.Since(start)

	// Check for errors / 에러 확인
	var errs []error
	for err := range errors {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		t.Errorf("encountered %d errors during concurrent operations", len(errs))
		for i, err := range errs {
			if i < 10 { // Log first 10 errors / 처음 10개 에러만 로그
				t.Logf("  %v", err)
			}
		}
	}

	// Verify total operations / 총 작업 수 검증
	expected := int64(numGoroutines * numOperations)
	actual := successCount.Load()
	if actual != expected {
		t.Errorf("success count = %d, want %d", actual, expected)
	}

	// Calculate throughput / 처리량 계산
	opsPerSec := float64(actual) / elapsed.Seconds()
	t.Logf("Completed %d multi-field validations in %v (%.2f ops/sec)",
		actual, elapsed, opsPerSec)

	// Minimum throughput requirement / 최소 처리량 요구사항
	// At least 5,000 operations per second
	// 초당 최소 5,000개 작업
	minOpsPerSec := 5000.0
	if opsPerSec < minOpsPerSec {
		t.Errorf("throughput = %.2f ops/sec, want >= %.2f ops/sec",
			opsPerSec, minOpsPerSec)
	}
}

// TestLoad_Throughput tests sustained throughput over time
// TestLoad_Throughput는 시간 경과에 따른 지속적인 처리량을 테스트합니다
func TestLoad_Throughput(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping load test in short mode")
	}

	duration := 5 * time.Second
	var opsCompleted atomic.Int64

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	numWorkers := 10
	var wg sync.WaitGroup

	// Launch worker goroutines / 워커 고루틴 실행
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			counter := 0
			for {
				select {
				case <-ctx.Done():
					return
				default:
					v := New(fmt.Sprintf("user%d-%d@test.com", workerID, counter), "email")
					v.Required().Email()
					_ = v.Validate()
					opsCompleted.Add(1)
					counter++
				}
			}
		}(i)
	}

	// Wait for completion / 완료 대기
	wg.Wait()

	total := opsCompleted.Load()
	opsPerSecond := float64(total) / duration.Seconds()

	t.Logf("Completed %d operations in %v", total, duration)
	t.Logf("Sustained throughput: %.2f ops/sec", opsPerSecond)

	// Minimum sustained throughput / 최소 지속 처리량
	// At least 10,000 operations per second sustained
	// 초당 최소 10,000개 작업 지속
	minOpsPerSecond := 10000.0
	if opsPerSecond < minOpsPerSecond {
		t.Errorf("throughput = %.2f ops/sec, want >= %.2f ops/sec",
			opsPerSecond, minOpsPerSecond)
	}
}

// TestLoad_RaceConditions tests for data race conditions
// TestLoad_RaceConditions는 데이터 경합 조건을 테스트합니다
func TestLoad_RaceConditions(t *testing.T) {
	// Run with: go test -race
	// 실행: go test -race

	numGoroutines := 100
	iterations := 100

	var wg sync.WaitGroup

	// Shared test data / 공유 테스트 데이터
	testData := []string{
		"test@example.com",
		"user@test.com",
		"admin@domain.com",
	}

	// Launch concurrent goroutines / 동시 고루틴 실행
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for j := 0; j < iterations; j++ {
				email := testData[j%len(testData)]
				v := New(email, "email")
				v.Required().Email()
				_ = v.Validate()
			}
		}(i)
	}

	wg.Wait()

	// If there are race conditions, go test -race will detect them
	// 경합 조건이 있으면 go test -race가 감지함
}

// TestLoad_ErrorHandling tests error handling under concurrent load
// TestLoad_ErrorHandling는 동시 부하 하에서 에러 처리를 테스트합니다
func TestLoad_ErrorHandling(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping load test in short mode")
	}

	numGoroutines := 50
	numOperations := 100

	var wg sync.WaitGroup
	var errorCount atomic.Int64
	var successCount atomic.Int64

	// Launch concurrent goroutines / 동시 고루틴 실행
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			for j := 0; j < numOperations; j++ {
				// Alternate between valid and invalid data
				// 유효한 데이터와 무효한 데이터를 번갈아 사용
				var email string
				if j%2 == 0 {
					email = fmt.Sprintf("valid%d-%d@example.com", workerID, j)
				} else {
					email = "invalid-email" // Invalid / 무효함
				}

				v := New(email, "email")
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

	// Wait for completion / 완료 대기
	wg.Wait()

	total := int64(numGoroutines * numOperations)
	errors := errorCount.Load()
	successes := successCount.Load()

	t.Logf("Total operations: %d", total)
	t.Logf("Successes: %d", successes)
	t.Logf("Errors: %d", errors)

	// Verify counts / 개수 검증
	if errors+successes != total {
		t.Errorf("count mismatch: errors(%d) + successes(%d) != total(%d)",
			errors, successes, total)
	}

	// Should have approximately 50% errors (every other validation fails)
	// 대략 50% 에러가 있어야 함 (두 번째 검증마다 실패)
	expectedErrors := int64(float64(total) * 0.5)
	tolerance := int64(float64(total) * 0.1) // 10% tolerance / 10% 허용오차
	if errors < expectedErrors-tolerance || errors > expectedErrors+tolerance {
		t.Errorf("error count = %d, expected approximately %d (±%d)",
			errors, expectedErrors, tolerance)
	}
}

// TestLoad_MemoryLeaks tests for memory leaks under sustained load
// TestLoad_MemoryLeaks는 지속적인 부하 하에서 메모리 누수를 테스트합니다
func TestLoad_MemoryLeaks(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping load test in short mode")
	}

	// Baseline memory / 기준 메모리
	var m1 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)
	baseline := m1.Alloc

	// Perform many operations / 많은 작업 수행
	iterations := 50000
	for i := 0; i < iterations; i++ {
		v := New("test@example.com", "email")
		v.Required().Email().MinLength(5).MaxLength(100)
		_ = v.Validate()

		// Force GC periodically / 주기적으로 GC 강제 실행
		if i%10000 == 0 {
			runtime.GC()
		}
	}

	// Final memory / 최종 메모리
	var m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m2)
	final := m2.Alloc

	// Calculate growth (handle underflow)
	// 증가량 계산 (언더플로우 처리)
	var growthMB float64
	if final > baseline {
		growth := final - baseline
		growthMB = float64(growth) / 1024 / 1024
	} else {
		growthMB = 0 // GC cleaned up, no growth
	}

	t.Logf("Memory baseline: %.2f MB", float64(baseline)/1024/1024)
	t.Logf("Memory final: %.2f MB", float64(final)/1024/1024)
	t.Logf("Memory growth: %.2f MB", growthMB)

	// Memory growth should be minimal (less than 10MB)
	// 메모리 증가는 최소화되어야 함 (10MB 미만)
	maxGrowthMB := 10.0
	if growthMB > maxGrowthMB {
		t.Errorf("memory growth = %.2f MB, expected less than %.2f MB",
			growthMB, maxGrowthMB)
	}
}

// TestLoad_MixedValidationTypes tests concurrent validation with mixed types
// TestLoad_MixedValidationTypes는 혼합 유형을 사용한 동시 검증을 테스트합니다
func TestLoad_MixedValidationTypes(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping load test in short mode")
	}

	numGoroutines := 20
	numOperations := 500

	var wg sync.WaitGroup
	var successCount atomic.Int64

	// Launch concurrent goroutines with different validation types
	// 다양한 검증 유형으로 동시 고루틴 실행
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			for j := 0; j < numOperations; j++ {
				// Rotate through different validation types
				// 다양한 검증 유형을 순환
				switch j % 5 {
				case 0: // Email validation / 이메일 검증
					v := New(fmt.Sprintf("user%d@example.com", j), "email")
					v.Email()
					_ = v.Validate()
				case 1: // Numeric validation / 숫자 검증
					v := New(j%100, "number")
					v.Min(0).Max(100)
					_ = v.Validate()
				case 2: // String validation / 문자열 검증
					v := New(fmt.Sprintf("string%d", j), "text")
					v.MinLength(5).MaxLength(20)
					_ = v.Validate()
				case 3: // Array validation / 배열 검증
					arr := []int{1, 2, 3, 4, 5}
					v := New(arr, "array")
					v.ArrayMinLength(1).ArrayMaxLength(10)
					_ = v.Validate()
				case 4: // Complex validation / 복합 검증
					mv := NewValidator()
					mv.Field("test@example.com", "email").Email()
					mv.Field(25, "age").Min(18)
					_ = mv.Validate()
				}
				successCount.Add(1)
			}
		}(i)
	}

	// Wait for completion / 완료 대기
	wg.Wait()

	expected := int64(numGoroutines * numOperations)
	actual := successCount.Load()
	if actual != expected {
		t.Errorf("success count = %d, want %d", actual, expected)
	}

	t.Logf("Successfully completed %d mixed-type validations", actual)
}
