package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"os/signal"
	"path/filepath"
	"sort"
	"strings"
	"syscall"
	"time"

	"github.com/arkd0ng/go-utils/fileutil"
	"github.com/arkd0ng/go-utils/logging"
	"github.com/arkd0ng/go-utils/websvrutil"
)

const logBaseName = "websvrutil-example"

var (
	exampleLogger *logging.Logger
	sectionIndex  int
)

func main() {
	exampleLogger = setupLogger()
	if exampleLogger != nil {
		defer exampleLogger.Close()
	}

	logSection("Initialization", "초기 설정")
	logDual("websvrutil package version: "+websvrutil.Version, "websvrutil 패키지 버전: "+websvrutil.Version)
	logDual("Shared log file: logs/"+logBaseName+".log", "공용 로그 파일: logs/"+logBaseName+".log")

	runBasicServerExamples()
	runRoutingExamples()
	runContextExamples()
	runBindingExamples()
	runResponseExamples()
	runMiddlewareExamples()
	runSessionExamples()
	runTemplateExamples()
	runCSRFExamples()
	runValidatorExamples()
	runFileUploadExamples()
	runStaticFileExamples()
	runTestingExamples()
	runGracefulShutdownExample()
	runProductionConfigExample()

	logSection("All Examples Completed", "모든 예제 완료")
	logDual("Every example finished successfully.", "모든 예제가 성공적으로 완료되었습니다.")
}

func setupLogger() *logging.Logger {
	if err := os.MkdirAll("logs", 0o755); err != nil {
		fmt.Fprintf(os.Stderr, "⚠️  Failed to create logs directory: %v\n", err)
		return nil
	}

	logFilePath := fmt.Sprintf("logs/%s.log", logBaseName)

	if fileutil.Exists(logFilePath) {
		if modTime, err := fileutil.ModTime(logFilePath); err == nil {
			backupName := fmt.Sprintf("logs/%s-%s.log", logBaseName, modTime.Format("20060102-150405"))
			if err := fileutil.CopyFile(logFilePath, backupName); err == nil {
				fmt.Fprintf(os.Stdout, "✅ Backed up previous log to: %s\n", backupName)
				fileutil.DeleteFile(logFilePath)
			}
		}

		backupPattern := fmt.Sprintf("logs/%s-*.log", logBaseName)
		if backupFiles, err := filepath.Glob(backupPattern); err == nil && len(backupFiles) > 5 {
			type fileInfo struct {
				path    string
				modTime time.Time
			}
			files := make([]fileInfo, 0, len(backupFiles))
			for _, f := range backupFiles {
				if mt, err := fileutil.ModTime(f); err == nil {
					files = append(files, fileInfo{path: f, modTime: mt})
				}
			}

			sort.Slice(files, func(i, j int) bool {
				return files[i].modTime.Before(files[j].modTime)
			})

			for i := 0; i < len(files)-5; i++ {
				fileutil.DeleteFile(files[i].path)
				fmt.Fprintf(os.Stdout, "🗑️  Deleted old backup: %s\n", files[i].path)
			}
		}
	}

	logger, err := logging.New(
		logging.WithFilePath(logFilePath),
		logging.WithStdout(true),
		logging.WithLevel(logging.INFO),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "⚠️  Failed to initialize logger: %v\n", err)
		return nil
	}

	logger.Banner("websvrutil Package Examples", websvrutil.Version)
	logger.Info("Logs mirror console output", "message", "로그가 콘솔 출력을 그대로 반영합니다")
	logger.Info("")

	return logger
}

func logSection(titleEn, titleKo string) {
	sectionIndex++
	divider := strings.Repeat("━", 78)
	logPrintln(divider)
	logPrintln(fmt.Sprintf("Section %02d: %s", sectionIndex, titleEn))
	logPrintln(fmt.Sprintf("섹션 %02d: %s", sectionIndex, titleKo))
	logPrintln(divider)
}

func logDual(messageEn, messageKo string) {
	logPrintln(messageEn)
	logPrintln(messageKo)
}

func logPrintln(args ...interface{}) {
	fmt.Fprintln(os.Stdout, args...)
	if exampleLogger != nil {
		msg := strings.TrimSuffix(fmt.Sprintln(args...), "\n")
		exampleLogger.Info(msg)
	}
}

func logPrintf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, format, args...)
	if exampleLogger != nil {
		exampleLogger.Info(strings.TrimSuffix(fmt.Sprintf(format, args...), "\n"))
	}
}

func runBasicServerExamples() {
	logSection("Basic Server", "기본 서버")

	app := websvrutil.New()
	logDual("Created app with default middleware (Logger, Recovery)", "기본 미들웨어(Logger, Recovery)가 활성화된 앱 생성")
	logDual("Recommended usage: app.Run(\":8080\")", "권장 실행 방법: app.Run(\":8080\")")

	customApp := websvrutil.New(
		websvrutil.WithReadTimeout(30*time.Second),
		websvrutil.WithWriteTimeout(30*time.Second),
		websvrutil.WithIdleTimeout(90*time.Second),
		websvrutil.WithMaxHeaderBytes(2<<20),
		websvrutil.WithTemplateDir("views"),
		websvrutil.WithStaticDir("public"),
		websvrutil.WithStaticPrefix("/assets"),
		websvrutil.WithAutoReload(true),
		websvrutil.WithLogger(false),
		websvrutil.WithRecovery(true),
	)
	logDual("Custom app configured with non-default options", "커스텀 옵션으로 구성된 앱 설정 완료")
	logDual("Key options: ReadTimeout=30s, WriteTimeout=30s, IdleTimeout=90s, MaxHeaderBytes=2MB", "주요 옵션: ReadTimeout=30초, WriteTimeout=30초, IdleTimeout=90초, MaxHeaderBytes=2MB")
	logDual("TemplateDir=views, StaticDir=public, StaticPrefix=/assets, AutoReload=true", "TemplateDir=views, StaticDir=public, StaticPrefix=/assets, AutoReload=true")
	logDual("Logger disabled, Recovery enabled", "Logger 비활성화, Recovery 활성화")

	_ = app
	_ = customApp
}

func runRoutingExamples() {
	logSection("Routing", "라우팅")

	type routeResult struct {
		Method string `json:"method"`
		Path   string `json:"path"`
		Status int    `json:"status"`
	}

	results := []routeResult{}

	routingApp := websvrutil.New()
	routingApp.GET("/resources", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	routingApp.POST("/resources", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	})
	routingApp.PUT("/resources/42", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	routingApp.PATCH("/resources/42", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
	})
	routingApp.DELETE("/resources/42", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	routingApp.HEAD("/resources", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	routingApp.OPTIONS("/resources", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Allow", "GET,POST,PUT,PATCH,DELETE,HEAD,OPTIONS")
		w.WriteHeader(http.StatusNoContent)
	})

	testCases := []struct {
		method string
		path   string
		expect int
	}{
		{"GET", "/resources", http.StatusOK},
		{"POST", "/resources", http.StatusCreated},
		{"PUT", "/resources/42", http.StatusOK},
		{"PATCH", "/resources/42", http.StatusAccepted},
		{"DELETE", "/resources/42", http.StatusNoContent},
		{"HEAD", "/resources", http.StatusOK},
		{"OPTIONS", "/resources", http.StatusNoContent},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest(tc.method, tc.path, nil)
		rec := httptest.NewRecorder()
		routingApp.ServeHTTP(rec, req)
		results = append(results, routeResult{Method: tc.method, Path: tc.path, Status: rec.Code})
	}

	// Route group with middleware
	apiGroup := routingApp.Group("/api")
	apiGroup.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-API-Version", "v1")
			next.ServeHTTP(w, r)
		})
	})
	apiGroup.GET("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	reqGroup := httptest.NewRequest("GET", "/api/health", nil)
	recGroup := httptest.NewRecorder()
	routingApp.ServeHTTP(recGroup, reqGroup)
	results = append(results, routeResult{Method: "GET", Path: "/api/health", Status: recGroup.Code})
	logDual("Route group middleware header X-API-Version: "+recGroup.Header().Get("X-API-Version"), "라우트 그룹 미들웨어 헤더 X-API-Version: "+recGroup.Header().Get("X-API-Version"))

	logDual("Routing matrix (method, path, status):", "라우팅 결과 (메서드, 경로, 상태):")
	for _, res := range results {
		logPrintf("  %s %s -> %d\n", res.Method, res.Path, res.Status)
	}
}

func runContextExamples() {
	logSection("Context Helpers", "Context 헬퍼")

	type ctxResult struct {
		PathParams map[string]string      `json:"path_params"`
		Query      map[string]string      `json:"query"`
		Headers    map[string]string      `json:"headers"`
		Values     map[string]interface{} `json:"values"`
	}

	result := ctxResult{
		PathParams: map[string]string{},
		Query:      map[string]string{},
		Headers:    map[string]string{},
		Values:     map[string]interface{}{},
	}

	app := websvrutil.New()

	app.GET("/users/:userId/posts/:postId", func(w http.ResponseWriter, r *http.Request) {
		ctx := websvrutil.GetContext(r)
		result.PathParams["userId"] = ctx.Param("userId")
		result.PathParams["postId"] = ctx.Param("postId")
		ctx.Set("authenticated", true)
		ctx.Set("roles", []string{"admin", "editor"})
		ctx.Set("requestID", "req-12345")
	})

	app.GET("/search", func(w http.ResponseWriter, r *http.Request) {
		ctx := websvrutil.GetContext(r)
		result.Query["q"] = ctx.Query("q")
		result.Query["page"] = ctx.QueryDefault("page", "1")
		result.Query["limit"] = ctx.QueryDefault("limit", "10")
		result.Headers["User-Agent"] = ctx.UserAgent()
		result.Headers["Content-Type"] = ctx.ContentType()
		result.Headers["Client-IP"] = ctx.ClientIP()
		result.Headers["Referer"] = ctx.Referer()
		if val, ok := ctx.Get("authenticated"); ok {
			result.Values["authenticated"] = val
		}
		if val, ok := ctx.Get("roles"); ok {
			result.Values["roles"] = val
		}
		if val, ok := ctx.Get("requestID"); ok {
			result.Values["requestID"] = val
		}
	})

	// Simulate path parameter request
	reqPath := httptest.NewRequest("GET", "/users/42/posts/99", nil)
	recPath := httptest.NewRecorder()
	app.ServeHTTP(recPath, reqPath)

	// Simulate query parameter request with headers
	reqQuery := httptest.NewRequest("GET", "/search?q=golang&page=2", nil)
	reqQuery.Header.Set("User-Agent", "ExampleClient/1.0")
	reqQuery.Header.Set("Content-Type", "application/json")
	reqQuery.Header.Set("X-Forwarded-For", "203.0.113.1, 70.41.3.18")
	reqQuery.Header.Set("Referer", "https://example.com/docs")
	recQuery := httptest.NewRecorder()
	app.ServeHTTP(recQuery, reqQuery)

	logDual("Context helper results (JSON):", "Context 헬퍼 결과(JSON):")
	encoded, _ := json.MarshalIndent(result, "  ", "  ")
	logPrintln("  " + string(encoded))
}

func runBindingExamples() {
	logSection("Request Binding", "요청 바인딩")

	type profilePayload struct {
		Name  string `json:"name" form:"name"`
		Email string `json:"email" form:"email"`
		Age   int    `json:"age" form:"age"`
	}

	type queryPayload struct {
		Query string `form:"q"`
		Page  int    `form:"page"`
		Limit int    `form:"limit"`
	}

	bindingApp := websvrutil.New()

	bindingApp.POST("/bind/json", func(w http.ResponseWriter, r *http.Request) {
		ctx := websvrutil.GetContext(r)
		var payload profilePayload
		if err := ctx.BindJSON(&payload); err != nil {
			logDual("BindJSON failed: "+err.Error(), "BindJSON 실패: "+err.Error())
			ctx.Status(http.StatusBadRequest)
			return
		}
		logDual(fmt.Sprintf("BindJSON result: %+v", payload), fmt.Sprintf("BindJSON 결과: %+v", payload))
		ctx.JSON(http.StatusCreated, payload)
	})

	bindingApp.POST("/bind/form", func(w http.ResponseWriter, r *http.Request) {
		ctx := websvrutil.GetContext(r)
		var payload profilePayload
		if err := ctx.BindForm(&payload); err != nil {
			logDual("BindForm failed: "+err.Error(), "BindForm 실패: "+err.Error())
			ctx.Status(http.StatusBadRequest)
			return
		}
		logDual(fmt.Sprintf("BindForm result: %+v", payload), fmt.Sprintf("BindForm 결과: %+v", payload))
		ctx.JSON(http.StatusOK, payload)
	})

	bindingApp.GET("/bind/query", func(w http.ResponseWriter, r *http.Request) {
		ctx := websvrutil.GetContext(r)
		var payload queryPayload
		if err := ctx.Bind(&payload); err != nil {
			logDual("Bind (query) failed: "+err.Error(), "Bind (query) 실패: "+err.Error())
			ctx.Status(http.StatusBadRequest)
			return
		}
		logDual(fmt.Sprintf("Bind(query) result: %+v", payload), fmt.Sprintf("Bind(query) 결과: %+v", payload))
		ctx.JSON(http.StatusOK, payload)
	})

	jsonBody := bytes.NewBufferString(`{"name":"Alice","email":"alice@example.com","age":29}`)
	reqJSON := httptest.NewRequest("POST", "/bind/json", jsonBody)
	reqJSON.Header.Set("Content-Type", "application/json")
	recJSON := httptest.NewRecorder()
	bindingApp.ServeHTTP(recJSON, reqJSON)
	logDual("/bind/json response body: "+strings.TrimSpace(recJSON.Body.String()), "/bind/json 응답 본문: "+strings.TrimSpace(recJSON.Body.String()))

	reqForm := httptest.NewRequest("POST", "/bind/form", strings.NewReader("name=Bob&email=bob%40example.com&age=34"))
	reqForm.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	recForm := httptest.NewRecorder()
	bindingApp.ServeHTTP(recForm, reqForm)
	logDual("/bind/form response body: "+strings.TrimSpace(recForm.Body.String()), "/bind/form 응답 본문: "+strings.TrimSpace(recForm.Body.String()))

	reqQuery := httptest.NewRequest("GET", "/bind/query?q=golang&page=3&limit=25", nil)
	recQuery := httptest.NewRecorder()
	bindingApp.ServeHTTP(recQuery, reqQuery)
	logDual("/bind/query response body: "+strings.TrimSpace(recQuery.Body.String()), "/bind/query 응답 본문: "+strings.TrimSpace(recQuery.Body.String()))
}

func runResponseExamples() {
	logSection("Response Helpers", "응답 헬퍼")

	tempDir, err := os.MkdirTemp("", "websvrutil-responses")
	if err != nil {
		logDual("Failed to create temp dir: "+err.Error(), "임시 디렉터리 생성 실패: "+err.Error())
		return
	}
	defer os.RemoveAll(tempDir)

	tempFile := filepath.Join(tempDir, "report.txt")
	os.WriteFile(tempFile, []byte("Quarterly Report"), 0o644)

	respApp := websvrutil.New()

	respApp.GET("/json", func(w http.ResponseWriter, r *http.Request) {
		ctx := websvrutil.GetContext(r)
		ctx.JSON(http.StatusOK, map[string]interface{}{"message": "success", "count": 3})
	})

	respApp.GET("/html", func(w http.ResponseWriter, r *http.Request) {
		ctx := websvrutil.GetContext(r)
		ctx.HTML(http.StatusOK, "<h1>Hello, websvrutil</h1>")
	})

	respApp.GET("/text", func(w http.ResponseWriter, r *http.Request) {
		ctx := websvrutil.GetContext(r)
		ctx.Text(http.StatusAccepted, "plain text response")
	})

	respApp.GET("/redirect", func(w http.ResponseWriter, r *http.Request) {
		ctx := websvrutil.GetContext(r)
		ctx.Redirect(http.StatusFound, "/json")
	})

	respApp.GET("/file", func(w http.ResponseWriter, r *http.Request) {
		ctx := websvrutil.GetContext(r)
		ctx.File(tempFile)
	})

	respApp.GET("/attachment", func(w http.ResponseWriter, r *http.Request) {
		ctx := websvrutil.GetContext(r)
		ctx.FileAttachment(tempFile, "report.txt")
	})

	testPaths := []string{"/json", "/html", "/text", "/redirect", "/file", "/attachment"}
	for _, path := range testPaths {
		req := httptest.NewRequest("GET", path, nil)
		rec := httptest.NewRecorder()
		respApp.ServeHTTP(rec, req)
		logDual(fmt.Sprintf("Response for %s: status=%d", path, rec.Code), fmt.Sprintf("%s 응답: 상태=%d", path, rec.Code))
		body := strings.TrimSpace(rec.Body.String())
		if body != "" {
			logDual("  Body: "+body, "  본문: "+body)
		}
		for key, vals := range rec.Header() {
			logDual(fmt.Sprintf("  Header %s: %v", key, vals), fmt.Sprintf("  헤더 %s: %v", key, vals))
		}
	}
}

func runMiddlewareExamples() {
	logSection("Middleware", "미들웨어")

	app := websvrutil.New()
	app.Use(websvrutil.RequestID())
	app.Use(websvrutil.Timeout(2 * time.Second))
	app.Use(websvrutil.CORS())
	app.Use(websvrutil.Recovery())

	app.GET("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("forced panic for Recovery middleware")
	})

	app.GET("/protected", func(w http.ResponseWriter, r *http.Request) {
		ctx := websvrutil.GetContext(r)
		ctx.Text(http.StatusOK, "protected content")
	})

	authApp := websvrutil.New()
	authApp.Use(websvrutil.BasicAuth("admin", "secret"))
	authApp.GET("/credentials", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(50 * time.Millisecond)
		ctx := websvrutil.GetContext(r)
		ctx.Text(http.StatusOK, "authorized")
	})

	// Recovery test
	reqPanic := httptest.NewRequest("GET", "/panic", nil)
	recPanic := httptest.NewRecorder()
	app.ServeHTTP(recPanic, reqPanic)
	logDual(fmt.Sprintf("Recovery middleware response: status=%d", recPanic.Code), fmt.Sprintf("Recovery 미들웨어 응답: 상태=%d", recPanic.Code))

	// RequestID + CORS test
	reqProtected := httptest.NewRequest("GET", "/protected", nil)
	recProtected := httptest.NewRecorder()
	app.ServeHTTP(recProtected, reqProtected)
	logDual("Headers set by RequestID/CORS middleware:", "RequestID/CORS 미들웨어가 설정한 헤더:")
	logDual(fmt.Sprintf("  X-Request-ID: %s", recProtected.Header().Get("X-Request-ID")), fmt.Sprintf("  X-Request-ID: %s", recProtected.Header().Get("X-Request-ID")))
	logDual(fmt.Sprintf("  Access-Control-Allow-Origin: %s", recProtected.Header().Get("Access-Control-Allow-Origin")), fmt.Sprintf("  Access-Control-Allow-Origin: %s", recProtected.Header().Get("Access-Control-Allow-Origin")))

	// BasicAuth test - missing credentials
	reqNoAuth := httptest.NewRequest("GET", "/credentials", nil)
	recNoAuth := httptest.NewRecorder()
	authApp.ServeHTTP(recNoAuth, reqNoAuth)
	logDual(fmt.Sprintf("BasicAuth without credentials -> status %d", recNoAuth.Code), fmt.Sprintf("인증 없이 BasicAuth 호출 -> 상태 %d", recNoAuth.Code))

	// BasicAuth test - correct credentials
	reqAuth := httptest.NewRequest("GET", "/credentials", nil)
	reqAuth.SetBasicAuth("admin", "secret")
	recAuth := httptest.NewRecorder()
	authApp.ServeHTTP(recAuth, reqAuth)
	logDual(fmt.Sprintf("BasicAuth with credentials -> status %d, body=%s", recAuth.Code, strings.TrimSpace(recAuth.Body.String())), fmt.Sprintf("인증 후 BasicAuth -> 상태 %d, 본문=%s", recAuth.Code, strings.TrimSpace(recAuth.Body.String())))
}

func runSessionExamples() {
	logSection("Session Store", "세션 저장소")

	store := websvrutil.NewSessionStore(websvrutil.DefaultSessionOptions())

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	session, _ := store.Get(req)
	session.Set("user", "alice")
	session.Set("role", "admin")
	store.Save(rec, session)

	logDual("Created new session with ID: "+session.ID, "새 세션 생성, ID: "+session.ID)
	logDual(fmt.Sprintf("Session cookie: %v", rec.Result().Cookies()), fmt.Sprintf("세션 쿠키: %v", rec.Result().Cookies()))

	cookie := rec.Result().Cookies()[0]
	req2 := httptest.NewRequest("GET", "/", nil)
	req2.AddCookie(cookie)

	session2, _ := store.Get(req2)
	logDual(fmt.Sprintf("Retrieved session data: user=%s, role=%s", session2.GetString("user"), session2.GetString("role")), fmt.Sprintf("세션 데이터 조회: user=%s, role=%s", session2.GetString("user"), session2.GetString("role")))

	session2.Delete("role")
	store.Save(httptest.NewRecorder(), session2)
	logDual("Updated session: removed role key", "세션 업데이트: role 키 삭제")
}

func runTemplateExamples() {
	logSection("Template Engine", "템플릿 엔진")

	tempDir, err := os.MkdirTemp("", "websvrutil-templates")
	if err != nil {
		logDual("Failed to create template dir: "+err.Error(), "템플릿 디렉터리 생성 실패: "+err.Error())
		return
	}
	defer os.RemoveAll(tempDir)

	layoutDir := filepath.Join(tempDir, "layouts")
	os.MkdirAll(layoutDir, 0o755)

	os.WriteFile(filepath.Join(tempDir, "home.html"), []byte("<h1>{{.Title}}</h1><p>{{.Message}}</p>"), 0o644)
	os.WriteFile(filepath.Join(layoutDir, "base.html"), []byte("<html><body>{{template \"content\" .}}</body></html>"), 0o644)

	engine := websvrutil.NewTemplateEngine(tempDir)
	engine.SetLayoutDir(layoutDir)
	engine.AddFunc("upper", strings.ToUpper)

	if err := engine.Load("home.html"); err != nil {
		logDual("Template load failed: "+err.Error(), "템플릿 로드 실패: "+err.Error())
		return
	}

	var buf bytes.Buffer
	err = engine.Render(&buf, "home.html", map[string]string{"Title": "Dashboard", "Message": "Rendered without layout"})
	if err == nil {
		logDual("Render result (no layout): "+buf.String(), "레이아웃 없이 렌더링 결과: "+buf.String())
	}

	var bufLayout bytes.Buffer
	err = engine.RenderWithLayout(&bufLayout, "base.html", "home.html", map[string]string{"Title": "Dashboard", "Message": "Rendered with layout"})
	if err == nil {
		logDual("Render with layout: "+bufLayout.String(), "레이아웃과 함께 렌더링: "+bufLayout.String())
	}
}

func runCSRFExamples() {
	logSection("CSRF Middleware", "CSRF 미들웨어")

	app := websvrutil.New()
	app.Use(websvrutil.CSRF())

	tokenCapture := ""

	app.GET("/form", func(w http.ResponseWriter, r *http.Request) {
		ctx := websvrutil.GetContext(r)
		token := websvrutil.GetCSRFToken(ctx)
		tokenCapture = token
		logDual("Issued CSRF token: "+token, "발급된 CSRF 토큰: "+token)
		ctx.Text(http.StatusOK, "token issued")
	})

	app.POST("/submit", func(w http.ResponseWriter, r *http.Request) {
		ctx := websvrutil.GetContext(r)
		ctx.Text(http.StatusOK, "CSRF validation passed")
	})

	reqGet := httptest.NewRequest("GET", "/form", nil)
	recGet := httptest.NewRecorder()
	app.ServeHTTP(recGet, reqGet)

	csrfCookie := recGet.Result().Cookies()
	logDual(fmt.Sprintf("CSRF cookies: %v", csrfCookie), fmt.Sprintf("CSRF 쿠키: %v", csrfCookie))

	reqPost := httptest.NewRequest("POST", "/submit", strings.NewReader("name=test"))
	reqPost.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	reqPost.Header.Set("X-CSRF-Token", tokenCapture)
	for _, c := range csrfCookie {
		reqPost.AddCookie(c)
	}
	recPost := httptest.NewRecorder()
	app.ServeHTTP(recPost, reqPost)
	logDual(fmt.Sprintf("POST /submit status: %d", recPost.Code), fmt.Sprintf("POST /submit 상태: %d", recPost.Code))
}

func runValidatorExamples() {
	logSection("Validator", "검증기")

	type userForm struct {
		Name  string `validate:"required,min=3"`
		Email string `validate:"required,email"`
		Age   int    `validate:"gte=18,lte=120"`
		Role  string `validate:"oneof=admin user guest"`
	}

	validator := websvrutil.DefaultValidator{}

	valid := userForm{Name: "Alice", Email: "alice@example.com", Age: 30, Role: "admin"}
	invalid := userForm{Name: "A", Email: "alice", Age: 10, Role: "unknown"}

	err := validator.Validate(valid)
	logDual("Valid form validation error: "+fmt.Sprint(err), "유효한 폼 검증 에러: "+fmt.Sprint(err))

	err = validator.Validate(invalid)
	logDual("Invalid form validation error: "+err.Error(), "유효하지 않은 폼 검증 에러: "+err.Error())
}

func runFileUploadExamples() {
	logSection("File Upload", "파일 업로드")

	tempDir, err := os.MkdirTemp("", "websvrutil-uploads")
	if err != nil {
		logDual("Failed to create upload dir: "+err.Error(), "업로드 디렉터리 생성 실패: "+err.Error())
		return
	}
	defer os.RemoveAll(tempDir)

	app := websvrutil.New()
	app.POST("/upload", func(w http.ResponseWriter, r *http.Request) {
		ctx := websvrutil.GetContext(r)
		fileHeader, err := ctx.FormFile("file")
		if err != nil {
			ctx.Text(http.StatusBadRequest, "missing file")
			return
		}
		destination := filepath.Join(tempDir, fileHeader.Filename)
		if err := ctx.SaveUploadedFile(fileHeader, destination); err != nil {
			ctx.Text(http.StatusInternalServerError, "upload failed")
			return
		}
		ctx.Text(http.StatusOK, "upload success")
	})

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fileWriter, _ := writer.CreateFormFile("file", "photo.png")
	fileWriter.Write([]byte("PNGDATA"))
	writer.Close()

	req := httptest.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	logDual(fmt.Sprintf("File upload status: %d", rec.Code), fmt.Sprintf("파일 업로드 상태: %d", rec.Code))
	logDual("Saved files: "+strings.Join(listDir(tempDir), ", "), "저장된 파일: "+strings.Join(listDir(tempDir), ", "))
}

func runStaticFileExamples() {
	logSection("Static File Serving", "정적 파일 제공")

	tempDir, err := os.MkdirTemp("", "websvrutil-static")
	if err != nil {
		logDual("Failed to create static dir: "+err.Error(), "정적 디렉터리 생성 실패: "+err.Error())
		return
	}
	defer os.RemoveAll(tempDir)

	staticFile := filepath.Join(tempDir, "style.css")
	os.WriteFile(staticFile, []byte("body { color: #333; }"), 0o644)

	app := websvrutil.New()
	app.Static("/assets", tempDir)

	req := httptest.NewRequest("GET", "/assets/style.css", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	logDual(fmt.Sprintf("Static file status: %d", rec.Code), fmt.Sprintf("정적 파일 상태: %d", rec.Code))
	logDual("Static file body: "+strings.TrimSpace(rec.Body.String()), "정적 파일 본문: "+strings.TrimSpace(rec.Body.String()))
}

func runTestingExamples() {
	logSection("Testing Patterns", "테스트 패턴")

	app := websvrutil.New()
	app.GET("/ping", func(w http.ResponseWriter, r *http.Request) {
		ctx := websvrutil.GetContext(r)
		ctx.JSON(http.StatusOK, map[string]string{"message": "pong"})
	})

	req := httptest.NewRequest("GET", "/ping", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)

	logDual(fmt.Sprintf("Testing example: status=%d, body=%s", rec.Code, strings.TrimSpace(rec.Body.String())), fmt.Sprintf("테스트 예제: 상태=%d, 본문=%s", rec.Code, strings.TrimSpace(rec.Body.String())))
}

func runGracefulShutdownExample() {
	logSection("Graceful Shutdown", "정상 종료")

	app := websvrutil.New()
	done := make(chan struct{})
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.Run(":0"); err != nil && !strings.Contains(err.Error(), "Server closed") {
			logDual("Server encountered error: "+err.Error(), "서버 실행 중 에러 발생: "+err.Error())
		}
		close(done)
	}()

	logDual("Server started in goroutine on ephemeral port", "에페메럴 포트로 고루틴에서 서버 시작")
	time.Sleep(200 * time.Millisecond)

	go func() {
		time.Sleep(200 * time.Millisecond)
		quit <- syscall.SIGINT
	}()

	received := <-quit
	logDual("Simulated signal received: "+received.String(), "시뮬레이션된 시그널 수신: "+received.String())
	signal.Stop(quit)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := app.Shutdown(ctx); err != nil {
		logDual("Graceful shutdown error: "+err.Error(), "정상 종료 중 에러: "+err.Error())
	} else {
		logDual("Graceful shutdown completed without error", "정상 종료가 오류 없이 완료")
	}

	<-done
	logDual("Server goroutine exited", "서버 고루틴 종료")
}

func runProductionConfigExample() {
	logSection("Production Checklist", "프로덕션 체크리스트")

	app := websvrutil.New(
		websvrutil.WithReadTimeout(10*time.Second),
		websvrutil.WithWriteTimeout(10*time.Second),
		websvrutil.WithIdleTimeout(120*time.Second),
		websvrutil.WithMaxHeaderBytes(1<<20),
		websvrutil.WithAutoReload(false),
	)

	logDual("Production server initialized with hardened timeouts and header limits", "보안 강화 타임아웃과 헤더 제한으로 프로덕션 서버 초기화 완료")
	app.Use(websvrutil.Logger())
	app.Use(websvrutil.Recovery())
	logDual("Enabled default Logger and Recovery middleware", "기본 Logger 및 Recovery 미들웨어 활성화")

	_ = app
}

func listDir(path string) []string {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil
	}
	var names []string
	for _, entry := range entries {
		names = append(names, entry.Name())
	}
	return names
}
