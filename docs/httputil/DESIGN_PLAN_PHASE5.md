# httputil Package - Phase 5 Design Plan
# httputil 패키지 - Phase 5 설계 계획

**Version / 버전**: v1.11.x
**Date / 날짜**: 2025-10-15
**Status / 상태**: Planning / 계획 중

---

## Table of Contents / 목차

1. [Overview / 개요](#1-overview--개요)
2. [Design Goals / 설계 목표](#2-design-goals--설계-목표)
3. [Phase 5a: Core Advanced Features](#3-phase-5a-core-advanced-features)
4. [Phase 5b: Resilience Features](#4-phase-5b-resilience-features)
5. [Phase 5c: Monitoring & Streaming](#5-phase-5c-monitoring--streaming)
6. [Phase 5d: Specialized Features](#6-phase-5d-specialized-features)
7. [Architecture Decisions / 아키텍처 결정](#7-architecture-decisions--아키텍처-결정)
8. [Implementation Order / 구현 순서](#8-implementation-order--구현-순서)

---

## 1. Overview / 개요

### 1.1 Purpose / 목적

Phase 5 extends httputil with advanced enterprise-grade features for production use cases including:
- Cookie management and session handling
- Request/response interceptors (middleware)
- Batch requests for parallel processing
- Request caching with TTL
- Proxy support (HTTP/SOCKS5)
- Circuit breaker for resilience
- OAuth2 authentication helper
- Request metrics and monitoring
- Streaming support (SSE, large files)
- GraphQL support
- WebSocket client
- Request mocking for testing

Phase 5는 다음을 포함한 프로덕션 사용 사례를 위한 고급 엔터프라이즈급 기능으로 httputil을 확장합니다:
- 쿠키 관리 및 세션 처리
- 요청/응답 인터셉터 (미들웨어)
- 병렬 처리를 위한 배치 요청
- TTL이 있는 요청 캐싱
- 프록시 지원 (HTTP/SOCKS5)
- 복원력을 위한 서킷 브레이커
- OAuth2 인증 헬퍼
- 요청 메트릭 및 모니터링
- 스트리밍 지원 (SSE, 대용량 파일)
- GraphQL 지원
- WebSocket 클라이언트
- 테스트를 위한 요청 모킹

### 1.2 Target Version / 목표 버전

- **v1.11.001**: Phase 5a (Cookie, Interceptors, Batch)
- **v1.11.002**: Phase 5b (Cache, Proxy, Circuit Breaker)
- **v1.11.003**: Phase 5c (OAuth2, Metrics, Streaming)
- **v1.11.004**: Phase 5d (GraphQL, WebSocket, Mocking)

---

## 2. Design Goals / 설계 목표

### 2.1 Core Principles / 핵심 원칙

1. **Backward Compatibility / 하위 호환성**
   - All Phase 1-4 features remain unchanged
   - New features are opt-in via options
   - 모든 Phase 1-4 기능은 변경되지 않음
   - 새 기능은 옵션을 통해 선택 가능

2. **Zero External Dependencies / 제로 외부 의존성**
   - Continue using only standard library
   - WebSocket: use golang.org/x/net/websocket (minimal)
   - 표준 라이브러리만 계속 사용
   - WebSocket: golang.org/x/net/websocket 사용 (최소)

3. **Performance / 성능**
   - Minimal overhead for unused features
   - Efficient caching and batching
   - 사용하지 않는 기능에 대한 최소 오버헤드
   - 효율적인 캐싱 및 배칭

4. **Simplicity / 간결함**
   - Maintain "30 lines → 2 lines" philosophy
   - Intuitive API design
   - "30줄 → 2줄" 철학 유지
   - 직관적인 API 설계

---

## 3. Phase 5a: Core Advanced Features

### 3.1 Cookie Management (cookie.go)

**Purpose / 목적**: Automatic cookie persistence and session management
**목적**: 자동 쿠키 지속성 및 세션 관리

**Design / 설계**:
```go
// Cookie jar with persistence
type CookieJar struct {
    jar      http.CookieJar
    filePath string  // For persistence
    mu       sync.RWMutex
}

// Client methods
func (c *Client) GetCookies(u *url.URL) []*http.Cookie
func (c *Client) SetCookie(u *url.URL, cookie *http.Cookie)
func (c *Client) ClearCookies()
func (c *Client) SaveCookies() error
func (c *Client) LoadCookies() error

// Options
func WithCookieJar() Option
func WithPersistentCookies(filePath string) Option
```

**Key Features / 주요 기능**:
- Automatic cookie storage per domain
- Optional persistence to file (JSON)
- Thread-safe operations
- 도메인별 자동 쿠키 저장
- 파일로 선택적 지속성 (JSON)
- 스레드 안전 작업

**Use Cases / 사용 사례**:
- Session-based authentication
- Shopping cart persistence
- Multi-step form workflows
- 세션 기반 인증
- 쇼핑 카트 지속성
- 다단계 폼 워크플로우

---

### 3.2 Request/Response Interceptors (interceptor.go)

**Purpose / 목적**: Middleware pattern for request/response processing
**목적**: 요청/응답 처리를 위한 미들웨어 패턴

**Design / 설계**:
```go
// Interceptor types
type RequestInterceptor func(*http.Request) error
type ResponseInterceptor func(*Response) error

// Interceptor chain
type InterceptorChain struct {
    requestInterceptors  []RequestInterceptor
    responseInterceptors []ResponseInterceptor
}

// Client integration
func WithRequestInterceptor(interceptor RequestInterceptor) Option
func WithResponseInterceptor(interceptor ResponseInterceptor) Option

// Built-in interceptors
func LoggingInterceptor() RequestInterceptor
func RateLimitInterceptor(limit int, window time.Duration) RequestInterceptor
func RetryInterceptor(maxRetries int) ResponseInterceptor
```

**Key Features / 주요 기능**:
- Multiple interceptors chaining
- Request modification before sending
- Response inspection before returning
- Built-in common interceptors
- 여러 인터셉터 체이닝
- 전송 전 요청 수정
- 반환 전 응답 검사
- 내장 일반 인터셉터

**Use Cases / 사용 사례**:
- Automatic token refresh
- Request/response logging
- Custom authentication
- Rate limiting
- 자동 토큰 갱신
- 요청/응답 로깅
- 커스텀 인증
- 속도 제한

---

### 3.3 Batch Requests (batch.go)

**Purpose / 목적**: Send multiple requests concurrently
**목적**: 여러 요청을 동시에 전송

**Design / 설계**:
```go
// Batch request definition
type BatchRequest struct {
    ID     string        // Optional identifier
    Method string
    URL    string
    Body   interface{}
    Opts   []Option
}

// Batch result
type BatchResult struct {
    ID       string
    Response *Response
    Error    error
    Duration time.Duration
}

// Client methods
func (c *Client) Batch(requests []BatchRequest) ([]BatchResult, error)
func (c *Client) BatchContext(ctx context.Context, requests []BatchRequest) ([]BatchResult, error)
func (c *Client) BatchWithLimit(requests []BatchRequest, maxConcurrent int) ([]BatchResult, error)

// Utility
func NewBatchRequest(method, url string, body interface{}) BatchRequest
```

**Key Features / 주요 기능**:
- Concurrent request execution
- Configurable concurrency limit
- Individual error handling
- Request timing per request
- 동시 요청 실행
- 설정 가능한 동시성 제한
- 개별 에러 처리
- 요청별 타이밍

**Use Cases / 사용 사례**:
- Fetching multiple resources at once
- Parallel API calls
- Dashboard data aggregation
- 한 번에 여러 리소스 가져오기
- 병렬 API 호출
- 대시보드 데이터 집계

---

## 4. Phase 5b: Resilience Features

### 4.1 Request Caching (cache.go)

**Purpose / 목적**: HTTP cache with TTL for reducing API calls
**목적**: API 호출을 줄이기 위한 TTL이 있는 HTTP 캐시

**Design / 설계**:
```go
// Cache entry
type CacheEntry struct {
    Response  *Response
    ExpiresAt time.Time
    ETag      string
    LastMod   string
}

// Cache implementation
type Cache struct {
    entries  map[string]*CacheEntry
    maxSize  int
    ttl      time.Duration
    mu       sync.RWMutex
    hits     int64
    misses   int64
}

// Cache configuration
type CacheConfig struct {
    TTL           time.Duration
    MaxSize       int
    RespectHeader bool  // Respect Cache-Control headers
}

// Client methods
func WithCache(config CacheConfig) Option
func (c *Client) ClearCache()
func (c *Client) GetCacheStats() CacheStats

// Cache stats
type CacheStats struct {
    Hits        int64
    Misses      int64
    Size        int
    HitRate     float64
}
```

**Key Features / 주요 기능**:
- In-memory LRU cache
- Configurable TTL per request
- Respect HTTP Cache-Control headers
- ETag and Last-Modified support
- Cache statistics
- 인메모리 LRU 캐시
- 요청별 설정 가능한 TTL
- HTTP Cache-Control 헤더 존중
- ETag 및 Last-Modified 지원
- 캐시 통계

**Use Cases / 사용 사례**:
- Reduce API rate limiting
- Improve performance
- Offline-first applications
- API 속도 제한 감소
- 성능 향상
- 오프라인 우선 애플리케이션

---

### 4.2 Proxy Support (proxy.go)

**Purpose / 목적**: HTTP/SOCKS5 proxy support
**목적**: HTTP/SOCKS5 프록시 지원

**Design / 설계**:
```go
// Proxy configuration
type ProxyConfig struct {
    URL      string
    Username string
    Password string
    NoProxy  []string  // Domains to bypass proxy
}

// Client methods
func WithProxy(proxyURL string) Option
func WithProxyAuth(proxyURL, username, password string) Option
func WithProxyFromEnv() Option  // Read from HTTP_PROXY, HTTPS_PROXY env
func WithNoProxy(domains ...string) Option

// Proxy utilities
func ParseProxyURL(proxyURL string) (*ProxyConfig, error)
func IsProxyRequired(url string, noProxy []string) bool
```

**Key Features / 주요 기능**:
- HTTP and HTTPS proxy
- SOCKS5 proxy support
- Proxy authentication
- No-proxy list for bypassing
- Environment variable support
- HTTP 및 HTTPS 프록시
- SOCKS5 프록시 지원
- 프록시 인증
- 우회를 위한 No-proxy 목록
- 환경 변수 지원

**Use Cases / 사용 사례**:
- Corporate networks
- Privacy and anonymity
- Geo-restricted APIs
- 기업 네트워크
- 개인정보 보호 및 익명성
- 지역 제한 API

---

### 4.3 Circuit Breaker (circuitbreaker.go)

**Purpose / 목적**: Prevent cascading failures in microservices
**목적**: 마이크로서비스에서 연쇄 실패 방지

**Design / 설계**:
```go
// Circuit breaker states
type CircuitState int
const (
    StateClosed CircuitState = iota  // Normal operation
    StateOpen                         // Failing, reject requests
    StateHalfOpen                     // Testing recovery
)

// Circuit breaker configuration
type CircuitBreakerConfig struct {
    FailureThreshold int           // Failures before opening
    SuccessThreshold int           // Successes before closing
    Timeout          time.Duration // Time in open state
    OnStateChange    func(from, to CircuitState)
}

// Circuit breaker
type CircuitBreaker struct {
    state            CircuitState
    failures         int
    successes        int
    lastFailureTime  time.Time
    config           CircuitBreakerConfig
    mu               sync.RWMutex
}

// Client methods
func WithCircuitBreaker(config CircuitBreakerConfig) Option
func (c *Client) GetCircuitState() CircuitState
func (c *Client) ResetCircuit()
```

**Key Features / 주요 기능**:
- Three-state circuit breaker (Closed, Open, Half-Open)
- Configurable thresholds
- Automatic recovery testing
- State change callbacks
- 3상태 서킷 브레이커 (Closed, Open, Half-Open)
- 설정 가능한 임계값
- 자동 복구 테스트
- 상태 변경 콜백

**Use Cases / 사용 사례**:
- Microservices resilience
- Prevent cascade failures
- Graceful degradation
- 마이크로서비스 복원력
- 연쇄 실패 방지
- 우아한 성능 저하

---

## 5. Phase 5c: Monitoring & Streaming

### 5.1 OAuth2 Helper (oauth.go)

**Purpose / 목적**: OAuth2 authentication flow automation
**목적**: OAuth2 인증 흐름 자동화

**Design / 설계**:
```go
// OAuth2 configuration
type OAuth2Config struct {
    ClientID     string
    ClientSecret string
    TokenURL     string
    AuthURL      string
    RedirectURL  string
    Scopes       []string
}

// Token storage
type OAuth2Token struct {
    AccessToken  string
    RefreshToken string
    TokenType    string
    ExpiresAt    time.Time
}

// Client methods
func WithOAuth2(config OAuth2Config) Option
func WithOAuth2Token(token OAuth2Token) Option
func (c *Client) RefreshToken() error
func (c *Client) GetAuthURL() string
func (c *Client) ExchangeCode(code string) error
```

**Key Features / 주요 기능**:
- Automatic token refresh
- Authorization code flow
- Client credentials flow
- Token persistence
- 자동 토큰 갱신
- 인증 코드 흐름
- 클라이언트 자격 증명 흐름
- 토큰 지속성

**Use Cases / 사용 사례**:
- OAuth2 API integration (Google, GitHub, etc.)
- Automatic token management
- Multi-tenant applications
- OAuth2 API 통합 (Google, GitHub 등)
- 자동 토큰 관리
- 멀티 테넌트 애플리케이션

---

### 5.2 Request Metrics (metrics.go)

**Purpose / 목적**: Collect and report HTTP request metrics
**목적**: HTTP 요청 메트릭 수집 및 보고

**Design / 설계**:
```go
// Metrics data
type Metrics struct {
    TotalRequests   int64
    SuccessRequests int64
    FailedRequests  int64
    AverageDuration time.Duration
    MinDuration     time.Duration
    MaxDuration     time.Duration
    TotalBytes      int64
    MethodCounts    map[string]int64
    StatusCounts    map[int]int64
}

// Client methods
func WithMetrics() Option
func (c *Client) GetMetrics() Metrics
func (c *Client) ResetMetrics()
func (c *Client) ExportMetrics(format string) ([]byte, error)  // JSON, CSV

// Metrics aggregation by URL
func (c *Client) GetMetricsByURL(url string) Metrics
```

**Key Features / 주요 기능**:
- Request counting and timing
- Success/failure tracking
- Method and status code distribution
- Data transfer tracking
- Export to JSON/CSV
- 요청 카운팅 및 타이밍
- 성공/실패 추적
- 메서드 및 상태 코드 분포
- 데이터 전송 추적
- JSON/CSV로 내보내기

**Use Cases / 사용 사례**:
- Performance monitoring
- API usage analysis
- Debugging and troubleshooting
- 성능 모니터링
- API 사용 분석
- 디버깅 및 문제 해결

---

### 5.3 Streaming Support (stream.go)

**Purpose / 목적**: Server-Sent Events and large file streaming
**목적**: Server-Sent Events 및 대용량 파일 스트리밍

**Design / 설계**:
```go
// SSE event
type SSEEvent struct {
    ID    string
    Event string
    Data  string
}

// SSE handler
type SSEHandler func(event SSEEvent) error

// Client methods for SSE
func (c *Client) StreamSSE(url string, handler SSEHandler) error
func (c *Client) StreamSSEContext(ctx context.Context, url string, handler SSEHandler) error

// Large file streaming
func (c *Client) DownloadStream(url string, writer io.Writer, progress ProgressFunc) error
func (c *Client) UploadStream(url string, reader io.Reader, size int64, progress ProgressFunc) error

// Chunked transfer
func (c *Client) PostChunked(url string, reader io.Reader, result interface{}) error
```

**Key Features / 주요 기능**:
- Server-Sent Events (SSE) support
- Chunked transfer encoding
- Memory-efficient large file handling
- Progress tracking for streams
- Server-Sent Events (SSE) 지원
- 청크 전송 인코딩
- 메모리 효율적인 대용량 파일 처리
- 스트림 진행 상황 추적

**Use Cases / 사용 사례**:
- Real-time updates (stock prices, notifications)
- Large file upload/download
- Video streaming
- 실시간 업데이트 (주식 가격, 알림)
- 대용량 파일 업로드/다운로드
- 비디오 스트리밍

---

## 6. Phase 5d: Specialized Features

### 6.1 GraphQL Support (graphql.go)

**Purpose / 목적**: GraphQL query execution
**목적**: GraphQL 쿼리 실행

**Design / 설계**:
```go
// GraphQL query
type GraphQLQuery struct {
    Query         string
    Variables     map[string]interface{}
    OperationName string
}

// GraphQL response
type GraphQLResponse struct {
    Data   interface{}
    Errors []GraphQLError
}

type GraphQLError struct {
    Message string
    Path    []string
}

// Client methods
func (c *Client) GraphQL(url string, query GraphQLQuery, result interface{}) error
func (c *Client) GraphQLContext(ctx context.Context, url string, query GraphQLQuery, result interface{}) error

// Query builder
func NewGraphQLQuery(query string) *GraphQLQuery
func (q *GraphQLQuery) WithVariables(vars map[string]interface{}) *GraphQLQuery
func (q *GraphQLQuery) WithOperation(name string) *GraphQLQuery
```

**Key Features / 주요 기능**:
- GraphQL query execution
- Variables support
- Error handling
- Query builder
- GraphQL 쿼리 실행
- 변수 지원
- 에러 처리
- 쿼리 빌더

**Use Cases / 사용 사례**:
- GraphQL API clients
- Complex data fetching
- GitHub, Shopify APIs
- GraphQL API 클라이언트
- 복잡한 데이터 가져오기
- GitHub, Shopify API

---

### 6.2 WebSocket Support (websocket.go)

**Purpose / 목적**: WebSocket client for real-time communication
**목적**: 실시간 통신을 위한 WebSocket 클라이언트

**Design / 설계**:
```go
// WebSocket connection
type WebSocket struct {
    conn   *websocket.Conn
    url    string
    mu     sync.Mutex
    closed bool
}

// Message types
type WSMessage struct {
    Type int
    Data []byte
}

// Client methods
func (c *Client) WebSocket(url string, opts ...Option) (*WebSocket, error)
func (ws *WebSocket) Send(data []byte) error
func (ws *WebSocket) SendJSON(v interface{}) error
func (ws *WebSocket) Receive() (*WSMessage, error)
func (ws *WebSocket) ReceiveJSON(v interface{}) error
func (ws *WebSocket) Close() error

// Handler pattern
type WSHandler func(*WSMessage) error
func (ws *WebSocket) Listen(handler WSHandler) error
```

**Key Features / 주요 기능**:
- Full-duplex communication
- Text and binary messages
- JSON message helpers
- Automatic reconnection
- 전이중 통신
- 텍스트 및 바이너리 메시지
- JSON 메시지 헬퍼
- 자동 재연결

**Use Cases / 사용 사례**:
- Real-time chat applications
- Live data feeds
- Gaming
- 실시간 채팅 애플리케이션
- 라이브 데이터 피드
- 게임

---

### 6.3 Request Mocking (mock.go)

**Purpose / 목적**: Mock HTTP responses for testing
**목적**: 테스트를 위한 HTTP 응답 모킹

**Design / 설계**:
```go
// Mock response
type MockResponse struct {
    StatusCode int
    Body       []byte
    Headers    map[string]string
    Delay      time.Duration
}

// Mock configuration
type MockConfig struct {
    responses map[string]*MockResponse  // URL pattern -> response
    recording bool
    recordFile string
}

// Client methods
func WithMock(url string, response MockResponse) Option
func WithMockFile(filePath string) Option
func (c *Client) AddMock(pattern string, response MockResponse)
func (c *Client) RemoveMock(pattern string)
func (c *Client) ClearMocks()

// Recording for replay
func (c *Client) StartRecording(filePath string)
func (c *Client) StopRecording() error
func (c *Client) LoadRecording(filePath string) error
```

**Key Features / 주요 기능**:
- URL pattern matching
- Response delay simulation
- Record and replay
- Multiple mock responses
- URL 패턴 매칭
- 응답 지연 시뮬레이션
- 기록 및 재생
- 여러 모킹 응답

**Use Cases / 사용 사례**:
- Unit testing without backend
- Integration testing
- Development without API
- 백엔드 없이 유닛 테스트
- 통합 테스트
- API 없이 개발

---

## 7. Architecture Decisions / 아키텍처 결정

### 7.1 Design Patterns / 디자인 패턴

1. **Functional Options Pattern**
   - All features opt-in via `WithXXX()` options
   - Backward compatible
   - 모든 기능은 `WithXXX()` 옵션을 통해 선택
   - 하위 호환성

2. **Middleware/Interceptor Pattern**
   - Chain of responsibility for request/response processing
   - Extensible by users
   - 요청/응답 처리를 위한 책임 체인
   - 사용자가 확장 가능

3. **Strategy Pattern**
   - Different caching strategies
   - Different proxy types
   - 다양한 캐싱 전략
   - 다양한 프록시 타입

4. **Observer Pattern**
   - Circuit breaker state changes
   - Metrics collection
   - 서킷 브레이커 상태 변경
   - 메트릭 수집

### 7.2 Thread Safety / 스레드 안전성

All new components must be thread-safe:
- Use `sync.RWMutex` for read-heavy operations (cache, cookies)
- Use `sync.Mutex` for write-heavy operations (metrics)
- Atomic operations for counters
- 읽기 집중 작업에 `sync.RWMutex` 사용 (캐시, 쿠키)
- 쓰기 집중 작업에 `sync.Mutex` 사용 (메트릭)
- 카운터에 원자적 연산

### 7.3 Performance Considerations / 성능 고려사항

1. **Zero-cost for unused features**
   - Features only activated when options provided
   - No performance penalty for existing code
   - 옵션 제공 시에만 기능 활성화
   - 기존 코드에 성능 패널티 없음

2. **Memory efficiency**
   - LRU cache with size limits
   - Streaming for large files
   - LRU 캐시와 크기 제한
   - 대용량 파일용 스트리밍

3. **Concurrency**
   - Batch requests with configurable limits
   - Goroutine pool for connection management
   - 설정 가능한 제한이 있는 배치 요청
   - 연결 관리를 위한 고루틴 풀

---

## 8. Implementation Order / 구현 순서

### Phase 5a (v1.11.001) - Week 1
1. ✅ Cookie Management (cookie.go)
2. ✅ Request/Response Interceptors (interceptor.go)
3. ✅ Batch Requests (batch.go)
4. ✅ Tests + Documentation + Examples

### Phase 5b (v1.11.002) - Week 2
5. ✅ Request Caching (cache.go)
6. ✅ Proxy Support (proxy.go)
7. ✅ Circuit Breaker (circuitbreaker.go)
8. ✅ Tests + Documentation + Examples

### Phase 5c (v1.11.003) - Week 3
9. ✅ OAuth2 Helper (oauth.go)
10. ✅ Request Metrics (metrics.go)
11. ✅ Streaming Support (stream.go)
12. ✅ Tests + Documentation + Examples

### Phase 5d (v1.11.004) - Week 4
13. ✅ GraphQL Support (graphql.go)
14. ✅ WebSocket Support (websocket.go)
15. ✅ Request Mocking (mock.go)
16. ✅ Tests + Documentation + Examples

### Final (v1.11.005) - Week 5
17. ✅ Comprehensive integration tests
18. ✅ Complete documentation update
19. ✅ Performance benchmarks
20. ✅ Release notes

---

## 9. Success Criteria / 성공 기준

### 9.1 Code Quality / 코드 품질
- [ ] All features implemented according to design
- [ ] 100% test coverage for new code
- [ ] All tests passing
- [ ] Zero external dependencies (except websocket)
- [ ] Backward compatible with Phase 1-4

### 9.2 Documentation / 문서
- [ ] Design document (this file)
- [ ] Work plan (WORK_PLAN_PHASE5.md)
- [ ] README.md updated
- [ ] USER_MANUAL.md updated
- [ ] DEVELOPER_GUIDE.md updated
- [ ] All examples working

### 9.3 Performance / 성능
- [ ] No performance regression for existing features
- [ ] Cache hit rate > 80% in typical scenarios
- [ ] Batch requests 5x faster than sequential
- [ ] Memory usage < 100MB for typical workloads

---

## 10. Risks and Mitigations / 위험 및 완화

### 10.1 Risks / 위험
1. **Complexity Increase**: Too many features may confuse users
2. **Performance Impact**: New features may slow down existing code
3. **Breaking Changes**: Accidental API changes
4. **Testing Complexity**: More features = more test cases

### 10.2 Mitigations / 완화
1. **Clear Documentation**: Comprehensive examples and guides
2. **Opt-in Design**: Features only active when enabled
3. **Backward Compatibility Tests**: Ensure Phase 1-4 unaffected
4. **Incremental Rollout**: Release in phases for feedback

---

**Document Status / 문서 상태**: ✅ Complete / 완료
**Next Step / 다음 단계**: Create WORK_PLAN_PHASE5.md
**Approved By / 승인자**: Development Team / 개발 팀
**Date / 날짜**: 2025-10-15
