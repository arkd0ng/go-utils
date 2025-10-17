package version

import (
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

// Config represents the app configuration structure.
// Config는 앱 설정 구조를 나타냅니다.
type Config struct {
	App struct {
		Name        string `yaml:"name"`
		Version     string `yaml:"version"`
		Description string `yaml:"description"`
	} `yaml:"app"`
}

var cachedVersion string

// Get returns the version from cfg/app.yaml.
// Get은 cfg/app.yaml에서 버전을 반환합니다.
func Get() string {
	if cachedVersion != "" {
		return cachedVersion
	}

	// Get the project root directory
	_, filename, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(filename), "../..")

	configPath := filepath.Join(projectRoot, "cfg", "app.yaml")

	data, err := os.ReadFile(configPath)
	if err != nil {
		// Fallback to default if can't read file
		return "v0.0.0-dev"
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return "v0.0.0-dev"
	}

	cachedVersion = config.App.Version
	return cachedVersion
}
