# String Parameters - Comprehensive Guide / 문자열 매개변수 - 종합 가이드

**NEW in v1.6.008!** / **v1.6.008 신규!**

## Overview / 개요

One of the most common pain points when working with time data is dealing with strings from databases, APIs, and user input. In v1.6.008, we've added comprehensive string parameter support that automatically parses 40+ time formats!

시간 데이터를 다룰 때 가장 흔한 문제 중 하나는 데이터베이스, API, 사용자 입력에서 오는 문자열을 처리하는 것입니다. v1.6.008에서 40개 이상의 시간 포맷을 자동으로 파싱하는 포괄적인 문자열 매개변수 지원을 추가했습니다!

---

## Why String Parameters? / 왜 문자열 매개변수인가?

### The Problem / 문제점

```go
// ❌ OLD WAY - Too much boilerplate!
layout := "2006-01-02 15:04:05.000"
t1, err := time.ParseInLocation(layout, "2024-10-04 08:34:42.324", timeutil.KST)
if err != nil {
    return err
}
t2, err := time.ParseInLocation(layout, "2024-10-14 14:56:23.789", timeutil.KST)
if err != nil {
    return err
}
diff := timeutil.SubTime(t1, t2)
```

### The Solution / 해결책

```go
// ✅ NEW WAY - So much simpler!
diff, err := timeutil.SubTimeString("2024-10-04 08:34:42.324", "2024-10-14 14:56:23.789")
if err != nil {
    return err
}
```

---

## ParseAny - The Magic Function / 마법의 함수

The `ParseAny` function is the heart of string parameter support. It automatically detects and parses 40+ time formats!

`ParseAny` 함수는 문자열 매개변수 지원의 핵심입니다. 40개 이상의 시간 포맷을 자동으로 감지하고 파싱합니다!

### Supported Formats / 지원되는 포맷

#### 1. Database Formats / 데이터베이스 포맷

```go
// MySQL DATETIME
t, _ := timeutil.ParseAny("2024-10-04 08:34:42")

// MySQL DATETIME with milliseconds
t, _ := timeutil.ParseAny("2024-10-04 08:34:42.324")

// PostgreSQL TIMESTAMP
t, _ := timeutil.ParseAny("2024-10-04 08:34:42.324567")

// SQLite DATETIME
t, _ := timeutil.ParseAny("2024-10-04 08:34:42")
```

#### 2. ISO8601 / RFC3339 Formats

```go
// ISO8601 with timezone
t, _ := timeutil.ParseAny("2024-10-04T08:34:42+09:00")

// RFC3339
t, _ := timeutil.ParseAny("2024-10-04T08:34:42Z")

// ISO8601 without timezone
t, _ := timeutil.ParseAny("2024-10-04T15:04:05")
```

#### 3. Common Date Formats / 일반 날짜 포맷

```go
// Standard date
t, _ := timeutil.ParseAny("2024-10-04")

// Date with slashes
t, _ := timeutil.ParseAny("2024/10/04")

// US format
t, _ := timeutil.ParseAny("10/04/2024")

// EU format
t, _ := timeutil.ParseAny("04-10-2024")

// German format
t, _ := timeutil.ParseAny("04.10.2024")
```

#### 4. Month Names / 월 이름

```go
// Full month name
t, _ := timeutil.ParseAny("October 04, 2024")

// Short month name
t, _ := timeutil.ParseAny("Oct 04, 2024")

// Different format
t, _ := timeutil.ParseAny("04-Oct-2024")
t, _ := timeutil.ParseAny("04 October 2024")
```

#### 5. Korean Formats / 한글 포맷 ⭐

```go
// Full Korean datetime
t, _ := timeutil.ParseAny("2024년 10월 04일 15시 30분 45초")

// Without seconds
t, _ := timeutil.ParseAny("2024년 10월 04일 15시 30분")

// Hour only
t, _ := timeutil.ParseAny("2024년 10월 04일 15시")

// Date only
t, _ := timeutil.ParseAny("2024년 10월 04일")

// Without leading zeros
t, _ := timeutil.ParseAny("2024년 1월 4일 9시 5분 3초")

// With AM/PM (오전/오후)
t, _ := timeutil.ParseAny("2024년 10월 04일 오전 9시 30분")
t, _ := timeutil.ParseAny("2024년 10월 04일 오후 3시 30분")
```

#### 6. Standard Go Time Formats

```go
// ANSIC
t, _ := timeutil.ParseAny("Mon Jan  2 15:04:05 2006")

// UnixDate
t, _ := timeutil.ParseAny("Mon Jan  2 15:04:05 MST 2006")

// RFC822
t, _ := timeutil.ParseAny("02 Jan 06 15:04 MST")

// RFC1123
t, _ := timeutil.ParseAny("Mon, 02 Jan 2006 15:04:05 MST")
```

---

## Parse Functions / 파싱 함수

### ParseAny

Automatically detects and parses 40+ formats.

40개 이상의 포맷을 자동으로 감지하고 파싱합니다.

```go
func ParseAny(s string) (time.Time, error)
```

**Example**:
```go
// Works with any format!
t1, _ := timeutil.ParseAny("2024-10-04 08:34:42.324")
t2, _ := timeutil.ParseAny("Oct 04, 2024")
t3, _ := timeutil.ParseAny("2024년 10월 04일")
t4, _ := timeutil.ParseAny("2024/10/04")
```

### ParseWithLayout

Parse with a custom layout.

사용자 지정 레이아웃으로 파싱합니다.

```go
func ParseWithLayout(s, layout string) (time.Time, error)
```

**Example**:
```go
t, err := timeutil.ParseWithLayout("2024-10-04 08:34:42.324", "2006-01-02 15:04:05.000")
```

### ParseMillis

Parse datetime with milliseconds (YYYY-MM-DD HH:mm:ss.SSS).

밀리초를 포함한 날짜시간을 파싱합니다.

```go
func ParseMillis(s string) (time.Time, error)
```

**Example**:
```go
t, err := timeutil.ParseMillis("2024-10-04 08:34:42.324")
```

### ParseMicros

Parse datetime with microseconds (YYYY-MM-DD HH:mm:ss.SSSSSS).

마이크로초를 포함한 날짜시간을 파싱합니다.

```go
func ParseMicros(s string) (time.Time, error)
```

**Example**:
```go
t, err := timeutil.ParseMicros("2024-10-04 08:34:42.324567")
```

---

## String Version Functions / String 버전 함수

All major timeutil functions now have String versions! Here's the complete list:

모든 주요 timeutil 함수에 이제 String 버전이 있습니다! 전체 목록은 다음과 같습니다:

### Time Difference / 시간 차이

```go
// Calculate difference between two time strings
diff, err := timeutil.SubTimeString("2024-10-04 08:34:42", "2024-10-14 14:56:23")
fmt.Println(diff.String())  // "10 days 6 hours 21 minutes"

// Get days between dates
days, err := timeutil.DiffInDaysString("2024-10-04", "2024-10-14")
fmt.Printf("%.2f days\n", days)  // "10.00 days"

// Get hours between times
hours, err := timeutil.DiffInHoursString("2024-10-04 08:00", "2024-10-04 14:30")
fmt.Printf("%.2f hours\n", hours)  // "6.50 hours"

// Get minutes
minutes, err := timeutil.DiffInMinutesString("08:00", "14:30")
```

### Age Calculation / 나이 계산

```go
// Get full age details
age, err := timeutil.AgeString("1990-01-15")
fmt.Printf("%d years, %d months, %d days\n", age.Years, age.Months, age.Days)

// Get just years
years, err := timeutil.AgeInYearsString("1990-01-15")
fmt.Printf("%d years old\n", years)
```

### Relative Time / 상대 시간

```go
// Get human-readable relative time
rel, err := timeutil.RelativeTimeString("2024-10-13 15:30:00")
fmt.Println(rel)  // "1 day ago"
```

### Business Days / 영업일

```go
// Check if date is a business day
isBiz, err := timeutil.IsBusinessDayString("2024-10-14")
fmt.Println(isBiz)  // true (Monday)

// Check if weekend
isWeekend, err := timeutil.IsWeekendString("2024-10-12")
fmt.Println(isWeekend)  // true (Saturday)
```

### Date Arithmetic / 날짜 연산

```go
// Add days
future, err := timeutil.AddDaysString("2024-10-04", 7)
fmt.Println(timeutil.FormatDate(future))  // "2024-10-11"

// Subtract days
past, err := timeutil.SubDaysString("2024-10-14", 7)
fmt.Println(timeutil.FormatDate(past))  // "2024-10-07"

// Add hours
later, err := timeutil.AddHoursString("2024-10-04 08:00", 6)

// Add minutes
soon, err := timeutil.AddMinutesString("15:30", 45)
```

### Formatting / 포맷팅

```go
// Convert between formats
result, err := timeutil.FormatString("Oct 04, 2024", "2006-01-02")
fmt.Println(result)  // "2024-10-04"

// Format as date
date, err := timeutil.FormatDateString("2024-10-04 15:30:00")
fmt.Println(date)  // "2024-10-04"

// Format as ISO8601
iso, err := timeutil.FormatISO8601String("Oct 04, 2024")
fmt.Println(iso)  // "2024-10-04T00:00:00+09:00"
```

### Timezone / 타임존

```go
// Convert timezone
ny, err := timeutil.ConvertTimezoneString("2024-10-04 15:30:00", "America/New_York")
```

### Time Boundaries / 시간 경계

```go
// Start of day
start, err := timeutil.StartOfDayString("2024-10-04 15:30:45")
// Returns: 2024-10-04 00:00:00

// End of day
end, err := timeutil.EndOfDayString("2024-10-04")
// Returns: 2024-10-04 23:59:59

// Start/End of week, month, year also available
startWeek, _ := timeutil.StartOfWeekString("2024-10-04")
endMonth, _ := timeutil.EndOfMonthString("2024-10-04")
startYear, _ := timeutil.StartOfYearString("2024-10-04")
```

### Weekdays / 요일

```go
// Get weekday name
day, err := timeutil.WeekdayString("2024-10-14")
fmt.Println(day)  // "Monday"

// Korean weekday
dayKor, err := timeutil.WeekdayKoreanString("2024-10-14")
fmt.Println(dayKor)  // "월요일"

// Short Korean weekday
dayShort, err := timeutil.WeekdayShortKoreanString("2024-10-14")
fmt.Println(dayShort)  // "월"

// Weekday number (0=Sunday, 6=Saturday)
num, err := timeutil.WeekdayNumberString("2024-10-14")
fmt.Println(num)  // 1 (Monday)
```

### Week/Month Info / 주/월 정보

```go
// Week of year
week, err := timeutil.WeekOfYearString("2024-10-14")

// Week of month
weekMonth, err := timeutil.WeekOfMonthString("2024-10-14")

// Days in month
days, err := timeutil.DaysInMonthString("2024-10-04")

// Days in year
daysYear, err := timeutil.DaysInYearString("2024-01-01")
```

### Month Info / 월 정보

```go
// Korean month name
month, err := timeutil.MonthKoreanString("2024-10-04")
fmt.Println(month)  // "10월"

// Full month name
monthName, err := timeutil.MonthNameString("2024-10-04")
fmt.Println(monthName)  // "October"

// Short month name
monthShort, err := timeutil.MonthNameShortString("2024-10-04")
fmt.Println(monthShort)  // "Oct"

// Quarter
quarter, err := timeutil.QuarterString("2024-10-04")
fmt.Println(quarter)  // 4
```

### Comparisons / 비교

```go
// Check if same day
same, err := timeutil.IsSameDayString("2024-10-04 08:00", "Oct 04, 2024")
fmt.Println(same)  // true

// Check if before
before, err := timeutil.IsBeforeString("2024-10-04", "2024-10-14")
fmt.Println(before)  // true

// Check if after
after, err := timeutil.IsAfterString("2024-10-14", "2024-10-04")
fmt.Println(after)  // true

// Check if between
between, err := timeutil.IsBetweenString("2024-10-10", "2024-10-04", "2024-10-14")
fmt.Println(between)  // true
```

### Leap Year / 윤년

```go
// Check if leap year
isLeap, err := timeutil.IsLeapYearString("2024-01-01")
fmt.Println(isLeap)  // true (2024 is a leap year)
```

---

## Complete Function List / 전체 함수 목록

### Parse Functions (4 new) / 파싱 함수 (4개 신규)

| Function | Description |
|----------|-------------|
| `ParseAny(s)` | Auto-detect 40+ formats / 40개 이상 포맷 자동 감지 |
| `ParseWithLayout(s, layout)` | Parse with custom layout / 커스텀 레이아웃으로 파싱 |
| `ParseMillis(s)` | Parse with milliseconds / 밀리초로 파싱 |
| `ParseMicros(s)` | Parse with microseconds / 마이크로초로 파싱 |

### String Functions (46 new) / String 함수 (46개 신규)

All return `(result, error)` / 모두 `(결과, error)` 반환

**Time Difference / 시간 차이**:
- `SubTimeString(s1, s2)` → `*TimeDiff`
- `DiffInDaysString(s1, s2)` → `float64`
- `DiffInHoursString(s1, s2)` → `float64`
- `DiffInMinutesString(s1, s2)` → `float64`

**Age Calculation / 나이 계산**:
- `AgeString(birthDate)` → `*AgeDetail`
- `AgeInYearsString(birthDate)` → `int`

**Relative Time / 상대 시간**:
- `RelativeTimeString(s)` → `string`

**Business Days / 영업일**:
- `IsBusinessDayString(s)` → `bool`
- `IsWeekendString(s)` → `bool`

**Date Arithmetic / 날짜 연산**:
- `AddDaysString(s, days)` → `time.Time`
- `AddHoursString(s, hours)` → `time.Time`
- `AddMinutesString(s, minutes)` → `time.Time`
- `SubDaysString(s, days)` → `time.Time`
- `SubHoursString(s, hours)` → `time.Time`
- `SubMinutesString(s, minutes)` → `time.Time`

**Formatting / 포맷팅**:
- `FormatString(s, layout)` → `string`
- `FormatDateString(s)` → `string`
- `FormatDateTimeString(s)` → `string`
- `FormatISO8601String(s)` → `string`

**Timezone / 타임존**:
- `ConvertTimezoneString(s, tz)` → `time.Time`

**Time Boundaries / 시간 경계**:
- `StartOfDayString(s)` → `time.Time`
- `EndOfDayString(s)` → `time.Time`
- `StartOfWeekString(s)` → `time.Time`
- `EndOfWeekString(s)` → `time.Time`
- `StartOfMonthString(s)` → `time.Time`
- `EndOfMonthString(s)` → `time.Time`
- `StartOfYearString(s)` → `time.Time`
- `EndOfYearString(s)` → `time.Time`

**Weekdays / 요일**:
- `WeekdayString(s)` → `string`
- `WeekdayKoreanString(s)` → `string`
- `WeekdayShortString(s)` → `string`
- `WeekdayShortKoreanString(s)` → `string`
- `WeekdayNumberString(s)` → `int`

**Week/Month Info / 주/월 정보**:
- `WeekOfYearString(s)` → `int`
- `WeekOfMonthString(s)` → `int`
- `DaysInMonthString(s)` → `int`
- `DaysInYearString(s)` → `int`

**Month Info / 월 정보**:
- `MonthKoreanString(s)` → `string`
- `MonthNameString(s)` → `string`
- `MonthNameShortString(s)` → `string`
- `QuarterString(s)` → `int`

**Leap Year / 윤년**:
- `IsLeapYearString(s)` → `bool`

**Comparisons / 비교**:
- `IsSameDayString(s1, s2)` → `bool`
- `IsBeforeString(s1, s2)` → `bool`
- `IsAfterString(s1, s2)` → `bool`
- `IsBetweenString(s, start, end)` → `bool`

---

## Use Cases / 사용 사례

### 1. Database Queries / 데이터베이스 쿼리

```go
// Fetch timestamps from database
var createdAt, updatedAt string
db.QueryRow("SELECT created_at, updated_at FROM users WHERE id = ?", userID).
    Scan(&createdAt, &updatedAt)

// Calculate time difference
diff, err := timeutil.SubTimeString(createdAt, updatedAt)
if err != nil {
    return err
}

fmt.Printf("User updated %s after creation\n", diff.Humanize())
```

### 2. API Responses / API 응답

```go
type APIResponse struct {
    Timestamp string `json:"timestamp"`
    ExpiresAt string `json:"expires_at"`
}

func handleResponse(resp APIResponse) error {
    // Parse timestamps from any format
    expiry, err := timeutil.ParseAny(resp.ExpiresAt)
    if err != nil {
        return err
    }

    // Check if expired
    if expiry.Before(time.Now()) {
        return fmt.Errorf("token expired")
    }

    return nil
}
```

### 3. CSV/File Parsing / CSV/파일 파싱

```go
// Parse dates from CSV
func parseCSVRow(row []string) (*Record, error) {
    // Dates can be in any format!
    startDate, err := timeutil.ParseAny(row[0])
    if err != nil {
        return nil, err
    }

    endDate, err := timeutil.ParseAny(row[1])
    if err != nil {
        return nil, err
    }

    return &Record{
        Start: startDate,
        End:   endDate,
    }, nil
}
```

### 4. User Input / 사용자 입력

```go
// Accept flexible date input from users
func calculateAge(birthDateInput string) (int, error) {
    // Users can enter: "1990-01-15", "Jan 15, 1990", "15-01-1990", etc.
    age, err := timeutil.AgeInYearsString(birthDateInput)
    if err != nil {
        return 0, fmt.Errorf("invalid date format: %w", err)
    }

    return age, nil
}
```

### 5. Log File Analysis / 로그 파일 분석

```go
// Parse log timestamps
func analyzeLogs(logLines []string) {
    for _, line := range logLines {
        // Extract timestamp (could be in any format)
        timestamp := extractTimestamp(line)

        t, err := timeutil.ParseAny(timestamp)
        if err != nil {
            continue
        }

        // Get relative time
        rel, _ := timeutil.RelativeTimeString(timestamp)
        fmt.Printf("Log from %s\n", rel)
    }
}
```

---

## Best Practices / 모범 사례

### 1. Always Check Errors / 항상 에러 확인

```go
// ✅ Good
diff, err := timeutil.SubTimeString(s1, s2)
if err != nil {
    return fmt.Errorf("failed to parse time: %w", err)
}

// ❌ Bad
diff, _ := timeutil.SubTimeString(s1, s2)  // Ignoring errors!
```

### 2. Use ParseAny for Unknown Formats / 알 수 없는 포맷에는 ParseAny 사용

```go
// ✅ When you don't know the exact format
t, err := timeutil.ParseAny(unknownFormatString)

// ✅ When you know the exact format (faster)
t, err := timeutil.ParseMillis("2024-10-04 08:34:42.324")
```

### 3. Validate Before Processing / 처리 전 검증

```go
func processDate(dateStr string) error {
    // Validate first
    _, err := timeutil.ParseAny(dateStr)
    if err != nil {
        return fmt.Errorf("invalid date format: %w", err)
    }

    // Now safe to use String functions
    age, _ := timeutil.AgeInYearsString(dateStr)
    // ...
}
```

### 4. Cache Parsed Times for Repeated Use / 반복 사용시 파싱 결과 캐싱

```go
// ❌ Bad - parsing multiple times
for i := 0; i < 1000; i++ {
    isLeap, _ := timeutil.IsLeapYearString("2024-01-01")
}

// ✅ Good - parse once, reuse
t, _ := timeutil.ParseAny("2024-01-01")
for i := 0; i < 1000; i++ {
    isLeap := timeutil.IsLeapYear(t)
}
```

---

## Performance Notes / 성능 참고사항

### ParseAny Performance / ParseAny 성능

`ParseAny` tries multiple formats sequentially, so it's slightly slower than using specific parse functions.

`ParseAny`는 여러 포맷을 순차적으로 시도하므로 특정 파싱 함수를 사용하는 것보다 약간 느립니다.

**Benchmarks / 벤치마크**:
```
ParseMillis:     ~200 ns/op  (fastest)
ParseAny:        ~800 ns/op  (still very fast!)
```

**Recommendation / 권장사항**:
- Use `ParseAny` for unknown formats / 알 수 없는 포맷에는 `ParseAny` 사용
- Use specific parsers (`ParseMillis`, `ParseMicros`) when format is known / 포맷을 알 때는 특정 파서 사용

---

## Troubleshooting / 문제 해결

### Q: ParseAny returns error for my format / ParseAny가 내 포맷에 에러 반환

**A**: If your format isn't in the 40+ supported formats, use `ParseWithLayout`:

**답변**: 40개 이상의 지원 포맷에 없는 경우, `ParseWithLayout` 사용:

```go
// Custom format
t, err := timeutil.ParseWithLayout("04.10.2024 15:30", "02.01.2006 15:04")
```

### Q: Korean format not working / 한글 포맷이 작동하지 않음

**A**: Make sure your string exactly matches one of the Korean formats:

**답변**: 문자열이 한글 포맷 중 하나와 정확히 일치하는지 확인:

```go
// ✅ Correct
timeutil.ParseAny("2024년 10월 04일 15시 30분 45초")

// ❌ Wrong - missing units
timeutil.ParseAny("2024 10 04 15 30 45")
```

### Q: Time parsed in wrong timezone / 잘못된 타임존으로 파싱됨

**A**: Use `ParseWithTimezone` or convert after parsing:

**답변**: `ParseWithTimezone` 사용 또는 파싱 후 변환:

```go
// Method 1: Parse in specific timezone
t, err := timeutil.ParseWithTimezone("2024-10-04 08:34:42", "America/New_York")

// Method 2: Convert after parsing
t, err := timeutil.ParseAny("2024-10-04 08:34:42")
t, err = timeutil.ConvertTimezone(t, "America/New_York")
```

---

## Migration Guide / 마이그레이션 가이드

### Migrating from Old Code / 기존 코드에서 마이그레이션

```go
// OLD (v1.6.007 and earlier)
layout := "2006-01-02 15:04:05"
t1, err := time.ParseInLocation(layout, str1, timeutil.KST)
if err != nil {
    return err
}
t2, err := time.ParseInLocation(layout, str2, timeutil.KST)
if err != nil {
    return err
}
diff := timeutil.SubTime(t1, t2)

// NEW (v1.6.008)
diff, err := timeutil.SubTimeString(str1, str2)
if err != nil {
    return err
}
```

---

## Summary / 요약

String parameter support in v1.6.008 brings:

v1.6.008의 문자열 매개변수 지원은 다음을 제공합니다:

✅ **40+ formats** automatically detected / 40개 이상 포맷 자동 감지
✅ **50+ String functions** for all major operations / 모든 주요 작업용 50개 이상 String 함수
✅ **Korean format** support / 한글 포맷 지원
✅ **Database-friendly** - works with MySQL, PostgreSQL, SQLite / 데이터베이스 친화적
✅ **API-friendly** - handles JSON timestamps / API 친화적
✅ **User-friendly** - accepts flexible input / 사용자 친화적

This makes timeutil the most comprehensive and easy-to-use time library in Go!

이것은 timeutil을 Go에서 가장 포괄적이고 사용하기 쉬운 시간 라이브러리로 만듭니다!

---

**Version**: v1.6.008
**Last Updated**: 2025-10-14
