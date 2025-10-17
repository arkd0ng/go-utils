package websvrutil

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"sort"
	"testing"
)

// TestContextStorageSetGet tests Set and Get methods
// Set과 Get 메서드 테스트
func TestContextStorageSetGet(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	ctx := NewContext(w, req)

	// Set and get string
	// 문자열 설정 및 가져오기
	ctx.Set("name", "John Doe")
	value, exists := ctx.Get("name")
	if !exists {
		t.Error("Expected key 'name' to exist")
	}
	if value != "John Doe" {
		t.Errorf("Expected 'John Doe', got %v", value)
	}

	// Get non-existent key
	// 존재하지 않는 키 가져오기
	_, exists = ctx.Get("nonexistent")
	if exists {
		t.Error("Expected key 'nonexistent' to not exist")
	}
}

// TestContextStorageMustGet tests MustGet method
// MustGet 메서드 테스트
func TestContextStorageMustGet(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	ctx := NewContext(w, req)

	ctx.Set("key", "value")
	value := ctx.MustGet("key")
	if value != "value" {
		t.Errorf("Expected 'value', got %v", value)
	}
}

// TestContextStorageMustGetPanic tests MustGet panic
// MustGet 패닉 테스트
func TestContextStorageMustGetPanic(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	ctx := NewContext(w, req)

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for non-existent key")
		}
	}()

	ctx.MustGet("nonexistent")
}

// TestContextGetString tests GetString method
// GetString 메서드 테스트
func TestContextGetString(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	ctx := NewContext(w, req)

	ctx.Set("name", "Alice")
	name := ctx.GetString("name")
	if name != "Alice" {
		t.Errorf("Expected 'Alice', got %q", name)
	}

	// Get non-existent key
	// 존재하지 않는 키
	empty := ctx.GetString("nonexistent")
	if empty != "" {
		t.Errorf("Expected empty string, got %q", empty)
	}

	// Get wrong type
	// 잘못된 타입
	ctx.Set("number", 123)
	str := ctx.GetString("number")
	if str != "" {
		t.Errorf("Expected empty string for wrong type, got %q", str)
	}
}

// TestContextGetInt tests GetInt method
// GetInt 메서드 테스트
func TestContextGetInt(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	ctx := NewContext(w, req)

	ctx.Set("age", 25)
	age := ctx.GetInt("age")
	if age != 25 {
		t.Errorf("Expected 25, got %d", age)
	}

	// Get non-existent key
	// 존재하지 않는 키
	zero := ctx.GetInt("nonexistent")
	if zero != 0 {
		t.Errorf("Expected 0, got %d", zero)
	}

	// Get wrong type
	// 잘못된 타입
	ctx.Set("str", "text")
	num := ctx.GetInt("str")
	if num != 0 {
		t.Errorf("Expected 0 for wrong type, got %d", num)
	}
}

// TestContextGetBool tests GetBool method
// GetBool 메서드 테스트
func TestContextGetBool(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	ctx := NewContext(w, req)

	ctx.Set("active", true)
	active := ctx.GetBool("active")
	if !active {
		t.Error("Expected true, got false")
	}

	ctx.Set("inactive", false)
	inactive := ctx.GetBool("inactive")
	if inactive {
		t.Error("Expected false, got true")
	}

	// Get non-existent key
	// 존재하지 않는 키
	def := ctx.GetBool("nonexistent")
	if def {
		t.Error("Expected false for non-existent key")
	}
}

// TestContextGetInt64 tests GetInt64 method
// GetInt64 메서드 테스트
func TestContextGetInt64(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	ctx := NewContext(w, req)

	var largeNum int64 = 9223372036854775807
	ctx.Set("large", largeNum)
	result := ctx.GetInt64("large")
	if result != largeNum {
		t.Errorf("Expected %d, got %d", largeNum, result)
	}

	// Get non-existent key
	// 존재하지 않는 키
	zero := ctx.GetInt64("nonexistent")
	if zero != 0 {
		t.Errorf("Expected 0, got %d", zero)
	}
}

// TestContextGetFloat64 tests GetFloat64 method
// GetFloat64 메서드 테스트
func TestContextGetFloat64(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	ctx := NewContext(w, req)

	ctx.Set("price", 19.99)
	price := ctx.GetFloat64("price")
	if price != 19.99 {
		t.Errorf("Expected 19.99, got %f", price)
	}

	// Get non-existent key
	// 존재하지 않는 키
	zero := ctx.GetFloat64("nonexistent")
	if zero != 0.0 {
		t.Errorf("Expected 0.0, got %f", zero)
	}
}

// TestContextGetStringSlice tests GetStringSlice method
// GetStringSlice 메서드 테스트
func TestContextGetStringSlice(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	ctx := NewContext(w, req)

	tags := []string{"golang", "web", "api"}
	ctx.Set("tags", tags)
	result := ctx.GetStringSlice("tags")
	if !reflect.DeepEqual(result, tags) {
		t.Errorf("Expected %v, got %v", tags, result)
	}

	// Get non-existent key
	// 존재하지 않는 키
	nilSlice := ctx.GetStringSlice("nonexistent")
	if nilSlice != nil {
		t.Errorf("Expected nil, got %v", nilSlice)
	}
}

// TestContextGetStringMap tests GetStringMap method
// GetStringMap 메서드 테스트
func TestContextGetStringMap(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	ctx := NewContext(w, req)

	user := map[string]interface{}{
		"id":   123,
		"name": "Bob",
		"role": "admin",
	}
	ctx.Set("user", user)
	result := ctx.GetStringMap("user")
	if !reflect.DeepEqual(result, user) {
		t.Errorf("Expected %v, got %v", user, result)
	}

	// Get non-existent key
	// 존재하지 않는 키
	nilMap := ctx.GetStringMap("nonexistent")
	if nilMap != nil {
		t.Errorf("Expected nil, got %v", nilMap)
	}
}

// TestContextExists tests Exists method
// Exists 메서드 테스트
func TestContextExists(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	ctx := NewContext(w, req)

	// Check non-existent key
	// 존재하지 않는 키 확인
	if ctx.Exists("key") {
		t.Error("Expected key to not exist")
	}

	// Set and check
	// 설정 및 확인
	ctx.Set("key", "value")
	if !ctx.Exists("key") {
		t.Error("Expected key to exist")
	}

	// Check with nil value
	// nil 값 확인
	ctx.Set("nil", nil)
	if !ctx.Exists("nil") {
		t.Error("Expected nil key to exist")
	}
}

// TestContextDelete tests Delete method
// Delete 메서드 테스트
func TestContextDelete(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	ctx := NewContext(w, req)

	// Set and delete
	// 설정 및 삭제
	ctx.Set("temp", "value")
	if !ctx.Exists("temp") {
		t.Error("Expected temp to exist")
	}

	ctx.Delete("temp")
	if ctx.Exists("temp") {
		t.Error("Expected temp to not exist after delete")
	}

	// Delete non-existent key (should not panic)
	// 존재하지 않는 키 삭제 (패닉하지 않아야 함)
	ctx.Delete("nonexistent")
}

// TestContextKeys tests Keys method
// Keys 메서드 테스트
func TestContextKeys(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	ctx := NewContext(w, req)

	// Empty context
	// 빈 컨텍스트
	keys := ctx.Keys()
	if len(keys) != 0 {
		t.Errorf("Expected 0 keys, got %d", len(keys))
	}

	// Set multiple values
	// 여러 값 설정
	ctx.Set("key1", "value1")
	ctx.Set("key2", "value2")
	ctx.Set("key3", "value3")

	keys = ctx.Keys()
	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}

	// Check all keys are present
	// 모든 키가 있는지 확인
	sort.Strings(keys)
	expected := []string{"key1", "key2", "key3"}
	if !reflect.DeepEqual(keys, expected) {
		t.Errorf("Expected %v, got %v", expected, keys)
	}
}

// TestContextStorageConcurrency tests thread-safety
// 스레드 안전성 테스트
func TestContextStorageConcurrency(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	ctx := NewContext(w, req)

	done := make(chan bool)
	iterations := 100

	// Concurrent writes
	// 동시 쓰기
	for i := 0; i < iterations; i++ {
		go func(n int) {
			ctx.Set("key", n)
			done <- true
		}(i)
	}

	// Wait for all goroutines
	// 모든 고루틴 대기
	for i := 0; i < iterations; i++ {
		<-done
	}

	// Concurrent reads
	// 동시 읽기
	for i := 0; i < iterations; i++ {
		go func() {
			ctx.Get("key")
			done <- true
		}()
	}

	for i := 0; i < iterations; i++ {
		<-done
	}
}

// TestContextStorageTypes tests various data types
// 다양한 데이터 타입 테스트
func TestContextStorageTypes(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	ctx := NewContext(w, req)

	testCases := []struct {
		key   string
		value interface{}
	}{
		{"string", "hello"},
		{"int", 42},
		{"int64", int64(9223372036854775807)},
		{"float64", 3.14159},
		{"bool", true},
		{"slice", []string{"a", "b", "c"}},
		{"map", map[string]interface{}{"nested": "value"}},
		{"nil", nil},
	}

	for _, tc := range testCases {
		t.Run(tc.key, func(t *testing.T) {
			ctx.Set(tc.key, tc.value)
			value, exists := ctx.Get(tc.key)
			if !exists {
				t.Errorf("Expected key %q to exist", tc.key)
			}
			if !reflect.DeepEqual(value, tc.value) {
				t.Errorf("Expected %v, got %v", tc.value, value)
			}
		})
	}
}

// BenchmarkContextSet benchmarks Set method
// Set 메서드 벤치마크
func BenchmarkContextSet(b *testing.B) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	ctx := NewContext(w, req)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx.Set("key", "value")
	}
}

// BenchmarkContextGet benchmarks Get method
// Get 메서드 벤치마크
func BenchmarkContextGet(b *testing.B) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	ctx := NewContext(w, req)
	ctx.Set("key", "value")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx.Get("key")
	}
}

// BenchmarkContextGetString benchmarks GetString method
// GetString 메서드 벤치마크
func BenchmarkContextGetString(b *testing.B) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	ctx := NewContext(w, req)
	ctx.Set("key", "value")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx.GetString("key")
	}
}
