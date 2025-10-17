# Validation Package Phase 1 Design Document
# 검증 패키지 Phase 1 설계 문서

**Version**: v1.13.016 - v1.13.018
**Date**: 2025-10-17
**Status**: In Development / 개발 중

---

## Overview / 개요

This document describes the design and implementation plan for Phase 1 enhancements to the validation package, focusing on three high-priority validator categories:

1. **Network Validators**: IP addresses, CIDR, MAC addresses
2. **DateTime Validators**: Date/time format validation
3. **Range Validators**: Enhanced range validation for various types

이 문서는 검증 패키지의 Phase 1 개선사항에 대한 설계 및 구현 계획을 설명합니다. 세 가지 높은 우선순위 검증기 카테고리에 초점을 맞춥니다.

---

## 1. Network Validators / 네트워크 검증기

**File**: `validation/rules_network.go`
**Target Version**: v1.13.016
**Estimated LOC**: ~200 lines

### 1.1 IPv4 Validator

**Function Signature**:
```go
func (v *Validator) IPv4() *Validator
```

**Purpose**: Validates that the value is a valid IPv4 address.
**목적**: 값이 유효한 IPv4 주소인지 검증합니다.

**Validation Rules**:
- Must be a string
- Must match IPv4 format: `xxx.xxx.xxx.xxx`
- Each octet must be 0-255
- No leading zeros (except for 0 itself)

**Examples**:
```go
// Valid IPv4
"192.168.1.1"
"10.0.0.1"
"255.255.255.255"
"0.0.0.0"

// Invalid IPv4
"256.1.1.1"      // octet > 255
"192.168.1"      // incomplete
"192.168.1.1.1"  // too many octets
"192.168.01.1"   // leading zero
```

**Implementation Approach**:
- Use `net.ParseIP()` from standard library
- Verify it's IPv4 (not IPv6)
- Additional validation for edge cases

**Error Message**:
```
"{field} must be a valid IPv4 address / {field}은(는) 유효한 IPv4 주소여야 합니다"
```

---

### 1.2 IPv6 Validator

**Function Signature**:
```go
func (v *Validator) IPv6() *Validator
```

**Purpose**: Validates that the value is a valid IPv6 address.
**목적**: 값이 유효한 IPv6 주소인지 검증합니다.

**Validation Rules**:
- Must be a string
- Must match IPv6 format (8 groups of 4 hex digits)
- Supports compressed notation (::)
- Supports mixed notation (IPv4-mapped IPv6)

**Examples**:
```go
// Valid IPv6
"2001:0db8:85a3:0000:0000:8a2e:0370:7334"
"2001:db8:85a3::8a2e:370:7334"  // compressed
"::1"                            // loopback
"fe80::1"                        // link-local
"::ffff:192.0.2.1"              // IPv4-mapped

// Invalid IPv6
"2001:0db8:85a3::8a2e::7334"    // double ::
"gggg::1"                        // invalid hex
```

**Implementation Approach**:
- Use `net.ParseIP()` from standard library
- Verify it's IPv6 (not IPv4)

**Error Message**:
```
"{field} must be a valid IPv6 address / {field}은(는) 유효한 IPv6 주소여야 합니다"
```

---

### 1.3 IP Validator (IPv4 or IPv6)

**Function Signature**:
```go
func (v *Validator) IP() *Validator
```

**Purpose**: Validates that the value is a valid IP address (IPv4 or IPv6).
**목적**: 값이 유효한 IP 주소(IPv4 또는 IPv6)인지 검증합니다.

**Validation Rules**:
- Must be a string
- Must be either valid IPv4 or valid IPv6

**Implementation Approach**:
- Use `net.ParseIP()` which handles both formats

**Error Message**:
```
"{field} must be a valid IP address / {field}은(는) 유효한 IP 주소여야 합니다"
```

---

### 1.4 CIDR Validator

**Function Signature**:
```go
func (v *Validator) CIDR() *Validator
```

**Purpose**: Validates that the value is a valid CIDR notation.
**목적**: 값이 유효한 CIDR 표기법인지 검증합니다.

**Validation Rules**:
- Must be a string
- Format: `<IP>/<prefix>`
- IP must be valid IPv4 or IPv6
- Prefix must be valid for the IP version:
  - IPv4: 0-32
  - IPv6: 0-128

**Examples**:
```go
// Valid CIDR
"192.168.1.0/24"
"10.0.0.0/8"
"2001:db8::/32"
"::1/128"

// Invalid CIDR
"192.168.1.0/33"     // prefix > 32 for IPv4
"192.168.1.0"        // missing prefix
"invalid/24"         // invalid IP
```

**Implementation Approach**:
- Use `net.ParseCIDR()` from standard library

**Error Message**:
```
"{field} must be a valid CIDR notation / {field}은(는) 유효한 CIDR 표기법이어야 합니다"
```

---

### 1.5 MAC Validator

**Function Signature**:
```go
func (v *Validator) MAC() *Validator
```

**Purpose**: Validates that the value is a valid MAC address.
**목적**: 값이 유효한 MAC 주소인지 검증합니다.

**Validation Rules**:
- Must be a string
- Supports multiple formats:
  - Colon-separated: `00:1A:2B:3C:4D:5E`
  - Hyphen-separated: `00-1A-2B-3C-4D-5E`
  - Dot-separated: `001A.2B3C.4D5E` (Cisco format)
- Case-insensitive hex digits
- Must be 6 octets (48 bits)

**Examples**:
```go
// Valid MAC
"00:1A:2B:3C:4D:5E"
"00-1a-2b-3c-4d-5e"
"001A.2B3C.4D5E"
"00:00:00:00:00:00"
"FF:FF:FF:FF:FF:FF"

// Invalid MAC
"00:1A:2B:3C:4D"     // too short
"GG:1A:2B:3C:4D:5E"  // invalid hex
"00:1A:2B:3C:4D:5E:6F" // too long
```

**Implementation Approach**:
- Use `net.ParseMAC()` from standard library

**Error Message**:
```
"{field} must be a valid MAC address / {field}은(는) 유효한 MAC 주소여야 합니다"
```

---

## 2. DateTime Validators / 날짜/시간 검증기

**File**: `validation/rules_datetime.go`
**Target Version**: v1.13.017
**Estimated LOC**: ~180 lines

### 2.1 DateFormat Validator

**Function Signature**:
```go
func (v *Validator) DateFormat(layout string) *Validator
```

**Purpose**: Validates that the string value matches the specified date/time format.
**목적**: 문자열 값이 지정된 날짜/시간 형식과 일치하는지 검증합니다.

**Parameters**:
- `layout`: Go time layout string (e.g., "2006-01-02", "15:04:05")

**Validation Rules**:
- Must be a string
- Must be parseable using `time.Parse()` with the given layout
- Supports custom formats using Go's reference time

**Examples**:
```go
// Date only
v.DateFormat("2006-01-02")  // "2024-12-31"
v.DateFormat("01/02/2006")  // "12/31/2024"
v.DateFormat("02-Jan-2006") // "31-Dec-2024"

// Time only
v.DateFormat("15:04:05")    // "23:59:59"
v.DateFormat("03:04 PM")    // "11:59 PM"

// DateTime
v.DateFormat("2006-01-02 15:04:05")           // "2024-12-31 23:59:59"
v.DateFormat("January 2, 2006 at 3:04pm MST") // "December 31, 2024 at 11:59pm PST"
```

**Implementation Approach**:
- Use `time.Parse(layout, value)` from standard library
- Return error if parsing fails

**Error Message**:
```
"{field} must match date format '{layout}' / {field}은(는) 날짜 형식 '{layout}'과 일치해야 합니다"
```

---

### 2.2 RFC3339 Validator

**Function Signature**:
```go
func (v *Validator) RFC3339() *Validator
```

**Purpose**: Validates that the string value is in RFC3339 format.
**목적**: 문자열 값이 RFC3339 형식인지 검증합니다.

**Validation Rules**:
- Must be a string
- Must match RFC3339 format: `YYYY-MM-DDTHH:MM:SSZ` or with timezone offset
- Format: `time.RFC3339` = "2006-01-02T15:04:05Z07:00"

**Examples**:
```go
// Valid RFC3339
"2024-12-31T23:59:59Z"
"2024-12-31T23:59:59+09:00"
"2024-12-31T23:59:59-05:00"
"2024-01-01T00:00:00Z"

// Invalid RFC3339
"2024-12-31 23:59:59"        // space instead of T
"2024-12-31T23:59:59"        // missing timezone
"12/31/2024T23:59:59Z"       // wrong date format
```

**Implementation Approach**:
- Use `time.Parse(time.RFC3339, value)`

**Error Message**:
```
"{field} must be in RFC3339 format (e.g., '2006-01-02T15:04:05Z07:00') / {field}은(는) RFC3339 형식이어야 합니다"
```

---

### 2.3 DateISO8601 Validator

**Function Signature**:
```go
func (v *Validator) DateISO8601() *Validator
```

**Purpose**: Validates that the string value is in ISO 8601 date format.
**목적**: 문자열 값이 ISO 8601 날짜 형식인지 검증합니다.

**Validation Rules**:
- Must be a string
- Must match ISO 8601 date format: `YYYY-MM-DD`
- No time component (date only)

**Examples**:
```go
// Valid ISO 8601 date
"2024-12-31"
"2024-01-01"
"2000-02-29"  // leap year

// Invalid ISO 8601 date
"2024/12/31"        // wrong separator
"12-31-2024"        // wrong order
"2024-12-31 10:00"  // includes time
"2024-2-9"          // no leading zeros
```

**Implementation Approach**:
- Use regex: `^\d{4}-\d{2}-\d{2}$`
- Then validate with `time.Parse("2006-01-02", value)`

**Error Message**:
```
"{field} must be in ISO 8601 date format (YYYY-MM-DD) / {field}은(는) ISO 8601 날짜 형식(YYYY-MM-DD)이어야 합니다"
```

---

### 2.4 TimeZone Validator

**Function Signature**:
```go
func (v *Validator) TimeZone() *Validator
```

**Purpose**: Validates that the string value is a valid IANA timezone name.
**목적**: 문자열 값이 유효한 IANA 시간대 이름인지 검증합니다.

**Validation Rules**:
- Must be a string
- Must be a valid IANA timezone identifier
- Can be loaded using `time.LoadLocation()`

**Examples**:
```go
// Valid timezones
"America/New_York"
"Europe/London"
"Asia/Seoul"
"Asia/Tokyo"
"UTC"
"GMT"
"Local"

// Invalid timezones
"EST"                // abbreviations not valid
"New York"           // spaces not allowed
"America/InvalidCity"
```

**Implementation Approach**:
- Use `time.LoadLocation(value)` to verify timezone exists

**Error Message**:
```
"{field} must be a valid timezone (e.g., 'America/New_York', 'Asia/Seoul') / {field}은(는) 유효한 시간대여야 합니다"
```

---

## 3. Range Validators / 범위 검증기

**File**: `validation/rules_range.go` (NEW)
**Target Version**: v1.13.018
**Estimated LOC**: ~150 lines

### 3.1 LengthBetween Validator

**Function Signature**:
```go
func (v *Validator) LengthBetween(min, max int) *Validator
```

**Purpose**: Validates that the string length is between min and max (inclusive).
**목적**: 문자열 길이가 min과 max 사이(포함)인지 검증합니다.

**Parameters**:
- `min`: Minimum length (inclusive)
- `max`: Maximum length (inclusive)

**Validation Rules**:
- Must be a string
- Length must be >= min AND <= max
- Uses rune count (not byte count) for Unicode support

**Examples**:
```go
v.LengthBetween(3, 10)

// Valid
"abc"      // length 3
"hello"    // length 5
"1234567890" // length 10

// Invalid
"ab"       // length 2 < 3
"12345678901" // length 11 > 10
""         // length 0 < 3
```

**Implementation Approach**:
- Check type is string
- Use `utf8.RuneCountInString()` for accurate length
- Compare with min and max

**Error Message**:
```
"{field} length must be between {min} and {max} characters / {field} 길이는 {min}자에서 {max}자 사이여야 합니다"
```

---

### 3.2 SizeBetween Validator

**Function Signature**:
```go
func (v *Validator) SizeBetween(min, max int) *Validator
```

**Purpose**: Validates that the collection size is between min and max (inclusive).
**목적**: 컬렉션 크기가 min과 max 사이(포함)인지 검증합니다.

**Parameters**:
- `min`: Minimum size (inclusive)
- `max`: Maximum size (inclusive)

**Validation Rules**:
- Must be a slice, array, or map
- Size must be >= min AND <= max
- Size = number of elements

**Examples**:
```go
v.SizeBetween(2, 5)

// Valid
[]int{1, 2}           // size 2
[]string{"a", "b", "c"} // size 3
map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5} // size 5

// Invalid
[]int{1}              // size 1 < 2
[]int{1,2,3,4,5,6}    // size 6 > 5
[]int{}               // size 0 < 2
```

**Implementation Approach**:
- Use reflection to check type (slice, array, map)
- Get length using `reflect.Value.Len()`
- Compare with min and max

**Error Message**:
```
"{field} size must be between {min} and {max} elements / {field} 크기는 {min}개에서 {max}개 사이여야 합니다"
```

---

### 3.3 DateBetween Validator

**Function Signature**:
```go
func (v *Validator) DateBetween(start, end time.Time) *Validator
```

**Purpose**: Validates that the date value is between start and end (inclusive).
**목적**: 날짜 값이 start와 end 사이(포함)인지 검증합니다.

**Parameters**:
- `start`: Start date (inclusive)
- `end`: End date (inclusive)

**Validation Rules**:
- Must be a `time.Time` value
- Date must be >= start AND <= end
- Compares full timestamp (date and time)

**Examples**:
```go
start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
end := time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)
v.DateBetween(start, end)

// Valid
time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)     // equals start
time.Date(2024, 6, 15, 12, 0, 0, 0, time.UTC)   // in range
time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC) // equals end

// Invalid
time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC) // before start
time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)      // after end
```

**Implementation Approach**:
- Check type is `time.Time`
- Use `time.Before()` and `time.After()` for comparison
- Use `!date.Before(start) && !date.After(end)` for inclusive range

**Error Message**:
```
"{field} must be between {start} and {end} / {field}은(는) {start}와 {end} 사이여야 합니다"
```

---

## Testing Strategy / 테스트 전략

### Test Coverage Requirements / 테스트 커버리지 요구사항
- **100% line coverage** for all new validators
- **모든 새 검증기에 대해 100% 라인 커버리지**

### Test Categories / 테스트 카테고리

#### 1. Unit Tests
- Valid input tests (should pass)
- Invalid input tests (should fail)
- Edge cases (boundary values, empty, nil)
- Type mismatch tests
- StopOnError path tests

#### 2. Benchmark Tests
- Performance measurement for each validator
- Comparison with similar validators

#### 3. Example Tests
- Documentation examples
- Real-world usage scenarios

### Test Files / 테스트 파일
```
validation/
├── rules_network_test.go       (IPv4, IPv6, IP, CIDR, MAC tests)
├── rules_datetime_test.go      (DateFormat, RFC3339, ISO8601, TimeZone tests)
└── rules_range_test.go         (LengthBetween, SizeBetween, DateBetween tests)
```

---

## Documentation Requirements / 문서 요구사항

### 1. Inline Documentation / 인라인 문서
- Godoc comments for each function (English)
- Korean translation as inline comment
- Usage examples in comments
- Parameter descriptions

### 2. User Manual / 사용자 매뉴얼
Update `docs/validation/USER_MANUAL.md`:
- Add new validator sections
- Add usage examples
- Add common pitfalls and solutions

### 3. Examples / 예제
Update `examples/validation/main.go`:
- Demonstrate each new validator
- Show real-world scenarios
- Include error handling

### 4. CHANGELOG / 변경 로그
Create entries for each version:
- v1.13.016: Network validators
- v1.13.017: DateTime validators
- v1.13.018: Range validators

---

## Implementation Order / 구현 순서

### Version v1.13.016 - Network Validators
**Estimated Time**: 1 day / 예상 소요 시간: 1일

1. Create `rules_network.go` with 5 validators
2. Create `rules_network_test.go` with comprehensive tests
3. Add benchmark tests
4. Add example tests
5. Update documentation

### Version v1.13.017 - DateTime Validators
**Estimated Time**: 1 day / 예상 소요 시간: 1일

1. Create `rules_datetime.go` with 4 validators
2. Create `rules_datetime_test.go` with comprehensive tests
3. Add benchmark tests
4. Add example tests
5. Update documentation

### Version v1.13.018 - Range Validators
**Estimated Time**: 0.5 day / 예상 소요 시간: 0.5일

1. Create `rules_range.go` with 3 validators
2. Create `rules_range_test.go` with comprehensive tests
3. Add benchmark tests
4. Add example tests
5. Update documentation

**Total Estimated Time**: 2.5 days / 총 예상 소요 시간: 2.5일

---

## Success Criteria / 성공 기준

### Quality Metrics / 품질 지표
- ✅ 100% test coverage maintained
- ✅ All tests pass
- ✅ All benchmarks complete successfully
- ✅ Zero breaking changes to existing API
- ✅ Bilingual documentation complete

### Performance Criteria / 성능 기준
- Network validators: < 1000 ns/op
- DateTime validators: < 5000 ns/op
- Range validators: < 500 ns/op

### Documentation Criteria / 문서 기준
- English and Korean comments for all functions
- Examples for all validators
- User manual updated
- CHANGELOG updated

---

## Risk Analysis / 위험 분석

### Potential Issues / 잠재적 문제

1. **Timezone Database**: TimeZone validator depends on system timezone database
   - **Mitigation**: Document requirement, provide fallback behavior

2. **Performance**: Network validators may be slower due to parsing
   - **Mitigation**: Benchmark and optimize, use caching if needed

3. **Unicode Handling**: LengthBetween must count runes, not bytes
   - **Mitigation**: Use `utf8.RuneCountInString()`

4. **Date Format Ambiguity**: DateFormat accepts any Go layout
   - **Mitigation**: Document common layouts, provide examples

---

## Appendix / 부록

### Standard Library Dependencies / 표준 라이브러리 종속성
- `net` - IP, CIDR, MAC parsing
- `time` - Date/time parsing and timezone handling
- `reflect` - Type checking for collections
- `unicode/utf8` - Unicode rune counting

### Reference Materials / 참고 자료
- [Go net package](https://pkg.go.dev/net)
- [Go time package](https://pkg.go.dev/time)
- [RFC 3339 - Date and Time on the Internet](https://www.rfc-editor.org/rfc/rfc3339)
- [ISO 8601 - Date and time format](https://en.wikipedia.org/wiki/ISO_8601)
- [IANA Time Zone Database](https://www.iana.org/time-zones)

---

**Document Version**: 1.0
**Last Updated**: 2025-10-17
**Author**: Claude Code
**Status**: Ready for Implementation / 구현 준비 완료
