package maputil

import (
	"testing"
)

func TestReduce(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	sum := Reduce(m, 0, func(acc int, k string, v int) int {
		return acc + v
	})

	if sum != 6 {
		t.Errorf("Expected sum 6, got %d", sum)
	}
}

func TestSum(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	total := Sum(m)

	if total != 6 {
		t.Errorf("Expected 6, got %d", total)
	}

	// Empty map / 빈 맵
	empty := map[string]int{}
	if Sum(empty) != 0 {
		t.Error("Sum of empty map should be 0")
	}
}

func TestMin(t *testing.T) {
	m := map[string]int{"a": 3, "b": 1, "c": 2}

	key, value, ok := Min(m)

	if !ok {
		t.Error("Expected ok to be true")
	}

	if value != 1 {
		t.Errorf("Expected min value 1, got %d", value)
	}

	if key != "b" {
		t.Errorf("Expected key 'b', got '%s'", key)
	}

	// Empty map / 빈 맵
	empty := map[string]int{}
	_, _, ok = Min(empty)
	if ok {
		t.Error("Expected ok to be false for empty map")
	}
}

func TestMax(t *testing.T) {
	m := map[string]int{"a": 3, "b": 1, "c": 2}

	key, value, ok := Max(m)

	if !ok {
		t.Error("Expected ok to be true")
	}

	if value != 3 {
		t.Errorf("Expected max value 3, got %d", value)
	}

	if key != "a" {
		t.Errorf("Expected key 'a', got '%s'", key)
	}

	// Empty map / 빈 맵
	empty := map[string]int{}
	_, _, ok = Max(empty)
	if ok {
		t.Error("Expected ok to be false for empty map")
	}
}

func TestMinBy(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	m := map[string]User{
		"alice": {Name: "Alice", Age: 30},
		"bob":   {Name: "Bob", Age: 25},
		"carol": {Name: "Carol", Age: 35},
	}

	key, user, ok := MinBy(m, func(u User) float64 {
		return float64(u.Age)
	})

	if !ok {
		t.Error("Expected ok to be true")
	}

	if user.Age != 25 {
		t.Errorf("Expected age 25, got %d", user.Age)
	}

	if key != "bob" {
		t.Errorf("Expected key 'bob', got '%s'", key)
	}
}

func TestMaxBy(t *testing.T) {
	type User struct {
		Name  string
		Score int
	}

	m := map[string]User{
		"alice": {Name: "Alice", Score: 95},
		"bob":   {Name: "Bob", Score: 88},
		"carol": {Name: "Carol", Score: 92},
	}

	key, user, ok := MaxBy(m, func(u User) float64 {
		return float64(u.Score)
	})

	if !ok {
		t.Error("Expected ok to be true")
	}

	if user.Score != 95 {
		t.Errorf("Expected score 95, got %d", user.Score)
	}

	if key != "alice" {
		t.Errorf("Expected key 'alice', got '%s'", key)
	}
}

func TestAverage(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	avg := Average(m)

	if avg != 2.5 {
		t.Errorf("Expected average 2.5, got %f", avg)
	}

	// Empty map / 빈 맵
	empty := map[string]int{}
	if Average(empty) != 0 {
		t.Error("Average of empty map should be 0")
	}
}

func TestGroupBy(t *testing.T) {
	type User struct {
		Name string
		City string
	}

	users := []User{
		{Name: "Alice", City: "Seoul"},
		{Name: "Bob", City: "Seoul"},
		{Name: "Charlie", City: "Busan"},
	}

	byCity := GroupBy[string, User, string](users, func(u User) string {
		return u.City
	})

	if len(byCity) != 2 {
		t.Errorf("Expected 2 cities, got %d", len(byCity))
	}

	if len(byCity["Seoul"]) != 2 {
		t.Errorf("Expected 2 users in Seoul, got %d", len(byCity["Seoul"]))
	}

	if len(byCity["Busan"]) != 1 {
		t.Errorf("Expected 1 user in Busan, got %d", len(byCity["Busan"]))
	}
}

func TestCountBy(t *testing.T) {
	type User struct {
		Name string
		City string
	}

	users := []User{
		{Name: "Alice", City: "Seoul"},
		{Name: "Bob", City: "Seoul"},
		{Name: "Charlie", City: "Busan"},
	}

	counts := CountBy[string, User, string](users, func(u User) string {
		return u.City
	})

	if counts["Seoul"] != 2 {
		t.Errorf("Expected 2 users in Seoul, got %d", counts["Seoul"])
	}

	if counts["Busan"] != 1 {
		t.Errorf("Expected 1 user in Busan, got %d", counts["Busan"])
	}
}

// Benchmark tests / 벤치마크 테스트

func BenchmarkSum(b *testing.B) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Sum(m)
	}
}

func BenchmarkAverage(b *testing.B) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Average(m)
	}
}

func BenchmarkGroupBy(b *testing.B) {
	type User struct {
		Name string
		City string
	}
	users := []User{
		{Name: "Alice", City: "Seoul"},
		{Name: "Bob", City: "Seoul"},
		{Name: "Charlie", City: "Busan"},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GroupBy[string, User, string](users, func(u User) string { return u.City })
	}
}
