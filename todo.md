# errorutil 개발 작업

**패키지**: errorutil (v1.12.x)
**최종 업데이트**: 2025-10-17

---

## 🔄 현재 작업

- [ ] 없음 - 계획 완료, 구현 시작 준비

---

## 📋 Phase 1: 핵심 타입 (Core Types)

- [ ] 에러 인터페이스 정의 (Unwrapper, Coder, StackTracer, Contexter)
- [ ] wrappedError 타입 구현 및 테스트
- [ ] codedError 타입 구현 및 테스트
- [ ] stackError 타입 구현 및 테스트
- [ ] contextError 타입 구현 및 테스트
- [ ] compositeError 타입 구현 및 테스트

---

## 📋 Phase 2: 에러 생성 (Error Creation)

- [ ] New() 함수 구현 및 테스트
- [ ] Newf() 함수 구현 및 테스트
- [ ] WithCode() 함수 구현 및 테스트
- [ ] WithNumericCode() 함수 구현 및 테스트

---

## 📋 Phase 3: 에러 래핑 (Error Wrapping)

- [ ] Wrap() 함수 구현 및 테스트
- [ ] Wrapf() 함수 구현 및 테스트
- [ ] WrapWithCode() 함수 구현 및 테스트
- [ ] WrapMany() 함수 구현 및 테스트

---

## 📋 Phase 4: 에러 검사 (Error Inspection)

- [ ] Unwrap() 함수 구현 및 테스트 (필요시)
- [ ] UnwrapAll() 함수 구현 및 테스트
- [ ] Root() 함수 구현 및 테스트
- [ ] HasCode() 함수 구현 및 테스트
- [ ] GetCode() 함수 구현 및 테스트
- [ ] GetNumericCode() 함수 구현 및 테스트
- [ ] Contains() 함수 구현 및 테스트

---

## 📋 Phase 5: 에러 분류 (Error Classification)

- [ ] 센티널 에러 정의 (ErrValidation, ErrNotFound 등)
- [ ] IsValidation() 함수 구현 및 테스트
- [ ] IsNotFound() 함수 구현 및 테스트
- [ ] IsPermission() 함수 구현 및 테스트
- [ ] IsNetwork() 함수 구현 및 테스트
- [ ] IsTimeout() 함수 구현 및 테스트
- [ ] IsDatabase() 함수 구현 및 테스트
- [ ] IsInternal() 함수 구현 및 테스트

---

## 📋 Phase 6: 에러 포매팅 (Error Formatting)

- [ ] Format() 함수 구현 (기본) 및 테스트
- [ ] Format() 상세 모드 구현 및 테스트
- [ ] FormatChain() 함수 구현 및 테스트
- [ ] ToJSON() 함수 구현 및 테스트
- [ ] ToMap() 함수 구현 및 테스트

---

## 📋 Phase 7: 스택 트레이스 (Stack Traces)

- [ ] Frame 타입 및 메서드 구현
- [ ] 스택 캡처 로직 구현
- [ ] WithStack() 함수 구현 및 테스트
- [ ] WrapWithStack() 함수 구현 및 테스트
- [ ] GetStack() 함수 구현 및 테스트
- [ ] FormatWithStack() 함수 구현 및 테스트
- [ ] 스택 캡처 성능 최적화

---

## 📋 Phase 8: 컨텍스트 에러 (Context Errors)

- [ ] WithContext() 함수 구현 및 테스트
- [ ] WrapWithContext() 함수 구현 및 테스트
- [ ] GetContext() 함수 구현 및 테스트
- [ ] 래핑된 에러의 컨텍스트 병합 구현
- [ ] JSON/Map 출력에 컨텍스트 추가

---

## 📋 Phase 9: 에러 단언 (Error Assertions)

- [ ] As() 함수 구현 및 테스트 (표준 확장시)
- [ ] Is() 함수 구현 및 테스트 (표준 확장시)
- [ ] Must() 함수 구현 및 테스트
- [ ] MustReturn[T]() 제네릭 함수 구현 및 테스트
- [ ] Assert() 함수 구현 및 테스트

---

## 📋 Phase 10: 문서화 (Documentation)

- [ ] errorutil/README.md 생성 (빠른 시작 포함)
- [ ] docs/errorutil/USER_MANUAL.md 생성 (모든 함수 포함)
- [ ] docs/errorutil/DEVELOPER_GUIDE.md 생성 (아키텍처 포함)
- [ ] 모든 내보내진 함수에 godoc 주석 추가
- [ ] API 참조 테이블 생성
- [ ] 표준 라이브러리에서의 마이그레이션 가이드 추가
- [ ] 모범 사례 섹션 추가

---

## 📋 Phase 11: 예제 (Examples)

- [ ] examples/errorutil/basic/ 생성 (생성 및 래핑)
- [ ] examples/errorutil/advanced/ 생성 (코드, 스택, 컨텍스트)
- [ ] examples/errorutil/http_handler/ 생성
- [ ] examples/errorutil/middleware/ 생성
- [ ] examples/errorutil/cli_app/ 생성
- [ ] 로깅 통합 예제 추가

---

## 📋 Phase 12: 테스트 및 마무리 (Testing & Polish)

- [ ] 전체 80% 이상 테스트 커버리지 달성
- [ ] 모든 중요 경로에 벤치마크 테스트 추가
- [ ] 통합 테스트 추가
- [ ] 성능 최적화 단계
- [ ] 메모리 누수 테스트
- [ ] 코드 리뷰 및 정리
- [ ] 최종 문서 검토
- [ ] v1.12.x 요약으로 루트 CHANGELOG.md 업데이트

---

## ✅ 완료

- [x] 패키지 초기화 (v1.12.001)
- [x] 이중 언어 문서화 표준 수립 (v1.12.001)
- [x] CHANGELOG 워크플로우 수립 (v1.12.001)
- [x] DESIGN_PLAN.md 생성 (v1.12.002)
- [x] WORK_PLAN.md 및 todo.md 생성 (v1.12.003)

---

## 📝 메모

### 작업 원칙
- 각 작업 = 함수/기능 하나 = 패치 버전 하나
- 버전 번호는 계획이 아닌 작업 시작 시 할당
- 항상 표준 작업 사이클 준수: 버전 증가 → 작업 → 테스트 → CHANGELOG → 커밋 → 푸시
- 각 작업 완료 후 이 파일 업데이트

### 작업 우선순위
- 생성/래핑 함수 전에 핵심 타입이 완료되어야 함
- 각 단계 내에서 작업은 유연한 순서로 진행 가능
- 테스트 및 문서화는 구현과 함께 진행되어야 함

### 추가 작업
- 새로운 작업이 식별되면 적절한 단계에 추가
- 추가된 작업은 🔀 표시로 마크
- 중요한 변경사항이 발생하면 WORK_PLAN.md 업데이트
