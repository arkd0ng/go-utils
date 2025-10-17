package validation

import (
	"fmt"
	"net"
)

// Network validator rules for IP addresses, CIDR notation, and MAC addresses.
// IP 주소, CIDR 표기법, MAC 주소에 대한 네트워크 검증 규칙입니다.

// IPv4 validates that the value is a valid IPv4 address.
// Uses net.ParseIP with IPv4-specific validation.
//
// IPv4는 값이 유효한 IPv4 주소인지 검증합니다.
// IPv4 전용 검증과 함께 net.ParseIP를 사용합니다.
//
// Parameters / 매개변수:
//   - None (operates on validator's value)
//     없음 (validator의 값에 대해 작동)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses net.ParseIP for format validation
//     형식 검증을 위해 net.ParseIP 사용
//   - Validates IPv4 format (xxx.xxx.xxx.xxx)
//     IPv4 형식 검증 (xxx.xxx.xxx.xxx)
//   - Each octet must be 0-255
//     각 옥텟은 0-255 범위
//   - No leading zeros except for 0 itself
//     0 자체를 제외한 선행 0 없음
//   - Rejects IPv6 addresses
//     IPv6 주소 거부
//   - Fails if value is not string
//     값이 문자열이 아니면 실패
//
// Use Cases / 사용 사례:
//   - Server IP validation / 서버 IP 검증
//   - Network configuration / 네트워크 구성
//   - Firewall rules / 방화벽 규칙
//   - IPv4-only environments / IPv4 전용 환경
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n), n = string length
//     시간 복잡도: O(n), n = 문자열 길이
//   - Single net.ParseIP call
//     단일 net.ParseIP 호출
//
// The validator checks for / 검증기 확인 사항:
// - Valid IPv4 format (xxx.xxx.xxx.xxx)
// - Each octet in range 0-255
// - No leading zeros (except for 0 itself)
//
// Example / 예제:
//
//	// Valid IPv4 addresses / 유효한 IPv4 주소
//	v := validation.New("192.168.1.1", "ip_address")
//	v.IPv4()  // Passes
//
//	v = validation.New("10.0.0.1", "server_ip")
//	v.IPv4()  // Passes
//
//	v = validation.New("255.255.255.255", "broadcast")
//	v.IPv4()  // Passes
//
//	// Invalid addresses / 무효한 주소
//	v = validation.New("256.1.1.1", "ip")
//	v.IPv4()  // Fails (octet > 255)
//
//	v = validation.New("2001:db8::1", "ip")
//	v.IPv4()  // Fails (IPv6 address)
func (v *Validator) IPv4() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("ipv4", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	ip := net.ParseIP(str)
	if ip == nil {
		v.addError("ipv4", fmt.Sprintf("%s must be a valid IPv4 address / %s은(는) 유효한 IPv4 주소여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Check if it's IPv4 (not IPv6)
	if ip.To4() == nil {
		v.addError("ipv4", fmt.Sprintf("%s must be a valid IPv4 address / %s은(는) 유효한 IPv4 주소여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	return v
}

// IPv6 validates that the value is a valid IPv6 address.
// Uses net.ParseIP with IPv6-specific validation.
//
// IPv6는 값이 유효한 IPv6 주소인지 검증합니다.
// IPv6 전용 검증과 함께 net.ParseIP를 사용합니다.
//
// Parameters / 매개변수:
//   - None (operates on validator's value)
//     없음 (validator의 값에 대해 작동)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses net.ParseIP for format validation
//     형식 검증을 위해 net.ParseIP 사용
//   - Supports full IPv6 format (8 groups of 4 hex digits)
//     전체 IPv6 형식 지원 (4자리 16진수 8그룹)
//   - Supports compressed notation (::)
//     압축 표기법 지원 (::)
//   - Supports mixed notation (IPv4-mapped IPv6)
//     혼합 표기법 지원 (IPv4 매핑 IPv6)
//   - Rejects IPv4 addresses
//     IPv4 주소 거부
//   - Fails if value is not string
//     값이 문자열이 아니면 실패
//
// Use Cases / 사용 사례:
//   - Modern network configuration / 최신 네트워크 구성
//   - IPv6-only environments / IPv6 전용 환경
//   - Dual-stack validation / 듀얼 스택 검증
//   - Future-proof addressing / 미래 대비 주소 지정
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n), n = string length
//     시간 복잡도: O(n), n = 문자열 길이
//   - Single net.ParseIP call
//     단일 net.ParseIP 호출
//
// The validator supports / 검증기 지원:
// - Full IPv6 format (8 groups of 4 hex digits)
// - Compressed notation (::)
// - Mixed notation (IPv4-mapped IPv6)
//
// Example / 예제:
//
//	// Full IPv6 format / 전체 IPv6 형식
//	v := validation.New("2001:0db8:85a3:0000:0000:8a2e:0370:7334", "ip")
//	v.IPv6()  // Passes
//
//	// Compressed notation / 압축 표기법
//	v = validation.New("2001:db8:85a3::8a2e:370:7334", "ip")
//	v.IPv6()  // Passes
//
//	v = validation.New("::1", "loopback")
//	v.IPv6()  // Passes (IPv6 loopback)
//
//	// IPv4-mapped IPv6 / IPv4 매핑 IPv6
//	v = validation.New("::ffff:192.0.2.1", "ip")
//	v.IPv6()  // Passes
//
//	// Invalid addresses / 무효한 주소
//	v = validation.New("192.168.1.1", "ip")
//	v.IPv6()  // Fails (IPv4 address)
func (v *Validator) IPv6() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("ipv6", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	ip := net.ParseIP(str)
	if ip == nil {
		v.addError("ipv6", fmt.Sprintf("%s must be a valid IPv6 address / %s은(는) 유효한 IPv6 주소여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	// Check if it's IPv6 (To4() returns nil for IPv6)
	if ip.To4() != nil {
		v.addError("ipv6", fmt.Sprintf("%s must be a valid IPv6 address / %s은(는) 유효한 IPv6 주소여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	return v
}

// IP validates that the value is a valid IP address (IPv4 or IPv6).
// Accepts both IPv4 and IPv6 formats. Use IPv4() or IPv6() for version-specific validation.
//
// IP는 값이 유효한 IP 주소(IPv4 또는 IPv6)인지 검증합니다.
// IPv4 및 IPv6 형식 모두 허용합니다. 버전별 검증은 IPv4() 또는 IPv6()를 사용하세요.
//
// Parameters / 매개변수:
//   - None (operates on validator's value)
//     없음 (validator의 값에 대해 작동)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses net.ParseIP for format validation
//     형식 검증을 위해 net.ParseIP 사용
//   - Accepts both IPv4 and IPv6 addresses
//     IPv4 및 IPv6 주소 모두 허용
//   - Validates format without version restriction
//     버전 제한 없이 형식 검증
//   - Fails if value is not string
//     값이 문자열이 아니면 실패
//
// Use Cases / 사용 사례:
//   - Version-agnostic IP validation / 버전 무관 IP 검증
//   - Dual-stack environments / 듀얼 스택 환경
//   - Generic IP configuration / 일반 IP 구성
//   - Protocol-independent validation / 프로토콜 독립적 검증
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n), n = string length
//     시간 복잡도: O(n), n = 문자열 길이
//   - Single net.ParseIP call
//     단일 net.ParseIP 호출
//
// This validator accepts both IPv4 and IPv6 addresses.
// Use IPv4() or IPv6() for specific version validation.
//
// Example / 예제:
//
//	// IPv4 addresses / IPv4 주소
//	v := validation.New("192.168.1.1", "ip_address")
//	v.IP()  // Passes
//
//	v = validation.New("10.0.0.1", "server_ip")
//	v.IP()  // Passes
//
//	// IPv6 addresses / IPv6 주소
//	v = validation.New("2001:db8::1", "ip_address")
//	v.IP()  // Passes
//
//	v = validation.New("::1", "loopback")
//	v.IP()  // Passes
//
//	// Invalid addresses / 무효한 주소
//	v = validation.New("999.999.999.999", "ip")
//	v.IP()  // Fails
//
//	v = validation.New("not-an-ip", "ip")
//	v.IP()  // Fails
func (v *Validator) IP() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("ip", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	ip := net.ParseIP(str)
	if ip == nil {
		v.addError("ip", fmt.Sprintf("%s must be a valid IP address / %s은(는) 유효한 IP 주소여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	return v
}

// CIDR validates that the value is a valid CIDR notation.
// CIDR format is <IP>/<prefix> for network address representation.
//
// CIDR는 값이 유효한 CIDR 표기법인지 검증합니다.
// CIDR 형식은 네트워크 주소 표현을 위한 <IP>/<접두사>입니다.
//
// Parameters / 매개변수:
//   - None (operates on validator's value)
//     없음 (validator의 값에 대해 작동)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses net.ParseCIDR for validation
//     검증을 위해 net.ParseCIDR 사용
//   - Accepts IPv4 CIDR (0-32 prefix)
//     IPv4 CIDR 허용 (0-32 접두사)
//   - Accepts IPv6 CIDR (0-128 prefix)
//     IPv6 CIDR 허용 (0-128 접두사)
//   - Validates both IP and prefix
//     IP와 접두사 모두 검증
//   - Fails if format is invalid
//     형식이 무효하면 실패
//   - Fails if value is not string
//     값이 문자열이 아니면 실패
//
// Use Cases / 사용 사례:
//   - Network configuration / 네트워크 구성
//   - Subnet validation / 서브넷 검증
//   - Firewall rule definition / 방화벽 규칙 정의
//   - IP range specification / IP 범위 지정
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n), n = string length
//     시간 복잡도: O(n), n = 문자열 길이
//   - Single net.ParseCIDR call
//     단일 net.ParseCIDR 호출
//
// CIDR format is <IP>/<prefix> where / CIDR 형식은 다음과 같습니다:
// - IP can be IPv4 or IPv6
// - prefix must be valid for the IP version (0-32 for IPv4, 0-128 for IPv6)
//
// Example / 예제:
//
//	// IPv4 CIDR / IPv4 CIDR
//	v := validation.New("192.168.1.0/24", "network")
//	v.CIDR()  // Passes
//
//	v = validation.New("10.0.0.0/8", "private_network")
//	v.CIDR()  // Passes
//
//	v = validation.New("172.16.0.0/12", "subnet")
//	v.CIDR()  // Passes
//
//	// IPv6 CIDR / IPv6 CIDR
//	v = validation.New("2001:db8::/32", "network")
//	v.CIDR()  // Passes
//
//	v = validation.New("fe80::/10", "link_local")
//	v.CIDR()  // Passes
//
//	// Invalid CIDR / 무효한 CIDR
//	v = validation.New("192.168.1.0/33", "network")
//	v.CIDR()  // Fails (prefix > 32 for IPv4)
//
//	v = validation.New("192.168.1.1", "ip")
//	v.CIDR()  // Fails (missing prefix)
func (v *Validator) CIDR() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("cidr", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	_, _, err := net.ParseCIDR(str)
	if err != nil {
		v.addError("cidr", fmt.Sprintf("%s must be a valid CIDR notation (e.g., '192.168.1.0/24') / %s은(는) 유효한 CIDR 표기법이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	return v
}

// MAC validates that the value is a valid MAC address.
// Supports multiple MAC address formats with case-insensitive validation.
//
// MAC는 값이 유효한 MAC 주소인지 검증합니다.
// 대소문자 구분 없는 검증으로 여러 MAC 주소 형식을 지원합니다.
//
// Parameters / 매개변수:
//   - None (operates on validator's value)
//     없음 (validator의 값에 대해 작동)
//
// Returns / 반환:
//   - *Validator: Returns self for method chaining
//     메서드 체이닝을 위해 자신을 반환
//
// Behavior / 동작:
//   - Uses net.ParseMAC for validation
//     검증을 위해 net.ParseMAC 사용
//   - Supports colon-separated format (00:1A:2B:3C:4D:5E)
//     콜론 구분 형식 지원 (00:1A:2B:3C:4D:5E)
//   - Supports hyphen-separated format (00-1A-2B-3C-4D-5E)
//     하이픈 구분 형식 지원 (00-1A-2B-3C-4D-5E)
//   - Supports dot-separated format (001A.2B3C.4D5E - Cisco)
//     점 구분 형식 지원 (001A.2B3C.4D5E - Cisco)
//   - Case-insensitive validation
//     대소문자 구분 없는 검증
//   - Fails if value is not string
//     값이 문자열이 아니면 실패
//
// Use Cases / 사용 사례:
//   - Network device identification / 네트워크 장치 식별
//   - Hardware address validation / 하드웨어 주소 검증
//   - Network configuration / 네트워크 구성
//   - Device registration / 장치 등록
//
// Thread Safety / 스레드 안전성:
//   - Thread-safe: No shared state / 스레드 안전: 공유 상태 없음
//
// Performance / 성능:
//   - Time complexity: O(n), n = string length
//     시간 복잡도: O(n), n = 문자열 길이
//   - Single net.ParseMAC call
//     단일 net.ParseMAC 호출
//
// Supports multiple MAC address formats / 여러 MAC 주소 형식 지원:
// - Colon-separated: 00:1A:2B:3C:4D:5E
// - Hyphen-separated: 00-1A-2B-3C-4D-5E
// - Dot-separated: 001A.2B3C.4D5E (Cisco format)
//
// The validation is case-insensitive.
//
// Example / 예제:
//
//	// Colon-separated (common) / 콜론 구분 (일반)
//	v := validation.New("00:1A:2B:3C:4D:5E", "mac_address")
//	v.MAC()  // Passes
//
//	// Hyphen-separated / 하이픈 구분
//	v = validation.New("00-1a-2b-3c-4d-5e", "mac_address")
//	v.MAC()  // Passes (case-insensitive)
//
//	// Cisco format (dot-separated) / Cisco 형식 (점 구분)
//	v = validation.New("001A.2B3C.4D5E", "mac_address")
//	v.MAC()  // Passes
//
//	// Case variations / 대소문자 변형
//	v = validation.New("AA:BB:CC:DD:EE:FF", "mac")
//	v.MAC()  // Passes
//
//	// Invalid MAC / 무효한 MAC
//	v = validation.New("GG:HH:II:JJ:KK:LL", "mac")
//	v.MAC()  // Fails (invalid hex)
//
//	v = validation.New("00:1A:2B:3C:4D", "mac")
//	v.MAC()  // Fails (too short)
func (v *Validator) MAC() *Validator {
	if v.stopOnError && len(v.errors) > 0 {
		return v
	}

	str, ok := v.value.(string)
	if !ok {
		v.addError("mac", fmt.Sprintf("%s must be a string / %s은(는) 문자열이어야 합니다", v.fieldName, v.fieldName))
		return v
	}

	_, err := net.ParseMAC(str)
	if err != nil {
		v.addError("mac", fmt.Sprintf("%s must be a valid MAC address (e.g., '00:1A:2B:3C:4D:5E') / %s은(는) 유효한 MAC 주소여야 합니다", v.fieldName, v.fieldName))
		return v
	}

	return v
}
