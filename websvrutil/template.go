package websvrutil

import (
	"fmt"
	"html/template"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// TemplateEngine manages HTML template rendering.
// TemplateEngine는 HTML 템플릿 렌더링을 관리합니다.
type TemplateEngine struct {
	// templates stores all parsed templates
	// templates는 모든 파싱된 템플릿을 저장합니다
	templates map[string]*template.Template

	// mu protects concurrent access to templates
	// mu는 템플릿에 대한 동시 액세스를 보호합니다
	mu sync.RWMutex

	// dir is the template directory
	// dir는 템플릿 디렉토리입니다
	dir string

	// funcMap contains custom template functions
	// funcMap은 커스텀 템플릿 함수를 포함합니다
	funcMap template.FuncMap

	// delims contains custom template delimiters
	// delims는 커스텀 템플릿 구분자를 포함합니다
	delims [2]string

	// layoutDir is the layout templates directory
	// layoutDir는 레이아웃 템플릿 디렉토리입니다
	layoutDir string

	// layouts stores layout templates
	// layouts는 레이아웃 템플릿을 저장합니다
	layouts map[string]*template.Template
}

// NewTemplateEngine creates a new template engine.
// NewTemplateEngine은 새로운 템플릿 엔진을 생성합니다.
//
// Example / 예제:
//
//	engine := NewTemplateEngine("views")
func NewTemplateEngine(dir string) *TemplateEngine {
	engine := &TemplateEngine{
		templates: make(map[string]*template.Template),
		layouts:   make(map[string]*template.Template),
		dir:       dir,
		layoutDir: filepath.Join(dir, "layouts"),
		funcMap:   make(template.FuncMap),
		delims:    [2]string{"{{", "}}"}, // Default delimiters
	}

	// Add built-in template functions
	// 내장 템플릿 함수 추가
	engine.addBuiltinFuncs()

	return engine
}

// addBuiltinFuncs adds built-in template functions.
// addBuiltinFuncs는 내장 템플릿 함수를 추가합니다.
func (e *TemplateEngine) addBuiltinFuncs() {
	e.funcMap["upper"] = strings.ToUpper
	e.funcMap["lower"] = strings.ToLower
	e.funcMap["title"] = strings.Title
	e.funcMap["trim"] = strings.TrimSpace
	e.funcMap["trimPrefix"] = strings.TrimPrefix
	e.funcMap["trimSuffix"] = strings.TrimSuffix
	e.funcMap["replace"] = strings.ReplaceAll
	e.funcMap["contains"] = strings.Contains
	e.funcMap["hasPrefix"] = strings.HasPrefix
	e.funcMap["hasSuffix"] = strings.HasSuffix
	e.funcMap["split"] = strings.Split
	e.funcMap["join"] = strings.Join
	e.funcMap["repeat"] = strings.Repeat

	// Date/Time functions
	// 날짜/시간 함수
	e.funcMap["now"] = time.Now
	e.funcMap["formatDate"] = func(t time.Time, layout string) string {
		return t.Format(layout)
	}
	e.funcMap["formatDateSimple"] = func(t time.Time) string {
		return t.Format("2006-01-02")
	}
	e.funcMap["formatDateTime"] = func(t time.Time) string {
		return t.Format("2006-01-02 15:04:05")
	}
	e.funcMap["formatTime"] = func(t time.Time) string {
		return t.Format("15:04:05")
	}

	// URL functions
	// URL 함수
	e.funcMap["urlEncode"] = url.QueryEscape
	e.funcMap["urlDecode"] = url.QueryUnescape

	// Safe HTML
	// 안전한 HTML
	e.funcMap["safeHTML"] = func(s string) template.HTML {
		return template.HTML(s)
	}
	e.funcMap["safeURL"] = func(s string) template.URL {
		return template.URL(s)
	}
	e.funcMap["safeJS"] = func(s string) template.JS {
		return template.JS(s)
	}

	// Utility functions
	// 유틸리티 함수
	e.funcMap["default"] = func(defaultVal, val interface{}) interface{} {
		if val == nil || val == "" {
			return defaultVal
		}
		return val
	}
	e.funcMap["len"] = func(v interface{}) int {
		switch val := v.(type) {
		case string:
			return len(val)
		case []interface{}:
			return len(val)
		case map[string]interface{}:
			return len(val)
		default:
			return 0
		}
	}
}

// SetDelimiters sets custom template delimiters.
// SetDelimiters는 커스텀 템플릿 구분자를 설정합니다.
//
// Example / 예제:
//
//	engine.SetDelimiters("[[", "]]")
func (e *TemplateEngine) SetDelimiters(left, right string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.delims = [2]string{left, right}
}

// AddFunc adds a custom template function.
// AddFunc는 커스텀 템플릿 함수를 추가합니다.
//
// Example / 예제:
//
//	engine.AddFunc("upper", strings.ToUpper)
func (e *TemplateEngine) AddFunc(name string, fn interface{}) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.funcMap[name] = fn
}

// AddFuncs adds multiple custom template functions.
// AddFuncs는 여러 커스텀 템플릿 함수를 추가합니다.
//
// Example / 예제:
//
//	engine.AddFuncs(template.FuncMap{
//	    "upper": strings.ToUpper,
//	    "lower": strings.ToLower,
//	})
func (e *TemplateEngine) AddFuncs(funcs template.FuncMap) {
	e.mu.Lock()
	defer e.mu.Unlock()
	for name, fn := range funcs {
		e.funcMap[name] = fn
	}
}

// Load loads a single template file.
// Load는 단일 템플릿 파일을 로드합니다.
//
// Example / 예제:
//
//	err := engine.Load("index.html")
func (e *TemplateEngine) Load(name string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	// Build full path
	// 전체 경로 생성
	path := filepath.Join(e.dir, name)

	// Check if file exists
	// 파일이 존재하는지 확인
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("template file not found: %s", path)
	}

	// Create new template
	// 새 템플릿 생성
	tmpl := template.New(name)

	// Set custom delimiters
	// 커스텀 구분자 설정
	tmpl.Delims(e.delims[0], e.delims[1])

	// Add custom functions
	// 커스텀 함수 추가
	if len(e.funcMap) > 0 {
		tmpl.Funcs(e.funcMap)
	}

	// Parse template file
	// 템플릿 파일 파싱
	tmpl, err := tmpl.ParseFiles(path)
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", name, err)
	}

	// Store template
	// 템플릿 저장
	e.templates[name] = tmpl

	return nil
}

// LoadGlob loads all templates matching the pattern.
// LoadGlob는 패턴과 일치하는 모든 템플릿을 로드합니다.
//
// Example / 예제:
//
//	err := engine.LoadGlob("*.html")
func (e *TemplateEngine) LoadGlob(pattern string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	// Build full pattern
	// 전체 패턴 생성
	fullPattern := filepath.Join(e.dir, pattern)

	// Find all matching files
	// 일치하는 모든 파일 찾기
	files, err := filepath.Glob(fullPattern)
	if err != nil {
		return fmt.Errorf("failed to glob pattern %s: %w", pattern, err)
	}

	if len(files) == 0 {
		return fmt.Errorf("no templates found matching pattern: %s", pattern)
	}

	// Load each file
	// 각 파일 로드
	for _, path := range files {
		// Get relative name
		// 상대 경로 이름 가져오기
		name, err := filepath.Rel(e.dir, path)
		if err != nil {
			return fmt.Errorf("failed to get relative path for %s: %w", path, err)
		}

		// Normalize path separators
		// 경로 구분자 정규화
		name = filepath.ToSlash(name)

		// Create new template
		// 새 템플릿 생성
		tmpl := template.New(name)

		// Set custom delimiters
		// 커스텀 구분자 설정
		tmpl.Delims(e.delims[0], e.delims[1])

		// Add custom functions
		// 커스텀 함수 추가
		if len(e.funcMap) > 0 {
			tmpl.Funcs(e.funcMap)
		}

		// Parse template file
		// 템플릿 파일 파싱
		tmpl, err = tmpl.ParseFiles(path)
		if err != nil {
			return fmt.Errorf("failed to parse template %s: %w", name, err)
		}

		// Store template
		// 템플릿 저장
		e.templates[name] = tmpl
	}

	return nil
}

// LoadAll loads all templates from the directory recursively.
// LoadAll은 디렉토리에서 모든 템플릿을 재귀적으로 로드합니다.
//
// Example / 예제:
//
//	err := engine.LoadAll()
func (e *TemplateEngine) LoadAll() error {
	e.mu.Lock()
	defer e.mu.Unlock()

	// Check if directory exists
	// 디렉토리가 존재하는지 확인
	if _, err := os.Stat(e.dir); os.IsNotExist(err) {
		return fmt.Errorf("template directory not found: %s", e.dir)
	}

	// Walk through directory
	// 디렉토리 순회
	return filepath.Walk(e.dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		// 디렉토리 건너뛰기
		if info.IsDir() {
			return nil
		}

		// Only load .html, .htm, .tmpl files
		// .html, .htm, .tmpl 파일만 로드
		ext := strings.ToLower(filepath.Ext(path))
		if ext != ".html" && ext != ".htm" && ext != ".tmpl" {
			return nil
		}

		// Get relative name
		// 상대 경로 이름 가져오기
		name, err := filepath.Rel(e.dir, path)
		if err != nil {
			return fmt.Errorf("failed to get relative path for %s: %w", path, err)
		}

		// Normalize path separators
		// 경로 구분자 정규화
		name = filepath.ToSlash(name)

		// Create new template
		// 새 템플릿 생성
		tmpl := template.New(name)

		// Set custom delimiters
		// 커스텀 구분자 설정
		tmpl.Delims(e.delims[0], e.delims[1])

		// Add custom functions
		// 커스텀 함수 추가
		if len(e.funcMap) > 0 {
			tmpl.Funcs(e.funcMap)
		}

		// Parse template file
		// 템플릿 파일 파싱
		tmpl, err = tmpl.ParseFiles(path)
		if err != nil {
			return fmt.Errorf("failed to parse template %s: %w", name, err)
		}

		// Store template
		// 템플릿 저장
		e.templates[name] = tmpl

		return nil
	})
}

// Render renders a template with data to the writer.
// Render는 템플릿을 데이터와 함께 writer에 렌더링합니다.
//
// Example / 예제:
//
//	err := engine.Render(w, "index.html", data)
func (e *TemplateEngine) Render(w io.Writer, name string, data interface{}) error {
	e.mu.RLock()
	tmpl, ok := e.templates[name]
	e.mu.RUnlock()

	if !ok {
		return fmt.Errorf("template not found: %s", name)
	}

	return tmpl.Execute(w, data)
}

// Has checks if a template exists.
// Has는 템플릿이 존재하는지 확인합니다.
func (e *TemplateEngine) Has(name string) bool {
	e.mu.RLock()
	defer e.mu.RUnlock()
	_, ok := e.templates[name]
	return ok
}

// List returns all loaded template names.
// List는 모든 로드된 템플릿 이름을 반환합니다.
func (e *TemplateEngine) List() []string {
	e.mu.RLock()
	defer e.mu.RUnlock()

	names := make([]string, 0, len(e.templates))
	for name := range e.templates {
		names = append(names, name)
	}
	return names
}

// Clear removes all loaded templates.
// Clear는 모든 로드된 템플릿을 제거합니다.
func (e *TemplateEngine) Clear() {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.templates = make(map[string]*template.Template)
}

// SetLayoutDir sets the layout templates directory.
// SetLayoutDir는 레이아웃 템플릿 디렉토리를 설정합니다.
//
// Example / 예제:
//
//	engine.SetLayoutDir("views/layouts")
func (e *TemplateEngine) SetLayoutDir(dir string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.layoutDir = dir
}

// LoadLayout loads a layout template.
// LoadLayout는 레이아웃 템플릿을 로드합니다.
//
// Example / 예제:
//
//	err := engine.LoadLayout("base.html")
func (e *TemplateEngine) LoadLayout(name string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	// Build full path
	// 전체 경로 생성
	path := filepath.Join(e.layoutDir, name)

	// Check if file exists
	// 파일이 존재하는지 확인
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("layout file not found: %s", path)
	}

	// Create new template
	// 새 템플릿 생성
	tmpl := template.New(name)

	// Set custom delimiters
	// 커스텀 구분자 설정
	tmpl.Delims(e.delims[0], e.delims[1])

	// Add custom functions
	// 커스텀 함수 추가
	if len(e.funcMap) > 0 {
		tmpl.Funcs(e.funcMap)
	}

	// Parse layout file
	// 레이아웃 파일 파싱
	tmpl, err := tmpl.ParseFiles(path)
	if err != nil {
		return fmt.Errorf("failed to parse layout %s: %w", name, err)
	}

	// Store layout
	// 레이아웃 저장
	e.layouts[name] = tmpl

	return nil
}

// LoadAllLayouts loads all layout templates from the layout directory.
// LoadAllLayouts는 레이아웃 디렉토리에서 모든 레이아웃 템플릿을 로드합니다.
//
// Example / 예제:
//
//	err := engine.LoadAllLayouts()
func (e *TemplateEngine) LoadAllLayouts() error {
	e.mu.Lock()
	defer e.mu.Unlock()

	// Check if layout directory exists
	// 레이아웃 디렉토리가 존재하는지 확인
	if _, err := os.Stat(e.layoutDir); os.IsNotExist(err) {
		// Layout directory doesn't exist, skip
		// 레이아웃 디렉토리가 존재하지 않으면 건너뜀
		return nil
	}

	// Walk through layout directory
	// 레이아웃 디렉토리 순회
	return filepath.Walk(e.layoutDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		// 디렉토리 건너뛰기
		if info.IsDir() {
			return nil
		}

		// Only load .html, .htm, .tmpl files
		// .html, .htm, .tmpl 파일만 로드
		ext := strings.ToLower(filepath.Ext(path))
		if ext != ".html" && ext != ".htm" && ext != ".tmpl" {
			return nil
		}

		// Get relative name
		// 상대 경로 이름 가져오기
		name, err := filepath.Rel(e.layoutDir, path)
		if err != nil {
			return fmt.Errorf("failed to get relative path for %s: %w", path, err)
		}

		// Normalize path separators
		// 경로 구분자 정규화
		name = filepath.ToSlash(name)

		// Create new template
		// 새 템플릿 생성
		tmpl := template.New(name)

		// Set custom delimiters
		// 커스텀 구분자 설정
		tmpl.Delims(e.delims[0], e.delims[1])

		// Add custom functions
		// 커스텀 함수 추가
		if len(e.funcMap) > 0 {
			tmpl.Funcs(e.funcMap)
		}

		// Parse layout file
		// 레이아웃 파일 파싱
		tmpl, err = tmpl.ParseFiles(path)
		if err != nil {
			return fmt.Errorf("failed to parse layout %s: %w", name, err)
		}

		// Store layout
		// 레이아웃 저장
		e.layouts[name] = tmpl

		return nil
	})
}

// RenderWithLayout renders a template with a layout.
// RenderWithLayout는 레이아웃과 함께 템플릿을 렌더링합니다.
//
// Example / 예제:
//
//	err := engine.RenderWithLayout(w, "base.html", "index.html", data)
func (e *TemplateEngine) RenderWithLayout(w io.Writer, layoutName, templateName string, data interface{}) error {
	e.mu.RLock()
	defer e.mu.RUnlock()

	// Get layout template
	// 레이아웃 템플릿 가져오기
	layout, ok := e.layouts[layoutName]
	if !ok {
		return fmt.Errorf("layout not found: %s", layoutName)
	}

	// Get content template
	// 콘텐츠 템플릿 가져오기
	content, ok := e.templates[templateName]
	if !ok {
		return fmt.Errorf("template not found: %s", templateName)
	}

	// Clone the layout
	// 레이아웃 복제
	layoutClone, err := layout.Clone()
	if err != nil {
		return fmt.Errorf("failed to clone layout: %w", err)
	}

	// Add content template as "content" to the layout
	// 콘텐츠 템플릿을 "content"로 레이아웃에 추가
	_, err = layoutClone.AddParseTree("content", content.Tree)
	if err != nil {
		return fmt.Errorf("failed to add content to layout: %w", err)
	}

	// Execute layout (execute the main template with layout name)
	// 레이아웃 실행 (레이아웃 이름으로 메인 템플릿 실행)
	return layoutClone.ExecuteTemplate(w, layoutName, data)
}

// HasLayout checks if a layout exists.
// HasLayout는 레이아웃이 존재하는지 확인합니다.
func (e *TemplateEngine) HasLayout(name string) bool {
	e.mu.RLock()
	defer e.mu.RUnlock()
	_, ok := e.layouts[name]
	return ok
}

// ListLayouts returns all loaded layout names.
// ListLayouts는 모든 로드된 레이아웃 이름을 반환합니다.
func (e *TemplateEngine) ListLayouts() []string {
	e.mu.RLock()
	defer e.mu.RUnlock()

	names := make([]string, 0, len(e.layouts))
	for name := range e.layouts {
		names = append(names, name)
	}
	return names
}
