package maputil

import (
	"fmt"
)

// Package maputil/nested.go provides comprehensive operations for nested map structures.
// This file contains functions for navigating, accessing, modifying, and validating
// deeply nested map[string]interface{} structures using path-based access.
//
// maputil/nested.go 패키지는 중첩된 맵 구조에 대한 포괄적인 작업을 제공합니다.
// 이 파일은 경로 기반 접근을 사용하여 깊이 중첩된 map[string]interface{} 구조를
// 탐색, 접근, 수정, 검증하는 함수들을 포함합니다.
//
// # Overview | 개요
//
// The nested.go file provides 6 nested map operations organized into 3 categories:
//
// nested.go 파일은 3개 카테고리로 구성된 6개 중첩 맵 작업을 제공합니다:
//
// 1. READING OPERATIONS | 읽기 작업
//   - GetNested: Retrieve value at path with boolean result (O(d))
//   - SafeGet: Retrieve value at path with error result (O(d))
//   - HasNested: Check if path exists (O(d))
//
// 2. WRITING OPERATIONS | 쓰기 작업
//   - SetNested: Set value at path, creating intermediate maps (O(d))
//   - DeleteNested: Remove value at path (O(d))
//
// 3. UTILITY OPERATIONS | 유틸리티 작업
//   - deepCopyMap: Deep copy nested map structure (internal helper)
//
// # Design Principles | 설계 원칙
//
// 1. PATH-BASED ACCESS | 경로 기반 접근
//   - Variadic path arguments: GetNested(m, "user", "address", "city")
//   - Intuitive navigation through nested structures
//   - No need for manual nil checks at each level
//   - Single function call for deep access
//
//   가변 경로 인자: GetNested(m, "user", "address", "city")
//   중첩 구조를 통한 직관적 탐색
//   각 레벨에서 수동 nil 검사 불필요
//   깊은 접근을 위한 단일 함수 호출
//
// 2. IMMUTABILITY BY DEFAULT | 기본 불변성
//   - SetNested, DeleteNested: Return new maps with modifications
//   - Original maps remain unchanged
//   - Deep copy ensures complete immutability
//   - Safe for concurrent read access
//
//   SetNested, DeleteNested: 수정된 새 맵 반환
//   원본 맵은 변경되지 않음
//   깊은 복사로 완전한 불변성 보장
//   동시 읽기 접근에 안전
//
// 3. AUTOMATIC INTERMEDIATE CREATION | 자동 중간 생성
//   - SetNested automatically creates missing intermediate maps
//   - No need to pre-initialize nested structure
//   - Simplifies deep property setting
//   - Handles arbitrary nesting depth
//
//   SetNested가 누락된 중간 맵을 자동 생성
//   중첩 구조를 미리 초기화할 필요 없음
//   깊은 속성 설정을 단순화
//   임의의 중첩 깊이 처리
//
// 4. TYPE SAFETY WITH TYPE ASSERTIONS | 타입 어설션을 통한 타입 안전성
//   - All operations validate map types at each level
//   - GetNested, SafeGet: Return (value, bool/error) for safe access
//   - HasNested: Validates entire path before returning true
//   - Graceful handling of non-map intermediate values
//
//   모든 작업이 각 레벨에서 맵 타입 검증
//   GetNested, SafeGet: 안전한 접근을 위해 (값, bool/error) 반환
//   HasNested: true 반환 전 전체 경로 검증
//   맵이 아닌 중간 값을 우아하게 처리
//
// # Function Categories | 함수 카테고리
//
// READING OPERATIONS | 읽기 작업
//
// GetNested(m, path...) retrieves a value from a nested map using a path of keys.
// It navigates through the nested structure step by step, performing type assertions
// at each level. Returns (value, true) if the path is valid and exists, or (nil, false)
// if any key is missing or intermediate value is not a map.
//
// GetNested(m, path...)는 키 경로를 사용하여 중첩 맵에서 값을 검색합니다.
// 중첩 구조를 단계별로 탐색하며 각 레벨에서 타입 어설션을 수행합니다.
// 경로가 유효하고 존재하면 (값, true)를, 키가 누락되거나 중간 값이 맵이 아니면
// (nil, false)를 반환합니다.
//
// Time Complexity: O(d) where d = path depth (number of keys)
// Space Complexity: O(1) - no allocations for reading
// Return Value: (value, true) if found, (nil, false) otherwise
// Type Constraint: Only map[string]interface{} at all levels
// Nil Handling: Returns (nil, false) for nil maps or missing keys
// Early Termination: Stops at first missing key or type mismatch
//
// 시간 복잡도: O(d) 여기서 d = 경로 깊이 (키 개수)
// 공간 복잡도: O(1) - 읽기를 위한 할당 없음
// 반환값: 찾으면 (값, true), 그렇지 않으면 (nil, false)
// 타입 제약: 모든 레벨에서 map[string]interface{}만
// Nil 처리: nil 맵이나 누락된 키에 대해 (nil, false) 반환
// 조기 종료: 첫 번째 누락 키나 타입 불일치에서 중단
//
// Use Case: Safe nested property access without panics
// 사용 사례: 패닉 없는 안전한 중첩 속성 접근
//
// Example:
//   config := map[string]interface{}{
//       "database": map[string]interface{}{
//           "host": "localhost",
//           "port": 3306,
//           "credentials": map[string]interface{}{
//               "user": "admin",
//               "password": "secret",
//           },
//       },
//   }
//   user, ok := GetNested(config, "database", "credentials", "user")
//   // user = "admin", ok = true
//   missing, ok := GetNested(config, "database", "timeout")
//   // missing = nil, ok = false
//
// SafeGet(m, path...) is similar to GetNested but returns an error instead of a boolean.
// This provides more detailed information about why the access failed, including the
// exact path location and type mismatch details. Useful for debugging and error reporting.
//
// SafeGet(m, path...)는 GetNested와 유사하지만 부울 대신 에러를 반환합니다.
// 정확한 경로 위치와 타입 불일치 세부사항을 포함하여 접근이 실패한 이유에 대한
// 더 자세한 정보를 제공합니다. 디버깅과 에러 보고에 유용합니다.
//
// Time Complexity: O(d) where d = path depth
// Space Complexity: O(1) for navigation, O(d) for error message
// Return Value: (value, nil) if found, (nil, error) otherwise
// Error Types:
//   - Empty path error
//   - "not a map" type mismatch error
//   - "key not found" missing key error
// Type Constraint: Accepts any type, checks map[string]interface{} at runtime
// Nil Handling: Returns error for nil maps or missing keys
//
// 시간 복잡도: O(d) 여기서 d = 경로 깊이
// 공간 복잡도: 탐색은 O(1), 에러 메시지는 O(d)
// 반환값: 찾으면 (값, nil), 그렇지 않으면 (nil, error)
// 에러 타입:
//   - 빈 경로 에러
//   - "not a map" 타입 불일치 에러
//   - "key not found" 누락 키 에러
// 타입 제약: 모든 타입 허용, 런타임에 map[string]interface{} 검사
// Nil 처리: nil 맵이나 누락된 키에 대해 에러 반환
//
// Use Case: Configuration parsing with detailed error reporting
// 사용 사례: 자세한 에러 보고를 포함한 구성 파싱
//
// Example:
//   config := map[string]interface{}{
//       "server": map[string]interface{}{
//           "host": "localhost",
//       },
//   }
//   host, err := SafeGet(config, "server", "host")
//   // host = "localhost", err = nil
//   port, err := SafeGet(config, "server", "port")
//   // port = nil, err = "key 'port' not found in map at path [server port]"
//
// HasNested(m, path...) checks if a complete path exists in the nested map.
// It validates that all keys in the path exist and that all intermediate values
// are maps. Returns true only if the entire path is valid and traversable.
//
// HasNested(m, path...)는 중첩 맵에 완전한 경로가 존재하는지 확인합니다.
// 경로의 모든 키가 존재하고 모든 중간 값이 맵인지 검증합니다.
// 전체 경로가 유효하고 탐색 가능한 경우에만 true를 반환합니다.
//
// Time Complexity: O(d) where d = path depth
// Space Complexity: O(1)
// Return Value: true if entire path exists, false otherwise
// Validation: Checks both key existence and map type at each level
// Use Before Access: Useful to check before GetNested/SetNested
// Empty Path: Returns false for empty path
//
// 시간 복잡도: O(d) 여기서 d = 경로 깊이
// 공간 복잡도: O(1)
// 반환값: 전체 경로가 존재하면 true, 그렇지 않으면 false
// 검증: 각 레벨에서 키 존재와 맵 타입 모두 확인
// 접근 전 사용: GetNested/SetNested 전 확인에 유용
// 빈 경로: 빈 경로에 대해 false 반환
//
// Use Case: Configuration validation, required field checking
// 사용 사례: 구성 검증, 필수 필드 확인
//
// Example:
//   config := map[string]interface{}{
//       "database": map[string]interface{}{
//           "host": "localhost",
//       },
//   }
//   HasNested(config, "database", "host") // true
//   HasNested(config, "database", "port") // false
//   HasNested(config, "cache")            // false
//
// WRITING OPERATIONS | 쓰기 작업
//
// SetNested(m, value, path...) sets a value in a nested map, automatically creating
// any missing intermediate maps. Returns a completely new nested structure with the
// modification, preserving immutability. The original map remains unchanged.
//
// SetNested(m, value, path...)는 중첩 맵에 값을 설정하고, 누락된 중간 맵을 자동으로
// 생성합니다. 불변성을 유지하면서 수정된 완전히 새로운 중첩 구조를 반환합니다.
// 원본 맵은 변경되지 않습니다.
//
// Time Complexity: O(d) for navigation + O(n) for deep copy = O(n)
// Space Complexity: O(n) where n = total entries in nested structure
// Immutability: Returns new map, original unchanged
// Auto-creation: Creates missing intermediate maps automatically
// Overwrite: Overwrites non-map intermediate values with maps
// Deep Copy: Uses deepCopyMap to ensure complete immutability
// Empty Path: Returns original map unchanged
//
// 시간 복잡도: 탐색 O(d) + 깊은 복사 O(n) = O(n)
// 공간 복잡도: O(n) 여기서 n = 중첩 구조의 총 항목
// 불변성: 새 맵 반환, 원본 변경 없음
// 자동 생성: 누락된 중간 맵을 자동으로 생성
// 덮어쓰기: 맵이 아닌 중간 값을 맵으로 덮어씀
// 깊은 복사: 완전한 불변성을 보장하기 위해 deepCopyMap 사용
// 빈 경로: 원본 맵을 변경하지 않고 반환
//
// Use Case: Dynamic configuration building, nested property initialization
// 사용 사례: 동적 구성 구축, 중첩 속성 초기화
//
// Example:
//   config := map[string]interface{}{}
//   config = SetNested(config, "localhost", "server", "host")
//   config = SetNested(config, 8080, "server", "port")
//   config = SetNested(config, true, "server", "ssl", "enabled")
//   // config = map[string]interface{}{
//   //   "server": map[string]interface{}{
//   //     "host": "localhost",
//   //     "port": 8080,
//   //     "ssl": map[string]interface{}{
//   //       "enabled": true,
//   //     },
//   //   },
//   // }
//
// DeleteNested(m, path...) removes a value from a nested map at the specified path.
// Returns a new map with the value deleted. Does NOT remove empty intermediate maps
// after deletion. The original map remains unchanged (immutable operation).
//
// DeleteNested(m, path...)는 지정된 경로의 중첩 맵에서 값을 제거합니다.
// 값이 삭제된 새 맵을 반환합니다. 삭제 후 빈 중간 맵은 제거하지 않습니다.
// 원본 맵은 변경되지 않습니다 (불변 작업).
//
// Time Complexity: O(d) for navigation + O(n) for deep copy = O(n)
// Space Complexity: O(n) for deep copy
// Immutability: Returns new map, original unchanged
// Empty Intermediates: Does not remove empty parent maps
// Missing Path: Returns original map unchanged if path doesn't exist
// Deep Copy: Uses deepCopyMap to ensure immutability
// Empty Path: Returns original map unchanged
//
// 시간 복잡도: 탐색 O(d) + 깊은 복사 O(n) = O(n)
// 공간 복잡도: 깊은 복사를 위한 O(n)
// 불변성: 새 맵 반환, 원본 변경 없음
// 빈 중간: 빈 부모 맵을 제거하지 않음
// 누락 경로: 경로가 존재하지 않으면 원본 맵을 변경하지 않고 반환
// 깊은 복사: 불변성을 보장하기 위해 deepCopyMap 사용
// 빈 경로: 원본 맵을 변경하지 않고 반환
//
// Use Case: Removing sensitive data, configuration cleanup, property removal
// 사용 사례: 민감한 데이터 제거, 구성 정리, 속성 제거
//
// Example:
//   config := map[string]interface{}{
//       "user": map[string]interface{}{
//           "name": "Alice",
//           "password": "secret",
//           "email": "alice@example.com",
//       },
//   }
//   sanitized := DeleteNested(config, "user", "password")
//   // sanitized = map[string]interface{}{
//   //   "user": map[string]interface{}{
//   //     "name": "Alice",
//   //     "email": "alice@example.com",
//   //   },
//   // }
//   // Original config still has password
//
// UTILITY OPERATIONS | 유틸리티 작업
//
// deepCopyMap(m) creates a complete deep copy of a nested map[string]interface{}
// structure. Recursively copies nested maps to ensure complete immutability.
// This is an internal helper function used by SetNested and DeleteNested.
//
// deepCopyMap(m)은 중첩 map[string]interface{} 구조의 완전한 깊은 복사본을 생성합니다.
// 완전한 불변성을 보장하기 위해 중첩 맵을 재귀적으로 복사합니다.
// SetNested와 DeleteNested가 사용하는 내부 헬퍼 함수입니다.
//
// Time Complexity: O(n) where n = total entries in all nested maps
// Space Complexity: O(n) - creates complete copy
// Recursion: Handles arbitrary nesting depth
// Type Handling: Only copies nested maps recursively, other types copied by reference
// Internal Use: Not exported, used internally by mutation operations
//
// 시간 복잡도: O(n) 여기서 n = 모든 중첩 맵의 총 항목
// 공간 복잡도: O(n) - 완전한 복사본 생성
// 재귀: 임의의 중첩 깊이 처리
// 타입 처리: 중첩 맵만 재귀적으로 복사, 다른 타입은 참조로 복사
// 내부 사용: 내보내지지 않음, 변경 작업에서 내부적으로 사용
//
// # Comparisons with Related Functions | 관련 함수와 비교
//
// GetNested vs. SafeGet:
//   - GetNested: Returns (value, bool), simpler error handling
//   - SafeGet: Returns (value, error), detailed error messages
//   - Use GetNested for existence checks: if val, ok := GetNested(...); ok { }
//   - Use SafeGet for debugging and detailed error reporting
//
// GetNested 대 SafeGet:
//   - GetNested: (값, bool) 반환, 더 간단한 에러 처리
//   - SafeGet: (값, error) 반환, 자세한 에러 메시지
//   - 존재 확인에 GetNested 사용: if val, ok := GetNested(...); ok { }
//   - 디버깅과 자세한 에러 보고에 SafeGet 사용
//
// GetNested vs. basic.Get:
//   - GetNested: Nested access with path: GetNested(m, "a", "b", "c")
//   - basic.Get: Single-level access: Get(m, "key")
//   - GetNested for deep structures, basic.Get for flat maps
//   - GetNested requires map[string]interface{}, basic.Get is generic
//
// GetNested 대 basic.Get:
//   - GetNested: 경로로 중첩 접근: GetNested(m, "a", "b", "c")
//   - basic.Get: 단일 레벨 접근: Get(m, "key")
//   - 깊은 구조에 GetNested, 평면 맵에 basic.Get
//   - GetNested는 map[string]interface{} 필요, basic.Get는 제네릭
//
// SetNested vs. basic.Set:
//   - SetNested: Creates nested structure automatically
//   - basic.Set: Single-level, requires pre-existing map
//   - SetNested returns new map (immutable), basic.Set returns new map too
//   - SetNested for deep property initialization
//
// SetNested 대 basic.Set:
//   - SetNested: 중첩 구조를 자동으로 생성
//   - basic.Set: 단일 레벨, 기존 맵 필요
//   - SetNested는 새 맵 반환 (불변), basic.Set도 새 맵 반환
//   - 깊은 속성 초기화에 SetNested
//
// SetNested vs. DeepMerge:
//   - SetNested: Sets specific value at path
//   - DeepMerge: Recursively merges entire nested structures
//   - SetNested for targeted updates, DeepMerge for combining configs
//   - SetNested more efficient for single property changes
//
// SetNested 대 DeepMerge:
//   - SetNested: 경로의 특정 값 설정
//   - DeepMerge: 전체 중첩 구조를 재귀적으로 병합
//   - 타겟 업데이트에 SetNested, 구성 결합에 DeepMerge
//   - 단일 속성 변경에 SetNested가 더 효율적
//
// DeleteNested vs. basic.Delete:
//   - DeleteNested: Navigates nested path to delete
//   - basic.Delete: Single-level deletion
//   - DeleteNested for deep property removal
//   - basic.Delete for flat map cleanup
//
// DeleteNested 대 basic.Delete:
//   - DeleteNested: 삭제를 위해 중첩 경로 탐색
//   - basic.Delete: 단일 레벨 삭제
//   - 깊은 속성 제거에 DeleteNested
//   - 평면 맵 정리에 basic.Delete
//
// HasNested vs. basic.Has:
//   - HasNested: Checks nested path existence
//   - basic.Has: Checks single key existence
//   - HasNested validates entire path, basic.Has only top-level
//   - HasNested for configuration validation
//
// HasNested 대 basic.Has:
//   - HasNested: 중첩 경로 존재 확인
//   - basic.Has: 단일 키 존재 확인
//   - HasNested는 전체 경로 검증, basic.Has는 최상위만
//   - 구성 검증에 HasNested
//
// # Performance Characteristics | 성능 특성
//
// Time Complexities:
//   - O(d): GetNested, SafeGet, HasNested (d = path depth)
//   - O(n): SetNested, DeleteNested (n = total nested map entries, due to deep copy)
//
// 시간 복잡도:
//   - O(d): GetNested, SafeGet, HasNested (d = 경로 깊이)
//   - O(n): SetNested, DeleteNested (n = 총 중첩 맵 항목, 깊은 복사 때문)
//
// Space Complexities:
//   - O(1): GetNested, HasNested (read-only, no allocations)
//   - O(d): SafeGet (error message includes path)
//   - O(n): SetNested, DeleteNested (deep copy entire structure)
//
// 공간 복잡도:
//   - O(1): GetNested, HasNested (읽기 전용, 할당 없음)
//   - O(d): SafeGet (에러 메시지가 경로 포함)
//   - O(n): SetNested, DeleteNested (전체 구조 깊은 복사)
//
// Optimization Tips:
//   - Use GetNested over SafeGet when error details not needed (no error string allocation)
//   - Cache HasNested results for repeated path validation
//   - Batch multiple SetNested calls using DeepMerge for better performance
//   - Avoid deep nesting when possible (flatter structures are faster)
//   - Consider mutable operations (not provided) if immutability not required
//
// 최적화 팁:
//   - 에러 세부사항이 필요하지 않을 때 SafeGet 대신 GetNested 사용 (에러 문자열 할당 없음)
//   - 반복된 경로 검증을 위해 HasNested 결과 캐시
//   - 더 나은 성능을 위해 DeepMerge로 여러 SetNested 호출 일괄 처리
//   - 가능하면 깊은 중첩 피하기 (더 평면적인 구조가 더 빠름)
//   - 불변성이 필요하지 않으면 가변 작업 고려 (제공되지 않음)
//
// # Common Usage Patterns | 일반적인 사용 패턴
//
// 1. Configuration File Parsing | 구성 파일 파싱
//
//	config := map[string]interface{}{
//	    "server": map[string]interface{}{
//	        "host": "localhost",
//	        "port": 8080,
//	    },
//	}
//	host, ok := maputil.GetNested(config, "server", "host")
//	if !ok {
//	    log.Fatal("server.host not found in config")
//	}
//	fmt.Printf("Server: %s:%v\n", host, config["server"].(map[string]interface{})["port"])
//
// 2. API Response Parsing | API 응답 파싱
//
//	response := map[string]interface{}{
//	    "data": map[string]interface{}{
//	        "user": map[string]interface{}{
//	            "id": 123,
//	            "name": "Alice",
//	        },
//	    },
//	}
//	userName, err := maputil.SafeGet(response, "data", "user", "name")
//	if err != nil {
//	    log.Printf("Failed to get user name: %v", err)
//	} else {
//	    fmt.Printf("User: %s\n", userName)
//	}
//
// 3. Dynamic Configuration Building | 동적 구성 구축
//
//	config := map[string]interface{}{}
//	config = maputil.SetNested(config, "localhost", "database", "host")
//	config = maputil.SetNested(config, 5432, "database", "port")
//	config = maputil.SetNested(config, "admin", "database", "credentials", "user")
//	config = maputil.SetNested(config, "secret", "database", "credentials", "password")
//	// config = {
//	//   "database": {
//	//     "host": "localhost",
//	//     "port": 5432,
//	//     "credentials": {
//	//       "user": "admin",
//	//       "password": "secret"
//	//     }
//	//   }
//	// }
//
// 4. Required Field Validation | 필수 필드 검증
//
//	config := map[string]interface{}{
//	    "database": map[string]interface{}{
//	        "host": "localhost",
//	    },
//	}
//	requiredFields := [][]string{
//	    {"database", "host"},
//	    {"database", "port"},
//	    {"database", "user"},
//	}
//	for _, path := range requiredFields {
//	    if !maputil.HasNested(config, path...) {
//	        log.Fatalf("Required field missing: %v", path)
//	    }
//	}
//
// 5. Sensitive Data Removal | 민감한 데이터 제거
//
//	userData := map[string]interface{}{
//	    "user": map[string]interface{}{
//	        "name": "Alice",
//	        "email": "alice@example.com",
//	        "password": "secret123",
//	        "ssn": "123-45-6789",
//	    },
//	}
//	sanitized := maputil.DeleteNested(userData, "user", "password")
//	sanitized = maputil.DeleteNested(sanitized, "user", "ssn")
//	// Safe to log or send over network
//	fmt.Printf("User data: %+v\n", sanitized)
//
// 6. Nested Property Update | 중첩 속성 업데이트
//
//	state := map[string]interface{}{
//	    "app": map[string]interface{}{
//	        "settings": map[string]interface{}{
//	            "theme": "dark",
//	            "language": "en",
//	        },
//	    },
//	}
//	state = maputil.SetNested(state, "light", "app", "settings", "theme")
//	// state.app.settings.theme is now "light"
//
// 7. Safe Navigation with Defaults | 기본값을 사용한 안전한 탐색
//
//	config := map[string]interface{}{
//	    "server": map[string]interface{}{
//	        "host": "localhost",
//	    },
//	}
//	port, ok := maputil.GetNested(config, "server", "port")
//	if !ok {
//	    port = 8080 // Default port
//	}
//	fmt.Printf("Port: %v\n", port)
//
// 8. Multi-Level Existence Check | 다중 레벨 존재 확인
//
//	data := map[string]interface{}{
//	    "company": map[string]interface{}{
//	        "departments": map[string]interface{}{
//	            "engineering": map[string]interface{}{
//	                "employees": []string{"Alice", "Bob"},
//	            },
//	        },
//	    },
//	}
//	if maputil.HasNested(data, "company", "departments", "engineering", "employees") {
//	    employees, _ := maputil.GetNested(data, "company", "departments", "engineering", "employees")
//	    fmt.Printf("Employees: %v\n", employees)
//	}
//
// 9. Error-Driven Configuration Validation | 에러 기반 구성 검증
//
//	config := map[string]interface{}{
//	    "database": map[string]interface{}{
//	        "host": "localhost",
//	    },
//	}
//	if _, err := maputil.SafeGet(config, "database", "port"); err != nil {
//	    log.Printf("Configuration warning: %v, using default", err)
//	    config = maputil.SetNested(config, 5432, "database", "port")
//	}
//
// 10. Nested State Management | 중첩 상태 관리
//
//	state := map[string]interface{}{}
//
//	// Initialize nested state
//	state = maputil.SetNested(state, 0, "counters", "pageViews")
//	state = maputil.SetNested(state, 0, "counters", "clicks")
//
//	// Update nested values
//	if current, ok := maputil.GetNested(state, "counters", "pageViews"); ok {
//	    state = maputil.SetNested(state, current.(int)+1, "counters", "pageViews")
//	}
//
//	// Check and delete
//	if maputil.HasNested(state, "counters", "clicks") {
//	    state = maputil.DeleteNested(state, "counters", "clicks")
//	}
//
// # Edge Cases and Nil Handling | 엣지 케이스와 Nil 처리
//
// Empty Path:
//   - GetNested, SafeGet: Return (nil, false) or error
//   - HasNested: Returns false
//   - SetNested, DeleteNested: Return original map unchanged
//
// 빈 경로:
//   - GetNested, SafeGet: (nil, false) 또는 에러 반환
//   - HasNested: false 반환
//   - SetNested, DeleteNested: 원본 맵을 변경하지 않고 반환
//
// Nil Maps:
//   - GetNested: Returns (nil, false)
//   - SafeGet: Returns error
//   - HasNested: Returns false
//   - SetNested, DeleteNested: Creates new structure or returns empty map
//
// Nil 맵:
//   - GetNested: (nil, false) 반환
//   - SafeGet: 에러 반환
//   - HasNested: false 반환
//   - SetNested, DeleteNested: 새 구조 생성 또는 빈 맵 반환
//
// Non-Map Intermediate Values:
//   - GetNested, SafeGet, HasNested: Return false/error (can't navigate further)
//   - SetNested: Overwrites non-map value with map to continue path
//
// 맵이 아닌 중간 값:
//   - GetNested, SafeGet, HasNested: false/에러 반환 (더 이상 탐색 불가)
//   - SetNested: 경로를 계속하기 위해 맵이 아닌 값을 맵으로 덮어씀
//
// Missing Keys:
//   - GetNested: Returns (nil, false)
//   - SafeGet: Returns error with key name
//   - HasNested: Returns false
//   - SetNested: Creates missing keys automatically
//   - DeleteNested: Returns original map unchanged
//
// 누락된 키:
//   - GetNested: (nil, false) 반환
//   - SafeGet: 키 이름과 함께 에러 반환
//   - HasNested: false 반환
//   - SetNested: 누락된 키를 자동으로 생성
//   - DeleteNested: 원본 맵을 변경하지 않고 반환
//
// Single-Level Path:
//   - All functions work with single key: GetNested(m, "key")
//   - Equivalent to basic.Get for single-level access
//
// 단일 레벨 경로:
//   - 모든 함수가 단일 키로 작동: GetNested(m, "key")
//   - 단일 레벨 접근은 basic.Get과 동등
//
// # Thread Safety | 스레드 안전성
//
// Read Operations (Safe for Concurrent Reads):
//   - GetNested, SafeGet, HasNested
//   - All are read-only, no modifications
//   - Safe when map has concurrent readers
//
// 읽기 작업 (동시 읽기 안전):
//   - GetNested, SafeGet, HasNested
//   - 모두 읽기 전용, 수정 없음
//   - 맵에 동시 읽기가 있을 때 안전
//
// Write Operations (Return New Maps):
//   - SetNested, DeleteNested
//   - Create new maps, don't modify originals
//   - Safe when original map has concurrent readers
//   - Returned map should not be shared until fully initialized
//
// 쓰기 작업 (새 맵 반환):
//   - SetNested, DeleteNested
//   - 새 맵 생성, 원본 수정 없음
//   - 원본 맵에 동시 읽기가 있을 때 안전
//   - 반환된 맵은 완전히 초기화될 때까지 공유하지 않아야 함
//
// Concurrent Modification Warning:
//   - Do not modify input maps while operations read them
//   - Deep copy in SetNested/DeleteNested protects original
//   - Use sync.RWMutex for concurrent access patterns
//
// 동시 수정 경고:
//   - 작업이 읽는 동안 입력 맵 수정 금지
//   - SetNested/DeleteNested의 깊은 복사가 원본 보호
//   - 동시 접근 패턴에 sync.RWMutex 사용
//
// # See Also | 참고
//
// Related files in maputil package:
//   - basic.go: Flat map operations (Get, Set, Delete)
//   - merge.go: DeepMerge for nested map merging
//   - transform.go: Flatten, Unflatten for nested <-> flat conversion
//   - filter.go: Filtering operations (complementary to nested access)
//
// maputil 패키지의 관련 파일:
//   - basic.go: 평면 맵 작업 (Get, Set, Delete)
//   - merge.go: 중첩 맵 병합을 위한 DeepMerge
//   - transform.go: 중첩 <-> 평면 변환을 위한 Flatten, Unflatten
//   - filter.go: 필터링 작업 (중첩 접근 보완)

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
