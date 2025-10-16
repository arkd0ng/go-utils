# Go-Utils Test Code Guide / Go-Utils 테스트 코드 가이드

**Version / 버전**: v1.11.x
**Last Updated / 최종 업데이트**: 2025-10-16

---

## Table of Contents / 목차

1. [Introduction / 소개](#introduction--소개)
2. [Testing Philosophy / 테스트 철학](#testing-philosophy--테스트-철학)
3. [Test Structure / 테스트 구조](#test-structure--테스트-구조)
4. [Unit Testing / 단위 테스트](#unit-testing--단위-테스트)
5. [Edge Cases and Error Paths / 엣지 케이스 및 에러 경로](#edge-cases-and-error-paths--엣지-케이스-및-에러-경로)
6. [Benchmark Testing / 벤치마크 테스트](#benchmark-testing--벤치마크-테스트)
7. [Integration Testing / 통합 테스트](#integration-testing--통합-테스트)
8. [Performance Testing / 성능 테스트](#performance-testing--성능-테스트)
9. [Load Testing / 부하 테스트](#load-testing--부하-테스트)
10. [Stress Testing / 스트레스 테스트](#stress-testing--스트레스-테스트)
11. [Security Testing / 보안 테스트](#security-testing--보안-테스트)
12. [Fuzz Testing / 퍼즈 테스트](#fuzz-testing--퍼즈-테스트)
13. [Property-Based Testing / 속성 기반 테스트](#property-based-testing--속성-기반-테스트)
14. [Test Coverage / 테스트 커버리지](#test-coverage--테스트-커버리지)
15. [Testing Best Practices / 테스트 모범 사례](#testing-best-practices--테스트-모범-사례)
16. [Continuous Integration / 지속적 통합](#continuous-integration--지속적-통합)

---

## Introduction / 소개

This guide provides comprehensive testing standards for the go-utils project.
이 가이드는 go-utils 프로젝트의 포괄적인 테스트 표준을 제공합니다.

**Testing Goals / 테스트 목표**:
- Achieve 100% test coverage for all packages / 모든 패키지에 대해 100% 테스트 커버리지 달성
- Ensure code reliability and correctness / 코드 신뢰성 및 정확성 보장
- Prevent regressions / 회귀 방지
- Document expected behavior through tests / 테스트를 통한 예상 동작 문서화
- Improve code quality and maintainability / 코드 품질 및 유지보수성 향상

---

## Testing Philosophy / 테스트 철학

### Core Principles / 핵심 원칙

1. **Comprehensive Testing / 포괄적 테스트**
   - Test all public functions and methods / 모든 공개 함수 및 메서드 테스트
   - Cover all code paths including errors / 에러를 포함한 모든 코드 경로 커버
   - Test edge cases and boundary conditions / 엣지 케이스 및 경계 조건 테스트

2. **Test Isolation / 테스트 격리**
   - Each test should be independent / 각 테스트는 독립적이어야 함
   - No shared state between tests / 테스트 간 공유 상태 없음
   - Use setup and teardown appropriately / 적절한 setup 및 teardown 사용

3. **Readability / 가독성**
   - Clear test names describing what is tested / 무엇을 테스트하는지 명확한 테스트 이름
   - Well-organized test cases / 잘 정리된 테스트 케이스
   - Bilingual comments (English/Korean) / 이중 언어 주석 (영문/한글)

4. **Maintainability / 유지보수성**
   - DRY (Don't Repeat Yourself) principle / DRY 원칙
   - Helper functions for common operations / 공통 작업을 위한 헬퍼 함수
   - Table-driven tests for multiple scenarios / 여러 시나리오를 위한 테이블 기반 테스트

---

## Test Structure / 테스트 구조

### File Organization / 파일 구성

```
package/
├── function.go          # Implementation / 구현
├── function_test.go     # Unit tests / 단위 테스트
├── benchmark_test.go    # Benchmarks / 벤치마크 (optional / 선택사항)
├── integration_test.go  # Integration tests / 통합 테스트 (if needed / 필요시)
└── testhelper_test.go   # Test helpers / 테스트 헬퍼
```

### Naming Conventions / 명명 규칙

**Test Files / 테스트 파일**:
- Unit tests: `*_test.go`
- Integration tests: `*_integration_test.go`
- Benchmark tests: Can be in regular `*_test.go` or separate `*_benchmark_test.go`

**Test Functions / 테스트 함수**:
```go
func TestFunctionName(t *testing.T)           // Unit test / 단위 테스트
func BenchmarkFunctionName(b *testing.B)      // Benchmark / 벤치마크
func ExampleFunctionName()                    // Example / 예제
func FuzzFunctionName(f *testing.F)           // Fuzz test / 퍼즈 테스트
```

### Test Function Template / 테스트 함수 템플릿

```go
// TestFunctionName tests the FunctionName function
// TestFunctionName은 FunctionName 함수를 테스트합니다
func TestFunctionName(t *testing.T) {
	// Table-driven test structure / 테이블 기반 테스트 구조
	tests := []struct {
		name     string        // Test case name / 테스트 케이스 이름
		input    InputType     // Input parameters / 입력 매개변수
		expected OutputType    // Expected result / 예상 결과
		wantErr  bool          // Expect error? / 에러 예상?
	}{
		{
			name:     "valid input",
			input:    validInput,
			expected: expectedOutput,
			wantErr:  false,
		},
		{
			name:     "invalid input",
			input:    invalidInput,
			expected: zeroValue,
			wantErr:  true,
		},
		// Add more test cases / 더 많은 테스트 케이스 추가
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Execute function / 함수 실행
			result, err := FunctionName(tt.input)

			// Check error / 에러 확인
			if (err != nil) != tt.wantErr {
				t.Errorf("FunctionName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Check result / 결과 확인
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("FunctionName() = %v, want %v", result, tt.expected)
			}
		})
	}
}
```

---

## Unit Testing / 단위 테스트

Unit tests verify individual functions or methods in isolation.
단위 테스트는 개별 함수 또는 메서드를 격리하여 검증합니다.

### Basic Unit Test Example / 기본 단위 테스트 예제

```go
// Example from sliceutil package
// sliceutil 패키지의 예제
func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		item     int
		expected bool
	}{
		{"found in middle", []int{1, 2, 3, 4, 5}, 3, true},
		{"found at start", []int{1, 2, 3, 4, 5}, 1, true},
		{"found at end", []int{1, 2, 3, 4, 5}, 5, true},
		{"not found", []int{1, 2, 3, 4, 5}, 10, false},
		{"empty slice", []int{}, 1, false},
		{"nil slice", nil, 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Contains(tt.slice, tt.item)
			if result != tt.expected {
				t.Errorf("Contains(%v, %d) = %v, want %v",
					tt.slice, tt.item, result, tt.expected)
			}
		})
	}
}
```

### Testing with Generics / 제네릭 테스트

```go
// Testing generic functions / 제네릭 함수 테스트
func TestMap(t *testing.T) {
	// Test with integers / 정수로 테스트
	t.Run("integers", func(t *testing.T) {
		input := []int{1, 2, 3}
		double := func(n int) int { return n * 2 }
		result := Map(input, double)
		expected := []int{2, 4, 6}

		if !Equal(result, expected) {
			t.Errorf("Map() = %v, want %v", result, expected)
		}
	})

	// Test with strings / 문자열로 테스트
	t.Run("strings", func(t *testing.T) {
		input := []string{"a", "b", "c"}
		upper := func(s string) string { return strings.ToUpper(s) }
		result := Map(input, upper)
		expected := []string{"A", "B", "C"}

		if !Equal(result, expected) {
			t.Errorf("Map() = %v, want %v", result, expected)
		}
	})
}
```

### Testing Error Handling / 에러 처리 테스트

```go
func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a, b    int
		want    int
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid division",
			a:       10,
			b:       2,
			want:    5,
			wantErr: false,
		},
		{
			name:    "division by zero",
			a:       10,
			b:       0,
			want:    0,
			wantErr: true,
			errMsg:  "division by zero",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)

			// Check error / 에러 확인
			if tt.wantErr {
				if err == nil {
					t.Error("expected error but got nil")
					return
				}
				if err.Error() != tt.errMsg {
					t.Errorf("error message = %q, want %q", err.Error(), tt.errMsg)
				}
				return
			}

			// Check no error / 에러 없음 확인
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			// Check result / 결과 확인
			if got != tt.want {
				t.Errorf("Divide(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
```

---

## Edge Cases and Error Paths / 엣지 케이스 및 에러 경로

Testing edge cases ensures robustness and prevents unexpected failures.
엣지 케이스 테스트는 견고성을 보장하고 예상치 못한 실패를 방지합니다.

### Common Edge Cases / 일반적인 엣지 케이스

1. **Empty Input / 빈 입력**
   ```go
   {"empty slice", []int{}, someFunc, expectedForEmpty},
   {"empty string", "", someFunc, expectedForEmpty},
   {"empty map", map[string]int{}, someFunc, expectedForEmpty},
   ```

2. **Nil Input / Nil 입력**
   ```go
   {"nil slice", nil, someFunc, expectedForNil},
   {"nil pointer", nil, someFunc, expectedForNil},
   ```

3. **Zero Values / 제로 값**
   ```go
   {"zero integer", 0, someFunc, expectedForZero},
   {"zero float", 0.0, someFunc, expectedForZero},
   {"zero time", time.Time{}, someFunc, expectedForZero},
   ```

4. **Boundary Values / 경계 값**
   ```go
   {"minimum value", math.MinInt64, someFunc, expected},
   {"maximum value", math.MaxInt64, someFunc, expected},
   {"just below limit", limit-1, someFunc, expected},
   {"at limit", limit, someFunc, expected},
   {"just above limit", limit+1, someFunc, expected},
   ```

5. **Single Element / 단일 요소**
   ```go
   {"single element", []int{1}, someFunc, expected},
   ```

6. **Large Input / 큰 입력**
   ```go
   {"large slice", make([]int, 10000), someFunc, expected},
   {"long string", strings.Repeat("a", 100000), someFunc, expected},
   ```

7. **Special Characters / 특수 문자**
   ```go
   {"unicode", "한글テスト🎉", someFunc, expected},
   {"special chars", "!@#$%^&*()", someFunc, expected},
   {"whitespace", "  \t\n  ", someFunc, expected},
   ```

8. **Negative Numbers / 음수**
   ```go
   {"negative", -10, someFunc, expected},
   ```

### Comprehensive Edge Case Example / 포괄적인 엣지 케이스 예제

```go
func TestTruncate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		maxLen   int
		expected string
		wantErr  bool
	}{
		// Normal cases / 일반 케이스
		{"normal truncate", "Hello World", 5, "Hello", false},
		{"no truncate needed", "Hi", 10, "Hi", false},
		{"exact length", "Hello", 5, "Hello", false},

		// Edge cases / 엣지 케이스
		{"empty string", "", 5, "", false},
		{"zero max length", "Hello", 0, "", false},
		{"negative max length", "Hello", -1, "", true},
		{"single character", "A", 1, "A", false},
		{"unicode emoji", "Hello🎉World", 6, "Hello🎉", false},
		{"korean text", "안녕하세요", 3, "안녕하", false},
		{"mixed unicode", "Hi안녕🎉", 4, "Hi안녕", false},
		{"only whitespace", "     ", 3, "   ", false},
		{"special chars", "!@#$%^&*()", 5, "!@#$%", false},
		{"very long string", strings.Repeat("a", 10000), 5, "aaaaa", false},

		// Boundary cases / 경계 케이스
		{"maxLen is 1", "Hello", 1, "H", false},
		{"maxLen equals length", "Test", 4, "Test", false},
		{"maxLen greater than length", "Test", 100, "Test", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Truncate(tt.input, tt.maxLen)

			if (err != nil) != tt.wantErr {
				t.Errorf("Truncate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if result != tt.expected {
				t.Errorf("Truncate(%q, %d) = %q, want %q",
					tt.input, tt.maxLen, result, tt.expected)
			}
		})
	}
}
```

### Error Path Testing / 에러 경로 테스트

```go
func TestValidateInput(t *testing.T) {
	tests := []struct {
		name    string
		input   Input
		wantErr error
	}{
		{
			name:    "valid input",
			input:   Input{Value: 10},
			wantErr: nil,
		},
		{
			name:    "negative value",
			input:   Input{Value: -1},
			wantErr: ErrNegativeValue,
		},
		{
			name:    "value too large",
			input:   Input{Value: 1000},
			wantErr: ErrValueTooLarge,
		},
		{
			name:    "missing required field",
			input:   Input{},
			wantErr: ErrMissingField,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateInput(tt.input)

			if tt.wantErr != nil {
				if err == nil {
					t.Error("expected error but got nil")
					return
				}
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("error = %v, want %v", err, tt.wantErr)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}
		})
	}
}
```

---

## Benchmark Testing / 벤치마크 테스트

Benchmark tests measure performance of functions.
벤치마크 테스트는 함수의 성능을 측정합니다.

### Basic Benchmark / 기본 벤치마크

```go
// BenchmarkAlnum benchmarks the Alnum method
// BenchmarkAlnum은 Alnum 메서드를 벤치마킹합니다
func BenchmarkAlnum(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = GenString.Alnum(32)
	}
}
```

### Parameterized Benchmarks / 매개변수화된 벤치마크

```go
func BenchmarkContains(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			slice := make([]int, size)
			for i := 0; i < size; i++ {
				slice[i] = i
			}
			searchItem := size / 2 // Middle element / 중간 요소

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = Contains(slice, searchItem)
			}
		})
	}
}
```

### Memory Benchmarks / 메모리 벤치마크

```go
func BenchmarkMap(b *testing.B) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = i
	}
	double := func(n int) int { return n * 2 }

	b.ResetTimer()
	b.ReportAllocs() // Report memory allocations / 메모리 할당 보고

	for i := 0; i < b.N; i++ {
		_ = Map(input, double)
	}
}
```

### Benchmark Comparison / 벤치마크 비교

```go
// Compare different implementations / 다른 구현 비교
func BenchmarkReverseV1(b *testing.B) {
	s := "Hello, World! 안녕하세요"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ReverseV1(s)
	}
}

func BenchmarkReverseV2(b *testing.B) {
	s := "Hello, World! 안녕하세요"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ReverseV2(s)
	}
}
```

### Running Benchmarks / 벤치마크 실행

```bash
# Run all benchmarks / 모든 벤치마크 실행
go test -bench=. ./...

# Run specific benchmark / 특정 벤치마크 실행
go test -bench=BenchmarkAlnum ./random

# With memory stats / 메모리 통계 포함
go test -bench=. -benchmem ./...

# Compare benchmarks / 벤치마크 비교
go test -bench=BenchmarkReverse -benchmem ./stringutil

# Save results / 결과 저장
go test -bench=. -benchmem ./... > bench.txt

# Compare with previous results using benchstat
# benchstat을 사용하여 이전 결과와 비교
go install golang.org/x/perf/cmd/benchstat@latest
benchstat old.txt new.txt
```

---

## Integration Testing / 통합 테스트

Integration tests verify that multiple components work together correctly.
통합 테스트는 여러 컴포넌트가 함께 올바르게 작동하는지 검증합니다.

### Database Integration Test / 데이터베이스 통합 테스트

```go
// +build integration

package mysql

import (
	"context"
	"testing"
	"time"
)

func TestMySQLIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	// Setup / 설정
	client := setupTestDatabase(t)
	defer client.Close()

	ctx := context.Background()

	t.Run("full CRUD lifecycle", func(t *testing.T) {
		// Create / 생성
		user := map[string]interface{}{
			"name":  "John Doe",
			"email": "john@example.com",
			"age":   30,
		}

		id, err := client.Insert(ctx, "users", user)
		if err != nil {
			t.Fatalf("Insert failed: %v", err)
		}

		// Read / 읽기
		var result map[string]interface{}
		err = client.SelectOne(ctx, &result, "users", "id = ?", id)
		if err != nil {
			t.Fatalf("SelectOne failed: %v", err)
		}

		if result["name"] != "John Doe" {
			t.Errorf("name = %v, want John Doe", result["name"])
		}

		// Update / 업데이트
		updates := map[string]interface{}{"age": 31}
		_, err = client.Update(ctx, "users", updates, "id = ?", id)
		if err != nil {
			t.Fatalf("Update failed: %v", err)
		}

		// Verify update / 업데이트 검증
		err = client.SelectOne(ctx, &result, "users", "id = ?", id)
		if err != nil {
			t.Fatalf("SelectOne after update failed: %v", err)
		}

		if int(result["age"].(int64)) != 31 {
			t.Errorf("age after update = %v, want 31", result["age"])
		}

		// Delete / 삭제
		_, err = client.Delete(ctx, "users", "id = ?", id)
		if err != nil {
			t.Fatalf("Delete failed: %v", err)
		}

		// Verify deletion / 삭제 검증
		exists, err := client.Exists(ctx, "users", "id = ?", id)
		if err != nil {
			t.Fatalf("Exists check failed: %v", err)
		}

		if exists {
			t.Error("record should not exist after deletion")
		}
	})
}

func setupTestDatabase(t *testing.T) *Client {
	client, err := New(
		WithDSN("root:password@tcp(localhost:3306)/testdb"),
		WithMaxOpenConns(10),
	)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	// Create test table / 테스트 테이블 생성
	_, err = client.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(100),
			email VARCHAR(100),
			age INT
		)
	`)
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}

	// Clean table / 테이블 정리
	_, err = client.Exec(context.Background(), "TRUNCATE TABLE users")
	if err != nil {
		t.Fatalf("failed to truncate table: %v", err)
	}

	return client
}
```

### Running Integration Tests / 통합 테스트 실행

```bash
# Run with integration tag / integration 태그로 실행
go test -tags=integration ./database/mysql

# Skip integration tests in short mode / short 모드에서 통합 테스트 건너뛰기
go test -short ./...

# Run only integration tests / 통합 테스트만 실행
go test -run Integration ./...
```

---

## Performance Testing / 성능 테스트

Performance tests measure execution time and resource usage.
성능 테스트는 실행 시간과 리소스 사용량을 측정합니다.

### Execution Time Test / 실행 시간 테스트

```go
func TestPerformance(t *testing.T) {
	t.Run("large dataset performance", func(t *testing.T) {
		size := 100000
		data := make([]int, size)
		for i := 0; i < size; i++ {
			data[i] = i
		}

		start := time.Now()
		result := Filter(data, func(n int) bool { return n%2 == 0 })
		elapsed := time.Since(start)

		t.Logf("Filtered %d elements in %v", len(result), elapsed)

		// Performance assertion / 성능 단언
		maxDuration := 100 * time.Millisecond
		if elapsed > maxDuration {
			t.Errorf("operation took %v, expected less than %v", elapsed, maxDuration)
		}
	})
}
```

### Memory Usage Test / 메모리 사용 테스트

```go
func TestMemoryUsage(t *testing.T) {
	var m runtime.MemStats

	runtime.GC()
	runtime.ReadMemStats(&m)
	before := m.Alloc

	// Perform operation / 작업 수행
	size := 10000
	result := make([][]int, size)
	for i := 0; i < size; i++ {
		result[i] = make([]int, 100)
	}

	runtime.ReadMemStats(&m)
	after := m.Alloc
	used := after - before

	t.Logf("Memory used: %d bytes (%.2f MB)", used, float64(used)/1024/1024)

	// Memory limit assertion / 메모리 제한 단언
	maxMemory := uint64(100 * 1024 * 1024) // 100 MB
	if used > maxMemory {
		t.Errorf("used %d bytes, expected less than %d bytes", used, maxMemory)
	}
}
```

---

## Load Testing / 부하 테스트

Load testing verifies system behavior under expected load.
부하 테스트는 예상 부하 하에서 시스템 동작을 검증합니다.

### Concurrent Operations Test / 동시 작업 테스트

```go
func TestConcurrentAccess(t *testing.T) {
	client := newTestClient(t)
	defer client.Close()

	numGoroutines := 100
	numOperations := 1000

	var wg sync.WaitGroup
	errors := make(chan error, numGoroutines)

	// Launch concurrent goroutines / 동시 고루틴 실행
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			for j := 0; j < numOperations; j++ {
				ctx := context.Background()
				user := map[string]interface{}{
					"name":  fmt.Sprintf("User%d-%d", workerID, j),
					"email": fmt.Sprintf("user%d-%d@test.com", workerID, j),
				}

				_, err := client.Insert(ctx, "users", user)
				if err != nil {
					errors <- fmt.Errorf("worker %d, op %d: %w", workerID, j, err)
					return
				}
			}
		}(i)
	}

	// Wait for completion / 완료 대기
	wg.Wait()
	close(errors)

	// Check for errors / 에러 확인
	var errs []error
	for err := range errors {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		t.Errorf("encountered %d errors during concurrent operations", len(errs))
		for _, err := range errs {
			t.Logf("  %v", err)
		}
	}

	// Verify total records / 총 레코드 수 검증
	count, err := client.Count(context.Background(), "users", "")
	if err != nil {
		t.Fatalf("Count failed: %v", err)
	}

	expected := int64(numGoroutines * numOperations)
	if count != expected {
		t.Errorf("count = %d, want %d", count, expected)
	}
}
```

### Throughput Test / 처리량 테스트

```go
func TestThroughput(t *testing.T) {
	client := newTestClient(t)
	defer client.Close()

	duration := 10 * time.Second
	var opsCompleted atomic.Int64

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	numWorkers := 10
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				default:
					_, err := client.SelectAll(ctx, "users")
					if err == nil {
						opsCompleted.Add(1)
					}
				}
			}
		}()
	}

	wg.Wait()

	total := opsCompleted.Load()
	opsPerSecond := float64(total) / duration.Seconds()

	t.Logf("Completed %d operations in %v", total, duration)
	t.Logf("Throughput: %.2f ops/sec", opsPerSecond)

	// Throughput requirement / 처리량 요구사항
	minOpsPerSecond := 100.0
	if opsPerSecond < minOpsPerSecond {
		t.Errorf("throughput = %.2f ops/sec, want >= %.2f ops/sec",
			opsPerSecond, minOpsPerSecond)
	}
}
```

---

## Stress Testing / 스트레스 테스트

Stress testing pushes the system beyond normal limits to find breaking points.
스트레스 테스트는 시스템을 정상 한계 이상으로 밀어 붙여 한계점을 찾습니다.

### Resource Exhaustion Test / 리소스 고갈 테스트

```go
func TestStressConnectionPool(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping stress test in short mode")
	}

	client, err := New(
		WithDSN(testDSN),
		WithMaxOpenConns(10), // Limited connections / 제한된 연결
	)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	// Try to exhaust connection pool / 연결 풀 고갈 시도
	numRequests := 1000
	timeout := 30 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var wg sync.WaitGroup
	errors := make(chan error, numRequests)

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			_, err := client.SelectAll(ctx, "users")
			if err != nil {
				errors <- fmt.Errorf("request %d: %w", id, err)
			}
		}(i)
	}

	wg.Wait()
	close(errors)

	// Collect errors / 에러 수집
	var failed int
	for err := range errors {
		failed++
		if failed <= 10 { // Log first 10 errors / 처음 10개 에러만 로그
			t.Logf("Error: %v", err)
		}
	}

	successRate := float64(numRequests-failed) / float64(numRequests) * 100
	t.Logf("Success rate: %.2f%% (%d/%d)", successRate, numRequests-failed, numRequests)

	// Minimum success rate / 최소 성공률
	minSuccessRate := 95.0
	if successRate < minSuccessRate {
		t.Errorf("success rate = %.2f%%, want >= %.2f%%", successRate, minSuccessRate)
	}
}
```

### Memory Stress Test / 메모리 스트레스 테스트

```go
func TestMemoryStress(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping stress test in short mode")
	}

	// Allocate large amounts of data / 대량의 데이터 할당
	numSlices := 10000
	sliceSize := 10000

	start := time.Now()
	var m runtime.MemStats

	runtime.GC()
	runtime.ReadMemStats(&m)
	startMem := m.Alloc

	slices := make([][]int, numSlices)
	for i := 0; i < numSlices; i++ {
		slices[i] = make([]int, sliceSize)
		for j := 0; j < sliceSize; j++ {
			slices[i][j] = j
		}

		// Process each slice / 각 슬라이스 처리
		_ = Map(slices[i], func(n int) int { return n * 2 })
	}

	runtime.ReadMemStats(&m)
	endMem := m.Alloc
	elapsed := time.Since(start)

	memUsed := endMem - startMem
	t.Logf("Processed %d slices of size %d in %v", numSlices, sliceSize, elapsed)
	t.Logf("Memory used: %.2f MB", float64(memUsed)/1024/1024)
	t.Logf("GC runs: %d", m.NumGC)

	// Verify no crashes or panics / 충돌이나 패닉 없음 검증
	if len(slices) != numSlices {
		t.Error("unexpected slice count after stress test")
	}
}
```

---

## Security Testing / 보안 테스트

Security testing identifies vulnerabilities and ensures secure code.
보안 테스트는 취약점을 식별하고 안전한 코드를 보장합니다.

### SQL Injection Prevention Test / SQL 인젝션 방지 테스트

```go
func TestSQLInjectionPrevention(t *testing.T) {
	client := newTestClient(t)
	defer client.Close()

	ctx := context.Background()

	// Insert test data / 테스트 데이터 삽입
	_, err := client.Insert(ctx, "users", map[string]interface{}{
		"name":  "John",
		"email": "john@test.com",
	})
	if err != nil {
		t.Fatal(err)
	}

	// Test SQL injection attempts / SQL 인젝션 시도 테스트
	injectionAttempts := []string{
		"' OR '1'='1",
		"'; DROP TABLE users; --",
		"1' UNION SELECT * FROM users--",
		"admin'--",
		"' OR 1=1--",
	}

	for _, attempt := range injectionAttempts {
		t.Run(fmt.Sprintf("injection_%s", attempt), func(t *testing.T) {
			var results []map[string]interface{}

			// This should be safe with parameterized queries
			// 매개변수화된 쿼리로 안전해야 함
			err := client.SelectAll(ctx, &results, "users", "name = ?", attempt)

			// Should return no results or error safely / 안전하게 결과 없음 또는 에러 반환
			if err != nil {
				t.Logf("Query error (expected): %v", err)
			}

			if len(results) > 0 {
				// Should not match unless there's actual data with this name
				// 이 이름의 실제 데이터가 없는 한 일치하지 않아야 함
				for _, r := range results {
					if r["name"] != attempt {
						t.Errorf("SQL injection vulnerability detected!")
					}
				}
			}
		})
	}

	// Verify database is still intact / 데이터베이스가 여전히 온전한지 검증
	count, err := client.Count(ctx, "users", "")
	if err != nil {
		t.Fatalf("database corruption detected: %v", err)
	}

	if count < 1 {
		t.Error("database appears to be corrupted or dropped")
	}
}
```

### Input Validation Test / 입력 검증 테스트

```go
func TestInputValidation(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"valid input", "hello", false},
		{"empty input", "", true},
		{"too long", strings.Repeat("a", 10000), true},
		{"null bytes", "hello\x00world", true},
		{"control characters", "hello\nworld\r\n", true},
		{"unicode", "안녕하세요", false},
		{"mixed", "Hello안녕🎉", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateInput(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateInput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
```

### Path Traversal Prevention Test / 경로 탐색 공격 방지 테스트

```go
func TestPathTraversalPrevention(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{"normal path", "file.txt", false},
		{"subdirectory", "sub/file.txt", false},
		{"parent directory", "../file.txt", true},
		{"absolute path", "/etc/passwd", true},
		{"windows absolute", "C:\\Windows\\System32", true},
		{"current directory", "./file.txt", false},
		{"multiple parent", "../../etc/passwd", true},
		{"encoded traversal", "..%2F..%2Fetc%2Fpasswd", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePath(tt.path)

			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePath(%q) error = %v, wantErr %v",
					tt.path, err, tt.wantErr)
			}
		})
	}
}
```

---

## Fuzz Testing / 퍼즈 테스트

Fuzz testing uses random inputs to find unexpected behaviors and crashes.
퍼즈 테스트는 랜덤 입력을 사용하여 예상치 못한 동작과 충돌을 찾습니다.

### Basic Fuzz Test / 기본 퍼즈 테스트

```go
// Go 1.18+ fuzzing support / Go 1.18+ 퍼징 지원
func FuzzReverse(f *testing.F) {
	// Seed corpus / 시드 코퍼스
	testcases := []string{
		"Hello, world",
		"안녕하세요",
		"",
		"!12345",
		"🎉🎊",
	}

	for _, tc := range testcases {
		f.Add(tc) // Add to corpus / 코퍼스에 추가
	}

	f.Fuzz(func(t *testing.T, input string) {
		// Test that Reverse doesn't panic / Reverse가 패닉하지 않는지 테스트
		result := Reverse(input)

		// Property: reversing twice gives original
		// 속성: 두 번 뒤집으면 원본
		doubleReverse := Reverse(result)
		if input != doubleReverse {
			t.Errorf("Reverse(Reverse(%q)) = %q, want %q", input, doubleReverse, input)
		}

		// Property: length is preserved
		// 속성: 길이 보존
		if len([]rune(input)) != len([]rune(result)) {
			t.Errorf("length changed: input=%d, result=%d",
				len([]rune(input)), len([]rune(result)))
		}
	})
}
```

### Advanced Fuzz Test / 고급 퍼즈 테스트

```go
func FuzzJSONParsing(f *testing.F) {
	// Seed with valid JSON / 유효한 JSON으로 시드
	f.Add(`{"name":"John","age":30}`)
	f.Add(`{"items":[1,2,3]}`)
	f.Add(`{}`)
	f.Add(`[]`)

	f.Fuzz(func(t *testing.T, data string) {
		var result interface{}

		// Should not panic on any input / 어떤 입력에도 패닉하지 않아야 함
		err := json.Unmarshal([]byte(data), &result)

		if err != nil {
			// Invalid JSON is expected, just ensure no panic
			// 유효하지 않은 JSON은 예상되며, 패닉만 없으면 됨
			return
		}

		// If parsing succeeded, should be able to marshal back
		// 파싱이 성공하면 다시 마샬링할 수 있어야 함
		_, err = json.Marshal(result)
		if err != nil {
			t.Errorf("failed to marshal parsed data: %v", err)
		}
	})
}
```

### Running Fuzz Tests / 퍼즈 테스트 실행

```bash
# Run fuzz test / 퍼즈 테스트 실행
go test -fuzz=FuzzReverse -fuzztime=30s ./stringutil

# Run with longer duration / 더 긴 시간으로 실행
go test -fuzz=FuzzReverse -fuzztime=5m ./stringutil

# Run all fuzz tests / 모든 퍼즈 테스트 실행
go test -fuzz=. ./...

# View corpus / 코퍼스 보기
ls testdata/fuzz/FuzzReverse/
```

---

## Property-Based Testing / 속성 기반 테스트

Property-based testing verifies that functions maintain certain properties.
속성 기반 테스트는 함수가 특정 속성을 유지하는지 검증합니다.

### Properties to Test / 테스트할 속성

1. **Idempotence / 멱등성**: `f(f(x)) == f(x)`
2. **Inverse / 역함수**: `f(g(x)) == x`
3. **Commutativity / 교환법칙**: `f(a, b) == f(b, a)`
4. **Associativity / 결합법칙**: `f(f(a, b), c) == f(a, f(b, c))`
5. **Identity / 항등원**: `f(x, identity) == x`

### Example Property Tests / 속성 테스트 예제

```go
func TestReverseProperties(t *testing.T) {
	testCases := generateRandomStrings(100)

	for i, s := range testCases {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			// Property 1: Reverse(Reverse(s)) == s
			// 속성 1: 두 번 뒤집으면 원본
			result := Reverse(Reverse(s))
			if result != s {
				t.Errorf("Reverse(Reverse(%q)) = %q, want %q", s, result, s)
			}

			// Property 2: Length is preserved
			// 속성 2: 길이 보존
			if len([]rune(s)) != len([]rune(Reverse(s))) {
				t.Errorf("length not preserved for %q", s)
			}

			// Property 3: Reversing empty string gives empty string
			// 속성 3: 빈 문자열 뒤집으면 빈 문자열
			if s == "" && Reverse(s) != "" {
				t.Error("reversing empty string should give empty string")
			}
		})
	}
}

func TestSortProperties(t *testing.T) {
	testSlices := generateRandomSlices(100)

	for i, slice := range testSlices {
		t.Run(fmt.Sprintf("slice_%d", i), func(t *testing.T) {
			original := Clone(slice)
			sorted := Sort(slice)

			// Property 1: Sorting is idempotent
			// 속성 1: 정렬은 멱등
			sortedTwice := Sort(sorted)
			if !Equal(sorted, sortedTwice) {
				t.Error("sorting is not idempotent")
			}

			// Property 2: Length is preserved
			// 속성 2: 길이 보존
			if len(original) != len(sorted) {
				t.Errorf("length changed: %d -> %d", len(original), len(sorted))
			}

			// Property 3: All elements from original are in sorted
			// 속성 3: 원본의 모든 요소가 정렬된 결과에 있음
			for _, v := range original {
				if !Contains(sorted, v) {
					t.Errorf("element %v missing after sort", v)
				}
			}

			// Property 4: Result is actually sorted
			// 속성 4: 결과가 실제로 정렬됨
			if !IsSorted(sorted) {
				t.Error("result is not sorted")
			}
		})
	}
}

func TestMapProperties(t *testing.T) {
	testSlices := generateRandomSlices(50)

	for i, slice := range testSlices {
		t.Run(fmt.Sprintf("slice_%d", i), func(t *testing.T) {
			// Property 1: Map with identity function returns same slice
			// 속성 1: 항등 함수로 Map하면 같은 슬라이스 반환
			identity := func(x int) int { return x }
			result := Map(slice, identity)
			if !Equal(slice, result) {
				t.Error("map with identity should return same slice")
			}

			// Property 2: Length is preserved
			// 속성 2: 길이 보존
			double := func(x int) int { return x * 2 }
			result = Map(slice, double)
			if len(slice) != len(result) {
				t.Errorf("length not preserved: %d -> %d", len(slice), len(result))
			}

			// Property 3: Composition of maps
			// 속성 3: Map 합성
			addOne := func(x int) int { return x + 1 }
			mulTwo := func(x int) int { return x * 2 }

			// Map(Map(s, f), g) == Map(s, compose(g, f))
			result1 := Map(Map(slice, addOne), mulTwo)
			composed := func(x int) int { return mulTwo(addOne(x)) }
			result2 := Map(slice, composed)

			if !Equal(result1, result2) {
				t.Error("map composition property violated")
			}
		})
	}
}

// Helper to generate random test data / 랜덤 테스트 데이터 생성 헬퍼
func generateRandomStrings(n int) []string {
	result := make([]string, n)
	for i := 0; i < n; i++ {
		length := rand.Intn(100)
		result[i] = randomString(length)
	}
	return result
}

func generateRandomSlices(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		size := rand.Intn(100)
		slice := make([]int, size)
		for j := 0; j < size; j++ {
			slice[j] = rand.Intn(1000)
		}
		result[i] = slice
	}
	return result
}
```

---

## Test Coverage / 테스트 커버리지

Test coverage measures how much code is tested.
테스트 커버리지는 얼마나 많은 코드가 테스트되는지 측정합니다.

### Running Coverage / 커버리지 실행

```bash
# Generate coverage report / 커버리지 보고서 생성
go test -cover ./...

# Generate detailed coverage profile / 상세 커버리지 프로파일 생성
go test -coverprofile=coverage.out ./...

# View coverage by function / 함수별 커버리지 보기
go tool cover -func=coverage.out

# Generate HTML coverage report / HTML 커버리지 보고서 생성
go tool cover -html=coverage.out -o coverage.html

# Coverage for specific package / 특정 패키지 커버리지
go test -cover ./stringutil

# Coverage with specific mode / 특정 모드로 커버리지
go test -covermode=atomic -coverprofile=coverage.out ./...
```

### Coverage Analysis / 커버리지 분석

```bash
# Check coverage percentage / 커버리지 비율 확인
$ go test -cover ./sliceutil
ok      github.com/arkd0ng/go-utils/sliceutil    0.156s  coverage: 100.0% of statements

# Detailed function coverage / 상세 함수 커버리지
$ go tool cover -func=coverage.out
github.com/arkd0ng/go-utils/sliceutil/basic.go:10:      Contains        100.0%
github.com/arkd0ng/go-utils/sliceutil/basic.go:20:      IndexOf         100.0%
github.com/arkd0ng/go-utils/sliceutil/transform.go:10:  Map             100.0%
github.com/arkd0ng/go-utils/sliceutil/transform.go:20:  Filter          100.0%
total:                                                  (statements)    100.0%
```

### Achieving 100% Coverage / 100% 커버리지 달성

```go
// Example: Function with multiple paths / 예제: 여러 경로가 있는 함수
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Test covering all paths / 모든 경로를 커버하는 테스트
func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a, b    int
		want    int
		wantErr bool
	}{
		// Success path / 성공 경로
		{"normal division", 10, 2, 5, false},
		{"negative result", -10, 2, -5, false},
		{"zero dividend", 0, 5, 0, false},

		// Error path / 에러 경로
		{"division by zero", 10, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)

			if tt.wantErr {
				if err == nil {
					t.Error("expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
```

### Coverage Guidelines / 커버리지 가이드라인

1. **Target 100% coverage / 100% 커버리지 목표**
   - All public functions must be tested / 모든 공개 함수는 테스트되어야 함
   - All error paths must be covered / 모든 에러 경로는 커버되어야 함

2. **Line coverage is not enough / 라인 커버리지만으로 부족**
   - Test all branches / 모든 분기 테스트
   - Test all conditions / 모든 조건 테스트
   - Test edge cases / 엣지 케이스 테스트

3. **Untested code / 테스트되지 않은 코드**
   ```go
   // Find uncovered code / 커버되지 않은 코드 찾기
   go test -coverprofile=coverage.out ./...
   go tool cover -html=coverage.out

   // Look for red-highlighted code in HTML report
   // HTML 보고서에서 빨간색으로 강조된 코드 찾기
   ```

---

## Testing Best Practices / 테스트 모범 사례

### 1. Use Table-Driven Tests / 테이블 기반 테스트 사용

```go
func TestFunction(t *testing.T) {
	tests := []struct {
		name     string
		input    Input
		expected Output
	}{
		{"case 1", input1, output1},
		{"case 2", input2, output2},
		// ... more cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// test logic
		})
	}
}
```

### 2. Use Subtests / 서브테스트 사용

```go
func TestFeature(t *testing.T) {
	t.Run("subfeature A", func(t *testing.T) {
		// test A
	})

	t.Run("subfeature B", func(t *testing.T) {
		// test B
	})
}
```

### 3. Use Test Helpers / 테스트 헬퍼 사용

```go
// testhelper_test.go
func newTestClient(t *testing.T) *Client {
	t.Helper() // Mark as helper / 헬퍼로 표시

	client, err := New(WithDSN(testDSN))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	return client
}

func assertEqual(t *testing.T, got, want interface{}) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
```

### 4. Clean Up Resources / 리소스 정리

```go
func TestWithResources(t *testing.T) {
	client := newTestClient(t)
	defer client.Close() // Always clean up / 항상 정리

	// ... test logic
}
```

### 5. Parallel Tests / 병렬 테스트

```go
func TestParallel(t *testing.T) {
	tests := []struct {
		name string
		// ...
	}{
		// test cases
	}

	for _, tt := range tests {
		tt := tt // Capture range variable / 범위 변수 캡처
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // Run in parallel / 병렬 실행
			// ... test logic
		})
	}
}
```

### 6. Test Naming / 테스트 이름 작성

```go
// Good / 좋음
func TestContains_EmptySlice_ReturnsFalse(t *testing.T)
func TestDivide_DivisionByZero_ReturnsError(t *testing.T)

// Bad / 나쁨
func TestFunc1(t *testing.T)
func TestCase2(t *testing.T)
```

### 7. Bilingual Comments / 이중 언어 주석

```go
// TestFunction tests the Function with various inputs
// TestFunction은 다양한 입력으로 Function을 테스트합니다
func TestFunction(t *testing.T) {
	// Test valid input / 유효한 입력 테스트
	// ...
}
```

### 8. Mock External Dependencies / 외부 의존성 모킹

```go
type mockClient struct {
	response interface{}
	err      error
}

func (m *mockClient) Call(ctx context.Context) (interface{}, error) {
	return m.response, m.err
}

func TestWithMock(t *testing.T) {
	mock := &mockClient{
		response: expectedResponse,
		err:      nil,
	}

	// Use mock in test / 테스트에서 모의 객체 사용
	result := someFunction(mock)
	// ... assertions
}
```

---

## Continuous Integration / 지속적 통합

### GitHub Actions Example / GitHub Actions 예제

```yaml
# .github/workflows/test.yml
name: Tests

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main, develop ]

jobs:
  test:
    runs-on: ubuntu-latest

    services:
      mysql:
        image: mysql:8.0
        env:
          MYSQL_ROOT_PASSWORD: testpassword
          MYSQL_DATABASE: testdb
        ports:
          - 3306:3306
        options: --health-cmd="mysqladmin ping" --health-interval=10s --health-timeout=5s --health-retries=3

      redis:
        image: redis:7-alpine
        ports:
          - 6379:6379
        options: --health-cmd="redis-cli ping" --health-interval=10s --health-timeout=5s --health-retries=3

    steps:
    - uses: actions/checkout@v3

    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.21'

    - name: Download dependencies
      run: go mod download

    - name: Run unit tests
      run: go test -v -race -coverprofile=coverage.out ./...

    - name: Run integration tests
      run: go test -v -tags=integration ./...
      env:
        MYSQL_DSN: root:testpassword@tcp(localhost:3306)/testdb
        REDIS_ADDR: localhost:6379

    - name: Check coverage
      run: |
        go tool cover -func=coverage.out
        COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print $3}' | sed 's/%//')
        echo "Total coverage: $COVERAGE%"
        if (( $(echo "$COVERAGE < 80" | bc -l) )); then
          echo "Coverage is below 80%"
          exit 1
        fi

    - name: Upload coverage to Codecov
      uses: codecov/codecov-action@v3
      with:
        files: ./coverage.out
        flags: unittests
        name: codecov-umbrella
```

---

## Summary / 요약

This guide covers comprehensive testing strategies for the go-utils project:
이 가이드는 go-utils 프로젝트의 포괄적인 테스트 전략을 다룹니다:

1. **Unit Testing** - Test individual functions / 개별 함수 테스트
2. **Edge Cases** - Test boundary conditions / 경계 조건 테스트
3. **Benchmarks** - Measure performance / 성능 측정
4. **Integration** - Test component interaction / 컴포넌트 상호작용 테스트
5. **Performance** - Verify speed and efficiency / 속도 및 효율성 검증
6. **Load Testing** - Test under expected load / 예상 부하 하 테스트
7. **Stress Testing** - Find breaking points / 한계점 찾기
8. **Security** - Identify vulnerabilities / 취약점 식별
9. **Fuzz Testing** - Find unexpected behaviors / 예상치 못한 동작 찾기
10. **Property-Based** - Verify invariants / 불변성 검증
11. **Coverage** - Achieve 100% goal / 100% 목표 달성

**Key Takeaways / 핵심 요점**:
- Write tests for all public functions / 모든 공개 함수에 대한 테스트 작성
- Cover all error paths / 모든 에러 경로 커버
- Use table-driven tests / 테이블 기반 테스트 사용
- Test edge cases thoroughly / 엣지 케이스 철저히 테스트
- Aim for 100% coverage / 100% 커버리지 목표
- Write clear, maintainable tests / 명확하고 유지보수 가능한 테스트 작성
- Use bilingual comments / 이중 언어 주석 사용

---

**Document Version / 문서 버전**: v1.11.x
**Last Updated / 최종 업데이트**: 2025-10-16
**Author / 작성자**: go-utils team
