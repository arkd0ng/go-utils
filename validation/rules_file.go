package validation

import (
	"fmt"
	"os"
	"path/filepath"
)

// FilePath checks if the value is a valid file path format.
// FilePath는 값이 유효한 파일 경로 형식인지 확인합니다.
func (v *Validator) FilePath() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("file_path", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Clean the path to normalize it
	cleaned := filepath.Clean(str)

	// Check if path is absolute or relative
	if cleaned == "" || cleaned == "." {
		v.addError("file_path", fmt.Sprintf("%s must be a valid file path / %s은(는) 유효한 파일 경로여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	return v
}

// FileExists checks if the file or directory exists at the given path.
// FileExists는 주어진 경로에 파일이나 디렉토리가 존재하는지 확인합니다.
func (v *Validator) FileExists() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("file_exists", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if _, err := os.Stat(str); os.IsNotExist(err) {
		v.addError("file_exists", fmt.Sprintf("%s must be an existing file or directory / %s은(는) 존재하는 파일 또는 디렉토리여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	return v
}

// FileReadable checks if the file is readable.
// FileReadable은 파일이 읽기 가능한지 확인합니다.
func (v *Validator) FileReadable() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("file_readable", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Try to open the file for reading
	file, err := os.Open(str)
	if err != nil {
		v.addError("file_readable", fmt.Sprintf("%s must be a readable file / %s은(는) 읽기 가능한 파일이어야 합니다", v.fieldName, v.fieldName))
		return v
	}
	defer file.Close()

	return v
}

// FileWritable checks if the file is writable (or its parent directory is writable for new files).
// FileWritable은 파일이 쓰기 가능한지 확인합니다 (새 파일의 경우 부모 디렉토리가 쓰기 가능한지 확인).
func (v *Validator) FileWritable() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("file_writable", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Check if file exists
	info, err := os.Stat(str)
	if err == nil {
		// File exists, check if we can write to it
		file, err := os.OpenFile(str, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			v.addError("file_writable", fmt.Sprintf("%s must be a writable file / %s은(는) 쓰기 가능한 파일이어야 합니다", v.fieldName, v.fieldName))
			return v
		}
		defer file.Close()
	} else if os.IsNotExist(err) {
		// File doesn't exist, check if parent directory is writable
		dir := filepath.Dir(str)
		dirInfo, err := os.Stat(dir)
		if err != nil || !dirInfo.IsDir() {
			v.addError("file_writable", fmt.Sprintf("%s parent directory must exist and be writable / %s의 부모 디렉토리가 존재하고 쓰기 가능해야 합니다", v.fieldName, v.fieldName))
			return v
		}

		// Try to create a temporary file in the directory to test writability
		tmpFile, err := os.CreateTemp(dir, ".write_test_*")
		if err != nil {
			v.addError("file_writable", fmt.Sprintf("%s parent directory must be writable / %s의 부모 디렉토리가 쓰기 가능해야 합니다", v.fieldName, v.fieldName))
			return v
		}
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	} else {
		// Other error (permission, etc.)
		if info != nil && info.IsDir() {
			v.addError("file_writable", fmt.Sprintf("%s must be a file, not a directory / %s은(는) 디렉토리가 아닌 파일이어야 합니다", v.fieldName, v.fieldName))
		} else {
			v.addError("file_writable", fmt.Sprintf("%s must be accessible / %s에 접근할 수 있어야 합니다", v.fieldName, v.fieldName))
		}
		return v
	}

	return v
}

// FileSize checks if the file size is within the specified range (in bytes).
// FileSize는 파일 크기가 지정된 범위 내에 있는지 확인합니다 (바이트 단위).
func (v *Validator) FileSize(min, max int64) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("file_size", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	info, err := os.Stat(str)
	if err != nil {
		v.addError("file_size", fmt.Sprintf("%s must be an existing file / %s은(는) 존재하는 파일이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	if info.IsDir() {
		v.addError("file_size", fmt.Sprintf("%s must be a file, not a directory / %s은(는) 디렉토리가 아닌 파일이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	size := info.Size()
	if size < min || size > max {
		v.addError("file_size", fmt.Sprintf("%s file size must be between %d and %d bytes / %s 파일 크기는 %d와 %d 바이트 사이여야 합니다", v.fieldName, min, max, v.fieldName, min, max))
		return v
	}

	return v
}

// FileExtension checks if the file has one of the specified extensions.
// FileExtension은 파일이 지정된 확장자 중 하나를 가지고 있는지 확인합니다.
func (v *Validator) FileExtension(extensions ...string) *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("file_extension", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	ext := filepath.Ext(str)
	if ext == "" {
		v.addError("file_extension", fmt.Sprintf("%s must have a file extension / %s은(는) 파일 확장자가 있어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Check if extension matches any of the allowed extensions
	found := false
	for _, allowedExt := range extensions {
		// Ensure extension starts with a dot
		if allowedExt != "" && allowedExt[0] != '.' {
			allowedExt = "." + allowedExt
		}
		if ext == allowedExt {
			found = true
			break
		}
	}

	if !found {
		v.addError("file_extension", fmt.Sprintf("%s must have one of the allowed extensions / %s은(는) 허용된 확장자 중 하나여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	return v
}
