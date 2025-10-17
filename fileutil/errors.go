package fileutil

import (
	"errors"
	"os"
)

// Common errors
// 일반적인 에러
var (
	// ErrNotFound indicates that the file or directory was not found
	// ErrNotFound는 파일 또는 디렉토리를 찾을 수 없음을 나타냅니다
	ErrNotFound = errors.New("fileutil: file or directory not found")

	// ErrNotFile indicates that the path is not a regular file
	// ErrNotFile은 경로가 일반 파일이 아님을 나타냅니다
	ErrNotFile = errors.New("fileutil: not a regular file")

	// ErrNotDirectory indicates that the path is not a directory
	// ErrNotDirectory는 경로가 디렉토리가 아님을 나타냅니다
	ErrNotDirectory = errors.New("fileutil: not a directory")

	// ErrPermission indicates a permission denied error
	// ErrPermission은 권한 거부 에러를 나타냅니다
	ErrPermission = errors.New("fileutil: permission denied")

	// ErrInvalidPath indicates that the path is invalid or unsafe
	// ErrInvalidPath는 경로가 유효하지 않거나 안전하지 않음을 나타냅니다
	ErrInvalidPath = errors.New("fileutil: invalid or unsafe path")

	// ErrAlreadyExists indicates that the file or directory already exists
	// ErrAlreadyExists는 파일 또는 디렉토리가 이미 존재함을 나타냅니다
	ErrAlreadyExists = errors.New("fileutil: file or directory already exists")

	// ErrIsDirectory indicates that the operation cannot be performed on a directory
	// ErrIsDirectory는 디렉토리에서 작업을 수행할 수 없음을 나타냅니다
	ErrIsDirectory = errors.New("fileutil: is a directory")

	// ErrNotEmpty indicates that the directory is not empty
	// ErrNotEmpty는 디렉토리가 비어 있지 않음을 나타냅니다
	ErrNotEmpty = errors.New("fileutil: directory not empty")

	// ErrInvalidFormat indicates an invalid file format
	// ErrInvalidFormat은 유효하지 않은 파일 형식을 나타냅니다
	ErrInvalidFormat = errors.New("fileutil: invalid file format")

	// ErrInvalidChecksum indicates that the checksum verification failed
	// ErrInvalidChecksum은 체크섬 검증 실패를 나타냅니다
	ErrInvalidChecksum = errors.New("fileutil: invalid checksum")
)

// IsNotFound checks if the error is a "file not found" error / IsNotFound는 에러가 "파일을 찾을 수 없음" 에러인지 확인합니다
//
// Example
// 예제:
//
// if fileutil.IsNotFound(err) {
// Handle file not found / 파일을 찾을 수 없음 처리
//	}
func IsNotFound(err error) bool {
	return errors.Is(err, os.ErrNotExist) || errors.Is(err, ErrNotFound)
}

// IsPermission checks if the error is a "permission denied" error / IsPermission은 에러가 "권한 거부됨" 에러인지 확인합니다
//
// Example
// 예제:
//
// if fileutil.IsPermission(err) {
// Handle permission denied / 권한 거부됨 처리
//	}
func IsPermission(err error) bool {
	return errors.Is(err, os.ErrPermission) || errors.Is(err, ErrPermission)
}

// IsExist checks if the error is an "already exists" error / IsExist는 에러가 "이미 존재함" 에러인지 확인합니다
//
// Example
// 예제:
//
// if fileutil.IsExist(err) {
// Handle already exists / 이미 존재함 처리
//	}
func IsExist(err error) bool {
	return errors.Is(err, os.ErrExist) || errors.Is(err, ErrAlreadyExists)
}

// IsInvalid checks if the error is an "invalid path" error / IsInvalid는 에러가 "유효하지 않은 경로" 에러인지 확인합니다
//
// Example
// 예제:
//
// if fileutil.IsInvalid(err) {
// Handle invalid path / 유효하지 않은 경로 처리
//	}
func IsInvalid(err error) bool {
	return errors.Is(err, ErrInvalidPath)
}
