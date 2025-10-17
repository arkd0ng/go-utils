package websvrutil

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestStaticFileServing tests the Static method
// Static 메서드 테스트
func TestStaticFileServing(t *testing.T) {
	// Create temporary directory with test files
	// 테스트 파일이 있는 임시 디렉토리 생성
	tmpDir := t.TempDir()

	// Create test files
	// 테스트 파일 생성
	testFiles := map[string]string{
		"index.html":    "<html><body>Index</body></html>",
		"style.css":     "body { margin: 0; }",
		"script.js":     "console.log('Hello');",
		"image.png":     "fake png content",
		"subdir/sub.txt": "subdirectory file",
	}

	for path, content := range testFiles {
		fullPath := filepath.Join(tmpDir, path)
		if strings.Contains(path, "/") {
			dir := filepath.Dir(fullPath)
			if err := os.MkdirAll(dir, 0755); err != nil {
				t.Fatalf("Failed to create directory: %v", err)
			}
		}
		if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
			t.Fatalf("Failed to create test file %s: %v", path, err)
		}
	}

	// Create app with static route
	// 정적 라우트가 있는 앱 생성
	app := New()
	app.Static("/static", tmpDir)

	// Create server to handle redirects
	// 리디렉트를 처리하는 서버 생성
	server := httptest.NewServer(app)
	defer server.Close()

	// Test index.html
	// index.html 테스트
	resp, err := http.Get(server.URL + "/static/index.html")
	if err != nil {
		t.Fatalf("Failed to get file: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	body := make([]byte, 1024)
	n, _ := resp.Body.Read(body)
	if !strings.Contains(string(body[:n]), "Index") {
		t.Errorf("Expected body to contain 'Index', got %q", string(body[:n]))
	}

	// Test CSS file
	// CSS 파일 테스트
	resp, err = http.Get(server.URL + "/static/style.css")
	if err != nil {
		t.Fatalf("Failed to get CSS file: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	body = make([]byte, 1024)
	n, _ = resp.Body.Read(body)
	if !strings.Contains(string(body[:n]), "margin") {
		t.Errorf("Expected body to contain 'margin', got %q", string(body[:n]))
	}

	// Test subdirectory file
	// 하위 디렉토리 파일 테스트
	resp, err = http.Get(server.URL + "/static/subdir/sub.txt")
	if err != nil {
		t.Fatalf("Failed to get subdirectory file: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	body = make([]byte, 1024)
	n, _ = resp.Body.Read(body)
	if !strings.Contains(string(body[:n]), "subdirectory") {
		t.Errorf("Expected body to contain 'subdirectory', got %q", string(body[:n]))
	}
}

// TestStaticNotFound tests 404 for non-existent static files
// 존재하지 않는 정적 파일에 대한 404 테스트
func TestStaticNotFound(t *testing.T) {
	tmpDir := t.TempDir()

	app := New()
	app.Static("/static", tmpDir)

	// Test non-existent file
	// 존재하지 않는 파일 테스트
	req := httptest.NewRequest(http.MethodGet, "/static/nonexistent.txt", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}

// TestStaticMultiplePrefixes tests multiple static directories
// 여러 정적 디렉토리 테스트
func TestStaticMultiplePrefixes(t *testing.T) {
	// Create two temporary directories
	// 두 개의 임시 디렉토리 생성
	tmpDir1 := t.TempDir()
	tmpDir2 := t.TempDir()

	// Create test files in each directory
	// 각 디렉토리에 테스트 파일 생성
	os.WriteFile(filepath.Join(tmpDir1, "file1.txt"), []byte("content1"), 0644)
	os.WriteFile(filepath.Join(tmpDir2, "file2.txt"), []byte("content2"), 0644)

	app := New()
	app.Static("/static1", tmpDir1)
	app.Static("/static2", tmpDir2)

	// Test first static directory
	// 첫 번째 정적 디렉토리 테스트
	req := httptest.NewRequest(http.MethodGet, "/static1/file1.txt", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	if w.Body.String() != "content1" {
		t.Errorf("Expected 'content1', got %q", w.Body.String())
	}

	// Test second static directory
	// 두 번째 정적 디렉토리 테스트
	req = httptest.NewRequest(http.MethodGet, "/static2/file2.txt", nil)
	w = httptest.NewRecorder()
	app.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	if w.Body.String() != "content2" {
		t.Errorf("Expected 'content2', got %q", w.Body.String())
	}
}

// TestFile tests the File method
// File 메서드 테스트
func TestFile(t *testing.T) {
	// Create temporary test file
	// 임시 테스트 파일 생성
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.html")
	testContent := "<html><body>Test File</body></html>"

	if err := os.WriteFile(testFile, []byte(testContent), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	app := New()
	app.GET("/file", func(w http.ResponseWriter, r *http.Request) {
		ctx := GetContext(r)
		ctx.File(testFile)
	})

	req := httptest.NewRequest(http.MethodGet, "/file", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	if w.Body.String() != testContent {
		t.Errorf("Expected %q, got %q", testContent, w.Body.String())
	}

	// Check Content-Type header
	// Content-Type 헤더 확인
	contentType := w.Header().Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		t.Errorf("Expected Content-Type to contain 'text/html', got %q", contentType)
	}
}

// TestFileNotFound tests File with non-existent file
// 존재하지 않는 파일로 File 테스트
func TestFileNotFound(t *testing.T) {
	app := New()
	app.GET("/file", func(w http.ResponseWriter, r *http.Request) {
		ctx := GetContext(r)
		ctx.File("/nonexistent/path/file.txt")
	})

	req := httptest.NewRequest(http.MethodGet, "/file", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}

// TestFileAttachment tests the FileAttachment method
// FileAttachment 메서드 테스트
func TestFileAttachment(t *testing.T) {
	// Create temporary test file
	// 임시 테스트 파일 생성
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "report.pdf")
	testContent := "PDF content here"

	if err := os.WriteFile(testFile, []byte(testContent), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	app := New()
	app.GET("/download", func(w http.ResponseWriter, r *http.Request) {
		ctx := GetContext(r)
		ctx.FileAttachment(testFile, "monthly-report.pdf")
	})

	req := httptest.NewRequest(http.MethodGet, "/download", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	if w.Body.String() != testContent {
		t.Errorf("Expected %q, got %q", testContent, w.Body.String())
	}

	// Check Content-Disposition header
	// Content-Disposition 헤더 확인
	disposition := w.Header().Get("Content-Disposition")
	expected := `attachment; filename="monthly-report.pdf"`
	if disposition != expected {
		t.Errorf("Expected Content-Disposition %q, got %q", expected, disposition)
	}
}

// TestFileAttachmentDifferentTypes tests FileAttachment with various file types
// 다양한 파일 타입으로 FileAttachment 테스트
func TestFileAttachmentDifferentTypes(t *testing.T) {
	tmpDir := t.TempDir()

	testCases := []struct {
		name        string
		filename    string
		content     string
		contentType string
	}{
		{"PDF", "document.pdf", "PDF content", "application/pdf"},
		{"ZIP", "archive.zip", "ZIP content", "application/zip"},
		{"Text", "notes.txt", "Text content", "text/plain"},
		{"Image", "photo.jpg", "JPEG content", "image/jpeg"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testFile := filepath.Join(tmpDir, tc.filename)
			if err := os.WriteFile(testFile, []byte(tc.content), 0644); err != nil {
				t.Fatalf("Failed to create test file: %v", err)
			}

			app := New()
			app.GET("/download", func(w http.ResponseWriter, r *http.Request) {
				ctx := GetContext(r)
				ctx.FileAttachment(testFile, tc.filename)
			})

			req := httptest.NewRequest(http.MethodGet, "/download", nil)
			w := httptest.NewRecorder()
			app.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Errorf("Expected status 200, got %d", w.Code)
			}

			disposition := w.Header().Get("Content-Disposition")
			expectedDisposition := `attachment; filename="` + tc.filename + `"`
			if disposition != expectedDisposition {
				t.Errorf("Expected Content-Disposition %q, got %q", expectedDisposition, disposition)
			}
		})
	}
}

// BenchmarkStaticFileServing benchmarks static file serving
// 정적 파일 서빙 벤치마크
func BenchmarkStaticFileServing(b *testing.B) {
	tmpDir := b.TempDir()
	testFile := filepath.Join(tmpDir, "bench.txt")
	os.WriteFile(testFile, []byte("benchmark content"), 0644)

	app := New()
	app.Static("/static", tmpDir)

	req := httptest.NewRequest(http.MethodGet, "/static/bench.txt", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		app.ServeHTTP(w, req)
	}
}

// BenchmarkFile benchmarks the File method
// File 메서드 벤치마크
func BenchmarkFile(b *testing.B) {
	tmpDir := b.TempDir()
	testFile := filepath.Join(tmpDir, "bench.html")
	os.WriteFile(testFile, []byte("<html><body>Benchmark</body></html>"), 0644)

	app := New()
	app.GET("/file", func(w http.ResponseWriter, r *http.Request) {
		ctx := GetContext(r)
		ctx.File(testFile)
	})

	req := httptest.NewRequest(http.MethodGet, "/file", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		app.ServeHTTP(w, req)
	}
}
