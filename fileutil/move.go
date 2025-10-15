package fileutil

import (
	"fmt"
	"os"
	"path/filepath"
)

// MoveFile moves a file from src to dst
// MoveFile은 src에서 dst로 파일을 이동합니다
//
// Example / 예제:
//
//	err := fileutil.MoveFile("old-path/file.txt", "new-path/file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
func MoveFile(src, dst string) error {
	// Check if source exists and is a file / 소스 존재 및 파일 확인
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("fileutil.MoveFile: %w", err)
	}
	if !srcInfo.Mode().IsRegular() {
		return fmt.Errorf("fileutil.MoveFile: source is not a regular file")
	}

	// Create destination directory / 대상 디렉토리 생성
	if err := os.MkdirAll(filepath.Dir(dst), DefaultDirMode); err != nil {
		return fmt.Errorf("fileutil.MoveFile: %w", err)
	}

	// Try to rename first (fast if on same filesystem) / 먼저 이름 변경 시도 (동일한 파일 시스템에서 빠름)
	if err := os.Rename(src, dst); err == nil {
		return nil
	}

	// If rename fails, copy and delete / 이름 변경 실패 시 복사 및 삭제
	if err := CopyFile(src, dst); err != nil {
		return fmt.Errorf("fileutil.MoveFile: %w", err)
	}

	if err := os.Remove(src); err != nil {
		return fmt.Errorf("fileutil.MoveFile: %w", err)
	}

	return nil
}

// MoveDir moves a directory from src to dst
// MoveDir는 src에서 dst로 디렉토리를 이동합니다
//
// Example / 예제:
//
//	err := fileutil.MoveDir("old-path/dir", "new-path/dir")
//	if err != nil {
//	    log.Fatal(err)
//	}
func MoveDir(src, dst string) error {
	// Check if source is a directory / 소스가 디렉토리인지 확인
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("fileutil.MoveDir: %w", err)
	}
	if !srcInfo.IsDir() {
		return fmt.Errorf("fileutil.MoveDir: %w", ErrNotDirectory)
	}

	// Try to rename first (fast if on same filesystem) / 먼저 이름 변경 시도 (동일한 파일 시스템에서 빠름)
	if err := os.Rename(src, dst); err == nil {
		return nil
	}

	// If rename fails, copy and delete / 이름 변경 실패 시 복사 및 삭제
	if err := CopyDir(src, dst); err != nil {
		return fmt.Errorf("fileutil.MoveDir: %w", err)
	}

	if err := os.RemoveAll(src); err != nil {
		return fmt.Errorf("fileutil.MoveDir: %w", err)
	}

	return nil
}

// Rename renames a file or directory
// Rename은 파일 또는 디렉토리의 이름을 변경합니다
//
// This is an alias for os.Rename.
// 이는 os.Rename의 별칭입니다.
//
// Example / 예제:
//
//	err := fileutil.Rename("old-name.txt", "new-name.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
func Rename(old, new string) error {
	// Create destination directory / 대상 디렉토리 생성
	if err := os.MkdirAll(filepath.Dir(new), DefaultDirMode); err != nil {
		return fmt.Errorf("fileutil.Rename: %w", err)
	}

	if err := os.Rename(old, new); err != nil {
		return fmt.Errorf("fileutil.Rename: %w", err)
	}
	return nil
}

// RenameExt renames a file by changing its extension
// RenameExt는 파일의 확장자를 변경하여 이름을 변경합니다
//
// Example / 예제:
//
//	err := fileutil.RenameExt("file.txt", ".md")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	// Result: "file.md"
func RenameExt(path, newExt string) error {
	newPath := ChangeExt(path, newExt)
	return Rename(path, newPath)
}

// SafeMove moves a file or directory safely by copying first and then deleting
// SafeMove는 먼저 복사한 다음 삭제하여 파일 또는 디렉토리를 안전하게 이동합니다
//
// This is slower than MoveFile/MoveDir but ensures data integrity.
// 이는 MoveFile/MoveDir보다 느리지만 데이터 무결성을 보장합니다.
//
// Example / 예제:
//
//	err := fileutil.SafeMove("source.txt", "destination.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
func SafeMove(src, dst string) error {
	// Check if source exists / 소스 존재 확인
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("fileutil.SafeMove: %w", err)
	}

	if srcInfo.IsDir() {
		// Copy directory / 디렉토리 복사
		if err := CopyDir(src, dst); err != nil {
			return fmt.Errorf("fileutil.SafeMove: %w", err)
		}

		// Verify copy succeeded by checking if dst exists / dst가 존재하는지 확인하여 복사 성공 검증
		if !Exists(dst) {
			return fmt.Errorf("fileutil.SafeMove: copy verification failed")
		}

		// Delete source / 소스 삭제
		if err := os.RemoveAll(src); err != nil {
			return fmt.Errorf("fileutil.SafeMove: %w", err)
		}
	} else {
		// Copy file / 파일 복사
		if err := CopyFile(src, dst); err != nil {
			return fmt.Errorf("fileutil.SafeMove: %w", err)
		}

		// Verify copy succeeded by checking if dst exists / dst가 존재하는지 확인하여 복사 성공 검증
		if !Exists(dst) {
			return fmt.Errorf("fileutil.SafeMove: copy verification failed")
		}

		// Delete source / 소스 삭제
		if err := os.Remove(src); err != nil {
			return fmt.Errorf("fileutil.SafeMove: %w", err)
		}
	}

	return nil
}
