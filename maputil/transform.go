package maputil

import (
	"fmt"
	"strings"
)

// Map transforms map values using the provided function.
// Map은 제공된 함수를 사용하여 맵 값을 변환합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	result := maputil.Map(m, func(k string, v int) string {
//	    return fmt.Sprintf("%s=%d", k, v)
//	}) // map[string]string{"a": "a=1", "b": "b=2", "c": "c=3"}
func Map[K comparable, V any, R any](m map[K]V, fn func(K, V) R) map[K]R {
	result := make(map[K]R, len(m))
	for k, v := range m {
		result[k] = fn(k, v)
	}
	return result
}

// MapKeys transforms map keys using the provided function.
// MapKeys는 제공된 함수를 사용하여 맵 키를 변환합니다.
//
// If the function produces duplicate keys, the last value wins.
// 함수가 중복 키를 생성하면 마지막 값이 우선합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	result := maputil.MapKeys(m, func(k string, v int) string {
//	    return strings.ToUpper(k)
//	}) // map[string]int{"A": 1, "B": 2, "C": 3}
func MapKeys[K comparable, V any, R comparable](m map[K]V, fn func(K, V) R) map[R]V {
	result := make(map[R]V, len(m))
	for k, v := range m {
		newKey := fn(k, v)
		result[newKey] = v
	}
	return result
}

// MapValues transforms map values using the provided function (value-only version).
// MapValues는 제공된 함수를 사용하여 맵 값을 변환합니다 (값만 사용하는 버전).
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	result := maputil.MapValues(m, func(v int) int {
//	    return v * 2
//	}) // map[string]int{"a": 2, "b": 4, "c": 6}
func MapValues[K comparable, V any, R any](m map[K]V, fn func(V) R) map[K]R {
	result := make(map[K]R, len(m))
	for k, v := range m {
		result[k] = fn(v)
	}
	return result
}

// MapEntries transforms both keys and values using the provided function.
// MapEntries는 제공된 함수를 사용하여 키와 값을 모두 변환합니다.
//
// If the function produces duplicate keys, the last value wins.
// 함수가 중복 키를 생성하면 마지막 값이 우선합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2}
// result := maputil.MapEntries(m, func(k string, v int) (int, string) {
// return v, k // Swap key and value types / 키와 값 타입 교환
//	}) // map[int]string{1: "a", 2: "b"}
func MapEntries[K1 comparable, V1 any, K2 comparable, V2 any](m map[K1]V1, fn func(K1, V1) (K2, V2)) map[K2]V2 {
	result := make(map[K2]V2, len(m))
	for k, v := range m {
		newKey, newValue := fn(k, v)
		result[newKey] = newValue
	}
	return result
}

// Invert swaps keys and values in the map.
// Invert는 맵의 키와 값을 교환합니다.
//
// If multiple keys have the same value, the last key-value pair wins.
// 여러 키가 같은 값을 가지면 마지막 키-값 쌍이 우선합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3}
//	result := maputil.Invert(m) // map[int]string{1: "a", 2: "b", 3: "c"}
func Invert[K comparable, V comparable](m map[K]V) map[V]K {
	result := make(map[V]K, len(m))
	for k, v := range m {
		result[v] = k
	}
	return result
}

// Flatten converts a nested map to a flat map using a delimiter.
// Flatten은 구분자를 사용하여 중첩된 맵을 평면 맵으로 변환합니다.
//
// Time complexity: O(n*m) where n is the number of keys and m is the average nested map size
// 시간 복잡도: O(n*m) 여기서 n은 키 개수, m은 평균 중첩 맵 크기
//
// Example
// 예제:
//
//	m := map[string]map[string]int{
//	    "user1": {"age": 25, "score": 100},
//	    "user2": {"age": 30, "score": 95},
//	}
//	result := maputil.Flatten(m, ".") // map[string]int{"user1.age": 25, "user1.score": 100, ...}
func Flatten[K comparable, V any](m map[K]map[string]V, delimiter string) map[string]V {
	result := make(map[string]V)
	for k, nested := range m {
		prefix := fmt.Sprintf("%v", k)
		for nestedKey, nestedValue := range nested {
			flatKey := prefix + delimiter + nestedKey
			result[flatKey] = nestedValue
		}
	}
	return result
}

// Unflatten converts a flat map to a nested map using a delimiter.
// Unflatten은 구분자를 사용하여 평면 맵을 중첩된 맵으로 변환합니다.
//
// Time complexity: O(n*k) where n is the number of keys and k is the average key depth
// 시간 복잡도: O(n*k) 여기서 n은 키 개수, k는 평균 키 깊이
//
// Example
// 예제:
//
//	m := map[string]int{"user.name": 1, "user.age": 25, "admin.name": 2}
//	result := maputil.Unflatten(m, ".") // map[string]interface{}{"user": {...}, "admin": {...}}
func Unflatten[V any](m map[string]V, delimiter string) map[string]interface{} {
	result := make(map[string]interface{})

	for key, value := range m {
		parts := strings.Split(key, delimiter)
		current := result

		for i := 0; i < len(parts)-1; i++ {
			part := parts[i]
			if _, exists := current[part]; !exists {
				current[part] = make(map[string]interface{})
			}
			current = current[part].(map[string]interface{})
		}

		current[parts[len(parts)-1]] = value
	}

	return result
}

// Chunk splits a map into multiple smaller maps of the specified size.
// Chunk는 맵을 지정된 크기의 여러 작은 맵으로 분할합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
//	result := maputil.Chunk(m, 2) // []map[string]int{{"a": 1, "b": 2}, {"c": 3, "d": 4}, {"e": 5}}
func Chunk[K comparable, V any](m map[K]V, size int) []map[K]V {
	if size <= 0 {
		return nil
	}

	numChunks := (len(m) + size - 1) / size
	result := make([]map[K]V, 0, numChunks)
	current := make(map[K]V, size)

	i := 0
	for k, v := range m {
		current[k] = v
		i++

		if i%size == 0 {
			result = append(result, current)
			current = make(map[K]V, size)
		}
	}

	if len(current) > 0 {
		result = append(result, current)
	}

	return result
}

// Partition splits a map into two maps based on a predicate function.
// Partition은 조건 함수를 기반으로 맵을 두 개의 맵으로 분할합니다.
//
// Returns two maps: the first contains entries that satisfy the predicate,
// the second contains entries that don't.
//
// 두 개의 맵을 반환합니다: 첫 번째는 조건을 만족하는 항목,
// 두 번째는 그렇지 않은 항목을 포함합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//	even, odd := maputil.Partition(m, func(k string, v int) bool {
//	    return v%2 == 0
//	}) // even = {"b": 2, "d": 4}, odd = {"a": 1, "c": 3}
func Partition[K comparable, V any](m map[K]V, fn func(K, V) bool) (map[K]V, map[K]V) {
	trueMap := make(map[K]V)
	falseMap := make(map[K]V)

	for k, v := range m {
		if fn(k, v) {
			trueMap[k] = v
		} else {
			falseMap[k] = v
		}
	}

	return trueMap, falseMap
}

// Compact removes entries with zero values from the map.
// Compact는 맵에서 zero 값을 가진 항목을 제거합니다.
//
// Time complexity: O(n)
// 시간 복잡도: O(n)
//
// Example
// 예제:
//
//	m := map[string]int{"a": 1, "b": 0, "c": 3, "d": 0}
//	result := maputil.Compact(m) // map[string]int{"a": 1, "c": 3}
func Compact[K comparable, V comparable](m map[K]V) map[K]V {
	var zero V
	result := make(map[K]V)

	for k, v := range m {
		if v != zero {
			result[k] = v
		}
	}

	return result
}
