package validation

// Validator represents a validation chain for a single value.
// Validator는 단일 값에 대한 검증 체인을 나타냅니다.
type Validator struct {
	value          interface{}
	fieldName      string
	errors         []ValidationError
	stopOnError    bool
	lastRule       string            // Track the last rule for WithMessage
	customMessages map[string]string // Custom messages for specific rules
}

// MultiValidator represents a validator for multiple fields.
// MultiValidator는 여러 필드에 대한 검증기를 나타냅니다.
type MultiValidator struct {
	validators []*Validator
	errors     []ValidationError
}

// RuleFunc is a function type for custom validation rules.
// RuleFunc는 사용자 정의 검증 규칙을 위한 함수 타입입니다.
type RuleFunc func(interface{}) bool

// MessageFunc is a function type for generating custom error messages.
// MessageFunc는 사용자 정의 에러 메시지를 생성하기 위한 함수 타입입니다.
type MessageFunc func(field string, value interface{}) string
