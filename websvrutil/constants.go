package websvrutil

import "time"

// Default timeout configurations
// 기본 타임아웃 설정
const (
	// DefaultReadTimeout is the maximum duration for reading the entire request, including the body.
	// DefaultReadTimeout은 본문을 포함한 전체 요청을 읽기 위한 최대 기간입니다.
	DefaultReadTimeout = 15 * time.Second

	// DefaultWriteTimeout is the maximum duration before timing out writes of the response.
	// DefaultWriteTimeout은 응답 쓰기 전 타임아웃까지의 최대 기간입니다.
	DefaultWriteTimeout = 15 * time.Second

	// DefaultIdleTimeout is the maximum amount of time to wait for the next request when keep-alives are enabled.
	// DefaultIdleTimeout은 keep-alive가 활성화되어 있을 때 다음 요청을 기다리는 최대 시간입니다.
	DefaultIdleTimeout = 60 * time.Second
)

// Default size limits
// 기본 크기 제한
const (
	// DefaultMaxHeaderBytes is the maximum number of bytes the server will read parsing the request header.
	// DefaultMaxHeaderBytes는 서버가 요청 헤더를 파싱할 때 읽을 최대 바이트 수입니다.
	DefaultMaxHeaderBytes = 1 << 20 // 1 MB

	// DefaultMaxBodySize is the maximum size of request body for JSON/form data.
	// DefaultMaxBodySize는 JSON/폼 데이터에 대한 요청 본문의 최대 크기입니다.
	// This limit helps prevent DoS attacks via large request bodies.
	// 이 제한은 큰 요청 본문을 통한 DoS 공격을 방지하는 데 도움이 됩니다.
	DefaultMaxBodySize = 10 << 20 // 10 MB

	// DefaultMaxUploadSize is the maximum size for file uploads.
	// DefaultMaxUploadSize는 파일 업로드의 최대 크기입니다.
	DefaultMaxUploadSize = 32 << 20 // 32 MB
)

// Default session configurations
// 기본 세션 설정
const (
	// DefaultSessionMaxAge is the maximum age of a session before it expires.
	// DefaultSessionMaxAge는 세션이 만료되기 전까지의 최대 유효 기간입니다.
	DefaultSessionMaxAge = 24 * time.Hour

	// DefaultSessionCookieName is the default name for the session cookie.
	// DefaultSessionCookieName은 세션 쿠키의 기본 이름입니다.
	DefaultSessionCookieName = "sessionid"

	// DefaultSessionCleanup is the interval at which expired sessions are cleaned up.
	// DefaultSessionCleanup은 만료된 세션이 정리되는 간격입니다.
	DefaultSessionCleanup = 5 * time.Minute
)

// Content-Type constants
// Content-Type 상수
const (
	// ContentTypeJSON represents JSON content type with UTF-8 charset.
	// ContentTypeJSON은 UTF-8 문자셋을 가진 JSON 콘텐츠 타입을 나타냅니다.
	ContentTypeJSON = "application/json; charset=utf-8"

	// ContentTypeHTML represents HTML content type with UTF-8 charset.
	// ContentTypeHTML은 UTF-8 문자셋을 가진 HTML 콘텐츠 타입을 나타냅니다.
	ContentTypeHTML = "text/html; charset=utf-8"

	// ContentTypeXML represents XML content type with UTF-8 charset.
	// ContentTypeXML은 UTF-8 문자셋을 가진 XML 콘텐츠 타입을 나타냅니다.
	ContentTypeXML = "application/xml; charset=utf-8"

	// ContentTypeText represents plain text content type with UTF-8 charset.
	// ContentTypeText는 UTF-8 문자셋을 가진 일반 텍스트 콘텐츠 타입을 나타냅니다.
	ContentTypeText = "text/plain; charset=utf-8"

	// ContentTypeForm represents URL-encoded form content type.
	// ContentTypeForm은 URL 인코딩된 폼 콘텐츠 타입을 나타냅니다.
	ContentTypeForm = "application/x-www-form-urlencoded"

	// ContentTypeMultipart represents multipart form data content type.
	// ContentTypeMultipart는 멀티파트 폼 데이터 콘텐츠 타입을 나타냅니다.
	ContentTypeMultipart = "multipart/form-data"
)

// HTTP header names
// HTTP 헤더 이름
const (
	// HeaderContentType is the Content-Type header name.
	// HeaderContentType은 Content-Type 헤더 이름입니다.
	HeaderContentType = "Content-Type"

	// HeaderAccept is the Accept header name.
	// HeaderAccept는 Accept 헤더 이름입니다.
	HeaderAccept = "Accept"

	// HeaderAuthorization is the Authorization header name.
	// HeaderAuthorization은 Authorization 헤더 이름입니다.
	HeaderAuthorization = "Authorization"

	// HeaderUserAgent is the User-Agent header name.
	// HeaderUserAgent는 User-Agent 헤더 이름입니다.
	HeaderUserAgent = "User-Agent"

	// HeaderXForwardedFor is the X-Forwarded-For header name.
	// HeaderXForwardedFor는 X-Forwarded-For 헤더 이름입니다.
	HeaderXForwardedFor = "X-Forwarded-For"

	// HeaderXRealIP is the X-Real-IP header name.
	// HeaderXRealIP는 X-Real-IP 헤더 이름입니다.
	HeaderXRealIP = "X-Real-IP"
)
