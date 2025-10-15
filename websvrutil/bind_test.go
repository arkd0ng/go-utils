package websvrutil

import (
	"bytes"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// TestBindJSON tests binding JSON request body
// TestBindJSON은 JSON 요청 본문 바인딩을 테스트합니다
func TestBindJSON(t *testing.T) {
	type User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Age   int    `json:"age"`
	}

	tests := []struct {
		name    string
		body    string
		want    User
		wantErr bool
	}{
		{
			name: "valid JSON",
			body: `{"name":"John","email":"john@example.com","age":30}`,
			want: User{Name: "John", Email: "john@example.com", Age: 30},
		},
		{
			name:    "invalid JSON",
			body:    `{"name":"John"`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/users", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			ctx := NewContext(w, req)

			var user User
			err := ctx.BindJSON(&user)

			if (err != nil) != tt.wantErr {
				t.Errorf("BindJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && user != tt.want {
				t.Errorf("BindJSON() got = %+v, want %+v", user, tt.want)
			}
		})
	}
}

// TestBindForm tests binding form data
// TestBindForm은 폼 데이터 바인딩을 테스트합니다
func TestBindForm(t *testing.T) {
	type LoginForm struct {
		Username string `form:"username"`
		Password string `form:"password"`
		Remember bool   `form:"remember"`
	}

	tests := []struct {
		name    string
		form    url.Values
		want    LoginForm
		wantErr bool
	}{
		{
			name: "valid form data",
			form: url.Values{
				"username": []string{"john"},
				"password": []string{"secret123"},
				"remember": []string{"true"},
			},
			want: LoginForm{Username: "john", Password: "secret123", Remember: true},
		},
		{
			name: "partial form data",
			form: url.Values{
				"username": []string{"john"},
			},
			want: LoginForm{Username: "john", Password: "", Remember: false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/login", strings.NewReader(tt.form.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			w := httptest.NewRecorder()

			ctx := NewContext(w, req)

			var form LoginForm
			err := ctx.BindForm(&form)

			if (err != nil) != tt.wantErr {
				t.Errorf("BindForm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && form != tt.want {
				t.Errorf("BindForm() got = %+v, want %+v", form, tt.want)
			}
		})
	}
}

// TestBindQuery tests binding query parameters
// TestBindQuery는 쿼리 매개변수 바인딩을 테스트합니다
func TestBindQuery(t *testing.T) {
	type SearchQuery struct {
		Q    string `form:"q"`
		Page int    `form:"page"`
		Size int    `form:"size"`
	}

	tests := []struct {
		name    string
		query   string
		want    SearchQuery
		wantErr bool
	}{
		{
			name:  "valid query params",
			query: "q=golang&page=2&size=10",
			want:  SearchQuery{Q: "golang", Page: 2, Size: 10},
		},
		{
			name:  "partial query params",
			query: "q=golang",
			want:  SearchQuery{Q: "golang", Page: 0, Size: 0},
		},
		{
			name:  "empty query",
			query: "",
			want:  SearchQuery{Q: "", Page: 0, Size: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/search?"+tt.query, nil)
			w := httptest.NewRecorder()

			ctx := NewContext(w, req)

			var query SearchQuery
			err := ctx.BindQuery(&query)

			if (err != nil) != tt.wantErr {
				t.Errorf("BindQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && query != tt.want {
				t.Errorf("BindQuery() got = %+v, want %+v", query, tt.want)
			}
		})
	}
}

// TestBind tests automatic binding based on Content-Type
// TestBind는 Content-Type에 따른 자동 바인딩을 테스트합니다
func TestBind(t *testing.T) {
	type Data struct {
		Name  string `json:"name" form:"name"`
		Value int    `json:"value" form:"value"`
	}

	tests := []struct {
		name        string
		contentType string
		body        string
		want        Data
		wantErr     bool
	}{
		{
			name:        "JSON content type",
			contentType: "application/json",
			body:        `{"name":"test","value":42}`,
			want:        Data{Name: "test", Value: 42},
		},
		{
			name:        "form content type",
			contentType: "application/x-www-form-urlencoded",
			body:        "name=test&value=42",
			want:        Data{Name: "test", Value: 42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/data", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", tt.contentType)
			w := httptest.NewRecorder()

			ctx := NewContext(w, req)

			var data Data
			err := ctx.Bind(&data)

			if (err != nil) != tt.wantErr {
				t.Errorf("Bind() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && data != tt.want {
				t.Errorf("Bind() got = %+v, want %+v", data, tt.want)
			}
		})
	}
}

// TestBindFormData tests the bindFormData helper function
// TestBindFormData는 bindFormData 헬퍼 함수를 테스트합니다
func TestBindFormData(t *testing.T) {
	type TestStruct struct {
		StringField string  `form:"str"`
		IntField    int     `form:"int"`
		FloatField  float64 `form:"float"`
		BoolField   bool    `form:"bool"`
	}

	tests := []struct {
		name    string
		values  url.Values
		want    TestStruct
		wantErr bool
	}{
		{
			name: "all fields",
			values: url.Values{
				"str":   []string{"hello"},
				"int":   []string{"42"},
				"float": []string{"3.14"},
				"bool":  []string{"true"},
			},
			want: TestStruct{
				StringField: "hello",
				IntField:    42,
				FloatField:  3.14,
				BoolField:   true,
			},
		},
		{
			name: "invalid int",
			values: url.Values{
				"int": []string{"not-a-number"},
			},
			wantErr: true,
		},
		{
			name: "invalid float",
			values: url.Values{
				"float": []string{"not-a-float"},
			},
			wantErr: true,
		},
		{
			name: "invalid bool",
			values: url.Values{
				"bool": []string{"not-a-bool"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var obj TestStruct
			err := bindFormData(&obj, tt.values)

			if (err != nil) != tt.wantErr {
				t.Errorf("bindFormData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && obj != tt.want {
				t.Errorf("bindFormData() got = %+v, want %+v", obj, tt.want)
			}
		})
	}
}

// TestBindFormDataError tests error cases for bindFormData
// TestBindFormDataError는 bindFormData의 에러 케이스를 테스트합니다
func TestBindFormDataError(t *testing.T) {
	t.Run("not a pointer", func(t *testing.T) {
		var str string
		err := bindFormData(str, url.Values{})
		if err == nil {
			t.Error("expected error for non-pointer")
		}
	})

	t.Run("not a struct", func(t *testing.T) {
		var str string
		err := bindFormData(&str, url.Values{})
		if err == nil {
			t.Error("expected error for non-struct")
		}
	})
}

// BenchmarkBindJSON benchmarks JSON binding
// BenchmarkBindJSON은 JSON 바인딩을 벤치마크합니다
func BenchmarkBindJSON(b *testing.B) {
	type User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Age   int    `json:"age"`
	}

	body := `{"name":"John","email":"john@example.com","age":30}`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("POST", "/users", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		ctx := NewContext(w, req)

		var user User
		ctx.BindJSON(&user)
	}
}

// BenchmarkBindForm benchmarks form binding
// BenchmarkBindForm은 폼 바인딩을 벤치마크합니다
func BenchmarkBindForm(b *testing.B) {
	type LoginForm struct {
		Username string `form:"username"`
		Password string `form:"password"`
	}

	form := url.Values{
		"username": []string{"john"},
		"password": []string{"secret123"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("POST", "/login", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()
		ctx := NewContext(w, req)

		var loginForm LoginForm
		ctx.BindForm(&loginForm)
	}
}
