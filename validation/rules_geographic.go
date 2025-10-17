package validation

import (
	"fmt"
	"strconv"
	"strings"
)

// Latitude validates latitude coordinates (-90 to 90).
// Accepts numeric types and string representations with automatic conversion.
//
// Latitude는 위도 좌표(-90 ~ 90)를 검증합니다.
// 자동 변환을 통해 숫자 타입과 문자열 표현을 허용합니다.
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
//   - Accepts float64, float32, int, int64, string
//     float64, float32, int, int64, string 허용
//   - Valid range: -90.0 to 90.0 degrees (inclusive)
//     유효 범위: -90.0 ~ 90.0도 (포함)
//   - String values parsed as float64
//     문자열 값은 float64로 파싱
//   - Fails if value not numeric or parseable
//     값이 숫자가 아니거나 파싱 불가능하면 실패
//
// Use Cases / 사용 사례:
//   - GPS coordinate validation / GPS 좌표 검증
//   - Location data validation / 위치 데이터 검증
//   - Map application input / 지도 애플리케이션 입력
//   - Geolocation services / 지오로케이션 서비스
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Type switch + range check
//     타입 스위치 + 범위 확인
//
// Valid range: -90.0 to 90.0 degrees
// 유효 범위: -90.0 ~ 90.0도
//
// Example / 예시:
//
//	// Seoul latitude / 서울 위도
//	v := validation.New(37.5665, "latitude")
//	v.Latitude()  // Passes
//
//	// North Pole / 북극
//	v = validation.New(90.0, "lat")
//	v.Latitude()  // Passes
//
//	// South Pole / 남극
//	v = validation.New(-90.0, "lat")
//	v.Latitude()  // Passes
//
//	// String value / 문자열 값
//	v = validation.New("37.5665", "latitude")
//	v.Latitude()  // Passes
//
//	// Invalid - out of range / 무효 - 범위 밖
//	v = validation.New(95.0, "latitude")
//	v.Latitude()  // Fails (> 90)
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
// Accepts numeric types and string representations with automatic conversion.
//
// Longitude는 경도 좌표(-180 ~ 180)를 검증합니다.
// 자동 변환을 통해 숫자 타입과 문자열 표현을 허용합니다.
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
//   - Accepts float64, float32, int, int64, string
//     float64, float32, int, int64, string 허용
//   - Valid range: -180.0 to 180.0 degrees (inclusive)
//     유효 범위: -180.0 ~ 180.0도 (포함)
//   - String values parsed as float64
//     문자열 값은 float64로 파싱
//   - Fails if value not numeric or parseable
//     값이 숫자가 아니거나 파싱 불가능하면 실패
//
// Use Cases / 사용 사례:
//   - GPS coordinate validation / GPS 좌표 검증
//   - Location data validation / 위치 데이터 검증
//   - Map application input / 지도 애플리케이션 입력
//   - Geolocation services / 지오로케이션 서비스
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(1)
//     시간 복잡도: O(1)
//   - Type switch + range check
//     타입 스위치 + 범위 확인
//
// Valid range: -180.0 to 180.0 degrees
// 유효 범위: -180.0 ~ 180.0도
//
// Example / 예시:
//
//	// Seoul longitude / 서울 경도
//	v := validation.New(126.9780, "longitude")
//	v.Longitude()  // Passes
//
//	// International Date Line (East) / 국제 날짜 변경선 (동)
//	v = validation.New(180.0, "lon")
//	v.Longitude()  // Passes
//
//	// International Date Line (West) / 국제 날짜 변경선 (서)
//	v = validation.New(-180.0, "lon")
//	v.Longitude()  // Passes
//
//	// String value / 문자열 값
//	v = validation.New("126.9780", "longitude")
//	v.Longitude()  // Passes
//
//	// Invalid - out of range / 무효 - 범위 밖
//	v = validation.New(185.0, "longitude")
//	v.Longitude()  // Fails (> 180)
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
// Parses and validates both latitude and longitude from comma-separated string.
//
// Coordinate는 "lat,lon" 형식의 좌표 문자열을 검증합니다.
// 쉼표로 구분된 문자열에서 위도와 경도를 모두 파싱하고 검증합니다.
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
//   - Format: "lat,lon" or "lat, lon" (spaces allowed)
//     형식: "lat,lon" 또는 "lat, lon" (공백 허용)
//   - Trims whitespace around values
//     값 주변 공백 제거
//   - Validates latitude range: -90 to 90
//     위도 범위 검증: -90 ~ 90
//   - Validates longitude range: -180 to 180
//     경도 범위 검증: -180 ~ 180
//   - Fails if not string or invalid format
//     문자열이 아니거나 형식이 무효하면 실패
//
// Use Cases / 사용 사례:
//   - GPS data validation / GPS 데이터 검증
//   - Location string parsing / 위치 문자열 파싱
//   - Map URL parameters / 지도 URL 매개변수
//   - Geolocation API input / 지오로케이션 API 입력
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n), n = string length
//     시간 복잡도: O(n), n = 문자열 길이
//   - String split + 2 float parses + 2 range checks
//     문자열 분할 + 2번의 float 파싱 + 2번의 범위 확인
//
// Accepts formats: "lat,lon" or "lat, lon" (with optional space)
// 허용 형식: "lat,lon" 또는 "lat, lon" (선택적 공백 포함)
//
// Example / 예시:
//
//	// Seoul coordinates / 서울 좌표
//	v := validation.New("37.5665,126.9780", "location")
//	v.Coordinate()  // Passes
//
//	// With spaces / 공백 포함
//	v = validation.New("37.5665, 126.9780", "location")
//	v.Coordinate()  // Passes
//
//	// Negative coordinates / 음수 좌표
//	v = validation.New("-33.8688,151.2093", "location")
//	v.Coordinate()  // Passes (Sydney)
//
//	// Boundaries / 경계값
//	v = validation.New("90,180", "location")
//	v.Coordinate()  // Passes
//
//	// Invalid - missing longitude / 무효 - 경도 누락
//	v = validation.New("37.5665", "location")
//	v.Coordinate()  // Fails
//
//	// Invalid - out of range / 무효 - 범위 밖
//	v = validation.New("95,126", "location")
//	v.Coordinate()  // Fails (lat > 90)
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
