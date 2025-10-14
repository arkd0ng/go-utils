# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 프로젝트 개요

`go-utils`는 Golang 개발을 위한 모듈화된 유틸리티 패키지 모음입니다. 라이브러리는 서브패키지 구조로 설계되어 사용자가 전체 라이브러리가 아닌 특정 유틸리티만 import할 수 있습니다.

**GitHub 저장소**: `github.com/arkd0ng/go-utils`
**현재 버전**: v1.4.005 (from cfg/app.yaml)
**Go 버전**: 1.24.6
**라이선스**: MIT

## 아키텍처

### 서브패키지 구조

라이브러리는 각 유틸리티 카테고리가 자체 디렉토리에 존재하는 서브패키지 아키텍처를 따릅니다:

```
go-utils/
├── random/              # 암호학적으로 안전한 랜덤 문자열 생성
│   ├── string.go       # 핵심 랜덤 문자열 생성 로직
│   ├── string_test.go  # 충돌 확률 테스트를 포함한 종합 테스트
│   └── README.md       # 패키지별 문서 (이중 언어: 영문/한글)
├── logging/             # 구조화된 로깅 및 파일 로테이션
│   ├── logger.go       # 핵심 로깅 로직
│   ├── options.go      # 함수형 옵션 패턴
│   └── README.md       # 패키지별 문서 (이중 언어: 영문/한글)
├── database/
│   ├── mysql/          # 극도로 간단한 MySQL/MariaDB 클라이언트
│   │   ├── client.go   # 클라이언트 핵심 로직
│   │   ├── simple.go   # 간단한 CRUD API (7개 메서드)
│   │   ├── builder.go  # Query Builder (Fluent API)
│   │   ├── transaction.go  # 트랜잭션 지원
│   │   ├── rotation.go # 자격 증명 순환 로직
│   │   └── README.md   # 패키지별 문서 (이중 언어: 영문/한글)
│   └── redis/          # 극도로 간단한 Redis 클라이언트
│       ├── client.go   # 클라이언트 핵심 로직
│       ├── string.go   # String 작업 (11개 메서드)
│       ├── hash.go     # Hash 작업 (10개 메서드)
│       ├── list.go     # List 작업 (9개 메서드)
│       ├── set.go      # Set 작업 (10개 메서드)
│       ├── zset.go     # Sorted Set 작업 (10개 메서드)
│       ├── key.go      # Key 작업 (11개 메서드)
│       ├── pipeline.go # Pipeline 지원
│       ├── transaction.go # Transaction 지원
│       ├── pubsub.go   # Pub/Sub 지원
│       └── README.md   # 패키지별 문서 (이중 언어: 영문/한글)
├── examples/
│   ├── random_string/  # 모든 14개 random 메서드 시연
│   ├── logging/        # 로깅 기능 및 설정 시연
│   ├── mysql/          # MySQL 패키지의 17개 예제 시연
│   └── redis/          # Redis 패키지의 8개 예제 시연
└── (향후 패키지)        # stringutil, sliceutil, maputil, fileutil 등
```

### 설계 원칙

1. **독립성**: 각 서브패키지는 교차 의존성 없이 자체 포함됩니다
2. **이중 언어 문서화**: 모든 문서, 주석, 테스트는 영문과 한글로 작성됩니다
3. **보안 우선**: 암호학적으로 안전한 랜덤 생성을 위해 `crypto/rand` 사용
4. **에러 처리**: 모든 메서드는 적절한 에러 처리를 위해 `(결과, error)`를 반환합니다
5. **가변 인자**: 메서드는 유연성을 위해 가변 `length` 인자를 받습니다:
   - 1개 인자: 고정 길이 (예: `Alnum(32)`는 정확히 32자 생성)
   - 2개 인자: 범위 (예: `Alnum(32, 128)`는 32-128자 생성)

### Random 패키지 아키텍처

`random` 패키지는 전역 싱글톤 패턴을 사용합니다:
- `stringGenerator` 구조체가 모든 생성 메서드를 제공
- `GenString`은 `random.GenString.메서드명()`으로 접근하는 전역 인스턴스
- 핵심 헬퍼 함수 `generateRandomString()`이 검증 및 생성 로직을 처리
- 문자 집합은 패키지 레벨 상수로 정의됨

**14개의 사용 가능한 메서드**:
- 기본: `Letters`, `Alnum`, `Digits`, `Complex`, `Standard`
- 대소문자 구분: `AlphaUpper`, `AlphaLower`, `AlnumUpper`, `AlnumLower`
- 16진수: `Hex`, `HexLower`
- 인코딩: `Base64`, `Base64URL`
- 사용자 정의: `Custom(charset string, length ...int)`

### Logging 패키지 아키텍처

`logging` 패키지는 구조화된 로깅과 파일 로테이션을 제공합니다:
- `Logger` 구조체가 모든 로깅 메서드 제공 (thread-safe with sync.Mutex)
- **Options Pattern**: 함수형 옵션으로 유연한 설정 (`WithFilePath`, `WithLevel` 등)
- **File Rotation**: lumberjack.v2 라이브러리 통합 (자동 크기 기반 로테이션)
- **Config Loading**: `cfg/app.yaml`에서 애플리케이션 정보 자동 로드
- **Multiple Writers**: 파일과 stdout에 동시 출력 지원

**주요 기능**:
- 로그 레벨: `DEBUG`, `INFO`, `WARN`, `ERROR`, `FATAL`
- 구조화 로깅: 키-값 쌍 지원 (`logger.Info("msg", "key", "value")`)
- 배너 출력: 애플리케이션 시작 시 정보 표시
- 색상 출력: stdout에 색상 적용 (파일에는 미적용)

### MySQL 패키지 아키텍처

`database/mysql` 패키지는 극도로 간단한 MySQL/MariaDB 클라이언트를 제공합니다:
- **30줄 → 2줄 코드 감소**: 보일러플레이트 코드 대폭 제거
- **자동 연결 관리**: 연결, 재시도, 재연결, 리소스 정리 모두 자동
- **무중단 자격 증명 순환**: 다중 연결 풀로 zero-downtime 자격 증명 교체
- **Options Pattern**: 함수형 옵션으로 유연한 설정
- **3가지 API 레벨**: Simple API, Query Builder, Raw SQL

**3가지 API 레벨**:

1. **Simple API (7개 메서드)**:
   - `SelectAll`, `SelectOne`, `Insert`, `Update`, `Delete`, `Count`, `Exists`
   - 가장 간단한 CRUD 작업을 1-2줄로 처리
   - Context 버전과 non-context 버전 제공 (예: `SelectAll`, `SelectAllContext`)
   - 추가 메서드: `SelectColumn`, `SelectColumns` (특정 컬럼만 선택)

2. **Query Builder (Fluent API)**:
   - 체이닝으로 복잡한 쿼리 구성 (JOIN, GROUP BY, HAVING 등)
   - 메서드: `Select().From().Where().OrderBy().Limit().All()`
   - 트랜잭션 내에서도 사용 가능

3. **SelectWhere API**:
   - Query Builder와 Simple API의 중간 레벨
   - 함수형 옵션으로 한 줄에 복잡한 쿼리 작성
   - 옵션: `WithColumns`, `WithOrderBy`, `WithLimit`, `WithGroupBy`, `WithHaving`, `WithDistinct`

4. **Raw SQL**:
   - `Query`, `Exec`, `QueryRow` 메서드로 직접 SQL 실행
   - 완전한 제어가 필요한 경우

**핵심 기능**:
- **자동 재시도**: 네트워크 오류 시 자동 재시도 (설정 가능)
- **헬스 체크**: 백그라운드에서 연결 상태 모니터링
- **트랜잭션**: `Transaction()` 메서드로 자동 commit/rollback
- **자격 증명 순환**: `WithCredentialRefresh()`로 동적 자격 증명 관리
- **자동 리소스 정리**: defer rows.Close() 불필요 (Simple API)

**고급 기능 (v1.3.010+)**:

1. **Batch Operations (batch.go)**:
   - `BatchInsert`: 한 쿼리로 여러 행 삽입
   - `BatchUpdate`: 여러 행을 다른 값으로 업데이트
   - `BatchDelete`: ID로 여러 행 삭제
   - `BatchSelectByIDs`: ID로 여러 행 선택

2. **Upsert Operations (upsert.go)**:
   - `Upsert`: INSERT ... ON DUPLICATE KEY UPDATE
   - `UpsertBatch`: 배치 Upsert
   - `Replace`: REPLACE INTO (DELETE + INSERT)

3. **Pagination (pagination.go)**:
   - `Paginate`: 간단한 페이지네이션 (메타데이터 포함)
   - `PaginateQuery`: 커스텀 쿼리 페이지네이션
   - `PaginationResult`: TotalRows, TotalPages, HasNext, HasPrev 제공

4. **Soft Delete (softdelete.go)**:
   - `SoftDelete`: deleted_at 컬럼으로 소프트 삭제
   - `Restore`: 소프트 삭제된 행 복원
   - `SelectAllWithTrashed`: 소프트 삭제된 행 포함
   - `SelectAllOnlyTrashed`: 소프트 삭제된 행만
   - `PermanentDelete`: 영구 삭제

5. **Query Statistics (stats.go)**:
   - `GetQueryStats`: 쿼리 통계 (TotalQueries, AverageDuration 등)
   - `ResetQueryStats`: 통계 초기화
   - `EnableSlowQueryLog`: 느린 쿼리 로깅 활성화
   - `GetSlowQueries`: 느린 쿼리 목록

6. **Pool Metrics (metrics.go)**:
   - `GetPoolMetrics`: 연결 풀 메트릭
   - `GetPoolHealthInfo`: 풀 상태 정보 (Healthy, Status, Message)
   - `GetConnectionUtilization`: 연결 사용률 (%)

7. **Schema Inspector (schema.go)**:
   - `GetTables`: 모든 테이블 목록
   - `GetColumns`: 테이블 컬럼 정보
   - `GetIndexes`: 테이블 인덱스 정보
   - `TableExists`: 테이블 존재 확인
   - `InspectTable`: 전체 테이블 정보

8. **Migration Helpers (migration.go)**:
   - `CreateTable`: 스키마로 테이블 생성
   - `DropTable`: 테이블 삭제
   - `TruncateTable`: 테이블 초기화
   - `AddColumn`: 컬럼 추가
   - `DropColumn`: 컬럼 삭제
   - `AddIndex`: 인덱스 추가
   - `AddForeignKey`: 외래 키 제약 조건 추가

9. **CSV Export/Import (export.go)**:
   - `ExportTableToCSV`: 전체 테이블을 CSV로 내보내기
   - `ImportFromCSV`: CSV를 테이블로 가져오기
   - `ExportQueryToCSV`: 쿼리 결과를 CSV로 내보내기
   - 컬럼 매핑 지원 (`WithColumnMapping`)

**파일 구조**:
- `client.go`: 클라이언트 핵심 로직 및 New() 생성자
- `simple.go`: 7개 Simple API 메서드 (SelectAll, Insert 등)
- `builder.go`: Query Builder (Fluent API)
- `select_options.go`: SelectWhere API 함수형 옵션
- `transaction.go`: 트랜잭션 지원
- `rotation.go`: 자격 증명 순환 로직
- `connection.go`: 연결 관리 및 헬스 체크
- `retry.go`: 재시도 로직
- `scan.go`: 행 스캔 및 타입 변환
- `errors.go`: 에러 타입 정의
- `config.go`: 설정 구조체
- `batch.go`: 배치 작업
- `upsert.go`: Upsert 작업
- `pagination.go`: 페이지네이션
- `softdelete.go`: 소프트 삭제
- `stats.go`: 쿼리 통계
- `metrics.go`: 풀 메트릭
- `schema.go`: 스키마 검사
- `migration.go`: 마이그레이션 헬퍼
- `export.go`: CSV 내보내기/가져오기

### Redis 패키지 아키텍처

`database/redis` 패키지는 극도로 간단한 Redis 클라이언트를 제공합니다:
- **20줄 → 2줄 코드 감소**: Redis 작업을 위한 보일러플레이트 코드 대폭 제거
- **자동 연결 관리**: 연결, 재시도, 재연결, 리소스 정리 모두 자동
- **Options Pattern**: 함수형 옵션으로 유연한 설정
- **60개 이상 메서드**: 6가지 데이터 타입 지원 + 고급 기능
- **타입 안전**: Generic 메서드로 타입 안전 작업 (`GetAs[T]`, `HGetAllAs[T]`)

**6가지 데이터 타입 작업**:

1. **String Operations (string.go - 11개 메서드)**:
   - 기본: `Set`, `Get`, `MGet`, `MSet`, `Append`, `GetRange`
   - 타입 안전: `GetAs[T]` (제네릭 메서드)
   - 숫자: `Incr`, `IncrBy`, `Decr`, `DecrBy`
   - 조건부: `SetNX` (Not Exists), `SetEX` (Expiration)

2. **Hash Operations (hash.go - 10개 메서드)**:
   - 기본: `HSet`, `HSetMap`, `HGet`, `HGetAll`, `HDel`
   - 타입 안전: `HGetAllAs[T]` (제네릭 메서드)
   - 유틸리티: `HExists`, `HLen`, `HKeys`, `HVals`
   - 숫자: `HIncrBy`, `HIncrByFloat`

3. **List Operations (list.go - 9개 메서드)**:
   - 추가: `LPush`, `RPush`
   - 제거: `LPop`, `RPop`, `LRem`, `LTrim`
   - 조회: `LRange`, `LIndex`, `LLen`
   - 수정: `LSet`

4. **Set Operations (set.go - 10개 메서드)**:
   - 기본: `SAdd`, `SRem`, `SMembers`, `SIsMember`, `SCard`
   - 집합 연산: `SUnion`, `SInter`, `SDiff`
   - 랜덤: `SPop`, `SRandMember`

5. **Sorted Set Operations (zset.go - 10개 메서드)**:
   - 기본: `ZAdd`, `ZAddMultiple`, `ZRem`, `ZCard`, `ZScore`
   - 범위 조회: `ZRange`, `ZRevRange`, `ZRangeByScore`
   - 순위: `ZRank`, `ZRevRank`
   - 증가: `ZIncrBy`

6. **Key Operations (key.go - 11개 메서드)**:
   - 기본: `Del`, `Exists`, `Type`, `Rename`, `RenameNX`
   - 만료: `Expire`, `ExpireAt`, `TTL`, `Persist`
   - 검색: `Keys`, `Scan`

**고급 기능**:

1. **Pipeline (pipeline.go)**:
   - `Pipeline()`: 배치 명령 실행
   - `TxPipeline()`: 트랜잭션 파이프라인
   - 네트워크 왕복 최소화

2. **Transaction (transaction.go)**:
   - `Transaction()`: WATCH/MULTI/EXEC 지원
   - 낙관적 잠금 (Optimistic Locking)
   - 자동 재시도 로직

3. **Pub/Sub (pubsub.go)**:
   - `Publish`: 메시지 발행
   - `Subscribe`: 채널 구독
   - `PSubscribe`: 패턴 구독
   - Channel 기반 메시지 수신

**핵심 기능**:
- **자동 재시도**: 네트워크 오류 시 exponential backoff 재시도
- **헬스 체크**: 백그라운드에서 연결 상태 모니터링
- **Connection Pooling**: go-redis 내장 연결 풀 활용
- **Context 지원**: 모든 메서드에서 context 지원 (취소 및 타임아웃)
- **타입 변환**: Generic 메서드로 자동 타입 변환
- **에러 처리**: 재시도 가능 에러 자동 감지 및 처리

**설계 철학** - "극도의 간결함 (Extreme Simplicity)":
- **Before**: 20줄 이상의 연결 관리, 에러 처리, 재시도 코드
  ```go
  // 20+ lines of code
  rdb := redis.NewClient(&redis.Options{
      Addr:     "localhost:6379",
      Password: "",
      DB:       0,
  })
  ctx := context.Background()
  err := rdb.Set(ctx, "key", "value", 0).Err()
  if err != nil {
      // retry logic
      // error handling
  }
  val, err := rdb.Get(ctx, "key").Result()
  if err != nil {
      // error handling
  }
  // More boilerplate...
  ```

- **After**: 2줄로 모든 작업 완료
  ```go
  client, _ := redis.New(redis.WithAddr("localhost:6379"))
  client.Set(context.Background(), "key", "value", 0)
  ```

**파일 구조**:
- `client.go`: 클라이언트 핵심 로직 및 New() 생성자
- `string.go`: String 작업 (11개 메서드)
- `hash.go`: Hash 작업 (10개 메서드)
- `list.go`: List 작업 (9개 메서드)
- `set.go`: Set 작업 (10개 메서드)
- `zset.go`: Sorted Set 작업 (10개 메서드)
- `key.go`: Key 작업 (11개 메서드)
- `pipeline.go`: Pipeline 지원
- `transaction.go`: Transaction 지원
- `pubsub.go`: Pub/Sub 지원
- `connection.go`: 연결 관리 및 헬스 체크
- `retry.go`: 재시도 로직 (exponential backoff)
- `errors.go`: 에러 타입 정의
- `config.go`: 설정 구조체
- `options.go`: 함수형 옵션 패턴
- `types.go`: 타입 정의

## 버전 관리 및 CHANGELOG 규칙

### 버전 관리

**버전 형식**: `vMAJOR.MINOR.PATCH` (예: v1.2.002)

**자동 패치 버전 증가 규칙**:
- 모든 작업 시작 전 패치 버전을 자동으로 1 증가시킵니다
- 예: v1.2.001 → v1.2.002 → v1.2.003
- 버전은 `cfg/app.yaml` 파일에서 관리합니다

**버전 히스토리**:
- **v1.0.x**: Random package 개발
- **v1.1.x**: Logging package 개발
- **v1.2.x**: Random 및 Logging 패키지 종합 문서화 (USER_MANUAL, DEVELOPER_GUIDE)
- **v1.3.x**: MySQL package 개발 및 문서화
- **v1.4.x** (현재): Redis package 개발 및 문서화

### 외부 의존성

프로젝트에서 사용하는 외부 라이브러리:
- **github.com/go-sql-driver/mysql**: MySQL 드라이버 (database/mysql 패키지)
- **github.com/redis/go-redis/v9**: Redis 클라이언트 (database/redis 패키지)
- **gopkg.in/natefinch/lumberjack.v2**: 파일 로테이션 (logging 패키지)
- **gopkg.in/yaml.v3**: YAML 설정 파일 파싱 (cfg/app.yaml, cfg/database.yaml, cfg/redis.yaml)

### CHANGELOG 관리

**파일 구조**:
```
go-utils/
├── CHANGELOG.md                     # Major/Minor 버전 개요만 포함
└── docs/
    └── CHANGELOG/
        ├── CHANGELOG-v1.0.md        # v1.0.x 상세 변경사항
        ├── CHANGELOG-v1.1.md        # v1.1.x 상세 변경사항
        ├── CHANGELOG-v1.2.md        # v1.2.x 상세 변경사항
        ├── CHANGELOG-v1.3.md        # v1.3.x 상세 변경사항
        └── CHANGELOG-v1.4.md        # v1.4.x 상세 변경사항
```

**CHANGELOG 규칙**:

1. **루트 CHANGELOG.md**:
   - Major/Minor 버전의 대략적인 내용만 언급
   - 각 버전별 상세 CHANGELOG 링크 제공 (예: `docs/CHANGELOG/CHANGELOG-v1.1.md`)

2. **버전별 CHANGELOG (docs/CHANGELOG/CHANGELOG-vX.Y.md)**:
   - 해당 Major.Minor 버전의 모든 패치 변경사항 포함
   - 각 패치별로 날짜, 변경 내용 기록 (이중 언어)
   - 최신 패치가 맨 위에 위치

3. **필수 업데이트 시점**:
   - 모든 패치 작업 후 GitHub 푸시 전
   - 모든 문서 작업 후 GitHub 푸시 전
   - **반드시 CHANGELOG를 업데이트한 후 커밋 및 푸시**

4. **CHANGELOG 형식**:
   ```markdown
   ## [v1.2.002] - 2025-10-10

   ### Added
   - 새로운 기능 추가 사항

   ### Changed
   - 변경된 기능

   ### Fixed
   - 버그 수정

   ### Removed
   - 제거된 기능
   ```

### Git 커밋 및 푸시 워크플로우

**작업 순서** (모든 작업에 적용):

1. **작업 시작**:
   - `cfg/app.yaml`의 패치 버전을 1 증가
   - 예: v1.2.001 → v1.2.002

2. **코드 작업 및 수정**

3. **컴파일 및 테스트**:
   ```bash
   go build ./...
   go test ./... -v
   ```

4. **문서 작업**:
   - README 업데이트
   - 필요시 예제 코드 업데이트

5. **CHANGELOG 업데이트**:
   - 해당 버전의 `docs/CHANGELOG-vX.Y.md` 업데이트
   - 변경사항을 명확하게 기록

6. **Git 커밋 및 푸시**:
   ```bash
   git add .
   git commit -m "타입: 간단한 설명"
   git push
   ```

**커밋 메시지 타입**:
- `Feat`: 새로운 기능
- `Fix`: 버그 수정
- `Docs`: 문서 변경
- `Refactor`: 리팩토링
- `Test`: 테스트 추가/수정
- `Chore`: 빌드, 설정 등

## 문서화 작업 워크플로우

### 패키지 문서화 표준 작업 순서

각 패키지에 대한 종합 문서를 작성할 때 다음 순서를 따릅니다:

**1. 버전 증가**:
```bash
# cfg/app.yaml의 패치 버전을 1 증가
# 예: v1.2.003 → v1.2.004
```

**2. 패키지 분석**:
- 패키지의 모든 코드 파일 검토 (`*.go`)
- 패키지 README.md 검토
- 테스트 파일 검토 (`*_test.go`)
- examples 디렉토리 검토

**3. 문서 디렉토리 생성**:
```bash
mkdir -p docs/{package_name}/
```

**4. 사용자 매뉴얼 작성** (`docs/{package}/USER_MANUAL.md`):

필수 섹션 (모두 이중 언어):
- **목차**: 모든 주요 섹션 링크
- **Introduction / 소개**: 패키지 개요, 주요 기능, 사용 사례
- **Installation / 설치**: 전제 조건, 패키지 설치, 임포트 방법
- **Quick Start / 빠른 시작**: 3-5개의 빠른 시작 예제
- **Configuration Reference / 설정 참조**: 모든 옵션, 메서드, 설정 테이블
- **Usage Patterns / 사용 패턴**: 일반적인 사용 패턴 5-10개
- **Common Use Cases / 일반적인 사용 사례**: 실제 사용 사례 5-10개 (전체 코드 포함)
- **Best Practices / 모범 사례**: 10-15개 모범 사례
- **Troubleshooting / 문제 해결**: 일반적인 문제 및 해결책
- **FAQ**: 10-15개 자주 묻는 질문

**5. 개발자 가이드 작성** (`docs/{package}/DEVELOPER_GUIDE.md`):

필수 섹션 (모두 이중 언어):
- **목차**: 모든 주요 섹션 링크
- **Architecture Overview / 아키텍처 개요**: 설계 원칙, 상위 수준 아키텍처 다이어그램
- **Package Structure / 패키지 구조**: 파일 구성, 파일별 책임
- **Core Components / 핵심 컴포넌트**: 주요 타입, 구조체, 인터페이스
- **Internal Implementation / 내부 구현**: 흐름 다이어그램, 상세 구현 설명
- **Design Patterns / 디자인 패턴**: 사용된 패턴 설명 (Singleton, Options 등)
- **Adding New Features / 새 기능 추가**: 단계별 가이드 및 예제
- **Testing Guide / 테스트 가이드**: 테스트 구조, 실행 방법, 작성 가이드
- **Performance / 성능**: 벤치마크, 최적화 기법
- **Contributing Guidelines / 기여 가이드라인**: 기여 프로세스, 체크리스트
- **Code Style / 코드 스타일**: 명명 규칙, 주석 스타일, 모범 사례

**6. 테스트 및 빌드**:
```bash
go build ./...
go test ./{package} -v
```

**7. CHANGELOG 업데이트**:
- `docs/CHANGELOG/CHANGELOG-v1.2.md`에 새 버전 항목 추가
- 생성된 문서 파일 나열
- 주요 섹션 요약

**8. Git 커밋 및 푸시**:
```bash
git add cfg/app.yaml docs/CHANGELOG/CHANGELOG-v1.2.md docs/{package}/ {package}/*_test.go
git commit -m "Docs: Add comprehensive {Package} package documentation (User Manual and Developer Guide)"
git push
```

### 문서 디렉토리 구조

```
go-utils/
├── docs/
│   ├── CHANGELOG/
│   │   ├── CHANGELOG-v1.0.md
│   │   ├── CHANGELOG-v1.1.md
│   │   ├── CHANGELOG-v1.2.md
│   │   ├── CHANGELOG-v1.3.md
│   │   └── CHANGELOG-v1.4.md
│   ├── random/
│   │   ├── USER_MANUAL.md      # ~600 lines, 완전한 사용자 가이드
│   │   └── DEVELOPER_GUIDE.md  # ~700 lines, 완전한 개발자 가이드
│   ├── logging/
│   │   ├── USER_MANUAL.md      # ~1000 lines, 완전한 사용자 가이드
│   │   └── DEVELOPER_GUIDE.md  # ~900 lines, 완전한 개발자 가이드
│   └── database/
│       ├── mysql/
│       │   ├── USER_MANUAL.md      # 완전한 사용자 가이드
│       │   ├── DEVELOPER_GUIDE.md  # 완전한 개발자 가이드
│       │   ├── DESIGN_PLAN.md      # 설계 계획 문서
│       │   └── WORK_PLAN.md        # 작업 계획 문서
│       └── redis/
│           ├── USER_MANUAL.md      # 완전한 사용자 가이드 (예정)
│           ├── DEVELOPER_GUIDE.md  # 완전한 개발자 가이드 (예정)
│           ├── DESIGN_PLAN.md      # 설계 계획 문서
│           └── WORK_PLAN.md        # 작업 계획 문서
```

### 문서 작성 지침

**이중 언어 형식**:
- 모든 제목: `## Section Title / 섹션 제목`
- 모든 설명: 영문 문장 다음에 한글 번역
- 코드 예제: 주석에 이중 언어 포함
- 테이블: 헤더와 내용 모두 이중 언어

**코드 예제**:
```go
// Create default logger / 기본 로거 생성
logger := logging.Default()
defer logger.Close()

// Log messages / 메시지 로깅
logger.Info("Application started")
```

**테이블 형식**:
```markdown
| Option / 옵션 | Type / 타입 | Default / 기본값 | Description / 설명 |
```

**일관성**:
- 모든 섹션 제목은 영문/한글 병기
- 모든 코드 주석은 이중 언어
- 파일 경로, 명령어, 코드는 원어 유지
- 기술 용어는 영문 후 한글 표기

## 개발 워크플로우

### 빌드 및 테스트

```bash
# 모든 테스트를 상세 출력으로 실행
go test ./... -v

# 특정 패키지에 대한 테스트 실행
go test ./random -v
go test ./logging -v
go test ./database/mysql -v
go test ./database/redis -v

# 단일 테스트 실행
go test ./random -v -run TestLetters
go test ./database/mysql -v -run TestSelectAll
go test ./database/redis -v -run TestStringOperations

# 벤치마크 실행
go test ./... -bench=.
go test ./random -bench=BenchmarkAlnum

# 커버리지와 함께 테스트 실행
go test ./... -cover
go test ./random -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### 예제 실행

```bash
# random string 예제 실행
go run examples/random_string/main.go

# logging 예제 실행
go run examples/logging/main.go

# MySQL 예제 실행 (MySQL 서버 필요)
go run examples/mysql/main.go

# Redis 예제 실행 (Redis 서버 필요)
go run examples/redis/main.go

# 예제 바이너리 빌드
go build -o bin/random_example examples/random_string/main.go
go build -o bin/logging_example examples/logging/main.go
go build -o bin/mysql_example examples/mysql/main.go
go build -o bin/redis_example examples/redis/main.go
```

**예제 디렉토리 구조**:
- `examples/random_string/main.go`: 모든 14개 random 메서드 시연
- `examples/logging/main.go`: 로깅 기능 및 설정 시연
- `examples/mysql/main.go`: MySQL 패키지의 17개 예제 시연
  - Simple API (SelectAll, Insert, Update, Delete, Count, Exists)
  - Query Builder (WHERE, JOIN, GROUP BY, HAVING, ORDER BY, LIMIT)
  - SelectWhere API (함수형 옵션)
  - SelectColumn, SelectColumns
  - Transaction
  - Raw SQL
- `examples/redis/main.go`: Redis 패키지의 8개 예제 시연
  - String Operations (Set, Get, GetAs, Incr, Decr)
  - Hash Operations (HSet, HGet, HGetAll, HGetAllAs)
  - List Operations (LPush, RPush, LPop, RPop, LRange)
  - Set Operations (SAdd, SMembers, SUnion, SInter)
  - Sorted Set Operations (ZAdd, ZRange, ZRangeByScore)
  - Key Operations (Del, Exists, Expire, TTL, Keys)
  - Pipeline (배치 명령 실행)
  - Pub/Sub (메시지 발행 및 구독)

**MySQL 예제 실행 요구사항**:
- MySQL 서버 실행 중이어야 함 (Docker: `./scripts/docker-mysql-start.sh`)
- `cfg/database.yaml` 설정 파일 필요
- 테스트 데이터베이스 및 테이블 (예제가 자동으로 설정)

**Redis 예제 실행 요구사항**:
- Redis 서버 실행 중이어야 함 (Docker: `./scripts/docker-redis-start.sh`)
- `cfg/redis.yaml` 설정 파일 필요
- 예제가 자동으로 데이터 생성 및 정리

### Docker 개발 워크플로우

**MySQL**:

```bash
# Docker MySQL 시작
./scripts/docker-mysql-start.sh

# Docker MySQL 중지
./scripts/docker-mysql-stop.sh

# Docker MySQL 로그 확인
./scripts/docker-mysql-logs.sh

# MySQL 클라이언트로 접속
docker exec -it go-utils-mysql mysql -u root -prootpassword

# Docker Compose로 직접 제어
docker compose up -d mysql
docker compose down mysql
```

**Redis**:

```bash
# Docker Redis 시작
./scripts/docker-redis-start.sh

# Docker Redis 중지
./scripts/docker-redis-stop.sh

# Docker Redis 로그 확인
./scripts/docker-redis-logs.sh

# Redis CLI로 접속
./scripts/docker-redis-cli.sh
# 또는
docker exec -it go-utils-redis redis-cli

# Docker Compose로 직접 제어
docker compose up -d redis
docker compose down redis
```

**MySQL 패키지 테스트**:

```bash
# MySQL 패키지만 테스트
go test ./database/mysql -v

# 특정 테스트 실행
go test ./database/mysql -v -run TestSelectAll
go test ./database/mysql -v -run TestTransaction

# 커버리지 확인
go test ./database/mysql -cover
go test ./database/mysql -coverprofile=coverage.out
go tool cover -html=coverage.out
```

**Redis 패키지 테스트**:

```bash
# Redis 패키지만 테스트
go test ./database/redis -v

# 특정 테스트 실행
go test ./database/redis -v -run TestStringOperations
go test ./database/redis -v -run TestPipeline

# 커버리지 확인
go test ./database/redis -cover
go test ./database/redis -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### 새로운 기능 추가

**`random` 패키지에 새로운 메서드를 추가할 때**:

1. 필요한 경우 `random/string.go` 상단에 문자 집합 상수 추가
2. `stringGenerator` 구조체에 메서드 생성
3. 포괄적인 이중 언어 문서 포함 (영문/한글)
4. 적절한 charset과 함께 `generateRandomString(charset, length...)` 호출
5. `random/string_test.go`에 해당 테스트 추가:
   - 기능 테스트 (길이, charset 정확성 검증)
   - 해당되는 경우 엣지 케이스 테스트
   - 벤치마크 테스트
6. `random/README.md`에 메서드 문서 업데이트
7. `examples/random_string/main.go`에 사용 예제 추가
8. 새 패키지나 주요 기능을 추가하는 경우 루트 `README.md` 업데이트

**`mysql` 패키지에 새로운 메서드를 추가할 때**:

1. 적절한 파일 선택:
   - Simple API → `simple.go`
   - Query Builder → `builder.go`
   - SelectWhere 옵션 → `select_options.go`
   - 트랜잭션 → `transaction.go`
2. Context 버전과 non-context 버전 모두 제공 (Simple API의 경우)
3. 이중 언어 문서 포함 (영문/한글)
4. 재시도 로직이 필요한 경우 `executeWithRetry()` 사용
5. `database/mysql/client_test.go`에 테스트 추가
6. `examples/mysql/main.go`에 예제 추가
7. `database/mysql/README.md` 업데이트

**`redis` 패키지에 새로운 메서드를 추가할 때**:

1. 적절한 파일 선택:
   - String 작업 → `string.go`
   - Hash 작업 → `hash.go`
   - List 작업 → `list.go`
   - Set 작업 → `set.go`
   - Sorted Set 작업 → `zset.go`
   - Key 작업 → `key.go`
   - Pipeline → `pipeline.go`
   - Transaction → `transaction.go`
   - Pub/Sub → `pubsub.go`
2. Context 지원 필수 (모든 메서드에서)
3. 이중 언어 문서 포함 (영문/한글)
4. 재시도 로직이 필요한 경우 `executeWithRetry()` 사용
5. `database/redis/client_test.go`에 테스트 추가
6. `examples/redis/main.go`에 예제 추가
7. `database/redis/README.md` 업데이트
8. Generic 타입이 필요한 경우 `GetAs[T]`, `HGetAllAs[T]` 패턴 사용

### 새로운 유틸리티 패키지 생성

새로운 유틸리티 패키지(예: `stringutil`, `sliceutil`)를 추가할 때:

1. 새 디렉토리 생성: `mkdir packagename`
2. 이중 언어 주석과 함께 패키지 파일 생성
3. 종합 테스트 파일 생성: `packagename_test.go`
4. 패키지 README 생성: `packagename/README.md` (이중 언어)
5. `examples/packagename/main.go`에 예제 추가
6. 새 패키지를 반영하도록 루트 `README.md` 업데이트
7. 패키지가 교차 의존성 없이 자체 포함되도록 보장

## 테스트 요구사항

모든 테스트는 다음을 포함해야 합니다:

1. **기능 테스트**: 예상 동작 및 출력 특성 검증
2. **엣지 케이스 테스트**: 음수 값, 잘못된 입력, 경계 조건
3. **랜덤성 테스트**: 고유성 및 적절한 분포 확인
4. **충돌 확률 테스트**: 랜덤 생성의 경우, 이론적 대 실제 충돌률 계산 및 검증
5. **벤치마크 테스트**: 모든 공개 메서드에 대한 성능 벤치마크
6. **이중 언어 주석**: 영문과 한글 설명 모두 포함

## 문서화 표준

### 코드 주석

모든 코드 주석은 이중 언어여야 합니다 (영문 먼저, 한글 두 번째):

```go
// Letters generates a random string containing only alphabetic characters (a-z, A-Z)
// Letters는 알파벳 문자(a-z, A-Z)만 포함하는 랜덤 문자열을 생성합니다
```

### README 파일

모든 README 파일은 병렬 구조로 이중 언어여야 합니다:
- 영문 문장 다음에 한글 번역
- 기술 용어는 두 언어로 표시
- 이중 언어 주석이 있는 코드 예제

## Import 경로

사용자는 특정 패키지를 import합니다:

```go
import "github.com/arkd0ng/go-utils/random"
```

루트 패키지를 import하지 않습니다:

```go
import "github.com/arkd0ng/go-utils"  // ❌ 이렇게 하지 마세요
```

## 에러 처리 패턴

모든 메서드는 `(결과, error)`를 반환합니다:

```go
str, err := random.GenString.Alnum(32)
if err != nil {
    log.Fatal(err)
}
```

에러 메시지는 영문으로만 작성됩니다 (이중 언어 아님).

## 버전 히스토리 컨텍스트

- **v0.1.0**: 루트 레벨 `GenRandomString`으로 첫 릴리스
- **v0.2.0**: Breaking change - 서브패키지 구조로 리팩토링
  - `GenRandomString`에서 `random.GenString`으로 변경
  - 9개의 새로운 메서드 추가 (총 14개)
  - 가변 인자 및 에러 처리 추가
  - 이중 언어 문서 추가
- **v1.0.x**: Random package 안정화 및 테스트 강화
- **v1.1.x**: Logging package 추가 (파일 로테이션, 구조화 로깅)
- **v1.2.x**: Random 및 Logging 패키지 종합 문서화 (USER_MANUAL, DEVELOPER_GUIDE)
- **v1.3.x**: MySQL package 추가
  - 극도로 간단한 MySQL/MariaDB 클라이언트 (30줄 → 2줄)
  - 3가지 API 레벨: Simple API, Query Builder, SelectWhere API
  - 무중단 자격 증명 순환 기능
  - 자동 연결 관리, 재시도, 리소스 정리
  - 종합 문서화 (USER_MANUAL, DEVELOPER_GUIDE)
- **v1.4.x** (현재): Redis package 추가
  - 극도로 간단한 Redis 클라이언트 (20줄 → 2줄)
  - 60개 이상 메서드: String, Hash, List, Set, Sorted Set, Key 작업
  - 고급 기능: Pipeline, Transaction, Pub/Sub
  - 타입 안전: Generic 메서드 (`GetAs[T]`, `HGetAllAs[T]`)
  - 자동 재시도, 연결 풀링, 헬스 체크
  - Docker 기반 개발 환경 (MySQL, Redis 통합)
  - 설계 문서 (DESIGN_PLAN, WORK_PLAN)

## 향후 로드맵

계획된 유틸리티 패키지 (README에 참조됨):
- `stringutil` - 문자열 처리 유틸리티
- `sliceutil` - 슬라이스/배열 헬퍼
- `maputil` - 맵 유틸리티
- `fileutil` - 파일/경로 유틸리티
- `httputil` - HTTP 헬퍼
- `timeutil` - 시간/날짜 유틸리티
- `validation` - 검증 유틸리티
- `errorutil` - 에러 처리 헬퍼
