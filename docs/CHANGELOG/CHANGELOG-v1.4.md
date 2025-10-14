# CHANGELOG - v1.4.x

All notable changes for version 1.4.x will be documented in this file.

v1.4.x 버전의 모든 주목할 만한 변경사항이 이 파일에 문서화됩니다.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

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
