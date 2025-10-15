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
// With options / 옵션 포함:
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
// With options / 옵션 포함:
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
