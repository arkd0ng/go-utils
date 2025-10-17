package validation

// Validator represents a validation chain for a single value.
// It maintains validation state, error collection, and configuration for a field.
//
// Validator는 단일 값에 대한 검증 체인을 나타냅니다.
// 필드에 대한 검증 상태, 에러 컬렉션 및 구성을 유지합니다.
//
// Lifecycle / 생명주기:
//   1. Creation: Use New() to create a new validator instance
//      생성: New()를 사용하여 새 검증기 인스턴스 생성
//   2. Rule Chaining: Chain validation rules using fluent API
//      규칙 체이닝: 유창한 API를 사용하여 검증 규칙 체이닝
//   3. Execution: Rules execute immediately when called
//      실행: 호출 시 규칙이 즉시 실행
//   4. Result: Call Validate() or GetErrors() to retrieve results
//      결과: Validate() 또는 GetErrors()를 호출하여 결과 검색
//
// State Management / 상태 관리:
//   - Each validator maintains independent state
//     각 검증기는 독립적인 상태 유지
//   - Errors accumulate unless StopOnError is set
//     StopOnError가 설정되지 않으면 에러 누적
//   - Custom messages can override default error messages
//     사용자 정의 메시지로 기본 에러 메시지 재정의 가능
//
// Thread Safety / 스레드 안전성:
//   - Not safe for concurrent modification
//     동시 수정 안전하지 않음
//   - Safe to create multiple validators concurrently
//     여러 검증기를 동시에 생성하는 것은 안전
//   - Do not share validator instances across goroutines
//     고루틴 간 검증기 인스턴스 공유하지 말 것
//
// Memory / 메모리:
//   - Small memory footprint (~100 bytes + error collection)
//     작은 메모리 공간 (~100바이트 + 에러 컬렉션)
//   - Error collection grows with failed validations
//     에러 컬렉션은 실패한 검증과 함께 증가
//   - Can be reused after calling Validate()
//     Validate() 호출 후 재사용 가능
//
// Example / 예제:
//   // Create and use validator / 검증기 생성 및 사용
//   v := validation.New("test@example.com", "email")
//   v.Required().Email().MaxLength(100)
//   
//   if err := v.Validate(); err != nil {
//       // Handle validation failure / 검증 실패 처리
//   }
//
//   // With custom configuration / 사용자 정의 구성과 함께
//   v := validation.New(data, "data").StopOnError()
//   v.Required().WithMessage("Data is required")
type Validator struct {
	value          interface{}       // Value being validated / 검증 중인 값
	fieldName      string            // Field name for error messages / 에러 메시지용 필드 이름
	errors         []ValidationError // Collected validation errors / 수집된 검증 에러
	stopOnError    bool              // Stop on first error flag / 첫 에러에서 중단 플래그
	lastRule       string            // Track the last rule for WithMessage / WithMessage용 마지막 규칙 추적
	customMessages map[string]string // Custom messages for specific rules / 특정 규칙용 사용자 정의 메시지
}

// MultiValidator represents a validator for multiple fields.
// It allows validation of multiple fields with independent error collection per field.
//
// MultiValidator는 여러 필드에 대한 검증기를 나타냅니다.
// 필드별 독립적인 에러 컬렉션으로 여러 필드의 검증을 허용합니다.
//
// Use Cases / 사용 사례:
//   - Form validation with multiple inputs
//     여러 입력이 있는 폼 검증
//   - API request validation with multiple parameters
//     여러 매개변수가 있는 API 요청 검증
//   - Complex object validation with nested fields
//     중첩 필드가 있는 복잡한 객체 검증
//   - Batch validation of related fields
//     관련 필드의 배치 검증
//
// Thread Safety / 스레드 안전성:
//   - Not safe for concurrent use
//     동시 사용 안전하지 않음
//   - Create separate instances for concurrent validation
//     동시 검증을 위해 별도 인스턴스 생성
//
// Example / 예제:
//   mv := validation.NewValidator()
//   mv.Field(user.Email, "email").Required().Email()
//   mv.Field(user.Age, "age").Required().Min(18).Max(120)
//   mv.Field(user.Name, "name").Required().MinLength(2)
//   
//   if err := mv.Validate(); err != nil {
//       // All field errors collected / 모든 필드 에러 수집됨
//   }
type MultiValidator struct {
	validators []*Validator      // Individual field validators / 개별 필드 검증기
	errors     []ValidationError // Aggregated errors from all fields / 모든 필드의 집계된 에러
}

// RuleFunc is a function type for custom validation rules.
// It receives a value and returns true if valid, false otherwise.
//
// RuleFunc는 사용자 정의 검증 규칙을 위한 함수 타입입니다.
// 값을 받아 유효하면 true, 아니면 false를 반환합니다.
//
// Parameters / 매개변수:
//   - interface{}: Value to validate (can be any type)
//     검증할 값 (모든 타입 가능)
//
// Returns / 반환값:
//   - bool: true if validation passes, false if fails
//     검증 통과 시 true, 실패 시 false
//
// Usage / 사용법:
//   - Used with Custom() method for business-specific validation
//     비즈니스별 검증을 위해 Custom() 메서드와 함께 사용
//   - Should perform type assertion if specific type expected
//     특정 타입이 예상되는 경우 타입 단언 수행해야 함
//   - Should return false for unexpected types
//     예상치 못한 타입에 대해 false 반환해야 함
//
// Example / 예제:
//   // Simple custom rule / 간단한 사용자 정의 규칙
//   isPositive := func(val interface{}) bool {
//       if num, ok := val.(int); ok {
//           return num > 0
//       }
//       return false
//   }
//   
//   v := validation.New(value, "number")
//   v.Custom(isPositive).WithMessage("Must be positive")
//
//   // Complex business rule / 복잡한 비즈니스 규칙
//   checkAvailability := func(val interface{}) bool {
//       username := val.(string)
//       return !db.UserExists(username)
//   }
//   
//   v := validation.New(username, "username")
//   v.Custom(checkAvailability).WithMessage("Username already taken")
type RuleFunc func(interface{}) bool

// MessageFunc is a function type for generating custom error messages.
// It receives field name and value, returning a custom error message.
//
// MessageFunc는 사용자 정의 에러 메시지를 생성하기 위한 함수 타입입니다.
// 필드 이름과 값을 받아 사용자 정의 에러 메시지를 반환합니다.
//
// Parameters / 매개변수:
//   - field: Name of the field that failed validation
//     검증 실패한 필드 이름
//   - value: The actual value that failed validation
//     검증 실패한 실제 값
//
// Returns / 반환값:
//   - string: Custom error message
//     사용자 정의 에러 메시지
//
// Usage / 사용법:
//   - Used for dynamic error message generation
//     동적 에러 메시지 생성에 사용
//   - Allows context-aware error messages
//     컨텍스트 인식 에러 메시지 허용
//   - Can include field value in message
//     메시지에 필드 값 포함 가능
//
// Example / 예제:
//   msgFunc := func(field string, value interface{}) string {
//       return fmt.Sprintf("Field '%s' with value '%v' is invalid", field, value)
//   }
//   
//   // Use with custom validation / 사용자 정의 검증과 함께 사용
//   v := validation.New(data, "field")
//   v.Custom(customRule).WithMessage(msgFunc("field", data))
type MessageFunc func(field string, value interface{}) string
