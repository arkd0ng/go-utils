package logging

// Option is a function that configures a Logger
// Option은 Logger를 설정하는 함수입니다
type Option func(*config) error

// config holds all configuration options for a Logger
// config는 Logger의 모든 설정 옵션을 보유합니다
type config struct {
	// Lumberjack file rotation settings
	// Lumberjack 파일 로테이션 설정
	// Log file path
	// 로그 파일 경로
	filename string
	// Maximum size in megabytes before rotation
	// 로테이션 전 최대 크기(MB)
	maxSize int
	// Maximum number of old log files to retain
	// 보관할 이전 로그 파일의 최대 개수
	maxBackups int
	// Maximum number of days to retain old log files
	// 이전 로그 파일을 보관할 최대 일수
	maxAge int
	// Whether to compress rotated files
	// 로테이션된 파일을 압축할지 여부
	compress bool

	// Logger settings
	// Logger 설정
	// Minimum log level
	// 최소 로그 레벨
	level Level
	// Log prefix
	// 로그 프리픽스
	prefix string
	// Enable color output for console
	// 콘솔 색상 출력 활성화
	enableColor bool
	// Enable stdout output
	// 표준 출력 활성화
	enableStdout bool
	// Enable file output
	// 파일 출력 활성화
	enableFile bool
	// Time format for log entries
	// 로그 항목의 시간 형식
	timeFormat string

	// Banner settings
	// 배너 설정
	// Automatically print banner on logger creation
	// 로거 생성 시 자동으로 배너 출력
	autoBanner bool
	// Application name for banner
	// 배너에 표시할 애플리케이션 이름
	appName string
	// Application version for banner
	// 배너에 표시할 애플리케이션 버전
	appVersion string
}

// defaultConfig returns the default configuration
// defaultConfig는 기본 설정을 반환합니다
func defaultConfig() *config {
	// Try to load version from app.yaml
	// app.yaml에서 버전 로드 시도
	appVersion := TryLoadAppVersion()
	if appVersion == "" {
		// Fallback to default
		// 기본값으로 대체
		appVersion = "v1.0.0"
	}

	// Try to load app name from app.yaml
	// app.yaml에서 앱 이름 로드 시도
	appName := TryLoadAppName()
	if appName == "" {
		// Fallback to default
		// 기본값으로 대체
		appName = "Application"
	}

	return &config{
		filename: "./logs/app.log",
		maxSize:  100, // 100 MB
		// Keep 3 backups
		// 3개 백업 유지
		maxBackups: 3,
		// 28 days
		// 28일
		maxAge: 28,
		// Compress rotated files
		// 로테이션된 파일 압축
		compress:    true,
		level:       INFO,
		prefix:      "",
		enableColor: true,
		// Disabled by default, file only
		// 기본적으로 비활성화, 파일만 출력
		enableStdout: false,
		enableFile:   true,
		timeFormat:   "2006-01-02 15:04:05",
		// Auto banner enabled by default
		// 기본적으로 자동 배너 활성화
		autoBanner: true,
		// From app.yaml or default
		// app.yaml 또는 기본값
		appName: appName,
		// From app.yaml or default
		// app.yaml 또는 기본값
		appVersion: appVersion,
	}
}

// WithFilePath sets the log file path
// WithFilePath는 로그 파일 경로를 설정합니다
//
// Parameters
// 매개변수:
// - path: log file path
// 로그 파일 경로
//
// Example
// 예제:
//
//	logger, _ := logging.New(logging.WithFilePath("./logs/myapp.log"))
func WithFilePath(path string) Option {
	return func(c *config) error {
		c.filename = path
		return nil
	}
}

// WithMaxSize sets the maximum size in megabytes of the log file before it gets rotated
// WithMaxSize는 로그 파일이 로테이션되기 전 최대 크기를 메가바이트로 설정합니다
//
// Parameters
// 매개변수:
// - mb: maximum size in megabytes
// 메가바이트 단위 최대 크기
//
// Example
// 예제:
//
//	logger, _ := logging.New(logging.WithMaxSize(50)) // 50 MB
func WithMaxSize(mb int) Option {
	return func(c *config) error {
		c.maxSize = mb
		return nil
	}
}

// WithMaxBackups sets the maximum number of old log files to retain
// WithMaxBackups는 보관할 이전 로그 파일의 최대 개수를 설정합니다
//
// Parameters
// 매개변수:
// - n: maximum number of backups
// 백업의 최대 개수
//
// Example
// 예제:
//
//	logger, _ := logging.New(logging.WithMaxBackups(5))
func WithMaxBackups(n int) Option {
	return func(c *config) error {
		c.maxBackups = n
		return nil
	}
}

// WithMaxAge sets the maximum number of days to retain old log files
// WithMaxAge는 이전 로그 파일을 보관할 최대 일수를 설정합니다
//
// Parameters
// 매개변수:
// - days: maximum age in days
// 일 단위 최대 보관 기간
//
// Example
// 예제:
//
// logger, _ := logging.New(logging.WithMaxAge(7)) // Keep for 7 days
// 7일 동안 보관
func WithMaxAge(days int) Option {
	return func(c *config) error {
		c.maxAge = days
		return nil
	}
}

// WithCompress sets whether to compress rotated log files
// WithCompress는 로테이션된 로그 파일을 압축할지 여부를 설정합니다
//
// Parameters
// 매개변수:
// - compress: true to enable compression
// 압축을 활성화하려면 true
//
// Example
// 예제:
//
//	logger, _ := logging.New(logging.WithCompress(false))
func WithCompress(compress bool) Option {
	return func(c *config) error {
		c.compress = compress
		return nil
	}
}

// WithLevel sets the minimum log level
// WithLevel은 최소 로그 레벨을 설정합니다
//
// Parameters
// 매개변수:
// - level: minimum log level
// 최소 로그 레벨
//
// Example
// 예제:
//
//	logger, _ := logging.New(logging.WithLevel(logging.DEBUG))
func WithLevel(level Level) Option {
	return func(c *config) error {
		c.level = level
		return nil
	}
}

// WithPrefix sets the log prefix
// WithPrefix는 로그 프리픽스를 설정합니다
//
// Parameters
// 매개변수:
// - prefix: log prefix
// 로그 프리픽스
//
// Example
// 예제:
//
//	logger, _ := logging.New(logging.WithPrefix("[APP]"))
func WithPrefix(prefix string) Option {
	return func(c *config) error {
		c.prefix = prefix
		return nil
	}
}

// WithColor enables or disables colored output for console
// WithColor는 콘솔의 색상 출력을 활성화하거나 비활성화합니다
//
// Parameters
// 매개변수:
// - enable: true to enable color
// 색상을 활성화하려면 true
//
// Example
// 예제:
//
//	logger, _ := logging.New(logging.WithColor(false))
func WithColor(enable bool) Option {
	return func(c *config) error {
		c.enableColor = enable
		return nil
	}
}

// WithStdout enables or disables stdout output
// WithStdout은 표준 출력을 활성화하거나 비활성화합니다
//
// Parameters
// 매개변수:
// - enable: true to enable stdout
// 표준 출력을 활성화하려면 true
//
// Example
// 예제:
//
//	logger, _ := logging.New(logging.WithStdout(false))
func WithStdout(enable bool) Option {
	return func(c *config) error {
		c.enableStdout = enable
		return nil
	}
}

// WithFile enables or disables file output
// WithFile은 파일 출력을 활성화하거나 비활성화합니다
//
// Parameters
// 매개변수:
// - enable: true to enable file output
// 파일 출력을 활성화하려면 true
//
// Example
// 예제:
//
//	logger, _ := logging.New(logging.WithFile(false))
func WithFile(enable bool) Option {
	return func(c *config) error {
		c.enableFile = enable
		return nil
	}
}

// WithFileOnly disables stdout and enables only file output
// WithFileOnly는 표준 출력을 비활성화하고 파일 출력만 활성화합니다
//
// Example
// 예제:
//
//	logger, _ := logging.New(logging.WithFileOnly())
func WithFileOnly() Option {
	return func(c *config) error {
		c.enableStdout = false
		c.enableFile = true
		return nil
	}
}

// WithStdoutOnly disables file output and enables only stdout
// WithStdoutOnly는 파일 출력을 비활성화하고 표준 출력만 활성화합니다
//
// Example
// 예제:
//
//	logger, _ := logging.New(logging.WithStdoutOnly())
func WithStdoutOnly() Option {
	return func(c *config) error {
		c.enableStdout = true
		c.enableFile = false
		return nil
	}
}

// WithTimeFormat sets the time format for log entries
// WithTimeFormat은 로그 항목의 시간 형식을 설정합니다
//
// Parameters
// 매개변수:
// - format: time format string (Go time format)
// 시간 형식 문자열 (Go 시간 형식)
//
// Example
// 예제:
//
//	logger, _ := logging.New(logging.WithTimeFormat("2006/01/02 15:04:05"))
func WithTimeFormat(format string) Option {
	return func(c *config) error {
		c.timeFormat = format
		return nil
	}
}

// WithAutoBanner enables or disables automatic banner printing on logger creation
// WithAutoBanner는 로거 생성 시 자동 배너 출력을 활성화하거나 비활성화합니다
//
// Parameters
// 매개변수:
// - enable: true to enable auto banner
// 자동 배너를 활성화하려면 true
//
// Example
// 예제:
//
// logger, _ := logging.New(logging.WithAutoBanner(false)) // Disable auto banner
// 자동 배너 비활성화
func WithAutoBanner(enable bool) Option {
	return func(c *config) error {
		c.autoBanner = enable
		return nil
	}
}

// WithAppName sets the application name for the banner
// WithAppName은 배너에 표시할 애플리케이션 이름을 설정합니다
//
// Parameters
// 매개변수:
// - name: application name
// 애플리케이션 이름
//
// Example
// 예제:
//
//	logger, _ := logging.New(logging.WithAppName("MyApp"))
func WithAppName(name string) Option {
	return func(c *config) error {
		c.appName = name
		return nil
	}
}

// WithAppVersion sets the application version for the banner
// WithAppVersion은 배너에 표시할 애플리케이션 버전을 설정합니다
//
// Parameters
// 매개변수:
// - version: application version
// 애플리케이션 버전
//
// Example
// 예제:
//
//	logger, _ := logging.New(logging.WithAppVersion("v2.0.0"))
func WithAppVersion(version string) Option {
	return func(c *config) error {
		c.appVersion = version
		return nil
	}
}

// WithBanner is a convenience function to set app name, version, and enable auto banner
// WithBanner는 앱 이름, 버전을 설정하고 자동 배너를 활성화하는 편의 함수입니다
//
// Parameters
// 매개변수:
// - name: application name
// 애플리케이션 이름
// - version: application version
// 애플리케이션 버전
//
// Example
// 예제:
//
//	logger, _ := logging.New(logging.WithBanner("MyApp", "v2.0.0"))
func WithBanner(name, version string) Option {
	return func(c *config) error {
		c.autoBanner = true
		c.appName = name
		c.appVersion = version
		return nil
	}
}
