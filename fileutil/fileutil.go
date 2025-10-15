// Package fileutil provides extreme simplicity file and path utilities for Golang.
// Package fileutil은 Golang을 위한 극도로 간단한 파일 및 경로 유틸리티를 제공합니다.
//
// This package reduces 20+ lines of repetitive file manipulation code to just 1-2 lines.
// 이 패키지는 20줄 이상의 반복적인 파일 조작 코드를 단 1-2줄로 줄입니다.
//
// Features / 특징:
//   - Safe file operations with automatic directory creation / 자동 디렉토리 생성을 사용한 안전한 파일 작업
//   - Cross-platform compatibility / 크로스 플랫폼 호환성
//   - Comprehensive file/directory operations / 포괄적인 파일/디렉토리 작업
//   - Path manipulation and validation / 경로 조작 및 검증
//   - File hashing and checksums / 파일 해싱 및 체크섬
//   - File compression (zip, tar, gzip) / 파일 압축 (zip, tar, gzip)
//   - Temporary file management / 임시 파일 관리
//   - Zero external dependencies (except yaml) / 외부 의존성 없음 (yaml 제외)
//
// Example / 예제:
//
//	// Write file with automatic directory creation / 자동 디렉토리 생성과 함께 파일 쓰기
//	err := fileutil.WriteString("path/to/file.txt", "Hello, World!")
//
//	// Read file / 파일 읽기
//	content, err := fileutil.ReadString("path/to/file.txt")
//
//	// Copy file / 파일 복사
//	err = fileutil.CopyFile("source.txt", "destination.txt")
//
//	// List all .go files recursively / 모든 .go 파일 재귀적으로 나열
//	files, err := fileutil.FindFiles(".", func(path string, info os.FileInfo) bool {
//	    return filepath.Ext(path) == ".go"
//	})
//
// Version: v1.9.001
package fileutil

import (
	"os"
)

// Version is the current version of the fileutil package
// Version은 fileutil 패키지의 현재 버전입니다
const Version = "v1.9.002"

// Default file and directory permissions / 기본 파일 및 디렉토리 권한
const (
	// DefaultFileMode is the default file permission (rw-r--r--)
	// DefaultFileMode는 기본 파일 권한입니다 (rw-r--r--)
	DefaultFileMode os.FileMode = 0644

	// DefaultDirMode is the default directory permission (rwxr-xr-x)
	// DefaultDirMode는 기본 디렉토리 권한입니다 (rwxr-xr-x)
	DefaultDirMode os.FileMode = 0755

	// DefaultExecMode is the default executable file permission (rwxr-xr-x)
	// DefaultExecMode는 기본 실행 파일 권한입니다 (rwxr-xr-x)
	DefaultExecMode os.FileMode = 0755
)

// Buffer sizes for I/O operations / I/O 작업을 위한 버퍼 크기
const (
	// DefaultBufferSize is the default buffer size for reading/writing (32KB)
	// DefaultBufferSize는 읽기/쓰기를 위한 기본 버퍼 크기입니다 (32KB)
	DefaultBufferSize = 32 * 1024

	// DefaultChunkSize is the default chunk size for streaming operations (1MB)
	// DefaultChunkSize는 스트리밍 작업을 위한 기본 청크 크기입니다 (1MB)
	DefaultChunkSize = 1024 * 1024
)

// Size units for human-readable file sizes / 사람이 읽기 쉬운 파일 크기를 위한 크기 단위
const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
	TB = 1024 * GB
)
