package fileutil

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

// ReadFile reads the entire file and returns its contents as a byte slice
// ReadFile은 전체 파일을 읽고 내용을 바이트 슬라이스로 반환합니다
//
// Example / 예제:
//
//	data, err := fileutil.ReadFile("path/to/file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(string(data))
func ReadFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("fileutil.ReadFile: %w", err)
	}
	return data, nil
}

// ReadString reads the entire file and returns its contents as a string
// ReadString은 전체 파일을 읽고 내용을 문자열로 반환합니다
//
// Example / 예제:
//
//	content, err := fileutil.ReadString("path/to/file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(content)
func ReadString(path string) (string, error) {
	data, err := ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ReadLines reads a file and returns its contents as a slice of lines
// ReadLines는 파일을 읽고 내용을 줄의 슬라이스로 반환합니다
//
// Example / 예제:
//
//	lines, err := fileutil.ReadLines("path/to/file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, line := range lines {
//	    fmt.Println(line)
//	}
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("fileutil.ReadLines: %w", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	// Set max buffer size to handle long lines / 긴 줄을 처리하기 위해 최대 버퍼 크기 설정
	buf := make([]byte, 0, DefaultBufferSize)
	scanner.Buffer(buf, 1024*1024) // 1MB max token size

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("fileutil.ReadLines: %w", err)
	}

	return lines, nil
}

// ReadJSON reads a JSON file and unmarshals it into the provided value
// ReadJSON은 JSON 파일을 읽고 제공된 값으로 언마샬합니다
//
// Example / 예제:
//
//	var config Config
//	err := fileutil.ReadJSON("config.json", &config)
//	if err != nil {
//	    log.Fatal(err)
//	}
func ReadJSON(path string, v interface{}) error {
	data, err := ReadFile(path)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("fileutil.ReadJSON: %w", err)
	}

	return nil
}

// ReadYAML reads a YAML file and unmarshals it into the provided value
// ReadYAML은 YAML 파일을 읽고 제공된 값으로 언마샬합니다
//
// Example / 예제:
//
//	var config Config
//	err := fileutil.ReadYAML("config.yaml", &config)
//	if err != nil {
//	    log.Fatal(err)
//	}
func ReadYAML(path string, v interface{}) error {
	data, err := ReadFile(path)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, v); err != nil {
		return fmt.Errorf("fileutil.ReadYAML: %w", err)
	}

	return nil
}

// ReadCSV reads a CSV file and returns its contents as a slice of string slices
// ReadCSV는 CSV 파일을 읽고 내용을 문자열 슬라이스의 슬라이스로 반환합니다
//
// Example / 예제:
//
//	records, err := fileutil.ReadCSV("data.csv")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, record := range records {
//	    fmt.Println(record)
//	}
func ReadCSV(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("fileutil.ReadCSV: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("fileutil.ReadCSV: %w", err)
	}

	return records, nil
}

// ReadBytes reads a specific number of bytes from a file at a given offset
// ReadBytes는 주어진 오프셋에서 파일로부터 특정 바이트 수를 읽습니다
//
// Example / 예제:
//
//	// Read 100 bytes starting from offset 1000
//	data, err := fileutil.ReadBytes("large-file.bin", 1000, 100)
//	if err != nil {
//	    log.Fatal(err)
//	}
func ReadBytes(path string, offset, size int64) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("fileutil.ReadBytes: %w", err)
	}
	defer file.Close()

	// Seek to offset / 오프셋으로 이동
	if _, err := file.Seek(offset, io.SeekStart); err != nil {
		return nil, fmt.Errorf("fileutil.ReadBytes: %w", err)
	}

	// Read size bytes / size 바이트 읽기
	data := make([]byte, size)
	n, err := io.ReadFull(file, data)
	if err != nil && err != io.ErrUnexpectedEOF {
		return nil, fmt.Errorf("fileutil.ReadBytes: %w", err)
	}

	return data[:n], nil
}

// ReadChunk reads a file in chunks and calls the provided function for each chunk
// ReadChunk는 파일을 청크 단위로 읽고 각 청크마다 제공된 함수를 호출합니다
//
// This is useful for processing large files without loading them entirely into memory.
// 이는 파일을 메모리에 전체 로드하지 않고 큰 파일을 처리하는 데 유용합니다.
//
// Example / 예제:
//
//	err := fileutil.ReadChunk("large-file.txt", 1024*1024, func(chunk []byte) error {
//	    // Process chunk / 청크 처리
//	    fmt.Printf("Read %d bytes\n", len(chunk))
//	    return nil
//	})
func ReadChunk(path string, chunkSize int64, fn func([]byte) error) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("fileutil.ReadChunk: %w", err)
	}
	defer file.Close()

	buffer := make([]byte, chunkSize)
	for {
		n, err := file.Read(buffer)
		if n > 0 {
			if err := fn(buffer[:n]); err != nil {
				return fmt.Errorf("fileutil.ReadChunk: %w", err)
			}
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("fileutil.ReadChunk: %w", err)
		}
	}

	return nil
}
