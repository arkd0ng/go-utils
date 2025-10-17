package validation

import (
	"fmt"
)

// ============================================================================
// LOGICAL/CONDITIONAL VALIDATORS
// ============================================================================
//
// This file provides logical and conditional validation functions for:
// - OneOf (value must match one of the provided values)
// - NotOneOf (value must not match any of the provided values)
// - When (conditional validation based on predicate)
// - Unless (conditional validation unless predicate is true)
//
// 이 파일은 다음을 위한 논리 및 조건부 검증 함수를 제공합니다:
// - OneOf (값이 제공된 값 중 하나와 일치해야 함)
// - NotOneOf (값이 제공된 값 중 어느 것과도 일치하지 않아야 함)
// - When (조건부 검증 - 조건이 참일 때)
// - Unless (조건부 검증 - 조건이 거짓일 때)
//
// ============================================================================

// OneOf validates that the value matches one of the provided values.
// Useful for enum-like validation and restricted choice fields.
//
// OneOf는 값이 제공된 값 중 하나와 일치하는지 검증합니다.
// 열거형 검증 및 제한된 선택 필드에 유용합니다.
//
// Parameters / 매개변수:
//   - values: Variadic list of allowed values
//     가변 인자로 전달된 허용 값 목록
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses == operator for comparison
//     == 연산자를 사용하여 비교
//   - Accepts any type (interface{})
//     모든 타입 허용 (interface{})
//   - At least one allowed value required
//     최소 하나의 허용 값 필요
//   - Passes if value equals any provided value
//     값이 제공된 값 중 하나와 같으면 통과
//   - Fails if no match found
//     일치하는 값이 없으면 실패
//
// Use Cases / 사용 사례:
//   - Status field validation / 상태 필드 검증
//   - Enum-like fields / 열거형 필드
//   - Restricted choice validation / 제한된 선택 검증
//   - Role validation / 역할 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n), n = number of allowed values
//     시간 복잡도: O(n), n = 허용 값 개수
//   - Linear search through allowed values
//     허용 값을 통한 선형 검색
//
// This validator checks if the value equals any of the allowed values.
// Useful for validating enum-like fields or restricted choice fields.
// 이 검증기는 값이 허용된 값 중 하나와 같은지 확인합니다.
// 열거형 필드 또는 제한된 선택 필드를 검증하는 데 유용합니다.
//
// Example / 예시:
//
//	// Status validation / 상태 검증
//	status := "active"
//	v := validation.New(status, "status")
//	v.OneOf("active", "inactive", "pending")  // Passes
//
//	// Role validation / 역할 검증
//	v = validation.New("admin", "role")
//	v.OneOf("admin", "user", "guest")  // Passes
//
//	// Numeric values / 숫자 값
//	v = validation.New(2, "priority")
//	v.OneOf(1, 2, 3, 4, 5)  // Passes
//
//	// Invalid value / 무효한 값
//	v = validation.New("banned", "status")
//	v.OneOf("active", "inactive", "pending")  // Fails
//
// Validation rules / 검증 규칙:
//   - Value must equal one of the provided values / 값이 제공된 값 중 하나와 같아야 함
//   - Comparison uses == operator / 비교는 == 연산자 사용
//   - At least one allowed value must be provided / 최소 하나의 허용 값이 제공되어야 함
func (v *Validator) OneOf(values ...interface{}) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	if len(values) == 0 {
		v.addError("one_of", fmt.Sprintf("%s OneOf requires at least one value / %s OneOf는 최소 하나의 값이 필요합니다", v.fieldName, v.fieldName))
		return v
	}

	// Check if value matches any of the provided values
	for _, allowed := range values {
		if v.value == allowed {
			return v // Match found
		}
	}

	// No match found - create error message
	v.addError("one_of", fmt.Sprintf("%s must be one of the allowed values / %s은(는) 허용된 값 중 하나여야 합니다", v.fieldName, v.fieldName))
	return v
}

// NotOneOf validates that the value does not match any of the provided values.
// Useful for blacklisting specific values and preventing reserved keywords.
//
// NotOneOf는 값이 제공된 값 중 어느 것과도 일치하지 않는지 검증합니다.
// 특정 값을 블랙리스트에 추가하고 예약어를 방지하는 데 유용합니다.
//
// Parameters / 매개변수:
//   - values: Variadic list of forbidden values
//     가변 인자로 전달된 금지 값 목록
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses == operator for comparison
//     == 연산자를 사용하여 비교
//   - Accepts any type (interface{})
//     모든 타입 허용 (interface{})
//   - At least one forbidden value required
//     최소 하나의 금지 값 필요
//   - Passes if value differs from all provided values
//     값이 제공된 모든 값과 다르면 통과
//   - Fails if any match found
//     일치하는 값이 있으면 실패
//
// Use Cases / 사용 사례:
//   - Reserved keyword prevention / 예약어 방지
//   - Blacklisted value validation / 블랙리스트 값 검증
//   - Username restriction / 사용자명 제한
//   - Forbidden value checking / 금지 값 확인
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n), n = number of forbidden values
//     시간 복잡도: O(n), n = 금지 값 개수
//   - Linear search through forbidden values
//     금지 값을 통한 선형 검색
//
// This validator checks if the value is different from all forbidden values.
// Useful for blacklisting specific values or preventing reserved keywords.
// 이 검증기는 값이 모든 금지된 값과 다른지 확인합니다.
// 특정 값을 블랙리스트에 추가하거나 예약어를 방지하는 데 유용합니다.
//
// Example / 예시:
//
//	// Reserved username prevention / 예약된 사용자명 방지
//	username := "john"
//	v := validation.New(username, "username")
//	v.NotOneOf("admin", "root", "administrator")  // Passes
//
//	// Forbidden status / 금지된 상태
//	v = validation.New("active", "status")
//	v.NotOneOf("deleted", "banned", "suspended")  // Passes
//
//	// Numeric blacklist / 숫자 블랙리스트
//	v = validation.New(42, "port")
//	v.NotOneOf(0, 80, 443, 8080)  // Passes
//
//	// Invalid - forbidden value / 무효 - 금지 값
//	v = validation.New("admin", "username")
//	v.NotOneOf("admin", "root", "administrator")  // Fails
//
// Validation rules / 검증 규칙:
//   - Value must not equal any of the provided values / 값이 제공된 값 중 어느 것과도 같지 않아야 함
//   - Comparison uses == operator / 비교는 == 연산자 사용
//   - At least one forbidden value must be provided / 최소 하나의 금지 값이 제공되어야 함
func (v *Validator) NotOneOf(values ...interface{}) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	if len(values) == 0 {
		v.addError("not_one_of", fmt.Sprintf("%s NotOneOf requires at least one value / %s NotOneOf는 최소 하나의 값이 필요합니다", v.fieldName, v.fieldName))
		return v
	}

	// Check if value matches any of the forbidden values
	for _, forbidden := range values {
		if v.value == forbidden {
			v.addError("not_one_of", fmt.Sprintf("%s must not be one of the forbidden values / %s은(는) 금지된 값 중 하나가 아니어야 합니다", v.fieldName, v.fieldName))
			return v
		}
	}

	return v // Value is not in forbidden list
}

// When executes the validation function only if the predicate returns true.
// Enables conditional validation based on runtime conditions.
//
// When은 조건이 참을 반환할 때만 검증 함수를 실행합니다.
// 런타임 조건에 따른 조건부 검증을 활성화합니다.
//
// Parameters / 매개변수:
//   - predicate: Boolean condition determining if validation should run
//     검증 실행 여부를 결정하는 부울 조건
//   - fn: Validation function to execute if predicate is true
//     조건이 참일 때 실행할 검증 함수
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - If predicate is true, executes fn(v)
//     조건이 참이면 fn(v) 실행
//   - If predicate is false, skips validation
//     조건이 거짓이면 검증 건너뜀
//   - fn receives validator instance
//     fn은 검증기 인스턴스를 받음
//   - Supports nested validation chains
//     중첩된 검증 체인 지원
//
// Use Cases / 사용 사례:
//   - Conditional field validation / 조건부 필드 검증
//   - Role-based validation / 역할 기반 검증
//   - Feature flag validation / 기능 플래그 검증
//   - Dynamic validation rules / 동적 검증 규칙
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1) + validation function cost
//     시간 복잡도: O(1) + 검증 함수 비용
//   - Single predicate check
//     단일 조건 확인
//
// This validator allows conditional validation based on runtime conditions.
// The validation function is only executed if the predicate evaluates to true.
// 이 검증기는 런타임 조건에 따라 조건부 검증을 허용합니다.
// 검증 함수는 조건이 참으로 평가될 때만 실행됩니다.
//
// Example / 예시:
//
//	// Age-based validation / 나이 기반 검증
//	age := 25
//	isAdult := age >= 18
//	v := validation.New(age, "age")
//	v.When(isAdult, func(val *Validator) {
//	    val.Min(18).Max(120)
//	})  // Passes (validated because isAdult is true)
//
//	// Role-based email requirement / 역할 기반 이메일 필수
//	email := "user@example.com"
//	requireEmail := role == "registered"
//	v = validation.New(email, "email")
//	v.When(requireEmail, func(val *Validator) {
//	    val.Required().Email()
//	})
//
//	// Feature flag validation / 기능 플래그 검증
//	v = validation.New(value, "feature")
//	v.When(featureEnabled, func(val *Validator) {
//	    val.MinLength(10).MaxLength(100)
//	})
//
//	// Skipped validation / 건너뛴 검증
//	v = validation.New(10, "age")
//	v.When(false, func(val *Validator) {
//	    val.Min(18)  // Not executed (predicate is false)
//	})  // Passes (validation skipped)
//
// Validation rules / 검증 규칙:
//   - If predicate is true, validation function is executed / 조건이 참이면 검증 함수 실행
//   - If predicate is false, validation is skipped / 조건이 거짓이면 검증 건너뜀
//   - Useful for complex conditional validation logic / 복잡한 조건부 검증 로직에 유용
func (v *Validator) When(predicate bool, fn func(*Validator)) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	if predicate {
		fn(v)
	}

	return v
}

// Unless executes the validation function only if the predicate returns false.
// Inverse of When - validates when condition is NOT met.
//
// Unless는 조건이 거짓을 반환할 때만 검증 함수를 실행합니다.
// When의 반대 - 조건이 충족되지 않을 때 검증합니다.
//
// Parameters / 매개변수:
//   - predicate: Boolean condition determining if validation should be skipped
//     검증을 건너뛸지 결정하는 부울 조건
//   - fn: Validation function to execute if predicate is false
//     조건이 거짓일 때 실행할 검증 함수
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - If predicate is false, executes fn(v)
//     조건이 거짓이면 fn(v) 실행
//   - If predicate is true, skips validation
//     조건이 참이면 검증 건너뜀
//   - Inverse logic of When validator
//     When 검증기의 반대 논리
//   - fn receives validator instance
//     fn은 검증기 인스턴스를 받음
//   - Supports nested validation chains
//     중첩된 검증 체인 지원
//
// Use Cases / 사용 사례:
//   - "Validate unless X" scenarios / "X가 아닌 경우 검증" 시나리오
//   - Guest vs. registered user validation / 게스트 vs. 등록 사용자 검증
//   - Optional field validation / 선택적 필드 검증
//   - Exemption-based validation / 면제 기반 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1) + validation function cost
//     시간 복잡도: O(1) + 검증 함수 비용
//   - Single predicate check
//     단일 조건 확인
//
// This validator is the inverse of When - it executes validation when condition is false.
// Useful for "validate unless X" scenarios.
// 이 검증기는 When의 반대입니다 - 조건이 거짓일 때 검증을 실행합니다.
// "X가 아닌 경우 검증" 시나리오에 유용합니다.
//
// Example / 예시:
//
//	// Email required unless guest / 게스트가 아닌 경우 이메일 필수
//	email := "user@example.com"
//	isGuest := false
//	v := validation.New(email, "email")
//	v.Unless(isGuest, func(val *Validator) {
//	    val.Required().Email()
//	})  // Passes (validated because not guest)
//
//	// Skip validation for admin / 관리자는 검증 건너뛰기
//	v = validation.New(username, "username")
//	v.Unless(isAdmin, func(val *Validator) {
//	    val.MinLength(5).MaxLength(20)
//	})
//
//	// Guest can skip email validation / 게스트는 이메일 검증 건너뛰기
//	v = validation.New("", "email")
//	v.Unless(true, func(val *Validator) {
//	    val.Required().Email()  // Not executed (predicate is true)
//	})  // Passes (validation skipped for guests)
//
//	// Validate unless premium user / 프리미엄 사용자가 아닌 경우 검증
//	isPremium := false
//	v = validation.New(uploadSize, "file_size")
//	v.Unless(isPremium, func(val *Validator) {
//	    val.Max(10485760)  // 10MB limit for non-premium
//	})
//
// Validation rules / 검증 규칙:
//   - If predicate is false, validation function is executed / 조건이 거짓이면 검증 함수 실행
//   - If predicate is true, validation is skipped / 조건이 참이면 검증 건너뜀
//   - Inverse of When validator / When 검증기의 반대
func (v *Validator) Unless(predicate bool, fn func(*Validator)) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	if !predicate {
		fn(v)
	}

	return v
}
