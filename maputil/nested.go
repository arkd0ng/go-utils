package maputil

import (
	"fmt"
)

// GetNested retrieves a value from a nested map using a path of keys.
// GetNested는 키 경로를 사용하여 중첩 맵에서 값을 검색합니다.
//
// This function navigates through nested map[string]interface{} structures
// using a sequence of keys. It returns the value at the final key and a boolean
// indicating whether the path was valid and the value exists.
//
// 이 함수는 키 시퀀스를 사용하여 중첩 map[string]interface{} 구조를 탐색합니다.
// 최종 키의 값과 경로가 유효하고 값이 존재하는지를 나타내는 부울을 반환합니다.
//
// Time Complexity
// 시간 복잡도: O(d) where d is depth (path length)
// Space Complexity
// 공간 복잡도: O(1)
//
// Parameters
// 매개변수:
// - m: The nested map to navigate
// 탐색할 중첩 맵
// - path: Sequence of keys to follow
// 따를 키 시퀀스
//
// Returns
// 반환값:
// - interface{}: The value at the path (nil if not found)
// 경로의 값 (찾을 수 없으면 nil)
// - bool: true if path exists, false otherwise
// 경로가 존재하면 true, 그렇지 않으면 false
//
// Example
// 예제:
//
//	data := map[string]interface{}{
//		"user": map[string]interface{}{
//			"name": "Alice",
//			"address": map[string]interface{}{
//				"city": "Seoul",
//				"zip": "12345",
//			},
//		},
//	}
//	city, ok := maputil.GetNested(data, "user", "address", "city")
//	// city = "Seoul", ok = true
//	missing, ok := maputil.GetNested(data, "user", "phone")
//	// missing = nil, ok = false
//
// Use Case
// 사용 사례:
// - JSON/YAML configuration access
// JSON/YAML 설정 접근
// - API response parsing
// API 응답 파싱
// - Nested data structure navigation
// 중첩 데이터 구조 탐색
// - Safe property access without panic
// panic 없는 안전한 속성 접근
func GetNested(m map[string]interface{}, path ...string) (interface{}, bool) {
	if len(path) == 0 {
		return nil, false
	}

	current := interface{}(m)

	for i, key := range path {
		// Type assert current value to map
		currentMap, ok := current.(map[string]interface{})
		if !ok {
			return nil, false
		}

		// Get value for this key
		value, exists := currentMap[key]
		if !exists {
			return nil, false
		}

		// If this is the last key, return the value
		if i == len(path)-1 {
			return value, true
		}

		// Otherwise, continue navigating
		current = value
	}

	return nil, false
}

// SetNested sets a value in a nested map, creating intermediate maps as needed.
// SetNested는 중첩 맵에 값을 설정하고, 필요한 경우 중간 맵을 생성합니다.
//
// This function navigates through or creates nested map[string]interface{} structures
// to set a value at the specified path. It creates any missing intermediate maps
// automatically. The function returns a new map with the modification (immutable).
//
// 이 함수는 중첩 map[string]interface{} 구조를 탐색하거나 생성하여
// 지정된 경로에 값을 설정합니다. 누락된 중간 맵을 자동으로 생성합니다.
// 함수는 수정된 새 맵을 반환합니다 (불변).
//
// Time Complexity
// 시간 복잡도: O(d) where d is depth (path length)
// Space Complexity
// 공간 복잡도: O(n) for deep copy
//
// Parameters
// 매개변수:
// - m: The nested map to modify
// 수정할 중첩 맵
// - value: The value to set at the path
// 경로에 설정할 값
// - path: Sequence of keys to follow
// 따를 키 시퀀스
//
// Returns
// 반환값:
// - map[string]interface{}: New map with the value set
// 값이 설정된 새 맵
//
// Example
// 예제:
//
//	data := map[string]interface{}{}
//	result := maputil.SetNested(data, "Seoul", "user", "address", "city")
//	// result = map[string]interface{}{
//	//   "user": map[string]interface{}{
//	//     "address": map[string]interface{}{
//	//       "city": "Seoul",
//	//     },
//	//   },
//	// }
//
// Use Case
// 사용 사례:
// - Dynamic configuration building
// 동적 설정 구축
// - API request body construction
// API 요청 본문 구성
// - Nested data structure initialization
// 중첩 데이터 구조 초기화
// - Deep property updates
// 깊은 속성 업데이트
func SetNested(m map[string]interface{}, value interface{}, path ...string) map[string]interface{} {
	if len(path) == 0 {
		return m
	}

	// Deep copy the map
	result := deepCopyMap(m)

	// Navigate to the parent of the final key
	current := result
	for i := 0; i < len(path)-1; i++ {
		key := path[i]

		// Check if the key exists
		if val, exists := current[key]; exists {
			// If exists, type assert to map
			if nestedMap, ok := val.(map[string]interface{}); ok {
				// Create a copy for immutability
				newMap := deepCopyMap(nestedMap)
				current[key] = newMap
				current = newMap
			} else {
				// Value exists but is not a map, overwrite it
				newMap := make(map[string]interface{})
				current[key] = newMap
				current = newMap
			}
		} else {
			// Key doesn't exist, create new map
			newMap := make(map[string]interface{})
			current[key] = newMap
			current = newMap
		}
	}

	// Set the final value
	finalKey := path[len(path)-1]
	current[finalKey] = value

	return result
}

// HasNested checks if a nested path exists in the map.
// HasNested는 중첩 경로가 맵에 존재하는지 확인합니다.
//
// This function verifies that all keys in the path exist and that intermediate
// values are maps. It returns true only if the entire path is valid.
//
// 이 함수는 경로의 모든 키가 존재하고 중간 값이 맵인지 확인합니다.
// 전체 경로가 유효한 경우에만 true를 반환합니다.
//
// Time Complexity
// 시간 복잡도: O(d) where d is depth (path length)
// Space Complexity
// 공간 복잡도: O(1)
//
// Parameters
// 매개변수:
// - m: The nested map to check
// 확인할 중첩 맵
// - path: Sequence of keys to verify
// 확인할 키 시퀀스
//
// Returns
// 반환값:
// - bool: true if entire path exists, false otherwise
// 전체 경로가 존재하면 true, 그렇지 않으면 false
//
// Example
// 예제:
//
//	data := map[string]interface{}{
//		"user": map[string]interface{}{
//			"name": "Alice",
//			"email": "alice@example.com",
//		},
//	}
//	maputil.HasNested(data, "user", "name")  // true
//	maputil.HasNested(data, "user", "phone") // false
//	maputil.HasNested(data, "admin")         // false
//
// Use Case
// 사용 사례:
// - Configuration validation
// 설정 검증
// - Required field checking
// 필수 필드 확인
// - API response validation
// API 응답 검증
// - Safe navigation guards
// 안전한 탐색 가드
func HasNested(m map[string]interface{}, path ...string) bool {
	if len(path) == 0 {
		return false
	}

	current := interface{}(m)

	for _, key := range path {
		// Type assert current value to map
		currentMap, ok := current.(map[string]interface{})
		if !ok {
			return false
		}

		// Check if key exists
		value, exists := currentMap[key]
		if !exists {
			return false
		}

		// Move to next level
		current = value
	}

	return true
}

// DeleteNested removes a value from a nested map at the specified path.
// DeleteNested는 지정된 경로의 중첩 맵에서 값을 제거합니다.
//
// This function navigates through the nested structure and removes the value
// at the final key. It does not remove intermediate maps, even if they become
// empty. Returns a new map with the modification (immutable).
//
// 이 함수는 중첩 구조를 탐색하고 최종 키의 값을 제거합니다.
// 중간 맵이 비어 있어도 제거하지 않습니다. 수정된 새 맵을 반환합니다 (불변).
//
// Time Complexity
// 시간 복잡도: O(d) where d is depth (path length)
// Space Complexity
// 공간 복잡도: O(n) for deep copy
//
// Parameters
// 매개변수:
// - m: The nested map to modify
// 수정할 중첩 맵
// - path: Sequence of keys to the value to delete
// 삭제할 값의 키 시퀀스
//
// Returns
// 반환값:
// - map[string]interface{}: New map with the value deleted
// 값이 삭제된 새 맵
//
// Example
// 예제:
//
//	data := map[string]interface{}{
//		"user": map[string]interface{}{
//			"name": "Alice",
//			"password": "secret123",
//		},
//	}
//	result := maputil.DeleteNested(data, "user", "password")
//	// result = map[string]interface{}{
//	//   "user": map[string]interface{}{
//	//     "name": "Alice",
//	//   },
//	// }
//
// Use Case
// 사용 사례:
// - Removing sensitive data
// 민감한 데이터 제거
// - Configuration cleanup
// 설정 정리
// - API response filtering
// API 응답 필터링
// - Nested property removal
// 중첩 속성 제거
func DeleteNested(m map[string]interface{}, path ...string) map[string]interface{} {
	if len(path) == 0 {
		return m
	}

	// Deep copy the map
	result := deepCopyMap(m)

	// Navigate to the parent of the final key
	current := result
	for i := 0; i < len(path)-1; i++ {
		key := path[i]

		// Check if the key exists
		val, exists := current[key]
		if !exists {
			return result // Path doesn't exist, return unchanged
		}

		// Type assert to map
		nestedMap, ok := val.(map[string]interface{})
		if !ok {
			return result // Not a map, can't continue
		}

		// Create a copy for immutability
		newMap := deepCopyMap(nestedMap)
		current[key] = newMap
		current = newMap
	}

	// Delete the final key
	finalKey := path[len(path)-1]
	delete(current, finalKey)

	return result
}

// SafeGet safely retrieves a value from a nested structure with error handling.
// SafeGet은 에러 처리와 함께 중첩 구조에서 안전하게 값을 검색합니다.
//
// This function is similar to GetNested but returns an error instead of a boolean.
// It provides more detailed error messages when the path is invalid or when
// type assertions fail. Unlike GetNested, this works with any input type.
//
// 이 함수는 GetNested와 유사하지만 부울 대신 에러를 반환합니다.
// 경로가 유효하지 않거나 타입 어설션이 실패할 때 더 자세한 에러 메시지를 제공합니다.
// GetNested와 달리 모든 입력 타입에서 작동합니다.
//
// Time Complexity
// 시간 복잡도: O(d) where d is depth (path length)
// Space Complexity
// 공간 복잡도: O(1)
//
// Parameters
// 매개변수:
// - m: The value to navigate (typically a map)
// 탐색할 값 (일반적으로 맵)
// - path: Sequence of keys to follow
// 따를 키 시퀀스
//
// Returns
// 반환값:
// - interface{}: The value at the path
// 경로의 값
// - error: Error if path is invalid or type assertion fails
// 경로가 유효하지 않거나 타입 어설션이 실패하면 에러
//
// Example
// 예제:
//
//	data := map[string]interface{}{
//		"server": map[string]interface{}{
//			"host": "localhost",
//			"port": 8080,
//		},
//	}
//	host, err := maputil.SafeGet(data, "server", "host")
//	// host = "localhost", err = nil
//	port, err := maputil.SafeGet(data, "server", "port")
//	// port = 8080, err = nil
//	invalid, err := maputil.SafeGet(data, "server", "timeout")
//	// invalid = nil, err = "key 'timeout' not found in map"
//
// Use Case
// 사용 사례:
// - Safe config access with error reporting
// 에러 보고와 함께 안전한 설정 접근
// - API response parsing with validation
// 검증과 함께 API 응답 파싱
// - Debugging nested data access
// 중첩 데이터 접근 디버깅
// - Error-driven nested navigation
// 에러 기반 중첩 탐색
func SafeGet(m interface{}, path ...string) (interface{}, error) {
	if len(path) == 0 {
		return nil, fmt.Errorf("path cannot be empty")
	}

	current := m

	for i, key := range path {
		// Type assert current value to map
		currentMap, ok := current.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("value at path %v is not a map (type: %T)", path[:i], current)
		}

		// Get value for this key
		value, exists := currentMap[key]
		if !exists {
			return nil, fmt.Errorf("key '%s' not found in map at path %v", key, path[:i+1])
		}

		// If this is the last key, return the value
		if i == len(path)-1 {
			return value, nil
		}

		// Otherwise, continue navigating
		current = value
	}

	return nil, fmt.Errorf("unexpected error navigating path")
}

// deepCopyMap creates a deep copy of a map[string]interface{}.
// deepCopyMap은 map[string]interface{}의 깊은 복사본을 생성합니다.
func deepCopyMap(m map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{}, len(m))
	for k, v := range m {
		if nestedMap, ok := v.(map[string]interface{}); ok {
			result[k] = deepCopyMap(nestedMap)
		} else {
			result[k] = v
		}
	}
	return result
}
