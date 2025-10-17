package fileutil

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// WriteFile writes data to a file, creating the directory if it doesn't exist
// WriteFile은 파일에 데이터를 쓰고, 디렉토리가 존재하지 않으면 생성합니다
//
// Example
// 예제:
//
//	err := fileutil.WriteFile("path/to/file.txt", []byte("Hello, World!"))
//	if err != nil {
//	    log.Fatal(err)
//	}
func WriteFile(path string, data []byte, perm ...os.FileMode) error {
	mode := DefaultFileMode
	if len(perm) > 0 {
		mode = perm[0]
	}

	// Create directory if it doesn't exist
	// 디렉토리가 존재하지 않으면 생성
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, DefaultDirMode); err != nil {
		return fmt.Errorf("fileutil.WriteFile: %w", err)
	}

	// Write file
	// 파일 쓰기
	if err := os.WriteFile(path, data, mode); err != nil {
		return fmt.Errorf("fileutil.WriteFile: %w", err)
	}

	return nil
}

// WriteString writes a string to a file, creating the directory if it doesn't exist
// WriteString은 파일에 문자열을 쓰고, 디렉토리가 존재하지 않으면 생성합니다
//
// Example
// 예제:
//
//	err := fileutil.WriteString("path/to/file.txt", "Hello, World!")
//	if err != nil {
//	    log.Fatal(err)
//	}
func WriteString(path string, content string, perm ...os.FileMode) error {
	return WriteFile(path, []byte(content), perm...)
}

// WriteLines writes a slice of lines to a file, creating the directory if it doesn't exist
// WriteLines는 줄의 슬라이스를 파일에 쓰고, 디렉토리가 존재하지 않으면 생성합니다
//
// Example
// 예제:
//
//	lines := []string{"Line 1", "Line 2", "Line 3"}
//	err := fileutil.WriteLines("path/to/file.txt", lines)
//	if err != nil {
//	    log.Fatal(err)
//	}
func WriteLines(path string, lines []string, perm ...os.FileMode) error {
	content := strings.Join(lines, "\n")
	if len(lines) > 0 {
		// Add trailing newline
		// 후행 개행 추가
		content += "\n"
	}
	return WriteString(path, content, perm...)
}

// WriteJSON marshals a value to JSON and writes it to a file
// WriteJSON은 값을 JSON으로 마샬하고 파일에 씁니다
//
// Example
// 예제:
//
//	config := Config{Name: "app", Version: "1.0.0"}
//	err := fileutil.WriteJSON("config.json", config)
//	if err != nil {
//	    log.Fatal(err)
//	}
func WriteJSON(path string, v interface{}, perm ...os.FileMode) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("fileutil.WriteJSON: %w", err)
	}

	return WriteFile(path, data, perm...)
}

// WriteYAML marshals a value to YAML and writes it to a file
// WriteYAML은 값을 YAML로 마샬하고 파일에 씁니다
//
// Example
// 예제:
//
//	config := Config{Name: "app", Version: "1.0.0"}
//	err := fileutil.WriteYAML("config.yaml", config)
//	if err != nil {
//	    log.Fatal(err)
//	}
func WriteYAML(path string, v interface{}, perm ...os.FileMode) error {
	data, err := yaml.Marshal(v)
	if err != nil {
		return fmt.Errorf("fileutil.WriteYAML: %w", err)
	}

	return WriteFile(path, data, perm...)
}

// WriteCSV writes records to a CSV file
// WriteCSV는 레코드를 CSV 파일에 씁니다
//
// Example
// 예제:
//
//	records := [][]string{
//	    {"Name", "Age", "City"},
//	    {"Alice", "30", "Seoul"},
//	    {"Bob", "25", "Busan"},
//	}
//	err := fileutil.WriteCSV("data.csv", records)
//	if err != nil {
//	    log.Fatal(err)
//	}
func WriteCSV(path string, records [][]string, perm ...os.FileMode) error {
	mode := DefaultFileMode
	if len(perm) > 0 {
		mode = perm[0]
	}

	// Create directory if it doesn't exist
	// 디렉토리가 존재하지 않으면 생성
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, DefaultDirMode); err != nil {
		return fmt.Errorf("fileutil.WriteCSV: %w", err)
	}

	// Create file
	// 파일 생성
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		return fmt.Errorf("fileutil.WriteCSV: %w", err)
	}
	defer file.Close()

	// Write CSV
	// CSV 쓰기
	writer := csv.NewWriter(file)
	if err := writer.WriteAll(records); err != nil {
		return fmt.Errorf("fileutil.WriteCSV: %w", err)
	}

	return nil
}

// WriteAtomic writes data to a file atomically by writing to a temporary file first
// and then renaming it to the target path
// WriteAtomic는 먼저 임시 파일에 쓴 다음 대상 경로로 이름을 변경하여 파일에 원자적으로 데이터를 씁니다
//
// This prevents corruption if the write operation is interrupted.
// 이는 쓰기 작업이 중단될 경우 손상을 방지합니다.
//
// Example
// 예제:
//
//	err := fileutil.WriteAtomic("important.txt", []byte("Critical data"))
//	if err != nil {
//	    log.Fatal(err)
//	}
func WriteAtomic(path string, data []byte, perm ...os.FileMode) error {
	mode := DefaultFileMode
	if len(perm) > 0 {
		mode = perm[0]
	}

	// Create directory if it doesn't exist
	// 디렉토리가 존재하지 않으면 생성
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, DefaultDirMode); err != nil {
		return fmt.Errorf("fileutil.WriteAtomic: %w", err)
	}

	// Write to temporary file
	// 임시 파일에 쓰기
	tempPath := path + ".tmp"
	if err := os.WriteFile(tempPath, data, mode); err != nil {
		return fmt.Errorf("fileutil.WriteAtomic: %w", err)
	}

	// Rename temporary file to target path
	// 임시 파일을 대상 경로로 이름 변경
	if err := os.Rename(tempPath, path); err != nil {
		// Clean up temp file on error
		// 에러 시 임시 파일 정리
		os.Remove(tempPath)
		return fmt.Errorf("fileutil.WriteAtomic: %w", err)
	}

	return nil
}

// AppendFile appends data to a file, creating it if it doesn't exist
// AppendFile은 파일에 데이터를 추가하고, 파일이 존재하지 않으면 생성합니다
//
// Example
// 예제:
//
//	err := fileutil.AppendFile("log.txt", []byte("New log entry\n"))
//	if err != nil {
//	    log.Fatal(err)
//	}
func AppendFile(path string, data []byte) error {
	// Create directory if it doesn't exist
	// 디렉토리가 존재하지 않으면 생성
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, DefaultDirMode); err != nil {
		return fmt.Errorf("fileutil.AppendFile: %w", err)
	}

	// Open file in append mode
	// 추가 모드로 파일 열기
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, DefaultFileMode)
	if err != nil {
		return fmt.Errorf("fileutil.AppendFile: %w", err)
	}
	defer file.Close()

	// Write data
	// 데이터 쓰기
	if _, err := file.Write(data); err != nil {
		return fmt.Errorf("fileutil.AppendFile: %w", err)
	}

	return nil
}

// AppendString appends a string to a file, creating it if it doesn't exist
// AppendString은 파일에 문자열을 추가하고, 파일이 존재하지 않으면 생성합니다
//
// Example
// 예제:
//
//	err := fileutil.AppendString("log.txt", "New log entry\n")
//	if err != nil {
//	    log.Fatal(err)
//	}
func AppendString(path string, content string) error {
	return AppendFile(path, []byte(content))
}

// AppendLines appends lines to a file, creating it if it doesn't exist
// AppendLines는 파일에 줄을 추가하고, 파일이 존재하지 않으면 생성합니다
//
// Example
// 예제:
//
//	lines := []string{"Line 1", "Line 2", "Line 3"}
//	err := fileutil.AppendLines("log.txt", lines)
//	if err != nil {
//	    log.Fatal(err)
//	}
func AppendLines(path string, lines []string) error {
	content := strings.Join(lines, "\n")
	if len(lines) > 0 {
		// Add trailing newline
		// 후행 개행 추가
		content += "\n"
	}
	return AppendString(path, content)
}

// AppendBytes is an alias for AppendFile
// AppendBytes는 AppendFile의 별칭입니다
func AppendBytes(path string, data []byte) error {
	return AppendFile(path, data)
}

// WriteStream writes data to a file using buffered I/O
// WriteStream은 버퍼링된 I/O를 사용하여 파일에 데이터를 씁니다
//
// This is useful for writing large amounts of data efficiently.
// 이는 대량의 데이터를 효율적으로 쓰는 데 유용합니다.
//
// Example
// 예제:
//
//	file, err := fileutil.CreateFile("large-file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer file.Close()
//
//	writer := bufio.NewWriter(file)
//	for _, line := range lines {
//	    writer.WriteString(line + "\n")
//	}
//	writer.Flush()
func CreateFile(path string, perm ...os.FileMode) (*os.File, error) {
	mode := DefaultFileMode
	if len(perm) > 0 {
		mode = perm[0]
	}

	// Create directory if it doesn't exist
	// 디렉토리가 존재하지 않으면 생성
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, DefaultDirMode); err != nil {
		return nil, fmt.Errorf("fileutil.CreateFile: %w", err)
	}

	// Create file
	// 파일 생성
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		return nil, fmt.Errorf("fileutil.CreateFile: %w", err)
	}

	return file, nil
}

// NewWriter creates a new buffered writer for the specified file
// NewWriter는 지정된 파일에 대한 새 버퍼링된 writer를 생성합니다
//
// Example
// 예제:
//
//	writer, file, err := fileutil.NewWriter("output.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer file.Close()
//	defer writer.Flush()
//
//	for _, line := range lines {
//	    writer.WriteString(line + "\n")
//	}
func NewWriter(path string, perm ...os.FileMode) (*bufio.Writer, *os.File, error) {
	file, err := CreateFile(path, perm...)
	if err != nil {
		return nil, nil, err
	}

	writer := bufio.NewWriter(file)
	return writer, file, nil
}
