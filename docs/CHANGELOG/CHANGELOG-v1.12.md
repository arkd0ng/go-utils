# CHANGELOG v1.12.x - errorutil Package / 에러 처리 유틸리티 패키지

Error handling utilities package for Go applications.

Go 애플리케이션을 위한 에러 처리 유틸리티 패키지입니다.

---

## [v1.12.015] - 2025-10-17

### Changed / 변경
- examples/errorutil/main.go의 printBanner 함수를 logger로 출력하도록 개선
- fmt.Println(banner) 제거하여 모든 출력을 logger로 일관성 있게 처리

### Files Changed / 변경된 파일
- `cfg/app.yaml` - 버전을 v1.12.014에서 v1.12.015로 증가
- `examples/errorutil/main.go` - printBanner 함수 리팩토링 (fmt.Println → logger.Info)
- `docs/CHANGELOG/CHANGELOG-v1.12.md` - v1.12.015 항목 추가

### Context / 컨텍스트

**User Request / 사용자 요청**:
"배너(printBanner)로 로그로 남기면 안되나요?"

**Why / 이유**:
- fmt.Println과 logger.Info를 혼용하면 일관성이 떨어짐
- logging.WithStdout(true) 사용 시 모든 출력이 logger를 통해 이루어져야 함
- 로그 파일과 콘솔 출력이 동일한 형식으로 일관성 있게 기록됨
- 배너도 로그로 남겨야 추후 로그 분석 시 유용함

**Implementation Details / 구현 세부사항**:

1. **printBanner 함수 리팩토링**:
   - fmt.Sprintf로 만든 배너 문자열 제거
   - fmt.Println(banner) 제거
   - 각 라인을 logger.Info()로 개별 출력
   - 빈 라인도 logger.Info("")로 처리

2. **출력 일관성 확보**:
   - 모든 출력이 logger를 통해 이루어짐
   - 콘솔과 파일에 동일한 형식으로 기록
   - 타임스탬프와 로그 레벨이 모든 라인에 포함

3. **결과**:
   - fmt.Println 완전 제거 (로거 초기화 전 2개만 유지)
   - 배너 출력도 구조화된 로그로 남음
   - 로그 파일 분석 시 배너 정보도 포함

**Impact / 영향**:
- 출력 일관성 향상 (모든 출력이 logger 통해 처리)
- 로그 파일에 배너 정보 포함으로 분석 용이
- logging.WithStdout(true)의 장점 완전 활용
- 미래의 모든 예제 코드가 이 패턴을 따라야 함

---

## [v1.12.014] - 2025-10-17

### Changed / 변경
- examples/errorutil/main.go의 버전 정보를 동적 로딩으로 변경
- fmt.Printf 제거하여 중복 출력 방지 (logging.WithStdout(true) 사용)
- CLAUDE.md에 버전 정보 관리 규칙 추가

### Files Changed / 변경된 파일
- `cfg/app.yaml` - 버전을 v1.12.013에서 v1.12.014로 증가
- `examples/errorutil/main.go` - logging.TryLoadAppVersion() 사용, fmt.Printf 제거 (41개 → 2개)
- `CLAUDE.md` - 버전 정보 관리 규칙 섹션 추가 (65줄)
- `docs/CHANGELOG/CHANGELOG-v1.12.md` - v1.12.014 항목 추가

### Context / 컨텍스트

**User Request / 사용자 요청**:
버전 정보는 반드시 cfg/app.yaml에서 동적으로 읽도록 수정.
하드코딩 금지 규칙을 CLAUDE.md에 추가하여 지속적으로 따르도록 함.

**Why / 이유**:
- 버전 정보 하드코딩은 유지보수 문제 발생 (여러 곳에서 업데이트 필요)
- 단일 진실의 원천(Single Source of Truth) 원칙 위반
- cfg/app.yaml이 유일한 버전 정보 원천이어야 함
- 문서에서도 배지를 통해 동적으로 버전 표시

**Implementation Details / 구현 세부사항**:

1. **examples/errorutil/main.go 수정**:
   - 하드코딩된 버전 "v1.12.013" 제거
   - `logging.TryLoadAppVersion()` 함수 사용으로 변경
   - Fallback 값 "unknown" 설정
   - fmt.Printf 제거 (logging.WithStdout(true)로 콘솔 출력)

2. **CLAUDE.md 규칙 추가** (65줄):
   - 버전 정보 관리 규칙 섹션 신설
   - 단일 진실의 원천 원칙 명시
   - 코드에서 사용법: `logging.TryLoadAppVersion()`
   - 문서에서 사용법: 배지 또는 참조
   - 절대 금지 사항: 하드코딩
   - 예외 사항: CHANGELOG, 커밋 메시지
   - 올바른 예제와 잘못된 예제 포함

3. **fmt.Printf 제거**:
   - logging.WithStdout(true) 활성화로 콘솔 출력
   - 중복 출력 방지 (logger와 fmt.Printf 동시 사용 X)
   - 로거 초기화 전 메시지만 fmt 유지 (2개)
   - 나머지 41개 fmt.Printf 제거

**Impact / 영향**:
- 버전 정보 관리 일관성 확보
- cfg/app.yaml 한 곳만 수정하면 모든 곳에 반영
- 미래의 모든 예제 코드가 이 규칙을 따름
- 문서화 품질 향상 (동적 배지 사용)
- 콘솔 출력 중복 제거로 가독성 향상

---

## [v1.12.013] - 2025-10-17

### Added / 추가
- examples/errorutil/main.go 작성 (포괄적인 예제 코드, 12개 예제)
- 18개 함수 전체를 시연하는 예제 구현
- 실제 사용 패턴 예제 (HTTP API, 데이터베이스, 검증, 다중 레이어)
- 에러 분류 시스템 예제
- 표준 라이브러리 호환성 데모 (errors.Is, errors.As)

### Changed / 변경
- N/A

### Fixed / 수정
- N/A

### Files Changed / 변경된 파일
- `cfg/app.yaml` - 버전을 v1.12.012에서 v1.12.013로 증가
- `examples/errorutil/main.go` - 새 파일 생성 (600+ 줄, 12개 예제 포함)
- `logs/errorutil-example.log` - 예제 실행 로그 생성 (22KB)
- `docs/CHANGELOG/CHANGELOG-v1.12.md` - v1.12.013 항목 추가

### Context / 컨텍스트

**User Request / 사용자 요청**:
errorutil 패키지 문서화 작업 완료 - 예제 코드 작성

**Why / 이유**:
- 사용자가 실제로 작동하는 코드를 보고 학습할 수 있도록 지원
- EXAMPLE_CODE_GUIDE.md의 모든 표준 준수
- 모든 기능을 커버하는 포괄적인 예제 제공
- 실제 사용 패턴을 통한 학습 곡선 단축

**Implementation Details / 구현 세부사항**:

1. **예제 구조** (12개 예제):
   - Example 1: Basic Error Creation (New, Newf)
   - Example 2: String Coded Errors (WithCode, WithCodef, HasCode)
   - Example 3: Numeric Coded Errors (WithNumericCode, WithNumericCodef, GetNumericCode)
   - Example 4: Error Wrapping (Wrap, Wrapf)
   - Example 5: Error Chain Walking (다중 레이어 체인)
   - Example 6: Error Inspection (HasCode, HasNumericCode, GetCode, GetNumericCode)
   - Example 7: HTTP API Error Handling (404, 500, 401 시나리오)
   - Example 8: Database Error Patterns (연결, 쿼리, Not Found)
   - Example 9: Validation Error Patterns (필수 필드, 형식, 범위)
   - Example 10: Error Classification System (에러 분류 및 처리)
   - Example 11: Multi-Layer Wrapping (4개 레이어: DB → Repo → Service → HTTP)
   - Example 12: Standard Library Compatibility (errors.Is, errors.As)

2. **로깅 시스템**:
   - EXAMPLE_CODE_GUIDE.md 표준 완전 준수
   - 로그 백업 관리 (최근 5개 유지)
   - 모든 작업을 영문/한글 병기로 로그
   - 극도로 상세한 로그 (모든 입력, 출력, 에러 기록)
   - 구조화된 로깅 (key-value 쌍)

3. **주요 특징**:
   - 모든 18개 함수 시연
   - 실전 패턴 (HTTP API, DB, 검증)
   - 에러 체인 탐색 데모
   - Standard library 호환성 증명
   - 유니코드 기호 사용 (✅, ❌, 📊)
   - 모든 주석 영문/한글 병기

4. **테스트 결과**:
   - 예제 성공적으로 컴파일 및 실행
   - 로그 파일 정상 생성 (22KB)
   - 모든 12개 예제 정상 작동
   - 에러 없이 완료

**Impact / 영향**:
- 사용자가 복사하여 바로 사용할 수 있는 예제 제공
- errorutil의 모든 기능을 실제 코드로 학습 가능
- 실전 패턴을 통해 즉시 프로젝트에 적용 가능
- errorutil 패키지 문서화 100% 완료 (README → USER_MANUAL → DEVELOPER_GUIDE → Examples)
- 다음 작업: 테스트 실행 및 커버리지 확인

---

## [v1.12.012] - 2025-10-17

### Added / 추가
- docs/errorutil/DEVELOPER_GUIDE.md 작성 (포괄적인 개발자 가이드)
- 10개 주요 섹션 포함: 아키텍처 개요, 패키지 구조, 핵심 컴포넌트, 디자인 패턴, 내부 구현, 새 기능 추가, 테스트 가이드, 성능, 기여 가이드라인, 코드 스타일
- 인터페이스 및 에러 타입 상세 설명 (6개 타입, 5개 인터페이스)
- 에러 생성/래핑/검사 흐름 다이어그램
- 새 기능 추가 단계별 가이드 (새 에러 타입, 새 검사 함수)

### Changed / 변경
- N/A

### Fixed / 수정
- N/A

### Files Changed / 변경된 파일
- `cfg/app.yaml` - 버전을 v1.12.011에서 v1.12.012로 증가
- `docs/errorutil/DEVELOPER_GUIDE.md` - 새 파일 생성 (800+ 줄, 포괄적인 개발자 가이드)
- `docs/CHANGELOG/CHANGELOG-v1.12.md` - v1.12.012 항목 추가

### Context / 컨텍스트

**User Request / 사용자 요청**:
errorutil 패키지 문서화 작업 계속 진행 - DEVELOPER_GUIDE.md 작성

**Why / 이유**:
- USER_MANUAL은 사용자용, DEVELOPER_GUIDE는 패키지 유지보수 및 기여자용
- 패키지의 내부 아키텍처 및 설계 결정 문서화
- 새 기능 추가 및 패키지 확장을 위한 가이드 제공
- 코드 스타일 및 기여 가이드라인 명확화

**Implementation Details / 구현 세부사항**:

1. **가이드 구조**:
   - Architecture Overview: 설계 원칙 7가지, 상위 수준 아키텍처, 컴포넌트 상호작용
   - Package Structure: 파일 구성, 파일별 책임 (~2,544 줄 코드)
   - Core Components: 5개 인터페이스, 6개 에러 타입, Frame 타입 상세 설명
   - Design Patterns: Factory, Decorator, Chain of Responsibility, Template Method 패턴
   - Internal Implementation: 에러 생성/래핑/검사 흐름 다이어그램
   - Adding New Features: 새 에러 타입 추가 5단계, 새 검사 함수 추가 가이드
   - Testing Guide: 테스트 구조, 4가지 테스트 카테고리, 커버리지 요구사항 (99.2%)
   - Performance: 할당 벤치마크, 성능 고려사항 3가지
   - Contributing Guidelines: 개발 워크플로우 4단계, Pull Request 체크리스트
   - Code Style: 명명 규칙, 주석 스타일 (이중 언어), 에러 메시지 형식

2. **주요 내용**:
   - 각 인터페이스의 목적 및 구현자 설명 (Unwrapper, Coder, NumericCoder, StackTracer, Contexter)
   - 6개 내부 에러 타입 상세 (wrappedError, codedError, numericCodedError, stackError, contextError, compositeError)
   - 에러 생성/래핑/검사의 전체 흐름 다이어그램
   - 에러 체인 탐색 메커니즘 설명 (errors.As 사용)
   - 디자인 패턴 적용 예제 (Factory, Decorator, Chain of Responsibility, Template Method)
   - 새 기능 추가 단계별 가이드 (코드 예제 포함)
   - 테스트 카테고리 및 실행 방법
   - 성능 벤치마크 및 최적화 고려사항

3. **개발자 지원 요소**:
   - 모든 섹션 영문/한글 병기
   - 코드 예제로 개념 설명
   - 다이어그램으로 복잡한 흐름 시각화
   - 실용적인 체크리스트 및 가이드라인
   - 인터페이스 호환성 매트릭스
   - 함수 참조 테이블

**Impact / 영향**:
- 패키지 유지보수자가 내부 구조를 완전히 이해 가능
- 새 기능 추가 시 일관된 패턴 유지
- 기여자가 코드 스타일 및 테스트 요구사항 준수
- 문서 품질: README (Quick Start) → USER_MANUAL (Usage) → DEVELOPER_GUIDE (Architecture)
- errorutil 예제 코드 작성 준비 완료

---

## [v1.12.011] - 2025-10-17

### Added / 추가
- docs/errorutil/USER_MANUAL.md 작성 (포괄적인 사용자 매뉴얼)
- 12개 주요 섹션 포함: 소개, 설치, 빠른 시작, 핵심 개념, 에러 생성, 에러 래핑, 에러 검사, 고급 사용법, 모범 사례, 일반 패턴, 문제 해결, API 참조
- 실전 사용 패턴 4가지 제공 (검증 에러, 데이터베이스 에러, 외부 서비스 에러, 재시도 로직)
- 완전한 API 참조 문서 (18개 함수, 매개변수, 반환값, 예제 포함)

### Changed / 변경
- N/A

### Fixed / 수정
- N/A

### Files Changed / 변경된 파일
- `cfg/app.yaml` - 버전을 v1.12.010에서 v1.12.011로 증가
- `docs/errorutil/USER_MANUAL.md` - 새 파일 생성 (1000+ 줄, 포괄적인 사용자 매뉴얼)
- `docs/CHANGELOG/CHANGELOG-v1.12.md` - v1.12.011 항목 추가

### Context / 컨텍스트

**User Request / 사용자 요청**:
errorutil 패키지 문서화 작업 계속 진행 - USER_MANUAL.md 작성

**Why / 이유**:
- README는 빠른 시작용, USER_MANUAL은 상세 학습용
- 사용자가 errorutil의 모든 기능을 완전히 이해할 수 있도록 지원
- 실전 패턴과 문제 해결 가이드 제공
- 초보자부터 고급 사용자까지 모든 수준에 대응

**Implementation Details / 구현 세부사항**:

1. **매뉴얼 구조**:
   - Introduction: errorutil 소개 및 사용 이유
   - Installation: 설치 및 확인 방법
   - Quick Start: 5가지 기본 사용 예제
   - Core Concepts: 에러 체인, 에러 코드, 에러 인터페이스
   - Error Creation: 6개 함수 상세 설명
   - Error Wrapping: 6개 함수 상세 설명 및 다중 레벨 래핑
   - Error Inspection: 4개 함수 상세 설명
   - Advanced Usage: 표준 라이브러리 호환성, HTTP API, 커스텀 타입, 분류 시스템
   - Best Practices: 6가지 모범 사례 (✅ Good / ❌ Bad 예제)
   - Common Patterns: 4가지 실전 패턴 (검증, DB, 외부 서비스, 재시도)
   - Troubleshooting: 4가지 일반 문제 해결
   - API Reference: 18개 함수 완전 문서화

2. **주요 내용**:
   - 각 함수에 대한 상세 설명, 매개변수, 반환값, 예제
   - Why use errorutil? - 표준 라이브러리 대비 장점
   - Error chains 개념 및 다중 레벨 래핑 예제
   - String vs Numeric codes 사용 시기
   - HTTP API 에러 처리 완전 예제
   - 커스텀 에러 타입 구현 방법
   - 에러 분류 시스템 구축 방법
   - 재시도 로직 구현 패턴

3. **교육적 요소**:
   - 모든 섹션에 영문/한글 병기
   - ✅ Good / ❌ Bad 비교 예제
   - 실전 시나리오 (HTTP handler, database, validation)
   - 문제 상황과 해결 방법 제시
   - 완전한 코드 예제 (복사하여 바로 사용 가능)

**Impact / 영향**:
- 사용자가 errorutil의 모든 기능을 단계적으로 학습 가능
- 실전 패턴을 통해 즉시 프로젝트에 적용 가능
- 문제 해결 섹션으로 일반적인 실수 예방
- DEVELOPER_GUIDE 작성을 위한 기반 마련
- 문서 품질: README (Quick Start) → USER_MANUAL (Complete Guide) → DEVELOPER_GUIDE (Architecture)

---

## [v1.12.010] - 2025-10-17

### Added / 추가
- errorutil/README.md 작성 (포괄적인 패키지 문서)
- 빠른 시작 가이드 및 예제 코드 포함
- API 참조 문서 (18개 함수 전체 설명)
- 모범 사례 섹션 추가

### Changed / 변경
- N/A

### Fixed / 수정
- N/A

### Files Changed / 변경된 파일
- `cfg/app.yaml` - 버전을 v1.12.009에서 v1.12.010으로 증가
- `errorutil/README.md` - 새 파일 생성 (450+ 줄, 포괄적인 문서)
- `docs/CHANGELOG/CHANGELOG-v1.12.md` - v1.12.010 항목 추가

### Context / 컨텍스트

**User Request / 사용자 요청**:
errorutil 패키지 문서화 작업 계속 진행

**Why / 이유**:
- 사용자가 패키지를 빠르게 이해하고 사용할 수 있도록 안내
- API 참조를 통한 모든 함수의 사용법 제공
- 실제 사용 사례를 보여주는 예제 코드 제공
- 모범 사례를 통한 올바른 사용 패턴 안내

**Implementation Details / 구현 세부사항**:

1. **README 구조**:
   - Overview: 패키지 소개 및 주요 기능
   - Features: 핵심 기능 목록
   - Quick Start: 즉시 사용 가능한 예제
   - API Reference: 18개 함수 전체 설명
   - Examples: 실제 사용 사례 (HTTP API, 에러 분류, 중첩 체인)
   - Best Practices: 모범 사례 4가지

2. **문서화 내용**:
   - Error Creation: 6개 함수 (New, Newf, WithCode, WithCodef, WithNumericCode, WithNumericCodef)
   - Error Wrapping: 6개 함수 (Wrap, Wrapf, WrapWithCode, WrapWithCodef, WrapWithNumericCode, WrapWithNumericCodef)
   - Error Inspection: 4개 함수 (HasCode, HasNumericCode, GetCode, GetNumericCode)

3. **예제 시나리오**:
   - HTTP API 에러 처리
   - 에러 코드를 사용한 에러 분류
   - 깊게 중첩된 에러 체인 처리

**Impact / 영향**:
- 사용자가 5분 내에 패키지 사용 시작 가능
- 모든 API가 명확히 문서화되어 참조 용이
- 실제 사용 사례를 통한 학습 곡선 단축
- 다음 단계(USER_MANUAL, DEVELOPER_GUIDE) 작성 준비 완료

---

## [v1.12.009] - 2025-10-17

### Added / 추가
- N/A

### Changed / 변경
- examples/logging/main.go의 모든 주석을 영문/한글 병기로 개선하고 초보자 친화적으로 작성
- examples/websvrutil/main.go의 인라인 주석에 한글 병기 추가
- 로그 출력 메시지를 영문/한글 병기로 변경하여 예제 명확성 향상
- 헤더 정보 섹션을 문서 표준에 맞게 이중 언어로 확장
- 미들웨어 테스트 설명을 영문과 한글로 명확히 개선

### Fixed / 수정
- N/A

### Files Changed / 변경된 파일
- `cfg/app.yaml` - 버전을 v1.12.008에서 v1.12.009로 증가
- `examples/logging/main.go` - 모든 주석과 로그 메시지를 영문/한글 병기로 개선
- `examples/websvrutil/main.go` - 인라인 주석 및 미들웨어 설명에 한글 추가
- `docs/CHANGELOG/CHANGELOG-v1.12.md` - v1.12.009 항목 추가

### Context / 컨텍스트

**User Request / 사용자 요청**:
별도 세션에서 진행된 문서/주석 보강 작업을 CHANGELOG에 반영 (CHANGELOG-specials.md 참조)

**Why / 이유**:
- CLAUDE.md의 언어 사용 정책(영문/한글 병기, 매우 상세하고 친절한 주석) 준수
- 예제 코드의 교육적 가치 향상
- 한국어 사용자의 접근성 개선
- 초보자도 쉽게 이해할 수 있는 친절한 주석 제공

**Implementation Details / 구현 세부사항**:

1. **examples/logging/main.go 개선**:
   - 모든 함수 주석을 더 상세하고 친절하게 작성
   - 모든 로그 메시지에 한글 병기
   - displayHeader 함수의 모든 정보를 이중 언어로 표시

2. **examples/websvrutil/main.go 개선**:
   - 영문 전용 인라인 주석에 한글 병기
   - 미들웨어 테스트 설명을 영문과 한글로 명확히 작성

**Impact / 영향**:
- 한국어 사용자의 예제 코드 이해도 대폭 향상
- 국제화된 코드베이스 표준 수립
- 별도 세션 문서 작업이 메인 CHANGELOG에 통합되는 프로세스 확립

---

## [v1.12.008] - 2025-10-17

### Added / 추가
- errorutil 패키지 Phase 4 (Error Inspection) 완료
- 6개의 에러 검사 함수 구현:
  - HasCode(err, code): 문자열 코드 존재 여부 확인
  - HasNumericCode(err, code): 숫자 코드 존재 여부 확인
  - GetCode(err): 문자열 코드 추출
  - GetNumericCode(err): 숫자 코드 추출
  - GetStackTrace(err): 스택 트레이스 추출
  - GetContext(err): 컨텍스트 데이터 추출
- 에러 체인 탐색 기능 (errors.As 사용)
- 모든 검사 함수에 대한 포괄적인 테스트 추가
- GetContext 불변성 테스트 추가

### Changed / 변경
- N/A

### Fixed / 수정
- N/A

### Files Changed / 변경된 파일
- `cfg/app.yaml` - 버전을 v1.12.007에서 v1.12.008로 증가
- `errorutil/inspect.go` - 새 파일 생성 (270+ 줄, 6개 검사 함수)
- `errorutil/inspect_test.go` - 새 파일 생성 (420+ 줄, 포괄적인 테스트)
- `docs/CHANGELOG/CHANGELOG-v1.12.md` - v1.12.008 항목 추가

### Context / 컨텍스트

**User Request / 사용자 요청**:
Phase 3 완료 후 자동으로 Phase 4로 진행

**Why / 이유**:
- Phase 4는 에러 정보를 추출하고 검사하는 핵심 기능
- 에러 체인을 탐색하여 원하는 정보를 찾는 유틸리티 함수 제공
- Go 표준 라이브러리의 errors.As, errors.Is와 유사하지만 더 구체적
- 에러 처리 로직에서 조건부 분기를 쉽게 구현 가능

**Implementation Details / 구현 세부사항**:

1. **코드 확인 함수**:
   - HasCode(): 에러 체인에 특정 문자열 코드 존재 여부 확인
   - HasNumericCode(): 에러 체인에 특정 숫자 코드 존재 여부 확인
   - 현재 에러를 먼저 확인한 후 errors.As로 체인 탐색

2. **코드 추출 함수**:
   - GetCode(): 에러 체인에서 첫 번째 문자열 코드 추출
   - GetNumericCode(): 에러 체인에서 첫 번째 숫자 코드 추출
   - 코드를 찾지 못하면 (빈 문자열/0, false) 반환

3. **메타데이터 추출 함수**:
   - GetStackTrace(): 에러 체인에서 스택 트레이스 추출
   - GetContext(): 에러 체인에서 컨텍스트 데이터 추출
   - 불변성 보장 (컨텍스트는 복사본 반환)

4. **에러 체인 탐색**:
   - 모든 함수는 현재 에러를 먼저 확인
   - errors.As를 사용하여 에러 체인의 모든 에러 검사
   - 깊게 래핑된 에러도 정확히 탐색

**Impact / 영향**:
- 에러 코드 기반 조건부 처리 가능
- HTTP 상태 코드 추출 및 응답 생성 용이
- 스택 트레이스를 통한 디버깅 지원
- 컨텍스트 데이터를 통한 구조화된 로깅 가능
- 다음 단계(Phase 5 이후)의 기초 제공
- 전체 커버리지 99.2%로 목표 80% 초과

### Test Results / 테스트 결과

```
PASS
ok  	github.com/arkd0ng/go-utils/errorutil	0.696s
coverage: 99.2% of statements
```

All 26 test functions passed with 99 subtests:
- TestNew (3 cases)
- TestNewf (4 cases)
- TestWithCode (4 cases)
- TestWithCodef (3 cases)
- TestWithNumericCode (4 cases)
- TestWithNumericCodef (3 cases)
- TestWrap (3 cases)
- TestWrapf (3 cases)
- TestWrapWithCode (3 cases)
- TestWrapWithCodef (3 cases)
- TestWrapWithNumericCode (3 cases)
- TestWrapWithNumericCodef (3 cases)
- TestHasCode (7 cases)
- TestHasNumericCode (7 cases)
- TestGetCode (6 cases)
- TestGetNumericCode (7 cases)
- TestGetStackTrace (5 cases)
- TestGetContext (6 cases)
- TestGetContextImmutability (1 case)
- + Phase 1 tests (7 functions, 14 subtests)

### Next Steps / 다음 단계

나머지 Phase들은 기본 기능 위에 추가되는 선택적 기능들입니다:
- Phase 5-9: 고급 기능 (Classification, Formatting, Stack Traces, Context, Assertions)
- Phase 10-12: 문서화 및 예제 (Documentation, Examples, Testing & Polish)

현재 Phase 1-4 완료로 errorutil의 핵심 기능은 모두 구현되었습니다.

---

## [v1.12.007] - 2025-10-17

### Added / 추가
- errorutil 패키지 Phase 3 (Error Wrapping) 완료
- 6개의 에러 래핑 함수 구현:
  - Wrap(cause, message): 기본 에러 래핑
  - Wrapf(cause, format, args...): 포맷된 에러 래핑
  - WrapWithCode(cause, code, message): 문자열 코드와 함께 래핑
  - WrapWithCodef(cause, code, format, args...): 문자열 코드와 포맷된 메시지로 래핑
  - WrapWithNumericCode(cause, code, message): 숫자 코드와 함께 래핑
  - WrapWithNumericCodef(cause, code, format, args...): 숫자 코드와 포맷된 메시지로 래핑
- nil 에러 처리 (nil을 래핑하면 nil 반환)
- 모든 래핑 함수에 대한 포괄적인 테스트 추가

### Changed / 변경
- N/A

### Fixed / 수정
- N/A

### Files Changed / 변경된 파일
- `cfg/app.yaml` - 버전을 v1.12.006에서 v1.12.007로 증가
- `errorutil/error.go` - 6개 래핑 함수 추가 (200+ 줄 추가, 총 370+ 줄)
- `errorutil/error_test.go` - 6개 래핑 함수 테스트 추가 (460+ 줄 추가, 총 830+ 줄)
- `docs/CHANGELOG/CHANGELOG-v1.12.md` - v1.12.007 항목 추가

### Context / 컨텍스트

**User Request / 사용자 요청**:
Phase 2 완료 후 자동으로 Phase 3로 진행

**Why / 이유**:
- Phase 3는 에러 컨텍스트를 추가하는 핵심 기능
- Wrap 함수들은 에러가 콜 스택을 올라가면서 컨텍스트를 추가하는 일반적인 패턴
- Go 표준 라이브러리의 fmt.Errorf("%w", err)와 유사하지만 더 많은 기능 제공
- 에러 코드를 추가하면서 래핑하여 에러 분류와 추적 동시 지원

**Implementation Details / 구현 세부사항**:

1. **기본 래핑**:
   - Wrap(): 단순 메시지로 기존 에러 래핑
   - Wrapf(): 포맷된 메시지로 래핑
   - nil 에러를 래핑하면 nil 반환 (안전성)

2. **코드가 있는 래핑**:
   - WrapWithCode(): 문자열 코드 추가하며 래핑
   - WrapWithNumericCode(): 숫자 코드 추가하며 래핑
   - 각각 포맷 변형(WrapWithCodef, WrapWithNumericCodef) 제공

3. **인터페이스 호환성**:
   - 모든 함수는 Unwrapper 인터페이스 구현
   - errors.Is, errors.As와 완전히 호환
   - 코드가 있는 래핑은 Coder/NumericCoder 인터페이스도 구현

4. **테스트 커버리지**:
   - nil 에러 래핑 테스트
   - 빈 메시지 테스트
   - 다중 인자 포맷 테스트
   - Unwrap 동작 검증

**Impact / 영향**:
- 에러 전파 시 컨텍스트 추가 가능
- 에러 체인을 통한 원인 추적 가능
- 코드 추가로 에러 분류 및 처리 용이
- 다음 단계(Phase 4: Error Inspection)의 기초 제공
- 전체 커버리지 98.6%로 목표 80% 초과

### Test Results / 테스트 결과

```
PASS
ok  	github.com/arkd0ng/go-utils/errorutil	1.557s
coverage: 98.6% of statements
```

All 19 test functions passed with 54 subtests:
- TestNew (3 cases)
- TestNewf (4 cases)
- TestWithCode (4 cases)
- TestWithCodef (3 cases)
- TestWithNumericCode (4 cases)
- TestWithNumericCodef (3 cases)
- TestWrap (3 cases)
- TestWrapf (3 cases)
- TestWrapWithCode (3 cases)
- TestWrapWithCodef (3 cases)
- TestWrapWithNumericCode (3 cases)
- TestWrapWithNumericCodef (3 cases)
- + Phase 1 tests (7 functions, 14 subtests)

### Next Steps / 다음 단계

Phase 4: Error Inspection (에러 검사 함수)
- HasCode() 함수 구현
- HasNumericCode() 함수 구현
- GetCode() 함수 구현
- GetNumericCode() 함수 구현
- GetStackTrace() 함수 구현
- GetContext() 함수 구현

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
