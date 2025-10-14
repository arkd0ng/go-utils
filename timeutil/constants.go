package timeutil

import "time"

// Default timezone / 기본 타임존
// All functions use KST (Asia/Seoul, GMT+9) as default timezone unless specified.
// 모든 함수는 별도 지정이 없으면 KST (Asia/Seoul, GMT+9)를 기본 타임존으로 사용합니다.
const (
	DefaultTimezone = "Asia/Seoul"
	DefaultLocation = "Asia/Seoul" // KST, GMT+9
)

var (
	// KST is the default timezone location (Asia/Seoul, GMT+9)
	// KST는 기본 타임존 위치입니다 (Asia/Seoul, GMT+9)
	KST *time.Location

	// defaultLocation is the current default location, can be changed
	// defaultLocation은 현재 기본 위치이며 변경 가능합니다
	defaultLocation *time.Location
)

// init loads the default timezone (KST)
// init은 기본 타임존(KST)을 로드합니다
func init() {
	var err error
	KST, err = time.LoadLocation(DefaultTimezone)
	if err != nil {
		// Fallback to UTC if KST cannot be loaded
		// KST를 로드할 수 없으면 UTC로 폴백
		KST = time.UTC
	}
	defaultLocation = KST
}

// Time constants / 시간 상수
const (
	SecondsPerMinute = 60
	SecondsPerHour   = 3600
	SecondsPerDay    = 86400
	DaysPerWeek      = 7
	MonthsPerYear    = 12
	HoursPerDay      = 24
	MinutesPerHour   = 60
)

// Common format layouts / 일반 포맷 레이아웃
const (
	// ISO8601 format / ISO8601 포맷
	ISO8601Layout = "2006-01-02T15:04:05Z07:00"

	// RFC3339 format / RFC3339 포맷
	RFC3339Layout = "2006-01-02T15:04:05Z07:00"

	// Date only format / 날짜만 포맷
	DateLayout = "2006-01-02"

	// DateTime format / 날짜시간 포맷
	DateTimeLayout = "2006-01-02 15:04:05"

	// Time only format / 시간만 포맷
	TimeLayout = "15:04:05"
)

// Custom format tokens for user-friendly formatting / 사용자 친화적 포맷팅을 위한 커스텀 포맷 토큰
// These tokens are translated to Go's standard layout format.
// 이 토큰들은 Go의 표준 레이아웃 포맷으로 변환됩니다.
//
// Supported tokens / 지원되는 토큰:
//   YYYY - 4-digit year (2006)
//   YY   - 2-digit year (06)
//   MM   - 2-digit month (01-12)
//   M    - 1 or 2-digit month (1-12)
//   DD   - 2-digit day (01-31)
//   D    - 1 or 2-digit day (1-31)
//   HH   - 2-digit hour 24h format (00-23)
//   hh   - 2-digit hour 12h format (01-12)
//   mm   - 2-digit minute (00-59)
//   ss   - 2-digit second (00-59)
var customFormatTokens = map[string]string{
	"YYYY": "2006",
	"YY":   "06",
	"MM":   "01",
	"M":    "1",
	"DD":   "02",
	"D":    "2",
	"HH":   "15",
	"hh":   "03",
	"mm":   "04",
	"ss":   "05",
}
