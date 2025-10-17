package maputil

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// Package maputil/convert.go provides comprehensive map conversion operations.
// This file contains functions for converting maps to/from slices, entries,
// JSON, YAML, and other data formats, enabling interoperability and serialization.
//
// maputil/convert.go 패키지는 포괄적인 맵 변환 작업을 제공합니다.
// 이 파일은 맵을 슬라이스, 항목, JSON, YAML 및 기타 데이터 형식으로/에서 변환하는
// 함수들을 포함하여 상호 운용성과 직렬화를 가능하게 합니다.
//
// # Overview | 개요
//
// The convert.go file provides 11 conversion operations organized into 4 categories:
//
// convert.go 파일은 4개 카테고리로 구성된 11개 변환 작업을 제공합니다:
//
// 1. EXTRACTION OPERATIONS | 추출 작업
//   - Keys: Extract all keys as slice (O(n))
//   - Values: Extract all values as slice (O(n))
//   - Entries: Extract key-value pairs as Entry structs (O(n))
//
// 2. SLICE CONVERSION | 슬라이스 변환
//   - FromEntries: Create map from Entry slice (O(n))
//   - FromSlice: Create map from slice with key extractor (O(n))
//   - FromSliceBy: Create map from slice with key+value extractors (O(n))
//   - ToSlice: Convert map to slice with transformer (O(n))
//
// 3. JSON SERIALIZATION | JSON 직렬화
//   - ToJSON: Serialize map to JSON string (O(n))
//   - FromJSON: Deserialize JSON string to map (O(n))
//
// 4. YAML SERIALIZATION | YAML 직렬화
//   - ToYAML: Serialize map to YAML string (O(n))
//   - FromYAML: Deserialize YAML string to map (O(n))
//
// # Design Principles | 설계 원칙
//
// 1. BIDIRECTIONAL CONVERSION | 양방향 변환
//   - ToX / FromX pairs for symmetric operations
//   - Entries <-> Map conversion
//   - JSON <-> Map conversion
//   - YAML <-> Map conversion
//   - Ensures data can be round-tripped
//
//   대칭 작업을 위한 ToX / FromX 쌍
//   Entries <-> Map 변환
//   JSON <-> Map 변환
//   YAML <-> Map 변환
//   데이터 왕복이 가능함을 보장
//
// 2. FLEXIBLE KEY/VALUE EXTRACTION | 유연한 키/값 추출
//   - FromSlice: Single key extractor function
//   - FromSliceBy: Separate key and value extractors
//   - ToSlice: Custom transformation function
//   - Supports complex data transformations
//
//   FromSlice: 단일 키 추출 함수
//   FromSliceBy: 별도 키 및 값 추출자
//   ToSlice: 커스텀 변환 함수
//   복잡한 데이터 변환 지원
//
// 3. DUPLICATE KEY HANDLING | 중복 키 처리
//   - Last-wins strategy in FromEntries, FromSlice, FromSliceBy
//   - Consistent behavior across all construction functions
//   - No error on duplicates, silent overwrite
//   - Predictable and simple semantics
//
//   FromEntries, FromSlice, FromSliceBy에서 마지막 우선 전략
//   모든 구성 함수에서 일관된 동작
//   중복에 에러 없음, 자동 덮어쓰기
//   예측 가능하고 간단한 의미론
//
// 4. FORMAT AGNOSTIC | 형식 독립적
//   - JSON: Standard web API format
//   - YAML: Human-readable configuration format
//   - Entry: Type-safe intermediate representation
//   - Slice: General-purpose sequential data
//
//   JSON: 표준 웹 API 형식
//   YAML: 사람이 읽을 수 있는 구성 형식
//   Entry: 타입 안전 중간 표현
//   Slice: 범용 순차 데이터
//
// # Function Categories | 함수 카테고리
//
// EXTRACTION OPERATIONS | 추출 작업
//
// Keys(m) extracts all keys from a map as a slice. This is one of the most fundamental
// map operations, providing a simple way to iterate over keys or pass them to other
// functions. Order is not guaranteed due to map's unordered nature.
//
// Keys(m)는 맵의 모든 키를 슬라이스로 추출합니다. 가장 기본적인 맵 작업 중 하나로,
// 키를 순회하거나 다른 함수에 전달하는 간단한 방법을 제공합니다. 맵의 순서 없는
// 특성으로 인해 순서는 보장되지 않습니다.
//
// Time Complexity: O(n) - iterates through all entries
// Space Complexity: O(n) - creates slice with all keys
// Order: Non-deterministic (map iteration order is random)
// Nil Handling: Returns empty slice for nil maps
// Pre-allocation: Pre-allocates slice with len(m) capacity
//
// 시간 복잡도: O(n) - 모든 항목 순회
// 공간 복잡도: O(n) - 모든 키를 가진 슬라이스 생성
// 순서: 비결정적 (맵 순회 순서는 무작위)
// Nil 처리: nil 맵에 대해 빈 슬라이스 반환
// 사전 할당: len(m) 용량으로 슬라이스 사전 할당
//
// Use Case: Key enumeration, iteration, set operations
// 사용 사례: 키 나열, 순회, 집합 연산
//
// Example:
//   config := map[string]int{"timeout": 30, "retries": 3, "cache": 100}
//   keys := Keys(config)
//   // keys = []string{"timeout", "retries", "cache"} (order may vary)
//
// Values(m) extracts all values from a map as a slice. Complementary to Keys(),
// this provides access to all values without their keys. Useful for aggregation
// or when key information is not needed.
//
// Values(m)는 맵의 모든 값을 슬라이스로 추출합니다. Keys()의 보완으로,
// 키 없이 모든 값에 대한 접근을 제공합니다. 집계나 키 정보가 필요하지 않을 때 유용합니다.
//
// Time Complexity: O(n) - iterates through all entries
// Space Complexity: O(n) - creates slice with all values
// Order: Non-deterministic (map iteration order is random)
// Nil Handling: Returns empty slice for nil maps
// Pre-allocation: Pre-allocates slice with len(m) capacity
//
// 시간 복잡도: O(n) - 모든 항목 순회
// 공간 복잡도: O(n) - 모든 값을 가진 슬라이스 생성
// 순서: 비결정적 (맵 순회 순서는 무작위)
// Nil 처리: nil 맵에 대해 빈 슬라이스 반환
// 사전 할당: len(m) 용량으로 슬라이스 사전 할당
//
// Use Case: Value aggregation, statistics, when keys are irrelevant
// 사용 사례: 값 집계, 통계, 키가 무관할 때
//
// Example:
//   scores := map[string]int{"alice": 95, "bob": 87, "charlie": 92}
//   allScores := Values(scores)
//   // allScores = []int{95, 87, 92} (order may vary)
//
// Entries(m) extracts all key-value pairs as Entry structs. This creates a type-safe
// intermediate representation that preserves both keys and values. Useful for
// serialization or when you need to pass map data to functions expecting slices.
//
// Entries(m)는 모든 키-값 쌍을 Entry 구조체로 추출합니다. 키와 값을 모두 보존하는
// 타입 안전 중간 표현을 생성합니다. 직렬화나 슬라이스를 기대하는 함수에 맵 데이터를
// 전달해야 할 때 유용합니다.
//
// Time Complexity: O(n) - iterates through all entries
// Space Complexity: O(n) - creates slice of Entry structs
// Order: Non-deterministic (map iteration order is random)
// Entry Structure: Entry[K, V]{Key: K, Value: V}
// Type Safety: Preserves key and value types
// Nil Handling: Returns empty slice for nil maps
//
// 시간 복잡도: O(n) - 모든 항목 순회
// 공간 복잡도: O(n) - Entry 구조체 슬라이스 생성
// 순서: 비결정적 (맵 순회 순서는 무작위)
// Entry 구조: Entry[K, V]{Key: K, Value: V}
// 타입 안전성: 키와 값 타입 보존
// Nil 처리: nil 맵에 대해 빈 슬라이스 반환
//
// Use Case: Serialization, sorting entries, type-safe conversion
// 사용 사례: 직렬화, 항목 정렬, 타입 안전 변환
//
// Example:
//   m := map[string]int{"a": 1, "b": 2}
//   entries := Entries(m)
//   // entries = []Entry[string, int]{
//   //     {Key: "a", Value: 1},
//   //     {Key: "b", Value: 2},
//   // } (order may vary)
//
// SLICE CONVERSION OPERATIONS | 슬라이스 변환 작업
//
// FromEntries(entries) creates a map from a slice of Entry structs. This is the
// inverse operation of Entries(), enabling round-trip conversion. If duplicate keys
// exist in the entries, the last entry wins (last-wins strategy).
//
// FromEntries(entries)는 Entry 구조체 슬라이스로부터 맵을 생성합니다. Entries()의
// 역연산으로, 왕복 변환을 가능하게 합니다. 항목에 중복 키가 존재하면, 마지막 항목이
// 우선합니다 (마지막 우선 전략).
//
// Time Complexity: O(n) - iterates through all entries
// Space Complexity: O(n) - creates map with all entries
// Duplicate Keys: Last entry wins (silent overwrite)
// Pre-allocation: Pre-allocates map with len(entries) capacity
// Round-trip: FromEntries(Entries(m)) recreates m (except order)
//
// 시간 복잡도: O(n) - 모든 항목 순회
// 공간 복잡도: O(n) - 모든 항목을 가진 맵 생성
// 중복 키: 마지막 항목 우선 (자동 덮어쓰기)
// 사전 할당: len(entries) 용량으로 맵 사전 할당
// 왕복: FromEntries(Entries(m))가 m 재생성 (순서 제외)
//
// Use Case: Reconstructing maps from serialized entries
// 사용 사례: 직렬화된 항목으로부터 맵 재구성
//
// Example:
//   entries := []Entry[string, int]{
//       {Key: "a", Value: 1},
//       {Key: "b", Value: 2},
//   }
//   m := FromEntries(entries)
//   // m = map[string]int{"a": 1, "b": 2}
//
// FromSlice(slice, keyFn) creates a map from a slice by extracting keys using a
// function. The values in the map are the original slice elements. This is useful
// for indexing a collection by a computed key.
//
// FromSlice(slice, keyFn)는 함수를 사용하여 키를 추출하여 슬라이스로부터 맵을 생성합니다.
// 맵의 값은 원본 슬라이스 요소입니다. 계산된 키로 컬렉션을 인덱싱하는 데 유용합니다.
//
// Time Complexity: O(n) - iterates through all slice elements
// Space Complexity: O(n) - creates map with all elements
// Key Extractor: func(V) K - computes key from value
// Duplicate Keys: Last value wins
// Use Case: Indexing collections by ID, name, or computed field
//
// 시간 복잡도: O(n) - 모든 슬라이스 요소 순회
// 공간 복잡도: O(n) - 모든 요소를 가진 맵 생성
// 키 추출자: func(V) K - 값으로부터 키 계산
// 중복 키: 마지막 값 우선
// 사용 사례: ID, 이름 또는 계산된 필드로 컬렉션 인덱싱
//
// Example:
//   type User struct { ID int; Name string }
//   users := []User{{ID: 1, Name: "Alice"}, {ID: 2, Name: "Bob"}}
//   userMap := FromSlice(users, func(u User) int { return u.ID })
//   // userMap = map[int]User{1: {ID: 1, Name: "Alice"}, 2: {ID: 2, Name: "Bob"}}
//
// FromSliceBy(slice, keyFn, valueFn) creates a map from a slice using both key and
// value extractors. This allows transformation of both keys and values during
// conversion, providing maximum flexibility.
//
// FromSliceBy(slice, keyFn, valueFn)는 키와 값 추출자를 모두 사용하여 슬라이스로부터
// 맵을 생성합니다. 변환 중 키와 값 모두를 변환할 수 있어 최대 유연성을 제공합니다.
//
// Time Complexity: O(n) - iterates through all slice elements
// Space Complexity: O(n) - creates map with all transformed entries
// Key Extractor: func(V) K - computes key from element
// Value Extractor: func(V) R - computes value from element
// Duplicate Keys: Last value wins
// Type Change: Can change value type (V -> R)
//
// 시간 복잡도: O(n) - 모든 슬라이스 요소 순회
// 공간 복잡도: O(n) - 모든 변환된 항목을 가진 맵 생성
// 키 추출자: func(V) K - 요소로부터 키 계산
// 값 추출자: func(V) R - 요소로부터 값 계산
// 중복 키: 마지막 값 우선
// 타입 변경: 값 타입 변경 가능 (V -> R)
//
// Use Case: Creating lookup maps with transformed values
// 사용 사례: 변환된 값을 가진 조회 맵 생성
//
// Example:
//   type User struct { ID int; Name string; Age int }
//   users := []User{{ID: 1, Name: "Alice", Age: 25}, {ID: 2, Name: "Bob", Age: 30}}
//   nameMap := FromSliceBy(users,
//       func(u User) int { return u.ID },
//       func(u User) string { return u.Name },
//   )
//   // nameMap = map[int]string{1: "Alice", 2: "Bob"}
//
// ToSlice(m, fn) converts a map to a slice using a transformation function.
// This provides a flexible way to transform map entries into any slice type,
// enabling custom aggregation or formatting.
//
// ToSlice(m, fn)는 변환 함수를 사용하여 맵을 슬라이스로 변환합니다.
// 맵 항목을 모든 슬라이스 타입으로 변환하는 유연한 방법을 제공하여,
// 커스텀 집계나 포맷팅을 가능하게 합니다.
//
// Time Complexity: O(n) - applies transformation to each entry
// Space Complexity: O(n) - creates slice with transformed elements
// Transformer: func(K, V) R - transforms key-value to result type
// Order: Non-deterministic (map iteration order)
// Use Case: Custom formatting, aggregation, projection
//
// 시간 복잡도: O(n) - 각 항목에 변환 적용
// 공간 복잡도: O(n) - 변환된 요소를 가진 슬라이스 생성
// 변환자: func(K, V) R - 키-값을 결과 타입으로 변환
// 순서: 비결정적 (맵 순회 순서)
// 사용 사례: 커스텀 포맷팅, 집계, 프로젝션
//
// Example:
//   m := map[string]int{"a": 1, "b": 2, "c": 3}
//   formatted := ToSlice(m, func(k string, v int) string {
//       return fmt.Sprintf("%s=%d", k, v)
//   })
//   // formatted = []string{"a=1", "b=2", "c=3"} (order may vary)
//
// JSON SERIALIZATION OPERATIONS | JSON 직렬화 작업
//
// ToJSON(m) serializes a map to a JSON string. This uses Go's standard encoding/json
// package to marshal the map. Keys must be JSON-serializable (typically strings).
// Returns error if marshaling fails.
//
// ToJSON(m)는 맵을 JSON 문자열로 직렬화합니다. Go의 표준 encoding/json 패키지를
// 사용하여 맵을 마샬링합니다. 키는 JSON 직렬화 가능해야 합니다 (일반적으로 문자열).
// 마샬링이 실패하면 에러를 반환합니다.
//
// Time Complexity: O(n) - marshals all entries
// Space Complexity: O(n) - creates JSON string
// Format: Standard JSON object notation
// Error Handling: Returns error if marshaling fails
// Use Case: API responses, configuration export, data persistence
//
// 시간 복잡도: O(n) - 모든 항목 마샬
// 공간 복잡도: O(n) - JSON 문자열 생성
// 형식: 표준 JSON 객체 표기법
// 에러 처리: 마샬링 실패 시 에러 반환
// 사용 사례: API 응답, 구성 내보내기, 데이터 지속성
//
// Example:
//   config := map[string]int{"timeout": 30, "retries": 3}
//   jsonStr, err := ToJSON(config)
//   // jsonStr = `{"timeout":30,"retries":3}`
//
// FromJSON(jsonStr, m) deserializes a JSON string into a map. This uses Go's
// standard encoding/json package to unmarshal. The map pointer must be provided
// to receive the result. Returns error if unmarshaling fails.
//
// FromJSON(jsonStr, m)는 JSON 문자열을 맵으로 역직렬화합니다. Go의 표준
// encoding/json 패키지를 사용하여 언마샬합니다. 결과를 받을 맵 포인터를
// 제공해야 합니다. 언마샬링이 실패하면 에러를 반환합니다.
//
// Time Complexity: O(n) - unmarshals all entries
// Space Complexity: O(n) - populates map
// Format: Standard JSON object notation
// Error Handling: Returns error if unmarshaling fails
// Pointer Required: Must pass *map[K]V to receive result
//
// 시간 복잡도: O(n) - 모든 항목 언마샬
// 공간 복잡도: O(n) - 맵 채움
// 형식: 표준 JSON 객체 표기법
// 에러 처리: 언마샬링 실패 시 에러 반환
// 포인터 필요: 결과를 받으려면 *map[K]V 전달 필요
//
// Use Case: API request parsing, configuration import
// 사용 사례: API 요청 파싱, 구성 가져오기
//
// Example:
//   var config map[string]int
//   err := FromJSON(`{"timeout":30,"retries":3}`, &config)
//   // config = map[string]int{"timeout": 30, "retries": 3}
//
// YAML SERIALIZATION OPERATIONS | YAML 직렬화 작업
//
// ToYAML(m) serializes a map to a YAML string. This uses the gopkg.in/yaml.v3
// package to marshal the map. YAML is more human-readable than JSON and supports
// nested structures naturally. Returns error if marshaling fails.
//
// ToYAML(m)는 맵을 YAML 문자열로 직렬화합니다. gopkg.in/yaml.v3 패키지를
// 사용하여 맵을 마샬링합니다. YAML은 JSON보다 사람이 읽기 쉽고 중첩 구조를
// 자연스럽게 지원합니다. 마샬링이 실패하면 에러를 반환합니다.
//
// Time Complexity: O(n) - marshals all entries
// Space Complexity: O(n) - creates YAML string
// Format: YAML 1.2 specification
// Human-Readable: Better for configuration files
// Error Handling: Returns error if marshaling fails
// Nested Support: Handles nested maps naturally
//
// 시간 복잡도: O(n) - 모든 항목 마샬
// 공간 복잡도: O(n) - YAML 문자열 생성
// 형식: YAML 1.2 사양
// 사람이 읽을 수 있음: 구성 파일에 더 좋음
// 에러 처리: 마샬링 실패 시 에러 반환
// 중첩 지원: 중첩 맵을 자연스럽게 처리
//
// Use Case: Configuration file generation, human-readable exports
// 사용 사례: 구성 파일 생성, 사람이 읽을 수 있는 내보내기
//
// Example:
//   config := map[string]interface{}{
//       "server": map[string]interface{}{
//           "host": "localhost",
//           "port": 8080,
//       },
//   }
//   yamlStr, err := ToYAML(config)
//   // yamlStr:
//   // server:
//   //   host: localhost
//   //   port: 8080
//
// FromYAML(yamlStr) deserializes a YAML string into a map[string]interface{}.
// This uses the gopkg.in/yaml.v3 package to unmarshal. Returns the map and error.
// Note: Unlike FromJSON, this returns the map directly (not via pointer).
//
// FromYAML(yamlStr)는 YAML 문자열을 map[string]interface{}로 역직렬화합니다.
// gopkg.in/yaml.v3 패키지를 사용하여 언마샬합니다. 맵과 에러를 반환합니다.
// 참고: FromJSON과 달리, 맵을 직접 반환합니다 (포인터 아님).
//
// Time Complexity: O(n) - unmarshals all entries
// Space Complexity: O(n) - creates map
// Format: YAML 1.2 specification
// Return Type: map[string]interface{} (always this type)
// Error Handling: Returns nil map and error on failure
// Direct Return: Returns map directly, not via pointer
//
// 시간 복잡도: O(n) - 모든 항목 언마샬
// 공간 복잡도: O(n) - 맵 생성
// 형식: YAML 1.2 사양
// 반환 타입: map[string]interface{} (항상 이 타입)
// 에러 처리: 실패 시 nil 맵과 에러 반환
// 직접 반환: 포인터가 아닌 맵 직접 반환
//
// Use Case: Configuration file loading, YAML config parsing
// 사용 사례: 구성 파일 로딩, YAML 설정 파싱
//
// Example:
//   yamlStr := `
//   server:
//     host: localhost
//     port: 8080
//   `
//   config, err := FromYAML(yamlStr)
//   // config = map[string]interface{}{
//   //   "server": map[string]interface{}{
//   //     "host": "localhost",
//   //     "port": 8080,
//   //   },
//   // }
//
// # Comparisons with Related Functions | 관련 함수와 비교
//
// Keys vs. KeysSlice (keys.go):
//   - Functionally identical (KeysSlice is an alias for Keys)
//   - Both return []K with all keys
//   - Use Keys in convert.go context, KeysSlice elsewhere
//
// Keys 대 KeysSlice (keys.go):
//   - 기능적으로 동일 (KeysSlice는 Keys의 별칭)
//   - 둘 다 모든 키를 가진 []K 반환
//   - convert.go 컨텍스트에서 Keys, 다른 곳에서 KeysSlice 사용
//
// Values vs. ValuesSlice (values.go):
//   - Functionally identical (ValuesSlice is an alias for Values)
//   - Both return []V with all values
//   - Use Values in convert.go context, ValuesSlice elsewhere
//
// Values 대 ValuesSlice (values.go):
//   - 기능적으로 동일 (ValuesSlice는 Values의 별칭)
//   - 둘 다 모든 값을 가진 []V 반환
//   - convert.go 컨텍스트에서 Values, 다른 곳에서 ValuesSlice 사용
//
// Entries vs. ToSlice:
//   - Entries: Creates []Entry[K, V] (fixed structure)
//   - ToSlice: Creates []R with custom transformer (flexible)
//   - Entries for type-safe serialization
//   - ToSlice for custom formatting
//
// Entries 대 ToSlice:
//   - Entries: []Entry[K, V] 생성 (고정 구조)
//   - ToSlice: 커스텀 변환자로 []R 생성 (유연)
//   - 타입 안전 직렬화에 Entries
//   - 커스텀 포맷팅에 ToSlice
//
// FromSlice vs. FromSliceBy:
//   - FromSlice: Single key extractor, values are original elements
//   - FromSliceBy: Separate key+value extractors, can transform values
//   - FromSlice simpler when values don't need transformation
//   - FromSliceBy more flexible for custom transformations
//
// FromSlice 대 FromSliceBy:
//   - FromSlice: 단일 키 추출자, 값은 원본 요소
//   - FromSliceBy: 별도 키+값 추출자, 값 변환 가능
//   - 값 변환이 필요하지 않을 때 FromSlice가 더 간단
//   - 커스텀 변환에 FromSliceBy가 더 유연
//
// ToJSON vs. ToYAML:
//   - ToJSON: Compact, machine-readable, standard web format
//   - ToYAML: Human-readable, better for configuration files
//   - JSON for APIs, YAML for config files
//   - YAML supports comments, JSON does not
//
// ToJSON 대 ToYAML:
//   - ToJSON: 간결, 기계 읽기 가능, 표준 웹 형식
//   - ToYAML: 사람이 읽을 수 있음, 구성 파일에 더 좋음
//   - API에 JSON, 구성 파일에 YAML
//   - YAML은 주석 지원, JSON은 미지원
//
// FromJSON vs. FromYAML:
//   - FromJSON: Requires pointer parameter (*map[K]V)
//   - FromYAML: Returns map directly (map[string]interface{})
//   - FromJSON more generic (any K, V types)
//   - FromYAML always returns map[string]interface{}
//
// FromJSON 대 FromYAML:
//   - FromJSON: 포인터 매개변수 필요 (*map[K]V)
//   - FromYAML: 맵을 직접 반환 (map[string]interface{})
//   - FromJSON이 더 범용 (모든 K, V 타입)
//   - FromYAML은 항상 map[string]interface{} 반환
//
// FromEntries vs. transform.MapEntries:
//   - FromEntries: []Entry -> map[K]V (construction)
//   - MapEntries: map[K]V -> map[K2]V2 (transformation)
//   - FromEntries for slice-to-map conversion
//   - MapEntries for map-to-map transformation
//
// FromEntries 대 transform.MapEntries:
//   - FromEntries: []Entry -> map[K]V (구성)
//   - MapEntries: map[K]V -> map[K2]V2 (변환)
//   - 슬라이스-맵 변환에 FromEntries
//   - 맵-맵 변환에 MapEntries
//
// # Performance Characteristics | 성능 특성
//
// Time Complexities:
//   - O(n): All functions (Keys, Values, Entries, FromEntries, FromSlice, FromSliceBy, ToSlice, ToJSON, FromJSON, ToYAML, FromYAML)
//
// 시간 복잡도:
//   - O(n): 모든 함수 (Keys, Values, Entries, FromEntries, FromSlice, FromSliceBy, ToSlice, ToJSON, FromJSON, ToYAML, FromYAML)
//
// Space Complexities:
//   - O(n): All functions create new data structures
//
// 공간 복잡도:
//   - O(n): 모든 함수가 새 데이터 구조 생성
//
// Optimization Tips:
//   - Pre-allocation used in Keys, Values, Entries, FromEntries for efficiency
//   - ToJSON/ToYAML: Cache serialized strings if map doesn't change
//   - FromJSON/FromYAML: Reuse map variables to avoid repeated allocations
//   - FromSlice/FromSliceBy: Use when O(1) lookup needed after conversion
//
// 최적화 팁:
//   - 효율성을 위해 Keys, Values, Entries, FromEntries에서 사전 할당 사용
//   - ToJSON/ToYAML: 맵이 변경되지 않으면 직렬화된 문자열 캐시
//   - FromJSON/FromYAML: 반복 할당 피하기 위해 맵 변수 재사용
//   - FromSlice/FromSliceBy: 변환 후 O(1) 조회가 필요할 때 사용
//
// # Common Usage Patterns | 일반적인 사용 패턴
//
// 1. Indexing Collection by ID | ID로 컬렉션 인덱싱
//
//	type User struct { ID int; Name string; Email string }
//	users := []User{
//	    {ID: 1, Name: "Alice", Email: "alice@example.com"},
//	    {ID: 2, Name: "Bob", Email: "bob@example.com"},
//	}
//	userMap := maputil.FromSlice(users, func(u User) int { return u.ID })
//	// Fast O(1) lookup: userMap[1] returns Alice
//
// 2. Creating Name-to-Age Lookup | 이름-나이 조회 생성
//
//	type Person struct { Name string; Age int; City string }
//	people := []Person{
//	    {Name: "Alice", Age: 25, City: "Seoul"},
//	    {Name: "Bob", Age: 30, City: "Busan"},
//	}
//	ageMap := maputil.FromSliceBy(people,
//	    func(p Person) string { return p.Name },
//	    func(p Person) int { return p.Age },
//	)
//	// ageMap = map[string]int{"Alice": 25, "Bob": 30}
//
// 3. Configuration Export to JSON | JSON으로 구성 내보내기
//
//	config := map[string]interface{}{
//	    "timeout": 30,
//	    "retries": 3,
//	    "debug":   true,
//	}
//	jsonStr, err := maputil.ToJSON(config)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	// Save to file or send in API response
//	ioutil.WriteFile("config.json", []byte(jsonStr), 0644)
//
// 4. Configuration Import from YAML | YAML로부터 구성 가져오기
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
//	if err != nil {
//	    log.Fatal(err)
//	}
//	// Access nested values
//	serverConfig := config["server"].(map[string]interface{})
//	host := serverConfig["host"].(string)
//
// 5. Formatting Map as String Slice | 맵을 문자열 슬라이스로 포맷팅
//
//	env := map[string]string{"PATH": "/usr/bin", "HOME": "/home/user"}
//	envVars := maputil.ToSlice(env, func(k, v string) string {
//	    return fmt.Sprintf("%s=%s", k, v)
//	})
//	// envVars = []string{"PATH=/usr/bin", "HOME=/home/user"}
//	// Useful for environment variable formatting
//
// 6. Round-trip Entries Conversion | 왕복 항목 변환
//
//	original := map[string]int{"a": 1, "b": 2, "c": 3}
//	entries := maputil.Entries(original)
//	// entries = []Entry[string, int]{{Key: "a", Value: 1}, ...}
//	restored := maputil.FromEntries(entries)
//	// restored equals original (order may differ)
//
// 7. Extract and Sort Keys | 키 추출 및 정렬
//
//	scores := map[string]int{"alice": 95, "charlie": 87, "bob": 92}
//	keys := maputil.Keys(scores)
//	sort.Strings(keys)
//	for _, name := range keys {
//	    fmt.Printf("%s: %d\n", name, scores[name])
//	}
//	// Output: alice: 95, bob: 92, charlie: 87
//
// 8. Aggregate Values | 값 집계
//
//	sales := map[string]int{"jan": 1000, "feb": 1500, "mar": 1200}
//	values := maputil.Values(sales)
//	total := 0
//	for _, v := range values {
//	    total += v
//	}
//	// total = 3700
//
// 9. API Response Serialization | API 응답 직렬화
//
//	response := map[string]interface{}{
//	    "status":  "success",
//	    "data":    map[string]int{"count": 42},
//	    "message": "Request processed",
//	}
//	jsonResponse, _ := maputil.ToJSON(response)
//	w.Header().Set("Content-Type", "application/json")
//	w.Write([]byte(jsonResponse))
//
// 10. Configuration Merge and Export | 구성 병합 및 내보내기
//
//	defaults := map[string]interface{}{"timeout": 30, "retries": 3}
//	userConfig := map[string]interface{}{"timeout": 60}
//	finalConfig := maputil.Merge(defaults, userConfig)
//	yamlConfig, _ := maputil.ToYAML(finalConfig)
//	ioutil.WriteFile("config.yaml", []byte(yamlConfig), 0644)
//
// # Edge Cases and Nil Handling | 엣지 케이스와 Nil 처리
//
// Empty Maps:
//   - Keys, Values, Entries: Return empty slices
//   - ToJSON: Returns "{}"
//   - ToYAML: Returns "{}\n"
//
// 빈 맵:
//   - Keys, Values, Entries: 빈 슬라이스 반환
//   - ToJSON: "{}" 반환
//   - ToYAML: "{}\n" 반환
//
// Nil Maps:
//   - Keys, Values, Entries: Return empty slices (no panic)
//   - ToJSON, ToYAML: Return "null" or error (depends on marshaler)
//
// Nil 맵:
//   - Keys, Values, Entries: 빈 슬라이스 반환 (패닉 없음)
//   - ToJSON, ToYAML: "null" 또는 에러 반환 (마샬러에 의존)
//
// Empty Slices:
//   - FromEntries, FromSlice, FromSliceBy: Return empty maps
//
// 빈 슬라이스:
//   - FromEntries, FromSlice, FromSliceBy: 빈 맵 반환
//
// Duplicate Keys:
//   - FromEntries, FromSlice, FromSliceBy: Last value wins (silent overwrite)
//
// 중복 키:
//   - FromEntries, FromSlice, FromSliceBy: 마지막 값 우선 (자동 덮어쓰기)
//
// Invalid JSON/YAML:
//   - FromJSON, FromYAML: Return error, nil/empty map
//
// 잘못된 JSON/YAML:
//   - FromJSON, FromYAML: 에러 반환, nil/빈 맵
//
// # Thread Safety | 스레드 안전성
//
// Read Operations (Safe for Concurrent Reads):
//   - Keys, Values, Entries, ToSlice, ToJSON, ToYAML
//   - All read from input maps without modification
//
// 읽기 작업 (동시 읽기 안전):
//   - Keys, Values, Entries, ToSlice, ToJSON, ToYAML
//   - 모두 수정 없이 입력 맵에서 읽음
//
// Write Operations (Create New Maps):
//   - FromEntries, FromSlice, FromSliceBy, FromJSON, FromYAML
//   - Create new maps, safe as long as inputs not modified concurrently
//
// 쓰기 작업 (새 맵 생성):
//   - FromEntries, FromSlice, FromSliceBy, FromJSON, FromYAML
//   - 새 맵 생성, 입력이 동시에 수정되지 않는 한 안전
//
// Concurrent Modification Warning:
//   - Do not modify maps while reading/converting them
//   - Use sync.RWMutex for concurrent access
//
// 동시 수정 경고:
//   - 읽기/변환 중 맵 수정 금지
//   - 동시 접근에 sync.RWMutex 사용
//
// # See Also | 참고
//
// Related files in maputil package:
//   - keys.go: KeysSlice (alias for Keys), KeysSorted, KeysBy
//   - values.go: ValuesSlice (alias for Values), ValuesSorted, ValuesBy
//   - transform.go: Map, MapKeys, MapValues, MapEntries for in-place transformations
//   - basic.go: Clone, Equal for map operations
//
// maputil 패키지의 관련 파일:
//   - keys.go: KeysSlice (Keys의 별칭), KeysSorted, KeysBy
//   - values.go: ValuesSlice (Values의 별칭), ValuesSorted, ValuesBy
//   - transform.go: 제자리 변환을 위한 Map, MapKeys, MapValues, MapEntries
//   - basic.go: 맵 작업을 위한 Clone, Equal

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
