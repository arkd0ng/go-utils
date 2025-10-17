package websvrutil

import "net/http"

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
//	v1 := app.Group("/api/v1")
//	v1.Use(AuthMiddleware())
//	v1.GET("/users", listUsers)
//	v1.POST("/users", createUser)
//
// // Create nested admin group
// 중첩된 admin 그룹 생성
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
//   1. Combine group prefix with route pattern
//   2. Wrap handler with group middleware
//   3. Register route with parent app
//
//   1. 그룹 접두사와 라우트 패턴 결합
//   2. 그룹 미들웨어로 핸들러 래핑
//   3. 부모 앱에 라우트 등록
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
