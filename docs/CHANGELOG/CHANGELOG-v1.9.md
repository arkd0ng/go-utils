# CHANGELOG - v1.9.x

This file contains detailed change logs for the v1.9.x releases of go-utils, focusing on the fileutil package.

이 파일은 fileutil 패키지에 중점을 둔 go-utils의 v1.9.x 릴리스에 대한 상세한 변경 로그를 포함합니다.

---

## [v1.9.001] - 2025-10-15

### Added / 추가됨

#### Fileutil Package - Complete Implementation / Fileutil 패키지 - 완전한 구현

**Package Structure / 패키지 구조:**
- `fileutil/fileutil.go` - Package documentation and constants (version v1.9.001)
- `fileutil/errors.go` - Custom error types and helper functions
- `fileutil/options.go` - Functional options pattern for copy operations
- `fileutil/read.go` - File reading functions (8 functions)
- `fileutil/write.go` - File writing and appending functions (11 functions)
- `fileutil/info.go` - File information functions (15 functions)
- `fileutil/path.go` - Path manipulation functions (18 functions)
- `fileutil/copy.go` - File/directory copying functions (4 functions)
- `fileutil/move.go` - File/directory moving functions (5 functions)
- `fileutil/delete.go` - File/directory deletion functions (7 functions)
- `fileutil/dir.go` - Directory operation functions (13 functions)
- `fileutil/hash.go` - File hashing functions (10 functions)
- `fileutil/fileutil_test.go` - Comprehensive test suite (7 test suites, 2 benchmarks)
- `fileutil/README.md` - Package documentation (bilingual)

**File Reading (8 functions) / 파일 읽기 (8개 함수):**
1. `ReadFile(path string) ([]byte, error)` - Read entire file
2. `ReadString(path string) (string, error)` - Read file as string
3. `ReadLines(path string) ([]string, error)` - Read file as lines
4. `ReadJSON(path string, v interface{}) error` - Read and unmarshal JSON
5. `ReadYAML(path string, v interface{}) error` - Read and unmarshal YAML
6. `ReadCSV(path string) ([][]string, error)` - Read CSV file
7. `ReadBytes(path string, offset, length int64) ([]byte, error)` - Read specific bytes
8. `ReadChunk(path string, chunkSize int64, fn func([]byte) error) error` - Read file in chunks

**File Writing (11 functions) / 파일 쓰기 (11개 함수):**
1. `WriteFile(path string, data []byte) error` - Write bytes to file
2. `WriteString(path string, s string) error` - Write string to file
3. `WriteLines(path string, lines []string) error` - Write lines to file
4. `WriteJSON(path string, v interface{}) error` - Marshal and write JSON
5. `WriteYAML(path string, v interface{}) error` - Marshal and write YAML
6. `WriteCSV(path string, records [][]string) error` - Write CSV file
7. `WriteAtomic(path string, data []byte) error` - Atomic write (temp + rename)
8. `AppendFile(path string, data []byte) error` - Append bytes to file
9. `AppendString(path string, s string) error` - Append string to file
10. `AppendLines(path string, lines []string) error` - Append lines to file
11. `AppendBytes(path string, data []byte) error` - Append bytes (alias)

**File Information (15 functions) / 파일 정보 (15개 함수):**
1. `Exists(path string) bool` - Check existence
2. `IsFile(path string) bool` - Check if file
3. `IsDir(path string) bool` - Check if directory
4. `IsSymlink(path string) bool` - Check if symlink
5. `Size(path string) (int64, error)` - Get file size
6. `SizeHuman(path string) (string, error)` - Get human-readable size
7. `Chmod(path string, mode os.FileMode) error` - Change permissions
8. `Chown(path string, uid, gid int) error` - Change owner
9. `IsReadable(path string) bool` - Check if readable
10. `IsWritable(path string) bool` - Check if writable
11. `IsExecutable(path string) bool` - Check if executable
12. `ModTime(path string) (time.Time, error)` - Get modification time
13. `AccessTime(path string) (time.Time, error)` - Get access time
14. `ChangeTime(path string) (time.Time, error)` - Get change time
15. `Touch(path string) error` - Update modification time

**Path Operations (18 functions) / 경로 작업 (18개 함수):**
1. `Join(elem ...string) string` - Join path elements
2. `Split(path string) (string, string)` - Split into dir and file
3. `Base(path string) string` - Get base name
4. `Dir(path string) string` - Get directory
5. `Ext(path string) string` - Get extension
6. `Abs(path string) (string, error)` - Get absolute path
7. `CleanPath(path string) string` - Clean path (renamed from Clean)
8. `Normalize(path string) (string, error)` - Normalize path
9. `ToSlash(path string) string` - Convert to forward slashes
10. `FromSlash(path string) string` - Convert to OS-specific
11. `IsAbs(path string) bool` - Check if absolute
12. `IsValid(path string) bool` - Validate path
13. `IsSafe(path, root string) bool` - Check path safety
14. `Match(pattern, name string) (bool, error)` - Match pattern
15. `Glob(pattern string) ([]string, error)` - Find by glob pattern
16. `Rel(basepath, targpath string) (string, error)` - Get relative path
17. `WithoutExt(path string) string` - Remove extension
18. `ChangeExt(path, newExt string) string` - Change extension
19. `HasExt(path string, exts ...string) bool` - Check extension

**File Copying (4 functions) / 파일 복사 (4개 함수):**
1. `CopyFile(src, dst string, opts ...CopyOption) error` - Copy single file
2. `CopyDir(src, dst string, opts ...CopyOption) error` - Copy directory
3. `CopyRecursive(src, dst string, opts ...CopyOption) error` - Copy recursively
4. `SyncDirs(src, dst string, opts ...CopyOption) error` - Sync two directories

**Copy Options / 복사 옵션:**
- `WithOverwrite(bool)` - Overwrite existing files
- `WithPreservePermissions(bool)` - Preserve file permissions
- `WithPreserveTimestamps(bool)` - Preserve timestamps
- `WithProgress(func(written, total int64))` - Progress callback
- `WithFilter(func(path string, info os.FileInfo) bool)` - File filter

**File Moving (5 functions) / 파일 이동 (5개 함수):**
1. `MoveFile(src, dst string) error` - Move file
2. `MoveDir(src, dst string) error` - Move directory
3. `Rename(oldPath, newPath string) error` - Rename file/directory
4. `RenameExt(path, newExt string) (string, error)` - Change file extension
5. `SafeMove(src, dst string) error` - Move with existence check

**File Deleting (7 functions) / 파일 삭제 (7개 함수):**
1. `DeleteFile(path string) error` - Delete single file
2. `DeleteDir(path string) error` - Delete empty directory
3. `DeleteRecursive(path string) error` - Delete recursively
4. `DeletePattern(dir, pattern string) error` - Delete files by pattern
5. `DeleteFiles(paths []string) error` - Delete multiple files
6. `Clean(path string) error` - Remove directory contents
7. `RemoveEmpty(path string) error` - Remove empty directories

**Directory Operations (13 functions) / 디렉토리 작업 (13개 함수):**
1. `MkdirAll(path string) error` - Create directory tree
2. `CreateTemp(dir, pattern string) (*os.File, error)` - Create temp file
3. `CreateTempDir(dir, pattern string) (string, error)` - Create temp directory
4. `IsEmpty(path string) (bool, error)` - Check if directory is empty
5. `DirSize(path string) (int64, error)` - Calculate directory size
6. `ListFiles(dir string) ([]string, error)` - List files only
7. `ListDirs(dir string) ([]string, error)` - List directories only
8. `ListAll(dir string) ([]string, error)` - List all entries
9. `Walk(root string, fn filepath.WalkFunc) error` - Walk directory tree
10. `WalkFiles(root string, fn func(string, os.FileInfo) error) error` - Walk files only
11. `WalkDirs(root string, fn func(string, os.FileInfo) error) error` - Walk directories only
12. `FindFiles(root string, filter func(string, interface{}) bool) ([]string, error)` - Find files by filter
13. `FilterFiles(root string, patterns []string) ([]string, error)` - Filter files by patterns

**File Hashing (10 functions) / 파일 해싱 (10개 함수):**
1. `MD5(path string) (string, error)` - Calculate MD5 hash
2. `SHA1(path string) (string, error)` - Calculate SHA1 hash
3. `SHA256(path string) (string, error)` - Calculate SHA256 hash
4. `SHA512(path string) (string, error)` - Calculate SHA512 hash
5. `Hash(path, algorithm string) (string, error)` - Calculate hash by algorithm
6. `HashBytes(data []byte, algorithm string) (string, error)` - Hash byte slice
7. `CompareFiles(path1, path2 string) (bool, error)` - Compare files byte-by-byte
8. `CompareHash(path1, path2 string) (bool, error)` - Compare files by hash
9. `Checksum(path string) (string, error)` - Calculate checksum (SHA256)
10. `VerifyChecksum(path, expected string) (bool, error)` - Verify checksum

**Custom Error Types / 사용자 정의 에러 타입:**
- `ErrNotFound` - File or directory not found
- `ErrNotFile` - Path is not a file
- `ErrNotDirectory` - Path is not a directory
- `ErrInvalidPath` - Invalid path
- `ErrPermissionDenied` - Permission denied
- `ErrAlreadyExists` - File or directory already exists
- `ErrNotEmpty` - Directory is not empty

**Error Helper Functions / 에러 헬퍼 함수:**
- `IsNotFound(err error) bool` - Check if error is "not found"
- `IsPermission(err error) bool` - Check if error is "permission denied"
- `IsExist(err error) bool` - Check if error is "already exists"
- `IsInvalid(err error) bool` - Check if error is "invalid path"

**Documentation / 문서:**
- Package README with quick start guide and function reference (bilingual)
- Comprehensive examples in `examples/fileutil/main.go` (7 example scenarios)
- Design plan document (`docs/fileutil/DESIGN_PLAN.md`)
- Work plan document (`docs/fileutil/WORK_PLAN.md`)

**Testing / 테스트:**
- 7 comprehensive test suites covering all function categories
- 2 benchmark tests for read/write operations
- 100% pass rate on all tests
- Test coverage across all major functionality

**Examples / 예제:**
- File Writing and Reading
- Path Operations
- File Information
- File Copying (with progress callback)
- File Hashing
- Directory Operations
- File Deletion

### Key Features / 주요 기능

1. **Automatic Directory Creation / 자동 디렉토리 생성:**
   - All write operations automatically create parent directories if they don't exist
   - 모든 쓰기 작업은 상위 디렉토리가 존재하지 않으면 자동으로 생성합니다

2. **Cross-Platform Compatibility / 크로스 플랫폼 호환성:**
   - All path operations use `filepath` package for OS-agnostic behavior
   - 모든 경로 작업은 OS에 구애받지 않는 동작을 위해 `filepath` 패키지를 사용합니다

3. **Buffered I/O / 버퍼링된 I/O:**
   - Default 32KB buffer size for optimal performance
   - 최적의 성능을 위한 기본 32KB 버퍼 크기

4. **Atomic Operations / 원자적 작업:**
   - `WriteAtomic` function for safe file updates (write to temp, then rename)
   - 안전한 파일 업데이트를 위한 `WriteAtomic` 함수 (임시 파일에 쓰기, 그 다음 이름 변경)

5. **Progress Callbacks / 진행 상황 콜백:**
   - Copy operations support progress callbacks for large files
   - 복사 작업은 대용량 파일에 대한 진행 상황 콜백을 지원합니다

6. **Flexible Copying / 유연한 복사:**
   - Functional options pattern for copy operations
   - 복사 작업을 위한 함수형 옵션 패턴
   - Options: overwrite, preserve permissions, preserve timestamps, progress, filter
   - 옵션: 덮어쓰기, 권한 보존, 타임스탬프 보존, 진행 상황, 필터

7. **Multiple Hash Algorithms / 여러 해시 알고리즘:**
   - Support for MD5, SHA1, SHA256, SHA512
   - MD5, SHA1, SHA256, SHA512 지원

8. **Path Safety / 경로 안전성:**
   - `IsSafe` function to prevent directory traversal attacks
   - 디렉토리 탐색 공격을 방지하기 위한 `IsSafe` 함수

9. **Human-Readable Output / 사람이 읽기 쉬운 출력:**
   - `SizeHuman` converts bytes to KB/MB/GB/TB
   - `SizeHuman`은 바이트를 KB/MB/GB/TB로 변환합니다

10. **Zero External Dependencies / 외부 의존성 없음:**
    - Only uses standard library (except gopkg.in/yaml.v3)
    - 표준 라이브러리만 사용 (gopkg.in/yaml.v3 제외)

### Design Philosophy / 설계 철학

**"20줄 → 1줄" (20 lines → 1 line)**

Reducing repetitive file manipulation code from 20+ lines to just 1-2 lines:

반복적인 파일 조작 코드를 20줄 이상에서 단 1-2줄로 줄입니다:

```go
// Before: 20+ lines
dir := filepath.Dir(path)
if err := os.MkdirAll(dir, 0755); err != nil {
    return err
}
file, err := os.Create(path)
if err != nil {
    return err
}
defer file.Close()
if _, err := file.WriteString(content); err != nil {
    return err
}

// After: 1 line
err := fileutil.WriteString(path, content)
```

### Fixed / 수정됨

1. **Function Name Collision / 함수 이름 충돌:**
   - Renamed `Clean` in `path.go` to `CleanPath` to avoid conflict with `Clean` in `delete.go`
   - `path.go`의 `Clean`을 `CleanPath`로 이름 변경하여 `delete.go`의 `Clean`과 충돌 방지
   - Updated references in `Normalize()` and `IsSafe()` functions
   - `Normalize()` 및 `IsSafe()` 함수의 참조 업데이트

2. **Missing Import / 누락된 임포트:**
   - Added `path/filepath` import to `hash.go` for `HashDir` function
   - `HashDir` 함수를 위해 `hash.go`에 `path/filepath` 임포트 추가

### Technical Details / 기술 세부사항

- **Go Version**: 1.18+ (uses standard library generics where appropriate)
- **Dependencies**: Standard library + `gopkg.in/yaml.v3`
- **Constants**:
  - `DefaultFileMode = 0644` - Default file permissions
  - `DefaultDirMode = 0755` - Default directory permissions
  - `DefaultBufferSize = 32 * 1024` - 32KB buffer for I/O operations
  - `DefaultChunkSize = 1024 * 1024` - 1MB chunk size for large file processing

### Performance / 성능

- Buffered I/O for all file operations (32KB buffer)
- 모든 파일 작업에 버퍼링된 I/O 사용 (32KB 버퍼)
- Efficient chunk-based processing for large files (1MB chunks)
- 대용량 파일에 대한 효율적인 청크 기반 처리 (1MB 청크)
- Optimized directory walking with filter support
- 필터 지원을 통한 최적화된 디렉토리 순회

---

**Total Functions Implemented / 구현된 총 함수 수: ~91 functions across 12 categories**

**Total Functions Implemented / 구현된 총 함수 수: 12개 카테고리에 걸쳐 약 91개 함수**
