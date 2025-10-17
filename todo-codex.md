# Bilingual Comment Overhaul – Master Checklist / 작업 마스터 체크리스트

이 문서는 세션과 관계없이 동일한 순서와 방식으로 주석 보강 작업을 진행하기 위한 **단일 진실 소스**입니다.  
아래 절차와 체크리스트에 따라 작업하면 언제든지 이어서 진행할 수 있습니다.

---

## 📋 Table of Contents / 목차
1. [Global Workflow](#global-workflow--전체-작업-절차)
2. [Comment Standards](#comment-standards--주석-작성-표준)
3. [High Priority Packages](#high-priority-packages-핵심-패키지)
4. [Core Utility Packages](#core-utility-packages-중간-우선순위)
5. [Supporting Packages](#supporting-packages-보조-패키지)
6. [Database Packages](#database-packages-데이터베이스-패키지)
7. [Examples Directory](#examples-directory-예제-디렉터리)
8. [Test Files](#test-files-테스트-파일)
9. [Verification Steps](#verification-steps-검증-단계)
10. [Progress Tracking](#progress-tracking-진행-상황-추적)

---

## Global Workflow / 전체 작업 절차

### 작업 시작 전 (Before Starting)
1. **작업 대상 선정**: `todo-codex.md`에서 다음 작업 대상을 고르고 체크박스를 `[-]`로 변경
2. **기존 코드 분석**: 대상 파일의 현재 주석 상태, 함수/메서드 구조 파악
3. **관련 문서 확인**: README.md, 기존 문서, 테스트 코드 확인

### 주석 보강 작업 (Comment Enhancement)
1. **패키지 레벨 주석**: 패키지의 목적, 주요 기능, 사용 시나리오를 영문/한글로 작성
2. **타입/구조체 주석**: 각 필드의 목적, 제약 조건, 예상 값 범위를 병기
3. **함수/메서드 주석**: 다음 정보를 **영문 블록** 후 **한글 블록**으로 작성
   - **Purpose** (목적): 함수가 수행하는 작업의 핵심 설명
   - **Parameters** (매개변수): 각 파라미터의 타입, 의미, 제약 조건
   - **Returns** (반환값): 반환 타입별 의미와 성공/실패 조건
   - **Errors** (에러 케이스): 발생 가능한 에러와 그 조건
   - **Example** (예제, 선택): 간단한 사용 예제 (복잡한 경우)
   - **Notes** (주의사항, 선택): Thread safety, 성능 특성, 제한사항 등
4. **인라인 주석**: 복잡한 로직에만 `// English comment / 한글 주석` 형태로 병기
5. **상수/변수 주석**: 목적과 사용처를 간단히 병기

### 검증 및 문서화 (Verification & Documentation)
1. **테스트 실행**: `go test ./[package]` 또는 `go test ./...` 실행하여 동작 확인
2. **주석 품질 검증**: 
   - 모든 public 함수/타입에 주석 있는지 확인
   - 영문/한글 병기가 올바른지 확인
   - 내용이 코드와 일치하는지 확인
3. **문서 업데이트**:
   - `docs/BILINGUAL_AUDIT.md`: 완료된 파일 체크, 위험 항목 업데이트
   - `docs/CHANGELOG/CHANGELOG-specials.md`: 작업 요약 및 주요 변경사항 기록
4. **완료 표시**: 체크박스를 `[x]`로 변경하고 완료 날짜 기록

### 세션 종료 전 (Before Ending Session)
1. 현재 진행 중인 파일의 상태를 `todo-codex.md`에 명확히 기록 (`[-]` 상태 유지)
2. 다음 세션에서 시작할 지점을 "Next Steps" 섹션에 기록
3. 특이사항이나 발견된 이슈를 별도로 기록

---

## Comment Standards / 주석 작성 표준

### 1. 패키지 레벨 주석 형식
```go
// Package [name] provides functionality for [purpose].
// It offers [key features] with support for [capabilities].
//
// Main features include:
//   - Feature 1: Description
//   - Feature 2: Description
//   - Feature 3: Description
//
// Usage example:
//   [simple example code]
//
// [name] 패키지는 [목적]을 위한 기능을 제공합니다.
// [주요 기능]을 제공하며 [기능들]을 지원합니다.
//
// 주요 기능:
//   - 기능 1: 설명
//   - 기능 2: 설명
//   - 기능 3: 설명
//
// 사용 예:
//   [간단한 예제 코드]
package packagename
```

### 2. 함수/메서드 주석 형식
```go
// FunctionName performs [specific action] with [specific behavior].
// It [detailed explanation of what it does].
//
// Parameters:
//   - param1: Description of param1, expected values, constraints
//   - param2: Description of param2, expected values, constraints
//
// Returns:
//   - type1: Description of return value under success conditions
//   - error: Specific error conditions (ErrXXX when YYY)
//
// Errors:
//   - ErrInvalidInput: when param1 is invalid
//   - ErrNotFound: when resource doesn't exist
//
// Example:
//   result, err := FunctionName(param1, param2)
//   if err != nil {
//       // handle error
//   }
//
// Notes:
//   - Thread-safe / Not thread-safe
//   - Performance characteristics
//   - Any important limitations
//
// FunctionName은 [특정 동작]을 [특정 방식]으로 수행합니다.
// [상세한 동작 설명]을 합니다.
//
// 매개변수:
//   - param1: param1의 설명, 예상 값, 제약 조건
//   - param2: param2의 설명, 예상 값, 제약 조건
//
// 반환값:
//   - type1: 성공 조건에서의 반환 값 설명
//   - error: 특정 에러 조건 (YYY일 때 ErrXXX)
//
// 에러:
//   - ErrInvalidInput: param1이 유효하지 않을 때
//   - ErrNotFound: 리소스가 존재하지 않을 때
//
// 예제:
//   result, err := FunctionName(param1, param2)
//   if err != nil {
//       // 에러 처리
//   }
//
// 주의사항:
//   - Thread-safe / Thread-safe하지 않음
//   - 성능 특성
//   - 중요한 제한사항
func FunctionName(param1, param2 type) (type1, error) {
    // implementation
}
```

### 3. 타입/구조체 주석 형식
```go
// TypeName represents [what it represents].
// It is used for [purpose and use cases].
//
// TypeName은 [표현하는 것]을 나타냅니다.
// [목적과 사용 사례]에 사용됩니다.
type TypeName struct {
    // Field1 describes [purpose and constraints]
    // Field1은 [목적과 제약조건]을 나타냅니다
    Field1 string
    
    // Field2 indicates [purpose and constraints]
    // Field2는 [목적과 제약조건]을 나타냅니다
    Field2 int
}
```

### 4. 상수/변수 주석 형식
```go
// ConstantName defines [purpose].
// Used for [specific use case].
//
// ConstantName은 [목적]을 정의합니다.
// [특정 사용 사례]에 사용됩니다.
const ConstantName = "value"
```

### 5. 인라인 주석 규칙
- 복잡한 로직이나 비직관적인 코드에만 사용
- 한 줄 형식: `// English explanation / 한글 설명`
- 간결하고 명확하게 작성
- 코드 자체가 명확하면 인라인 주석 생략

### 6. 주석 품질 기준
✅ **좋은 주석**:
- 코드를 읽지 않아도 API 사용법을 이해할 수 있음
- 예외 상황과 에러 케이스가 명확히 설명됨
- 성능 특성이나 제약사항이 명시됨
- 초보자도 이해할 수 있을 정도로 상세함

❌ **피해야 할 주석**:
- 코드를 그대로 반복: `// Add x and y` for `result := x + y`
- 모호하거나 불완전한 설명
- 영문만 있거나 한글만 있는 경우
- 오래되어 코드와 일치하지 않는 주석

---

## High Priority Packages (핵심 패키지 – 순서대로 진행)

### 1. websvrutil 패키지
- [-] Overview  
      작업 내용: 모든 미들웨어·헬퍼·컨텍스트 파일을 정리, 테스트 포함
  - [ ] `websvrutil/middleware.go`
  - [ ] `websvrutil/app.go`
  - [ ] `websvrutil/router.go`
  - [ ] `websvrutil/options.go`
  - [ ] `websvrutil/context.go`
  - [ ] `websvrutil/context_bind.go`
  - [ ] `websvrutil/context_helpers.go`
  - [ ] `websvrutil/context_request.go`
  - [ ] `websvrutil/context_response.go`
  - [ ] `websvrutil/session.go`
  - [ ] `websvrutil/template.go`
  - [ ] `websvrutil/validator.go`
  - [ ] 관련 테스트 파일(`*_test.go`)

### 2. examples 디렉터리
- [ ] Overview  
      작업 내용: 예제 흐름 주석, 출력 예시, 학습 포인트를 병기
  - [ ] `examples/websvrutil/main.go`
  - [ ] `examples/logging/main.go`
  - [ ] `examples/mysql/main.go`
  - [ ] `examples/redis/main.go`
  - [ ] `examples/timeutil/main.go`
  - [ ] 기타 예제 서브 디렉터리

### 3. database/mysql 패키지
- [ ] Overview  
      작업 내용: 배치/마이그레이션/스키마/연결 관리 등 함수 주석 병기
  - [ ] `database/mysql/batch.go`
  - [ ] `database/mysql/migration.go`
  - [ ] `database/mysql/schema.go`
  - [ ] `database/mysql/options.go`
  - [ ] `database/mysql/client.go`
  - [ ] `database/mysql/metrics.go`
  - [ ] 기타 mysql 관련 파일 및 테스트

### 4. database/redis 패키지
- [ ] Overview  
      작업 내용: Redis 명령별 사용 시나리오와 오류 처리 전략 병기
  - [ ] `database/redis/client.go`
  - [ ] `database/redis/pipeline.go`
  - [ ] `database/redis/hash.go`
  - [ ] `database/redis/set.go`
  - [ ] `database/redis/string.go`
  - [ ] 기타 redis 관련 파일 및 테스트

---

## Core Utility Packages (중간 우선순위)
- [ ] `fileutil/*` – 파일 권한, 예외 처리, 플랫폼 주의 사항
- [ ] `maputil/*` – 시간 복잡도, 불변성 여부, 예제 추가
- [ ] `sliceutil/*` – 재할당 상황, 성능 팁, 에러 케이스
- [ ] `stringutil/*` – 국제화, 입력 검증, 예제 코드
- [ ] `timeutil/*` – 시간대/Locale/DST 설명 및 주의 사항
- [ ] `random/*` – 난수 특성, 시드 관리, 테스트 전략

---

## Supporting Packages (보조 패키지)
- [ ] `httputil/*` – 재시도 정책, Timeout/Context, 에러 처리
- [ ] `logging/*` – 로테이션, 색상 출력, 배너 활용법
- [ ] `errorutil/*` – 래핑 깊이, 코드/컨텍스트, stdlib 호환성

---

## Shared Tests & Helpers (공통 테스트 및 헬퍼)
- [ ] `*_test.go` 전체 – Given/When/Then 구분 등 단계별 병기
- [ ] 기타 헬퍼(`internal/*` 등 존재 시) – 목적, 입력, 출력 설명

---

## Verification & Documentation (검증 및 문서화)
- [ ] 주석 검사 스크립트 재실행 (영문-only / 한글-only 라인 체크)
- [ ] `docs/BILINGUAL_AUDIT.md` 최신화 (완료/잔여 항목 기록)
- [ ] `docs/CHANGELOG/CHANGELOG-specials.md` 업데이트 (작업 요약)
- [ ] `go test ./...` 또는 범위 테스트 실행 (로그/결과 기록)

---

## Commenting Guidelines / 주석 작성 규칙
- 함수/메서드 주석은 영문 블록 다음 한글 블록으로 구성합니다.  
  예)  
  `// AddUser adds a new user to the repository.`  
  `// AddUser는 새 사용자를 저장소에 추가합니다.`
- 블록 주석은 목적, 파라미터, 반환값, 에러, 예제 순으로 정리합니다.
- 인라인 주석은 `// do something / 무언가 수행`처럼 한 줄 병기를 허용합니다.
- 설명은 초보자도 이해할 수 있을 정도로 친절하고 자세하게 작성합니다.
- 주석 작성 후 반드시 테스트(`go test ./...`)로 기본 동작을 확인합니다.
- 모든 변경 사항은 Bilingual Audit과 Changelog에 기록합니다.

---
