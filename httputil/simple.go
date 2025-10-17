package httputil

import (
	"context"
)

// defaultClient is the default HTTP client used by package-level functions.
// defaultClient는 패키지 레벨 함수에서 사용되는 기본 HTTP 클라이언트입니다.
var defaultClient = NewClient()

// Get performs a GET request using the default client and decodes the JSON response into result.
// Get은 기본 클라이언트를 사용하여 GET 요청을 수행하고 JSON 응답을 result로 디코딩합니다.
//
// Example:
//
//	var result MyStruct
//	err := httputil.Get("https://api.example.com/data", &result)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// With options
// 옵션 포함:
//
//	err := httputil.Get("https://api.example.com/data", &result,
//	    httputil.WithBearerToken("your-token"),
//	    httputil.WithTimeout(30*time.Second),
//	)
func Get(url string, result interface{}, opts ...Option) error {
	return defaultClient.Get(url, result, opts...)
}

// GetContext performs a GET request with context using the default client.
// GetContext는 기본 클라이언트를 사용하여 context와 함께 GET 요청을 수행합니다.
func GetContext(ctx context.Context, url string, result interface{}, opts ...Option) error {
	return defaultClient.GetContext(ctx, url, result, opts...)
}

// Post performs a POST request with body using the default client and decodes the JSON response into result.
// Post는 기본 클라이언트를 사용하여 body와 함께 POST 요청을 수행하고 JSON 응답을 result로 디코딩합니다.
//
// Example:
//
//	payload := MyPayload{Name: "test"}
//	var response MyResponse
//	err := httputil.Post("https://api.example.com/create", payload, &response)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// With options
// 옵션 포함:
//
//	err := httputil.Post("https://api.example.com/create", payload, &response,
//	    httputil.WithBearerToken("your-token"),
//	    httputil.WithRetry(3),
//	)
func Post(url string, body, result interface{}, opts ...Option) error {
	return defaultClient.Post(url, body, result, opts...)
}

// PostContext performs a POST request with context using the default client.
// PostContext는 기본 클라이언트를 사용하여 context와 함께 POST 요청을 수행합니다.
func PostContext(ctx context.Context, url string, body, result interface{}, opts ...Option) error {
	return defaultClient.PostContext(ctx, url, body, result, opts...)
}

// Put performs a PUT request with body using the default client and decodes the JSON response into result.
// Put은 기본 클라이언트를 사용하여 body와 함께 PUT 요청을 수행하고 JSON 응답을 result로 디코딩합니다.
//
// Example:
//
//	payload := MyPayload{ID: 1, Name: "updated"}
//	var response MyResponse
//	err := httputil.Put("https://api.example.com/update/1", payload, &response)
//	if err != nil {
//	    log.Fatal(err)
//	}
func Put(url string, body, result interface{}, opts ...Option) error {
	return defaultClient.Put(url, body, result, opts...)
}

// PutContext performs a PUT request with context using the default client.
// PutContext는 기본 클라이언트를 사용하여 context와 함께 PUT 요청을 수행합니다.
func PutContext(ctx context.Context, url string, body, result interface{}, opts ...Option) error {
	return defaultClient.PutContext(ctx, url, body, result, opts...)
}

// Patch performs a PATCH request with body using the default client and decodes the JSON response into result.
// Patch는 기본 클라이언트를 사용하여 body와 함께 PATCH 요청을 수행하고 JSON 응답을 result로 디코딩합니다.
//
// Example:
//
//	payload := map[string]interface{}{"name": "patched"}
//	var response MyResponse
//	err := httputil.Patch("https://api.example.com/update/1", payload, &response)
//	if err != nil {
//	    log.Fatal(err)
//	}
func Patch(url string, body, result interface{}, opts ...Option) error {
	return defaultClient.Patch(url, body, result, opts...)
}

// PatchContext performs a PATCH request with context using the default client.
// PatchContext는 기본 클라이언트를 사용하여 context와 함께 PATCH 요청을 수행합니다.
func PatchContext(ctx context.Context, url string, body, result interface{}, opts ...Option) error {
	return defaultClient.PatchContext(ctx, url, body, result, opts...)
}

// Delete performs a DELETE request using the default client and decodes the JSON response into result.
// Delete는 기본 클라이언트를 사용하여 DELETE 요청을 수행하고 JSON 응답을 result로 디코딩합니다.
//
// Example:
//
//	var response MyResponse
//	err := httputil.Delete("https://api.example.com/delete/1", &response)
//	if err != nil {
//	    log.Fatal(err)
//	}
func Delete(url string, result interface{}, opts ...Option) error {
	return defaultClient.Delete(url, result, opts...)
}

// DeleteContext performs a DELETE request with context using the default client.
// DeleteContext는 기본 클라이언트를 사용하여 context와 함께 DELETE 요청을 수행합니다.
func DeleteContext(ctx context.Context, url string, result interface{}, opts ...Option) error {
	return defaultClient.DeleteContext(ctx, url, result, opts...)
}

// SetDefaultClient sets the default client used by package-level functions.
// SetDefaultClient는 패키지 레벨 함수에서 사용되는 기본 클라이언트를 설정합니다.
//
// This is useful if you want to configure the default client with custom options.
// 기본 클라이언트를 사용자 정의 옵션으로 구성하려는 경우 유용합니다.
//
// Example:
//
//	httputil.SetDefaultClient(httputil.NewClient(
//	    httputil.WithTimeout(60*time.Second),
//	    httputil.WithRetry(5),
//	))
func SetDefaultClient(client *Client) {
	defaultClient = client
}

// DownloadFile downloads a file using the default client.
// DownloadFile은 기본 클라이언트를 사용하여 파일을 다운로드합니다.
func DownloadFile(url, filepath string, opts ...Option) error {
	return defaultClient.DownloadFile(url, filepath, opts...)
}

// DownloadFileContext downloads a file with context using the default client.
// DownloadFileContext는 기본 클라이언트를 사용하여 context와 함께 파일을 다운로드합니다.
func DownloadFileContext(ctx context.Context, url, filepath string, progress ProgressFunc, opts ...Option) error {
	return defaultClient.DownloadFileContext(ctx, url, filepath, progress, opts...)
}

// Download downloads data using the default client.
// Download는 기본 클라이언트를 사용하여 데이터를 다운로드합니다.
func Download(url string, opts ...Option) ([]byte, error) {
	return defaultClient.Download(url, opts...)
}

// DownloadContext downloads data with context using the default client.
// DownloadContext는 기본 클라이언트를 사용하여 context와 함께 데이터를 다운로드합니다.
func DownloadContext(ctx context.Context, url string, opts ...Option) ([]byte, error) {
	return defaultClient.DownloadContext(ctx, url, opts...)
}

// UploadFile uploads a file using the default client.
// UploadFile은 기본 클라이언트를 사용하여 파일을 업로드합니다.
func UploadFile(url, fieldName, filepath string, result interface{}, opts ...Option) error {
	return defaultClient.UploadFile(url, fieldName, filepath, result, opts...)
}

// UploadFileContext uploads a file with context using the default client.
// UploadFileContext는 기본 클라이언트를 사용하여 context와 함께 파일을 업로드합니다.
func UploadFileContext(ctx context.Context, url, fieldName, filepath string, result interface{}, progress ProgressFunc, opts ...Option) error {
	return defaultClient.UploadFileContext(ctx, url, fieldName, filepath, result, progress, opts...)
}

// UploadFiles uploads multiple files using the default client.
// UploadFiles는 기본 클라이언트를 사용하여 여러 파일을 업로드합니다.
func UploadFiles(url string, files map[string]string, result interface{}, opts ...Option) error {
	return defaultClient.UploadFiles(url, files, result, opts...)
}

// UploadFilesContext uploads multiple files with context using the default client.
// UploadFilesContext는 기본 클라이언트를 사용하여 context와 함께 여러 파일을 업로드합니다.
func UploadFilesContext(ctx context.Context, url string, files map[string]string, result interface{}, opts ...Option) error {
	return defaultClient.UploadFilesContext(ctx, url, files, result, opts...)
}

// PostForm performs a POST request with form data using the default client.
// PostForm은 기본 클라이언트를 사용하여 폼 데이터와 함께 POST 요청을 수행합니다.
func PostForm(url string, data map[string]string, result interface{}, opts ...Option) error {
	return defaultClient.PostForm(url, data, result, opts...)
}

// PostFormContext performs a POST request with form data and context using the default client.
// PostFormContext는 기본 클라이언트를 사용하여 context 및 폼 데이터와 함께 POST 요청을 수행합니다.
func PostFormContext(ctx context.Context, url string, data map[string]string, result interface{}, opts ...Option) error {
	return defaultClient.PostFormContext(ctx, url, data, result, opts...)
}

// DoRaw performs an HTTP request and returns the raw response using the default client.
// DoRaw는 기본 클라이언트를 사용하여 HTTP 요청을 수행하고 원시 응답을 반환합니다.
func DoRaw(method, url string, body interface{}, opts ...Option) (*Response, error) {
	return defaultClient.DoRaw(method, url, body, opts...)
}

// DoRawContext performs an HTTP request with context and returns the raw response using the default client.
// DoRawContext는 기본 클라이언트를 사용하여 context와 함께 HTTP 요청을 수행하고 원시 응답을 반환합니다.
func DoRawContext(ctx context.Context, method, url string, body interface{}, opts ...Option) (*Response, error) {
	return defaultClient.DoRawContext(ctx, method, url, body, opts...)
}
