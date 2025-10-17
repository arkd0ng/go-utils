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
