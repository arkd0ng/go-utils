# CHANGELOG - v1.4.x

All notable changes for version 1.4.x will be documented in this file.

v1.4.x 버전의 모든 주목할 만한 변경사항이 이 파일에 문서화됩니다.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

---

## [v1.4.021] - 2025-10-14

### Changed / 변경

- **CLEANUP**: Removed personal development files from repository
- **정리**: 저장소에서 개인 개발 파일 제거

### Notes / 참고사항

- Current version: v1.4.021
- 현재 버전: v1.4.021

---

## [v1.4.019] - 2025-10-14

### Changed / 변경

- **REFACTOR**: Updated random string examples to use logging package
- **리팩터링**: random 문자열 예제를 logging 패키지 사용하도록 업데이트
- Random examples now use structured logging with automatic log file creation
- Random 예제는 이제 자동 로그 파일 생성과 함께 구조화된 로깅 사용
- Log files saved to `./results/logs/random_example_YYYYMMDD_HHMMSS.log`
- 로그 파일은 `./results/logs/random_example_YYYYMMDD_HHMMSS.log`에 저장됨
- Added banner display at startup showing app version
- 시작 시 앱 버전을 표시하는 배너 추가
- All output goes to both console (with colors) and log file (without colors)
- 모든 출력은 콘솔(색상 포함)과 로그 파일(색상 제외) 모두에 기록됨

### Benefits / 장점

- Consistent logging pattern across all examples (random, logging, mysql, redis)
- 모든 예제에서 일관된 로깅 패턴 (random, logging, mysql, redis)
- Better traceability with timestamped log files
- 타임스탬프가 찍힌 로그 파일로 더 나은 추적성
- Professional output format with structured logging
- 구조화된 로깅으로 전문적인 출력 형식

### Updated Files / 업데이트된 파일

- `examples/random_string/main.go` - Complete rewrite with logging integration

### Notes / 참고사항

- Random examples now follow same pattern as MySQL and Redis examples
- Random 예제는 이제 MySQL 및 Redis 예제와 동일한 패턴 사용
- Results directory automatically created if not exists
- Results 디렉토리가 없으면 자동으로 생성됨
- Current version: v1.4.019
- 현재 버전: v1.4.019

---

## [v1.4.018] - 2025-10-14

### Fixed / 수정

- **FIX**: Updated logging package test to check for correct version number
- **수정**: 올바른 버전 번호를 확인하도록 logging 패키지 테스트 업데이트
- Fixed `TestBanner()` in `logging/logger_test.go` to expect v1.4.017 instead of v1.4.014
- `logging/logger_test.go`의 `TestBanner()`를 v1.4.014 대신 v1.4.017을 예상하도록 수정
- Test was failing because version expectation was outdated
- 버전 기대값이 오래되어 테스트가 실패했었음

### Verified / 검증

- All 4 packages tested and passing ✅
- 4개 패키지 모두 테스트되고 통과 ✅
  - Random package: 17 tests ✅
  - Logging package: 14 tests ✅
  - MySQL package: All tests ✅
  - Redis package: 18 tests ✅
- All tests verified after directory reorganization
- 디렉토리 재구성 후 모든 테스트 검증 완료
- No breaking changes to public API
- 공개 API에 대한 중단 변경 없음

### Updated Files / 업데이트된 파일

- `logging/logger_test.go` - Updated version expectation in TestBanner

### Notes / 참고사항

- Final verification of all packages after v1.4.016 directory reorganization
- v1.4.016 디렉토리 재구성 후 모든 패키지의 최종 검증
- Repository is now in clean state with all tests passing
- 저장소는 이제 모든 테스트가 통과하는 깨끗한 상태
- Current version: v1.4.018
- 현재 버전: v1.4.018

---

## [v1.4.017] - 2025-10-14

### Fixed / 수정

- **FIX**: Updated Redis test helper to use new `.docker/` directory path
- **수정**: 새 `.docker/` 디렉토리 경로를 사용하도록 Redis 테스트 헬퍼 업데이트
- Fixed `startDockerRedis()` to run docker compose from `.docker/` directory
- `.docker/` 디렉토리에서 docker compose를 실행하도록 `startDockerRedis()` 수정
- Fixed `stopDockerRedis()` to run docker compose from `.docker/` directory
- `.docker/` 디렉토리에서 docker compose를 실행하도록 `stopDockerRedis()` 수정
- Test helper now correctly locates docker-compose.yml in new location
- 테스트 헬퍼가 이제 새 위치에서 docker-compose.yml을 올바르게 찾음

### Verified / 검증

- All Redis package tests pass (18 tests) ✅
- 모든 Redis 패키지 테스트 통과 (18개 테스트) ✅
- All MySQL package tests pass ✅
- 모든 MySQL 패키지 테스트 통과 ✅
- Redis examples run successfully with new paths ✅
- Redis 예제가 새 경로에서 성공적으로 실행 ✅
- Docker auto-start/stop works correctly from test suite ✅
- 테스트 스위트에서 Docker 자동 시작/중지 정상 작동 ✅

### Updated Files / 업데이트된 파일

- `database/redis/testhelper_test.go` - Updated docker compose directory paths

### Notes / 참고사항

- All tests and examples verified after directory reorganization
- 디렉토리 재구성 후 모든 테스트 및 예제 검증 완료
- No breaking changes to public API
- 공개 API에 대한 중단 변경 없음
- Current version: v1.4.017
- 현재 버전: v1.4.017

---

## [v1.4.016] - 2025-10-14

### Changed / 변경

- **REFACTOR**: Reorganized Docker files into `.docker/` directory
- **리팩터링**: Docker 파일을 `.docker/` 디렉토리로 재구성
- Moved `mysql/` → `.docker/mysql/` (MySQL configuration and init scripts)
- `mysql/`를 `.docker/mysql/`로 이동 (MySQL 설정 및 초기화 스크립트)
- Moved `redis/` → `.docker/redis/` (Redis configuration)
- `redis/`를 `.docker/redis/`로 이동 (Redis 설정)
- Moved `scripts/` → `.docker/scripts/` (Docker management scripts)
- `scripts/`를 `.docker/scripts/`로 이동 (Docker 관리 스크립트)
- Moved `docker-compose.yml` → `.docker/docker-compose.yml`
- `docker-compose.yml`을 `.docker/docker-compose.yml`로 이동
- Updated all script references throughout codebase
- 코드베이스 전체의 모든 스크립트 참조 업데이트
- All scripts now referenced as `./.docker/scripts/docker-*`
- 모든 스크립트는 이제 `./.docker/scripts/docker-*`로 참조됨

### Added / 추가

- **GITIGNORE**: Added `.claude/` and `CLAUDE.md` to .gitignore
- **GITIGNORE**: `.claude/` 및 `CLAUDE.md`를 .gitignore에 추가
- Personal Claude Code configuration files no longer tracked by git
- 개인 Claude Code 설정 파일은 더 이상 git에서 추적되지 않음
- Local files preserved, only removed from remote repository
- 로컬 파일은 유지되고 원격 저장소에서만 제거됨

### Updated Files / 업데이트된 파일

- `.gitignore` - Added Claude Code exclusions
- `CLAUDE.md` - Updated all Docker script paths
- `database/redis/README.md` - Updated script paths
- `docs/CHANGELOG/CHANGELOG-v1.3.md` - Updated script paths
- `docs/database/redis/DEVELOPER_GUIDE.md` - Updated script paths
- `docs/database/redis/USER_MANUAL.md` - Updated script paths
- `docs/database/redis/WORK_PLAN.md` - Updated script paths
- `examples/redis/README.md` - Updated script paths
- `examples/redis/main.go` - Updated script paths

### Benefits / 장점

- Cleaner project root directory
- 더 깨끗한 프로젝트 루트 디렉토리
- All Docker-related files organized in one place
- 모든 Docker 관련 파일이 한 곳에 정리됨
- Easier to find and manage Docker configurations
- Docker 설정을 찾고 관리하기 더 쉬움
- Personal settings not shared in repository
- 개인 설정이 저장소에 공유되지 않음

### Notes / 참고사항

- Scripts maintain same functionality with new paths
- 스크립트는 새 경로에서 동일한 기능 유지
- Docker containers still named `go-utils-mysql` and `go-utils-redis`
- Docker 컨테이너는 여전히 `go-utils-mysql` 및 `go-utils-redis`로 명명됨
- Current version: v1.4.016
- 현재 버전: v1.4.016

---

## [v1.4.015] - 2025-10-14

### Fixed / 수정

- **FIX**: Fixed config file path resolution in examples/redis/main.go
- **수정**: examples/redis/main.go의 설정 파일 경로 해석 수정
- Changed from hardcoded relative path to dynamic project root resolution
- 하드코딩된 상대 경로에서 동적 프로젝트 루트 해석으로 변경
- Uses `filepath.Join(projectRoot, "cfg", "database-redis.yaml")` like MySQL examples
- MySQL 예제와 동일하게 `filepath.Join(projectRoot, "cfg", "database-redis.yaml")` 사용
- Added `path/filepath` import for cross-platform path handling
- 크로스 플랫폼 경로 처리를 위한 `path/filepath` import 추가

### Updated / 업데이트

- **DOCS**: Updated examples/redis/README.md documentation links
- **문서**: examples/redis/README.md 문서 링크 업데이트
- Removed "Coming soon" markers from USER_MANUAL.md and DEVELOPER_GUIDE.md links
- USER_MANUAL.md 및 DEVELOPER_GUIDE.md 링크에서 "Coming soon" 표시 제거
- All comprehensive documentation is now available and accessible
- 모든 포괄적인 문서가 이제 사용 가능하고 접근 가능함

### Verified / 검증

- All Redis package tests pass successfully (18 tests)
- 모든 Redis 패키지 테스트 성공적으로 통과 (18개 테스트)
- All Redis examples run correctly (8 examples)
- 모든 Redis 예제 정상 실행 (8개 예제)
- Docker auto-start/stop functionality verified
- Docker 자동 시작/중지 기능 검증
- Configuration loading from YAML works correctly
- YAML에서 설정 로드 정상 작동
- Examples work from any directory with proper path resolution
- 적절한 경로 해석으로 모든 디렉토리에서 예제 작동

### Notes / 참고사항

- Redis package (v1.4.x) comprehensive review and verification complete
- Redis 패키지 (v1.4.x) 종합 검토 및 검증 완료
- All documentation (README, USER_MANUAL, DEVELOPER_GUIDE) verified
- 모든 문서 (README, USER_MANUAL, DEVELOPER_GUIDE) 검증 완료
- Package is production-ready
- 패키지 프로덕션 준비 완료
- Current version: v1.4.015
- 현재 버전: v1.4.015

---

## [v1.4.014] - 2025-10-14

### Added / 추가
- **TESTS**: Replaced MySQL package tests with end-to-end coverage powered by Docker MySQL auto orchestration
- **테스트**: Docker MySQL 자동 오케스트레이션을 활용한 MySQL 패키지 엔드투엔드 테스트로 교체
- **INFRA**: `TestMain`-based helper that starts MySQL once for the suite and tears it down after completion
- **인프라**: 테스트 전용 `TestMain` 헬퍼 추가로 MySQL을 한 번만 시작하고 종료 시 자동 정리
- **DOCS**: Updated Redis manuals to v1.4.014 to reflect the latest release number
- **문서**: 최신 릴리스를 반영하도록 Redis 매뉴얼 버전을 v1.4.014로 갱신

### Changed / 변경
- **REFACTOR**: Removed legacy MySQL unit tests that assumed missing databases and rewrote them as deterministic integration tests
- **리팩터링**: 데이터베이스 부재를 가정했던 기존 MySQL 단위 테스트를 제거하고 결정적인 통합 테스트로 재작성
- **CLEANUP**: Ensured MySQL test runs start from a clean schema using table truncation helpers
- **정리**: 테이블 초기화 헬퍼로 MySQL 테스트가 항상 깨끗한 스키마에서 시작되도록 보장

### Fixed / 수정
- **FIX**: Fixed config file path in examples/redis/main.go
- **수정**: examples/redis/main.go의 설정 파일 경로 수정
- Changed from hardcoded `../../cfg/database-redis.yaml` to dynamic path resolution
- 하드코딩된 `../../cfg/database-redis.yaml`에서 동적 경로 해석으로 변경
- Now uses `filepath.Join(projectRoot, "cfg", "database-redis.yaml")` like MySQL examples
- MySQL 예제와 같이 `filepath.Join(projectRoot, "cfg", "database-redis.yaml")` 사용
- **DOCS**: Updated examples/redis/README.md documentation links
- **문서**: examples/redis/README.md 문서 링크 업데이트
- Removed "Coming soon" markers from USER_MANUAL.md and DEVELOPER_GUIDE.md links
- USER_MANUAL.md 및 DEVELOPER_GUIDE.md 링크에서 "Coming soon" 표시 제거

### Verified / 검증
- All Redis package tests pass (18 tests)
- 모든 Redis 패키지 테스트 통과 (18개 테스트)
- All Redis examples run successfully (8 examples)
- 모든 Redis 예제 성공적으로 실행 (8개 예제)
- Redis documentation complete and accurate
- Redis 문서 완전하고 정확함

### Notes / 참고사항
- Redis package (v1.4.x) comprehensive review complete
- Redis 패키지 (v1.4.x) 종합 검토 완료
- Current version: v1.4.014
- 현재 버전: v1.4.014

---

## [v1.4.013] - 2025-10-14

### Added / 추가
- **TESTS**: Comprehensive Redis test suite covering strings, hashes, lists, sets, sorted sets, keys, pipelines, transactions, and pub/sub
- **테스트**: 문자열, 해시, 리스트, 집합, 정렬 집합, 키, 파이프라인, 트랜잭션, pub/sub을 아우르는 Redis 테스트 스위트 추가
- **INFRA**: Shared Redis test helper that auto-starts Docker Redis when needed and flushes databases between tests
- **인프라**: 필요 시 Docker Redis를 자동으로 시작하고 각 테스트 사이 DB를 초기화하는 테스트 헬퍼 추가

### Changed / 변경
- **REFACTOR**: Reworked `client_test.go` to validate configuration options, ping health checks, and close behavior
- **리팩터링**: `client_test.go`를 재구성하여 설정 옵션, Ping 헬스 체크, Close 동작을 검증하도록 개선
- **DOCS**: Bumped Redis documentation headers to v1.4.013 to match latest release
- **문서**: 최신 릴리스에 맞춰 Redis 문서 헤더를 v1.4.013으로 업데이트

### Notes / 참고사항
- Current version: v1.4.013
- 현재 버전: v1.4.013

---

## [v1.4.012] - 2025-10-14

### Changed / 변경
- **DOCS**: Updated Redis documentation to reflect `cfg/database-redis.yaml`
- **문서**: Redis 문서를 `cfg/database-redis.yaml` 명명 규칙에 맞게 업데이트
- Clarified configuration instructions in `database/redis/README.md`
- `database/redis/README.md`의 설정 지침을 명확히 설명
- Updated Redis Go compatibility badge to Go 1.18+
- Redis Go 호환 배지를 Go 1.18+로 업데이트
- Bumped documentation version headers to v1.4.012 for Redis manuals
- Redis 매뉴얼 문서 버전 헤더를 v1.4.012로 갱신

### Notes / 참고사항
- Current version: v1.4.012
- 현재 버전: v1.4.012

---

## [v1.4.011] - 2025-10-14

### Added / 추가
- **DOCS**: Created comprehensive DEVELOPER_GUIDE.md for Redis package
- **문서**: Redis 패키지에 대한 포괄적인 DEVELOPER_GUIDE.md 생성
- Complete developer guide with 1550+ lines covering all technical details
- 모든 기술 세부사항을 다루는 1550줄 이상의 완전한 개발자 가이드
- Table of contents with 10 major sections
- 10개 주요 섹션이 있는 목차
- Architecture overview with design philosophy and high-level diagrams
- 설계 철학 및 상위 수준 다이어그램이 포함된 아키텍처 개요
- Complete package structure with file responsibilities (17 files, ~1888 lines)
- 파일 책임이 포함된 완전한 패키지 구조 (17개 파일, 약 1888줄)
- Core components documentation:
- 핵심 컴포넌트 문서화:
  - Client structure and lifecycle
  - Configuration and options pattern
  - Error handling and classification
  - Retry logic with exponential backoff
  - Health check implementation
- Internal implementation details:
- 내부 구현 세부사항:
  - String operations flow
  - Generic type-safe methods (GetAs[T], HGetAllAs[T])
  - Pipeline operations
  - Transaction operations with optimistic locking
  - Pub/Sub operations
- Design patterns documentation:
- 디자인 패턴 문서화:
  - Functional Options Pattern
  - Type Alias Pattern (hiding go-redis dependency)
  - Retry with Exponential Backoff
  - Health Check Pattern
  - Generic Methods Pattern
- Adding new features guide with step-by-step examples
- 단계별 예제가 포함된 새 기능 추가 가이드
- Comprehensive testing guide:
- 포괄적인 테스트 가이드:
  - Running tests and benchmarks
  - Test structure and categories
  - Test best practices
  - Integration testing with Docker
- Performance optimization guide:
- 성능 최적화 가이드:
  - Benchmarking tips
  - Connection pooling recommendations
  - Pipeline usage for batch operations
  - Timeout guidelines
  - Performance monitoring
- Contributing guidelines with code review checklist
- 코드 리뷰 체크리스트가 포함된 기여 가이드라인
- Code style guide:
- 코드 스타일 가이드:
  - Naming conventions
  - Documentation style
  - Code organization
  - Error handling patterns
  - Testing style
- Appendix with glossary and useful links
- 용어집 및 유용한 링크가 포함된 부록
- All content is bilingual (English/Korean)
- 모든 내용은 이중 언어 (영문/한글)

### Created Files / 생성된 파일
- `docs/database/redis/DEVELOPER_GUIDE.md` (1551 lines)

### Notes / 참고사항
- DEVELOPER_GUIDE.md completes Redis package documentation
- DEVELOPER_GUIDE.md로 Redis 패키지 문서화 완료
- Together with USER_MANUAL.md, provides complete documentation for users and developers
- USER_MANUAL.md와 함께 사용자 및 개발자를 위한 완전한 문서 제공
- Redis package documentation work is now complete
- Redis 패키지 문서화 작업 완료
- Current version: v1.4.011
- 현재 버전: v1.4.011

---

## [v1.4.010] - 2025-10-14

### Added / 추가
- **DOCS**: Created comprehensive USER_MANUAL.md for Redis package
- **문서**: Redis 패키지에 대한 포괄적인 USER_MANUAL.md 생성
- Complete user manual with 1200+ lines covering all Redis package features
- Redis 패키지의 모든 기능을 다루는 1200줄 이상의 완전한 사용자 매뉴얼
- Table of contents with 11 major sections
- 11개 주요 섹션이 있는 목차
- Introduction explaining extreme simplicity philosophy (20+ lines → 2 lines)
- 극도의 간결함 철학 설명 소개 (20줄 이상 → 2줄)
- Installation instructions for Docker and local Redis setup
- Docker 및 로컬 Redis 설정을 위한 설치 지침
- Quick Start with 3 progressive examples
- 3개 단계별 예제가 포함된 빠른 시작
- Configuration reference with detailed table of all 10 options
- 10개 모든 옵션의 상세 테이블이 포함된 설정 참조
- Core operations documentation for all 6 Redis data types:
- 6가지 Redis 데이터 타입 모두에 대한 핵심 작업 문서화:
  - String Operations (Set, Get, MSet, MGet, Incr, Decr, Append, SetNX, SetEX)
  - Hash Operations (HSet, HGet, HGetAll, HGetAllAs[T], HSetMap, HDel, HExists, HLen, HIncrBy)
  - List Operations (LPush, RPush, LPop, RPop, LRange, LLen, LIndex, LSet, LRem)
  - Set Operations (SAdd, SRem, SMembers, SIsMember, SCard, SUnion, SInter, SDiff)
  - Sorted Set Operations (ZAdd, ZRange, ZRangeByScore, ZRem, ZScore, ZIncrBy, ZCard, ZRank)
  - Key Operations (Del, Exists, Expire, TTL, Keys, Scan, Rename, Type)
- Advanced features documentation:
- 고급 기능 문서화:
  - Pipeline operations for batch command execution
  - Transactions with optimistic locking (WATCH/MULTI/EXEC)
  - Pub/Sub for message publishing and subscribing
- 6 usage patterns with code examples:
- 코드 예제가 포함된 6개 사용 패턴:
  - Session Storage, Cache with TTL, Rate Limiting, Distributed Lock, Task Queue, Leaderboard
- 5 common use cases with complete implementation code:
- 완전한 구현 코드가 포함된 5개 일반 사용 사례:
  - User Session Management, API Rate Limiting, Real-time Chat, Leaderboard System, Job Queue
- 10 best practices with good/bad examples
- 좋은/나쁜 예제가 포함된 10개 모범 사례
- Troubleshooting section with 6 common problems and solutions
- 6개 일반 문제 및 해결책이 포함된 문제 해결 섹션
- FAQ with 15 questions and detailed answers
- 15개 질문 및 상세 답변이 포함된 FAQ
- All content is bilingual (English/Korean)
- 모든 내용은 이중 언어 (영문/한글)

### Created Files / 생성된 파일
- `docs/database/redis/USER_MANUAL.md` (1200+ lines)

### Notes / 참고사항
- USER_MANUAL.md follows same comprehensive format as MySQL package
- USER_MANUAL.md는 MySQL 패키지와 동일한 포괄적인 형식 사용
- Next: DEVELOPER_GUIDE.md creation planned for v1.4.011
- 다음: DEVELOPER_GUIDE.md 생성 v1.4.011에 계획됨
- Current version: v1.4.010
- 현재 버전: v1.4.010

---

## [v1.4.009] - 2025-10-14

### Changed / 변경
- **FIX**: Remove unnecessary go-redis import dependency from examples
- **수정**: 예제에서 불필요한 go-redis import 의존성 제거
- Exported `Pipeliner` type as type alias in database/redis package
- database/redis 패키지에서 `Pipeliner` 타입을 타입 별칭으로 export
- Users no longer need to import `github.com/redis/go-redis/v9` directly
- 사용자는 더 이상 `github.com/redis/go-redis/v9`를 직접 import할 필요 없음
- Updated Pipeline() and TxPipeline() to use redis.Pipeliner instead of go-redis.Pipeliner
- Pipeline() 및 TxPipeline()을 go-redis.Pipeliner 대신 redis.Pipeliner 사용하도록 업데이트
- Updated Transaction Exec() to use redis.Pipeliner
- Transaction Exec()을 redis.Pipeliner 사용하도록 업데이트

### Updated Files / 업데이트된 파일
- `database/redis/types.go`: Added Pipeliner type alias
- `database/redis/pipeline.go`: Updated to use Pipeliner type, removed redis import
- `database/redis/transaction.go`: Updated Exec() to use Pipeliner type
- `examples/redis/main.go`: Removed go-redis import, use redis.Pipeliner

### Benefits / 장점
- Cleaner API - users only need to import our package
- 더 깨끗한 API - 사용자는 우리 패키지만 import하면 됨
- Reduced dependency exposure
- 의존성 노출 감소
- Simpler example code
- 더 간단한 예제 코드

### Notes / 참고사항
- This is a non-breaking change as Pipeliner is a type alias
- Pipeliner가 타입 별칭이므로 중단되지 않는 변경사항
- Current version: v1.4.009
- 현재 버전: v1.4.009

---

## [v1.4.008] - 2025-10-14

### Changed / 변경
- **REFACTOR**: Updated examples/redis/main.go to use logging package and results directory
- **리팩토링**: examples/redis/main.go를 logging 패키지 사용 및 results 디렉토리 사용하도록 업데이트
- Redis examples now use structured logging with automatic log file creation
- Redis 예제는 이제 자동 로그 파일 생성과 함께 구조화된 로깅 사용
- Added Docker auto-start and auto-stop functionality (like MySQL examples)
- Docker 자동 시작 및 자동 중지 기능 추가 (MySQL 예제와 동일)
- All example functions now accept logger parameter for consistent logging
- 모든 예제 함수는 이제 일관된 로깅을 위해 logger 매개변수 수용
- Results saved to `./results/logs/redis_example_YYYYMMDD_HHMMSS.log`
- 결과는 `./results/logs/redis_example_YYYYMMDD_HHMMSS.log`에 저장됨

### Changed / 변경 (Config Files)
- **BREAKING**: Renamed `cfg/database.yaml` to `cfg/database-mysql.yaml`
- **중단**: `cfg/database.yaml`을 `cfg/database-mysql.yaml`로 이름 변경
- **BREAKING**: Renamed `cfg/redis.yaml` to `cfg/database-redis.yaml`
- **중단**: `cfg/redis.yaml`을 `cfg/database-redis.yaml`로 이름 변경
- Updated all references in code and documentation to use new config file names
- 코드 및 문서의 모든 참조를 새 설정 파일 이름 사용하도록 업데이트
- Consistent naming convention: `database-{provider}.yaml` for all database configs
- 일관된 명명 규칙: 모든 데이터베이스 설정에 대해 `database-{provider}.yaml`

### Updated Files / 업데이트된 파일
- `examples/redis/main.go`: Complete rewrite with logging integration
- `examples/mysql/main.go`: Updated config file path references
- `CLAUDE.md`: Updated all config file path references
- `cfg/database.yaml` → `cfg/database-mysql.yaml`
- `cfg/redis.yaml` → `cfg/database-redis.yaml`

### Notes / 참고사항
- Redis examples now follow same pattern as MySQL examples for consistency
- Redis 예제는 이제 일관성을 위해 MySQL 예제와 동일한 패턴 사용
- All database configuration files now use `database-{provider}.yaml` naming
- 모든 데이터베이스 설정 파일은 이제 `database-{provider}.yaml` 명명 사용
- Current version: v1.4.008
- 현재 버전: v1.4.008

---

## [v1.4.007] - 2025-10-14

### Added / 추가
- **DOCS**: Created comprehensive examples/redis/README.md
- **문서**: 포괄적인 examples/redis/README.md 생성
- Documented all 8 example categories with detailed descriptions
- 상세 설명과 함께 8개 예제 카테고리 모두 문서화
- Added running instructions and prerequisites (Docker/Local Redis)
- 실행 지침 및 전제 조건 추가 (Docker/로컬 Redis)
- Added configuration section for customizing Redis connection
- Redis 연결 사용자 정의를 위한 설정 섹션 추가
- Added output example showing what users can expect
- 사용자가 기대할 수 있는 출력 예제 추가
- Added comprehensive troubleshooting section
- 포괄적인 문제 해결 섹션 추가
- Added links to additional resources and documentation
- 추가 리소스 및 문서 링크 추가

### Examples Documented / 문서화된 예제
1. String Operations - Set, Get, MSet, MGet, Incr
2. Hash Operations - HSet, HGet, HGetAll, HGetAllAs[T], HIncrBy
3. List Operations - LPush, RPush, LPop, RPop, LRange, LLen
4. Set Operations - SAdd, SMembers, SCard, SUnion, SInter, SDiff
5. Sorted Set Operations - ZAdd, ZRange, ZRangeByScore, ZScore
6. Key Operations - Del, Exists, Expire, TTL, Keys, Type
7. Pipeline Operations - Batch command execution
8. Transaction Operations - Optimistic locking with WATCH/MULTI/EXEC

### Notes / 참고사항
- README follows same format as MySQL examples for consistency
- README는 일관성을 위해 MySQL 예제와 동일한 형식 사용
- All examples include bilingual descriptions (English/Korean)
- 모든 예제는 이중 언어 설명 포함 (영문/한글)
- Current version: v1.4.007
- 현재 버전: v1.4.007

---

## [v1.4.006] - 2025-10-14

### Changed / 변경
- **DOCS**: Updated CLAUDE.md with comprehensive Redis package architecture
- **문서**: Redis 패키지 아키텍처를 포함하여 CLAUDE.md 업데이트
- Added Redis package to subpackage structure section
- 서브패키지 구조 섹션에 Redis 패키지 추가
- Added complete Redis package architecture documentation (similar to MySQL section)
- MySQL 섹션과 유사한 완전한 Redis 패키지 아키텍처 문서화 추가
- Documented all 6 data types: String, Hash, List, Set, Sorted Set, Key operations
- 6가지 데이터 타입 문서화: String, Hash, List, Set, Sorted Set, Key 작업
- Documented advanced features: Pipeline, Transaction, Pub/Sub
- 고급 기능 문서화: Pipeline, Transaction, Pub/Sub
- Added design philosophy with Before/After code examples (20+ lines → 2 lines)
- Before/After 코드 예제를 포함한 설계 철학 추가 (20줄 이상 → 2줄)
- Updated version history to v1.4.x as current version
- 버전 히스토리를 v1.4.x로 현재 버전 업데이트
- Updated external dependencies to include redis/go-redis/v9
- redis/go-redis/v9를 포함하도록 외부 의존성 업데이트
- Updated CHANGELOG file structure to include CHANGELOG-v1.4.md
- CHANGELOG-v1.4.md를 포함하도록 CHANGELOG 파일 구조 업데이트
- Updated documentation directory structure with redis package docs
- redis 패키지 문서를 포함하도록 문서 디렉토리 구조 업데이트
- Updated build and test commands to include redis package
- redis 패키지를 포함하도록 빌드 및 테스트 명령 업데이트
- Updated example execution section with Redis examples
- Redis 예제를 포함하도록 예제 실행 섹션 업데이트
- Renamed "MySQL 개발 워크플로우" to "Docker 개발 워크플로우"
- "MySQL 개발 워크플로우"를 "Docker 개발 워크플로우"로 변경
- Added Docker Redis workflow with management scripts
- 관리 스크립트를 포함한 Docker Redis 워크플로우 추가
- Added Redis package testing guidelines
- Redis 패키지 테스트 가이드라인 추가
- Added Redis package development guidelines with file selection guide
- 파일 선택 가이드를 포함한 Redis 패키지 개발 가이드라인 추가
- Updated version history context to include v1.4.x details
- v1.4.x 세부 정보를 포함하도록 버전 히스토리 컨텍스트 업데이트

### Notes / 참고사항
- CLAUDE.md now has complete documentation for all 4 major packages
- CLAUDE.md에 이제 4개 주요 패키지 모두에 대한 완전한 문서 포함
- Redis package section follows same format as MySQL package for consistency
- Redis 패키지 섹션은 일관성을 위해 MySQL 패키지와 동일한 형식 사용
- Current version: v1.4.006
- 현재 버전: v1.4.006

---

## [v1.4.005] - 2025-10-14

### Added / 추가
- **COMPLETE**: Implemented full database/redis package with all core features
- **완료**: 모든 핵심 기능을 갖춘 database/redis 패키지 완전 구현
- Core client with connection management, retry logic, and health check
- 연결 관리, 재시도 로직, 헬스 체크를 갖춘 핵심 클라이언트
- String operations: Set, Get, MGet, MSet, Incr, Decr, Append, SetNX, SetEX
- 문자열 작업: Set, Get, MGet, MSet, Incr, Decr, Append, SetNX, SetEX
- Hash operations: HSet, HGet, HGetAll, HSetMap, HDel, HExists, HLen, HIncrBy
- 해시 작업: HSet, HGet, HGetAll, HSetMap, HDel, HExists, HLen, HIncrBy
- List operations: LPush, RPush, LPop, RPop, LRange, LLen, LIndex, LSet, LRem
- 리스트 작업: LPush, RPush, LPop, RPop, LRange, LLen, LIndex, LSet, LRem
- Set operations: SAdd, SRem, SMembers, SIsMember, SCard, SUnion, SInter, SDiff
- 집합 작업: SAdd, SRem, SMembers, SIsMember, SCard, SUnion, SInter, SDiff
- Sorted set operations: ZAdd, ZRange, ZRangeByScore, ZRem, ZScore, ZIncrBy
- 정렬 집합 작업: ZAdd, ZRange, ZRangeByScore, ZRem, ZScore, ZIncrBy
- Key operations: Del, Exists, Expire, TTL, Keys, Scan, Rename, Type
- 키 작업: Del, Exists, Expire, TTL, Keys, Scan, Rename, Type
- Pipeline support for batch operations
- 배치 작업을 위한 파이프라인 지원
- Transaction support with optimistic locking (WATCH/MULTI/EXEC)
- 낙관적 잠금을 사용한 트랜잭션 지원 (WATCH/MULTI/EXEC)
- Pub/Sub support for message publishing and subscribing
- 메시지 발행 및 구독을 위한 Pub/Sub 지원
- Type-safe generic methods: GetAs[T], HGetAllAs[T]
- 타입 안전 제네릭 메서드: GetAs[T], HGetAllAs[T]
- Comprehensive test suite with 8 test cases
- 8개 테스트 케이스를 포함한 종합 테스트 스위트
- Complete examples demonstrating all features (8 example functions)
- 모든 기능을 시연하는 완전한 예제 (8개 예제 함수)
- Comprehensive README with API reference and usage examples
- API 참조 및 사용 예제가 포함된 종합 README

### Features / 기능
- Auto-retry with exponential backoff for network errors
- 네트워크 에러에 대한 지수 백오프를 사용한 자동 재시도
- Connection pooling for high performance
- 고성능을 위한 연결 풀링
- Background health checking
- 백그라운드 헬스 체크
- Context support for cancellation and timeout
- 취소 및 타임아웃을 위한 Context 지원
- Options pattern for flexible configuration
- 유연한 설정을 위한 옵션 패턴

---

## [v1.4.004] - 2025-10-14

### Added / 추가
- Docker Redis setup with docker-compose.yml configuration
- Docker Redis 설정 및 docker-compose.yml 구성
- Created redis.conf for Redis server configuration
- Redis 서버 설정을 위한 redis.conf 생성
- Created 4 Redis management scripts:
- 4개의 Redis 관리 스크립트 생성:
  - `docker-redis-start.sh` - Start Docker Redis / Docker Redis 시작
  - `docker-redis-stop.sh` - Stop and cleanup Docker Redis / Docker Redis 중지 및 정리
  - `docker-redis-logs.sh` - View Redis logs / Redis 로그 확인
  - `docker-redis-cli.sh` - Connect to Redis CLI / Redis CLI 연결
- Created cfg/redis.yaml for Redis package configuration
- Redis 패키지 설정을 위한 cfg/redis.yaml 생성
- Created redis/README.md with usage and troubleshooting guide
- 사용법 및 문제 해결 가이드가 포함된 redis/README.md 생성

### Configuration / 설정
- Redis 7-alpine image with persistent data volume
- 영구 데이터 볼륨이 있는 Redis 7-alpine 이미지
- AOF (Append Only File) persistence enabled
- AOF (Append Only File) 영속성 활성화
- Health check with redis-cli ping
- redis-cli ping을 사용한 헬스 체크
- 16 databases (0-15) available
- 16개 데이터베이스 (0-15) 사용 가능

---

## [v1.4.003] - 2025-10-14

### Added / 추가
- Created WORK_PLAN.md for Redis package with comprehensive implementation roadmap
- Redis 패키지에 대한 포괄적인 구현 로드맵이 포함된 WORK_PLAN.md 생성
- Defined 5 phases: Foundation, Core Features, Advanced Features, Testing & Documentation, Release
- 5단계 정의: 기초, 핵심 기능, 고급 기능, 테스팅 및 문서화, 릴리스
- Planned 20+ tasks with clear acceptance criteria and dependencies
- 명확한 수용 기준 및 의존성이 있는 20개 이상의 작업 계획
- Quality checklist for code, testing, documentation, performance, and Docker
- 코드, 테스팅, 문서화, 성능, Docker에 대한 품질 체크리스트

---

## [v1.4.002] - 2025-10-14

### Added / 추가
- **NEW Package**: `database/redis` package - Extreme simplicity Redis client
- **새로운 패키지**: `database/redis` 패키지 - 극도로 간단한 Redis 클라이언트
- Created DESIGN_PLAN.md for Redis package with comprehensive architecture design
- Redis 패키지에 대한 포괄적인 아키텍처 설계가 포함된 DESIGN_PLAN.md 생성

### Documentation / 문서
- Documented Redis package design philosophy: "If it's not dramatically simpler, don't build it"
- Redis 패키지 설계 철학 문서화: "극적으로 간단하지 않으면 만들지 마세요"
- Planned Simple API for String, Hash, List, Set, Sorted Set operations
- String, Hash, List, Set, Sorted Set 작업을 위한 Simple API 계획
- Designed Pipeline and Transaction support
- Pipeline 및 Transaction 지원 설계
- Designed Pub/Sub support
- Pub/Sub 지원 설계

---

## [v1.4.001] - 2025-10-14

### Changed / 변경
- Version bumped to v1.4.001 to start new Redis package development
- Redis 패키지 개발 시작을 위해 버전을 v1.4.001로 증가
- Started v1.4.x series for database/redis package
- database/redis 패키지를 위한 v1.4.x 시리즈 시작

---

## Version Overview / 버전 개요

**v1.4.x Series Goals / v1.4.x 시리즈 목표**:
- Implement `database/redis` package with extreme simplicity (20+ lines → 2 lines)
- 극도의 간결함으로 `database/redis` 패키지 구현 (20줄 이상 → 2줄)
- Auto-everything: connection management, retry, reconnect, resource cleanup
- 모든 것 자동화: 연결 관리, 재시도, 재연결, 리소스 정리
- Type-safe API with generics
- 제네릭을 사용한 타입 안전 API
- Simple transaction and pipeline support
- 간단한 트랜잭션 및 파이프라인 지원
- Pub/Sub support
- Pub/Sub 지원
- Docker-based testing with redis:alpine
- redis:alpine을 사용한 Docker 기반 테스트
- Comprehensive documentation (README, USER_MANUAL, DEVELOPER_GUIDE)
- 포괄적인 문서화 (README, USER_MANUAL, DEVELOPER_GUIDE)
