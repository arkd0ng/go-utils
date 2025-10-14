# CHANGELOG - v1.5.x

All notable changes for version 1.5.x will be documented in this file.

v1.5.x 버전의 모든 주목할 만한 변경사항이 이 파일에 문서화됩니다.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

---

## [v1.5.008] - 2025-10-14 16:04

### Added / 추가

- **DOCS**: Created comprehensive DEVELOPER_GUIDE.md for stringutil package
- **문서**: stringutil 패키지에 대한 포괄적인 DEVELOPER_GUIDE.md 생성
- **CONTENT**: ~1,750 lines of detailed developer documentation
- **내용**: 약 1,750줄의 상세한 개발자 문서

### Documentation Structure / 문서 구조

- **Table of Contents**: 10 major sections
- **목차**: 10개 주요 섹션
- **Architecture Overview**: Design philosophy, high-level architecture, design decisions
- **아키텍처 개요**: 설계 철학, 상위 수준 아키텍처, 설계 결정
- **Package Structure**: File organization, dependencies, responsibility breakdown
- **패키지 구조**: 파일 구성, 의존성, 책임 분류
- **Core Components**: 5 detailed component implementations
- **핵심 컴포넌트**: 5개 상세 컴포넌트 구현
  - Smart Word Splitting algorithm
  - Unicode-Safe Truncation
  - Practical Email Validation
  - Map and Filter (Functional Programming)
  - Multi-Pattern Replace
- **Internal Implementation**: Flow diagrams for key functions
- **내부 구현**: 주요 함수에 대한 흐름 다이어그램
- **Design Patterns**: 5 patterns with examples
- **디자인 패턴**: 예제가 있는 5개 패턴
  - Helper Function Pattern
  - Wrapper Pattern
  - Higher-Order Function Pattern
  - Predicate Pattern
  - Builder Pattern
- **Adding New Features**: 7-step guide with templates
- **새 기능 추가**: 템플릿이 있는 7단계 가이드
- **Testing Guide**: Test structure, running tests, writing good tests
- **테스트 가이드**: 테스트 구조, 테스트 실행, 좋은 테스트 작성
- **Performance**: Time/space complexity, optimization techniques
- **성능**: 시간/공간 복잡도, 최적화 기법
- **Contributing Guidelines**: Complete workflow and review checklist
- **기여 가이드라인**: 완전한 워크플로우 및 리뷰 체크리스트
- **Code Style**: Naming conventions, comment style, error handling
- **코드 스타일**: 명명 규칙, 주석 스타일, 에러 처리
- **Appendix**: Complete function reference
- **부록**: 완전한 함수 참조

### Key Topics Covered / 다룬 주요 주제

- Design philosophy: "20 lines → 1 line"
- 설계 철학: "20줄 → 1줄"
- Unicode safety implementation ([]rune vs byte)
- 유니코드 안전 구현 ([]rune vs byte)
- Zero dependencies approach
- 제로 의존성 접근
- Practical validation vs RFC compliance
- 실용적 검증 vs RFC 준수
- Smart case conversion algorithm
- 스마트 케이스 변환 알고리즘
- Performance optimization techniques
- 성능 최적화 기법
- Functional programming patterns (Map/Filter)
- 함수형 프로그래밍 패턴 (Map/Filter)

### Files Created / 생성된 파일

- `docs/stringutil/DEVELOPER_GUIDE.md` (~1,750 lines)
- Architecture diagrams in ASCII art
- 아스키 아트로 된 아키텍처 다이어그램
- Flow diagrams for key algorithms
- 주요 알고리즘에 대한 흐름 다이어그램
- Complete code examples with explanations
- 설명이 있는 완전한 코드 예제
- All sections in bilingual format (English/Korean)
- 모든 섹션은 이중 언어 형식 (영문/한글)

### Notes / 참고사항

- Comprehensive developer documentation complete
- 포괄적인 개발자 문서 완성
- Package documentation now complete (USER_MANUAL + DEVELOPER_GUIDE)
- 패키지 문서화 완료 (USER_MANUAL + DEVELOPER_GUIDE)
- Total documentation: ~3,600 lines
- 전체 문서: 약 3,600줄
- Current version: v1.5.008
- 현재 버전: v1.5.008

---

## [v1.5.007] - 2025-10-14 16:01

### Fixed / 수정

- **FIX**: Fixed MySQL example Docker cleanup to handle container not found errors
- **수정**: MySQL 예제 Docker 정리에서 컨테이너 찾을 수 없음 에러 처리 수정
- Added conditional check before stopping/removing Docker containers
- Docker 컨테이너 중지/제거 전 조건 확인 추가
- Improved error messages for Docker operations
- Docker 작업에 대한 에러 메시지 개선

### Files Modified / 수정된 파일

- `examples/mysql/main.go` - Enhanced Docker cleanup logic

### Notes / 참고사항

- Current version: v1.5.007
- 현재 버전: v1.5.007

---

## [v1.5.005] - 2025-10-14 15:57

### Added / 추가

- **DOCS**: Created comprehensive USER_MANUAL.md for stringutil package
- **문서**: stringutil 패키지에 대한 포괄적인 USER_MANUAL.md 생성
- **CONTENT**: ~1,850 lines of detailed user documentation
- **내용**: 약 1,850줄의 상세한 사용자 문서

### Documentation Structure / 문서 구조

- **Table of Contents**: 9 major sections with subsections
- **목차**: 하위 섹션이 있는 9개 주요 섹션
- **Introduction**: Package overview, key features, use cases
- **소개**: 패키지 개요, 주요 기능, 사용 사례
- **Installation**: Prerequisites, package installation, verification
- **설치**: 전제 조건, 패키지 설치, 확인
- **Quick Start**: 5 quick examples covering all categories
- **빠른 시작**: 모든 카테고리를 다루는 5개의 빠른 예제
- **Configuration Reference**: Complete API documentation for all 37 functions
- **설정 참조**: 모든 37개 함수에 대한 완전한 API 문서
  - Case Conversion (5 functions)
  - String Manipulation (9 functions)
  - Validation (8 functions)
  - Search & Replace (6 functions)
  - Utilities (9 functions)
- **Usage Patterns**: 10 common patterns with complete code examples
- **사용 패턴**: 완전한 코드 예제가 있는 10개의 일반적인 패턴
- **Common Use Cases**: 10 real-world scenarios with full implementations
- **일반적인 사용 사례**: 전체 구현이 있는 10개의 실제 시나리오
- **Best Practices**: 15 practical recommendations
- **모범 사례**: 15개의 실용적인 권장 사항
- **Troubleshooting**: 10 common problems and solutions
- **문제 해결**: 10개의 일반적인 문제 및 해결책
- **FAQ**: 15 frequently asked questions
- **FAQ**: 15개의 자주 묻는 질문
- **Appendix**: Function reference table
- **부록**: 함수 참조 표

### Key Features Documented / 문서화된 주요 기능

- Unicode safety explanation and examples
- 유니코드 안전 설명 및 예제
- Practical validation approach (vs RFC compliance)
- 실용적인 검증 접근 방식 (RFC 준수 대비)
- Smart case conversion handling
- 스마트 케이스 변환 처리
- Map/Filter functional programming patterns
- Map/Filter 함수형 프로그래밍 패턴
- Multi-pattern search and replace
- 다중 패턴 검색 및 치환

### Files Created / 생성된 파일

- `docs/stringutil/USER_MANUAL.md` (~1,850 lines)
- All sections in bilingual format (English/Korean)
- 모든 섹션은 이중 언어 형식 (영문/한글)

### Notes / 참고사항

- Comprehensive user-facing documentation complete
- 포괄적인 사용자 대상 문서 완성
- Next: DEVELOPER_GUIDE.md
- 다음: DEVELOPER_GUIDE.md
- Current version: v1.5.005
- 현재 버전: v1.5.005

---

## [v1.5.006] - 2025-10-14 15:51

### Fixed / 수정

- **FIX**: Fixed MySQL example initialization to ensure sample data exists
- **수정**: 샘플 데이터가 존재하도록 MySQL 예제 초기화를 수정
- **ENHANCEMENT**: Added `initializeDatabaseIfNeeded()` function to automatically:
- **개선**: `initializeDatabaseIfNeeded()` 함수를 추가하여 자동으로:
  - Check if `users` table exists and create it if needed
  - users 테이블이 존재하는지 확인하고 필요시 생성
  - Verify sample data by checking for known user (john@example.com)
  - 알려진 사용자(john@example.com)를 확인하여 샘플 데이터 검증
  - Truncate and reinitialize table if sample data is missing or incomplete
  - 샘플 데이터가 없거나 불완전한 경우 테이블을 비우고 재초기화
  - Insert 11 sample users for consistent example execution
  - 일관된 예제 실행을 위해 11명의 샘플 사용자 삽입

### Problem Solved / 해결된 문제

- MySQL Docker volumes persist between container restarts
- MySQL Docker 볼륨이 컨테이너 재시작 간에 유지됨
- Init scripts in `/docker-entrypoint-initdb.d` only run on first initialization
- `/docker-entrypoint-initdb.d`의 초기화 스크립트는 첫 초기화 시에만 실행됨
- Previous test runs left incomplete or incorrect data in the database
- 이전 테스트 실행이 데이터베이스에 불완전하거나 잘못된 데이터를 남김
- Examples failed with "no rows in result set" error
- 예제가 "no rows in result set" 오류로 실패

### Files Changed / 변경된 파일

- `examples/mysql/main.go`:
  - Added `initializeDatabaseIfNeeded()` function (~90 lines)
  - Call initialization before running examples
  - Smart detection of sample data presence
  - Automatic table creation and data population

### Test Results / 테스트 결과

```
✅ All 35 MySQL examples completed successfully
✅ 모든 35개 MySQL 예제가 성공적으로 완료되었습니다
- Example 1-8: Basic operations (SelectAll, SelectOne, Insert, Update, Count, Exists, Transaction, Delete)
- Example 9-17: Query operations (Raw SQL, Query Builder, SelectWhere, SelectColumn, SelectColumns)
- Example 18-24: Advanced operations (Batch, Upsert, Pagination, Soft Delete)
- Example 25-30: Monitoring (QueryStats, SlowQueryLog, PoolMetrics)
- Example 31-35: Schema operations (GetTables, InspectTable, Migration, ExportCSV)
```

### Notes / 참고사항

- MySQL example now works reliably regardless of previous test runs
- MySQL 예제가 이전 테스트 실행에 관계없이 안정적으로 작동
- Initialization happens automatically and is idempotent
- 초기화가 자동으로 발생하며 멱등성을 가짐
- Sample data includes 11 users across 7 Korean cities
- 샘플 데이터는 7개 한국 도시에 걸쳐 11명의 사용자 포함
- Current version: v1.5.006
- 현재 버전: v1.5.006

---

## [v1.5.004] - 2025-10-14 15:48

### Added / 추가

- **TEST**: Created comprehensive test suite for stringutil package
- **테스트**: stringutil 패키지에 대한 포괄적인 테스트 스위트 생성
- **DOCS**: Created README.md with complete API documentation
- **문서**: 완전한 API 문서가 포함된 README.md 생성
- **EXAMPLES**: Created working examples with logging integration
- **예제**: 로깅 통합이 포함된 작동 예제 생성

### Test Files / 테스트 파일

- `stringutil/case_test.go` - 3 test functions (ToSnakeCase, ToCamelCase, ToPascalCase)
- `stringutil/manipulation_test.go` - 3 test functions (Truncate, Reverse, Clean)
- `stringutil/validation_test.go` - 3 test functions (IsEmail, IsURL, IsAlphanumeric)
- All 9 tests passing with Unicode validation
- 모든 9개 테스트 통과 (유니코드 검증 포함)

### Documentation / 문서

- `stringutil/README.md` - Comprehensive package documentation with:
  - Installation instructions
  - Quick start examples
  - Complete API reference for all 37 functions
  - 5 categories with detailed tables
  - Usage examples for each function
- All documentation in bilingual format (English/Korean)
- 모든 문서는 이중 언어 형식 (영문/한글)

### Examples / 예제

- `examples/stringutil/main.go` - Complete working example demonstrating:
  - All 5 categories of functions
  - Logging integration with file and stdout output
  - Banner display with version info
  - Practical use cases for each category
- Example tested and runs successfully
- 예제 테스트 완료 및 성공적으로 실행됨

### Test Results / 테스트 결과

```
PASS: TestToSnakeCase
PASS: TestToCamelCase
PASS: TestToPascalCase
PASS: TestTruncate (with Unicode "안녕하세요")
PASS: TestReverse (with Unicode "안녕")
PASS: TestClean
PASS: TestIsEmail
PASS: TestIsURL
PASS: TestIsAlphanumeric
ok  github.com/arkd0ng/go-utils/stringutil  0.697s
```

### Notes / 참고사항

- Phase 4 testing complete: 9 tests passing
- Phase 4 테스팅 완료: 9개 테스트 통과
- Basic documentation complete (README)
- 기본 문서화 완료 (README)
- Package ready for basic use
- 패키지 기본 사용 준비 완료
- Next: Comprehensive documentation (USER_MANUAL, DEVELOPER_GUIDE)
- 다음: 포괄적인 문서화 (USER_MANUAL, DEVELOPER_GUIDE)
- Current version: v1.5.004
- 현재 버전: v1.5.004

---

## [v1.5.003] - 2025-10-14 15:46

### Added / 추가

- **FEAT**: Implemented all core stringutil functions (Phase 2 complete)
- **기능**: 모든 핵심 stringutil 함수 구현 완료 (Phase 2 완료)
- **Case Conversion** (5 functions): ToSnakeCase, ToCamelCase, ToKebabCase, ToPascalCase, ToScreamingSnakeCase
- **케이스 변환** (5개 함수): ToSnakeCase, ToCamelCase, ToKebabCase, ToPascalCase, ToScreamingSnakeCase
- **String Manipulation** (9 functions): Truncate, TruncateWithSuffix, Reverse, Capitalize, CapitalizeFirst, RemoveDuplicates, RemoveSpaces, RemoveSpecialChars, Clean
- **문자열 조작** (9개 함수): Truncate, TruncateWithSuffix, Reverse, Capitalize, CapitalizeFirst, RemoveDuplicates, RemoveSpaces, RemoveSpecialChars, Clean
- **Validation** (8 functions): IsEmail, IsURL, IsAlphanumeric, IsNumeric, IsAlpha, IsBlank, IsLower, IsUpper
- **유효성 검사** (8개 함수): IsEmail, IsURL, IsAlphanumeric, IsNumeric, IsAlpha, IsBlank, IsLower, IsUpper
- **Search & Replace** (6 functions): ContainsAny, ContainsAll, StartsWithAny, EndsWithAny, ReplaceAll, ReplaceIgnoreCase
- **검색 및 치환** (6개 함수): ContainsAny, ContainsAll, StartsWithAny, EndsWithAny, ReplaceAll, ReplaceIgnoreCase
- **Utilities** (9 functions): CountWords, CountOccurrences, Join, Map, Filter, PadLeft, PadRight, Lines, Words
- **유틸리티** (9개 함수): CountWords, CountOccurrences, Join, Map, Filter, PadLeft, PadRight, Lines, Words

### Implementation Details / 구현 세부사항

- All functions are Unicode-safe (using rune, not byte)
- 모든 함수는 유니코드 안전 (byte 대신 rune 사용)
- Zero external dependencies (standard library only)
- 외부 의존성 제로 (표준 라이브러리만)
- Bilingual documentation for all functions
- 모든 함수에 이중 언어 문서
- Smart case conversion handles multiple input formats
- 스마트 케이스 변환은 여러 입력 형식 처리
- Practical email/URL validation (99% use cases)
- 실용적인 이메일/URL 검증 (99% 사용 사례)

### Files Created / 생성된 파일

- `stringutil/stringutil.go` - Package documentation
- `stringutil/case.go` - Case conversion (163 lines)
- `stringutil/manipulation.go` - String manipulation (139 lines)
- `stringutil/validation.go` - Validation (170 lines)
- `stringutil/search.go` - Search and replace (114 lines)
- `stringutil/utils.go` - Utility functions (128 lines)

### Notes / 참고사항

- Phase 2 complete: 37 functions implemented
- Phase 2 완료: 37개 함수 구현됨
- Next: Phase 4 - Testing (skip Phase 3 Builder for now)
- 다음: Phase 4 - 테스팅 (Phase 3 Builder는 일단 건너뜀)
- Current version: v1.5.003
- 현재 버전: v1.5.003

---

## [v1.5.002] - 2025-10-14 15:41

### Added / 추가

- **DOCS**: Created comprehensive WORK_PLAN.md for stringutil package
- **문서**: stringutil 패키지에 대한 포괄적인 WORK_PLAN.md 생성
- Defined 5 phases with 14 tasks and estimated 16.5 work units
- 14개 작업과 16.5 작업 단위가 예상되는 5개 단계 정의
- Detailed task breakdown for each phase:
- 각 단계에 대한 상세 작업 분류:
  - Phase 1: Foundation (1 task, 0.5 units)
  - Phase 2: Core Functions (5 tasks, 7.5 units)
  - Phase 3: Advanced Features (2 tasks, 1.5 units)
  - Phase 4: Testing & Documentation (4 tasks, 6.0 units)
  - Phase 5: Release (2 tasks, 1.0 units)
- Each task has clear acceptance criteria and subtasks
- 각 작업에 명확한 수용 기준 및 하위 작업 있음
- Task dependencies documented with visual flow
- 시각적 흐름과 함께 작업 의존성 문서화
- Quality checklist for code, testing, and documentation
- 코드, 테스팅, 문서화를 위한 품질 체크리스트

### Notes / 참고사항

- Next: Begin Phase 1 - Project Structure Setup
- 다음: 1단계 시작 - 프로젝트 구조 설정
- Current version: v1.5.002
- 현재 버전: v1.5.002

---

## [v1.5.001] - 2025-10-14 15:38

### Added / 추가

- **NEW Package**: `stringutil` package - Extreme simplicity string utilities
- **새로운 패키지**: `stringutil` 패키지 - 극도로 간단한 문자열 유틸리티
- Created DESIGN_PLAN.md for stringutil package with comprehensive architecture design
- stringutil 패키지에 대한 포괄적인 아키텍처 설계가 포함된 DESIGN_PLAN.md 생성

### Documentation / 문서

- Documented stringutil package design philosophy: "20 lines → 1 line"
- stringutil 패키지 설계 철학 문서화: "20줄 → 1줄"
- Planned 5 categories of functions:
- 5개 카테고리의 함수 계획:
  - Case Conversion (ToSnakeCase, ToCamelCase, ToKebabCase, etc.)
  - String Manipulation (Truncate, Reverse, Capitalize, Clean, etc.)
  - Validation (IsEmail, IsURL, IsAlphanumeric, IsNumeric, etc.)
  - Search & Replace (ContainsAny, ContainsAll, ReplaceAll, etc.)
  - Utilities (CountWords, PadLeft, Lines, Words, etc.)
- Unicode-safe operations with rune support
- rune 지원으로 유니코드 안전 작업
- Zero external dependencies (standard library only)
- 외부 의존성 제로 (표준 라이브러리만)

### Notes / 참고사항

- Started v1.5.x series for stringutil package
- stringutil 패키지를 위한 v1.5.x 시리즈 시작
- Next: WORK_PLAN.md creation
- 다음: WORK_PLAN.md 생성
- Current version: v1.5.001
- 현재 버전: v1.5.001

---

## Version Overview / 버전 개요

**v1.5.x Series Goals / v1.5.x 시리즈 목표**:
- Implement `stringutil` package with extreme simplicity (20 lines → 1 line)
- 극도의 간결함으로 `stringutil` 패키지 구현 (20줄 → 1줄)
- Case conversions: snake_case, camelCase, kebab-case, PascalCase
- 케이스 변환: snake_case, camelCase, kebab-case, PascalCase
- String manipulation: truncate, reverse, capitalize, clean
- 문자열 조작: 자르기, 뒤집기, 대문자화, 정리
- Validation: email, URL, alphanumeric, numeric
- 검증: 이메일, URL, 영숫자, 숫자
- Search & replace: contains, starts/ends with, replace
- 검색 및 치환: 포함, 시작/끝, 치환
- Utilities: word count, padding, splitting
- 유틸리티: 단어 개수, 패딩, 분할
- Unicode-safe with 100% test coverage
- 유니코드 안전 및 100% 테스트 커버리지
- Comprehensive documentation (README, USER_MANUAL, DEVELOPER_GUIDE)
- 포괄적인 문서화 (README, USER_MANUAL, DEVELOPER_GUIDE)
