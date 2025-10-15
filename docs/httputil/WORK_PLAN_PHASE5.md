# httputil Package - Phase 5 Work Plan
# httputil 패키지 - Phase 5 작업 계획

**Version / 버전**: v1.11.x
**Date / 날짜**: 2025-10-15
**Estimated Duration / 예상 기간**: 5 weeks / 5주

---

## Table of Contents / 목차

1. [Overview / 개요](#1-overview--개요)
2. [Task Breakdown / 작업 분류](#2-task-breakdown--작업-분류)
3. [Phase 5a Tasks](#3-phase-5a-tasks)
4. [Phase 5b Tasks](#4-phase-5b-tasks)
5. [Phase 5c Tasks](#5-phase-5c-tasks)
6. [Phase 5d Tasks](#6-phase-5d-tasks)
7. [Documentation Tasks](#7-documentation-tasks)
8. [Testing Strategy](#8-testing-strategy)
9. [Timeline](#9-timeline)

---

## 1. Overview / 개요

This work plan outlines the detailed tasks for implementing Phase 5 features in the httputil package. Each feature will follow the same workflow:

이 작업 계획은 httputil 패키지에서 Phase 5 기능을 구현하기 위한 상세 작업을 설명합니다. 각 기능은 동일한 워크플로우를 따릅니다:

**Standard Workflow / 표준 워크플로우**:
1. Design API and data structures / API 및 데이터 구조 설계
2. Implement core functionality / 핵심 기능 구현
3. Write unit tests / 유닛 테스트 작성
4. Update documentation / 문서 업데이트
5. Create examples / 예제 생성
6. Integration testing / 통합 테스트

---

## 2. Task Breakdown / 작업 분류

### 2.1 Implementation Tasks / 구현 작업

Each feature requires:
- **File Creation**: New .go file for feature
- **Client Integration**: Add methods to Client struct
- **Options**: Create WithXXX() option function
- **Simple API**: Add package-level convenience functions
- **Error Handling**: Custom error types if needed

각 기능에는 다음이 필요합니다:
- **파일 생성**: 기능을 위한 새 .go 파일
- **클라이언트 통합**: Client 구조체에 메서드 추가
- **옵션**: WithXXX() 옵션 함수 생성
- **간단한 API**: 패키지 레벨 편의 함수 추가
- **에러 처리**: 필요시 커스텀 에러 타입

### 2.2 Testing Tasks / 테스트 작업

Each feature requires:
- **Unit Tests**: Test individual functions (80%+ coverage)
- **Integration Tests**: Test with other features
- **Error Cases**: Test all error paths
- **Benchmarks**: Performance testing

각 기능에는 다음이 필요합니다:
- **유닛 테스트**: 개별 함수 테스트 (80% 이상 커버리지)
- **통합 테스트**: 다른 기능과 함께 테스트
- **에러 케이스**: 모든 에러 경로 테스트
- **벤치마크**: 성능 테스트

### 2.3 Documentation Tasks / 문서 작업

Each feature requires:
- **README.md**: API reference section
- **USER_MANUAL.md**: Usage patterns + examples
- **DEVELOPER_GUIDE.md**: Architecture details
- **Examples**: Working code in examples/httputil/

각 기능에는 다음이 필요합니다:
- **README.md**: API 참조 섹션
- **USER_MANUAL.md**: 사용 패턴 + 예제
- **DEVELOPER_GUIDE.md**: 아키텍처 세부사항
- **예제**: examples/httputil/에 작동하는 코드

---

## 3. Phase 5a Tasks

### 3.1 Cookie Management (cookie.go)

**Estimated Time / 예상 시간**: 8 hours / 8시간

#### Implementation / 구현
- [ ] **Task 3.1.1**: Create cookie.go file (1 hour)
  - Define CookieJar struct
  - Implement GetCookies() method
  - Implement SetCookie() method
  - Implement ClearCookies() method

- [ ] **Task 3.1.2**: Add persistence support (1 hour)
  - Implement SaveCookies() - save to JSON file
  - Implement LoadCookies() - load from JSON file
  - Handle file I/O errors

- [ ] **Task 3.1.3**: Client integration (1 hour)
  - Add cookieJar field to Client struct
  - Add WithCookieJar() option
  - Add WithPersistentCookies() option
  - Update request execution to use cookie jar

- [ ] **Task 3.1.4**: Simple API (0.5 hour)
  - Add package-level cookie functions
  - GetCookies(), SetCookie(), ClearCookies()

#### Testing / 테스트
- [ ] **Task 3.1.5**: Write tests (2 hours)
  - TestCookieJar - basic operations
  - TestCookiePersistence - save/load
  - TestCookieAutomation - automatic cookie handling
  - TestCookieThreadSafety - concurrent access

#### Documentation / 문서
- [ ] **Task 3.1.6**: Update README.md (0.5 hour)
  - Add Cookie Management section
  - API reference
  - Quick example

- [ ] **Task 3.1.7**: Update USER_MANUAL.md (1 hour)
  - Usage patterns section
  - Common use cases with examples

- [ ] **Task 3.1.8**: Update DEVELOPER_GUIDE.md (1 hour)
  - Architecture section
  - Implementation details

#### Examples / 예제
- [ ] **Task 3.1.9**: Create example (1 hour)
  - Session-based authentication example
  - Cookie persistence example

**Total Time / 총 시간**: 8 hours / 8시간

---

### 3.2 Request/Response Interceptors (interceptor.go)

**Estimated Time / 예상 시간**: 10 hours / 10시간

#### Implementation / 구현
- [ ] **Task 3.2.1**: Create interceptor.go file (1.5 hours)
  - Define RequestInterceptor type
  - Define ResponseInterceptor type
  - Define InterceptorChain struct
  - Implement chain execution logic

- [ ] **Task 3.2.2**: Built-in interceptors (2 hours)
  - LoggingInterceptor() - request/response logging
  - RateLimitInterceptor() - rate limiting
  - RetryInterceptor() - retry on failure
  - TokenRefreshInterceptor() - automatic token refresh

- [ ] **Task 3.2.3**: Client integration (1 hour)
  - Add interceptor chains to Client struct
  - Add WithRequestInterceptor() option
  - Add WithResponseInterceptor() option
  - Integrate with request execution flow

- [ ] **Task 3.2.4**: Simple API (0.5 hour)
  - Add package-level interceptor functions
  - SetRequestInterceptor(), SetResponseInterceptor()

#### Testing / 테스트
- [ ] **Task 3.2.5**: Write tests (2.5 hours)
  - TestRequestInterceptor - single interceptor
  - TestResponseInterceptor - single interceptor
  - TestInterceptorChain - multiple interceptors
  - TestBuiltInInterceptors - all built-in interceptors
  - TestInterceptorErrors - error handling

#### Documentation / 문서
- [ ] **Task 3.2.6**: Update README.md (0.5 hour)
- [ ] **Task 3.2.7**: Update USER_MANUAL.md (1 hour)
- [ ] **Task 3.2.8**: Update DEVELOPER_GUIDE.md (1 hour)

#### Examples / 예제
- [ ] **Task 3.2.9**: Create examples (1 hour)
  - Logging interceptor example
  - Token refresh interceptor example
  - Custom authentication interceptor example

**Total Time / 총 시간**: 10 hours / 10시간

---

### 3.3 Batch Requests (batch.go)

**Estimated Time / 예상 시간**: 8 hours / 8시간

#### Implementation / 구현
- [ ] **Task 3.3.1**: Create batch.go file (1.5 hours)
  - Define BatchRequest struct
  - Define BatchResult struct
  - Implement Batch() method
  - Implement BatchContext() method
  - Implement BatchWithLimit() method

- [ ] **Task 3.3.2**: Concurrency management (1 hour)
  - Worker pool implementation
  - Configurable concurrency limit
  - Context cancellation support

- [ ] **Task 3.3.3**: Client integration (0.5 hour)
  - Add batch methods to Client
  - Update options if needed

- [ ] **Task 3.3.4**: Simple API (0.5 hour)
  - Add package-level batch functions

#### Testing / 테스트
- [ ] **Task 3.3.5**: Write tests (2 hours)
  - TestBatch - basic batch execution
  - TestBatchConcurrency - concurrent execution
  - TestBatchErrors - individual error handling
  - TestBatchCancellation - context cancellation
  - TestBatchTiming - performance

#### Documentation / 문서
- [ ] **Task 3.3.6**: Update README.md (0.5 hour)
- [ ] **Task 3.3.7**: Update USER_MANUAL.md (1 hour)
- [ ] **Task 3.3.8**: Update DEVELOPER_GUIDE.md (1 hour)

#### Examples / 예제
- [ ] **Task 3.3.9**: Create example (1 hour)
  - Fetching multiple resources example
  - Dashboard data aggregation example

**Total Time / 총 시간**: 8 hours / 8시간

---

## 4. Phase 5b Tasks

### 4.1 Request Caching (cache.go)

**Estimated Time / 예상 시간**: 12 hours / 12시간

#### Implementation / 구현
- [ ] **Task 4.1.1**: Create cache.go file (2 hours)
  - Define CacheEntry struct
  - Define Cache struct with LRU
  - Implement Get() method
  - Implement Set() method
  - Implement eviction logic

- [ ] **Task 4.1.2**: Cache-Control header support (1.5 hours)
  - Parse Cache-Control headers
  - Respect max-age, no-cache, no-store
  - ETag and Last-Modified support

- [ ] **Task 4.1.3**: Client integration (1 hour)
  - Add cache field to Client struct
  - Add WithCache() option
  - Integrate with request execution

- [ ] **Task 4.1.4**: Cache management (1 hour)
  - ClearCache() method
  - GetCacheStats() method
  - Export cache statistics

- [ ] **Task 4.1.5**: Simple API (0.5 hour)
  - Package-level cache functions

#### Testing / 테스트
- [ ] **Task 4.1.6**: Write tests (3 hours)
  - TestCache - basic caching
  - TestCacheTTL - expiration
  - TestCacheLRU - eviction
  - TestCacheHeaders - Cache-Control
  - TestCacheStats - statistics
  - TestCacheThreadSafety - concurrent access

#### Documentation / 문서
- [ ] **Task 4.1.7**: Update README.md (0.5 hour)
- [ ] **Task 4.1.8**: Update USER_MANUAL.md (1.5 hours)
- [ ] **Task 4.1.9**: Update DEVELOPER_GUIDE.md (1 hour)

#### Examples / 예제
- [ ] **Task 4.1.10**: Create example (1 hour)
  - API response caching example
  - Cache statistics monitoring example

**Total Time / 총 시간**: 12 hours / 12시간

---

### 4.2 Proxy Support (proxy.go)

**Estimated Time / 예상 시간**: 8 hours / 8시간

#### Implementation / 구현
- [ ] **Task 4.2.1**: Create proxy.go file (2 hours)
  - Define ProxyConfig struct
  - Implement HTTP proxy support
  - Implement HTTPS proxy support
  - Implement SOCKS5 proxy support

- [ ] **Task 4.2.2**: Proxy authentication (1 hour)
  - Add username/password support
  - Handle proxy authentication errors

- [ ] **Task 4.2.3**: No-proxy list (0.5 hour)
  - Parse no-proxy patterns
  - Check if URL should bypass proxy

- [ ] **Task 4.2.4**: Environment variables (0.5 hour)
  - Read HTTP_PROXY, HTTPS_PROXY, NO_PROXY
  - Implement WithProxyFromEnv() option

- [ ] **Task 4.2.5**: Client integration (0.5 hour)
  - Add proxy configuration to Client
  - Add WithProxy() and related options

- [ ] **Task 4.2.6**: Simple API (0.5 hour)
  - Package-level proxy functions

#### Testing / 테스트
- [ ] **Task 4.2.7**: Write tests (2 hours)
  - TestProxy - basic proxy usage
  - TestProxyAuth - authentication
  - TestNoProxy - bypass list
  - TestProxyFromEnv - environment variables

#### Documentation / 문서
- [ ] **Task 4.2.8**: Update README.md (0.5 hour)
- [ ] **Task 4.2.9**: Update USER_MANUAL.md (0.5 hour)
- [ ] **Task 4.2.10**: Update DEVELOPER_GUIDE.md (0.5 hour)

#### Examples / 예제
- [ ] **Task 4.2.11**: Create example (1 hour)
  - Corporate proxy example
  - Proxy with authentication example

**Total Time / 총 시간**: 8 hours / 8시간

---

### 4.3 Circuit Breaker (circuitbreaker.go)

**Estimated Time / 예상 시간**: 10 hours / 10시간

#### Implementation / 구현
- [ ] **Task 4.3.1**: Create circuitbreaker.go file (2.5 hours)
  - Define CircuitBreaker struct
  - Implement state machine (Closed/Open/Half-Open)
  - Implement failure tracking
  - Implement success tracking
  - Implement timeout logic

- [ ] **Task 4.3.2**: State transitions (1 hour)
  - Closed → Open transition
  - Open → Half-Open transition
  - Half-Open → Closed/Open transition

- [ ] **Task 4.3.3**: Client integration (1 hour)
  - Add circuit breaker to Client
  - Add WithCircuitBreaker() option
  - Integrate with request execution

- [ ] **Task 4.3.4**: Callbacks (0.5 hour)
  - OnStateChange callback
  - State change notifications

- [ ] **Task 4.3.5**: Simple API (0.5 hour)
  - Package-level circuit breaker functions

#### Testing / 테스트
- [ ] **Task 4.3.6**: Write tests (2.5 hours)
  - TestCircuitBreakerClosed - normal operation
  - TestCircuitBreakerOpen - failure state
  - TestCircuitBreakerHalfOpen - recovery
  - TestCircuitBreakerTransitions - state changes
  - TestCircuitBreakerCallback - callbacks

#### Documentation / 문서
- [ ] **Task 4.3.7**: Update README.md (0.5 hour)
- [ ] **Task 4.3.8**: Update USER_MANUAL.md (1 hour)
- [ ] **Task 4.3.9**: Update DEVELOPER_GUIDE.md (1 hour)

#### Examples / 예제
- [ ] **Task 4.3.10**: Create example (1 hour)
  - Microservice resilience example
  - Circuit breaker with fallback example

**Total Time / 총 시간**: 10 hours / 10시간

---

## 5. Phase 5c Tasks

### 5.1 OAuth2 Helper (oauth.go)

**Estimated Time / 예상 시간**: 12 hours / 12시간

#### Implementation / 구현
- [ ] **Task 5.1.1**: Create oauth.go file (2.5 hours)
  - Define OAuth2Config struct
  - Define OAuth2Token struct
  - Implement token storage
  - Implement token refresh logic

- [ ] **Task 5.1.2**: Authorization flows (2 hours)
  - Authorization code flow
  - Client credentials flow
  - Refresh token flow

- [ ] **Task 5.1.3**: Client integration (1 hour)
  - Add OAuth2 support to Client
  - Add WithOAuth2() option
  - Automatic token refresh

- [ ] **Task 5.1.4**: Token persistence (1 hour)
  - Save tokens to file
  - Load tokens from file

- [ ] **Task 5.1.5**: Simple API (0.5 hour)
  - Package-level OAuth2 functions

#### Testing / 테스트
- [ ] **Task 5.1.6**: Write tests (3 hours)
  - TestOAuth2AuthorizationCode
  - TestOAuth2ClientCredentials
  - TestOAuth2TokenRefresh
  - TestOAuth2TokenPersistence

#### Documentation / 문서
- [ ] **Task 5.1.7**: Update README.md (0.5 hour)
- [ ] **Task 5.1.8**: Update USER_MANUAL.md (1 hour)
- [ ] **Task 5.1.9**: Update DEVELOPER_GUIDE.md (0.5 hour)

#### Examples / 예제
- [ ] **Task 5.1.10**: Create examples (1 hour)
  - Google OAuth2 example
  - GitHub OAuth2 example

**Total Time / 총 시간**: 12 hours / 12시간

---

### 5.2 Request Metrics (metrics.go)

**Estimated Time / 예상 시간**: 8 hours / 8시간

#### Implementation / 구현
- [ ] **Task 5.2.1**: Create metrics.go file (1.5 hours)
  - Define Metrics struct
  - Implement counter tracking
  - Implement timing tracking
  - Implement aggregation

- [ ] **Task 5.2.2**: Metrics collection (1 hour)
  - Hook into request execution
  - Track success/failure
  - Track duration
  - Track bytes transferred

- [ ] **Task 5.2.3**: Metrics export (1 hour)
  - ExportMetrics() - JSON format
  - ExportMetrics() - CSV format

- [ ] **Task 5.2.4**: Client integration (0.5 hour)
  - Add WithMetrics() option
  - GetMetrics() method
  - ResetMetrics() method

- [ ] **Task 5.2.5**: Simple API (0.5 hour)
  - Package-level metrics functions

#### Testing / 테스트
- [ ] **Task 5.2.6**: Write tests (2 hours)
  - TestMetricsCollection
  - TestMetricsAggregation
  - TestMetricsExport
  - TestMetricsThreadSafety

#### Documentation / 문서
- [ ] **Task 5.2.7**: Update README.md (0.5 hour)
- [ ] **Task 5.2.8**: Update USER_MANUAL.md (0.5 hour)
- [ ] **Task 5.2.9**: Update DEVELOPER_GUIDE.md (0.5 hour)

#### Examples / 예제
- [ ] **Task 5.2.10**: Create example (1 hour)
  - Metrics monitoring example
  - Performance analysis example

**Total Time / 총 시간**: 8 hours / 8시간

---

### 5.3 Streaming Support (stream.go)

**Estimated Time / 예상 시간**: 12 hours / 12시간

#### Implementation / 구현
- [ ] **Task 5.3.1**: Create stream.go file (2.5 hours)
  - Define SSEEvent struct
  - Implement SSE parser
  - Implement StreamSSE() method
  - Implement StreamSSEContext() method

- [ ] **Task 5.3.2**: Large file streaming (2 hours)
  - DownloadStream() - memory efficient download
  - UploadStream() - memory efficient upload
  - Progress tracking

- [ ] **Task 5.3.3**: Chunked transfer (1 hour)
  - PostChunked() method
  - Chunked encoding support

- [ ] **Task 5.3.4**: Client integration (0.5 hour)
  - Add streaming methods to Client

- [ ] **Task 5.3.5**: Simple API (0.5 hour)
  - Package-level streaming functions

#### Testing / 테스트
- [ ] **Task 5.3.6**: Write tests (3 hours)
  - TestSSE - Server-Sent Events
  - TestDownloadStream - large file download
  - TestUploadStream - large file upload
  - TestChunkedTransfer - chunked encoding

#### Documentation / 문서
- [ ] **Task 5.3.7**: Update README.md (0.5 hour)
- [ ] **Task 5.3.8**: Update USER_MANUAL.md (1.5 hours)
- [ ] **Task 5.3.9**: Update DEVELOPER_GUIDE.md (0.5 hour)

#### Examples / 예제
- [ ] **Task 5.3.10**: Create examples (1 hour)
  - SSE real-time updates example
  - Large file streaming example

**Total Time / 총 시간**: 12 hours / 12시간

---

## 6. Phase 5d Tasks

### 6.1 GraphQL Support (graphql.go)

**Estimated Time / 예상 시간**: 8 hours / 8시간

#### Implementation / 구현
- [ ] **Task 6.1.1**: Create graphql.go file (2 hours)
  - Define GraphQLQuery struct
  - Define GraphQLResponse struct
  - Define GraphQLError struct
  - Implement GraphQL() method
  - Implement GraphQLContext() method

- [ ] **Task 6.1.2**: Query builder (1 hour)
  - NewGraphQLQuery() constructor
  - WithVariables() method
  - WithOperation() method

- [ ] **Task 6.1.3**: Client integration (0.5 hour)
  - Add GraphQL methods to Client

- [ ] **Task 6.1.4**: Simple API (0.5 hour)
  - Package-level GraphQL functions

#### Testing / 테스트
- [ ] **Task 6.1.5**: Write tests (2 hours)
  - TestGraphQL - basic query
  - TestGraphQLVariables - with variables
  - TestGraphQLErrors - error handling

#### Documentation / 문서
- [ ] **Task 6.1.6**: Update README.md (0.5 hour)
- [ ] **Task 6.1.7**: Update USER_MANUAL.md (1 hour)
- [ ] **Task 6.1.8**: Update DEVELOPER_GUIDE.md (0.5 hour)

#### Examples / 예제
- [ ] **Task 6.1.9**: Create example (1 hour)
  - GitHub GraphQL API example
  - Complex query example

**Total Time / 총 시간**: 8 hours / 8시간

---

### 6.2 WebSocket Support (websocket.go)

**Estimated Time / 예상 시간**: 12 hours / 12시간

**Note**: Requires golang.org/x/net/websocket dependency
**참고**: golang.org/x/net/websocket 의존성 필요

#### Implementation / 구현
- [ ] **Task 6.2.1**: Create websocket.go file (2.5 hours)
  - Define WebSocket struct
  - Define WSMessage struct
  - Implement connection management
  - Implement Send() method
  - Implement Receive() method

- [ ] **Task 6.2.2**: JSON helpers (1 hour)
  - SendJSON() method
  - ReceiveJSON() method

- [ ] **Task 6.2.3**: Handler pattern (1 hour)
  - Listen() method with handler
  - Automatic reconnection

- [ ] **Task 6.2.4**: Client integration (0.5 hour)
  - Add WebSocket() method to Client

- [ ] **Task 6.2.5**: Simple API (0.5 hour)
  - Package-level WebSocket functions

#### Testing / 테스트
- [ ] **Task 6.2.6**: Write tests (3 hours)
  - TestWebSocketConnect
  - TestWebSocketSendReceive
  - TestWebSocketJSON
  - TestWebSocketReconnect

#### Documentation / 문서
- [ ] **Task 6.2.7**: Update README.md (0.5 hour)
- [ ] **Task 6.2.8**: Update USER_MANUAL.md (2 hours)
- [ ] **Task 6.2.9**: Update DEVELOPER_GUIDE.md (1 hour)

#### Examples / 예제
- [ ] **Task 6.2.10**: Create examples (1 hour)
  - Chat client example
  - Live data feed example

**Total Time / 총 시간**: 12 hours / 12시간

---

### 6.3 Request Mocking (mock.go)

**Estimated Time / 예상 시간**: 10 hours / 10시간

#### Implementation / 구현
- [ ] **Task 6.3.1**: Create mock.go file (2 hours)
  - Define MockResponse struct
  - Define MockConfig struct
  - Implement URL pattern matching
  - Implement response selection

- [ ] **Task 6.3.2**: Recording (1.5 hours)
  - StartRecording() method
  - StopRecording() method
  - Save recordings to file

- [ ] **Task 6.3.3**: Playback (1 hour)
  - LoadRecording() method
  - Replay recorded responses

- [ ] **Task 6.3.4**: Client integration (1 hour)
  - Add WithMock() option
  - Add WithMockFile() option
  - Mock management methods

- [ ] **Task 6.3.5**: Simple API (0.5 hour)
  - Package-level mock functions

#### Testing / 테스트
- [ ] **Task 6.3.6**: Write tests (2 hours)
  - TestMock - basic mocking
  - TestMockPattern - URL patterns
  - TestRecording - record and replay
  - TestMockDelay - response delays

#### Documentation / 문서
- [ ] **Task 6.3.7**: Update README.md (0.5 hour)
- [ ] **Task 6.3.8**: Update USER_MANUAL.md (1 hour)
- [ ] **Task 6.3.9**: Update DEVELOPER_GUIDE.md (0.5 hour)

#### Examples / 예제
- [ ] **Task 6.3.10**: Create example (1 hour)
  - Unit test with mocking example
  - Record and replay example

**Total Time / 총 시간**: 10 hours / 10시간

---

## 7. Documentation Tasks

### 7.1 README.md Updates

- [ ] **Task 7.1.1**: Update Features section (1 hour)
  - Add all Phase 5 features

- [ ] **Task 7.1.2**: Add Quick Start examples (2 hours)
  - 5-10 new examples for Phase 5 features

- [ ] **Task 7.1.3**: Update API Reference (3 hours)
  - Add sections for all new features

- [ ] **Task 7.1.4**: Update Configuration Options (1 hour)
  - Document all new options

**Total Time / 총 시간**: 7 hours / 7시간

---

### 7.2 USER_MANUAL.md Updates

- [ ] **Task 7.2.1**: Update Introduction (0.5 hour)
  - Add Phase 5 features overview

- [ ] **Task 7.2.2**: Add Usage Patterns (4 hours)
  - 5.10-5.22: New subsections for Phase 5

- [ ] **Task 7.2.3**: Add Common Use Cases (4 hours)
  - 6.8-6.18: Real-world examples

- [ ] **Task 7.2.4**: Update Best Practices (1 hour)
  - Best practices for new features

**Total Time / 총 시간**: 9.5 hours / 9.5시간

---

### 7.3 DEVELOPER_GUIDE.md Updates

- [ ] **Task 7.3.1**: Update Package Structure (1 hour)
  - Add 12 new files

- [ ] **Task 7.3.2**: Add Core Components (3 hours)
  - 3.9-3.20: Architecture for new features

- [ ] **Task 7.3.3**: Update Design Patterns (1 hour)
  - Document new patterns used

- [ ] **Task 7.3.4**: Update Testing Guide (1 hour)
  - Testing strategies for new features

**Total Time / 총 시간**: 6 hours / 6시간

---

## 8. Testing Strategy

### 8.1 Unit Tests

**Coverage Goal / 커버리지 목표**: 80%+ per feature

Each feature requires:
- Positive test cases (happy path)
- Negative test cases (error handling)
- Edge cases
- Thread safety tests (if applicable)

각 기능에는 다음이 필요합니다:
- 긍정 테스트 케이스 (정상 경로)
- 부정 테스트 케이스 (에러 처리)
- 엣지 케이스
- 스레드 안전성 테스트 (해당되는 경우)

### 8.2 Integration Tests

- [ ] **Task 8.2.1**: Phase interaction tests (4 hours)
  - Test Phase 5 features with Phase 1-4
  - Test Phase 5 features together

### 8.3 Performance Tests

- [ ] **Task 8.3.1**: Benchmarks (3 hours)
  - Benchmark all new features
  - Compare with baseline (Phase 1-4)

- [ ] **Task 8.3.2**: Load tests (2 hours)
  - Test under concurrent load
  - Memory profiling

**Total Time / 총 시간**: 9 hours / 9시간

---

## 9. Timeline

### Week 1: Phase 5a (v1.11.001)
- **Day 1-2**: Cookie Management (8h)
- **Day 3-4**: Interceptors (10h)
- **Day 5**: Batch Requests (8h)
- **Total**: 26 hours

### Week 2: Phase 5b (v1.11.002)
- **Day 1-2**: Request Caching (12h)
- **Day 3**: Proxy Support (8h)
- **Day 4-5**: Circuit Breaker (10h)
- **Total**: 30 hours

### Week 3: Phase 5c (v1.11.003)
- **Day 1-2**: OAuth2 Helper (12h)
- **Day 3**: Request Metrics (8h)
- **Day 4-5**: Streaming Support (12h)
- **Total**: 32 hours

### Week 4: Phase 5d (v1.11.004)
- **Day 1**: GraphQL Support (8h)
- **Day 2-3**: WebSocket Support (12h)
- **Day 4-5**: Request Mocking (10h)
- **Total**: 30 hours

### Week 5: Documentation & Testing (v1.11.005)
- **Day 1-2**: README updates (7h)
- **Day 3**: USER_MANUAL updates (9.5h)
- **Day 4**: DEVELOPER_GUIDE updates (6h)
- **Day 5**: Integration & performance tests (9h)
- **Total**: 31.5 hours

**Grand Total / 총합**: ~150 hours / 약 150시간

---

## 10. Deliverables / 산출물

### 10.1 Code / 코드
- [ ] 12 new .go files (~3000 lines)
- [ ] Updated Client struct
- [ ] 50+ new functions/methods
- [ ] All tests passing

### 10.2 Documentation / 문서
- [ ] DESIGN_PLAN_PHASE5.md ✅
- [ ] WORK_PLAN_PHASE5.md ✅
- [ ] Updated README.md
- [ ] Updated USER_MANUAL.md
- [ ] Updated DEVELOPER_GUIDE.md

### 10.3 Tests / 테스트
- [ ] 50+ unit tests
- [ ] 10+ integration tests
- [ ] 20+ benchmarks
- [ ] 80%+ code coverage

### 10.4 Examples / 예제
- [ ] 12+ feature examples
- [ ] Updated examples/httputil/main.go
- [ ] All examples working

### 10.5 Release / 릴리스
- [ ] v1.11.001: Phase 5a
- [ ] v1.11.002: Phase 5b
- [ ] v1.11.003: Phase 5c
- [ ] v1.11.004: Phase 5d
- [ ] v1.11.005: Final release

---

**Document Status / 문서 상태**: ✅ Complete / 완료
**Next Step / 다음 단계**: Begin Phase 5a implementation
**Start Date / 시작 날짜**: 2025-10-15
