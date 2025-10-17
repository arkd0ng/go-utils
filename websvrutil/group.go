package websvrutil

import "net/http"

// group.go provides route grouping functionality for organizing related routes.
//
// This file implements the Group type which allows organizing routes under common
// path prefixes with shared middleware:
//
// Route Grouping Features:
//   - Prefix Management: All routes in a group share a common path prefix
//   - Middleware Stacking: Group-specific middleware applied to all routes
//   - Nested Groups: Groups can be nested to create hierarchical structures
//   - Method Chaining: Fluent API for registering multiple routes
//
// Group Structure:
//   - Group.prefix: Path prefix prepended to all routes
//   - Group.middleware: Stack of middleware functions
//   - Group.app: Reference to parent App for route registration
//
// Core Methods:
//   - App.Group(prefix): Create new top-level group
//   - Group.Group(prefix): Create nested group (inherits parent prefix and middleware)
//   - Group.Use(middleware...): Add middleware to group
//   - HTTP method shortcuts: GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD
//   - registerRoute(): Internal method combining prefix, middleware, and handler
//
// Middleware Inheritance:
//   - Child groups inherit parent middleware
//   - Middleware is copied (not shared) to prevent cross-contamination
//   - Middleware is applied in registration order
//
// Use Cases:
//   - API versioning: /api/v1, /api/v2
//   - Feature-based organization: /users, /products, /orders
//   - Permission-based routes: /public, /authenticated, /admin
//   - Microservice-style routing: /auth, /billing, /notifications
//
// Example usage:
//
//	app := New()
//
//	// Public API routes
//	public := app.Group("/api/public")
//	public.GET("/health", healthCheck)
//
//	// Authenticated API routes
//	api := app.Group("/api")
//	api.Use(AuthMiddleware())
//	api.GET("/users", listUsers)
//
//	// Admin routes with nested groups
//	admin := api.Group("/admin")
//	admin.Use(AdminMiddleware())
//	admin.GET("/stats", getStats)        // Route: /api/admin/stats
//	admin.POST("/settings", updateSettings)
//
// Performance:
// - Group creation is lightweight (no route pre-registration)
// - Middleware is applied once during route registration (not per request)
// - Prefix concatenation happens at registration time (efficient request handling)
//
// group.go는 관련 라우트를 구성하기 위한 라우트 그룹화 기능을 제공합니다.
//
// 이 파일은 공통 경로 접두사와 공유 미들웨어 아래에 라우트를 구성할 수 있는
// Group 타입을 구현합니다:
//
// 라우트 그룹화 기능:
//   - 접두사 관리: 그룹의 모든 라우트가 공통 경로 접두사 공유
//   - 미들웨어 스택: 모든 라우트에 적용되는 그룹별 미들웨어
//   - 중첩 그룹: 계층적 구조를 만들기 위한 그룹 중첩
//   - 메서드 체이닝: 여러 라우트 등록을 위한 유창한 API
//
// Group 구조:
//   - Group.prefix: 모든 라우트 앞에 추가되는 경로 접두사
//   - Group.middleware: 미들웨어 함수 스택
//   - Group.app: 라우트 등록을 위한 부모 App 참조
//
// 핵심 메서드:
//   - App.Group(prefix): 새 최상위 그룹 생성
//   - Group.Group(prefix): 중첩 그룹 생성 (부모 접두사 및 미들웨어 상속)
//   - Group.Use(middleware...): 그룹에 미들웨어 추가
//   - HTTP 메서드 단축키: GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD
//   - registerRoute(): 접두사, 미들웨어, 핸들러를 결합하는 내부 메서드
//
// 미들웨어 상속:
//   - 자식 그룹은 부모 미들웨어 상속
//   - 미들웨어는 복사됨 (공유되지 않음) - 교차 오염 방지
//   - 미들웨어는 등록 순서대로 적용
//
// 사용 사례:
//   - API 버전 관리: /api/v1, /api/v2
//   - 기능 기반 구성: /users, /products, /orders
//   - 권한 기반 라우트: /public, /authenticated, /admin
//   - 마이크로서비스 스타일 라우팅: /auth, /billing, /notifications
//
// 사용 예제:
//
//	app := New()
//
//	// 공개 API 라우트
//	public := app.Group("/api/public")
//	public.GET("/health", healthCheck)
//
//	// 인증된 API 라우트
//	api := app.Group("/api")
//	api.Use(AuthMiddleware())
//	api.GET("/users", listUsers)
//
//	// 중첩 그룹이 있는 관리자 라우트
//	admin := api.Group("/admin")
//	admin.Use(AdminMiddleware())
//	admin.GET("/stats", getStats)        // 라우트: /api/admin/stats
//	admin.POST("/settings", updateSettings)
//
// 성능:
// - 그룹 생성은 가벼움 (라우트 사전 등록 없음)
// - 미들웨어는 라우트 등록 중 한 번 적용됨 (요청당 아님)
// - 접두사 연결은 등록 시점에 발생 (효율적인 요청 처리)

// Group represents a route group with a common prefix and middleware.
// Group은 공통 접두사와 미들웨어를 가진 라우트 그룹을 나타냅니다.
//
// Route groups allow organizing related routes under a common path prefix
// and applying middleware to all routes in the group.
// 라우트 그룹을 사용하면 관련 라우트를 공통 경로 접두사 아래로 구성하고
// 그룹의 모든 라우트에 미들웨어를 적용할 수 있습니다.
//
// Features
// 기능:
// - Prefix: All routes in the group share a common path prefix
// - 접두사: 그룹의 모든 라우트가 공통 경로 접두사 공유
// - Middleware: Group-specific middleware applied to all routes
// - 미들웨어: 모든 라우트에 적용되는 그룹별 미들웨어
// - Nesting: Groups can be nested to create hierarchical route structures
// - 중첩: 그룹을 중첩하여 계층적 라우트 구조 생성 가능
//
// Example
// 예제:
//
// // Create API v1 group
// API v1 그룹 생성
//
//	v1 := app.Group("/api/v1")
//	v1.Use(AuthMiddleware())
//	v1.GET("/users", listUsers)
//	v1.POST("/users", createUser)
//
// // Create nested admin group
// 중첩된 admin 그룹 생성
//
//	admin := v1.Group("/admin")
//	admin.Use(AdminMiddleware())
//	admin.GET("/stats", getStats)
//	// Results in route: /api/v1/admin/stats
type Group struct {
	// prefix is the path prefix for all routes in this group
	// prefix는 이 그룹의 모든 라우트에 대한 경로 접두사입니다
	prefix string

	// middleware stores group-specific middleware
	// middleware는 그룹별 미들웨어를 저장합니다
	middleware []MiddlewareFunc

	// app is a reference to the parent App
	// app은 부모 App에 대한 참조입니다
	app *App
}

// Group creates a new route group with the given prefix.
// Group은 주어진 접두사로 새 라우트 그룹을 생성합니다.
//
// The prefix is prepended to all routes registered in the group.
// 접두사는 그룹에 등록된 모든 라우트 앞에 추가됩니다.
//
// Parameters
// 매개변수:
//   - prefix: Path prefix for the group (e.g., "/api/v1")
//
// Returns
// 반환:
//   - *Group: New route group for method chaining
//
// Example
// 예제:
//
//	api := app.Group("/api")
//	api.GET("/users", listUsers)     // Route: /api/users
//	api.POST("/users", createUser)   // Route: /api/users
func (a *App) Group(prefix string) *Group {
	return &Group{
		prefix:     prefix,
		middleware: make([]MiddlewareFunc, 0),
		app:        a,
	}
}

// Group creates a nested route group with the given prefix.
// Group은 주어진 접두사로 중첩된 라우트 그룹을 생성합니다.
//
// The new group inherits the parent group's prefix and middleware.
// 새 그룹은 부모 그룹의 접두사와 미들웨어를 상속합니다.
//
// Parameters
// 매개변수:
//   - prefix: Additional path prefix for the nested group
//
// Returns
// 반환:
//   - *Group: New nested route group
//
// Example
// 예제:
//
//	api := app.Group("/api")
//	v1 := api.Group("/v1")           // Prefix: /api/v1
//	admin := v1.Group("/admin")      // Prefix: /api/v1/admin
//	admin.GET("/users", listUsers)   // Route: /api/v1/admin/users
func (g *Group) Group(prefix string) *Group {
	newGroup := &Group{
		prefix:     g.prefix + prefix,
		middleware: make([]MiddlewareFunc, len(g.middleware)),
		app:        g.app,
	}
	// Copy parent middleware to new group
	// 부모 미들웨어를 새 그룹으로 복사
	copy(newGroup.middleware, g.middleware)
	return newGroup
}

// Use adds middleware to the group.
// Use는 그룹에 미들웨어를 추가합니다.
//
// Middleware is applied to all routes in the group in the order they are added.
// 미들웨어는 추가된 순서대로 그룹의 모든 라우트에 적용됩니다.
//
// Parameters
// 매개변수:
//   - middleware: One or more middleware functions
//
// Returns
// 반환:
//   - *Group: The group for method chaining
//
// Example
// 예제:
//
//	api := app.Group("/api")
//	api.Use(AuthMiddleware(), LoggingMiddleware())
//	api.GET("/users", listUsers) // Both middleware applied
func (g *Group) Use(middleware ...MiddlewareFunc) *Group {
	g.middleware = append(g.middleware, middleware...)
	return g
}

// GET registers a GET route in the group.
// GET은 그룹에 GET 라우트를 등록합니다.
func (g *Group) GET(pattern string, handler http.HandlerFunc) *Group {
	g.registerRoute("GET", pattern, handler)
	return g
}

// POST registers a POST route in the group.
// POST는 그룹에 POST 라우트를 등록합니다.
func (g *Group) POST(pattern string, handler http.HandlerFunc) *Group {
	g.registerRoute("POST", pattern, handler)
	return g
}

// PUT registers a PUT route in the group.
// PUT은 그룹에 PUT 라우트를 등록합니다.
func (g *Group) PUT(pattern string, handler http.HandlerFunc) *Group {
	g.registerRoute("PUT", pattern, handler)
	return g
}

// PATCH registers a PATCH route in the group.
// PATCH는 그룹에 PATCH 라우트를 등록합니다.
func (g *Group) PATCH(pattern string, handler http.HandlerFunc) *Group {
	g.registerRoute("PATCH", pattern, handler)
	return g
}

// DELETE registers a DELETE route in the group.
// DELETE는 그룹에 DELETE 라우트를 등록합니다.
func (g *Group) DELETE(pattern string, handler http.HandlerFunc) *Group {
	g.registerRoute("DELETE", pattern, handler)
	return g
}

// OPTIONS registers an OPTIONS route in the group.
// OPTIONS는 그룹에 OPTIONS 라우트를 등록합니다.
func (g *Group) OPTIONS(pattern string, handler http.HandlerFunc) *Group {
	g.registerRoute("OPTIONS", pattern, handler)
	return g
}

// HEAD registers a HEAD route in the group.
// HEAD는 그룹에 HEAD 라우트를 등록합니다.
func (g *Group) HEAD(pattern string, handler http.HandlerFunc) *Group {
	g.registerRoute("HEAD", pattern, handler)
	return g
}

// registerRoute registers a route with the group's prefix and middleware.
// registerRoute는 그룹의 접두사와 미들웨어로 라우트를 등록합니다.
//
// Process
// 프로세스:
//
//  1. Combine group prefix with route pattern
//
//  2. Wrap handler with group middleware
//
//  3. Register route with parent app
//
//  1. 그룹 접두사와 라우트 패턴 결합
//
//  2. 그룹 미들웨어로 핸들러 래핑
//
//  3. 부모 앱에 라우트 등록
func (g *Group) registerRoute(method, pattern string, handler http.HandlerFunc) {
	// Combine group prefix with pattern
	// 그룹 접두사와 패턴 결합
	fullPattern := g.prefix + pattern

	// Wrap handler with group middleware
	// 그룹 미들웨어로 핸들러 래핑
	wrappedHandler := handler
	for i := len(g.middleware) - 1; i >= 0; i-- {
		wrappedHandler = g.middleware[i](http.HandlerFunc(wrappedHandler)).ServeHTTP
	}

	// Register route with parent app
	// 부모 앱에 라우트 등록
	g.app.registerRoute(method, fullPattern, wrappedHandler)
}
