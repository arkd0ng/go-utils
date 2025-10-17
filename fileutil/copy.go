package fileutil

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CopyFile copies a file from src to dst with optional settings
// CopyFile은 선택적 설정을 사용하여 src에서 dst로 파일을 복사합니다
//
// Example
// 예제:
//
// // Simple copy
// 간단한 복사
//	err := fileutil.CopyFile("source.txt", "destination.txt")
//
// // Copy with overwrite
// 덮어쓰기와 함께 복사
//	err = fileutil.CopyFile("source.txt", "destination.txt",
//	    fileutil.WithOverwrite(true))
//
// // Copy with progress callback
// 진행 상황 콜백과 함께 복사
//	err = fileutil.CopyFile("source.txt", "destination.txt",
//	    fileutil.WithProgress(func(written, total int64) {
//	        fmt.Printf("Progress: %d/%d (%.2f%%)\n",
//	            written, total, float64(written)/float64(total)*100)
//	    }))
func CopyFile(src, dst string, options ...CopyOption) error {
	opts := defaultCopyOptions()
	for _, option := range options {
		option(opts)
	}

	// Check if source exists and is a file
	// 소스 존재 및 파일 확인
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("fileutil.CopyFile: %w", err)
	}
	if !srcInfo.Mode().IsRegular() {
		return fmt.Errorf("fileutil.CopyFile: source is not a regular file")
	}

	// Check if destination exists
	// 대상 존재 확인
	if !opts.overwrite {
		if _, err := os.Stat(dst); err == nil {
			return fmt.Errorf("fileutil.CopyFile: %w", ErrAlreadyExists)
		}
	}

	// Create destination directory
	// 대상 디렉토리 생성
	if err := os.MkdirAll(filepath.Dir(dst), DefaultDirMode); err != nil {
		return fmt.Errorf("fileutil.CopyFile: %w", err)
	}

	// Open source file
	// 소스 파일 열기
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("fileutil.CopyFile: %w", err)
	}
	defer srcFile.Close()

	// Create destination file
	// 대상 파일 생성
	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("fileutil.CopyFile: %w", err)
	}
	defer dstFile.Close()

	// Copy file contents
	// 파일 내용 복사
	if opts.progress != nil {
		// Copy with progress
		// 진행 상황과 함께 복사
		totalSize := srcInfo.Size()
		written := int64(0)

		buf := make([]byte, DefaultBufferSize)
		for {
			nr, er := srcFile.Read(buf)
			if nr > 0 {
				nw, ew := dstFile.Write(buf[0:nr])
				if nw > 0 {
					written += int64(nw)
					opts.progress(written, totalSize)
				}
				if ew != nil {
					return fmt.Errorf("fileutil.CopyFile: %w", ew)
				}
				if nr != nw {
					return fmt.Errorf("fileutil.CopyFile: short write")
				}
			}
			if er != nil {
				if er != io.EOF {
					return fmt.Errorf("fileutil.CopyFile: %w", er)
				}
				break
			}
		}
	} else {
		// Simple copy
		// 간단한 복사
		if _, err := io.Copy(dstFile, srcFile); err != nil {
			return fmt.Errorf("fileutil.CopyFile: %w", err)
		}
	}

	// Preserve permissions
	// 권한 보존
	if opts.preservePermissions {
		if err := os.Chmod(dst, srcInfo.Mode()); err != nil {
			return fmt.Errorf("fileutil.CopyFile: %w", err)
		}
	}

	// Preserve timestamps
	// 타임스탬프 보존
	if opts.preserveTimestamps {
		if err := os.Chtimes(dst, srcInfo.ModTime(), srcInfo.ModTime()); err != nil {
			return fmt.Errorf("fileutil.CopyFile: %w", err)
		}
	}

	return nil
}

// CopyDir copies a directory from src to dst with optional settings
// CopyDir는 선택적 설정을 사용하여 src에서 dst로 디렉토리를 복사합니다
//
// Example
// 예제:
//
// // Copy directory
// 디렉토리 복사
//	err := fileutil.CopyDir("source-dir", "destination-dir")
//
// // Copy with filter (exclude hidden files)
// 필터와 함께 복사 (숨겨진 파일 제외)
//	err = fileutil.CopyDir("source-dir", "destination-dir",
//	    fileutil.WithFilter(func(path string, info os.FileInfo) bool {
//	        return !strings.HasPrefix(info.Name(), ".")
//	    }))
func CopyDir(src, dst string, options ...CopyOption) error {
	opts := defaultCopyOptions()
	for _, option := range options {
		option(opts)
	}

	// Check if source is a directory
	// 소스가 디렉토리인지 확인
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("fileutil.CopyDir: %w", err)
	}
	if !srcInfo.IsDir() {
		return fmt.Errorf("fileutil.CopyDir: %w", ErrNotDirectory)
	}

	// Create destination directory
	// 대상 디렉토리 생성
	if err := os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return fmt.Errorf("fileutil.CopyDir: %w", err)
	}

	// Read directory entries
	// 디렉토리 항목 읽기
	entries, err := os.ReadDir(src)
	if err != nil {
		return fmt.Errorf("fileutil.CopyDir: %w", err)
	}

	// Copy each entry
	// 각 항목 복사
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		// Get entry info for filter
		// 필터를 위한 항목 정보 가져오기
		entryInfo, err := entry.Info()
		if err != nil {
			return fmt.Errorf("fileutil.CopyDir: %w", err)
		}

		// Apply filter if provided
		// 필터가 제공된 경우 적용
		if opts.filter != nil && !opts.filter(srcPath, entryInfo) {
			continue
		}

		if entry.IsDir() {
			// Recursively copy subdirectory
			// 하위 디렉토리 재귀적으로 복사
			if err := CopyDir(srcPath, dstPath, options...); err != nil {
				return err
			}
		} else {
			// Copy file
			// 파일 복사
			if err := CopyFile(srcPath, dstPath, options...); err != nil {
				return err
			}
		}
	}

	return nil
}

// CopyRecursive is an alias for CopyDir
// CopyRecursive는 CopyDir의 별칭입니다
func CopyRecursive(src, dst string, options ...CopyOption) error {
	return CopyDir(src, dst, options...)
}

// SyncDirs synchronizes the contents of two directories
// SyncDirs는 두 디렉토리의 내용을 동기화합니다
//
// This function copies new and updated files from src to dst.
// Files in dst that don't exist in src are not removed.
// 이 함수는 src에서 dst로 새 파일과 업데이트된 파일을 복사합니다.
// src에 존재하지 않는 dst의 파일은 제거되지 않습니다.
//
// Example
// 예제:
//
//	err := fileutil.SyncDirs("source-dir", "destination-dir")
//	if err != nil {
//	    log.Fatal(err)
//	}
func SyncDirs(src, dst string) error {
	// Check if source is a directory
	// 소스가 디렉토리인지 확인
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("fileutil.SyncDirs: %w", err)
	}
	if !srcInfo.IsDir() {
		return fmt.Errorf("fileutil.SyncDirs: %w", ErrNotDirectory)
	}

	// Create destination directory if it doesn't exist
	// 대상 디렉토리가 존재하지 않으면 생성
	if err := os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return fmt.Errorf("fileutil.SyncDirs: %w", err)
	}

	// Walk through source directory
	// 소스 디렉토리 순회
	return filepath.Walk(src, func(srcPath string, srcInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Get relative path
		// 상대 경로 가져오기
		relPath, err := filepath.Rel(src, srcPath)
		if err != nil {
			return err
		}

		// Skip root directory
		// 루트 디렉토리 건너뛰기
		if relPath == "." {
			return nil
		}

		dstPath := filepath.Join(dst, relPath)

		if srcInfo.IsDir() {
			// Create directory if it doesn't exist
			// 디렉토리가 존재하지 않으면 생성
			if err := os.MkdirAll(dstPath, srcInfo.Mode()); err != nil {
				return err
			}
		} else {
			// Check if file needs to be copied
			// 파일을 복사해야 하는지 확인
			shouldCopy := true
			dstInfo, err := os.Stat(dstPath)
			if err == nil {
				// File exists, check if source is newer
				// 파일 존재, 소스가 더 새로운지 확인
				if !srcInfo.ModTime().After(dstInfo.ModTime()) {
					shouldCopy = false
				}
			}

			if shouldCopy {
				// Copy file
				// 파일 복사
				if err := CopyFile(srcPath, dstPath); err != nil {
					return err
				}
			}
		}

		return nil
	})
}
