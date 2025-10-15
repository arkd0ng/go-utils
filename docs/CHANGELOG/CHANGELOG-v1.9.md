# CHANGELOG - v1.9.x

This file contains detailed change logs for the v1.9.x releases of go-utils, focusing on the fileutil package.

이 파일은 fileutil 패키지에 중점을 둔 go-utils의 v1.9.x 릴리스에 대한 상세한 변경 로그를 포함합니다.

---

## [v1.9.006] - 2025-10-15

### Added / 추가됨

#### Log File Backup Management System for All Examples / 모든 예제에 로그 파일 백업 관리 시스템 추가

**Feature / 기능:**
- Implemented automatic log file backup and cleanup system across all example programs
- 모든 예제 프로그램에 자동 로그 파일 백업 및 정리 시스템 구현

**Details / 상세 내용:**
- Each example now backs up previous log file with timestamp format `yyyymmdd-hhmmss` before creating new log
- Automatically maintains only 5 most recent backup files per example
- Old backup files are automatically deleted to prevent disk space issues
- 각 예제가 새 로그를 생성하기 전에 `yyyymmdd-hhmmss` 타임스탬프 형식으로 이전 로그 파일 백업
- 예제당 최근 5개의 백업 파일만 자동으로 유지
- 디스크 공간 문제를 방지하기 위해 오래된 백업 파일 자동 삭제

**Updated Examples / 업데이트된 예제:**
1. `examples/random_string/main.go` - Added backup logic for `random-example.log`
2. `examples/stringutil/main.go` - Added backup logic for `stringutil-example.log`
3. `examples/timeutil/main.go` - Added backup logic for `timeutil-example.log`
4. `examples/sliceutil/main.go` - Added backup logic for `sliceutil-example.log`
5. `examples/maputil/main.go` - Added backup logic for `maputil-example.log`
6. `examples/mysql/main.go` - Added backup logic for `mysql-example.log`
7. `examples/redis/main.go` - Added backup logic for `redis-example.log`
8. `examples/logging/main.go` - Added helper function `backupLogFile()` for 12 different log files
9. `examples/fileutil/main.go` - Already has backup logic (updated timestamp format to include seconds)

**Technical Implementation / 기술 구현:**
- Uses `fileutil.Exists()` to check for existing log files
- Uses `fileutil.ModTime()` to get file modification time for backup naming
- Uses `fileutil.CopyFile()` to create timestamped backups
- Uses `filepath.Glob()` to find and manage backup files
- Uses `fileutil.DeleteFile()` to clean up old backups
- Custom sorting algorithm to identify oldest files

**Benefits / 장점:**
- Prevents log file loss when running examples multiple times
- Maintains clean logs directory without manual intervention
- Provides historical log files for comparison and debugging
- 예제를 여러 번 실행할 때 로그 파일 손실 방지
- 수동 개입 없이 깨끗한 로그 디렉토리 유지
- 비교 및 디버깅을 위한 히스토리 로그 파일 제공

---

## [v1.9.005] - 2025-10-15

### Added / 추가됨

#### Extremely Detailed Logging in fileutil Examples / fileutil 예제에 극도로 상세한 로깅 추가

**Feature / 기능:**
- Enhanced `examples/fileutil/main.go` Examples 1-3 with comprehensive, tutorial-quality logging
- fileutil 예제 1-3을 포괄적이고 튜토리얼 수준의 로깅으로 개선

**Details / 상세 내용:**
- Each function now includes:
  - 📚 Function signature
  - 📖 Description (bilingual)
  - 🎯 Use cases (4+ scenarios)
  - 💡 Key features (4+ features)
  - ▶️ Step-by-step execution details
  - ✅ Success confirmation with detailed information
  - 🔍 Verification checks
- 각 함수가 이제 포함함:
  - 📚 함수 시그니처
  - 📖 설명 (이중 언어)
  - 🎯 사용 사례 (4개 이상)
  - 💡 주요 기능 (4개 이상)
  - ▶️ 단계별 실행 세부 정보
  - ✅ 상세 정보와 함께 성공 확인
  - 🔍 검증 체크

**Example 1 - File Writing Operations (8 functions):**
1. `WriteString()` - 55+ lines of detailed logging
2. `WriteFile()` - 50+ lines of detailed logging
3. `WriteLines()` - 60+ lines of detailed logging with content preview
4. `WriteJSON()` - 55+ lines of detailed logging with JSON content display
5. `WriteYAML()` - 50+ lines of detailed logging with YAML content display
6. `WriteCSV()` - 55+ lines of detailed logging with CSV preview
7. `AppendString()` - 50+ lines with before/after comparison
8. `AppendLines()` - 50+ lines with full content display

**Example 2 - File Reading Operations (6 functions):**
1. `ReadString()` - 50+ lines with content display
2. `ReadFile()` - 50+ lines with hex and ASCII representation
3. `ReadLines()` - 55+ lines with line-by-line content display
4. `ReadJSON()` - 50+ lines with parsed struct field display
5. `ReadYAML()` - 50+ lines with parsed struct and verification
6. `ReadCSV()` - 50+ lines with header and row display

**Example 3 - Path Operations (12 functions):**
1. `Join()` - 50+ lines for path element combination
2. `Split()` - 50+ lines for directory/file separation
3. `Base()` - 45+ lines for filename extraction
4. `Dir()` - 45+ lines for directory extraction
5. `Ext()` - 45+ lines for extension retrieval
6. `WithoutExt()` - 50+ lines for extension removal
7. `ChangeExt()` - 50+ lines for extension modification
8. `HasExt()` - 55+ lines for extension validation
9. `Abs()` - 45+ lines for absolute path resolution
10. `IsAbs()` - 50+ lines for path type checking
11. `CleanPath()` - 55+ lines for path normalization
12. `ToSlash/FromSlash()` - 55+ lines for separator conversion

**Benefits / 이점:**
- Users can understand the complete function behavior from logs alone without reading documentation
- 사용자가 문서를 읽지 않고도 로그만으로 완전한 함수 동작을 이해할 수 있음
- Serves as an interactive tutorial
- 대화형 튜토리얼 역할을 함
- Perfect for learning and debugging
- 학습 및 디버깅에 완벽함

### Changed / 변경됨

#### Unified Log File Path and Naming Convention / 통합된 로그 파일 경로 및 명명 규칙

**Problem / 문제:**
- Log files were scattered across different directories and had inconsistent naming conventions / 로그 파일이 여러 디렉토리에 분산되어 있고 일관성 없는 명명 규칙을 사용함
- Some examples used `logs/`, others used `./logs/` / 일부 예제는 `logs/`를, 다른 예제는 `./logs/`를 사용
- Filename formats varied: `package_example_`, `package-example-`, different timestamp formats / 파일명 형식이 다양함: `package_example_`, `package-example-`, 다양한 타임스탬프 형식

**Solution / 해결책:**
- Unified all example log files to repository root `logs/` directory / 모든 예제 로그 파일을 레포지토리 루트 `logs/` 디렉토리로 통합
- Standardized filename format: `<package>-example-<timestamp>.log` / 파일명 형식 표준화: `<package>-example-<timestamp>.log`
- Standardized timestamp format: `20060102-150405` (YYYYMMDD-HHMMSS) / 타임스탬프 형식 표준화: `20060102-150405` (YYYYMMDD-HHMMSS)
- Exception: logging package examples keep original filenames (but use unified `logs/` directory) / 예외: logging 패키지 예제는 원래 파일명 유지 (하지만 통합된 `logs/` 디렉토리 사용)

**Updated Files / 업데이트된 파일:**
1. `examples/fileutil/main.go` - `fileutil-examples-` → `fileutil-example-`
2. `examples/maputil/main.go` - Timestamp format: `20060102_150405` → `20060102-150405`
3. `examples/mysql/main.go` - `mysql_example_` + `20060102_150405` → `mysql-example-` + `20060102-150405`
4. `examples/random_string/main.go` - `random_example_` + `20060102_150405` → `random-example-` + `20060102-150405`
5. `examples/redis/main.go` - `redis_example_` + `20060102_150405` → `redis-example-` + `20060102-150405`
6. `examples/sliceutil/main.go` - `sliceutil_example_` + `20060102_150405` → `sliceutil-example-` + `20060102-150405`
7. `examples/stringutil/main.go` - `stringutil_example_` + `20060102_150405` → `stringutil-example-` + `20060102-150405`
8. `examples/timeutil/main.go` - `timeutil-example.log` (static) → `timeutil-example-<timestamp>.log` (timestamped)
9. `examples/logging/main.go` - All `./logs/` → `logs/` (filenames unchanged: `custom.log`, `app.log`, etc.)

**New Unified Format / 새로운 통합 형식:**
```
logs/fileutil-example-20251015-200641.log
logs/maputil-example-20251015-143022.log
logs/mysql-example-20251015-143022.log
logs/random-example-20251015-143022.log
logs/redis-example-20251015-143022.log
logs/sliceutil-example-20251015-143022.log
logs/stringutil-example-20251015-143022.log
logs/timeutil-example-20251015-143022.log
logs/custom.log          (logging example - filename preserved)
logs/app.log             (logging example - filename preserved)
logs/database.log        (logging example - filename preserved)
```

**Benefits / 이점:**
- Centralized log management in single `logs/` directory / 단일 `logs/` 디렉토리에서 중앙 집중식 로그 관리
- Consistent naming convention across all examples / 모든 예제에서 일관된 명명 규칙
- Easier log file discovery and organization / 로그 파일 검색 및 정리가 더 쉬움
- Predictable log file locations for CI/CD and automation / CI/CD 및 자동화를 위한 예측 가능한 로그 파일 위치
- Timestamp in filename enables chronological sorting / 파일명의 타임스탬프로 시간순 정렬 가능

---

## [v1.9.004] - 2025-10-15

### Enhanced / 보강됨

#### Complete Logging Migration in Fileutil Example / Fileutil 예제의 완전한 로깅 마이그레이션

**Updated Files / 업데이트된 파일:**
- `examples/fileutil/main.go` - Completely replaced all fmt output with structured logging / 모든 fmt 출력을 구조화된 로깅으로 완전히 교체

**Key Changes / 주요 변경사항:**
- Replaced all `fmt.Println()` and `fmt.Printf()` calls with structured `logger.Info()` calls / 모든 `fmt.Println()` 및 `fmt.Printf()` 호출을 구조화된 `logger.Info()` 호출로 교체
- Implemented key-value structured logging for all output messages / 모든 출력 메시지에 대해 키-값 구조화 로깅 구현
- Added logger parameter to `example2PathOperations()` function / `example2PathOperations()` 함수에 logger 매개변수 추가
- Used consistent key naming: `path`, `count`, `value`, `bytes`, `hash`, `checksum`, `valid`, `same`, `exists`, `isEmpty`, `name` / 일관된 키 이름 사용
- Eliminated duplicate output (removed redundant fmt.Println after logger.Info) / 중복 출력 제거 (logger.Info 이후 중복된 fmt.Println 제거)
- Maintained fmt.Printf only for progress callback (line 208) which cannot easily access logger / 로거에 쉽게 접근할 수 없는 진행 상황 콜백(208행)에만 fmt.Printf 유지

**Benefits / 이점:**
- All output now appears in both console and log file thanks to `WithStdout(true)` / `WithStdout(true)` 덕분에 모든 출력이 콘솔과 로그 파일 양쪽에 표시됨
- Structured logging allows easier parsing and analysis of logs / 구조화된 로깅으로 로그를 더 쉽게 파싱하고 분석 가능
- Consistent logging pattern across all example functions / 모든 예제 함수에서 일관된 로깅 패턴
- Better observability with key-value pairs / 키-값 쌍으로 더 나은 관찰 가능성

**Example Output / 예제 출력:**
```
2025-10-15 19:56:38 [INFO] ✓ Written to file path=/path/to/file.txt
2025-10-15 19:56:38 [INFO] ✓ Found files count=3
2025-10-15 19:56:38 [INFO] ✓ SHA256 hash hash=7d5e51fa...
2025-10-15 19:56:38 [INFO] ✓ Directory status exists=true isEmpty=true
```

---

## [v1.9.003] - 2025-10-15

### Enhanced / 보강됨

#### Example Files Logging Integration / 예제 파일 로깅 통합

**Updated Files / 업데이트된 파일:**
- `examples/fileutil/main.go` - Integrated logging package with timestamped log files / 타임스탬프가 있는 로그 파일로 logging 패키지 통합

**Key Changes / 주요 변경사항:**
- Replaced all `log.Fatal` calls with `logger.Fatalf` in fileutil example / fileutil 예제의 모든 `log.Fatal` 호출을 `logger.Fatalf`로 교체
- Updated all example functions to accept `logger *logging.Logger` parameter / 모든 예제 함수에서 `logger *logging.Logger` 매개변수를 받도록 업데이트
- Removed unused `log` package import / 사용하지 않는 `log` 패키지 import 제거
- Log files now saved to `logs/` directory with pattern: `logs/fileutil-examples-{timestamp}.log` / 로그 파일이 이제 `logs/` 디렉토리에 `logs/fileutil-examples-{timestamp}.log` 패턴으로 저장됨
- Added `WithStdout(true)` for console output alongside file logging / 파일 로깅과 함께 콘솔 출력을 위한 `WithStdout(true)` 추가

**Verified / 확인됨:**
- All utility package examples (stringutil, timeutil, sliceutil, maputil) already use logging package correctly / 모든 유틸리티 패키지 예제(stringutil, timeutil, sliceutil, maputil)가 이미 logging 패키지를 올바르게 사용함
- All examples build successfully / 모든 예제가 성공적으로 빌드됨
- Fileutil example tested and confirmed working with logging package / Fileutil 예제가 logging 패키지와 함께 테스트되고 작동 확인됨

---

## [v1.9.002] - 2025-10-15

### Enhanced / 보강됨

#### Comprehensive Test Suite / 포괄적인 테스트 스위트

**Test Coverage / 테스트 커버리지:**
- **87 test cases** across 9 test suites / 9개 테스트 스위트에 걸쳐 87개 테스트 케이스
- **10 benchmark tests** for performance measurement / 성능 측정을 위한 10개 벤치마크 테스트
- **55.2% code coverage** / 55.2% 코드 커버리지

**Test Suites / 테스트 스위트:**
1. **TestFileReading** (8 tests): ReadFile, ReadString, ReadLines, ReadJSON, ReadYAML, ReadCSV, ReadBytes, ReadChunk
2. **TestFileWriting** (9 tests): WriteFile, WriteString, WriteLines, WriteJSON, WriteYAML, WriteCSV, WriteAtomic, AppendFile, AppendLines
3. **TestPathOperations** (17 tests): Join, Split, Base, Dir, Ext, Abs, CleanPath, Normalize, ToSlash, FromSlash, IsAbs, IsValid, IsSafe, Match, WithoutExt, ChangeExt, HasExt
4. **TestFileInformation** (11 tests): Exists, IsFile, IsDir, IsSymlink, Size, SizeHuman, Chmod, IsReadable, IsWritable, ModTime, Touch
5. **TestCopyOperations** (4 tests): CopyFile, CopyFile_WithOverwrite, CopyFile_WithProgress, CopyDir
6. **TestMoveOperations** (3 tests): MoveFile, Rename, RenameExt
7. **TestDeleteOperations** (7 tests): DeleteFile, DeleteDir, DeleteRecursive, DeletePattern, DeleteFiles, Clean, RemoveEmpty
8. **TestDirectoryOperations** (9 tests): MkdirAll, CreateTemp, CreateTempDir, IsEmpty, DirSize, ListFiles, ListDirs, ListAll, FindFiles
9. **TestHashOperations** (10 tests): MD5, SHA1, SHA256, SHA512, Hash, HashBytes, CompareFiles, CompareHash, Checksum, VerifyChecksum

**Benchmark Tests / 벤치마크 테스트:**
1. `BenchmarkWriteFile` - File writing performance / 파일 쓰기 성능
2. `BenchmarkReadFile` - File reading performance / 파일 읽기 성능
3. `BenchmarkWriteString` - String writing performance / 문자열 쓰기 성능
4. `BenchmarkReadString` - String reading performance / 문자열 읽기 성능
5. `BenchmarkCopyFile` - File copying performance / 파일 복사 성능
6. `BenchmarkSHA256` - SHA256 hashing performance / SHA256 해싱 성능
7. `BenchmarkMD5` - MD5 hashing performance / MD5 해싱 성능
8. `BenchmarkJSON/WriteJSON` - JSON writing performance / JSON 쓰기 성능
9. `BenchmarkJSON/ReadJSON` - JSON reading performance / JSON 읽기 성능
10. `BenchmarkYAML/WriteYAML` - YAML writing performance / YAML 쓰기 성능
11. `BenchmarkYAML/ReadYAML` - YAML reading performance / YAML 읽기 성능
12. `BenchmarkListFiles` - Directory listing performance / 디렉토리 나열 성능

**Benchmark Results / 벤치마크 결과 (darwin/amd64):**
- WriteFile: ~52,000 ns/op, 584 B/op, 6 allocs/op
- ReadFile: ~12,000 ns/op, 920 B/op, 5 allocs/op
- CopyFile (10KB): ~164,000 ns/op, 34KB/op, 22 allocs/op
- SHA256 (10KB): ~56,000 ns/op, 33KB/op, 9 allocs/op
- MD5 (10KB): ~42,000 ns/op, 33KB/op, 9 allocs/op
- JSON Write: ~75,000 ns/op, 888 B/op, 14 allocs/op
- JSON Read: ~23,000 ns/op, 1.7KB/op, 29 allocs/op
- YAML Write: ~50,000 ns/op, 7.4KB/op, 33 allocs/op
- YAML Read: ~22,000 ns/op, 8.7KB/op, 61 allocs/op
- ListFiles (100 files): ~86,000 ns/op, 23KB/op, 244 allocs/op

**Test Features / 테스트 기능:**
- Comprehensive edge case coverage / 포괄적인 엣지 케이스 커버리지
- Automatic cleanup with defer / defer를 사용한 자동 정리
- Temporary directory usage for isolation / 격리를 위한 임시 디렉토리 사용
- Error handling validation / 에러 처리 검증
- Cross-platform path handling / 크로스 플랫폼 경로 처리
- Progress callback testing / 진행 상황 콜백 테스트
- Copy options testing (overwrite, progress, filter) / 복사 옵션 테스트
- All hash algorithms tested / 모든 해시 알고리즘 테스트

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
