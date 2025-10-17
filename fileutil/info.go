package fileutil

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

// Exists checks if a file or directory exists
// Exists는 파일 또는 디렉토리가 존재하는지 확인합니다
//
// Example
// 예제:
//
//	if fileutil.Exists("path/to/file.txt") {
//	    fmt.Println("File exists")
//	}
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// IsFile checks if the path is a regular file
// IsFile은 경로가 일반 파일인지 확인합니다
//
// Example
// 예제:
//
//	if fileutil.IsFile("path/to/file.txt") {
//	    fmt.Println("Is a file")
//	}
func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.Mode().IsRegular()
}

// IsDir checks if the path is a directory
// IsDir는 경로가 디렉토리인지 확인합니다
//
// Example
// 예제:
//
//	if fileutil.IsDir("path/to/directory") {
//	    fmt.Println("Is a directory")
//	}
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// IsSymlink checks if the path is a symbolic link
// IsSymlink은 경로가 심볼릭 링크인지 확인합니다
//
// Example
// 예제:
//
//	if fileutil.IsSymlink("path/to/link") {
//	    fmt.Println("Is a symbolic link")
//	}
func IsSymlink(path string) bool {
	info, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return info.Mode()&os.ModeSymlink != 0
}

// Size returns the size of a file in bytes
// Size는 파일의 크기를 바이트로 반환합니다
//
// Example
// 예제:
//
//	size, err := fileutil.Size("path/to/file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("File size: %d bytes\n", size)
func Size(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, fmt.Errorf("fileutil.Size: %w", err)
	}
	return info.Size(), nil
}

// SizeHuman returns the file size in a human-readable format (e.g., "1.5 MB")
// SizeHuman은 사람이 읽기 쉬운 형식으로 파일 크기를 반환합니다 (예: "1.5 MB")
//
// Example
// 예제:
//
//	size, err := fileutil.SizeHuman("path/to/file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("File size: %s\n", size)
func SizeHuman(path string) (string, error) {
	size, err := Size(path)
	if err != nil {
		return "", err
	}

	return formatSize(size), nil
}

// formatSize formats a size in bytes to a human-readable string
// formatSize는 바이트 크기를 사람이 읽기 쉬운 문자열로 포맷합니다
func formatSize(size int64) string {
	if size < KB {
		return fmt.Sprintf("%d B", size)
	} else if size < MB {
		return fmt.Sprintf("%.2f KB", float64(size)/float64(KB))
	} else if size < GB {
		return fmt.Sprintf("%.2f MB", float64(size)/float64(MB))
	} else if size < TB {
		return fmt.Sprintf("%.2f GB", float64(size)/float64(GB))
	}
	return fmt.Sprintf("%.2f TB", float64(size)/float64(TB))
}

// Chmod changes the mode of a file or directory
// Chmod는 파일 또는 디렉토리의 모드를 변경합니다
//
// Example
// 예제:
//
//	err := fileutil.Chmod("path/to/file.txt", 0644)
//	if err != nil {
//	    log.Fatal(err)
//	}
func Chmod(path string, mode os.FileMode) error {
	if err := os.Chmod(path, mode); err != nil {
		return fmt.Errorf("fileutil.Chmod: %w", err)
	}
	return nil
}

// Chown changes the owner and group of a file or directory
// Chown은 파일 또는 디렉토리의 소유자와 그룹을 변경합니다
//
// Example
// 예제:
//
//	err := fileutil.Chown("path/to/file.txt", 1000, 1000)
//	if err != nil {
//	    log.Fatal(err)
//	}
func Chown(path string, uid, gid int) error {
	if err := os.Chown(path, uid, gid); err != nil {
		return fmt.Errorf("fileutil.Chown: %w", err)
	}
	return nil
}

// IsReadable checks if a file or directory is readable
// IsReadable은 파일 또는 디렉토리가 읽기 가능한지 확인합니다
//
// Example
// 예제:
//
//	if fileutil.IsReadable("path/to/file.txt") {
//	    fmt.Println("File is readable")
//	}
func IsReadable(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.Mode()&0400 != 0 // Check read permission for owner
}

// IsWritable checks if a file or directory is writable
// IsWritable은 파일 또는 디렉토리가 쓰기 가능한지 확인합니다
//
// Example
// 예제:
//
//	if fileutil.IsWritable("path/to/file.txt") {
//	    fmt.Println("File is writable")
//	}
func IsWritable(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.Mode()&0200 != 0 // Check write permission for owner
}

// IsExecutable checks if a file is executable
// IsExecutable은 파일이 실행 가능한지 확인합니다
//
// Example
// 예제:
//
//	if fileutil.IsExecutable("path/to/script.sh") {
//	    fmt.Println("File is executable")
//	}
func IsExecutable(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.Mode()&0100 != 0 // Check execute permission for owner
}

// ModTime returns the modification time of a file or directory
// ModTime은 파일 또는 디렉토리의 수정 시간을 반환합니다
//
// Example
// 예제:
//
//	modTime, err := fileutil.ModTime("path/to/file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Last modified: %s\n", modTime)
func ModTime(path string) (time.Time, error) {
	info, err := os.Stat(path)
	if err != nil {
		return time.Time{}, fmt.Errorf("fileutil.ModTime: %w", err)
	}
	return info.ModTime(), nil
}

// AccessTime returns the access time of a file or directory
// AccessTime은 파일 또는 디렉토리의 액세스 시간을 반환합니다
//
// Note: This function is platform-specific and may not work on all systems.
// 참고: 이 함수는 플랫폼별로 다르며 모든 시스템에서 작동하지 않을 수 있습니다.
//
// Example
// 예제:
//
//	accessTime, err := fileutil.AccessTime("path/to/file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Last accessed: %s\n", accessTime)
func AccessTime(path string) (time.Time, error) {
	info, err := os.Stat(path)
	if err != nil {
		return time.Time{}, fmt.Errorf("fileutil.AccessTime: %w", err)
	}

	// Try to get access time from system-specific info
	// 시스템별 정보에서 액세스 시간 가져오기 시도
	if stat, ok := info.Sys().(*syscall.Stat_t); ok {
		return time.Unix(stat.Atimespec.Sec, stat.Atimespec.Nsec), nil
	}

	// Fallback to modification time if access time is not available
	// 액세스 시간을 사용할 수 없는 경우 수정 시간으로 대체
	return info.ModTime(), nil
}

// ChangeTime returns the change time of a file or directory
// ChangeTime은 파일 또는 디렉토리의 변경 시간을 반환합니다
//
// Note: This function is platform-specific and may not work on all systems.
// 참고: 이 함수는 플랫폼별로 다르며 모든 시스템에서 작동하지 않을 수 있습니다.
//
// Example
// 예제:
//
//	changeTime, err := fileutil.ChangeTime("path/to/file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Last changed: %s\n", changeTime)
func ChangeTime(path string) (time.Time, error) {
	info, err := os.Stat(path)
	if err != nil {
		return time.Time{}, fmt.Errorf("fileutil.ChangeTime: %w", err)
	}

	// Try to get change time from system-specific info
	// 시스템별 정보에서 변경 시간 가져오기 시도
	if stat, ok := info.Sys().(*syscall.Stat_t); ok {
		return time.Unix(stat.Ctimespec.Sec, stat.Ctimespec.Nsec), nil
	}

	// Fallback to modification time if change time is not available
	// 변경 시간을 사용할 수 없는 경우 수정 시간으로 대체
	return info.ModTime(), nil
}

// Touch creates a file if it doesn't exist, or updates its modification time if it does
// Touch는 파일이 존재하지 않으면 생성하고, 존재하면 수정 시간을 업데이트합니다
//
// Example
// 예제:
//
//	err := fileutil.Touch("path/to/file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
func Touch(path string) error {
	// Check if file exists
	// 파일 존재 확인
	if !Exists(path) {
		// Create file
		// 파일 생성
		return WriteFile(path, []byte{})
	}

	// Update modification time
	// 수정 시간 업데이트
	now := time.Now()
	if err := os.Chtimes(path, now, now); err != nil {
		return fmt.Errorf("fileutil.Touch: %w", err)
	}

	return nil
}

// FileInfo returns the FileInfo for a file or directory
// FileInfo는 파일 또는 디렉토리의 FileInfo를 반환합니다
//
// This is an alias for os.Stat.
// 이는 os.Stat의 별칭입니다.
//
// Example
// 예제:
//
//	info, err := fileutil.FileInfo("path/to/file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Name: %s, Size: %d\n", info.Name(), info.Size())
func FileInfo(path string) (os.FileInfo, error) {
	return Stat(path)
}

// Stat returns the FileInfo for a file or directory
// Stat은 파일 또는 디렉토리의 FileInfo를 반환합니다
//
// Example
// 예제:
//
//	info, err := fileutil.Stat("path/to/file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Name: %s, Size: %d\n", info.Name(), info.Size())
func Stat(path string) (os.FileInfo, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("fileutil.Stat: %w", err)
	}
	return info, nil
}

// Lstat returns the FileInfo for a file or directory (does not follow symbolic links)
// Lstat은 파일 또는 디렉토리의 FileInfo를 반환합니다 (심볼릭 링크를 따라가지 않음)
//
// Example
// 예제:
//
//	info, err := fileutil.Lstat("path/to/link")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Name: %s, Size: %d\n", info.Name(), info.Size())
func Lstat(path string) (os.FileInfo, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return nil, fmt.Errorf("fileutil.Lstat: %w", err)
	}
	return info, nil
}
