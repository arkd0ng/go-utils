# Websvrutil Package - Developer Guide / 개발자 가이드

**Package**: `github.com/arkd0ng/go-utils/websvrutil`  
**Version**: v1.11.040  
**Last Updated**: 2025-10-16

---

## Table of Contents / 목차

1. [Architecture Overview / 아키텍처 개요](#architecture-overview--아키텍처-개요)  
2. [Package Layout / 패키지 구조](#package-layout--패키지-구조)  
3. [Core Components / 핵심 컴포넌트](#core-components--핵심-컴포넌트)  
4. [Options & Configuration / 옵션 및 구성](#options--configuration--옵션-및-구성)  
5. [Middleware Internals / 미들웨어 내부 구조](#middleware-internals--미들웨어-내부-구조)  
6. [Template Engine / 템플릿 엔진](#template-engine--템플릿-엔진)  
7. [Session Store / 세션 저장소](#session-store--세션-저장소)  
8. [Testing Notes / 테스트 참고사항](#testing-notes--테스트-참고사항)  
9. [Contribution Checklist / 기여 체크리스트](#contribution-checklist--기여-체크리스트)

---

## Architecture Overview / 아키텍처 개요

`websvrutil` wraps `net/http` while keeping handlers and middleware transparent.  
`websvrutil`은 `net/http`를 래핑하면서 핸들러와 미들웨어를 투명하게 유지합니다.

```
Incoming Request / 들어오는 요청
      │
      ▼
App.ServeHTTP (store App in request context) / App.ServeHTTP (요청 컨텍스트에 App 저장)
      │
      ▼
Middleware Chain (outer → inner) / 미들웨어 체인 (외부 → 내부)
      │
      ▼
Router.ServeHTTP (method map, pattern match) / Router.ServeHTTP (메서드 맵, 패턴 매칭)
      │
      ▼
Context creation + association / 컨텍스트 생성 및 연결
      │
      ▼
Handler execution / 핸들러 실행
```

Key design notes / 설계 특징:

- Built on standard library primitives for maximum compatibility / 표준 라이브러리 기본 요소 위에 구축되어 높은 호환성을 제공합니다.
- Context objects are per-request and stored inside `http.Request.Context()` / 컨텍스트 객체는 요청마다 생성되며 `http.Request.Context()`에 저장됩니다.
- Middleware registration order is deterministic; last added runs first / 미들웨어 등록 순서는 결정적이며 마지막에 추가된 미들웨어가 먼저 실행됩니다.
- Router keeps a map of method → routes and performs sequential segment matching / 라우터는 메서드별 라우트 맵을 유지하며 순차적으로 세그먼트를 매칭합니다.
- Helpers (binding, cookies, uploads) live on `Context` to stay close to handler code / 바인딩, 쿠키, 업로드 등의 헬퍼는 `Context`에 있어 핸들러 코드와 가깝습니다.

---

## Package Layout / 패키지 구조

```
websvrutil/
├── app.go
├── bind.go
├── constants.go
├── context.go
├── context_bind.go
├── context_helpers.go
├── context_request.go
├── context_response.go
├── csrf.go
├── group.go
├── middleware.go
├── options.go
├── router.go
├── session.go
├── template.go
├── validator.go
├── websvrutil.go
└── *_test.go
```

Responsibilities / 파일 역할:

- `app.go` – application lifecycle, middleware composition, HTTP server bootstrap / 애플리케이션 생명주기, 미들웨어 구성, HTTP 서버 초기화
- `bind.go` – form/query binding via reflection helpers / 리플렉션 헬퍼를 이용한 폼/쿼리 바인딩
- `constants.go` – default timeouts, limits, header constants / 기본 타임아웃, 제한, 헤더 상수 정의
- `context*.go` – context struct plus request, response, helper methods / 컨텍스트 구조체와 요청·응답·헬퍼 메서드 구현
- `csrf.go` – CSRF middleware and token store / CSRF 미들웨어 및 토큰 저장소
- `group.go` – route grouping with middleware inheritance / 미들웨어 상속을 포함한 라우트 그룹 관리
- `middleware.go` – built-in middleware implementations / 내장 미들웨어 구현체
- `options.go` – functional options and defaults / 함수형 옵션과 기본값 정의
- `router.go` – route registration, pattern parsing, dispatch / 라우트 등록, 패턴 파싱, 요청 분배
- `session.go` – in-memory session store / 인메모리 세션 저장소
- `template.go` – HTML template engine with layouts and reload / 레이아웃 및 재로드 기능을 갖춘 HTML 템플릿 엔진
- `validator.go` – lightweight validation utilities / 경량 검증 유틸리티
- `*_test.go` – unit, integration, benchmark tests / 단위 테스트, 통합 테스트, 벤치마크 테스트

---

## Core Components / 핵심 컴포넌트

### App / 앱

```go
type App struct {
    router     http.Handler
    middleware []MiddlewareFunc
    templates  *TemplateEngine
    options    *Options
    server     *http.Server
    mu         sync.RWMutex
    running    bool
}
```

- `New(opts ...Option)` seeds defaults from `defaultOptions()` then applies overrides / `New(opts ...Option)`은 `defaultOptions()`로 기본값을 설정한 뒤 전달된 옵션을 적용합니다.
- `Use(m ...MiddlewareFunc)` appends middleware; panics if the server is already running / `Use(m ...MiddlewareFunc)`는 미들웨어를 추가하며 서버가 실행 중이면 패닉을 발생시킵니다.
- Route helpers (`GET`, `POST`, etc.) call a private `registerRoute` that updates the underlying `Router` / 라우트 헬퍼(`GET`, `POST` 등)는 내부 `registerRoute`를 호출해 실제 `Router`를 갱신합니다.
- `NotFound` swaps out the router's fallback handler / `NotFound`는 라우터의 폴백 핸들러를 교체합니다.
- `Static` registers a GET handler that serves files from disk / `Static`은 디스크의 파일을 제공하는 GET 핸들러를 등록합니다.
- `Run`, `RunWithGracefulShutdown`, and `Shutdown` manage the HTTP server lifecycle / `Run`, `RunWithGracefulShutdown`, `Shutdown`은 HTTP 서버 생명주기를 관리합니다.
- `ServeHTTP` allows `App` to satisfy `http.Handler` and stores the app pointer in request context / `ServeHTTP`는 `App`이 `http.Handler`를 구현하도록 하고 요청 컨텍스트에 App 포인터를 저장합니다.

### Router / 라우터

```go
type Router struct {
    routes          map[string][]*Route
    notFoundHandler http.HandlerFunc
    mu              sync.RWMutex
}
```

- `Handle` normalizes HTTP verbs, parses patterns (`parsePattern`), and appends to the method slice / `Handle`은 HTTP 메서드를 정규화하고 패턴을 파싱하여 메서드별 슬라이스에 추가합니다.
- `Route.match` supports literal segments, `:param`, and `*` wildcards / `Route.match`는 리터럴 세그먼트, `:param`, `*` 와일드카드를 지원합니다.
- `ServeHTTP` fetches routes for the method, matches sequentially, creates a `Context`, stores it in the request, and invokes the handler / `ServeHTTP`는 메서드별 라우트를 가져와 순차 매칭하고 `Context`를 생성해 요청에 저장한 뒤 핸들러를 호출합니다.
- Thread safety is ensured with `sync.RWMutex`, allowing concurrent reads after startup / `sync.RWMutex`로 스레드 안전성을 확보하여 시작 후 동시 읽기를 허용합니다.

### Context / 컨텍스트

```go
type Context struct {
    Request        *http.Request
    ResponseWriter http.ResponseWriter
    params         map[string]string
    values         map[string]interface{}
    app            *App
    mu             sync.RWMutex
}
```

- `NewContext` lazily allocates internal maps; `values` stays nil until `Set` is called / `NewContext`는 내부 맵을 지연 할당하며 `values`는 `Set`이 호출될 때까지 nil 상태로 유지됩니다.
- Accessors provide request data: `Param`, `Params`, `Method`, `Path`, `Query`, `QueryDefault`, `GetHeader`, etc. / 접근자 메서드로 `Param`, `Params`, `Method`, `Path`, `Query`, `QueryDefault`, `GetHeader` 등을 통해 요청 데이터를 제공합니다.
- Storage helpers (`Set`, `Get`, `MustGet`, typed getters, `Exists`, `Delete`, `Keys`) manage per-request state / 저장소 헬퍼(`Set`, `Get`, `MustGet`, 타입별 getter, `Exists`, `Delete`, `Keys`)는 요청 단위 상태를 관리합니다.
- Response helpers (`Status`, `Write`, `Text`, `JSON`, `HTML`, `XML`, `File`, `FileAttachment`, `ErrorJSON`, etc.) keep handler code concise / 응답 헬퍼(`Status`, `Write`, `Text`, `JSON`, `HTML`, `XML`, `File`, `FileAttachment`, `ErrorJSON` 등)는 핸들러 코드를 간결하게 유지합니다.
- Binding helpers (`Bind`, `BindJSON`, `BindForm`, `BindQuery`) apply body limits using `io.LimitReader` based on `Options` / 바인딩 헬퍼(`Bind`, `BindJSON`, `BindForm`, `BindQuery`)는 `Options`에 따라 `io.LimitReader`로 본문 제한을 적용합니다.
- File helpers (`FormFile`, `MultipartForm`, `SaveUploadedFile`) reuse upload size settings / 파일 헬퍼(`FormFile`, `MultipartForm`, `SaveUploadedFile`)는 업로드 크기 설정을 재사용합니다.
- `GetContext` retrieves the stored context or returns a fallback wrapper when invoked manually / `GetContext`는 저장된 컨텍스트를 가져오거나 수동 호출 시 폴백 래퍼를 반환합니다.

### Groups / 그룹

- `Group` holds a prefix, middleware slice, and pointer to the parent app / `Group`은 접두사, 미들웨어 슬라이스, 부모 앱 포인터를 보유합니다.
- Nested groups concatenate prefixes and copy middleware so inheritance is explicit / 중첩 그룹은 접두사를 이어 붙이고 미들웨어를 복사하여 상속을 명시적으로 처리합니다.
- Route registration wraps handlers with group middleware before delegating to the parent `App` / 라우트 등록 시 그룹 미들웨어로 핸들러를 감싼 후 부모 `App`에 위임합니다.

### Validator / 검증기

- `validator.go` defines rules like `required`, `email`, `min`, `max`, `len`, `oneof`, `alpha`, `numeric` / `validator.go`는 `required`, `email`, `min`, `max`, `len`, `oneof`, `alpha`, `numeric` 등의 규칙을 정의합니다.
- Validation aggregates errors into `ValidationErrors` for easy reporting / 검증 결과는 `ValidationErrors`로 모아 보고하기 쉽게 만듭니다.
- Binding code can reuse the validator to enforce struct tags / 바인딩 코드에서 검증기를 재사용해 구조체 태그를 검증할 수 있습니다.

---

## Options & Configuration / 옵션 및 구성

`Options` centralize server and framework configuration.  
`Options`는 서버와 프레임워크 구성을 한 곳에서 관리합니다.

Key fields / 주요 필드:

- `ReadTimeout`, `WriteTimeout`, `IdleTimeout` – request lifecycle limits / 요청 생명주기 제한 시간
- `MaxHeaderBytes`, `MaxBodySize`, `MaxUploadSize` – resource limits / 리소스 제한 값
- `TemplateDir`, `StaticDir`, `StaticPrefix`, `EnableAutoReload` – asset and template settings / 에셋 및 템플릿 설정
- `EnableLogger`, `EnableRecovery` – default middleware toggles / 기본 미들웨어 토글

Functional options (`WithReadTimeout`, `WithTemplateDir`, `WithMaxBodySize`, etc.) mutate the struct before the app is used.  
함수형 옵션(`WithReadTimeout`, `WithTemplateDir`, `WithMaxBodySize` 등)은 앱 사용 전에 구조체를 수정합니다.

Context helpers and middleware read these values at runtime (e.g., `BindJSON` uses `MaxBodySize`).  
컨텍스트 헬퍼와 미들웨어는 런타임에 값을 읽어 사용합니다(예: `BindJSON`은 `MaxBodySize`를 사용).

---

## Middleware Internals / 미들웨어 내부 구조

`MiddlewareFunc` is defined as `func(http.Handler) http.Handler`, matching the standard net/http style.  
`MiddlewareFunc`는 `func(http.Handler) http.Handler`로 정의되어 표준 net/http 스타일과 일치합니다.

Implementation details / 구현 세부사항:

- Middleware is stored in `App.middleware` and composed in reverse order inside `buildHandler` / 미들웨어는 `App.middleware`에 저장되고 `buildHandler` 내부에서 역순으로 조합됩니다.
- Recovery, logger, compression, and timeout middleware wrap the writer to capture status and body size / Recovery, Logger, Compression, Timeout 미들웨어는 상태와 본문 크기를 기록하기 위해 writer를 감쌉니다.
- `RateLimiter` uses tickers and mutex-protected counters for thread safety / `RateLimiter`는 티커와 mutex로 보호되는 카운터를 사용하여 스레드 안전성을 유지합니다.
- `CSRF` middleware stores tokens in a global map with periodic cleanup triggered by `init` / `CSRF` 미들웨어는 토큰을 전역 맵에 저장하고 `init`에서 주기적으로 정리합니다.
- Static and redirect middleware reuse the same response helpers exposed on `Context` / 정적 파일 및 리다이렉트 미들웨어는 `Context`의 응답 헬퍼를 재사용합니다.

---

## Template Engine / 템플릿 엔진

`TemplateEngine` builds on `html/template` and registers convenience functions.  
`TemplateEngine`은 `html/template` 위에 구축되며 편의 함수를 등록합니다.

Highlights / 주요 기능:

- `NewTemplateEngine(dir string)` loads templates relative to the root directory / `NewTemplateEngine(dir string)`은 루트 디렉터리를 기준으로 템플릿을 로드합니다.
- `AddFunc`, `AddFuncs` merge custom functions into the template set / `AddFunc`, `AddFuncs`는 사용자 정의 함수를 템플릿 세트에 병합합니다.
- Layout support via `SetLayoutDir`, `LoadLayout`, `LoadAllLayouts`, `RenderWithLayout` / `SetLayoutDir`, `LoadLayout`, `LoadAllLayouts`, `RenderWithLayout`을 통해 레이아웃을 지원합니다.
- `EnableAutoReload` starts a watcher goroutine; `DisableAutoReload` stops it / `EnableAutoReload`는 감시 고루틴을 시작하고 `DisableAutoReload`는 이를 종료합니다.
- `ReloadTemplates`, `LoadTemplate`, `LoadTemplates`, `LoadAll` provide manual refresh hooks / `ReloadTemplates`, `LoadTemplate`, `LoadTemplates`, `LoadAll`은 수동 새로 고침 기능을 제공합니다.
- `IsAutoReloadEnabled` reports whether watchers are active / `IsAutoReloadEnabled`는 감시자가 활성 상태인지 확인합니다.

---

## Session Store / 세션 저장소

`SessionStore` offers an in-memory implementation suited for development or single-instance deployments.  
`SessionStore`는 개발 및 단일 인스턴스 배포에 적합한 인메모리 구현을 제공합니다.

- `NewSessionStore` accepts `SessionOptions` and launches a cleanup goroutine / `NewSessionStore`는 `SessionOptions`를 받아 정리 고루틴을 시작합니다.
- Session IDs are 32 random bytes encoded with base64 / 세션 ID는 32바이트 랜덤 값을 base64로 인코딩합니다.
- `DefaultSessionOptions` sets cookie name, max age, SameSite, cleanup interval, path, and domain defaults / `DefaultSessionOptions`는 쿠키 이름, 최대 수명, SameSite, 정리 간격, Path, Domain 기본값을 설정합니다.
- `Get` fetches an existing session or creates a fresh one when absent / `Get`은 기존 세션을 가져오거나 없으면 새 세션을 생성합니다.
- `Save` refreshes expiration and writes the cookie; `Destroy` deletes session data and clears the cookie / `Save`는 만료 시간을 갱신하고 쿠키를 쓰며, `Destroy`는 세션 데이터를 삭제하고 쿠키를 제거합니다.
- Session helpers (`Set`, `Get`, `GetString`, `GetInt`, `GetBool`, etc.) mirror context storage patterns / 세션 헬퍼(`Set`, `Get`, `GetString`, `GetInt`, `GetBool` 등)는 컨텍스트 저장 패턴과 동일합니다.

---

## Example Suite & Logging / 예제 모음 및 로깅

`examples/websvrutil/main.go` is now a comprehensive workbook that exercises every major feature.  
`examples/websvrutil/main.go`는 주요 기능 전부를 다루는 종합 워크북 형태로 구성되어 있습니다.

- **Feature coverage / 기능 커버리지**: routing (all verbs + groups), context helpers, request binding, responses, middleware, sessions, template engine, CSRF, validator, file uploads, static files, graceful shutdown, testing scenarios.  
  **기능 커버리지**: 라우팅(모든 메서드 및 그룹), 컨텍스트 헬퍼, 요청 바인딩, 응답, 미들웨어, 세션, 템플릿 엔진, CSRF, 검증기, 파일 업로드, 정적 파일, 정상 종료, 테스트 시나리오.  
- **Log-first narrative / 로그 우선 학습**: every console line is duplicated in `logs/websvrutil-example.log` with English first and the Korean translation immediately after.  
  **로그 기반 학습**: 모든 콘솔 출력이 영어 → 한국어 순으로 `logs/websvrutil-example.log`에 기록됩니다.  
- **Structured context / 구조화된 정보**: request method, path, headers, payloads, status codes, and resulting artifacts (e.g., saved files, CSRF tokens) are logged alongside each step.  
  **구조화된 정보**: 요청 메서드·경로·헤더·페이로드·상태 코드와 결과물(저장된 파일, CSRF 토큰 등)을 단계별로 기록합니다.  
- **Shared rotation / 공용 로그 로테이션**: existing logs are backed up with timestamps (`logs/websvrutil-example-YYYYMMDD-HHMMSS.log`) and only the latest five are kept.  
  **공용 로그 로테이션**: 기존 로그는 타임스탬프 백업(`logs/websvrutil-example-YYYYMMDD-HHMMSS.log`) 후 최근 5개만 유지합니다.

Implementation notes / 구현 메모:

1. `setupLogger` ensures the `logs/` directory exists, rotates the current log, and initializes `logging.Logger` with stdout mirroring.  
   `setupLogger`는 `logs/` 디렉터리를 생성하고 현재 로그를 회전시킨 뒤 `logging.Logger`를 stdout 미러링과 함께 초기화합니다.
2. `logSection`, `logDual`, `logPrintln`, `logPrintf` helper functions centralize bilingual logging so every example step stays consistent.  
   `logSection`, `logDual`, `logPrintln`, `logPrintf` 헬퍼로 이중 언어 로그를 일관되게 출력합니다.
3. Each functional area is encapsulated in `runXYZExamples` to keep the file navigable and make it easy to cherry-pick snippets.  
   각 기능 영역은 `runXYZExamples` 함수에 캡슐화되어 있어 탐색과 코드 재사용이 용이합니다.

When authoring new examples for other packages, follow the pattern documented in `docs/EXAMPLE_CODE_GUIDE.md`—shared `logs/`, timestamped backups, and bilingual, data-rich logging.  
다른 패키지의 예제를 작성할 때도 `docs/EXAMPLE_CODE_GUIDE.md`에 정리된 패턴(공용 `logs/`, 타임스탬프 백업, 이중 언어 상세 로그)을 준수하세요.

---

## Testing Notes / 테스트 참고사항

- Tests cover routing, context helpers, middleware, template rendering, sessions, uploads, and CSRF / 테스트는 라우팅, 컨텍스트 헬퍼, 미들웨어, 템플릿 렌더링, 세션, 업로드, CSRF를 다룹니다.
- Benchmark tests (e.g., middleware, session) provide baseline performance data / 벤치마크 테스트(예: 미들웨어, 세션)는 기본 성능 데이터를 제공합니다.
- Use table-driven tests and bilingual comments to stay consistent with repository style / 테이블 기반 테스트와 한·영 병기 주석을 사용해 저장소 스타일을 유지하세요.
- Run `go test ./websvrutil/...` before submitting changes / 변경 사항을 제출하기 전에 `go test ./websvrutil/...`을 실행하세요.

```bash
go test ./websvrutil/...
```

---

## Contribution Checklist / 기여 체크리스트

1. **API Consistency / API 일관성** – follow naming conventions and add bilingual comments.  
2. **Thread Safety / 스레드 안전성** – guard shared maps and state with the appropriate mutex.  
3. **Options Integration / 옵션 연동** – expose new behavior through `Options` or dedicated config structs.  
4. **Documentation / 문서화** – update `USER_MANUAL.md`, `DEVELOPER_GUIDE.md`, and examples when APIs change.  
5. **Testing / 테스트** – add or adjust tests; ensure `go test ./websvrutil/...` passes.  
6. **Version Control / 버전 관리** – do not bump `websvrutil.go` version outside scheduled releases.

Following this checklist keeps contributions predictable and easy to review.  
이 체크리스트를 따르면 기여 내용이 예측 가능해지고 검토가 쉬워집니다.
