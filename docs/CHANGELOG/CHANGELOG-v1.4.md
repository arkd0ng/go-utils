# CHANGELOG - v1.4.x

All notable changes for version 1.4.x will be documented in this file.

v1.4.x 버전의 모든 주목할 만한 변경사항이 이 파일에 문서화됩니다.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

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
