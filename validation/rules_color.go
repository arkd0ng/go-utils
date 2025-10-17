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
// HexColor는 값이 유효한 16진수 색상 코드인지 검증합니다.
//
// Supports both 3-digit (#RGB) and 6-digit (#RRGGBB) hex colors, with or without # prefix.
// 3자리(#RGB) 및 6자리(#RRGGBB) 16진수 색상을 모두 지원하며, # 접두사는 선택 사항입니다.
//
// Example / 예시:
//
//	color := "#FF5733"
//	v := validation.New(color, "brand_color")
//	v.HexColor()
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
// RGB는 값이 유효한 RGB 색상 형식인지 검증합니다.
//
// Format: rgb(r, g, b) where r, g, b are integers 0-255.
// 형식: rgb(r, g, b) 여기서 r, g, b는 0-255 범위의 정수입니다.
//
// Example / 예시:
//
//	color := "rgb(255, 87, 51)"
//	v := validation.New(color, "background_color")
//	v.RGB()
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
// RGBA는 값이 유효한 RGBA 색상 형식인지 검증합니다.
//
// Format: rgba(r, g, b, a) where r, g, b are integers 0-255 and a is 0-1.
// 형식: rgba(r, g, b, a) 여기서 r, g, b는 0-255 정수이고 a는 0-1 범위입니다.
//
// Example / 예시:
//
//	color := "rgba(255, 87, 51, 0.8)"
//	v := validation.New(color, "overlay_color")
//	v.RGBA()
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
// HSL은 값이 유효한 HSL 색상 형식인지 검증합니다.
//
// Format: hsl(h, s%, l%) where h is 0-360, s and l are 0-100%.
// 형식: hsl(h, s%, l%) 여기서 h는 0-360, s와 l은 0-100% 범위입니다.
//
// Example / 예시:
//
//	color := "hsl(9, 100%, 60%)"
//	v := validation.New(color, "theme_color")
//	v.HSL()
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
