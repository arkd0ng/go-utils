# Changelog v1.1.x - Logging Package / 로깅 패키지

All notable changes to the Logging package (v1.1.x) are documented in this file.

Logging 패키지(v1.1.x)의 모든 주요 변경사항이 이 파일에 기록됩니다.

## [v1.1.007] - 2025-10-10

**Commit**: `8a7a2c3`

### Added / 추가
- Added app.yaml version management for logging package / 로깅 패키지를 위한 app.yaml 버전 관리 추가
- Created `cfg/app.yaml` for centralized app name and version management / 중앙집중식 앱 이름 및 버전 관리를 위한 `cfg/app.yaml` 생성
- Added `logging/appconfig.go` to automatically load app.yaml from multiple search paths / 여러 검색 경로에서 app.yaml을 자동으로 로드하는 `logging/appconfig.go` 추가
- Logger now automatically loads app name and version from app.yaml / 로거가 이제 app.yaml에서 앱 이름과 버전을 자동으로 로드
- Search paths include: cfg/, apps/, root directory and parent directories / 검색 경로 포함: cfg/, apps/, 루트 디렉토리 및 상위 디렉토리
- Added comprehensive test for app.yaml integration (`TestAppYamlIntegration`) / app.yaml 통합을 위한 종합 테스트 추가 (`TestAppYamlIntegration`)
- Added yaml.v3 dependency for YAML parsing / YAML 파싱을 위한 yaml.v3 의존성 추가

### Changed / 변경
- Updated logging banner to display app name and version from app.yaml / app.yaml의 앱 이름과 버전을 표시하도록 로깅 배너 업데이트
- Enhanced logger to search for app.yaml in current and parent directories / 현재 및 상위 디렉토리에서 app.yaml을 검색하도록 로거 향상
- Updated .gitignore to exclude log files and directories / 로그 파일 및 디렉토리를 제외하도록 .gitignore 업데이트

### Documentation / 문서
- Updated README with app.yaml documentation and usage examples / app.yaml 문서 및 사용 예제로 README 업데이트
- Added version management section to Features / Features에 버전 관리 섹션 추가
- Added Dependencies section with yaml.v3 / yaml.v3가 포함된 Dependencies 섹션 추가

## [v1.1.006] - 2025-10-10

**Commit**: `58ce152`

### Changed / 변경
- Increased banner width from 40 to 60 characters for better visual appearance / 더 나은 시각적 외관을 위해 배너 너비를 40에서 60 문자로 증가
- Updated all banner methods to use new minimum width / 새로운 최소 너비를 사용하도록 모든 배너 메서드 업데이트

## [v1.1.005] - 2025-10-10

**Commit**: `0f6054e`

### Added / 추가
- Auto-extract app name from log filename for banner / 배너를 위해 로그 파일명에서 앱 이름 자동 추출
- When app name is not specified, automatically extract from log file path / 앱 이름이 지정되지 않은 경우 로그 파일 경로에서 자동 추출
- Example: `./logs/database.log` → banner shows "database" / 예: `./logs/database.log` → 배너에 "database" 표시
- Smart detection works with complex paths: `./logs/api-server.log` → "api-server" / 복잡한 경로에서도 스마트 감지 작동: `./logs/api-server.log` → "api-server"

### Changed / 변경
- Banner now prioritizes: custom name > filename extraction > default "Application" / 배너가 이제 우선순위 적용: 사용자 정의 이름 > 파일명 추출 > 기본값 "Application"

## [v1.1.004] - 2025-10-10

**Commit**: `e8d2532`

### Changed / 변경
- **Breaking Change / 주요 변경**: Changed default output to file-only (no console output) / 기본 출력을 파일 전용으로 변경 (콘솔 출력 없음)
- Console output now disabled by default (`enableStdout: false`) / 콘솔 출력이 이제 기본적으로 비활성화 (`enableStdout: false`)
- Users must explicitly enable console output with `WithStdout(true)` / 사용자는 `WithStdout(true)`로 명시적으로 콘솔 출력을 활성화해야 함
- This change improves production deployment defaults / 이 변경으로 프로덕션 배포 기본값이 개선됨

## [v1.1.003] - 2025-10-10

**Commit**: `9fb7f16`

### Added / 추가
- Added Printf-style logging support for all log levels / 모든 로그 레벨에 대한 Printf 스타일 로깅 지원 추가:
  - `Debugf(format, ...args)`
  - `Infof(format, ...args)`
  - `Warnf(format, ...args)`
  - `Errorf(format, ...args)`
  - `Fatalf(format, ...args)`
- Familiar `fmt.Printf` syntax for quick logging / 빠른 로깅을 위한 친숙한 `fmt.Printf` 문법
- Both structured (key-value) and Printf-style logging now supported / 구조화 (키-값) 및 Printf 스타일 로깅 모두 지원

### Documentation / 문서
- Added comprehensive documentation comparing structured vs Printf-style logging / 구조화 대 Printf 스타일 로깅을 비교하는 종합 문서 추가
- Added usage examples and best practices / 사용 예제 및 모범 사례 추가
- Added recommendation table for when to use each style / 각 스타일을 언제 사용할지에 대한 권장 테이블 추가

## [v1.1.002] - 2025-10-10

**Commit**: `945b0ac`

### Changed / 변경
- Changed log format to put timestamp first / 타임스탬프를 먼저 배치하도록 로그 형식 변경
- Old format / 이전 형식: `[LEVEL] timestamp message`
- New format / 새 형식: `timestamp [LEVEL] message`
- Improved readability and consistency with standard logging practices / 가독성 향상 및 표준 로깅 관행과의 일관성 개선

## [v1.1.001] - 2025-10-10

**Commit**: `e137798`

### Added / 추가
- Added automatic banner feature to logging package / 로깅 패키지에 자동 배너 기능 추가
- Banner automatically prints on logger creation by default / 기본적으로 로거 생성 시 배너가 자동으로 출력됨
- Multiple banner styles / 다양한 배너 스타일:
  - `Banner(appName, version)` - Standard box banner / 표준 박스 배너
  - `SimpleBanner(appName, version)` - Simple separator banner / 간단한 구분선 배너
  - `DoubleBanner(appName, version, description)` - Banner with description / 설명이 있는 배너
  - `CustomBanner(lines)` - Custom ASCII art / 사용자 정의 ASCII 아트
  - `SeparatorLine(char, width)` - Separator line / 구분선
- Configuration options / 설정 옵션:
  - `WithAutoBanner(bool)` - Enable/disable auto banner / 자동 배너 활성화/비활성화
  - `WithAppName(string)` - Set application name / 애플리케이션 이름 설정
  - `WithAppVersion(string)` - Set application version / 애플리케이션 버전 설정
  - `WithBanner(name, version)` - Convenience function / 편의 함수

### Documentation / 문서
- Added banner documentation to README / README에 배너 문서 추가
- Added banner usage examples / 배너 사용 예제 추가

## [v1.1.000] - 2025-10-10

**Commits**: `13740ab`, `cb6a7bd`

### Added / 추가
- Initial logging package with file rotation and structured logging / 파일 로테이션 및 구조화된 로깅을 갖춘 초기 로깅 패키지
- Integration with lumberjack v2.2.1 for automatic log file rotation / 자동 로그 파일 로테이션을 위한 lumberjack v2.2.1 통합
- Multiple log levels: DEBUG, INFO, WARN, ERROR, FATAL / 다중 로그 레벨: DEBUG, INFO, WARN, ERROR, FATAL
- Structured logging with key-value pairs / 키-값 쌍을 사용한 구조화된 로깅
- Thread-safe logging with mutex locks / 뮤텍스 잠금을 사용한 스레드 안전 로깅
- ANSI color-coded console output / ANSI 색상 코딩된 콘솔 출력
- Configurable options via Options pattern / 옵션 패턴을 통한 설정 가능한 옵션:
  - File rotation settings (MaxSize, MaxBackups, MaxAge, Compress) / 파일 로테이션 설정
  - Log level filtering / 로그 레벨 필터링
  - Custom prefix / 사용자 정의 프리픽스
  - Color output toggle / 색상 출력 토글
  - Stdout/File output control / Stdout/파일 출력 제어
  - Custom time format / 사용자 정의 시간 형식
- Support for multiple independent loggers / 여러 독립 로거 지원
- Comprehensive test suite with 15+ test cases / 15개 이상의 테스트 케이스를 포함한 종합 테스트 스위트
- Benchmark tests for performance validation / 성능 검증을 위한 벤치마크 테스트

### Documentation / 문서
- Created comprehensive README.md with bilingual content / 이중 언어 내용으로 종합 README.md 생성
- Added usage examples and best practices / 사용 예제 및 모범 사례 추가
- Documented all configuration options / 모든 설정 옵션 문서화
- Added examples directory with working code / 작동하는 코드가 포함된 examples 디렉토리 추가

---

**Version Summary / 버전 요약**: v1.1.x focused on developing the Logging package with automatic file rotation, structured logging, Printf-style support, auto banner, and app.yaml version management.

v1.1.x는 자동 파일 로테이션, 구조화된 로깅, Printf 스타일 지원, 자동 배너, app.yaml 버전 관리를 갖춘 Logging 패키지 개발에 초점을 맞추었습니다.
