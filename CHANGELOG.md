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

## [v1.11.x] - Web Server Utilities Package / 웹 서버 유틸리티 패키지 (개발 중 / In Development)

**Focus / 초점**: Extreme simplicity web server utilities / 극도로 간단한 웹 서버 유틸리티

**Status / 상태**: In Development / 개발 중
**Branch / 브랜치**: `feature/v1.11.x-websvrutil`
**Latest Patch / 최신 패치**: v1.11.039 – comprehensive example suite with bilingual logging and documentation alignment

**Detailed Changes / 상세 변경사항**: See / 참조 [docs/CHANGELOG/CHANGELOG-v1.11.md](docs/CHANGELOG/CHANGELOG-v1.11.md)

### Planned Features / 계획된 기능
- **Simple Router**: RESTful routing with path parameters / 경로 매개변수를 가진 RESTful 라우팅
- **Middleware**: CORS, logging, recovery, auth, rate limiting / 미들웨어
- **Handler Helpers**: JSON response, error response, file serving / 핸들러 헬퍼
- **Request/Response Utilities**: Body binding, cookie, headers / 요청/응답 유틸리티
- **Server Management**: Graceful shutdown, hot reload, health check / 서버 관리

**Key Design Principles / 주요 설계 원칙**:
- Extreme Simplicity: 50+ lines → 5 lines / 극도의 간결함: 50줄 이상 → 5줄
- Zero Configuration: Sensible defaults for 99% of use cases / 제로 설정: 99% 사용 사례에 대한 합리적인 기본값
- Standard Compatible: Works with standard net/http / 표준 net/http와 호환
- Middleware Chain: Easy middleware composition / 쉬운 미들웨어 조합

**Current Version / 현재 버전**: v1.11.001 (2025-10-16)

**Roadmap / 로드맵**:
- Phase 1 (v1.11.001-005): Core router and basic handlers
- Phase 2 (v1.11.006-010): Middleware implementation
- Phase 3 (v1.11.011-015): Request/Response utilities
- Phase 4 (v1.11.016-020): Server management features
- Phase 5 (v1.11.021-025): Documentation and examples

---

## [v1.10.x] - HTTP Client Utilities Package / HTTP 클라이언트 유틸리티 패키지 (완료 / Completed)

**Focus / 초점**: Extreme simplicity HTTP client utilities / 극도로 간단한 HTTP 클라이언트 유틸리티

**Detailed Changes / 상세 변경사항**: See / 참조 [docs/CHANGELOG/CHANGELOG-v1.10.md](docs/CHANGELOG/CHANGELOG-v1.10.md)

### Highlights / 주요 사항
- **Complete httputil package (Phase 1-5)**: 30+ lines → 2-3 lines of code / 완전한 httputil 패키지 (Phase 1-5): 30줄 이상 → 2-3줄 코드
- **RESTful HTTP methods**: GET, POST, PUT, PATCH, DELETE + Context variants / RESTful HTTP 메서드: GET, POST, PUT, PATCH, DELETE + Context 변형
- **Simple API (26+ functions)**: Package-level convenience functions / 간단한 API (26개 이상 함수): 패키지 레벨 편의 함수
- **Response helpers (20+ methods)**: Status checks, body access, headers / 응답 헬퍼 (20개 이상 메서드): 상태 확인, 본문 접근, 헤더
- **File operations**: Upload/download with progress tracking / 파일 작업: 진행 상황 추적이 있는 업로드/다운로드
- **URL Builder**: Fluent API for building URLs with parameters / URL 빌더: 매개변수와 함께 URL을 구축하기 위한 Fluent API
- **Form Builder**: Fluent API for building forms with conditional fields / Form 빌더: 조건부 필드가 있는 폼을 구축하기 위한 Fluent API
- **Cookie Management**: In-memory and persistent cookie jars / 쿠키 관리: 메모리 내 및 지속성 쿠키 저장소
- **Automatic JSON handling**: Request/response encoding and decoding / 자동 JSON 처리: 요청/응답 인코딩 및 디코딩
- **Smart retry logic**: Exponential backoff with jitter for network errors and 5xx / 스마트 재시도 로직: 네트워크 에러 및 5xx에 대한 지터가 있는 지수 백오프
- **14 configuration options**: Timeout, headers, auth, retry, cookies, base URL, etc. / 14개 설정 옵션: 타임아웃, 헤더, 인증, 재시도, 쿠키, 기본 URL 등
- **Rich error types**: HTTPError, RetryError, TimeoutError with full context / 풍부한 에러 타입: 전체 컨텍스트가 있는 HTTPError, RetryError, TimeoutError
- **Options pattern**: Flexible configuration without breaking API / 옵션 패턴: API를 깨뜨리지 않는 유연한 설정
- **Zero external dependencies**: Standard library only (net/http, encoding/json, mime/multipart) / 외부 의존성 없음: 표준 라이브러리만 (net/http, encoding/json, mime/multipart)
- **Comprehensive documentation**: README, USER_MANUAL, DEVELOPER_GUIDE, WORK_PLAN / 포괄적인 문서화: README, USER_MANUAL, DEVELOPER_GUIDE, WORK_PLAN
- **60.9% test coverage**: 17+ tests, 100+ sub-tests / 60.9% 테스트 커버리지: 17개 이상 테스트, 100개 이상 하위 테스트

**Key Design Principles / 주요 설계 원칙**:
- Extreme Simplicity: 30+ lines → 2-3 lines / 극도의 간결함: 30줄 이상 → 2-3줄
- Auto Everything: Automatic JSON handling, retry, error wrapping / 모든 것 자동화: 자동 JSON 처리, 재시도, 에러 래핑
- Type Safety: Rich error types with context / 타입 안전성: 컨텍스트가 있는 풍부한 에러 타입
- Zero Configuration: Sensible defaults for 99% of use cases / 제로 설정: 99% 사용 사례에 대한 합리적인 기본값

**Latest Version / 최신 버전**: v1.10.004 (2025-10-16)

**Completed Phases / 완료된 단계**:
- ✅ Phase 1 (v1.10.001): Core HTTP client, RESTful methods, retry logic / 핵심 HTTP 클라이언트, RESTful 메서드, 재시도 로직
- ✅ Phase 2-4 (v1.10.002-003): Response helpers, file operations, URL/Form builders / 응답 헬퍼, 파일 작업, URL/Form 빌더
- ✅ Phase 5 (v1.10.004): Cookie management (in-memory and persistent) / 쿠키 관리 (메모리 내 및 지속성)

---

## [v1.9.x] - File Utilities Package / 파일 유틸리티 패키지

**Focus / 초점**: Extreme simplicity file and path utilities / 극도로 간단한 파일 및 경로 유틸리티

**Detailed Changes / 상세 변경사항**: See / 참조 [docs/CHANGELOG/CHANGELOG-v1.9.md](docs/CHANGELOG/CHANGELOG-v1.9.md)

### Highlights / 주요 사항
- **Complete fileutil package**: 20 lines → 1-2 lines of code / 완전한 fileutil 패키지: 20줄 → 1-2줄 코드
- **~91 functions across 12 categories**: Complete coverage of file/directory operations / 12개 카테고리에 걸쳐 약 91개 함수: 파일/디렉토리 작업의 완전한 커버리지
- **Cross-platform compatibility**: All path operations use filepath for OS-agnostic behavior / 크로스 플랫폼 호환성: 모든 경로 작업이 OS에 구애받지 않는 filepath 사용
- **Automatic directory creation**: All write operations auto-create parent directories / 자동 디렉토리 생성: 모든 쓰기 작업이 상위 디렉토리 자동 생성
- **Buffered I/O**: Default 32KB buffer for optimal performance / 버퍼링된 I/O: 최적의 성능을 위한 기본 32KB 버퍼
- **Atomic operations**: WriteAtomic for safe file updates (temp + rename) / 원자적 작업: 안전한 파일 업데이트를 위한 WriteAtomic (임시 + 이름 변경)
- **Progress callbacks**: Copy operations support progress tracking for large files / 진행 상황 콜백: 대용량 파일에 대한 진행 상황 추적을 지원하는 복사 작업
- **Multiple hash algorithms**: MD5, SHA1, SHA256, SHA512 / 여러 해시 알고리즘: MD5, SHA1, SHA256, SHA512
- **Path safety**: IsSafe function to prevent directory traversal attacks / 경로 안전성: 디렉토리 탐색 공격 방지를 위한 IsSafe 함수
- **JSON/YAML/CSV support**: Direct read/write for structured data formats / JSON/YAML/CSV 지원: 구조화된 데이터 형식의 직접 읽기/쓰기
- **Zero external dependencies**: Standard library only (except gopkg.in/yaml.v3) / 외부 의존성 없음: 표준 라이브러리만 (gopkg.in/yaml.v3 제외)
- **Comprehensive documentation**: Package README with quick start guide and function reference / 포괄적인 문서화: 빠른 시작 가이드 및 함수 참조를 포함한 패키지 README

**Key Design Principles / 주요 설계 원칙**:
- Extreme Simplicity: 20 lines → 1-2 lines / 극도의 간결함: 20줄 → 1-2줄
- Safety First: Automatic directory creation, atomic writes, path validation / 안전 우선: 자동 디렉토리 생성, 원자적 쓰기, 경로 검증
- Cross-Platform: OS-agnostic path operations using filepath / 크로스 플랫폼: filepath를 사용한 OS에 구애받지 않는 경로 작업
- Zero Configuration: No setup required / 제로 설정: 설정 불필요

**Latest Version / 최신 버전**: v1.9.018 (2025-10-15)

---

## [v1.8.x] - Map Utilities Package / 맵 유틸리티 패키지

**Focus / 초점**: Extreme simplicity map utilities with Go 1.18+ generics / Go 1.18+ 제네릭을 사용한 극도로 간단한 맵 유틸리티

**Detailed Changes / 상세 변경사항**: See / 참조 [docs/CHANGELOG/CHANGELOG-v1.8.md](docs/CHANGELOG/CHANGELOG-v1.8.md)

### Highlights / 주요 사항
- **Complete maputil package**: 20 lines → 1-2 lines of code / 완전한 maputil 패키지: 20줄 → 1-2줄 코드
- **99 functions across 14 categories**: Complete coverage of common map operations / 14개 카테고리에 걸쳐 99개 함수: 일반적인 맵 작업의 완전한 커버리지
- **Go 1.18+ generics**: Type-safe map operations with generic type parameters / Go 1.18+ 제네릭: 제네릭 타입 파라미터를 사용한 타입 안전 맵 작업
- **Functional programming**: Map, Filter, Reduce, GroupBy and more / 함수형 프로그래밍: Map, Filter, Reduce, GroupBy 등
- **Merge operations**: Merge, Union, Intersection, Difference / 병합 작업: 병합, 합집합, 교집합, 차집합
- **Nested map support**: GetNested, SetNested, HasNested, DeleteNested, SafeGet / 중첩 맵 지원
- **Default value management**: GetOrSet, SetDefault, Defaults / 기본값 관리
- **Statistics functions**: Median, Frequencies / 통계 함수
- **YAML conversion**: ToYAML, FromYAML (in addition to JSON) / YAML 변환: ToYAML, FromYAML (JSON에 추가)
- **Comparison functions**: Diff, Compare, EqualFunc / 비교 함수: Diff, Compare, EqualFunc
- **Immutable operations**: All functions return new maps (no mutation) / 불변 작업: 모든 함수는 새 맵을 반환 (변경 없음)
- **Zero dependencies**: Standard library only (except gopkg.in/yaml.v3) / 제로 의존성: 표준 라이브러리만 (gopkg.in/yaml.v3 제외)
- **Comprehensive tests**: 92.8% test coverage with 90+ sub-tests and 17+ benchmarks / 포괄적인 테스트: 90개 이상의 하위 테스트 및 17개 이상의 벤치마크로 92.8% 테스트 커버리지
- **Comprehensive documentation**: USER_MANUAL (2,207 lines), DEVELOPER_GUIDE (2,356 lines), complete examples (1,676 lines) / 포괄적인 문서화: 사용자 매뉴얼 (2,207줄), 개발자 가이드 (2,356줄), 완전한 예제 (1,676줄)

**Key Design Principles / 주요 설계 원칙**:
- Extreme Simplicity: 20 lines → 1-2 lines / 극도의 간결함: 20줄 → 1-2줄
- Type Safety: Generic type parameters for compile-time safety / 타입 안전: 컴파일 타임 안전을 위한 제네릭 타입 파라미터
- Functional Style: Inspired by JavaScript, Lodash, Python dict methods / 함수형 스타일: JavaScript, Lodash, Python dict 메서드에서 영감
- Zero Configuration: No setup required / 제로 설정: 설정 불필요
- Nested Map Support: Safe navigation of deeply nested structures / 중첩 맵 지원: 깊이 중첩된 구조의 안전한 탐색

**Latest Version / 최신 버전**: v1.8.017 (2025-10-15)

---

## [v1.7.x] - Slice Utilities Package / 슬라이스 유틸리티 패키지

**Focus / 초점**: Extreme simplicity slice utilities with Go 1.18+ generics / Go 1.18+ 제네릭을 사용한 극도로 간단한 슬라이스 유틸리티

**Detailed Changes / 상세 변경사항**: See / 참조 [docs/CHANGELOG/CHANGELOG-v1.7.md](docs/CHANGELOG/CHANGELOG-v1.7.md)

### Highlights / 주요 사항
- **Complete sliceutil package**: 20 lines → 1 line of code / 완전한 sliceutil 패키지: 20줄 → 1줄 코드
- **95 functions across 14 categories**: Complete coverage of common slice operations / 14개 카테고리에 걸쳐 95개 함수: 일반적인 슬라이스 작업의 완전한 커버리지
- **Go 1.18+ generics**: Type-safe slice operations with generic type parameters / Go 1.18+ 제네릭: 제네릭 타입 파라미터를 사용한 타입 안전 슬라이스 작업
- **Functional programming**: Map, Filter, Reduce, Scan, ZipWith and more / 함수형 프로그래밍: Map, Filter, Reduce, Scan, ZipWith 등
- **Statistics functions**: Median, Mode, StandardDeviation, Variance, Percentile / 통계 함수: 중앙값, 최빈값, 표준편차, 분산, 백분위수
- **Diff operations**: Compare slices, track changes, EqualUnordered / Diff 작업: 슬라이스 비교, 변경 추적, 순서 무관 비교
- **Immutable operations**: All functions return new slices (no mutation) / 불변 작업: 모든 함수는 새 슬라이스를 반환 (변경 없음)
- **Zero dependencies**: Standard library only (except golang.org/x/exp for constraints) / 제로 의존성: 표준 라이브러리만 (제약조건을 위한 golang.org/x/exp 제외)
- **100% test coverage**: 260+ test cases with comprehensive edge case coverage / 100% 테스트 커버리지: 포괄적인 엣지 케이스 커버를 갖춘 260개 이상 테스트 케이스
- **Comprehensive documentation**: USER_MANUAL (3,887 lines), DEVELOPER_GUIDE (2,205 lines), PERFORMANCE_BENCHMARKS / 포괄적인 문서화: 사용자 매뉴얼 (3,887줄), 개발자 가이드 (2,205줄), 성능 벤치마크

**Key Design Principles / 주요 설계 원칙**:
- Extreme Simplicity: 20 lines → 1 line / 극도의 간결함: 20줄 → 1줄
- Type Safety: Generic type parameters for compile-time safety / 타입 안전: 컴파일 타임 안전을 위한 제네릭 타입 파라미터
- Functional Style: Inspired by JavaScript, Python, Ruby array methods / 함수형 스타일: JavaScript, Python, Ruby 배열 메서드에서 영감
- Zero Configuration: No setup required / 제로 설정: 설정 불필요

**Latest Version / 최신 버전**: v1.7.023 (2025-10-15)

---

## [v1.6.x] - Time Utilities Package / 시간 유틸리티 패키지

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
