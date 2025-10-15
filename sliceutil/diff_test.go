package sliceutil

import (
	"testing"
)

// TestDiff tests the Diff function.
// TestDiff는 Diff 함수를 테스트합니다.
func TestDiff(t *testing.T) {
	t.Run("basic diff", func(t *testing.T) {
		old := []int{1, 2, 3, 4}
		new := []int{2, 3, 4, 5}
		diff := Diff(old, new)

		if !Equal(diff.Added, []int{5}) {
			t.Errorf("Diff().Added = %v, want [5]", diff.Added)
		}
		if !Equal(diff.Removed, []int{1}) {
			t.Errorf("Diff().Removed = %v, want [1]", diff.Removed)
		}
		if !EqualUnordered(diff.Unchanged, []int{2, 3, 4}) {
			t.Errorf("Diff().Unchanged = %v, want [2, 3, 4]", diff.Unchanged)
		}
	})

	t.Run("all added", func(t *testing.T) {
		old := []int{}
		new := []int{1, 2, 3}
		diff := Diff(old, new)

		if !Equal(diff.Added, []int{1, 2, 3}) {
			t.Errorf("Diff().Added = %v, want [1, 2, 3]", diff.Added)
		}
		if len(diff.Removed) != 0 {
			t.Errorf("Diff().Removed should be empty, got %v", diff.Removed)
		}
		if len(diff.Unchanged) != 0 {
			t.Errorf("Diff().Unchanged should be empty, got %v", diff.Unchanged)
		}
	})

	t.Run("all removed", func(t *testing.T) {
		old := []int{1, 2, 3}
		new := []int{}
		diff := Diff(old, new)

		if len(diff.Added) != 0 {
			t.Errorf("Diff().Added should be empty, got %v", diff.Added)
		}
		if !Equal(diff.Removed, []int{1, 2, 3}) {
			t.Errorf("Diff().Removed = %v, want [1, 2, 3]", diff.Removed)
		}
		if len(diff.Unchanged) != 0 {
			t.Errorf("Diff().Unchanged should be empty, got %v", diff.Unchanged)
		}
	})

	t.Run("no changes", func(t *testing.T) {
		old := []int{1, 2, 3}
		new := []int{1, 2, 3}
		diff := Diff(old, new)

		if len(diff.Added) != 0 {
			t.Errorf("Diff().Added should be empty, got %v", diff.Added)
		}
		if len(diff.Removed) != 0 {
			t.Errorf("Diff().Removed should be empty, got %v", diff.Removed)
		}
		if !EqualUnordered(diff.Unchanged, []int{1, 2, 3}) {
			t.Errorf("Diff().Unchanged = %v, want [1, 2, 3]", diff.Unchanged)
		}
	})

	t.Run("both empty", func(t *testing.T) {
		old := []int{}
		new := []int{}
		diff := Diff(old, new)

		if len(diff.Added) != 0 || len(diff.Removed) != 0 || len(diff.Unchanged) != 0 {
			t.Errorf("Diff() of empty slices should have all empty results")
		}
	})

	t.Run("strings", func(t *testing.T) {
		old := []string{"apple", "banana", "cherry"}
		new := []string{"banana", "cherry", "date"}
		diff := Diff(old, new)

		if !Equal(diff.Added, []string{"date"}) {
			t.Errorf("Diff().Added = %v, want [date]", diff.Added)
		}
		if !Equal(diff.Removed, []string{"apple"}) {
			t.Errorf("Diff().Removed = %v, want [apple]", diff.Removed)
		}
	})

	t.Run("with duplicates", func(t *testing.T) {
		old := []int{1, 2, 2, 3}
		new := []int{2, 2, 3, 4}
		diff := Diff(old, new)

		if !Contains(diff.Added, 4) {
			t.Errorf("Diff().Added should contain 4, got %v", diff.Added)
		}
		if !Contains(diff.Removed, 1) {
			t.Errorf("Diff().Removed should contain 1, got %v", diff.Removed)
		}
	})
}

// BenchmarkDiff benchmarks the Diff function.
// BenchmarkDiff는 Diff 함수를 벤치마크합니다.
func BenchmarkDiff(b *testing.B) {
	old := make([]int, 1000)
	new := make([]int, 1000)
	for i := range old {
		old[i] = i
		new[i] = i + 100
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Diff(old, new)
	}
}

// TestDiffBy tests the DiffBy function.
// TestDiffBy는 DiffBy 함수를 테스트합니다.
func TestDiffBy(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	t.Run("diff by ID", func(t *testing.T) {
		old := []User{{1, "Alice"}, {2, "Bob"}, {3, "Charlie"}}
		new := []User{{2, "Bob"}, {3, "Charlie"}, {4, "David"}}
		diff := DiffBy(old, new, func(u User) int { return u.ID })

		// Check added
		if len(diff.Added) != 1 {
			t.Fatalf("DiffBy().Added should have 1 element, got %v", len(diff.Added))
		}
		if diff.Added[0].ID != 4 {
			t.Errorf("DiffBy().Added[0].ID = %v, want 4", diff.Added[0].ID)
		}

		// Check removed
		if len(diff.Removed) != 1 {
			t.Fatalf("DiffBy().Removed should have 1 element, got %v", len(diff.Removed))
		}
		if diff.Removed[0].ID != 1 {
			t.Errorf("DiffBy().Removed[0].ID = %v, want 1", diff.Removed[0].ID)
		}

		// Check unchanged
		if len(diff.Unchanged) != 2 {
			t.Errorf("DiffBy().Unchanged should have 2 elements, got %v", len(diff.Unchanged))
		}
	})

	t.Run("all added", func(t *testing.T) {
		old := []User{}
		new := []User{{1, "Alice"}, {2, "Bob"}}
		diff := DiffBy(old, new, func(u User) int { return u.ID })

		if len(diff.Added) != 2 {
			t.Errorf("DiffBy().Added should have 2 elements, got %v", len(diff.Added))
		}
		if len(diff.Removed) != 0 {
			t.Errorf("DiffBy().Removed should be empty, got %v", diff.Removed)
		}
	})

	t.Run("all removed", func(t *testing.T) {
		old := []User{{1, "Alice"}, {2, "Bob"}}
		new := []User{}
		diff := DiffBy(old, new, func(u User) int { return u.ID })

		if len(diff.Added) != 0 {
			t.Errorf("DiffBy().Added should be empty, got %v", diff.Added)
		}
		if len(diff.Removed) != 2 {
			t.Errorf("DiffBy().Removed should have 2 elements, got %v", len(diff.Removed))
		}
	})

	t.Run("no changes", func(t *testing.T) {
		old := []User{{1, "Alice"}, {2, "Bob"}}
		new := []User{{1, "Alice"}, {2, "Bob"}}
		diff := DiffBy(old, new, func(u User) int { return u.ID })

		if len(diff.Added) != 0 {
			t.Errorf("DiffBy().Added should be empty, got %v", diff.Added)
		}
		if len(diff.Removed) != 0 {
			t.Errorf("DiffBy().Removed should be empty, got %v", diff.Removed)
		}
		if len(diff.Unchanged) != 2 {
			t.Errorf("DiffBy().Unchanged should have 2 elements, got %v", len(diff.Unchanged))
		}
	})

	t.Run("name changed but ID same", func(t *testing.T) {
		old := []User{{1, "Alice"}}
		new := []User{{1, "Alicia"}} // Same ID, different name
		diff := DiffBy(old, new, func(u User) int { return u.ID })

		// Same ID, so should be in unchanged
		if len(diff.Added) != 0 {
			t.Errorf("DiffBy().Added should be empty, got %v", diff.Added)
		}
		if len(diff.Removed) != 0 {
			t.Errorf("DiffBy().Removed should be empty, got %v", diff.Removed)
		}
		if len(diff.Unchanged) != 1 {
			t.Errorf("DiffBy().Unchanged should have 1 element, got %v", len(diff.Unchanged))
		}
	})
}

// BenchmarkDiffBy benchmarks the DiffBy function.
// BenchmarkDiffBy는 DiffBy 함수를 벤치마크합니다.
func BenchmarkDiffBy(b *testing.B) {
	type User struct {
		ID   int
		Name string
	}

	old := make([]User, 1000)
	new := make([]User, 1000)
	for i := range old {
		old[i] = User{i, "User"}
		new[i] = User{i + 100, "User"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DiffBy(old, new, func(u User) int { return u.ID })
	}
}

// TestEqualUnordered tests the EqualUnordered function.
// TestEqualUnordered는 EqualUnordered 함수를 테스트합니다.
func TestEqualUnordered(t *testing.T) {
	t.Run("equal unordered", func(t *testing.T) {
		a := []int{1, 2, 3, 2}
		b := []int{3, 2, 1, 2}
		if !EqualUnordered(a, b) {
			t.Error("EqualUnordered() should return true for equal unordered slices")
		}
	})

	t.Run("different lengths", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{1, 2}
		if EqualUnordered(a, b) {
			t.Error("EqualUnordered() should return false for different lengths")
		}
	})

	t.Run("different elements", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{1, 2, 4}
		if EqualUnordered(a, b) {
			t.Error("EqualUnordered() should return false for different elements")
		}
	})

	t.Run("different frequencies", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{1, 2, 3, 3}
		if EqualUnordered(a, b) {
			t.Error("EqualUnordered() should return false for different frequencies")
		}
	})

	t.Run("both empty", func(t *testing.T) {
		a := []int{}
		b := []int{}
		if !EqualUnordered(a, b) {
			t.Error("EqualUnordered() should return true for two empty slices")
		}
	})

	t.Run("same slice same order", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}
		b := []int{1, 2, 3, 4, 5}
		if !EqualUnordered(a, b) {
			t.Error("EqualUnordered() should return true for same slices")
		}
	})

	t.Run("strings", func(t *testing.T) {
		a := []string{"apple", "banana", "cherry"}
		b := []string{"cherry", "apple", "banana"}
		if !EqualUnordered(a, b) {
			t.Error("EqualUnordered() should return true for equal unordered string slices")
		}
	})

	t.Run("one empty one not", func(t *testing.T) {
		a := []int{}
		b := []int{1}
		if EqualUnordered(a, b) {
			t.Error("EqualUnordered() should return false when one is empty and one is not")
		}
	})

	t.Run("different unique count", func(t *testing.T) {
		a := []int{1, 1, 1}
		b := []int{1, 2, 3}
		if EqualUnordered(a, b) {
			t.Error("EqualUnordered() should return false when unique element counts differ")
		}
	})
}

// BenchmarkEqualUnordered benchmarks the EqualUnordered function.
// BenchmarkEqualUnordered는 EqualUnordered 함수를 벤치마크합니다.
func BenchmarkEqualUnordered(b *testing.B) {
	a := make([]int, 1000)
	b2 := make([]int, 1000)
	for i := range a {
		a[i] = i
		b2[999-i] = i // Reverse order
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EqualUnordered(a, b2)
	}
}

// TestHasDuplicates tests the HasDuplicates function.
// TestHasDuplicates는 HasDuplicates 함수를 테스트합니다.
func TestHasDuplicates(t *testing.T) {
	t.Run("has duplicates", func(t *testing.T) {
		numbers := []int{1, 2, 3, 2, 4}
		if !HasDuplicates(numbers) {
			t.Error("HasDuplicates() should return true for slice with duplicates")
		}
	})

	t.Run("no duplicates", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		if HasDuplicates(numbers) {
			t.Error("HasDuplicates() should return false for slice without duplicates")
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		if HasDuplicates(numbers) {
			t.Error("HasDuplicates() should return false for empty slice")
		}
	})

	t.Run("single element", func(t *testing.T) {
		numbers := []int{42}
		if HasDuplicates(numbers) {
			t.Error("HasDuplicates() should return false for single element")
		}
	})

	t.Run("all same", func(t *testing.T) {
		numbers := []int{5, 5, 5, 5}
		if !HasDuplicates(numbers) {
			t.Error("HasDuplicates() should return true when all elements are same")
		}
	})

	t.Run("strings with duplicates", func(t *testing.T) {
		words := []string{"apple", "banana", "apple", "cherry"}
		if !HasDuplicates(words) {
			t.Error("HasDuplicates() should return true for strings with duplicates")
		}
	})

	t.Run("strings without duplicates", func(t *testing.T) {
		words := []string{"apple", "banana", "cherry"}
		if HasDuplicates(words) {
			t.Error("HasDuplicates() should return false for strings without duplicates")
		}
	})
}

// BenchmarkHasDuplicates benchmarks the HasDuplicates function.
// BenchmarkHasDuplicates는 HasDuplicates 함수를 벤치마크합니다.
func BenchmarkHasDuplicates(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HasDuplicates(numbers)
	}
}
