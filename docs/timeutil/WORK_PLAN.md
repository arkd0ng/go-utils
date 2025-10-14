# Timeutil Package - Work Plan / 작업 계획서
# timeutil 패키지 - 작업 계획서

**Version / 버전**: v1.6.x
**Author / 작성자**: arkd0ng
**Created / 작성일**: 2025-10-14
**Status / 상태**: Planning / 계획 중

---

## Table of Contents / 목차

1. [Overview / 개요](#overview--개요)
2. [Work Phases / 작업 단계](#work-phases--작업-단계)
3. [Phase 1: Foundation / 1단계: 기초](#phase-1-foundation--1단계-기초)
4. [Phase 2: Core Features / 2단계: 핵심 기능](#phase-2-core-features--2단계-핵심-기능)
5. [Phase 3: Advanced Features / 3단계: 고급 기능](#phase-3-advanced-features--3단계-고급-기능)
6. [Phase 4: Testing & Documentation / 4단계: 테스팅 및 문서화](#phase-4-testing--documentation--4단계-테스팅-및-문서화)
7. [Phase 5: Release / 5단계: 릴리스](#phase-5-release--5단계-릴리스)
8. [Task Dependencies / 작업 의존성](#task-dependencies--작업-의존성)
9. [Quality Checklist / 품질 체크리스트](#quality-checklist--품질-체크리스트)

---

## Overview / 개요

This work plan outlines the detailed implementation steps for the `timeutil` package. Each phase is broken down into specific tasks with clear acceptance criteria.

이 작업 계획은 `timeutil` 패키지의 상세한 구현 단계를 설명합니다. 각 단계는 명확한 수용 기준과 함께 구체적인 작업으로 나뉩니다.

### Project Timeline / 프로젝트 타임라인

- **Phase 1**: Foundation / 기초 (1-2 작업 단위)
- **Phase 2**: Core Features / 핵심 기능 (8-10 작업 단위)
- **Phase 3**: Advanced Features / 고급 기능 (2-3 작업 단위)
- **Phase 4**: Testing & Documentation / 테스팅 및 문서화 (3-4 작업 단위)
- **Phase 5**: Release / 릴리스 (1-2 작업 단위)

**Total Estimated Work Units / 총 예상 작업 단위**: 15-21 units

---

## Work Phases / 작업 단계

### Priority Legend / 우선순위 범례

- 🔴 **P0**: Critical / 필수 - Must have for MVP / MVP를 위해 반드시 필요
- 🟡 **P1**: High / 높음 - Important for production readiness / 프로덕션 준비를 위해 중요
- 🟢 **P2**: Medium / 보통 - Nice to have / 있으면 좋음
- 🔵 **P3**: Low / 낮음 - Future enhancement / 향후 개선사항

---

## Phase 1: Foundation / 1단계: 기초

### Task 1.1: Project Structure Setup / 프로젝트 구조 설정

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Create the basic directory structure and initialize the package files.

기본 디렉토리 구조를 생성하고 패키지 파일을 초기화합니다.

**Subtasks / 하위 작업**:

1. Create directory structure / 디렉토리 구조 생성:
   ```bash
   mkdir -p timeutil
   mkdir -p examples/timeutil
   mkdir -p docs/timeutil
   mkdir -p docs/CHANGELOG
   ```

2. Create initial package files / 초기 패키지 파일 생성:
   - `timeutil/timeutil.go` (package doc and core types)
   - `timeutil/diff.go` (time difference functions)
   - `timeutil/timezone.go` (timezone operations)
   - `timeutil/arithmetic.go` (date arithmetic)
   - `timeutil/format.go` (date formatting)
   - `timeutil/parse.go` (time parsing)
   - `timeutil/business.go` (business days)
   - `timeutil/comparison.go` (time comparisons)
   - `timeutil/age.go` (age calculations)
   - `timeutil/relative.go` (relative time)
   - `timeutil/unix.go` (unix timestamp)
   - `timeutil/constants.go` (constants and helpers)

3. Create test files / 테스트 파일 생성:
   - `timeutil/diff_test.go`
   - `timeutil/timezone_test.go`
   - `timeutil/arithmetic_test.go`
   - `timeutil/format_test.go`
   - `timeutil/parse_test.go`
   - `timeutil/business_test.go`
   - `timeutil/comparison_test.go`
   - `timeutil/age_test.go`
   - `timeutil/relative_test.go`
   - `timeutil/unix_test.go`

4. Add package documentation / 패키지 문서 추가:
   - `timeutil/README.md` (initial version)
   - `docs/CHANGELOG/CHANGELOG-v1.6.md` (initial version)

**Acceptance Criteria / 수용 기준**:
- [ ] All directories created / 모든 디렉토리 생성됨
- [ ] All initial files created with package declarations / 패키지 선언과 함께 모든 초기 파일 생성됨
- [ ] Package compiles without errors / 패키지가 에러 없이 컴파일됨
- [ ] Initial README.md created / 초기 README.md 생성됨
- [ ] CHANGELOG-v1.6.md created / CHANGELOG-v1.6.md 생성됨

**Version / 버전**: v1.6.001

---

## Phase 2: Core Features / 2단계: 핵심 기능

### Task 2.1: Core Types and Constants / 핵심 타입 및 상수

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Implement core types (TimeDiff, Age) and constants.

핵심 타입(TimeDiff, Age) 및 상수를 구현합니다.

**Files / 파일**:
- `timeutil/timeutil.go`
- `timeutil/constants.go`

**Implementation / 구현**:

```go
// timeutil/timeutil.go
package timeutil

import "time"

// TimeDiff represents the difference between two times
// TimeDiff는 두 시간 사이의 차이를 나타냅니다
type TimeDiff struct {
    Duration time.Duration
}

// Methods for TimeDiff / TimeDiff 메서드
func (td *TimeDiff) Seconds() float64
func (td *TimeDiff) Minutes() float64
func (td *TimeDiff) Hours() float64
func (td *TimeDiff) Days() float64
func (td *TimeDiff) Weeks() float64
func (td *TimeDiff) String() string
func (td *TimeDiff) Humanize() string
func (td *TimeDiff) Abs() *TimeDiff

// Age represents a person's age in years, months, and days
// Age는 년, 월, 일 단위의 나이를 나타냅니다
type Age struct {
    Years  int
    Months int
    Days   int
}

// Methods for Age / Age 메서드
func (a *Age) String() string
func (a *Age) Humanize() string

// timeutil/constants.go
package timeutil

const (
    SecondsPerMinute = 60
    SecondsPerHour   = 3600
    SecondsPerDay    = 86400
    DaysPerWeek      = 7
    MonthsPerYear    = 12
)
```

**Tests / 테스트**:
- Test TimeDiff methods / TimeDiff 메서드 테스트
- Test Age methods / Age 메서드 테스트
- Benchmark tests / 벤치마크 테스트

**Acceptance Criteria / 수용 기준**:
- [ ] TimeDiff type implemented with all methods / 모든 메서드와 함께 TimeDiff 타입 구현됨
- [ ] Age type implemented with all methods / 모든 메서드와 함께 Age 타입 구현됨
- [ ] Constants defined / 상수 정의됨
- [ ] All tests passing / 모든 테스트 통과
- [ ] Benchmark tests added / 벤치마크 테스트 추가됨

**Version / 버전**: v1.6.002

---

### Task 2.2: Time Difference Functions / 시간 차이 함수

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Implement time difference calculation functions.

시간 차이 계산 함수를 구현합니다.

**Files / 파일**:
- `timeutil/diff.go`
- `timeutil/diff_test.go`

**Functions / 함수** (8 functions):
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

**Tests / 테스트**:
- Test all diff functions / 모든 diff 함수 테스트
- Test edge cases (same time, negative diff) / 엣지 케이스 테스트 (같은 시간, 음수 차이)
- Test with different timezones / 다른 타임존으로 테스트
- Benchmark tests / 벤치마크 테스트

**Acceptance Criteria / 수용 기준**:
- [ ] All 8 functions implemented / 8개 함수 모두 구현됨
- [ ] All tests passing with edge cases / 엣지 케이스 포함 모든 테스트 통과
- [ ] Benchmark tests added / 벤치마크 테스트 추가됨
- [ ] Bilingual documentation / 이중 언어 문서

**Version / 버전**: v1.6.003

---

### Task 2.3: Timezone Operations / 타임존 작업

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Implement timezone conversion and management functions.

타임존 변환 및 관리 함수를 구현합니다.

**Files / 파일**:
- `timeutil/timezone.go`
- `timeutil/timezone_test.go`

**Functions / 함수** (5 functions):
```go
ConvertTimezone(t time.Time, tz string) (time.Time, error)
GetTimezoneOffset(tz string) (int, error)
ListTimezones() []string
IsValidTimezone(tz string) bool
GetLocalTimezone() string
```

**Implementation Details / 구현 세부사항**:
- Cache loaded timezones for performance / 성능을 위해 로드된 타임존 캐시
- Thread-safe timezone cache / 스레드 안전 타임존 캐시
- Support all IANA timezone names / 모든 IANA 타임존 이름 지원

**Tests / 테스트**:
- Test timezone conversion / 타임존 변환 테스트
- Test invalid timezones / 잘못된 타임존 테스트
- Test caching mechanism / 캐싱 메커니즘 테스트
- Benchmark tests / 벤치마크 테스트

**Acceptance Criteria / 수용 기준**:
- [ ] All 5 functions implemented / 5개 함수 모두 구현됨
- [ ] Timezone caching working / 타임존 캐싱 작동
- [ ] All tests passing / 모든 테스트 통과
- [ ] Benchmark tests added / 벤치마크 테스트 추가됨

**Version / 버전**: v1.6.004

---

### Task 2.4: Date Arithmetic / 날짜 연산

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Implement date arithmetic functions (add/subtract days, weeks, months, etc.).

날짜 연산 함수를 구현합니다 (일, 주, 월 등 더하기/빼기).

**Files / 파일**:
- `timeutil/arithmetic.go`
- `timeutil/arithmetic_test.go`

**Functions / 함수** (16 functions):
```go
// Add/Subtract time units / 시간 단위 더하기/빼기
AddSeconds(t time.Time, seconds int) time.Time
AddMinutes(t time.Time, minutes int) time.Time
AddHours(t time.Time, hours int) time.Time
AddDays(t time.Time, days int) time.Time
AddWeeks(t time.Time, weeks int) time.Time
AddMonths(t time.Time, months int) time.Time
AddYears(t time.Time, years int) time.Time

// Start/End of period / 기간의 시작/끝
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

**Tests / 테스트**:
- Test all arithmetic functions / 모든 연산 함수 테스트
- Test edge cases (month boundaries, leap years) / 엣지 케이스 테스트 (월 경계, 윤년)
- Test with different timezones / 다른 타임존으로 테스트
- Benchmark tests / 벤치마크 테스트

**Acceptance Criteria / 수용 기준**:
- [ ] All 16 functions implemented / 16개 함수 모두 구현됨
- [ ] Month/year boundary handling correct / 월/년 경계 처리 정확
- [ ] All tests passing / 모든 테스트 통과
- [ ] Benchmark tests added / 벤치마크 테스트 추가됨

**Version / 버전**: v1.6.005

---

### Task 2.5: Date Formatting / 날짜 포맷팅

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Implement date formatting functions with intuitive format strings.

직관적인 포맷 문자열로 날짜 포맷팅 함수를 구현합니다.

**Files / 파일**:
- `timeutil/format.go`
- `timeutil/format_test.go`

**Functions / 함수** (8 functions):
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

**Custom Format Tokens / 커스텀 포맷 토큰**:
- `YYYY` - 4-digit year
- `YY` - 2-digit year
- `MM` - 2-digit month
- `M` - 1 or 2-digit month
- `DD` - 2-digit day
- `D` - 1 or 2-digit day
- `HH` - 2-digit hour (24h)
- `hh` - 2-digit hour (12h)
- `mm` - 2-digit minute
- `ss` - 2-digit second

**Tests / 테스트**:
- Test all format functions / 모든 포맷 함수 테스트
- Test custom format strings / 커스텀 포맷 문자열 테스트
- Test with different timezones / 다른 타임존으로 테스트
- Benchmark tests / 벤치마크 테스트

**Acceptance Criteria / 수용 기준**:
- [ ] All 8 functions implemented / 8개 함수 모두 구현됨
- [ ] Custom format tokens working / 커스텀 포맷 토큰 작동
- [ ] All tests passing / 모든 테스트 통과
- [ ] Benchmark tests added / 벤치마크 테스트 추가됨

**Version / 버전**: v1.6.006

---

### Task 2.6: Time Parsing / 시간 파싱

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Implement time parsing functions for common formats.

일반 포맷에 대한 시간 파싱 함수를 구현합니다.

**Files / 파일**:
- `timeutil/parse.go`
- `timeutil/parse_test.go`

**Functions / 함수** (6 functions):
```go
ParseISO8601(s string) (time.Time, error)
ParseRFC3339(s string) (time.Time, error)
ParseDate(s string) (time.Time, error)
ParseDateTime(s string) (time.Time, error)
Parse(s string) (time.Time, error)
ParseWithTimezone(s, tz string) (time.Time, error)
```

**Implementation Details / 구현 세부사항**:
- Smart parse: auto-detect format / 스마트 파싱: 포맷 자동 감지
- Support multiple common formats / 여러 일반 포맷 지원
- Clear error messages / 명확한 에러 메시지

**Tests / 테스트**:
- Test all parse functions / 모든 파싱 함수 테스트
- Test invalid formats / 잘못된 포맷 테스트
- Test auto-detection / 자동 감지 테스트
- Benchmark tests / 벤치마크 테스트

**Acceptance Criteria / 수용 기준**:
- [ ] All 6 functions implemented / 6개 함수 모두 구현됨
- [ ] Smart parse working / 스마트 파싱 작동
- [ ] All tests passing / 모든 테스트 통과
- [ ] Benchmark tests added / 벤치마크 테스트 추가됨

**Version / 버전**: v1.6.007

---

### Task 2.7: Time Comparisons / 시간 비교

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Implement time comparison functions.

시간 비교 함수를 구현합니다.

**Files / 파일**:
- `timeutil/comparison.go`
- `timeutil/comparison_test.go`

**Functions / 함수** (18 functions):
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

**Tests / 테스트**:
- Test all comparison functions / 모든 비교 함수 테스트
- Test edge cases / 엣지 케이스 테스트
- Test with different timezones / 다른 타임존으로 테스트
- Benchmark tests / 벤치마크 테스트

**Acceptance Criteria / 수용 기준**:
- [ ] All 18 functions implemented / 18개 함수 모두 구현됨
- [ ] All tests passing / 모든 테스트 통과
- [ ] Benchmark tests added / 벤치마크 테스트 추가됨

**Version / 버전**: v1.6.008

---

### Task 2.8: Unix Timestamp / Unix 타임스탬프

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Implement Unix timestamp operations.

Unix 타임스탬프 작업을 구현합니다.

**Files / 파일**:
- `timeutil/unix.go`
- `timeutil/unix_test.go`

**Functions / 함수** (8 functions):
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

**Tests / 테스트**:
- Test all unix functions / 모든 unix 함수 테스트
- Test conversion accuracy / 변환 정확도 테스트
- Benchmark tests / 벤치마크 테스트

**Acceptance Criteria / 수용 기준**:
- [ ] All 8 functions implemented / 8개 함수 모두 구현됨
- [ ] All tests passing / 모든 테스트 통과
- [ ] Benchmark tests added / 벤치마크 테스트 추가됨

**Version / 버전**: v1.6.009

---

## Phase 3: Advanced Features / 3단계: 고급 기능

### Task 3.1: Business Days / 영업일

**Priority / 우선순위**: 🟡 P1

**Description / 설명**:
Implement business day calculations.

영업일 계산을 구현합니다.

**Files / 파일**:
- `timeutil/business.go`
- `timeutil/business_test.go`

**Functions / 함수** (6 functions):
```go
AddBusinessDays(t time.Time, days int) time.Time
IsBusinessDay(t time.Time) bool
CountBusinessDays(start, end time.Time) int
NextBusinessDay(t time.Time) time.Time
PreviousBusinessDay(t time.Time) time.Time
IsHoliday(t time.Time) bool
```

**Holiday Support / 공휴일 지원**:
```go
SetHolidays(holidays []time.Time)
GetHolidays() []time.Time
ClearHolidays()
```

**Tests / 테스트**:
- Test business day calculations / 영업일 계산 테스트
- Test with custom holidays / 커스텀 공휴일로 테스트
- Test edge cases (weekends, holidays) / 엣지 케이스 테스트
- Benchmark tests / 벤치마크 테스트

**Acceptance Criteria / 수용 기준**:
- [ ] All 6 functions implemented / 6개 함수 모두 구현됨
- [ ] Holiday support working / 공휴일 지원 작동
- [ ] All tests passing / 모든 테스트 통과
- [ ] Benchmark tests added / 벤치마크 테스트 추가됨

**Version / 버전**: v1.6.010

---

### Task 3.2: Age Calculations / 나이 계산

**Priority / 우선순위**: 🟡 P1

**Description / 설명**:
Implement age calculation functions.

나이 계산 함수를 구현합니다.

**Files / 파일**:
- `timeutil/age.go`
- `timeutil/age_test.go`

**Functions / 함수** (4 functions):
```go
AgeInYears(birthDate time.Time) int
AgeInMonths(birthDate time.Time) int
AgeInDays(birthDate time.Time) int
Age(birthDate time.Time) *Age
```

**Tests / 테스트**:
- Test age calculations / 나이 계산 테스트
- Test edge cases (leap years, different timezones) / 엣지 케이스 테스트
- Benchmark tests / 벤치마크 테스트

**Acceptance Criteria / 수용 기준**:
- [ ] All 4 functions implemented / 4개 함수 모두 구현됨
- [ ] Accurate age calculations / 정확한 나이 계산
- [ ] All tests passing / 모든 테스트 통과
- [ ] Benchmark tests added / 벤치마크 테스트 추가됨

**Version / 버전**: v1.6.011

---

### Task 3.3: Relative Time / 상대 시간

**Priority / 우선순위**: 🟡 P1

**Description / 설명**:
Implement relative time functions (e.g., "2 hours ago", "in 3 days").

상대 시간 함수를 구현합니다 (예: "2 hours ago", "in 3 days").

**Files / 파일**:
- `timeutil/relative.go`
- `timeutil/relative_test.go`

**Functions / 함수** (3 functions):
```go
RelativeTime(t time.Time) string
RelativeTimeShort(t time.Time) string
TimeAgo(t time.Time) string
```

**Implementation Details / 구현 세부사항**:
- Support both past and future / 과거와 미래 모두 지원
- Human-readable strings / 사람이 읽기 쉬운 문자열
- Short format option / 짧은 포맷 옵션

**Tests / 테스트**:
- Test past and future times / 과거 및 미래 시간 테스트
- Test different time ranges / 다른 시간 범위 테스트
- Test short format / 짧은 포맷 테스트
- Benchmark tests / 벤치마크 테스트

**Acceptance Criteria / 수용 기준**:
- [ ] All 3 functions implemented / 3개 함수 모두 구현됨
- [ ] Human-readable output / 사람이 읽기 쉬운 출력
- [ ] All tests passing / 모든 테스트 통과
- [ ] Benchmark tests added / 벤치마크 테스트 추가됨

**Version / 버전**: v1.6.012

---

## Phase 4: Testing & Documentation / 4단계: 테스팅 및 문서화

### Task 4.1: Comprehensive Testing / 종합 테스팅

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Ensure 100% test coverage with comprehensive test cases.

포괄적인 테스트 케이스로 100% 테스트 커버리지를 확보합니다.

**Subtasks / 하위 작업**:

1. Review all test files / 모든 테스트 파일 검토
2. Add missing test cases / 누락된 테스트 케이스 추가
3. Add edge case tests / 엣지 케이스 테스트 추가
4. Add benchmark tests / 벤치마크 테스트 추가
5. Run coverage analysis / 커버리지 분석 실행

**Commands / 명령어**:
```bash
# Run all tests / 모든 테스트 실행
go test ./timeutil -v

# Run with coverage / 커버리지와 함께 실행
go test ./timeutil -cover
go test ./timeutil -coverprofile=coverage.out
go tool cover -html=coverage.out

# Run benchmarks / 벤치마크 실행
go test ./timeutil -bench=.
```

**Acceptance Criteria / 수용 기준**:
- [ ] All tests passing / 모든 테스트 통과
- [ ] Test coverage ≥ 90% / 테스트 커버리지 ≥ 90%
- [ ] All edge cases covered / 모든 엣지 케이스 커버
- [ ] Benchmark tests for all functions / 모든 함수에 대한 벤치마크 테스트

**Version / 버전**: v1.6.013

---

### Task 4.2: Example Code / 예제 코드

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Create comprehensive example code demonstrating all features.

모든 기능을 보여주는 포괄적인 예제 코드를 작성합니다.

**Files / 파일**:
- `examples/timeutil/main.go`

**Example Sections / 예제 섹션**:
1. Time difference calculations / 시간 차이 계산
2. Timezone conversions / 타임존 변환
3. Date arithmetic / 날짜 연산
4. Date formatting / 날짜 포맷팅
5. Time parsing / 시간 파싱
6. Business days / 영업일
7. Time comparisons / 시간 비교
8. Age calculations / 나이 계산
9. Relative time / 상대 시간
10. Unix timestamp / Unix 타임스탬프

**Acceptance Criteria / 수용 기준**:
- [ ] Example code created / 예제 코드 생성됨
- [ ] All 10 sections covered / 10개 섹션 모두 커버
- [ ] Example runs without errors / 예제가 에러 없이 실행됨
- [ ] Bilingual comments / 이중 언어 주석

**Version / 버전**: v1.6.014

---

### Task 4.3: Documentation / 문서화

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Create comprehensive documentation (README, USER_MANUAL, DEVELOPER_GUIDE).

포괄적인 문서를 작성합니다 (README, USER_MANUAL, DEVELOPER_GUIDE).

**Files / 파일**:
- `timeutil/README.md` (comprehensive version)
- `docs/timeutil/USER_MANUAL.md`
- `docs/timeutil/DEVELOPER_GUIDE.md`

**README.md Sections / README.md 섹션**:
1. Package overview / 패키지 개요
2. Installation / 설치
3. Quick start / 빠른 시작
4. All function categories / 모든 함수 카테고리
5. Examples / 예제
6. Testing / 테스팅
7. Documentation links / 문서 링크

**USER_MANUAL.md Sections / USER_MANUAL.md 섹션** (~1500 lines):
1. Introduction / 소개
2. Installation / 설치
3. Quick Start / 빠른 시작
4. Configuration Reference / 설정 참조 (all functions)
5. Usage Patterns / 사용 패턴
6. Common Use Cases / 일반적인 사용 사례
7. Best Practices / 모범 사례
8. Troubleshooting / 문제 해결
9. FAQ

**DEVELOPER_GUIDE.md Sections / DEVELOPER_GUIDE.md 섹션** (~1000 lines):
1. Architecture Overview / 아키텍처 개요
2. Package Structure / 패키지 구조
3. Core Components / 핵심 컴포넌트
4. Internal Implementation / 내부 구현
5. Design Patterns / 디자인 패턴
6. Adding New Features / 새 기능 추가
7. Testing Guide / 테스트 가이드
8. Performance / 성능
9. Contributing Guidelines / 기여 가이드라인
10. Code Style / 코드 스타일

**Acceptance Criteria / 수용 기준**:
- [ ] README.md completed / README.md 완성됨
- [ ] USER_MANUAL.md completed (~1500 lines) / USER_MANUAL.md 완성됨
- [ ] DEVELOPER_GUIDE.md completed (~1000 lines) / DEVELOPER_GUIDE.md 완성됨
- [ ] All documentation bilingual / 모든 문서 이중 언어
- [ ] Code examples in documentation / 문서의 코드 예제

**Version / 버전**: v1.6.015

---

## Phase 5: Release / 5단계: 릴리스

### Task 5.1: Final Review & Release / 최종 검토 및 릴리스

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Final review, update version, and commit to GitHub.

최종 검토, 버전 업데이트 및 GitHub에 커밋합니다.

**Subtasks / 하위 작업**:

1. Update version in cfg/app.yaml / cfg/app.yaml의 버전 업데이트
2. Update root README.md / 루트 README.md 업데이트
3. Update root CHANGELOG.md / 루트 CHANGELOG.md 업데이트
4. Update CLAUDE.md / CLAUDE.md 업데이트
5. Final testing / 최종 테스팅:
   ```bash
   go build ./...
   go test ./timeutil -v
   go test ./timeutil -cover
   go run examples/timeutil/main.go
   ```

6. Git commit and push / Git 커밋 및 푸시:
   ```bash
   git add .
   git commit -m "Feat: Add timeutil package with 80+ time/date utility functions (v1.6.015)"
   git push
   ```

**Acceptance Criteria / 수용 기준**:
- [ ] Version updated in cfg/app.yaml / cfg/app.yaml의 버전 업데이트됨
- [ ] Root README.md updated / 루트 README.md 업데이트됨
- [ ] Root CHANGELOG.md updated / 루트 CHANGELOG.md 업데이트됨
- [ ] CLAUDE.md updated / CLAUDE.md 업데이트됨
- [ ] All tests passing / 모든 테스트 통과
- [ ] Committed and pushed to GitHub / GitHub에 커밋 및 푸시됨

**Version / 버전**: v1.6.015 (Final release version)

---

## Task Dependencies / 작업 의존성

```
Phase 1: Foundation
└── Task 1.1 (v1.6.001)

Phase 2: Core Features
├── Task 2.1 (v1.6.002) → Depends on 1.1
├── Task 2.2 (v1.6.003) → Depends on 2.1
├── Task 2.3 (v1.6.004) → Depends on 2.1
├── Task 2.4 (v1.6.005) → Depends on 2.1
├── Task 2.5 (v1.6.006) → Depends on 2.1
├── Task 2.6 (v1.6.007) → Depends on 2.1, 2.5
├── Task 2.7 (v1.6.008) → Depends on 2.1
└── Task 2.8 (v1.6.009) → Depends on 2.1

Phase 3: Advanced Features
├── Task 3.1 (v1.6.010) → Depends on 2.4
├── Task 3.2 (v1.6.011) → Depends on 2.1, 2.2
└── Task 3.3 (v1.6.012) → Depends on 2.1, 2.2

Phase 4: Testing & Documentation
├── Task 4.1 (v1.6.013) → Depends on all Phase 2 & 3
├── Task 4.2 (v1.6.014) → Depends on all Phase 2 & 3
└── Task 4.3 (v1.6.015) → Depends on all Phase 2 & 3

Phase 5: Release
└── Task 5.1 (v1.6.015) → Depends on all Phase 4
```

---

## Quality Checklist / 품질 체크리스트

### Code Quality / 코드 품질

- [ ] All functions have bilingual documentation / 모든 함수에 이중 언어 문서
- [ ] No external dependencies (standard library only) / 외부 의존성 없음 (표준 라이브러리만)
- [ ] All functions are thread-safe / 모든 함수가 스레드 안전
- [ ] No panic (sensible defaults) / 패닉 없음 (합리적 기본값)
- [ ] Consistent naming conventions / 일관된 명명 규칙
- [ ] Error messages are clear / 에러 메시지 명확

### Testing / 테스팅

- [ ] Test coverage ≥ 90% / 테스트 커버리지 ≥ 90%
- [ ] All edge cases tested / 모든 엣지 케이스 테스트됨
- [ ] Benchmark tests for all functions / 모든 함수에 대한 벤치마크 테스트
- [ ] Tests with different timezones / 다른 타임존으로 테스트
- [ ] All tests passing / 모든 테스트 통과

### Documentation / 문서화

- [ ] README.md complete and accurate / README.md 완성 및 정확
- [ ] USER_MANUAL.md comprehensive (~1500 lines) / USER_MANUAL.md 포괄적 (~1500줄)
- [ ] DEVELOPER_GUIDE.md comprehensive (~1000 lines) / DEVELOPER_GUIDE.md 포괄적 (~1000줄)
- [ ] All documentation bilingual / 모든 문서 이중 언어
- [ ] Code examples in documentation / 문서의 코드 예제
- [ ] CHANGELOG updated / CHANGELOG 업데이트됨

### Release Readiness / 릴리스 준비

- [ ] Version updated in cfg/app.yaml / cfg/app.yaml의 버전 업데이트됨
- [ ] Root README.md updated / 루트 README.md 업데이트됨
- [ ] Root CHANGELOG.md updated / 루트 CHANGELOG.md 업데이트됨
- [ ] CLAUDE.md updated / CLAUDE.md 업데이트됨
- [ ] Example code runs without errors / 예제 코드 에러 없이 실행
- [ ] All tests passing / 모든 테스트 통과
- [ ] Committed and pushed to GitHub / GitHub에 커밋 및 푸시됨

---

**Work Plan Status / 작업 계획 상태**: ✅ **APPROVED - Ready to Begin / 승인됨 - 시작 준비 완료**

**Total Functions / 총 함수 수**: ~80+ functions across 10 categories / 10개 카테고리에 걸쳐 약 80개 이상의 함수

**Estimated Completion / 예상 완료**: 15-21 work units / 15-21 작업 단위
