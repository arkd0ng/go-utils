package httputil

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// ProgressFunc is a callback function for tracking progress.
// ProgressFunc는 진행 상황을 추적하기 위한 콜백 함수입니다.
type ProgressFunc func(bytesRead, totalBytes int64)

// DownloadFile downloads a file from the given URL and saves it to the specified path.
// DownloadFile은 주어진 URL에서 파일을 다운로드하고 지정된 경로에 저장합니다.
func (c *Client) DownloadFile(url, filePath string, opts ...Option) error {
	return c.DownloadFileContext(context.Background(), url, filePath, nil, opts...)
}

// DownloadFileContext downloads a file with context and optional progress callback.
// DownloadFileContext는 context 및 선택적 진행 상황 콜백과 함께 파일을 다운로드합니다.
func (c *Client) DownloadFileContext(ctx context.Context, url, filePath string, progress ProgressFunc, opts ...Option) error {
	// Merge client config with request-specific options
	// 클라이언트 설정을 요청별 옵션과 병합
	cfg := *c.config
	cfg.apply(opts)

	// Create request / 요청 생성
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers / 헤더 설정
	req.Header.Set("User-Agent", cfg.userAgent)
	for k, v := range cfg.headers {
		req.Header.Set(k, v)
	}

	// Set authentication / 인증 설정
	if cfg.bearerToken != "" {
		req.Header.Set("Authorization", "Bearer "+cfg.bearerToken)
	}
	if cfg.basicAuthUser != "" {
		req.SetBasicAuth(cfg.basicAuthUser, cfg.basicAuthPass)
	}

	// Execute request / 요청 실행
	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to download file: %w", err)
	}
	defer resp.Body.Close()

	// Check status code / 상태 코드 확인
	if resp.StatusCode != http.StatusOK {
		return &HTTPError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
			URL:        url,
			Method:     http.MethodGet,
		}
	}

	// Create file / 파일 생성
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	// Copy with progress / 진행 상황과 함께 복사
	if progress != nil {
		totalBytes := resp.ContentLength
		reader := &progressReader{
			reader:   resp.Body,
			progress: progress,
			total:    totalBytes,
		}
		if _, err := io.Copy(out, reader); err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}
	} else {
		if _, err := io.Copy(out, resp.Body); err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}
	}

	return nil
}

// Download downloads data from the given URL and returns it as bytes.
// Download는 주어진 URL에서 데이터를 다운로드하고 바이트로 반환합니다.
func (c *Client) Download(url string, opts ...Option) ([]byte, error) {
	return c.DownloadContext(context.Background(), url, opts...)
}

// DownloadContext downloads data with context.
// DownloadContext는 context와 함께 데이터를 다운로드합니다.
func (c *Client) DownloadContext(ctx context.Context, url string, opts ...Option) ([]byte, error) {
	resp, err := c.DoRawContext(ctx, http.MethodGet, url, nil, opts...)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, &HTTPError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
			Body:       resp.String(),
			URL:        url,
			Method:     http.MethodGet,
		}
	}

	return resp.Body(), nil
}

// UploadFile uploads a file to the given URL using multipart/form-data.
// UploadFile은 multipart/form-data를 사용하여 주어진 URL에 파일을 업로드합니다.
func (c *Client) UploadFile(url, fieldName, filePath string, result interface{}, opts ...Option) error {
	return c.UploadFileContext(context.Background(), url, fieldName, filePath, result, nil, opts...)
}

// UploadFileContext uploads a file with context and optional progress callback.
// UploadFileContext는 context 및 선택적 진행 상황 콜백과 함께 파일을 업로드합니다.
func (c *Client) UploadFileContext(ctx context.Context, url, fieldName, filePath string, result interface{}, progress ProgressFunc, opts ...Option) error {
	// Open file / 파일 열기
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Get file info / 파일 정보 가져오기
	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	// Create multipart writer / multipart 작성기 생성
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create form file / 폼 파일 생성
	part, err := writer.CreateFormFile(fieldName, fileInfo.Name())
	if err != nil {
		return fmt.Errorf("failed to create form file: %w", err)
	}

	// Copy file content / 파일 내용 복사
	if progress != nil {
		reader := &progressReader{
			reader:   file,
			progress: progress,
			total:    fileInfo.Size(),
		}
		if _, err := io.Copy(part, reader); err != nil {
			return fmt.Errorf("failed to copy file: %w", err)
		}
	} else {
		if _, err := io.Copy(part, file); err != nil {
			return fmt.Errorf("failed to copy file: %w", err)
		}
	}

	// Close writer / 작성기 닫기
	if err := writer.Close(); err != nil {
		return fmt.Errorf("failed to close writer: %w", err)
	}

	// Merge client config / 클라이언트 설정 병합
	cfg := *c.config
	cfg.apply(opts)

	// Create request / 요청 생성
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers / 헤더 설정
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("User-Agent", cfg.userAgent)
	for k, v := range cfg.headers {
		req.Header.Set(k, v)
	}

	// Set authentication / 인증 설정
	if cfg.bearerToken != "" {
		req.Header.Set("Authorization", "Bearer "+cfg.bearerToken)
	}
	if cfg.basicAuthUser != "" {
		req.SetBasicAuth(cfg.basicAuthUser, cfg.basicAuthPass)
	}

	// Execute request / 요청 실행
	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}
	defer resp.Body.Close()

	// Check status code / 상태 코드 확인
	if resp.StatusCode >= 400 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return &HTTPError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
			Body:       string(bodyBytes),
			URL:        url,
			Method:     http.MethodPost,
		}
	}

	// Decode response / 응답 디코딩
	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

// UploadFiles uploads multiple files to the given URL using multipart/form-data.
// UploadFiles는 multipart/form-data를 사용하여 주어진 URL에 여러 파일을 업로드합니다.
func (c *Client) UploadFiles(url string, files map[string]string, result interface{}, opts ...Option) error {
	return c.UploadFilesContext(context.Background(), url, files, result, opts...)
}

// UploadFilesContext uploads multiple files with context.
// UploadFilesContext는 context와 함께 여러 파일을 업로드합니다.
func (c *Client) UploadFilesContext(ctx context.Context, url string, files map[string]string, result interface{}, opts ...Option) error {
	// Create multipart writer / multipart 작성기 생성
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add each file / 각 파일 추가
	for fieldName, filepath := range files {
		// Open file / 파일 열기
		file, err := os.Open(filepath)
		if err != nil {
			return fmt.Errorf("failed to open file %s: %w", filepath, err)
		}

		// Get file info / 파일 정보 가져오기
		fileInfo, err := file.Stat()
		if err != nil {
			file.Close()
			return fmt.Errorf("failed to get file info %s: %w", filepath, err)
		}

		// Create form file / 폼 파일 생성
		part, err := writer.CreateFormFile(fieldName, fileInfo.Name())
		if err != nil {
			file.Close()
			return fmt.Errorf("failed to create form file: %w", err)
		}

		// Copy file content / 파일 내용 복사
		if _, err := io.Copy(part, file); err != nil {
			file.Close()
			return fmt.Errorf("failed to copy file: %w", err)
		}

		file.Close()
	}

	// Close writer / 작성기 닫기
	if err := writer.Close(); err != nil {
		return fmt.Errorf("failed to close writer: %w", err)
	}

	// Merge client config / 클라이언트 설정 병합
	cfg := *c.config
	cfg.apply(opts)

	// Create request / 요청 생성
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers / 헤더 설정
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("User-Agent", cfg.userAgent)
	for k, v := range cfg.headers {
		req.Header.Set(k, v)
	}

	// Set authentication / 인증 설정
	if cfg.bearerToken != "" {
		req.Header.Set("Authorization", "Bearer "+cfg.bearerToken)
	}
	if cfg.basicAuthUser != "" {
		req.SetBasicAuth(cfg.basicAuthUser, cfg.basicAuthPass)
	}

	// Execute request / 요청 실행
	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to upload files: %w", err)
	}
	defer resp.Body.Close()

	// Check status code / 상태 코드 확인
	if resp.StatusCode >= 400 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return &HTTPError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
			Body:       string(bodyBytes),
			URL:        url,
			Method:     http.MethodPost,
		}
	}

	// Decode response / 응답 디코딩
	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

// progressReader wraps an io.Reader and calls a progress callback.
// progressReader는 io.Reader를 래핑하고 진행 상황 콜백을 호출합니다.
type progressReader struct {
	reader   io.Reader
	progress ProgressFunc
	total    int64
	read     int64
}

// Read implements io.Reader interface / io.Reader 인터페이스 구현
func (pr *progressReader) Read(p []byte) (int, error) {
	n, err := pr.reader.Read(p)
	pr.read += int64(n)
	if pr.progress != nil {
		pr.progress(pr.read, pr.total)
	}
	return n, err
}
