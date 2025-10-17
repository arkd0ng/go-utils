package validation

import "github.com/arkd0ng/go-utils/internal/version"

// Version returns the current version from cfg/app.yaml.
// Version은 cfg/app.yaml에서 현재 버전을 반환합니다.
var Version = version.Get()
