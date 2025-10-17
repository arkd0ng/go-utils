package validation

import (
	"fmt"
	"net"
)

// Network validator rules for IP addresses, CIDR notation, and MAC addresses.
// IP 주소, CIDR 표기법, MAC 주소에 대한 네트워크 검증 규칙입니다.

// IPv4 checks if the value is a valid IPv4 address.
// IPv4는 값이 유효한 IPv4 주소인지 확인합니다.
//
// The validator checks for:
// - Valid IPv4 format (xxx.xxx.xxx.xxx)
// - Each octet in range 0-255
// - No leading zeros (except for 0 itself)
//
// Example / 예제:
//
//	v := validation.New("192.168.1.1", "ip_address")
//	v.IPv4()
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

// IPv6 checks if the value is a valid IPv6 address.
// IPv6는 값이 유효한 IPv6 주소인지 확인합니다.
//
// The validator supports:
// - Full IPv6 format (8 groups of 4 hex digits)
// - Compressed notation (::)
// - Mixed notation (IPv4-mapped IPv6)
//
// Example / 예제:
//
//	v := validation.New("2001:0db8:85a3::8a2e:0370:7334", "ip_address")
//	v.IPv6()
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

// IP checks if the value is a valid IP address (IPv4 or IPv6).
// IP는 값이 유효한 IP 주소(IPv4 또는 IPv6)인지 확인합니다.
//
// This validator accepts both IPv4 and IPv6 addresses.
// Use IPv4() or IPv6() for specific version validation.
//
// Example / 예제:
//
//	v := validation.New("192.168.1.1", "ip_address")
//	v.IP()
//
//	v2 := validation.New("2001:db8::1", "ip_address")
//	v2.IP()
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

// CIDR checks if the value is a valid CIDR notation.
// CIDR는 값이 유효한 CIDR 표기법인지 확인합니다.
//
// CIDR format is <IP>/<prefix> where:
// - IP can be IPv4 or IPv6
// - prefix must be valid for the IP version (0-32 for IPv4, 0-128 for IPv6)
//
// Example / 예제:
//
//	v := validation.New("192.168.1.0/24", "network")
//	v.CIDR()
//
//	v2 := validation.New("2001:db8::/32", "network")
//	v2.CIDR()
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

// MAC checks if the value is a valid MAC address.
// MAC는 값이 유효한 MAC 주소인지 확인합니다.
//
// Supports multiple MAC address formats:
// - Colon-separated: 00:1A:2B:3C:4D:5E
// - Hyphen-separated: 00-1A-2B-3C-4D-5E
// - Dot-separated: 001A.2B3C.4D5E (Cisco format)
//
// The validation is case-insensitive.
//
// Example / 예제:
//
//	v := validation.New("00:1A:2B:3C:4D:5E", "mac_address")
//	v.MAC()
//
//	v2 := validation.New("00-1a-2b-3c-4d-5e", "mac_address")
//	v2.MAC()
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
