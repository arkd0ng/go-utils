package errorutil

import (
	"errors"
	"reflect"
	"testing"
)

// TestHasCode tests the HasCode function
// TestHasCode는 HasCode 함수를 테스트합니다
func TestHasCode(t *testing.T) {
	tests := []struct {
		name string
		err  error
		code string
		want bool
	}{
		{
			name: "coded error with matching code",
			err:  WithCode("ERR001", "test error"),
			code: "ERR001",
			want: true,
		},
		{
			name: "coded error with non-matching code",
			err:  WithCode("ERR001", "test error"),
			code: "ERR002",
			want: false,
		},
		{
			name: "wrapped coded error",
			err:  Wrap(WithCode("ERR001", "original"), "wrapped"),
			code: "ERR001",
			want: true,
		},
		{
			name: "non-coded error",
			err:  New("plain error"),
			code: "ERR001",
			want: false,
		},
		{
			name: "nil error",
			err:  nil,
			code: "ERR001",
			want: false,
		},
		{
			name: "deeply wrapped coded error",
			err:  Wrap(Wrap(WithCode("ERR001", "original"), "wrap1"), "wrap2"),
			code: "ERR001",
			want: true,
		},
		{
			name: "standard error",
			err:  errors.New("standard error"),
			code: "ERR001",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HasCode(tt.err, tt.code)
			if got != tt.want {
				t.Errorf("HasCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestHasNumericCode tests the HasNumericCode function
// TestHasNumericCode는 HasNumericCode 함수를 테스트합니다
func TestHasNumericCode(t *testing.T) {
	tests := []struct {
		name string
		err  error
		code int
		want bool
	}{
		{
			name: "numeric coded error with matching code",
			err:  WithNumericCode(404, "not found"),
			code: 404,
			want: true,
		},
		{
			name: "numeric coded error with non-matching code",
			err:  WithNumericCode(404, "not found"),
			code: 500,
			want: false,
		},
		{
			name: "wrapped numeric coded error",
			err:  Wrap(WithNumericCode(500, "internal error"), "wrapped"),
			code: 500,
			want: true,
		},
		{
			name: "non-coded error",
			err:  New("plain error"),
			code: 404,
			want: false,
		},
		{
			name: "nil error",
			err:  nil,
			code: 404,
			want: false,
		},
		{
			name: "deeply wrapped numeric coded error",
			err:  Wrap(Wrap(WithNumericCode(503, "service unavailable"), "wrap1"), "wrap2"),
			code: 503,
			want: true,
		},
		{
			name: "zero code",
			err:  WithNumericCode(0, "zero code"),
			code: 0,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HasNumericCode(tt.err, tt.code)
			if got != tt.want {
				t.Errorf("HasNumericCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestGetCode tests the GetCode function
// TestGetCode는 GetCode 함수를 테스트합니다
func TestGetCode(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		wantCode string
		wantOk   bool
	}{
		{
			name:     "coded error",
			err:      WithCode("ERR001", "validation failed"),
			wantCode: "ERR001",
			wantOk:   true,
		},
		{
			name:     "wrapped coded error",
			err:      Wrap(WithCode("ERR002", "database error"), "failed to save"),
			wantCode: "ERR002",
			wantOk:   true,
		},
		{
			name:     "non-coded error",
			err:      New("plain error"),
			wantCode: "",
			wantOk:   false,
		},
		{
			name:     "nil error",
			err:      nil,
			wantCode: "",
			wantOk:   false,
		},
		{
			name:     "deeply wrapped coded error",
			err:      Wrap(Wrap(WithCode("DEEP_ERR", "original"), "wrap1"), "wrap2"),
			wantCode: "DEEP_ERR",
			wantOk:   true,
		},
		{
			name:     "empty code",
			err:      WithCode("", "error with empty code"),
			wantCode: "",
			wantOk:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCode, gotOk := GetCode(tt.err)
			if gotCode != tt.wantCode {
				t.Errorf("GetCode() code = %q, want %q", gotCode, tt.wantCode)
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetCode() ok = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

// TestGetNumericCode tests the GetNumericCode function
// TestGetNumericCode는 GetNumericCode 함수를 테스트합니다
func TestGetNumericCode(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		wantCode int
		wantOk   bool
	}{
		{
			name:     "numeric coded error",
			err:      WithNumericCode(404, "not found"),
			wantCode: 404,
			wantOk:   true,
		},
		{
			name:     "wrapped numeric coded error",
			err:      Wrap(WithNumericCode(500, "internal error"), "failed to process"),
			wantCode: 500,
			wantOk:   true,
		},
		{
			name:     "non-coded error",
			err:      New("plain error"),
			wantCode: 0,
			wantOk:   false,
		},
		{
			name:     "nil error",
			err:      nil,
			wantCode: 0,
			wantOk:   false,
		},
		{
			name:     "deeply wrapped numeric coded error",
			err:      Wrap(Wrap(WithNumericCode(503, "service unavailable"), "wrap1"), "wrap2"),
			wantCode: 503,
			wantOk:   true,
		},
		{
			name:     "zero code",
			err:      WithNumericCode(0, "zero code error"),
			wantCode: 0,
			wantOk:   true,
		},
		{
			name:     "negative code",
			err:      WithNumericCode(-1, "negative code"),
			wantCode: -1,
			wantOk:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCode, gotOk := GetNumericCode(tt.err)
			if gotCode != tt.wantCode {
				t.Errorf("GetNumericCode() code = %d, want %d", gotCode, tt.wantCode)
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetNumericCode() ok = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

// TestGetStackTrace tests the GetStackTrace function
// TestGetStackTrace는 GetStackTrace 함수를 테스트합니다
func TestGetStackTrace(t *testing.T) {
	// Create a stack error for testing
	// 테스트를 위한 스택 에러 생성
	stackErr := &stackError{
		msg:   "test error",
		cause: nil,
		stack: []Frame{
			{File: "test.go", Line: 10, Function: "TestFunc"},
			{File: "main.go", Line: 20, Function: "main"},
		},
	}

	tests := []struct {
		name      string
		err       error
		wantStack []Frame
		wantOk    bool
	}{
		{
			name:      "stack error",
			err:       stackErr,
			wantStack: stackErr.stack,
			wantOk:    true,
		},
		{
			name:      "wrapped stack error",
			err:       Wrap(stackErr, "wrapped"),
			wantStack: stackErr.stack,
			wantOk:    true,
		},
		{
			name:      "non-stack error",
			err:       New("plain error"),
			wantStack: nil,
			wantOk:    false,
		},
		{
			name:      "nil error",
			err:       nil,
			wantStack: nil,
			wantOk:    false,
		},
		{
			name:      "deeply wrapped stack error",
			err:       Wrap(Wrap(stackErr, "wrap1"), "wrap2"),
			wantStack: stackErr.stack,
			wantOk:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStack, gotOk := GetStackTrace(tt.err)
			if !reflect.DeepEqual(gotStack, tt.wantStack) {
				t.Errorf("GetStackTrace() stack = %v, want %v", gotStack, tt.wantStack)
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetStackTrace() ok = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

// TestGetContext tests the GetContext function
// TestGetContext는 GetContext 함수를 테스트합니다
func TestGetContext(t *testing.T) {
	// Create a context error for testing
	// 테스트를 위한 컨텍스트 에러 생성
	ctx := map[string]interface{}{
		"user_id": 123,
		"action":  "login",
	}
	ctxErr := &contextError{
		msg:   "context error",
		cause: nil,
		ctx:   ctx,
	}

	tests := []struct {
		name    string
		err     error
		wantCtx map[string]interface{}
		wantOk  bool
	}{
		{
			name:    "context error",
			err:     ctxErr,
			wantCtx: ctx,
			wantOk:  true,
		},
		{
			name:    "wrapped context error",
			err:     Wrap(ctxErr, "wrapped"),
			wantCtx: ctx,
			wantOk:  true,
		},
		{
			name:    "non-context error",
			err:     New("plain error"),
			wantCtx: nil,
			wantOk:  false,
		},
		{
			name:    "nil error",
			err:     nil,
			wantCtx: nil,
			wantOk:  false,
		},
		{
			name:    "deeply wrapped context error",
			err:     Wrap(Wrap(ctxErr, "wrap1"), "wrap2"),
			wantCtx: ctx,
			wantOk:  true,
		},
		{
			name: "empty context",
			err: &contextError{
				msg:   "empty context",
				cause: nil,
				ctx:   map[string]interface{}{},
			},
			wantCtx: map[string]interface{}{},
			wantOk:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCtx, gotOk := GetContext(tt.err)
			if !reflect.DeepEqual(gotCtx, tt.wantCtx) {
				t.Errorf("GetContext() context = %v, want %v", gotCtx, tt.wantCtx)
			}
			if gotOk != tt.wantOk {
				t.Errorf("GetContext() ok = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

// TestGetContextImmutability tests that GetContext returns a copy
// TestGetContextImmutability는 GetContext가 복사본을 반환하는지 테스트합니다
func TestGetContextImmutability(t *testing.T) {
	ctx := map[string]interface{}{
		"user_id": 123,
	}
	ctxErr := &contextError{
		msg:   "test",
		ctx:   ctx,
		cause: nil,
	}

	// Get context and modify it
	// 컨텍스트를 가져와서 수정
	gotCtx, ok := GetContext(ctxErr)
	if !ok {
		t.Fatal("GetContext() failed")
	}

	gotCtx["user_id"] = 999
	gotCtx["new_key"] = "new_value"

	// Get context again and check it's unchanged
	// 컨텍스트를 다시 가져와서 변경되지 않았는지 확인
	gotCtx2, ok := GetContext(ctxErr)
	if !ok {
		t.Fatal("GetContext() failed")
	}

	if gotCtx2["user_id"] != 123 {
		t.Errorf("Context was modified: user_id = %v, want 123", gotCtx2["user_id"])
	}
	if _, exists := gotCtx2["new_key"]; exists {
		t.Error("Context was modified: new_key should not exist")
	}
}
