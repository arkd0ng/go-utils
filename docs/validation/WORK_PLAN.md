# validation Package - Work Plan / 작업 계획

Detailed implementation plan for the validation utility package.

검증 유틸리티 패키지의 상세 구현 계획입니다.

**Version / 버전**: v1.13.001
**Created / 생성**: 2025-10-17
**Status / 상태**: In Progress / 진행 중

---

## Overview / 개요

This work plan breaks down the validation package development into manageable tasks, each corresponding to a single patch version increment.

이 작업 계획은 validation 패키지 개발을 관리 가능한 작업으로 나누며, 각 작업은 단일 패치 버전 증가에 해당합니다.

**Total Estimated Tasks / 총 예상 작업**: 60 patches (v1.13.001-060)
**Target Coverage / 목표 커버리지**: 100%
**Target Completion / 목표 완료**: 2025-10-20

---

## Phase 1: Core Implementation / 핵심 구현 (v1.13.001-020)

### ✅ v1.13.001 - Project Setup / 프로젝트 설정

- [x] Create feature branch `feature/v1.13.x-validation`
- [x] Bump version to v1.13.001
- [x] Create DESIGN_PLAN.md
- [x] Create WORK_PLAN.md

**Status / 상태**: ✅ COMPLETED / 완료

---

### v1.13.002 - Package Structure / 패키지 구조

**Tasks / 작업:**
- [ ] Create `validation/` directory
- [ ] Create `version.go` with version constant
- [ ] Create `types.go` with basic type definitions
- [ ] Create `errors.go` with error types
- [ ] Create initial test files

**Files / 파일:**
- `validation/version.go`
- `validation/types.go`
- `validation/errors.go`
- `validation/types_test.go`
- `validation/errors_test.go`

**Deliverables / 결과물:**
```go
// types.go
type Validator struct {
    value       interface{}
    fieldName   string
    errors      []ValidationError
    stopOnError bool
}

type ValidationError struct {
    Field   string
    Value   interface{}
    Rule    string
    Message string
}

type ValidationErrors []ValidationError
```

---

### v1.13.003 - Validator Core / Validator 핵심

**Tasks / 작업:**
- [ ] Implement `Validator` struct
- [ ] Implement `New()` constructor
- [ ] Implement `Validate()` method
- [ ] Implement `GetErrors()` method
- [ ] Implement `StopOnError()` method

**Files / 파일:**
- `validation/validator.go`
- `validation/validator_test.go`

**Deliverables / 결과물:**
```go
func New(value interface{}, fieldName string) *Validator
func (v *Validator) Validate() error
func (v *Validator) GetErrors() []ValidationError
func (v *Validator) StopOnError() *Validator
```

---

### v1.13.004 - Basic String Validators (Part 1) / 기본 문자열 검증기 (Part 1)

**Tasks / 작업:**
- [ ] Implement `Required()` validator
- [ ] Implement `MinLength(n)` validator
- [ ] Implement `MaxLength(n)` validator
- [ ] Implement `Length(n)` validator
- [ ] Add comprehensive tests

**Files / 파일:**
- `validation/rules_string.go`
- `validation/rules_string_test.go`

**Deliverables / 결과물:**
```go
func (v *Validator) Required() *Validator
func (v *Validator) MinLength(n int) *Validator
func (v *Validator) MaxLength(n int) *Validator
func (v *Validator) Length(n int) *Validator
```

---

### v1.13.005 - Basic String Validators (Part 2) / 기본 문자열 검증기 (Part 2)

**Tasks / 작업:**
- [ ] Implement `Email()` validator
- [ ] Implement `URL()` validator
- [ ] Implement `Alpha()` validator
- [ ] Implement `Alphanumeric()` validator
- [ ] Implement `Numeric()` validator

**Files / 파일:**
- `validation/rules_string.go`
- `validation/rules_string_test.go`

**Deliverables / 결과물:**
```go
func (v *Validator) Email() *Validator
func (v *Validator) URL() *Validator
func (v *Validator) Alpha() *Validator
func (v *Validator) Alphanumeric() *Validator
func (v *Validator) Numeric() *Validator
```

---

### v1.13.006 - String Pattern Validators / 문자열 패턴 검증기

**Tasks / 작업:**
- [ ] Implement `StartsWith(prefix)` validator
- [ ] Implement `EndsWith(suffix)` validator
- [ ] Implement `Contains(substring)` validator
- [ ] Implement `Regex(pattern)` validator
- [ ] Add tests

**Files / 파일:**
- `validation/rules_string.go`
- `validation/rules_string_test.go`

**Deliverables / 결과물:**
```go
func (v *Validator) StartsWith(prefix string) *Validator
func (v *Validator) EndsWith(suffix string) *Validator
func (v *Validator) Contains(substring string) *Validator
func (v *Validator) Regex(pattern string) *Validator
```

---

### v1.13.007 - Format Validators / 형식 검증기

**Tasks / 작업:**
- [ ] Implement `UUID()` validator
- [ ] Implement `JSON()` validator
- [ ] Implement `Base64()` validator
- [ ] Implement `Phone()` validator
- [ ] Add tests

**Files / 파일:**
- `validation/rules_string.go`
- `validation/rules_string_test.go`

**Deliverables / 결과물:**
```go
func (v *Validator) UUID() *Validator
func (v *Validator) JSON() *Validator
func (v *Validator) Base64() *Validator
func (v *Validator) Phone() *Validator
```

---

### v1.13.008 - Case Validators / 대소문자 검증기

**Tasks / 작업:**
- [ ] Implement `Lowercase()` validator
- [ ] Implement `Uppercase()` validator
- [ ] Implement `CreditCard()` validator
- [ ] Add tests

**Files / 파일:**
- `validation/rules_string.go`
- `validation/rules_string_test.go`

**Deliverables / 결과물:**
```go
func (v *Validator) Lowercase() *Validator
func (v *Validator) Uppercase() *Validator
func (v *Validator) CreditCard() *Validator
```

---

### v1.13.009 - Numeric Validators (Part 1) / 숫자 검증기 (Part 1)

**Tasks / 작업:**
- [ ] Implement `Min(n)` validator
- [ ] Implement `Max(n)` validator
- [ ] Implement `Between(min, max)` validator
- [ ] Implement `Positive()` validator
- [ ] Implement `Negative()` validator

**Files / 파일:**
- `validation/rules_numeric.go`
- `validation/rules_numeric_test.go`

**Deliverables / 결과물:**
```go
func (v *Validator) Min(n float64) *Validator
func (v *Validator) Max(n float64) *Validator
func (v *Validator) Between(min, max float64) *Validator
func (v *Validator) Positive() *Validator
func (v *Validator) Negative() *Validator
```

---

### v1.13.010 - Numeric Validators (Part 2) / 숫자 검증기 (Part 2)

**Tasks / 작업:**
- [ ] Implement `Zero()` validator
- [ ] Implement `NonZero()` validator
- [ ] Implement `Even()` validator
- [ ] Implement `Odd()` validator
- [ ] Implement `MultipleOf(n)` validator

**Files / 파일:**
- `validation/rules_numeric.go`
- `validation/rules_numeric_test.go`

**Deliverables / 결과물:**
```go
func (v *Validator) Zero() *Validator
func (v *Validator) NonZero() *Validator
func (v *Validator) Even() *Validator
func (v *Validator) Odd() *Validator
func (v *Validator) MultipleOf(n float64) *Validator
```

---

### v1.13.011 - Error Handling Enhancement / 에러 처리 개선

**Tasks / 작업:**
- [ ] Implement `ValidationErrors.Error()` method
- [ ] Implement `ValidationErrors.HasField()` method
- [ ] Implement `ValidationErrors.GetField()` method
- [ ] Implement `ValidationErrors.ToMap()` method
- [ ] Add bilingual error messages

**Files / 파일:**
- `validation/errors.go`
- `validation/errors_test.go`

**Deliverables / 결과물:**
```go
func (ve ValidationErrors) Error() string
func (ve ValidationErrors) HasField(field string) bool
func (ve ValidationErrors) GetField(field string) []ValidationError
func (ve ValidationErrors) ToMap() map[string][]string
```

---

### v1.13.012 - Custom Validators / 사용자 정의 검증기

**Tasks / 작업:**
- [ ] Implement `Custom(fn, message)` validator
- [ ] Implement `WithMessage(message)` method
- [ ] Add tests for custom validators
- [ ] Add examples

**Files / 파일:**
- `validation/validator.go`
- `validation/validator_test.go`

**Deliverables / 결과물:**
```go
func (v *Validator) Custom(fn func(interface{}) bool, message string) *Validator
func (v *Validator) WithMessage(message string) *Validator
```

---

### v1.13.013 - Helper Functions / 도우미 함수

**Tasks / 작업:**
- [ ] Implement `IsEmail(s string) bool`
- [ ] Implement `IsURL(s string) bool`
- [ ] Implement `IsUUID(s string) bool`
- [ ] Implement `IsJSON(s string) bool`
- [ ] Add tests

**Files / 파일:**
- `validation/helpers.go`
- `validation/helpers_test.go`

**Deliverables / 결과물:**
```go
func IsEmail(s string) bool
func IsURL(s string) bool
func IsUUID(s string) bool
func IsJSON(s string) bool
```

---

### v1.13.014 - Examples (Part 1) / 예제 (Part 1)

**Tasks / 작업:**
- [ ] Create `examples/validation/` directory
- [ ] Create `main.go` with logging setup
- [ ] Add example for string validators
- [ ] Add example for numeric validators

**Files / 파일:**
- `examples/validation/main.go`

**Deliverables / 결과물:**
- Working examples with detailed logging
- Banner and section formatting

---

### v1.13.015 - Examples (Part 2) / 예제 (Part 2)

**Tasks / 작업:**
- [ ] Add example for custom validators
- [ ] Add example for error handling
- [ ] Add example for fluent API chaining

**Files / 파일:**
- `examples/validation/main.go`

---

### v1.13.016 - Package README (Part 1) / 패키지 README (Part 1)

**Tasks / 작업:**
- [ ] Create `validation/README.md`
- [ ] Add package overview
- [ ] Add installation instructions
- [ ] Add quick start guide

**Files / 파일:**
- `validation/README.md`

---

### v1.13.017 - Package README (Part 2) / 패키지 README (Part 2)

**Tasks / 작업:**
- [ ] Add API reference table
- [ ] Add all validators documentation
- [ ] Add usage examples
- [ ] Add error handling guide

**Files / 파일:**
- `validation/README.md`

---

### v1.13.018 - Package README (Part 3) / 패키지 README (Part 3)

**Tasks / 작업:**
- [ ] Add before/after comparison
- [ ] Add best practices
- [ ] Add FAQ section
- [ ] Bilingual formatting check

**Files / 파일:**
- `validation/README.md`

---

### v1.13.019 - CHANGELOG Creation / CHANGELOG 생성

**Tasks / 작업:**
- [ ] Create `docs/CHANGELOG/CHANGELOG-v1.13.md`
- [ ] Document all changes from v1.13.001-018
- [ ] Update root `CHANGELOG.md`

**Files / 파일:**
- `docs/CHANGELOG/CHANGELOG-v1.13.md`
- `CHANGELOG.md`

---

### v1.13.020 - Phase 1 Review & Testing / Phase 1 검토 및 테스트

**Tasks / 작업:**
- [ ] Run all tests (`go test ./validation -v`)
- [ ] Check coverage (`go test ./validation -cover`)
- [ ] Run examples
- [ ] Fix any issues found
- [ ] Update documentation

**Target / 목표:**
- ✅ All tests passing
- ✅ 80%+ coverage
- ✅ Examples working

---

## Phase 2: Advanced Features / 고급 기능 (v1.13.021-040)

### v1.13.021 - Date/Time Validators (Part 1) / 날짜/시간 검증기 (Part 1)

**Tasks / 작업:**
- [ ] Implement `After(date)` validator
- [ ] Implement `Before(date)` validator
- [ ] Implement `Between(start, end)` validator
- [ ] Add tests

**Files / 파일:**
- `validation/rules_time.go`
- `validation/rules_time_test.go`

---

### v1.13.022 - Date/Time Validators (Part 2) / 날짜/시간 검증기 (Part 2)

**Tasks / 작업:**
- [ ] Implement `Today()` validator
- [ ] Implement `Past()` validator
- [ ] Implement `Future()` validator
- [ ] Add tests

**Files / 파일:**
- `validation/rules_time.go`
- `validation/rules_time_test.go`

---

### v1.13.023 - Date/Time Validators (Part 3) / 날짜/시간 검증기 (Part 3)

**Tasks / 작업:**
- [ ] Implement `Weekday()` validator
- [ ] Implement `Weekend()` validator
- [ ] Add comprehensive tests

**Files / 파일:**
- `validation/rules_time.go`
- `validation/rules_time_test.go`

---

### v1.13.024 - Collection Validators (Part 1) / 컬렉션 검증기 (Part 1)

**Tasks / 작업:**
- [ ] Implement `MinItems(n)` validator
- [ ] Implement `MaxItems(n)` validator
- [ ] Implement `Unique()` validator
- [ ] Add tests

**Files / 파일:**
- `validation/rules_collection.go`
- `validation/rules_collection_test.go`

---

### v1.13.025 - Collection Validators (Part 2) / 컬렉션 검증기 (Part 2)

**Tasks / 작업:**
- [ ] Implement `In(values)` validator
- [ ] Implement `NotIn(values)` validator
- [ ] Implement `Empty()` validator
- [ ] Add tests

**Files / 파일:**
- `validation/rules_collection.go`
- `validation/rules_collection_test.go`

---

### v1.13.026 - Collection Validators (Part 3) / 컬렉션 검증기 (Part 3)

**Tasks / 작업:**
- [ ] Implement `Each(validator)` validator
- [ ] Add comprehensive tests
- [ ] Add examples

**Files / 파일:**
- `validation/rules_collection.go`
- `validation/rules_collection_test.go`

---

### v1.13.027 - Comparison Validators (Part 1) / 비교 검증기 (Part 1)

**Tasks / 작업:**
- [ ] Implement `Equal(value)` validator
- [ ] Implement `NotEqual(value)` validator
- [ ] Implement `GreaterThan(value)` validator
- [ ] Add tests

**Files / 파일:**
- `validation/rules_comparison.go`
- `validation/rules_comparison_test.go`

---

### v1.13.028 - Comparison Validators (Part 2) / 비교 검증기 (Part 2)

**Tasks / 작업:**
- [ ] Implement `LessThan(value)` validator
- [ ] Implement `OneOf(values)` validator
- [ ] Add comprehensive tests

**Files / 파일:**
- `validation/rules_comparison.go`
- `validation/rules_comparison_test.go`

---

### v1.13.029 - Struct Validation (Part 1) / 구조체 검증 (Part 1)

**Tasks / 작업:**
- [ ] Implement `ValidateStruct(s interface{}) error`
- [ ] Implement tag parsing
- [ ] Support basic validators in tags
- [ ] Add tests

**Files / 파일:**
- `validation/struct.go`
- `validation/struct_test.go`

---

### v1.13.030 - Struct Validation (Part 2) / 구조체 검증 (Part 2)

**Tasks / 작업:**
- [ ] Support all validators in tags
- [ ] Implement nested struct validation
- [ ] Add comprehensive tests
- [ ] Add examples

**Files / 파일:**
- `validation/struct.go`
- `validation/struct_test.go`

---

### v1.13.031 - Struct Validation (Part 3) / 구조체 검증 (Part 3)

**Tasks / 작업:**
- [ ] Implement `optional` tag
- [ ] Implement complex validation scenarios
- [ ] Add edge case tests

**Files / 파일:**
- `validation/struct.go`
- `validation/struct_test.go`

---

### v1.13.032 - Multi-Field Validator / 다중 필드 검증기

**Tasks / 작업:**
- [ ] Implement `NewValidator()` for multiple fields
- [ ] Implement `Field(value, name)` method
- [ ] Add tests
- [ ] Add examples

**Files / 파일:**
- `validation/validator.go`
- `validation/validator_test.go`

**Deliverables / 결과물:**
```go
func NewValidator() *MultiValidator
func (mv *MultiValidator) Field(value interface{}, name string) *Validator
func (mv *MultiValidator) Validate() error
```

---

### v1.13.033 - Options Pattern / 옵션 패턴

**Tasks / 작업:**
- [ ] Create `options.go`
- [ ] Implement `WithLanguage(lang string)` option
- [ ] Implement `WithStopOnError(stop bool)` option
- [ ] Add tests

**Files / 파일:**
- `validation/options.go`
- `validation/options_test.go`

---

### v1.13.034 - Custom Error Messages / 사용자 정의 에러 메시지

**Tasks / 작업:**
- [ ] Implement custom message templates
- [ ] Implement message placeholders
- [ ] Add bilingual message support
- [ ] Add tests

**Files / 파일:**
- `validation/errors.go`
- `validation/errors_test.go`

---

### v1.13.035 - Advanced Examples (Part 1) / 고급 예제 (Part 1)

**Tasks / 작업:**
- [ ] Add struct validation examples
- [ ] Add nested struct examples
- [ ] Add multi-field examples

**Files / 파일:**
- `examples/validation/main.go`

---

### v1.13.036 - Advanced Examples (Part 2) / 고급 예제 (Part 2)

**Tasks / 작업:**
- [ ] Add date/time validation examples
- [ ] Add collection validation examples
- [ ] Add custom validator examples

**Files / 파일:**
- `examples/validation/main.go`

---

### v1.13.037 - Benchmark Tests (Part 1) / 벤치마크 테스트 (Part 1)

**Tasks / 작업:**
- [ ] Create benchmark tests for string validators
- [ ] Create benchmark tests for numeric validators
- [ ] Analyze performance

**Files / 파일:**
- `validation/benchmark_test.go`

---

### v1.13.038 - Benchmark Tests (Part 2) / 벤치마크 테스트 (Part 2)

**Tasks / 작업:**
- [ ] Create benchmark tests for struct validation
- [ ] Create benchmark tests for complex scenarios
- [ ] Document performance results

**Files / 파일:**
- `validation/benchmark_test.go`

---

### v1.13.039 - Phase 2 Testing / Phase 2 테스트

**Tasks / 작업:**
- [ ] Run all tests (`go test ./validation -v`)
- [ ] Check coverage (`go test ./validation -cover`)
- [ ] Run benchmarks
- [ ] Fix any issues

**Target / 목표:**
- ✅ All tests passing
- ✅ 95%+ coverage

---

### v1.13.040 - Phase 2 Documentation Update / Phase 2 문서 업데이트

**Tasks / 작업:**
- [ ] Update `validation/README.md` with new features
- [ ] Update CHANGELOG
- [ ] Review all documentation

**Files / 파일:**
- `validation/README.md`
- `docs/CHANGELOG/CHANGELOG-v1.13.md`

---

## Phase 3: Documentation & Finalization / 문서화 및 마무리 (v1.13.041-060)

### v1.13.041-048 - USER_MANUAL.md

**Tasks / 작업 (8 patches):**
- [ ] v1.13.041: Structure and introduction / 구조 및 소개
- [ ] v1.13.042: String validators section / 문자열 검증기 섹션
- [ ] v1.13.043: Numeric validators section / 숫자 검증기 섹션
- [ ] v1.13.044: Date/time validators section / 날짜/시간 검증기 섹션
- [ ] v1.13.045: Collection validators section / 컬렉션 검증기 섹션
- [ ] v1.13.046: Struct validation section / 구조체 검증 섹션
- [ ] v1.13.047: Advanced patterns section / 고급 패턴 섹션
- [ ] v1.13.048: Examples and best practices / 예제 및 모범 사례

**Target / 목표**: 2000+ lines comprehensive guide / 2000줄 이상 포괄적 가이드

---

### v1.13.049-054 - DEVELOPER_GUIDE.md

**Tasks / 작업 (6 patches):**
- [ ] v1.13.049: Architecture overview / 아키텍처 개요
- [ ] v1.13.050: Core components / 핵심 구성요소
- [ ] v1.13.051: Adding custom validators / 사용자 정의 검증기 추가
- [ ] v1.13.052: Testing guide / 테스트 가이드
- [ ] v1.13.053: Contributing guidelines / 기여 가이드라인
- [ ] v1.13.054: Troubleshooting / 문제 해결

**Target / 목표**: 1500+ lines technical documentation / 1500줄 이상 기술 문서

---

### v1.13.055 - Performance Documentation / 성능 문서

**Tasks / 작업:**
- [ ] Create performance benchmarks document
- [ ] Document optimization strategies
- [ ] Add performance tips

**Files / 파일:**
- `docs/validation/PERFORMANCE_BENCHMARKS.md`

---

### v1.13.056 - Root README Update / 루트 README 업데이트

**Tasks / 작업:**
- [ ] Add validation package to root README.md
- [ ] Add usage examples
- [ ] Update package list

**Files / 파일:**
- `README.md`

---

### v1.13.057 - Final Testing & Coverage / 최종 테스트 및 커버리지

**Tasks / 작업:**
- [ ] Run all tests with race detection
- [ ] Verify 100% coverage
- [ ] Test all examples
- [ ] Fix any remaining issues

**Target / 목표:**
- ✅ 100% test coverage
- ✅ All tests passing
- ✅ All examples working
- ✅ No race conditions

---

### v1.13.058 - Root CHANGELOG Update / 루트 CHANGELOG 업데이트

**Tasks / 작업:**
- [ ] Update `CHANGELOG.md` with v1.13.x summary
- [ ] Finalize `docs/CHANGELOG/CHANGELOG-v1.13.md`
- [ ] Add feature highlights

**Files / 파일:**
- `CHANGELOG.md`
- `docs/CHANGELOG/CHANGELOG-v1.13.md`

---

### v1.13.059 - Pre-Merge Review / 병합 전 검토

**Tasks / 작업:**
- [ ] Final code review
- [ ] Documentation review
- [ ] Test coverage verification
- [ ] Bilingual check

**Checklist / 체크리스트:**
- [ ] All files follow go-utils conventions
- [ ] All documentation is bilingual
- [ ] All tests pass
- [ ] 100% coverage achieved
- [ ] Examples are comprehensive
- [ ] No external dependencies

---

### v1.13.060 - Merge to Main / 메인에 병합

**Tasks / 작업:**
- [ ] Final commit on feature branch
- [ ] Switch to main branch
- [ ] Merge `feature/v1.13.x-validation` to `main`
- [ ] Push to remote
- [ ] Create git tag `v1.13.060`
- [ ] Push tag

**Commands / 명령어:**
```bash
git checkout main
git merge feature/v1.13.x-validation
git push origin main
git tag v1.13.060
git push origin v1.13.060
```

---

## Progress Tracking / 진행 상황 추적

### Current Status / 현재 상태

**Phase / 단계**: Phase 1 - Core Implementation
**Current Version / 현재 버전**: v1.13.001
**Completed Tasks / 완료 작업**: 1 / 60
**Completion Rate / 완료율**: 1.7%

### Milestones / 마일스톤

- [x] **v1.13.001** - Project setup complete / 프로젝트 설정 완료
- [ ] **v1.13.020** - Phase 1 complete (Core Implementation)
- [ ] **v1.13.040** - Phase 2 complete (Advanced Features)
- [ ] **v1.13.060** - Phase 3 complete (Documentation & Merge)

---

## Quality Metrics / 품질 지표

### Target Metrics / 목표 지표

| Metric / 지표 | Target / 목표 | Current / 현재 |
|--------------|--------------|---------------|
| Test Coverage / 테스트 커버리지 | 100% | 0% |
| Documentation / 문서화 | 5000+ lines | 0 lines |
| Examples / 예제 | 20+ | 0 |
| Validators / 검증기 | 50+ | 0 |
| Performance / 성능 | < 100ns/op | - |

---

## Risk Management / 위험 관리

### Potential Risks / 잠재적 위험

1. **Complex Struct Validation / 복잡한 구조체 검증**
   - **Risk / 위험**: Nested structs and reflection complexity
   - **Mitigation / 완화**: Incremental implementation, extensive testing

2. **Performance / 성능**
   - **Risk / 위험**: Slow validation for large datasets
   - **Mitigation / 완화**: Benchmark early, optimize critical paths

3. **Error Message Localization / 에러 메시지 현지화**
   - **Risk / 위험**: Maintaining bilingual messages
   - **Mitigation / 완화**: Message template system

---

## Notes / 참고사항

### Development Guidelines / 개발 가이드라인

1. **Always increment version BEFORE work / 작업 전 항상 버전 증가**
2. **Write tests alongside implementation / 구현과 함께 테스트 작성**
3. **Update CHANGELOG for every commit / 모든 커밋마다 CHANGELOG 업데이트**
4. **Bilingual documentation / 이중 언어 문서화**
5. **Run `go test ./... -v` before commit / 커밋 전 테스트 실행**

### References / 참조

- [DESIGN_PLAN.md](./DESIGN_PLAN.md) - Package design / 패키지 설계
- [DEVELOPMENT_WORKFLOW_GUIDE.md](../DEVELOPMENT_WORKFLOW_GUIDE.md) - Standard workflow / 표준 워크플로우
- [PACKAGE_DEVELOPMENT_GUIDE.md](../PACKAGE_DEVELOPMENT_GUIDE.md) - Package development / 패키지 개발

---

**Document Version / 문서 버전**: v1.13.001
**Last Updated / 최종 업데이트**: 2025-10-17
**Author / 작성자**: go-utils team
**Status / 상태**: Active / 활성
