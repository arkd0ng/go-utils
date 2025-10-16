# errorutil Development Tasks / errorutil 개발 작업

**Package / 패키지**: errorutil (v1.12.x)  
**Last Updated / 최종 업데이트**: 2025-10-16

---

## 🔄 Current Task / 현재 작업

- [ ] None - Planning complete, ready to start implementation  
      없음 - 계획 완료, 구현 시작 준비

---

## 📋 Phase 1: Core Types / 핵심 타입

- [ ] Define error interfaces (Unwrapper, Coder, StackTracer, Contexter)  
      에러 인터페이스 정의
- [ ] Implement wrappedError type with tests  
      테스트와 함께 wrappedError 타입 구현
- [ ] Implement codedError type with tests  
      테스트와 함께 codedError 타입 구현
- [ ] Implement stackError type with tests  
      테스트와 함께 stackError 타입 구현
- [ ] Implement contextError type with tests  
      테스트와 함께 contextError 타입 구현
- [ ] Implement compositeError type with tests  
      테스트와 함께 compositeError 타입 구현

---

## 📋 Phase 2: Error Creation / 에러 생성

- [ ] Implement New() function with tests  
      테스트와 함께 New() 함수 구현
- [ ] Implement Newf() function with tests  
      테스트와 함께 Newf() 함수 구현
- [ ] Implement WithCode() function with tests  
      테스트와 함께 WithCode() 함수 구현
- [ ] Implement WithNumericCode() function with tests  
      테스트와 함께 WithNumericCode() 함수 구현

---

## 📋 Phase 3: Error Wrapping / 에러 래핑

- [ ] Implement Wrap() function with tests  
      테스트와 함께 Wrap() 함수 구현
- [ ] Implement Wrapf() function with tests  
      테스트와 함께 Wrapf() 함수 구현
- [ ] Implement WrapWithCode() function with tests  
      테스트와 함께 WrapWithCode() 함수 구현
- [ ] Implement WrapMany() function with tests  
      테스트와 함께 WrapMany() 함수 구현

---

## 📋 Phase 4: Error Inspection / 에러 검사

- [ ] Implement Unwrap() function with tests (if needed)  
      테스트와 함께 Unwrap() 함수 구현 (필요 시)
- [ ] Implement UnwrapAll() function with tests  
      테스트와 함께 UnwrapAll() 함수 구현
- [ ] Implement Root() function with tests  
      테스트와 함께 Root() 함수 구현
- [ ] Implement HasCode() function with tests  
      테스트와 함께 HasCode() 함수 구현
- [ ] Implement GetCode() function with tests  
      테스트와 함께 GetCode() 함수 구현
- [ ] Implement GetNumericCode() function with tests  
      테스트와 함께 GetNumericCode() 함수 구현
- [ ] Implement Contains() function with tests  
      테스트와 함께 Contains() 함수 구현

---

## 📋 Phase 5: Error Classification / 에러 분류

- [ ] Define sentinel errors (ErrValidation, ErrNotFound, etc.)  
      센티널 에러 정의
- [ ] Implement IsValidation() function with tests  
      테스트와 함께 IsValidation() 함수 구현
- [ ] Implement IsNotFound() function with tests  
      테스트와 함께 IsNotFound() 함수 구현
- [ ] Implement IsPermission() function with tests  
      테스트와 함께 IsPermission() 함수 구현
- [ ] Implement IsNetwork() function with tests  
      테스트와 함께 IsNetwork() 함수 구현
- [ ] Implement IsTimeout() function with tests  
      테스트와 함께 IsTimeout() 함수 구현
- [ ] Implement IsDatabase() function with tests  
      테스트와 함께 IsDatabase() 함수 구현
- [ ] Implement IsInternal() function with tests  
      테스트와 함께 IsInternal() 함수 구현

---

## 📋 Phase 6: Error Formatting / 에러 포매팅

- [ ] Implement Format() function (basic) with tests  
      테스트와 함께 Format() 함수 구현 (기본)
- [ ] Implement Format() with verbose mode with tests  
      테스트와 함께 상세 모드 Format() 구현
- [ ] Implement FormatChain() function with tests  
      테스트와 함께 FormatChain() 함수 구현
- [ ] Implement ToJSON() function with tests  
      테스트와 함께 ToJSON() 함수 구현
- [ ] Implement ToMap() function with tests  
      테스트와 함께 ToMap() 함수 구현

---

## 📋 Phase 7: Stack Traces / 스택 트레이스

- [ ] Implement Frame type and methods  
      Frame 타입 및 메서드 구현
- [ ] Implement stack capture logic  
      스택 캡처 로직 구현
- [ ] Implement WithStack() function with tests  
      테스트와 함께 WithStack() 함수 구현
- [ ] Implement WrapWithStack() function with tests  
      테스트와 함께 WrapWithStack() 함수 구현
- [ ] Implement GetStack() function with tests  
      테스트와 함께 GetStack() 함수 구현
- [ ] Implement FormatWithStack() function with tests  
      테스트와 함께 FormatWithStack() 함수 구현
- [ ] Optimize stack capture performance  
      스택 캡처 성능 최적화

---

## 📋 Phase 8: Context Errors / 컨텍스트 에러

- [ ] Implement WithContext() function with tests  
      테스트와 함께 WithContext() 함수 구현
- [ ] Implement WrapWithContext() function with tests  
      테스트와 함께 WrapWithContext() 함수 구현
- [ ] Implement GetContext() function with tests  
      테스트와 함께 GetContext() 함수 구현
- [ ] Implement context merging for wrapped errors  
      래핑된 에러의 컨텍스트 병합 구현
- [ ] Add context to JSON/Map output  
      JSON/Map 출력에 컨텍스트 추가

---

## 📋 Phase 9: Error Assertions / 에러 단언

- [ ] Implement As() function with tests (if extending standard)  
      테스트와 함께 As() 함수 구현 (표준 확장 시)
- [ ] Implement Is() function with tests (if extending standard)  
      테스트와 함께 Is() 함수 구현 (표준 확장 시)
- [ ] Implement Must() function with tests  
      테스트와 함께 Must() 함수 구현
- [ ] Implement MustReturn[T]() generic function with tests  
      테스트와 함께 MustReturn[T]() 제네릭 함수 구현
- [ ] Implement Assert() function with tests  
      테스트와 함께 Assert() 함수 구현

---

## 📋 Phase 10: Documentation / 문서화

- [ ] Create errorutil/README.md with quick start  
      빠른 시작이 포함된 errorutil/README.md 생성
- [ ] Create docs/errorutil/USER_MANUAL.md with all functions  
      모든 함수가 포함된 USER_MANUAL.md 생성
- [ ] Create docs/errorutil/DEVELOPER_GUIDE.md with architecture  
      아키텍처가 포함된 DEVELOPER_GUIDE.md 생성
- [ ] Add godoc comments to all exported functions  
      모든 내보내진 함수에 godoc 주석 추가
- [ ] Create API reference table  
      API 참조 테이블 생성
- [ ] Add migration guide from standard library  
      표준 라이브러리에서의 마이그레이션 가이드 추가
- [ ] Add best practices section  
      모범 사례 섹션 추가

---

## 📋 Phase 11: Examples / 예제

- [ ] Create examples/errorutil/basic/ (creation and wrapping)  
      basic/ 예제 생성 (생성 및 래핑)
- [ ] Create examples/errorutil/advanced/ (codes, stack, context)  
      advanced/ 예제 생성 (코드, 스택, 컨텍스트)
- [ ] Create examples/errorutil/http_handler/  
      http_handler/ 예제 생성
- [ ] Create examples/errorutil/middleware/  
      middleware/ 예제 생성
- [ ] Create examples/errorutil/cli_app/  
      cli_app/ 예제 생성
- [ ] Add logging integration example  
      로깅 통합 예제 추가

---

## 📋 Phase 12: Testing & Polish / 테스트 및 마무리

- [ ] Achieve 80%+ overall test coverage  
      전체 80% 이상 테스트 커버리지 달성
- [ ] Add benchmark tests for all critical paths  
      모든 중요 경로에 벤치마크 테스트 추가
- [ ] Add integration tests  
      통합 테스트 추가
- [ ] Performance optimization pass  
      성능 최적화 단계
- [ ] Memory leak testing  
      메모리 누수 테스트
- [ ] Code review and cleanup  
      코드 리뷰 및 정리
- [ ] Final documentation review  
      최종 문서 검토
- [ ] Update root CHANGELOG.md with v1.12.x summary  
      v1.12.x 요약으로 루트 CHANGELOG.md 업데이트

---

## ✅ Completed / 완료

- [x] Package initialization (v1.12.001)  
      패키지 초기화
- [x] Bilingual documentation standards established (v1.12.001)  
      이중 언어 문서화 표준 수립
- [x] CHANGELOG workflow established (v1.12.001)  
      CHANGELOG 워크플로우 수립
- [x] Create DESIGN_PLAN.md (v1.12.002)  
      DESIGN_PLAN.md 생성
- [x] Create WORK_PLAN.md and todo.md (v1.12.003)  
      WORK_PLAN.md 및 todo.md 생성

---

## 📝 Notes / 메모

### Work Principles / 작업 원칙
- Each task = One function/feature = One patch version  
  각 작업 = 함수/기능 하나 = 패치 버전 하나
- Version number assigned when work starts, not during planning  
  버전 번호는 계획이 아닌 작업 시작 시 할당
- Always follow standard work cycle: Version Bump → Work → Test → CHANGELOG → Commit → Push  
  항상 표준 작업 사이클 준수: 버전 증가 → 작업 → 테스트 → CHANGELOG → 커밋 → 푸시
- Update this file after each completed task  
  각 작업 완료 후 이 파일 업데이트

### Task Prioritization / 작업 우선순위
- Core types must be completed before creation/wrapping functions  
  생성/래핑 함수 전에 핵심 타입이 완료되어야 함
- Within each phase, tasks can be done in flexible order  
  각 단계 내에서 작업은 유연한 순서로 진행 가능
- Testing and documentation should be done alongside implementation  
  테스트 및 문서화는 구현과 함께 진행되어야 함

### Additional Tasks / 추가 작업
- When new tasks are identified, add them to appropriate phase  
  새로운 작업이 식별되면 적절한 단계에 추가
- Mark added tasks with 🔀 indicator  
  추가된 작업은 🔀 표시로 마크
- Update WORK_PLAN.md if significant changes occur  
  중요한 변경사항이 발생하면 WORK_PLAN.md 업데이트
