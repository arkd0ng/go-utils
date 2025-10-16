# errorutil Package Work Plan / errorutil 패키지 작업 계획서

Development work plan for errorutil package with detailed task breakdown.

errorutil 패키지의 상세한 작업 분류를 포함한 개발 작업 계획서입니다.

**Version / 버전**: v1.12.x  
**Status / 상태**: In Progress / 진행 중  
**Last Updated / 최종 업데이트**: 2025-10-16

---

## Table of Contents / 목차

- [Overview / 개요](#overview--개요)
- [Development Phases / 개발 단계](#development-phases--개발-단계)
- [Phase 1: Core Types / 핵심 타입](#phase-1-core-types--핵심-타입)
- [Phase 2: Error Creation / 에러 생성](#phase-2-error-creation--에러-생성)
- [Phase 3: Error Wrapping / 에러 래핑](#phase-3-error-wrapping--에러-래핑)
- [Phase 4: Error Inspection / 에러 검사](#phase-4-error-inspection--에러-검사)
- [Phase 5: Error Classification / 에러 분류](#phase-5-error-classification--에러-분류)
- [Phase 6: Error Formatting / 에러 포매팅](#phase-6-error-formatting--에러-포매팅)
- [Phase 7: Stack Traces / 스택 트레이스](#phase-7-stack-traces--스택-트레이스)
- [Phase 8: Context Errors / 컨텍스트 에러](#phase-8-context-errors--컨텍스트-에러)
- [Phase 9: Error Assertions / 에러 단언](#phase-9-error-assertions--에러-단언)
- [Phase 10: Documentation / 문서화](#phase-10-documentation--문서화)
- [Phase 11: Examples / 예제](#phase-11-examples--예제)
- [Phase 12: Testing & Polish / 테스트 및 마무리](#phase-12-testing--polish--테스트-및-마무리)
- [Task Tracking / 작업 추적](#task-tracking--작업-추적)

---

## Overview / 개요

### Work Principles / 작업 원칙

**🚨 CRITICAL RULES / 핵심 규칙:**
- One function/feature = One patch version / 함수/기능 하나 = 패치 버전 하나
- Do NOT assign patch numbers to tasks / 작업에 패치 번호 할당 금지
- Version numbers assigned during actual work / 버전 번호는 실제 작업 중 할당
- Use todo.md for task tracking / todo.md로 작업 추적
- Follow standard work cycle for EVERY task / 모든 작업에 표준 작업 사이클 적용

### Standard Work Cycle / 표준 작업 사이클

```
For EACH task / 각 작업마다:
1. Bump version in cfg/app.yaml / cfg/app.yaml에서 버전 증가
2. Implement function + tests + example / 함수 + 테스트 + 예제 구현
3. Run tests (go test ./...) / 테스트 실행
4. Update CHANGELOG / CHANGELOG 업데이트
5. Commit & push / 커밋 및 푸시
```

### Development Strategy / 개발 전략

- **Bottom-up approach / 상향식 접근**: Start with basic types, build up to complex features / 기본 타입부터 시작해서 복잡한 기능으로 구축
- **Test-driven / 테스트 주도**: Write tests alongside implementation / 구현과 함께 테스트 작성
- **Incremental / 점진적**: Each task adds working, tested functionality / 각 작업이 작동하고 테스트된 기능 추가
- **Independent tasks / 독립적 작업**: Tasks can be done in flexible order within phases / 단계 내에서 유연한 순서로 작업 가능

---

## Development Phases / 개발 단계

### Phase 1: Core Types / 핵심 타입
**Goal / 목표**: Define fundamental error types and interfaces  
기본 에러 타입 및 인터페이스 정의

**Files / 파일**: `types.go`, `types_test.go`

**Tasks / 작업:**
1. ⏳ Define error interfaces (Unwrapper, Coder, StackTracer, Contexter)  
   에러 인터페이스 정의
2. ⏳ Implement wrappedError type  
   wrappedError 타입 구현
3. ⏳ Implement codedError type  
   codedError 타입 구현
4. ⏳ Implement stackError type  
   stackError 타입 구현
5. ⏳ Implement contextError type  
   contextError 타입 구현
6. ⏳ Implement compositeError type  
   compositeError 타입 구현

**Completion Criteria / 완료 기준:**
- ✅ All error types implement error interface / 모든 에러 타입이 error 인터페이스 구현
- ✅ All error types have Unwrap() method / 모든 에러 타입이 Unwrap() 메서드 보유
- ✅ Tests for each type / 각 타입에 대한 테스트
- ✅ 80%+ test coverage / 80% 이상 테스트 커버리지

---

### Phase 2: Error Creation / 에러 생성
**Goal / 목표**: Implement basic error creation functions  
기본 에러 생성 함수 구현

**Files / 파일**: `create.go`, `create_test.go`

**Tasks / 작업:**
1. ⏳ Implement New() function  
   New() 함수 구현
2. ⏳ Implement Newf() function  
   Newf() 함수 구현
3. ⏳ Implement WithCode() function  
   WithCode() 함수 구현
4. ⏳ Implement WithNumericCode() function  
   WithNumericCode() 함수 구현

**Completion Criteria / 완료 기준:**
- ✅ All creation functions working / 모든 생성 함수 작동
- ✅ Each function has 5+ test cases / 각 함수에 5개 이상 테스트 케이스
- ✅ Example code for each function / 각 함수에 대한 예제 코드
- ✅ Error messages follow guidelines / 에러 메시지가 가이드라인 준수

---

### Phase 3: Error Wrapping / 에러 래핑
**Goal / 목표**: Implement error wrapping with context preservation  
컨텍스트 보존과 함께 에러 래핑 구현

**Files / 파일**: `wrap.go`, `wrap_test.go`

**Tasks / 작업:**
1. ⏳ Implement Wrap() function  
   Wrap() 함수 구현
2. ⏳ Implement Wrapf() function  
   Wrapf() 함수 구현
3. ⏳ Implement WrapWithCode() function  
   WrapWithCode() 함수 구현
4. ⏳ Implement WrapMany() function (wrap multiple errors)  
   WrapMany() 함수 구현 (다중 에러 래핑)

**Completion Criteria / 완료 기준:**
- ✅ Wrapping preserves original error / 래핑이 원본 에러 보존
- ✅ Error chain traversable / 에러 체인 순회 가능
- ✅ Compatible with errors.Is() and errors.As() / errors.Is() 및 errors.As()와 호환
- ✅ Tests for nested wrapping / 중첩 래핑 테스트

---

### Phase 4: Error Inspection / 에러 검사
**Goal / 목표**: Implement functions to inspect and extract error information  
에러 정보를 검사하고 추출하는 함수 구현

**Files / 파일**: `inspect.go`, `inspect_test.go`

**Tasks / 작업:**
1. ⏳ Implement Unwrap() function (if not using standard)  
   Unwrap() 함수 구현 (표준 미사용 시)
2. ⏳ Implement UnwrapAll() function  
   UnwrapAll() 함수 구현
3. ⏳ Implement Root() function  
   Root() 함수 구현
4. ⏳ Implement HasCode() function  
   HasCode() 함수 구현
5. ⏳ Implement GetCode() function  
   GetCode() 함수 구현
6. ⏳ Implement GetNumericCode() function  
   GetNumericCode() 함수 구현
7. ⏳ Implement Contains() function  
   Contains() 함수 구현

**Completion Criteria / 완료 기준:**
- ✅ All inspection functions work with error chains / 모든 검사 함수가 에러 체인과 작동
- ✅ Handle nil errors gracefully / nil 에러를 우아하게 처리
- ✅ Tests for deep error chains / 깊은 에러 체인 테스트
- ✅ Performance tests for chain traversal / 체인 순회 성능 테스트

---

### Phase 5: Error Classification / 에러 분류
**Goal / 목표**: Define standard error categories and classification functions  
표준 에러 카테고리 및 분류 함수 정의

**Files / 파일**: `classify.go`, `classify_test.go`, `errors.go`

**Tasks / 작업:**
1. ⏳ Define sentinel errors (ErrValidation, ErrNotFound, etc.)  
   센티널 에러 정의
2. ⏳ Implement IsValidation() function  
   IsValidation() 함수 구현
3. ⏳ Implement IsNotFound() function  
   IsNotFound() 함수 구현
4. ⏳ Implement IsPermission() function  
   IsPermission() 함수 구현
5. ⏳ Implement IsNetwork() function  
   IsNetwork() 함수 구현
6. ⏳ Implement IsTimeout() function  
   IsTimeout() 함수 구현
7. ⏳ Implement IsDatabase() function  
   IsDatabase() 함수 구현
8. ⏳ Implement IsInternal() function  
   IsInternal() 함수 구현

**Completion Criteria / 완료 기준:**
- ✅ All sentinel errors defined / 모든 센티널 에러 정의됨
- ✅ Classification works with wrapped errors / 분류가 래핑된 에러와 작동
- ✅ Tests for each classification function / 각 분류 함수에 대한 테스트
- ✅ Examples for common use cases / 일반적인 사용 사례 예제

---

### Phase 6: Error Formatting / 에러 포매팅
**Goal / 목표**: Implement flexible error formatting and output  
유연한 에러 포매팅 및 출력 구현

**Files / 파일**: `format.go`, `format_test.go`

**Tasks / 작업:**
1. ⏳ Implement Format() function (basic)  
   Format() 함수 구현 (기본)
2. ⏳ Implement Format() with verbose mode  
   상세 모드가 있는 Format() 구현
3. ⏳ Implement FormatChain() function  
   FormatChain() 함수 구현
4. ⏳ Implement ToJSON() function  
   ToJSON() 함수 구현
5. ⏳ Implement ToMap() function  
   ToMap() 함수 구현

**Completion Criteria / 완료 기준:**
- ✅ Multiple output formats supported / 다중 출력 형식 지원
- ✅ JSON output valid and parseable / JSON 출력이 유효하고 파싱 가능
- ✅ Format preserves all error information / 포맷이 모든 에러 정보 보존
- ✅ Tests for all format variations / 모든 포맷 변형에 대한 테스트

---

### Phase 7: Stack Traces / 스택 트레이스
**Goal / 목표**: Implement stack trace capture and formatting  
스택 트레이스 캡처 및 포매팅 구현

**Files / 파일**: `stack.go`, `stack_test.go`

**Tasks / 작업:**
1. ⏳ Implement Frame type and methods  
   Frame 타입 및 메서드 구현
2. ⏳ Implement stack capture logic  
   스택 캡처 로직 구현
3. ⏳ Implement WithStack() function  
   WithStack() 함수 구현
4. ⏳ Implement WrapWithStack() function  
   WrapWithStack() 함수 구현
5. ⏳ Implement GetStack() function  
   GetStack() 함수 구현
6. ⏳ Implement FormatWithStack() function  
   FormatWithStack() 함수 구현
7. ⏳ Optimize stack capture performance  
   스택 캡처 성능 최적화

**Completion Criteria / 완료 기준:**
- ✅ Stack traces captured correctly / 스택 트레이스가 올바르게 캡처됨
- ✅ Configurable stack depth / 설정 가능한 스택 깊이
- ✅ Skip internal frames / 내부 프레임 건너뛰기
- ✅ Performance benchmarks / 성능 벤치마크
- ✅ Memory usage acceptable / 메모리 사용량 허용 가능

---

### Phase 8: Context Errors / 컨텍스트 에러
**Goal / 목표**: Implement errors with structured contextual data  
구조화된 컨텍스트 데이터를 가진 에러 구현

**Files / 파일**: `context.go` (or in `create.go`), tests

**Tasks / 작업:**
1. ⏳ Implement WithContext() function  
   WithContext() 함수 구현
2. ⏳ Implement WrapWithContext() function  
   WrapWithContext() 함수 구현
3. ⏳ Implement GetContext() function  
   GetContext() 함수 구현
4. ⏳ Implement context merging for wrapped errors  
   래핑된 에러의 컨텍스트 병합 구현
5. ⏳ Add context to JSON/Map output  
   JSON/Map 출력에 컨텍스트 추가

**Completion Criteria / 완료 기준:**
- ✅ Context data preserved through wrapping / 래핑을 통해 컨텍스트 데이터 보존
- ✅ Context accessible at any level / 모든 레벨에서 컨텍스트 접근 가능
- ✅ Tests with complex context data / 복잡한 컨텍스트 데이터 테스트
- ✅ Thread-safe context handling / 스레드 안전 컨텍스트 처리

---

### Phase 9: Error Assertions / 에러 단언
**Goal / 목표**: Implement error assertion and must-pattern utilities  
에러 단언 및 must-패턴 유틸리티 구현

**Files / 파일**: `assert.go`, `assert_test.go`

**Tasks / 작업:**
1. ⏳ Implement As() function (if extending standard)  
   As() 함수 구현 (표준 확장 시)
2. ⏳ Implement Is() function (if extending standard)  
   Is() 함수 구현 (표준 확장 시)
3. ⏳ Implement Must() function  
   Must() 함수 구현
4. ⏳ Implement MustReturn[T]() generic function  
   MustReturn[T]() 제네릭 함수 구현
5. ⏳ Implement Assert() function  
   Assert() 함수 구현

**Completion Criteria / 완료 기준:**
- ✅ Must functions panic correctly / Must 함수가 올바르게 패닉
- ✅ Assert creates proper errors / Assert가 적절한 에러 생성
- ✅ Examples for initialization patterns / 초기화 패턴 예제
- ✅ Tests for panic recovery / 패닉 복구 테스트

---

### Phase 10: Documentation / 문서화
**Goal / 목표**: Create comprehensive bilingual documentation  
포괄적인 이중 언어 문서 생성

**Files / 파일**: `README.md`, `USER_MANUAL.md`, `DEVELOPER_GUIDE.md`

**Tasks / 작업:**
1. ⏳ Create errorutil/README.md with quick start  
   빠른 시작이 포함된 errorutil/README.md 생성
2. ⏳ Create USER_MANUAL.md with all functions  
   모든 함수가 포함된 USER_MANUAL.md 생성
3. ⏳ Create DEVELOPER_GUIDE.md with architecture  
   아키텍처가 포함된 DEVELOPER_GUIDE.md 생성
4. ⏳ Add godoc comments to all exported functions  
   모든 내보내진 함수에 godoc 주석 추가
5. ⏳ Create API reference table  
   API 참조 테이블 생성
6. ⏳ Add migration guide from standard library  
   표준 라이브러리에서의 마이그레이션 가이드 추가
7. ⏳ Add best practices section  
   모범 사례 섹션 추가

**Completion Criteria / 완료 기준:**
- ✅ All documentation bilingual (English/Korean) / 모든 문서 이중 언어
- ✅ Every function documented with examples / 모든 함수가 예제와 함께 문서화
- ✅ README with clear installation and usage / 명확한 설치 및 사용이 있는 README
- ✅ Architecture diagrams where helpful / 유용한 곳에 아키텍처 다이어그램

---

### Phase 11: Examples / 예제
**Goal / 목표**: Create comprehensive example applications  
포괄적인 예제 애플리케이션 생성

**Directories / 디렉토리**: `examples/errorutil/`

**Tasks / 작업:**
1. ⏳ Create basic/ example (creation and wrapping)  
   basic/ 예제 생성 (생성 및 래핑)
2. ⏳ Create advanced/ example (codes, stack, context)  
   advanced/ 예제 생성 (코드, 스택, 컨텍스트)
3. ⏳ Create http_handler/ example  
   http_handler/ 예제 생성
4. ⏳ Create middleware/ example  
   middleware/ 예제 생성
5. ⏳ Create cli_app/ example  
   cli_app/ 예제 생성
6. ⏳ Add logging integration example  
   로깅 통합 예제 추가

**Completion Criteria / 완료 기준:**
- ✅ All examples runnable / 모든 예제 실행 가능
- ✅ Examples use logging package / 예제가 logging 패키지 사용
- ✅ Real-world scenarios demonstrated / 실제 시나리오 시연
- ✅ Comments explain each step / 주석이 각 단계 설명

---

### Phase 12: Testing & Polish / 테스트 및 마무리
**Goal / 목표**: Achieve high test coverage and production readiness  
높은 테스트 커버리지 및 프로덕션 준비 달성

**Tasks / 작업:**
1. ⏳ Achieve 80%+ overall test coverage  
   전체 80% 이상 테스트 커버리지 달성
2. ⏳ Add benchmark tests for all critical paths  
   모든 중요 경로에 벤치마크 테스트 추가
3. ⏳ Add integration tests  
   통합 테스트 추가
4. ⏳ Performance optimization pass  
   성능 최적화 단계
5. ⏳ Memory leak testing  
   메모리 누수 테스트
6. ⏳ Code review and cleanup  
   코드 리뷰 및 정리
7. ⏳ Final documentation review  
   최종 문서 검토
8. ⏳ Update root CHANGELOG.md  
   루트 CHANGELOG.md 업데이트

**Completion Criteria / 완료 기준:**
- ✅ 80%+ test coverage / 80% 이상 테스트 커버리지
- ✅ All benchmarks within acceptable range / 모든 벤치마크가 허용 범위 내
- ✅ No memory leaks / 메모리 누수 없음
- ✅ All documentation complete and accurate / 모든 문서 완료 및 정확
- ✅ Ready for v1.12.x release / v1.12.x 릴리스 준비 완료

---

## Task Tracking / 작업 추적

### How to Track Progress / 진행 상황 추적 방법

**Use todo.md in root directory / 루트 디렉토리의 todo.md 사용:**

```markdown
# errorutil Development Tasks

## Current Task / 현재 작업
- [ ] Task description with version number when started

## Phase 1: Core Types / 핵심 타입
- [ ] Define error interfaces
- [ ] Implement wrappedError type
- [ ] Implement codedError type
...

## Completed / 완료
- [x] Version bump to v1.12.002 (DESIGN_PLAN.md)
- [x] Version bump to v1.12.003 (WORK_PLAN.md)
```

### Task Status Indicators / 작업 상태 표시

- ⏳ **Not Started / 시작 안 함**: Task planned but not yet begun / 계획되었으나 아직 시작 안 함
- 🔄 **In Progress / 진행 중**: Currently being worked on / 현재 작업 중
- ✅ **Completed / 완료**: Task finished and committed / 작업 완료 및 커밋됨
- ⏸️ **Blocked / 차단**: Waiting on dependency / 의존성 대기 중
- 🔀 **Modified / 수정**: Task changed from original plan / 원래 계획에서 변경됨

### Version Tracking / 버전 추적

**Version numbers assigned during work, not in planning / 버전 번호는 계획이 아닌 작업 중 할당:**

- When starting a task, bump version in cfg/app.yaml / 작업 시작 시 cfg/app.yaml에서 버전 증가
- Record version in todo.md when task starts / 작업 시작 시 todo.md에 버전 기록
- Update CHANGELOG with actual version used / 실제 사용된 버전으로 CHANGELOG 업데이트

---

## Summary / 요약

This work plan provides a structured approach to developing the errorutil package with clear phases and completion criteria. Each task follows the standard work cycle, and version numbers are assigned dynamically during actual development work.

이 작업 계획은 명확한 단계와 완료 기준으로 errorutil 패키지 개발에 대한 구조화된 접근 방식을 제공합니다. 각 작업은 표준 작업 사이클을 따르며, 버전 번호는 실제 개발 작업 중 동적으로 할당됩니다.

**Key Points / 주요 사항:**
- 🎯 12 phases with clear goals / 명확한 목표를 가진 12개 단계
- 📋 60+ individual tasks / 60개 이상의 개별 작업
- 🔄 Flexible task ordering within phases / 단계 내 유연한 작업 순서
- 📝 Track progress in todo.md / todo.md에서 진행 상황 추적
- ✅ Clear completion criteria for each phase / 각 단계에 대한 명확한 완료 기준
- 🚀 Focus on incremental, tested progress / 점진적이고 테스트된 진행에 집중

**Next Steps / 다음 단계:**
1. Create todo.md with initial task list / 초기 작업 목록으로 todo.md 생성
2. Start Phase 1: Core Types / Phase 1 시작: 핵심 타입
3. Follow standard work cycle for each task / 각 작업에 대한 표준 작업 사이클 준수
4. Update todo.md as tasks progress / 작업 진행에 따라 todo.md 업데이트
