# Timeutil Package - User Manual / 사용자 매뉴얼

**Version / 버전**: v1.6.008
**Package / 패키지**: `github.com/arkd0ng/go-utils/timeutil`
**Go Version / Go 버전**: 1.16+

---

## Table of Contents / 목차

1. [Introduction / 소개](#introduction--소개)
2. [Installation / 설치](#installation--설치)
3. [Quick Start / 빠른 시작](#quick-start--빠른-시작)
4. [Core Concepts / 핵심 개념](#core-concepts--핵심-개념)
5. [Function Reference / 함수 참조](#function-reference--함수-참조)
   - [Time Difference / 시간 차이](#time-difference--시간-차이)
   - [Timezone Operations / 타임존 작업](#timezone-operations--타임존-작업)
   - [Date Arithmetic / 날짜 연산](#date-arithmetic--날짜-연산)
   - [Date Formatting / 날짜 포맷팅](#date-formatting--날짜-포맷팅)
   - [Time Parsing / 시간 파싱](#time-parsing--시간-파싱)
   - [Time Comparisons / 시간 비교](#time-comparisons--시간-비교)
   - [Age Calculations / 나이 계산](#age-calculations--나이-계산)
   - [Relative Time / 상대 시간](#relative-time--상대-시간)
   - [Unix Timestamp / Unix 타임스탬프](#unix-timestamp--unix-타임스탬프)
   - [Business Days / 영업일](#business-days--영업일)
   - [**NEW! String Parameters** / 문자열 매개변수](#string-parameters--문자열-매개변수)
6. [Common Use Cases / 일반적인 사용 사례](#common-use-cases--일반적인-사용-사례)
7. [Best Practices / 모범 사례](#best-practices--모범-사례)
8. [Troubleshooting / 문제 해결](#troubleshooting--문제-해결)
9. [FAQ / 자주 묻는 질문](#faq--자주-묻는-질문)

---

## Introduction / 소개

The `timeutil` package provides an extremely simple and intuitive API for time and date operations in Go. It reduces 20+ lines of repetitive time manipulation code to just 1-2 lines.

`timeutil` 패키지는 Go에서 시간 및 날짜 작업을 위한 극도로 간단하고 직관적인 API를 제공합니다. 20줄 이상의 반복적인 시간 조작 코드를 단 1-2줄로 줄입니다.

### Key Features / 주요 기능

- **150+ functions** organized into 11 categories / 11개 카테고리로 구성된 150개 이상의 함수
- **NEW in v1.6.008**: String parameter support - parse any format automatically! / 문자열 매개변수 지원 - 모든 포맷 자동 파싱!
- **40+ time formats** automatically detected (including Korean!) / 40개 이상의 시간 포맷 자동 감지 (한글 포함!)
- **KST (GMT+9) default timezone** for Korean users / 한국 사용자를 위한 KST (GMT+9) 기본 타임존
- **Custom format tokens** like YYYY-MM-DD instead of Go's confusing 2006-01-02 / Go의 혼란스러운 2006-01-02 대신 YYYY-MM-DD 같은 커스텀 포맷 토큰
- **Business day support** with Korean holiday calendar / 한국 공휴일 달력이 포함된 영업일 지원
- **Human-readable output** for time differences and age / 시간 차이 및 나이에 대한 사람이 읽기 쉬운 출력
- **Thread-safe** timezone caching / 스레드 안전 타임존 캐싱
- **Zero dependencies** - standard library only / 제로 의존성 - 표준 라이브러리만 사용

### Design Philosophy / 설계 철학

**"20 lines → 1 line"** - Extreme simplicity

The goal is to make time operations so simple that you never need to Google "how to format date in Go" again.

목표는 시간 작업을 너무 간단하게 만들어서 "Go에서 날짜를 포맷하는 방법"을 다시 Google에서 검색할 필요가 없도록 하는 것입니다.

---

## Installation / 설치

### Prerequisites / 전제 조건

- Go 1.16 or higher / Go 1.16 이상
- No external dependencies / 외부 의존성 없음

### Install Package / 패키지 설치

```bash
go get github.com/arkd0ng/go-utils/timeutil
```

### Import / 임포트

```go
import "github.com/arkd0ng/go-utils/timeutil"
```

---

## Quick Start / 빠른 시작

### Example 1: Time Difference / 예제 1: 시간 차이

```go
package main

import (
    "fmt"
    "time"
    "github.com/arkd0ng/go-utils/timeutil"
)

func main() {
    start := time.Date(2025, 1, 1, 9, 0, 0, 0, time.UTC)
    end := time.Date(2025, 1, 3, 15, 30, 0, 0, time.UTC)

    // Calculate time difference / 시간 차이 계산
    diff := timeutil.SubTime(start, end)

    fmt.Println(diff.String())    // "2 days 6 hours 30 minutes"
    fmt.Println(diff.Humanize())  // "2d 6h 30m"
    fmt.Println(diff.Days())      // 2.270833333333333
}
```

### Example 2: Date Formatting / 예제 2: 날짜 포맷팅

```go
now := time.Now()

// Custom format tokens (no more 2006-01-02!)
// 커스텀 포맷 토큰 (더 이상 2006-01-02 안함!)
formatted := timeutil.Format(now, "YYYY-MM-DD HH:mm:ss")
fmt.Println(formatted) // "2025-10-14 15:04:05"

// Korean format / 한국어 포맷
korean := timeutil.FormatKorean(now)
fmt.Println(korean) // "2025년 10월 14일 15시 04분 05초"
```

### Example 3: Timezone Operations / 예제 3: 타임존 작업

```go
// Get current time in KST (default timezone)
// KST (기본 타임존)로 현재 시간 가져오기
kstNow := timeutil.NowKST()
fmt.Println(kstNow) // 2025-10-14 15:04:05 +0900 KST

// Convert to different timezone / 다른 타임존으로 변환
nyTime, _ := timeutil.ConvertTimezone(time.Now(), "America/New_York")
fmt.Println(nyTime) // 2025-10-14 02:04:05 -0400 EDT
```

### Example 4: Business Days / 예제 4: 영업일

```go
// Add Korean holidays for 2025 / 2025년 한국 공휴일 추가
timeutil.AddKoreanHolidays(2025)

today := time.Now()

// Check if today is a business day / 오늘이 영업일인지 확인
if timeutil.IsBusinessDay(today) {
    fmt.Println("Today is a business day") // "오늘은 영업일입니다"
}

// Add 5 business days (skips weekends and holidays)
// 5 영업일 추가 (주말 및 공휴일 건너뜀)
nextBizDay := timeutil.AddBusinessDays(today, 5)
fmt.Println(nextBizDay)
```

### Example 5: Relative Time / 예제 5: 상대 시간

```go
past := time.Now().Add(-2 * time.Hour)
future := time.Now().Add(3 * 24 * time.Hour)

fmt.Println(timeutil.RelativeTime(past))   // "2 hours ago"
fmt.Println(timeutil.RelativeTime(future)) // "in 3 days"

// Short format / 짧은 포맷
fmt.Println(timeutil.RelativeTimeShort(past)) // "2h ago"
```

---

## Core Concepts / 핵심 개념

### 1. KST Default Timezone / KST 기본 타임존

All functions use **Asia/Seoul (KST, GMT+9)** as the default timezone unless explicitly specified.

모든 함수는 명시적으로 지정하지 않는 한 **Asia/Seoul (KST, GMT+9)**를 기본 타임존으로 사용합니다.

```go
// These all use KST by default / 이것들은 모두 기본적으로 KST를 사용합니다
now := timeutil.NowKST()
formatted := timeutil.Format(time.Now(), "YYYY-MM-DD")
parsed, _ := timeutil.Parse("2025-10-14")
```

You can change the default timezone:

기본 타임존을 변경할 수 있습니다:

```go
// Change default to UTC / 기본값을 UTC로 변경
timeutil.SetDefaultTimezone("UTC")

// Change default to New York / 기본값을 뉴욕으로 변경
timeutil.SetDefaultTimezone("America/New_York")
```

### 2. Custom Format Tokens / 커스텀 포맷 토큰

Instead of Go's confusing reference date (2006-01-02), use intuitive tokens:

Go의 혼란스러운 참조 날짜(2006-01-02) 대신 직관적인 토큰을 사용하세요:

| Token / 토큰 | Meaning / 의미 | Example / 예제 |
|--------------|----------------|----------------|
| `YYYY` | 4-digit year / 4자리 연도 | 2025 |
| `YY` | 2-digit year / 2자리 연도 | 25 |
| `MM` | Month (01-12) / 월 (01-12) | 10 |
| `DD` | Day (01-31) / 일 (01-31) | 14 |
| `HH` | Hour (00-23) / 시 (00-23) | 15 |
| `mm` | Minute (00-59) / 분 (00-59) | 04 |
| `ss` | Second (00-59) / 초 (00-59) | 05 |

```go
// Easy to remember! / 기억하기 쉬움!
timeutil.Format(time.Now(), "YYYY-MM-DD HH:mm:ss")
timeutil.Format(time.Now(), "YYYY/MM/DD")
timeutil.Format(time.Now(), "DD-MM-YYYY")
```

### 3. TimeDiff Type / TimeDiff 타입

The `TimeDiff` type wraps `time.Duration` and provides human-readable methods:

`TimeDiff` 타입은 `time.Duration`을 래핑하고 사람이 읽기 쉬운 메서드를 제공합니다:

```go
diff := timeutil.SubTime(start, end)

// Get duration in different units / 다양한 단위로 기간 가져오기
seconds := diff.Seconds()       // 216000.0
minutes := diff.Minutes()       // 3600.0
hours := diff.Hours()           // 60.0
days := diff.Days()             // 2.5

// Human-readable output / 사람이 읽기 쉬운 출력
fmt.Println(diff.String())      // "2 days 6 hours 30 minutes"
fmt.Println(diff.Humanize())    // "2d 6h 30m"
```

### 4. AgeDetail Type / AgeDetail 타입

The `AgeDetail` type provides detailed age breakdown:

`AgeDetail` 타입은 상세한 나이 분석을 제공합니다:

```go
birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
age := timeutil.Age(birthDate)

fmt.Println(age.Years)         // 35
fmt.Println(age.Months)        // 4
fmt.Println(age.Days)          // 29
fmt.Println(age.String())      // "35 years 4 months 29 days"
```

### 5. Business Days / 영업일

Business day calculations automatically skip weekends (Saturday and Sunday) and registered holidays.

영업일 계산은 자동으로 주말(토요일 및 일요일)과 등록된 공휴일을 건너뜁니다.

```go
// Add Korean holidays / 한국 공휴일 추가
timeutil.AddKoreanHolidays(2025)

// Manually add custom holidays / 수동으로 커스텀 공휴일 추가
holidays := []time.Time{
    time.Date(2025, 12, 25, 0, 0, 0, 0, timeutil.KST), // Christmas
}
timeutil.SetHolidays(holidays)

// Check if day is a business day / 영업일인지 확인
if timeutil.IsBusinessDay(today) {
    fmt.Println("Open for business")
}
```

---

## Function Reference / 함수 참조

### Time Difference / 시간 차이

Calculate the difference between two times and get human-readable output.

두 시간 사이의 차이를 계산하고 사람이 읽기 쉬운 출력을 얻습니다.

#### SubTime(t1, t2 time.Time) *TimeDiff

Returns the difference between two times as a `TimeDiff` object.

두 시간 사이의 차이를 `TimeDiff` 객체로 반환합니다.

```go
start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
end := time.Date(2025, 1, 10, 12, 30, 45, 0, time.UTC)

diff := timeutil.SubTime(start, end)
fmt.Println(diff.String()) // "9 days 12 hours 30 minutes 45 seconds"
```

#### DiffInSeconds(t1, t2 time.Time) float64

Returns the difference in seconds.

초 단위로 차이를 반환합니다.

```go
seconds := timeutil.DiffInSeconds(start, end)
fmt.Println(seconds) // 820245.0
```

#### DiffInMinutes(t1, t2 time.Time) float64

Returns the difference in minutes.

분 단위로 차이를 반환합니다.

```go
minutes := timeutil.DiffInMinutes(start, end)
fmt.Println(minutes) // 13670.75
```

#### DiffInHours(t1, t2 time.Time) float64

Returns the difference in hours.

시간 단위로 차이를 반환합니다.

```go
hours := timeutil.DiffInHours(start, end)
fmt.Println(hours) // 228.5125
```

#### DiffInDays(t1, t2 time.Time) float64

Returns the difference in days.

일 단위로 차이를 반환합니다.

```go
days := timeutil.DiffInDays(start, end)
fmt.Println(days) // 9.521354166666666
```

#### DiffInWeeks(t1, t2 time.Time) float64

Returns the difference in weeks.

주 단위로 차이를 반환합니다.

```go
weeks := timeutil.DiffInWeeks(start, end)
fmt.Println(weeks) // 1.3601934523809523
```

#### DiffInMonths(t1, t2 time.Time) int

Returns the approximate difference in months.

대략적인 월 단위 차이를 반환합니다.

```go
months := timeutil.DiffInMonths(start, end)
fmt.Println(months) // 0 (less than a month)
```

#### DiffInYears(t1, t2 time.Time) int

Returns the approximate difference in years.

대략적인 년 단위 차이를 반환합니다.

```go
start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
end := time.Date(2025, 6, 15, 0, 0, 0, 0, time.UTC)
years := timeutil.DiffInYears(start, end)
fmt.Println(years) // 5
```

---

### Timezone Operations / 타임존 작업

Convert times between timezones and get timezone information.

타임존 간 시간을 변환하고 타임존 정보를 가져옵니다.

#### ConvertTimezone(t time.Time, tz string) (time.Time, error)

Converts a time to a different timezone.

시간을 다른 타임존으로 변환합니다.

```go
now := time.Now()
tokyo, err := timeutil.ConvertTimezone(now, "Asia/Tokyo")
if err != nil {
    log.Fatal(err)
}
fmt.Println(tokyo) // 2025-10-14 16:04:05 +0900 JST
```

#### ToUTC(t time.Time) time.Time

Converts time to UTC timezone.

시간을 UTC 타임존으로 변환합니다.

```go
kstTime := time.Now().In(timeutil.KST)
utcTime := timeutil.ToUTC(kstTime)
fmt.Println(utcTime) // 2025-10-14 06:04:05 +0000 UTC
```

#### ToKST(t time.Time) time.Time

Converts time to KST (Asia/Seoul) timezone.

시간을 KST (Asia/Seoul) 타임존으로 변환합니다.

```go
utcTime := time.Now().UTC()
kstTime := timeutil.ToKST(utcTime)
fmt.Println(kstTime) // 2025-10-14 15:04:05 +0900 KST
```

#### NowKST() time.Time

Returns the current time in KST timezone.

KST 타임존의 현재 시간을 반환합니다.

```go
now := timeutil.NowKST()
fmt.Println(now) // 2025-10-14 15:04:05 +0900 KST
```

#### GetTimezoneOffset(t time.Time) int

Returns the timezone offset in seconds from UTC.

UTC로부터의 타임존 오프셋을 초 단위로 반환합니다.

```go
offset := timeutil.GetTimezoneOffset(time.Now())
fmt.Println(offset) // 32400 (9 hours * 3600 seconds)
```

#### SetDefaultTimezone(tz string) error

Sets the default timezone for all package functions.

모든 패키지 함수의 기본 타임존을 설정합니다.

```go
// Change to UTC / UTC로 변경
err := timeutil.SetDefaultTimezone("UTC")
if err != nil {
    log.Fatal(err)
}

// Now all functions use UTC by default
// 이제 모든 함수가 기본적으로 UTC를 사용합니다
```

#### GetDefaultTimezone() string

Returns the current default timezone name.

현재 기본 타임존 이름을 반환합니다.

```go
tz := timeutil.GetDefaultTimezone()
fmt.Println(tz) // "Asia/Seoul"
```

#### GetLocalTimezone() string

Returns the system's local timezone name.

시스템의 로컬 타임존 이름을 반환합니다.

```go
local := timeutil.GetLocalTimezone()
fmt.Println(local) // "Asia/Seoul" (depends on system)
```

#### IsValidTimezone(tz string) bool

Checks if a timezone string is valid.

타임존 문자열이 유효한지 확인합니다.

```go
if timeutil.IsValidTimezone("America/New_York") {
    fmt.Println("Valid timezone")
}

if !timeutil.IsValidTimezone("Invalid/Timezone") {
    fmt.Println("Invalid timezone")
}
```

#### ListTimezones() []string

Returns a list of common timezone names.

일반적인 타임존 이름 목록을 반환합니다.

```go
timezones := timeutil.ListTimezones()
for _, tz := range timezones {
    fmt.Println(tz)
}
// Output:
// UTC
// America/New_York
// America/Los_Angeles
// Europe/London
// Asia/Tokyo
// Asia/Seoul
// ...
```

---

### Date Arithmetic / 날짜 연산

Add or subtract time units and get start/end of periods.

시간 단위를 더하거나 빼고 기간의 시작/끝을 가져옵니다.

#### AddSeconds(t time.Time, seconds int) time.Time

Adds seconds to a time.

시간에 초를 더합니다.

```go
now := time.Now()
later := timeutil.AddSeconds(now, 30)
fmt.Println(later) // 30 seconds later / 30초 후
```

#### AddMinutes(t time.Time, minutes int) time.Time

Adds minutes to a time.

시간에 분을 더합니다.

```go
later := timeutil.AddMinutes(now, 15)
fmt.Println(later) // 15 minutes later / 15분 후
```

#### AddHours(t time.Time, hours int) time.Time

Adds hours to a time.

시간에 시간을 더합니다.

```go
later := timeutil.AddHours(now, 2)
fmt.Println(later) // 2 hours later / 2시간 후
```

#### AddDays(t time.Time, days int) time.Time

Adds days to a time.

시간에 일을 더합니다.

```go
tomorrow := timeutil.AddDays(now, 1)
yesterday := timeutil.AddDays(now, -1)
```

#### AddWeeks(t time.Time, weeks int) time.Time

Adds weeks to a time.

시간에 주를 더합니다.

```go
nextWeek := timeutil.AddWeeks(now, 1)
fmt.Println(nextWeek) // 7 days later / 7일 후
```

#### AddMonths(t time.Time, months int) time.Time

Adds months to a time.

시간에 월을 더합니다.

```go
nextMonth := timeutil.AddMonths(now, 1)
lastYear := timeutil.AddMonths(now, -12)
```

#### AddYears(t time.Time, years int) time.Time

Adds years to a time.

시간에 년을 더합니다.

```go
nextYear := timeutil.AddYears(now, 1)
decade := timeutil.AddYears(now, 10)
```

#### StartOfDay(t time.Time) time.Time

Returns the start of the day (00:00:00).

하루의 시작(00:00:00)을 반환합니다.

```go
start := timeutil.StartOfDay(time.Now())
fmt.Println(start) // 2025-10-14 00:00:00 +0900 KST
```

#### EndOfDay(t time.Time) time.Time

Returns the end of the day (23:59:59).

하루의 끝(23:59:59)을 반환합니다.

```go
end := timeutil.EndOfDay(time.Now())
fmt.Println(end) // 2025-10-14 23:59:59 +0900 KST
```

#### StartOfWeek(t time.Time) time.Time

Returns the start of the week (Monday 00:00:00).

주의 시작(월요일 00:00:00)을 반환합니다.

```go
start := timeutil.StartOfWeek(time.Now())
fmt.Println(start) // 2025-10-13 00:00:00 +0900 KST (Monday)
```

#### EndOfWeek(t time.Time) time.Time

Returns the end of the week (Sunday 23:59:59).

주의 끝(일요일 23:59:59)을 반환합니다.

```go
end := timeutil.EndOfWeek(time.Now())
fmt.Println(end) // 2025-10-19 23:59:59 +0900 KST (Sunday)
```

#### StartOfMonth(t time.Time) time.Time

Returns the start of the month (1st day 00:00:00).

월의 시작(1일 00:00:00)을 반환합니다.

```go
start := timeutil.StartOfMonth(time.Now())
fmt.Println(start) // 2025-10-01 00:00:00 +0900 KST
```

#### EndOfMonth(t time.Time) time.Time

Returns the end of the month (last day 23:59:59).

월의 끝(마지막 날 23:59:59)을 반환합니다.

```go
end := timeutil.EndOfMonth(time.Now())
fmt.Println(end) // 2025-10-31 23:59:59 +0900 KST
```

#### StartOfYear(t time.Time) time.Time

Returns the start of the year (Jan 1 00:00:00).

년의 시작(1월 1일 00:00:00)을 반환합니다.

```go
start := timeutil.StartOfYear(time.Now())
fmt.Println(start) // 2025-01-01 00:00:00 +0900 KST
```

#### EndOfYear(t time.Time) time.Time

Returns the end of the year (Dec 31 23:59:59).

년의 끝(12월 31일 23:59:59)을 반환합니다.

```go
end := timeutil.EndOfYear(time.Now())
fmt.Println(end) // 2025-12-31 23:59:59 +0900 KST
```

#### StartOfQuarter(t time.Time) time.Time

Returns the start of the quarter.

분기의 시작을 반환합니다.

```go
start := timeutil.StartOfQuarter(time.Now())
fmt.Println(start) // 2025-10-01 00:00:00 +0900 KST (Q4)
```

---

### Date Formatting / 날짜 포맷팅

Format times using intuitive custom tokens or standard formats.

직관적인 커스텀 토큰 또는 표준 포맷을 사용하여 시간을 포맷합니다.

#### Format(t time.Time, layout string) string

Formats time using custom tokens (YYYY-MM-DD).

커스텀 토큰(YYYY-MM-DD)을 사용하여 시간을 포맷합니다.

```go
now := time.Now()

// Various formats / 다양한 포맷
fmt.Println(timeutil.Format(now, "YYYY-MM-DD"))             // "2025-10-14"
fmt.Println(timeutil.Format(now, "YYYY-MM-DD HH:mm:ss"))   // "2025-10-14 15:04:05"
fmt.Println(timeutil.Format(now, "DD/MM/YYYY"))            // "14/10/2025"
fmt.Println(timeutil.Format(now, "YYYY년 MM월 DD일"))       // "2025년 10월 14일"
fmt.Println(timeutil.Format(now, "HH:mm"))                  // "15:04"
```

#### FormatISO8601(t time.Time) string

Formats time in ISO 8601 format (YYYY-MM-DD).

ISO 8601 포맷(YYYY-MM-DD)으로 시간을 포맷합니다.

```go
iso := timeutil.FormatISO8601(time.Now())
fmt.Println(iso) // "2025-10-14"
```

#### FormatRFC3339(t time.Time) string

Formats time in RFC 3339 format.

RFC 3339 포맷으로 시간을 포맷합니다.

```go
rfc := timeutil.FormatRFC3339(time.Now())
fmt.Println(rfc) // "2025-10-14T15:04:05+09:00"
```

#### FormatDate(t time.Time) string

Formats time as date only (YYYY-MM-DD).

날짜만 포맷합니다(YYYY-MM-DD).

```go
date := timeutil.FormatDate(time.Now())
fmt.Println(date) // "2025-10-14"
```

#### FormatDateTime(t time.Time) string

Formats time as date and time (YYYY-MM-DD HH:mm:ss).

날짜와 시간을 포맷합니다(YYYY-MM-DD HH:mm:ss).

```go
datetime := timeutil.FormatDateTime(time.Now())
fmt.Println(datetime) // "2025-10-14 15:04:05"
```

#### FormatTime(t time.Time) string

Formats time as time only (HH:mm:ss).

시간만 포맷합니다(HH:mm:ss).

```go
timeOnly := timeutil.FormatTime(time.Now())
fmt.Println(timeOnly) // "15:04:05"
```

#### FormatKorean(t time.Time) string

Formats time in Korean style.

한국 스타일로 시간을 포맷합니다.

```go
korean := timeutil.FormatKorean(time.Now())
fmt.Println(korean) // "2025년 10월 14일 15시 04분 05초"
```

#### FormatCustom(t time.Time, layout string) string

Formats time using Go's standard layout (2006-01-02).

Go의 표준 레이아웃(2006-01-02)을 사용하여 시간을 포맷합니다.

```go
custom := timeutil.FormatCustom(time.Now(), "Jan 2, 2006 at 3:04 PM")
fmt.Println(custom) // "Oct 14, 2025 at 3:04 PM"
```

#### FormatWithTimezone(t time.Time, tz string) (string, error)

Formats time in a specific timezone using default DateTime format (YYYY-MM-DD HH:mm:ss).

기본 DateTime 포맷(YYYY-MM-DD HH:mm:ss)을 사용하여 특정 타임존으로 시간을 포맷합니다.

```go
formatted, err := timeutil.FormatWithTimezone(time.Now(), "America/New_York")
if err != nil {
    log.Fatal(err)
}
fmt.Println(formatted) // "2025-10-14 02:04:05" (EDT)
```

#### WeekdayKorean(t time.Time) string

Returns the Korean name of the weekday.

요일의 한글 이름을 반환합니다.

```go
t := time.Date(2025, 10, 14, 0, 0, 0, 0, time.UTC) // Tuesday
weekday := timeutil.WeekdayKorean(t)
fmt.Println(weekday) // "화요일"

// All weekdays / 모든 요일:
// Sunday    -> "일요일"
// Monday    -> "월요일"
// Tuesday   -> "화요일"
// Wednesday -> "수요일"
// Thursday  -> "목요일"
// Friday    -> "금요일"
// Saturday  -> "토요일"
```

#### WeekdayKoreanShort(t time.Time) string

Returns the short Korean name of the weekday.

요일의 짧은 한글 이름을 반환합니다.

```go
t := time.Date(2025, 10, 14, 0, 0, 0, 0, time.UTC) // Tuesday
weekday := timeutil.WeekdayKoreanShort(t)
fmt.Println(weekday) // "화"

// All short weekdays / 모든 짧은 요일:
// Sunday    -> "일"
// Monday    -> "월"
// Tuesday   -> "화"
// Wednesday -> "수"
// Thursday  -> "목"
// Friday    -> "금"
// Saturday  -> "토"
```

#### FormatKoreanDateTime(t time.Time) string

Formats a time in Korean format with weekday (YYYY년 MM월 DD일 (요일) HH시 mm분 ss초).

요일을 포함한 한국어 포맷으로 시간을 포맷합니다(YYYY년 MM월 DD일 (요일) HH시 mm분 ss초).

```go
t := time.Date(2025, 10, 14, 15, 30, 0, 0, time.UTC)
formatted := timeutil.FormatKoreanDateTime(t)
fmt.Println(formatted) // "2025년 10월 14일 (화요일) 15시 30분 00초"
```

#### FormatKoreanDateWithWeekday(t time.Time) string

Formats a date in Korean format with full weekday (YYYY년 MM월 DD일 (요일)).

전체 요일을 포함한 한국어 날짜 포맷으로 포맷합니다(YYYY년 MM월 DD일 (요일)).

```go
t := time.Date(2025, 10, 14, 0, 0, 0, 0, time.UTC)
formatted := timeutil.FormatKoreanDateWithWeekday(t)
fmt.Println(formatted) // "2025년 10월 14일 (화요일)"
```

#### FormatKoreanDateShort(t time.Time) string

Formats a date in Korean format with short weekday (YYYY년 MM월 DD일 (요일)).

짧은 요일을 포함한 한국어 날짜 포맷으로 포맷합니다(YYYY년 MM월 DD일 (요일)).

```go
t := time.Date(2025, 10, 14, 0, 0, 0, 0, time.UTC)
formatted := timeutil.FormatKoreanDateShort(t)
fmt.Println(formatted) // "2025년 10월 14일 (화)"
```

---

### Time Parsing / 시간 파싱

Parse strings into time.Time objects with automatic format detection.

자동 포맷 감지를 통해 문자열을 time.Time 객체로 파싱합니다.

#### Parse(s string) (time.Time, error)

Parses a time string with automatic format detection.

자동 포맷 감지를 사용하여 시간 문자열을 파싱합니다.

```go
// Automatically detects format / 자동으로 포맷 감지
t1, _ := timeutil.Parse("2025-10-14")
t2, _ := timeutil.Parse("2025-10-14 15:04:05")
t3, _ := timeutil.Parse("2025-10-14T15:04:05+09:00")

fmt.Println(t1) // 2025-10-14 00:00:00 +0900 KST
fmt.Println(t2) // 2025-10-14 15:04:05 +0900 KST
fmt.Println(t3) // 2025-10-14 15:04:05 +0900 KST
```

#### ParseISO8601(s string) (time.Time, error)

Parses ISO 8601 format (YYYY-MM-DD).

ISO 8601 포맷(YYYY-MM-DD)을 파싱합니다.

```go
t, err := timeutil.ParseISO8601("2025-10-14")
if err != nil {
    log.Fatal(err)
}
fmt.Println(t) // 2025-10-14 00:00:00 +0900 KST
```

#### ParseRFC3339(s string) (time.Time, error)

Parses RFC 3339 format.

RFC 3339 포맷을 파싱합니다.

```go
t, err := timeutil.ParseRFC3339("2025-10-14T15:04:05+09:00")
if err != nil {
    log.Fatal(err)
}
fmt.Println(t) // 2025-10-14 15:04:05 +0900 KST
```

#### ParseDate(s string) (time.Time, error)

Parses date string (YYYY-MM-DD).

날짜 문자열(YYYY-MM-DD)을 파싱합니다.

```go
date, err := timeutil.ParseDate("2025-10-14")
if err != nil {
    log.Fatal(err)
}
fmt.Println(date) // 2025-10-14 00:00:00 +0900 KST
```

#### ParseDateTime(s string) (time.Time, error)

Parses datetime string (YYYY-MM-DD HH:mm:ss).

날짜시간 문자열(YYYY-MM-DD HH:mm:ss)을 파싱합니다.

```go
dt, err := timeutil.ParseDateTime("2025-10-14 15:04:05")
if err != nil {
    log.Fatal(err)
}
fmt.Println(dt) // 2025-10-14 15:04:05 +0900 KST
```

#### ParseWithTimezone(s, tz string) (time.Time, error)

Parses time string in a specific timezone.

특정 타임존으로 시간 문자열을 파싱합니다.

```go
t, err := timeutil.ParseWithTimezone("2025-10-14 15:04:05", "America/New_York")
if err != nil {
    log.Fatal(err)
}
fmt.Println(t) // 2025-10-14 15:04:05 -0400 EDT
```

---

### Time Comparisons / 시간 비교

Compare times and check if times fall within specific periods.

시간을 비교하고 시간이 특정 기간 내에 있는지 확인합니다.

#### IsBefore(t1, t2 time.Time) bool

Checks if t1 is before t2.

t1이 t2보다 이전인지 확인합니다.

```go
past := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
future := time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC)

if timeutil.IsBefore(past, future) {
    fmt.Println("past is before future")
}
```

#### IsAfter(t1, t2 time.Time) bool

Checks if t1 is after t2.

t1이 t2보다 이후인지 확인합니다.

```go
if timeutil.IsAfter(future, past) {
    fmt.Println("future is after past")
}
```

#### IsBetween(t, start, end time.Time) bool

Checks if t is between start and end (inclusive).

t가 start와 end 사이에 있는지 확인합니다(포함).

```go
now := time.Now()
start := timeutil.StartOfDay(now)
end := timeutil.EndOfDay(now)

if timeutil.IsBetween(now, start, end) {
    fmt.Println("now is within today")
}
```

#### IsToday(t time.Time) bool

Checks if time is today.

시간이 오늘인지 확인합니다.

```go
if timeutil.IsToday(time.Now()) {
    fmt.Println("This is today")
}
```

#### IsYesterday(t time.Time) bool

Checks if time is yesterday.

시간이 어제인지 확인합니다.

```go
yesterday := timeutil.AddDays(time.Now(), -1)
if timeutil.IsYesterday(yesterday) {
    fmt.Println("This is yesterday")
}
```

#### IsTomorrow(t time.Time) bool

Checks if time is tomorrow.

시간이 내일인지 확인합니다.

```go
tomorrow := timeutil.AddDays(time.Now(), 1)
if timeutil.IsTomorrow(tomorrow) {
    fmt.Println("This is tomorrow")
}
```

#### IsWeekend(t time.Time) bool

Checks if time is on a weekend (Saturday or Sunday).

시간이 주말(토요일 또는 일요일)인지 확인합니다.

```go
if timeutil.IsWeekend(time.Now()) {
    fmt.Println("It's the weekend!")
}
```

#### IsWeekday(t time.Time) bool

Checks if time is on a weekday (Monday to Friday).

시간이 평일(월요일부터 금요일)인지 확인합니다.

```go
if timeutil.IsWeekday(time.Now()) {
    fmt.Println("It's a weekday")
}
```

#### IsThisWeek(t time.Time) bool

Checks if time is in the current week.

시간이 현재 주에 있는지 확인합니다.

```go
if timeutil.IsThisWeek(time.Now()) {
    fmt.Println("This week")
}
```

#### IsThisMonth(t time.Time) bool

Checks if time is in the current month.

시간이 현재 월에 있는지 확인합니다.

```go
if timeutil.IsThisMonth(time.Now()) {
    fmt.Println("This month")
}
```

#### IsThisYear(t time.Time) bool

Checks if time is in the current year.

시간이 현재 년에 있는지 확인합니다.

```go
if timeutil.IsThisYear(time.Now()) {
    fmt.Println("This year")
}
```

#### IsSameDay(t1, t2 time.Time) bool

Checks if two times are on the same day.

두 시간이 같은 날인지 확인합니다.

```go
now1 := time.Now()
now2 := time.Now().Add(1 * time.Hour)

if timeutil.IsSameDay(now1, now2) {
    fmt.Println("Same day")
}
```

#### IsSameWeek(t1, t2 time.Time) bool

Checks if two times are in the same week.

두 시간이 같은 주에 있는지 확인합니다.

```go
if timeutil.IsSameWeek(now1, now2) {
    fmt.Println("Same week")
}
```

#### IsSameMonth(t1, t2 time.Time) bool

Checks if two times are in the same month.

두 시간이 같은 월에 있는지 확인합니다.

```go
if timeutil.IsSameMonth(now1, now2) {
    fmt.Println("Same month")
}
```

#### IsSameYear(t1, t2 time.Time) bool

Checks if two times are in the same year.

두 시간이 같은 년에 있는지 확인합니다.

```go
if timeutil.IsSameYear(now1, now2) {
    fmt.Println("Same year")
}
```

#### IsLeapYear(t time.Time) bool

Checks if the year is a leap year.

년이 윤년인지 확인합니다.

```go
t := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
if timeutil.IsLeapYear(t) {
    fmt.Println("2024 is a leap year")
}
```

#### IsPast(t time.Time) bool

Checks if time is in the past.

시간이 과거인지 확인합니다.

```go
past := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
if timeutil.IsPast(past) {
    fmt.Println("This is in the past")
}
```

#### IsFuture(t time.Time) bool

Checks if time is in the future.

시간이 미래인지 확인합니다.

```go
future := time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC)
if timeutil.IsFuture(future) {
    fmt.Println("This is in the future")
}
```

---

### Age Calculations / 나이 계산

Calculate ages with detailed breakdown.

상세 분석과 함께 나이를 계산합니다.

#### AgeInYears(birthDate time.Time) int

Returns age in years.

년 단위로 나이를 반환합니다.

```go
birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
years := timeutil.AgeInYears(birthDate)
fmt.Println(years) // 35
```

#### AgeInMonths(birthDate time.Time) int

Returns age in months.

월 단위로 나이를 반환합니다.

```go
months := timeutil.AgeInMonths(birthDate)
fmt.Println(months) // 424 (35 years * 12 + 4 months)
```

#### AgeInDays(birthDate time.Time) int

Returns age in days.

일 단위로 나이를 반환합니다.

```go
days := timeutil.AgeInDays(birthDate)
fmt.Println(days) // 12906
```

#### Age(birthDate time.Time) *AgeDetail

Returns detailed age breakdown.

상세한 나이 분석을 반환합니다.

```go
birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
age := timeutil.Age(birthDate)

fmt.Println(age.Years)   // 35
fmt.Println(age.Months)  // 4
fmt.Println(age.Days)    // 29
fmt.Println(age.String()) // "35 years 4 months 29 days"
```

---

### Relative Time / 상대 시간

Get human-readable relative time strings like "2 hours ago" or "in 3 days".

"2 hours ago" 또는 "in 3 days" 같은 사람이 읽기 쉬운 상대 시간 문자열을 가져옵니다.

#### RelativeTime(t time.Time) string

Returns relative time string (long format).

상대 시간 문자열을 반환합니다(긴 포맷).

```go
// Past times / 과거 시간
past1 := time.Now().Add(-30 * time.Second)
fmt.Println(timeutil.RelativeTime(past1)) // "30 seconds ago"

past2 := time.Now().Add(-2 * time.Hour)
fmt.Println(timeutil.RelativeTime(past2)) // "2 hours ago"

past3 := time.Now().Add(-3 * 24 * time.Hour)
fmt.Println(timeutil.RelativeTime(past3)) // "3 days ago"

// Future times / 미래 시간
future1 := time.Now().Add(5 * time.Minute)
fmt.Println(timeutil.RelativeTime(future1)) // "in 5 minutes"

future2 := time.Now().Add(2 * 24 * time.Hour)
fmt.Println(timeutil.RelativeTime(future2)) // "in 2 days"
```

#### RelativeTimeShort(t time.Time) string

Returns relative time string (short format).

상대 시간 문자열을 반환합니다(짧은 포맷).

```go
past := time.Now().Add(-2 * time.Hour)
fmt.Println(timeutil.RelativeTimeShort(past)) // "2h ago"

future := time.Now().Add(3 * 24 * time.Hour)
fmt.Println(timeutil.RelativeTimeShort(future)) // "in 3d"
```

#### TimeAgo(t time.Time) string

Alias for RelativeTime (commonly used name).

RelativeTime의 별칭(일반적으로 사용되는 이름).

```go
past := time.Now().Add(-1 * time.Hour)
fmt.Println(timeutil.TimeAgo(past)) // "1 hour ago"
```

#### HumanizeDuration(d time.Duration) string

Converts duration to human-readable string.

기간을 사람이 읽기 쉬운 문자열로 변환합니다.

```go
duration := 2*time.Hour + 30*time.Minute + 45*time.Second
fmt.Println(timeutil.HumanizeDuration(duration)) // "2 hours 30 minutes 45 seconds"
```

---

### Unix Timestamp / Unix 타임스탬프

Work with Unix timestamps (seconds, milliseconds, microseconds, nanoseconds).

Unix 타임스탬프(초, 밀리초, 마이크로초, 나노초)로 작업합니다.

#### Now() int64

Returns current Unix timestamp in seconds.

초 단위의 현재 Unix 타임스탬프를 반환합니다.

```go
timestamp := timeutil.Now()
fmt.Println(timestamp) // 1729052645
```

#### NowMilli() int64

Returns current Unix timestamp in milliseconds.

밀리초 단위의 현재 Unix 타임스탬프를 반환합니다.

```go
millis := timeutil.NowMilli()
fmt.Println(millis) // 1729052645123
```

#### NowMicro() int64

Returns current Unix timestamp in microseconds.

마이크로초 단위의 현재 Unix 타임스탬프를 반환합니다.

```go
micros := timeutil.NowMicro()
fmt.Println(micros) // 1729052645123456
```

#### NowNano() int64

Returns current Unix timestamp in nanoseconds.

나노초 단위의 현재 Unix 타임스탬프를 반환합니다.

```go
nanos := timeutil.NowNano()
fmt.Println(nanos) // 1729052645123456789
```

#### FromUnix(sec int64) time.Time

Converts Unix timestamp (seconds) to time.Time.

Unix 타임스탬프(초)를 time.Time으로 변환합니다.

```go
t := timeutil.FromUnix(1729052645)
fmt.Println(t) // 2025-10-14 15:04:05 +0900 KST
```

#### FromUnixMilli(msec int64) time.Time

Converts Unix timestamp (milliseconds) to time.Time.

Unix 타임스탬프(밀리초)를 time.Time으로 변환합니다.

```go
t := timeutil.FromUnixMilli(1729052645123)
fmt.Println(t) // 2025-10-14 15:04:05.123 +0900 KST
```

#### FromUnixMicro(usec int64) time.Time

Converts Unix timestamp (microseconds) to time.Time.

Unix 타임스탬프(마이크로초)를 time.Time으로 변환합니다.

```go
t := timeutil.FromUnixMicro(1729052645123456)
fmt.Println(t) // 2025-10-14 15:04:05.123456 +0900 KST
```

#### FromUnixNano(nsec int64) time.Time

Converts Unix timestamp (nanoseconds) to time.Time.

Unix 타임스탬프(나노초)를 time.Time으로 변환합니다.

```go
t := timeutil.FromUnixNano(1729052645123456789)
fmt.Println(t) // 2025-10-14 15:04:05.123456789 +0900 KST
```

#### ToUnix(t time.Time) int64

Converts time.Time to Unix timestamp (seconds).

time.Time을 Unix 타임스탬프(초)로 변환합니다.

```go
timestamp := timeutil.ToUnix(time.Now())
fmt.Println(timestamp) // 1729052645
```

#### ToUnixMilli(t time.Time) int64

Converts time.Time to Unix timestamp (milliseconds).

time.Time을 Unix 타임스탬프(밀리초)로 변환합니다.

```go
millis := timeutil.ToUnixMilli(time.Now())
fmt.Println(millis) // 1729052645123
```

#### ToUnixMicro(t time.Time) int64

Converts time.Time to Unix timestamp (microseconds).

time.Time을 Unix 타임스탬프(마이크로초)로 변환합니다.

```go
micros := timeutil.ToUnixMicro(time.Now())
fmt.Println(micros) // 1729052645123456
```

#### ToUnixNano(t time.Time) int64

Converts time.Time to Unix timestamp (nanoseconds).

time.Time을 Unix 타임스탬프(나노초)로 변환합니다.

```go
nanos := timeutil.ToUnixNano(time.Now())
fmt.Println(nanos) // 1729052645123456789
```

---

### Business Days / 영업일

Calculate business days, skipping weekends and holidays.

주말 및 공휴일을 건너뛰고 영업일을 계산합니다.

#### IsBusinessDay(t time.Time) bool

Checks if time is a business day (Monday-Friday, not a holiday).

시간이 영업일(월-금, 공휴일 아님)인지 확인합니다.

```go
if timeutil.IsBusinessDay(time.Now()) {
    fmt.Println("Open for business")
}
```

#### IsHoliday(t time.Time) bool

Checks if time is a registered holiday.

시간이 등록된 공휴일인지 확인합니다.

```go
newYear := time.Date(2025, 1, 1, 0, 0, 0, 0, timeutil.KST)
if timeutil.IsHoliday(newYear) {
    fmt.Println("It's a holiday")
}
```

#### AddBusinessDays(t time.Time, days int) time.Time

Adds business days, skipping weekends and holidays.

주말 및 공휴일을 건너뛰고 영업일을 더합니다.

```go
today := time.Now()
nextBizDay := timeutil.AddBusinessDays(today, 5)
fmt.Println(nextBizDay) // 5 business days later
```

#### NextBusinessDay(t time.Time) time.Time

Returns the next business day.

다음 영업일을 반환합니다.

```go
next := timeutil.NextBusinessDay(time.Now())
fmt.Println(next)
```

#### PreviousBusinessDay(t time.Time) time.Time

Returns the previous business day.

이전 영업일을 반환합니다.

```go
prev := timeutil.PreviousBusinessDay(time.Now())
fmt.Println(prev)
```

#### CountBusinessDays(start, end time.Time) int

Counts business days between two times.

두 시간 사이의 영업일을 계산합니다.

```go
start := time.Date(2025, 10, 13, 0, 0, 0, 0, timeutil.KST) // Monday
end := time.Date(2025, 10, 17, 0, 0, 0, 0, timeutil.KST)   // Friday
count := timeutil.CountBusinessDays(start, end)
fmt.Println(count) // 5 business days
```

#### SetHolidays(holidays []time.Time)

Sets custom holidays for business day calculations.

영업일 계산을 위한 커스텀 공휴일을 설정합니다.

```go
holidays := []time.Time{
    time.Date(2025, 12, 25, 0, 0, 0, 0, timeutil.KST), // Christmas
    time.Date(2025, 1, 1, 0, 0, 0, 0, timeutil.KST),   // New Year
}
timeutil.SetHolidays(holidays)
```

#### GetHolidays() []time.Time

Returns the list of registered holidays.

등록된 공휴일 목록을 반환합니다.

```go
holidays := timeutil.GetHolidays()
for _, h := range holidays {
    fmt.Println(h)
}
```

#### ClearHolidays()

Clears all registered holidays.

모든 등록된 공휴일을 지웁니다.

```go
timeutil.ClearHolidays()
```

#### AddKoreanHolidays(year int)

Adds Korean public holidays for the specified year.

지정된 년도의 한국 공휴일을 추가합니다.

```go
// Add holidays for 2025 / 2025년 공휴일 추가
timeutil.AddKoreanHolidays(2025)

// Korean holidays include:
// 한국 공휴일 포함:
// - 1월 1일: 신정 (New Year's Day)
// - 2월 9-11일: 설날 (Lunar New Year)
// - 3월 1일: 삼일절 (Independence Movement Day)
// - 5월 5일: 어린이날 (Children's Day)
// - 5월 15일: 부처님오신날 (Buddha's Birthday)
// - 6월 6일: 현충일 (Memorial Day)
// - 8월 15일: 광복절 (Liberation Day)
// - 9월 16-18일: 추석 (Harvest Moon Festival)
// - 10월 3일: 개천절 (National Foundation Day)
// - 10월 9일: 한글날 (Hangeul Day)
// - 12월 25일: 성탄절 (Christmas)
```

---

## Common Use Cases / 일반적인 사용 사례

### Use Case 1: Calculate Session Duration / 사용 사례 1: 세션 기간 계산

```go
package main

import (
    "fmt"
    "time"
    "github.com/arkd0ng/go-utils/timeutil"
)

func main() {
    sessionStart := time.Now()

    // ... user session ...
    time.Sleep(2 * time.Hour) // Simulate session

    sessionEnd := time.Now()
    duration := timeutil.SubTime(sessionStart, sessionEnd)

    fmt.Printf("Session duration: %s\n", duration.String())
    // Output: "Session duration: 2 hours"
}
```

### Use Case 2: Format Dates for API Response / 사용 사례 2: API 응답을 위한 날짜 포맷

```go
package main

import (
    "encoding/json"
    "time"
    "github.com/arkd0ng/go-utils/timeutil"
)

type User struct {
    ID        int    `json:"id"`
    Name      string `json:"name"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}

func main() {
    now := time.Now()

    user := User{
        ID:        1,
        Name:      "John Doe",
        CreatedAt: timeutil.FormatISO8601(now),
        UpdatedAt: timeutil.FormatRFC3339(now),
    }

    jsonData, _ := json.Marshal(user)
    fmt.Println(string(jsonData))
    // Output: {"id":1,"name":"John Doe","created_at":"2025-10-14","updated_at":"2025-10-14T15:04:05+09:00"}
}
```

### Use Case 3: Schedule Next Business Day / 사용 사례 3: 다음 영업일 예약

```go
package main

import (
    "fmt"
    "time"
    "github.com/arkd0ng/go-utils/timeutil"
)

func scheduleDelivery(orderTime time.Time) time.Time {
    // Add Korean holidays / 한국 공휴일 추가
    timeutil.AddKoreanHolidays(orderTime.Year())

    // Schedule delivery 3 business days later
    // 3 영업일 후에 배송 예약
    deliveryDate := timeutil.AddBusinessDays(orderTime, 3)

    return deliveryDate
}

func main() {
    orderTime := time.Now()
    delivery := scheduleDelivery(orderTime)

    fmt.Printf("Order time: %s\n", timeutil.FormatDateTime(orderTime))
    fmt.Printf("Delivery date: %s\n", timeutil.FormatDateTime(delivery))
}
```

### Use Case 4: Display User Age / 사용 사례 4: 사용자 나이 표시

```go
package main

import (
    "fmt"
    "time"
    "github.com/arkd0ng/go-utils/timeutil"
)

type Profile struct {
    Name      string
    BirthDate time.Time
}

func (p *Profile) DisplayAge() string {
    age := timeutil.Age(p.BirthDate)
    return age.String()
}

func main() {
    profile := Profile{
        Name:      "Jane Doe",
        BirthDate: time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC),
    }

    fmt.Printf("%s is %s old\n", profile.Name, profile.DisplayAge())
    // Output: "Jane Doe is 35 years 4 months 29 days old"
}
```

### Use Case 5: Parse and Convert Timezone / 사용 사례 5: 타임존 파싱 및 변환

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/timeutil"
)

func main() {
    // Parse UTC time string / UTC 시간 문자열 파싱
    utcStr := "2025-10-14T06:00:00Z"
    t, err := timeutil.ParseRFC3339(utcStr)
    if err != nil {
        panic(err)
    }

    // Convert to KST / KST로 변환
    kstTime := timeutil.ToKST(t)

    // Format in Korean / 한국어로 포맷
    formatted := timeutil.FormatKorean(kstTime)

    fmt.Printf("UTC: %s\n", utcStr)
    fmt.Printf("KST: %s\n", formatted)
    // Output: "KST: 2025년 10월 14일 15시 00분 00초"
}
```

### Use Case 6: Generate Activity Timestamps / 사용 사례 6: 활동 타임스탬프 생성

```go
package main

import (
    "fmt"
    "time"
    "github.com/arkd0ng/go-utils/timeutil"
)

type Activity struct {
    UserID    int
    Action    string
    Timestamp int64  // Unix milliseconds
}

func logActivity(userID int, action string) Activity {
    return Activity{
        UserID:    userID,
        Action:    action,
        Timestamp: timeutil.NowMilli(),
    }
}

func displayActivity(activity Activity) {
    t := timeutil.FromUnixMilli(activity.Timestamp)
    relative := timeutil.RelativeTime(t)

    fmt.Printf("User %d %s %s\n", activity.UserID, activity.Action, relative)
}

func main() {
    // Log activity / 활동 로그
    activity := logActivity(123, "logged in")

    // Simulate time passing / 시간 경과 시뮬레이션
    time.Sleep(2 * time.Second)

    // Display activity / 활동 표시
    displayActivity(activity)
    // Output: "User 123 logged in 2 seconds ago"
}
```

### Use Case 7: Report Generation Period / 사용 사례 7: 보고서 생성 기간

```go
package main

import (
    "fmt"
    "time"
    "github.com/arkd0ng/go-utils/timeutil"
)

type Report struct {
    StartDate time.Time
    EndDate   time.Time
    Data      map[string]interface{}
}

func generateMonthlyReport() Report {
    now := time.Now()

    // Get start and end of current month
    // 현재 월의 시작과 끝 가져오기
    start := timeutil.StartOfMonth(now)
    end := timeutil.EndOfMonth(now)

    return Report{
        StartDate: start,
        EndDate:   end,
        Data:      make(map[string]interface{}),
    }
}

func (r *Report) Display() {
    fmt.Printf("Monthly Report\n")
    fmt.Printf("Period: %s to %s\n",
        timeutil.FormatDate(r.StartDate),
        timeutil.FormatDate(r.EndDate))

    days := timeutil.DiffInDays(r.StartDate, r.EndDate)
    fmt.Printf("Total days: %.0f\n", days)
}

func main() {
    report := generateMonthlyReport()
    report.Display()
    // Output:
    // Monthly Report
    // Period: 2025-10-01 to 2025-10-31
    // Total days: 31
}
```

### Use Case 8: Deadline Checker / 사용 사례 8: 마감일 확인

```go
package main

import (
    "fmt"
    "time"
    "github.com/arkd0ng/go-utils/timeutil"
)

type Task struct {
    Name     string
    Deadline time.Time
}

func (t *Task) CheckStatus() string {
    if timeutil.IsPast(t.Deadline) {
        duration := timeutil.SubTime(t.Deadline, time.Now())
        return fmt.Sprintf("Overdue by %s", duration.Humanize())
    }

    if timeutil.IsToday(t.Deadline) {
        return "Due today!"
    }

    relative := timeutil.RelativeTime(t.Deadline)
    return fmt.Sprintf("Due %s", relative)
}

func main() {
    tasks := []Task{
        {Name: "Submit report", Deadline: timeutil.AddDays(time.Now(), -2)},
        {Name: "Review code", Deadline: time.Now()},
        {Name: "Team meeting", Deadline: timeutil.AddDays(time.Now(), 3)},
    }

    for _, task := range tasks {
        fmt.Printf("%s: %s\n", task.Name, task.CheckStatus())
    }
    // Output:
    // Submit report: Overdue by 2d
    // Review code: Due today!
    // Team meeting: Due in 3 days
}
```

---

## Best Practices / 모범 사례

### 1. Use KST as Default for Korean Applications / 한국 애플리케이션에는 KST를 기본값으로 사용

```go
// The package defaults to KST, no configuration needed
// 패키지는 기본적으로 KST를 사용하므로 설정이 필요 없습니다
now := timeutil.NowKST()
formatted := timeutil.Format(now, "YYYY-MM-DD")
```

### 2. Use Custom Format Tokens for Readability / 가독성을 위해 커스텀 포맷 토큰 사용

```go
// ✅ Good: Easy to read and remember
// ✅ 좋음: 읽기 쉽고 기억하기 쉬움
formatted := timeutil.Format(time.Now(), "YYYY-MM-DD HH:mm:ss")

// ❌ Bad: Hard to remember Go's reference date
// ❌ 나쁨: Go의 참조 날짜를 기억하기 어려움
formatted := time.Now().Format("2006-01-02 15:04:05")
```

### 3. Use Parse() for Auto-Detection / 자동 감지를 위해 Parse() 사용

```go
// ✅ Good: Automatically detects format
// ✅ 좋음: 자동으로 포맷 감지
t, err := timeutil.Parse("2025-10-14 15:04:05")

// ❌ Less flexible: Must know exact format
// ❌ 덜 유연함: 정확한 포맷을 알아야 함
t, err := time.Parse("2006-01-02 15:04:05", "2025-10-14 15:04:05")
```

### 4. Register Holidays Once at Startup / 시작 시 한 번만 공휴일 등록

```go
func init() {
    // Register holidays once when application starts
    // 애플리케이션 시작 시 한 번만 공휴일 등록
    currentYear := time.Now().Year()
    timeutil.AddKoreanHolidays(currentYear)
    timeutil.AddKoreanHolidays(currentYear + 1)
}
```

### 5. Use TimeDiff for Human-Readable Output / 사람이 읽기 쉬운 출력을 위해 TimeDiff 사용

```go
// ✅ Good: Human-readable
// ✅ 좋음: 사람이 읽기 쉬움
diff := timeutil.SubTime(start, end)
fmt.Println(diff.String()) // "2 days 6 hours 30 minutes"

// ❌ Less readable: Raw duration
// ❌ 덜 읽기 쉬움: 원시 기간
duration := end.Sub(start)
fmt.Println(duration) // "54h30m0s"
```

### 6. Use Business Days for Scheduling / 일정 예약을 위해 영업일 사용

```go
// Always register holidays first / 항상 먼저 공휴일 등록
timeutil.AddKoreanHolidays(2025)

// Then calculate business days / 그런 다음 영업일 계산
deliveryDate := timeutil.AddBusinessDays(orderDate, 3)
```

### 7. Handle Timezone Conversions Explicitly / 타임존 변환을 명시적으로 처리

```go
// ✅ Good: Explicit timezone conversion
// ✅ 좋음: 명시적 타임존 변환
kstTime := timeutil.ToKST(time.Now().UTC())

// ❌ Implicit: May cause confusion
// ❌ 암시적: 혼란을 야기할 수 있음
localTime := time.Now()
```

### 8. Use StartOf/EndOf for Period Boundaries / 기간 경계를 위해 StartOf/EndOf 사용

```go
// ✅ Good: Clear period boundaries
// ✅ 좋음: 명확한 기간 경계
start := timeutil.StartOfDay(time.Now())
end := timeutil.EndOfDay(time.Now())

// ❌ Manual: Error-prone
// ❌ 수동: 오류 발생 가능
start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
```

### 9. Use Relative Time for User-Facing Displays / 사용자 대면 디스플레이에 상대 시간 사용

```go
// ✅ Good: User-friendly
// ✅ 좋음: 사용자 친화적
lastSeen := timeutil.RelativeTime(user.LastLoginAt)
fmt.Printf("Last seen %s", lastSeen) // "Last seen 2 hours ago"

// ❌ Less friendly: Absolute time
// ❌ 덜 친화적: 절대 시간
fmt.Printf("Last seen %s", user.LastLoginAt)
```

### 10. Use Unix Timestamps for Storage / 저장을 위해 Unix 타임스탬프 사용

```go
// ✅ Good: Store as Unix timestamp (timezone-independent)
// ✅ 좋음: Unix 타임스탬프로 저장(타임존 독립적)
timestamp := timeutil.ToUnixMilli(time.Now())

// Later: Convert back to time.Time
// 나중에: time.Time으로 다시 변환
t := timeutil.FromUnixMilli(timestamp)
```

### 11. Check Errors from Timezone Conversions / 타임존 변환에서 에러 확인

```go
// ✅ Good: Handle errors
// ✅ 좋음: 에러 처리
t, err := timeutil.ConvertTimezone(time.Now(), "America/New_York")
if err != nil {
    log.Printf("Invalid timezone: %v", err)
    return
}

// ❌ Bad: Ignore errors
// ❌ 나쁨: 에러 무시
t, _ := timeutil.ConvertTimezone(time.Now(), "Invalid/Timezone")
```

### 12. Use Comparison Functions for Clarity / 명확성을 위해 비교 함수 사용

```go
// ✅ Good: Clear intent
// ✅ 좋음: 명확한 의도
if timeutil.IsToday(t) {
    fmt.Println("Today!")
}

// ❌ Verbose: Manual comparison
// ❌ 장황함: 수동 비교
if t.Year() == time.Now().Year() && t.Month() == time.Now().Month() && t.Day() == time.Now().Day() {
    fmt.Println("Today!")
}
```

---

## Troubleshooting / 문제 해결

### Problem: Incorrect Timezone / 문제: 잘못된 타임존

**Symptom / 증상**: Times are showing in wrong timezone.

**Solution / 해결책**:

```go
// Check default timezone / 기본 타임존 확인
fmt.Println(timeutil.GetDefaultTimezone())

// Change if needed / 필요시 변경
timeutil.SetDefaultTimezone("Asia/Seoul")
```

### Problem: Invalid Timezone Name / 문제: 잘못된 타임존 이름

**Symptom / 증상**: Error "unknown time zone" when converting timezones.

**Solution / 해결책**:

```go
// Check if timezone is valid / 타임존이 유효한지 확인
if !timeutil.IsValidTimezone("Invalid/Timezone") {
    fmt.Println("Invalid timezone")
}

// List available timezones / 사용 가능한 타임존 나열
timezones := timeutil.ListTimezones()
for _, tz := range timezones {
    fmt.Println(tz)
}
```

### Problem: Business Days Not Skipping Holidays / 문제: 영업일이 공휴일을 건너뛰지 않음

**Symptom / 증상**: AddBusinessDays doesn't skip holidays.

**Solution / 해결책**:

```go
// Make sure holidays are registered / 공휴일이 등록되었는지 확인
timeutil.AddKoreanHolidays(2025)

// Or set custom holidays / 또는 커스텀 공휴일 설정
holidays := []time.Time{
    time.Date(2025, 12, 25, 0, 0, 0, 0, timeutil.KST),
}
timeutil.SetHolidays(holidays)

// Check if holidays are registered / 공휴일이 등록되었는지 확인
fmt.Println(len(timeutil.GetHolidays()))
```

### Problem: Parse Fails with Custom Format / 문제: 커스텀 포맷으로 파싱 실패

**Symptom / 증상**: Parse() returns error with custom date format.

**Solution / 해결책**:

```go
// Use Parse() for common formats / 일반 포맷에 Parse() 사용
t1, _ := timeutil.Parse("2025-10-14")
t2, _ := timeutil.Parse("2025-10-14 15:04:05")

// For custom formats, use specific parsers / 커스텀 포맷의 경우 특정 파서 사용
t3, _ := timeutil.ParseISO8601("2025-10-14")
t4, _ := timeutil.ParseRFC3339("2025-10-14T15:04:05+09:00")
```

### Problem: Age Calculation Seems Wrong / 문제: 나이 계산이 잘못된 것 같음

**Symptom / 증상**: Age() returns unexpected values.

**Solution / 해결책**:

```go
// Make sure birthdate is in the past / 생년월일이 과거인지 확인
birthDate := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
age := timeutil.Age(birthDate)

// Check each component / 각 구성 요소 확인
fmt.Printf("Years: %d, Months: %d, Days: %d\n", age.Years, age.Months, age.Days)

// Use AgeInYears for simple year count / 간단한 년 수를 위해 AgeInYears 사용
years := timeutil.AgeInYears(birthDate)
```

### Problem: Format Tokens Not Working / 문제: 포맷 토큰이 작동하지 않음

**Symptom / 증상**: Custom format tokens (YYYY, MM, DD) not being replaced.

**Solution / 해결책**:

```go
// ✅ Use Format() with custom tokens
// ✅ 커스텀 토큰으로 Format() 사용
formatted := timeutil.Format(time.Now(), "YYYY-MM-DD HH:mm:ss")

// ❌ Don't use standard time.Format() with custom tokens
// ❌ 커스텀 토큰으로 표준 time.Format() 사용 안함
// formatted := time.Now().Format("YYYY-MM-DD") // Won't work!
```

---

## FAQ / 자주 묻는 질문

### Q1: Why is KST the default timezone? / 왜 KST가 기본 타임존인가요?

**A**: This package was designed primarily for Korean users and applications. However, you can easily change the default timezone using `SetDefaultTimezone()`.

**답변**: 이 패키지는 주로 한국 사용자 및 애플리케이션을 위해 설계되었습니다. 그러나 `SetDefaultTimezone()`을 사용하여 기본 타임존을 쉽게 변경할 수 있습니다.

```go
timeutil.SetDefaultTimezone("UTC")
timeutil.SetDefaultTimezone("America/New_York")
```

### Q2: Can I use this package with other timezone libraries? / 다른 타임존 라이브러리와 함께 사용할 수 있나요?

**A**: Yes! This package uses Go's standard `time.Time` type, so it's compatible with any other time library.

**답변**: 네! 이 패키지는 Go의 표준 `time.Time` 타입을 사용하므로 다른 시간 라이브러리와 호환됩니다.

### Q3: How do I add my own custom holidays? / 나만의 커스텀 공휴일을 어떻게 추가하나요?

**A**: Use `SetHolidays()` to set your custom holiday list.

**답변**: `SetHolidays()`를 사용하여 커스텀 공휴일 목록을 설정하세요.

```go
holidays := []time.Time{
    time.Date(2025, 12, 25, 0, 0, 0, 0, timeutil.KST),
    time.Date(2025, 1, 1, 0, 0, 0, 0, timeutil.KST),
}
timeutil.SetHolidays(holidays)
```

### Q4: Is this package thread-safe? / 이 패키지는 스레드 안전한가요?

**A**: Yes! The package uses `sync.RWMutex` for timezone caching and holiday management, making it safe for concurrent use.

**답변**: 네! 패키지는 타임존 캐싱 및 공휴일 관리를 위해 `sync.RWMutex`를 사용하므로 동시 사용에 안전합니다.

### Q5: Can I use custom format tokens with time.Format()? / time.Format()과 함께 커스텀 포맷 토큰을 사용할 수 있나요?

**A**: No. Custom tokens (YYYY, MM, DD) only work with `timeutil.Format()`. Use Go's standard reference date (2006-01-02) with `time.Format()`.

**답변**: 아니오. 커스텀 토큰(YYYY, MM, DD)은 `timeutil.Format()`에서만 작동합니다. `time.Format()`과 함께 Go의 표준 참조 날짜(2006-01-02)를 사용하세요.

### Q6: How accurate are the business day calculations? / 영업일 계산은 얼마나 정확한가요?

**A**: Business day calculations skip weekends (Saturday and Sunday) and any holidays you've registered. Make sure to register all holidays for accurate results.

**답변**: 영업일 계산은 주말(토요일 및 일요일) 및 등록한 공휴일을 건너뜁니다. 정확한 결과를 위해 모든 공휴일을 등록하세요.

### Q7: Does Parse() work with all date formats? / Parse()는 모든 날짜 포맷에서 작동하나요?

**A**: Parse() supports common formats like ISO 8601, RFC 3339, date-only, and datetime. For custom formats, use the specific parser or `time.Parse()`.

**답변**: Parse()는 ISO 8601, RFC 3339, 날짜만, 날짜시간과 같은 일반 포맷을 지원합니다. 커스텀 포맷의 경우 특정 파서 또는 `time.Parse()`를 사용하세요.

### Q8: Can I change the week start day? / 주 시작일을 변경할 수 있나요?

**A**: Currently, the package uses Monday as the week start day (ISO 8601 standard). This is not configurable in the current version.

**답변**: 현재 패키지는 월요일을 주 시작일로 사용합니다(ISO 8601 표준). 현재 버전에서는 구성할 수 없습니다.

### Q9: What's the difference between RelativeTime() and TimeAgo()? / RelativeTime()과 TimeAgo()의 차이점은 무엇인가요?

**A**: They are the same! `TimeAgo()` is just an alias for `RelativeTime()` for familiarity with other libraries.

**답변**: 같습니다! `TimeAgo()`는 다른 라이브러리와의 친숙성을 위해 `RelativeTime()`의 별칭일 뿐입니다.

### Q10: How do I handle leap years? / 윤년을 어떻게 처리하나요?

**A**: The package automatically handles leap years in all date calculations. You can also explicitly check with `IsLeapYear()`.

**답변**: 패키지는 모든 날짜 계산에서 자동으로 윤년을 처리합니다. `IsLeapYear()`로 명시적으로 확인할 수도 있습니다.

```go
t := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
if timeutil.IsLeapYear(t) {
    fmt.Println("2024 is a leap year")
}
```

---

## Conclusion / 결론

The `timeutil` package provides a comprehensive and intuitive API for time and date operations in Go. By reducing boilerplate code and providing human-readable output, it makes working with time much easier.

`timeutil` 패키지는 Go에서 시간 및 날짜 작업을 위한 포괄적이고 직관적인 API를 제공합니다. 보일러플레이트 코드를 줄이고 사람이 읽기 쉬운 출력을 제공함으로써 시간 작업을 훨씬 쉽게 만듭니다.

For more information, see:
- [Package README](../../timeutil/README.md)
- [Developer Guide](./DEVELOPER_GUIDE.md)
- [GitHub Repository](https://github.com/arkd0ng/go-utils)

자세한 정보는 다음을 참조하세요:
- [패키지 README](../../timeutil/README.md)
- [개발자 가이드](./DEVELOPER_GUIDE.md)
- [GitHub 저장소](https://github.com/arkd0ng/go-utils)

---

**Version / 버전**: v1.6.006
**Last Updated / 마지막 업데이트**: 2025-10-14


---

## String Parameters / 문자열 매개변수

**NEW in v1.6.008!** / **v1.6.008 신규!**

For comprehensive documentation on string parameter support, see:
- [String Parameters Guide](./STRING_PARAMETERS.md) - Complete guide with 40+ format examples

문자열 매개변수 지원에 대한 포괄적인 문서는 다음을 참조하세요:
- [문자열 매개변수 가이드](./STRING_PARAMETERS.md) - 40개 이상의 포맷 예제가 포함된 완전한 가이드

### Quick Overview / 빠른 개요

```go
// OLD WAY - Too much boilerplate
layout := "2006-01-02 15:04:05.000"
t1, err := time.ParseInLocation(layout, "2024-10-04 08:34:42.324", timeutil.KST)
t2, err := time.ParseInLocation(layout, "2024-10-14 14:56:23.789", timeutil.KST)
diff := timeutil.SubTime(t1, t2)

// NEW WAY - So much simpler!
diff, err := timeutil.SubTimeString("2024-10-04 08:34:42.324", "2024-10-14 14:56:23.789")
```

### ParseAny - The Magic Function / 마법의 함수

Automatically detects and parses 40+ formats including:
- Database formats (MySQL, PostgreSQL, SQLite)
- ISO8601, RFC3339
- Date formats (YYYY-MM-DD, MM/DD/YYYY, etc.)
- Month names (Oct 04, 2024, October 04, 2024)
- **Korean formats** (2024년 10월 04일, 오전 9시, 오후 3시)

40개 이상의 포맷을 자동으로 감지하고 파싱합니다:
- 데이터베이스 포맷 (MySQL, PostgreSQL, SQLite)
- ISO8601, RFC3339
- 날짜 포맷 (YYYY-MM-DD, MM/DD/YYYY 등)
- 월 이름 (Oct 04, 2024, October 04, 2024)
- **한글 포맷** (2024년 10월 04일, 오전 9시, 오후 3시)

```go
// Works with ANY format!
t1, _ := timeutil.ParseAny("2024-10-04 08:34:42.324")
t2, _ := timeutil.ParseAny("Oct 04, 2024")
t3, _ := timeutil.ParseAny("2024년 10월 04일")
t4, _ := timeutil.ParseAny("2024/10/04")
```

### 50+ String Functions / 50개 이상의 String 함수

All major timeutil functions now accept strings:
- `SubTimeString`, `DiffInDaysString`, `AgeString`
- `AddDaysString`, `SubDaysString`, `FormatString`
- `IsSameDayString`, `IsBeforeString`, `WeekdayKoreanString`
- And 40+ more!

모든 주요 timeutil 함수가 이제 문자열을 받습니다:
- `SubTimeString`, `DiffInDaysString`, `AgeString`
- `AddDaysString`, `SubDaysString`, `FormatString`
- `IsSameDayString`, `IsBeforeString`, `WeekdayKoreanString`
- 그리고 40개 이상!

See [STRING_PARAMETERS.md](./STRING_PARAMETERS.md) for complete documentation.

전체 문서는 [STRING_PARAMETERS.md](./STRING_PARAMETERS.md)를 참조하세요.

---

**Version / 버전**: v1.6.008
**Last Updated / 마지막 업데이트**: 2025-10-14

