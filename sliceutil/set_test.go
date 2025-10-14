package sliceutil

import (
	"testing"
)

// TestUnion tests the Union function.
// TestUnion은 Union 함수를 테스트합니다.
func TestUnion(t *testing.T) {
	t.Run("union with overlap", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{3, 4, 5}
		result := Union(a, b)
		expected := []int{1, 2, 3, 4, 5}
		if !Equal(result, expected) {
			t.Errorf("Union() = %v, want %v", result, expected)
		}
	})

	t.Run("union no overlap", func(t *testing.T) {
		a := []int{1, 2}
		b := []int{3, 4}
		result := Union(a, b)
		if len(result) != 4 {
			t.Errorf("Union() length = %v, want 4", len(result))
		}
	})

	t.Run("union with duplicates in same slice", func(t *testing.T) {
		a := []int{1, 1, 2, 2}
		b := []int{2, 3, 3}
		result := Union(a, b)
		expected := []int{1, 2, 3}
		if !Equal(result, expected) {
			t.Errorf("Union() = %v, want %v", result, expected)
		}
	})

	t.Run("union empty slices", func(t *testing.T) {
		a := []int{}
		b := []int{}
		result := Union(a, b)
		if len(result) != 0 {
			t.Errorf("Union() with empty slices should return empty, got %v", result)
		}
	})

	t.Run("union one empty", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{}
		result := Union(a, b)
		expected := []int{1, 2, 3}
		if !Equal(result, expected) {
			t.Errorf("Union() = %v, want %v", result, expected)
		}
	})
}

// BenchmarkUnion benchmarks the Union function.
// BenchmarkUnion은 Union 함수를 벤치마크합니다.
func BenchmarkUnion(b *testing.B) {
	a := make([]int, 500)
	for i := range a {
		a[i] = i
	}
	c := make([]int, 500)
	for i := range c {
		c[i] = i + 250
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Union(a, c)
	}
}

// TestIntersection tests the Intersection function.
// TestIntersection은 Intersection 함수를 테스트합니다.
func TestIntersection(t *testing.T) {
	t.Run("intersection with overlap", func(t *testing.T) {
		a := []int{1, 2, 3, 4}
		b := []int{3, 4, 5, 6}
		result := Intersection(a, b)
		expected := []int{3, 4}
		if !Equal(result, expected) {
			t.Errorf("Intersection() = %v, want %v", result, expected)
		}
	})

	t.Run("intersection no overlap", func(t *testing.T) {
		a := []int{1, 2}
		b := []int{3, 4}
		result := Intersection(a, b)
		if len(result) != 0 {
			t.Errorf("Intersection() should be empty, got %v", result)
		}
	})

	t.Run("intersection with duplicates", func(t *testing.T) {
		a := []int{1, 2, 2, 3, 3}
		b := []int{2, 2, 3}
		result := Intersection(a, b)
		expected := []int{2, 3}
		if !Equal(result, expected) {
			t.Errorf("Intersection() = %v, want %v", result, expected)
		}
	})

	t.Run("intersection all same", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{1, 2, 3}
		result := Intersection(a, b)
		expected := []int{1, 2, 3}
		if !Equal(result, expected) {
			t.Errorf("Intersection() = %v, want %v", result, expected)
		}
	})

	t.Run("intersection empty slices", func(t *testing.T) {
		a := []int{}
		b := []int{}
		result := Intersection(a, b)
		if len(result) != 0 {
			t.Errorf("Intersection() with empty slices should return empty, got %v", result)
		}
	})
}

// BenchmarkIntersection benchmarks the Intersection function.
// BenchmarkIntersection은 Intersection 함수를 벤치마크합니다.
func BenchmarkIntersection(b *testing.B) {
	a := make([]int, 500)
	for i := range a {
		a[i] = i
	}
	c := make([]int, 500)
	for i := range c {
		c[i] = i + 250
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Intersection(a, c)
	}
}

// TestDifference tests the Difference function.
// TestDifference는 Difference 함수를 테스트합니다.
func TestDifference(t *testing.T) {
	t.Run("difference with overlap", func(t *testing.T) {
		a := []int{1, 2, 3, 4}
		b := []int{3, 4, 5, 6}
		result := Difference(a, b)
		expected := []int{1, 2}
		if !Equal(result, expected) {
			t.Errorf("Difference() = %v, want %v", result, expected)
		}
	})

	t.Run("difference no overlap", func(t *testing.T) {
		a := []int{1, 2}
		b := []int{3, 4}
		result := Difference(a, b)
		expected := []int{1, 2}
		if !Equal(result, expected) {
			t.Errorf("Difference() = %v, want %v", result, expected)
		}
	})

	t.Run("difference with duplicates", func(t *testing.T) {
		a := []int{1, 1, 2, 3, 3}
		b := []int{2}
		result := Difference(a, b)
		expected := []int{1, 3}
		if !Equal(result, expected) {
			t.Errorf("Difference() = %v, want %v", result, expected)
		}
	})

	t.Run("difference all removed", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{1, 2, 3, 4, 5}
		result := Difference(a, b)
		if len(result) != 0 {
			t.Errorf("Difference() should be empty, got %v", result)
		}
	})

	t.Run("difference empty slices", func(t *testing.T) {
		a := []int{}
		b := []int{1, 2}
		result := Difference(a, b)
		if len(result) != 0 {
			t.Errorf("Difference() with empty a should return empty, got %v", result)
		}
	})
}

// BenchmarkDifference benchmarks the Difference function.
// BenchmarkDifference는 Difference 함수를 벤치마크합니다.
func BenchmarkDifference(b *testing.B) {
	a := make([]int, 500)
	for i := range a {
		a[i] = i
	}
	c := make([]int, 500)
	for i := range c {
		c[i] = i + 250
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Difference(a, c)
	}
}

// TestSymmetricDifference tests the SymmetricDifference function.
// TestSymmetricDifference는 SymmetricDifference 함수를 테스트합니다.
func TestSymmetricDifference(t *testing.T) {
	t.Run("symmetric difference with overlap", func(t *testing.T) {
		a := []int{1, 2, 3, 4}
		b := []int{3, 4, 5, 6}
		result := SymmetricDifference(a, b)
		expected := []int{1, 2, 5, 6}
		if !Equal(result, expected) {
			t.Errorf("SymmetricDifference() = %v, want %v", result, expected)
		}
	})

	t.Run("symmetric difference no overlap", func(t *testing.T) {
		a := []int{1, 2}
		b := []int{3, 4}
		result := SymmetricDifference(a, b)
		if len(result) != 4 {
			t.Errorf("SymmetricDifference() length = %v, want 4", len(result))
		}
	})

	t.Run("symmetric difference all same", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{1, 2, 3}
		result := SymmetricDifference(a, b)
		if len(result) != 0 {
			t.Errorf("SymmetricDifference() should be empty, got %v", result)
		}
	})

	t.Run("symmetric difference with duplicates", func(t *testing.T) {
		a := []int{1, 1, 2}
		b := []int{2, 2, 3}
		result := SymmetricDifference(a, b)
		expected := []int{1, 3}
		if !Equal(result, expected) {
			t.Errorf("SymmetricDifference() = %v, want %v", result, expected)
		}
	})

	t.Run("symmetric difference empty slices", func(t *testing.T) {
		a := []int{}
		b := []int{}
		result := SymmetricDifference(a, b)
		if len(result) != 0 {
			t.Errorf("SymmetricDifference() with empty slices should return empty, got %v", result)
		}
	})
}

// BenchmarkSymmetricDifference benchmarks the SymmetricDifference function.
// BenchmarkSymmetricDifference는 SymmetricDifference 함수를 벤치마크합니다.
func BenchmarkSymmetricDifference(b *testing.B) {
	a := make([]int, 500)
	for i := range a {
		a[i] = i
	}
	c := make([]int, 500)
	for i := range c {
		c[i] = i + 250
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SymmetricDifference(a, c)
	}
}

// TestIsSubset tests the IsSubset function.
// TestIsSubset은 IsSubset 함수를 테스트합니다.
func TestIsSubset(t *testing.T) {
	t.Run("is subset true", func(t *testing.T) {
		a := []int{1, 2}
		b := []int{1, 2, 3, 4}
		if !IsSubset(a, b) {
			t.Errorf("IsSubset() should be true")
		}
	})

	t.Run("is subset false", func(t *testing.T) {
		a := []int{1, 5}
		b := []int{1, 2, 3, 4}
		if IsSubset(a, b) {
			t.Errorf("IsSubset() should be false")
		}
	})

	t.Run("is subset same", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{1, 2, 3}
		if !IsSubset(a, b) {
			t.Errorf("IsSubset() should be true for same sets")
		}
	})

	t.Run("is subset empty", func(t *testing.T) {
		a := []int{}
		b := []int{1, 2, 3}
		if !IsSubset(a, b) {
			t.Errorf("IsSubset() empty set should be subset of any set")
		}
	})

	t.Run("is subset with duplicates", func(t *testing.T) {
		a := []int{1, 1, 2}
		b := []int{1, 2, 3}
		if !IsSubset(a, b) {
			t.Errorf("IsSubset() should be true")
		}
	})
}

// BenchmarkIsSubset benchmarks the IsSubset function.
// BenchmarkIsSubset은 IsSubset 함수를 벤치마크합니다.
func BenchmarkIsSubset(b *testing.B) {
	a := make([]int, 250)
	for i := range a {
		a[i] = i
	}
	c := make([]int, 500)
	for i := range c {
		c[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsSubset(a, c)
	}
}

// TestIsSuperset tests the IsSuperset function.
// TestIsSuperset은 IsSuperset 함수를 테스트합니다.
func TestIsSuperset(t *testing.T) {
	t.Run("is superset true", func(t *testing.T) {
		a := []int{1, 2, 3, 4}
		b := []int{1, 2}
		if !IsSuperset(a, b) {
			t.Errorf("IsSuperset() should be true")
		}
	})

	t.Run("is superset false", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{1, 5}
		if IsSuperset(a, b) {
			t.Errorf("IsSuperset() should be false")
		}
	})

	t.Run("is superset same", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{1, 2, 3}
		if !IsSuperset(a, b) {
			t.Errorf("IsSuperset() should be true for same sets")
		}
	})

	t.Run("is superset empty", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{}
		if !IsSuperset(a, b) {
			t.Errorf("IsSuperset() any set should be superset of empty set")
		}
	})

	t.Run("is superset with duplicates", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{1, 1, 2}
		if !IsSuperset(a, b) {
			t.Errorf("IsSuperset() should be true")
		}
	})
}

// BenchmarkIsSuperset benchmarks the IsSuperset function.
// BenchmarkIsSuperset은 IsSuperset 함수를 벤치마크합니다.
func BenchmarkIsSuperset(b *testing.B) {
	a := make([]int, 500)
	for i := range a {
		a[i] = i
	}
	c := make([]int, 250)
	for i := range c {
		c[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsSuperset(a, c)
	}
}
