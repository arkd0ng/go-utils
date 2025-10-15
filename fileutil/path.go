package fileutil

import (
	"fmt"
	"path/filepath"
	"strings"
)

// Join joins path elements into a single path using the OS-specific separator
// Join은 OS별 구분자를 사용하여 경로 요소를 단일 경로로 결합합니다
//
// This is an alias for filepath.Join.
// 이는 filepath.Join의 별칭입니다.
//
// Example / 예제:
//
//	path := fileutil.Join("home", "user", "documents", "file.txt")
//	// Result: "home/user/documents/file.txt" (on Unix)
//	// Result: "home\\user\\documents\\file.txt" (on Windows)
func Join(elem ...string) string {
	return filepath.Join(elem...)
}

// Split splits a path into directory and file components
// Split은 경로를 디렉토리와 파일 구성 요소로 분할합니다
//
// This is an alias for filepath.Split.
// 이는 filepath.Split의 별칭입니다.
//
// Example / 예제:
//
//	dir, file := fileutil.Split("path/to/file.txt")
//	// dir = "path/to/", file = "file.txt"
func Split(path string) (dir, file string) {
	return filepath.Split(path)
}

// Base returns the last element of a path
// Base는 경로의 마지막 요소를 반환합니다
//
// This is an alias for filepath.Base.
// 이는 filepath.Base의 별칭입니다.
//
// Example / 예제:
//
//	base := fileutil.Base("path/to/file.txt")
//	// Result: "file.txt"
func Base(path string) string {
	return filepath.Base(path)
}

// Dir returns the directory part of a path
// Dir는 경로의 디렉토리 부분을 반환합니다
//
// This is an alias for filepath.Dir.
// 이는 filepath.Dir의 별칭입니다.
//
// Example / 예제:
//
//	dir := fileutil.Dir("path/to/file.txt")
//	// Result: "path/to"
func Dir(path string) string {
	return filepath.Dir(path)
}

// Ext returns the file extension of a path
// Ext는 경로의 파일 확장자를 반환합니다
//
// This is an alias for filepath.Ext.
// 이는 filepath.Ext의 별칭입니다.
//
// Example / 예제:
//
//	ext := fileutil.Ext("file.txt")
//	// Result: ".txt"
func Ext(path string) string {
	return filepath.Ext(path)
}

// Abs returns the absolute path of a file or directory
// Abs는 파일 또는 디렉토리의 절대 경로를 반환합니다
//
// Example / 예제:
//
//	abs, err := fileutil.Abs("relative/path/file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(abs) // "/current/dir/relative/path/file.txt"
func Abs(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("fileutil.Abs: %w", err)
	}
	return absPath, nil
}

// CleanPath cleans a path by simplifying it
// CleanPath는 경로를 단순화하여 정리합니다
//
// This is an alias for filepath.Clean.
// 이는 filepath.Clean의 별칭입니다.
//
// Example / 예제:
//
//	clean := fileutil.CleanPath("path/to/../file.txt")
//	// Result: "path/file.txt"
func CleanPath(path string) string {
	return filepath.Clean(path)
}

// Normalize normalizes a path by cleaning it and making it absolute if possible
// Normalize는 경로를 정리하고 가능하면 절대 경로로 만들어 정규화합니다
//
// Example / 예제:
//
//	normalized, err := fileutil.Normalize("./path/../file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(normalized) // "/current/dir/file.txt"
func Normalize(path string) (string, error) {
	// Clean the path first / 먼저 경로 정리
	cleanPath := CleanPath(path)

	// Make it absolute / 절대 경로로 만들기
	absPath, err := filepath.Abs(cleanPath)
	if err != nil {
		return cleanPath, fmt.Errorf("fileutil.Normalize: %w", err)
	}

	return absPath, nil
}

// ToSlash converts a path to use forward slashes
// ToSlash는 경로를 슬래시로 변환합니다
//
// This is an alias for filepath.ToSlash.
// 이는 filepath.ToSlash의 별칭입니다.
//
// Example / 예제:
//
//	slashPath := fileutil.ToSlash("path\\to\\file.txt")
//	// Result: "path/to/file.txt"
func ToSlash(path string) string {
	return filepath.ToSlash(path)
}

// FromSlash converts a path from forward slashes to OS-specific separators
// FromSlash는 경로를 슬래시에서 OS별 구분자로 변환합니다
//
// This is an alias for filepath.FromSlash.
// 이는 filepath.FromSlash의 별칭입니다.
//
// Example / 예제:
//
//	osPath := fileutil.FromSlash("path/to/file.txt")
//	// Result: "path/to/file.txt" (on Unix)
//	// Result: "path\\to\\file.txt" (on Windows)
func FromSlash(path string) string {
	return filepath.FromSlash(path)
}

// IsAbs checks if a path is absolute
// IsAbs는 경로가 절대 경로인지 확인합니다
//
// This is an alias for filepath.IsAbs.
// 이는 filepath.IsAbs의 별칭입니다.
//
// Example / 예제:
//
//	if fileutil.IsAbs("/absolute/path") {
//	    fmt.Println("Path is absolute")
//	}
func IsAbs(path string) bool {
	return filepath.IsAbs(path)
}

// IsValid checks if a path is valid (not empty and contains no invalid characters)
// IsValid는 경로가 유효한지 확인합니다 (비어 있지 않고 유효하지 않은 문자를 포함하지 않음)
//
// Example / 예제:
//
//	if fileutil.IsValid("path/to/file.txt") {
//	    fmt.Println("Path is valid")
//	}
func IsValid(path string) bool {
	if path == "" {
		return false
	}

	// Check for invalid characters / 유효하지 않은 문자 확인
	invalidChars := []string{"\x00"} // Null character
	for _, char := range invalidChars {
		if strings.Contains(path, char) {
			return false
		}
	}

	return true
}

// IsSafe checks if a path is safe (no path traversal attempts)
// IsSafe는 경로가 안전한지 확인합니다 (경로 순회 시도 없음)
//
// This function checks if the path contains ".." or absolute paths outside the root.
// 이 함수는 경로에 ".."가 포함되어 있거나 루트 외부의 절대 경로가 있는지 확인합니다.
//
// Example / 예제:
//
//	if fileutil.IsSafe("path/to/file.txt", "/root/dir") {
//	    fmt.Println("Path is safe")
//	}
func IsSafe(path, root string) bool {
	// Clean the path / 경로 정리
	cleanPath := CleanPath(path)

	// Check for path traversal attempts / 경로 순회 시도 확인
	if strings.Contains(cleanPath, "..") {
		return false
	}

	// If root is provided, check if path is within root / 루트가 제공되면 경로가 루트 내에 있는지 확인
	if root != "" {
		// Make both paths absolute / 두 경로 모두 절대 경로로 만들기
		absPath, err := filepath.Abs(cleanPath)
		if err != nil {
			return false
		}

		absRoot, err := filepath.Abs(root)
		if err != nil {
			return false
		}

		// Check if path is within root / 경로가 루트 내에 있는지 확인
		relPath, err := filepath.Rel(absRoot, absPath)
		if err != nil {
			return false
		}

		// If relative path starts with "..", it's outside root / 상대 경로가 ".."로 시작하면 루트 외부
		if strings.HasPrefix(relPath, "..") {
			return false
		}
	}

	return true
}

// Match checks if a filename matches a pattern
// Match는 파일 이름이 패턴과 일치하는지 확인합니다
//
// This is an alias for filepath.Match.
// 이는 filepath.Match의 별칭입니다.
//
// Example / 예제:
//
//	matched, err := fileutil.Match("*.txt", "file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	if matched {
//	    fmt.Println("File matches pattern")
//	}
func Match(pattern, name string) (bool, error) {
	matched, err := filepath.Match(pattern, name)
	if err != nil {
		return false, fmt.Errorf("fileutil.Match: %w", err)
	}
	return matched, nil
}

// Glob returns all files matching a pattern
// Glob는 패턴과 일치하는 모든 파일을 반환합니다
//
// This is an alias for filepath.Glob.
// 이는 filepath.Glob의 별칭입니다.
//
// Example / 예제:
//
//	files, err := fileutil.Glob("*.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, file := range files {
//	    fmt.Println(file)
//	}
func Glob(pattern string) ([]string, error) {
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, fmt.Errorf("fileutil.Glob: %w", err)
	}
	return files, nil
}

// Rel returns the relative path from basepath to targpath
// Rel은 basepath에서 targpath로의 상대 경로를 반환합니다
//
// Example / 예제:
//
//	relPath, err := fileutil.Rel("/base/path", "/base/path/to/file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(relPath) // "to/file.txt"
func Rel(basepath, targpath string) (string, error) {
	relPath, err := filepath.Rel(basepath, targpath)
	if err != nil {
		return "", fmt.Errorf("fileutil.Rel: %w", err)
	}
	return relPath, nil
}

// WithoutExt returns the path without the file extension
// WithoutExt는 파일 확장자를 제외한 경로를 반환합니다
//
// Example / 예제:
//
//	path := fileutil.WithoutExt("file.txt")
//	// Result: "file"
func WithoutExt(path string) string {
	ext := filepath.Ext(path)
	if ext == "" {
		return path
	}
	return path[:len(path)-len(ext)]
}

// ChangeExt changes the file extension
// ChangeExt는 파일 확장자를 변경합니다
//
// Example / 예제:
//
//	path := fileutil.ChangeExt("file.txt", ".md")
//	// Result: "file.md"
func ChangeExt(path, newExt string) string {
	base := WithoutExt(path)
	if !strings.HasPrefix(newExt, ".") {
		newExt = "." + newExt
	}
	return base + newExt
}

// HasExt checks if a path has one of the specified extensions
// HasExt는 경로가 지정된 확장자 중 하나를 가지고 있는지 확인합니다
//
// Example / 예제:
//
//	if fileutil.HasExt("file.txt", ".txt", ".md") {
//	    fmt.Println("File has a valid extension")
//	}
func HasExt(path string, exts ...string) bool {
	ext := filepath.Ext(path)
	for _, e := range exts {
		if !strings.HasPrefix(e, ".") {
			e = "." + e
		}
		if ext == e {
			return true
		}
	}
	return false
}
