package validation

import (
	"os"
	"testing"
	"time"
)

// Benchmark tests for validation package
// validation 패키지의 벤치마크 테스트

// BenchmarkRequired benchmarks the Required validator
// BenchmarkRequired는 Required 검증기를 벤치마크합니다
func BenchmarkRequired(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("test value", "field")
		v.Required()
		_ = v.Validate()
	}
}

// BenchmarkMinLength benchmarks the MinLength validator
// BenchmarkMinLength는 MinLength 검증기를 벤치마크합니다
func BenchmarkMinLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("test string value", "field")
		v.MinLength(5)
		_ = v.Validate()
	}
}

// BenchmarkMaxLength benchmarks the MaxLength validator
// BenchmarkMaxLength는 MaxLength 검증기를 벤치마크합니다
func BenchmarkMaxLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("test", "field")
		v.MaxLength(10)
		_ = v.Validate()
	}
}

// BenchmarkEmail benchmarks the Email validator
// BenchmarkEmail는 Email 검증기를 벤치마크합니다
func BenchmarkEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("test@example.com", "email")
		v.Email()
		_ = v.Validate()
	}
}

// BenchmarkURL benchmarks the URL validator
// BenchmarkURL는 URL 검증기를 벤치마크합니다
func BenchmarkURL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("https://example.com", "url")
		v.URL()
		_ = v.Validate()
	}
}

// BenchmarkMin benchmarks the Min validator
// BenchmarkMin는 Min 검증기를 벤치마크합니다
func BenchmarkMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New(100, "value")
		v.Min(0)
		_ = v.Validate()
	}
}

// BenchmarkMax benchmarks the Max validator
// BenchmarkMax는 Max 검증기를 벤치마크합니다
func BenchmarkMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New(50, "value")
		v.Max(100)
		_ = v.Validate()
	}
}

// BenchmarkBetween benchmarks the Min and Max validators together
// BenchmarkBetween는 Min과 Max 검증기를 함께 벤치마크합니다
func BenchmarkBetween(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New(50, "value")
		v.Min(0).Max(100)
		_ = v.Validate()
	}
}

// BenchmarkIn benchmarks the In validator
// BenchmarkIn는 In 검증기를 벤치마크합니다
func BenchmarkIn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("apple", "fruit")
		v.In("apple", "banana", "orange", "grape", "melon")
		_ = v.Validate()
	}
}

// BenchmarkNotIn benchmarks the NotIn validator
// BenchmarkNotIn는 NotIn 검증기를 벤치마크합니다
func BenchmarkNotIn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("pear", "fruit")
		v.NotIn("apple", "banana", "orange")
		_ = v.Validate()
	}
}

// BenchmarkArrayLength benchmarks the ArrayLength validator
// BenchmarkArrayLength는 ArrayLength 검증기를 벤치마크합니다
func BenchmarkArrayLength(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		v := New(arr, "numbers")
		v.ArrayLength(5)
		_ = v.Validate()
	}
}

// BenchmarkArrayUnique benchmarks the ArrayUnique validator
// BenchmarkArrayUnique는 ArrayUnique 검증기를 벤치마크합니다
func BenchmarkArrayUnique(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		v := New(arr, "numbers")
		v.ArrayUnique()
		_ = v.Validate()
	}
}

// BenchmarkMapHasKey benchmarks the MapHasKey validator
// BenchmarkMapHasKey는 MapHasKey 검증기를 벤치마크합니다
func BenchmarkMapHasKey(b *testing.B) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for i := 0; i < b.N; i++ {
		v := New(m, "data")
		v.MapHasKey("a")
		_ = v.Validate()
	}
}

// BenchmarkMapHasKeys benchmarks the MapHasKeys validator
// BenchmarkMapHasKeys는 MapHasKeys 검증기를 벤치마크합니다
func BenchmarkMapHasKeys(b *testing.B) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	for i := 0; i < b.N; i++ {
		v := New(m, "data")
		v.MapHasKeys("a", "b", "c")
		_ = v.Validate()
	}
}

// BenchmarkEquals benchmarks the Equals validator
// BenchmarkEquals는 Equals 검증기를 벤치마크합니다
func BenchmarkEquals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New(100, "value")
		v.Equals(100)
		_ = v.Validate()
	}
}

// BenchmarkBefore benchmarks the Before validator
// BenchmarkBefore는 Before 검증기를 벤치마크합니다
func BenchmarkBefore(b *testing.B) {
	now := time.Now()
	yesterday := now.Add(-24 * time.Hour)
	for i := 0; i < b.N; i++ {
		v := New(yesterday, "date")
		v.Before(now)
		_ = v.Validate()
	}
}

// BenchmarkAfter benchmarks the After validator
// BenchmarkAfter는 After 검증기를 벤치마크합니다
func BenchmarkAfter(b *testing.B) {
	now := time.Now()
	tomorrow := now.Add(24 * time.Hour)
	for i := 0; i < b.N; i++ {
		v := New(tomorrow, "date")
		v.After(now)
		_ = v.Validate()
	}
}

// BenchmarkCustom benchmarks the Custom validator
// BenchmarkCustom는 Custom 검증기를 벤치마크합니다
func BenchmarkCustom(b *testing.B) {
	customFunc := func(val interface{}) bool {
		str, ok := val.(string)
		return ok && len(str) > 0
	}
	for i := 0; i < b.N; i++ {
		v := New("test", "field")
		v.Custom(customFunc, "custom error")
		_ = v.Validate()
	}
}

// BenchmarkMultipleValidators benchmarks chaining multiple validators
// BenchmarkMultipleValidators는 여러 검증기 체이닝을 벤치마크합니다
func BenchmarkMultipleValidators(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("test@example.com", "email")
		v.Required().MinLength(5).MaxLength(50).Email()
		_ = v.Validate()
	}
}

// BenchmarkStopOnError benchmarks validation with StopOnError
// BenchmarkStopOnError는 StopOnError를 사용한 검증을 벤치마크합니다
func BenchmarkStopOnError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("", "field")
		v.StopOnError()
		v.Required().MinLength(5).MaxLength(50)
		_ = v.Validate()
	}
}

// BenchmarkMultiValidator benchmarks MultiValidator with multiple fields
// BenchmarkMultiValidator는 여러 필드를 가진 MultiValidator를 벤치마크합니다
func BenchmarkMultiValidator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mv := NewValidator()
		mv.Field("john@example.com", "email").Required().Email()
		mv.Field("John Doe", "name").Required().MinLength(2)
		mv.Field(25, "age").Required().Min(18)
		_ = mv.Validate()
	}
}

// BenchmarkValidationErrors benchmarks error collection
// BenchmarkValidationErrors는 에러 수집을 벤치마크합니다
func BenchmarkValidationErrors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("", "field")
		v.Required()
		v.MinLength(10)
		v.Email()
		err := v.Validate()
		if err != nil {
			verrs := err.(ValidationErrors)
			_ = verrs.Error()
		}
	}
}
// BenchmarkIPv4 benchmarks the IPv4 validator
// BenchmarkIPv4는 IPv4 검증기를 벤치마크합니다
func BenchmarkIPv4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("192.168.1.1", "ip_address")
		v.IPv4()
		_ = v.Validate()
	}
}

// BenchmarkIPv6 benchmarks the IPv6 validator
// BenchmarkIPv6는 IPv6 검증기를 벤치마크합니다
func BenchmarkIPv6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("2001:0db8:85a3:0000:0000:8a2e:0370:7334", "ip_address")
		v.IPv6()
		_ = v.Validate()
	}
}

// BenchmarkIP benchmarks the IP validator
// BenchmarkIP는 IP 검증기를 벤치마크합니다
func BenchmarkIP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("192.168.1.1", "ip_address")
		v.IP()
		_ = v.Validate()
	}
}

// BenchmarkCIDR benchmarks the CIDR validator
// BenchmarkCIDR는 CIDR 검증기를 벤치마크합니다
func BenchmarkCIDR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("192.168.1.0/24", "network")
		v.CIDR()
		_ = v.Validate()
	}
}

// BenchmarkMAC benchmarks the MAC validator
// BenchmarkMAC는 MAC 검증기를 벤치마크합니다
func BenchmarkMAC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("00:1A:2B:3C:4D:5E", "mac_address")
		v.MAC()
		_ = v.Validate()
	}
}

// BenchmarkDateFormat benchmarks the DateFormat validator
// BenchmarkDateFormat는 DateFormat 검증기를 벤치마크합니다
func BenchmarkDateFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("2025-10-17", "date")
		v.DateFormat("2006-01-02")
		_ = v.Validate()
	}
}

// BenchmarkTimeFormat benchmarks the TimeFormat validator
// BenchmarkTimeFormat는 TimeFormat 검증기를 벤치마크합니다
func BenchmarkTimeFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("14:30:00", "time")
		v.TimeFormat("15:04:05")
		_ = v.Validate()
	}
}

// BenchmarkDateBefore benchmarks the DateBefore validator
// BenchmarkDateBefore는 DateBefore 검증기를 벤치마크합니다
func BenchmarkDateBefore(b *testing.B) {
	baseDate := time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC)
	testDate := time.Date(2025, 10, 17, 12, 0, 0, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		v := New(testDate, "date")
		v.DateBefore(baseDate)
		_ = v.Validate()
	}
}

// BenchmarkDateAfter benchmarks the DateAfter validator
// BenchmarkDateAfter는 DateAfter 검증기를 벤치마크합니다
func BenchmarkDateAfter(b *testing.B) {
	baseDate := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	testDate := time.Date(2025, 10, 17, 12, 0, 0, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		v := New(testDate, "date")
		v.DateAfter(baseDate)
		_ = v.Validate()
	}
}

// BenchmarkIntRange benchmarks the IntRange validator
// BenchmarkIntRange는 IntRange 검증기를 벤치마크합니다
func BenchmarkIntRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New(25, "age")
		v.IntRange(18, 65)
		_ = v.Validate()
	}
}

// BenchmarkFloatRange benchmarks the FloatRange validator
// BenchmarkFloatRange는 FloatRange 검증기를 벤치마크합니다
func BenchmarkFloatRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New(98.6, "temperature")
		v.FloatRange(95.0, 105.0)
		_ = v.Validate()
	}
}

// BenchmarkDateRange benchmarks the DateRange validator
// BenchmarkDateRange는 DateRange 검증기를 벤치마크합니다
func BenchmarkDateRange(b *testing.B) {
	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC)
	testDate := time.Date(2025, 6, 15, 12, 0, 0, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		v := New(testDate, "date")
		v.DateRange(start, end)
		_ = v.Validate()
	}
}

// BenchmarkUUIDv4 benchmarks the UUIDv4 validator
// BenchmarkUUIDv4는 UUIDv4 검증기를 벤치마크합니다
func BenchmarkUUIDv4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("550e8400-e29b-41d4-a716-446655440000", "uuid")
		v.UUIDv4()
		_ = v.Validate()
	}
}

// BenchmarkXML benchmarks the XML validator
// BenchmarkXML는 XML 검증기를 벤치마크합니다
func BenchmarkXML(b *testing.B) {
	xmlData := `<?xml version="1.0"?><root><child>content</child></root>`
	for i := 0; i < b.N; i++ {
		v := New(xmlData, "xml")
		v.XML()
		_ = v.Validate()
	}
}

// BenchmarkHex benchmarks the Hex validator
// BenchmarkHex는 Hex 검증기를 벤치마크합니다
func BenchmarkHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("0xdeadbeef", "hex")
		v.Hex()
		_ = v.Validate()
	}
}

// BenchmarkFilePath benchmarks the FilePath validator
// BenchmarkFilePath는 FilePath 검증기를 벤치마크합니다
func BenchmarkFilePath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("/usr/bin/test", "file_path")
		v.FilePath()
		_ = v.Validate()
	}
}

// BenchmarkFileExists benchmarks the FileExists validator
// BenchmarkFileExists는 FileExists 검증기를 벤치마크합니다
func BenchmarkFileExists(b *testing.B) {
	// Create a temporary file for benchmarking
	tmpFile, _ := os.CreateTemp("", "bench_file_*.txt")
	tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := New(tmpFile.Name(), "file_path")
		v.FileExists()
		_ = v.Validate()
	}
}

// BenchmarkFileReadable benchmarks the FileReadable validator
// BenchmarkFileReadable는 FileReadable 검증기를 벤치마크합니다
func BenchmarkFileReadable(b *testing.B) {
	// Create a temporary file for benchmarking
	tmpFile, _ := os.CreateTemp("", "bench_file_*.txt")
	tmpFile.WriteString("test content")
	tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := New(tmpFile.Name(), "file_path")
		v.FileReadable()
		_ = v.Validate()
	}
}

// BenchmarkFileSize benchmarks the FileSize validator
// BenchmarkFileSize는 FileSize 검증기를 벤치마크합니다
func BenchmarkFileSize(b *testing.B) {
	// Create a temporary file for benchmarking
	tmpFile, _ := os.CreateTemp("", "bench_file_*.txt")
	tmpFile.WriteString("Hello, World!")
	tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := New(tmpFile.Name(), "file_path")
		v.FileSize(0, 100)
		_ = v.Validate()
	}
}

// BenchmarkFileExtension benchmarks the FileExtension validator
// BenchmarkFileExtension는 FileExtension 검증기를 벤치마크합니다
func BenchmarkFileExtension(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("test.txt", "file_path")
		v.FileExtension(".txt", ".md")
		_ = v.Validate()
	}
}

// BenchmarkCreditCard benchmarks the CreditCard validator
// BenchmarkCreditCard는 CreditCard 검증기를 벤치마크합니다
func BenchmarkCreditCard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("4532015112830366", "card")
		v.CreditCard()
		_ = v.Validate()
	}
}

// BenchmarkCreditCardType benchmarks the CreditCardType validator
// BenchmarkCreditCardType는 CreditCardType 검증기를 벤치마크합니다
func BenchmarkCreditCardType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("4532015112830366", "card")
		v.CreditCardType("visa")
		_ = v.Validate()
	}
}

// BenchmarkLuhn benchmarks the Luhn validator
// BenchmarkLuhn는 Luhn 검증기를 벤치마크합니다
func BenchmarkLuhn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("4532015112830366", "luhn")
		v.Luhn()
		_ = v.Validate()
	}
}

// BenchmarkISBN benchmarks the ISBN validator
// BenchmarkISBN는 ISBN 검증기를 벤치마크합니다
func BenchmarkISBN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("978-0-596-52068-7", "isbn")
		v.ISBN()
		_ = v.Validate()
	}
}

// BenchmarkISSN benchmarks the ISSN validator
// BenchmarkISSN는 ISSN 검증기를 벤치마크합니다
func BenchmarkISSN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("2049-3630", "issn")
		v.ISSN()
		_ = v.Validate()
	}
}

// BenchmarkEAN benchmarks the EAN validator
// BenchmarkEAN는 EAN 검증기를 벤치마크합니다
func BenchmarkEAN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("4006381333931", "ean")
		v.EAN()
		_ = v.Validate()
	}
}

// BenchmarkLatitude benchmarks the Latitude validator
// BenchmarkLatitude는 Latitude 검증기를 벤치마크합니다
func BenchmarkLatitude(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New(37.5665, "latitude")
		v.Latitude()
		_ = v.Validate()
	}
}

// BenchmarkLongitude benchmarks the Longitude validator
// BenchmarkLongitude는 Longitude 검증기를 벤치마크합니다
func BenchmarkLongitude(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New(126.9780, "longitude")
		v.Longitude()
		_ = v.Validate()
	}
}

// BenchmarkCoordinate benchmarks the Coordinate validator
// BenchmarkCoordinate는 Coordinate 검증기를 벤치마크합니다
func BenchmarkCoordinate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("37.5665,126.9780", "coordinate")
		v.Coordinate()
		_ = v.Validate()
	}
}

// BenchmarkJWT benchmarks the JWT validator
// BenchmarkJWT는 JWT 검증기를 벤치마크합니다
func BenchmarkJWT(b *testing.B) {
	jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIn0.dozjgNryP4J3jVmNHl0w5N_XgL0n3I9PlFUP0THsR8U"
	for i := 0; i < b.N; i++ {
		v := New(jwt, "token")
		v.JWT()
		_ = v.Validate()
	}
}

// BenchmarkBCrypt benchmarks the BCrypt validator
// BenchmarkBCrypt는 BCrypt 검증기를 벤치마크합니다
func BenchmarkBCrypt(b *testing.B) {
	hash := "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy"
	for i := 0; i < b.N; i++ {
		v := New(hash, "password")
		v.BCrypt()
		_ = v.Validate()
	}
}

// BenchmarkMD5 benchmarks the MD5 validator
// BenchmarkMD5는 MD5 검증기를 벤치마크합니다
func BenchmarkMD5(b *testing.B) {
	hash := "5d41402abc4b2a76b9719d911017c592"
	for i := 0; i < b.N; i++ {
		v := New(hash, "hash")
		v.MD5()
		_ = v.Validate()
	}
}

// BenchmarkSHA1 benchmarks the SHA1 validator
// BenchmarkSHA1는 SHA1 검증기를 벤치마크합니다
func BenchmarkSHA1(b *testing.B) {
	hash := "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"
	for i := 0; i < b.N; i++ {
		v := New(hash, "hash")
		v.SHA1()
		_ = v.Validate()
	}
}

// BenchmarkSHA256 benchmarks the SHA256 validator
// BenchmarkSHA256는 SHA256 검증기를 벤치마크합니다
func BenchmarkSHA256(b *testing.B) {
	hash := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	for i := 0; i < b.N; i++ {
		v := New(hash, "hash")
		v.SHA256()
		_ = v.Validate()
	}
}

// BenchmarkSHA512 benchmarks the SHA512 validator
// BenchmarkSHA512는 SHA512 검증기를 벤치마크합니다
func BenchmarkSHA512(b *testing.B) {
	hash := "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"
	for i := 0; i < b.N; i++ {
		v := New(hash, "hash")
		v.SHA512()
		_ = v.Validate()
	}
}

// BenchmarkHexColor benchmarks the HexColor validator
// BenchmarkHexColor는 HexColor 검증기를 벤치마크합니다
func BenchmarkHexColor(b *testing.B) {
	color := "#FF5733"
	for i := 0; i < b.N; i++ {
		v := New(color, "color")
		v.HexColor()
		_ = v.Validate()
	}
}

// BenchmarkRGB benchmarks the RGB validator
// BenchmarkRGB는 RGB 검증기를 벤치마크합니다
func BenchmarkRGB(b *testing.B) {
	color := "rgb(255, 87, 51)"
	for i := 0; i < b.N; i++ {
		v := New(color, "color")
		v.RGB()
		_ = v.Validate()
	}
}

// BenchmarkRGBA benchmarks the RGBA validator
// BenchmarkRGBA는 RGBA 검증기를 벤치마크합니다
func BenchmarkRGBA(b *testing.B) {
	color := "rgba(255, 87, 51, 0.8)"
	for i := 0; i < b.N; i++ {
		v := New(color, "color")
		v.RGBA()
		_ = v.Validate()
	}
}

// BenchmarkHSL benchmarks the HSL validator
// BenchmarkHSL는 HSL 검증기를 벤치마크합니다
func BenchmarkHSL(b *testing.B) {
	color := "hsl(9, 100%, 60%)"
	for i := 0; i < b.N; i++ {
		v := New(color, "color")
		v.HSL()
		_ = v.Validate()
	}
}

// BenchmarkASCII benchmarks the ASCII validator
// BenchmarkASCII는 ASCII 검증기를 벤치마크합니다
func BenchmarkASCII(b *testing.B) {
	text := "Hello World 123"
	for i := 0; i < b.N; i++ {
		v := New(text, "text")
		v.ASCII()
		_ = v.Validate()
	}
}

// BenchmarkPrintable benchmarks the Printable validator
// BenchmarkPrintable는 Printable 검증기를 벤치마크합니다
func BenchmarkPrintable(b *testing.B) {
	text := "Hello World! 123"
	for i := 0; i < b.N; i++ {
		v := New(text, "text")
		v.Printable()
		_ = v.Validate()
	}
}

// BenchmarkWhitespace benchmarks the Whitespace validator
// BenchmarkWhitespace는 Whitespace 검증기를 벤치마크합니다
func BenchmarkWhitespace(b *testing.B) {
	text := " \t\n  "
	for i := 0; i < b.N; i++ {
		v := New(text, "text")
		v.Whitespace()
		_ = v.Validate()
	}
}

// BenchmarkAlphaSpace benchmarks the AlphaSpace validator
// BenchmarkAlphaSpace는 AlphaSpace 검증기를 벤치마크합니다
func BenchmarkAlphaSpace(b *testing.B) {
	text := "John Doe"
	for i := 0; i < b.N; i++ {
		v := New(text, "text")
		v.AlphaSpace()
		_ = v.Validate()
	}
}

// BenchmarkOneOf benchmarks the OneOf validator
// BenchmarkOneOf는 OneOf 검증기를 벤치마크합니다
func BenchmarkOneOf(b *testing.B) {
	status := "active"
	for i := 0; i < b.N; i++ {
		v := New(status, "status")
		v.OneOf("active", "inactive", "pending")
		_ = v.Validate()
	}
}

// BenchmarkNotOneOf benchmarks the NotOneOf validator
// BenchmarkNotOneOf는 NotOneOf 검증기를 벤치마크합니다
func BenchmarkNotOneOf(b *testing.B) {
	username := "user123"
	for i := 0; i < b.N; i++ {
		v := New(username, "username")
		v.NotOneOf("admin", "root", "administrator")
		_ = v.Validate()
	}
}

// BenchmarkWhen benchmarks the When validator
// BenchmarkWhen는 When 검증기를 벤치마크합니다
func BenchmarkWhen(b *testing.B) {
	email := "user@example.com"
	isRequired := true
	for i := 0; i < b.N; i++ {
		v := New(email, "email")
		v.When(isRequired, func(val *Validator) {
			val.Required().Email()
		})
		_ = v.Validate()
	}
}

// BenchmarkUnless benchmarks the Unless validator
// BenchmarkUnless는 Unless 검증기를 벤치마크합니다
func BenchmarkUnless(b *testing.B) {
	email := "user@example.com"
	isGuest := false
	for i := 0; i < b.N; i++ {
		v := New(email, "email")
		v.Unless(isGuest, func(val *Validator) {
			val.Required().Email()
		})
		_ = v.Validate()
	}
}

// BenchmarkTrue benchmarks the True validator
// BenchmarkTrue는 True 검증기를 벤치마크합니다
func BenchmarkTrue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New(true, "accepted")
		v.True()
		_ = v.Validate()
	}
}

// BenchmarkFalse benchmarks the False validator
// BenchmarkFalse는 False 검증기를 벤치마크합니다
func BenchmarkFalse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New(false, "declined")
		v.False()
		_ = v.Validate()
	}
}

// BenchmarkNil benchmarks the Nil validator
// BenchmarkNil는 Nil 검증기를 벤치마크합니다
func BenchmarkNil(b *testing.B) {
	var ptr *string
	for i := 0; i < b.N; i++ {
		v := New(ptr, "optional")
		v.Nil()
		_ = v.Validate()
	}
}

// BenchmarkNotNil benchmarks the NotNil validator
// BenchmarkNotNil는 NotNil 검증기를 벤치마크합니다
func BenchmarkNotNil(b *testing.B) {
	str := "value"
	ptr := &str
	for i := 0; i < b.N; i++ {
		v := New(ptr, "required_ptr")
		v.NotNil()
		_ = v.Validate()
	}
}

// BenchmarkType benchmarks the Type validator
// BenchmarkType는 Type 검증기를 벤치마크합니다
func BenchmarkType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("test string", "text")
		v.Type("string")
		_ = v.Validate()
	}
}

// BenchmarkEmpty benchmarks the Empty validator
// BenchmarkEmpty는 Empty 검증기를 벤치마크합니다
func BenchmarkEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("", "optional_field")
		v.Empty()
		_ = v.Validate()
	}
}

// BenchmarkNotEmpty benchmarks the NotEmpty validator
// BenchmarkNotEmpty는 NotEmpty 검증기를 벤치마크합니다
func BenchmarkNotEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("value", "required_field")
		v.NotEmpty()
		_ = v.Validate()
	}
}

// BenchmarkBetweenTime benchmarks the BetweenTime validator
// BenchmarkBetweenTime는 BetweenTime 검증기를 벤치마크합니다
func BenchmarkBetweenTime(b *testing.B) {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)
	middle := time.Date(2024, 6, 15, 12, 0, 0, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		v := New(middle, "date")
		v.BetweenTime(start, end)
		_ = v.Validate()
	}
}

// BenchmarkWithCustomMessage benchmarks the WithCustomMessage method
// BenchmarkWithCustomMessage는 WithCustomMessage 메서드를 벤치마크합니다
func BenchmarkWithCustomMessage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("", "email")
		v.WithCustomMessage("required", "Please enter your email address")
		v.Required()
		_ = v.Validate()
	}
}

// BenchmarkWithCustomMessages benchmarks the WithCustomMessages method
// BenchmarkWithCustomMessages는 WithCustomMessages 메서드를 벤치마크합니다
func BenchmarkWithCustomMessages(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := New("", "password")
		v.WithCustomMessages(map[string]string{
			"required":  "비밀번호를 입력해주세요",
			"minlength": "비밀번호는 8자 이상이어야 합니다",
			"maxlength": "비밀번호는 20자 이하여야 합니다",
		})
		v.StopOnError().Required().MinLength(8).MaxLength(20)
		_ = v.Validate()
	}
}

// BenchmarkCustomMessageVsDefault compares custom message vs default message performance
// BenchmarkCustomMessageVsDefault는 커스텀 메시지와 기본 메시지의 성능을 비교합니다
func BenchmarkCustomMessageVsDefault(b *testing.B) {
	b.Run("Default message", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			v := New("", "field")
			v.Required()
			_ = v.Validate()
		}
	})

	b.Run("Custom message", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			v := New("", "field")
			v.WithCustomMessage("required", "Custom required message")
			v.Required()
			_ = v.Validate()
		}
	})
}
