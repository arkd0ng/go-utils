package timeutil

import "time"

// ============================================================================
// FILE OVERVIEW / 파일 개요
// ============================================================================
//
// Package: timeutil/month.go
// Purpose: Month name and quarter operations
//          월 이름 및 분기 연산
//
// This file provides operations for extracting month names in different formats
// and determining the quarter of the year. It supports both Korean and English
// month names, with full and abbreviated forms for English. The quarter function
// divides the year into four 3-month periods following standard business quarters.
// These operations are useful for reports, calendars, localized date displays,
// and financial/business period calculations.
//
// 이 파일은 다양한 형식으로 월 이름을 추출하고 연도의 분기를 결정하기 위한
// 연산을 제공합니다. 한국어 및 영어 월 이름을 지원하며, 영어의 경우 전체 형식과
// 약어 형식을 제공합니다. quarter 함수는 표준 비즈니스 분기를 따라 연도를
// 4개의 3개월 기간으로 나눕니다. 이러한 연산은 보고서, 달력, 현지화된 날짜
// 표시 및 재무/비즈니스 기간 계산에 유용합니다.
//
// ============================================================================
// KEY FEATURES / 주요 기능
// ============================================================================
//
// 1. KOREAN MONTH NAME (한국어 월 이름)
//    - MonthKorean: Returns "N월" format
//      "N월" 형식 반환
//    - Examples: "1월", "10월", "12월"
//      예시: "1월", "10월", "12월"
//    - Native Korean representation
//      네이티브 한국어 표현
//
// 2. ENGLISH MONTH NAME (영어 월 이름)
//    - MonthName: Full English name
//      전체 영어 이름
//    - Examples: "January", "October", "December"
//      예시: "January", "October", "December"
//    - Standard English month names
//      표준 영어 월 이름
//
// 3. SHORT ENGLISH MONTH NAME (짧은 영어 월 이름)
//    - MonthNameShort: 3-letter abbreviation
//      3글자 약어
//    - Examples: "Jan", "Oct", "Dec"
//      예시: "Jan", "Oct", "Dec"
//    - Compact format for space-constrained displays
//      공간 제약 디스플레이를 위한 컴팩트 형식
//
// 4. QUARTER CALCULATION (분기 계산)
//    - Quarter: Returns quarter number (1-4)
//      분기 번호 반환 (1-4)
//    - Q1: Jan-Mar, Q2: Apr-Jun, Q3: Jul-Sep, Q4: Oct-Dec
//      Q1: 1-3월, Q2: 4-6월, Q3: 7-9월, Q4: 10-12월
//    - Standard business quarters
//      표준 비즈니스 분기
//
// ============================================================================
// DESIGN PHILOSOPHY / 설계 철학
// ============================================================================
//
// 1. LOCALIZATION SUPPORT (현지화 지원)
//    - Provides both Korean and English month names
//      한국어 및 영어 월 이름 모두 제공
//    - Easy to add more languages if needed
//      필요시 더 많은 언어 추가 용이
//    - Uses Go's native formatting where possible
//      가능한 경우 Go의 네이티브 형식 사용
//
// 2. MULTIPLE FORMATS (여러 형식)
//    - Full names for formal displays
//      공식 디스플레이용 전체 이름
//    - Short names for compact UIs
//      컴팩트 UI용 짧은 이름
//    - Korean format follows local convention ("N월")
//      한국어 형식은 로컬 규칙 따름 ("N월")
//
// 3. STANDARD BUSINESS QUARTERS (표준 비즈니스 분기)
//    - Calendar year quarters (Jan 1 - Dec 31)
//      역년 분기 (1월 1일 - 12월 31일)
//    - Not fiscal year (which may differ)
//      회계연도 아님 (다를 수 있음)
//    - Q1-Q4 numbering (1-based)
//      Q1-Q4 번호 (1 기반)
//
// 4. SIMPLICITY OVER FLEXIBILITY (유연성보다 단순성)
//    - Returns strings directly, not Month enum
//      Month 열거형이 아닌 문자열 직접 반환
//    - Quarter as simple integer (1-4)
//      간단한 정수로 분기 (1-4)
//    - Easy to use without extra conversions
//      추가 변환 없이 사용 용이
//
// 5. ZERO DEPENDENCIES ON EXTERNAL I18N (외부 I18N에 대한 제로 의존성)
//    - No external localization library
//      외부 현지화 라이브러리 없음
//    - Built-in support for Korean/English
//      한국어/영어 내장 지원
//    - Lightweight and self-contained
//      경량 및 독립적
//
// ============================================================================
// MONTH OPERATIONS OVERVIEW / 월 연산 개요
// ============================================================================
//
// MONTH NAMES (월 이름) - 3 functions
// ├─ MonthKorean      : Korean month name ("1월", "10월")
// ├─ MonthName        : Full English name ("January", "October")
// └─ MonthNameShort   : Short English name ("Jan", "Oct")
//
// QUARTER (분기) - 1 function
// └─ Quarter : Quarter of year (1-4)
//
// Total: 4 month/quarter functions
// 총: 4개의 월/분기 함수
//
// ============================================================================
// PERFORMANCE CHARACTERISTICS / 성능 특성
// ============================================================================
//
// TIME COMPLEXITY (시간 복잡도):
//
// MONTHKOREAN: O(1)
//   Simple Format() call
//   간단한 Format() 호출
//   ~100-200 nanoseconds
//
// MONTHNAME: O(1)
//   Returns string from time.Month.String()
//   time.Month.String()에서 문자열 반환
//   ~50-100 nanoseconds
//
// MONTHNAMESHORT: O(1)
//   Simple Format() call
//   간단한 Format() 호출
//   ~100-200 nanoseconds
//
// QUARTER: O(1)
//   Arithmetic: (month-1)/3 + 1
//   산술: (월-1)/3 + 1
//   ~10-20 nanoseconds
//
// SPACE COMPLEXITY (공간 복잡도):
// - All functions: O(1) - return pre-allocated strings
//   모든 함수: O(1) - 미리 할당된 문자열 반환
//
// PERFORMANCE NOTES:
// 성능 참고:
// 1. Quarter is fastest (simple arithmetic)
//    Quarter가 가장 빠름 (간단한 산술)
// 2. Month name functions use formatting (slightly slower)
//    월 이름 함수는 형식 사용 (약간 느림)
// 3. All operations are extremely efficient
//    모든 연산이 매우 효율적
// 4. No allocations after initial string creation
//    초기 문자열 생성 후 할당 없음
//
// ============================================================================
// MONTH NAME FORMATS / 월 이름 형식
// ============================================================================
//
// KOREAN FORMAT (한국어 형식):
// - "N월" where N is month number (1-12)
//   N은 월 번호 (1-12)
// - Examples:
//   * January   → "1월"
//   * February  → "2월"
//   * December  → "12월"
//
// ENGLISH FULL FORMAT (영어 전체 형식):
// - Full month name capitalized
//   대문자로 시작하는 전체 월 이름
// - Examples:
//   * Month 1  → "January"
//   * Month 2  → "February"
//   * Month 12 → "December"
//
// ENGLISH SHORT FORMAT (영어 짧은 형식):
// - First 3 letters capitalized
//   대문자로 시작하는 처음 3글자
// - Examples:
//   * January   → "Jan"
//   * February  → "Feb"
//   * September → "Sep"
//   * December  → "Dec"
//
// ============================================================================
// QUARTER SYSTEM / 분기 시스템
// ============================================================================
//
// QUARTER DEFINITION:
// 분기 정의:
//
// Q1 (Quarter 1): January, February, March
//                 1월, 2월, 3월
//
// Q2 (Quarter 2): April, May, June
//                 4월, 5월, 6월
//
// Q3 (Quarter 3): July, August, September
//                 7월, 8월, 9월
//
// Q4 (Quarter 4): October, November, December
//                 10월, 11월, 12월
//
// QUARTER CALCULATION:
// 분기 계산:
// Formula: (month - 1) / 3 + 1
// 공식: (월 - 1) / 3 + 1
//
// Examples:
// 예시:
//   January   (1): (1-1)/3 + 1 = 0/3 + 1 = 0 + 1 = 1 (Q1)
//   April     (4): (4-1)/3 + 1 = 3/3 + 1 = 1 + 1 = 2 (Q2)
//   July      (7): (7-1)/3 + 1 = 6/3 + 1 = 2 + 1 = 3 (Q3)
//   October  (10): (10-1)/3 + 1 = 9/3 + 1 = 3 + 1 = 4 (Q4)
//
// NOTE: This is CALENDAR YEAR quarters, not fiscal year.
// 참고: 이것은 역년 분기이지 회계연도가 아닙니다.
// Some companies/countries use different fiscal year start dates.
// 일부 회사/국가는 다른 회계연도 시작 날짜를 사용합니다.
//
// ============================================================================
// USAGE PATTERNS / 사용 패턴
// ============================================================================
//
// PATTERN 1: Localized Calendar Display (현지화된 달력 표시)
// Use case: Show month name in user's language
// 사용 사례: 사용자의 언어로 월 이름 표시
//
//   t := time.Now()
//   var monthName string
//
//   if userLanguage == "ko" {
//       monthName = timeutil.MonthKorean(t)
//   } else {
//       monthName = timeutil.MonthName(t)
//   }
//
//   fmt.Printf("Current month: %s\n", monthName)
//
// PATTERN 2: Compact Date Display (컴팩트 날짜 표시)
// Use case: Space-constrained UI (mobile, cards)
// 사용 사례: 공간 제약 UI (모바일, 카드)
//
//   t := time.Date(2024, 10, 15, 0, 0, 0, 0, timeutil.KST)
//   shortDate := fmt.Sprintf("%s %d, %d",
//       timeutil.MonthNameShort(t),
//       t.Day(),
//       t.Year(),
//   )
//   fmt.Println(shortDate)  // "Oct 15, 2024"
//
// PATTERN 3: Quarterly Report (분기별 보고서)
// Use case: Group data by business quarter
// 사용 사례: 비즈니스 분기별로 데이터 그룹화
//
//   type QuarterlyReport struct {
//       Year    int
//       Quarter int
//       Revenue float64
//   }
//
//   reports := make(map[string]*QuarterlyReport)
//
//   for _, transaction := range transactions {
//       quarter := timeutil.Quarter(transaction.Date)
//       year := transaction.Date.Year()
//       key := fmt.Sprintf("%d-Q%d", year, quarter)
//
//       if _, exists := reports[key]; !exists {
//           reports[key] = &QuarterlyReport{
//               Year:    year,
//               Quarter: quarter,
//           }
//       }
//       reports[key].Revenue += transaction.Amount
//   }
//
// PATTERN 4: Month Dropdown/Selector (월 드롭다운/선택기)
// Use case: Month picker in date selector
// 사용 사례: 날짜 선택기의 월 피커
//
//   func getMonthOptions(language string) []string {
//       months := make([]string, 12)
//       for i := 1; i <= 12; i++ {
//           t := time.Date(2024, time.Month(i), 1, 0, 0, 0, 0, timeutil.KST)
//           if language == "ko" {
//               months[i-1] = timeutil.MonthKorean(t)
//           } else {
//               months[i-1] = timeutil.MonthName(t)
//           }
//       }
//       return months
//   }
//
// PATTERN 5: Financial Period Analysis (재무 기간 분석)
// Use case: Calculate quarter-over-quarter growth
// 사용 사례: 분기별 성장률 계산
//
//   currentQuarter := timeutil.Quarter(time.Now())
//   currentYearStart := timeutil.StartOfYear(time.Now())
//
//   // Group sales by quarter
//   quarterSales := make(map[int]float64)
//   for q := 1; q <= 4; q++ {
//       quarterSales[q] = calculateQuarterSales(currentYearStart, q)
//   }
//
//   // Compare current quarter with previous
//   if currentQuarter > 1 {
//       growth := (quarterSales[currentQuarter] - quarterSales[currentQuarter-1]) /
//                  quarterSales[currentQuarter-1] * 100
//       fmt.Printf("Q%d vs Q%d: %.2f%% growth\n",
//           currentQuarter, currentQuarter-1, growth)
//   }
//
// PATTERN 6: Event Calendar (이벤트 달력)
// Use case: Display events grouped by month
// 사용 사례: 월별로 그룹화된 이벤트 표시
//
//   type Event struct {
//       Title string
//       Date  time.Time
//   }
//
//   // Group events by month
//   eventsByMonth := make(map[string][]Event)
//
//   for _, event := range events {
//       monthKey := timeutil.MonthKorean(event.Date)
//       eventsByMonth[monthKey] = append(eventsByMonth[monthKey], event)
//   }
//
//   // Display
//   for month, monthEvents := range eventsByMonth {
//       fmt.Printf("\n%s:\n", month)
//       for _, event := range monthEvents {
//           fmt.Printf("  - %s\n", event.Title)
//       }
//   }
//
// PATTERN 7: Chart Labels (차트 레이블)
// Use case: X-axis labels for monthly/quarterly charts
// 사용 사례: 월별/분기별 차트의 X축 레이블
//
//   // Monthly chart
//   labels := make([]string, 12)
//   for i := 0; i < 12; i++ {
//       t := time.Date(2024, time.Month(i+1), 1, 0, 0, 0, 0, timeutil.KST)
//       labels[i] = timeutil.MonthNameShort(t)
//   }
//   // ["Jan", "Feb", "Mar", ..., "Dec"]
//
//   // Quarterly chart
//   quarterLabels := []string{"Q1", "Q2", "Q3", "Q4"}
//
// PATTERN 8: File/Directory Naming (파일/디렉토리 명명)
// Use case: Organize backups by month
// 사용 사례: 월별로 백업 정리
//
//   t := time.Now()
//   dirName := fmt.Sprintf("%d-%s",
//       t.Year(),
//       timeutil.MonthName(t),
//   )
//   // "2024-October"
//
//   backupPath := fmt.Sprintf("/backups/%s/", dirName)
//   os.MkdirAll(backupPath, 0755)
//
// PATTERN 9: Quarter-based Filtering (분기 기반 필터링)
// Use case: Show data for current quarter
// 사용 사례: 현재 분기 데이터 표시
//
//   currentQuarter := timeutil.Quarter(time.Now())
//   currentYear := time.Now().Year()
//
//   filteredData := make([]DataPoint, 0)
//   for _, point := range allData {
//       pointQuarter := timeutil.Quarter(point.Timestamp)
//       pointYear := point.Timestamp.Year()
//
//       if pointYear == currentYear && pointQuarter == currentQuarter {
//           filteredData = append(filteredData, point)
//       }
//   }
//
//   fmt.Printf("Showing data for %d Q%d\n", currentYear, currentQuarter)
//
// PATTERN 10: Bilingual Display (이중 언어 표시)
// Use case: Show both Korean and English month names
// 사용 사례: 한국어 및 영어 월 이름 모두 표시
//
//   t := time.Now()
//   korean := timeutil.MonthKorean(t)
//   english := timeutil.MonthName(t)
//
//   fmt.Printf("%s (%s)\n", korean, english)
//   // "10월 (October)"
//
// ============================================================================
// EDGE CASES / 경계 사례
// ============================================================================
//
// ZERO TIME:
// 제로 시간:
//   var t time.Time
//   month := timeutil.MonthKorean(t)  // "1월" (January)
//   quarter := timeutil.Quarter(t)    // 1 (Q1)
//   // Zero time is January 1, year 1
//   // 제로 시간은 1년 1월 1일
//
// DIFFERENT TIMEZONES:
// 다른 타임존:
//   // Month/quarter based on date component only
//   // Timezone doesn't affect the result
//   // 월/분기는 날짜 구성 요소만 기반
//   // 타임존은 결과에 영향 없음
//   t1 := time.Date(2024, 10, 15, 23, 0, 0, 0, time.UTC)
//   t2 := t1.In(timeutil.KST)
//   // Both have same month and quarter
//   // 둘 다 같은 월 및 분기
//
// LEAP YEAR:
// 윤년:
//   // Month/quarter not affected by leap year
//   // February is still Q1 regardless of 28 or 29 days
//   // 월/분기는 윤년의 영향을 받지 않음
//   // 2월은 28일이든 29일이든 여전히 Q1
//
// FISCAL YEAR:
// 회계연도:
//   // Quarter() returns CALENDAR year quarters (Jan-based)
//   // Not fiscal year (which may start in different month)
//   // Quarter()는 역년 분기 반환 (1월 기반)
//   // 회계연도 아님 (다른 월에 시작할 수 있음)
//   // Example: Fiscal year starting in July would need custom logic
//   // 예시: 7월에 시작하는 회계연도는 사용자 정의 로직 필요
//
// ============================================================================
// THREAD SAFETY / 스레드 안전성
// ============================================================================
//
// THREAD-SAFE FUNCTIONS:
// 스레드 안전 함수:
// - All month/quarter functions are thread-safe
//   모든 월/분기 함수는 스레드 안전
// - Read-only operations on time.Time
//   time.Time에 대한 읽기 전용 연산
// - No shared mutable state
//   공유 변경 가능 상태 없음
// - Safe for concurrent use
//   동시 사용에 안전
//
// CONCURRENT USAGE:
// 동시 사용:
//   var wg sync.WaitGroup
//   dates := []time.Time{ /* ... */ }
//
//   for _, date := range dates {
//       wg.Add(1)
//       go func(t time.Time) {
//           defer wg.Done()
//           month := timeutil.MonthKorean(t)
//           quarter := timeutil.Quarter(t)
//           fmt.Printf("%s, Q%d\n", month, quarter)
//       }(date)
//   }
//   wg.Wait()
//
// ============================================================================
// DEPENDENCIES / 의존성
// ============================================================================
//
// This file depends on:
// 이 파일이 의존하는 항목:
//
// FROM constants.go:
// - MonthsPerQuarter: For quarter calculation (value: 3)
//
// STANDARD LIBRARY:
// - time.Time: Base time type
// - time.Month: Month enum
// - time.Time.Format(): For Korean/short names
// - time.Month.String(): For English full names
//
// USED BY (사용처):
// - Reports (monthly/quarterly reports)
//   보고서 (월별/분기별 보고서)
// - UI displays (calendar, date pickers)
//   UI 디스플레이 (달력, 날짜 선택기)
// - Localization (multi-language apps)
//   현지화 (다중 언어 앱)
// - Analytics (period-based data grouping)
//   분석 (기간 기반 데이터 그룹화)
//
// ============================================================================
// BEST PRACTICES / 모범 사례
// ============================================================================
//
// 1. USE KOREAN FOR KOREAN USERS
//    한국 사용자에게 한국어 사용
//    monthName := timeutil.MonthKorean(t)
//
// 2. USE SHORT NAMES FOR COMPACT DISPLAYS
//    컴팩트 디스플레이에 짧은 이름 사용
//    label := timeutil.MonthNameShort(t)
//
// 3. USE QUARTER FOR BUSINESS REPORTING
//    비즈니스 보고에 Quarter 사용
//    q := timeutil.Quarter(salesDate)
//
// 4. CACHE MONTH OPTIONS FOR DROPDOWNS
//    드롭다운용 월 옵션 캐시
//    // Generate once, reuse
//
// 5. USE MONTHNAME FOR FORMAL DISPLAYS
//    공식 디스플레이에 MonthName 사용
//    fmt.Printf("%s %d, %d", timeutil.MonthName(t), t.Day(), t.Year())
//
// 6. GROUP DATA BY QUARTER FOR TRENDS
//    트렌드를 위해 분기별로 데이터 그룹화
//    dataByQuarter[timeutil.Quarter(t)] = append(...)
//
// 7. CONSIDER USER LOCALE FOR MONTH NAMES
//    월 이름에 사용자 로케일 고려
//    if locale == "ko" { use MonthKorean } else { use MonthName }
//
// 8. USE CONSISTENT QUARTER DEFINITION
//    일관된 분기 정의 사용
//    // Calendar year (Jan-Dec) or fiscal year (custom)
//
// ============================================================================

// MonthKorean returns the Korean name of the month.
// MonthKorean은 월의 한글 이름을 반환합니다.
//
// Example
// 예제:
//
//	t := time.Date(2025, 10, 14, 0, 0, 0, 0, KST)
//
// month := timeutil.MonthKorean(t)
// fmt.Println(month) // "10월"
func MonthKorean(t time.Time) string {
	return t.Format("1월")
}

// MonthName returns the English name of the month.
// MonthName은 월의 영문 이름을 반환합니다.
//
// Example
// 예제:
//
//	t := time.Date(2025, 10, 14, 0, 0, 0, 0, KST)
//	month := timeutil.MonthName(t)
//	fmt.Println(month) // "October"
func MonthName(t time.Time) string {
	return t.Month().String()
}

// MonthNameShort returns the short English name of the month (3 letters).
// MonthNameShort는 월의 짧은 영문 이름을 반환합니다 (3글자).
//
// Example
// 예제:
//
//	t := time.Date(2025, 10, 14, 0, 0, 0, 0, KST)
//	month := timeutil.MonthNameShort(t)
//	fmt.Println(month) // "Oct"
func MonthNameShort(t time.Time) string {
	return t.Format("Jan")
}

// Quarter returns the quarter of the year (1-4).
// Quarter는 년의 분기를 반환합니다 (1-4).
//
// Q1: January-March, Q2: April-June, Q3: July-September, Q4: October-December
// Q1: 1-3월, Q2: 4-6월, Q3: 7-9월, Q4: 10-12월
//
// Example
// 예제:
//
//	t := time.Date(2025, 10, 14, 0, 0, 0, 0, KST)
//	quarter := timeutil.Quarter(t)
//	fmt.Println(quarter) // 4
func Quarter(t time.Time) int {
	month := int(t.Month())
	return (month-1)/MonthsPerQuarter + 1
}
