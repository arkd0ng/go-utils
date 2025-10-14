# CHANGELOG - v1.5.x

All notable changes for version 1.5.x will be documented in this file.

v1.5.x 버전의 모든 주목할 만한 변경사항이 이 파일에 문서화됩니다.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

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
