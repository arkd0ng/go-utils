package maputil

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// Keys returns all keys from the map as a slice.
// Keys는 맵의 모든 키를 슬라이스로 반환합니다.
//
// The order of keys is not guaranteed (maps are unordered).
// 키의 순서는 보장되지 않습니다 (맵은 순서가 없음).
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	keys := maputil.Keys(m) // []string{"a", "b", "c"} (order may vary)
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Values returns all values from the map as a slice.
// Values는 맵의 모든 값을 슬라이스로 반환합니다.
//
// The order of values is not guaranteed (maps are unordered).
// 값의 순서는 보장되지 않습니다 (맵은 순서가 없음).
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	values := maputil.Values(m) // []int{1, 2, 3} (order may vary)
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// Entries returns all key-value pairs from the map as a slice of Entry structs.
// Entries는 맵의 모든 키-값 쌍을 Entry 구조체의 슬라이스로 반환합니다.
//
// The order of entries is not guaranteed (maps are unordered).
// 항목의 순서는 보장되지 않습니다 (맵은 순서가 없음).
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2}
//	entries := maputil.Entries(m)
//	// []Entry[string, int]{
//	//     {Key: "a", Value: 1},
//	//     {Key: "b", Value: 2},
//	// } (order may vary)
func Entries[K comparable, V any](m map[K]V) []Entry[K, V] {
	entries := make([]Entry[K, V], 0, len(m))
	for k, v := range m {
		entries = append(entries, Entry[K, V]{Key: k, Value: v})
	}
	return entries
}

// FromEntries creates a map from a slice of Entry structs.
// FromEntries는 Entry 구조체의 슬라이스로부터 맵을 생성합니다.
//
// If duplicate keys exist, the last entry wins.
// 중복 키가 있으면 마지막 항목이 우선합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	entries := []Entry[string, int]{
//	    {Key: "a", Value: 1},
//	    {Key: "b", Value: 2},
//	}
//	m := maputil.FromEntries(entries) // map[string]int{"a": 1, "b": 2}
func FromEntries[K comparable, V any](entries []Entry[K, V]) map[K]V {
	result := make(map[K]V, len(entries))
	for _, entry := range entries {
		result[entry.Key] = entry.Value
	}
	return result
}

// FromSlice creates a map from a slice using a key extraction function.
// FromSlice는 키 추출 함수를 사용하여 슬라이스로부터 맵을 생성합니다.
//
// If duplicate keys exist, the last value wins.
// 중복 키가 있으면 마지막 값이 우선합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	type User struct { ID int; Name string }
//	users := []User{{ID: 1, Name: "Alice"}, {ID: 2, Name: "Bob"}}
//	m := maputil.FromSlice(users, func(u User) int {
//	    return u.ID
//	}) // map[int]User{1: {ID: 1, Name: "Alice"}, 2: {ID: 2, Name: "Bob"}}
func FromSlice[K comparable, V any](slice []V, fn func(V) K) map[K]V {
	result := make(map[K]V, len(slice))
	for _, item := range slice {
		key := fn(item)
		result[key] = item
	}
	return result
}

// FromSliceBy creates a map from a slice using key and value extraction functions.
// FromSliceBy는 키와 값 추출 함수를 사용하여 슬라이스로부터 맵을 생성합니다.
//
// If duplicate keys exist, the last value wins.
// 중복 키가 있으면 마지막 값이 우선합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	type User struct { ID int; Name string; Age int }
//	users := []User{{ID: 1, Name: "Alice", Age: 25}, {ID: 2, Name: "Bob", Age: 30}}
//	m := maputil.FromSliceBy(users,
//	    func(u User) int { return u.ID },
//	    func(u User) string { return u.Name },
//	) // map[int]string{1: "Alice", 2: "Bob"}
func FromSliceBy[K comparable, V any, R any](slice []V, keyFn func(V) K, valueFn func(V) R) map[K]R {
	result := make(map[K]R, len(slice))
	for _, item := range slice {
		key := keyFn(item)
		value := valueFn(item)
		result[key] = value
	}
	return result
}

// ToSlice converts a map to a slice using a transformation function.
// ToSlice는 변환 함수를 사용하여 맵을 슬라이스로 변환합니다.
//
// The order of elements is not guaranteed (maps are unordered).
// 요소의 순서는 보장되지 않습니다 (맵은 순서가 없음).
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	slice := maputil.ToSlice(m, func(k string, v int) string {
//	    return fmt.Sprintf("%s=%d", k, v)
//	}) // []string{"a=1", "b=2", "c=3"} (order may vary)
func ToSlice[K comparable, V any, R any](m map[K]V, fn func(K, V) R) []R {
	result := make([]R, 0, len(m))
	for k, v := range m {
		result = append(result, fn(k, v))
	}
	return result
}

// ToJSON converts a map to a JSON string.
// ToJSON은 맵을 JSON 문자열로 변환합니다.
//
// Returns an error if the map cannot be marshaled to JSON.
// 맵을 JSON으로 마샬링할 수 없으면 에러를 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	json, err := maputil.ToJSON(m) // `{"a":1,"b":2,"c":3}`
func ToJSON[K comparable, V any](m map[K]V) (string, error) {
	bytes, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// FromJSON parses a JSON string into a map.
// FromJSON은 JSON 문자열을 맵으로 파싱합니다.
//
// Returns an error if the JSON string cannot be unmarshaled.
// JSON 문자열을 언마샬링할 수 없으면 에러를 반환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	var m map[string]int
//	err := maputil.FromJSON(`{"a":1,"b":2,"c":3}`, &m)
//	// m = map[string]int{"a": 1, "b": 2, "c": 3}
func FromJSON[K comparable, V any](jsonStr string, m *map[K]V) error {
	return json.Unmarshal([]byte(jsonStr), m)
}

// ToYAML converts a map to a YAML string.
// ToYAML은 맵을 YAML 문자열로 변환합니다.
//
// This function serializes a map into YAML format using the gopkg.in/yaml.v3 package.
// Returns an error if the map cannot be marshaled to YAML.
//
// 이 함수는 gopkg.in/yaml.v3 패키지를 사용하여 맵을 YAML 형식으로 직렬화합니다.
// 맵을 YAML로 마샬링할 수 없으면 에러를 반환합니다.
//
// Time Complexity
// 시간 복잡도: O(n)
// Space Complexity
// 공간 복잡도: O(n)
//
// Parameters
// 매개변수:
// - m: The map to convert
// 변환할 맵
//
// Returns
// 반환값:
// - string: YAML string representation
// YAML 문자열 표현
// - error: Error if marshaling fails
// 마샬링 실패 시 에러
//
// Example
// 예제:
//
//	config := map[string]interface{}{
//		"server": map[string]interface{}{
//			"host": "localhost",
//			"port": 8080,
//		},
//		"database": map[string]interface{}{
//			"host": "localhost",
//			"port": 5432,
//		},
//	}
//	yamlStr, err := maputil.ToYAML(config)
//	// yamlStr:
//	// server:
//	//   host: localhost
//	//   port: 8080
//	// database:
//	//   host: localhost
//	//   port: 5432
//
// Use Case
// 사용 사례:
// - Configuration file generation
// 설정 파일 생성
// - API response formatting
// API 응답 포맷팅
// - Data serialization
// 데이터 직렬화
// - Config export
// 설정 내보내기
func ToYAML[K comparable, V any](m map[K]V) (string, error) {
	bytes, err := yaml.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// FromYAML parses a YAML string into a map.
// FromYAML은 YAML 문자열을 맵으로 파싱합니다.
//
// This function deserializes a YAML string into a map[string]interface{}.
// Returns an error if the YAML string cannot be unmarshaled.
//
// 이 함수는 YAML 문자열을 map[string]interface{}로 역직렬화합니다.
// YAML 문자열을 언마샬링할 수 없으면 에러를 반환합니다.
//
// Time Complexity
// 시간 복잡도: O(n)
// Space Complexity
// 공간 복잡도: O(n)
//
// Parameters
// 매개변수:
// - yamlStr: YAML string to parse
// 파싱할 YAML 문자열
//
// Returns
// 반환값:
// - map[string]interface{}: Parsed map
// 파싱된 맵
// - error: Error if unmarshaling fails
// 언마샬링 실패 시 에러
//
// Example
// 예제:
//
//	yamlStr := `
//	server:
//	  host: localhost
//	  port: 8080
//	database:
//	  host: localhost
//	  port: 5432
//	`
//	config, err := maputil.FromYAML(yamlStr)
//	// config = map[string]interface{}{
//	//   "server": map[string]interface{}{
//	//     "host": "localhost",
//	//     "port": 8080,
//	//   },
//	//   "database": map[string]interface{}{
//	//     "host": "localhost",
//	//     "port": 5432,
//	//   },
//	// }
//
// Use Case
// 사용 사례:
// - Configuration file loading
// 설정 파일 로딩
// - API request parsing
// API 요청 파싱
// - Data deserialization
// 데이터 역직렬화
// - Config import
// 설정 가져오기
func FromYAML(yamlStr string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := yaml.Unmarshal([]byte(yamlStr), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
