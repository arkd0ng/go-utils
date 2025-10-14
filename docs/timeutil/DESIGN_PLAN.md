# Timeutil Package - Design Plan / 설계 계획서
# timeutil 패키지 - 설계 계획서

**Version / 버전**: v1.6.x
**Author / 작성자**: arkd0ng
**Created / 작성일**: 2025-10-14
**Status / 상태**: Final Design - Extreme Simplicity / 최종 설계 - 극도의 간결함

---

## Table of Contents / 목차

1. [Why This Package Exists / 왜 이 패키지가 존재하는가](#why-this-package-exists--왜-이-패키지가-존재하는가)
2. [Design Philosophy / 설계 철학](#design-philosophy--설계-철학)
3. [What Users Get / 사용자가 얻는 것](#what-users-get--사용자가-얻는-것)
4. [API Design / API 설계](#api-design--api-설계)
5. [Implementation Architecture / 구현 아키텍처](#implementation-architecture--구현-아키텍처)
6. [File Structure / 파일 구조](#file-structure--파일-구조)
7. [Detailed Features / 상세 기능](#detailed-features--상세-기능)

---

## Why This Package Exists / 왜 이 패키지가 존재하는가

### The Problem / 문제점

Working with time and dates in Go using the standard `time` package requires:

Go의 표준 `time` 패키지를 사용하여 시간 및 날짜를 다루려면:

1. **Complex timezone handling / 복잡한 타임존 처리**:
   ```go
   // 타임존 변환
   loc, err := time.LoadLocation("America/New_York")
   if err != nil {
       return err
   }
   nyTime := time.Now().In(loc)

   // 다른 타임존으로 변환
   loc2, err := time.LoadLocation("Asia/Seoul")
   if err != nil {
       return err
   }
   seoulTime := nyTime.In(loc2)
   ```

2. **Manual time difference calculations / 수동 시간 차이 계산**:
   ```go
   // 두 시간의 차이 계산
   start := time.Now()
   // ... some work ...
   end := time.Now()
   duration := end.Sub(start)

   // 초, 분, 시간, 일 등으로 변환
   seconds := duration.Seconds()
   minutes := duration.Minutes()
   hours := duration.Hours()
   days := hours / 24

   // 사람이 읽기 쉬운 형식으로 변환
   if days > 0 {
       fmt.Printf("%d days %d hours", int(days), int(hours)%24)
   } else if hours > 0 {
       fmt.Printf("%d hours %d minutes", int(hours), int(minutes)%60)
   }
   // ... 20+ lines of code
   ```

3. **Complex date formatting / 복잡한 날짜 포맷팅**:
   ```go
   // ISO8601, RFC3339, Custom formats
   t := time.Now()

   // ISO8601
   iso := t.Format("2006-01-02T15:04:05Z07:00")

   // Custom format
   custom := t.Format("2006-01-02 15:04:05")

   // 매번 "2006-01-02" 같은 이상한 포맷 문자열 기억해야 함
   ```

4. **Manual date arithmetic / 수동 날짜 연산**:
   ```go
   // 날짜 더하기/빼기
   tomorrow := time.Now().Add(24 * time.Hour)
   nextWeek := time.Now().Add(7 * 24 * time.Hour)

   // 월 계산은 더 복잡
   now := time.Now()
   nextMonth := now.AddDate(0, 1, 0)

   // 월말, 월초 계산
   year, month, _ := now.Date()
   firstDay := time.Date(year, month, 1, 0, 0, 0, 0, now.Location())
   lastDay := time.Date(year, month+1, 0, 23, 59, 59, 0, now.Location())
   ```

5. **Business day calculations / 영업일 계산**:
   ```go
   // 주말 제외, 공휴일 제외 계산
   // 개발자가 직접 로직 작성
   current := start
   businessDays := 0
   for businessDays < targetDays {
       current = current.Add(24 * time.Hour)
       if current.Weekday() != time.Saturday && current.Weekday() != time.Sunday {
           businessDays++
       }
   }
   // ... 복잡한 로직
   ```

6. **Time range checks / 시간 범위 체크**:
   ```go
   // 시간이 특정 범위 내에 있는지 확인
   now := time.Now()
   if now.After(start) && now.Before(end) {
       // in range
   }

   // 오늘인지 확인
   today := time.Now()
   y1, m1, d1 := today.Date()
   y2, m2, d2 := checkDate.Date()
   if y1 == y2 && m1 == m2 && d1 == d2 {
       // is today
   }
   ```

### The Solution / 해결책

**이 패키지는 위의 모든 번거로움을 제거합니다**:

```go
// 1. Timezone conversion - 간단하게
seoulTime := timeutil.ConvertTimezone(time.Now(), "Asia/Seoul")
nyTime := timeutil.ConvertTimezone(seoulTime, "America/New_York")

// 2. Time difference - 사람이 읽기 쉽게
diff := timeutil.SubTime(start, end)
fmt.Println(diff.String()) // "2 days 3 hours 15 minutes"
fmt.Println(diff.Days())    // 2
fmt.Println(diff.Hours())   // 51

// 3. Date formatting - 직관적으로
iso := timeutil.FormatISO8601(time.Now())
rfc := timeutil.FormatRFC3339(time.Now())
custom := timeutil.Format(time.Now(), "YYYY-MM-DD HH:mm:ss")

// 4. Date arithmetic - 간단하게
tomorrow := timeutil.AddDays(time.Now(), 1)
nextWeek := timeutil.AddWeeks(time.Now(), 1)
nextMonth := timeutil.AddMonths(time.Now(), 1)
startOfMonth := timeutil.StartOfMonth(time.Now())
endOfMonth := timeutil.EndOfMonth(time.Now())

// 5. Business day calculations - 자동으로
nextBusinessDay := timeutil.AddBusinessDays(time.Now(), 5)
isBusinessDay := timeutil.IsBusinessDay(time.Now())

// 6. Time range checks - 직관적으로
if timeutil.IsBetween(now, start, end) {
    // in range
}
if timeutil.IsToday(checkDate) {
    // is today
}
```

### If It's Not This Simple, Don't Build It / 이 정도로 간단하지 않으면 만들지 마세요

**The Rule**: If a time/date operation takes more than 5 lines in standard Go, it should be 1 line in this package.

**규칙**: 표준 Go에서 5줄 이상 걸리는 시간/날짜 작업은 이 패키지에서 1줄이어야 합니다.

---

## Design Philosophy / 설계 철학

### 1. Extreme Simplicity: "20 lines → 1 line" / 극도의 간결함: "20줄 → 1줄"

**Before (Standard Go) / 이전 (표준 Go)**:
```go
// Calculate age in years / 나이 계산 (년)
birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
now := time.Now()

years := now.Year() - birthDate.Year()
if now.YearDay() < birthDate.YearDay() {
    years--
}
// 5+ lines
```

**After (This Package) / 이후 (이 패키지)**:
```go
age := timeutil.AgeInYears(birthDate)
// 1 line
```

### 2. Human-Readable / 사람이 읽기 쉽게

**Function names should read like natural language**:
**함수 이름은 자연어처럼 읽혀야 합니다**:

- `timeutil.IsToday(t)` - "Is today?"
- `timeutil.IsBefore(t1, t2)` - "Is before?"
- `timeutil.AddDays(t, 7)` - "Add days"
- `timeutil.StartOfWeek(t)` - "Start of week"

### 3. Zero Configuration / 제로 설정

**No setup required. Import and use**:
**설정 불필요. 임포트하고 사용**:

```go
import "github.com/arkd0ng/go-utils/timeutil"

// No New(), no Init(), just use / New(), Init() 없이 바로 사용
result := timeutil.AddDays(time.Now(), 7)
```

### 4. Unicode-Safe & i18n Ready / 유니코드 안전 및 국제화 지원

All functions should work correctly with any timezone and locale:

모든 함수는 모든 타임존과 로케일에서 올바르게 작동해야 합니다:

```go
// Works with any timezone / 모든 타임존에서 작동
seoulTime := timeutil.ConvertTimezone(time.Now(), "Asia/Seoul")
londonTime := timeutil.ConvertTimezone(time.Now(), "Europe/London")
```

### 5. Practical First / 실용성 우선

Focus on the 99% use cases that developers encounter daily:

개발자가 매일 마주치는 99% 사용 사례에 집중:

- Time difference calculations / 시간 차이 계산
- Timezone conversions / 타임존 변환
- Date arithmetic / 날짜 연산
- Business day calculations / 영업일 계산
- Time formatting / 시간 포맷팅

---

## What Users Get / 사용자가 얻는 것

### 1. Time Difference / 시간 차이

**Duration operations made simple**:
**Duration 작업을 간단하게**:

```go
// Calculate time difference / 시간 차이 계산
start := time.Date(2025, 1, 1, 9, 0, 0, 0, time.UTC)
end := time.Date(2025, 1, 3, 15, 30, 0, 0, time.UTC)

diff := timeutil.SubTime(start, end)
fmt.Println(diff.Days())       // 2
fmt.Println(diff.Hours())      // 54
fmt.Println(diff.Minutes())    // 3270
fmt.Println(diff.String())     // "2 days 6 hours 30 minutes"
fmt.Println(diff.Humanize())   // "2d 6h 30m"
```

### 2. Timezone Operations / 타임존 작업

**Timezone handling without pain**:
**고통 없는 타임존 처리**:

```go
// Convert to different timezone / 다른 타임존으로 변환
seoulTime := timeutil.ConvertTimezone(time.Now(), "Asia/Seoul")
nyTime := timeutil.ConvertTimezone(time.Now(), "America/New_York")

// Get timezone offset / 타임존 오프셋 가져오기
offset := timeutil.GetTimezoneOffset("Asia/Seoul") // +9

// List all timezones / 모든 타임존 나열
zones := timeutil.ListTimezones()
```

### 3. Date Arithmetic / 날짜 연산

**Add/subtract dates easily**:
**날짜 더하기/빼기 쉽게**:

```go
// Add/subtract time units / 시간 단위 더하기/빼기
tomorrow := timeutil.AddDays(time.Now(), 1)
yesterday := timeutil.AddDays(time.Now(), -1)
nextWeek := timeutil.AddWeeks(time.Now(), 1)
nextMonth := timeutil.AddMonths(time.Now(), 1)
nextYear := timeutil.AddYears(time.Now(), 1)

// Start/End of period / 기간의 시작/끝
startOfDay := timeutil.StartOfDay(time.Now())
endOfDay := timeutil.EndOfDay(time.Now())
startOfWeek := timeutil.StartOfWeek(time.Now())
endOfWeek := timeutil.EndOfWeek(time.Now())
startOfMonth := timeutil.StartOfMonth(time.Now())
endOfMonth := timeutil.EndOfMonth(time.Now())
startOfYear := timeutil.StartOfYear(time.Now())
endOfYear := timeutil.EndOfYear(time.Now())
```

### 4. Date Formatting / 날짜 포맷팅

**Human-readable format functions**:
**사람이 읽기 쉬운 포맷 함수**:

```go
// Common formats / 일반 포맷
iso := timeutil.FormatISO8601(time.Now())      // "2025-10-14T15:04:05+09:00"
rfc := timeutil.FormatRFC3339(time.Now())      // "2025-10-14T15:04:05+09:00"
date := timeutil.FormatDate(time.Now())        // "2025-10-14"
datetime := timeutil.FormatDateTime(time.Now()) // "2025-10-14 15:04:05"
time := timeutil.FormatTime(time.Now())        // "15:04:05"

// Custom format (no more "2006-01-02") / 커스텀 포맷
custom := timeutil.Format(time.Now(), "YYYY-MM-DD HH:mm:ss")
```

### 5. Time Parsing / 시간 파싱

**Parse time strings easily**:
**시간 문자열 쉽게 파싱**:

```go
// Parse common formats / 일반 포맷 파싱
t1, _ := timeutil.ParseISO8601("2025-10-14T15:04:05+09:00")
t2, _ := timeutil.ParseRFC3339("2025-10-14T15:04:05+09:00")
t3, _ := timeutil.ParseDate("2025-10-14")
t4, _ := timeutil.ParseDateTime("2025-10-14 15:04:05")

// Smart parse (auto-detect format) / 스마트 파싱 (포맷 자동 감지)
t5, _ := timeutil.Parse("2025-10-14")
t6, _ := timeutil.Parse("2025-10-14 15:04:05")
```

### 6. Business Days / 영업일

**Business day calculations**:
**영업일 계산**:

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
```

### 7. Time Comparisons / 시간 비교

**Intuitive comparison functions**:
**직관적인 비교 함수**:

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
```

### 8. Age Calculations / 나이 계산

**Age calculations made easy**:
**쉬운 나이 계산**:

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
```

### 9. Relative Time / 상대 시간

**Human-readable relative time**:
**사람이 읽기 쉬운 상대 시간**:

```go
// Relative time strings / 상대 시간 문자열
past := time.Now().Add(-2 * time.Hour)
fmt.Println(timeutil.RelativeTime(past))     // "2 hours ago"

future := time.Now().Add(3 * time.Hour)
fmt.Println(timeutil.RelativeTime(future))   // "in 3 hours"

// Short format / 짧은 포맷
fmt.Println(timeutil.RelativeTimeShort(past)) // "2h ago"
```

### 10. Unix Timestamp / Unix 타임스탬프

**Unix timestamp operations**:
**Unix 타임스탬프 작업**:

```go
// Current timestamp / 현재 타임스탬프
unix := timeutil.Now()                    // Unix timestamp in seconds
unixMilli := timeutil.NowMilli()          // Unix timestamp in milliseconds
unixMicro := timeutil.NowMicro()          // Unix timestamp in microseconds
unixNano := timeutil.NowNano()            // Unix timestamp in nanoseconds

// Convert from timestamp / 타임스탬프에서 변환
t := timeutil.FromUnix(1634198400)
t := timeutil.FromUnixMilli(1634198400000)

// Convert to timestamp / 타임스탬프로 변환
unix := timeutil.ToUnix(time.Now())
unixMilli := timeutil.ToUnixMilli(time.Now())
```

---

## API Design / API 설계

### Function Categories / 함수 카테고리

The package is organized into 10 logical categories:

패키지는 10개의 논리적 카테고리로 구성됩니다:

#### 1. Time Difference / 시간 차이 (8 functions)

```go
SubTime(t1, t2 time.Time) *TimeDiff
DiffInSeconds(t1, t2 time.Time) float64
DiffInMinutes(t1, t2 time.Time) float64
DiffInHours(t1, t2 time.Time) float64
DiffInDays(t1, t2 time.Time) float64
DiffInWeeks(t1, t2 time.Time) float64
DiffInMonths(t1, t2 time.Time) int
DiffInYears(t1, t2 time.Time) int
```

#### 2. Timezone Operations / 타임존 작업 (5 functions)

```go
ConvertTimezone(t time.Time, tz string) (time.Time, error)
GetTimezoneOffset(tz string) (int, error)
ListTimezones() []string
IsValidTimezone(tz string) bool
GetLocalTimezone() string
```

#### 3. Date Arithmetic / 날짜 연산 (16 functions)

```go
AddSeconds(t time.Time, seconds int) time.Time
AddMinutes(t time.Time, minutes int) time.Time
AddHours(t time.Time, hours int) time.Time
AddDays(t time.Time, days int) time.Time
AddWeeks(t time.Time, weeks int) time.Time
AddMonths(t time.Time, months int) time.Time
AddYears(t time.Time, years int) time.Time

StartOfDay(t time.Time) time.Time
EndOfDay(t time.Time) time.Time
StartOfWeek(t time.Time) time.Time
EndOfWeek(t time.Time) time.Time
StartOfMonth(t time.Time) time.Time
EndOfMonth(t time.Time) time.Time
StartOfYear(t time.Time) time.Time
EndOfYear(t time.Time) time.Time
StartOfQuarter(t time.Time) time.Time
```

#### 4. Date Formatting / 날짜 포맷팅 (8 functions)

```go
FormatISO8601(t time.Time) string
FormatRFC3339(t time.Time) string
FormatDate(t time.Time) string
FormatDateTime(t time.Time) string
FormatTime(t time.Time) string
Format(t time.Time, layout string) string
FormatCustom(t time.Time, layout string) string
FormatWithTimezone(t time.Time, tz string) (string, error)
```

#### 5. Time Parsing / 시간 파싱 (6 functions)

```go
ParseISO8601(s string) (time.Time, error)
ParseRFC3339(s string) (time.Time, error)
ParseDate(s string) (time.Time, error)
ParseDateTime(s string) (time.Time, error)
Parse(s string) (time.Time, error)
ParseWithTimezone(s, tz string) (time.Time, error)
```

#### 6. Business Days / 영업일 (6 functions)

```go
AddBusinessDays(t time.Time, days int) time.Time
IsBusinessDay(t time.Time) bool
CountBusinessDays(start, end time.Time) int
NextBusinessDay(t time.Time) time.Time
PreviousBusinessDay(t time.Time) time.Time
IsHoliday(t time.Time) bool
```

#### 7. Time Comparisons / 시간 비교 (18 functions)

```go
IsBefore(t1, t2 time.Time) bool
IsAfter(t1, t2 time.Time) bool
IsBetween(t, start, end time.Time) bool
IsToday(t time.Time) bool
IsYesterday(t time.Time) bool
IsTomorrow(t time.Time) bool
IsThisWeek(t time.Time) bool
IsThisMonth(t time.Time) bool
IsThisYear(t time.Time) bool
IsWeekend(t time.Time) bool
IsWeekday(t time.Time) bool
IsSameDay(t1, t2 time.Time) bool
IsSameWeek(t1, t2 time.Time) bool
IsSameMonth(t1, t2 time.Time) bool
IsSameYear(t1, t2 time.Time) bool
IsLeapYear(t time.Time) bool
IsPast(t time.Time) bool
IsFuture(t time.Time) bool
```

#### 8. Age Calculations / 나이 계산 (4 functions)

```go
AgeInYears(birthDate time.Time) int
AgeInMonths(birthDate time.Time) int
AgeInDays(birthDate time.Time) int
Age(birthDate time.Time) *Age
```

#### 9. Relative Time / 상대 시간 (3 functions)

```go
RelativeTime(t time.Time) string
RelativeTimeShort(t time.Time) string
TimeAgo(t time.Time) string
```

#### 10. Unix Timestamp / Unix 타임스탬프 (8 functions)

```go
Now() int64
NowMilli() int64
NowMicro() int64
NowNano() int64
FromUnix(sec int64) time.Time
FromUnixMilli(msec int64) time.Time
ToUnix(t time.Time) int64
ToUnixMilli(t time.Time) int64
```

### Total: ~80+ functions across 10 categories
### 총: 10개 카테고리에 걸쳐 약 80개 이상의 함수

---

## Implementation Architecture / 구현 아키텍처

### Package Structure / 패키지 구조

```go
package timeutil

import (
    "time"
)

// Core types / 핵심 타입
type TimeDiff struct {
    Duration time.Duration
    // Helper methods
}

type Age struct {
    Years  int
    Months int
    Days   int
}

// Helper constants / 헬퍼 상수
const (
    SecondsPerMinute = 60
    SecondsPerHour   = 3600
    SecondsPerDay    = 86400
    DaysPerWeek      = 7
    MonthsPerYear    = 12
)
```

### Key Implementation Decisions / 주요 구현 결정

1. **Pure Functions / 순수 함수**:
   - All functions are stateless / 모든 함수는 상태 비저장
   - No global state / 전역 상태 없음
   - Thread-safe by design / 설계상 스레드 안전

2. **Error Handling / 에러 처리**:
   - Functions return (result, error) for timezone operations
   - 타임존 작업은 (결과, error) 반환
   - Other functions panic-free with sensible defaults
   - 다른 함수는 합리적인 기본값으로 패닉 없음

3. **Performance / 성능**:
   - No unnecessary allocations / 불필요한 할당 없음
   - Efficient time calculations / 효율적인 시간 계산
   - Cached timezone locations / 캐시된 타임존 위치

4. **Unicode Support / 유니코드 지원**:
   - All functions work correctly with any locale / 모든 함수가 모든 로케일에서 올바르게 작동
   - UTF-8 safe / UTF-8 안전

---

## File Structure / 파일 구조

```
timeutil/
├── timeutil.go         # Package doc and core types / 패키지 문서 및 핵심 타입
├── diff.go             # Time difference functions / 시간 차이 함수
├── timezone.go         # Timezone operations / 타임존 작업
├── arithmetic.go       # Date arithmetic / 날짜 연산
├── format.go           # Date formatting / 날짜 포맷팅
├── parse.go            # Time parsing / 시간 파싱
├── business.go         # Business days / 영업일
├── comparison.go       # Time comparisons / 시간 비교
├── age.go              # Age calculations / 나이 계산
├── relative.go         # Relative time / 상대 시간
├── unix.go             # Unix timestamp / Unix 타임스탬프
├── constants.go        # Constants and helpers / 상수 및 헬퍼
├── diff_test.go        # Tests for diff.go
├── timezone_test.go    # Tests for timezone.go
├── arithmetic_test.go  # Tests for arithmetic.go
├── format_test.go      # Tests for format.go
├── parse_test.go       # Tests for parse.go
├── business_test.go    # Tests for business.go
├── comparison_test.go  # Tests for comparison.go
├── age_test.go         # Tests for age.go
├── relative_test.go    # Tests for relative.go
├── unix_test.go        # Tests for unix.go
└── README.md           # Package documentation / 패키지 문서
```

---

## Detailed Features / 상세 기능

### 1. TimeDiff Type / TimeDiff 타입

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

### 2. Age Type / Age 타입

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

### 3. Custom Format Strings / 커스텀 포맷 문자열

Instead of Go's confusing "2006-01-02" format, use intuitive format strings:

Go의 혼란스러운 "2006-01-02" 포맷 대신 직관적인 포맷 문자열 사용:

```go
// Custom format tokens / 커스텀 포맷 토큰
YYYY - 4-digit year / 4자리 년
YY   - 2-digit year / 2자리 년
MM   - 2-digit month / 2자리 월
M    - 1 or 2-digit month / 1 또는 2자리 월
DD   - 2-digit day / 2자리 일
D    - 1 or 2-digit day / 1 또는 2자리 일
HH   - 2-digit hour (24h) / 2자리 시 (24시간)
hh   - 2-digit hour (12h) / 2자리 시 (12시간)
mm   - 2-digit minute / 2자리 분
ss   - 2-digit second / 2자리 초

// Example / 예제
timeutil.Format(t, "YYYY-MM-DD HH:mm:ss")  // "2025-10-14 15:04:05"
timeutil.Format(t, "MM/DD/YYYY")           // "10/14/2025"
timeutil.Format(t, "YYYY년 MM월 DD일")      // "2025년 10월 14일"
```

### 4. Business Days Configuration / 영업일 설정

```go
// Default: Monday-Friday are business days / 기본: 월-금이 영업일
// Can be customized with holidays / 공휴일로 커스터마이징 가능

// Set custom holidays / 커스텀 공휴일 설정
holidays := []time.Time{
    time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),  // New Year
    time.Date(2025, 12, 25, 0, 0, 0, 0, time.UTC), // Christmas
}
timeutil.SetHolidays(holidays)

// Check if holiday / 공휴일 확인
if timeutil.IsHoliday(someDate) {
    // It's a holiday
}
```

### 5. Timezone Caching / 타임존 캐싱

```go
// Internally cache loaded timezones for performance
// 성능을 위해 로드된 타임존을 내부적으로 캐시

var timezoneCache = make(map[string]*time.Location)
var timezoneCacheMu sync.RWMutex

func loadTimezone(tz string) (*time.Location, error) {
    timezoneCacheMu.RLock()
    if loc, ok := timezoneCache[tz]; ok {
        timezoneCacheMu.RUnlock()
        return loc, nil
    }
    timezoneCacheMu.RUnlock()

    loc, err := time.LoadLocation(tz)
    if err != nil {
        return nil, err
    }

    timezoneCacheMu.Lock()
    timezoneCache[tz] = loc
    timezoneCacheMu.Unlock()

    return loc, nil
}
```

---

## Design Principles Summary / 설계 원칙 요약

1. **Extreme Simplicity / 극도의 간결함**: "20 lines → 1 line"
2. **Human-Readable / 사람이 읽기 쉽게**: Function names read like natural language
3. **Zero Configuration / 제로 설정**: No setup, just import and use
4. **Unicode-Safe / 유니코드 안전**: Works with any timezone and locale
5. **Practical First / 실용성 우선**: Focus on 99% use cases
6. **Performance / 성능**: Efficient with caching and zero allocations
7. **Thread-Safe / 스레드 안전**: All functions are thread-safe
8. **Error-Friendly / 에러 친화적**: Clear error messages, no panics

---

## Success Criteria / 성공 기준

This package is successful if:

이 패키지가 성공적이려면:

1. ✅ Time difference calculations: 10+ lines → 1 line
2. ✅ Timezone conversions: 5+ lines → 1 line
3. ✅ Date arithmetic: 5+ lines → 1 line
4. ✅ Business day calculations: 20+ lines → 1 line
5. ✅ Time formatting: Intuitive format strings (no "2006-01-02")
6. ✅ Age calculations: 10+ lines → 1 line
7. ✅ ~80+ functions covering all common time/date operations
8. ✅ Zero external dependencies (standard library only)
9. ✅ 100% test coverage with comprehensive benchmarks
10. ✅ Bilingual documentation (English/Korean)

---

**Design Status / 설계 상태**: ✅ **APPROVED - Ready for Implementation / 승인됨 - 구현 준비 완료**

**Next Step / 다음 단계**: Create WORK_PLAN.md / 작업 계획서 작성
