package websvrutil

import (
	"testing"
	"time"
)

// TestWithReadTimeout tests the WithReadTimeout option.
// TestWithReadTimeout은 WithReadTimeout 옵션을 테스트합니다.
func TestWithReadTimeout(t *testing.T) {
	customTimeout := 30 * time.Second
	app := New(WithReadTimeout(customTimeout))

	if app.options.ReadTimeout != customTimeout {
		t.Errorf("ReadTimeout = %v, want %v", app.options.ReadTimeout, customTimeout)
	}
}

// TestWithWriteTimeout tests the WithWriteTimeout option.
// TestWithWriteTimeout은 WithWriteTimeout 옵션을 테스트합니다.
func TestWithWriteTimeout(t *testing.T) {
	customTimeout := 45 * time.Second
	app := New(WithWriteTimeout(customTimeout))

	if app.options.WriteTimeout != customTimeout {
		t.Errorf("WriteTimeout = %v, want %v", app.options.WriteTimeout, customTimeout)
	}
}

// TestWithIdleTimeout tests the WithIdleTimeout option.
// TestWithIdleTimeout은 WithIdleTimeout 옵션을 테스트합니다.
func TestWithIdleTimeout(t *testing.T) {
	customTimeout := 120 * time.Second
	app := New(WithIdleTimeout(customTimeout))

	if app.options.IdleTimeout != customTimeout {
		t.Errorf("IdleTimeout = %v, want %v", app.options.IdleTimeout, customTimeout)
	}
}

// TestWithMaxHeaderBytes tests the WithMaxHeaderBytes option.
// TestWithMaxHeaderBytes는 WithMaxHeaderBytes 옵션을 테스트합니다.
func TestWithMaxHeaderBytes(t *testing.T) {
	customSize := 2 << 20 // 2 MB
	app := New(WithMaxHeaderBytes(customSize))

	if app.options.MaxHeaderBytes != customSize {
		t.Errorf("MaxHeaderBytes = %v, want %v", app.options.MaxHeaderBytes, customSize)
	}
}

// TestWithTemplateDir tests the WithTemplateDir option.
// TestWithTemplateDir은 WithTemplateDir 옵션을 테스트합니다.
func TestWithTemplateDir(t *testing.T) {
	customDir := "custom/templates"
	app := New(WithTemplateDir(customDir))

	if app.options.TemplateDir != customDir {
		t.Errorf("TemplateDir = %v, want %v", app.options.TemplateDir, customDir)
	}
}

// TestWithStaticDir tests the WithStaticDir option.
// TestWithStaticDir은 WithStaticDir 옵션을 테스트합니다.
func TestWithStaticDir(t *testing.T) {
	customDir := "custom/static"
	app := New(WithStaticDir(customDir))

	if app.options.StaticDir != customDir {
		t.Errorf("StaticDir = %v, want %v", app.options.StaticDir, customDir)
	}
}

// TestWithStaticPrefix tests the WithStaticPrefix option.
// TestWithStaticPrefix는 WithStaticPrefix 옵션을 테스트합니다.
func TestWithStaticPrefix(t *testing.T) {
	customPrefix := "/assets"
	app := New(WithStaticPrefix(customPrefix))

	if app.options.StaticPrefix != customPrefix {
		t.Errorf("StaticPrefix = %v, want %v", app.options.StaticPrefix, customPrefix)
	}
}

// TestWithAutoReload tests the WithAutoReload option.
// TestWithAutoReload은 WithAutoReload 옵션을 테스트합니다.
func TestWithAutoReload(t *testing.T) {
	tests := []struct {
		name   string
		enable bool
	}{
		{"Enable", true},
		{"Disable", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := New(WithAutoReload(tt.enable))

			if app.options.EnableAutoReload != tt.enable {
				t.Errorf("EnableAutoReload = %v, want %v", app.options.EnableAutoReload, tt.enable)
			}
		})
	}
}

// TestWithLogger tests the WithLogger option.
// TestWithLogger는 WithLogger 옵션을 테스트합니다.
func TestWithLogger(t *testing.T) {
	tests := []struct {
		name   string
		enable bool
	}{
		{"Enable", true},
		{"Disable", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := New(WithLogger(tt.enable))

			if app.options.EnableLogger != tt.enable {
				t.Errorf("EnableLogger = %v, want %v", app.options.EnableLogger, tt.enable)
			}
		})
	}
}

// TestWithRecovery tests the WithRecovery option.
// TestWithRecovery는 WithRecovery 옵션을 테스트합니다.
func TestWithRecovery(t *testing.T) {
	tests := []struct {
		name   string
		enable bool
	}{
		{"Enable", true},
		{"Disable", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := New(WithRecovery(tt.enable))

			if app.options.EnableRecovery != tt.enable {
				t.Errorf("EnableRecovery = %v, want %v", app.options.EnableRecovery, tt.enable)
			}
		})
	}
}

// TestMultipleOptions tests applying multiple options at once.
// TestMultipleOptions는 여러 옵션을 동시에 적용하는 것을 테스트합니다.
func TestMultipleOptions(t *testing.T) {
	customReadTimeout := 20 * time.Second
	customWriteTimeout := 30 * time.Second
	customIdleTimeout := 90 * time.Second
	customMaxHeaderBytes := 3 << 20 // 3 MB
	customTemplateDir := "views"
	customStaticDir := "public"
	customStaticPrefix := "/public"

	app := New(
		WithReadTimeout(customReadTimeout),
		WithWriteTimeout(customWriteTimeout),
		WithIdleTimeout(customIdleTimeout),
		WithMaxHeaderBytes(customMaxHeaderBytes),
		WithTemplateDir(customTemplateDir),
		WithStaticDir(customStaticDir),
		WithStaticPrefix(customStaticPrefix),
		WithAutoReload(true),
		WithLogger(false),
		WithRecovery(false),
	)

	tests := []struct {
		name string
		got  interface{}
		want interface{}
	}{
		{"ReadTimeout", app.options.ReadTimeout, customReadTimeout},
		{"WriteTimeout", app.options.WriteTimeout, customWriteTimeout},
		{"IdleTimeout", app.options.IdleTimeout, customIdleTimeout},
		{"MaxHeaderBytes", app.options.MaxHeaderBytes, customMaxHeaderBytes},
		{"TemplateDir", app.options.TemplateDir, customTemplateDir},
		{"StaticDir", app.options.StaticDir, customStaticDir},
		{"StaticPrefix", app.options.StaticPrefix, customStaticPrefix},
		{"EnableAutoReload", app.options.EnableAutoReload, true},
		{"EnableLogger", app.options.EnableLogger, false},
		{"EnableRecovery", app.options.EnableRecovery, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("%s = %v, want %v", tt.name, tt.got, tt.want)
			}
		})
	}
}

// TestOptionsOverride tests that later options override earlier ones.
// TestOptionsOverride는 나중 옵션이 이전 옵션을 재정의하는지 테스트합니다.
func TestOptionsOverride(t *testing.T) {
	firstTimeout := 10 * time.Second
	secondTimeout := 20 * time.Second

	app := New(
		WithReadTimeout(firstTimeout),
		WithReadTimeout(secondTimeout), // This should override the first one
	)

	if app.options.ReadTimeout != secondTimeout {
		t.Errorf("ReadTimeout = %v, want %v (should be overridden)", app.options.ReadTimeout, secondTimeout)
	}
}

// TestOptionsImmutability tests that modifying options doesn't affect other instances.
// TestOptionsImmutability는 옵션 수정이 다른 인스턴스에 영향을 주지 않는지 테스트합니다.
func TestOptionsImmutability(t *testing.T) {
	app1 := New(WithReadTimeout(10 * time.Second))
	app2 := New(WithReadTimeout(20 * time.Second))

	if app1.options.ReadTimeout == app2.options.ReadTimeout {
		t.Error("options should be independent between instances")
	}

	// Verify each has its own value
	// 각각 자체 값을 가지는지 확인
	if app1.options.ReadTimeout != 10*time.Second {
		t.Errorf("app1.ReadTimeout = %v, want %v", app1.options.ReadTimeout, 10*time.Second)
	}

	if app2.options.ReadTimeout != 20*time.Second {
		t.Errorf("app2.ReadTimeout = %v, want %v", app2.options.ReadTimeout, 20*time.Second)
	}
}

// TestZeroValueOptions tests that zero values can be set.
// TestZeroValueOptions는 제로 값을 설정할 수 있는지 테스트합니다.
func TestZeroValueOptions(t *testing.T) {
	app := New(
		WithReadTimeout(0),
		WithMaxHeaderBytes(0),
		WithTemplateDir(""),
	)

	if app.options.ReadTimeout != 0 {
		t.Errorf("ReadTimeout = %v, want 0", app.options.ReadTimeout)
	}

	if app.options.MaxHeaderBytes != 0 {
		t.Errorf("MaxHeaderBytes = %v, want 0", app.options.MaxHeaderBytes)
	}

	if app.options.TemplateDir != "" {
		t.Errorf("TemplateDir = %v, want empty string", app.options.TemplateDir)
	}
}

// BenchmarkWithReadTimeout benchmarks the WithReadTimeout option.
// BenchmarkWithReadTimeout은 WithReadTimeout 옵션을 벤치마크합니다.
func BenchmarkWithReadTimeout(b *testing.B) {
	timeout := 30 * time.Second

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = New(WithReadTimeout(timeout))
	}
}

// BenchmarkMultipleOptions benchmarks applying multiple options.
// BenchmarkMultipleOptions는 여러 옵션 적용을 벤치마크합니다.
func BenchmarkMultipleOptions(b *testing.B) {
	opts := []Option{
		WithReadTimeout(30 * time.Second),
		WithWriteTimeout(30 * time.Second),
		WithIdleTimeout(90 * time.Second),
		WithTemplateDir("views"),
		WithStaticDir("public"),
		WithLogger(false),
		WithRecovery(false),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = New(opts...)
	}
}
