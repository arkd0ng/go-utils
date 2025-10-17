package sliceutil

import (
	"strings"
	"testing"
)

// TestForEach tests the ForEach function.
// TestForEach는 ForEach 함수를 테스트합니다.
func TestForEach(t *testing.T) {
	t.Run("execute for each element", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		sum := 0
		ForEach(numbers, func(n int) {
			sum += n
		})
		expected := 15
		if sum != expected {
			t.Errorf("ForEach() sum = %d, want %d", sum, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		count := 0
		ForEach(numbers, func(n int) {
			count++
		})
		if count != 0 {
			t.Errorf("ForEach() with empty slice should not execute function")
		}
	})

	t.Run("side effects", func(t *testing.T) {
		words := []string{"hello", "world"}
		result := []string{}
		ForEach(words, func(word string) {
			result = append(result, strings.ToUpper(word))
		})
		expected := []string{"HELLO", "WORLD"}
		if !Equal(result, expected) {
			t.Errorf("ForEach() side effects = %v, want %v", result, expected)
		}
	})
}

// BenchmarkForEach benchmarks the ForEach function.
// BenchmarkForEach는 ForEach 함수를 벤치마크합니다.
func BenchmarkForEach(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ForEach(numbers, func(n int) {
			_ = n * 2
		})
	}
}

// TestForEachIndexed tests the ForEachIndexed function.
// TestForEachIndexed는 ForEachIndexed 함수를 테스트합니다.
func TestForEachIndexed(t *testing.T) {
	t.Run("execute with index", func(t *testing.T) {
		words := []string{"apple", "banana", "cherry"}
		indices := []int{}
		values := []string{}
		ForEachIndexed(words, func(i int, word string) {
			indices = append(indices, i)
			values = append(values, word)
		})
		expectedIndices := []int{0, 1, 2}
		if !Equal(indices, expectedIndices) {
			t.Errorf("ForEachIndexed() indices = %v, want %v", indices, expectedIndices)
		}
		if !Equal(values, words) {
			t.Errorf("ForEachIndexed() values = %v, want %v", values, words)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		count := 0
		ForEachIndexed(numbers, func(i int, n int) {
			count++
		})
		if count != 0 {
			t.Errorf("ForEachIndexed() with empty slice should not execute function")
		}
	})
}

// BenchmarkForEachIndexed benchmarks the ForEachIndexed function.
// BenchmarkForEachIndexed는 ForEachIndexed 함수를 벤치마크합니다.
func BenchmarkForEachIndexed(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ForEachIndexed(numbers, func(idx int, n int) {
			_ = idx + n
		})
	}
}

// TestJoin tests the Join function.
// TestJoin은 Join 함수를 테스트합니다.
func TestJoin(t *testing.T) {
	t.Run("join integers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Join(numbers, ", ")
		expected := "1, 2, 3, 4, 5"
		if result != expected {
			t.Errorf("Join() = %v, want %v", result, expected)
		}
	})

	t.Run("join strings", func(t *testing.T) {
		words := []string{"apple", "banana", "cherry"}
		result := Join(words, "-")
		expected := "apple-banana-cherry"
		if result != expected {
			t.Errorf("Join() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		result := Join(numbers, ", ")
		if result != "" {
			t.Errorf("Join() with empty slice should return empty string, got %v", result)
		}
	})

	t.Run("single element", func(t *testing.T) {
		numbers := []int{42}
		result := Join(numbers, ", ")
		expected := "42"
		if result != expected {
			t.Errorf("Join() = %v, want %v", result, expected)
		}
	})

	t.Run("empty separator", func(t *testing.T) {
		words := []string{"a", "b", "c"}
		result := Join(words, "")
		expected := "abc"
		if result != expected {
			t.Errorf("Join() = %v, want %v", result, expected)
		}
	})
}

// BenchmarkJoin benchmarks the Join function.
// BenchmarkJoin은 Join 함수를 벤치마크합니다.
func BenchmarkJoin(b *testing.B) {
	numbers := make([]int, 100)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(numbers, ", ")
	}
}

// TestClone tests the Clone function.
// TestClone은 Clone 함수를 테스트합니다.
func TestClone(t *testing.T) {
	t.Run("clone slice", func(t *testing.T) {
		original := []int{1, 2, 3, 4, 5}
		cloned := Clone(original)
		if !Equal(original, cloned) {
			t.Errorf("Clone() = %v, want %v", cloned, original)
		}
		// Modify cloned
		// 복제된 슬라이스 수정
		cloned[0] = 99
		if original[0] == 99 {
			t.Error("Clone() should create independent copy")
		}
	})

	t.Run("clone empty slice", func(t *testing.T) {
		original := []int{}
		cloned := Clone(original)
		if len(cloned) != 0 {
			t.Errorf("Clone() with empty slice should return empty slice")
		}
	})

	t.Run("clone nil slice", func(t *testing.T) {
		var original []int
		cloned := Clone(original)
		if cloned != nil {
			t.Errorf("Clone() with nil slice should return nil")
		}
	})

	t.Run("clone strings", func(t *testing.T) {
		original := []string{"apple", "banana", "cherry"}
		cloned := Clone(original)
		if !Equal(original, cloned) {
			t.Errorf("Clone() = %v, want %v", cloned, original)
		}
	})
}

// BenchmarkClone benchmarks the Clone function.
// BenchmarkClone은 Clone 함수를 벤치마크합니다.
func BenchmarkClone(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Clone(slice)
	}
}

// TestFill tests the Fill function.
// TestFill은 Fill 함수를 테스트합니다.
func TestFill(t *testing.T) {
	t.Run("fill with value", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		filled := Fill(slice, 0)
		expected := []int{0, 0, 0, 0, 0}
		if !Equal(filled, expected) {
			t.Errorf("Fill() = %v, want %v", filled, expected)
		}
		// Verify original is unchanged
		// 원본이 변경되지 않았는지 확인
		if slice[0] == 0 {
			t.Error("Fill() should not modify original slice")
		}
	})

	t.Run("fill strings", func(t *testing.T) {
		slice := []string{"a", "b", "c"}
		filled := Fill(slice, "x")
		expected := []string{"x", "x", "x"}
		if !Equal(filled, expected) {
			t.Errorf("Fill() = %v, want %v", filled, expected)
		}
	})

	t.Run("fill empty slice", func(t *testing.T) {
		slice := []int{}
		filled := Fill(slice, 42)
		if len(filled) != 0 {
			t.Errorf("Fill() with empty slice should return empty slice")
		}
	})

	t.Run("fill single element", func(t *testing.T) {
		slice := []int{99}
		filled := Fill(slice, 42)
		expected := []int{42}
		if !Equal(filled, expected) {
			t.Errorf("Fill() = %v, want %v", filled, expected)
		}
	})
}

// BenchmarkFill benchmarks the Fill function.
// BenchmarkFill은 Fill 함수를 벤치마크합니다.
func BenchmarkFill(b *testing.B) {
	slice := make([]int, 1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Fill(slice, 42)
	}
}

// TestInsert tests the Insert function.
// TestInsert은 Insert 함수를 테스트합니다.
func TestInsert(t *testing.T) {
	t.Run("insert in middle", func(t *testing.T) {
		slice := []int{1, 2, 5, 6}
		result := Insert(slice, 2, 3, 4)
		expected := []int{1, 2, 3, 4, 5, 6}
		if !Equal(result, expected) {
			t.Errorf("Insert() = %v, want %v", result, expected)
		}
	})

	t.Run("insert at beginning", func(t *testing.T) {
		slice := []int{2, 3, 4}
		result := Insert(slice, 0, 1)
		expected := []int{1, 2, 3, 4}
		if !Equal(result, expected) {
			t.Errorf("Insert() = %v, want %v", result, expected)
		}
	})

	t.Run("insert at end", func(t *testing.T) {
		slice := []int{1, 2, 3}
		result := Insert(slice, 3, 4, 5)
		expected := []int{1, 2, 3, 4, 5}
		if !Equal(result, expected) {
			t.Errorf("Insert() = %v, want %v", result, expected)
		}
	})

	t.Run("insert with negative index", func(t *testing.T) {
		slice := []int{2, 3, 4}
		result := Insert(slice, -1, 1)
		expected := []int{1, 2, 3, 4}
		if !Equal(result, expected) {
			t.Errorf("Insert() with negative index should insert at beginning, got %v", result)
		}
	})

	t.Run("insert with large index", func(t *testing.T) {
		slice := []int{1, 2, 3}
		result := Insert(slice, 100, 4, 5)
		expected := []int{1, 2, 3, 4, 5}
		if !Equal(result, expected) {
			t.Errorf("Insert() with large index should append, got %v", result)
		}
	})

	t.Run("insert into empty slice", func(t *testing.T) {
		slice := []int{}
		result := Insert(slice, 0, 1, 2, 3)
		expected := []int{1, 2, 3}
		if !Equal(result, expected) {
			t.Errorf("Insert() = %v, want %v", result, expected)
		}
	})

	t.Run("insert no items", func(t *testing.T) {
		slice := []int{1, 2, 3}
		result := Insert(slice, 1)
		if !Equal(result, slice) {
			t.Errorf("Insert() with no items should return copy of original")
		}
	})
}

// BenchmarkInsert benchmarks the Insert function.
// BenchmarkInsert은 Insert 함수를 벤치마크합니다.
func BenchmarkInsert(b *testing.B) {
	slice := make([]int, 100)
	for i := range slice {
		slice[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Insert(slice, 50, 999)
	}
}

// TestRemove tests the Remove function.
// TestRemove은 Remove 함수를 테스트합니다.
func TestRemove(t *testing.T) {
	t.Run("remove from middle", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		result := Remove(slice, 2)
		expected := []int{1, 2, 4, 5}
		if !Equal(result, expected) {
			t.Errorf("Remove() = %v, want %v", result, expected)
		}
	})

	t.Run("remove first element", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		result := Remove(slice, 0)
		expected := []int{2, 3, 4, 5}
		if !Equal(result, expected) {
			t.Errorf("Remove() = %v, want %v", result, expected)
		}
	})

	t.Run("remove last element", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		result := Remove(slice, 4)
		expected := []int{1, 2, 3, 4}
		if !Equal(result, expected) {
			t.Errorf("Remove() = %v, want %v", result, expected)
		}
	})

	t.Run("remove with negative index", func(t *testing.T) {
		slice := []int{1, 2, 3}
		result := Remove(slice, -1)
		if !Equal(result, slice) {
			t.Error("Remove() with negative index should return copy of original")
		}
	})

	t.Run("remove with large index", func(t *testing.T) {
		slice := []int{1, 2, 3}
		result := Remove(slice, 100)
		if !Equal(result, slice) {
			t.Error("Remove() with large index should return copy of original")
		}
	})

	t.Run("original unchanged", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		_ = Remove(slice, 2)
		expected := []int{1, 2, 3, 4, 5}
		if !Equal(slice, expected) {
			t.Error("Remove() should not modify original slice")
		}
	})
}

// BenchmarkRemove benchmarks the Remove function.
// BenchmarkRemove은 Remove 함수를 벤치마크합니다.
func BenchmarkRemove(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Remove(slice, 500)
	}
}

// TestRemoveAll tests the RemoveAll function.
// TestRemoveAll은 RemoveAll 함수를 테스트합니다.
func TestRemoveAll(t *testing.T) {
	t.Run("remove all occurrences", func(t *testing.T) {
		slice := []int{1, 2, 3, 2, 4, 2, 5}
		result := RemoveAll(slice, 2)
		expected := []int{1, 3, 4, 5}
		if !Equal(result, expected) {
			t.Errorf("RemoveAll() = %v, want %v", result, expected)
		}
	})

	t.Run("remove all strings", func(t *testing.T) {
		words := []string{"apple", "banana", "apple", "cherry", "apple"}
		result := RemoveAll(words, "apple")
		expected := []string{"banana", "cherry"}
		if !Equal(result, expected) {
			t.Errorf("RemoveAll() = %v, want %v", result, expected)
		}
	})

	t.Run("remove non-existent item", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		result := RemoveAll(slice, 99)
		if !Equal(result, slice) {
			t.Errorf("RemoveAll() with non-existent item should return copy of original")
		}
	})

	t.Run("remove all elements", func(t *testing.T) {
		slice := []int{5, 5, 5, 5}
		result := RemoveAll(slice, 5)
		if len(result) != 0 {
			t.Errorf("RemoveAll() should return empty slice when all elements removed")
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		slice := []int{}
		result := RemoveAll(slice, 1)
		if len(result) != 0 {
			t.Errorf("RemoveAll() with empty slice should return empty slice")
		}
	})
}

// BenchmarkRemoveAll benchmarks the RemoveAll function.
// BenchmarkRemoveAll은 RemoveAll 함수를 벤치마크합니다.
func BenchmarkRemoveAll(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i % 10
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveAll(slice, 5)
	}
}

// TestShuffle tests the Shuffle function.
// TestShuffle은 Shuffle 함수를 테스트합니다.
func TestShuffle(t *testing.T) {
	t.Run("shuffle changes order", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		shuffled := Shuffle(slice)

		// Should have same length
		// 같은 길이여야 함
		if len(shuffled) != len(slice) {
			t.Errorf("Shuffle() length = %d, want %d", len(shuffled), len(slice))
		}

		// Should contain same elements
		// 같은 요소를 포함해야 함
		for _, v := range slice {
			if !Contains(shuffled, v) {
				t.Errorf("Shuffle() missing element %d", v)
			}
		}
	})

	t.Run("original unchanged", func(t *testing.T) {
		original := []int{1, 2, 3, 4, 5}
		expected := Clone(original)
		_ = Shuffle(original)
		if !Equal(original, expected) {
			t.Error("Shuffle() should not modify original slice")
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		slice := []int{}
		shuffled := Shuffle(slice)
		if len(shuffled) != 0 {
			t.Errorf("Shuffle() with empty slice should return empty slice")
		}
	})

	t.Run("single element", func(t *testing.T) {
		slice := []int{42}
		shuffled := Shuffle(slice)
		expected := []int{42}
		if !Equal(shuffled, expected) {
			t.Errorf("Shuffle() with single element = %v, want %v", shuffled, expected)
		}
	})
}

// BenchmarkShuffle benchmarks the Shuffle function.
// BenchmarkShuffle은 Shuffle 함수를 벤치마크합니다.
func BenchmarkShuffle(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Shuffle(slice)
	}
}

// TestZip tests the Zip function.
// TestZip은 Zip 함수를 테스트합니다.
func TestZip(t *testing.T) {
	t.Run("zip two slices", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		words := []string{"one", "two", "three"}
		zipped := Zip(numbers, words)

		if len(zipped) != 3 {
			t.Errorf("Zip() length = %d, want 3", len(zipped))
		}

		if zipped[0][0].(int) != 1 || zipped[0][1].(string) != "one" {
			t.Errorf("Zip() first pair = [%v, %v], want [1, one]", zipped[0][0], zipped[0][1])
		}
		if zipped[1][0].(int) != 2 || zipped[1][1].(string) != "two" {
			t.Errorf("Zip() second pair = [%v, %v], want [2, two]", zipped[1][0], zipped[1][1])
		}
	})

	t.Run("zip different lengths", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		words := []string{"one", "two"}
		zipped := Zip(numbers, words)

		// Should use minimum length
		// 최소 길이를 사용해야 함
		if len(zipped) != 2 {
			t.Errorf("Zip() with different lengths should use minimum, got %d", len(zipped))
		}
	})

	t.Run("zip empty slices", func(t *testing.T) {
		numbers := []int{}
		words := []string{}
		zipped := Zip(numbers, words)

		if len(zipped) != 0 {
			t.Errorf("Zip() with empty slices should return empty")
		}
	})

	t.Run("zip one empty", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		words := []string{}
		zipped := Zip(numbers, words)

		if len(zipped) != 0 {
			t.Errorf("Zip() with one empty slice should return empty")
		}
	})
}

// BenchmarkZip benchmarks the Zip function.
// BenchmarkZip은 Zip 함수를 벤치마크합니다.
func BenchmarkZip(b *testing.B) {
	slice1 := make([]int, 1000)
	slice2 := make([]string, 1000)
	for i := range slice1 {
		slice1[i] = i
		slice2[i] = "value"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Zip(slice1, slice2)
	}
}

// TestUnzip tests the Unzip function.
// TestUnzip은 Unzip 함수를 테스트합니다.
func TestUnzip(t *testing.T) {
	t.Run("unzip pairs", func(t *testing.T) {
		zipped := [][2]any{{1, "one"}, {2, "two"}, {3, "three"}}
		numbers, words := Unzip[int, string](zipped)

		expectedNumbers := []int{1, 2, 3}
		expectedWords := []string{"one", "two", "three"}

		if !Equal(numbers, expectedNumbers) {
			t.Errorf("Unzip() numbers = %v, want %v", numbers, expectedNumbers)
		}
		if !Equal(words, expectedWords) {
			t.Errorf("Unzip() words = %v, want %v", words, expectedWords)
		}
	})

	t.Run("unzip empty", func(t *testing.T) {
		zipped := [][2]any{}
		numbers, words := Unzip[int, string](zipped)

		if len(numbers) != 0 || len(words) != 0 {
			t.Errorf("Unzip() with empty should return empty slices")
		}
	})

	t.Run("zip and unzip roundtrip", func(t *testing.T) {
		originalNumbers := []int{1, 2, 3, 4, 5}
		originalWords := []string{"a", "b", "c", "d", "e"}

		zipped := Zip(originalNumbers, originalWords)
		numbers, words := Unzip[int, string](zipped)

		if !Equal(numbers, originalNumbers) {
			t.Errorf("Roundtrip numbers = %v, want %v", numbers, originalNumbers)
		}
		if !Equal(words, originalWords) {
			t.Errorf("Roundtrip words = %v, want %v", words, originalWords)
		}
	})
}

// BenchmarkUnzip benchmarks the Unzip function.
// BenchmarkUnzip은 Unzip 함수를 벤치마크합니다.
func BenchmarkUnzip(b *testing.B) {
	zipped := make([][2]any, 1000)
	for i := range zipped {
		zipped[i] = [2]any{i, "value"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Unzip[int, string](zipped)
	}
}

// TestWindow tests the Window function.
// TestWindow는 Window 함수를 테스트합니다.
func TestWindow(t *testing.T) {
	t.Run("window size 2", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Window(numbers, 2)
		expected := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
		if len(result) != len(expected) {
			t.Fatalf("Window() length = %v, want %v", len(result), len(expected))
		}
		for i := range expected {
			if !Equal(result[i], expected[i]) {
				t.Errorf("Window() window %d = %v, want %v", i, result[i], expected[i])
			}
		}
	})

	t.Run("window size 3", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Window(numbers, 3)
		expected := [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}
		if len(result) != len(expected) {
			t.Fatalf("Window() length = %v, want %v", len(result), len(expected))
		}
		for i := range expected {
			if !Equal(result[i], expected[i]) {
				t.Errorf("Window() window %d = %v, want %v", i, result[i], expected[i])
			}
		}
	})

	t.Run("window size equals slice length", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Window(numbers, 3)
		expected := [][]int{{1, 2, 3}}
		if len(result) != 1 {
			t.Fatalf("Window() should return 1 window, got %v", len(result))
		}
		if !Equal(result[0], expected[0]) {
			t.Errorf("Window() = %v, want %v", result[0], expected[0])
		}
	})

	t.Run("window size 0", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Window(numbers, 0)
		if len(result) != 0 {
			t.Errorf("Window() with size 0 should return empty slice, got %v", result)
		}
	})

	t.Run("window size greater than slice", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Window(numbers, 5)
		if len(result) != 0 {
			t.Errorf("Window() with size > length should return empty slice, got %v", result)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		result := Window(numbers, 2)
		if len(result) != 0 {
			t.Errorf("Window() with empty slice should return empty slice")
		}
	})
}

// TestTap tests the Tap function.
// TestTap는 Tap 함수를 테스트합니다.
func TestTap(t *testing.T) {
	t.Run("tap executes function", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		executed := false
		result := Tap(numbers, func(s []int) {
			executed = true
		})
		if !executed {
			t.Error("Tap() should execute the provided function")
		}
		if !Equal(result, numbers) {
			t.Errorf("Tap() should return original slice unchanged")
		}
	})

	t.Run("tap with side effect", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		var sum int
		result := Tap(numbers, func(s []int) {
			for _, n := range s {
				sum += n
			}
		})
		if sum != 15 {
			t.Errorf("Tap() side effect: sum = %v, want 15", sum)
		}
		if !Equal(result, numbers) {
			t.Errorf("Tap() should return original slice unchanged")
		}
	})

	t.Run("tap returns same slice reference", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Tap(numbers, func(s []int) {})
		// Check if it's the same slice (not a copy)
		numbers[0] = 999
		if result[0] != 999 {
			t.Error("Tap() should return the same slice reference, not a copy")
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		executed := false
		result := Tap(numbers, func(s []int) {
			executed = true
		})
		if !executed {
			t.Error("Tap() should execute function even for empty slice")
		}
		if len(result) != 0 {
			t.Error("Tap() with empty slice should return empty slice")
		}
	})
}
