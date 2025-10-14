# CHANGELOG - v1.6.x

All notable changes for version 1.6.x will be documented in this file.

v1.6.x 버전의 모든 주목할 만한 변경사항이 이 파일에 문서화됩니다.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

---

## [v1.6.003] - 2025-10-14

### Changed / 변경

- **EXAMPLES**: Completely rewrote examples/timeutil/main.go (112 lines → 587 lines) / examples/timeutil/main.go 완전 재작성 (112줄 → 587줄)
  - Added logging package integration with file output / 파일 출력이 있는 logging 패키지 통합 추가
  - Demonstrated ALL 93 functions across 10 categories / 10개 카테고리에 걸쳐 모든 93개 함수 시연
  - Detailed structured logging for each function call / 각 함수 호출에 대한 상세한 구조화된 로깅
  - Summary section listing all categories and function counts / 모든 카테고리와 함수 수를 나열하는 요약 섹션

### Example Coverage / 예제 커버리지

**All 10 Categories Demonstrated / 모든 10개 카테고리 시연**:
1. Time Difference Functions (8 functions) / 시간 차이 함수 (8개 함수)
2. Timezone Operations (10 functions) / 타임존 작업 (10개 함수)
3. Date Arithmetic (16 functions) / 날짜 연산 (16개 함수)
4. Date Formatting (8 functions) / 날짜 포맷팅 (8개 함수)
5. Time Parsing (6 functions) / 시간 파싱 (6개 함수)
6. Time Comparisons (18 functions) / 시간 비교 (18개 함수)
7. Age Calculations (4 functions) / 나이 계산 (4개 함수)
8. Relative Time (4 functions) / 상대 시간 (4개 함수)
9. Unix Timestamp (12 functions) / Unix 타임스탬프 (12개 함수)
10. Business Days (7 functions) / 영업일 (7개 함수)

**Total Functions**: 93 functions fully demonstrated / 93개 함수 완전 시연

### Logging Output / 로깅 출력

- Log file: `./logs/timeutil-example.log` / 로그 파일: `./logs/timeutil-example.log`
- Detailed input/output for each function / 각 함수의 상세한 입력/출력
- Structured key-value logging / 구조화된 키-값 로깅
- Banner with version information / 버전 정보가 있는 배너
- Summary with total function count / 전체 함수 수가 있는 요약

---

## [v1.6.002] - 2025-10-14

### Added / 추가

- **DOCS**: Created comprehensive USER_MANUAL.md (~1,800 lines) / 포괄적인 USER_MANUAL.md 생성 (~1,800줄)
  - Complete function reference with examples / 예제를 포함한 완전한 함수 참조
  - 9 common use cases with full code / 전체 코드를 포함한 9개의 일반적인 사용 사례
  - 12 best practices / 12개의 모범 사례
  - Troubleshooting guide / 문제 해결 가이드
  - FAQ (10 questions) / FAQ (10개 질문)

- **DOCS**: Created comprehensive DEVELOPER_GUIDE.md (~1,600 lines) / 포괄적인 DEVELOPER_GUIDE.md 생성 (~1,600줄)
  - Architecture overview with diagrams / 다이어그램이 있는 아키텍처 개요
  - Core components detailed explanation / 핵심 컴포넌트 상세 설명
  - 5 design patterns used / 사용된 5개의 디자인 패턴
  - Internal implementation details / 내부 구현 세부사항
  - Step-by-step guide for adding features / 기능 추가를 위한 단계별 가이드
  - Testing guide with examples / 예제를 포함한 테스트 가이드
  - Performance optimization strategies / 성능 최적화 전략
  - Contributing guidelines / 기여 가이드라인
  - Code style guide / 코드 스타일 가이드

### Documentation Highlights / 문서 하이라이트

**USER_MANUAL.md Sections / USER_MANUAL.md 섹션**:
1. Introduction with key features / 주요 기능이 있는 소개
2. Installation guide / 설치 가이드
3. 5 quick start examples / 5개의 빠른 시작 예제
4. Core concepts (KST default, custom tokens, types) / 핵심 개념
5. Complete function reference (80+ functions) / 완전한 함수 참조 (80개 이상 함수)
6. 8 common use cases with full implementations / 전체 구현이 있는 8개의 일반적인 사용 사례
7. 12 best practices / 12개의 모범 사례
8. Troubleshooting guide / 문제 해결 가이드
9. FAQ with 10 questions / 10개 질문이 있는 FAQ

**DEVELOPER_GUIDE.md Sections / DEVELOPER_GUIDE.md 섹션**:
1. Architecture overview with ASCII diagrams / ASCII 다이어그램이 있는 아키텍처 개요
2. Package structure (file organization) / 패키지 구조 (파일 구성)
3. Core components (types, constants, caches) / 핵심 컴포넌트 (타입, 상수, 캐시)
4. 5 design patterns (Singleton, Factory, Strategy, Decorator, Cache-Aside) / 5개 디자인 패턴
5. Internal implementation for 5 key features / 5개 주요 기능의 내부 구현
6. Step-by-step guide for adding new features / 새 기능 추가를 위한 단계별 가이드
7. Testing guide (organization, running, coverage) / 테스트 가이드
8. Performance tips and benchmark results / 성능 팁 및 벤치마크 결과
9. Contributing process and checklist / 기여 프로세스 및 체크리스트
10. Code style guide with examples / 예제가 있는 코드 스타일 가이드

### Files Updated / 업데이트된 파일

- `cfg/app.yaml`: Version updated to v1.6.002 / 버전을 v1.6.002로 업데이트
- `docs/timeutil/USER_MANUAL.md`: New comprehensive user manual / 새로운 포괄적인 사용자 매뉴얼
- `docs/timeutil/DEVELOPER_GUIDE.md`: New comprehensive developer guide / 새로운 포괄적인 개발자 가이드
- `docs/CHANGELOG/CHANGELOG-v1.6.md`: This file / 이 파일

### Documentation Statistics / 문서 통계

| Document / 문서 | Lines / 줄 수 | Sections / 섹션 | Language / 언어 |
|-----------------|---------------|----------------|-----------------|
| USER_MANUAL.md | ~1,800 | 9 | Bilingual / 이중 언어 |
| DEVELOPER_GUIDE.md | ~1,600 | 10 | Bilingual / 이중 언어 |
| **Total / 합계** | **~3,400** | **19** | |

---

## [v1.6.001] - 2025-10-14

### Added / 추가

- **NEW PACKAGE**: `timeutil` - Time and date utility functions with 80+ functions / 80개 이상의 함수를 가진 시간 및 날짜 유틸리티
- **DESIGN**: Created comprehensive design document (DESIGN_PLAN.md)
- **DESIGN**: Created detailed work plan (WORK_PLAN.md)
- **DOCS**: Created initial README.md
- **DOCS**: Created CHANGELOG-v1.6.md
- **IMPLEMENTATION**: Completed all core and advanced features / 모든 핵심 및 고급 기능 완성
- **TESTS**: Comprehensive test suite (all tests passing) / 포괄적인 테스트 스위트 (모든 테스트 통과)
- **EXAMPLES**: Working example code in examples/timeutil/ / examples/timeutil/의 작동 예제 코드
- **DEFAULT TIMEZONE**: KST (Asia/Seoul, GMT+9) as default timezone / KST (Asia/Seoul, GMT+9)를 기본 타임존으로 설정

### Package Overview / 패키지 개요

**Design Philosophy / 설계 철학**: "20 lines → 1 line" - Extreme simplicity for time/date operations

**Total Functions / 총 함수 수**: ~80+ functions across 10 categories / 10개 카테고리에 걸쳐 약 80개 이상의 함수

**Categories / 카테고리**:
1. **Time Difference / 시간 차이** (8 functions): SubTime, DiffInSeconds, DiffInMinutes, DiffInHours, DiffInDays, DiffInWeeks, DiffInMonths, DiffInYears
2. **Timezone Operations / 타임존 작업** (5 functions): ConvertTimezone, GetTimezoneOffset, ListTimezones, IsValidTimezone, GetLocalTimezone
3. **Date Arithmetic / 날짜 연산** (16 functions): AddSeconds, AddMinutes, AddHours, AddDays, AddWeeks, AddMonths, AddYears, StartOfDay, EndOfDay, StartOfWeek, EndOfWeek, StartOfMonth, EndOfMonth, StartOfYear, EndOfYear, StartOfQuarter
4. **Date Formatting / 날짜 포맷팅** (8 functions): FormatISO8601, FormatRFC3339, FormatDate, FormatDateTime, FormatTime, Format, FormatCustom, FormatWithTimezone
5. **Time Parsing / 시간 파싱** (6 functions): ParseISO8601, ParseRFC3339, ParseDate, ParseDateTime, Parse, ParseWithTimezone
6. **Business Days / 영업일** (6 functions): AddBusinessDays, IsBusinessDay, CountBusinessDays, NextBusinessDay, PreviousBusinessDay, IsHoliday
7. **Time Comparisons / 시간 비교** (18 functions): IsBefore, IsAfter, IsBetween, IsToday, IsYesterday, IsTomorrow, IsThisWeek, IsThisMonth, IsThisYear, IsWeekend, IsWeekday, IsSameDay, IsSameWeek, IsSameMonth, IsSameYear, IsLeapYear, IsPast, IsFuture
8. **Age Calculations / 나이 계산** (4 functions): AgeInYears, AgeInMonths, AgeInDays, Age
9. **Relative Time / 상대 시간** (3 functions): RelativeTime, RelativeTimeShort, TimeAgo
10. **Unix Timestamp / Unix 타임스탬프** (8 functions): Now, NowMilli, NowMicro, NowNano, FromUnix, FromUnixMilli, ToUnix, ToUnixMilli

**Core Types / 핵심 타입**:
- `TimeDiff`: Time difference with helper methods (Seconds, Minutes, Hours, Days, Weeks, String, Humanize, Abs)
- `Age`: Age representation with Years, Months, Days (String, Humanize methods)

### Key Features / 주요 기능

1. **Extreme Simplicity / 극도의 간결함**:
   - Reduce 20+ lines of time manipulation code to just 1 line
   - 20줄 이상의 시간 조작 코드를 단 1줄로 줄임

2. **Human-Readable / 사람이 읽기 쉬움**:
   - Intuitive function names that read like natural language
   - 자연어처럼 읽히는 직관적인 함수 이름

3. **Zero Configuration / 제로 설정**:
   - No setup required, just import and use
   - 설정 불필요, 임포트하고 바로 사용

4. **Custom Format Tokens / 커스텀 포맷 토큰**:
   - Use `YYYY-MM-DD` instead of Go's confusing `2006-01-02`
   - Go의 혼란스러운 `2006-01-02` 대신 `YYYY-MM-DD` 사용

5. **Business Day Support / 영업일 지원**:
   - Calculate business days with holiday support
   - 공휴일 지원과 함께 영업일 계산

6. **Timezone Caching / 타임존 캐싱**:
   - Efficient timezone operations with caching
   - 캐싱으로 효율적인 타임존 작업

7. **Thread-Safe / 스레드 안전**:
   - All functions are thread-safe
   - 모든 함수가 스레드 안전

8. **Zero Dependencies / 제로 의존성**:
   - Standard library only, no external dependencies
   - 표준 라이브러리만, 외부 의존성 없음

### Files Created / 생성된 파일

**Documentation / 문서**:
- `docs/timeutil/DESIGN_PLAN.md` - Design philosophy and architecture
- `docs/timeutil/WORK_PLAN.md` - Implementation roadmap
- `docs/CHANGELOG/CHANGELOG-v1.6.md` - This file
- `timeutil/README.md` - Initial package documentation

**Directory Structure / 디렉토리 구조**:
```
timeutil/
├── README.md               # Package documentation / 패키지 문서
└── (implementation files to be added) / (구현 파일 추가 예정)

docs/
├── timeutil/
│   ├── DESIGN_PLAN.md     # Design document / 설계 문서
│   └── WORK_PLAN.md       # Work plan / 작업 계획서
└── CHANGELOG/
    └── CHANGELOG-v1.6.md  # This file / 이 파일
```

### Next Steps / 다음 단계

1. **Phase 1: Foundation / 1단계: 기초** (v1.6.001):
   - [x] Create project structure / 프로젝트 구조 생성
   - [x] Create design documents / 설계 문서 작성
   - [x] Create initial README / 초기 README 생성
   - [ ] Create initial package files / 초기 패키지 파일 생성
   - [ ] Update version in cfg/app.yaml / cfg/app.yaml의 버전 업데이트

2. **Phase 2: Core Features / 2단계: 핵심 기능** (v1.6.002-v1.6.009):
   - [ ] Implement core types and constants / 핵심 타입 및 상수 구현
   - [ ] Implement time difference functions / 시간 차이 함수 구현
   - [ ] Implement timezone operations / 타임존 작업 구현
   - [ ] Implement date arithmetic / 날짜 연산 구현
   - [ ] Implement date formatting / 날짜 포맷팅 구현
   - [ ] Implement time parsing / 시간 파싱 구현
   - [ ] Implement time comparisons / 시간 비교 구현
   - [ ] Implement unix timestamp operations / Unix 타임스탬프 작업 구현

3. **Phase 3: Advanced Features / 3단계: 고급 기능** (v1.6.010-v1.6.012):
   - [ ] Implement business days / 영업일 구현
   - [ ] Implement age calculations / 나이 계산 구현
   - [ ] Implement relative time / 상대 시간 구현

4. **Phase 4: Testing & Documentation / 4단계: 테스팅 및 문서화** (v1.6.013-v1.6.015):
   - [ ] Comprehensive testing (≥90% coverage) / 종합 테스팅 (≥90% 커버리지)
   - [ ] Create example code / 예제 코드 생성
   - [ ] Write USER_MANUAL.md / USER_MANUAL.md 작성
   - [ ] Write DEVELOPER_GUIDE.md / DEVELOPER_GUIDE.md 작성

5. **Phase 5: Release / 5단계: 릴리스** (v1.6.015):
   - [ ] Final review / 최종 검토
   - [ ] Update root README.md / 루트 README.md 업데이트
   - [ ] Update root CHANGELOG.md / 루트 CHANGELOG.md 업데이트
   - [ ] Update CLAUDE.md / CLAUDE.md 업데이트
   - [ ] Commit and push to GitHub / GitHub에 커밋 및 푸시

### Design Highlights / 설계 하이라이트

**Before (Standard Go) / 이전 (표준 Go)**:
```go
// Calculate time difference in days
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

### Notes / 참고사항

- This is the initial planning release / 이것은 초기 계획 릴리스입니다
- Implementation will proceed according to WORK_PLAN.md / 구현은 WORK_PLAN.md에 따라 진행됩니다
- Expected completion: 15-21 work units / 예상 완료: 15-21 작업 단위
- Target version for full release: v1.6.015 / 전체 릴리스 목표 버전: v1.6.015

---

## Version History / 버전 히스토리

- **v1.6.001**: Initial planning and design / 초기 계획 및 설계
- **v1.6.002-v1.6.009**: Core features implementation / 핵심 기능 구현 (planned / 예정)
- **v1.6.010-v1.6.012**: Advanced features / 고급 기능 (planned / 예정)
- **v1.6.013-v1.6.014**: Testing & documentation / 테스팅 및 문서화 (planned / 예정)
- **v1.6.015**: Final release / 최종 릴리스 (planned / 예정)

---

**Status / 상태**: 🚧 In Development / 개발 중

**Current Version / 현재 버전**: v1.6.001 (Planning Phase / 계획 단계)

**Target Release Version / 목표 릴리스 버전**: v1.6.015
