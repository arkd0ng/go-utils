// Package websvrutil provides extreme simplicity web server utilities.
// It reduces 50+ lines of server setup code to just 5 lines.
//
// websvrutil 패키지는 극도로 간단한 웹 서버 유틸리티를 제공합니다.
// 50줄 이상의 서버 설정 코드를 단 5줄로 줄입니다.
//
// Version information is loaded dynamically from cfg/app.yaml.
// 버전 정보는 cfg/app.yaml에서 동적으로 로드됩니다.
package websvrutil

import "github.com/arkd0ng/go-utils/logging"

// Version is the current version of the websvrutil package.
// Version은 websvrutil 패키지의 현재 버전입니다.
var Version = getVersion()

// getVersion loads the application version from cfg/app.yaml.
// getVersion은 cfg/app.yaml에서 애플리케이션 버전을 로드합니다.
func getVersion() string {
	version := logging.TryLoadAppVersion()
	if version == "" {
		return "unknown"
	}
	return version
}
