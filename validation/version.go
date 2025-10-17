package validation

import "github.com/arkd0ng/go-utils/internal/version"

// Version returns the current package version string from cfg/app.yaml.
// It provides the semantic version number for the validation package.
//
// Version은 cfg/app.yaml에서 현재 패키지 버전 문자열을 반환합니다.
// validation 패키지의 시맨틱 버전 번호를 제공합니다.
//
// Format / 형식:
//   - Semantic versioning: "v1.13.x"
//     시맨틱 버저닝: "v1.13.x"
//
// Usage / 사용법:
//   version := validation.Version
//   fmt.Printf("Validation package version: %s\n", version)
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: Read-only variable
//     스레드 안전: 읽기 전용 변수
var Version = version.Get()
