# CHANGELOG v1.12.x - errorutil Package / 에러 처리 유틸리티 패키지

Error handling utilities package for Go applications.

Go 애플리케이션을 위한 에러 처리 유틸리티 패키지입니다.

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
