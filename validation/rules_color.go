package validation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ============================================================================
// COLOR/CSS VALIDATORS
// ============================================================================
//
// This file provides color and CSS-related validation functions for:
// - HexColor (3 or 6 digit hex color codes)
// - RGB (Red, Green, Blue color format)
// - RGBA (RGB with Alpha channel)
// - HSL (Hue, Saturation, Lightness color format)
//
// 이 파일은 다음을 위한 색상 및 CSS 관련 검증 함수를 제공합니다:
// - HexColor (3자리 또는 6자리 16진수 색상 코드)
// - RGB (빨강, 초록, 파랑 색상 형식)
// - RGBA (알파 채널이 있는 RGB)
// - HSL (색조, 채도, 명도 색상 형식)
//
// ============================================================================

// HexColor validates that the value is a valid hexadecimal color code.
// Supports 3-digit (#RGB) and 6-digit (#RRGGBB) formats, with optional # prefix.
//
// HexColor는 값이 유효한 16진수 색상 코드인지 검증합니다.
// 3자리(#RGB) 및 6자리(#RRGGBB) 형식을 지원하며, # 접두사는 선택 사항입니다.
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
//   - Accepts 3-digit hex (#RGB or RGB)
//     3자리 16진수 허용 (#RGB 또는 RGB)
//   - Accepts 6-digit hex (#RRGGBB or RRGGBB)
//     6자리 16진수 허용 (#RRGGBB 또는 RRGGBB)
//   - # prefix is optional
//     # 접두사 선택 사항
//   - Case-insensitive (A-F or a-f)
//     대소문자 구분 없음 (A-F 또는 a-f)
//   - Fails if value is not string
//     값이 문자열이 아니면 실패
//
// Use Cases / 사용 사례:
//   - Web color validation / 웹 색상 검증
//   - CSS color input / CSS 색상 입력
//   - Design system colors / 디자인 시스템 색상
//   - Brand color validation / 브랜드 색상 검증
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
// Supports both 3-digit (#RGB) and 6-digit (#RRGGBB) hex colors, with or without # prefix.
// 3자리(#RGB) 및 6자리(#RRGGBB) 16진수 색상을 모두 지원하며, # 접두사는 선택 사항입니다.
//
// Example / 예시:
//
//	// 6-digit with # / # 포함 6자리
//	color := "#FF5733"
//	v := validation.New(color, "brand_color")
//	v.HexColor()  // Passes
//
//	// 6-digit without # / # 없는 6자리
//	v = validation.New("FF5733", "color")
//	v.HexColor()  // Passes
//
//	// 3-digit shorthand / 3자리 축약형
//	v = validation.New("#F53", "color")
//	v.HexColor()  // Passes (equivalent to #FF5533)
//
//	// Case-insensitive / 대소문자 구분 없음
//	v = validation.New("#ff5733", "color")
//	v.HexColor()  // Passes
//
//	// Invalid formats / 무효한 형식
//	v = validation.New("#GG5733", "color")
//	v.HexColor()  // Fails (invalid hex)
//
//	v = validation.New("#F5", "color")
//	v.HexColor()  // Fails (2 digits, not 3 or 6)
//
// Validation rules / 검증 규칙:
//   - Must be a string / 문자열이어야 함
//   - Format: #RGB or #RRGGBB (# is optional) / 형식: #RGB 또는 #RRGGBB (# 선택)
//   - RGB values must be hexadecimal (0-9, A-F) / RGB 값은 16진수여야 함
func (v *Validator) HexColor() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("hex_color", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Remove # prefix if present
	color := strings.TrimPrefix(str, "#")

	// Validate 3-digit or 6-digit hex color
	hexColorRegex := regexp.MustCompile(`^([A-Fa-f0-9]{3}|[A-Fa-f0-9]{6})$`)
	if !hexColorRegex.MatchString(color) {
		v.addError("hex_color", fmt.Sprintf("%s must be a valid hex color (#RGB or #RRGGBB) / %s은(는) 유효한 16진수 색상이어야 합니다 (#RGB 또는 #RRGGBB)", v.fieldName, v.fieldName))
	}

	return v
}

// RGB validates that the value is a valid RGB color format.
// Format: rgb(r, g, b) where r, g, b are integers 0-255.
//
// RGB는 값이 유효한 RGB 색상 형식인지 검증합니다.
// 형식: rgb(r, g, b) 여기서 r, g, b는 0-255 범위의 정수입니다.
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
//   - Validates rgb(r, g, b) format
//     rgb(r, g, b) 형식 검증
//   - Red component: 0-255
//     빨강 구성 요소: 0-255
//   - Green component: 0-255
//     초록 구성 요소: 0-255
//   - Blue component: 0-255
//     파랑 구성 요소: 0-255
//   - Spaces around commas optional
//     쉼표 주변 공백 선택 사항
//   - Fails if value is not string
//     값이 문자열이 아니면 실패
//
// Use Cases / 사용 사례:
//   - CSS rgb() color validation / CSS rgb() 색상 검증
//   - Web styling / 웹 스타일링
//   - Color picker values / 색상 선택기 값
//   - Design tool output / 디자인 도구 출력
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n), n = string length
//     시간 복잡도: O(n), n = 문자열 길이
//   - Single regex match + 3 range checks
//     단일 정규식 매칭 + 3번의 범위 확인
//
// Format: rgb(r, g, b) where r, g, b are integers 0-255.
// 형식: rgb(r, g, b) 여기서 r, g, b는 0-255 범위의 정수입니다.
//
// Example / 예시:
//
//	// Valid RGB colors / 유효한 RGB 색상
//	color := "rgb(255, 87, 51)"
//	v := validation.New(color, "background_color")
//	v.RGB()  // Passes
//
//	// With extra spaces / 추가 공백 포함
//	v = validation.New("rgb( 255 , 87 , 51 )", "color")
//	v.RGB()  // Passes (spaces allowed)
//
//	// Black / 검정색
//	v = validation.New("rgb(0, 0, 0)", "color")
//	v.RGB()  // Passes
//
//	// White / 흰색
//	v = validation.New("rgb(255, 255, 255)", "color")
//	v.RGB()  // Passes
//
//	// Invalid values / 무효한 값
//	v = validation.New("rgb(256, 87, 51)", "color")
//	v.RGB()  // Fails (red > 255)
//
//	v = validation.New("rgb(255, 87)", "color")
//	v.RGB()  // Fails (missing blue)
//
//	v = validation.New("255, 87, 51", "color")
//	v.RGB()  // Fails (missing rgb() wrapper)
//
// Validation rules / 검증 규칙:
//   - Must be a string / 문자열이어야 함
//   - Format: rgb(r, g, b) / 형식: rgb(r, g, b)
//   - Each component must be 0-255 / 각 구성 요소는 0-255 범위
//   - Spaces around commas are optional / 쉼표 주변 공백은 선택 사항
func (v *Validator) RGB() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("rgb", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Extract values from rgb(r, g, b) format
	rgbRegex := regexp.MustCompile(`^rgb\s*\(\s*(\d+)\s*,\s*(\d+)\s*,\s*(\d+)\s*\)$`)
	matches := rgbRegex.FindStringSubmatch(str)

	if matches == nil {
		v.addError("rgb", fmt.Sprintf("%s must be a valid RGB format (rgb(r, g, b)) / %s은(는) 유효한 RGB 형식이어야 합니다 (rgb(r, g, b))", v.fieldName, v.fieldName))
		return v
	}

	// Validate each component is 0-255
	for i := 1; i <= 3; i++ {
		val, _ := strconv.Atoi(matches[i])
		if val < 0 || val > 255 {
			component := []string{"red", "green", "blue"}[i-1]
			v.addError("rgb", fmt.Sprintf("%s RGB %s value must be between 0 and 255 / %s RGB %s 값은 0과 255 사이여야 합니다", v.fieldName, component, v.fieldName, component))
			return v
		}
	}

	return v
}

// RGBA validates that the value is a valid RGBA color format.
// Format: rgba(r, g, b, a) where r, g, b are 0-255 integers and a is 0-1 float.
//
// RGBA는 값이 유효한 RGBA 색상 형식인지 검증합니다.
// 형식: rgba(r, g, b, a) 여기서 r, g, b는 0-255 정수이고 a는 0-1 부동소수점입니다.
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
//   - Validates rgba(r, g, b, a) format
//     rgba(r, g, b, a) 형식 검증
//   - Red, Green, Blue: 0-255 (integers)
//     빨강, 초록, 파랑: 0-255 (정수)
//   - Alpha: 0.0-1.0 (float, opacity)
//     알파: 0.0-1.0 (부동소수점, 불투명도)
//   - Spaces around commas optional
//     쉼표 주변 공백 선택 사항
//   - Fails if value is not string
//     값이 문자열이 아니면 실패
//
// Use Cases / 사용 사례:
//   - CSS rgba() color validation / CSS rgba() 색상 검증
//   - Transparency effects / 투명도 효과
//   - Overlay colors / 오버레이 색상
//   - Semi-transparent UI elements / 반투명 UI 요소
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n), n = string length
//     시간 복잡도: O(n), n = 문자열 길이
//   - Single regex match + 4 range checks
//     단일 정규식 매칭 + 4번의 범위 확인
//
// Format: rgba(r, g, b, a) where r, g, b are integers 0-255 and a is 0-1.
// 형식: rgba(r, g, b, a) 여기서 r, g, b는 0-255 정수이고 a는 0-1 범위입니다.
//
// Example / 예시:
//
//	// Valid RGBA colors / 유효한 RGBA 색상
//	color := "rgba(255, 87, 51, 0.8)"
//	v := validation.New(color, "overlay_color")
//	v.RGBA()  // Passes (80% opaque)
//
//	// Fully opaque / 완전 불투명
//	v = validation.New("rgba(255, 87, 51, 1)", "color")
//	v.RGBA()  // Passes (alpha = 1.0)
//
//	// Fully transparent / 완전 투명
//	v = validation.New("rgba(255, 87, 51, 0)", "color")
//	v.RGBA()  // Passes (alpha = 0.0)
//
//	// Semi-transparent / 반투명
//	v = validation.New("rgba(0, 0, 0, 0.5)", "shadow")
//	v.RGBA()  // Passes (50% opacity)
//
//	// With spaces / 공백 포함
//	v = validation.New("rgba( 255 , 87 , 51 , 0.8 )", "color")
//	v.RGBA()  // Passes
//
//	// Invalid values / 무효한 값
//	v = validation.New("rgba(256, 87, 51, 0.8)", "color")
//	v.RGBA()  // Fails (red > 255)
//
//	v = validation.New("rgba(255, 87, 51, 1.5)", "color")
//	v.RGBA()  // Fails (alpha > 1.0)
//
//	v = validation.New("rgba(255, 87, 51)", "color")
//	v.RGBA()  // Fails (missing alpha)
//
// Validation rules / 검증 규칙:
//   - Must be a string / 문자열이어야 함
//   - Format: rgba(r, g, b, a) / 형식: rgba(r, g, b, a)
//   - RGB components must be 0-255 / RGB 구성 요소는 0-255 범위
//   - Alpha must be 0-1 (0.0 to 1.0) / 알파는 0-1 범위 (0.0 ~ 1.0)
func (v *Validator) RGBA() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("rgba", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Extract values from rgba(r, g, b, a) format
	rgbaRegex := regexp.MustCompile(`^rgba\s*\(\s*(\d+)\s*,\s*(\d+)\s*,\s*(\d+)\s*,\s*([\d.]+)\s*\)$`)
	matches := rgbaRegex.FindStringSubmatch(str)

	if matches == nil {
		v.addError("rgba", fmt.Sprintf("%s must be a valid RGBA format (rgba(r, g, b, a)) / %s은(는) 유효한 RGBA 형식이어야 합니다 (rgba(r, g, b, a))", v.fieldName, v.fieldName))
		return v
	}

	// Validate RGB components (0-255)
	for i := 1; i <= 3; i++ {
		val, _ := strconv.Atoi(matches[i])
		if val < 0 || val > 255 {
			component := []string{"red", "green", "blue"}[i-1]
			v.addError("rgba", fmt.Sprintf("%s RGBA %s value must be between 0 and 255 / %s RGBA %s 값은 0과 255 사이여야 합니다", v.fieldName, component, v.fieldName, component))
			return v
		}
	}

	// Validate alpha component (0-1)
	alpha, err := strconv.ParseFloat(matches[4], 64)
	if err != nil || alpha < 0.0 || alpha > 1.0 {
		v.addError("rgba", fmt.Sprintf("%s RGBA alpha value must be between 0 and 1 / %s RGBA 알파 값은 0과 1 사이여야 합니다", v.fieldName, v.fieldName))
	}

	return v
}

// HSL validates that the value is a valid HSL color format.
// Format: hsl(h, s%, l%) where h is 0-360 degrees, s and l are 0-100%.
//
// HSL은 값이 유효한 HSL 색상 형식인지 검증합니다.
// 형식: hsl(h, s%, l%) 여기서 h는 0-360도, s와 l은 0-100% 범위입니다.
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
//   - Validates hsl(h, s%, l%) format
//     hsl(h, s%, l%) 형식 검증
//   - Hue: 0-360 degrees (color wheel position)
//     색조: 0-360도 (색상환 위치)
//   - Saturation: 0-100% (color intensity)
//     채도: 0-100% (색상 강도)
//   - Lightness: 0-100% (brightness)
//     명도: 0-100% (밝기)
//   - % symbol required for s and l
//     s와 l에 % 기호 필수
//   - Fails if value is not string
//     값이 문자열이 아니면 실패
//
// Use Cases / 사용 사례:
//   - CSS hsl() color validation / CSS hsl() 색상 검증
//   - Design systems / 디자인 시스템
//   - Color harmony calculations / 색상 조화 계산
//   - Theme generation / 테마 생성
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n), n = string length
//     시간 복잡도: O(n), n = 문자열 길이
//   - Single regex match + 3 range checks
//     단일 정규식 매칭 + 3번의 범위 확인
//
// Format: hsl(h, s%, l%) where h is 0-360, s and l are 0-100%.
// 형식: hsl(h, s%, l%) 여기서 h는 0-360, s와 l은 0-100% 범위입니다.
//
// Example / 예시:
//
//	// Valid HSL colors / 유효한 HSL 색상
//	color := "hsl(9, 100%, 60%)"
//	v := validation.New(color, "theme_color")
//	v.HSL()  // Passes (orange-red)
//
//	// Red / 빨강
//	v = validation.New("hsl(0, 100%, 50%)", "color")
//	v.HSL()  // Passes
//
//	// Green / 초록
//	v = validation.New("hsl(120, 100%, 50%)", "color")
//	v.HSL()  // Passes
//
//	// Blue / 파랑
//	v = validation.New("hsl(240, 100%, 50%)", "color")
//	v.HSL()  // Passes
//
//	// Gray (no saturation) / 회색 (채도 없음)
//	v = validation.New("hsl(0, 0%, 50%)", "color")
//	v.HSL()  // Passes
//
//	// White / 흰색
//	v = validation.New("hsl(0, 0%, 100%)", "color")
//	v.HSL()  // Passes
//
//	// Black / 검정
//	v = validation.New("hsl(0, 0%, 0%)", "color")
//	v.HSL()  // Passes
//
//	// Invalid values / 무효한 값
//	v = validation.New("hsl(361, 100%, 50%)", "color")
//	v.HSL()  // Fails (hue > 360)
//
//	v = validation.New("hsl(0, 101%, 50%)", "color")
//	v.HSL()  // Fails (saturation > 100)
//
//	v = validation.New("hsl(0, 100, 50)", "color")
//	v.HSL()  // Fails (missing % symbols)
//
// Validation rules / 검증 규칙:
//   - Must be a string / 문자열이어야 함
//   - Format: hsl(h, s%, l%) / 형식: hsl(h, s%, l%)
//   - Hue must be 0-360 degrees / 색조는 0-360도 범위
//   - Saturation and Lightness must be 0-100% / 채도와 명도는 0-100% 범위
func (v *Validator) HSL() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("hsl", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Extract values from hsl(h, s%, l%) format
	hslRegex := regexp.MustCompile(`^hsl\s*\(\s*(\d+)\s*,\s*(\d+)%\s*,\s*(\d+)%\s*\)$`)
	matches := hslRegex.FindStringSubmatch(str)

	if matches == nil {
		v.addError("hsl", fmt.Sprintf("%s must be a valid HSL format (hsl(h, s%%, l%%)) / %s은(는) 유효한 HSL 형식이어야 합니다 (hsl(h, s%%, l%%))", v.fieldName, v.fieldName))
		return v
	}

	// Validate hue (0-360)
	hue, _ := strconv.Atoi(matches[1])
	if hue < 0 || hue > 360 {
		v.addError("hsl", fmt.Sprintf("%s HSL hue must be between 0 and 360 / %s HSL 색조는 0과 360 사이여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Validate saturation (0-100)
	saturation, _ := strconv.Atoi(matches[2])
	if saturation < 0 || saturation > 100 {
		v.addError("hsl", fmt.Sprintf("%s HSL saturation must be between 0 and 100 / %s HSL 채도는 0과 100 사이여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Validate lightness (0-100)
	lightness, _ := strconv.Atoi(matches[3])
	if lightness < 0 || lightness > 100 {
		v.addError("hsl", fmt.Sprintf("%s HSL lightness must be between 0 and 100 / %s HSL 명도는 0과 100 사이여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	return v
}
