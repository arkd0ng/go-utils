package fileutil

import (
	"fmt"
	"os"
	"path/filepath"
)

// DeleteFile deletes a file
// DeleteFile은 파일을 삭제합니다
//
// Example
// 예제:
//
//	err := fileutil.DeleteFile("path/to/file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
func DeleteFile(path string) error {
	// Check if path is a file
	// 경로가 파일인지 확인
	if !IsFile(path) {
		return fmt.Errorf("fileutil.DeleteFile: %w", ErrNotFile)
	}

	if err := os.Remove(path); err != nil {
		return fmt.Errorf("fileutil.DeleteFile: %w", err)
	}

	return nil
}

// DeleteDir deletes an empty directory
// DeleteDir는 빈 디렉토리를 삭제합니다
//
// Example
// 예제:
//
//	err := fileutil.DeleteDir("path/to/empty-dir")
//	if err != nil {
//	    log.Fatal(err)
//	}
func DeleteDir(path string) error {
	// Check if path is a directory
	// 경로가 디렉토리인지 확인
	if !IsDir(path) {
		return fmt.Errorf("fileutil.DeleteDir: %w", ErrNotDirectory)
	}

	if err := os.Remove(path); err != nil {
		return fmt.Errorf("fileutil.DeleteDir: %w", err)
	}

	return nil
}

// DeleteRecursive deletes a file or directory recursively
// DeleteRecursive는 파일 또는 디렉토리를 재귀적으로 삭제합니다
//
// This is an alias for os.RemoveAll.
// 이는 os.RemoveAll의 별칭입니다.
//
// Example
// 예제:
//
//	err := fileutil.DeleteRecursive("path/to/directory")
//	if err != nil {
//	    log.Fatal(err)
//	}
func DeleteRecursive(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return fmt.Errorf("fileutil.DeleteRecursive: %w", err)
	}
	return nil
}

// DeletePattern deletes all files matching a glob pattern
// DeletePattern은 glob 패턴과 일치하는 모든 파일을 삭제합니다
//
// Example
// 예제:
//
// // Delete all .tmp files
// 모든 .tmp 파일 삭제
//
//	err := fileutil.DeletePattern("*.tmp")
//	if err != nil {
//	    log.Fatal(err)
//	}
func DeletePattern(pattern string) error {
	files, err := filepath.Glob(pattern)
	if err != nil {
		return fmt.Errorf("fileutil.DeletePattern: %w", err)
	}

	for _, file := range files {
		if err := os.Remove(file); err != nil {
			return fmt.Errorf("fileutil.DeletePattern: %w", err)
		}
	}

	return nil
}

// DeleteFiles deletes multiple files
// DeleteFiles는 여러 파일을 삭제합니다
//
// Example
// 예제:
//
//	err := fileutil.DeleteFiles("file1.txt", "file2.txt", "file3.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
func DeleteFiles(paths ...string) error {
	for _, path := range paths {
		if err := os.Remove(path); err != nil {
			return fmt.Errorf("fileutil.DeleteFiles: %w", err)
		}
	}
	return nil
}

// Clean removes all files and subdirectories in a directory but keeps the directory itself
// Clean은 디렉토리의 모든 파일과 하위 디렉토리를 제거하지만 디렉토리 자체는 유지합니다
//
// Example
// 예제:
//
//	err := fileutil.Clean("path/to/directory")
//	if err != nil {
//	    log.Fatal(err)
//	}
func Clean(path string) error {
	// Check if path is a directory
	// 경로가 디렉토리인지 확인
	if !IsDir(path) {
		return fmt.Errorf("fileutil.Clean: %w", ErrNotDirectory)
	}

	// Read directory entries
	// 디렉토리 항목 읽기
	entries, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("fileutil.Clean: %w", err)
	}

	// Delete each entry
	// 각 항목 삭제
	for _, entry := range entries {
		entryPath := filepath.Join(path, entry.Name())
		if err := os.RemoveAll(entryPath); err != nil {
			return fmt.Errorf("fileutil.Clean: %w", err)
		}
	}

	return nil
}

// RemoveEmpty removes all empty directories recursively under the given path
// RemoveEmpty는 주어진 경로 아래의 모든 빈 디렉토리를 재귀적으로 제거합니다
//
// Example
// 예제:
//
//	err := fileutil.RemoveEmpty("path/to/directory")
//	if err != nil {
//	    log.Fatal(err)
//	}
func RemoveEmpty(path string) error {
	// Check if path is a directory
	// 경로가 디렉토리인지 확인
	if !IsDir(path) {
		return fmt.Errorf("fileutil.RemoveEmpty: %w", ErrNotDirectory)
	}

	// Walk directory tree in reverse order (bottom-up)
	// 디렉토리 트리를 역순으로 순회 (상향식)
	var dirsToCheck []string
	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			dirsToCheck = append(dirsToCheck, p)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("fileutil.RemoveEmpty: %w", err)
	}

	// Check and remove empty directories from bottom to top
	// 하단에서 상단으로 빈 디렉토리 확인 및 제거
	for i := len(dirsToCheck) - 1; i >= 0; i-- {
		dir := dirsToCheck[i]
		isEmpty, err := IsEmpty(dir)
		if err != nil {
			return fmt.Errorf("fileutil.RemoveEmpty: %w", err)
		}
		// Don't remove the root path
		// 루트 경로는 제거하지 않음
		if isEmpty && dir != path {
			if err := os.Remove(dir); err != nil {
				return fmt.Errorf("fileutil.RemoveEmpty: %w", err)
			}
		}
	}

	return nil
}
