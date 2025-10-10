# Changelog v1.0.x - Random Package / 랜덤 패키지

All notable changes to the Random package (v1.0.x) are documented in this file.

Random 패키지(v1.0.x)의 모든 주요 변경사항이 이 파일에 기록됩니다.

## [v1.0.008] - 2025-10-10

**Commit**: `8215696`

### Added / 추가
- Added variadic parameters support for all random string generation methods / 모든 랜덤 문자열 생성 메서드에 가변 인자 지원 추가
- Methods now accept flexible length parameters / 메서드가 이제 유연한 길이 파라미터를 받습니다:
  - Single parameter: Fixed length (e.g., `Alnum(32)` generates exactly 32 characters) / 단일 파라미터: 고정 길이 (예: `Alnum(32)`는 정확히 32자 생성)
  - Two parameters: Range (e.g., `Alnum(32, 128)` generates 32-128 characters) / 두 파라미터: 범위 (예: `Alnum(32, 128)`는 32-128자 생성)
- Comprehensive error handling for all methods / 모든 메서드에 대한 포괄적인 에러 처리
- All methods now return `(string, error)` instead of just `string` / 모든 메서드가 이제 `string` 대신 `(string, error)`를 반환

### Changed / 변경
- **Breaking Change / 주요 변경**: Method signatures changed to return `(string, error)` / 메서드 시그니처가 `(string, error)` 반환으로 변경
- Improved validation for input parameters / 입력 파라미터에 대한 개선된 검증

## [v1.0.007] - 2025-10-10

**Commit**: `e729ef4`

### Changed / 변경
- Updated random string package to handle new error return values / 새로운 에러 반환값을 처리하도록 랜덤 문자열 패키지 업데이트
- Updated all examples and documentation to use new error handling pattern / 새로운 에러 처리 패턴을 사용하도록 모든 예제 및 문서 업데이트

## [v1.0.006] - 2025-10-10

**Commit**: `1351c48`

### Added / 추가
- Added 9 new random string generation methods / 9개의 새로운 랜덤 문자열 생성 메서드 추가:
  - `AlphaUpper()` - Uppercase letters only (A-Z) / 대문자만 (A-Z)
  - `AlphaLower()` - Lowercase letters only (a-z) / 소문자만 (a-z)
  - `AlnumUpper()` - Uppercase alphanumeric (A-Z, 0-9) / 대문자 알파뉴메릭 (A-Z, 0-9)
  - `AlnumLower()` - Lowercase alphanumeric (a-z, 0-9) / 소문자 알파뉴메릭 (a-z, 0-9)
  - `Hex()` - Uppercase hexadecimal (0-9, A-F) / 대문자 16진수 (0-9, A-F)
  - `HexLower()` - Lowercase hexadecimal (0-9, a-f) / 소문자 16진수 (0-9, a-f)
  - `Base64()` - Base64 characters / Base64 문자
  - `Base64URL()` - URL-safe Base64 characters / URL 안전 Base64 문자
  - `Complex()` - Complex characters including special symbols / 특수 기호를 포함한 복잡한 문자

### Changed / 변경
- Expanded total methods from 5 to 14 / 전체 메서드를 5개에서 14개로 확장

## [v1.0.005] - 2025-10-10

**Commit**: `a44c6c7`

### Changed / 변경
- Renamed methods for consistency / 일관성을 위해 메서드 이름 변경:
  - `GenLetters` → `Letters`
  - `GenAlnum` → `Alnum`
  - `GenDigits` → `Digits`
  - `GenCustom` → `Custom`
  - `GenStandard` → `Standard`

### Added / 추가
- Added collision probability testing / 충돌 확률 테스트 추가
- Implemented theoretical vs actual collision rate calculation / 이론적 대 실제 충돌률 계산 구현
- Added comprehensive test suite for randomness validation / 랜덤성 검증을 위한 종합 테스트 스위트 추가

## [v1.0.004] - 2025-10-10

**Commit**: `56f9454`

### Changed / 변경
- Reorganized documentation structure / 문서 구조 재구성
- Improved README formatting / README 형식 개선
- Enhanced code examples with better explanations / 더 나은 설명과 함께 코드 예제 향상

## [v1.0.003] - 2025-10-10

**Commit**: `1a1bb38`

### Added / 추가
- Added bilingual comments (English/Korean) throughout codebase / 코드베이스 전체에 이중 언어 주석 (영문/한글) 추가
- All function documentation now includes both English and Korean descriptions / 모든 함수 문서에 영문 및 한글 설명 포함
- README updated with bilingual content / 이중 언어 내용으로 README 업데이트

### Changed / 변경
- Improved documentation readability / 문서 가독성 개선
- Enhanced code comments for better understanding / 더 나은 이해를 위해 코드 주석 향상

## [v1.0.002] - 2025-10-10

**Commit**: `f1436e3`

### Changed / 변경
- **Breaking Change / 주요 변경**: Migrated to subpackage structure / 서브패키지 구조로 마이그레이션
- Moved from root-level `GenRandomString` to `random.GenString` / 루트 레벨 `GenRandomString`에서 `random.GenString`으로 이동
- Package now accessed via `import "github.com/arkd0ng/go-utils/random"` / 이제 패키지는 `import "github.com/arkd0ng/go-utils/random"`을 통해 접근
- Improved modularity and maintainability / 모듈성 및 유지보수성 개선

### Added / 추가
- Created dedicated `random/` subpackage / 전용 `random/` 서브패키지 생성
- Added package-level documentation / 패키지 레벨 문서 추가
- Created `examples/random_string/` directory / `examples/random_string/` 디렉토리 생성

## [v1.0.001] - 2025-10-10

**Commit**: `7ceb110`

### Added / 추가
- Initial commit with `GenRandomString` utility / `GenRandomString` 유틸리티를 포함한 초기 커밋
- Basic random string generation functionality / 기본 랜덤 문자열 생성 기능
- Cryptographically secure random generation using `crypto/rand` / `crypto/rand`를 사용한 암호학적으로 안전한 랜덤 생성
- Initial test suite / 초기 테스트 스위트
- MIT License / MIT 라이선스

---

**Version Summary / 버전 요약**: v1.0.x focused on developing the Random package with cryptographically secure random string generation, supporting 14 different methods and flexible length parameters.

v1.0.x는 암호학적으로 안전한 랜덤 문자열 생성을 갖춘 Random 패키지 개발에 초점을 맞추었으며, 14가지 다양한 메서드와 유연한 길이 파라미터를 지원합니다.
