package maputil

// GetOrSet retrieves a value if it exists, otherwise sets and returns the default value.
// GetOrSet은 값이 존재하면 검색하고, 그렇지 않으면 기본값을 설정하고 반환합니다.
//
// This function provides a convenient way to ensure a key always has a value.
// If the key exists, it returns the existing value. If not, it sets the key
// to the default value and returns it. The original map is modified.
//
// 이 함수는 키가 항상 값을 가지도록 보장하는 편리한 방법을 제공합니다.
// 키가 존재하면 기존 값을 반환합니다. 존재하지 않으면 키를
// 기본값으로 설정하고 반환합니다. 원본 맵이 수정됩니다.
//
// Time Complexity / 시간 복잡도: O(1)
// Space Complexity / 공간 복잡도: O(1)
//
// Parameters / 매개변수:
//   - m: The input map (will be modified) / 입력 맵 (수정됨)
//   - key: The key to look up / 조회할 키
//   - defaultValue: The default value to set if key doesn't exist / 키가 존재하지 않으면 설정할 기본값
//
// Returns / 반환값:
//   - V: The existing value or the default value / 기존 값 또는 기본값
//
// Example / 예제:
//
//	cache := map[string]int{"a": 1, "b": 2}
//	maputil.GetOrSet(cache, "a", 10) // Returns 1 (exists)
//	maputil.GetOrSet(cache, "c", 10) // Returns 10 (doesn't exist, sets c=10)
//	// cache is now: map[string]int{"a": 1, "b": 2, "c": 10}
//
// Use Case / 사용 사례:
//   - Cache initialization / 캐시 초기화
//   - Default value management / 기본값 관리
//   - Configuration with fallbacks / 폴백이 있는 설정
//   - Lazy initialization / 지연 초기화
func GetOrSet[K comparable, V any](m map[K]V, key K, defaultValue V) V {
	if value, exists := m[key]; exists {
		return value
	}
	m[key] = defaultValue
	return defaultValue
}

// SetDefault sets a key to a default value only if the key doesn't already exist.
// SetDefault는 키가 아직 존재하지 않는 경우에만 키를 기본값으로 설정합니다.
//
// This function provides a way to initialize a key with a default value without
// overwriting existing values. It returns whether the key was set (true if the
// key didn't exist before, false if it already existed).
//
// 이 함수는 기존 값을 덮어쓰지 않고 기본값으로 키를 초기화하는 방법을 제공합니다.
// 키가 설정되었는지 여부를 반환합니다 (키가 이전에 존재하지 않았으면 true,
// 이미 존재했으면 false).
//
// Time Complexity / 시간 복잡도: O(1)
// Space Complexity / 공간 복잡도: O(1)
//
// Parameters / 매개변수:
//   - m: The input map (will be modified) / 입력 맵 (수정됨)
//   - key: The key to set / 설정할 키
//   - defaultValue: The default value to set if key doesn't exist / 키가 존재하지 않으면 설정할 기본값
//
// Returns / 반환값:
//   - bool: true if key was set, false if key already existed / 키가 설정되었으면 true, 이미 존재했으면 false
//
// Example / 예제:
//
//	config := map[string]string{"host": "localhost"}
//	maputil.SetDefault(config, "port", "8080")    // Returns true, sets port=8080
//	maputil.SetDefault(config, "host", "0.0.0.0") // Returns false, host remains localhost
//	// config: map[string]string{"host": "localhost", "port": "8080"}
//
// Use Case / 사용 사례:
//   - Safe initialization without overwriting / 덮어쓰지 않는 안전한 초기화
//   - Default configuration setup / 기본 설정 설정
//   - Conditional key creation / 조건부 키 생성
func SetDefault[K comparable, V any](m map[K]V, key K, defaultValue V) bool {
	if _, exists := m[key]; exists {
		return false
	}
	m[key] = defaultValue
	return true
}

// Defaults returns a new map with default values for missing keys.
// Defaults는 누락된 키에 대해 기본값이 있는 새 맵을 반환합니다.
//
// This function creates a new map that contains all entries from the original map
// plus any keys from the defaults map that weren't in the original. Existing keys
// in the original map are not overwritten.
//
// 이 함수는 원본 맵의 모든 항목과 원본에 없는 defaults 맵의 키를
// 포함하는 새 맵을 생성합니다. 원본 맵의 기존 키는 덮어쓰지 않습니다.
//
// Time Complexity / 시간 복잡도: O(n + d) where n is original map size, d is defaults size
// Space Complexity / 공간 복잡도: O(n + d)
//
// Parameters / 매개변수:
//   - m: The original map / 원본 맵
//   - defaults: Map containing default key-value pairs / 기본 키-값 쌍을 포함하는 맵
//
// Returns / 반환값:
//   - map[K]V: New map with defaults applied / 기본값이 적용된 새 맵
//
// Example / 예제:
//
//	config := map[string]string{"host": "localhost"}
//	defaults := map[string]string{"host": "0.0.0.0", "port": "8080", "timeout": "30s"}
//	result := maputil.Defaults(config, defaults)
//	// result: map[string]string{"host": "localhost", "port": "8080", "timeout": "30s"}
//	// Note: host keeps original value, port and timeout are added from defaults
//
// Use Case / 사용 사례:
//   - Configuration management with defaults / 기본값이 있는 설정 관리
//   - Template rendering with default values / 기본값이 있는 템플릿 렌더링
//   - API response with fallback values / 폴백 값이 있는 API 응답
//   - User preferences with system defaults / 시스템 기본값이 있는 사용자 기본 설정
func Defaults[K comparable, V any](m, defaults map[K]V) map[K]V {
	result := make(map[K]V, len(m)+len(defaults))

	// First, copy all defaults
	for k, v := range defaults {
		result[k] = v
	}

	// Then, overwrite with original values (original takes precedence)
	for k, v := range m {
		result[k] = v
	}

	return result
}
