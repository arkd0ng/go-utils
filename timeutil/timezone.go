package timeutil

import (
	"fmt"
	"sync"
	"time"
)

// Timezone cache for performance
// 성능을 위한 타임존 캐시
var (
	timezoneCache   = make(map[string]*time.Location)
	timezoneCacheMu sync.RWMutex
)

// SetDefaultTimezone sets the default timezone for all timeutil functions.
// SetDefaultTimezone은 모든 timeutil 함수의 기본 타임존을 설정합니다.
//
// Default is "Asia/Seoul" (KST, GMT+9). / 기본값은 "Asia/Seoul" (KST, GMT+9)입니다.
//
// Example
// 예제:
//
//	timeutil.SetDefaultTimezone("America/New_York")
func SetDefaultTimezone(tz string) error {
	loc, err := loadTimezone(tz)
	if err != nil {
		return fmt.Errorf("failed to set default timezone: %w", err)
	}
	defaultLocation = loc
	return nil
}

// GetDefaultTimezone returns the current default timezone name.
// GetDefaultTimezone은 현재 기본 타임존 이름을 반환합니다.
//
// Example
// 예제:
//
//	tz := timeutil.GetDefaultTimezone() // "Asia/Seoul"
func GetDefaultTimezone() string {
	return defaultLocation.String()
}

// ResetDefaultTimezone resets the default timezone to KST (Asia/Seoul).
// ResetDefaultTimezone은 기본 타임존을 KST (Asia/Seoul)로 재설정합니다.
func ResetDefaultTimezone() {
	defaultLocation = KST
}

// ConvertTimezone converts a time to a different timezone.
// ConvertTimezone은 시간을 다른 타임존으로 변환합니다.
//
// Example
// 예제:
//
//	now := time.Now()
//	seoulTime, _ := timeutil.ConvertTimezone(now, "Asia/Seoul")
//	nyTime, _ := timeutil.ConvertTimezone(now, "America/New_York")
func ConvertTimezone(t time.Time, tz string) (time.Time, error) {
	loc, err := loadTimezone(tz)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to convert timezone: %w", err)
	}
	return t.In(loc), nil
}

// ToKST converts a time to KST (Asia/Seoul, GMT+9).
// ToKST는 시간을 KST (Asia/Seoul, GMT+9)로 변환합니다.
//
// This is a convenience function for ConvertTimezone(t, "Asia/Seoul"). / 이것은 ConvertTimezone(t, "Asia/Seoul")의 편의 함수입니다.
//
// Example
// 예제:
//
//	kstTime := timeutil.ToKST(time.Now())
func ToKST(t time.Time) time.Time {
	return t.In(KST)
}

// ToUTC converts a time to UTC.
// ToUTC는 시간을 UTC로 변환합니다.
//
// Example
// 예제:
//
//	utcTime := timeutil.ToUTC(time.Now())
func ToUTC(t time.Time) time.Time {
	return t.UTC()
}

// GetTimezoneOffset returns the timezone offset in hours from UTC.
// GetTimezoneOffset는 UTC로부터의 타임존 오프셋을 시간 단위로 반환합니다.
//
// Example
// 예제:
//
//	offset, _ := timeutil.GetTimezoneOffset("Asia/Seoul") // +9
func GetTimezoneOffset(tz string) (int, error) {
	loc, err := loadTimezone(tz)
	if err != nil {
		return 0, fmt.Errorf("failed to get timezone offset: %w", err)
	}

	// Get offset at current time
	// 현재 시간의 오프셋 가져오기
	_, offset := time.Now().In(loc).Zone()
	return offset / SecondsPerHour, nil
}

// ListTimezones returns a list of common timezone names.
// ListTimezones는 일반적인 타임존 이름 목록을 반환합니다.
//
// Note: This returns a curated list of commonly used timezones.
// 참고: 이것은 일반적으로 사용되는 타임존의 선별된 목록을 반환합니다.
func ListTimezones() []string {
	return []string{
		// Asia
		// 아시아
		"Asia/Seoul",        // KST (GMT+9)
		"Asia/Tokyo",        // JST (GMT+9)
		"Asia/Shanghai",     // CST (GMT+8)
		"Asia/Hong_Kong",    // HKT (GMT+8)
		"Asia/Singapore",    // SGT (GMT+8)
		"Asia/Bangkok",      // ICT (GMT+7)
		"Asia/Dubai",        // GST (GMT+4)
		"Asia/Kolkata",      // IST (GMT+5:30)
		"Asia/Jakarta",      // WIB (GMT+7)
		"Asia/Manila",       // PHT (GMT+8)
		"Asia/Taipei",       // CST (GMT+8)
		"Asia/Ho_Chi_Minh",  // ICT (GMT+7)
		"Asia/Kuala_Lumpur", // MYT (GMT+8)

		// Europe
		// 유럽
		"Europe/London",    // GMT/BST (GMT+0/+1)
		"Europe/Paris",     // CET/CEST (GMT+1/+2)
		"Europe/Berlin",    // CET/CEST (GMT+1/+2)
		"Europe/Rome",      // CET/CEST (GMT+1/+2)
		"Europe/Madrid",    // CET/CEST (GMT+1/+2)
		"Europe/Amsterdam", // CET/CEST (GMT+1/+2)
		"Europe/Brussels",  // CET/CEST (GMT+1/+2)
		"Europe/Vienna",    // CET/CEST (GMT+1/+2)
		"Europe/Zurich",    // CET/CEST (GMT+1/+2)
		"Europe/Stockholm", // CET/CEST (GMT+1/+2)
		"Europe/Moscow",    // MSK (GMT+3)

		// Americas
		// 아메리카
		"America/New_York",      // EST/EDT (GMT-5/-4)
		"America/Chicago",       // CST/CDT (GMT-6/-5)
		"America/Denver",        // MST/MDT (GMT-7/-6)
		"America/Los_Angeles",   // PST/PDT (GMT-8/-7)
		"America/Toronto",       // EST/EDT (GMT-5/-4)
		"America/Vancouver",     // PST/PDT (GMT-8/-7)
		"America/Mexico_City",   // CST/CDT (GMT-6/-5)
		"America/Sao_Paulo",     // BRT/BRST (GMT-3/-2)
		"America/Buenos_Aires",  // ART (GMT-3)
		"America/Santiago",      // CLT/CLST (GMT-4/-3)
		"America/Bogota",        // COT (GMT-5)
		"America/Lima",          // PET (GMT-5)
		"America/Caracas",       // VET (GMT-4)
		"America/Panama",        // EST (GMT-5)
		"America/Havana",        // CST/CDT (GMT-5/-4)
		"America/Port-au-Prince", // EST/EDT (GMT-5/-4)

		// Pacific
		// 태평양
		"Pacific/Auckland",  // NZST/NZDT (GMT+12/+13)
		"Pacific/Fiji",      // FJT/FJST (GMT+12/+13)
		"Pacific/Honolulu",  // HST (GMT-10)
		"Pacific/Guam",      // ChST (GMT+10)
		"Pacific/Pago_Pago", // SST (GMT-11)
		"Pacific/Tahiti",    // TAHT (GMT-10)

		// Australia
		// 호주
		"Australia/Sydney",    // AEST/AEDT (GMT+10/+11)
		"Australia/Melbourne", // AEST/AEDT (GMT+10/+11)
		"Australia/Brisbane",  // AEST (GMT+10)
		"Australia/Perth",     // AWST (GMT+8)
		"Australia/Adelaide",  // ACST/ACDT (GMT+9:30/+10:30)

		// Africa
		// 아프리카
		"Africa/Cairo",        // EET/EEST (GMT+2/+3)
		"Africa/Johannesburg", // SAST (GMT+2)
		"Africa/Lagos",        // WAT (GMT+1)
		"Africa/Nairobi",      // EAT (GMT+3)
		"Africa/Casablanca",   // WET/WEST (GMT+0/+1)

		// Middle East
		// 중동
		"Asia/Jerusalem",  // IST/IDT (GMT+2/+3)
		"Asia/Riyadh",     // AST (GMT+3)
		"Asia/Tehran",     // IRST/IRDT (GMT+3:30/+4:30)
		"Asia/Baghdad",    // AST (GMT+3)
		"Asia/Kuwait",     // AST (GMT+3)
		"Asia/Doha",       // AST (GMT+3)
		"Asia/Muscat",     // GST (GMT+4)
		"Asia/Karachi",    // PKT (GMT+5)
		"Asia/Dhaka",      // BST (GMT+6)
		"Asia/Yangon",     // MMT (GMT+6:30)
		"Asia/Kathmandu",  // NPT (GMT+5:45)

		// UTC
		"UTC",
	}
}

// IsValidTimezone checks if a timezone name is valid.
// IsValidTimezone은 타임존 이름이 유효한지 확인합니다.
//
// Example
// 예제:
//
//	if timeutil.IsValidTimezone("Asia/Seoul") {
//	    // Valid timezone
//	}
func IsValidTimezone(tz string) bool {
	_, err := time.LoadLocation(tz)
	return err == nil
}

// GetLocalTimezone returns the local system timezone name.
// GetLocalTimezone은 로컬 시스템 타임존 이름을 반환합니다.
//
// Example
// 예제:
//
//	local := timeutil.GetLocalTimezone()
func GetLocalTimezone() string {
	return time.Local.String()
}

// NowInTimezone returns the current time in the specified timezone.
// NowInTimezone은 지정된 타임존의 현재 시간을 반환합니다.
//
// Example
// 예제:
//
//	seoulNow, _ := timeutil.NowInTimezone("Asia/Seoul")
func NowInTimezone(tz string) (time.Time, error) {
	loc, err := loadTimezone(tz)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to get current time: %w", err)
	}
	return time.Now().In(loc), nil
}

// NowKST returns the current time in KST (Asia/Seoul, GMT+9).
// NowKST는 KST (Asia/Seoul, GMT+9)의 현재 시간을 반환합니다.
//
// This is the default timezone for all timeutil functions.
// 이것은 모든 timeutil 함수의 기본 타임존입니다.
//
// Example
// 예제:
//
//	now := timeutil.NowKST()
func NowKST() time.Time {
	return time.Now().In(KST)
}

// loadTimezone loads a timezone location with caching.
// loadTimezone은 캐싱과 함께 타임존 위치를 로드합니다.
func loadTimezone(tz string) (*time.Location, error) {
	// Check cache first
	// 먼저 캐시 확인
	timezoneCacheMu.RLock()
	if loc, ok := timezoneCache[tz]; ok {
		timezoneCacheMu.RUnlock()
		return loc, nil
	}
	timezoneCacheMu.RUnlock()

	// Load timezone
	// 타임존 로드
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return nil, fmt.Errorf("invalid timezone '%s': %w", tz, err)
	}

	// Cache it
	// 캐시에 저장
	timezoneCacheMu.Lock()
	timezoneCache[tz] = loc
	timezoneCacheMu.Unlock()

	return loc, nil
}
