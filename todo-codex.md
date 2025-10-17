# Bilingual Comment Overhaul – Master Checklist / 작업 마스터 체크리스트

이 문서는 세션과 관계없이 동일한 순서와 방식으로 주석 보강 작업을 진행하기 위한 **단일 진실 소스**입니다.  
아래 절차와 체크리스트에 따라 작업하면 언제든지 이어서 진행할 수 있습니다.

---

## 📋 Table of Contents / 목차
1. [Global Workflow](#global-workflow--전체-작업-절차)
2. [Comment Quality Standards](#comment-quality-standards--주석-품질-기준)
3. [Comment Writing Standards](#comment-writing-standards--주석-작성-표준)
4. [Complete File Checklist](#complete-file-checklist--전체-파일-체크리스트)
   - [websvrutil Package](#1-websvrutil-package)
   - [sliceutil Package](#2-sliceutil-package)
   - [maputil Package](#3-maputil-package)
   - [stringutil Package](#4-stringutil-package)
   - [timeutil Package](#5-timeutil-package)
   - [fileutil Package](#6-fileutil-package)
   - [httputil Package](#7-httputil-package)
   - [logging Package](#8-logging-package)
   - [errorutil Package](#9-errorutil-package)
   - [random Package](#10-random-package)
   - [database/mysql Package](#11-databasemysql-package)
   - [database/redis Package](#12-databaseredis-package)
   - [validation Package](#13-validation-package)
   - [examples Directory](#14-examples-directory)
5. [Verification Steps](#verification-steps-검증-단계)
6. [Progress Tracking](#progress-tracking-진행-상황-추적)

---

## Global Workflow / 전체 작업 절차

### ⚠️ 핵심 원칙 (Core Principles)
**모든 파일은 동등하게 중요합니다. 누락 없이 완료하는 것이 최우선 목표입니다.**

**주석 작성의 철학**:
- 📖 **충분히 자세하게**: 코드를 보지 않아도 동작을 완전히 이해할 수 있어야 함
- 👨‍🎓 **매우 친절하게**: Go 언어 초보자도 쉽게 이해할 수 있어야 함
- 🔍 **포괄적으로**: 엣지 케이스, 에러 상황, 성능 특성 모두 설명
- 💡 **실용적으로**: 실제 사용 예시와 주의사항 포함
- 🌐 **이중 언어**: 영문과 한글 모두 동일한 수준의 상세함 유지

### 작업 시작 전 (Before Starting)
1. **작업 대상 선정**: `todo-codex.md`에서 다음 작업 대상을 고르고 체크박스를 `[-]`로 변경
2. **기존 코드 분석**: 대상 파일의 현재 주석 상태, 함수/메서드 구조 파악
3. **관련 문서 확인**: README.md, 기존 문서, 테스트 코드 확인
4. **파일 목적 이해**: 파일이 해결하려는 문제와 제공하는 가치 파악

### 주석 보강 작업 (Comment Enhancement)

#### 1. 패키지 레벨 주석 (Package-Level Comments)
**작성 원칙**: 패키지의 존재 이유와 전체적인 그림을 그릴 수 있어야 함

포함할 내용:
- 패키지가 해결하는 문제
- 주요 기능과 제공하는 타입/함수 개요
- 일반적인 사용 시나리오 (최소 2-3개)
- 다른 패키지와의 관계
- 간단한 사용 예제 (Getting Started)
- 특별한 주의사항이나 제한사항

#### 2. 타입/구조체 주석 (Type/Struct Comments)
**작성 원칙**: 타입의 목적과 올바른 사용 방법을 명확히 전달

포함할 내용:
- 타입이 표현하는 개념
- 각 필드의 의미와 목적
- 필드의 제약 조건 (nil 가능 여부, 범위, 포맷 등)
- 타입의 불변성(immutability) 여부
- 동시성 안전성 (thread-safe 여부)
- 타입 생성 방법 (생성자 함수 안내)

#### 3. 함수/메서드 주석 (Function/Method Comments)
**작성 원칙**: 함수를 사용하는 개발자가 알아야 할 모든 것을 제공

**필수 항목** (모든 함수):
- **Purpose** (목적): 
  - 함수가 수행하는 작업을 명확하게 설명
  - WHY를 포함: 왜 이 함수가 필요한가?
  - WHAT을 포함: 무엇을 하는가?
  
- **Parameters** (매개변수):
  - 각 파라미터의 의미와 역할
  - 예상되는 값의 범위나 형식
  - nil 값 허용 여부
  - 특수한 값의 의미 (예: 0이면 무제한, -1이면 기본값)
  
- **Returns** (반환값):
  - 각 반환값의 의미
  - 성공/실패 시나리오별 반환값
  - nil 반환 조건
  
- **Errors** (에러):
  - 발생 가능한 모든 에러 타입
  - 각 에러가 발생하는 구체적인 조건
  - 에러 처리 권장 방법

**선택 항목** (복잡도에 따라):
- **Behavior** (동작 설명):
  - 함수의 내부 동작 흐름
  - 중요한 알고리즘 설명
  - 특수한 처리 로직
  
- **Example** (예제):
  - 기본 사용 예제
  - 일반적인 사용 패턴
  - 엣지 케이스 처리 예제
  
- **Performance** (성능):
  - 시간 복잡도 (Big-O)
  - 공간 복잡도
  - 성능 최적화 팁
  
- **Notes** (주의사항):
  - Thread safety
  - Goroutine 안전성
  - 메모리 할당 여부
  - 플랫폼별 차이점
  - 알려진 제한사항
  - 사용 시 주의할 점
  
- **See Also** (관련 항목):
  - 관련된 다른 함수
  - 대체 가능한 함수
  - 함께 사용하면 좋은 함수

#### 4. 상수/변수 주석 (Constant/Variable Comments)
**작성 원칙**: 값의 의미와 사용 목적을 명확히 전달

포함할 내용:
- 상수/변수의 목적
- 값의 의미와 단위
- 언제 사용하는지
- 변경하면 안 되는 이유 (상수의 경우)

#### 5. 인라인 주석 (Inline Comments)
**작성 원칙**: 코드만으로 이해하기 어려운 부분에만 사용

사용 시기:
- 복잡한 알고리즘이나 로직
- 비직관적인 코드
- 임시 해결책 (workaround)
- 성능 최적화를 위한 특수 처리
- 플랫폼별 분기 처리

형식: `// English explanation / 한글 설명`

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

## Comment Quality Standards / 주석 품질 기준

### 📊 주석 완성도 체크리스트
각 파일 작업 완료 시 다음을 확인하세요:

- [ ] **완전성 (Completeness)**: 모든 public 함수/타입/상수에 주석이 있는가?
- [ ] **상세성 (Detail)**: 초보자가 읽고 완전히 이해할 수 있는가?
- [ ] **정확성 (Accuracy)**: 주석이 실제 코드 동작과 일치하는가?
- [ ] **이중언어 (Bilingual)**: 영문과 한글 주석이 모두 동일한 수준으로 상세한가?
- [ ] **예제 (Examples)**: 복잡한 함수에 사용 예제가 있는가?
- [ ] **에러처리 (Error Handling)**: 모든 에러 케이스가 문서화되었는가?
- [ ] **성능 (Performance)**: 성능 특성이 명시되었는가? (필요한 경우)
- [ ] **안전성 (Safety)**: Thread-safety, 동시성 관련 주의사항이 있는가? (필요한 경우)

### ✅ 좋은 주석의 특징

1. **자기 완결적 (Self-Contained)**
   ```go
   // Good: 주석만 읽어도 함수를 사용할 수 있음
   // ParseDuration converts a duration string to time.Duration.
   // It supports formats like "1h", "30m", "45s", "1h30m45s".
   // Returns error if the format is invalid or the value is negative.
   //
   // ParseDuration은 기간 문자열을 time.Duration으로 변환합니다.
   // "1h", "30m", "45s", "1h30m45s"와 같은 형식을 지원합니다.
   // 형식이 잘못되었거나 값이 음수인 경우 에러를 반환합니다.
   ```

2. **구체적 (Specific)**
   ```go
   // Good: 구체적인 값과 조건 명시
   // MaxRetries defines the maximum number of retry attempts (1-10).
   // If set to 0, no retries will be performed.
   // Values greater than 10 will be capped at 10.
   //
   // MaxRetries는 최대 재시도 횟수를 정의합니다 (1-10).
   // 0으로 설정하면 재시도를 수행하지 않습니다.
   // 10보다 큰 값은 10으로 제한됩니다.
   ```

3. **실용적 (Practical)**
   ```go
   // Good: 사용 예제와 주의사항 포함
   // Connect establishes a connection to the database.
   //
   // Example:
   //   db, err := Connect("localhost:3306", opts)
   //   if err != nil {
   //       log.Fatal(err)
   //   }
   //   defer db.Close()
   //
   // Note: Always call Close() when done to prevent connection leaks.
   //
   // Connect는 데이터베이스에 연결을 설정합니다.
   //
   // 예제:
   //   db, err := Connect("localhost:3306", opts)
   //   if err != nil {
   //       log.Fatal(err)
   //   }
   //   defer db.Close()
   //
   // 주의: 연결 누수를 방지하기 위해 사용 후 반드시 Close()를 호출하세요.
   ```

### ❌ 피해야 할 주석

1. **불충분한 주석 (Insufficient)**
   ```go
   // Bad: 너무 간략함
   // Add adds two numbers.
   // Add는 두 숫자를 더합니다.
   func Add(a, b int) int
   
   // Good: 충분히 상세함
   // Add returns the sum of two integers.
   // It performs standard integer addition without overflow checking.
   // For large numbers that might overflow, consider using math/big package.
   //
   // Add는 두 정수의 합을 반환합니다.
   // 오버플로우 검사 없이 표준 정수 덧셈을 수행합니다.
   // 오버플로우가 발생할 수 있는 큰 숫자의 경우 math/big 패키지 사용을 고려하세요.
   ```

2. **모호한 주석 (Vague)**
   ```go
   // Bad: 모호한 설명
   // Process processes data.
   // Process는 데이터를 처리합니다.
   
   // Good: 명확한 설명
   // Process validates, transforms, and stores the input data.
   // It returns the processed data ID and any validation errors.
   //
   // Process는 입력 데이터를 검증, 변환 및 저장합니다.
   // 처리된 데이터의 ID와 발생한 검증 에러를 반환합니다.
   ```

3. **불완전한 이중언어 (Incomplete Bilingual)**
   ```go
   // Bad: 한쪽 언어만 상세함
   // ParseConfig reads and parses a YAML configuration file.
   // It supports environment variable expansion using ${VAR} syntax.
   // Returns ErrInvalidFormat if the YAML is malformed.
   // Returns ErrFileNotFound if the file doesn't exist.
   //
   // ParseConfig는 설정 파일을 파싱합니다.
   
   // Good: 양쪽 언어 모두 상세함
   // ParseConfig reads and parses a YAML configuration file.
   // It supports environment variable expansion using ${VAR} syntax.
   // Returns ErrInvalidFormat if the YAML is malformed.
   // Returns ErrFileNotFound if the file doesn't exist.
   //
   // ParseConfig는 YAML 설정 파일을 읽고 파싱합니다.
   // ${VAR} 문법을 사용한 환경 변수 확장을 지원합니다.
   // YAML 형식이 잘못된 경우 ErrInvalidFormat을 반환합니다.
   // 파일이 존재하지 않는 경우 ErrFileNotFound를 반환합니다.
   ```

---

## Comment Writing Standards / 주석 작성 표준

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

**기본 형식** (모든 함수에 적용):
```go
// FunctionName performs [specific action] with [specific behavior].
// It [detailed explanation of what it does and why it's needed].
// 
// This function is useful when [use case 1], [use case 2], etc.
// [Additional context about design decisions or implementation details]
//
// Parameters:
//   - param1: [Detailed description of param1]
//     * Expected values: [range, format, or specific values]
//     * Constraints: [any limitations or requirements]
//     * Special values: [e.g., nil means default, 0 means unlimited]
//   - param2: [Detailed description of param2]
//     * Expected values: [range, format, or specific values]
//     * Constraints: [any limitations or requirements]
//
// Returns:
//   - type1: [Description of return value under success conditions]
//     * Possible values: [what values can be returned]
//     * nil case: [when nil is returned]
//   - error: [General description of error cases]
//     * nil: indicates success
//     * non-nil: indicates failure (see Errors section)
//
// Errors:
//   - ErrInvalidInput: when param1 is invalid or out of range
//     * Example: param1 < 0 or param1 > 100
//   - ErrNotFound: when the requested resource doesn't exist
//   - ErrTimeout: when operation exceeds the timeout duration
//   - [any other possible errors]
//
// Example:
//   // Basic usage / 기본 사용법
//   result, err := FunctionName(10, "test")
//   if err != nil {
//       log.Printf("error: %v", err)
//       return
//   }
//   fmt.Printf("result: %v\n", result)
//
//   // Advanced usage / 고급 사용법
//   result, err := FunctionName(0, "") // uses defaults
//
// Performance:
//   - Time complexity: O(n) where n is [description]
//   - Space complexity: O(1) / O(n)
//   - [Any performance considerations]
//
// Notes:
//   - Thread-safe: [Yes/No] - [explanation]
//   - Goroutine-safe: [Yes/No] - [explanation]
//   - Memory allocation: [describe allocation behavior]
//   - Platform differences: [any OS-specific behavior]
//   - Known limitations: [any known issues or constraints]
//
// See Also:
//   - RelatedFunction: [how it relates]
//   - AlternativeFunction: [when to use instead]
//
// FunctionName은 [특정 동작]을 [특정 방식]으로 수행합니다.
// [함수가 수행하는 작업과 필요한 이유에 대한 상세한 설명]을 합니다.
//
// 이 함수는 [사용 사례 1], [사용 사례 2] 등에서 유용합니다.
// [설계 결정이나 구현 세부사항에 대한 추가 컨텍스트]
//
// 매개변수:
//   - param1: [param1에 대한 상세한 설명]
//     * 예상 값: [범위, 형식 또는 특정 값]
//     * 제약 조건: [제한사항 또는 요구사항]
//     * 특수 값: [예: nil은 기본값, 0은 무제한]
//   - param2: [param2에 대한 상세한 설명]
//     * 예상 값: [범위, 형식 또는 특정 값]
//     * 제약 조건: [제한사항 또는 요구사항]
//
// 반환값:
//   - type1: [성공 조건에서의 반환 값에 대한 설명]
//     * 가능한 값: [반환될 수 있는 값]
//     * nil 케이스: [nil이 반환되는 경우]
//   - error: [에러 케이스에 대한 일반적인 설명]
//     * nil: 성공을 나타냄
//     * non-nil: 실패를 나타냄 (에러 섹션 참조)
//
// 에러:
//   - ErrInvalidInput: param1이 유효하지 않거나 범위를 벗어날 때
//     * 예: param1 < 0 또는 param1 > 100
//   - ErrNotFound: 요청한 리소스가 존재하지 않을 때
//   - ErrTimeout: 작업이 타임아웃 시간을 초과할 때
//   - [발생 가능한 기타 에러]
//
### 6. 주석 길이 및 상세도 가이드

**목표**: 주석만 읽고 코드를 완전히 이해하고 사용할 수 있어야 합니다.

#### 추천 주석 길이:

1. **간단한 함수** (1-5줄 코드):
   - 최소 10-15줄의 주석
   - 영문 5-7줄 + 한글 5-7줄

2. **중간 복잡도 함수** (5-20줄 코드):
   - 최소 20-30줄의 주석
   - 영문 10-15줄 + 한글 10-15줄

3. **복잡한 함수** (20줄 이상):
   - 최소 30-50줄의 주석
   - 영문 15-25줄 + 한글 15-25줄
   - 예제 코드 포함 필수

4. **패키지 레벨 주석**:
   - 최소 30-50줄
   - 개요, 기능 목록, 사용 예제, 주의사항 모두 포함

#### 주석이 충분히 상세한지 확인하는 질문:

✅ **작성 후 자가 점검 질문**:
- [ ] 이 코드를 처음 보는 Go 초보자가 이해할 수 있는가?
- [ ] 모든 파라미터의 제약 조건이 명시되어 있는가?
- [ ] 가능한 모든 에러 케이스가 설명되어 있는가?
- [ ] 특수한 입력 값(nil, 0, 빈 문자열)의 동작이 설명되어 있는가?
- [ ] 사용 예제가 있는가? (복잡한 경우)
- [ ] 성능 특성이 설명되어 있는가? (필요한 경우)
- [ ] Thread-safety가 명시되어 있는가? (필요한 경우)
- [ ] 영문과 한글 주석의 상세도가 동일한가?
- [ ] 주석의 길이가 코드보다 짧지 않은가?

💡 **경험 법칙**: "주석이 너무 길다고 생각되면, 그제야 적당합니다!"
// 성능:
//   - 시간 복잡도: O(n) (n은 [설명])
//   - 공간 복잡도: O(1) / O(n)
//   - [성능 고려사항]
//
// 주의사항:
//   - Thread-safe: [예/아니오] - [설명]
//   - Goroutine-safe: [예/아니오] - [설명]
//   - 메모리 할당: [할당 동작 설명]
//   - 플랫폼 차이: [OS별 특정 동작]
//   - 알려진 제한사항: [알려진 이슈나 제약]
//
// 참고:
//   - RelatedFunction: [관계 설명]
//   - AlternativeFunction: [대신 사용할 경우]
func FunctionName(param1, param2 type) (type1, error) {
    // implementation
}
```

**간단한 함수** (에러 없고 복잡도 낮음):
```go
// SimpleName returns [what it returns].
// It [what it does in detail].
//
// Parameters:
//   - param: [description with constraints]
//
// Returns:
//   - [return type and meaning]
//
// SimpleName은 [반환하는 것]을 반환합니다.
// [상세한 동작 설명]을 수행합니다.
//
// 매개변수:
//   - param: [제약조건을 포함한 설명]
//
// 반환값:
//   - [반환 타입과 의미]
func SimpleName(param type) returnType {
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

## High Priority Packages (핵심 패키지)
**우선순위**: ⭐⭐⭐⭐⭐ (가장 높음)  
**작업 순서**: 순서대로 진행 권장

### 1. websvrutil 패키지 (Web Server Utilities)
**패키지 설명**: 웹 서버 애플리케이션 개발을 위한 핵심 유틸리티  
**작업 중요도**: Critical - 사용자가 가장 많이 사용하는 패키지

#### Core Files (핵심 파일)
- [ ] `websvrutil/websvrutil.go` - 패키지 메인 파일 및 기본 정의
- [ ] `websvrutil/app.go` - 애플리케이션 구조체 및 생명주기 관리
- [ ] `websvrutil/router.go` - 라우팅 로직 및 경로 매칭
- [ ] `websvrutil/options.go` - 설정 옵션 및 빌더 패턴
- [ ] `websvrutil/constants.go` - 상수 정의

#### Context Management (컨텍스트 관리)
- [ ] `websvrutil/context.go` - 기본 컨텍스트 구조체 및 메서드
- [ ] `websvrutil/context_request.go` - 요청 처리 관련 메서드
- [ ] `websvrutil/context_response.go` - 응답 처리 관련 메서드
- [ ] `websvrutil/context_bind.go` - 데이터 바인딩 기능
- [ ] `websvrutil/context_helpers.go` - 컨텍스트 헬퍼 함수들

#### Middleware & Security (미들웨어 및 보안)
- [ ] `websvrutil/middleware.go` - 미들웨어 체인 및 핸들러
- [ ] `websvrutil/csrf.go` - CSRF 토큰 생성 및 검증
- [ ] `websvrutil/session.go` - 세션 관리 및 저장소

#### Additional Features (추가 기능)
- [ ] `websvrutil/group.go` - 라우트 그룹화
- [ ] `websvrutil/bind.go` - 요청 데이터 바인딩
- [ ] `websvrutil/template.go` - 템플릿 엔진 통합
- [ ] `websvrutil/validator.go` - 입력 검증

#### Test Files (테스트 파일)
- [ ] `websvrutil/app_test.go`
- [ ] `websvrutil/router_test.go`
- [ ] `websvrutil/context_test.go`
- [ ] `websvrutil/middleware_test.go`
- [ ] `websvrutil/bind_test.go`
- [ ] `websvrutil/session_test.go`
- [ ] `websvrutil/csrf_test.go`
- [ ] `websvrutil/template_test.go`
- [ ] `websvrutil/validator_test.go`
- [ ] `websvrutil/group_test.go`
- [ ] `websvrutil/options_test.go`
- [ ] `websvrutil/error_test.go`
- [ ] `websvrutil/method_test.go`
- [ ] `websvrutil/upload_test.go`
- [ ] `websvrutil/cookie_test.go`
- [ ] `websvrutil/static_test.go`
- [ ] `websvrutil/shutdown_test.go`
- [ ] `websvrutil/storage_test.go`
- [ ] `websvrutil/coverage_test.go`
- [ ] `websvrutil/coverage_complete_test.go`
- [ ] `websvrutil/coverage_additional_test.go`
- [ ] `websvrutil/integration_test.go`
- [ ] `websvrutil/benchmark_test.go`
- [ ] `websvrutil/example_test.go`

**작업 노트**:
- 미들웨어 실행 순서와 체인 구조 명확히 설명
- 컨텍스트 생명주기와 스레드 안전성 문서화
- CSRF/세션 보안 고려사항 상세 기술
- 성능 특성 및 최적화 팁 포함

---

## Core Utility Packages (핵심 유틸리티 패키지)
**우선순위**: ⭐⭐⭐⭐ (높음)  
**작업 순서**: 병렬 진행 가능

### 2. sliceutil 패키지 (Slice Utilities)
**패키지 설명**: 슬라이스 조작 및 변환을 위한 유틸리티 함수 모음  
**작업 중요도**: High - 자주 사용되는 범용 유틸리티

#### Core Files
- [ ] `sliceutil/sliceutil.go` - 패키지 메인 및 기본 정의
- [ ] `sliceutil/basic.go` - 기본 슬라이스 연산
- [ ] `sliceutil/advanced.go` - 고급 슬라이스 연산
- [ ] `sliceutil/slice.go` - 범용 슬라이스 함수

#### Functional Operations
- [ ] `sliceutil/transform.go` - Map, Filter 등 변환 함수
- [ ] `sliceutil/aggregate.go` - Reduce, Sum 등 집계 함수
- [ ] `sliceutil/predicate.go` - 조건 검사 함수
- [ ] `sliceutil/conditional.go` - 조건부 연산

#### Set Operations
- [ ] `sliceutil/set.go` - 집합 연산 (Union, Intersection 등)
- [ ] `sliceutil/diff.go` - 차이 비교 함수

#### Indexing & Sorting
- [ ] `sliceutil/index.go` - 인덱스 검색 및 탐색
- [ ] `sliceutil/sort.go` - 정렬 관련 함수

#### Combinatorial Operations
- [ ] `sliceutil/combinatorial.go` - 조합, 순열 등
- [ ] `sliceutil/statistics.go` - 통계 함수

#### Utility Functions
- [ ] `sliceutil/util.go` - 기타 유틸리티 함수

#### Test Files
- [ ] `sliceutil/sliceutil_test.go`
- [ ] `sliceutil/basic_test.go`
- [ ] `sliceutil/advanced_test.go`
- [ ] `sliceutil/slice_test.go`
- [ ] `sliceutil/transform_test.go`
- [ ] `sliceutil/aggregate_test.go`
- [ ] `sliceutil/predicate_test.go`
- [ ] `sliceutil/conditional_test.go`
- [ ] `sliceutil/set_test.go`
- [ ] `sliceutil/diff_test.go`
- [ ] `sliceutil/index_test.go`
- [ ] `sliceutil/sort_test.go`
- [ ] `sliceutil/combinatorial_test.go`
- [ ] `sliceutil/statistics_test.go`
- [ ] `sliceutil/util_test.go`

**작업 노트**:
- 슬라이스 재할당 조건 명확히 설명
- 시간/공간 복잡도 명시
- nil 슬라이스 vs 빈 슬라이스 처리 방식 문서화
- 대용량 데이터 처리 시 성능 고려사항

### 3. maputil 패키지 (Map Utilities)
**패키지 설명**: 맵 조작 및 변환을 위한 유틸리티 함수 모음

#### Core Files
- [ ] `maputil/maputil.go` - 패키지 메인 및 기본 정의
- [ ] `maputil/basic.go` - 기본 맵 연산

#### Functional Operations
- [ ] `maputil/transform.go` - 맵 변환 함수
- [ ] `maputil/filter.go` - 필터링 함수
- [ ] `maputil/aggregate.go` - 집계 함수
- [ ] `maputil/predicate.go` - 조건 검사

#### Key/Value Operations
- [ ] `maputil/keys.go` - 키 관련 연산
- [ ] `maputil/values.go` - 값 관련 연산

#### Advanced Operations
- [ ] `maputil/merge.go` - 맵 병합 함수
- [ ] `maputil/nested.go` - 중첩 맵 처리
- [ ] `maputil/comparison.go` - 맵 비교
- [ ] `maputil/convert.go` - 타입 변환
- [ ] `maputil/default.go` - 기본값 처리

#### Utility Functions
- [ ] `maputil/util.go` - 기타 유틸리티

#### Test Files
- [ ] `maputil/maputil_test.go`
- [ ] `maputil/basic_test.go`
- [ ] `maputil/transform_test.go`
- [ ] `maputil/filter_test.go`
- [ ] `maputil/aggregate_test.go`
- [ ] `maputil/predicate_test.go`
- [ ] `maputil/keys_test.go`
- [ ] `maputil/values_test.go`
- [ ] `maputil/merge_test.go`
- [ ] `maputil/nested_test.go`
- [ ] `maputil/comparison_test.go`
- [ ] `maputil/convert_test.go`
- [ ] `maputil/default_test.go`
- [ ] `maputil/util_test.go`

**작업 노트**:
- 맵의 불변성 여부 명시
- nil 맵 처리 방식 문서화
- 동시성 안전성 주의사항
- 대용량 맵 처리 성능 팁

### 4. stringutil 패키지 (String Utilities)
**패키지 설명**: 문자열 조작, 검증, 변환을 위한 종합 유틸리티

#### Core Files
- [ ] `stringutil/stringutil.go` - 패키지 메인

#### String Manipulation
- [ ] `stringutil/manipulation.go` - 문자열 조작
- [ ] `stringutil/case.go` - 대소문자 변환
- [ ] `stringutil/formatting.go` - 포맷팅 함수
- [ ] `stringutil/builder.go` - 문자열 빌더 유틸리티

#### String Analysis
- [ ] `stringutil/search.go` - 검색 함수
- [ ] `stringutil/comparison.go` - 비교 함수
- [ ] `stringutil/distance.go` - 거리 계산 (Levenshtein 등)
- [ ] `stringutil/validation.go` - 검증 함수

#### Encoding & Unicode
- [ ] `stringutil/encoding.go` - 인코딩 변환
- [ ] `stringutil/unicode.go` - 유니코드 처리

#### Utility Functions
- [ ] `stringutil/utils.go` - 기타 유틸리티

#### Test Files
- [ ] `stringutil/manipulation_test.go`
- [ ] `stringutil/case_test.go`
- [ ] `stringutil/formatting_test.go`
- [ ] `stringutil/builder_test.go`
- [ ] `stringutil/search_test.go`
- [ ] `stringutil/comparison_test.go`
- [ ] `stringutil/distance_test.go`
- [ ] `stringutil/validation_test.go`
- [ ] `stringutil/encoding_test.go`
- [ ] `stringutil/unicode_test.go`
- [ ] `stringutil/utils_test.go`

**작업 노트**:
- UTF-8 인코딩 처리 방식 명확히 설명
- 국제화(i18n) 고려사항 문서화
- 정규표현식 패턴 설명
- 성능 최적화 팁 (strings.Builder 사용 등)

### 5. timeutil 패키지 (Time Utilities)
**패키지 설명**: 시간 처리, 변환, 포맷팅을 위한 유틸리티

#### Core Files
- [ ] `timeutil/timeutil.go` - 패키지 메인
- [ ] `timeutil/constants.go` - 시간 관련 상수

#### Time Operations
- [ ] `timeutil/parse.go` - 시간 파싱
- [ ] `timeutil/format.go` - 시간 포맷팅
- [ ] `timeutil/format_korean_test.go` - 한국어 포맷 테스트
- [ ] `timeutil/string.go` - 문자열 변환

#### Time Calculations
- [ ] `timeutil/arithmetic.go` - 시간 연산
- [ ] `timeutil/diff.go` - 시간 차이 계산
- [ ] `timeutil/comparison.go` - 시간 비교
- [ ] `timeutil/age.go` - 나이 계산

#### Special Time Types
- [ ] `timeutil/week.go` - 주(week) 관련 함수
- [ ] `timeutil/month.go` - 월(month) 관련 함수
- [ ] `timeutil/business.go` - 영업일 계산
- [ ] `timeutil/relative.go` - 상대 시간

#### Time Utilities
- [ ] `timeutil/unix.go` - Unix 타임스탬프
- [ ] `timeutil/timezone.go` - 시간대 처리
- [ ] `timeutil/sleep.go` - Sleep 유틸리티

#### Test Files
- [ ] `timeutil/parse_test.go`
- [ ] `timeutil/string_test.go`
- [ ] `timeutil/week_test.go`
- [ ] `timeutil/month_test.go`
- [ ] `timeutil/sleep_test.go`
- [ ] `timeutil/timeutil_comprehensive_test.go`

**작업 노트**:
- 시간대(Timezone) 처리 방식 상세 설명
- DST(Daylight Saving Time) 고려사항
- Locale별 포맷팅 차이점
- 성능 고려사항 (time.Now() 호출 최소화 등)

### 6. fileutil 패키지 (File Utilities)
**패키지 설명**: 파일 및 디렉터리 조작을 위한 유틸리티

#### Core Files
- [ ] `fileutil/fileutil.go` - 패키지 메인
- [ ] `fileutil/errors.go` - 에러 정의
- [ ] `fileutil/options.go` - 옵션 설정

#### File Operations
- [ ] `fileutil/read.go` - 파일 읽기
- [ ] `fileutil/write.go` - 파일 쓰기
- [ ] `fileutil/copy.go` - 파일 복사
- [ ] `fileutil/move.go` - 파일 이동
- [ ] `fileutil/delete.go` - 파일 삭제

#### Directory Operations
- [ ] `fileutil/dir.go` - 디렉터리 관리

#### File Information
- [ ] `fileutil/info.go` - 파일 정보 조회
- [ ] `fileutil/path.go` - 경로 처리
- [ ] `fileutil/hash.go` - 파일 해시 계산

#### Test Files
- [ ] `fileutil/fileutil_test.go`

**작업 노트**:
- 파일 권한 처리 방식 (Unix vs Windows)
- 심볼릭 링크 처리 주의사항
- 대용량 파일 처리 전략
- 에러 처리 및 복구 방법

---

## Supporting Packages (지원 패키지)
**우선순위**: ⭐⭐⭐ (중간)

### 7. httputil 패키지 (HTTP Utilities)
**패키지 설명**: HTTP 클라이언트 및 요청 처리 유틸리티

#### Core Files
- [ ] `httputil/httputil.go` - 패키지 메인
- [ ] `httputil/client.go` - HTTP 클라이언트
- [ ] `httputil/options.go` - 클라이언트 옵션
- [ ] `httputil/errors.go` - 에러 정의

#### Request/Response Handling
- [ ] `httputil/simple.go` - 간단한 요청 함수
- [ ] `httputil/response.go` - 응답 처리
- [ ] `httputil/form.go` - 폼 데이터 처리
- [ ] `httputil/file.go` - 파일 업로드/다운로드

#### HTTP Features
- [ ] `httputil/cookie.go` - 쿠키 관리
- [ ] `httputil/url.go` - URL 처리

#### Test Files
- [ ] `httputil/httputil_test.go`
- [ ] `httputil/cookie_test.go`

**작업 노트**:
- Timeout 및 Context 처리 방식
- 재시도 정책 및 백오프 전략
- TLS/SSL 설정 방법
- 에러 처리 및 로깅 전략

### 8. logging 패키지 (Logging Utilities)
**패키지 설명**: 구조화된 로깅 및 로그 관리

#### Core Files
- [ ] `logging/logger.go` - 로거 구현
- [ ] `logging/level.go` - 로그 레벨
- [ ] `logging/options.go` - 로거 옵션
- [ ] `logging/appconfig.go` - 애플리케이션 설정
- [ ] `logging/banner.go` - 배너 출력

#### Test Files
- [ ] `logging/logger_test.go`

**작업 노트**:
- 로그 레벨별 사용 시나리오
- 로그 로테이션 설정
- 구조화된 로깅 (structured logging) 방식
- 성능 최적화 (비동기 로깅 등)

### 9. errorutil 패키지 (Error Utilities)
**패키지 설명**: 에러 생성, 래핑, 검사를 위한 유틸리티

#### Core Files
- [ ] `errorutil/error.go` - 에러 생성 및 래핑
- [ ] `errorutil/types.go` - 에러 타입 정의
- [ ] `errorutil/inspect.go` - 에러 검사 및 분석

#### Test Files
- [ ] `errorutil/error_test.go`
- [ ] `errorutil/types_test.go`
- [ ] `errorutil/inspect_test.go`

**작업 노트**:
- errors.Is, errors.As 사용법
- 에러 래핑 깊이 제한
- 컨텍스트 정보 추가 방법
- 표준 라이브러리와의 호환성

### 10. random 패키지 (Random Utilities)
**패키지 설명**: 난수 생성 및 랜덤 문자열 생성

#### Core Files
- [ ] `random/string.go` - 랜덤 문자열 생성

#### Test Files
- [ ] `random/string_test.go`

**작업 노트**:
- 암호학적 안전성 여부 명시
- 시드 관리 방법
- 문자 세트 커스터마이징
- 성능 특성

---

## Database Packages (데이터베이스 패키지)
**우선순위**: ⭐⭐⭐⭐ (높음)

### 11. database/mysql 패키지
**패키지 설명**: MySQL 데이터베이스 연동 및 관리

#### Core Files
- [ ] `database/mysql/client.go` - MySQL 클라이언트
- [ ] `database/mysql/options.go` - 연결 옵션
- [ ] `database/mysql/errors.go` - 에러 정의

#### Database Operations
- [ ] `database/mysql/batch.go` - 배치 작업
- [ ] `database/mysql/migration.go` - 마이그레이션
- [ ] `database/mysql/schema.go` - 스키마 관리
- [ ] `database/mysql/pagination.go` - 페이지네이션
- [ ] `database/mysql/transaction.go` - 트랜잭션 관리

#### Monitoring & Utilities
- [ ] `database/mysql/metrics.go` - 메트릭 수집
- [ ] `database/mysql/stats.go` - 통계 정보
- [ ] `database/mysql/pool.go` - 연결 풀 관리
- [ ] `database/mysql/query.go` - 쿼리 빌더
- [ ] `database/mysql/helper.go` - 헬퍼 함수

#### Test Files
- [ ] `database/mysql/stats_test.go`
- [ ] (기타 테스트 파일들)

**작업 노트**:
- 연결 풀 설정 최적화 가이드
- 트랜잭션 격리 수준 설명
- 데드락 처리 방법
- 쿼리 성능 최적화 팁
- 마이그레이션 롤백 전략

### 12. database/redis 패키지
**패키지 설명**: Redis 연동 및 명령 실행

#### Core Files
- [ ] `database/redis/client.go` - Redis 클라이언트
- [ ] `database/redis/options.go` - 연결 옵션
- [ ] `database/redis/errors.go` - 에러 정의

#### Redis Operations
- [ ] `database/redis/string.go` - String 명령
- [ ] `database/redis/hash.go` - Hash 명령
- [ ] `database/redis/set.go` - Set 명령
- [ ] `database/redis/list.go` - List 명령
- [ ] `database/redis/sortedset.go` - Sorted Set 명령
- [ ] `database/redis/pipeline.go` - 파이프라인
- [ ] `database/redis/pubsub.go` - Pub/Sub

#### Advanced Features
- [ ] `database/redis/transaction.go` - 트랜잭션
- [ ] `database/redis/cluster.go` - 클러스터
- [ ] `database/redis/sentinel.go` - Sentinel
- [ ] `database/redis/scan.go` - SCAN 명령

#### Monitoring & Utilities
- [ ] `database/redis/pool.go` - 연결 풀
- [ ] `database/redis/metrics.go` - 메트릭
- [ ] `database/redis/helper.go` - 헬퍼 함수

#### Test Files
- [ ] (Redis 테스트 파일들)

**작업 노트**:
- 각 데이터 타입별 사용 시나리오
- 파이프라인 vs 트랜잭션 비교
- 연결 풀 최적화 전략
- 클러스터/Sentinel 설정 가이드
- 메모리 관리 주의사항

---

## Examples Directory (예제 디렉터리)
**우선순위**: ⭐⭐⭐ (중간)  
**작업 특성**: 교육적 내용 중심

### 13. Examples - 모든 예제 파일
**작업 목적**: 초보자도 이해할 수 있는 명확한 예제 제공

#### Example Files
- [ ] `examples/errorutil/main.go` - errorutil 사용 예제
- [ ] `examples/fileutil/main.go` - fileutil 사용 예제
- [ ] `examples/httputil/main.go` - httputil 사용 예제
- [ ] `examples/logging/main.go` - logging 사용 예제
- [ ] `examples/maputil/main.go` - maputil 사용 예제
- [ ] `examples/mysql/main.go` - MySQL 연동 예제
- [ ] `examples/random_string/main.go` - 난수 생성 예제
- [ ] `examples/redis/main.go` - Redis 연동 예제
- [ ] `examples/sliceutil/main.go` - sliceutil 사용 예제
- [ ] `examples/stringutil/main.go` - stringutil 사용 예제
- [ ] `examples/timeutil/main.go` - timeutil 사용 예제
- [ ] `examples/websvrutil/main.go` - websvrutil 사용 예제

**예제 주석 작성 가이드**:
1. **전체 흐름 설명**: 예제의 목적과 실행 흐름을 상단에 명시
2. **단계별 주석**: 각 코드 블록마다 무엇을 하는지 설명
3. **출력 예시**: 예상되는 출력 결과를 주석으로 표시
4. **학습 포인트**: 핵심 개념이나 주의사항을 강조
5. **실행 방법**: 예제 실행 명령어 및 필요한 환경 설정

**예제 주석 템플릿**:
```go
// Example: [Example Name]
// This example demonstrates [purpose and key concepts].
//
// Learning points:
//   - Point 1
//   - Point 2
//
// How to run:
//   go run examples/[package]/main.go
//
// Expected output:
//   [output description]
//
// 예제: [예제 이름]
// 이 예제는 [목적과 핵심 개념]을 보여줍니다.
//
// 학습 포인트:
//   - 포인트 1
//   - 포인트 2
//
// 실행 방법:
//   go run examples/[package]/main.go
//
// 예상 출력:
//   [출력 설명]
```

---

## Test Files (테스트 파일)
**우선순위**: ⭐⭐⭐ (중간)  
**작업 전략**: 주요 패키지 작업 시 함께 진행

### 테스트 파일 주석 작성 원칙
1. **테스트 목적**: 각 테스트 함수가 무엇을 검증하는지 명시
2. **테스트 시나리오**: Given-When-Then 패턴으로 구조화
3. **경계 조건**: 엣지 케이스와 에러 케이스 설명
4. **설정 정보**: 테스트에 필요한 환경 설정이나 전제 조건

**테스트 주석 템플릿**:
```go
// TestFunctionName tests [what is being tested].
//
// Test scenarios:
//   - Scenario 1: [description]
//   - Scenario 2: [description]
//
// TestFunctionName은 [테스트 대상]을 검증합니다.
//
// 테스트 시나리오:
//   - 시나리오 1: [설명]
//   - 시나리오 2: [설명]
func TestFunctionName(t *testing.T) {
    // Given: [setup description] / 준비: [설정 설명]
    
    // When: [action description] / 실행: [동작 설명]
    
    // Then: [assertion description] / 검증: [검증 설명]
}
```

---

## Verification Steps (검증 단계)
**완료 조건**: 모든 패키지 작업 완료 후 실행

### Phase 1: 코드 품질 검증
- [ ] **Go Vet 실행**: `go vet ./...`
- [ ] **Go Fmt 검사**: `go fmt ./...`
- [ ] **Golint 실행**: `golint ./...` (설치된 경우)
- [ ] **주석 커버리지 검사**: 모든 public 함수/타입 주석 확인

### Phase 2: 테스트 검증
- [ ] **전체 테스트**: `go test ./...`
- [ ] **패키지별 테스트**: `go test ./[package]` (주요 패키지)
- [ ] **Race Detector**: `go test -race ./...`
- [ ] **Coverage 측정**: `go test -cover ./...`

### Phase 3: 문서 검증
- [ ] **README 업데이트**: 변경사항 반영
- [ ] **BILINGUAL_AUDIT.md 업데이트**: 완료 항목 체크
- [ ] **CHANGELOG 작성**: 주요 변경사항 기록
- [ ] **API 문서 생성**: `godoc` 또는 `pkgsite` 확인

### Phase 4: 최종 검토
- [ ] **일관성 검사**: 주석 스타일 통일성 확인
- [ ] **예제 동작 확인**: 모든 예제 실행 및 출력 검증
- [ ] **링크 검증**: 문서 내 모든 링크 작동 확인
- [ ] **오타 검사**: 영문/한글 오타 확인

---

## Progress Tracking (진행 상황 추적)

### 현재 진행 상황
- **작업 시작일**: [YYYY-MM-DD]
- **현재 작업 중인 파일**: 
- **완료된 패키지**: 
- **전체 진행률**: 0% (0/[total] 파일)

### Next Steps (다음 단계)
1. [ ] [다음 작업 항목]
2. [ ] [다음 작업 항목]
3. [ ] [다음 작업 항목]

### 발견된 이슈 및 특이사항
- [이슈 1]
- [이슈 2]

### 작업 히스토리
| 날짜 | 패키지/파일 | 작업 내용 | 상태 |
|------|------------|----------|------|
| 2025-10-17 | todo-codex.md | 마스터 체크리스트 보강 | ✅ |
|  |  |  |  |

---

## Automation & Tools (자동화 및 도구)

### 주석 검증 스크립트
```bash
# 영문만 있는 주석 찾기
grep -r "^// [^/]*$" --include="*.go" --exclude-dir="vendor"

# 한글만 있는 주석 찾기
grep -r "^// .*[ㄱ-ㅎㅏ-ㅣ가-힣].*$" --include="*.go" --exclude-dir="vendor" | grep -v "/ "

# public 함수 중 주석 없는 것 찾기
grep -r "^func [A-Z]" --include="*.go" --exclude-dir="vendor" -B1 | grep -v "^//"
```

### 통계 수집
```bash
# 전체 .go 파일 수
find . -name "*.go" -not -path "./vendor/*" | wc -l

# 전체 함수 수
grep -r "^func " --include="*.go" --exclude-dir="vendor" | wc -l

# 주석 라인 수
grep -r "^//" --include="*.go" --exclude-dir="vendor" | wc -l
```

---
