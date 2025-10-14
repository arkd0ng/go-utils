package stringutil

import (
	"encoding/base64"
	"html"
	"net/url"
)

// Base64Encode encodes a string to base64.
// Base64Encode는 문자열을 base64로 인코딩합니다.
//
// Example / 예제:
//
//	Base64Encode("hello")  // "aGVsbG8="
func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Base64Decode decodes a base64 string.
// Base64Decode는 base64 문자열을 디코딩합니다.
//
// Returns an error if the input is not valid base64.
// 입력이 유효한 base64가 아니면 에러를 반환합니다.
//
// Example / 예제:
//
//	Base64Decode("aGVsbG8=")  // "hello", nil
//	Base64Decode("invalid!")  // "", error
func Base64Decode(s string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

// Base64URLEncode encodes a string to URL-safe base64.
// Base64URLEncode는 문자열을 URL 안전 base64로 인코딩합니다.
//
// URL-safe base64 uses '-' and '_' instead of '+' and '/'.
// URL 안전 base64는 '+'와 '/' 대신 '-'와 '_'를 사용합니다.
//
// Example / 예제:
//
//	Base64URLEncode("hello?world")  // "aGVsbG8_d29ybGQ="
func Base64URLEncode(s string) string {
	return base64.URLEncoding.EncodeToString([]byte(s))
}

// Base64URLDecode decodes a URL-safe base64 string.
// Base64URLDecode는 URL 안전 base64 문자열을 디코딩합니다.
//
// Returns an error if the input is not valid URL-safe base64.
// 입력이 유효한 URL 안전 base64가 아니면 에러를 반환합니다.
//
// Example / 예제:
//
//	Base64URLDecode("aGVsbG8_d29ybGQ=")  // "hello?world", nil
func Base64URLDecode(s string) (string, error) {
	decoded, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

// URLEncode encodes a string for safe use in URLs.
// URLEncode는 URL에서 안전하게 사용하기 위해 문자열을 인코딩합니다.
//
// Example / 예제:
//
//	URLEncode("hello world")  // "hello+world"
//	URLEncode("hello/world")  // "hello%2Fworld"
func URLEncode(s string) string {
	return url.QueryEscape(s)
}

// URLDecode decodes a URL-encoded string.
// URLDecode는 URL 인코딩된 문자열을 디코딩합니다.
//
// Returns an error if the input is not valid URL encoding.
// 입력이 유효한 URL 인코딩이 아니면 에러를 반환합니다.
//
// Example / 예제:
//
//	URLDecode("hello+world")  // "hello world", nil
//	URLDecode("hello%2Fworld")  // "hello/world", nil
func URLDecode(s string) (string, error) {
	return url.QueryUnescape(s)
}

// HTMLEscape escapes special HTML characters.
// HTMLEscape는 특수 HTML 문자를 이스케이프합니다.
//
// Escapes: <, >, &, ", '
// 이스케이프: <, >, &, ", '
//
// Example / 예제:
//
//	HTMLEscape("<div>hello</div>")  // "&lt;div&gt;hello&lt;/div&gt;"
//	HTMLEscape("'hello' & \"world\"")  // "&#39;hello&#39; &amp; &#34;world&#34;"
func HTMLEscape(s string) string {
	return html.EscapeString(s)
}

// HTMLUnescape unescapes HTML entities.
// HTMLUnescape는 HTML 엔티티를 언이스케이프합니다.
//
// Example / 예제:
//
//	HTMLUnescape("&lt;div&gt;hello&lt;/div&gt;")  // "<div>hello</div>"
//	HTMLUnescape("&#39;hello&#39; &amp; &#34;world&#34;")  // "'hello' & \"world\""
func HTMLUnescape(s string) string {
	return html.UnescapeString(s)
}
