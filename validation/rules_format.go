package validation

import (
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"
)

// UUIDv4 checks if the value is a valid UUID version 4.
// UUIDv4는 값이 유효한 UUID 버전 4인지 확인합니다.
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

// XML checks if the value is a valid XML string.
// XML은 값이 유효한 XML 문자열인지 확인합니다.
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

// Hex checks if the value is a valid hexadecimal string.
// Hex는 값이 유효한 16진수 문자열인지 확인합니다.
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
