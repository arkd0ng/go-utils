package websvrutil

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

// TestFormFile tests the FormFile method
// FormFile 메서드 테스트
func TestFormFile(t *testing.T) {
	// Create multipart form with file
	// 파일이 포함된 multipart 폼 생성
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add file field
	// 파일 필드 추가
	fileWriter, err := writer.CreateFormFile("upload", "test.txt")
	if err != nil {
		t.Fatalf("Failed to create form file: %v", err)
	}

	// Write test content
	// 테스트 내용 작성
	fileContent := []byte("Hello, World!")
	if _, err := fileWriter.Write(fileContent); err != nil {
		t.Fatalf("Failed to write file content: %v", err)
	}

	writer.Close()

	// Create HTTP request
	// HTTP 요청 생성
	req := httptest.NewRequest(http.MethodPost, "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Create context
	// 컨텍스트 생성
	w := httptest.NewRecorder()
	ctx := NewContext(w, req)

	// Test FormFile
	// FormFile 테스트
	fileHeader, err := ctx.FormFile("upload")
	if err != nil {
		t.Fatalf("FormFile() error = %v", err)
	}

	if fileHeader.Filename != "test.txt" {
		t.Errorf("Expected filename 'test.txt', got '%s'", fileHeader.Filename)
	}

	if fileHeader.Size != int64(len(fileContent)) {
		t.Errorf("Expected size %d, got %d", len(fileContent), fileHeader.Size)
	}
}

// TestFormFileNotFound tests FormFile with non-existent field
// 존재하지 않는 필드로 FormFile 테스트
func TestFormFileNotFound(t *testing.T) {
	// Create empty multipart form
	// 빈 multipart 폼 생성
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	ctx := NewContext(w, req)

	// Test FormFile with non-existent field
	// 존재하지 않는 필드로 FormFile 테스트
	_, err := ctx.FormFile("nonexistent")
	if err == nil {
		t.Error("Expected error for non-existent field, got nil")
	}
}

// TestMultipartForm tests the MultipartForm method
// MultipartForm 메서드 테스트
func TestMultipartForm(t *testing.T) {
	// Create multipart form with file and fields
	// 파일과 필드가 포함된 multipart 폼 생성
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add file
	// 파일 추가
	fileWriter, _ := writer.CreateFormFile("file", "test.txt")
	fileWriter.Write([]byte("test content"))

	// Add form fields
	// 폼 필드 추가
	writer.WriteField("name", "John Doe")
	writer.WriteField("email", "john@example.com")

	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	ctx := NewContext(w, req)

	// Test MultipartForm
	// MultipartForm 테스트
	form, err := ctx.MultipartForm()
	if err != nil {
		t.Fatalf("MultipartForm() error = %v", err)
	}

	// Check file
	// 파일 확인
	if len(form.File["file"]) != 1 {
		t.Errorf("Expected 1 file, got %d", len(form.File["file"]))
	}

	// Check form fields
	// 폼 필드 확인
	if name := form.Value["name"][0]; name != "John Doe" {
		t.Errorf("Expected name 'John Doe', got '%s'", name)
	}

	if email := form.Value["email"][0]; email != "john@example.com" {
		t.Errorf("Expected email 'john@example.com', got '%s'", email)
	}
}

// TestMultipartFormMultipleFiles tests MultipartForm with multiple files
// 여러 파일로 MultipartForm 테스트
func TestMultipartFormMultipleFiles(t *testing.T) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add multiple files
	// 여러 파일 추가
	for i := 1; i <= 3; i++ {
		fileWriter, _ := writer.CreateFormFile("files", "test"+string(rune('0'+i))+".txt")
		fileWriter.Write([]byte("content " + string(rune('0'+i))))
	}

	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	ctx := NewContext(w, req)

	form, err := ctx.MultipartForm()
	if err != nil {
		t.Fatalf("MultipartForm() error = %v", err)
	}

	// Check file count
	// 파일 수 확인
	if len(form.File["files"]) != 3 {
		t.Errorf("Expected 3 files, got %d", len(form.File["files"]))
	}
}

// TestSaveUploadedFile tests the SaveUploadedFile method
// SaveUploadedFile 메서드 테스트
func TestSaveUploadedFile(t *testing.T) {
	// Create temporary directory for test
	// 테스트용 임시 디렉토리 생성
	tmpDir := t.TempDir()

	// Create multipart form with file
	// 파일이 포함된 multipart 폼 생성
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fileContent := []byte("This is test file content")
	fileWriter, _ := writer.CreateFormFile("upload", "test.txt")
	fileWriter.Write(fileContent)
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	ctx := NewContext(w, req)

	// Get file header
	// 파일 헤더 가져오기
	fileHeader, err := ctx.FormFile("upload")
	if err != nil {
		t.Fatalf("FormFile() error = %v", err)
	}

	// Save file
	// 파일 저장
	dstPath := filepath.Join(tmpDir, "saved.txt")
	if err := ctx.SaveUploadedFile(fileHeader, dstPath); err != nil {
		t.Fatalf("SaveUploadedFile() error = %v", err)
	}

	// Verify file exists
	// 파일 존재 확인
	if _, err := os.Stat(dstPath); os.IsNotExist(err) {
		t.Error("Saved file does not exist")
	}

	// Verify file content
	// 파일 내용 확인
	savedContent, err := os.ReadFile(dstPath)
	if err != nil {
		t.Fatalf("Failed to read saved file: %v", err)
	}

	if !bytes.Equal(savedContent, fileContent) {
		t.Errorf("File content mismatch. Expected %q, got %q", fileContent, savedContent)
	}
}

// TestSaveUploadedFileLargeFile tests saving a large file
// 큰 파일 저장 테스트
func TestSaveUploadedFileLargeFile(t *testing.T) {
	tmpDir := t.TempDir()

	// Create 1MB file content
	// 1MB 파일 내용 생성
	largeContent := make([]byte, 1<<20) // 1 MB
	for i := range largeContent {
		largeContent[i] = byte(i % 256)
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fileWriter, _ := writer.CreateFormFile("upload", "large.bin")
	fileWriter.Write(largeContent)
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	ctx := NewContext(w, req)

	fileHeader, err := ctx.FormFile("upload")
	if err != nil {
		t.Fatalf("FormFile() error = %v", err)
	}

	dstPath := filepath.Join(tmpDir, "large.bin")
	if err := ctx.SaveUploadedFile(fileHeader, dstPath); err != nil {
		t.Fatalf("SaveUploadedFile() error = %v", err)
	}

	// Verify file size
	// 파일 크기 확인
	fileInfo, err := os.Stat(dstPath)
	if err != nil {
		t.Fatalf("Failed to stat saved file: %v", err)
	}

	if fileInfo.Size() != int64(len(largeContent)) {
		t.Errorf("File size mismatch. Expected %d, got %d", len(largeContent), fileInfo.Size())
	}
}

// TestSaveUploadedFileError tests error handling in SaveUploadedFile
// SaveUploadedFile 에러 처리 테스트
func TestSaveUploadedFileError(t *testing.T) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fileWriter, _ := writer.CreateFormFile("upload", "test.txt")
	fileWriter.Write([]byte("test"))
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	ctx := NewContext(w, req)

	fileHeader, _ := ctx.FormFile("upload")

	// Try to save to invalid path
	// 잘못된 경로에 저장 시도
	invalidPath := "/invalid/path/that/does/not/exist/file.txt"
	err := ctx.SaveUploadedFile(fileHeader, invalidPath)
	if err == nil {
		t.Error("Expected error when saving to invalid path, got nil")
	}
}

// TestMultipartFormWithMaxUploadSize tests custom max upload size
// 커스텀 최대 업로드 크기 테스트
func TestMultipartFormWithMaxUploadSize(t *testing.T) {
	// Create app with custom max upload size
	// 커스텀 최대 업로드 크기로 앱 생성
	app := New(WithMaxUploadSize(1 << 20)) // 1 MB limit

	// Create small file (within limit)
	// 작은 파일 생성 (제한 내)
	content := make([]byte, 1024) // 1 KB

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fileWriter, _ := writer.CreateFormFile("file", "small.txt")
	fileWriter.Write(content)
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	ctx := NewContext(w, req)
	// Set app reference
	// 앱 참조 설정
	ctx.app = app

	// Should succeed with custom size limit
	// 커스텀 크기 제한으로 성공해야 함
	form, err := ctx.MultipartForm()
	if err != nil {
		t.Errorf("Expected success with custom max upload size, got error: %v", err)
	}

	if len(form.File["file"]) != 1 {
		t.Errorf("Expected 1 file, got %d", len(form.File["file"]))
	}
}

// BenchmarkFormFile benchmarks the FormFile method
// FormFile 메서드 벤치마크
func BenchmarkFormFile(b *testing.B) {
	// Prepare multipart form
	// multipart 폼 준비
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fileWriter, _ := writer.CreateFormFile("upload", "test.txt")
	fileWriter.Write([]byte("benchmark content"))
	writer.Close()

	bodyBytes := body.Bytes()
	contentType := writer.FormDataContentType()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodPost, "/upload", bytes.NewReader(bodyBytes))
		req.Header.Set("Content-Type", contentType)
		w := httptest.NewRecorder()
		ctx := NewContext(w, req)
		ctx.FormFile("upload")
	}
}

// BenchmarkSaveUploadedFile benchmarks the SaveUploadedFile method
// SaveUploadedFile 메서드 벤치마크
func BenchmarkSaveUploadedFile(b *testing.B) {
	tmpDir := b.TempDir()

	// Prepare multipart form
	// multipart 폼 준비
	fileContent := make([]byte, 1024) // 1 KB
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fileWriter, _ := writer.CreateFormFile("upload", "test.txt")
	fileWriter.Write(fileContent)
	writer.Close()

	bodyBytes := body.Bytes()
	contentType := writer.FormDataContentType()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodPost, "/upload", bytes.NewReader(bodyBytes))
		req.Header.Set("Content-Type", contentType)
		w := httptest.NewRecorder()
		ctx := NewContext(w, req)

		fileHeader, _ := ctx.FormFile("upload")
		dstPath := filepath.Join(tmpDir, "bench.txt")
		ctx.SaveUploadedFile(fileHeader, dstPath)
		// Clean up
		// 정리
		os.Remove(dstPath)
	}
}
