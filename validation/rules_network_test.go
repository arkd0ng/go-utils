package validation

import (
	"testing"
)

// Tests for network validators
// 네트워크 검증기에 대한 테스트

// TestIPv4 tests the IPv4 validator
// TestIPv4는 IPv4 검증기를 테스트합니다
func TestIPv4(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid IPv4 addresses
		{"valid basic", "192.168.1.1", false},
		{"valid loopback", "127.0.0.1", false},
		{"valid zeros", "0.0.0.0", false},
		{"valid broadcast", "255.255.255.255", false},
		{"valid private", "10.0.0.1", false},
		{"valid class B", "172.16.0.1", false},

		// Invalid IPv4 addresses
		{"invalid octet too large", "256.1.1.1", true},
		{"invalid incomplete", "192.168.1", true},
		{"invalid too many octets", "192.168.1.1.1", true},
		{"invalid IPv6", "2001:db8::1", true},
		{"invalid non-numeric", "abc.def.ghi.jkl", true},
		{"invalid empty", "", true},
		{"invalid not string", 12345, true},
		{"invalid nil", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "ip_address")
			v.IPv4()
			err := v.Validate()

			if tt.wantError && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tt.wantError && err != nil {
				t.Errorf("expected no error but got %v", err)
			}
		})
	}
}

// TestIPv6 tests the IPv6 validator
// TestIPv6는 IPv6 검증기를 테스트합니다
func TestIPv6(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid IPv6 addresses
		{"valid full", "2001:0db8:85a3:0000:0000:8a2e:0370:7334", false},
		{"valid compressed", "2001:db8:85a3::8a2e:370:7334", false},
		{"valid loopback", "::1", false},
		{"valid link-local", "fe80::1", false},
		{"valid all zeros", "::", false},
		{"valid multicast", "ff02::1", false},

		// Invalid IPv6 addresses
		{"invalid double colon twice", "2001:0db8:85a3::8a2e::7334", true},
		{"invalid hex", "gggg::1", true},
		{"invalid IPv4", "192.168.1.1", true},
		{"invalid incomplete", "2001:db8", true},
		{"invalid empty", "", true},
		{"invalid not string", 12345, true},
		{"invalid nil", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "ip_address")
			v.IPv6()
			err := v.Validate()

			if tt.wantError && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tt.wantError && err != nil {
				t.Errorf("expected no error but got %v", err)
			}
		})
	}
}

// TestIP tests the IP validator (IPv4 or IPv6)
// TestIP는 IP 검증기를 테스트합니다 (IPv4 또는 IPv6)
func TestIP(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid IP addresses (IPv4 and IPv6)
		{"valid IPv4", "192.168.1.1", false},
		{"valid IPv6", "2001:db8::1", false},
		{"valid IPv4 loopback", "127.0.0.1", false},
		{"valid IPv6 loopback", "::1", false},
		{"valid IPv4 zeros", "0.0.0.0", false},
		{"valid IPv6 zeros", "::", false},

		// Invalid IP addresses
		{"invalid format", "not-an-ip", true},
		{"invalid empty", "", true},
		{"invalid incomplete IPv4", "192.168.1", true},
		{"invalid incomplete IPv6", "2001:db8", true},
		{"invalid not string", 12345, true},
		{"invalid nil", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "ip_address")
			v.IP()
			err := v.Validate()

			if tt.wantError && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tt.wantError && err != nil {
				t.Errorf("expected no error but got %v", err)
			}
		})
	}
}

// TestCIDR tests the CIDR validator
// TestCIDR는 CIDR 검증기를 테스트합니다
func TestCIDR(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid CIDR notation
		{"valid IPv4 /24", "192.168.1.0/24", false},
		{"valid IPv4 /32", "192.168.1.1/32", false},
		{"valid IPv4 /0", "0.0.0.0/0", false},
		{"valid IPv4 /8", "10.0.0.0/8", false},
		{"valid IPv6 /32", "2001:db8::/32", false},
		{"valid IPv6 /128", "::1/128", false},
		{"valid IPv6 /64", "fe80::/64", false},

		// Invalid CIDR notation
		{"invalid no prefix", "192.168.1.0", true},
		{"invalid prefix too large IPv4", "192.168.1.0/33", true},
		{"invalid prefix too large IPv6", "2001:db8::/129", true},
		{"invalid IP", "invalid/24", true},
		{"invalid negative prefix", "192.168.1.0/-1", true},
		{"invalid empty", "", true},
		{"invalid not string", 12345, true},
		{"invalid nil", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "network")
			v.CIDR()
			err := v.Validate()

			if tt.wantError && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tt.wantError && err != nil {
				t.Errorf("expected no error but got %v", err)
			}
		})
	}
}

// TestMAC tests the MAC validator
// TestMAC는 MAC 검증기를 테스트합니다
func TestMAC(t *testing.T) {
	tests := []struct {
		name      string
		value     interface{}
		wantError bool
	}{
		// Valid MAC addresses
		{"valid colon uppercase", "00:1A:2B:3C:4D:5E", false},
		{"valid colon lowercase", "00:1a:2b:3c:4d:5e", false},
		{"valid hyphen", "00-1A-2B-3C-4D-5E", false},
		{"valid dot", "001A.2B3C.4D5E", false},
		{"valid all zeros", "00:00:00:00:00:00", false},
		{"valid all Fs", "FF:FF:FF:FF:FF:FF", false},

		// Invalid MAC addresses
		{"invalid too short", "00:1A:2B:3C:4D", true},
		{"invalid too long", "00:1A:2B:3C:4D:5E:6F", true},
		{"invalid hex chars", "GG:1A:2B:3C:4D:5E", true},
		{"invalid format", "00:1A:2B:3C:4D:5", true},
		{"invalid empty", "", true},
		{"invalid not string", 12345, true},
		{"invalid nil", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New(tt.value, "mac_address")
			v.MAC()
			err := v.Validate()

			if tt.wantError && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tt.wantError && err != nil {
				t.Errorf("expected no error but got %v", err)
			}
		})
	}
}

// TestIPv4StopOnError tests StopOnError with IPv4 validator
// TestIPv4StopOnError는 IPv4 검증기의 StopOnError를 테스트합니다
func TestIPv4StopOnError(t *testing.T) {
	v := New("invalid", "ip")
	v.StopOnError()
	v.IPv4()     // First error
	v.IPv4()     // Should hit stopOnError return
	err := v.Validate()

	if err == nil {
		t.Error("expected error from IPv4")
	}

	verrs := err.(ValidationErrors)
	if len(verrs) != 1 {
		t.Errorf("expected 1 error with StopOnError, got %d", len(verrs))
	}
}

// TestIPv6StopOnError tests StopOnError with IPv6 validator
// TestIPv6StopOnError는 IPv6 검증기의 StopOnError를 테스트합니다
func TestIPv6StopOnError(t *testing.T) {
	v := New("invalid", "ip")
	v.StopOnError()
	v.IPv6()     // First error
	v.IPv6()     // Should hit stopOnError return
	err := v.Validate()

	if err == nil {
		t.Error("expected error from IPv6")
	}

	verrs := err.(ValidationErrors)
	if len(verrs) != 1 {
		t.Errorf("expected 1 error with StopOnError, got %d", len(verrs))
	}
}

// TestIPStopOnError tests StopOnError with IP validator
// TestIPStopOnError는 IP 검증기의 StopOnError를 테스트합니다
func TestIPStopOnError(t *testing.T) {
	v := New("invalid", "ip")
	v.StopOnError()
	v.IP()     // First error
	v.IP()     // Should hit stopOnError return
	err := v.Validate()

	if err == nil {
		t.Error("expected error from IP")
	}

	verrs := err.(ValidationErrors)
	if len(verrs) != 1 {
		t.Errorf("expected 1 error with StopOnError, got %d", len(verrs))
	}
}

// TestCIDRStopOnError tests StopOnError with CIDR validator
// TestCIDRStopOnError는 CIDR 검증기의 StopOnError를 테스트합니다
func TestCIDRStopOnError(t *testing.T) {
	v := New("invalid", "network")
	v.StopOnError()
	v.CIDR()     // First error
	v.CIDR()     // Should hit stopOnError return
	err := v.Validate()

	if err == nil {
		t.Error("expected error from CIDR")
	}

	verrs := err.(ValidationErrors)
	if len(verrs) != 1 {
		t.Errorf("expected 1 error with StopOnError, got %d", len(verrs))
	}
}

// TestMACStopOnError tests StopOnError with MAC validator
// TestMACStopOnError는 MAC 검증기의 StopOnError를 테스트합니다
func TestMACStopOnError(t *testing.T) {
	v := New("invalid", "mac")
	v.StopOnError()
	v.MAC()     // First error
	v.MAC()     // Should hit stopOnError return
	err := v.Validate()

	if err == nil {
		t.Error("expected error from MAC")
	}

	verrs := err.(ValidationErrors)
	if len(verrs) != 1 {
		t.Errorf("expected 1 error with StopOnError, got %d", len(verrs))
	}
}

// TestNetworkValidatorsWithNonStringType tests network validators with non-string types
// TestNetworkValidatorsWithNonStringType는 문자열이 아닌 타입으로 네트워크 검증기를 테스트합니다
func TestNetworkValidatorsWithNonStringType(t *testing.T) {
	testValues := []interface{}{
		123,
		45.67,
		true,
		[]string{"test"},
		map[string]int{"test": 1},
	}

	for _, val := range testValues {
		// Test IPv4
		v1 := New(val, "field")
		v1.IPv4()
		if err := v1.Validate(); err == nil {
			t.Errorf("IPv4 should fail for non-string type %T", val)
		}

		// Test IPv6
		v2 := New(val, "field")
		v2.IPv6()
		if err := v2.Validate(); err == nil {
			t.Errorf("IPv6 should fail for non-string type %T", val)
		}

		// Test IP
		v3 := New(val, "field")
		v3.IP()
		if err := v3.Validate(); err == nil {
			t.Errorf("IP should fail for non-string type %T", val)
		}

		// Test CIDR
		v4 := New(val, "field")
		v4.CIDR()
		if err := v4.Validate(); err == nil {
			t.Errorf("CIDR should fail for non-string type %T", val)
		}

		// Test MAC
		v5 := New(val, "field")
		v5.MAC()
		if err := v5.Validate(); err == nil {
			t.Errorf("MAC should fail for non-string type %T", val)
		}
	}
}
