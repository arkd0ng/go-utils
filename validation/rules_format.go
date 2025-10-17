package validation

import (
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"
)

// UUIDv4 validates that the value is a valid UUID version 4.
// Specifically checks for UUID v4 format with correct version and variant bits.
//
// UUIDv4는 값이 유효한 UUID 버전 4인지 검증합니다.
// 올바른 버전 및 변형 비트를 가진 UUID v4 형식을 구체적으로 확인합니다.
//
// Parameters / 매개변수:
//   - None (operates on validator's value)
//     없음 (validator의 값에 대해 작동)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Validates UUID v4 format (8-4-4-4-12 hex digits)
//     UUID v4 형식 검증 (8-4-4-4-12 16진수)
//   - Version field must be '4' (13th character after hyphens)
//     버전 필드는 '4'여야 함 (하이픈 이후 13번째 문자)
//   - Variant field must be 8, 9, a, or b
//     변형 필드는 8, 9, a 또는 b여야 함
//   - Case-insensitive hex validation
//     대소문자 구분 없는 16진수 검증
//   - Fails if value is not string
//     값이 문자열이 아니면 실패
//
// Use Cases / 사용 사례:
//   - Unique identifier validation / 고유 식별자 검증
//   - API key validation / API 키 검증
//   - Resource ID validation / 리소스 ID 검증
//   - Database key validation / 데이터베이스 키 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n), n = string length
//     시간 복잡도: O(n), n = 문자열 길이
//   - Single regex match
//     단일 정규식 매칭
//
// UUID v4 Format / UUID v4 형식:
//   - Pattern: xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx
//   - Version bit: '4' (13th position)
//   - Variant bit: 8, 9, a, or b (17th position)
//
// Example / 예제:
//
//	// Valid UUID v4 / 유효한 UUID v4
//	uuid := "550e8400-e29b-41d4-a716-446655440000"
//	v := validation.New(uuid, "user_id")
//	v.UUIDv4()  // Passes
//
//	// Lowercase / 소문자
//	v = validation.New("f47ac10b-58cc-4372-a567-0e02b2c3d479", "id")
//	v.UUIDv4()  // Passes
//
//	// Uppercase / 대문자
//	v = validation.New("F47AC10B-58CC-4372-A567-0E02B2C3D479", "id")
//	v.UUIDv4()  // Passes
//
//	// Invalid - wrong version / 무효 - 잘못된 버전
//	v = validation.New("550e8400-e29b-31d4-a716-446655440000", "id")
//	v.UUIDv4()  // Fails (version 3, not 4)
//
//	// Invalid - wrong variant / 무효 - 잘못된 변형
//	v = validation.New("550e8400-e29b-41d4-f716-446655440000", "id")
//	v.UUIDv4()  // Fails (variant f, not 8-b)
func (v *Validator) UUIDv4() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("uuid_v4", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// UUIDv4 regex pattern (version 4 has '4' in the version position)
	// Format: xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx where y is 8, 9, a, or b
	uuidv4Pattern := `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`
	matched, _ := regexp.MatchString(uuidv4Pattern, str)

	if !matched {
		v.addError("uuid_v4", fmt.Sprintf("%s must be a valid UUID v4 / %s은(는) 유효한 UUID v4여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	return v
}

// XML validates that the value is a valid XML string.
// Uses Go's xml.Unmarshal to verify XML structure validity.
//
// XML은 값이 유효한 XML 문자열인지 검증합니다.
// Go의 xml.Unmarshal을 사용하여 XML 구조 유효성을 확인합니다.
//
// Parameters / 매개변수:
//   - None (operates on validator's value)
//     없음 (validator의 값에 대해 작동)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Trims whitespace before validation
//     검증 전 공백 제거
//   - Uses xml.Unmarshal for parsing
//     xml.Unmarshal을 사용하여 파싱
//   - Validates XML structure and syntax
//     XML 구조 및 구문 검증
//   - Checks well-formedness
//     올바른 형식 확인
//   - Fails if value is not string
//     값이 문자열이 아니면 실패
//
// Use Cases / 사용 사례:
//   - XML configuration validation / XML 구성 검증
//   - SOAP message validation / SOAP 메시지 검증
//   - XML data input validation / XML 데이터 입력 검증
//   - API request/response validation / API 요청/응답 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n), n = XML length
//     시간 복잡도: O(n), n = XML 길이
//   - Full XML parsing required
//     전체 XML 파싱 필요
//
// Example / 예제:
//
//	// Valid XML / 유효한 XML
//	xmlStr := "<root><item>value</item></root>"
//	v := validation.New(xmlStr, "xml_data")
//	v.XML()  // Passes
//
//	// With attributes / 속성 포함
//	v = validation.New(`<root id="1"><item>value</item></root>`, "data")
//	v.XML()  // Passes
//
//	// With whitespace / 공백 포함
//	v = validation.New("  <root>text</root>  ", "xml")
//	v.XML()  // Passes (whitespace trimmed)
//
//	// Self-closing tag / 자체 닫기 태그
//	v = validation.New("<element />", "element")
//	v.XML()  // Passes
//
//	// Invalid - unclosed tag / 무효 - 닫히지 않은 태그
//	v = validation.New("<root><item>value</root>", "xml")
//	v.XML()  // Fails (mismatched tags)
//
//	// Invalid - not XML / 무효 - XML 아님
//	v = validation.New("plain text", "xml")
//	v.XML()  // Fails
func (v *Validator) XML() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("xml", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Trim whitespace for validation
	str = strings.TrimSpace(str)

	// Check if it's a valid XML by attempting to unmarshal
	var xmlData interface{}
	if err := xml.Unmarshal([]byte(str), &xmlData); err != nil {
		v.addError("xml", fmt.Sprintf("%s must be valid XML / %s은(는) 유효한 XML이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	return v
}

// Hex validates that the value is a valid hexadecimal string.
// Supports optional '0x' or '0X' prefix and case-insensitive hex digits.
//
// Hex는 값이 유효한 16진수 문자열인지 검증합니다.
// 선택적 '0x' 또는 '0X' 접두사와 대소문자 구분 없는 16진수를 지원합니다.
//
// Parameters / 매개변수:
//   - None (operates on validator's value)
//     없음 (validator의 값에 대해 작동)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Removes optional '0x' or '0X' prefix
//     선택적 '0x' 또는 '0X' 접두사 제거
//   - Uses hex.DecodeString for validation
//     hex.DecodeString을 사용하여 검증
//   - Validates hex digits (0-9, A-F, a-f)
//     16진수 검증 (0-9, A-F, a-f)
//   - Requires even number of characters
//     짝수 개의 문자 필요
//   - Case-insensitive
//     대소문자 구분 없음
//   - Fails if value is not string
//     값이 문자열이 아니면 실패
//
// Use Cases / 사용 사례:
//   - Hex encoded data validation / 16진수 인코딩 데이터 검증
//   - Binary data representation / 바이너리 데이터 표현
//   - Cryptographic hash validation / 암호화 해시 검증
//   - Memory address validation / 메모리 주소 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n), n = string length
//     시간 복잡도: O(n), n = 문자열 길이
//   - Single hex.DecodeString call
//     단일 hex.DecodeString 호출
//
// Example / 예제:
//
//	// Valid hex strings / 유효한 16진수 문자열
//	v := validation.New("48656c6c6f", "hex_data")
//	v.Hex()  // Passes (even length)
//
//	// With 0x prefix / 0x 접두사 포함
//	v = validation.New("0x48656c6c6f", "hex")
//	v.Hex()  // Passes (prefix removed)
//
//	// Uppercase / 대문자
//	v = validation.New("ABCDEF123456", "hex")
//	v.Hex()  // Passes
//
//	// Mixed case / 대소문자 혼합
//	v = validation.New("AbCdEf123456", "hex")
//	v.Hex()  // Passes
//
//	// With 0X prefix / 0X 접두사 포함
//	v = validation.New("0XABCDEF", "hex")
//	v.Hex()  // Passes
//
//	// Invalid - odd length / 무효 - 홀수 길이
//	v = validation.New("ABC", "hex")
//	v.Hex()  // Fails (must be even length)
//
//	// Invalid - non-hex character / 무효 - 16진수가 아닌 문자
//	v = validation.New("GHIJKL", "hex")
//	v.Hex()  // Fails (G-L not hex)
func (v *Validator) Hex() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("hex", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Remove optional 0x prefix
	str = strings.TrimPrefix(str, "0x")
	str = strings.TrimPrefix(str, "0X")

	// Check if it's valid hexadecimal
	if _, err := hex.DecodeString(str); err != nil {
		v.addError("hex", fmt.Sprintf("%s must be valid hexadecimal / %s은(는) 유효한 16진수여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	return v
}
