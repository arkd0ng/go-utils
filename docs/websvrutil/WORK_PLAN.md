# websvrutil Work Plan / 작업 계획

**Version / 버전**: v1.11.001
**Created / 생성일**: 2025-10-16
**Package / 패키지**: `github.com/arkd0ng/go-utils/websvrutil`

---

## Table of Contents / 목차

- [Overview / 개요](#overview--개요)
- [Development Phases / 개발 단계](#development-phases--개발-단계)
- [Task Breakdown / 작업 분류](#task-breakdown--작업-분류)
- [Progress Tracking / 진행 상황 추적](#progress-tracking--진행-상황-추적)

---

## Overview / 개요

This work plan outlines the development phases for the `websvrutil` package, focusing on **developer convenience** over raw performance.

이 작업 계획은 순수 성능보다 **개발자 편의성**에 초점을 맞춘 `websvrutil` 패키지의 개발 단계를 설명합니다.

### Goals / 목표

- **Reduce 50+ lines to 5 lines** / 50줄 이상을 5줄로 줄이기
- **Intuitive API** / 직관적인 API
- **Auto template discovery** / 자동 템플릿 발견
- **Easy middleware chaining** / 쉬운 미들웨어 체이닝
- **Smart defaults** / 스마트 기본값

---

## Development Phases / 개발 단계

### Phase 1: Core Foundation (v1.11.001-005) ✅ Planning

**Focus / 초점**: App, Router, Context, basic handlers

**Core Components / 핵심 컴포넌트**:
1. App struct and constructor / App 구조체 및 생성자
2. Router with path matching / 경로 매칭을 가진 라우터
3. Context wrapper / Context 래퍼
4. Basic response helpers / 기본 응답 헬퍼
5. Options pattern / 옵션 패턴

### Phase 2: Middleware System (v1.11.006-010)

**Focus / 초점**: Built-in middleware and middleware chain

**Components / 컴포넌트**:
1. Middleware chain mechanism / 미들웨어 체인 메커니즘
2. Recovery middleware / Recovery 미들웨어
3. Logger middleware / Logger 미들웨어
4. CORS middleware / CORS 미들웨어
5. Auth middleware / Auth 미들웨어

### Phase 3: Template System (v1.11.011-015)

**Focus / 초점**: Template engine with auto-discovery

**Components / 컴포넌트**:
1. Template engine core / 템플릿 엔진 핵심
2. Auto template discovery / 자동 템플릿 발견
3. Layout system / 레이아웃 시스템
4. Custom template functions / 사용자 정의 템플릿 함수
5. Hot reload / 핫 리로드

### Phase 4: Advanced Features (v1.11.016-020)

**Focus / 초점**: Request/Response utilities

**Components / 컴포넌트**:
1. Request binding (JSON, Form) / 요청 바인딩
2. Cookie helpers / 쿠키 헬퍼
3. Header helpers / 헤더 헬퍼
4. File upload / 파일 업로드
5. Static file serving / 정적 파일 서빙

### Phase 5: Server Management (v1.11.021-025)

**Focus / 초점**: Server lifecycle and utilities

**Components / 컴포넌트**:
1. Graceful shutdown / 우아한 종료
2. Health check endpoint / 헬스 체크 엔드포인트
3. Route groups / 라우트 그룹
4. Error handling / 에러 처리
5. Server utilities / 서버 유틸리티

### Phase 6: Documentation & Polish (v1.11.026-030)

**Focus / 초점**: Documentation, examples, testing

**Components / 컴포넌트**:
1. USER_MANUAL.md / 사용자 매뉴얼
2. DEVELOPER_GUIDE.md / 개발자 가이드
3. Comprehensive examples / 포괄적인 예제
4. Benchmark tests / 벤치마크 테스트
5. Final polish / 최종 마무리

---

## Task Breakdown / 작업 분류

### Phase 1: Core Foundation (v1.11.001-005)

#### v1.11.001 - Project Setup ✅
- [x] Create package structure / 패키지 구조 생성
- [x] Create DESIGN_PLAN.md / 설계 계획 작성
- [x] Create WORK_PLAN.md / 작업 계획 작성
- [x] Create websvrutil.go (package info) / 패키지 정보 파일 생성

#### v1.11.002 - App & Options
- [ ] Create `app.go` (App struct) / App 구조체 생성
- [ ] Create `options.go` (Options pattern) / 옵션 패턴 생성
- [ ] Implement `New()` constructor / 생성자 구현
- [ ] Implement `Run()` method / Run 메서드 구현
- [ ] Implement `Shutdown()` method / Shutdown 메서드 구현
- [ ] Write tests for App / App 테스트 작성
- [ ] Update documentation / 문서 업데이트

#### v1.11.003 - Router
- [ ] Create `router.go` (Router struct) / Router 구조체 생성
- [ ] Implement route registration / 라우트 등록 구현
  - [ ] `GET()`, `POST()`, `PUT()`, `PATCH()`, `DELETE()`
- [ ] Implement route matching / 라우트 매칭 구현
  - [ ] Static routes / 정적 라우트
  - [ ] Path parameters (`:id`) / 경로 매개변수
  - [ ] Wildcard (`*`) / 와일드카드
- [ ] Write tests for Router / Router 테스트 작성
- [ ] Update documentation / 문서 업데이트

#### v1.11.004 - Context (Part 1)
- [ ] Create `context.go` (Context struct) / Context 구조체 생성
- [ ] Implement request methods / 요청 메서드 구현
  - [ ] `Param()`, `QueryParam()`, `QueryParamDefault()`
  - [ ] `Header()`, `SetHeader()`
  - [ ] `Cookie()`, `SetCookie()`
- [ ] Write tests for Context request methods / Context 요청 메서드 테스트 작성
- [ ] Update documentation / 문서 업데이트

#### v1.11.005 - Response Helpers
- [ ] Create `response.go` / response.go 생성
- [ ] Implement response methods / 응답 메서드 구현
  - [ ] `JSON()` - JSON response / JSON 응답
  - [ ] `String()` - Text response / 텍스트 응답
  - [ ] `Data()` - Binary data / 바이너리 데이터
  - [ ] `Redirect()` - Redirect / 리디렉션
  - [ ] `Error()` - Error response / 에러 응답
- [ ] Write tests for response helpers / 응답 헬퍼 테스트 작성
- [ ] Create simple example / 간단한 예제 생성
- [ ] Update documentation / 문서 업데이트

---

### Phase 2: Middleware System (v1.11.006-010)

#### v1.11.006 - Middleware Chain
- [ ] Create `middleware.go` / middleware.go 생성
- [ ] Implement middleware chain mechanism / 미들웨어 체인 메커니즘 구현
  - [ ] `MiddlewareFunc` type / MiddlewareFunc 타입
  - [ ] `Use()` method / Use 메서드
  - [ ] `Next()` in Context / Context의 Next
  - [ ] `Abort()` in Context / Context의 Abort
- [ ] Write tests for middleware chain / 미들웨어 체인 테스트 작성
- [ ] Update documentation / 문서 업데이트

#### v1.11.007 - Recovery Middleware
- [ ] Implement `Recovery()` middleware / Recovery 미들웨어 구현
  - [ ] Panic recovery / 패닉 복구
  - [ ] Error logging / 에러 로깅
  - [ ] Custom error response / 사용자 정의 에러 응답
- [ ] Write tests for Recovery / Recovery 테스트 작성
- [ ] Update documentation / 문서 업데이트

#### v1.11.008 - Logger Middleware
- [ ] Implement `Logger()` middleware / Logger 미들웨어 구현
  - [ ] Request logging / 요청 로깅
  - [ ] Response time / 응답 시간
  - [ ] Status code / 상태 코드
  - [ ] Integration with logging package / logging 패키지 통합
- [ ] Write tests for Logger / Logger 테스트 작성
- [ ] Update documentation / 문서 업데이트

#### v1.11.009 - CORS Middleware
- [ ] Implement `CORS()` middleware / CORS 미들웨어 구현
  - [ ] Default CORS settings / 기본 CORS 설정
  - [ ] Custom CORS options / 사용자 정의 CORS 옵션
  - [ ] Preflight handling / 프리플라이트 처리
- [ ] Write tests for CORS / CORS 테스트 작성
- [ ] Update documentation / 문서 업데이트

#### v1.11.010 - Auth Middleware
- [ ] Implement `Auth()` middleware / Auth 미들웨어 구현
  - [ ] Token validation / 토큰 검증
  - [ ] Bearer token support / Bearer 토큰 지원
  - [ ] Custom validator / 사용자 정의 검증자
- [ ] Write tests for Auth / Auth 테스트 작성
- [ ] Create middleware example / 미들웨어 예제 생성
- [ ] Update documentation / 문서 업데이트

---

### Phase 3: Template System (v1.11.011-015)

#### v1.11.011 - Template Engine Core
- [ ] Create `template.go` (TemplateEngine struct) / TemplateEngine 구조체 생성
- [ ] Implement template loading / 템플릿 로딩 구현
- [ ] Implement template rendering / 템플릿 렌더링 구현
- [ ] Write tests for template engine / 템플릿 엔진 테스트 작성
- [ ] Update documentation / 문서 업데이트

#### v1.11.012 - Auto Template Discovery
- [ ] Implement auto-discovery of templates / 템플릿 자동 발견 구현
- [ ] Support nested directories / 중첩 디렉토리 지원
- [ ] Cache compiled templates / 컴파일된 템플릿 캐시
- [ ] Write tests for auto-discovery / 자동 발견 테스트 작성
- [ ] Update documentation / 문서 업데이트

#### v1.11.013 - Layout System
- [ ] Implement layout support / 레이아웃 지원 구현
- [ ] Base layout template / 기본 레이아웃 템플릿
- [ ] Partial templates / 파셜 템플릿
- [ ] Template inheritance / 템플릿 상속
- [ ] Write tests for layout system / 레이아웃 시스템 테스트 작성
- [ ] Update documentation / 문서 업데이트

#### v1.11.014 - Custom Template Functions
- [ ] Implement `AddTemplateFunc()` / AddTemplateFunc 구현
- [ ] Add built-in template functions / 내장 템플릿 함수 추가
  - [ ] `formatDate`, `upper`, `lower`, `safeHTML`
- [ ] Write tests for template functions / 템플릿 함수 테스트 작성
- [ ] Update documentation / 문서 업데이트

#### v1.11.015 - Hot Reload
- [ ] Implement hot reload in development / 개발 시 핫 리로드 구현
- [ ] File watcher / 파일 감시자
- [ ] Auto recompile on change / 변경 시 자동 재컴파일
- [ ] Write tests for hot reload / 핫 리로드 테스트 작성
- [ ] Create template example / 템플릿 예제 생성
- [ ] Update documentation / 문서 업데이트

---

### Phase 4: Advanced Features (v1.11.016-020)

#### v1.11.016 - Request Binding
- [ ] Implement `Bind()` method / Bind 메서드 구현
- [ ] Implement `BindJSON()` / BindJSON 구현
- [ ] Implement `BindForm()` / BindForm 구현
- [ ] Implement `BindQuery()` / BindQuery 구현
- [ ] Write tests for request binding / 요청 바인딩 테스트 작성
- [ ] Update documentation / 문서 업데이트

#### v1.11.017 - Cookie & Header Helpers
- [ ] Enhanced cookie methods / 향상된 쿠키 메서드
  - [ ] `GetCookie()`, `SetCookie()`, `DeleteCookie()`
- [ ] Enhanced header methods / 향상된 헤더 메서드
  - [ ] `GetHeader()`, `SetHeader()`, `AddHeader()`
- [ ] Write tests for helpers / 헬퍼 테스트 작성
- [ ] Update documentation / 문서 업데이트

#### v1.11.018 - File Upload
- [ ] Implement `FormFile()` / FormFile 구현
- [ ] Implement `MultipartForm()` / MultipartForm 구현
- [ ] Implement `SaveUploadedFile()` / SaveUploadedFile 구현
- [ ] File size limits / 파일 크기 제한
- [ ] Write tests for file upload / 파일 업로드 테스트 작성
- [ ] Update documentation / 문서 업데이트

#### v1.11.019 - Static File Serving
- [ ] Implement `Static()` method / Static 메서드 구현
- [ ] Implement `File()` in Context / Context의 File
- [ ] File system abstraction / 파일 시스템 추상화
- [ ] Write tests for static files / 정적 파일 테스트 작성
- [ ] Update documentation / 문서 업데이트

#### v1.11.020 - Context Storage
- [ ] Implement `Set()`, `Get()`, `MustGet()` / Set, Get, MustGet 구현
- [ ] Context value storage / Context 값 저장소
- [ ] Middleware data passing / 미들웨어 데이터 전달
- [ ] Write tests for context storage / Context 저장소 테스트 작성
- [ ] Create advanced features example / 고급 기능 예제 생성
- [ ] Update documentation / 문서 업데이트

---

### Phase 5: Server Management (v1.11.021-025)

#### v1.11.021 - Graceful Shutdown
- [ ] Implement graceful shutdown / 우아한 종료 구현
- [ ] Signal handling (SIGINT, SIGTERM) / 시그널 처리
- [ ] Connection draining / 연결 드레이닝
- [ ] Shutdown timeout / 종료 타임아웃
- [ ] Write tests for graceful shutdown / 우아한 종료 테스트 작성
- [ ] Update documentation / 문서 업데이트

#### v1.11.022 - Health Check
- [ ] Implement `Health()` endpoint / Health 엔드포인트 구현
- [ ] Custom health checks / 사용자 정의 헬스 체크
- [ ] Readiness probe / Readiness 프로브
- [ ] Liveness probe / Liveness 프로브
- [ ] Write tests for health check / 헬스 체크 테스트 작성
- [ ] Update documentation / 문서 업데이트

#### v1.11.023 - Route Groups
- [ ] Implement `Group()` method / Group 메서드 구현
- [ ] Prefix-based grouping / 접두사 기반 그룹화
- [ ] Group-level middleware / 그룹 레벨 미들웨어
- [ ] Nested groups / 중첩 그룹
- [ ] Write tests for route groups / 라우트 그룹 테스트 작성
- [ ] Update documentation / 문서 업데이트

#### v1.11.024 - Error Handling
- [ ] Create `errors.go` (HTTPError) / errors.go 생성
- [ ] Implement error types / 에러 타입 구현
- [ ] Custom error handler / 사용자 정의 에러 핸들러
- [ ] Error response formatting / 에러 응답 포맷팅
- [ ] Write tests for error handling / 에러 처리 테스트 작성
- [ ] Update documentation / 문서 업데이트

#### v1.11.025 - Server Utilities
- [ ] Implement `HTML()` method (template rendering) / HTML 메서드 구현
- [ ] Rate limiting middleware / Rate limiting 미들웨어
- [ ] Timeout middleware / Timeout 미들웨어
- [ ] Compression middleware / Compression 미들웨어
- [ ] Write tests for server utilities / 서버 유틸리티 테스트 작성
- [ ] Create server management example / 서버 관리 예제 생성
- [ ] Update documentation / 문서 업데이트

---

### Phase 6: Documentation & Polish (v1.11.026-030)

#### v1.11.026 - USER_MANUAL.md
- [ ] Create USER_MANUAL.md / 사용자 매뉴얼 생성
  - [ ] Introduction / 소개
  - [ ] Installation / 설치
  - [ ] Quick Start / 빠른 시작
  - [ ] API Reference / API 참조
  - [ ] Configuration / 설정
  - [ ] Middleware Guide / 미들웨어 가이드
  - [ ] Template Guide / 템플릿 가이드
  - [ ] Best Practices / 모범 사례
  - [ ] Troubleshooting / 문제 해결
  - [ ] FAQ / 자주 묻는 질문

#### v1.11.027 - DEVELOPER_GUIDE.md
- [ ] Create DEVELOPER_GUIDE.md / 개발자 가이드 생성
  - [ ] Architecture Overview / 아키텍처 개요
  - [ ] Package Structure / 패키지 구조
  - [ ] Core Components / 핵심 컴포넌트
  - [ ] Internal Implementation / 내부 구현
  - [ ] Design Patterns / 디자인 패턴
  - [ ] Adding New Features / 새 기능 추가
  - [ ] Testing Guide / 테스트 가이드
  - [ ] Performance / 성능
  - [ ] Contributing / 기여

#### v1.11.028 - Comprehensive Examples
- [ ] Create REST API example / REST API 예제 생성
- [ ] Create web application example / 웹 애플리케이션 예제 생성
- [ ] Create middleware example / 미들웨어 예제 생성
- [ ] Create template example / 템플릿 예제 생성
- [ ] Create full-featured example / 전체 기능 예제 생성
- [ ] All examples with detailed logging / 모든 예제에 상세 로깅

#### v1.11.029 - Testing & Benchmarks
- [ ] Ensure 80%+ test coverage / 80% 이상 테스트 커버리지 확인
- [ ] Add benchmark tests / 벤치마크 테스트 추가
  - [ ] Router benchmarks / 라우터 벤치마크
  - [ ] Context benchmarks / Context 벤치마크
  - [ ] Middleware benchmarks / 미들웨어 벤치마크
  - [ ] Template benchmarks / 템플릿 벤치마크
- [ ] Performance comparison / 성능 비교
- [ ] Update documentation with benchmarks / 벤치마크로 문서 업데이트

#### v1.11.030 - Final Polish
- [ ] Code review and refactoring / 코드 리뷰 및 리팩토링
- [ ] Update all documentation / 모든 문서 업데이트
- [ ] Update README.md / README.md 업데이트
- [ ] Update CHANGELOG.md / CHANGELOG.md 업데이트
- [ ] Final testing / 최종 테스트
- [ ] Prepare for v1.11.0 release / v1.11.0 릴리스 준비

---

## Progress Tracking / 진행 상황 추적

### Phase 1: Core Foundation (v1.11.001-005)

| Version | Task | Status | Completed |
|---------|------|--------|-----------|
| v1.11.001 | Project Setup | ✅ Done | 2025-10-16 |
| v1.11.002 | App & Options | 📝 Planned | - |
| v1.11.003 | Router | 📝 Planned | - |
| v1.11.004 | Context (Part 1) | 📝 Planned | - |
| v1.11.005 | Response Helpers | 📝 Planned | - |

### Phase 2: Middleware System (v1.11.006-010)

| Version | Task | Status | Completed |
|---------|------|--------|-----------|
| v1.11.006 | Middleware Chain | 📝 Planned | - |
| v1.11.007 | Recovery Middleware | 📝 Planned | - |
| v1.11.008 | Logger Middleware | 📝 Planned | - |
| v1.11.009 | CORS Middleware | 📝 Planned | - |
| v1.11.010 | Auth Middleware | 📝 Planned | - |

### Phase 3: Template System (v1.11.011-015)

| Version | Task | Status | Completed |
|---------|------|--------|-----------|
| v1.11.011 | Template Engine Core | 📝 Planned | - |
| v1.11.012 | Auto Template Discovery | 📝 Planned | - |
| v1.11.013 | Layout System | 📝 Planned | - |
| v1.11.014 | Custom Template Functions | 📝 Planned | - |
| v1.11.015 | Hot Reload | 📝 Planned | - |

### Phase 4: Advanced Features (v1.11.016-020)

| Version | Task | Status | Completed |
|---------|------|--------|-----------|
| v1.11.016 | Request Binding | 📝 Planned | - |
| v1.11.017 | Cookie & Header Helpers | 📝 Planned | - |
| v1.11.018 | File Upload | 📝 Planned | - |
| v1.11.019 | Static File Serving | 📝 Planned | - |
| v1.11.020 | Context Storage | 📝 Planned | - |

### Phase 5: Server Management (v1.11.021-025)

| Version | Task | Status | Completed |
|---------|------|--------|-----------|
| v1.11.021 | Graceful Shutdown | 📝 Planned | - |
| v1.11.022 | Health Check | 📝 Planned | - |
| v1.11.023 | Route Groups | 📝 Planned | - |
| v1.11.024 | Error Handling | 📝 Planned | - |
| v1.11.025 | Server Utilities | 📝 Planned | - |

### Phase 6: Documentation & Polish (v1.11.026-030)

| Version | Task | Status | Completed |
|---------|------|--------|-----------|
| v1.11.026 | USER_MANUAL.md | 📝 Planned | - |
| v1.11.027 | DEVELOPER_GUIDE.md | 📝 Planned | - |
| v1.11.028 | Comprehensive Examples | 📝 Planned | - |
| v1.11.029 | Testing & Benchmarks | 📝 Planned | - |
| v1.11.030 | Final Polish | 📝 Planned | - |

---

## Estimated Timeline / 예상 일정

- **Phase 1** (5 tasks): ~1 week / 약 1주
- **Phase 2** (5 tasks): ~1 week / 약 1주
- **Phase 3** (5 tasks): ~1 week / 약 1주
- **Phase 4** (5 tasks): ~1 week / 약 1주
- **Phase 5** (5 tasks): ~1 week / 약 1주
- **Phase 6** (5 tasks): ~1 week / 약 1주

**Total Estimated Time / 총 예상 시간**: ~6 weeks / 약 6주

---

## Notes / 참고사항

1. **Follow PACKAGE_DEVELOPMENT_GUIDE.md** / PACKAGE_DEVELOPMENT_GUIDE.md 따르기
   - Increment patch version before each task / 각 작업 전에 패치 버전 증가
   - Write comprehensive tests / 포괄적인 테스트 작성
   - Update documentation / 문서 업데이트
   - Commit after each unit task / 각 단위 작업 후 커밋

2. **Testing Priority / 테스트 우선순위**
   - Aim for 80%+ coverage / 80% 이상 커버리지 목표
   - Test all edge cases / 모든 엣지 케이스 테스트
   - Add benchmarks for performance / 성능을 위한 벤치마크 추가

3. **Documentation Priority / 문서 우선순위**
   - All documentation in bilingual (English/Korean) / 모든 문서 이중 언어
   - Example code with detailed logging / 상세 로깅이 있는 예제 코드
   - Update docs after each feature / 각 기능 후 문서 업데이트

4. **Performance Considerations / 성능 고려사항**
   - Developer convenience first / 개발자 편의성 우선
   - Performance optimization secondary / 성능 최적화는 부차적
   - Benchmark critical paths / 중요 경로 벤치마크

---

## Conclusion / 결론

This work plan provides a structured approach to developing the `websvrutil` package with focus on **developer convenience**.

이 작업 계획은 **개발자 편의성**에 초점을 맞춘 `websvrutil` 패키지 개발을 위한 구조화된 접근 방식을 제공합니다.

Each phase builds upon the previous one, ensuring a solid foundation while maintaining simplicity.

각 단계는 이전 단계를 기반으로 하여, 간결함을 유지하면서 견고한 기반을 보장합니다.

**Let's make web server development enjoyable! / 웹 서버 개발을 즐겁게 만들어봅시다!** 🚀
