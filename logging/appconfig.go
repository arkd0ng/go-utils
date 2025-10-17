package logging

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// AppConfig represents the application configuration from app.yaml
// AppConfig는 app.yaml의 애플리케이션 설정을 나타냅니다
type AppConfig struct {
	App struct {
		Name        string `yaml:"name"`
		Version     string `yaml:"version"`
		Description string `yaml:"description"`
	} `yaml:"app"`
}

// LoadAppConfig loads configuration from app.yaml file
// LoadAppConfig는 app.yaml 파일에서 설정을 로드합니다
//
// It searches for app.yaml in the following locations (current dir and parent dirs):
// 다음 위치에서 app.yaml을 검색합니다 (현재 디렉토리 및 상위 디렉토리):
//   1. cfg/app.yaml
//   2. apps/app.yaml
//   3. app.yaml (root directory)
//   4. ../cfg/app.yaml (parent directory)
//   5. ../apps/app.yaml (parent directory)
//   6. ../app.yaml (parent directory)
//
// Returns
// 반환값:
// - *AppConfig: loaded configuration
// 로드된 설정
// - error: error if file not found or parsing failed
// 파일을 찾지 못하거나 파싱 실패 시 에러
func LoadAppConfig() (*AppConfig, error) {
	// Search paths in order
	// 순서대로 검색 경로
	searchPaths := []string{
		"cfg/app.yaml",
		"apps/app.yaml",
		"app.yaml",
		"../cfg/app.yaml",
		"../apps/app.yaml",
		"../app.yaml",
		"../../cfg/app.yaml",
		"../../apps/app.yaml",
		"../../app.yaml",
	}

	var lastErr error
	for _, path := range searchPaths {
		config, err := loadFromPath(path)
		if err == nil {
			return config, nil
		}
		lastErr = err
	}

	// If no file found, return default config
	// 파일을 찾지 못하면 기본 설정 반환
	return &AppConfig{}, fmt.Errorf("app.yaml not found in cfg/, apps/, or root directory: %w", lastErr)
}

// loadFromPath loads configuration from a specific path
// loadFromPath는 특정 경로에서 설정을 로드합니다
func loadFromPath(path string) (*AppConfig, error) {
	// Make path absolute
	// 절대 경로로 변환
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	// Read file
	// 파일 읽기
	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, err
	}

	// Parse YAML
	// YAML 파싱
	var config AppConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse %s: %w", path, err)
	}

	return &config, nil
}

// TryLoadAppVersion attempts to load app version from app.yaml
// TryLoadAppVersion은 app.yaml에서 앱 버전을 로드 시도합니다
//
// Returns the version string if successful, empty string otherwise
// 성공하면 버전 문자열을 반환하고, 실패하면 빈 문자열 반환
func TryLoadAppVersion() string {
	config, err := LoadAppConfig()
	if err != nil {
		return ""
	}
	return config.App.Version
}

// TryLoadAppName attempts to load app name from app.yaml
// TryLoadAppName은 app.yaml에서 앱 이름을 로드 시도합니다
//
// Returns the app name if successful, empty string otherwise
// 성공하면 앱 이름을 반환하고, 실패하면 빈 문자열 반환
func TryLoadAppName() string {
	config, err := LoadAppConfig()
	if err != nil {
		return ""
	}
	return config.App.Name
}
