package validation

import (
	"fmt"
	"reflect"
)

// In validates that the value exists in the given list of allowed values.
// Uses deep equality for comparison, supporting any comparable type.
//
// In은 값이 주어진 허용 값 목록에 존재하는지 검증합니다.
// 깊은 동등성 비교를 사용하며 비교 가능한 모든 타입을 지원합니다.
//
// Parameters / 매개변수:
//   - values: Variadic list of allowed values
//     허용되는 값들의 가변 인자 목록
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses reflect.DeepEqual for comparison
//     reflect.DeepEqual을 사용하여 비교
//   - Supports any type (strings, numbers, structs, etc.)
//     모든 타입 지원 (문자열, 숫자, 구조체 등)
//   - Case-sensitive for strings
//     문자열은 대소문자 구분
//   - Fails if value not found in list
//     목록에서 값을 찾지 못하면 실패
//
// Use Cases / 사용 사례:
//   - Enum-like validation / 열거형 유사 검증
//   - Whitelist validation / 화이트리스트 검증
//   - Status/state validation / 상태/상태 검증
//   - Role-based validation / 역할 기반 검증
//   - Allowed option selection / 허용된 옵션 선택
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n) where n is number of allowed values
//     시간 복잡도: O(n) (n은 허용 값의 수)
//   - Linear search through values list
//     값 목록을 선형 검색
//   - DeepEqual can be expensive for complex types
//     DeepEqual은 복잡한 타입에서 비용이 클 수 있음
//
// Example / 예제:
//   // String values / 문자열 값
//   v := validation.New("red", "color")
//   v.In("red", "green", "blue")  // Passes / 성공
//
//   v := validation.New("yellow", "color")
//   v.In("red", "green", "blue")  // Fails / 실패
//
//   // Numeric values / 숫자 값
//   v := validation.New(2, "status")
//   v.In(1, 2, 3)  // Passes / 성공
//
//   // Status validation / 상태 검증
//   v := validation.New(status, "order_status")
//   v.Required().In("pending", "processing", "completed", "cancelled")
//
//   // Role validation / 역할 검증
//   v := validation.New(role, "user_role")
//   v.In("admin", "editor", "viewer")
func (v *Validator) In(values ...interface{}) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	found := false
	for _, val := range values {
		if reflect.DeepEqual(v.value, val) {
			found = true
			break
		}
	}

	if !found {
		v.addError("in", fmt.Sprintf("%s must be one of the allowed values / %s은(는) 허용된 값 중 하나여야 합니다", v.fieldName, v.fieldName))
	}

	return v
}

// NotIn validates that the value does NOT exist in the given list of forbidden values.
// Uses deep equality for comparison, supporting any comparable type.
//
// NotIn은 값이 주어진 금지 값 목록에 존재하지 않는지 검증합니다.
// 깊은 동등성 비교를 사용하며 비교 가능한 모든 타입을 지원합니다.
//
// Parameters / 매개변수:
//   - values: Variadic list of forbidden values
//     금지된 값들의 가변 인자 목록
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses reflect.DeepEqual for comparison
//     reflect.DeepEqual을 사용하여 비교
//   - Supports any type (strings, numbers, structs, etc.)
//     모든 타입 지원 (문자열, 숫자, 구조체 등)
//   - Case-sensitive for strings
//     문자열은 대소문자 구분
//   - Fails if value found in forbidden list
//     금지 목록에서 값을 찾으면 실패
//
// Use Cases / 사용 사례:
//   - Blacklist validation / 블랙리스트 검증
//   - Reserved value prevention / 예약된 값 방지
//   - Forbidden status/state validation / 금지된 상태/상태 검증
//   - Role restriction / 역할 제한
//   - Banned input prevention / 금지된 입력 방지
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n) where n is number of forbidden values
//     시간 복잡도: O(n) (n은 금지 값의 수)
//   - Linear search through values list
//     값 목록을 선형 검색
//   - DeepEqual can be expensive for complex types
//     DeepEqual은 복잡한 타입에서 비용이 클 수 있음
//
// Example / 예제:
//   // Reserved usernames / 예약된 사용자명
//   v := validation.New("user123", "username")
//   v.NotIn("admin", "root", "system")  // Passes / 성공
//
//   v := validation.New("admin", "username")
//   v.NotIn("admin", "root", "system")  // Fails / 실패
//
//   // Role restriction / 역할 제한
//   v := validation.New("editor", "role")
//   v.NotIn("guest", "anonymous")  // Passes / 성공
//
//   // Forbidden status / 금지된 상태
//   v := validation.New(newStatus, "status")
//   v.Required().NotIn("deleted", "archived")
//
//   // Reserved words / 예약어
//   v := validation.New(tableName, "table_name")
//   v.NotIn("select", "delete", "drop", "insert")
func (v *Validator) NotIn(values ...interface{}) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	found := false
	for _, val := range values {
		if reflect.DeepEqual(v.value, val) {
			found = true
			break
		}
	}

	if found {
		v.addError("notin", fmt.Sprintf("%s must not be one of the forbidden values / %s은(는) 금지된 값이 아니어야 합니다", v.fieldName, v.fieldName))
	}

	return v
}

// ArrayLength validates that the array or slice has exactly the specified length.
// Works with both arrays and slices. Fails if value is not an array/slice type.
//
// ArrayLength는 배열 또는 슬라이스가 정확히 지정된 길이를 가지는지 검증합니다.
// 배열과 슬라이스 모두에서 작동합니다. 값이 배열/슬라이스 타입이 아니면 실패합니다.
//
// Parameters / 매개변수:
//   - length: Exact required length
//     필요한 정확한 길이
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses reflection to get length
//     반사를 사용하여 길이 확인
//   - Works with both arrays and slices
//     배열과 슬라이스 모두에서 작동
//   - Fails if not array/slice type
//     배열/슬라이스 타입이 아니면 실패
//   - Exact match required (not >=, not <=)
//     정확히 일치 필요 (이상/이하 아님)
//
// Use Cases / 사용 사례:
//   - Fixed-size collection validation / 고정 크기 컬렉션 검증
//   - Coordinate tuples (x, y, z) / 좌표 튜플 (x, y, z)
//   - RGB color arrays [R, G, B] / RGB 색상 배열
//   - Fixed parameter counts / 고정 매개변수 수
//   - Matrix dimensions / 행렬 차원
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Reflection overhead for type checking
//     타입 검사를 위한 반사 오버헤드
//
// Example / 예제:
//   // Exact length / 정확한 길이
//   v := validation.New([]int{1, 2, 3}, "numbers")
//   v.ArrayLength(3)  // Passes / 성공
//
//   v := validation.New([]int{1, 2}, "numbers")
//   v.ArrayLength(3)  // Fails (too short) / 실패 (너무 짧음)
//
//   // RGB color validation / RGB 색상 검증
//   v := validation.New([3]int{255, 128, 64}, "rgb")
//   v.ArrayLength(3)  // Passes / 성공
//
//   // Coordinate validation / 좌표 검증
//   v := validation.New(coords, "coordinates")
//   v.ArrayNotEmpty().ArrayLength(3) // Must be [x, y, z]
//
//   // Invalid type / 유효하지 않은 타입
//   v := validation.New("string", "value")
//   v.ArrayLength(3)  // Fails (not array/slice) / 실패 (배열/슬라이스 아님)
func (v *Validator) ArrayLength(length int) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	val := reflect.ValueOf(v.value)
	if val.Kind() != reflect.Slice && val.Kind() != reflect.Array {
		v.addError("arraylength", fmt.Sprintf("%s must be an array or slice / %s은(는) 배열 또는 슬라이스여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if val.Len() != length {
		v.addError("arraylength", fmt.Sprintf("%s length must be exactly %d / %s 길이는 정확히 %d여야 합니다", v.fieldName, length, v.fieldName, length))
	}

	return v
}

// ArrayMinLength validates that the array or slice has at least the minimum length.
// Works with both arrays and slices. Fails if value is not an array/slice type.
//
// ArrayMinLength는 배열 또는 슬라이스가 최소 길이를 가지는지 검증합니다.
// 배열과 슬라이스 모두에서 작동합니다. 값이 배열/슬라이스 타입이 아니면 실패합니다.
//
// Parameters / 매개변수:
//   - min: Minimum required length (inclusive)
//     필요한 최소 길이 (포함)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses reflection to get length
//     반사를 사용하여 길이 확인
//   - Works with both arrays and slices
//     배열과 슬라이스 모두에서 작동
//   - Fails if not array/slice type
//     배열/슬라이스 타입이 아니면 실패
//   - Inclusive comparison (length == min is valid)
//     포함 비교 (길이 == 최소값 유효)
//
// Use Cases / 사용 사례:
//   - Minimum item requirements / 최소 항목 요구사항
//   - Batch processing (minimum batch size)
//     배치 처리 (최소 배치 크기)
//   - Form validation (select at least N items)
//     폼 검증 (최소 N개 항목 선택)
//   - Required options selection / 필수 옵션 선택
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Reflection overhead for type checking
//     타입 검사를 위한 반사 오버헤드
//
// Example / 예제:
//   // Minimum items / 최소 항목
//   v := validation.New([]string{"a", "b", "c"}, "items")
//   v.ArrayMinLength(2)  // Passes / 성공
//
//   v := validation.New([]string{"a"}, "items")
//   v.ArrayMinLength(2)  // Fails / 실패
//
//   v := validation.New([]int{1, 2}, "numbers")
//   v.ArrayMinLength(2)  // Passes (boundary) / 성공 (경계값)
//
//   // Multiple selections / 다중 선택
//   v := validation.New(selectedItems, "selections")
//   v.ArrayNotEmpty().ArrayMinLength(1).ArrayMaxLength(10)
//
//   // Batch validation / 배치 검증
//   v := validation.New(batch, "batch")
//   v.ArrayMinLength(10) // Minimum 10 items per batch
func (v *Validator) ArrayMinLength(min int) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	val := reflect.ValueOf(v.value)
	if val.Kind() != reflect.Slice && val.Kind() != reflect.Array {
		v.addError("arrayminlength", fmt.Sprintf("%s must be an array or slice / %s은(는) 배열 또는 슬라이스여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if val.Len() < min {
		v.addError("arrayminlength", fmt.Sprintf("%s length must be at least %d / %s 길이는 최소 %d여야 합니다", v.fieldName, min, v.fieldName, min))
	}

	return v
}

// ArrayMaxLength validates that the array or slice has at most the maximum length.
// Works with both arrays and slices. Fails if value is not an array/slice type.
//
// ArrayMaxLength는 배열 또는 슬라이스가 최대 길이를 넘지 않는지 검증합니다.
// 배열과 슬라이스 모두에서 작동합니다. 값이 배열/슬라이스 타입이 아니면 실패합니다.
//
// Parameters / 매개변수:
//   - max: Maximum allowed length (inclusive)
//     허용되는 최대 길이 (포함)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses reflection to get length
//     반사를 사용하여 길이 확인
//   - Works with both arrays and slices
//     배열과 슬라이스 모두에서 작동
//   - Fails if not array/slice type
//     배열/슬라이스 타입이 아니면 실패
//   - Inclusive comparison (length == max is valid)
//     포함 비교 (길이 == 최대값 유효)
//
// Use Cases / 사용 사례:
//   - Maximum item limits / 최대 항목 제한
//   - API response size limits / API 응답 크기 제한
//   - Form validation (select at most N items)
//     폼 검증 (최대 N개 항목 선택)
//   - Batch size limits / 배치 크기 제한
//   - Memory constraints / 메모리 제약
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Reflection overhead for type checking
//     타입 검사를 위한 반사 오버헤드
//
// Example / 예제:
//   // Maximum items / 최대 항목
//   v := validation.New([]int{1, 2, 3}, "numbers")
//   v.ArrayMaxLength(5)  // Passes / 성공
//
//   v := validation.New([]int{1, 2, 3, 4, 5, 6}, "numbers")
//   v.ArrayMaxLength(5)  // Fails / 실패
//
//   v := validation.New([]string{"a", "b", "c"}, "items")
//   v.ArrayMaxLength(3)  // Passes (boundary) / 성공 (경계값)
//
//   // Limit selections / 선택 제한
//   v := validation.New(selectedOptions, "options")
//   v.ArrayNotEmpty().ArrayMaxLength(3)
//
//   // Batch size limit / 배치 크기 제한
//   v := validation.New(batch, "batch")
//   v.ArrayMinLength(1).ArrayMaxLength(100)
func (v *Validator) ArrayMaxLength(max int) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	val := reflect.ValueOf(v.value)
	if val.Kind() != reflect.Slice && val.Kind() != reflect.Array {
		v.addError("arraymaxlength", fmt.Sprintf("%s must be an array or slice / %s은(는) 배열 또는 슬라이스여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if val.Len() > max {
		v.addError("arraymaxlength", fmt.Sprintf("%s length must be at most %d / %s 길이는 최대 %d여야 합니다", v.fieldName, max, v.fieldName, max))
	}

	return v
}

// ArrayNotEmpty validates that the array or slice is not empty (has at least one element).
// Works with both arrays and slices. Fails if value is not an array/slice type.
//
// ArrayNotEmpty는 배열 또는 슬라이스가 비어있지 않은지 (최소 하나의 요소 포함) 검증합니다.
// 배열과 슬라이스 모두에서 작동합니다. 값이 배열/슬라이스 타입이 아니면 실패합니다.
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses reflection to check length
//     반사를 사용하여 길이 확인
//   - Works with both arrays and slices
//     배열과 슬라이스 모두에서 작동
//   - Fails if not array/slice type
//     배열/슬라이스 타입이 아니면 실패
//   - Fails if length is zero
//     길이가 0이면 실패
//
// Use Cases / 사용 사례:
//   - Required collection validation / 필수 컬렉션 검증
//   - Ensure selection made / 선택 확인
//   - Non-empty result validation / 비어있지 않은 결과 검증
//   - Required items in cart / 장바구니의 필수 항목
//   - Attachment requirement / 첨부 파일 요구사항
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Reflection overhead for type checking
//     타입 검사를 위한 반사 오버헤드
//
// Note / 참고:
//   - Equivalent to ArrayMinLength(1)
//     ArrayMinLength(1)과 동등
//   - More semantic and readable
//     더 의미론적이고 읽기 쉬움
//
// Example / 예제:
//   // Must have items / 항목 필수
//   v := validation.New([]string{"item"}, "items")
//   v.ArrayNotEmpty()  // Passes / 성공
//
//   v := validation.New([]string{}, "items")
//   v.ArrayNotEmpty()  // Fails / 실패
//
//   v := validation.New([]int{0}, "numbers")
//   v.ArrayNotEmpty()  // Passes (has one element) / 성공 (요소 하나 있음)
//
//   // Shopping cart / 장바구니
//   v := validation.New(cartItems, "cart")
//   v.ArrayNotEmpty() // Cart must have items
//
//   // File attachments / 파일 첨부
//   v := validation.New(attachments, "attachments")
//   v.ArrayNotEmpty().ArrayMaxLength(5)
func (v *Validator) ArrayNotEmpty() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	val := reflect.ValueOf(v.value)
	if val.Kind() != reflect.Slice && val.Kind() != reflect.Array {
		v.addError("arraynotempty", fmt.Sprintf("%s must be an array or slice / %s은(는) 배열 또는 슬라이스여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if val.Len() == 0 {
		v.addError("arraynotempty", fmt.Sprintf("%s must not be empty / %s은(는) 비어있지 않아야 합니다", v.fieldName, v.fieldName))
	}

	return v
}

// ArrayUnique validates that all elements in the array or slice are unique (no duplicates).
// Works with both arrays and slices. Fails if value is not an array/slice type.
//
// ArrayUnique는 배열 또는 슬라이스의 모든 요소가 고유한지 (중복 없음) 검증합니다.
// 배열과 슬라이스 모두에서 작동합니다. 값이 배열/슬라이스 타입이 아니면 실패합니다.
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses map to track seen elements
//     본 요소를 추적하기 위해 맵 사용
//   - Works with both arrays and slices
//     배열과 슬라이스 모두에서 작동
//   - Fails if not array/slice type
//     배열/슬라이스 타입이 아니면 실패
//   - Fails on first duplicate found
//     첫 번째 중복을 찾으면 실패
//   - Empty arrays/slices pass validation
//     빈 배열/슬라이스는 검증 통과
//
// Use Cases / 사용 사례:
//   - Unique ID validation / 고유 ID 검증
//   - Unique tag/category lists / 고유 태그/카테고리 목록
//   - Prevent duplicate selections / 중복 선택 방지
//   - Unique email lists / 고유 이메일 목록
//   - Set-like behavior validation / 집합 유사 동작 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n) where n is array length
//     시간 복잡도: O(n) (n은 배열 길이)
//   - Space complexity: O(n) for tracking seen elements
//     공간 복잡도: O(n) (본 요소 추적)
//   - Early termination on duplicate found
//     중복 발견 시 조기 종료
//
// Limitations / 제한사항:
//   - Elements must be map-key compatible
//     요소는 맵 키 호환 가능해야 함
//   - Complex types (slices, maps) may not work as expected
//     복잡한 타입 (슬라이스, 맵)은 예상대로 작동하지 않을 수 있음
//
// Example / 예제:
//   // Unique numbers / 고유 숫자
//   v := validation.New([]int{1, 2, 3}, "numbers")
//   v.ArrayUnique()  // Passes / 성공
//
//   v := validation.New([]int{1, 2, 2, 3}, "numbers")
//   v.ArrayUnique()  // Fails (duplicate 2) / 실패 (중복 2)
//
//   // Unique strings / 고유 문자열
//   v := validation.New([]string{"a", "b", "c"}, "tags")
//   v.ArrayUnique()  // Passes / 성공
//
//   // Empty array / 빈 배열
//   v := validation.New([]int{}, "numbers")
//   v.ArrayUnique()  // Passes (no duplicates) / 성공 (중복 없음)
//
//   // Unique IDs / 고유 ID
//   v := validation.New(userIDs, "user_ids")
//   v.ArrayNotEmpty().ArrayUnique()
//
//   // Unique tags / 고유 태그
//   v := validation.New(tags, "tags")
//   v.ArrayUnique().ArrayMaxLength(10)
func (v *Validator) ArrayUnique() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	val := reflect.ValueOf(v.value)
	if val.Kind() != reflect.Slice && val.Kind() != reflect.Array {
		v.addError("arrayunique", fmt.Sprintf("%s must be an array or slice / %s은(는) 배열 또는 슬라이스여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	seen := make(map[interface{}]bool)
	for i := 0; i < val.Len(); i++ {
		item := val.Index(i).Interface()
		if seen[item] {
			v.addError("arrayunique", fmt.Sprintf("%s must contain only unique elements / %s은(는) 고유한 요소만 포함해야 합니다", v.fieldName, v.fieldName))
			return v
		}
		seen[item] = true
	}

	return v
}

// MapHasKey validates that the map contains the specified key.
// Fails if value is not a map type or if the key is not found.
//
// MapHasKey는 맵이 지정된 키를 포함하는지 검증합니다.
// 값이 맵 타입이 아니거나 키를 찾지 못하면 실패합니다.
//
// Parameters / 매개변수:
//   - key: The key that must exist in the map
//     맵에 존재해야 하는 키
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses reflection to check map key existence
//     반사를 사용하여 맵 키 존재 확인
//   - Works with any map type
//     모든 맵 타입에서 작동
//   - Fails if not a map type
//     맵 타입이 아니면 실패
//   - Fails if key not found
//     키를 찾지 못하면 실패
//   - Does not validate the value associated with the key
//     키와 연관된 값은 검증하지 않음
//
// Use Cases / 사용 사례:
//   - Required configuration keys / 필수 설정 키
//   - API request body validation / API 요청 본문 검증
//   - JSON/struct field presence / JSON/구조체 필드 존재
//   - Required metadata / 필수 메타데이터
//   - Config validation / 설정 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1) average for hash map
//     시간 복잡도: 해시 맵 평균 O(1)
//   - Reflection overhead for type checking
//     타입 검사를 위한 반사 오버헤드
//
// Example / 예제:
//   // Single key check / 단일 키 확인
//   v := validation.New(map[string]int{"age": 25}, "data")
//   v.MapHasKey("age")  // Passes / 성공
//
//   v := validation.New(map[string]int{"name": 1}, "data")
//   v.MapHasKey("age")  // Fails / 실패
//
//   // Configuration validation / 설정 검증
//   v := validation.New(config, "config")
//   v.MapHasKey("database_url")
//
//   // API request / API 요청
//   v := validation.New(requestData, "request")
//   v.MapNotEmpty().MapHasKey("user_id")
//
//   // Different key types / 다른 키 타입
//   v := validation.New(map[int]string{1: "one"}, "data")
//   v.MapHasKey(1)  // Passes / 성공
func (v *Validator) MapHasKey(key interface{}) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	val := reflect.ValueOf(v.value)
	if val.Kind() != reflect.Map {
		v.addError("maphaskey", fmt.Sprintf("%s must be a map / %s은(는) 맵이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	keyVal := reflect.ValueOf(key)
	if !val.MapIndex(keyVal).IsValid() {
		v.addError("maphaskey", fmt.Sprintf("%s must contain key '%v' / %s은(는) 키 '%v'를 포함해야 합니다", v.fieldName, key, v.fieldName, key))
	}

	return v
}

// MapHasKeys validates that the map contains all of the specified keys.
// Fails if value is not a map type or if any required key is missing.
//
// MapHasKeys는 맵이 지정된 모든 키를 포함하는지 검증합니다.
// 값이 맵 타입이 아니거나 필수 키가 하나라도 없으면 실패합니다.
//
// Parameters / 매개변수:
//   - keys: Variadic list of keys that must all exist
//     모두 존재해야 하는 키들의 가변 인자 목록
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses reflection to check map key existence
//     반사를 사용하여 맵 키 존재 확인
//   - Works with any map type
//     모든 맵 타입에서 작동
//   - Fails if not a map type
//     맵 타입이 아니면 실패
//   - Fails if any key is missing
//     키가 하나라도 없으면 실패
//   - Lists all missing keys in error message
//     오류 메시지에 모든 누락된 키 나열
//   - Does not validate the values associated with keys
//     키와 연관된 값은 검증하지 않음
//
// Use Cases / 사용 사례:
//   - Required configuration keys / 필수 설정 키들
//   - API request body validation / API 요청 본문 검증
//   - Required struct fields / 필수 구조체 필드들
//   - Complete data validation / 완전한 데이터 검증
//   - Schema validation / 스키마 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(k) where k is number of keys to check
//     시간 복잡도: O(k) (k는 확인할 키의 수)
//   - Each key lookup is O(1) average
//     각 키 조회는 평균 O(1)
//   - Reflection overhead for type checking
//     타입 검사를 위한 반사 오버헤드
//
// Example / 예제:
//   // Multiple keys / 여러 키
//   v := validation.New(map[string]string{"name": "John", "email": "j@e.com"}, "user")
//   v.MapHasKeys("name", "email")  // Passes / 성공
//
//   v := validation.New(map[string]string{"name": "John"}, "user")
//   v.MapHasKeys("name", "email")  // Fails (missing email) / 실패 (email 누락)
//
//   // Configuration validation / 설정 검증
//   v := validation.New(config, "config")
//   v.MapNotEmpty().MapHasKeys("host", "port", "database")
//
//   // API request / API 요청
//   v := validation.New(requestData, "request")
//   v.MapHasKeys("user_id", "action", "timestamp")
//
//   // Error message shows all missing keys
//   // 오류 메시지는 모든 누락된 키 표시
//   data := map[string]int{"a": 1}
//   v := validation.New(data, "data")
//   v.MapHasKeys("a", "b", "c")  // Error: missing keys: [b c]
func (v *Validator) MapHasKeys(keys ...interface{}) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	val := reflect.ValueOf(v.value)
	if val.Kind() != reflect.Map {
		v.addError("maphaskeys", fmt.Sprintf("%s must be a map / %s은(는) 맵이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	var missingKeys []interface{}
	for _, key := range keys {
		keyVal := reflect.ValueOf(key)
		if !val.MapIndex(keyVal).IsValid() {
			missingKeys = append(missingKeys, key)
		}
	}

	if len(missingKeys) > 0 {
		v.addError("maphaskeys", fmt.Sprintf("%s is missing required keys: %v / %s에 필수 키가 없습니다: %v", v.fieldName, missingKeys, v.fieldName, missingKeys))
	}

	return v
}

// MapNotEmpty validates that the map is not empty (has at least one key-value pair).
// Fails if value is not a map type or if the map has no entries.
//
// MapNotEmpty는 맵이 비어있지 않은지 (최소 하나의 키-값 쌍 포함) 검증합니다.
// 값이 맵 타입이 아니거나 맵에 항목이 없으면 실패합니다.
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses reflection to check map length
//     반사를 사용하여 맵 길이 확인
//   - Works with any map type
//     모든 맵 타입에서 작동
//   - Fails if not a map type
//     맵 타입이 아니면 실패
//   - Fails if map length is zero
//     맵 길이가 0이면 실패
//
// Use Cases / 사용 사례:
//   - Required data validation / 필수 데이터 검증
//   - Non-empty configuration / 비어있지 않은 설정
//   - API response validation / API 응답 검증
//   - Metadata presence / 메타데이터 존재
//   - Parameter validation / 매개변수 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Reflection overhead for type checking
//     타입 검사를 위한 반사 오버헤드
//
// Note / 참고:
//   - Does not validate specific keys or values
//     특정 키나 값은 검증하지 않음
//   - Only checks that map has at least one entry
//     맵에 최소 하나의 항목이 있는지만 확인
//   - Use MapHasKeys for specific key validation
//     특정 키 검증은 MapHasKeys 사용
//
// Example / 예제:
//   // Non-empty map / 비어있지 않은 맵
//   v := validation.New(map[string]int{"count": 1}, "data")
//   v.MapNotEmpty()  // Passes / 성공
//
//   v := validation.New(map[string]int{}, "data")
//   v.MapNotEmpty()  // Fails / 실패
//
//   // Configuration validation / 설정 검증
//   v := validation.New(config, "config")
//   v.MapNotEmpty()
//
//   // API request data / API 요청 데이터
//   v := validation.New(requestData, "data")
//   v.MapNotEmpty().MapHasKeys("action")
//
//   // Metadata validation / 메타데이터 검증
//   v := validation.New(metadata, "metadata")
//   v.MapNotEmpty()
//
//   // Nil map check / nil 맵 확인
//   var nilMap map[string]int
//   v := validation.New(nilMap, "data")
//   v.MapNotEmpty()  // Fails (nil map has length 0) / 실패 (nil 맵은 길이 0)
func (v *Validator) MapNotEmpty() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	val := reflect.ValueOf(v.value)
	if val.Kind() != reflect.Map {
		v.addError("mapnotempty", fmt.Sprintf("%s must be a map / %s은(는) 맵이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if val.Len() == 0 {
		v.addError("mapnotempty", fmt.Sprintf("%s must not be empty / %s은(는) 비어있지 않아야 합니다", v.fieldName, v.fieldName))
	}

	return v
}
