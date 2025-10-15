// Package httputil provides extreme simplicity HTTP utilities for Go.
// 패키지 httputil은 Go를 위한 극도로 간단한 HTTP 유틸리티를 제공합니다.
//
// This package reduces 30+ lines of repetitive HTTP code to just 2-3 lines
// with automatic JSON handling, retry logic, and type-safe operations.
//
// 이 패키지는 30줄 이상의 반복적인 HTTP 코드를 자동 JSON 처리, 재시도 로직,
// 타입 안전 작업을 통해 단 2-3줄로 줄입니다.
//
// # Key Features / 주요 기능
//
//   - Simple HTTP methods (GET, POST, PUT, PATCH, DELETE) / 간단한 HTTP 메서드
//   - Automatic JSON encoding/decoding / 자동 JSON 인코딩/디코딩
//   - Automatic retry with exponential backoff / 지수 백오프를 사용한 자동 재시도
//   - Type-safe response parsing with generics / 제네릭을 사용한 타입 안전 응답 파싱
//   - Context support for cancellation and timeouts / 취소 및 타임아웃을 위한 Context 지원
//   - Rich error types with debugging information / 디버깅 정보를 포함한 풍부한 에러 타입
//   - Zero external dependencies / 외부 의존성 제로
//
// # Quick Start / 빠른 시작
//
// Simple GET request:
//
//	var result MyStruct
//	err := httputil.Get("https://api.example.com/data", &result)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// Simple POST request with options:
//
//	payload := MyPayload{Name: "test"}
//	var response MyResponse
//	err := httputil.Post("https://api.example.com/create", payload, &response,
//	    httputil.WithBearerToken("your-token"),
//	    httputil.WithTimeout(30*time.Second),
//	    httputil.WithRetry(3),
//	)
//
// Using a client for multiple requests:
//
//	client := httputil.NewClient(
//	    httputil.WithBaseURL("https://api.example.com"),
//	    httputil.WithBearerToken("your-token"),
//	    httputil.WithRetry(3),
//	)
//
//	var result MyStruct
//	err := client.Get("/data", &result)
//
// # Categories / 카테고리
//
//   - Simple HTTP Methods (10 functions): Get, Post, Put, Patch, Delete with Context variants
//   - Request Builders (8 functions): Query params, headers, auth, form data
//   - Response Helpers (10 functions): JSON parsing, status checking, body reading
//   - Client Configuration (12 functions): Timeout, retry, proxy, TLS config
//   - Download/Upload (6 functions): File download and upload with progress
//   - Utilities (8 functions): URL building, query params, content type helpers
//
// Total: ~54 functions across 6 categories
//
// # Design Philosophy / 설계 철학
//
// "30 lines → 2-3 lines" - Extreme Simplicity
//
//   - Auto everything: JSON handling, retries, error wrapping
//   - Type-safe with generics
//   - Zero configuration needed
//   - Context support everywhere
//
// # Version / 버전
//
// Current version: v1.10.001
//
// # License / 라이선스
//
// MIT License
package httputil

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Version is the current version of the httputil package.
// Version은 httputil 패키지의 현재 버전입니다.
//
// The version is automatically loaded from cfg/app.yaml at package initialization.
// If the version cannot be loaded, it defaults to "unknown".
//
// 버전은 패키지 초기화 시 cfg/app.yaml에서 자동으로 로드됩니다.
// 버전을 로드할 수 없는 경우 "unknown"으로 기본 설정됩니다.
var Version = getVersion()

// getVersion loads the version from cfg/app.yaml
// getVersion은 cfg/app.yaml에서 버전을 로드합니다
func getVersion() string {
	// Try to find cfg/app.yaml relative to the current working directory
	// 현재 작업 디렉토리를 기준으로 cfg/app.yaml을 찾습니다
	configPaths := []string{
		"cfg/app.yaml",
		"../cfg/app.yaml",
		"../../cfg/app.yaml",
	}

	for _, path := range configPaths {
		if data, err := os.ReadFile(path); err == nil {
			var config struct {
				App struct {
					Version string `yaml:"version"`
				} `yaml:"app"`
			}
			if err := yaml.Unmarshal(data, &config); err == nil && config.App.Version != "" {
				return config.App.Version
			}
		}
	}

	// If we can't find the config, try to find it relative to this source file
	// 설정을 찾을 수 없는 경우 이 소스 파일을 기준으로 찾습니다
	if execPath, err := os.Executable(); err == nil {
		configPath := filepath.Join(filepath.Dir(execPath), "..", "cfg", "app.yaml")
		if data, err := os.ReadFile(configPath); err == nil {
			var config struct {
				App struct {
					Version string `yaml:"version"`
				} `yaml:"app"`
			}
			if err := yaml.Unmarshal(data, &config); err == nil && config.App.Version != "" {
				return config.App.Version
			}
		}
	}

	return "unknown"
}

// Example demonstrates basic usage of the httputil package.
// Example은 httputil 패키지의 기본 사용법을 보여줍니다.
func Example() {
	// This is a placeholder for package-level examples
	// 이것은 패키지 레벨 예제를 위한 플레이스홀더입니다
	fmt.Println("httputil package v" + Version)
	fmt.Println("Extreme simplicity HTTP utilities for Go")
	fmt.Println("30 lines → 2-3 lines")
}
