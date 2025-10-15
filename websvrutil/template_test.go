package websvrutil

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestNewTemplateEngine tests creating a new template engine
// TestNewTemplateEngine은 새 템플릿 엔진 생성을 테스트합니다
func TestNewTemplateEngine(t *testing.T) {
	engine := NewTemplateEngine("views")
	if engine == nil {
		t.Fatal("Expected non-nil engine")
	}
	if engine.dir != "views" {
		t.Errorf("Expected dir 'views', got %s", engine.dir)
	}
	if engine.templates == nil {
		t.Error("Expected non-nil templates map")
	}
	if engine.funcMap == nil {
		t.Error("Expected non-nil funcMap")
	}
}

// TestSetDelimiters tests setting custom delimiters
// TestSetDelimiters는 커스텀 구분자 설정을 테스트합니다
func TestSetDelimiters(t *testing.T) {
	engine := NewTemplateEngine("views")
	engine.SetDelimiters("[[", "]]")

	if engine.delims[0] != "[[" || engine.delims[1] != "]]" {
		t.Errorf("Expected delimiters '[[' and ']]', got '%s' and '%s'", engine.delims[0], engine.delims[1])
	}
}

// TestAddFunc tests adding a custom function
// TestAddFunc는 커스텀 함수 추가를 테스트합니다
func TestAddFunc(t *testing.T) {
	engine := NewTemplateEngine("views")
	engine.AddFunc("upper", strings.ToUpper)

	if _, ok := engine.funcMap["upper"]; !ok {
		t.Error("Expected 'upper' function in funcMap")
	}
}

// TestAddFuncs tests adding multiple custom functions
// TestAddFuncs는 여러 커스텀 함수 추가를 테스트합니다
func TestAddFuncs(t *testing.T) {
	engine := NewTemplateEngine("views")
	engine.AddFuncs(template.FuncMap{
		"upper": strings.ToUpper,
		"lower": strings.ToLower,
	})

	if _, ok := engine.funcMap["upper"]; !ok {
		t.Error("Expected 'upper' function in funcMap")
	}
	if _, ok := engine.funcMap["lower"]; !ok {
		t.Error("Expected 'lower' function in funcMap")
	}
}

// TestLoad tests loading a single template
// TestLoad는 단일 템플릿 로드를 테스트합니다
func TestLoad(t *testing.T) {
	// Create temp directory
	// 임시 디렉토리 생성
	tmpDir := t.TempDir()

	// Create test template
	// 테스트 템플릿 생성
	tmplPath := filepath.Join(tmpDir, "test.html")
	content := "<h1>Hello {{.Name}}</h1>"
	os.WriteFile(tmplPath, []byte(content), 0644)

	// Create engine and load template
	// 엔진 생성 및 템플릿 로드
	engine := NewTemplateEngine(tmpDir)
	err := engine.Load("test.html")
	if err != nil {
		t.Fatalf("Failed to load template: %v", err)
	}

	// Check if template exists
	// 템플릿이 존재하는지 확인
	if !engine.Has("test.html") {
		t.Error("Expected template 'test.html' to be loaded")
	}
}

// TestLoadNonExistent tests loading a non-existent template
// TestLoadNonExistent는 존재하지 않는 템플릿 로드를 테스트합니다
func TestLoadNonExistent(t *testing.T) {
	tmpDir := t.TempDir()
	engine := NewTemplateEngine(tmpDir)

	err := engine.Load("nonexistent.html")
	if err == nil {
		t.Error("Expected error when loading non-existent template")
	}
}

// TestLoadGlob tests loading templates with glob pattern
// TestLoadGlob는 glob 패턴으로 템플릿 로드를 테스트합니다
func TestLoadGlob(t *testing.T) {
	tmpDir := t.TempDir()

	// Create multiple templates
	// 여러 템플릿 생성
	os.WriteFile(filepath.Join(tmpDir, "one.html"), []byte("<h1>One</h1>"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "two.html"), []byte("<h1>Two</h1>"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "three.txt"), []byte("Three"), 0644) // Should be ignored

	engine := NewTemplateEngine(tmpDir)
	err := engine.LoadGlob("*.html")
	if err != nil {
		t.Fatalf("Failed to load templates: %v", err)
	}

	if !engine.Has("one.html") {
		t.Error("Expected template 'one.html' to be loaded")
	}
	if !engine.Has("two.html") {
		t.Error("Expected template 'two.html' to be loaded")
	}
	if engine.Has("three.txt") {
		t.Error("Did not expect 'three.txt' to be loaded")
	}
}

// TestLoadAll tests loading all templates from directory
// TestLoadAll은 디렉토리에서 모든 템플릿 로드를 테스트합니다
func TestLoadAll(t *testing.T) {
	tmpDir := t.TempDir()

	// Create nested directory structure
	// 중첩 디렉토리 구조 생성
	subDir := filepath.Join(tmpDir, "sub")
	os.Mkdir(subDir, 0755)

	// Create templates
	// 템플릿 생성
	os.WriteFile(filepath.Join(tmpDir, "index.html"), []byte("<h1>Index</h1>"), 0644)
	os.WriteFile(filepath.Join(subDir, "page.html"), []byte("<h1>Page</h1>"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "readme.txt"), []byte("README"), 0644) // Should be ignored

	engine := NewTemplateEngine(tmpDir)
	err := engine.LoadAll()
	if err != nil {
		t.Fatalf("Failed to load all templates: %v", err)
	}

	if !engine.Has("index.html") {
		t.Error("Expected template 'index.html' to be loaded")
	}
	if !engine.Has("sub/page.html") {
		t.Error("Expected template 'sub/page.html' to be loaded")
	}
	if engine.Has("readme.txt") {
		t.Error("Did not expect 'readme.txt' to be loaded")
	}
}

// TestRender tests rendering a template
// TestRender는 템플릿 렌더링을 테스트합니다
func TestRender(t *testing.T) {
	tmpDir := t.TempDir()

	// Create template
	// 템플릿 생성
	tmplPath := filepath.Join(tmpDir, "hello.html")
	content := "<h1>Hello {{.Name}}</h1>"
	os.WriteFile(tmplPath, []byte(content), 0644)

	// Load and render template
	// 템플릿 로드 및 렌더링
	engine := NewTemplateEngine(tmpDir)
	engine.Load("hello.html")

	var buf bytes.Buffer
	data := map[string]string{"Name": "World"}
	err := engine.Render(&buf, "hello.html", data)
	if err != nil {
		t.Fatalf("Failed to render template: %v", err)
	}

	expected := "<h1>Hello World</h1>"
	if buf.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, buf.String())
	}
}

// TestRenderNonExistent tests rendering a non-existent template
// TestRenderNonExistent는 존재하지 않는 템플릿 렌더링을 테스트합니다
func TestRenderNonExistent(t *testing.T) {
	engine := NewTemplateEngine("views")

	var buf bytes.Buffer
	err := engine.Render(&buf, "nonexistent.html", nil)
	if err == nil {
		t.Error("Expected error when rendering non-existent template")
	}
}

// TestRenderWithCustomFunc tests rendering with custom function
// TestRenderWithCustomFunc는 커스텀 함수로 렌더링을 테스트합니다
func TestRenderWithCustomFunc(t *testing.T) {
	tmpDir := t.TempDir()

	// Create template with custom function
	// 커스텀 함수가 있는 템플릿 생성
	tmplPath := filepath.Join(tmpDir, "func.html")
	content := "<h1>{{upper .Name}}</h1>"
	os.WriteFile(tmplPath, []byte(content), 0644)

	// Add custom function and load template
	// 커스텀 함수 추가 및 템플릿 로드
	engine := NewTemplateEngine(tmpDir)
	engine.AddFunc("upper", strings.ToUpper)
	engine.Load("func.html")

	var buf bytes.Buffer
	data := map[string]string{"Name": "world"}
	err := engine.Render(&buf, "func.html", data)
	if err != nil {
		t.Fatalf("Failed to render template: %v", err)
	}

	expected := "<h1>WORLD</h1>"
	if buf.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, buf.String())
	}
}

// TestHas tests checking if template exists
// TestHas는 템플릿 존재 확인을 테스트합니다
func TestHas(t *testing.T) {
	tmpDir := t.TempDir()
	os.WriteFile(filepath.Join(tmpDir, "test.html"), []byte("<h1>Test</h1>"), 0644)

	engine := NewTemplateEngine(tmpDir)
	engine.Load("test.html")

	if !engine.Has("test.html") {
		t.Error("Expected template 'test.html' to exist")
	}
	if engine.Has("nonexistent.html") {
		t.Error("Did not expect 'nonexistent.html' to exist")
	}
}

// TestList tests listing all loaded templates
// TestList는 모든 로드된 템플릿 목록 작성을 테스트합니다
func TestList(t *testing.T) {
	tmpDir := t.TempDir()

	os.WriteFile(filepath.Join(tmpDir, "one.html"), []byte("<h1>One</h1>"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "two.html"), []byte("<h1>Two</h1>"), 0644)

	engine := NewTemplateEngine(tmpDir)
	engine.LoadGlob("*.html")

	names := engine.List()
	if len(names) != 2 {
		t.Errorf("Expected 2 templates, got %d", len(names))
	}
}

// TestClear tests clearing all templates
// TestClear는 모든 템플릿 제거를 테스트합니다
func TestClear(t *testing.T) {
	tmpDir := t.TempDir()
	os.WriteFile(filepath.Join(tmpDir, "test.html"), []byte("<h1>Test</h1>"), 0644)

	engine := NewTemplateEngine(tmpDir)
	engine.Load("test.html")

	if !engine.Has("test.html") {
		t.Error("Expected template to be loaded")
	}

	engine.Clear()

	if engine.Has("test.html") {
		t.Error("Expected template to be cleared")
	}
}

// BenchmarkLoad benchmarks loading a template
// BenchmarkLoad는 템플릿 로드를 벤치마크합니다
func BenchmarkLoad(b *testing.B) {
	tmpDir := b.TempDir()
	tmplPath := filepath.Join(tmpDir, "test.html")
	content := "<h1>Hello {{.Name}}</h1>"
	os.WriteFile(tmplPath, []byte(content), 0644)

	engine := NewTemplateEngine(tmpDir)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		engine.Clear()
		engine.Load("test.html")
	}
}

// BenchmarkRender benchmarks rendering a template
// BenchmarkRender는 템플릿 렌더링을 벤치마크합니다
func BenchmarkRender(b *testing.B) {
	tmpDir := b.TempDir()
	tmplPath := filepath.Join(tmpDir, "test.html")
	content := "<h1>Hello {{.Name}}</h1>"
	os.WriteFile(tmplPath, []byte(content), 0644)

	engine := NewTemplateEngine(tmpDir)
	engine.Load("test.html")

	data := map[string]string{"Name": "World"}
	var buf bytes.Buffer

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		engine.Render(&buf, "test.html", data)
	}
}

// TestBuiltinFuncs tests built-in template functions
// TestBuiltinFuncs는 내장 템플릿 함수를 테스트합니다
func TestBuiltinFuncs(t *testing.T) {
	tmpDir := t.TempDir()

	// Test upper function
	t.Run("Upper", func(t *testing.T) {
		tmplPath := filepath.Join(tmpDir, "upper.html")
		content := "{{upper .Name}}"
		os.WriteFile(tmplPath, []byte(content), 0644)

		engine := NewTemplateEngine(tmpDir)
		engine.Load("upper.html")

		var buf bytes.Buffer
		data := map[string]string{"Name": "hello"}
		err := engine.Render(&buf, "upper.html", data)
		if err != nil {
			t.Fatalf("Failed to render: %v", err)
		}

		if buf.String() != "HELLO" {
			t.Errorf("Expected 'HELLO', got '%s'", buf.String())
		}
	})

	// Test lower function
	t.Run("Lower", func(t *testing.T) {
		tmplPath := filepath.Join(tmpDir, "lower.html")
		content := "{{lower .Name}}"
		os.WriteFile(tmplPath, []byte(content), 0644)

		engine := NewTemplateEngine(tmpDir)
		engine.Load("lower.html")

		var buf bytes.Buffer
		data := map[string]string{"Name": "HELLO"}
		err := engine.Render(&buf, "lower.html", data)
		if err != nil {
			t.Fatalf("Failed to render: %v", err)
		}

		if buf.String() != "hello" {
			t.Errorf("Expected 'hello', got '%s'", buf.String())
		}
	})

	// Test safeHTML function
	t.Run("SafeHTML", func(t *testing.T) {
		tmplPath := filepath.Join(tmpDir, "safe.html")
		content := "{{safeHTML .HTML}}"
		os.WriteFile(tmplPath, []byte(content), 0644)

		engine := NewTemplateEngine(tmpDir)
		engine.Load("safe.html")

		var buf bytes.Buffer
		data := map[string]string{"HTML": "<b>Bold</b>"}
		err := engine.Render(&buf, "safe.html", data)
		if err != nil {
			t.Fatalf("Failed to render: %v", err)
		}

		if buf.String() != "<b>Bold</b>" {
			t.Errorf("Expected '<b>Bold</b>', got '%s'", buf.String())
		}
	})
}

// TestLoadLayout tests loading a layout template
// TestLoadLayout는 레이아웃 템플릿 로드를 테스트합니다
func TestLoadLayout(t *testing.T) {
	tmpDir := t.TempDir()
	layoutDir := filepath.Join(tmpDir, "layouts")
	os.Mkdir(layoutDir, 0755)

	// Create layout template
	layoutPath := filepath.Join(layoutDir, "base.html")
	layoutContent := `<!DOCTYPE html><html><body>{{template "content" .}}</body></html>`
	os.WriteFile(layoutPath, []byte(layoutContent), 0644)

	engine := NewTemplateEngine(tmpDir)
	err := engine.LoadLayout("base.html")
	if err != nil {
		t.Fatalf("Failed to load layout: %v", err)
	}

	if !engine.HasLayout("base.html") {
		t.Error("Expected layout 'base.html' to be loaded")
	}
}

// TestLoadAllLayouts tests loading all layouts
// TestLoadAllLayouts는 모든 레이아웃 로드를 테스트합니다
func TestLoadAllLayouts(t *testing.T) {
	tmpDir := t.TempDir()
	layoutDir := filepath.Join(tmpDir, "layouts")
	os.Mkdir(layoutDir, 0755)

	// Create multiple layouts
	os.WriteFile(filepath.Join(layoutDir, "base.html"), []byte("Base Layout"), 0644)
	os.WriteFile(filepath.Join(layoutDir, "admin.html"), []byte("Admin Layout"), 0644)

	engine := NewTemplateEngine(tmpDir)
	err := engine.LoadAllLayouts()
	if err != nil {
		t.Fatalf("Failed to load layouts: %v", err)
	}

	if !engine.HasLayout("base.html") {
		t.Error("Expected layout 'base.html' to be loaded")
	}
	if !engine.HasLayout("admin.html") {
		t.Error("Expected layout 'admin.html' to be loaded")
	}
}

// TestRenderWithLayout tests rendering with layout
// TestRenderWithLayout는 레이아웃과 함께 렌더링을 테스트합니다
func TestRenderWithLayout(t *testing.T) {
	tmpDir := t.TempDir()
	layoutDir := filepath.Join(tmpDir, "layouts")
	os.Mkdir(layoutDir, 0755)

	// Create layout
	layoutPath := filepath.Join(layoutDir, "base.html")
	layoutContent := `<html><body>{{template "content" .}}</body></html>`
	os.WriteFile(layoutPath, []byte(layoutContent), 0644)

	// Create content template
	contentPath := filepath.Join(tmpDir, "index.html")
	contentContent := `<h1>{{.Title}}</h1>`
	os.WriteFile(contentPath, []byte(contentContent), 0644)

	engine := NewTemplateEngine(tmpDir)
	engine.LoadLayout("base.html")
	engine.Load("index.html")

	var buf bytes.Buffer
	data := map[string]string{"Title": "Hello"}
	err := engine.RenderWithLayout(&buf, "base.html", "index.html", data)
	if err != nil {
		t.Fatalf("Failed to render with layout: %v", err)
	}

	expected := "<html><body><h1>Hello</h1></body></html>"
	if buf.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, buf.String())
	}
}

// TestSetLayoutDir tests setting layout directory
// TestSetLayoutDir는 레이아웃 디렉토리 설정을 테스트합니다
func TestSetLayoutDir(t *testing.T) {
	engine := NewTemplateEngine("views")
	engine.SetLayoutDir("custom/layouts")

	if engine.layoutDir != "custom/layouts" {
		t.Errorf("Expected layout dir 'custom/layouts', got '%s'", engine.layoutDir)
	}
}

// TestListLayouts tests listing all layouts
// TestListLayouts는 모든 레이아웃 목록을 테스트합니다
func TestListLayouts(t *testing.T) {
	tmpDir := t.TempDir()
	layoutDir := filepath.Join(tmpDir, "layouts")
	os.Mkdir(layoutDir, 0755)

	os.WriteFile(filepath.Join(layoutDir, "one.html"), []byte("One"), 0644)
	os.WriteFile(filepath.Join(layoutDir, "two.html"), []byte("Two"), 0644)

	engine := NewTemplateEngine(tmpDir)
	engine.LoadAllLayouts()

	layouts := engine.ListLayouts()
	if len(layouts) != 2 {
		t.Errorf("Expected 2 layouts, got %d", len(layouts))
	}
}

// BenchmarkBuiltinFuncs benchmarks built-in template functions
// BenchmarkBuiltinFuncs는 내장 템플릿 함수를 벤치마크합니다
func BenchmarkBuiltinFuncs(b *testing.B) {
	tmpDir := b.TempDir()
	tmplPath := filepath.Join(tmpDir, "test.html")
	content := "{{upper .Name}}"
	os.WriteFile(tmplPath, []byte(content), 0644)

	engine := NewTemplateEngine(tmpDir)
	engine.Load("test.html")

	data := map[string]string{"Name": "hello"}
	var buf bytes.Buffer

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		engine.Render(&buf, "test.html", data)
	}
}

// BenchmarkRenderWithLayout benchmarks rendering with layout
// BenchmarkRenderWithLayout는 레이아웃과 함께 렌더링을 벤치마크합니다
func BenchmarkRenderWithLayout(b *testing.B) {
	tmpDir := b.TempDir()
	layoutDir := filepath.Join(tmpDir, "layouts")
	os.Mkdir(layoutDir, 0755)

	layoutPath := filepath.Join(layoutDir, "base.html")
	layoutContent := `<html><body>{{template "content" .}}</body></html>`
	os.WriteFile(layoutPath, []byte(layoutContent), 0644)

	contentPath := filepath.Join(tmpDir, "index.html")
	contentContent := `<h1>{{.Title}}</h1>`
	os.WriteFile(contentPath, []byte(contentContent), 0644)

	engine := NewTemplateEngine(tmpDir)
	engine.LoadLayout("base.html")
	engine.Load("index.html")

	data := map[string]string{"Title": "Hello"}
	var buf bytes.Buffer

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		engine.RenderWithLayout(&buf, "base.html", "index.html", data)
	}
}

// TestEnableAutoReload tests enabling auto-reload
// TestEnableAutoReload은 자동 재로드 활성화를 테스트합니다
func TestEnableAutoReload(t *testing.T) {
	tmpDir := t.TempDir()
	engine := NewTemplateEngine(tmpDir)

	if engine.IsAutoReloadEnabled() {
		t.Error("Expected auto-reload to be disabled by default")
	}

	err := engine.EnableAutoReload()
	if err != nil {
		t.Fatalf("Failed to enable auto-reload: %v", err)
	}

	if !engine.IsAutoReloadEnabled() {
		t.Error("Expected auto-reload to be enabled")
	}

	// Enable again should not error
	// 다시 활성화해도 에러가 발생하지 않아야 함
	err = engine.EnableAutoReload()
	if err != nil {
		t.Fatalf("Failed to enable auto-reload again: %v", err)
	}

	// Cleanup / 정리
	engine.DisableAutoReload()
}

// TestDisableAutoReload tests disabling auto-reload
// TestDisableAutoReload은 자동 재로드 비활성화를 테스트합니다
func TestDisableAutoReload(t *testing.T) {
	tmpDir := t.TempDir()
	engine := NewTemplateEngine(tmpDir)

	engine.EnableAutoReload()
	if !engine.IsAutoReloadEnabled() {
		t.Error("Expected auto-reload to be enabled")
	}

	engine.DisableAutoReload()
	if engine.IsAutoReloadEnabled() {
		t.Error("Expected auto-reload to be disabled")
	}

	// Disable again should not error
	// 다시 비활성화해도 에러가 발생하지 않아야 함
	engine.DisableAutoReload()
}

// TestIsAutoReloadEnabled tests checking auto-reload status
// TestIsAutoReloadEnabled는 자동 재로드 상태 확인을 테스트합니다
func TestIsAutoReloadEnabled(t *testing.T) {
	tmpDir := t.TempDir()
	engine := NewTemplateEngine(tmpDir)

	// Initially disabled / 초기에는 비활성화
	if engine.IsAutoReloadEnabled() {
		t.Error("Expected auto-reload to be disabled initially")
	}

	// Enable / 활성화
	engine.EnableAutoReload()
	if !engine.IsAutoReloadEnabled() {
		t.Error("Expected auto-reload to be enabled after calling EnableAutoReload")
	}

	// Disable / 비활성화
	engine.DisableAutoReload()
	if engine.IsAutoReloadEnabled() {
		t.Error("Expected auto-reload to be disabled after calling DisableAutoReload")
	}
}
