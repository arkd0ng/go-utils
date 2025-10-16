# CHANGELOG v1.12.x - errorutil Package / 에러 처리 유틸리티 패키지

Error handling utilities package for Go applications.

Go 애플리케이션을 위한 에러 처리 유틸리티 패키지입니다.

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
