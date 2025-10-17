package fileutil

import (
	"os"
)

// CopyOption is a functional option for file copy operations
// CopyOption은 파일 복사 작업을 위한 함수형 옵션입니다
type CopyOption func(*copyOptions)

// copyOptions holds the configuration for file copy operations
// copyOptions는 파일 복사 작업을 위한 설정을 보유합니다
type copyOptions struct {
	overwrite           bool
	preservePermissions bool
	preserveTimestamps  bool
	progress            func(written, total int64)
	filter              func(path string, info os.FileInfo) bool
}

// defaultCopyOptions returns the default copy options
// defaultCopyOptions는 기본 복사 옵션을 반환합니다
func defaultCopyOptions() *copyOptions {
	return &copyOptions{
		overwrite:           false,
		preservePermissions: true,
		preserveTimestamps:  false,
		progress:            nil,
		filter:              nil,
	}
}

// WithOverwrite sets whether to overwrite existing files
// WithOverwrite는 기존 파일을 덮어쓸지 여부를 설정합니다
//
// Example
// 예제:
//
//	err := fileutil.CopyFile(src, dst, fileutil.WithOverwrite(true))
func WithOverwrite(overwrite bool) CopyOption {
	return func(opts *copyOptions) {
		opts.overwrite = overwrite
	}
}

// WithPreservePermissions sets whether to preserve file permissions
// WithPreservePermissions는 파일 권한을 보존할지 여부를 설정합니다
//
// Example
// 예제:
//
//	err := fileutil.CopyFile(src, dst, fileutil.WithPreservePermissions(true))
func WithPreservePermissions(preserve bool) CopyOption {
	return func(opts *copyOptions) {
		opts.preservePermissions = preserve
	}
}

// WithPreserveTimestamps sets whether to preserve file timestamps
// WithPreserveTimestamps는 파일 타임스탬프를 보존할지 여부를 설정합니다
//
// Example
// 예제:
//
//	err := fileutil.CopyFile(src, dst, fileutil.WithPreserveTimestamps(true))
func WithPreserveTimestamps(preserve bool) CopyOption {
	return func(opts *copyOptions) {
		opts.preserveTimestamps = preserve
	}
}

// WithProgress sets a progress callback function
// WithProgress는 진행 상황 콜백 함수를 설정합니다
//
// Example
// 예제:
//
//	err := fileutil.CopyFile(src, dst, fileutil.WithProgress(func(written, total int64) {
//	    fmt.Printf("Progress: %d/%d bytes\n", written, total)
//	}))
func WithProgress(fn func(written, total int64)) CopyOption {
	return func(opts *copyOptions) {
		opts.progress = fn
	}
}

// WithFilter sets a filter function to include/exclude files during copy
// WithFilter는 복사 중 파일을 포함/제외하는 필터 함수를 설정합니다
//
// Example
// 예제:
//
//	err := fileutil.CopyDir(src, dst, fileutil.WithFilter(func(path string, info os.FileInfo) bool {
//	    return !strings.HasPrefix(info.Name(), ".")
//	}))
func WithFilter(fn func(path string, info os.FileInfo) bool) CopyOption {
	return func(opts *copyOptions) {
		opts.filter = fn
	}
}
