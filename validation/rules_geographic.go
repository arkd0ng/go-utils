package validation

import (
	"fmt"
	"strconv"
	"strings"
)

// Latitude validates latitude coordinates (-90 to 90).
// Latitude는 위도 좌표(-90 ~ 90)를 검증합니다.
//
// Valid range: -90.0 to 90.0 degrees
// 유효 범위: -90.0 ~ 90.0도
//
// Example / 예시:
//   v := validation.New(37.5665, "latitude")
//   v.Latitude()
func (v *Validator) Latitude() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	var lat float64
	switch val := v.value.(type) {
	case float64:
		lat = val
	case float32:
		lat = float64(val)
	case int:
		lat = float64(val)
	case int64:
		lat = float64(val)
	case string:
		parsed, err := strconv.ParseFloat(val, 64)
		if err != nil {
			v.addError("latitude", fmt.Sprintf("%s must be a valid number / %s은(는) 유효한 숫자여야 합니다", v.fieldName, v.fieldName))
			return v
		}
		lat = parsed
	default:
		v.addError("latitude", fmt.Sprintf("%s must be a number / %s은(는) 숫자여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if lat < -90.0 || lat > 90.0 {
		v.addError("latitude", fmt.Sprintf("%s must be between -90 and 90 / %s은(는) -90과 90 사이여야 합니다", v.fieldName, v.fieldName))
	}

	return v
}

// Longitude validates longitude coordinates (-180 to 180).
// Longitude는 경도 좌표(-180 ~ 180)를 검증합니다.
//
// Valid range: -180.0 to 180.0 degrees
// 유효 범위: -180.0 ~ 180.0도
//
// Example / 예시:
//   v := validation.New(126.9780, "longitude")
//   v.Longitude()
func (v *Validator) Longitude() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	var lon float64
	switch val := v.value.(type) {
	case float64:
		lon = val
	case float32:
		lon = float64(val)
	case int:
		lon = float64(val)
	case int64:
		lon = float64(val)
	case string:
		parsed, err := strconv.ParseFloat(val, 64)
		if err != nil {
			v.addError("longitude", fmt.Sprintf("%s must be a valid number / %s은(는) 유효한 숫자여야 합니다", v.fieldName, v.fieldName))
			return v
		}
		lon = parsed
	default:
		v.addError("longitude", fmt.Sprintf("%s must be a number / %s은(는) 숫자여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if lon < -180.0 || lon > 180.0 {
		v.addError("longitude", fmt.Sprintf("%s must be between -180 and 180 / %s은(는) -180과 180 사이여야 합니다", v.fieldName, v.fieldName))
	}

	return v
}

// Coordinate validates coordinate string in "lat,lon" format.
// Coordinate는 "lat,lon" 형식의 좌표 문자열을 검증합니다.
//
// Accepts formats: "lat,lon" or "lat, lon" (with optional space)
// 허용 형식: "lat,lon" 또는 "lat, lon" (선택적 공백 포함)
//
// Example / 예시:
//   v := validation.New("37.5665,126.9780", "location")
//   v.Coordinate()
func (v *Validator) Coordinate() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("coordinate", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Split by comma
	parts := strings.Split(str, ",")
	if len(parts) != 2 {
		v.addError("coordinate", fmt.Sprintf("%s must be in 'lat,lon' format / %s은(는) 'lat,lon' 형식이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Parse latitude
	latStr := strings.TrimSpace(parts[0])
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		v.addError("coordinate", fmt.Sprintf("%s latitude must be a valid number / %s 위도는 유효한 숫자여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Parse longitude
	lonStr := strings.TrimSpace(parts[1])
	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		v.addError("coordinate", fmt.Sprintf("%s longitude must be a valid number / %s 경도는 유효한 숫자여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Validate ranges
	if lat < -90.0 || lat > 90.0 {
		v.addError("coordinate", fmt.Sprintf("%s latitude must be between -90 and 90 / %s 위도는 -90과 90 사이여야 합니다", v.fieldName, v.fieldName))
	}

	if lon < -180.0 || lon > 180.0 {
		v.addError("coordinate", fmt.Sprintf("%s longitude must be between -180 and 180 / %s 경도는 -180과 180 사이여야 합니다", v.fieldName, v.fieldName))
	}

	return v
}
