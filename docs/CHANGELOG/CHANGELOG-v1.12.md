# CHANGELOG v1.12.x - errorutil Package / 에러 처리 유틸리티 패키지

Error handling utilities package for Go applications.

Go 애플리케이션을 위한 에러 처리 유틸리티 패키지입니다.

---

## [v1.12.006] - 2025-10-17

### Added / 추가
- errorutil 패키지 Phase 2 (Error Creation) 완료
- 6개의 에러 생성 함수 구현:
  - New(message): 기본 에러 생성
  - Newf(format, args...): 포맷된 에러 생성
  - WithCode(code, message): 문자열 코드가 있는 에러 생성
  - WithCodef(code, format, args...): 문자열 코드와 포맷된 메시지
  - WithNumericCode(code, message): 숫자 코드가 있는 에러 생성
  - WithNumericCodef(code, format, args...): 숫자 코드와 포맷된 메시지
- 모든 에러 생성 함수에 대한 포괄적인 테스트 추가

### Changed / 변경
- N/A

### Fixed / 수정
- N/A

### Files Changed / 변경된 파일
- `cfg/app.yaml` - 버전을 v1.12.005에서 v1.12.006으로 증가
- `errorutil/error.go` - 새 파일 생성 (180+ 줄, 에러 생성 함수들)
- `errorutil/error_test.go` - 새 파일 생성 (380+ 줄, 포괄적인 테스트)
- `docs/CHANGELOG/CHANGELOG-v1.12.md` - v1.12.006 항목 추가

### Context / 컨텍스트

**User Request / 사용자 요청**:
Phase 1 완료 후 자동으로 Phase 2로 진행

**Why / 이유**:
- Phase 2는 사용자가 에러를 생성하는 공개 API의 첫 단계
- New()와 Newf()는 errors.New, fmt.Errorf와 유사하지만 unwrapping 지원
- WithCode와 WithNumericCode는 에러 분류 및 API 응답에 필수
- 포맷 변형(Newf, WithCodef, WithNumericCodef)은 동적 메시지 생성 지원

**Implementation Details / 구현 세부사항**:

1. **기본 에러 생성**:
   - New(): 단순 메시지로 wrappedError 반환
   - Newf(): fmt.Sprintf로 포맷된 메시지의 wrappedError 반환

2. **코드가 있는 에러**:
   - WithCode(): 문자열 코드("ERR001", "VALIDATION_ERROR" 등)
   - WithNumericCode(): 숫자 코드(404, 500 등 HTTP 상태 코드)
   - 각각 포맷 변형(WithCodef, WithNumericCodef) 제공

3. **인터페이스 호환성**:
   - 모든 함수는 Phase 1의 타입(wrappedError, codedError, numericCodedError) 반환
   - Coder, NumericCoder 인터페이스 구현
   - Unwrapper 인터페이스 구현 (Go 표준 라이브러리 호환)

**Impact / 영향**:
- 사용자가 간단하게 에러 생성 가능
- 에러 코드를 통한 분류 가능
- 다음 단계(Phase 3: Error Wrapping)의 기초 제공
- 전체 커버리지 98.1%로 목표 80% 초과

### Test Results / 테스트 결과

```
PASS
ok  	github.com/arkd0ng/go-utils/errorutil	0.760s
coverage: 98.1% of statements
```

All 13 test functions passed with 33 subtests:
- TestNew (3 cases)
- TestNewf (4 cases)
- TestWithCode (4 cases)
- TestWithCodef (3 cases)
- TestWithNumericCode (4 cases)
- TestWithNumericCodef (3 cases)
- + Phase 1 tests (7 functions, 14 subtests)

### Next Steps / 다음 단계

Phase 3: Error Wrapping (에러 래핑 함수)
- Wrap() 함수 구현
- Wrapf() 함수 구현
- WrapWithCode() 함수 구현
- WrapWithNumericCode() 함수 구현

---

## [v1.12.005] - 2025-10-17

### Added / 추가
- errorutil 패키지 Phase 1 (Core Types) 완료
- 5개의 핵심 인터페이스 정의 (Unwrapper, Coder, NumericCoder, StackTracer, Contexter)
- Frame 구조체 추가 (스택 트레이스용)
- 6개의 에러 타입 구현:
  - wrappedError: 기본 에러 래핑
  - codedError: 문자열 코드를 가진 에러
  - numericCodedError: 숫자 코드를 가진 에러
  - stackError: 스택 트레이스를 캡처하는 에러
  - contextError: 컨텍스트 데이터를 전달하는 에러
  - compositeError: 모든 기능을 결합한 에러
- 모든 에러 타입에 대한 포괄적인 테스트 추가 (97.8% 커버리지)

### Changed / 변경
- N/A

### Fixed / 수정
- types.go:261-263의 문법 오류 수정 (함수 시그니처와 중괄호 사이의 불필요한 개행 제거)

### Files Changed / 변경된 파일
- `cfg/app.yaml` - 버전을 v1.12.004에서 v1.12.005로 증가
- `errorutil/types.go` - 새 파일 생성 (350+ 줄, 모든 핵심 타입)
- `errorutil/types_test.go` - 새 파일 생성 (450+ 줄, 포괄적인 테스트)
- `docs/CHANGELOG/CHANGELOG-v1.12.md` - v1.12.005 항목 추가

### Context / 컨텍스트

**User Request / 사용자 요청**:
"현재 errorutil패키지 작업중이었습니다. CHANGELOG와 기타 문서들을 확인하고 패키지를 완성해줘"

**Why / 이유**:
- errorutil 패키지는 12개 Phase로 계획된 대규모 작업
- Phase 1(Core Types)은 모든 후속 Phase의 기초가 되는 핵심 구현
- 에러 처리는 Go 애플리케이션의 핵심 기능이며, 표준 라이브러리보다 향상된 기능 제공
- 에러 코드(문자열/숫자), 스택 트레이스, 컨텍스트 데이터 등 다양한 에러 처리 패턴 지원
- Go 1.13+ errors 패키지와 완전히 호환되는 Unwrap 인터페이스 구현

**Implementation Details / 구현 세부사항**:

1. **인터페이스 설계**:
   - Unwrapper: Go 표준 라이브러리 호환 (errors.Is, errors.As 지원)
   - Coder: API 응답 및 에러 분류를 위한 문자열 코드
   - NumericCoder: HTTP 상태 코드 등 숫자 코드
   - StackTracer: 디버깅을 위한 스택 트레이스
   - Contexter: 구조화된 컨텍스트 데이터 전달

2. **불변성 보장**:
   - contextError와 compositeError의 Context() 메서드는 복사본 반환
   - 외부 수정으로부터 내부 데이터 보호

3. **테스트 전략**:
   - 테이블 기반 테스트로 다양한 시나리오 검증
   - cause가 있는/없는 경우 모두 테스트
   - 불변성 테스트 포함
   - 97.8% 커버리지 달성 (목표 80% 초과)

**Impact / 영향**:
- Phase 2-12의 모든 공개 API가 이 핵심 타입들을 기반으로 구축됨
- New(), Wrap(), WithCode() 등 공개 함수들이 이 타입들을 반환
- 사용자는 인터페이스를 통해 에러 특성 검사 가능
- Go 표준 라이브러리와 완벽히 호환되어 기존 코드와 통합 용이
- 다음 단계(Phase 2: Error Creation)로 진행 가능

### Test Results / 테스트 결과

```
PASS
ok  	github.com/arkd0ng/go-utils/errorutil	0.791s
coverage: 97.8% of statements
```

All 7 test functions passed with 14 subtests:
- TestWrappedError (2 cases)
- TestCodedError (2 cases)
- TestNumericCodedError (2 cases)
- TestStackError (2 cases)
- TestContextError (2 cases)
- TestCompositeError (3 cases)
- TestFrame (2 cases)

### Next Steps / 다음 단계

Phase 2: Error Creation (에러 생성 함수)
- New() 함수 구현
- Newf() 함수 구현
- WithCode() 함수 구현
- WithNumericCode() 함수 구현

---

## [v1.12.004] - 2025-10-17

### Added / 추가
- 언어 사용 정책을 CLAUDE.md에 명확히 정의 및 문서화
- 공개 문서(영문/한글 병기)와 비공개 문서(한글 전용) 구분 명시
- 코드 주석, 로그 메시지, Git 커밋 메시지에 대한 이중 언어 규칙 명문화
- 주석을 "매우 상세하고 친절하게" 작성하는 원칙 추가
- 규칙 위반 시 조치 방법 명시

### Changed / 변경
- CLAUDE.md를 영문에서 한글 전용으로 완전히 재작성
- todo.md를 영문/한글 병기에서 한글 전용으로 변환
- CLAUDE.md 내용을 더 간결하고 실용적으로 재구성 (618줄 → 463줄)
- 언어 사용 규칙을 최우선 섹션으로 배치

### Fixed / 수정
- N/A

### Files Changed / 변경된 파일
- `cfg/app.yaml` - 버전을 v1.12.003에서 v1.12.004로 증가
- `CLAUDE.md` - 한글 전용으로 완전히 재작성, 언어 규칙 섹션 추가
- `todo.md` - 한글 전용으로 변환
- `docs/CHANGELOG/CHANGELOG-v1.12.md` - v1.12.004 항목 추가

### Context / 컨텍스트

**User Request / 사용자 요청**:
"한글로 설명해줘.. 앞으로도. 문서는 영문/한글을 병기하도록 합니다. 코드내 주석문/로그도 영문/한글을 병기하도록 합니다. 또한 주석문은 매우 자세하고 친절하게 작성합니다. 공개하지 않는 문서(CLAUDE.md등)는 한글로 작성합니다. 이는 매우 중요한 규칙입니다. 이 규칙을 매번 설명하지 않도록 CLAUDE.md등에 명기하여 다시 지시하지 않아도 계속 따르도록 합니다. 이 규칙에 위배된 내용이 있으면 전체 코드와 문서를 확인하고 수정 및 보완합니다."

**Why / 이유**:
- 언어 사용 규칙을 명확히 하여 일관성 유지
- AI 어시스턴트가 매번 지시 없이도 규칙을 따르도록 문서화
- 공개/비공개 문서의 언어 정책을 명확히 구분
- 코드의 이해도를 높이기 위해 주석을 매우 상세하게 작성하는 원칙 수립
- 이중 언어 지원으로 국내외 개발자 모두가 쉽게 사용 가능

**Impact / 영향**:
- CLAUDE.md가 명확한 언어 정책 가이드 역할 수행
- 향후 모든 작업에서 자동으로 언어 규칙 준수
- 비공개 문서(CLAUDE.md, todo.md)는 한글로 작성하여 가독성 향상
- 공개 문서는 영문/한글 병기로 유지하여 접근성 확보
- 주석의 상세함 기준이 명확해져 코드 품질 향상
- 일관된 언어 정책으로 프로젝트 전체의 통일성 증대

### Verification / 검증

레포지토리 전체 파일 검증 결과:
- ✅ 코드 주석: random/string.go, logging/logger.go, stringutil/stringutil.go, maputil/maputil.go 모두 영문/한글 병기 확인
- ✅ 로그 메시지: examples/httputil/main.go 등에서 영문/한글 병기 확인
- ✅ Git 커밋 메시지: 최근 20개 커밋 모두 영문/한글 병기 형식 준수 확인
- ✅ 공개 문서: README.md, USER_MANUAL.md 등 영문/한글 병기 확인
- ✅ 비공개 문서: CLAUDE.md, todo.md 한글 전용으로 변환 완료

### Commits / 커밋

1. **db4afca** - `Chore: Bump version to v1.12.004 / v1.12.004로 버전 증가`
   - 버전 증가만

2. **(pending)** - `Docs: Update CLAUDE.md and todo.md with language policy / 언어 정책으로 CLAUDE.md 및 todo.md 업데이트 (v1.12.004)`
   - 언어 사용 규칙 문서화 및 적용

---

## [v1.12.003] - 2025-10-16

### Added / 추가
- Created errorutil package WORK_PLAN.md with comprehensive task breakdown / 포괄적인 작업 분류를 포함한 errorutil 패키지 WORK_PLAN.md 생성
- Created todo.md for task tracking and progress management / 작업 추적 및 진행 관리를 위한 todo.md 생성
- Defined 12 development phases with clear goals / 명확한 목표를 가진 12개 개발 단계 정의
- Documented 60+ individual tasks across all phases / 모든 단계에 걸쳐 60개 이상의 개별 작업 문서화
- Added completion criteria for each phase / 각 단계에 대한 완료 기준 추가
- Established flexible task ordering within phases / 단계 내 유연한 작업 순서 수립

### Changed / 변경
- N/A

### Fixed / 수정
- N/A

### Files Changed / 변경된 파일
- `cfg/app.yaml` - Version bumped from v1.12.002 to v1.12.003 / 버전을 v1.12.002에서 v1.12.003로 증가
- `docs/errorutil/WORK_PLAN.md` - Created comprehensive work plan with 12 phases / 12개 단계를 포함한 포괄적인 작업 계획서 생성
- `todo.md` - Created task tracking file with all planned tasks / 모든 계획된 작업이 포함된 작업 추적 파일 생성
- `docs/CHANGELOG/CHANGELOG-v1.12.md` - Added v1.12.003 entry / v1.12.003 항목 추가

### Context / 컨텍스트

**User Request / 사용자 요청**: 
"진행해주세요. 단위작업을 잘 만들어 주세요. 필요하다면 todo.md를 만들어 진행할 수 있도록 해주세요. 함수/기능 하나에 패치 하나를 기준으로 합니다. 단위작업에 패치 번호를 할당하지 마세요. 중간에 추가작업이 있을 수 있습니다."
"Please proceed. Create good unit tasks. Create todo.md if needed to proceed. One function/feature = one patch. Don't assign patch numbers to unit tasks. There may be additional tasks in between."

**Why / 이유**: 
- Provide clear roadmap for errorutil package implementation / errorutil 패키지 구현을 위한 명확한 로드맵 제공
- Break down development into manageable, trackable tasks / 개발을 관리 가능하고 추적 가능한 작업으로 분할
- Allow flexibility for adding tasks during development / 개발 중 작업 추가를 위한 유연성 허용
- Follow principle: one function/feature = one patch version / 원칙 준수: 함수/기능 하나 = 패치 버전 하나
- Enable clear progress tracking via todo.md / todo.md를 통한 명확한 진행 상황 추적 가능
- Avoid rigid version number assignment that limits flexibility / 유연성을 제한하는 엄격한 버전 번호 할당 회피

**Impact / 영향**: 
- Clear development path with 12 well-defined phases / 12개의 잘 정의된 단계를 가진 명확한 개발 경로
- 60+ tasks ready to be executed incrementally / 점진적으로 실행할 준비가 된 60개 이상의 작업
- Flexible task ordering allows parallel work when possible / 유연한 작업 순서로 가능한 경우 병렬 작업 허용
- Version numbers assigned during actual work, not planning / 버전 번호는 계획이 아닌 실제 작업 중 할당
- Easy to add new tasks without disrupting version sequence / 버전 순서를 방해하지 않고 새 작업 추가 용이
- todo.md serves as central progress tracking document / todo.md가 중앙 진행 상황 추적 문서로 역할
- Ready to start Phase 1: Core Types implementation / Phase 1: 핵심 타입 구현 시작 준비 완료

**Work Plan Highlights / 작업 계획 주요 사항**:
- **Phase 1**: Core Types (6 tasks) - Error type definitions / 핵심 타입 (6개 작업) - 에러 타입 정의
- **Phase 2**: Error Creation (4 tasks) - Basic creation functions / 에러 생성 (4개 작업) - 기본 생성 함수
- **Phase 3**: Error Wrapping (4 tasks) - Context preservation / 에러 래핑 (4개 작업) - 컨텍스트 보존
- **Phase 4**: Error Inspection (7 tasks) - Information extraction / 에러 검사 (7개 작업) - 정보 추출
- **Phase 5**: Error Classification (8 tasks) - Error categorization / 에러 분류 (8개 작업) - 에러 범주화
- **Phase 6**: Error Formatting (5 tasks) - Output formatting / 에러 포매팅 (5개 작업) - 출력 포매팅
- **Phase 7**: Stack Traces (7 tasks) - Stack capture and display / 스택 트레이스 (7개 작업) - 스택 캡처 및 표시
- **Phase 8**: Context Errors (5 tasks) - Structured data / 컨텍스트 에러 (5개 작업) - 구조화된 데이터
- **Phase 9**: Error Assertions (5 tasks) - Must patterns / 에러 단언 (5개 작업) - Must 패턴
- **Phase 10**: Documentation (7 tasks) - Comprehensive docs / 문서화 (7개 작업) - 포괄적인 문서
- **Phase 11**: Examples (6 tasks) - Real-world scenarios / 예제 (6개 작업) - 실제 시나리오
- **Phase 12**: Testing & Polish (8 tasks) - Production readiness / 테스트 및 마무리 (8개 작업) - 프로덕션 준비

### Commits / 커밋

1. **67465cf** - `Chore: Bump version to v1.12.003 / v1.12.003로 버전 증가`
   - Version bump only / 버전 증가만

2. **(pending)** - `Docs: Create WORK_PLAN.md and todo.md for errorutil development / errorutil 개발을 위한 WORK_PLAN.md 및 todo.md 생성 (v1.12.003)`
   - Created comprehensive work plan and task tracking / 포괄적인 작업 계획 및 작업 추적 생성

---

## [v1.12.002] - 2025-10-16

### Added / 추가
- Created errorutil package DESIGN_PLAN.md with full bilingual format / 완전한 이중 언어 형식의 errorutil 패키지 DESIGN_PLAN.md 생성
- Comprehensive package design documentation (14 sections) / 포괄적인 패키지 설계 문서 (14개 섹션)
- Error types hierarchy and architecture / 에러 타입 계층 및 아키텍처
- Six feature modules with detailed API design / 상세한 API 설계를 포함한 6개 기능 모듈
- Performance considerations and optimization strategies / 성능 고려사항 및 최적화 전략
- Testing strategy with 80%+ coverage target / 80% 이상 커버리지 목표를 가진 테스트 전략
- Migration path from standard library / 표준 라이브러리에서의 마이그레이션 경로
- Version plan (v1.12.001-070) / 버전 계획

### Changed / 변경
- N/A

### Fixed / 수정
- N/A

### Files Changed / 변경된 파일
- `cfg/app.yaml` - Version bumped from v1.12.001 to v1.12.002 / 버전을 v1.12.001에서 v1.12.002로 증가
- `docs/errorutil/DESIGN_PLAN.md` - Created comprehensive design plan with bilingual documentation / 이중 언어 문서로 포괄적인 설계 계획서 생성
- `docs/CHANGELOG/CHANGELOG-v1.12.md` - Added v1.12.002 entry / v1.12.002 항목 추가

### Context / 컨텍스트

**User Request / 사용자 요청**: 
"errorutil 패키지의 DESIGN_PLAN.md를 영문/한글 병기 형식으로 작성해주세요"
"Please create errorutil package DESIGN_PLAN.md with bilingual (English/Korean) format"

**Why / 이유**: 
- Follow the newly established bilingual documentation standards / 새로 수립된 이중 언어 문서화 표준 준수
- Provide comprehensive design documentation before implementation / 구현 전 포괄적인 설계 문서 제공
- Define clear architecture and API design for errorutil package / errorutil 패키지의 명확한 아키텍처 및 API 설계 정의
- Ensure all stakeholders can understand the design (English and Korean speakers) / 모든 이해관계자가 설계를 이해할 수 있도록 보장 (영어 및 한국어 사용자)

**Impact / 영향**: 
- Clear roadmap for errorutil package development / errorutil 패키지 개발을 위한 명확한 로드맵
- Comprehensive design serves as reference during implementation / 포괄적인 설계가 구현 중 참조 자료로 활용
- Bilingual format ensures accessibility for international contributors / 이중 언어 형식으로 국제 기여자의 접근성 보장
- Follows all project documentation standards / 모든 프로젝트 문서화 표준 준수
- Ready to proceed with WORK_PLAN.md creation / WORK_PLAN.md 생성 준비 완료

**Design Highlights / 설계 주요 사항**:
- 6 feature modules: Creation, Wrapping, Inspection, Classification, Formatting, Assertion / 6개 기능 모듈
- 5 error types: Wrapped, Coded, Stack, Context, Composite / 5개 에러 타입
- 40+ planned functions / 40개 이상 계획된 함수
- Zero external dependencies / 외부 의존성 없음
- Standard library compatible / 표준 라이브러리 호환
- 80%+ test coverage target / 80% 이상 테스트 커버리지 목표

### Commits / 커밋

1. **9f67011** - `Chore: Bump version to v1.12.002 / v1.12.002로 버전 증가`
   - Version bump only / 버전 증가만

2. **(pending)** - `Docs: Create errorutil DESIGN_PLAN.md with bilingual format / 이중 언어 형식의 errorutil DESIGN_PLAN.md 생성 (v1.12.002)`
   - Created comprehensive design documentation / 포괄적인 설계 문서 생성

---

## [v1.12.001] - 2025-10-16

### Added / 추가
- Started errorutil package development / errorutil 패키지 개발 시작
- Created errorutil directory structure / errorutil 디렉토리 구조 생성
- Added bilingual (English/Korean) requirements to all development guide documents / 모든 개발 가이드 문서에 이중 언어(영문/한글) 요구사항 추가
- Added detailed CHANGELOG requirements and workflow / 상세한 CHANGELOG 요구사항 및 워크플로우 추가
- Created initial errorutil DESIGN_PLAN.md (English only, to be updated with bilingual version) / 초기 errorutil DESIGN_PLAN.md 생성 (영문만, 이중 언어 버전으로 업데이트 예정)

### Changed / 변경
- Updated PACKAGE_DEVELOPMENT_GUIDE.md with explicit bilingual documentation standards / PACKAGE_DEVELOPMENT_GUIDE.md에 명시적인 이중 언어 문서화 표준 추가
  - Added section "What Must Be Bilingual" / "병기가 필요한 항목" 섹션 추가
  - Added section "What Can Be English-Only" / "영문만 사용 가능한 항목" 섹션 추가
  - Added documentation format examples / 문서 형식 예제 추가
  - Added detailed bilingual commit message format with correct/incorrect examples / 올바른/잘못된 예제와 함께 상세한 이중 언어 커밋 메시지 형식 추가
  - Added comprehensive CHANGELOG writing guidelines (Step 6 expanded) / 포괄적인 CHANGELOG 작성 가이드라인 추가 (Step 6 확장)

- Updated DEVELOPMENT_WORKFLOW_GUIDE.md with bilingual format requirements / DEVELOPMENT_WORKFLOW_GUIDE.md에 이중 언어 형식 요구사항 추가
  - Added "What Must Be Bilingual" section / "반드시 병기해야 하는 항목" 섹션 추가
  - Added "Exceptions (English Only)" section / "예외 (영문만)" 섹션 추가
  - Updated commit message format with bilingual examples / 이중 언어 예제로 커밋 메시지 형식 업데이트
  - Added correct/incorrect commit message examples / 올바른/잘못된 커밋 메시지 예제 추가
  - Added CHANGELOG requirements summary / CHANGELOG 요구사항 요약 추가

- Updated CLAUDE.md with critical bilingual and CHANGELOG requirements / CLAUDE.md에 핵심 이중 언어 및 CHANGELOG 요구사항 추가
  - Added "Bilingual Requirements" section at top / 상단에 "이중 언어 요구사항" 섹션 추가
  - Added "CHANGELOG Requirements" section / "CHANGELOG 요구사항" 섹션 추가
  - Listed what must be bilingual vs. English-only / 병기 필수 항목 vs. 영문만 항목 나열
  - Added commit message format examples / 커밋 메시지 형식 예제 추가

### Fixed / 수정
- N/A

### Files Changed / 변경된 파일
- `cfg/app.yaml` - Version bumped from v1.11.046 to v1.12.001 / 버전을 v1.11.046에서 v1.12.001로 증가
- `CLAUDE.md` - Added bilingual and CHANGELOG requirements sections / 이중 언어 및 CHANGELOG 요구사항 섹션 추가
- `docs/DEVELOPMENT_WORKFLOW_GUIDE.md` - Enhanced bilingual format and commit message guidelines / 이중 언어 형식 및 커밋 메시지 가이드라인 강화
- `docs/PACKAGE_DEVELOPMENT_GUIDE.md` - Added comprehensive bilingual and CHANGELOG documentation / 포괄적인 이중 언어 및 CHANGELOG 문서화 추가
- `docs/errorutil/DESIGN_PLAN.md` - Created initial design plan (English only) / 초기 설계 계획서 생성 (영문만)
- `errorutil/` - Created package directory / 패키지 디렉토리 생성
- `docs/errorutil/` - Created documentation directory / 문서 디렉토리 생성
- `examples/errorutil/` - Created examples directory / 예제 디렉토리 생성
- `docs/CHANGELOG/CHANGELOG-v1.12.md` - Created this changelog file / 이 변경 로그 파일 생성

### Context / 컨텍스트

**User Request / 사용자 요청**: 
1. "문서는 영문과 한글을 항상 병기해야 합니다. 규칙에도 추가해 주세요. 코드내 주석도 마찬가지입니다."
   "Documentation must always include both English and Korean. Please add this to the rules. Same for code comments."

2. "깃헙의 커밋 메시지도 앞으로는 병기했으면 좋겠습니다."
   "I'd like GitHub commit messages to also be bilingual going forward."

3. "깃헙(커밋과 푸쉬등 작업)을 하기 전에 반드시 CHANGELOG를 작성해야 합니다. 어떤 파일이 어떻게 바뀌었고, 왜 바뀌었고, 무슨 요청이 었고 등등등.. 또한 루트의 CHANGELOG.md는 메이저와 마이너 버젼별 아웃룩한 부분만 명시하고, 'docs/CHANGELOG/' 에 각 마이너 버젼별로 파일이 있습니다.(없으면 만들어서) 여기에 자세히 적는겁니다. 이또한 규칙에 넣어두어서 제가 다시 언급안해도 되게 해주세요."
   "CHANGELOG must be written before any GitHub work (commit, push, etc.). Include what files changed, how they changed, why they changed, what the request was, etc. The root CHANGELOG.md should only show high-level overview by major/minor version, while 'docs/CHANGELOG/' should have detailed files for each minor version (create if not exists). Please add this to the rules so I don't have to mention it again."

**Why / 이유**: 
- Establish consistent bilingual documentation standards across the entire project / 전체 프로젝트에 걸쳐 일관된 이중 언어 문서화 표준 수립
- Make bilingual requirements explicit so they are automatically followed / 이중 언어 요구사항을 명시적으로 만들어 자동으로 따르도록 함
- Ensure comprehensive change tracking with detailed CHANGELOG for better project history / 상세한 CHANGELOG로 포괄적인 변경 추적을 보장하여 더 나은 프로젝트 이력 확보
- Prevent having to repeatedly ask for bilingual documentation and proper CHANGELOG updates / 이중 언어 문서화 및 적절한 CHANGELOG 업데이트를 반복적으로 요청하지 않도록 방지
- Start errorutil package development with proper foundation / 적절한 기반으로 errorutil 패키지 개발 시작

**Impact / 영향**: 
- All future documentation will automatically be bilingual / 향후 모든 문서가 자동으로 이중 언어로 작성됨
- All future commit messages will be bilingual / 향후 모든 커밋 메시지가 이중 언어로 작성됨
- All changes will be thoroughly documented in CHANGELOG before commits / 모든 변경사항이 커밋 전 CHANGELOG에 철저히 문서화됨
- Better project history and traceability / 더 나은 프로젝트 이력 및 추적성
- Improved international accessibility (English and Korean speakers) / 향상된 국제 접근성 (영어 및 한국어 사용자)
- New errorutil package ready for feature development / 새로운 errorutil 패키지가 기능 개발 준비 완료

### Commits / 커밋

1. **17108ee** - `Chore: Bump version to v1.12.001 - Start errorutil package development`
   - Version bump only / 버전 증가만

2. **3fc650c** - `Docs: Add bilingual requirements to development guides / 개발 가이드에 이중 언어 요구사항 추가 (v1.12.001)`
   - Added bilingual and CHANGELOG rules to guide documents / 가이드 문서에 이중 언어 및 CHANGELOG 규칙 추가
   - Created initial errorutil DESIGN_PLAN.md / 초기 errorutil DESIGN_PLAN.md 생성

---

## Version Summary / 버전 요약

- **v1.12.001**: Package initialization, bilingual requirements, CHANGELOG workflow / 패키지 초기화, 이중 언어 요구사항, CHANGELOG 워크플로우

---

**Next Steps / 다음 단계**:
1. Update errorutil DESIGN_PLAN.md with bilingual format / errorutil DESIGN_PLAN.md를 이중 언어 형식으로 업데이트
2. Create errorutil WORK_PLAN.md / errorutil WORK_PLAN.md 생성
3. Begin implementing core error types / 핵심 에러 타입 구현 시작
