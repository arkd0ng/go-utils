package websvrutil_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/arkd0ng/go-utils/websvrutil"
)

// Example demonstrates a complete application setup with middleware and routes.
// Example은 미들웨어와 라우트가 포함된 완전한 애플리케이션 설정을 보여줍니다.
func Example() {
	app := websvrutil.New(websvrutil.WithTemplateDir(""))

	// Add middleware
	// 미들웨어 추가
	app.Use(websvrutil.Logger())
	app.Use(websvrutil.Recovery())

	// Register routes
	// 라우트 등록
	app.GET("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := websvrutil.GetContext(r)
		ctx.JSON(200, map[string]string{"message": "Hello World"})
	})

	app.GET("/users/:id", func(w http.ResponseWriter, r *http.Request) {
		ctx := websvrutil.GetContext(r)
		id := ctx.Param("id")
		ctx.JSON(200, map[string]string{"user_id": id})
	})

	fmt.Println("Application configured successfully")
	// Output: Application configured successfully
}

// ExampleNew demonstrates creating a new web application.
// ExampleNew는 새 웹 애플리케이션을 생성하는 방법을 보여줍니다.
func ExampleNew() {
	app := websvrutil.New(websvrutil.WithTemplateDir(""))
	fmt.Printf("App created: %T\n", app)
	// Output: App created: *websvrutil.App
}

// ExampleNew_withOptions demonstrates creating an app with custom options.
// ExampleNew_withOptions는 커스텀 옵션으로 앱을 생성하는 방법을 보여줍니다.
func ExampleNew_withOptions() {
	app := websvrutil.New(
		websvrutil.WithTemplateDir(""),
		websvrutil.WithStaticDir("custom/static"),
		websvrutil.WithStaticPrefix("/assets"),
	)
	fmt.Printf("App created with options: %T\n", app)
	// Output: App created with options: *websvrutil.App
}

// ExampleApp_GET demonstrates registering a GET route.
// ExampleApp_GET는 GET 라우트를 등록하는 방법을 보여줍니다.
func ExampleApp_GET() {
	app := websvrutil.New(websvrutil.WithTemplateDir(""))

	app.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	fmt.Println("GET route registered")
	// Output: GET route registered
}

// ExampleApp_POST demonstrates registering a POST route.
// ExampleApp_POST는 POST 라우트를 등록하는 방법을 보여줍니다.
func ExampleApp_POST() {
	app := websvrutil.New(websvrutil.WithTemplateDir(""))

	app.POST("/users", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(201)
		w.Write([]byte("User created"))
	})

	fmt.Println("POST route registered")
	// Output: POST route registered
}

// ExampleApp_Group demonstrates creating route groups.
// ExampleApp_Group는 라우트 그룹을 생성하는 방법을 보여줍니다.
func ExampleApp_Group() {
	app := websvrutil.New(websvrutil.WithTemplateDir(""))

	// Create API group
	// API 그룹 생성
	api := app.Group("/api")
	api.GET("/users", func(w http.ResponseWriter, r *http.Request) {})
	api.GET("/posts", func(w http.ResponseWriter, r *http.Request) {})

	// Create v1 subgroup
	// v1 하위 그룹 생성
	v1 := api.Group("/v1")
	v1.GET("/products", func(w http.ResponseWriter, r *http.Request) {})

	fmt.Println("Route groups created")
	// Output: Route groups created
}

// ExampleContext_Param demonstrates retrieving URL parameters.
// ExampleContext_Param은 URL 매개변수를 가져오는 방법을 보여줍니다.
func ExampleContext_Param() {
	app := websvrutil.New(websvrutil.WithTemplateDir(""))

	app.GET("/users/:id", func(w http.ResponseWriter, r *http.Request) {
		ctx := websvrutil.GetContext(r)
		id := ctx.Param("id")
		fmt.Printf("User ID: %s\n", id)
	})

	fmt.Println("Route with parameter configured")
	// Output: Route with parameter configured
}

// ExampleContext_JSON demonstrates sending JSON responses.
// ExampleContext_JSON은 JSON 응답을 전송하는 방법을 보여줍니다.
func ExampleContext_JSON() {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	ctx := websvrutil.NewContext(rec, req)

	data := map[string]interface{}{
		"message": "success",
		"code":    200,
	}

	ctx.JSON(200, data)

	fmt.Println("Content-Type:", rec.Header().Get("Content-Type"))
	// Output: Content-Type: application/json; charset=utf-8
}

// ExampleContext_BindJSON demonstrates binding JSON request bodies.
// ExampleContext_BindJSON은 JSON 요청 본문을 바인딩하는 방법을 보여줍니다.
func ExampleContext_BindJSON() {
	type User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	app := websvrutil.New(websvrutil.WithTemplateDir(""))

	app.POST("/users", func(w http.ResponseWriter, r *http.Request) {
		ctx := websvrutil.GetContext(r)
		var user User
		if err := ctx.BindJSON(&user); err != nil {
			ctx.ErrorJSON(400, "Invalid JSON")
			return
		}
		ctx.JSON(201, user)
	})

	fmt.Println("JSON binding route configured")
	// Output: JSON binding route configured
}

// ExampleContext_Query demonstrates retrieving query parameters.
// ExampleContext_Query는 쿼리 매개변수를 가져오는 방법을 보여줍니다.
func ExampleContext_Query() {
	req := httptest.NewRequest("GET", "/?page=1&limit=10", nil)
	rec := httptest.NewRecorder()
	ctx := websvrutil.NewContext(rec, req)

	page := ctx.Query("page")
	limit := ctx.Query("limit")

	fmt.Printf("Page: %s, Limit: %s\n", page, limit)
	// Output: Page: 1, Limit: 10
}

// ExampleContext_SetCookie demonstrates setting cookies.
// ExampleContext_SetCookie는 쿠키를 설정하는 방법을 보여줍니다.
func ExampleContext_SetCookie() {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	ctx := websvrutil.NewContext(rec, req)

	cookie := &http.Cookie{
		Name:  "session",
		Value: "abc123",
		Path:  "/",
	}

	ctx.SetCookie(cookie)

	fmt.Println("Cookie set successfully")
	// Output: Cookie set successfully
}

// ExampleLogger demonstrates using the Logger middleware.
// ExampleLogger는 Logger 미들웨어 사용 방법을 보여줍니다.
func ExampleLogger() {
	app := websvrutil.New(websvrutil.WithTemplateDir(""))

	// Add logger middleware
	// 로거 미들웨어 추가
	app.Use(websvrutil.Logger())

	app.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	fmt.Println("Logger middleware configured")
	// Output: Logger middleware configured
}

// ExampleRecovery demonstrates using the Recovery middleware.
// ExampleRecovery는 Recovery 미들웨어 사용 방법을 보여줍니다.
func ExampleRecovery() {
	app := websvrutil.New(websvrutil.WithTemplateDir(""))

	// Add recovery middleware
	// 복구 미들웨어 추가
	app.Use(websvrutil.Recovery())

	app.GET("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("something went wrong")
	})

	fmt.Println("Recovery middleware configured")
	// Output: Recovery middleware configured
}

// ExampleCORS demonstrates using the CORS middleware.
// ExampleCORS는 CORS 미들웨어 사용 방법을 보여줍니다.
func ExampleCORS() {
	app := websvrutil.New(websvrutil.WithTemplateDir(""))

	// Add CORS middleware with custom config
	// 커스텀 설정으로 CORS 미들웨어 추가
	app.Use(websvrutil.CORSWithConfig(websvrutil.CORSConfig{
		AllowOrigins: []string{"https://example.com"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	fmt.Println("CORS middleware configured")
	// Output: CORS middleware configured
}

// ExampleNewSessionStore demonstrates creating a session store.
// ExampleNewSessionStore는 세션 저장소 생성 방법을 보여줍니다.
func ExampleNewSessionStore() {
	store := websvrutil.NewSessionStore(websvrutil.DefaultSessionOptions())

	fmt.Printf("Session store created: %T\n", store)
	// Output: Session store created: *websvrutil.SessionStore
}
