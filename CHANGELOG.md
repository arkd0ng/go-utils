# Changelog / 변경 이력

All notable changes to this project will be documented in this file.

이 프로젝트의 모든 주요 변경사항이 이 파일에 기록됩니다.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

형식은 [Keep a Changelog](https://keepachangelog.com/en/1.0.0/)를 따르며,
이 프로젝트는 [Semantic Versioning](https://semver.org/spec/v2.0.0.html)을 준수합니다.

## Version Overview / 버전 개요

This file contains a high-level overview of major and minor versions. For detailed patch-level changes, please refer to the version-specific changelog files.

이 파일은 메이저 및 마이너 버전의 개요만 포함합니다. 패치 레벨의 상세 변경사항은 버전별 changelog 파일을 참조하세요.

---

## [v1.6.x] - Time Utilities Package / 시간 유틸리티 패키지 (현재 / Current)

**Focus / 초점**: Extreme simplicity time and date utilities with KST default timezone / KST 기본 타임존을 갖춘 극도로 간단한 시간 및 날짜 유틸리티

**Detailed Changes / 상세 변경사항**: See / 참조 [docs/CHANGELOG/CHANGELOG-v1.6.md](docs/CHANGELOG/CHANGELOG-v1.6.md)

### Highlights / 주요 사항
- **Complete timeutil package**: 20 lines → 1 line of code / 완전한 timeutil 패키지: 20줄 → 1줄 코드
- **80+ functions**: Time difference, timezone, arithmetic, formatting, parsing, comparison, age, relative time, unix timestamp, business days / 80개 이상 함수
- **KST default timezone**: Asia/Seoul (GMT+9) as package-wide default / KST 기본 타임존: Asia/Seoul (GMT+9)를 패키지 전체 기본값으로
- **Custom format tokens**: YYYY-MM-DD instead of Go's 2006-01-02 / 커스텀 포맷 토큰: Go의 2006-01-02 대신 YYYY-MM-DD
- **Business day support**: Weekend and holiday-aware date calculations / 영업일 지원: 주말 및 공휴일을 고려한 날짜 계산
- **Korean localization**: Korean format and holiday support / 한국어 지역화: 한국 형식 및 공휴일 지원
- **Thread-safe**: Timezone caching with sync.RWMutex / 스레드 안전: sync.RWMutex를 사용한 타임존 캐싱
- **Zero dependencies**: Standard library only / 제로 의존성: 표준 라이브러리만

**Key Design Principles / 주요 설계 원칙**:
- Extreme Simplicity: 20 lines → 1 line / 극도의 간결함: 20줄 → 1줄
- Human-Readable: Intuitive function names / 사람이 읽기 쉬움: 직관적인 함수 이름
- Zero Configuration: No setup required / 제로 설정: 설정 불필요

**Latest Version / 최신 버전**: v1.6.001 (2025-10-14)

---

## [v1.5.x] - String Utilities Package / 문자열 유틸리티 패키지

**Focus / 초점**: Extreme simplicity string manipulation utilities / 극도로 간단한 문자열 처리 유틸리티

**Detailed Changes / 상세 변경사항**: See / 참조 [docs/CHANGELOG/CHANGELOG-v1.5.md](docs/CHANGELOG/CHANGELOG-v1.5.md)

### Highlights / 주요 사항
- **Complete stringutil package**: 20 lines → 1 line of code / 완전한 stringutil 패키지: 20줄 → 1줄 코드
- **37 functions**: Case conversion, manipulation, validation, search & replace, utilities / 37개 함수
- **Unicode-safe**: All operations use rune (not byte) for proper Unicode support / 유니코드 안전
- **Zero dependencies**: Standard library only / 제로 의존성: 표준 라이브러리만
- **Functional programming**: Map and Filter for functional-style operations / 함수형 프로그래밍
- **Comprehensive documentation**: USER_MANUAL and DEVELOPER_GUIDE / 포괄적인 문서화

**Key Design Principles / 주요 설계 원칙**:
- Extreme Simplicity: 20 lines → 1 line / 극도의 간결함: 20줄 → 1줄
- Unicode First: Works with Korean, emoji, and all Unicode characters / 유니코드 우선
- Practical Over Perfect: 99% use cases covered / 실용성 중심: 99% 사용 사례 커버

**Latest Version / 최신 버전**: v1.5.009 (2025-10-14)

---

## [v1.4.x] - Redis Package / Redis 패키지

**Focus / 초점**: Extreme simplicity Redis client with comprehensive operations support / 포괄적인 작업 지원을 갖춘 극도로 간단한 Redis 클라이언트

**Detailed Changes / 상세 변경사항**: See / 참조 [docs/CHANGELOG/CHANGELOG-v1.4.md](docs/CHANGELOG/CHANGELOG-v1.4.md)

### Highlights / 주요 사항
- **Complete Redis package**: 20 lines → 2 lines of code / 완전한 Redis 패키지: 20줄 → 2줄 코드
- **60+ methods**: String, Hash, List, Set, Sorted Set, Key operations / 60개 이상 메서드
- **Advanced features**: Pipeline, Transaction, Pub/Sub / 고급 기능: 파이프라인, 트랜잭션, Pub/Sub
- **Auto-retry**: Exponential backoff for network errors / 자동 재시도: 네트워크 에러에 대한 지수 백오프
- **Connection pooling**: Built-in connection pool for high performance / 연결 풀링: 고성능을 위한 내장 연결 풀
- **Health check**: Background health monitoring / 헬스 체크: 백그라운드 헬스 모니터링
- **Type-safe**: Generic methods for type-safe operations / 타입 안전: 타입 안전 작업을 위한 제네릭 메서드
- **Docker setup**: Automated Docker Redis with management scripts / Docker 설정: 관리 스크립트를 사용한 자동화된 Docker Redis

**Key Design Principles / 주요 설계 원칙**:
- Extreme Simplicity: If not dramatically simpler, don't build it / 극도의 간결함: 극적으로 간단하지 않으면 만들지 마세요
- Auto Everything: Connection, retry, reconnect, cleanup / 모든 것 자동: 연결, 재시도, 재연결, 정리
- Context Support: All methods support context for cancellation / Context 지원: 모든 메서드가 취소를 위한 context 지원

**Latest Version / 최신 버전**: v1.4.005 (2025-10-14)

---

## [v1.3.x] - MySQL Package / MySQL 패키지

**Focus / 초점**: Extreme simplicity MySQL/MariaDB package with zero-downtime credential rotation / 무중단 자격 증명 순환을 갖춘 극도로 간단한 MySQL/MariaDB 패키지

**Detailed Changes / 상세 변경사항**: See / 참조 [docs/CHANGELOG/CHANGELOG-v1.3.md](docs/CHANGELOG/CHANGELOG-v1.3.md)

### Highlights / 주요 사항
- **Complete MySQL package**: 30 lines → 2 lines of code / 완전한 MySQL 패키지: 30줄 → 2줄 코드
- **Three-layer API**: Simple, Query Builder, Raw SQL / 3계층 API: 간단, 쿼리 빌더, Raw SQL
- **Zero-downtime credential rotation**: Multiple connection pools with rolling rotation / 무중단 자격 증명 순환: 순환 교체 방식의 다중 연결 풀
- **Advanced features**: Batch, Upsert, Pagination, Soft Delete, Query Stats, Schema Inspector, Migration Helpers, CSV Export/Import / 고급 기능
- **Auto everything**: Connection management, retry, cleanup / 모든 것 자동: 연결 관리, 재시도, 정리

**Key Design Principles / 주요 설계 원칙**:
- Zero Mental Overhead: Connect once, forget about DB state / 한 번 연결하면 DB 상태를 잊어버려도 됨
- SQL-Like API: Close to actual SQL syntax / SQL 문법에 가까운 API
- "If not 10x simpler, don't build it" / "10배 간단하지 않으면 만들지 마세요"

**Latest Version / 최신 버전**: v1.3.017 (2025-10-14)

---

## [v1.2.x] - Documentation Work / 문서화 작업

**Focus / 초점**: Comprehensive documentation, CHANGELOG system, and project management / 종합 문서화, CHANGELOG 시스템, 프로젝트 관리

**Detailed Changes / 상세 변경사항**: See / 참조 [docs/CHANGELOG/CHANGELOG-v1.2.md](docs/CHANGELOG/CHANGELOG-v1.2.md)

### Highlights / 주요 사항
- Established CHANGELOG system with version-specific documentation / 버전별 문서화를 포함한 CHANGELOG 시스템 구축
- Created comprehensive version management rules / 포괄적인 버전 관리 규칙 생성
- Documented Git workflow and commit conventions / Git 워크플로우 및 커밋 규칙 문서화
- Improved project documentation structure / 프로젝트 문서 구조 개선

**Latest Version / 최신 버전**: v1.2.004 (2025-10-10)

---

## [v1.1.x] - Logging Package / 로깅 패키지

**Focus / 초점**: Enterprise-grade logging package with file rotation and structured logging / 파일 로테이션과 구조화된 로깅을 갖춘 엔터프라이즈급 로깅 패키지

**Detailed Changes / 상세 변경사항**: See / 참조 [docs/CHANGELOG/CHANGELOG-v1.1.md](docs/CHANGELOG/CHANGELOG-v1.1.md)

### Highlights / 주요 사항
- Automatic file rotation with lumberjack integration / lumberjack 통합 자동 파일 로테이션
- Structured logging with key-value pairs / 키-값 쌍을 사용한 구조화된 로깅
- Printf-style logging support / Printf 스타일 로깅 지원
- Automatic banner with app.yaml version management / app.yaml 버전 관리를 통한 자동 배너
- Multiple log levels (DEBUG, INFO, WARN, ERROR, FATAL) / 다중 로그 레벨
- Thread-safe concurrent logging / 스레드 안전 동시 로깅
- Dual output support (console and file) / 이중 출력 지원 (콘솔 및 파일)
- Colored console output / 색상 콘솔 출력
- Auto-extract app name from log filename / 로그 파일명에서 앱 이름 자동 추출

**Key Features / 주요 기능**:
- 7 patches (v1.1.000 to v1.1.007) / 7개 패치
- app.yaml version management / app.yaml 버전 관리
- Both structured and Printf-style logging / 구조화 및 Printf 스타일 로깅 모두 지원
- Comprehensive test suite (15+ tests) / 종합 테스트 스위트 (15개 이상)
- Production-ready with best practices / 모범 사례를 적용한 프로덕션 준비 완료

**Latest Version / 최신 버전**: v1.1.007 (2025-10-10)

---

## [v1.0.x] - Random Package / 랜덤 패키지

**Focus / 초점**: Cryptographically secure random string generation / 암호학적으로 안전한 랜덤 문자열 생성

**Detailed Changes / 상세 변경사항**: See / 참조 [docs/CHANGELOG/CHANGELOG-v1.0.md](docs/CHANGELOG/CHANGELOG-v1.0.md)

### Highlights / 주요 사항
- Cryptographically secure random string generation using crypto/rand / crypto/rand를 사용한 암호학적으로 안전한 랜덤 문자열 생성
- 14 different generation methods / 14가지 다양한 생성 메서드
- Flexible length parameters (fixed or range) / 유연한 길이 파라미터 (고정 또는 범위)
- Comprehensive error handling / 포괄적인 에러 처리
- Subpackage architecture / 서브패키지 아키텍처
- Bilingual documentation (English/Korean) / 이중 언어 문서화 (영문/한글)

**Available Methods / 사용 가능한 메서드**:
- Basic / 기본: Letters, Alnum, Digits, Complex, Standard
- Case-specific / 대소문자 구분: AlphaUpper, AlphaLower, AlnumUpper, AlnumLower
- Encoding / 인코딩: Hex, HexLower, Base64, Base64URL
- Custom / 사용자 정의: Custom(charset, length...)

**Key Features / 주요 기능**:
- 8 patches (v1.0.001 to v1.0.008) / 8개 패치
- Variadic parameters for flexible length / 유연한 길이를 위한 가변 인자
- Collision probability testing / 충돌 확률 테스트
- Breaking change: Migrated to subpackage structure / 주요 변경: 서브패키지 구조로 마이그레이션
- Breaking change: Added error return values / 주요 변경: 에러 반환값 추가

**Latest Version / 최신 버전**: v1.0.008 (2025-10-10)

---

## Version Numbering / 버전 번호 체계

This project uses semantic versioning: `vMAJOR.MINOR.PATCH`

이 프로젝트는 시맨틱 버저닝을 사용합니다: `vMAJOR.MINOR.PATCH`

- **MAJOR / 메이저**: Breaking changes / 호환성이 깨지는 변경사항
- **MINOR / 마이너**: New features (backwards compatible) / 새로운 기능 (하위 호환)
- **PATCH / 패치**: Bug fixes and minor improvements / 버그 수정 및 소규모 개선

### Patch Version Format / 패치 버전 형식
Patches use 3-digit format: v1.2.001, v1.2.002, etc.

패치는 3자리 형식을 사용합니다: v1.2.001, v1.2.002 등

---

## Links / 링크

- [GitHub Repository / 저장소](https://github.com/arkd0ng/go-utils)
- [Random Package Documentation / 랜덤 패키지 문서](random/README.md)
- [Logging Package Documentation / 로깅 패키지 문서](logging/README.md)
- [Project Documentation / 프로젝트 문서](CLAUDE.md)

---

**Maintained by / 관리자**: arkd0ng
**License / 라이선스**: MIT
