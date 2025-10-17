package fileutil

import (
	"fmt"
	"os"
	"path/filepath"
)

// MkdirAll creates a directory and all necessary parents
// MkdirAll은 디렉토리와 모든 필요한 부모 디렉토리를 생성합니다
//
// Example
// 예제:
//
//	err := fileutil.MkdirAll("path/to/nested/directory")
//	if err != nil {
//	    log.Fatal(err)
//	}
func MkdirAll(path string, perm ...os.FileMode) error {
	mode := DefaultDirMode
	if len(perm) > 0 {
		mode = perm[0]
	}

	if err := os.MkdirAll(path, mode); err != nil {
		return fmt.Errorf("fileutil.MkdirAll: %w", err)
	}

	return nil
}

// CreateTemp creates a temporary file in the specified directory
// CreateTemp는 지정된 디렉토리에 임시 파일을 생성합니다
//
// Example
// 예제:
//
//	tempPath, err := fileutil.CreateTemp("", "temp-*.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println("Temp file:", tempPath)
func CreateTemp(dir, pattern string) (string, error) {
	file, err := os.CreateTemp(dir, pattern)
	if err != nil {
		return "", fmt.Errorf("fileutil.CreateTemp: %w", err)
	}
	defer file.Close()

	return file.Name(), nil
}

// CreateTempDir creates a temporary directory in the specified directory
// CreateTempDir는 지정된 디렉토리에 임시 디렉토리를 생성합니다
//
// Example
// 예제:
//
//	tempDir, err := fileutil.CreateTempDir("", "temp-dir-*")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println("Temp directory:", tempDir)
func CreateTempDir(dir, pattern string) (string, error) {
	tempDir, err := os.MkdirTemp(dir, pattern)
	if err != nil {
		return "", fmt.Errorf("fileutil.CreateTempDir: %w", err)
	}

	return tempDir, nil
}

// IsEmpty checks if a directory is empty
// IsEmpty는 디렉토리가 비어 있는지 확인합니다
//
// Example
// 예제:
//
//	empty, err := fileutil.IsEmpty("path/to/directory")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	if empty {
//	    fmt.Println("Directory is empty")
//	}
func IsEmpty(path string) (bool, error) {
	// Check if path is a directory
	// 경로가 디렉토리인지 확인
	if !IsDir(path) {
		return false, fmt.Errorf("fileutil.IsEmpty: %w", ErrNotDirectory)
	}

	// Read directory entries
	// 디렉토리 항목 읽기
	entries, err := os.ReadDir(path)
	if err != nil {
		return false, fmt.Errorf("fileutil.IsEmpty: %w", err)
	}

	return len(entries) == 0, nil
}

// DirSize calculates the total size of all files in a directory recursively
// DirSize는 디렉토리의 모든 파일 크기를 재귀적으로 계산합니다
//
// Example
// 예제:
//
//	size, err := fileutil.DirSize("path/to/directory")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Directory size: %d bytes\n", size)
func DirSize(path string) (int64, error) {
	var totalSize int64

	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			totalSize += info.Size()
		}
		return nil
	})

	if err != nil {
		return 0, fmt.Errorf("fileutil.DirSize: %w", err)
	}

	return totalSize, nil
}

// ListFiles lists all files in a directory
// ListFiles는 디렉토리의 모든 파일을 나열합니다
//
// If recursive is true, it lists files in all subdirectories.
// recursive가 true이면 모든 하위 디렉토리의 파일을 나열합니다.
//
// Example
// 예제:
//
// // List files in current directory only
// 현재 디렉토리의 파일만 나열
//	files, err := fileutil.ListFiles(".")
//
// // List files recursively
// 재귀적으로 파일 나열
//	files, err = fileutil.ListFiles(".", true)
func ListFiles(dir string, recursive ...bool) ([]string, error) {
	isRecursive := false
	if len(recursive) > 0 {
		isRecursive = recursive[0]
	}

	var files []string

	if isRecursive {
		// Recursive listing
		// 재귀적 나열
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				files = append(files, path)
			}
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("fileutil.ListFiles: %w", err)
		}
	} else {
		// Non-recursive listing
		// 비재귀적 나열
		entries, err := os.ReadDir(dir)
		if err != nil {
			return nil, fmt.Errorf("fileutil.ListFiles: %w", err)
		}

		for _, entry := range entries {
			if !entry.IsDir() {
				files = append(files, filepath.Join(dir, entry.Name()))
			}
		}
	}

	return files, nil
}

// ListDirs lists all subdirectories in a directory
// ListDirs는 디렉토리의 모든 하위 디렉토리를 나열합니다
//
// If recursive is true, it lists all subdirectories recursively.
// recursive가 true이면 모든 하위 디렉토리를 재귀적으로 나열합니다.
//
// Example
// 예제:
//
// // List subdirectories in current directory only
// 현재 디렉토리의 하위 디렉토리만 나열
//	dirs, err := fileutil.ListDirs(".")
//
// // List subdirectories recursively
// 재귀적으로 하위 디렉토리 나열
//	dirs, err = fileutil.ListDirs(".", true)
func ListDirs(dir string, recursive ...bool) ([]string, error) {
	isRecursive := false
	if len(recursive) > 0 {
		isRecursive = recursive[0]
	}

	var dirs []string

	if isRecursive {
		// Recursive listing
		// 재귀적 나열
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() && path != dir {
				dirs = append(dirs, path)
			}
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("fileutil.ListDirs: %w", err)
		}
	} else {
		// Non-recursive listing
		// 비재귀적 나열
		entries, err := os.ReadDir(dir)
		if err != nil {
			return nil, fmt.Errorf("fileutil.ListDirs: %w", err)
		}

		for _, entry := range entries {
			if entry.IsDir() {
				dirs = append(dirs, filepath.Join(dir, entry.Name()))
			}
		}
	}

	return dirs, nil
}

// ListAll lists all files and directories
// ListAll은 모든 파일과 디렉토리를 나열합니다
//
// If recursive is true, it lists all entries recursively.
// recursive가 true이면 모든 항목을 재귀적으로 나열합니다.
//
// Example
// 예제:
//
// // List all entries in current directory only
// 현재 디렉토리의 모든 항목만 나열
//	entries, err := fileutil.ListAll(".")
//
// // List all entries recursively
// 재귀적으로 모든 항목 나열
//	entries, err = fileutil.ListAll(".", true)
func ListAll(dir string, recursive ...bool) ([]string, error) {
	isRecursive := false
	if len(recursive) > 0 {
		isRecursive = recursive[0]
	}

	var entries []string

	if isRecursive {
		// Recursive listing
		// 재귀적 나열
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if path != dir {
				entries = append(entries, path)
			}
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("fileutil.ListAll: %w", err)
		}
	} else {
		// Non-recursive listing
		// 비재귀적 나열
		dirEntries, err := os.ReadDir(dir)
		if err != nil {
			return nil, fmt.Errorf("fileutil.ListAll: %w", err)
		}

		for _, entry := range dirEntries {
			entries = append(entries, filepath.Join(dir, entry.Name()))
		}
	}

	return entries, nil
}

// Walk walks the file tree rooted at root, calling fn for each file or directory
// Walk는 root를 루트로 하는 파일 트리를 순회하며 각 파일 또는 디렉토리에 대해 fn을 호출합니다
//
// This is an alias for filepath.Walk.
// 이는 filepath.Walk의 별칭입니다.
//
// Example
// 예제:
//
//	err := fileutil.Walk(".", func(path string, info os.FileInfo, err error) error {
//	    if err != nil {
//	        return err
//	    }
//	    fmt.Println(path)
//	    return nil
//	})
func Walk(root string, fn filepath.WalkFunc) error {
	if err := filepath.Walk(root, fn); err != nil {
		return fmt.Errorf("fileutil.Walk: %w", err)
	}
	return nil
}

// WalkFiles walks the file tree and calls fn only for files
// WalkFiles는 파일 트리를 순회하며 파일에 대해서만 fn을 호출합니다
//
// Example
// 예제:
//
//	err := fileutil.WalkFiles(".", func(path string, info os.FileInfo) error {
//	    fmt.Println("File:", path)
//	    return nil
//	})
func WalkFiles(root string, fn func(path string, info os.FileInfo) error) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return fn(path, info)
		}
		return nil
	})
}

// WalkDirs walks the file tree and calls fn only for directories
// WalkDirs는 파일 트리를 순회하며 디렉토리에 대해서만 fn을 호출합니다
//
// Example
// 예제:
//
//	err := fileutil.WalkDirs(".", func(path string, info os.FileInfo) error {
//	    fmt.Println("Directory:", path)
//	    return nil
//	})
func WalkDirs(root string, fn func(path string, info os.FileInfo) error) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return fn(path, info)
		}
		return nil
	})
}

// FindFiles finds all files matching a predicate function
// FindFiles는 조건자 함수와 일치하는 모든 파일을 찾습니다
//
// Example
// 예제:
//
// // Find all .go files
// 모든 .go 파일 찾기
//	files, err := fileutil.FindFiles(".", func(path string, info os.FileInfo) bool {
//	    return filepath.Ext(path) == ".go"
//	})
func FindFiles(root string, predicate func(path string, info os.FileInfo) bool) ([]string, error) {
	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && predicate(path, info) {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("fileutil.FindFiles: %w", err)
	}

	return files, nil
}

// FilterFiles filters a slice of file paths using a predicate function
// FilterFiles는 조건자 함수를 사용하여 파일 경로 슬라이스를 필터링합니다
//
// Example
// 예제:
//
// // Filter large files (>1MB)
// 큰 파일 필터링 (>1MB)
//	large, err := fileutil.FilterFiles(files, func(path string) bool {
//	    info, _ := os.Stat(path)
//	    return info.Size() > 1024*1024
//	})
func FilterFiles(files []string, predicate func(path string) bool) ([]string, error) {
	var filtered []string

	for _, file := range files {
		if predicate(file) {
			filtered = append(filtered, file)
		}
	}

	return filtered, nil
}
