# Timeutil Package - Developer Guide / 개발자 가이드

**Version / 버전**: v1.6.002
**Package / 패키지**: `github.com/arkd0ng/go-utils/timeutil`
**Go Version / Go 버전**: 1.16+

---

## Table of Contents / 목차

1. [Architecture Overview / 아키텍처 개요](#architecture-overview--아키텍처-개요)
2. [Package Structure / 패키지 구조](#package-structure--패키지-구조)
3. [Core Components / 핵심 컴포넌트](#core-components--핵심-컴포넌트)
4. [Design Patterns / 디자인 패턴](#design-patterns--디자인-패턴)
5. [Internal Implementation / 내부 구현](#internal-implementation--내부-구현)
6. [Adding New Features / 새 기능 추가](#adding-new-features--새-기능-추가)
7. [Testing Guide / 테스트 가이드](#testing-guide--테스트-가이드)
8. [Performance / 성능](#performance--성능)
9. [Contributing Guidelines / 기여 가이드라인](#contributing-guidelines--기여-가이드라인)
10. [Code Style / 코드 스타일](#code-style--코드-스타일)

---

## Architecture Overview / 아키텍처 개요

### Design Principles / 설계 원칙

The `timeutil` package follows these core design principles:

`timeutil` 패키지는 다음과 같은 핵심 설계 원칙을 따릅니다:

1. **Extreme Simplicity / 극도의 간결함**
   - Reduce 20+ lines of code to 1-2 lines / 20줄 이상의 코드를 1-2줄로 줄임
   - Intuitive function names / 직관적인 함수 이름
   - No configuration required / 설정 불필요

2. **Human-Readable Output / 사람이 읽기 쉬운 출력**
   - Custom types with String() methods / String() 메서드가 있는 커스텀 타입
   - Relative time strings ("2 hours ago") / 상대 시간 문자열 ("2 hours ago")
   - Custom format tokens (YYYY-MM-DD) / 커스텀 포맷 토큰 (YYYY-MM-DD)

3. **KST Default Timezone / KST 기본 타임존**
   - Asia/Seoul (GMT+9) as package-wide default / Asia/Seoul (GMT+9)을 패키지 전체 기본값으로
   - Configurable via SetDefaultTimezone() / SetDefaultTimezone()으로 구성 가능
   - Thread-safe timezone caching / 스레드 안전 타임존 캐싱

4. **Zero Dependencies / 제로 의존성**
   - Standard library only / 표준 라이브러리만 사용
   - No external packages required / 외부 패키지 불필요

5. **Thread-Safe / 스레드 안전**
   - sync.RWMutex for shared state / 공유 상태를 위한 sync.RWMutex
   - Safe for concurrent use / 동시 사용에 안전

### High-Level Architecture / 상위 수준 아키텍처

```
┌─────────────────────────────────────────────────────────────┐
│                     timeutil Package                        │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐    │
│  │   Core Types │  │   Constants  │  │   Timezone   │    │
│  │  (TimeDiff,  │  │     (KST,    │  │    Cache     │    │
│  │  AgeDetail)  │  │  Locations)  │  │  (RWMutex)   │    │
│  └──────────────┘  └──────────────┘  └──────────────┘    │
│                                                             │
│  ┌─────────────────────────────────────────────────────┐  │
│  │              Function Categories                     │  │
│  ├─────────────────────────────────────────────────────┤  │
│  │ • Time Difference    • Timezone Operations          │  │
│  │ • Date Arithmetic    • Date Formatting              │  │
│  │ • Time Parsing       • Time Comparisons             │  │
│  │ • Age Calculations   • Relative Time                │  │
│  │ • Unix Timestamp     • Business Days                │  │
│  └─────────────────────────────────────────────────────┘  │
│                                                             │
│  ┌─────────────────────────────────────────────────────┐  │
│  │            Helper Functions                          │  │
│  ├─────────────────────────────────────────────────────┤  │
│  │ • Format token conversion (YYYY → 2006)             │  │
│  │ • Timezone loading and caching                       │  │
│  │ • Holiday management (thread-safe map)               │  │
│  └─────────────────────────────────────────────────────┘  │
│                                                             │
└─────────────────────────────────────────────────────────────┘
                            │
                            ▼
                 Go Standard Library
                   (time package)
```

---

## Package Structure / 패키지 구조

### File Organization / 파일 구성

The package is organized into 12 core files, each responsible for a specific category of functionality:

패키지는 12개의 핵심 파일로 구성되며, 각각 특정 기능 카테고리를 담당합니다:

```
timeutil/
├── timeutil.go          # Package documentation and core types
│                        # 패키지 문서 및 핵심 타입
├── constants.go         # Package constants and initialization
│                        # 패키지 상수 및 초기화
├── diff.go              # Time difference calculations
│                        # 시간 차이 계산
├── timezone.go          # Timezone operations and caching
│                        # 타임존 작업 및 캐싱
├── arithmetic.go        # Date arithmetic operations
│                        # 날짜 연산 작업
├── format.go            # Date formatting with custom tokens
│                        # 커스텀 토큰을 사용한 날짜 포맷팅
├── parse.go             # Time parsing with auto-detection
│                        # 자동 감지를 사용한 시간 파싱
├── comparison.go        # Time comparison functions
│                        # 시간 비교 함수
├── age.go               # Age calculation functions
│                        # 나이 계산 함수
├── relative.go          # Relative time strings
│                        # 상대 시간 문자열
├── unix.go              # Unix timestamp operations
│                        # Unix 타임스탬프 작업
├── business.go          # Business day calculations
│                        # 영업일 계산
├── timeutil_test.go     # Comprehensive test suite
│                        # 종합 테스트 스위트
└── README.md            # Package documentation
                         # 패키지 문서
```

### File Responsibilities / 파일 책임

| File / 파일 | Lines / 줄 수 | Responsibility / 책임 | Functions / 함수 수 |
|-------------|---------------|----------------------|---------------------|
| `timeutil.go` | ~100 | Core types, package doc / 핵심 타입, 패키지 문서 | 2 types |
| `constants.go` | ~50 | Constants, init() / 상수, init() | - |
| `diff.go` | ~150 | Time difference / 시간 차이 | 8 |
| `timezone.go` | ~200 | Timezone ops / 타임존 작업 | 10 |
| `arithmetic.go` | ~200 | Date arithmetic / 날짜 연산 | 16 |
| `format.go` | ~150 | Formatting / 포맷팅 | 8 |
| `parse.go` | ~120 | Parsing / 파싱 | 6 |
| `comparison.go` | ~250 | Comparisons / 비교 | 18 |
| `age.go` | ~100 | Age / 나이 | 4 |
| `relative.go` | ~150 | Relative time / 상대 시간 | 4 |
| `unix.go` | ~150 | Unix timestamps / Unix 타임스탬프 | 12 |
| `business.go` | ~200 | Business days / 영업일 | 7 |
| **Total / 합계** | **~1,820** | | **95+** |

---

## Core Components / 핵심 컴포넌트

### 1. Core Types / 핵심 타입

#### TimeDiff Type / TimeDiff 타입

The `TimeDiff` type wraps `time.Duration` to provide human-readable output.

`TimeDiff` 타입은 `time.Duration`을 래핑하여 사람이 읽기 쉬운 출력을 제공합니다.

```go
// TimeDiff represents the difference between two times
// TimeDiff는 두 시간 사이의 차이를 나타냅니다
type TimeDiff struct {
    Duration time.Duration
}

// Helper methods / 헬퍼 메서드
func (td *TimeDiff) Seconds() float64
func (td *TimeDiff) Minutes() float64
func (td *TimeDiff) Hours() float64
func (td *TimeDiff) Days() float64
func (td *TimeDiff) String() string       // "2 days 6 hours 30 minutes"
func (td *TimeDiff) Humanize() string     // "2d 6h 30m"
```

**Design Rationale / 설계 근거**:
- Provides convenience methods for common time units / 일반적인 시간 단위를 위한 편의 메서드 제공
- Human-readable output via String() and Humanize() / String() 및 Humanize()를 통한 사람이 읽기 쉬운 출력
- Wraps standard Duration for compatibility / 호환성을 위해 표준 Duration 래핑

#### AgeDetail Type / AgeDetail 타입

The `AgeDetail` type provides detailed age breakdown.

`AgeDetail` 타입은 상세한 나이 분석을 제공합니다.

```go
// AgeDetail represents a detailed age breakdown
// AgeDetail은 상세한 나이 분석을 나타냅니다
type AgeDetail struct {
    Years  int
    Months int
    Days   int
}

// Helper method / 헬퍼 메서드
func (a *AgeDetail) String() string  // "35 years 4 months 29 days"
```

**Design Rationale / 설계 근거**:
- Provides detailed age information beyond just years / 단순 년수를 넘어 상세한 나이 정보 제공
- Useful for displaying ages in various formats / 다양한 포맷으로 나이를 표시하는 데 유용
- String() method for easy display / 쉬운 표시를 위한 String() 메서드

### 2. Package Constants / 패키지 상수

Located in `constants.go`:

`constants.go`에 위치:

```go
const (
    // DefaultTimezone is the default timezone for all package functions
    // DefaultTimezone은 모든 패키지 함수의 기본 타임존입니다
    DefaultTimezone = "Asia/Seoul"
    DefaultLocation = "Asia/Seoul" // KST, GMT+9
)

var (
    // KST is the pre-loaded Korea Standard Time location
    // KST는 사전 로드된 한국 표준시 위치입니다
    KST *time.Location

    // defaultLocation is the current default timezone
    // defaultLocation은 현재 기본 타임존입니다
    defaultLocation *time.Location
)

// init() loads KST timezone on package initialization
// init()은 패키지 초기화 시 KST 타임존을 로드합니다
func init() {
    var err error
    KST, err = time.LoadLocation(DefaultTimezone)
    if err != nil {
        // Fallback to UTC if KST cannot be loaded
        // KST를 로드할 수 없는 경우 UTC로 폴백
        KST = time.UTC
    }
    defaultLocation = KST
}
```

**Design Rationale / 설계 근거**:
- Pre-load KST in init() for performance / 성능을 위해 init()에서 KST 사전 로드
- Provide package-wide KST constant for convenience / 편의를 위해 패키지 전체 KST 상수 제공
- Fallback to UTC if timezone loading fails / 타임존 로드 실패 시 UTC로 폴백

### 3. Timezone Cache / 타임존 캐시

Located in `timezone.go`:

`timezone.go`에 위치:

```go
var (
    // timezoneCache caches loaded timezones for performance
    // timezoneCache는 성능을 위해 로드된 타임존을 캐시합니다
    timezoneCache   = make(map[string]*time.Location)
    timezoneCacheMu sync.RWMutex
)

// loadTimezone loads a timezone with caching
// loadTimezone은 캐싱과 함께 타임존을 로드합니다
func loadTimezone(tz string) (*time.Location, error) {
    // Check cache first (read lock)
    // 먼저 캐시 확인 (읽기 잠금)
    timezoneCacheMu.RLock()
    if loc, ok := timezoneCache[tz]; ok {
        timezoneCacheMu.RUnlock()
        return loc, nil
    }
    timezoneCacheMu.RUnlock()

    // Load timezone (write lock)
    // 타임존 로드 (쓰기 잠금)
    timezoneCacheMu.Lock()
    defer timezoneCacheMu.Unlock()

    // Double-check after acquiring write lock
    // 쓰기 잠금 획득 후 다시 확인
    if loc, ok := timezoneCache[tz]; ok {
        return loc, nil
    }

    // Load and cache
    // 로드 및 캐시
    loc, err := time.LoadLocation(tz)
    if err != nil {
        return nil, err
    }
    timezoneCache[tz] = loc
    return loc, nil
}
```

**Design Rationale / 설계 근거**:
- Cache loaded timezones to avoid repeated disk I/O / 반복된 디스크 I/O를 피하기 위해 로드된 타임존 캐시
- Use RWMutex for thread-safe concurrent access / 스레드 안전 동시 액세스를 위해 RWMutex 사용
- Double-check locking pattern to prevent race conditions / 경쟁 조건을 방지하기 위한 이중 확인 잠금 패턴

### 4. Holiday Management / 공휴일 관리

Located in `business.go`:

`business.go`에 위치:

```go
var (
    // holidays is a thread-safe map of registered holidays
    // holidays는 등록된 공휴일의 스레드 안전 맵입니다
    holidays   = make(map[string]bool)
    holidaysMu sync.RWMutex
)

// holidayKey generates a unique key for a date
// holidayKey는 날짜에 대한 고유 키를 생성합니다
func holidayKey(t time.Time) string {
    return t.Format("2006-01-02")
}

// SetHolidays sets the list of holidays
// SetHolidays는 공휴일 목록을 설정합니다
func SetHolidays(dates []time.Time) {
    holidaysMu.Lock()
    defer holidaysMu.Unlock()

    holidays = make(map[string]bool)
    for _, date := range dates {
        key := holidayKey(date)
        holidays[key] = true
    }
}

// IsHoliday checks if a date is a holiday
// IsHoliday는 날짜가 공휴일인지 확인합니다
func IsHoliday(t time.Time) bool {
    holidaysMu.RLock()
    defer holidaysMu.RUnlock()

    key := holidayKey(t)
    return holidays[key]
}
```

**Design Rationale / 설계 근거**:
- Use map for O(1) holiday lookup / O(1) 공휴일 조회를 위해 맵 사용
- Thread-safe with RWMutex / RWMutex로 스레드 안전
- Date-only key format (YYYY-MM-DD) / 날짜만 있는 키 포맷 (YYYY-MM-DD)

### 5. Custom Format Tokens / 커스텀 포맷 토큰

Located in `format.go`:

`format.go`에 위치:

```go
var customFormatTokens = map[string]string{
    "YYYY": "2006",  // 4-digit year / 4자리 연도
    "YY":   "06",    // 2-digit year / 2자리 연도
    "MM":   "01",    // Month (01-12) / 월 (01-12)
    "DD":   "02",    // Day (01-31) / 일 (01-31)
    "HH":   "15",    // Hour (00-23) / 시 (00-23)
    "mm":   "04",    // Minute (00-59) / 분 (00-59)
    "ss":   "05",    // Second (00-59) / 초 (00-59)
}

// convertFormatTokens converts custom tokens to Go's format
// convertFormatTokens는 커스텀 토큰을 Go의 포맷으로 변환합니다
func convertFormatTokens(layout string) string {
    goLayout := layout
    for token, goToken := range customFormatTokens {
        goLayout = strings.ReplaceAll(goLayout, token, goToken)
    }
    return goLayout
}

// Format formats time using custom tokens
// Format은 커스텀 토큰을 사용하여 시간을 포맷합니다
func Format(t time.Time, layout string) string {
    goLayout := convertFormatTokens(layout)
    return t.In(defaultLocation).Format(goLayout)
}
```

**Design Rationale / 설계 근거**:
- Intuitive tokens that don't require memorization / 암기가 필요 없는 직관적인 토큰
- Simple string replacement for conversion / 변환을 위한 간단한 문자열 치환
- Converts to Go's standard format under the hood / 내부적으로 Go의 표준 포맷으로 변환

---

## Design Patterns / 디자인 패턴

### 1. Singleton Pattern / 싱글톤 패턴

**Used in / 사용 위치**: Package-level constants and caches

**Implementation / 구현**:

```go
var (
    KST             *time.Location        // Global constant
    defaultLocation *time.Location        // Global state
    timezoneCache   map[string]*time.Location  // Shared cache
    holidays        map[string]bool       // Shared state
)

func init() {
    // Initialize once on package load
    // 패키지 로드 시 한 번만 초기화
    KST, _ = time.LoadLocation("Asia/Seoul")
    defaultLocation = KST
}
```

**Benefits / 이점**:
- Single source of truth for shared state / 공유 상태의 단일 소스
- Initialized once for efficiency / 효율성을 위해 한 번만 초기화
- Easy global access / 쉬운 전역 액세스

### 2. Factory Pattern / 팩토리 패턴

**Used in / 사용 위치**: Time parsing functions

**Implementation / 구현**:

```go
// Parse automatically detects format and creates time.Time
// Parse는 자동으로 포맷을 감지하고 time.Time을 생성합니다
func Parse(s string) (time.Time, error) {
    // Try different formats in order
    // 순서대로 다른 포맷 시도
    formats := []string{
        time.RFC3339,
        "2006-01-02 15:04:05",
        "2006-01-02",
    }

    for _, format := range formats {
        if t, err := time.Parse(format, s); err == nil {
            return t.In(defaultLocation), nil
        }
    }

    return time.Time{}, fmt.Errorf("unable to parse time: %s", s)
}
```

**Benefits / 이점**:
- Encapsulates complex creation logic / 복잡한 생성 로직 캡슐화
- Automatic format detection / 자동 포맷 감지
- Single interface for multiple formats / 여러 포맷에 대한 단일 인터페이스

### 3. Strategy Pattern / 전략 패턴

**Used in / 사용 위치**: Format conversion

**Implementation / 구현**:

```go
// Different format strategies / 다른 포맷 전략
func FormatISO8601(t time.Time) string {
    return Format(t, "YYYY-MM-DD")
}

func FormatRFC3339(t time.Time) string {
    return t.In(defaultLocation).Format(time.RFC3339)
}

func FormatKorean(t time.Time) string {
    return Format(t, "YYYY년 MM월 DD일 HH시 mm분 ss초")
}

func FormatCustom(t time.Time, layout string) string {
    return t.In(defaultLocation).Format(layout)
}
```

**Benefits / 이점**:
- Interchangeable formatting strategies / 교체 가능한 포맷팅 전략
- Easy to add new formats / 새 포맷 추가 용이
- Clear separation of concerns / 명확한 관심사 분리

### 4. Decorator Pattern / 데코레이터 패턴

**Used in / 사용 위치**: TimeDiff type methods

**Implementation / 구현**:

```go
type TimeDiff struct {
    Duration time.Duration  // Wraps standard Duration
}

// Decorates Duration with additional methods
// Duration을 추가 메서드로 데코레이트
func (td *TimeDiff) Days() float64 {
    return td.Hours() / 24
}

func (td *TimeDiff) Humanize() string {
    // Enhanced human-readable format
    // 향상된 사람이 읽기 쉬운 포맷
    // ...
}
```

**Benefits / 이점**:
- Extends standard Duration without modification / 수정 없이 표준 Duration 확장
- Adds human-readable methods / 사람이 읽기 쉬운 메서드 추가
- Maintains compatibility with Duration / Duration과의 호환성 유지

### 5. Cache-Aside Pattern / 캐시-어사이드 패턴

**Used in / 사용 위치**: Timezone loading

**Implementation / 구현**:

```go
func loadTimezone(tz string) (*time.Location, error) {
    // 1. Check cache first / 먼저 캐시 확인
    timezoneCacheMu.RLock()
    if loc, ok := timezoneCache[tz]; ok {
        timezoneCacheMu.RUnlock()
        return loc, nil
    }
    timezoneCacheMu.RUnlock()

    // 2. Load from source (disk) / 소스(디스크)에서 로드
    timezoneCacheMu.Lock()
    defer timezoneCacheMu.Unlock()

    // 3. Double-check and cache / 다시 확인 및 캐시
    if loc, ok := timezoneCache[tz]; ok {
        return loc, nil
    }

    loc, err := time.LoadLocation(tz)
    if err != nil {
        return nil, err
    }
    timezoneCache[tz] = loc
    return loc, nil
}
```

**Benefits / 이점**:
- Improves performance by caching / 캐싱으로 성능 향상
- Thread-safe with double-check locking / 이중 확인 잠금으로 스레드 안전
- Reduces disk I/O / 디스크 I/O 감소

---

## Internal Implementation / 내부 구현

### 1. Time Difference Calculation / 시간 차이 계산

**File / 파일**: `diff.go`

**Core Algorithm / 핵심 알고리즘**:

```go
func SubTime(t1, t2 time.Time) *TimeDiff {
    return &TimeDiff{
        Duration: t2.Sub(t1),
    }
}

func (td *TimeDiff) String() string {
    d := td.Duration
    if d < 0 {
        d = -d
    }

    days := int(d.Hours() / 24)
    hours := int(d.Hours()) % 24
    minutes := int(d.Minutes()) % 60
    seconds := int(d.Seconds()) % 60

    parts := []string{}
    if days > 0 {
        parts = append(parts, fmt.Sprintf("%d day%s", days, plural(days)))
    }
    if hours > 0 {
        parts = append(parts, fmt.Sprintf("%d hour%s", hours, plural(hours)))
    }
    if minutes > 0 {
        parts = append(parts, fmt.Sprintf("%d minute%s", minutes, plural(minutes)))
    }
    if seconds > 0 || len(parts) == 0 {
        parts = append(parts, fmt.Sprintf("%d second%s", seconds, plural(seconds)))
    }

    return strings.Join(parts, " ")
}

func plural(n int) string {
    if n == 1 {
        return ""
    }
    return "s"
}
```

**Key Points / 핵심 포인트**:
- Uses standard time.Duration for calculation / 계산을 위해 표준 time.Duration 사용
- Breaks down duration into days, hours, minutes, seconds / 기간을 일, 시간, 분, 초로 분해
- Automatically handles plural forms / 자동으로 복수형 처리

### 2. Timezone Conversion / 타임존 변환

**File / 파일**: `timezone.go`

**Core Algorithm / 핵심 알고리즘**:

```go
func ConvertTimezone(t time.Time, tz string) (time.Time, error) {
    loc, err := loadTimezone(tz)
    if err != nil {
        return time.Time{}, fmt.Errorf("invalid timezone %s: %w", tz, err)
    }
    return t.In(loc), nil
}

func SetDefaultTimezone(tz string) error {
    loc, err := loadTimezone(tz)
    if err != nil {
        return fmt.Errorf("invalid timezone %s: %w", tz, err)
    }

    timezoneCacheMu.Lock()
    defer timezoneCacheMu.Unlock()

    defaultLocation = loc
    return nil
}
```

**Key Points / 핵심 포인트**:
- Validates timezone before conversion / 변환 전 타임존 검증
- Uses cached locations for performance / 성능을 위해 캐시된 위치 사용
- Thread-safe default timezone update / 스레드 안전 기본 타임존 업데이트

### 3. Business Day Calculation / 영업일 계산

**File / 파일**: `business.go`

**Core Algorithm / 핵심 알고리즘**:

```go
func IsBusinessDay(t time.Time) bool {
    // Check if weekend / 주말인지 확인
    weekday := t.Weekday()
    if weekday == time.Saturday || weekday == time.Sunday {
        return false
    }

    // Check if holiday / 공휴일인지 확인
    if IsHoliday(t) {
        return false
    }

    return true
}

func AddBusinessDays(t time.Time, days int) time.Time {
    if days == 0 {
        return t
    }

    // Direction / 방향
    step := 1
    if days < 0 {
        step = -1
        days = -days
    }

    result := t
    count := 0

    for count < days {
        result = result.AddDate(0, 0, step)
        if IsBusinessDay(result) {
            count++
        }
    }

    return result
}

func AddKoreanHolidays(year int) {
    koreanHolidays := []struct {
        month int
        day   int
    }{
        {1, 1},   // New Year's Day / 신정
        {3, 1},   // Independence Movement Day / 삼일절
        {5, 5},   // Children's Day / 어린이날
        {6, 6},   // Memorial Day / 현충일
        {8, 15},  // Liberation Day / 광복절
        {10, 3},  // National Foundation Day / 개천절
        {10, 9},  // Hangeul Day / 한글날
        {12, 25}, // Christmas / 성탄절
    }

    holidaysMu.Lock()
    defer holidaysMu.Unlock()

    for _, h := range koreanHolidays {
        date := time.Date(year, time.Month(h.month), h.day, 0, 0, 0, 0, KST)
        key := holidayKey(date)
        holidays[key] = true
    }
}
```

**Key Points / 핵심 포인트**:
- Checks both weekends and holidays / 주말 및 공휴일 모두 확인
- Iterates forward/backward to skip non-business days / 비영업일을 건너뛰기 위해 앞/뒤로 반복
- Korean holidays hard-coded for convenience / 편의를 위해 한국 공휴일 하드코딩

### 4. Relative Time String Generation / 상대 시간 문자열 생성

**File / 파일**: `relative.go`

**Core Algorithm / 핵심 알고리즘**:

```go
func RelativeTime(t time.Time) string {
    now := time.Now().In(defaultLocation)
    t = t.In(defaultLocation)
    duration := now.Sub(t)

    isPast := duration > 0
    if !isPast {
        duration = -duration
    }

    seconds := int(duration.Seconds())
    minutes := int(duration.Minutes())
    hours := int(duration.Hours())
    days := int(duration.Hours() / 24)
    weeks := days / 7
    months := days / 30
    years := days / 365

    var timeStr string
    switch {
    case seconds < 60:
        timeStr = fmt.Sprintf("%d second%s", seconds, plural(seconds))
    case minutes < 60:
        timeStr = fmt.Sprintf("%d minute%s", minutes, plural(minutes))
    case hours < 24:
        timeStr = fmt.Sprintf("%d hour%s", hours, plural(hours))
    case days < 7:
        timeStr = fmt.Sprintf("%d day%s", days, plural(days))
    case weeks < 4:
        timeStr = fmt.Sprintf("%d week%s", weeks, plural(weeks))
    case months < 12:
        timeStr = fmt.Sprintf("%d month%s", months, plural(months))
    default:
        timeStr = fmt.Sprintf("%d year%s", years, plural(years))
    }

    if isPast {
        return timeStr + " ago"
    }
    return "in " + timeStr
}
```

**Key Points / 핵심 포인트**:
- Determines if time is past or future / 시간이 과거인지 미래인지 판단
- Selects appropriate unit based on magnitude / 크기에 따라 적절한 단위 선택
- Handles both "ago" and "in" prefixes / "ago" 및 "in" 접두사 모두 처리

### 5. Age Calculation / 나이 계산

**File / 파일**: `age.go`

**Core Algorithm / 핵심 알고리즘**:

```go
func Age(birthDate time.Time) *AgeDetail {
    now := time.Now().In(defaultLocation)
    birthDate = birthDate.In(defaultLocation)

    years := now.Year() - birthDate.Year()
    months := int(now.Month()) - int(birthDate.Month())
    days := now.Day() - birthDate.Day()

    // Adjust for negative days / 음수 일 조정
    if days < 0 {
        months--
        // Get days in previous month / 이전 월의 일수 가져오기
        prevMonth := now.AddDate(0, -1, 0)
        days += daysInMonth(prevMonth.Year(), prevMonth.Month())
    }

    // Adjust for negative months / 음수 월 조정
    if months < 0 {
        years--
        months += 12
    }

    return &AgeDetail{
        Years:  years,
        Months: months,
        Days:   days,
    }
}

func daysInMonth(year int, month time.Month) int {
    // First day of next month minus one day
    // 다음 달 첫날에서 하루 빼기
    return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
```

**Key Points / 핵심 포인트**:
- Calculates exact years, months, and days / 정확한 년, 월, 일 계산
- Handles month/day overflow correctly / 월/일 오버플로우 올바르게 처리
- Accounts for varying days in months / 월별 다른 일수 고려

---

## Adding New Features / 새 기능 추가

### Step-by-Step Guide / 단계별 가이드

#### 1. Choose the Right File / 올바른 파일 선택

Based on the function category:

함수 카테고리에 따라:

- Time difference → `diff.go`
- Timezone operations → `timezone.go`
- Date arithmetic → `arithmetic.go`
- Formatting → `format.go`
- Parsing → `parse.go`
- Comparisons → `comparison.go`
- Age → `age.go`
- Relative time → `relative.go`
- Unix timestamps → `unix.go`
- Business days → `business.go`

#### 2. Write the Function / 함수 작성

**Template / 템플릿**:

```go
// FunctionName does something with time
// FunctionName은 시간으로 무언가를 합니다
//
// Parameters / 매개변수:
//   - t: the time to process / 처리할 시간
//
// Returns / 반환:
//   - result: the processed result / 처리된 결과
//
// Example / 예제:
//
//     t := time.Now()
//     result := timeutil.FunctionName(t)
//     fmt.Println(result)
//
func FunctionName(t time.Time) ReturnType {
    // 1. Validate input / 입력 검증
    // 2. Apply default timezone if needed / 필요시 기본 타임존 적용
    t = t.In(defaultLocation)

    // 3. Perform calculation / 계산 수행
    // ...

    // 4. Return result / 결과 반환
    return result
}
```

**Best Practices / 모범 사례**:
- Always use bilingual comments (English/Korean) / 항상 이중 언어 주석 사용 (영문/한글)
- Apply default timezone consistently / 기본 타임존을 일관되게 적용
- Handle edge cases / 엣지 케이스 처리
- Return errors when appropriate / 적절한 경우 에러 반환

#### 3. Add Tests / 테스트 추가

**Template / 템플릿**:

```go
// TestFunctionName tests the FunctionName function
// TestFunctionName은 FunctionName 함수를 테스트합니다
func TestFunctionName(t *testing.T) {
    // Test cases / 테스트 케이스
    tests := []struct {
        name     string
        input    time.Time
        expected ReturnType
    }{
        {
            name:     "basic case",
            input:    time.Date(2025, 10, 14, 15, 0, 0, 0, KST),
            expected: expectedValue,
        },
        // Add more test cases / 더 많은 테스트 케이스 추가
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := FunctionName(tt.input)
            if result != tt.expected {
                t.Errorf("FunctionName() = %v, want %v", result, tt.expected)
            }
        })
    }
}
```

#### 4. Add Benchmarks / 벤치마크 추가

```go
// BenchmarkFunctionName benchmarks the FunctionName function
// BenchmarkFunctionName은 FunctionName 함수를 벤치마크합니다
func BenchmarkFunctionName(b *testing.B) {
    t := time.Now()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        FunctionName(t)
    }
}
```

#### 5. Update Documentation / 문서 업데이트

Update the following files:

다음 파일 업데이트:

- `README.md`: Add function to appropriate category / 적절한 카테고리에 함수 추가
- `USER_MANUAL.md`: Add detailed usage examples / 상세한 사용 예제 추가
- `DEVELOPER_GUIDE.md`: Add implementation notes / 구현 노트 추가

#### 6. Example Implementation / 예제 구현

Let's add a new function `IsWorkday()`:

새 함수 `IsWorkday()`를 추가해봅시다:

```go
// File: comparison.go

// IsWorkday checks if the time is on a workday (alias for IsWeekday)
// IsWorkday는 시간이 평일인지 확인합니다 (IsWeekday의 별칭)
//
// A workday is Monday through Friday, regardless of holidays.
// Use IsBusinessDay() if you want to exclude holidays.
//
// 평일은 월요일부터 금요일까지이며 공휴일과 무관합니다.
// 공휴일을 제외하려면 IsBusinessDay()를 사용하세요.
//
// Parameters / 매개변수:
//   - t: the time to check / 확인할 시간
//
// Returns / 반환:
//   - true if workday, false otherwise / 평일이면 true, 아니면 false
//
// Example / 예제:
//
//     now := time.Now()
//     if timeutil.IsWorkday(now) {
//         fmt.Println("It's a workday")
//     }
//
func IsWorkday(t time.Time) bool {
    return IsWeekday(t)
}
```

```go
// File: timeutil_test.go

func TestIsWorkday(t *testing.T) {
    tests := []struct {
        name     string
        time     time.Time
        expected bool
    }{
        {
            name:     "Monday is workday",
            time:     time.Date(2025, 10, 13, 0, 0, 0, 0, KST), // Monday
            expected: true,
        },
        {
            name:     "Saturday is not workday",
            time:     time.Date(2025, 10, 18, 0, 0, 0, 0, KST), // Saturday
            expected: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := IsWorkday(tt.time)
            if result != tt.expected {
                t.Errorf("IsWorkday() = %v, want %v", result, tt.expected)
            }
        })
    }
}
```

---

## Testing Guide / 테스트 가이드

### Test Organization / 테스트 구성

All tests are in `timeutil_test.go`, organized by category:

모든 테스트는 `timeutil_test.go`에 있으며 카테고리별로 구성됩니다:

```go
// Test functions / 테스트 함수
func TestTimeDiff(t *testing.T)        // Time difference tests
func TestTimezone(t *testing.T)        // Timezone tests
func TestArithmetic(t *testing.T)      // Arithmetic tests
func TestFormat(t *testing.T)          // Format tests
func TestParse(t *testing.T)           // Parse tests
func TestComparison(t *testing.T)      // Comparison tests
func TestAge(t *testing.T)             // Age tests
func TestRelativeTime(t *testing.T)    // Relative time tests
func TestUnix(t *testing.T)            // Unix timestamp tests
func TestBusinessDays(t *testing.T)    // Business days tests

// Benchmark functions / 벤치마크 함수
func BenchmarkFormat(b *testing.B)
func BenchmarkParse(b *testing.B)
func BenchmarkTimeDiff(b *testing.B)
```

### Running Tests / 테스트 실행

```bash
# Run all tests / 모든 테스트 실행
go test ./timeutil -v

# Run specific test / 특정 테스트 실행
go test ./timeutil -v -run TestTimeDiff

# Run with coverage / 커버리지와 함께 실행
go test ./timeutil -cover
go test ./timeutil -coverprofile=coverage.out
go tool cover -html=coverage.out

# Run benchmarks / 벤치마크 실행
go test ./timeutil -bench=.
go test ./timeutil -bench=BenchmarkFormat
```

### Test Coverage Goals / 테스트 커버리지 목표

- **Minimum / 최소**: 70% statement coverage / 70% 문 커버리지
- **Target / 목표**: 85% statement coverage / 85% 문 커버리지
- **Ideal / 이상적**: 95% statement coverage / 95% 문 커버리지

**Current Coverage / 현재 커버리지**: 31.4% (needs improvement / 개선 필요)

### Test Best Practices / 테스트 모범 사례

1. **Use Table-Driven Tests / 테이블 기반 테스트 사용**

```go
func TestFormat(t *testing.T) {
    tests := []struct {
        name     string
        time     time.Time
        layout   string
        expected string
    }{
        {"ISO8601", time.Date(2025, 10, 14, 0, 0, 0, 0, KST), "YYYY-MM-DD", "2025-10-14"},
        {"DateTime", time.Date(2025, 10, 14, 15, 4, 5, 0, KST), "YYYY-MM-DD HH:mm:ss", "2025-10-14 15:04:05"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Format(tt.time, tt.layout)
            if result != tt.expected {
                t.Errorf("Format() = %v, want %v", result, tt.expected)
            }
        })
    }
}
```

2. **Test Edge Cases / 엣지 케이스 테스트**

```go
func TestDiffInDays_EdgeCases(t *testing.T) {
    // Same time / 같은 시간
    t1 := time.Now()
    if diff := DiffInDays(t1, t1); diff != 0 {
        t.Errorf("Same time should have 0 days difference")
    }

    // Negative duration / 음수 기간
    t2 := t1.Add(24 * time.Hour)
    if diff := DiffInDays(t2, t1); diff >= 0 {
        t.Errorf("Negative duration should be negative")
    }
}
```

3. **Use Subtests / 서브테스트 사용**

```go
func TestTimezone(t *testing.T) {
    t.Run("ConvertTimezone", func(t *testing.T) {
        // Test timezone conversion
    })

    t.Run("InvalidTimezone", func(t *testing.T) {
        // Test error handling
    })

    t.Run("ToKST", func(t *testing.T) {
        // Test KST conversion
    })
}
```

4. **Clean Up Resources / 리소스 정리**

```go
func TestSetDefaultTimezone(t *testing.T) {
    // Save original / 원본 저장
    original := GetDefaultTimezone()

    // Restore after test / 테스트 후 복원
    defer SetDefaultTimezone(original)

    // Test / 테스트
    err := SetDefaultTimezone("UTC")
    if err != nil {
        t.Fatalf("SetDefaultTimezone() error = %v", err)
    }
}
```

---

## Performance / 성능

### Optimization Strategies / 최적화 전략

#### 1. Timezone Caching / 타임존 캐싱

**Problem / 문제**: `time.LoadLocation()` is slow (disk I/O)

**Solution / 해결책**: Cache loaded timezones in memory

**Impact / 영향**: ~10x faster for repeated timezone conversions

**Implementation / 구현**:

```go
var (
    timezoneCache   = make(map[string]*time.Location)
    timezoneCacheMu sync.RWMutex
)
```

#### 2. Pre-loaded KST / 사전 로드된 KST

**Problem / 문제**: KST is used frequently as default

**Solution / 해결책**: Load KST once in init()

**Impact / 영향**: Eliminates repeated loads for default timezone

**Implementation / 구현**:

```go
func init() {
    KST, _ = time.LoadLocation("Asia/Seoul")
    defaultLocation = KST
}
```

#### 3. Holiday Map Lookup / 공휴일 맵 조회

**Problem / 문제**: Checking holidays in a slice is O(n)

**Solution / 해결책**: Use map for O(1) lookup

**Impact / 영향**: ~100x faster for large holiday lists

**Implementation / 구현**:

```go
holidays = make(map[string]bool)  // O(1) lookup instead of O(n)
```

#### 4. Avoid Unnecessary Conversions / 불필요한 변환 방지

**Problem / 문제**: Repeated timezone conversions are expensive

**Solution / 해결책**: Convert once and reuse

**Example / 예제**:

```go
// ❌ Bad: Multiple conversions
// ❌ 나쁨: 여러 번 변환
for i := 0; i < 1000; i++ {
    t := time.Now().In(KST)
    // Use t
}

// ✅ Good: Convert once
// ✅ 좋음: 한 번만 변환
t := time.Now().In(KST)
for i := 0; i < 1000; i++ {
    // Use t
}
```

### Benchmark Results / 벤치마크 결과

```bash
go test ./timeutil -bench=.
```

**Expected Results / 예상 결과**:

```
BenchmarkFormat-8           5000000    250 ns/op    48 B/op    2 allocs/op
BenchmarkParse-8            2000000    600 ns/op    96 B/op    4 allocs/op
BenchmarkTimeDiff-8        10000000    100 ns/op    24 B/op    1 allocs/op
BenchmarkConvertTimezone-8  3000000    400 ns/op    32 B/op    1 allocs/op
BenchmarkIsBusinessDay-8   20000000     50 ns/op     0 B/op    0 allocs/op
```

### Performance Tips / 성능 팁

1. **Reuse time.Time Objects / time.Time 객체 재사용**

```go
// ✅ Good / 좋음
now := time.Now()
for i := 0; i < 1000; i++ {
    tomorrow := AddDays(now, 1)
}
```

2. **Cache Formatted Strings / 포맷된 문자열 캐시**

```go
// If formatting the same time repeatedly / 같은 시간을 반복적으로 포맷하는 경우
formatted := Format(now, "YYYY-MM-DD")
// Use 'formatted' multiple times
```

3. **Use Appropriate Precision / 적절한 정밀도 사용**

```go
// Use Unix seconds for most cases / 대부분의 경우 Unix 초 사용
timestamp := Now()  // Faster than NowNano()
```

---

## Contributing Guidelines / 기여 가이드라인

### Contribution Process / 기여 프로세스

1. **Fork the Repository / 저장소 포크**
   ```bash
   git clone https://github.com/your-username/go-utils.git
   ```

2. **Create Feature Branch / 기능 브랜치 생성**
   ```bash
   git checkout -b feature/add-new-function
   ```

3. **Make Changes / 변경**
   - Add function with bilingual comments / 이중 언어 주석과 함께 함수 추가
   - Add tests with 80%+ coverage / 80% 이상 커버리지의 테스트 추가
   - Update documentation / 문서 업데이트

4. **Run Tests / 테스트 실행**
   ```bash
   go test ./timeutil -v
   go test ./timeutil -cover
   ```

5. **Commit Changes / 변경사항 커밋**
   ```bash
   git add .
   git commit -m "Feat: Add IsWorkday function"
   ```

6. **Push and Create PR / 푸시 및 PR 생성**
   ```bash
   git push origin feature/add-new-function
   ```

### Contribution Checklist / 기여 체크리스트

- [ ] Function has bilingual comments (English/Korean) / 함수에 이중 언어 주석 있음 (영문/한글)
- [ ] Tests added with 80%+ coverage / 80% 이상 커버리지의 테스트 추가됨
- [ ] Benchmarks added / 벤치마크 추가됨
- [ ] Documentation updated (README, USER_MANUAL, DEVELOPER_GUIDE) / 문서 업데이트됨
- [ ] All tests pass / 모든 테스트 통과
- [ ] Code follows style guide / 코드가 스타일 가이드 따름
- [ ] No external dependencies added / 외부 의존성 추가 안됨

---

## Code Style / 코드 스타일

### Naming Conventions / 명명 규칙

1. **Function Names / 함수 이름**
   - Use clear, descriptive names / 명확하고 설명적인 이름 사용
   - Start with verb (Add, Get, Is, Format, etc.) / 동사로 시작 (Add, Get, Is, Format 등)
   - Use PascalCase for exported functions / 내보낸 함수에 PascalCase 사용

```go
// ✅ Good / 좋음
func AddDays(t time.Time, days int) time.Time
func IsBusinessDay(t time.Time) bool
func FormatISO8601(t time.Time) string

// ❌ Bad / 나쁨
func add(t time.Time, d int) time.Time  // Not descriptive
func businessday(t time.Time) bool      // Not clear
func iso8601(t time.Time) string        // Missing verb
```

2. **Variable Names / 변수 이름**
   - Short names for local variables / 로컬 변수에 짧은 이름
   - Descriptive names for package-level variables / 패키지 레벨 변수에 설명적인 이름

```go
// ✅ Good / 좋음
func AddDays(t time.Time, days int) time.Time {
    return t.AddDate(0, 0, days)
}

var defaultLocation *time.Location  // Package-level

// ❌ Bad / 나쁨
func AddDays(timeValue time.Time, numberOfDays int) time.Time {
    return timeValue.AddDate(0, 0, numberOfDays)
}
```

3. **Constants / 상수**
   - Use PascalCase for exported constants / 내보낸 상수에 PascalCase 사용
   - Use SCREAMING_SNAKE_CASE for internal constants / 내부 상수에 SCREAMING_SNAKE_CASE 사용

```go
// Exported / 내보낸
const DefaultTimezone = "Asia/Seoul"
const DefaultLocation = "Asia/Seoul"

// Internal / 내부
const (
    secondsPerMinute = 60
    minutesPerHour   = 60
    hoursPerDay      = 24
)
```

### Comment Style / 주석 스타일

1. **Package Comments / 패키지 주석**

```go
// Package timeutil provides time and date utility functions with extreme simplicity.
// It reduces 20+ lines of repetitive time code to just 1-2 lines.
//
// 패키지 timeutil은 극도의 간결함으로 시간 및 날짜 유틸리티 함수를 제공합니다.
// 20줄 이상의 반복적인 시간 코드를 단 1-2줄로 줄입니다.
//
// Example usage:
//
//     import "github.com/arkd0ng/go-utils/timeutil"
//
//     // Time difference / 시간 차이
//     diff := timeutil.SubTime(start, end)
//     fmt.Println(diff.String())  // "2 days 6 hours"
//
package timeutil
```

2. **Function Comments / 함수 주석**

```go
// AddDays adds the specified number of days to a time.
// AddDays는 시간에 지정된 일 수를 더합니다.
//
// Negative values subtract days.
// 음수 값은 일을 뺍니다.
//
// Parameters / 매개변수:
//   - t: the time to add days to / 일을 더할 시간
//   - days: number of days to add (can be negative) / 더할 일 수 (음수 가능)
//
// Returns / 반환:
//   - the resulting time / 결과 시간
//
// Example / 예제:
//
//     now := time.Now()
//     tomorrow := timeutil.AddDays(now, 1)
//     yesterday := timeutil.AddDays(now, -1)
//
func AddDays(t time.Time, days int) time.Time {
    return t.AddDate(0, 0, days)
}
```

3. **Inline Comments / 인라인 주석**

```go
func Format(t time.Time, layout string) string {
    // Convert custom tokens to Go format / 커스텀 토큰을 Go 포맷으로 변환
    goLayout := convertFormatTokens(layout)

    // Apply default timezone / 기본 타임존 적용
    t = t.In(defaultLocation)

    return t.Format(goLayout)
}
```

### Error Handling / 에러 처리

1. **Return Errors / 에러 반환**

```go
// ✅ Good: Return error / 좋음: 에러 반환
func ConvertTimezone(t time.Time, tz string) (time.Time, error) {
    loc, err := loadTimezone(tz)
    if err != nil {
        return time.Time{}, fmt.Errorf("invalid timezone %s: %w", tz, err)
    }
    return t.In(loc), nil
}

// ❌ Bad: Panic on error / 나쁨: 에러 시 패닉
func ConvertTimezone(t time.Time, tz string) time.Time {
    loc, err := loadTimezone(tz)
    if err != nil {
        panic(err)  // Don't panic!
    }
    return t.In(loc)
}
```

2. **Wrap Errors / 에러 래핑**

```go
// ✅ Good: Wrap error with context / 좋음: 컨텍스트와 함께 에러 래핑
if err != nil {
    return time.Time{}, fmt.Errorf("failed to parse time %s: %w", s, err)
}

// ❌ Bad: Lose error context / 나쁨: 에러 컨텍스트 손실
if err != nil {
    return time.Time{}, err
}
```

### Code Formatting / 코드 포맷팅

Use `gofmt` and `goimports`:

`gofmt` 및 `goimports` 사용:

```bash
# Format code / 코드 포맷
gofmt -w .

# Organize imports / import 정리
goimports -w .
```

---

## Appendix / 부록

### Useful Resources / 유용한 리소스

- **Go Time Package**: https://pkg.go.dev/time
- **Go Testing**: https://pkg.go.dev/testing
- **Go Benchmarking**: https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go

### Common Issues / 일반적인 문제

1. **Timezone Data Not Found / 타임존 데이터를 찾을 수 없음**
   - Install tzdata package / tzdata 패키지 설치
   - On Windows, may need to bundle timezone data / Windows에서는 타임존 데이터 번들 필요

2. **Race Conditions / 경쟁 조건**
   - Always use mutex for shared state / 공유 상태에 항상 뮤텍스 사용
   - Run `go test -race` to detect / `go test -race`로 감지

3. **Memory Leaks / 메모리 누수**
   - Clean up goroutines / 고루틴 정리
   - Clear large maps when done / 완료 시 큰 맵 정리

---

## Conclusion / 결론

The `timeutil` package is designed with extreme simplicity, performance, and usability in mind. By following this developer guide, you can understand the internal architecture, contribute new features, and maintain the package effectively.

`timeutil` 패키지는 극도의 간결함, 성능 및 사용성을 염두에 두고 설계되었습니다. 이 개발자 가이드를 따르면 내부 아키텍처를 이해하고, 새로운 기능을 기여하고, 패키지를 효과적으로 유지 관리할 수 있습니다.

For questions or suggestions, please open an issue on GitHub:
https://github.com/arkd0ng/go-utils/issues

질문이나 제안 사항이 있으면 GitHub에서 이슈를 열어주세요:
https://github.com/arkd0ng/go-utils/issues

---

**Version / 버전**: v1.6.002
**Last Updated / 마지막 업데이트**: 2025-10-14
