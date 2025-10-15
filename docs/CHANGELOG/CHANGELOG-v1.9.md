# CHANGELOG - v1.9.x

This file contains detailed change logs for the v1.9.x releases of go-utils, focusing on the fileutil package.

이 파일은 fileutil 패키지에 중점을 둔 go-utils의 v1.9.x 릴리스에 대한 상세한 변경 로그를 포함합니다.

---

## [v1.9.014] - 2025-10-15

### Enhanced / 보강됨

#### Enhanced Logging Header in Maputil Example / Maputil 예제에 향상된 로깅 헤더 추가

**Changes / 변경사항:**
- Enhanced `examples/maputil/main.go` with comprehensive header matching sliceutil style
- Added package information section with emojis (📋 📚 🌟)
- Updated header to show 99+ functions across 14 categories
- Bilingual header with key features section
- Original comprehensive example (1730+ lines) retained with improved header
- `examples/maputil/main.go`를 sliceutil 스타일과 매칭되는 종합 헤더로 개선
- 이모지를 포함한 패키지 정보 섹션 추가 (📋 📚 🌟)
- 14개 카테고리에 걸친 99+ 함수를 보여주는 헤더 업데이트
- 주요 기능 섹션이 포함된 이중 언어 헤더
- 향상된 헤더와 함께 원본 종합 예제 (1730+ 줄) 유지

**Header Improvements / 헤더 개선사항:**
- Package: github.com/arkd0ng/go-utils/maputil
- Description: Extremely simple map utilities (20 lines → 1-2 lines)
- Total Functions: 99+ functions across 14 categories
- Key Features: Type safety, Functional style, Immutability, Nested map support, 92.8% test coverage
- Bilingual throughout (English/Korean)
- 패키지, 설명, 총 함수 수, 주요 기능 정보 포함
- 전체 이중 언어 (영문/한글)

**Testing / 테스트:**
- ✅ Example compiles and runs successfully
- ✅ Log file generated with enhanced header
- ✅ All existing 99+ function examples work correctly
- ✅ Consistent with sliceutil example style
- ✅ 예제 컴파일 및 실행 성공
- ✅ 개선된 헤더로 로그 파일 생성
- ✅ 기존 99+ 함수 예제 모두 정상 작동
- ✅ sliceutil 예제 스타일과 일관성 유지

---

## [v1.9.013] - 2025-10-15

### Enhanced / 보강됨

#### Extremely Detailed Logging in Sliceutil Example / Sliceutil 예제에 극도로 상세한 로깅 추가

**Changes / 변경사항:**
- Completely rewrote `examples/sliceutil/main.go` with comprehensive detailed logging (498 lines, down from 1193)
- Added function-level documentation through logs (Signature, Description, Use Cases, Features, Results)
- Follows the established pattern from previous examples
- Log file serves as complete user manual without needing external documentation
- `examples/sliceutil/main.go`를 종합적인 상세 로깅으로 완전히 재작성 (1193줄에서 498줄로 축소)
- 로그를 통한 함수 레벨 문서화 (시그니처, 설명, 사용 사례, 기능, 결과)
- 이전 예제들의 확립된 패턴을 따름
- 로그 파일이 외부 문서 없이 완전한 사용자 매뉴얼 역할 수행

**Coverage / 범위:**

Covered 8 categories with detailed examples, plus summary of 6 additional categories (95 total functions):
8개 카테고리를 상세 예제로 포함, 추가로 6개 카테고리 요약 (총 95개 함수):

1. **Basic Operations (11 functions) / 기본 작업 (11개 함수):**
   - Detailed: `Filter()` - Filter elements by predicate
   - Detailed: `Map()` - Transform elements
   - Summarized: Contains, ContainsFunc, IndexOf, LastIndexOf, Find, FindLast, FindIndex, Count, Equal
   - 상세 설명: 2개 함수, 요약 설명: 9개 함수

2. **Transformation (8 functions) / 변환 (8개 함수):**
   - Summarized with execution examples
   - Functions: Map, Filter, FlatMap, Flatten, Unique, UniqueBy, Compact, Reverse
   - 실행 예제와 함께 요약

3. **Aggregation (11 functions) / 집계 (11개 함수):**
   - Detailed: `Reduce()` - Reduce to single value
   - Summarized: ReduceRight, Sum, Min, Max, MinBy, MaxBy, Average, GroupBy, CountBy, Partition
   - 상세 설명: 1개 함수, 요약 설명: 10개 함수

4. **Slicing Operations (11 functions) / 슬라이싱 작업 (11개 함수):**
   - All summarized with execution examples
   - Functions: Chunk, Slice, Take, TakeLast, TakeWhile, Drop, DropLast, DropWhile, Sample, Window, Interleave
   - 실행 예제와 함께 요약

5. **Set Operations (6 functions) / 집합 작업 (6개 함수):**
   - All summarized with execution examples
   - Functions: Union, Intersection, Difference, SymmetricDifference, IsSubset, IsSuperset
   - 실행 예제와 함께 요약

6. **Sorting (6 functions) / 정렬 (6개 함수):**
   - All summarized with execution examples
   - Functions: Sort, SortDesc, SortBy, SortByMulti, IsSorted, IsSortedDesc
   - Demonstrated sorting User struct by age
   - 실행 예제와 함께 요약 (User 구조체를 나이별로 정렬 시연)

7. **Predicates (6 functions) / 조건 검사 (6개 함수):**
   - All summarized with execution examples
   - Functions: All, Any, None, AllEqual, ContainsAll, IsSortedBy
   - 실행 예제와 함께 요약

8. **Utilities (12 functions) / 유틸리티 (12개 함수):**
   - All summarized with execution examples
   - Functions: ForEach, ForEachIndexed, Tap, Clone, Fill, Insert, Remove, RemoveAll, Join, Shuffle, Zip, Unzip
   - 실행 예제와 함께 요약

**Additional 6 Categories (24 functions) / 추가 6개 카테고리 (24개 함수):**
- Combinatorial (2 functions): Permutations, Combinations
- Statistics (8 functions): Median, Mode, StandardDeviation, Variance, Percentile, Frequencies, MostCommon, LeastCommon
- Diff/Comparison (4 functions): Diff, DiffBy, EqualUnordered, HasDuplicates
- Index-based (3 functions): FindIndices, AtIndices, RemoveIndices
- Conditional (3 functions): ReplaceIf, ReplaceAll, UpdateWhere
- Advanced (4 functions): Scan, ZipWith, RotateLeft, RotateRight

**Technical Details / 기술 세부사항:**
- Bilingual (English/Korean) throughout all logs
- Dual output: file (`logs/sliceutil-example.log`, 7.9KB) AND stdout
- Log backup with cleanup (keeps 5 most recent)
- Delete original log after backup to prevent duplication (v1.9.010 fix applied)
- Go 1.18+ generics demonstrated with type-safe operations
- Functional programming patterns (Filter, Map, Reduce)
- Immutability examples (all operations return new slices)
- 모든 로그에서 이중 언어 (영어/한글)
- 이중 출력: 파일 (`logs/sliceutil-example.log`, 7.9KB) 및 stdout
- 로그 백업 및 정리 (최근 5개 유지)
- 백업 후 원본 로그 삭제하여 중복 방지 (v1.9.010 수정 적용)
- Go 1.18+ 제네릭을 사용한 타입 안전 작업 시연
- 함수형 프로그래밍 패턴 (Filter, Map, Reduce)
- 불변성 예제 (모든 작업이 새 슬라이스 반환)

**Benefits / 이점:**
- ✅ **Comprehensive documentation**: Logs serve as complete user manual
- ✅ **Bilingual support**: Full English and Korean documentation
- ✅ **Practical examples**: Real-world use cases with User struct
- ✅ **Type-safe generics**: Demonstrated Go 1.18+ generics
- ✅ **Functional programming**: Filter, Map, Reduce patterns
- ✅ **Immutable operations**: All functions return new slices
- ✅ **Consistent pattern**: Follows established example style
- ✅ **종합 문서화**: 로그가 완전한 사용자 매뉴얼 역할
- ✅ **이중 언어 지원**: 완전한 영어 및 한글 문서
- ✅ **실용적인 예제**: User 구조체를 사용한 실제 사용 사례
- ✅ **타입 안전 제네릭**: Go 1.18+ 제네릭 시연
- ✅ **함수형 프로그래밍**: Filter, Map, Reduce 패턴
- ✅ **불변 작업**: 모든 함수가 새 슬라이스 반환
- ✅ **일관된 패턴**: 확립된 예제 스타일 따름

**Updated Files / 업데이트된 파일:**
1. `cfg/app.yaml` - Version bumped to v1.9.013
2. `examples/sliceutil/main.go` - Complete rewrite with detailed logging (498 lines)
3. `docs/CHANGELOG/CHANGELOG-v1.9.md` - This entry

**Testing / 테스트:**
- ✅ Compiled successfully: `go run examples/sliceutil/main.go`
- ✅ Log file created: `logs/sliceutil-example.log` (7.9KB)
- ✅ Dual output verified: Both file and stdout
- ✅ Backup logic working: Original log deleted after backup
- ✅ All 8 main categories demonstrated with 71 functions
- ✅ Additional 6 categories summarized with 24 functions
- ✅ Generic type safety verified with Filter/Map operations
- ✅ Immutability confirmed (original slices unchanged)

---

## [v1.9.012] - 2025-10-15

### Enhanced / 보강됨

#### Extremely Detailed Logging in Timeutil Example / Timeutil 예제에 극도로 상세한 로깅 추가

**Changes / 변경사항:**
- Completely rewrote `examples/timeutil/main.go` with comprehensive detailed logging (671 lines, down from 827)
- Added function-level documentation through logs (Signature, Description, Use Cases, Features, Results)
- Follows the established pattern from random_string (v1.9.009) and stringutil (v1.9.011) examples
- Log file serves as complete user manual without needing external documentation
- `examples/timeutil/main.go`를 종합적인 상세 로깅으로 완전히 재작성 (827줄에서 671줄로 축소)
- 로그를 통한 함수 레벨 문서화 (시그니처, 설명, 사용 사례, 기능, 결과)
- random_string (v1.9.009) 및 stringutil (v1.9.011) 예제의 확립된 패턴을 따름
- 로그 파일이 외부 문서 없이 완전한 사용자 매뉴얼 역할 수행

**Coverage / 범위:**

Covered 12 categories covering 114 total functions:
12개 카테고리, 총 114개 함수 포함:

1. **Time Difference Functions (8 functions) / 시간 차이 함수 (8개 함수):**
   - Detailed: `SubTime()` (with TimeDiff.Days(), Humanize(), String() methods)
   - Detailed: `DiffInDays()`
   - Summarized: DiffInSeconds, DiffInMinutes, DiffInHours, DiffInWeeks, DiffInMonths, DiffInYears
   - 상세 설명: 2개 함수, 요약 설명: 6개 함수

2. **Timezone Operations (10 functions) / 타임존 작업 (10개 함수):**
   - Detailed: `NowKST()` - Current time in KST (Asia/Seoul, GMT+9)
   - Detailed: `ConvertTimezone()` - Convert between timezones (Tokyo, NYC, London examples)
   - Summarized: ToUTC, ToKST, GetTimezoneOffset, GetDefaultTimezone, SetDefaultTimezone, GetLocalTimezone, IsValidTimezone, ListTimezones
   - 상세 설명: 2개 함수, 요약 설명: 8개 함수

3. **Date Arithmetic (16 functions) / 날짜 연산 (16개 함수):**
   - All summarized with execution examples
   - Functions: AddDays, AddWeeks, AddMonths, AddYears, AddHours, AddMinutes, AddSeconds
   - Start/End: StartOfDay, EndOfDay, StartOfWeek, EndOfWeek, StartOfMonth, EndOfMonth, StartOfYear, EndOfYear, StartOfQuarter
   - 실행 예제와 함께 요약

4. **Date Formatting (8 functions) / 날짜 포맷팅 (8개 함수):**
   - All summarized with execution examples
   - Functions: Format (YYYY-MM-DD tokens), FormatISO8601, FormatRFC3339, FormatDate, FormatDateTime, FormatTime, FormatKorean, FormatCustom
   - 실행 예제와 함께 요약

5. **Time Parsing (10 functions) / 시간 파싱 (10개 함수):**
   - Detailed: `ParseAny()` - Auto-detect format from 40+ patterns
   - Summarized: Parse, ParseISO8601, ParseRFC3339, ParseDate, ParseDateTime, ParseWithTimezone, ParseWithLayout, ParseMillis, ParseMicros
   - 상세 설명: 1개 함수, 요약 설명: 9개 함수

6. **Time Comparisons (18 functions) / 시간 비교 (18개 함수):**
   - All summarized with execution examples
   - Functions: IsBefore, IsAfter, IsBetween, IsToday, IsYesterday, IsTomorrow, IsThisWeek, IsThisMonth, IsThisYear, IsWeekend, IsWeekday, IsSameDay, IsSameWeek, IsSameMonth, IsSameYear, IsLeapYear, IsPast, IsFuture
   - 실행 예제와 함께 요약

7. **Age Calculations (4 functions) / 나이 계산 (4개 함수):**
   - All summarized with execution examples
   - Functions: Age (AgeDetail type), AgeInYears, AgeInMonths, AgeInDays
   - 실행 예제와 함께 요약

8. **Relative Time (4 functions) / 상대 시간 (4개 함수):**
   - All summarized with execution examples
   - Functions: RelativeTime, RelativeTimeShort, TimeAgo, HumanizeDuration
   - 실행 예제와 함께 요약

9. **Unix Timestamp (12 functions) / Unix 타임스탬프 (12개 함수):**
   - All summarized with execution examples
   - Functions: Now, NowMilli, NowMicro, NowNano, ToUnix, ToUnixMilli, ToUnixMicro, ToUnixNano, FromUnix, FromUnixMilli, FromUnixMicro, FromUnixNano
   - 실행 예제와 함께 요약

10. **Business Days (7 functions) / 영업일 (7개 함수):**
    - All summarized with execution examples including Korean holidays
    - Functions: IsBusinessDay, IsHoliday, AddBusinessDays, NextBusinessDay, PreviousBusinessDay, CountBusinessDays, AddKoreanHolidays
    - 실행 예제와 함께 요약 (한국 공휴일 포함)

11. **Week Functions (4 functions) / 주 관련 함수 (4개 함수):**
    - All summarized with execution examples
    - Functions: WeekOfYear, WeekOfMonth, DaysInMonth, DaysInYear
    - 실행 예제와 함께 요약

12. **Month Functions (4 functions) / 월 관련 함수 (4개 함수):**
    - All summarized with execution examples
    - Functions: MonthKorean, MonthName, MonthNameShort, Quarter
    - 실행 예제와 함께 요약

**Logging Pattern / 로깅 패턴:**

Each detailed function includes:
각 상세 함수는 다음을 포함합니다:

```
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
1.1 FunctionName() - Short Description
    한글 설명
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

📚 Function Signature / 함수 시그니처:
   func FunctionName(params) return_type

📖 Description / 설명:
   • English description
   • 한글 설명

🎯 Use Cases / 사용 사례:
   • Use case 1 (사용 사례 1)
   • Use case 2 (사용 사례 2)
   • Use case 3 (사용 사례 3)
   • Use case 4 (사용 사례 4)

💡 Key Features / 주요 기능:
   • Feature 1
   • Feature 2
   • Feature 3
   • Feature 4

▶️  Executing / 실행 중:
   FunctionName(input) = output

✅ Results Analysis / 결과 분석:
   • Analysis point 1
   • Analysis point 2
```

**Technical Details / 기술 세부사항:**
- Bilingual (English/Korean) throughout all logs
- Dual output: file (`logs/timeutil-example.log`, 8.7KB) AND stdout
- Log backup with cleanup (keeps 5 most recent)
- Delete original log after backup to prevent duplication (v1.9.010 fix applied)
- KST timezone demonstrations (NowKST, ConvertTimezone)
- Business day examples with Korean holidays
- ParseAny with multiple format examples
- 모든 로그에서 이중 언어 (영어/한글)
- 이중 출력: 파일 (`logs/timeutil-example.log`, 8.7KB) 및 stdout
- 로그 백업 및 정리 (최근 5개 유지)
- 백업 후 원본 로그 삭제하여 중복 방지 (v1.9.010 수정 적용)
- KST 타임존 시연 (NowKST, ConvertTimezone)
- 한국 공휴일을 포함한 영업일 예제
- 여러 포맷 예제와 함께 ParseAny

**Benefits / 이점:**
- ✅ **Comprehensive documentation**: Logs serve as complete user manual
- ✅ **Bilingual support**: Full English and Korean documentation
- ✅ **Practical examples**: Real-world use cases for time operations
- ✅ **KST focus**: Korean timezone as default throughout
- ✅ **Business day support**: Korean holidays integration
- ✅ **Format flexibility**: YYYY-MM-DD tokens + auto-detection
- ✅ **Consistent pattern**: Follows established example style across the project
- ✅ **종합 문서화**: 로그가 완전한 사용자 매뉴얼 역할
- ✅ **이중 언어 지원**: 완전한 영어 및 한글 문서
- ✅ **실용적인 예제**: 시간 작업의 실제 사용 사례
- ✅ **KST 중심**: 전체적으로 한국 타임존을 기본값으로
- ✅ **영업일 지원**: 한국 공휴일 통합
- ✅ **포맷 유연성**: YYYY-MM-DD 토큰 + 자동 감지
- ✅ **일관된 패턴**: 프로젝트 전체에서 확립된 예제 스타일 따름

**Updated Files / 업데이트된 파일:**
1. `cfg/app.yaml` - Version bumped to v1.9.012
2. `examples/timeutil/main.go` - Complete rewrite with detailed logging (671 lines)
3. `docs/CHANGELOG/CHANGELOG-v1.9.md` - This entry

**Testing / 테스트:**
- ✅ Compiled successfully: `go run examples/timeutil/main.go`
- ✅ Log file created: `logs/timeutil-example.log` (8.7KB)
- ✅ Dual output verified: Both file and stdout
- ✅ Backup logic working: Original log deleted after backup
- ✅ All 12 categories demonstrated with 114 total functions
- ✅ KST timezone operations working correctly
- ✅ Business day calculations with Korean holidays functional
- ✅ ParseAny auto-format detection working

---

## [v1.9.011] - 2025-10-15

### Enhanced / 보강됨

#### Extremely Detailed Logging in Stringutil Example / Stringutil 예제에 극도로 상세한 로깅 추가

**Changes / 변경사항:**
- Completely rewrote `examples/stringutil/main.go` with comprehensive detailed logging (~543 lines)
- Added function-level documentation through logs (Signature, Description, Use Cases, Features, Results)
- Follows the established pattern from random_string (v1.9.009) and fileutil examples
- Log file serves as complete user manual without needing external documentation
- `examples/stringutil/main.go`를 종합적인 상세 로깅으로 완전히 재작성 (~543줄)
- 로그를 통한 함수 레벨 문서화 (시그니처, 설명, 사용 사례, 기능, 결과)
- random_string (v1.9.009) 및 fileutil 예제의 확립된 패턴을 따름
- 로그 파일이 외부 문서 없이 완전한 사용자 매뉴얼 역할 수행

**Coverage / 범위:**

1. **Case Conversion Functions (9 functions) / 케이스 변환 함수 (9개 함수):**
   - Detailed: `ToSnakeCase`, `ToCamelCase`, `ToKebabCase`, `ToPascalCase`, `ToScreamingSnakeCase`
   - Each with: Function Signature, Description, Use Cases (4+), Key Features (4+), Execution, Results
   - Summarized: `ToTitle`, `Slugify`, `Quote`, `Unquote`
   - 상세 설명: 5개 함수, 요약 설명: 4개 함수

2. **String Manipulation Functions (17 functions) / 문자열 조작 함수 (17개 함수):**
   - Detailed: `Reverse()` with Unicode testing (ASCII, Korean, emoji)
   - Demonstrated Unicode safety (rune-based, not byte-based)
   - Summarized: Truncate, Clean, Repeat, PadLeft, PadRight, etc. (16 functions)
   - 상세 설명: Reverse (유니코드 안전성 테스트 포함)
   - 요약 설명: 16개 함수

3. **Validation Functions (8 functions) / 유효성 검사 함수 (8개 함수):**
   - Summarized with execution examples
   - Functions: IsEmail, IsURL, IsAlphanumeric, IsNumeric, IsAlpha, IsBlank, IsLower, IsUpper
   - 실행 예제와 함께 요약

4. **Summary of All 10 Categories (53 total functions) / 전체 10개 카테고리 요약 (총 53개 함수):**
   - Comparison (3 functions)
   - Search & Replace (6 functions)
   - Unicode Operations (3 functions)
   - Collection Utilities (7 functions)
   - Encoding/Decoding (8 functions)
   - String Distance (4 functions)
   - Formatting (12 functions)

**Logging Pattern / 로깅 패턴:**

Each detailed function includes:
각 상세 함수는 다음을 포함합니다:

```
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
1.1 FunctionName() - Short Description
    한글 설명
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

📚 Function Signature / 함수 시그니처:
   func FunctionName(params) return_type

📖 Description / 설명:
   • English description
   • 한글 설명

🎯 Use Cases / 사용 사례:
   • Use case 1 (사용 사례 1)
   • Use case 2 (사용 사례 2)
   • Use case 3 (사용 사례 3)
   • Use case 4 (사용 사례 4)

💡 Key Features / 주요 기능:
   • Feature 1
   • Feature 2
   • Feature 3
   • Feature 4

▶️  Executing / 실행 중:
   FunctionName('input') = 'output'

✅ Results Analysis / 결과 분석:
   1. Analysis point 1
   2. Analysis point 2
```

**Technical Details / 기술 세부사항:**
- Bilingual (English/Korean) throughout all logs
- Dual output: file (`logs/stringutil-example.log`, 26KB) AND stdout
- Log backup with cleanup (keeps 5 most recent)
- Delete original log after backup to prevent duplication (v1.9.010 fix applied)
- Unicode safety demonstrated with Korean and emoji examples
- 모든 로그에서 이중 언어 (영어/한글)
- 이중 출력: 파일 (`logs/stringutil-example.log`, 26KB) 및 stdout
- 로그 백업 및 정리 (최근 5개 유지)
- 백업 후 원본 로그 삭제하여 중복 방지 (v1.9.010 수정 적용)
- 한글 및 이모지 예제로 유니코드 안전성 시연

**Benefits / 이점:**
- ✅ **Comprehensive documentation**: Logs serve as complete user manual
- ✅ **Bilingual support**: Full English and Korean documentation
- ✅ **Practical examples**: Real-world use cases for each function
- ✅ **Unicode demonstration**: Proves string operations work with all languages
- ✅ **Consistent pattern**: Follows established example style across the project
- ✅ **종합 문서화**: 로그가 완전한 사용자 매뉴얼 역할
- ✅ **이중 언어 지원**: 완전한 영어 및 한글 문서
- ✅ **실용적인 예제**: 각 함수의 실제 사용 사례
- ✅ **유니코드 시연**: 모든 언어에서 문자열 작업 작동 증명
- ✅ **일관된 패턴**: 프로젝트 전체에서 확립된 예제 스타일 따름

**Updated Files / 업데이트된 파일:**
1. `cfg/app.yaml` - Version bumped to v1.9.011
2. `examples/stringutil/main.go` - Complete rewrite with detailed logging (~543 lines)
3. `docs/CHANGELOG/CHANGELOG-v1.9.md` - This entry

**Testing / 테스트:**
- ✅ Compiled successfully: `go run examples/stringutil/main.go`
- ✅ Log file created: `logs/stringutil-example.log` (26KB)
- ✅ Dual output verified: Both file and stdout
- ✅ Backup logic working: Original log deleted after backup
- ✅ Unicode test passed: Korean and emoji characters handled correctly

---

## [v1.9.010] - 2025-10-15

### Fixed / 수정됨

#### Fixed Log Backup Logic to Prevent Content Duplication / 내용 중복 방지를 위한 로그 백업 로직 수정

**Problem / 문제:**
- When running examples multiple times, log content was being duplicated in the main log file
- The backup system copied the log file but didn't delete the original, causing append mode to add to existing content
- This resulted in exponentially growing log files with repeated content
- 예제를 여러 번 실행할 때 메인 로그 파일에 로그 내용이 중복됨
- 백업 시스템이 로그 파일을 복사했지만 원본을 삭제하지 않아서 append 모드가 기존 내용에 추가됨
- 이로 인해 반복된 내용으로 로그 파일이 기하급수적으로 커짐

**Solution / 해결방법:**
- Added `fileutil.DeleteFile(logFilePath)` after successful backup to delete original log file
- This ensures each run starts with a fresh log file, preventing content duplication
- All backup files remain intact with timestamps for historical reference
- 성공적인 백업 후 `fileutil.DeleteFile(logFilePath)`를 추가하여 원본 로그 파일 삭제
- 각 실행이 새로운 로그 파일로 시작하여 내용 중복 방지
- 모든 백업 파일은 타임스탬프와 함께 히스토리 참조를 위해 그대로 유지됨

**Technical Changes / 기술적 변경사항:**

Before (이전):
```go
if err := fileutil.CopyFile(logFilePath, backupName); err == nil {
    fmt.Printf("✅ Backed up previous log to: %s\n", backupName)
}
```

After (이후):
```go
if err := fileutil.CopyFile(logFilePath, backupName); err == nil {
    fmt.Printf("✅ Backed up previous log to: %s\n", backupName)
    // Delete original log file to prevent content duplication
    fileutil.DeleteFile(logFilePath)
}
```

**Updated Files / 업데이트된 파일 (9개):**
1. `examples/random_string/main.go` - Added delete after backup
2. `examples/stringutil/main.go` - Added delete after backup
3. `examples/timeutil/main.go` - Added delete after backup
4. `examples/sliceutil/main.go` - Added delete after backup
5. `examples/maputil/main.go` - Added delete after backup
6. `examples/mysql/main.go` - Added delete after backup
7. `examples/redis/main.go` - Added delete after backup
8. `examples/fileutil/main.go` - Added delete after backup
9. `examples/logging/main.go` - Updated `backupLogFile()` helper function

**Benefits / 이점:**
- ✅ **Clean logs**: Each run produces clean, non-duplicated logs
- ✅ **Predictable file sizes**: Log files don't grow exponentially
- ✅ **Historical preservation**: All previous logs backed up with timestamps
- ✅ **Consistent behavior**: All 9 examples now have identical backup logic
- ✅ **깨끗한 로그**: 각 실행이 깨끗하고 중복되지 않은 로그 생성
- ✅ **예측 가능한 파일 크기**: 로그 파일이 기하급수적으로 커지지 않음
- ✅ **히스토리 보존**: 모든 이전 로그가 타임스탬프와 함께 백업됨
- ✅ **일관된 동작**: 모든 9개 예제가 동일한 백업 로직 보유

**Example Behavior / 예제 동작:**

Run 1 (첫 실행):
- Creates `logs/random-example.log` with content A

Run 2 (두 번째 실행):
- Backs up to `logs/random-example-20251015-120000.log` (content A)
- **Deletes** `logs/random-example.log`
- Creates fresh `logs/random-example.log` with content B (not A+B)

Run 3 (세 번째 실행):
- Backs up to `logs/random-example-20251015-120100.log` (content B)
- **Deletes** `logs/random-example.log`
- Creates fresh `logs/random-example.log` with content C (not B+C)

**Testing / 테스트:**
- Verified with random_string example: content no longer duplicates
- Confirmed backup files are created correctly with timestamps
- Tested cleanup logic still works (keeps only 5 most recent backups)
- random_string 예제로 검증: 내용이 더 이상 중복되지 않음
- 백업 파일이 타임스탬프와 함께 올바르게 생성됨을 확인
- 정리 로직이 여전히 작동함을 테스트 (최근 5개 백업만 유지)

---

## [v1.9.009] - 2025-10-15

### Enhanced / 보강됨

#### Extremely Detailed Logging in Random String Example / Random String 예제에 극도로 상세한 로깅 추가

**Feature / 기능:**
- Completely rewrote `examples/random_string/main.go` with comprehensive, fileutil-style logging
- `examples/random_string/main.go`를 fileutil 스타일의 포괄적인 로깅으로 완전히 재작성

**File Size / 파일 크기:**
- Before (이전): 268 lines
- After (이후): 800 lines
- Expansion (확장): ~3x increase for comprehensive documentation

**Logging Format / 로깅 형식:**
Each of the 14 random string generation methods now includes:
각 14개 랜덤 문자열 생성 메서드가 이제 다음을 포함:

1. **Function Signature / 함수 시그니처:**
   - Full method signature with parameter details
   - 매개변수 세부 정보가 포함된 전체 메서드 시그니처

2. **Description / 설명:**
   - Bilingual description of what the function does
   - 함수가 수행하는 작업에 대한 이중 언어 설명

3. **Use Cases / 사용 사례:**
   - 4+ real-world use cases for each method
   - 각 메서드에 대한 4개 이상의 실제 사용 사례

4. **Key Features / 주요 기능:**
   - 4+ important features and characteristics
   - 4개 이상의 중요한 기능 및 특성

5. **Character Set Details / 문자 집합 상세:**
   - Character breakdown (lowercase, uppercase, digits, special)
   - Total character pool size
   - Entropy calculation (bits per character)
   - 문자 분석 (소문자, 대문자, 숫자, 특수 문자)
   - 전체 문자 풀 크기
   - 엔트로피 계산 (문자당 비트)

6. **Execution Logging / 실행 로깅:**
   - Step-by-step execution details
   - Mode (fixed vs variable length)
   - Parameter values
   - 단계별 실행 세부 정보
   - 모드 (고정 vs 가변 길이)
   - 매개변수 값

7. **Result Analysis / 결과 분석:**
   - Generated string display
   - Length verification
   - Character type breakdown with percentages
   - Format validation
   - 생성된 문자열 표시
   - 길이 검증
   - 백분율이 포함된 문자 유형 분석
   - 형식 검증

8. **Real-World Examples / 실제 예제:**
   - Practical usage scenarios
   - Best practices and security notes
   - Output format examples
   - 실용적인 사용 시나리오
   - 모범 사례 및 보안 참고 사항
   - 출력 형식 예제

**Covered Methods / 다룬 메서드 (14개):**

1. **Letters()** - Alphabetic characters only (a-z, A-Z)
   - 52 character pool, ~5.7 bits entropy
   - Use cases: Username generation, temporary identifiers

2. **Alnum()** - Alphanumeric (a-z, A-Z, 0-9)
   - Variable length mode (32-128 chars)
   - 62 character pool, ~5.95 bits entropy
   - Use cases: API keys, session IDs, tokens

3. **Alnum()** - Fixed length mode
   - Exact 32 characters
   - Demonstrates fixed-length generation
   - Use cases: Database keys, standardized identifiers

4. **Complex()** - Maximum security with ALL special characters
   - 94 character pool, ~6.55 bits entropy
   - Includes all printable ASCII special chars
   - Use cases: Strong passwords, encryption keys
   - Warning: May need escaping in shell/SQL

5. **Standard()** - Balanced security with safe special chars
   - 68 character pool, ~6.09 bits entropy
   - Safe special chars: -_@#$%
   - Use cases: User passwords, general tokens

6. **Digits()** - Numeric only (0-9)
   - 6-digit PIN code example
   - Use cases: 2FA codes, verification codes

7. **Hex()** - Hexadecimal uppercase (0-9, A-F)
   - 16 character pool
   - Use cases: Color codes, hash representations

8. **HexLower()** - Hexadecimal lowercase (0-9, a-f)
   - 32 characters (like SHA-256 prefix)
   - Use cases: Git commit hashes, CSS colors

9. **Base64()** - Standard Base64 (a-z, A-Z, 0-9, +, /)
   - 64 character pool
   - Use cases: Binary encoding, API authentication

10. **Base64URL()** - URL-safe Base64 (a-z, A-Z, 0-9, -, _)
    - No URL encoding needed
    - Use cases: JWT tokens, URL parameters

11. **Case Variants** - AlphaUpper, AlphaLower, AlnumUpper, AlnumLower
    - Uppercase/lowercase control
    - Use cases: Case-sensitive systems, formatting requirements

12. **Custom()** - User-defined character set
    - Custom charset: "ABC123!@#"
    - Flexible length range
    - Use cases: Custom requirements, specialized tokens

**Example Output Patterns / 예제 출력 패턴:**

```
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
1️⃣  Letters() - Alphabetic characters only
   알파벳 문자만 생성
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

📚 Function Signature / 함수 시그니처:
   func Letters(length ...int) (string, error)

📖 Description / 설명:
   Generates random string with uppercase and lowercase letters only
   대소문자 알파벳만 포함하는 랜덤 문자열을 생성합니다

🎯 Use Cases / 사용 사례:
   • Username generation (사용자명 생성)
   • Temporary identifiers (임시 식별자)
   • Alphabetic tokens (알파벳 토큰)
   • Name placeholders (이름 플레이스홀더)

💡 Key Features / 주요 기능:
   • Character set: a-z, A-Z (52 characters)
   • Variable length: 1 arg = fixed, 2 args = range
   • Cryptographically secure (crypto/rand)
   • URL-safe: no special characters

📊 Character Set Details / 문자 집합 상세:
   • Lowercase: a-z (26 characters)
   • Uppercase: A-Z (26 characters)
   • Total pool: 52 possible characters
   • Entropy: ~5.7 bits per character

▶️  Executing / 실행 중:
   str, err := random.GenString.Letters(8, 12)
   • Mode: Variable length (가변 길이)
   • Min length: 8 characters
   • Max length: 12 characters

✅ Generation Successful / 생성 성공
   📝 Result: mPiVtbbWsEB
   📏 Length: 11 characters
   🔤 Type: Alphabetic only
   ✓ In Range: true (8-12 characters)

🔍 Character Analysis / 문자 분석:
   • Lowercase letters: 6 (54.5%)
   • Uppercase letters: 5 (45.5%)
   • Total: 11
```

**Technical Implementation / 기술 구현:**
- Uses emoji icons for visual hierarchy and readability
- Bilingual logging (English/Korean) throughout
- Character analysis with percentage calculations
- Step-by-step execution flow documentation
- Real-world usage examples with context
- Security notes and best practices
- 이모지 아이콘을 사용하여 시각적 계층 구조 및 가독성 향상
- 전반에 걸쳐 이중 언어 로깅 (영문/한글)
- 백분율 계산이 포함된 문자 분석
- 단계별 실행 흐름 문서화
- 컨텍스트가 포함된 실제 사용 예제
- 보안 참고 사항 및 모범 사례

**Benefits / 이점:**
- **Self-Documenting**: Logs alone serve as complete tutorial
- **Educational**: Teaches cryptographic best practices
- **Comprehensive**: All 14 methods fully documented
- **Practical**: Real-world use cases for each method
- **Bilingual**: Accessible to both English and Korean speakers
- **자체 문서화**: 로그만으로도 완전한 튜토리얼 역할
- **교육적**: 암호학적 모범 사례 교육
- **포괄적**: 모든 14개 메서드 완전 문서화
- **실용적**: 각 메서드의 실제 사용 사례
- **이중 언어**: 영어 및 한국어 사용자 모두 접근 가능

**File Structure / 파일 구조:**
```
examples/random_string/main.go:
  - Package introduction (50+ lines)
  - Security features overview
  - 14 method demonstrations (750+ lines)
    - Each method: ~50-70 lines of detailed logging
  - Character analysis functions
  - Real-world usage examples
  - Best practices and security notes
```

**Testing / 테스트:**
- Verified successful execution with all methods
- Confirmed dual output (file + screen) with `WithStdout(true)`
- Log backup system working correctly
- All generated strings validated for length and character set
- 모든 메서드로 성공적인 실행 검증
- `WithStdout(true)`로 이중 출력 (파일 + 화면) 확인
- 로그 백업 시스템 정상 작동
- 생성된 모든 문자열의 길이 및 문자 집합 검증 완료

**Design Philosophy / 설계 철학:**
Following the fileutil example pattern: "Logs alone should serve as complete manual"
- No need to read external documentation
- Self-contained tutorial in log format
- Learn by reading execution traces
- fileutil 예제 패턴 따름: "로그만으로도 완전한 매뉴얼 역할"
- 외부 문서를 읽을 필요 없음
- 로그 형식의 자체 포함 튜토리얼
- 실행 추적을 읽으며 학습

---

## [v1.9.008] - 2025-10-15

### Fixed / 수정됨

#### Fixed defaultExample() to Use Consistent Logging Naming / defaultExample()이 일관된 로깅 명명 규칙을 사용하도록 수정

**Issue / 문제:**
- `defaultExample()` function was using `logging.Default()` which creates `logs/app.log`
- This was inconsistent with the `logging-example-` prefix used by all other examples
- `defaultExample()` 함수가 `logging.Default()`를 사용하여 `logs/app.log` 생성
- 다른 모든 예제에서 사용하는 `logging-example-` 접두사와 불일치

**Solution / 해결방법:**
- Changed `defaultExample()` to explicitly create logger with `logging-example-default.log` path
- Updated log message to reflect new file path
- Added `logging-example-default.log` to backup list in main()
- `defaultExample()`을 명시적으로 `logging-example-default.log` 경로로 로거 생성하도록 변경
- 로그 메시지를 새 파일 경로로 업데이트
- main()의 백업 목록에 `logging-example-default.log` 추가

**Technical Changes / 기술적 변경사항:**
```go
// Before (이전)
func defaultExample() {
    logger := logging.Default()
    defer logger.Close()
    logger.Info("Logs to ./logs/app.log by default")
}

// After (이후)
func defaultExample() {
    logger, _ := logging.New(
        logging.WithFilePath("logs/logging-example-default.log"),
    )
    defer logger.Close()
    logger.Info("Logs to ./logs/logging-example-default.log")
}
```

**Result / 결과:**
- Now all 13 log files in logging example use consistent `logging-example-` prefix
- No more `logs/app.log` file created by logging example
- Complete standardization across all example programs
- 이제 로깅 예제의 모든 13개 로그 파일이 일관된 `logging-example-` 접두사 사용
- 로깅 예제에서 더 이상 `logs/app.log` 파일 생성 안 함
- 모든 예제 프로그램에 걸쳐 완전한 표준화

**Total Log Files / 전체 로그 파일 개수:**
- Previous: 12 files (이전: 12개 파일)
- Current: 13 files (현재: 13개 파일)
- New file: `logs/logging-example-default.log`

---

## [v1.9.007] - 2025-10-15

### Changed / 변경됨

#### Standardized Logging Example Log File Naming Convention / 로깅 예제 로그 파일 명명 규칙 표준화

**Feature / 기능:**
- Updated all log file names in logging example to use consistent `logging-example-` prefix
- 로깅 예제의 모든 로그 파일명을 일관된 `logging-example-` 접두사를 사용하도록 업데이트

**Details / 상세 내용:**
- All 12 log files in `examples/logging/main.go` now follow the naming pattern: `logs/logging-example-{name}.log`
- This matches the naming convention used by all other example programs
- `examples/logging/main.go`의 모든 12개 로그 파일이 이제 명명 패턴을 따름: `logs/logging-example-{name}.log`
- 이는 다른 모든 예제 프로그램에서 사용하는 명명 규칙과 일치함

**Updated Log File Names / 업데이트된 로그 파일명:**
```
Before (이전)                      →  After (이후)
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
logs/app.log                       →  logs/logging-example-app.log
logs/custom.log                    →  logs/logging-example-custom.log
logs/database.log                  →  logs/logging-example-database.log
logs/api.log                       →  logs/logging-example-api.log
logs/levels.log                    →  logs/logging-example-levels.log
logs/structured.log                →  logs/logging-example-structured.log
logs/auto_banner_default.log       →  logs/logging-example-auto_banner_default.log
logs/auto_banner_custom.log        →  logs/logging-example-auto_banner_custom.log
logs/auto_banner_convenience.log   →  logs/logging-example-auto_banner_convenience.log
logs/auto_banner_disabled.log      →  logs/logging-example-auto_banner_disabled.log
logs/manual_banner_only.log        →  logs/logging-example-manual_banner_only.log
logs/banners.log                   →  logs/logging-example-banners.log
```

**Technical Changes / 기술적 변경사항:**
- Updated `logFiles` array in `main()` function with new file paths (12 entries)
- Updated all 12 `logging.WithFilePath()` calls across 7 example functions:
  - `customExample()` - 1 logger
  - `multipleLoggersExample()` - 3 loggers (app, database, api)
  - `logLevelsExample()` - 1 logger
  - `structuredLoggingExample()` - 1 logger
  - `autoBannerExample()` - 5 loggers (default, custom, convenience, disabled, manual)
  - `bannerExample()` - 1 logger
- Backup and cleanup functionality automatically works with new names
- `main()` 함수의 `logFiles` 배열을 새 파일 경로로 업데이트 (12개 항목)
- 7개 예제 함수의 모든 12개 `logging.WithFilePath()` 호출 업데이트
- 백업 및 정리 기능이 새 이름으로 자동 작동

**Modified Functions / 수정된 함수:**
1. `main()` - Updated logFiles array
2. `customExample()` - Updated file path
3. `multipleLoggersExample()` - Updated 3 file paths
4. `logLevelsExample()` - Updated file path
5. `structuredLoggingExample()` - Updated file path
6. `autoBannerExample()` - Updated 5 file paths
7. `bannerExample()` - Updated file path

**Benefits / 장점:**
- Consistent naming convention across all 9 example programs
- Easy identification of logging example files in logs directory
- Better organization and clarity
- Automatic backup management works seamlessly with new names
- 모든 9개 예제 프로그램에 걸쳐 일관된 명명 규칙
- logs 디렉토리에서 로깅 예제 파일 쉽게 식별
- 더 나은 조직화 및 명확성
- 자동 백업 관리가 새 이름으로 원활하게 작동

**Compatibility / 호환성:**
- Old log files are not automatically migrated (manual cleanup recommended)
- New runs will create files with new naming convention
- Backup system will handle new files from first run onwards
- 이전 로그 파일은 자동으로 마이그레이션되지 않음 (수동 정리 권장)
- 새 실행은 새 명명 규칙으로 파일 생성
- 백업 시스템은 첫 실행부터 새 파일 처리

---

## [v1.9.006] - 2025-10-15

### Added / 추가됨

#### Log File Backup Management System for All Examples / 모든 예제에 로그 파일 백업 관리 시스템 추가

**Feature / 기능:**
- Implemented automatic log file backup and cleanup system across all example programs
- 모든 예제 프로그램에 자동 로그 파일 백업 및 정리 시스템 구현

**Details / 상세 내용:**
- Each example now backs up previous log file with timestamp format `yyyymmdd-hhmmss` before creating new log
- Automatically maintains only 5 most recent backup files per example
- Old backup files are automatically deleted to prevent disk space issues
- 각 예제가 새 로그를 생성하기 전에 `yyyymmdd-hhmmss` 타임스탬프 형식으로 이전 로그 파일 백업
- 예제당 최근 5개의 백업 파일만 자동으로 유지
- 디스크 공간 문제를 방지하기 위해 오래된 백업 파일 자동 삭제

**Updated Examples / 업데이트된 예제:**
1. `examples/random_string/main.go` - Added backup logic for `random-example.log`
2. `examples/stringutil/main.go` - Added backup logic for `stringutil-example.log`
3. `examples/timeutil/main.go` - Added backup logic for `timeutil-example.log`
4. `examples/sliceutil/main.go` - Added backup logic for `sliceutil-example.log`
5. `examples/maputil/main.go` - Added backup logic for `maputil-example.log`
6. `examples/mysql/main.go` - Added backup logic for `mysql-example.log`
7. `examples/redis/main.go` - Added backup logic for `redis-example.log`
8. `examples/logging/main.go` - Added helper function `backupLogFile()` for 12 different log files
9. `examples/fileutil/main.go` - Already has backup logic (updated timestamp format to include seconds)

**Technical Implementation / 기술 구현:**
- Uses `fileutil.Exists()` to check for existing log files
- Uses `fileutil.ModTime()` to get file modification time for backup naming
- Uses `fileutil.CopyFile()` to create timestamped backups
- Uses `filepath.Glob()` to find and manage backup files
- Uses `fileutil.DeleteFile()` to clean up old backups
- Custom sorting algorithm to identify oldest files

**Benefits / 장점:**
- Prevents log file loss when running examples multiple times
- Maintains clean logs directory without manual intervention
- Provides historical log files for comparison and debugging
- 예제를 여러 번 실행할 때 로그 파일 손실 방지
- 수동 개입 없이 깨끗한 로그 디렉토리 유지
- 비교 및 디버깅을 위한 히스토리 로그 파일 제공

---

## [v1.9.005] - 2025-10-15

### Added / 추가됨

#### Extremely Detailed Logging in fileutil Examples / fileutil 예제에 극도로 상세한 로깅 추가

**Feature / 기능:**
- Enhanced `examples/fileutil/main.go` Examples 1-3 with comprehensive, tutorial-quality logging
- fileutil 예제 1-3을 포괄적이고 튜토리얼 수준의 로깅으로 개선

**Details / 상세 내용:**
- Each function now includes:
  - 📚 Function signature
  - 📖 Description (bilingual)
  - 🎯 Use cases (4+ scenarios)
  - 💡 Key features (4+ features)
  - ▶️ Step-by-step execution details
  - ✅ Success confirmation with detailed information
  - 🔍 Verification checks
- 각 함수가 이제 포함함:
  - 📚 함수 시그니처
  - 📖 설명 (이중 언어)
  - 🎯 사용 사례 (4개 이상)
  - 💡 주요 기능 (4개 이상)
  - ▶️ 단계별 실행 세부 정보
  - ✅ 상세 정보와 함께 성공 확인
  - 🔍 검증 체크

**Example 1 - File Writing Operations (8 functions):**
1. `WriteString()` - 55+ lines of detailed logging
2. `WriteFile()` - 50+ lines of detailed logging
3. `WriteLines()` - 60+ lines of detailed logging with content preview
4. `WriteJSON()` - 55+ lines of detailed logging with JSON content display
5. `WriteYAML()` - 50+ lines of detailed logging with YAML content display
6. `WriteCSV()` - 55+ lines of detailed logging with CSV preview
7. `AppendString()` - 50+ lines with before/after comparison
8. `AppendLines()` - 50+ lines with full content display

**Example 2 - File Reading Operations (6 functions):**
1. `ReadString()` - 50+ lines with content display
2. `ReadFile()` - 50+ lines with hex and ASCII representation
3. `ReadLines()` - 55+ lines with line-by-line content display
4. `ReadJSON()` - 50+ lines with parsed struct field display
5. `ReadYAML()` - 50+ lines with parsed struct and verification
6. `ReadCSV()` - 50+ lines with header and row display

**Example 3 - Path Operations (12 functions):**
1. `Join()` - 50+ lines for path element combination
2. `Split()` - 50+ lines for directory/file separation
3. `Base()` - 45+ lines for filename extraction
4. `Dir()` - 45+ lines for directory extraction
5. `Ext()` - 45+ lines for extension retrieval
6. `WithoutExt()` - 50+ lines for extension removal
7. `ChangeExt()` - 50+ lines for extension modification
8. `HasExt()` - 55+ lines for extension validation
9. `Abs()` - 45+ lines for absolute path resolution
10. `IsAbs()` - 50+ lines for path type checking
11. `CleanPath()` - 55+ lines for path normalization
12. `ToSlash/FromSlash()` - 55+ lines for separator conversion

**Benefits / 이점:**
- Users can understand the complete function behavior from logs alone without reading documentation
- 사용자가 문서를 읽지 않고도 로그만으로 완전한 함수 동작을 이해할 수 있음
- Serves as an interactive tutorial
- 대화형 튜토리얼 역할을 함
- Perfect for learning and debugging
- 학습 및 디버깅에 완벽함

### Changed / 변경됨

#### Unified Log File Path and Naming Convention / 통합된 로그 파일 경로 및 명명 규칙

**Problem / 문제:**
- Log files were scattered across different directories and had inconsistent naming conventions / 로그 파일이 여러 디렉토리에 분산되어 있고 일관성 없는 명명 규칙을 사용함
- Some examples used `logs/`, others used `./logs/` / 일부 예제는 `logs/`를, 다른 예제는 `./logs/`를 사용
- Filename formats varied: `package_example_`, `package-example-`, different timestamp formats / 파일명 형식이 다양함: `package_example_`, `package-example-`, 다양한 타임스탬프 형식

**Solution / 해결책:**
- Unified all example log files to repository root `logs/` directory / 모든 예제 로그 파일을 레포지토리 루트 `logs/` 디렉토리로 통합
- Standardized filename format: `<package>-example-<timestamp>.log` / 파일명 형식 표준화: `<package>-example-<timestamp>.log`
- Standardized timestamp format: `20060102-150405` (YYYYMMDD-HHMMSS) / 타임스탬프 형식 표준화: `20060102-150405` (YYYYMMDD-HHMMSS)
- Exception: logging package examples keep original filenames (but use unified `logs/` directory) / 예외: logging 패키지 예제는 원래 파일명 유지 (하지만 통합된 `logs/` 디렉토리 사용)

**Updated Files / 업데이트된 파일:**
1. `examples/fileutil/main.go` - `fileutil-examples-` → `fileutil-example-`
2. `examples/maputil/main.go` - Timestamp format: `20060102_150405` → `20060102-150405`
3. `examples/mysql/main.go` - `mysql_example_` + `20060102_150405` → `mysql-example-` + `20060102-150405`
4. `examples/random_string/main.go` - `random_example_` + `20060102_150405` → `random-example-` + `20060102-150405`
5. `examples/redis/main.go` - `redis_example_` + `20060102_150405` → `redis-example-` + `20060102-150405`
6. `examples/sliceutil/main.go` - `sliceutil_example_` + `20060102_150405` → `sliceutil-example-` + `20060102-150405`
7. `examples/stringutil/main.go` - `stringutil_example_` + `20060102_150405` → `stringutil-example-` + `20060102-150405`
8. `examples/timeutil/main.go` - `timeutil-example.log` (static) → `timeutil-example-<timestamp>.log` (timestamped)
9. `examples/logging/main.go` - All `./logs/` → `logs/` (filenames unchanged: `custom.log`, `app.log`, etc.)

**New Unified Format / 새로운 통합 형식:**
```
logs/fileutil-example-20251015-200641.log
logs/maputil-example-20251015-143022.log
logs/mysql-example-20251015-143022.log
logs/random-example-20251015-143022.log
logs/redis-example-20251015-143022.log
logs/sliceutil-example-20251015-143022.log
logs/stringutil-example-20251015-143022.log
logs/timeutil-example-20251015-143022.log
logs/custom.log          (logging example - filename preserved)
logs/app.log             (logging example - filename preserved)
logs/database.log        (logging example - filename preserved)
```

**Benefits / 이점:**
- Centralized log management in single `logs/` directory / 단일 `logs/` 디렉토리에서 중앙 집중식 로그 관리
- Consistent naming convention across all examples / 모든 예제에서 일관된 명명 규칙
- Easier log file discovery and organization / 로그 파일 검색 및 정리가 더 쉬움
- Predictable log file locations for CI/CD and automation / CI/CD 및 자동화를 위한 예측 가능한 로그 파일 위치
- Timestamp in filename enables chronological sorting / 파일명의 타임스탬프로 시간순 정렬 가능

---

## [v1.9.004] - 2025-10-15

### Enhanced / 보강됨

#### Complete Logging Migration in Fileutil Example / Fileutil 예제의 완전한 로깅 마이그레이션

**Updated Files / 업데이트된 파일:**
- `examples/fileutil/main.go` - Completely replaced all fmt output with structured logging / 모든 fmt 출력을 구조화된 로깅으로 완전히 교체

**Key Changes / 주요 변경사항:**
- Replaced all `fmt.Println()` and `fmt.Printf()` calls with structured `logger.Info()` calls / 모든 `fmt.Println()` 및 `fmt.Printf()` 호출을 구조화된 `logger.Info()` 호출로 교체
- Implemented key-value structured logging for all output messages / 모든 출력 메시지에 대해 키-값 구조화 로깅 구현
- Added logger parameter to `example2PathOperations()` function / `example2PathOperations()` 함수에 logger 매개변수 추가
- Used consistent key naming: `path`, `count`, `value`, `bytes`, `hash`, `checksum`, `valid`, `same`, `exists`, `isEmpty`, `name` / 일관된 키 이름 사용
- Eliminated duplicate output (removed redundant fmt.Println after logger.Info) / 중복 출력 제거 (logger.Info 이후 중복된 fmt.Println 제거)
- Maintained fmt.Printf only for progress callback (line 208) which cannot easily access logger / 로거에 쉽게 접근할 수 없는 진행 상황 콜백(208행)에만 fmt.Printf 유지

**Benefits / 이점:**
- All output now appears in both console and log file thanks to `WithStdout(true)` / `WithStdout(true)` 덕분에 모든 출력이 콘솔과 로그 파일 양쪽에 표시됨
- Structured logging allows easier parsing and analysis of logs / 구조화된 로깅으로 로그를 더 쉽게 파싱하고 분석 가능
- Consistent logging pattern across all example functions / 모든 예제 함수에서 일관된 로깅 패턴
- Better observability with key-value pairs / 키-값 쌍으로 더 나은 관찰 가능성

**Example Output / 예제 출력:**
```
2025-10-15 19:56:38 [INFO] ✓ Written to file path=/path/to/file.txt
2025-10-15 19:56:38 [INFO] ✓ Found files count=3
2025-10-15 19:56:38 [INFO] ✓ SHA256 hash hash=7d5e51fa...
2025-10-15 19:56:38 [INFO] ✓ Directory status exists=true isEmpty=true
```

---

## [v1.9.003] - 2025-10-15

### Enhanced / 보강됨

#### Example Files Logging Integration / 예제 파일 로깅 통합

**Updated Files / 업데이트된 파일:**
- `examples/fileutil/main.go` - Integrated logging package with timestamped log files / 타임스탬프가 있는 로그 파일로 logging 패키지 통합

**Key Changes / 주요 변경사항:**
- Replaced all `log.Fatal` calls with `logger.Fatalf` in fileutil example / fileutil 예제의 모든 `log.Fatal` 호출을 `logger.Fatalf`로 교체
- Updated all example functions to accept `logger *logging.Logger` parameter / 모든 예제 함수에서 `logger *logging.Logger` 매개변수를 받도록 업데이트
- Removed unused `log` package import / 사용하지 않는 `log` 패키지 import 제거
- Log files now saved to `logs/` directory with pattern: `logs/fileutil-examples-{timestamp}.log` / 로그 파일이 이제 `logs/` 디렉토리에 `logs/fileutil-examples-{timestamp}.log` 패턴으로 저장됨
- Added `WithStdout(true)` for console output alongside file logging / 파일 로깅과 함께 콘솔 출력을 위한 `WithStdout(true)` 추가

**Verified / 확인됨:**
- All utility package examples (stringutil, timeutil, sliceutil, maputil) already use logging package correctly / 모든 유틸리티 패키지 예제(stringutil, timeutil, sliceutil, maputil)가 이미 logging 패키지를 올바르게 사용함
- All examples build successfully / 모든 예제가 성공적으로 빌드됨
- Fileutil example tested and confirmed working with logging package / Fileutil 예제가 logging 패키지와 함께 테스트되고 작동 확인됨

---

## [v1.9.002] - 2025-10-15

### Enhanced / 보강됨

#### Comprehensive Test Suite / 포괄적인 테스트 스위트

**Test Coverage / 테스트 커버리지:**
- **87 test cases** across 9 test suites / 9개 테스트 스위트에 걸쳐 87개 테스트 케이스
- **10 benchmark tests** for performance measurement / 성능 측정을 위한 10개 벤치마크 테스트
- **55.2% code coverage** / 55.2% 코드 커버리지

**Test Suites / 테스트 스위트:**
1. **TestFileReading** (8 tests): ReadFile, ReadString, ReadLines, ReadJSON, ReadYAML, ReadCSV, ReadBytes, ReadChunk
2. **TestFileWriting** (9 tests): WriteFile, WriteString, WriteLines, WriteJSON, WriteYAML, WriteCSV, WriteAtomic, AppendFile, AppendLines
3. **TestPathOperations** (17 tests): Join, Split, Base, Dir, Ext, Abs, CleanPath, Normalize, ToSlash, FromSlash, IsAbs, IsValid, IsSafe, Match, WithoutExt, ChangeExt, HasExt
4. **TestFileInformation** (11 tests): Exists, IsFile, IsDir, IsSymlink, Size, SizeHuman, Chmod, IsReadable, IsWritable, ModTime, Touch
5. **TestCopyOperations** (4 tests): CopyFile, CopyFile_WithOverwrite, CopyFile_WithProgress, CopyDir
6. **TestMoveOperations** (3 tests): MoveFile, Rename, RenameExt
7. **TestDeleteOperations** (7 tests): DeleteFile, DeleteDir, DeleteRecursive, DeletePattern, DeleteFiles, Clean, RemoveEmpty
8. **TestDirectoryOperations** (9 tests): MkdirAll, CreateTemp, CreateTempDir, IsEmpty, DirSize, ListFiles, ListDirs, ListAll, FindFiles
9. **TestHashOperations** (10 tests): MD5, SHA1, SHA256, SHA512, Hash, HashBytes, CompareFiles, CompareHash, Checksum, VerifyChecksum

**Benchmark Tests / 벤치마크 테스트:**
1. `BenchmarkWriteFile` - File writing performance / 파일 쓰기 성능
2. `BenchmarkReadFile` - File reading performance / 파일 읽기 성능
3. `BenchmarkWriteString` - String writing performance / 문자열 쓰기 성능
4. `BenchmarkReadString` - String reading performance / 문자열 읽기 성능
5. `BenchmarkCopyFile` - File copying performance / 파일 복사 성능
6. `BenchmarkSHA256` - SHA256 hashing performance / SHA256 해싱 성능
7. `BenchmarkMD5` - MD5 hashing performance / MD5 해싱 성능
8. `BenchmarkJSON/WriteJSON` - JSON writing performance / JSON 쓰기 성능
9. `BenchmarkJSON/ReadJSON` - JSON reading performance / JSON 읽기 성능
10. `BenchmarkYAML/WriteYAML` - YAML writing performance / YAML 쓰기 성능
11. `BenchmarkYAML/ReadYAML` - YAML reading performance / YAML 읽기 성능
12. `BenchmarkListFiles` - Directory listing performance / 디렉토리 나열 성능

**Benchmark Results / 벤치마크 결과 (darwin/amd64):**
- WriteFile: ~52,000 ns/op, 584 B/op, 6 allocs/op
- ReadFile: ~12,000 ns/op, 920 B/op, 5 allocs/op
- CopyFile (10KB): ~164,000 ns/op, 34KB/op, 22 allocs/op
- SHA256 (10KB): ~56,000 ns/op, 33KB/op, 9 allocs/op
- MD5 (10KB): ~42,000 ns/op, 33KB/op, 9 allocs/op
- JSON Write: ~75,000 ns/op, 888 B/op, 14 allocs/op
- JSON Read: ~23,000 ns/op, 1.7KB/op, 29 allocs/op
- YAML Write: ~50,000 ns/op, 7.4KB/op, 33 allocs/op
- YAML Read: ~22,000 ns/op, 8.7KB/op, 61 allocs/op
- ListFiles (100 files): ~86,000 ns/op, 23KB/op, 244 allocs/op

**Test Features / 테스트 기능:**
- Comprehensive edge case coverage / 포괄적인 엣지 케이스 커버리지
- Automatic cleanup with defer / defer를 사용한 자동 정리
- Temporary directory usage for isolation / 격리를 위한 임시 디렉토리 사용
- Error handling validation / 에러 처리 검증
- Cross-platform path handling / 크로스 플랫폼 경로 처리
- Progress callback testing / 진행 상황 콜백 테스트
- Copy options testing (overwrite, progress, filter) / 복사 옵션 테스트
- All hash algorithms tested / 모든 해시 알고리즘 테스트

---

## [v1.9.001] - 2025-10-15

### Added / 추가됨

#### Fileutil Package - Complete Implementation / Fileutil 패키지 - 완전한 구현

**Package Structure / 패키지 구조:**
- `fileutil/fileutil.go` - Package documentation and constants (version v1.9.001)
- `fileutil/errors.go` - Custom error types and helper functions
- `fileutil/options.go` - Functional options pattern for copy operations
- `fileutil/read.go` - File reading functions (8 functions)
- `fileutil/write.go` - File writing and appending functions (11 functions)
- `fileutil/info.go` - File information functions (15 functions)
- `fileutil/path.go` - Path manipulation functions (18 functions)
- `fileutil/copy.go` - File/directory copying functions (4 functions)
- `fileutil/move.go` - File/directory moving functions (5 functions)
- `fileutil/delete.go` - File/directory deletion functions (7 functions)
- `fileutil/dir.go` - Directory operation functions (13 functions)
- `fileutil/hash.go` - File hashing functions (10 functions)
- `fileutil/fileutil_test.go` - Comprehensive test suite (7 test suites, 2 benchmarks)
- `fileutil/README.md` - Package documentation (bilingual)

**File Reading (8 functions) / 파일 읽기 (8개 함수):**
1. `ReadFile(path string) ([]byte, error)` - Read entire file
2. `ReadString(path string) (string, error)` - Read file as string
3. `ReadLines(path string) ([]string, error)` - Read file as lines
4. `ReadJSON(path string, v interface{}) error` - Read and unmarshal JSON
5. `ReadYAML(path string, v interface{}) error` - Read and unmarshal YAML
6. `ReadCSV(path string) ([][]string, error)` - Read CSV file
7. `ReadBytes(path string, offset, length int64) ([]byte, error)` - Read specific bytes
8. `ReadChunk(path string, chunkSize int64, fn func([]byte) error) error` - Read file in chunks

**File Writing (11 functions) / 파일 쓰기 (11개 함수):**
1. `WriteFile(path string, data []byte) error` - Write bytes to file
2. `WriteString(path string, s string) error` - Write string to file
3. `WriteLines(path string, lines []string) error` - Write lines to file
4. `WriteJSON(path string, v interface{}) error` - Marshal and write JSON
5. `WriteYAML(path string, v interface{}) error` - Marshal and write YAML
6. `WriteCSV(path string, records [][]string) error` - Write CSV file
7. `WriteAtomic(path string, data []byte) error` - Atomic write (temp + rename)
8. `AppendFile(path string, data []byte) error` - Append bytes to file
9. `AppendString(path string, s string) error` - Append string to file
10. `AppendLines(path string, lines []string) error` - Append lines to file
11. `AppendBytes(path string, data []byte) error` - Append bytes (alias)

**File Information (15 functions) / 파일 정보 (15개 함수):**
1. `Exists(path string) bool` - Check existence
2. `IsFile(path string) bool` - Check if file
3. `IsDir(path string) bool` - Check if directory
4. `IsSymlink(path string) bool` - Check if symlink
5. `Size(path string) (int64, error)` - Get file size
6. `SizeHuman(path string) (string, error)` - Get human-readable size
7. `Chmod(path string, mode os.FileMode) error` - Change permissions
8. `Chown(path string, uid, gid int) error` - Change owner
9. `IsReadable(path string) bool` - Check if readable
10. `IsWritable(path string) bool` - Check if writable
11. `IsExecutable(path string) bool` - Check if executable
12. `ModTime(path string) (time.Time, error)` - Get modification time
13. `AccessTime(path string) (time.Time, error)` - Get access time
14. `ChangeTime(path string) (time.Time, error)` - Get change time
15. `Touch(path string) error` - Update modification time

**Path Operations (18 functions) / 경로 작업 (18개 함수):**
1. `Join(elem ...string) string` - Join path elements
2. `Split(path string) (string, string)` - Split into dir and file
3. `Base(path string) string` - Get base name
4. `Dir(path string) string` - Get directory
5. `Ext(path string) string` - Get extension
6. `Abs(path string) (string, error)` - Get absolute path
7. `CleanPath(path string) string` - Clean path (renamed from Clean)
8. `Normalize(path string) (string, error)` - Normalize path
9. `ToSlash(path string) string` - Convert to forward slashes
10. `FromSlash(path string) string` - Convert to OS-specific
11. `IsAbs(path string) bool` - Check if absolute
12. `IsValid(path string) bool` - Validate path
13. `IsSafe(path, root string) bool` - Check path safety
14. `Match(pattern, name string) (bool, error)` - Match pattern
15. `Glob(pattern string) ([]string, error)` - Find by glob pattern
16. `Rel(basepath, targpath string) (string, error)` - Get relative path
17. `WithoutExt(path string) string` - Remove extension
18. `ChangeExt(path, newExt string) string` - Change extension
19. `HasExt(path string, exts ...string) bool` - Check extension

**File Copying (4 functions) / 파일 복사 (4개 함수):**
1. `CopyFile(src, dst string, opts ...CopyOption) error` - Copy single file
2. `CopyDir(src, dst string, opts ...CopyOption) error` - Copy directory
3. `CopyRecursive(src, dst string, opts ...CopyOption) error` - Copy recursively
4. `SyncDirs(src, dst string, opts ...CopyOption) error` - Sync two directories

**Copy Options / 복사 옵션:**
- `WithOverwrite(bool)` - Overwrite existing files
- `WithPreservePermissions(bool)` - Preserve file permissions
- `WithPreserveTimestamps(bool)` - Preserve timestamps
- `WithProgress(func(written, total int64))` - Progress callback
- `WithFilter(func(path string, info os.FileInfo) bool)` - File filter

**File Moving (5 functions) / 파일 이동 (5개 함수):**
1. `MoveFile(src, dst string) error` - Move file
2. `MoveDir(src, dst string) error` - Move directory
3. `Rename(oldPath, newPath string) error` - Rename file/directory
4. `RenameExt(path, newExt string) (string, error)` - Change file extension
5. `SafeMove(src, dst string) error` - Move with existence check

**File Deleting (7 functions) / 파일 삭제 (7개 함수):**
1. `DeleteFile(path string) error` - Delete single file
2. `DeleteDir(path string) error` - Delete empty directory
3. `DeleteRecursive(path string) error` - Delete recursively
4. `DeletePattern(dir, pattern string) error` - Delete files by pattern
5. `DeleteFiles(paths []string) error` - Delete multiple files
6. `Clean(path string) error` - Remove directory contents
7. `RemoveEmpty(path string) error` - Remove empty directories

**Directory Operations (13 functions) / 디렉토리 작업 (13개 함수):**
1. `MkdirAll(path string) error` - Create directory tree
2. `CreateTemp(dir, pattern string) (*os.File, error)` - Create temp file
3. `CreateTempDir(dir, pattern string) (string, error)` - Create temp directory
4. `IsEmpty(path string) (bool, error)` - Check if directory is empty
5. `DirSize(path string) (int64, error)` - Calculate directory size
6. `ListFiles(dir string) ([]string, error)` - List files only
7. `ListDirs(dir string) ([]string, error)` - List directories only
8. `ListAll(dir string) ([]string, error)` - List all entries
9. `Walk(root string, fn filepath.WalkFunc) error` - Walk directory tree
10. `WalkFiles(root string, fn func(string, os.FileInfo) error) error` - Walk files only
11. `WalkDirs(root string, fn func(string, os.FileInfo) error) error` - Walk directories only
12. `FindFiles(root string, filter func(string, interface{}) bool) ([]string, error)` - Find files by filter
13. `FilterFiles(root string, patterns []string) ([]string, error)` - Filter files by patterns

**File Hashing (10 functions) / 파일 해싱 (10개 함수):**
1. `MD5(path string) (string, error)` - Calculate MD5 hash
2. `SHA1(path string) (string, error)` - Calculate SHA1 hash
3. `SHA256(path string) (string, error)` - Calculate SHA256 hash
4. `SHA512(path string) (string, error)` - Calculate SHA512 hash
5. `Hash(path, algorithm string) (string, error)` - Calculate hash by algorithm
6. `HashBytes(data []byte, algorithm string) (string, error)` - Hash byte slice
7. `CompareFiles(path1, path2 string) (bool, error)` - Compare files byte-by-byte
8. `CompareHash(path1, path2 string) (bool, error)` - Compare files by hash
9. `Checksum(path string) (string, error)` - Calculate checksum (SHA256)
10. `VerifyChecksum(path, expected string) (bool, error)` - Verify checksum

**Custom Error Types / 사용자 정의 에러 타입:**
- `ErrNotFound` - File or directory not found
- `ErrNotFile` - Path is not a file
- `ErrNotDirectory` - Path is not a directory
- `ErrInvalidPath` - Invalid path
- `ErrPermissionDenied` - Permission denied
- `ErrAlreadyExists` - File or directory already exists
- `ErrNotEmpty` - Directory is not empty

**Error Helper Functions / 에러 헬퍼 함수:**
- `IsNotFound(err error) bool` - Check if error is "not found"
- `IsPermission(err error) bool` - Check if error is "permission denied"
- `IsExist(err error) bool` - Check if error is "already exists"
- `IsInvalid(err error) bool` - Check if error is "invalid path"

**Documentation / 문서:**
- Package README with quick start guide and function reference (bilingual)
- Comprehensive examples in `examples/fileutil/main.go` (7 example scenarios)
- Design plan document (`docs/fileutil/DESIGN_PLAN.md`)
- Work plan document (`docs/fileutil/WORK_PLAN.md`)

**Testing / 테스트:**
- 7 comprehensive test suites covering all function categories
- 2 benchmark tests for read/write operations
- 100% pass rate on all tests
- Test coverage across all major functionality

**Examples / 예제:**
- File Writing and Reading
- Path Operations
- File Information
- File Copying (with progress callback)
- File Hashing
- Directory Operations
- File Deletion

### Key Features / 주요 기능

1. **Automatic Directory Creation / 자동 디렉토리 생성:**
   - All write operations automatically create parent directories if they don't exist
   - 모든 쓰기 작업은 상위 디렉토리가 존재하지 않으면 자동으로 생성합니다

2. **Cross-Platform Compatibility / 크로스 플랫폼 호환성:**
   - All path operations use `filepath` package for OS-agnostic behavior
   - 모든 경로 작업은 OS에 구애받지 않는 동작을 위해 `filepath` 패키지를 사용합니다

3. **Buffered I/O / 버퍼링된 I/O:**
   - Default 32KB buffer size for optimal performance
   - 최적의 성능을 위한 기본 32KB 버퍼 크기

4. **Atomic Operations / 원자적 작업:**
   - `WriteAtomic` function for safe file updates (write to temp, then rename)
   - 안전한 파일 업데이트를 위한 `WriteAtomic` 함수 (임시 파일에 쓰기, 그 다음 이름 변경)

5. **Progress Callbacks / 진행 상황 콜백:**
   - Copy operations support progress callbacks for large files
   - 복사 작업은 대용량 파일에 대한 진행 상황 콜백을 지원합니다

6. **Flexible Copying / 유연한 복사:**
   - Functional options pattern for copy operations
   - 복사 작업을 위한 함수형 옵션 패턴
   - Options: overwrite, preserve permissions, preserve timestamps, progress, filter
   - 옵션: 덮어쓰기, 권한 보존, 타임스탬프 보존, 진행 상황, 필터

7. **Multiple Hash Algorithms / 여러 해시 알고리즘:**
   - Support for MD5, SHA1, SHA256, SHA512
   - MD5, SHA1, SHA256, SHA512 지원

8. **Path Safety / 경로 안전성:**
   - `IsSafe` function to prevent directory traversal attacks
   - 디렉토리 탐색 공격을 방지하기 위한 `IsSafe` 함수

9. **Human-Readable Output / 사람이 읽기 쉬운 출력:**
   - `SizeHuman` converts bytes to KB/MB/GB/TB
   - `SizeHuman`은 바이트를 KB/MB/GB/TB로 변환합니다

10. **Zero External Dependencies / 외부 의존성 없음:**
    - Only uses standard library (except gopkg.in/yaml.v3)
    - 표준 라이브러리만 사용 (gopkg.in/yaml.v3 제외)

### Design Philosophy / 설계 철학

**"20줄 → 1줄" (20 lines → 1 line)**

Reducing repetitive file manipulation code from 20+ lines to just 1-2 lines:

반복적인 파일 조작 코드를 20줄 이상에서 단 1-2줄로 줄입니다:

```go
// Before: 20+ lines
dir := filepath.Dir(path)
if err := os.MkdirAll(dir, 0755); err != nil {
    return err
}
file, err := os.Create(path)
if err != nil {
    return err
}
defer file.Close()
if _, err := file.WriteString(content); err != nil {
    return err
}

// After: 1 line
err := fileutil.WriteString(path, content)
```

### Fixed / 수정됨

1. **Function Name Collision / 함수 이름 충돌:**
   - Renamed `Clean` in `path.go` to `CleanPath` to avoid conflict with `Clean` in `delete.go`
   - `path.go`의 `Clean`을 `CleanPath`로 이름 변경하여 `delete.go`의 `Clean`과 충돌 방지
   - Updated references in `Normalize()` and `IsSafe()` functions
   - `Normalize()` 및 `IsSafe()` 함수의 참조 업데이트

2. **Missing Import / 누락된 임포트:**
   - Added `path/filepath` import to `hash.go` for `HashDir` function
   - `HashDir` 함수를 위해 `hash.go`에 `path/filepath` 임포트 추가

### Technical Details / 기술 세부사항

- **Go Version**: 1.18+ (uses standard library generics where appropriate)
- **Dependencies**: Standard library + `gopkg.in/yaml.v3`
- **Constants**:
  - `DefaultFileMode = 0644` - Default file permissions
  - `DefaultDirMode = 0755` - Default directory permissions
  - `DefaultBufferSize = 32 * 1024` - 32KB buffer for I/O operations
  - `DefaultChunkSize = 1024 * 1024` - 1MB chunk size for large file processing

### Performance / 성능

- Buffered I/O for all file operations (32KB buffer)
- 모든 파일 작업에 버퍼링된 I/O 사용 (32KB 버퍼)
- Efficient chunk-based processing for large files (1MB chunks)
- 대용량 파일에 대한 효율적인 청크 기반 처리 (1MB 청크)
- Optimized directory walking with filter support
- 필터 지원을 통한 최적화된 디렉토리 순회

---

**Total Functions Implemented / 구현된 총 함수 수: ~91 functions across 12 categories**

**Total Functions Implemented / 구현된 총 함수 수: 12개 카테고리에 걸쳐 약 91개 함수**
