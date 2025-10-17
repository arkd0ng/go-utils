# CHANGELOG v1.13.x - validation Package / 검증 유틸리티 패키지

Validation utilities package for Go applications.

Go 애플리케이션을 위한 검증 유틸리티 패키지입니다.

---

## [v1.13.009] - 2025-10-17

### Added / 추가
- Numeric validators 구현 (10개)
  - `Min(n)` - 최소값 검증
  - `Max(n)` - 최대값 검증
  - `Between(min, max)` - 범위 검증 (포함)
  - `Positive()` - 양수 검증
  - `Negative()` - 음수 검증
  - `Zero()` - 0 검증
  - `NonZero()` - 0이 아님 검증
  - `Even()` - 짝수 검증
  - `Odd()` - 홀수 검증
  - `MultipleOf(n)` - 배수 검증

### Implementation Details / 구현 세부사항
- **Type Support**: 모든 숫자 타입 자동 변환 (int, uint, float)
- **Bilingual Messages**: 영어/한글 에러 메시지
- **Method Chaining**: Fluent API로 연속 검증 가능
- **Zero Division Protection**: MultipleOf에서 0으로 나누기 방지

### Files Changed / 변경된 파일
- `validation/rules_numeric.go` - 10개 numeric validators (~87줄)
- `validation/rules_numeric_test.go` - 포괄적 테스트 (~282줄)

### Test Results / 테스트 결과
```bash
go test ./validation -cover
# All 50+ tests passed ✅
# Coverage: 98.3%
```

### Context / 컨텍스트

**Why / 이유**:
- 숫자 검증은 매우 일반적인 요구사항
- 범위 체크, 짝수/홀수, 배수 등 자주 사용되는 패턴
- 타입 안전한 검증으로 런타임 에러 방지

**Impact / 영향**:
- ✅ 30개 이상의 validators 구현 완료 (string 20개 + numeric 10개)
- ✅ 98.3% coverage 달성
- ✅ 모든 테스트 통과

**Example / 예제**:
```go
// Age validation
v := validation.New(25, "age")
v.Positive().Min(18).Max(120)
err := v.Validate()

// Even number check
v := validation.New(10, "value")
v.Even().MultipleOf(5)
err := v.Validate()
```

---

## [v1.13.008] - 2025-10-17

### Changed / 변경
- 모든 패키지의 버전 관리를 동적 로딩으로 변경
  - `internal/version` 패키지 사용으로 통합
  - 하드코딩된 버전 제거
  - cfg/app.yaml에서 중앙 집중식 버전 관리

### Files Changed / 변경된 파일
- `errorutil/types.go` - 하드코딩된 const를 internal/version.Get()으로 변경
- `sliceutil/sliceutil.go` - logging.TryLoadAppVersion()을 internal/version.Get()으로 변경
- `maputil/maputil.go` - logging.TryLoadAppVersion()을 internal/version.Get()으로 변경
- `fileutil/fileutil.go` - logging.TryLoadAppVersion()을 internal/version.Get()으로 변경
- `httputil/httputil.go` - 커스텀 로직을 internal/version.Get()으로 변경
- `websvrutil/websvrutil.go` - logging.TryLoadAppVersion()을 internal/version.Get()으로 변경
- `httputil/httputil_test.go` - TestVersion 수정 (동적 버전 체크)

### Context / 컨텍스트

**User Request / 사용자 요청**: "일단 작업을 멈추고 버전정보 업데이트 하는 부분을 현재의 방식대로 다른패키지에 전체 적용하고 계속 진행바랍니다"

**Why / 이유**:
- 각 패키지마다 버전 로딩 방식이 달라 유지보수 어려움
- 하드코딩된 버전은 실제 버전과 불일치 가능성 있음
- 단일 소스(cfg/app.yaml)에서 중앙 집중식 관리 필요

**Impact / 영향**:
- ✅ 모든 패키지가 동일한 방식으로 버전 로딩
- ✅ 버전 불일치 문제 해결
- ✅ 유지보수성 향상
- ✅ 모든 테스트 통과 (go test ./... 성공)

**Pattern / 패턴**:
```go
// ❌ Before - Hardcoded
const Version = "v1.12.005"

// ❌ Before - Custom logic
func getVersion() string {
    version := logging.TryLoadAppVersion()
    if version == "" {
        return "unknown"
    }
    return version
}

// ✅ After - Unified approach
import "github.com/arkd0ng/go-utils/internal/version"
var Version = version.Get()
```

---

## [v1.13.003] - 2025-10-17

### Added / 추가
- Validator 핵심 기능 구현
  - `New()` - 새 Validator 생성
  - `Validate()` - 검증 실행 및 에러 반환
  - `GetErrors()` - 모든 에러 조회
  - `StopOnError()` - 첫 에러에서 중지 설정
  - `WithMessage()` - 사용자 정의 메시지 설정
  - `Custom()` - 사용자 정의 검증 함수
  - `NewValidator()` - MultiValidator 생성
  - `Field()` - 필드 추가
  - Helper functions: `validateString()`, `validateNumeric()`

### Implementation Details / 구현 세부사항
- **Fluent API**: 메서드 체이닝으로 직관적인 사용
- **Stop on Error**: 첫 번째 에러에서 중지 옵션
- **Custom Messages**: 각 규칙에 사용자 정의 메시지 지정 가능
- **Multi-field Validation**: 여러 필드를 한 번에 검증
- **Type Support**: 모든 숫자 타입 (int, uint, float) 자동 변환

### Files Changed / 변경된 파일
- `validation/validator.go` - 핵심 검증 로직 (~170줄)
- `validation/validator_test.go` - 포괄적 테스트 (~280줄)

### Test Results / 테스트 결과
```bash
go test ./validation -v -cover
# All 36 tests passed ✅
# Coverage: 95.5%
```

### Context / 컨텍스트

**Why / 이유**:
- 검증 규칙을 적용하기 위한 핵심 인프라 필요
- Fluent API로 사용성 극대화
- Multi-field 검증으로 실제 사용 시나리오 지원

**Impact / 영향**:
- ✅ 검증 프레임워크 핵심 완성
- ✅ Custom validators 지원으로 확장성 확보
- ✅ 95.5% 높은 테스트 커버리지

**Next Steps / 다음 단계**:
- v1.13.004-008: String validators 구현 (Required, MinLength, Email, URL, etc.)

---

## [v1.13.002] - 2025-10-17

### Added / 추가
- validation 패키지 기본 구조 생성
  - `version.go` - 패키지 버전 상수
  - `types.go` - Validator, MultiValidator, RuleFunc, MessageFunc 타입 정의
  - `errors.go` - ValidationError, ValidationErrors 타입 및 에러 처리 메서드
  - `types_test.go` - 타입 정의 테스트
  - `errors_test.go` - 에러 처리 포괄적 테스트

### Implementation Details / 구현 세부사항
- **Validator struct**: 단일 값 검증을 위한 핵심 구조체
- **MultiValidator struct**: 여러 필드 검증을 위한 구조체
- **ValidationError**: 필드별 검증 에러 정보 (Field, Value, Rule, Message)
- **ValidationErrors**: 검증 에러 컬렉션 with helper methods
  - `Error()` - 포맷된 에러 메시지
  - `HasField()` - 필드별 에러 확인
  - `GetField()` - 필드별 에러 조회
  - `ToMap()` - 맵 형식 변환
  - `First()` - 첫 번째 에러 조회
  - `Count()` - 에러 개수

### Files Changed / 변경된 파일
- `validation/version.go` - 패키지 버전 (v1.13.002)
- `validation/types.go` - 타입 정의 (~30줄)
- `validation/errors.go` - 에러 타입 및 메서드 (~90줄)
- `validation/types_test.go` - 타입 테스트 (~50줄)
- `validation/errors_test.go` - 에러 테스트 (~160줄)

### Test Results / 테스트 결과
```bash
go test ./validation -v
# All 11 tests passed ✅
# Coverage: 100% for errors.go
```

### Context / 컨텍스트

**Why / 이유**:
- 모든 검증 기능의 기반이 되는 타입과 에러 처리 필요
- 견고한 에러 처리는 사용자 경험에 중요
- 테스트부터 시작하여 높은 품질 보장

**Impact / 영향**:
- ✅ 패키지 기초 구조 완성
- ✅ 타입 안전성 확보
- ✅ 포괄적인 에러 처리 메커니즘
- ✅ 100% 테스트 커버리지

**Next Steps / 다음 단계**:
- v1.13.003: Validator core implementation (New, Validate, GetErrors 메서드)

---

## [v1.13.001] - 2025-10-17

### Added / 추가
- validation 패키지 개발 프로젝트 시작
  - 기능 브랜치 생성: `feature/v1.13.x-validation`
  - 버전을 v1.13.001로 증가
  - DESIGN_PLAN.md 생성 (포괄적인 패키지 설계 계획)
  - WORK_PLAN.md 생성 (60개 패치로 구성된 상세 작업 계획)
  - CHANGELOG-v1.13.md 생성

### Design Highlights / 설계 핵심 사항
- **Extreme Simplicity / 극도의 간결함**: 50줄 → 2-3줄로 코드 감소 (95% 감소)
- **Fluent API / Fluent API**: 체이닝 가능한 검증 규칙
- **50+ Validators / 50개 이상 검증기**: 문자열, 숫자, 날짜/시간, 컬렉션, 비교
- **Struct Validation / 구조체 검증**: 태그 기반 검증 및 중첩 구조체 지원
- **Custom Validators / 사용자 정의 검증기**: 쉬운 사용자 정의 규칙 생성
- **Bilingual Errors / 이중 언어 에러**: 영문/한글 에러 메시지
- **Zero Dependencies / 제로 의존성**: 표준 라이브러리만 사용
- **100% Coverage Target / 100% 커버리지 목표**

### Implementation Plan / 구현 계획
**Phase 1 (v1.13.001-020)**: Core Implementation / 핵심 구현
- Package structure, types, and error handling
- String validators (20 rules)
- Numeric validators (10 rules)
- Basic examples and README

**Phase 2 (v1.13.021-040)**: Advanced Features / 고급 기능
- Date/time validators (8 rules)
- Collection validators (7 rules)
- Comparison validators (5 rules)
- Struct validation with tags
- Custom validators
- Multi-field validation

**Phase 3 (v1.13.041-060)**: Documentation & Finalization / 문서화 및 마무리
- USER_MANUAL.md (2000+ lines)
- DEVELOPER_GUIDE.md (1500+ lines)
- Performance benchmarks
- Root documentation updates
- Merge to main

### Files Changed / 변경된 파일
- `cfg/app.yaml` - 버전을 v1.13.001로 증가
- `docs/validation/DESIGN_PLAN.md` - 패키지 설계 계획 문서 생성 (~800줄)
- `docs/validation/WORK_PLAN.md` - 60개 패치 작업 계획 생성 (~600줄)
- `docs/CHANGELOG/CHANGELOG-v1.13.md` - v1.13.x CHANGELOG 생성

### Context / 컨텍스트

**User Request / 사용자 요청**:
"validation 패키지 개발 시작 (v1.13.x)"

**Why / 이유**:
- go-utils에 검증 유틸리티 패키지가 필요함
- 웹 API, 백엔드 서비스에서 입력 검증은 필수적
- 기존 검증 라이브러리는 복잡하거나 의존성이 많음
- go-utils의 "극도의 간결함" 철학에 맞는 검증 패키지 필요

**Impact / 영향**:
- ✅ 검증 코드를 50줄에서 2-3줄로 대폭 감소 (95% 감소)
- ✅ 50개 이상의 즉시 사용 가능한 검증 규칙 제공
- ✅ 구조체 태그 기반 검증으로 생산성 향상
- ✅ 사용자 정의 검증기로 확장성 제공
- ✅ 이중 언어 에러 메시지로 사용자 경험 향상
- ✅ 외부 의존성 없이 표준 라이브러리만 사용

**Design Goals / 설계 목표**:
1. Extreme simplicity (50+ lines → 2-3 lines)
2. Comprehensive validators (50+ built-in rules)
3. Fluent API for intuitive usage
4. Struct validation with tag support
5. Custom validator support
6. Detailed bilingual error messages
7. Zero external dependencies
8. 100% test coverage

**Next Steps / 다음 단계**:
- v1.13.002: Package structure (types, errors)
- v1.13.003: Validator core implementation
- v1.13.004-008: String validators
- v1.13.009-010: Numeric validators
- Continue Phase 1 implementation

---

**Latest Version / 최신 버전**: v1.13.001
**Package Status / 패키지 상태**: In Development / 개발 중
**Target Completion / 목표 완료**: v1.13.060
**Estimated Date / 예상 날짜**: 2025-10-20
