package fileutil

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// MD5 calculates the MD5 hash of a file
// MD5는 파일의 MD5 해시를 계산합니다
//
// Example
// 예제:
//
//	hash, err := fileutil.MD5("path/to/file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println("MD5:", hash)
func MD5(path string) (string, error) {
	return Hash(path, "md5")
}

// SHA1 calculates the SHA1 hash of a file
// SHA1은 파일의 SHA1 해시를 계산합니다
//
// Example
// 예제:
//
//	hash, err := fileutil.SHA1("path/to/file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println("SHA1:", hash)
func SHA1(path string) (string, error) {
	return Hash(path, "sha1")
}

// SHA256 calculates the SHA256 hash of a file
// SHA256은 파일의 SHA256 해시를 계산합니다
//
// Example
// 예제:
//
//	hash, err := fileutil.SHA256("path/to/file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println("SHA256:", hash)
func SHA256(path string) (string, error) {
	return Hash(path, "sha256")
}

// SHA512 calculates the SHA512 hash of a file
// SHA512는 파일의 SHA512 해시를 계산합니다
//
// Example
// 예제:
//
//	hash, err := fileutil.SHA512("path/to/file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println("SHA512:", hash)
func SHA512(path string) (string, error) {
	return Hash(path, "sha512")
}

// Hash calculates the hash of a file using the specified algorithm
// Hash는 지정된 알고리즘을 사용하여 파일의 해시를 계산합니다
//
// Supported algorithms: md5, sha1, sha256, sha512
// 지원되는 알고리즘: md5, sha1, sha256, sha512
//
// Example
// 예제:
//
//	hash, err := fileutil.Hash("path/to/file.txt", "sha256")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println("Hash:", hash)
func Hash(path string, algorithm string) (string, error) {
	// Open file
	// 파일 열기
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("fileutil.Hash: %w", err)
	}
	defer file.Close()

	// Create hasher based on algorithm
	// 알고리즘에 따라 해셔 생성
	var hasher hash.Hash
	switch strings.ToLower(algorithm) {
	case "md5":
		hasher = md5.New()
	case "sha1":
		hasher = sha1.New()
	case "sha256":
		hasher = sha256.New()
	case "sha512":
		hasher = sha512.New()
	default:
		return "", fmt.Errorf("fileutil.Hash: unsupported algorithm: %s", algorithm)
	}

	// Copy file contents to hasher
	// 파일 내용을 해셔로 복사
	if _, err := io.Copy(hasher, file); err != nil {
		return "", fmt.Errorf("fileutil.Hash: %w", err)
	}

	// Return hex-encoded hash
	// 16진수 인코딩된 해시 반환
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

// HashBytes calculates the hash of a file and returns it as a byte slice
// HashBytes는 파일의 해시를 계산하고 바이트 슬라이스로 반환합니다
//
// This is useful when you need the raw hash bytes instead of a hex string.
// 이는 16진수 문자열 대신 원시 해시 바이트가 필요할 때 유용합니다.
//
// Example
// 예제:
//
//	hashBytes, err := fileutil.HashBytes("path/to/file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
func HashBytes(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("fileutil.HashBytes: %w", err)
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return nil, fmt.Errorf("fileutil.HashBytes: %w", err)
	}

	return hasher.Sum(nil), nil
}

// CompareFiles compares two files byte by byte
// CompareFiles는 두 파일을 바이트 단위로 비교합니다
//
// Returns true if the files are identical, false otherwise.
// 파일이 동일하면 true를, 그렇지 않으면 false를 반환합니다.
//
// Example
// 예제:
//
//	same, err := fileutil.CompareFiles("file1.txt", "file2.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	if same {
//	    fmt.Println("Files are identical")
//	}
func CompareFiles(path1, path2 string) (bool, error) {
	// Check if both files exist
	// 두 파일 모두 존재하는지 확인
	info1, err := os.Stat(path1)
	if err != nil {
		return false, fmt.Errorf("fileutil.CompareFiles: %w", err)
	}
	info2, err := os.Stat(path2)
	if err != nil {
		return false, fmt.Errorf("fileutil.CompareFiles: %w", err)
	}

	// Quick check: if sizes are different, files are different
	// 빠른 확인: 크기가 다르면 파일이 다름
	if info1.Size() != info2.Size() {
		return false, nil
	}

	// Open both files
	// 두 파일 모두 열기
	file1, err := os.Open(path1)
	if err != nil {
		return false, fmt.Errorf("fileutil.CompareFiles: %w", err)
	}
	defer file1.Close()

	file2, err := os.Open(path2)
	if err != nil {
		return false, fmt.Errorf("fileutil.CompareFiles: %w", err)
	}
	defer file2.Close()

	// Compare contents in chunks
	// 청크 단위로 내용 비교
	buf1 := make([]byte, DefaultBufferSize)
	buf2 := make([]byte, DefaultBufferSize)

	for {
		n1, err1 := file1.Read(buf1)
		n2, err2 := file2.Read(buf2)

		if n1 != n2 {
			return false, nil
		}

		if n1 > 0 {
			if string(buf1[:n1]) != string(buf2[:n2]) {
				return false, nil
			}
		}

		if err1 == io.EOF && err2 == io.EOF {
			break
		}

		if err1 != nil {
			return false, fmt.Errorf("fileutil.CompareFiles: %w", err1)
		}
		if err2 != nil {
			return false, fmt.Errorf("fileutil.CompareFiles: %w", err2)
		}
	}

	return true, nil
}

// CompareHash compares two files by comparing their SHA256 hashes
// CompareHash는 두 파일의 SHA256 해시를 비교하여 파일을 비교합니다
//
// This is faster than byte-by-byte comparison for large files.
// 이는 큰 파일에 대해 바이트 단위 비교보다 빠릅니다.
//
// Example
// 예제:
//
//	same, err := fileutil.CompareHash("file1.txt", "file2.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	if same {
//	    fmt.Println("Files have the same hash")
//	}
func CompareHash(path1, path2 string) (bool, error) {
	hash1, err := SHA256(path1)
	if err != nil {
		return false, err
	}

	hash2, err := SHA256(path2)
	if err != nil {
		return false, err
	}

	return hash1 == hash2, nil
}

// Checksum calculates the SHA256 checksum of a file
// Checksum은 파일의 SHA256 체크섬을 계산합니다
//
// This is an alias for SHA256.
// 이는 SHA256의 별칭입니다.
//
// Example
// 예제:
//
//	checksum, err := fileutil.Checksum("path/to/file.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println("Checksum:", checksum)
func Checksum(path string) (string, error) {
	return SHA256(path)
}

// VerifyChecksum verifies that a file's checksum matches the expected value
// VerifyChecksum은 파일의 체크섬이 예상 값과 일치하는지 확인합니다
//
// Example
// 예제:
//
//	expected := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
//	valid, err := fileutil.VerifyChecksum("path/to/file.txt", expected)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	if !valid {
//	    fmt.Println("Checksum mismatch!")
//	}
func VerifyChecksum(path, expectedChecksum string) (bool, error) {
	actualChecksum, err := Checksum(path)
	if err != nil {
		return false, err
	}

	// Compare checksums (case-insensitive)
	// 체크섬 비교 (대소문자 구분 안함)
	return strings.EqualFold(actualChecksum, expectedChecksum), nil
}

// HashDir calculates the combined hash of all files in a directory
// HashDir는 디렉토리의 모든 파일의 결합된 해시를 계산합니다
//
// This is useful for detecting changes in a directory.
// 이는 디렉토리의 변경 사항을 감지하는 데 유용합니다.
//
// Example
// 예제:
//
//	hash, err := fileutil.HashDir("path/to/directory")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println("Directory hash:", hash)
func HashDir(path string) (string, error) {
	// Check if path is a directory
	// 경로가 디렉토리인지 확인
	if !IsDir(path) {
		return "", fmt.Errorf("fileutil.HashDir: %w", ErrNotDirectory)
	}

	hasher := sha256.New()

	// Walk directory and hash each file
	// 디렉토리 순회 및 각 파일 해시
	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			// Hash filename
			// 파일 이름 해시
			relPath, err := filepath.Rel(path, filePath)
			if err != nil {
				return err
			}
			hasher.Write([]byte(relPath))

			// Hash file contents
			// 파일 내용 해시
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer file.Close()

			if _, err := io.Copy(hasher, file); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return "", fmt.Errorf("fileutil.HashDir: %w", err)
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
