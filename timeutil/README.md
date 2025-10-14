# timeutil - Time and Date Utilities / 시간 및 날짜 유틸리티

**v1.6.008** - Now with String Parameter Support! 🎉

Extreme simplicity time and date utility functions for Go - reduce 20 lines of time manipulation code to just 1 line.

극도로 간단한 Go용 시간 및 날짜 유틸리티 함수 - 20줄의 시간 조작 코드를 단 1줄로 줄입니다.

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.16-blue)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## ⭐ NEW in v1.6.008!

**String Parameter Support** - Parse any time format automatically!
- 🚀 **40+ formats** automatically detected (including Korean!)
- 🎯 **50+ String functions** - all major functions now accept strings
- 💾 **Database-friendly** - perfect for MySQL, PostgreSQL, SQLite timestamps
- 🌏 **한글 지원** - "2024년 10월 04일", "오전 9시", "오후 3시"

```go
// OLD WAY - 10+ lines of boilerplate
layout := "2006-01-02 15:04:05.000"
t1, err := time.ParseInLocation(layout, "2024-10-04 08:34:42.324", timeutil.KST)
// ... more parsing code ...

// NEW WAY - 1 line! 🎉
diff, err := timeutil.SubTimeString("2024-10-04 08:34:42.324", "2024-10-14 14:56:23.789")
```

## Overview / 개요

The `timeutil` package provides **150+ intuitive functions** for common time and date operations in Go. Stop writing repetitive time manipulation code and start using human-readable function names.

`timeutil` 패키지는 Go에서 일반적인 시간 및 날짜 작업을 위한 **150개 이상의 직관적인 함수**를 제공합니다. 반복적인 시간 조작 코드 작성을 멈추고 사람이 읽기 쉬운 함수 이름을 사용하세요.

### Design Philosophy: "20 lines → 1 line" / 설계 철학: "20줄 → 1줄"

**Before (Standard Go) / 이전 (표준 Go)**:
```go
// Calculate time difference in days / 일 단위로 시간 차이 계산
start := time.Date(2025, 1, 1, 9, 0, 0, 0, time.UTC)
end := time.Date(2025, 1, 3, 15, 30, 0, 0, time.UTC)

duration := end.Sub(start)
hours := duration.Hours()
days := hours / 24

if days > 0 {
    fmt.Printf("%d days %d hours", int(days), int(hours)%24)
} else if hours > 0 {
    fmt.Printf("%d hours %d minutes", int(hours), int(duration.Minutes())%60)
}
// 10+ lines
```

**After (This Package) / 이후 (이 패키지)**:
```go
diff := timeutil.SubTime(start, end)
fmt.Println(diff.String()) // "2 days 6 hours 30 minutes"
// 1-2 lines
```

## Installation / 설치

```bash
go get github.com/arkd0ng/go-utils/timeutil
```

## Quick Start / 빠른 시작

```go
package main

import (
    "fmt"
    "time"
    "github.com/arkd0ng/go-utils/timeutil"
)

func main() {
    // Time difference / 시간 차이
    start := time.Date(2025, 1, 1, 9, 0, 0, 0, time.UTC)
    end := time.Date(2025, 1, 3, 15, 30, 0, 0, time.UTC)
    diff := timeutil.SubTime(start, end)
    fmt.Println(diff.Days())    // 2
    fmt.Println(diff.Hours())   // 54.5
    fmt.Println(diff.String())  // "2 days 6 hours 30 minutes"

    // Timezone conversion / 타임존 변환
    seoulTime, _ := timeutil.ConvertTimezone(time.Now(), "Asia/Seoul")
    nyTime, _ := timeutil.ConvertTimezone(time.Now(), "America/New_York")

    // Date arithmetic / 날짜 연산
    tomorrow := timeutil.AddDays(time.Now(), 1)
    nextWeek := timeutil.AddWeeks(time.Now(), 1)
    startOfMonth := timeutil.StartOfMonth(time.Now())

    // Date formatting / 날짜 포맷팅
    iso := timeutil.FormatISO8601(time.Now())
    custom := timeutil.Format(time.Now(), "YYYY-MM-DD HH:mm:ss")

    // Time comparisons / 시간 비교
    if timeutil.IsToday(someDate) {
        fmt.Println("It's today!")
    }

    // Age calculations / 나이 계산
    birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
    age := timeutil.AgeInYears(birthDate)
    fmt.Printf("Age: %d years\n", age)

    // Business days / 영업일
    nextBusinessDay := timeutil.AddBusinessDays(time.Now(), 5)
    if timeutil.IsBusinessDay(time.Now()) {
        fmt.Println("It's a business day!")
    }

    // Relative time / 상대 시간
    past := time.Now().Add(-2 * time.Hour)
    fmt.Println(timeutil.RelativeTime(past)) // "2 hours ago"

    // ⭐ NEW! String Parameters - Parse any format automatically!
    // 문자열 매개변수 - 모든 포맷 자동 파싱!

    // Parse any time format
    t1, _ := timeutil.ParseAny("2024-10-04 08:34:42.324")  // MySQL
    t2, _ := timeutil.ParseAny("Oct 04, 2024")            // Month name
    t3, _ := timeutil.ParseAny("2024년 10월 04일")          // Korean!

    // Calculate difference from strings directly
    diff2, _ := timeutil.SubTimeString("2024-10-04", "2024-10-14")

    // Get age from string
    age2, _ := timeutil.AgeInYearsString("1990-01-15")

    // Format conversion
    formatted, _ := timeutil.FormatString("Oct 04, 2024", "2006-01-02")

    // Unix timestamp / Unix 타임스탬프
    unix := timeutil.Now()
    unixMilli := timeutil.NowMilli()
}
```

## Features / 기능

### 1. Time Difference / 시간 차이 (8 functions)

Calculate time differences easily with human-readable output.

사람이 읽기 쉬운 출력으로 시간 차이를 쉽게 계산합니다.

```go
// Time difference / 시간 차이
diff := timeutil.SubTime(start, end)
fmt.Println(diff.Days())       // 2
fmt.Println(diff.Hours())      // 54.5
fmt.Println(diff.Minutes())    // 3270
fmt.Println(diff.String())     // "2 days 6 hours 30 minutes"
fmt.Println(diff.Humanize())   // "2d 6h 30m"

// Direct difference functions / 직접 차이 함수
seconds := timeutil.DiffInSeconds(t1, t2)
minutes := timeutil.DiffInMinutes(t1, t2)
hours := timeutil.DiffInHours(t1, t2)
days := timeutil.DiffInDays(t1, t2)
weeks := timeutil.DiffInWeeks(t1, t2)
months := timeutil.DiffInMonths(t1, t2)
years := timeutil.DiffInYears(t1, t2)
```

### 2. Timezone Operations / 타임존 작업 (5 functions)

Handle timezones without pain.

고통 없이 타임존을 처리합니다.

```go
// Convert to different timezone / 다른 타임존으로 변환
seoulTime, _ := timeutil.ConvertTimezone(time.Now(), "Asia/Seoul")
nyTime, _ := timeutil.ConvertTimezone(time.Now(), "America/New_York")

// Get timezone offset / 타임존 오프셋 가져오기
offset, _ := timeutil.GetTimezoneOffset("Asia/Seoul") // +9

// List all timezones / 모든 타임존 나열
zones := timeutil.ListTimezones()

// Validate timezone / 타임존 검증
if timeutil.IsValidTimezone("Asia/Seoul") {
    // Valid timezone
}

// Get local timezone / 로컬 타임존 가져오기
local := timeutil.GetLocalTimezone()
```

### 3. Date Arithmetic / 날짜 연산 (16 functions)

Add/subtract time units easily.

시간 단위를 쉽게 더하기/빼기합니다.

```go
// Add/subtract time units / 시간 단위 더하기/빼기
tomorrow := timeutil.AddDays(time.Now(), 1)
yesterday := timeutil.AddDays(time.Now(), -1)
nextWeek := timeutil.AddWeeks(time.Now(), 1)
nextMonth := timeutil.AddMonths(time.Now(), 1)
nextYear := timeutil.AddYears(time.Now(), 1)

// Also available: AddSeconds, AddMinutes, AddHours
// 또한 사용 가능: AddSeconds, AddMinutes, AddHours

// Start/End of period / 기간의 시작/끝
startOfDay := timeutil.StartOfDay(time.Now())
endOfDay := timeutil.EndOfDay(time.Now())
startOfWeek := timeutil.StartOfWeek(time.Now())
endOfWeek := timeutil.EndOfWeek(time.Now())
startOfMonth := timeutil.StartOfMonth(time.Now())
endOfMonth := timeutil.EndOfMonth(time.Now())
startOfYear := timeutil.StartOfYear(time.Now())
endOfYear := timeutil.EndOfYear(time.Now())
startOfQuarter := timeutil.StartOfQuarter(time.Now())
```

### 4. Date Formatting / 날짜 포맷팅 (8 functions)

Format dates with intuitive format strings (no more "2006-01-02").

직관적인 포맷 문자열로 날짜를 포맷합니다 ("2006-01-02"는 이제 그만).

```go
// Common formats / 일반 포맷
iso := timeutil.FormatISO8601(time.Now())      // "2025-10-14T15:04:05+09:00"
rfc := timeutil.FormatRFC3339(time.Now())      // "2025-10-14T15:04:05+09:00"
date := timeutil.FormatDate(time.Now())        // "2025-10-14"
datetime := timeutil.FormatDateTime(time.Now()) // "2025-10-14 15:04:05"
timeStr := timeutil.FormatTime(time.Now())     // "15:04:05"

// Custom format (intuitive tokens!) / 커스텀 포맷 (직관적인 토큰!)
custom := timeutil.Format(time.Now(), "YYYY-MM-DD HH:mm:ss")
custom2 := timeutil.Format(time.Now(), "YYYY년 MM월 DD일")

// Format with timezone / 타임존과 함께 포맷
formatted, _ := timeutil.FormatWithTimezone(time.Now(), "Asia/Seoul")
```

**Custom Format Tokens / 커스텀 포맷 토큰**:
- `YYYY` - 4-digit year / 4자리 년
- `YY` - 2-digit year / 2자리 년
- `MM` - 2-digit month / 2자리 월
- `M` - 1 or 2-digit month / 1 또는 2자리 월
- `DD` - 2-digit day / 2자리 일
- `D` - 1 or 2-digit day / 1 또는 2자리 일
- `HH` - 2-digit hour (24h) / 2자리 시 (24시간)
- `hh` - 2-digit hour (12h) / 2자리 시 (12시간)
- `mm` - 2-digit minute / 2자리 분
- `ss` - 2-digit second / 2자리 초

### 5. Time Parsing / 시간 파싱 (6 functions)

Parse time strings with automatic format detection.

자동 포맷 감지로 시간 문자열을 파싱합니다.

```go
// Parse common formats / 일반 포맷 파싱
t1, _ := timeutil.ParseISO8601("2025-10-14T15:04:05+09:00")
t2, _ := timeutil.ParseRFC3339("2025-10-14T15:04:05+09:00")
t3, _ := timeutil.ParseDate("2025-10-14")
t4, _ := timeutil.ParseDateTime("2025-10-14 15:04:05")

// Smart parse (auto-detect format) / 스마트 파싱 (포맷 자동 감지)
t5, _ := timeutil.Parse("2025-10-14")
t6, _ := timeutil.Parse("2025-10-14 15:04:05")

// Parse with timezone / 타임존과 함께 파싱
t7, _ := timeutil.ParseWithTimezone("2025-10-14 15:04:05", "Asia/Seoul")
```

### 6. Business Days / 영업일 (6 functions)

Calculate business days (excluding weekends and holidays).

영업일을 계산합니다 (주말 및 공휴일 제외).

```go
// Add/subtract business days / 영업일 더하기/빼기
nextBusinessDay := timeutil.AddBusinessDays(time.Now(), 5)
prevBusinessDay := timeutil.AddBusinessDays(time.Now(), -3)

// Check if business day / 영업일인지 확인
if timeutil.IsBusinessDay(time.Now()) {
    // It's a weekday
}

// Count business days between two dates / 두 날짜 사이의 영업일 수
days := timeutil.CountBusinessDays(start, end)

// Next/Previous business day / 다음/이전 영업일
next := timeutil.NextBusinessDay(time.Now())
prev := timeutil.PreviousBusinessDay(time.Now())

// Check if holiday / 공휴일 확인
if timeutil.IsHoliday(someDate) {
    // It's a holiday
}

// Set custom holidays / 커스텀 공휴일 설정
holidays := []time.Time{
    time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),  // New Year
    time.Date(2025, 12, 25, 0, 0, 0, 0, time.UTC), // Christmas
}
timeutil.SetHolidays(holidays)
```

### 7. Time Comparisons / 시간 비교 (18 functions)

Intuitive time comparison functions.

직관적인 시간 비교 함수입니다.

```go
// Basic comparisons / 기본 비교
if timeutil.IsBefore(t1, t2) { /* ... */ }
if timeutil.IsAfter(t1, t2) { /* ... */ }
if timeutil.IsBetween(t, start, end) { /* ... */ }

// Special checks / 특별 체크
if timeutil.IsToday(t) { /* ... */ }
if timeutil.IsYesterday(t) { /* ... */ }
if timeutil.IsTomorrow(t) { /* ... */ }
if timeutil.IsThisWeek(t) { /* ... */ }
if timeutil.IsThisMonth(t) { /* ... */ }
if timeutil.IsThisYear(t) { /* ... */ }

// Weekend/Weekday / 주말/평일
if timeutil.IsWeekend(t) { /* ... */ }
if timeutil.IsWeekday(t) { /* ... */ }

// Same day/week/month/year / 같은 날/주/월/년
if timeutil.IsSameDay(t1, t2) { /* ... */ }
if timeutil.IsSameWeek(t1, t2) { /* ... */ }
if timeutil.IsSameMonth(t1, t2) { /* ... */ }
if timeutil.IsSameYear(t1, t2) { /* ... */ }

// Other checks / 기타 체크
if timeutil.IsLeapYear(t) { /* ... */ }
if timeutil.IsPast(t) { /* ... */ }
if timeutil.IsFuture(t) { /* ... */ }
```

### 8. Age Calculations / 나이 계산 (4 functions)

Calculate age easily.

나이를 쉽게 계산합니다.

```go
birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)

// Age in different units / 다양한 단위로 나이
years := timeutil.AgeInYears(birthDate)       // 35
months := timeutil.AgeInMonths(birthDate)     // 420
days := timeutil.AgeInDays(birthDate)         // 12800

// Detailed age / 상세 나이
age := timeutil.Age(birthDate)
fmt.Println(age.Years)   // 35
fmt.Println(age.Months)  // 5
fmt.Println(age.Days)    // 14
fmt.Println(age.String()) // "35 years 5 months 14 days"
```

### 9. Relative Time / 상대 시간 (3 functions)

Human-readable relative time strings.

사람이 읽기 쉬운 상대 시간 문자열입니다.

```go
// Relative time strings / 상대 시간 문자열
past := time.Now().Add(-2 * time.Hour)
fmt.Println(timeutil.RelativeTime(past))     // "2 hours ago"

future := time.Now().Add(3 * time.Hour)
fmt.Println(timeutil.RelativeTime(future))   // "in 3 hours"

// Short format / 짧은 포맷
fmt.Println(timeutil.RelativeTimeShort(past)) // "2h ago"

// Time ago / 시간 전
fmt.Println(timeutil.TimeAgo(past)) // "2 hours ago"
```

### 10. Unix Timestamp / Unix 타임스탬프 (8 functions)

Unix timestamp operations made easy.

Unix 타임스탬프 작업을 쉽게 만듭니다.

```go
// Current timestamp / 현재 타임스탬프
unix := timeutil.Now()                    // Unix timestamp in seconds
unixMilli := timeutil.NowMilli()          // Unix timestamp in milliseconds
unixMicro := timeutil.NowMicro()          // Unix timestamp in microseconds
unixNano := timeutil.NowNano()            // Unix timestamp in nanoseconds

// Convert from timestamp / 타임스탬프에서 변환
t1 := timeutil.FromUnix(1634198400)
t2 := timeutil.FromUnixMilli(1634198400000)

// Convert to timestamp / 타임스탬프로 변환
unix := timeutil.ToUnix(time.Now())
unixMilli := timeutil.ToUnixMilli(time.Now())
```

## Core Types / 핵심 타입

### TimeDiff Type / TimeDiff 타입

```go
type TimeDiff struct {
    Duration time.Duration
}

// Methods / 메서드
func (td *TimeDiff) Seconds() float64
func (td *TimeDiff) Minutes() float64
func (td *TimeDiff) Hours() float64
func (td *TimeDiff) Days() float64
func (td *TimeDiff) Weeks() float64
func (td *TimeDiff) String() string        // "2 days 3 hours 15 minutes"
func (td *TimeDiff) Humanize() string      // "2d 3h 15m"
func (td *TimeDiff) Abs() *TimeDiff        // Absolute value / 절대값
```

### Age Type / Age 타입

```go
type Age struct {
    Years  int
    Months int
    Days   int
}

// Methods / 메서드
func (a *Age) String() string    // "35 years 5 months 14 days"
func (a *Age) Humanize() string  // "35y 5m 14d"
```

## Package Statistics / 패키지 통계

- **Total Functions / 총 함수 수**: ~80+ functions / 약 80개 이상의 함수
- **Categories / 카테고리**: 10 categories / 10개 카테고리
- **Test Coverage / 테스트 커버리지**: ≥90%
- **Dependencies / 의존성**: Zero external dependencies / 외부 의존성 제로 (standard library only)
- **Thread-Safe / 스레드 안전**: All functions are thread-safe / 모든 함수가 스레드 안전
- **Code Lines / 코드 라인**: ~2000+ lines of code / 약 2000줄 이상의 코드
- **Test Lines / 테스트 라인**: ~3000+ lines of tests / 약 3000줄 이상의 테스트

## Testing / 테스트

```bash
# Run all tests / 모든 테스트 실행
go test ./timeutil -v

# Run with coverage / 커버리지와 함께 실행
go test ./timeutil -cover
go test ./timeutil -coverprofile=coverage.out
go tool cover -html=coverage.out

# Run benchmarks / 벤치마크 실행
go test ./timeutil -bench=.

# Run specific test / 특정 테스트 실행
go test ./timeutil -v -run TestSubTime
```

## Examples / 예제

See [examples/timeutil/main.go](../examples/timeutil/main.go) for comprehensive examples of all features.

모든 기능의 포괄적인 예제는 [examples/timeutil/main.go](../examples/timeutil/main.go)를 참조하세요.

```bash
# Run example / 예제 실행
go run examples/timeutil/main.go
```

## Documentation / 문서

- **[User Manual / 사용자 매뉴얼](../docs/timeutil/USER_MANUAL.md)** - Comprehensive user guide (~1500 lines) / 포괄적인 사용자 가이드
- **[Developer Guide / 개발자 가이드](../docs/timeutil/DEVELOPER_GUIDE.md)** - Architecture and implementation details (~1000 lines) / 아키텍처 및 구현 세부사항
- **[Design Plan / 설계 계획서](../docs/timeutil/DESIGN_PLAN.md)** - Design philosophy and decisions / 설계 철학 및 결정
- **[Work Plan / 작업 계획서](../docs/timeutil/WORK_PLAN.md)** - Implementation roadmap / 구현 로드맵

## Use Cases / 사용 사례

1. **Time Difference Calculations / 시간 차이 계산**: Calculate elapsed time in human-readable format / 사람이 읽기 쉬운 형식으로 경과 시간 계산
2. **Timezone Conversions / 타임존 변환**: Convert times across different timezones / 다른 타임존 간 시간 변환
3. **Date Arithmetic / 날짜 연산**: Add/subtract days, weeks, months easily / 일, 주, 월을 쉽게 더하기/빼기
4. **Business Day Calculations / 영업일 계산**: Calculate business days for scheduling / 일정 관리를 위한 영업일 계산
5. **Age Calculations / 나이 계산**: Calculate age from birth date / 생년월일로 나이 계산
6. **Date Formatting / 날짜 포맷팅**: Format dates with intuitive tokens / 직관적인 토큰으로 날짜 포맷
7. **Time Range Checks / 시간 범위 체크**: Check if time is in range, today, this week, etc. / 시간이 범위 내인지, 오늘인지, 이번 주인지 등 확인
8. **Relative Time Display / 상대 시간 표시**: Show "2 hours ago" style timestamps / "2시간 전" 스타일 타임스탬프 표시
9. **Unix Timestamp Operations / Unix 타임스탬프 작업**: Convert between time.Time and Unix timestamps / time.Time과 Unix 타임스탬프 간 변환
10. **Schedule Management / 일정 관리**: Manage schedules with start/end of period functions / 기간의 시작/끝 함수로 일정 관리

## Best Practices / 모범 사례

1. **Use timezone-aware functions / 타임존 인식 함수 사용**: Always consider timezones when working with dates / 날짜 작업 시 항상 타임존 고려
2. **Handle errors properly / 에러 적절히 처리**: Timezone and parsing functions return errors / 타임존 및 파싱 함수는 에러를 반환
3. **Use custom format tokens / 커스텀 포맷 토큰 사용**: Use `YYYY-MM-DD` instead of `2006-01-02` / `2006-01-02` 대신 `YYYY-MM-DD` 사용
4. **Set holidays for business days / 영업일에 공휴일 설정**: Configure holidays for accurate business day calculations / 정확한 영업일 계산을 위해 공휴일 설정
5. **Use relative time for UI / UI에 상대 시간 사용**: Show relative times for better UX / 더 나은 UX를 위해 상대 시간 표시

## Contributing / 기여하기

Contributions are welcome! Please see the [Developer Guide](../docs/timeutil/DEVELOPER_GUIDE.md) for details.

기여를 환영합니다! 자세한 내용은 [개발자 가이드](../docs/timeutil/DEVELOPER_GUIDE.md)를 참조하세요.

## License / 라이선스

This project is licensed under the MIT License - see the [LICENSE](../LICENSE) file for details.

이 프로젝트는 MIT 라이선스에 따라 배포됩니다 - 자세한 내용은 [LICENSE](../LICENSE) 파일을 참조하세요.

## Author / 작성자

**arkd0ng**

- GitHub: [@arkd0ng](https://github.com/arkd0ng)

## Changelog / 변경 이력

For detailed version history, see [CHANGELOG-v1.6.md](../docs/CHANGELOG/CHANGELOG-v1.6.md).

상세한 버전 히스토리는 [CHANGELOG-v1.6.md](../docs/CHANGELOG/CHANGELOG-v1.6.md)를 참조하세요.

## Version / 버전

Current version: **v1.6.x** (in development)

현재 버전: **v1.6.x** (개발 중)
