package sliceutil

import (
	"math"
	"testing"
)

// TestMedian tests the Median function.
// TestMedian은 Median 함수를 테스트합니다.
func TestMedian(t *testing.T) {
	t.Run("odd length slice", func(t *testing.T) {
		numbers := []int{3, 1, 4, 1, 5, 9, 2}
		median, err := Median(numbers)
		if err != nil {
			t.Fatalf("Median() error = %v", err)
		}
		expected := 3.0
		if median != expected {
			t.Errorf("Median() = %v, want %v", median, expected)
		}
	})

	t.Run("even length slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		median, err := Median(numbers)
		if err != nil {
			t.Fatalf("Median() error = %v", err)
		}
		expected := 2.5
		if median != expected {
			t.Errorf("Median() = %v, want %v", median, expected)
		}
	})

	t.Run("single element", func(t *testing.T) {
		numbers := []int{42}
		median, err := Median(numbers)
		if err != nil {
			t.Fatalf("Median() error = %v", err)
		}
		expected := 42.0
		if median != expected {
			t.Errorf("Median() = %v, want %v", median, expected)
		}
	})

	t.Run("floats", func(t *testing.T) {
		numbers := []float64{1.5, 2.5, 3.5, 4.5, 5.5}
		median, err := Median(numbers)
		if err != nil {
			t.Fatalf("Median() error = %v", err)
		}
		expected := 3.5
		if median != expected {
			t.Errorf("Median() = %v, want %v", median, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		_, err := Median(numbers)
		if err == nil {
			t.Error("Median() should return error for empty slice")
		}
	})

	t.Run("unsorted slice", func(t *testing.T) {
		numbers := []int{10, 1, 5, 3, 8}
		median, err := Median(numbers)
		if err != nil {
			t.Fatalf("Median() error = %v", err)
		}
		expected := 5.0
		if median != expected {
			t.Errorf("Median() = %v, want %v", median, expected)
		}
	})
}

// BenchmarkMedian benchmarks the Median function.
// BenchmarkMedian은 Median 함수를 벤치마크합니다.
func BenchmarkMedian(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Median(numbers)
	}
}

// TestMode tests the Mode function.
// TestMode는 Mode 함수를 테스트합니다.
func TestMode(t *testing.T) {
	t.Run("clear mode", func(t *testing.T) {
		numbers := []int{1, 2, 2, 3, 3, 3, 4}
		mode, err := Mode(numbers)
		if err != nil {
			t.Fatalf("Mode() error = %v", err)
		}
		expected := 3
		if mode != expected {
			t.Errorf("Mode() = %v, want %v", mode, expected)
		}
	})

	t.Run("multiple modes returns first", func(t *testing.T) {
		numbers := []int{1, 1, 2, 2, 3}
		mode, err := Mode(numbers)
		if err != nil {
			t.Fatalf("Mode() error = %v", err)
		}
		// Either 1 or 2 is acceptable (both appear twice)
		if mode != 1 && mode != 2 {
			t.Errorf("Mode() = %v, want 1 or 2", mode)
		}
	})

	t.Run("all unique", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		mode, err := Mode(numbers)
		if err != nil {
			t.Fatalf("Mode() error = %v", err)
		}
		// Any value is acceptable since all appear once
		if !Contains(numbers, mode) {
			t.Errorf("Mode() = %v, not in slice", mode)
		}
	})

	t.Run("single element", func(t *testing.T) {
		numbers := []int{42}
		mode, err := Mode(numbers)
		if err != nil {
			t.Fatalf("Mode() error = %v", err)
		}
		expected := 42
		if mode != expected {
			t.Errorf("Mode() = %v, want %v", mode, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		_, err := Mode(numbers)
		if err == nil {
			t.Error("Mode() should return error for empty slice")
		}
	})

	t.Run("strings", func(t *testing.T) {
		words := []string{"apple", "banana", "banana", "cherry", "banana"}
		mode, err := Mode(words)
		if err != nil {
			t.Fatalf("Mode() error = %v", err)
		}
		expected := "banana"
		if mode != expected {
			t.Errorf("Mode() = %v, want %v", mode, expected)
		}
	})
}

// BenchmarkMode benchmarks the Mode function.
// BenchmarkMode는 Mode 함수를 벤치마크합니다.
func BenchmarkMode(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i % 100
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Mode(numbers)
	}
}

// TestFrequencies tests the Frequencies function.
// TestFrequencies는 Frequencies 함수를 테스트합니다.
func TestFrequencies(t *testing.T) {
	t.Run("basic frequencies", func(t *testing.T) {
		numbers := []int{1, 2, 2, 3, 3, 3, 4}
		freq := Frequencies(numbers)
		if freq[1] != 1 {
			t.Errorf("Frequencies()[1] = %v, want 1", freq[1])
		}
		if freq[2] != 2 {
			t.Errorf("Frequencies()[2] = %v, want 2", freq[2])
		}
		if freq[3] != 3 {
			t.Errorf("Frequencies()[3] = %v, want 3", freq[3])
		}
		if freq[4] != 1 {
			t.Errorf("Frequencies()[4] = %v, want 1", freq[4])
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		freq := Frequencies(numbers)
		if len(freq) != 0 {
			t.Errorf("Frequencies() should return empty map for empty slice, got %v", freq)
		}
	})

	t.Run("all unique", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		freq := Frequencies(numbers)
		if len(freq) != 5 {
			t.Errorf("Frequencies() should have 5 entries, got %v", len(freq))
		}
		for k, v := range freq {
			if v != 1 {
				t.Errorf("Frequencies()[%v] = %v, want 1", k, v)
			}
		}
	})

	t.Run("strings", func(t *testing.T) {
		words := []string{"a", "b", "b", "c", "c", "c"}
		freq := Frequencies(words)
		if freq["a"] != 1 || freq["b"] != 2 || freq["c"] != 3 {
			t.Errorf("Frequencies() = %v, incorrect counts", freq)
		}
	})
}

// BenchmarkFrequencies benchmarks the Frequencies function.
// BenchmarkFrequencies는 Frequencies 함수를 벤치마크합니다.
func BenchmarkFrequencies(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i % 100
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Frequencies(numbers)
	}
}

// TestPercentile tests the Percentile function.
// TestPercentile은 Percentile 함수를 테스트합니다.
func TestPercentile(t *testing.T) {
	t.Run("50th percentile (median)", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		p50, err := Percentile(numbers, 50)
		if err != nil {
			t.Fatalf("Percentile() error = %v", err)
		}
		expected := 5.5
		if math.Abs(p50-expected) > 0.01 {
			t.Errorf("Percentile(50) = %v, want %v", p50, expected)
		}
	})

	t.Run("75th percentile", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		p75, err := Percentile(numbers, 75)
		if err != nil {
			t.Fatalf("Percentile() error = %v", err)
		}
		expected := 7.75
		if math.Abs(p75-expected) > 0.01 {
			t.Errorf("Percentile(75) = %v, want %v", p75, expected)
		}
	})

	t.Run("90th percentile", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		p90, err := Percentile(numbers, 90)
		if err != nil {
			t.Fatalf("Percentile() error = %v", err)
		}
		expected := 9.1
		if math.Abs(p90-expected) > 0.01 {
			t.Errorf("Percentile(90) = %v, want %v", p90, expected)
		}
	})

	t.Run("0th percentile (min)", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		p0, err := Percentile(numbers, 0)
		if err != nil {
			t.Fatalf("Percentile() error = %v", err)
		}
		expected := 1.0
		if p0 != expected {
			t.Errorf("Percentile(0) = %v, want %v", p0, expected)
		}
	})

	t.Run("100th percentile (max)", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		p100, err := Percentile(numbers, 100)
		if err != nil {
			t.Fatalf("Percentile() error = %v", err)
		}
		expected := 5.0
		if p100 != expected {
			t.Errorf("Percentile(100) = %v, want %v", p100, expected)
		}
	})

	t.Run("invalid percentile negative", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		_, err := Percentile(numbers, -1)
		if err == nil {
			t.Error("Percentile() should return error for negative percentile")
		}
	})

	t.Run("invalid percentile over 100", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		_, err := Percentile(numbers, 101)
		if err == nil {
			t.Error("Percentile() should return error for percentile > 100")
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		_, err := Percentile(numbers, 50)
		if err == nil {
			t.Error("Percentile() should return error for empty slice")
		}
	})
}

// BenchmarkPercentile benchmarks the Percentile function.
// BenchmarkPercentile은 Percentile 함수를 벤치마크합니다.
func BenchmarkPercentile(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Percentile(numbers, 75)
	}
}

// TestStandardDeviation tests the StandardDeviation function.
// TestStandardDeviation은 StandardDeviation 함수를 테스트합니다.
func TestStandardDeviation(t *testing.T) {
	t.Run("known standard deviation", func(t *testing.T) {
		numbers := []float64{2, 4, 4, 4, 5, 5, 7, 9}
		stddev, err := StandardDeviation(numbers)
		if err != nil {
			t.Fatalf("StandardDeviation() error = %v", err)
		}
		expected := 2.0
		if math.Abs(stddev-expected) > 0.01 {
			t.Errorf("StandardDeviation() = %v, want %v", stddev, expected)
		}
	})

	t.Run("all same values", func(t *testing.T) {
		numbers := []int{5, 5, 5, 5, 5}
		stddev, err := StandardDeviation(numbers)
		if err != nil {
			t.Fatalf("StandardDeviation() error = %v", err)
		}
		expected := 0.0
		if stddev != expected {
			t.Errorf("StandardDeviation() = %v, want %v", stddev, expected)
		}
	})

	t.Run("single element", func(t *testing.T) {
		numbers := []int{42}
		stddev, err := StandardDeviation(numbers)
		if err != nil {
			t.Fatalf("StandardDeviation() error = %v", err)
		}
		expected := 0.0
		if stddev != expected {
			t.Errorf("StandardDeviation() = %v, want %v", stddev, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		_, err := StandardDeviation(numbers)
		if err == nil {
			t.Error("StandardDeviation() should return error for empty slice")
		}
	})
}

// BenchmarkStandardDeviation benchmarks the StandardDeviation function.
// BenchmarkStandardDeviation은 StandardDeviation 함수를 벤치마크합니다.
func BenchmarkStandardDeviation(b *testing.B) {
	numbers := make([]float64, 1000)
	for i := range numbers {
		numbers[i] = float64(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StandardDeviation(numbers)
	}
}

// TestVariance tests the Variance function.
// TestVariance는 Variance 함수를 테스트합니다.
func TestVariance(t *testing.T) {
	t.Run("known variance", func(t *testing.T) {
		numbers := []float64{2, 4, 4, 4, 5, 5, 7, 9}
		variance, err := Variance(numbers)
		if err != nil {
			t.Fatalf("Variance() error = %v", err)
		}
		expected := 4.0
		if math.Abs(variance-expected) > 0.01 {
			t.Errorf("Variance() = %v, want %v", variance, expected)
		}
	})

	t.Run("all same values", func(t *testing.T) {
		numbers := []int{5, 5, 5, 5, 5}
		variance, err := Variance(numbers)
		if err != nil {
			t.Fatalf("Variance() error = %v", err)
		}
		expected := 0.0
		if variance != expected {
			t.Errorf("Variance() = %v, want %v", variance, expected)
		}
	})

	t.Run("simple case", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		variance, err := Variance(numbers)
		if err != nil {
			t.Fatalf("Variance() error = %v", err)
		}
		// Mean = 3, variance = ((1-3)^2 + (2-3)^2 + (3-3)^2 + (4-3)^2 + (5-3)^2) / 5 = 10/5 = 2
		expected := 2.0
		if math.Abs(variance-expected) > 0.01 {
			t.Errorf("Variance() = %v, want %v", variance, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		_, err := Variance(numbers)
		if err == nil {
			t.Error("Variance() should return error for empty slice")
		}
	})
}

// BenchmarkVariance benchmarks the Variance function.
// BenchmarkVariance는 Variance 함수를 벤치마크합니다.
func BenchmarkVariance(b *testing.B) {
	numbers := make([]float64, 1000)
	for i := range numbers {
		numbers[i] = float64(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Variance(numbers)
	}
}

// TestMostCommon tests the MostCommon function.
// TestMostCommon은 MostCommon 함수를 테스트합니다.
func TestMostCommon(t *testing.T) {
	t.Run("top 2 most common", func(t *testing.T) {
		numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
		top2 := MostCommon(numbers, 2)
		if len(top2) != 2 {
			t.Fatalf("MostCommon(2) should return 2 elements, got %v", len(top2))
		}
		if top2[0] != 4 {
			t.Errorf("MostCommon(2)[0] = %v, want 4", top2[0])
		}
		if top2[1] != 3 {
			t.Errorf("MostCommon(2)[1] = %v, want 3", top2[1])
		}
	})

	t.Run("top 1 most common", func(t *testing.T) {
		numbers := []int{1, 1, 1, 2, 2, 3}
		top1 := MostCommon(numbers, 1)
		if len(top1) != 1 {
			t.Fatalf("MostCommon(1) should return 1 element, got %v", len(top1))
		}
		if top1[0] != 1 {
			t.Errorf("MostCommon(1)[0] = %v, want 1", top1[0])
		}
	})

	t.Run("n greater than unique count", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		top := MostCommon(numbers, 10)
		if len(top) != 3 {
			t.Errorf("MostCommon(10) should return 3 elements, got %v", len(top))
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		result := MostCommon(numbers, 2)
		if len(result) != 0 {
			t.Errorf("MostCommon() on empty slice should return empty slice, got %v", result)
		}
	})

	t.Run("n is zero", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := MostCommon(numbers, 0)
		if len(result) != 0 {
			t.Errorf("MostCommon(0) should return empty slice, got %v", result)
		}
	})

	t.Run("ties ordered by first occurrence", func(t *testing.T) {
		numbers := []int{1, 2, 3, 1, 2, 3}
		top := MostCommon(numbers, 3)
		// All have same frequency, should be ordered by first occurrence: 1, 2, 3
		expected := []int{1, 2, 3}
		if !Equal(top, expected) {
			t.Errorf("MostCommon() = %v, want %v", top, expected)
		}
	})
}

// BenchmarkMostCommon benchmarks the MostCommon function.
// BenchmarkMostCommon은 MostCommon 함수를 벤치마크합니다.
func BenchmarkMostCommon(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i % 100
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MostCommon(numbers, 10)
	}
}

// TestLeastCommon tests the LeastCommon function.
// TestLeastCommon은 LeastCommon 함수를 테스트합니다.
func TestLeastCommon(t *testing.T) {
	t.Run("bottom 2 least common", func(t *testing.T) {
		numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
		bottom2 := LeastCommon(numbers, 2)
		if len(bottom2) != 2 {
			t.Fatalf("LeastCommon(2) should return 2 elements, got %v", len(bottom2))
		}
		if bottom2[0] != 1 {
			t.Errorf("LeastCommon(2)[0] = %v, want 1", bottom2[0])
		}
		if bottom2[1] != 2 {
			t.Errorf("LeastCommon(2)[1] = %v, want 2", bottom2[1])
		}
	})

	t.Run("bottom 1 least common", func(t *testing.T) {
		numbers := []int{1, 2, 2, 3, 3, 3}
		bottom1 := LeastCommon(numbers, 1)
		if len(bottom1) != 1 {
			t.Fatalf("LeastCommon(1) should return 1 element, got %v", len(bottom1))
		}
		if bottom1[0] != 1 {
			t.Errorf("LeastCommon(1)[0] = %v, want 1", bottom1[0])
		}
	})

	t.Run("n greater than unique count", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		bottom := LeastCommon(numbers, 10)
		if len(bottom) != 3 {
			t.Errorf("LeastCommon(10) should return 3 elements, got %v", len(bottom))
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		result := LeastCommon(numbers, 2)
		if len(result) != 0 {
			t.Errorf("LeastCommon() on empty slice should return empty slice, got %v", result)
		}
	})

	t.Run("n is zero", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := LeastCommon(numbers, 0)
		if len(result) != 0 {
			t.Errorf("LeastCommon(0) should return empty slice, got %v", result)
		}
	})

	t.Run("ties ordered by first occurrence", func(t *testing.T) {
		numbers := []int{3, 2, 1, 3, 2, 1}
		bottom := LeastCommon(numbers, 3)
		// All have same frequency, should be ordered by first occurrence: 3, 2, 1
		expected := []int{3, 2, 1}
		if !Equal(bottom, expected) {
			t.Errorf("LeastCommon() = %v, want %v", bottom, expected)
		}
	})
}

// BenchmarkLeastCommon benchmarks the LeastCommon function.
// BenchmarkLeastCommon은 LeastCommon 함수를 벤치마크합니다.
func BenchmarkLeastCommon(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i % 100
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LeastCommon(numbers, 10)
	}
}
