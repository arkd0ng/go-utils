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

## Complete File Checklist / 전체 파일 체크리스트

### 📌 중요 안내
- **모든 파일은 동등하게 중요합니다** - 우선순위 없음
- **누락 없이 완료**하는 것이 목표입니다
- **순서는 제안일 뿐** - 편한 순서로 작업 가능
- **체크박스 상태**:
  - `[ ]` : 작업 대기 중
  - `[-]` : 현재 작업 중
  - `[x]` : 작업 완료
- **각 파일마다** 충분히 자세하고 친절한 주석 작성 필수

### 🎯 작업 완료 기준
각 파일은 다음 기준을 **모두** 충족해야 완료로 표시:
1. ✅ 모든 public 함수/타입/상수에 상세한 주석
2. ✅ 영문/한글 주석 모두 동일한 수준의 상세함
3. ✅ 복잡한 함수에 사용 예제 포함
4. ✅ 모든 에러 케이스 문서화
5. ✅ Thread-safety, 성능 특성 명시 (해당되는 경우)
6. ✅ 테스트 실행 성공
7. ✅ 관련 문서 업데이트 (BILINGUAL_AUDIT.md, CHANGELOG)

---

## 1. websvrutil Package
**파일 수**: 총 51개 파일 (소스 27개 + 테스트 24개)  
**패키지 설명**: 웹 서버 애플리케이션 개발을 위한 핵심 유틸리티  
**특별 주의사항**: 
- 미들웨어 실행 순서 명확히 설명
- 컨텍스트 생명주기 상세 문서화
- CSRF/세션 보안 고려사항 강조
- 동시성 안전성 명시

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
- 각 함수마다 실제 사용 예제 작성
- 초보자를 위한 상세한 설명 필수

---

## 2. sliceutil Package
**파일 수**: 총 32개 파일 (소스 16개 + 테스트 16개)  
**패키지 설명**: 슬라이스 조작 및 변환을 위한 유틸리티 함수 모음  
**특별 주의사항**:
- 슬라이스 재할당 조건 명확히 설명
- 시간/공간 복잡도 모든 함수에 명시
- nil 슬라이스 vs 빈 슬라이스 처리 방식
- 대용량 데이터 처리 시 성능 고려사항

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
- 각 함수의 메모리 할당 동작 설명
- 실제 사용 예제를 풍부하게 제공

---

## 3. maputil Package
**파일 수**: 총 28개 파일 (소스 14개 + 테스트 14개)  
**패키지 설명**: 맵 조작 및 변환을 위한 유틸리티 함수 모음  
**특별 주의사항**:
- 맵의 불변성 여부 명시
- nil 맵 처리 방식 문서화
- 동시성 안전성 주의사항
- 대용량 맵 처리 성능 팁

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
- 키/값 타입 제약 조건 설명
- 중첩 맵 처리 시 주의사항 상세 기술

---

## 4. stringutil Package
**파일 수**: 총 22개 파일 (소스 11개 + 테스트 11개)  
**패키지 설명**: 문자열 조작, 검증, 변환을 위한 종합 유틸리티  
**특별 주의사항**:
- UTF-8 인코딩 처리 방식 명확히 설명
- 국제화(i18n) 고려사항 문서화
- 정규표현식 패턴 상세 설명
- 성능 최적화 팁 (strings.Builder 사용 등)

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
**작업 노트**:
- UTF-8 인코딩 처리 방식 명확히 설명
- 국제화(i18n) 고려사항 문서화
- 정규표현식 패턴 설명
- 성능 최적화 팁 (strings.Builder 사용 등)
- 유니코드 처리 시 주의사항
- 다양한 문자열 조작 예제 제공

---

## 5. timeutil Package
**파일 수**: 총 24개 파일 (소스 12개 + 테스트 12개)  
**패키지 설명**: 시간 처리, 변환, 포맷팅을 위한 유틸리티  
**특별 주의사항**:
- 시간대(Timezone) 처리 방식 상세 설명
- DST(Daylight Saving Time) 고려사항
- Locale별 포맷팅 차이점
- 성능 고려사항 (time.Now() 호출 최소화 등)
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
**작업 노트**:
- 시간대(Timezone) 처리 방식 상세 설명
- DST(Daylight Saving Time) 고려사항
- Locale별 포맷팅 차이점
- 성능 고려사항 (time.Now() 호출 최소화 등)
- 시간 계산 시 정확도 이슈
- 다양한 시간 형식 파싱 예제

---

## 6. fileutil Package
**파일 수**: 총 20개 파일 (소스 10개 + 테스트 10개)  
**패키지 설명**: 파일 및 디렉터리 조작을 위한 유틸리티  
**특별 주의사항**:
- 파일 권한 처리 방식 (Unix vs Windows)
- 심볼릭 링크 처리 주의사항
- 대용량 파일 처리 전략
- 에러 처리 및 복구 방법
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
**작업 노트**:
- 파일 권한 처리 방식 (Unix vs Windows)
- 심볼릭 링크 처리 주의사항
- 대용량 파일 처리 전략
- 에러 처리 및 복구 방법
- 파일 시스템별 차이점 설명
- 안전한 파일 작업 패턴 제시

---

## 7. httputil Package
**파일 수**: 총 20개 파일 (소스 10개 + 테스트 10개)  
**패키지 설명**: HTTP 클라이언트 및 요청 처리 유틸리티  
**특별 주의사항**:
- Timeout 및 Context 처리 방식
- 재시도 정책 및 백오프 전략
- TLS/SSL 설정 방법
- 에러 처리 및 로깅 전략
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
**작업 노트**:
- Timeout 및 Context 처리 방식
- 재시도 정책 및 백오프 전략
- TLS/SSL 설정 방법
- 에러 처리 및 로깅 전략
- HTTP 클라이언트 풀링 설명
- 다양한 HTTP 시나리오 예제

---

## 8. logging Package
**파일 수**: 총 12개 파일 (소스 6개 + 테스트 6개)  
**패키지 설명**: 구조화된 로깅 및 로그 관리  
**특별 주의사항**:
- 로그 레벨별 사용 시나리오
- 로그 로테이션 설정
- 구조화된 로깅 방식
- 성능 최적화 (비동기 로깅 등)
- [ ] `httputil/cookie.go` - 쿠키 관리
- [ ] `httputil/url.go` - URL 처리

#### Test Files
- [ ] `httputil/httputil_test.go`
- [ ] `httputil/cookie_test.go`

**작업 노트**:
- Timeout 및 Context 처리 방식
- 재시도 정책 및 백오프 전략
- TLS/SSL 설정 방법
**작업 노트**:
- 로그 레벨별 사용 시나리오
- 로그 로테이션 설정
- 구조화된 로깅 (structured logging) 방식
- 성능 최적화 (비동기 로깅 등)
- 로그 포맷 커스터마이징
- 프로덕션 환경 로깅 모범 사례

---

## 9. errorutil Package
**파일 수**: 총 6개 파일 (소스 3개 + 테스트 3개)  
**패키지 설명**: 에러 생성, 래핑, 검사를 위한 유틸리티  
**특별 주의사항**:
- errors.Is, errors.As 사용법
- 에러 래핑 깊이 제한
- 컨텍스트 정보 추가 방법
- 표준 라이브러리와의 호환성
- [ ] `logging/options.go` - 로거 옵션
- [ ] `logging/appconfig.go` - 애플리케이션 설정
- [ ] `logging/banner.go` - 배너 출력

#### Test Files
- [ ] `logging/logger_test.go`
**작업 노트**:
- errors.Is, errors.As 사용법
- 에러 래핑 깊이 제한
- 컨텍스트 정보 추가 방법
- 표준 라이브러리와의 호환성
- 에러 체인 추적 방법
- 사용자 정의 에러 타입 설계 가이드

---

## 10. random Package
**파일 수**: 총 2개 파일 (소스 1개 + 테스트 1개)  
**패키지 설명**: 난수 생성 및 랜덤 문자열 생성  
**특별 주의사항**:
- 암호학적 안전성 여부 명시
**작업 노트**:
- 암호학적 안전성 여부 명시
- 시드 관리 방법
- 문자 세트 커스터마이징
- 성능 특성
- 보안 용도 vs 일반 용도 구분
- 다양한 난수 생성 예제

---

## 11. database/mysql Package
**파일 수**: 총 36개 파일 (소스 18개 + 테스트 18개)  
**패키지 설명**: MySQL 데이터베이스 연동 및 관리  
**특별 주의사항**:
- 연결 풀 설정 최적화 가이드
- 트랜잭션 격리 수준 설명
- 데드락 처리 방법
- 쿼리 성능 최적화 팁
- 마이그레이션 롤백 전략

#### Core Files
- [ ] `database/mysql/client.go` - MySQL 클라이언트 메인
- [ ] `database/mysql/connection.go` - 연결 관리
- [ ] `database/mysql/config.go` - 설정 관리
- [ ] `database/mysql/options.go` - 옵션 설정
- [ ] `database/mysql/types.go` - 타입 정의
- [ ] `database/mysql/errors.go` - 에러 정의

#### Query & Transaction
- [ ] `database/mysql/simple.go` - 간단한 쿼리 함수
- [ ] `database/mysql/builder.go` - 쿼리 빌더
- [ ] `database/mysql/select_options.go` - SELECT 옵션
- [ ] `database/mysql/transaction.go` - 트랜잭션 관리
- [ ] `database/mysql/scan.go` - 스캔 유틸리티

#### Advanced Operations
- [ ] `database/mysql/batch.go` - 배치 작업
- [ ] `database/mysql/upsert.go` - UPSERT 작업
- [ ] `database/mysql/softdelete.go` - 소프트 삭제
- [ ] `database/mysql/pagination.go` - 페이지네이션

#### Database Management
- [ ] `database/mysql/migration.go` - 마이그레이션
- [ ] `database/mysql/schema.go` - 스키마 관리
- [ ] `database/mysql/export.go` - 데이터 내보내기

#### Monitoring & Utilities
- [ ] `database/mysql/metrics.go` - 메트릭 수집
- [ ] `database/mysql/stats.go` - 통계 정보
- [ ] `database/mysql/retry.go` - 재시도 로직
- [ ] `database/mysql/rotation.go` - 로테이션

#### Test Files
- [ ] `database/mysql/client_test.go`
- [ ] `database/mysql/batch_test.go`
- [ ] `database/mysql/upsert_test.go`
- [ ] `database/mysql/softdelete_test.go`
- [ ] `database/mysql/pagination_test.go`
- [ ] `database/mysql/migration_test.go`
- [ ] `database/mysql/schema_test.go`
- [ ] `database/mysql/export_test.go`
- [ ] `database/mysql/metrics_test.go`
- [ ] `database/mysql/stats_test.go`
- [ ] `database/mysql/mysql_integration_test.go`
- [ ] `database/mysql/test_utils_test.go`
- [ ] `database/mysql/testhelper_test.go`

**작업 노트**:
- 연결 풀 관리 상세 설명
- 트랜잭션 사용 패턴 및 주의사항
- 쿼리 빌더 사용 예제
- 마이그레이션 전략 및 롤백 방법
- 성능 모니터링 및 최적화
- SQL 인젝션 방지 방법

---

## 12. database/redis Package
**파일 수**: 총 28개 파일 (소스 14개 + 테스트 14개)  
**패키지 설명**: Redis 연동 및 명령 실행  
**특별 주의사항**:
- 각 데이터 타입별 사용 시나리오
- 파이프라인 vs 트랜잭션 비교
- 연결 풀 최적화 전략
- 클러스터 설정 가이드
- 메모리 관리 주의사항

#### Core Files
- [ ] `database/redis/client.go` - Redis 클라이언트 메인
- [ ] `database/redis/connection.go` - 연결 관리
- [ ] `database/redis/config.go` - 설정 관리
- [ ] `database/redis/options.go` - 옵션 설정
- [ ] `database/redis/types.go` - 타입 정의
- [ ] `database/redis/errors.go` - 에러 정의

#### Data Type Operations
- [ ] `database/redis/string.go` - String 타입 명령
- [ ] `database/redis/hash.go` - Hash 타입 명령
- [ ] `database/redis/list.go` - List 타입 명령
- [ ] `database/redis/set.go` - Set 타입 명령
- [ ] `database/redis/zset.go` - Sorted Set 타입 명령

#### Advanced Features
- [ ] `database/redis/key.go` - 키 관리
- [ ] `database/redis/pipeline.go` - 파이프라인
- [ ] `database/redis/transaction.go` - 트랜잭션
- [ ] `database/redis/pubsub.go` - Pub/Sub
- [ ] `database/redis/retry.go` - 재시도 로직

#### Test Files
- [ ] `database/redis/client_test.go`
- [ ] `database/redis/string_test.go`
- [ ] `database/redis/hash_test.go`
- [ ] `database/redis/list_test.go`
- [ ] `database/redis/set_test.go`
- [ ] `database/redis/zset_test.go`
- [ ] `database/redis/key_test.go`
- [ ] `database/redis/pipeline_test.go`
- [ ] `database/redis/transaction_test.go`
- [ ] `database/redis/pubsub_test.go`
- [ ] `database/redis/testhelper_test.go`

**작업 노트**:
- 각 데이터 타입의 특징과 사용 시나리오
- 파이프라인 사용을 통한 성능 최적화
- 트랜잭션 WATCH/MULTI/EXEC 설명
- Pub/Sub 패턴 구현 가이드
- 연결 풀링 및 재시도 전략
- TTL 및 메모리 관리

---

## 13. validation Package
**파일 수**: 총 8개 파일 (소스 4개 + 테스트 4개)  
**패키지 설명**: 데이터 검증 및 유효성 검사  
**특별 주의사항**:
- 검증 규칙 작성 방법
- 커스텀 검증기 구현
- 에러 메시지 국제화
- 성능 최적화

#### Core Files
- [ ] `validation/validator.go` - 메인 검증기
- [ ] `validation/types.go` - 타입 정의
- [ ] `validation/errors.go` - 에러 정의
- [ ] `validation/version.go` - 버전 정보

#### Test Files
- [ ] `validation/validator_test.go`
- [ ] `validation/types_test.go`
- [ ] `validation/errors_test.go`

**작업 노트**:
- 다양한 검증 규칙 예제
- 중첩된 구조체 검증 방법
- 조건부 검증 구현
- 커스텀 검증 함수 작성 가이드
- 에러 메시지 커스터마이징
- 성능 고려사항

---

## 14. examples Directory
**파일 수**: 총 12개 예제 파일  
**패키지 설명**: 각 패키지의 사용 예제  
**특별 주의사항**:
- 초보자 친화적인 설명
- 단계별 주석
- 예상 출력 명시
- 실행 방법 안내
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

**예제 주석 작성 템플릿**:
```go
// Example: [Example Name]
//
// Purpose:
// This example demonstrates [specific purpose and key concepts].
// It shows how to [main task] using [package name].
//
// Prerequisites:
//   - [any required setup, e.g., database running]
//   - [environment variables needed]
//
// Learning points:
//   - [Key concept 1]: [explanation]
//   - [Key concept 2]: [explanation]
//   - [Key concept 3]: [explanation]
//
// How to run:
//   go run examples/[package]/main.go
//
// Expected output:
//   [detailed description of expected output]
//   [example output text]
//
// Common issues:
//   - [Issue 1]: [solution]
//   - [Issue 2]: [solution]
//
// 예제: [예제 이름]
//
// 목적:
// 이 예제는 [구체적인 목적과 핵심 개념]을 보여줍니다.
// [패키지 이름]을 사용하여 [주요 작업]을 수행하는 방법을 보여줍니다.
//
// 사전 요구사항:
//   - [필요한 설정, 예: 데이터베이스 실행 중]
//   - [필요한 환경 변수]
//
// 학습 포인트:
//   - [핵심 개념 1]: [설명]
//   - [핵심 개념 2]: [설명]
//   - [핵심 개념 3]: [설명]
//
// 실행 방법:
//   go run examples/[package]/main.go
//
// 예상 출력:
//   [예상 출력에 대한 상세 설명]
//   [출력 예시 텍스트]
//
// 일반적인 문제:
//   - [문제 1]: [해결 방법]
//   - [문제 2]: [해결 방법]

package main

import (
    // imports with comments explaining why each is needed
    // 각 import가 필요한 이유를 주석으로 설명
)

func main() {
    // Step 1: [First step description]
    // 1단계: [첫 번째 단계 설명]
    
    // Step 2: [Second step description]
    // 2단계: [두 번째 단계 설명]
    
    // ... and so on
}
```

**작업 노트**:
- 각 예제의 전체 흐름을 상단에 명확히 설명
- 코드 블록마다 영문/한글 주석 병기
- 초보자가 막힐 수 있는 부분 미리 안내
- 실제 실행 가능한 완전한 예제 제공
- 출력 결과를 주석으로 표시
- 일반적인 에러와 해결 방법 포함

---

## Test Files (테스트 파일)
**총 파일 수**: 100개 이상 (각 패키지의 테스트 파일)  
**작업 특성**: 주요 패키지 작업 시 함께 진행

### 테스트 파일 주석 작성 원칙

#### 1. 테스트 함수 주석
**목표**: 각 테스트가 무엇을 검증하는지, 왜 중요한지 명확히 설명

```go
// TestFunctionName verifies [what is being tested].
//
// Purpose:
// This test ensures that [specific behavior or requirement].
// It validates [expected behavior] under [specific conditions].
//
// Test coverage:
//   - [Scenario 1]: [what is tested]
//   - [Scenario 2]: [what is tested]
//   - [Edge case 1]: [what is tested]
//   - [Error case 1]: [what is tested]
//
// Test methodology:
// Uses [testing approach, e.g., table-driven tests, mocking]
// to verify [aspect being tested].
//
// Important assumptions:
//   - [Assumption 1]
//   - [Assumption 2]
//
// TestFunctionName은 [테스트 대상]을 검증합니다.
//
// 목적:
// 이 테스트는 [특정 동작이나 요구사항]을 보장합니다.
// [특정 조건] 하에서 [예상 동작]을 검증합니다.
//
// 테스트 범위:
//   - [시나리오 1]: [테스트 내용]
//   - [시나리오 2]: [테스트 내용]
//   - [엣지 케이스 1]: [테스트 내용]
//   - [에러 케이스 1]: [테스트 내용]
//
// 테스트 방법론:
// [테스트 방식, 예: 테이블 기반 테스트, 모킹]을 사용하여
// [테스트 대상 측면]을 검증합니다.
//
// 중요한 가정:
//   - [가정 1]
//   - [가정 2]
func TestFunctionName(t *testing.T) {
    // Given: [setup and preconditions]
    // 준비: [설정 및 사전 조건]
    
    // When: [action being tested]
    // 실행: [테스트할 동작]
    
    // Then: [expected outcomes and assertions]
    // 검증: [예상 결과 및 단언]
}
```

#### 2. 테스트 케이스 주석
테이블 기반 테스트의 각 케이스에도 주석 추가:

```go
tests := []struct {
    name    string // Test case name / 테스트 케이스 이름
    input   string // Input description / 입력 설명
    want    string // Expected output / 예상 출력
    wantErr bool   // Should return error / 에러 반환 여부
}{
    {
        name:    "valid input",
        input:   "test", // Tests normal case / 정상 케이스 테스트
        want:    "TEST",
        wantErr: false,
    },
    {
        name:    "empty string",
        input:   "", // Edge case: empty input / 엣지 케이스: 빈 입력
        want:    "",
        wantErr: true,
    },
}
```

#### 3. 테스트 헬퍼 함수 주석
```go
// setupTestDB creates a test database connection for integration tests.
// It returns the database connection and a cleanup function.
// The cleanup function should be called with defer to ensure proper cleanup.
//
// Parameters:
//   - t: testing.T instance for logging and failing
//
// Returns:
//   - *sql.DB: test database connection
//   - func(): cleanup function to close connection and remove test data
//
// setupTestDB는 통합 테스트를 위한 테스트 데이터베이스 연결을 생성합니다.
// 데이터베이스 연결과 정리 함수를 반환합니다.
// 정리 함수는 defer로 호출하여 적절한 정리를 보장해야 합니다.
//
// 매개변수:
//   - t: 로깅 및 실패를 위한 testing.T 인스턴스
//
// 반환값:
//   - *sql.DB: 테스트 데이터베이스 연결
//   - func(): 연결을 닫고 테스트 데이터를 제거하는 정리 함수
func setupTestDB(t *testing.T) (*sql.DB, func()) {
    // implementation
}
```

### 테스트 파일 체크리스트
- [ ] 모든 테스트 함수에 목적 설명
- [ ] Given-When-Then 패턴 명확히 표시
- [ ] 테스트 케이스별 의도 설명
- [ ] 엣지 케이스와 에러 케이스 문서화
- [ ] 테스트 헬퍼 함수 상세 설명
- [ ] 필요한 환경 설정 명시
- [ ] 영문/한글 주석 병기

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

### 📊 전체 진행 상황 요약
**업데이트 날짜**: 2025-10-17

| 패키지 | 총 파일 수 | 완료 파일 | 진행률 | 상태 |
|--------|-----------|----------|--------|------|
| websvrutil | 51 | 0 | 0% | 대기 |
| sliceutil | 32 | 0 | 0% | 대기 |
| maputil | 28 | 0 | 0% | 대기 |
| stringutil | 22 | 0 | 0% | 대기 |
| timeutil | 24 | 0 | 0% | 대기 |
| fileutil | 20 | 0 | 0% | 대기 |
| httputil | 20 | 0 | 0% | 대기 |
| logging | 12 | 0 | 0% | 대기 |
| errorutil | 6 | 0 | 0% | 대기 |
| random | 2 | 0 | 0% | 대기 |
| database/mysql | 36 | 0 | 0% | 대기 |
| database/redis | 28 | 0 | 0% | 대기 |
| validation | 8 | 0 | 0% | 대기 |
| examples | 12 | 0 | 0% | 대기 |
| **전체** | **~301** | **0** | **0%** | **시작 전** |

### 📝 현재 작업 상태
- **작업 시작일**: [미시작]
- **마지막 업데이트**: 2025-10-17
- **현재 작업 중인 파일**: 없음
- **현재 세션 진행 상황**: todo-codex.md 마스터 체크리스트 보강 완료

### 🎯 Next Steps (다음 작업 계획)
1. [ ] 작업 순서 결정 (권장: websvrutil 또는 errorutil부터 시작)
2. [ ] 첫 번째 파일 선택 및 체크박스 `[-]`로 변경
3. [ ] 기존 주석 분석 및 보강 계획 수립
4. [ ] 주석 작성 시작

### 💡 작업 팁
- **한 번에 한 파일씩**: 완전히 끝낸 후 다음 파일로
- **테스트와 함께**: 소스 파일 완료 후 바로 테스트 파일 작업
- **자주 커밋**: 파일 2-3개 완료 시마다 커밋
- **정기적 문서화**: 5-10개 파일 완료 시마다 BILINGUAL_AUDIT.md 업데이트
- **품질 > 속도**: 빠르게 하기보다 충분히 자세하게 작성

### 📌 발견된 이슈 및 특이사항
_작업 중 발견되는 이슈를 여기에 기록_

- [날짜] [파일명]: [이슈 설명]

### 📚 작업 히스토리
| 날짜 | 패키지/파일 | 작업 내용 | 파일 수 | 상태 |
|------|------------|----------|---------|------|
| 2025-10-17 | todo-codex.md | 마스터 체크리스트 보강 작업 | 1 | ✅ 완료 |
|  |  |  |  |  |

### 🏆 마일스톤
- [ ] **마일스톤 1**: 첫 10개 파일 완료
- [ ] **마일스톤 2**: 첫 패키지 완전 완료
- [ ] **마일스톤 3**: 100개 파일 완료
- [ ] **마일스톤 4**: 모든 소스 파일 완료
- [ ] **마일스톤 5**: 모든 테스트 파일 완료
- [ ] **마일스톤 6**: 전체 프로젝트 완료 및 최종 검증

### 📈 주간 진행 목표
_각 주의 목표 파일 수를 설정하고 추적_

- **1주차**: [목표 파일 수] (예: 20-30 파일)
- **2주차**: [목표 파일 수]
- **3주차**: [목표 파일 수]
- **완료 예상일**: [예상 날짜]

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

### 파일별 주석 비율 확인
```bash
# 특정 파일의 주석 비율 확인
count_comments() {
    local file=$1
    local total=$(wc -l < "$file")
    local comments=$(grep -c "^[[:space:]]*\/\/" "$file" || echo 0)
    local ratio=$(awk "BEGIN {printf \"%.1f\", ($comments/$total)*100}")
    echo "$file: $comments/$total lines ($ratio%)"
}

# 사용 예
count_comments "sliceutil/slice.go"
```

### 진행률 계산
```bash
# 완료된 파일 수 계산 (체크리스트에서 [x] 카운트)
grep -c "\[x\]" todo-codex.md

# 전체 항목 수 계산
grep -c "\[ \]\|\[-\]\|\[x\]" todo-codex.md
```

---

## Quick Reference (빠른 참조)

### 🚀 작업 시작하기
1. `todo-codex.md` 열기
2. 작업할 파일 선택하고 `[ ]` → `[-]`로 변경
3. 파일 열고 기존 주석 검토
4. 주석 작성 표준에 따라 주석 보강
5. 테스트 실행: `go test ./[package]`
6. 완료 시 `[-]` → `[x]`로 변경
7. 문서 업데이트 (BILINGUAL_AUDIT.md, CHANGELOG)

### 📝 주석 작성 체크리스트 (빠른 확인)
작업 완료 전 반드시 확인:

**필수 사항**:
- [ ] 모든 public 함수에 주석
- [ ] 영문/한글 모두 동일한 수준으로 상세
- [ ] Purpose, Parameters, Returns, Errors 모두 포함
- [ ] 특수 값(nil, 0 등) 처리 방법 명시
- [ ] Thread-safety 명시 (필요한 경우)

**권장 사항**:
- [ ] 복잡한 함수에 사용 예제
- [ ] 성능 특성 (시간/공간 복잡도)
- [ ] 알려진 제한사항
- [ ] 관련 함수 참조

**품질 확인**:
- [ ] 초보자가 이해할 수 있는가?
- [ ] 주석이 코드보다 짧지 않은가?
- [ ] 실제 코드와 일치하는가?

### 🎨 주석 템플릿 (복사하여 사용)

#### 간단한 함수
```go
// FunctionName [동사] [명사구].
// It [상세 설명 2-3문장].
//
// Parameters:
//   - param: [설명, 제약조건]
//
// Returns:
//   - [타입]: [설명]
//
// FunctionName은 [동사] [명사구].
// [상세 설명 2-3문장].
//
// 매개변수:
//   - param: [설명, 제약조건]
//
// 반환값:
//   - [타입]: [설명]
```

#### 에러를 반환하는 함수
```go
// FunctionName [동사] [명사구].
// [상세 설명 및 목적].
//
// Parameters:
//   - param1: [설명]
//   - param2: [설명]
//
// Returns:
//   - [타입]: [성공 시 반환값]
//   - error: [에러 설명]
//
// Errors:
//   - ErrXXX: when [조건]
//   - ErrYYY: when [조건]
//
// Example:
//   result, err := FunctionName(arg1, arg2)
//   if err != nil {
//       // handle error
//   }
//
// [동일 내용 한글]
```

### 📊 진행 상황 업데이트 방법
```markdown
1. 파일 완료 시:
   - [ ] → [x] 변경
   
2. 패키지 완료 시:
   - Progress Tracking 테이블 업데이트
   - BILINGUAL_AUDIT.md 업데이트
   
3. 세션 종료 시:
   - 현재 작업 파일 [-] 상태 유지
   - "현재 작업 상태" 섹션 업데이트
   - "Next Steps" 업데이트
```

### 🔧 자주 사용하는 명령어
```bash
# 특정 패키지 테스트
go test ./sliceutil -v

# 전체 테스트
go test ./...

# 레이스 검사와 함께 테스트
go test -race ./...

# 커버리지 확인
go test -cover ./sliceutil

# 코드 포맷팅
go fmt ./...

# Vet 실행
go vet ./...
```

### 💬 GitHub Copilot에게 요청하는 방법
```
[좋은 요청 예시]

"sliceutil/slice.go 파일의 Map 함수에 대해 충분히 자세하고 
매우 친절한 영문/한글 병기 주석을 작성해주세요. 
다음을 포함해주세요:
- 함수의 목적과 사용 시나리오
- 모든 파라미터의 의미와 제약 조건
- 반환값 설명
- 시간/공간 복잡도
- nil 슬라이스 처리 방법
- 사용 예제
- thread-safety 여부"
```

### 📖 참고 문서
- `docs/BILINGUAL_AUDIT.md` - 주석 감사 결과
- `docs/CHANGELOG/CHANGELOG-specials.md` - 변경 이력
- `README.md` - 각 패키지 README
- [Effective Go](https://go.dev/doc/effective_go) - Go 주석 가이드

---

## Appendix (부록)

### A. 패키지별 특수 고려사항

**websvrutil**:
- HTTP 핸들러 체인의 실행 순서
- 컨텍스트 값의 생명주기
- 미들웨어 작성 패턴

**database/mysql, database/redis**:
- 연결 풀 관리 방법
- 트랜잭션 사용 패턴
- 에러 재시도 전략

**sliceutil, maputil**:
- 제네릭 타입 설명
- 성능 특성 (O-notation)
- 메모리 할당 패턴

**timeutil**:
- 시간대 처리
- DST 고려사항
- time.Time의 zero value 처리

**fileutil**:
- 파일 권한 (Unix vs Windows)
- 심볼릭 링크 처리
- 대용량 파일 처리

### B. 용어집 (Glossary)

**영문 → 한글**:
- Thread-safe → Thread-safe (스레드 안전)
- Goroutine-safe → Goroutine-safe (고루틴 안전)
- Immutable → 불변
- Mutable → 가변
- Time complexity → 시간 복잡도
- Space complexity → 공간 복잡도
- Edge case → 엣지 케이스 / 경계 조건
- Corner case → 코너 케이스 / 특수 상황

### C. 문서 버전 관리

| 버전 | 날짜 | 변경 내용 | 작성자 |
|------|------|----------|--------|
| 1.0.0 | 2025-10-17 | 초기 마스터 체크리스트 생성 | AI |
|  |  |  |  |

---

## 📞 문의 및 이슈

문서나 작업 진행에 대한 질문이나 이슈가 있으면 다음과 같이 기록:

1. GitHub Issues 생성
2. 또는 이 문서의 "발견된 이슈 및 특이사항" 섹션에 기록

---

**마지막 업데이트**: 2025-10-17  
**문서 상태**: 활성화 (Active)  
**다음 리뷰 예정일**: [작업 시작 시 설정]

---
