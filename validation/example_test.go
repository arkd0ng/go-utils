package validation_test

import (
	"fmt"
	"os"
	"time"

	"github.com/arkd0ng/go-utils/validation"
)

// Example tests for validation package documentation
// validation 패키지 문서화를 위한 Example 테스트

// ExampleNew demonstrates creating a new validator
// ExampleNew는 새 검증기 생성을 보여줍니다
func ExampleNew() {
	v := validation.New("john@example.com", "email")
	v.Required().Email()
	err := v.Validate()
	if err != nil {
		fmt.Println("Validation failed")
	} else {
		fmt.Println("Validation passed")
	}
	// Output: Validation passed
}

// ExampleValidator_Required demonstrates the Required validator
// ExampleValidator_Required는 Required 검증기를 보여줍니다
func ExampleValidator_Required() {
	v := validation.New("", "username")
	v.Required()
	err := v.Validate()
	if err != nil {
		fmt.Println("Error: field is required")
	}
	// Output: Error: field is required
}

// ExampleValidator_MinLength demonstrates the MinLength validator
// ExampleValidator_MinLength는 MinLength 검증기를 보여줍니다
func ExampleValidator_MinLength() {
	v := validation.New("abc", "password")
	v.MinLength(8)
	err := v.Validate()
	if err != nil {
		fmt.Println("Error: password too short")
	}
	// Output: Error: password too short
}

// ExampleValidator_MaxLength demonstrates the MaxLength validator
// ExampleValidator_MaxLength는 MaxLength 검증기를 보여줍니다
func ExampleValidator_MaxLength() {
	v := validation.New("short", "username")
	v.MaxLength(20)
	err := v.Validate()
	if err != nil {
		fmt.Println("Error: username too long")
	} else {
		fmt.Println("Valid username")
	}
	// Output: Valid username
}

// ExampleValidator_Email demonstrates the Email validator
// ExampleValidator_Email는 Email 검증기를 보여줍니다
func ExampleValidator_Email() {
	v := validation.New("user@example.com", "email")
	v.Email()
	err := v.Validate()
	if err != nil {
		fmt.Println("Invalid email")
	} else {
		fmt.Println("Valid email")
	}
	// Output: Valid email
}

// ExampleValidator_URL demonstrates the URL validator
// ExampleValidator_URL는 URL 검증기를 보여줍니다
func ExampleValidator_URL() {
	v := validation.New("https://example.com", "website")
	v.URL()
	err := v.Validate()
	if err != nil {
		fmt.Println("Invalid URL")
	} else {
		fmt.Println("Valid URL")
	}
	// Output: Valid URL
}

// ExampleValidator_Min demonstrates the Min validator
// ExampleValidator_Min는 Min 검증기를 보여줍니다
func ExampleValidator_Min() {
	v := validation.New(25, "age")
	v.Min(18)
	err := v.Validate()
	if err != nil {
		fmt.Println("Too young")
	} else {
		fmt.Println("Valid age")
	}
	// Output: Valid age
}

// ExampleValidator_Max demonstrates the Max validator
// ExampleValidator_Max는 Max 검증기를 보여줍니다
func ExampleValidator_Max() {
	v := validation.New(150, "age")
	v.Max(120)
	err := v.Validate()
	if err != nil {
		fmt.Println("Age exceeds maximum")
	}
	// Output: Age exceeds maximum
}

// Example_rangeValidation demonstrates using Min and Max together for range validation
// Example_rangeValidation는 범위 검증을 위해 Min과 Max를 함께 사용하는 것을 보여줍니다
func Example_rangeValidation() {
	v := validation.New(50, "percentage")
	v.Min(0).Max(100)
	err := v.Validate()
	if err != nil {
		fmt.Println("Out of range")
	} else {
		fmt.Println("Valid percentage")
	}
	// Output: Valid percentage
}

// ExampleValidator_In demonstrates the In validator
// ExampleValidator_In는 In 검증기를 보여줍니다
func ExampleValidator_In() {
	v := validation.New("apple", "fruit")
	v.In("apple", "banana", "orange")
	err := v.Validate()
	if err != nil {
		fmt.Println("Invalid fruit")
	} else {
		fmt.Println("Valid fruit")
	}
	// Output: Valid fruit
}

// ExampleValidator_NotIn demonstrates the NotIn validator
// ExampleValidator_NotIn는 NotIn 검증기를 보여줍니다
func ExampleValidator_NotIn() {
	v := validation.New("admin", "role")
	v.NotIn("guest", "anonymous")
	err := v.Validate()
	if err != nil {
		fmt.Println("Forbidden role")
	} else {
		fmt.Println("Allowed role")
	}
	// Output: Allowed role
}

// ExampleValidator_ArrayLength demonstrates the ArrayLength validator
// ExampleValidator_ArrayLength는 ArrayLength 검증기를 보여줍니다
func ExampleValidator_ArrayLength() {
	v := validation.New([]int{1, 2, 3}, "numbers")
	v.ArrayLength(3)
	err := v.Validate()
	if err != nil {
		fmt.Println("Invalid array length")
	} else {
		fmt.Println("Valid array length")
	}
	// Output: Valid array length
}

// ExampleValidator_ArrayMinLength demonstrates the ArrayMinLength validator
// ExampleValidator_ArrayMinLength는 ArrayMinLength 검증기를 보여줍니다
func ExampleValidator_ArrayMinLength() {
	v := validation.New([]string{"a", "b", "c"}, "items")
	v.ArrayMinLength(2)
	err := v.Validate()
	if err != nil {
		fmt.Println("Array too short")
	} else {
		fmt.Println("Valid array length")
	}
	// Output: Valid array length
}

// ExampleValidator_ArrayMaxLength demonstrates the ArrayMaxLength validator
// ExampleValidator_ArrayMaxLength는 ArrayMaxLength 검증기를 보여줍니다
func ExampleValidator_ArrayMaxLength() {
	v := validation.New([]string{"a", "b"}, "items")
	v.ArrayMaxLength(5)
	err := v.Validate()
	if err != nil {
		fmt.Println("Array too long")
	} else {
		fmt.Println("Valid array length")
	}
	// Output: Valid array length
}

// ExampleValidator_ArrayUnique demonstrates the ArrayUnique validator
// ExampleValidator_ArrayUnique는 ArrayUnique 검증기를 보여줍니다
func ExampleValidator_ArrayUnique() {
	v := validation.New([]int{1, 2, 3, 4, 5}, "numbers")
	v.ArrayUnique()
	err := v.Validate()
	if err != nil {
		fmt.Println("Duplicate values found")
	} else {
		fmt.Println("All values unique")
	}
	// Output: All values unique
}

// ExampleValidator_MapHasKey demonstrates the MapHasKey validator
// ExampleValidator_MapHasKey는 MapHasKey 검증기를 보여줍니다
func ExampleValidator_MapHasKey() {
	data := map[string]int{"age": 25, "score": 100}
	v := validation.New(data, "data")
	v.MapHasKey("age")
	err := v.Validate()
	if err != nil {
		fmt.Println("Key not found")
	} else {
		fmt.Println("Key exists")
	}
	// Output: Key exists
}

// ExampleValidator_MapHasKeys demonstrates the MapHasKeys validator
// ExampleValidator_MapHasKeys는 MapHasKeys 검증기를 보여줍니다
func ExampleValidator_MapHasKeys() {
	data := map[string]int{"name": 1, "age": 25, "city": 3}
	v := validation.New(data, "data")
	v.MapHasKeys("name", "age")
	err := v.Validate()
	if err != nil {
		fmt.Println("Missing required keys")
	} else {
		fmt.Println("All keys present")
	}
	// Output: All keys present
}

// ExampleValidator_Equals demonstrates the Equals validator
// ExampleValidator_Equals는 Equals 검증기를 보여줍니다
func ExampleValidator_Equals() {
	v := validation.New(100, "value")
	v.Equals(100)
	err := v.Validate()
	if err != nil {
		fmt.Println("Values not equal")
	} else {
		fmt.Println("Values are equal")
	}
	// Output: Values are equal
}

// ExampleValidator_NotEquals demonstrates the NotEquals validator
// ExampleValidator_NotEquals는 NotEquals 검증기를 보여줍니다
func ExampleValidator_NotEquals() {
	v := validation.New("newpassword", "password")
	v.NotEquals("oldpassword")
	err := v.Validate()
	if err != nil {
		fmt.Println("Passwords must be different")
	} else {
		fmt.Println("Valid new password")
	}
	// Output: Valid new password
}

// ExampleValidator_Before demonstrates the Before validator
// ExampleValidator_Before는 Before 검증기를 보여줍니다
func ExampleValidator_Before() {
	now := time.Now()
	yesterday := now.Add(-24 * time.Hour)
	v := validation.New(yesterday, "date")
	v.Before(now)
	err := v.Validate()
	if err != nil {
		fmt.Println("Date not before now")
	} else {
		fmt.Println("Valid date")
	}
	// Output: Valid date
}

// ExampleValidator_After demonstrates the After validator
// ExampleValidator_After는 After 검증기를 보여줍니다
func ExampleValidator_After() {
	now := time.Now()
	tomorrow := now.Add(24 * time.Hour)
	v := validation.New(tomorrow, "date")
	v.After(now)
	err := v.Validate()
	if err != nil {
		fmt.Println("Date not after now")
	} else {
		fmt.Println("Valid date")
	}
	// Output: Valid date
}

// ExampleValidator_Custom demonstrates the Custom validator
// ExampleValidator_Custom는 Custom 검증기를 보여줍니다
func ExampleValidator_Custom() {
	v := validation.New("password123", "password")
	v.Custom(func(val interface{}) bool {
		str, ok := val.(string)
		if !ok {
			return false
		}
		// Check if password contains numbers
		for _, ch := range str {
			if ch >= '0' && ch <= '9' {
				return true
			}
		}
		return false
	}, "password must contain at least one number")
	err := v.Validate()
	if err != nil {
		fmt.Println("Invalid password")
	} else {
		fmt.Println("Valid password")
	}
	// Output: Valid password
}

// ExampleValidator_StopOnError demonstrates StopOnError behavior
// ExampleValidator_StopOnError는 StopOnError 동작을 보여줍니다
func ExampleValidator_StopOnError() {
	v := validation.New("", "email")
	v.StopOnError()
	v.Required().Email().MinLength(5)
	err := v.Validate()
	if err != nil {
		verrs := err.(validation.ValidationErrors)
		fmt.Printf("Number of errors: %d\n", len(verrs))
	}
	// Output: Number of errors: 1
}

// ExampleNewValidator demonstrates MultiValidator for multiple fields
// ExampleNewValidator는 여러 필드를 위한 MultiValidator를 보여줍니다
func ExampleNewValidator() {
	mv := validation.NewValidator()
	mv.Field("john", "name").Required().MinLength(2)
	mv.Field(25, "age").Min(18).Max(150)

	err := mv.Validate()
	if err != nil {
		fmt.Println("Validation failed")
	} else {
		fmt.Println("All fields valid")
	}
	// Output: All fields valid
}

// ExampleMultiValidator_Field demonstrates adding fields to MultiValidator
// ExampleMultiValidator_Field는 MultiValidator에 필드 추가를 보여줍니다
func ExampleMultiValidator_Field() {
	mv := validation.NewValidator()
	mv.Field("TestUser", "username").Required().MinLength(3)
	mv.Field(100, "score").Min(0).Max(200)

	err := mv.Validate()
	if err != nil {
		fmt.Println("Validation errors found")
	} else {
		fmt.Println("All validations passed")
	}
	// Output: All validations passed
}

// ExampleValidationErrors demonstrates handling validation errors
// ExampleValidationErrors는 검증 에러 처리를 보여줍니다
func ExampleValidationErrors() {
	v := validation.New("", "email")
	v.StopOnError() // Stop at first error
	v.Required()
	v.Email()

	err := v.Validate()
	if err != nil {
		verrs := err.(validation.ValidationErrors)
		for _, verr := range verrs {
			fmt.Printf("Field: %s, Rule: %s\n", verr.Field, verr.Rule)
		}
	}
	// Output: Field: email, Rule: required
}

// Example_chainedValidation demonstrates chaining multiple validators
// Example_chainedValidation는 여러 검증기 체이닝을 보여줍니다
func Example_chainedValidation() {
	v := validation.New("user@example.com", "email")
	v.Required().MinLength(5).MaxLength(100).Email()

	err := v.Validate()
	if err != nil {
		fmt.Println("Validation failed")
	} else {
		fmt.Println("Email is valid")
	}
	// Output: Email is valid
}

// Example_complexValidation demonstrates complex validation scenario
// Example_complexValidation는 복잡한 검증 시나리오를 보여줍니다
func Example_complexValidation() {
	// User registration validation
	mv := validation.NewValidator()

	// Username validation
	mv.Field("john_doe", "username").
		Required().
		MinLength(3).
		MaxLength(50)

	// Password validation
	mv.Field("SecurePass123!", "password").
		Required().
		MinLength(8).
		MaxLength(128)

	// Age validation
	mv.Field(25, "age").
		Min(18).
		Max(150)

	// Role validation
	mv.Field("user", "role").
		Required().
		In("admin", "user", "moderator")

	err := mv.Validate()
	if err != nil {
		fmt.Println("Registration validation failed")
	} else {
		fmt.Println("Registration data is valid")
	}
	// Output: Registration data is valid
}

// ExampleValidator_IPv4 demonstrates the IPv4 validator
// ExampleValidator_IPv4는 IPv4 검증기를 보여줍니다
func ExampleValidator_IPv4() {
	v := validation.New("192.168.1.1", "ip_address")
	v.IPv4()
	err := v.Validate()
	if err != nil {
		fmt.Println("Invalid IPv4 address")
	} else {
		fmt.Println("Valid IPv4 address")
	}
	// Output: Valid IPv4 address
}

// ExampleValidator_IPv6 demonstrates the IPv6 validator
// ExampleValidator_IPv6는 IPv6 검증기를 보여줍니다
func ExampleValidator_IPv6() {
	v := validation.New("2001:db8::1", "ip_address")
	v.IPv6()
	err := v.Validate()
	if err != nil {
		fmt.Println("Invalid IPv6 address")
	} else {
		fmt.Println("Valid IPv6 address")
	}
	// Output: Valid IPv6 address
}

// ExampleValidator_IP demonstrates the IP validator
// ExampleValidator_IP는 IP 검증기를 보여줍니다
func ExampleValidator_IP() {
	v := validation.New("10.0.0.1", "ip_address")
	v.IP()
	err := v.Validate()
	if err != nil {
		fmt.Println("Invalid IP address")
	} else {
		fmt.Println("Valid IP address")
	}
	// Output: Valid IP address
}

// ExampleValidator_CIDR demonstrates the CIDR validator
// ExampleValidator_CIDR는 CIDR 검증기를 보여줍니다
func ExampleValidator_CIDR() {
	v := validation.New("192.168.1.0/24", "network")
	v.CIDR()
	err := v.Validate()
	if err != nil {
		fmt.Println("Invalid CIDR notation")
	} else {
		fmt.Println("Valid CIDR notation")
	}
	// Output: Valid CIDR notation
}

// ExampleValidator_MAC demonstrates the MAC validator
// ExampleValidator_MAC는 MAC 검증기를 보여줍니다
func ExampleValidator_MAC() {
	v := validation.New("00:1A:2B:3C:4D:5E", "mac_address")
	v.MAC()
	err := v.Validate()
	if err != nil {
		fmt.Println("Invalid MAC address")
	} else {
		fmt.Println("Valid MAC address")
	}
	// Output: Valid MAC address
}

// Example_networkValidation demonstrates validating network configuration
// Example_networkValidation는 네트워크 구성 검증을 보여줍니다
func Example_networkValidation() {
	// Validate server configuration
	mv := validation.NewValidator()
	mv.Field("192.168.1.10", "server_ip").Required().IPv4()
	mv.Field("192.168.1.0/24", "subnet").Required().CIDR()
	mv.Field("00:1A:2B:3C:4D:5E", "mac").Required().MAC()

	err := mv.Validate()
	if err != nil {
		fmt.Println("Invalid network configuration")
	} else {
		fmt.Println("Valid network configuration")
	}
	// Output: Valid network configuration
}

// ExampleValidator_DateFormat demonstrates date format validation
// ExampleValidator_DateFormat는 날짜 형식 검증을 보여줍니다
func ExampleValidator_DateFormat() {
	v := validation.New("2025-10-17", "birth_date")
	v.DateFormat("2006-01-02")
	err := v.Validate()
	if err != nil {
		fmt.Println("Invalid date format")
	} else {
		fmt.Println("Valid date format")
	}
	// Output: Valid date format
}

// ExampleValidator_TimeFormat demonstrates time format validation
// ExampleValidator_TimeFormat는 시간 형식 검증을 보여줍니다
func ExampleValidator_TimeFormat() {
	v := validation.New("14:30:00", "meeting_time")
	v.TimeFormat("15:04:05")
	err := v.Validate()
	if err != nil {
		fmt.Println("Invalid time format")
	} else {
		fmt.Println("Valid time format")
	}
	// Output: Valid time format
}

// ExampleValidator_DateBefore demonstrates date before validation
// ExampleValidator_DateBefore는 날짜 이전 검증을 보여줍니다
func ExampleValidator_DateBefore() {
	maxDate := time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC)
	testDate := time.Date(2025, 10, 17, 12, 0, 0, 0, time.UTC)

	v := validation.New(testDate, "expiry_date")
	v.DateBefore(maxDate)
	err := v.Validate()
	if err != nil {
		fmt.Println("Date is not before max date")
	} else {
		fmt.Println("Date is before max date")
	}
	// Output: Date is before max date
}

// ExampleValidator_DateAfter demonstrates date after validation
// ExampleValidator_DateAfter는 날짜 이후 검증을 보여줍니다
func ExampleValidator_DateAfter() {
	minDate := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	testDate := time.Date(2025, 10, 17, 12, 0, 0, 0, time.UTC)

	v := validation.New(testDate, "start_date")
	v.DateAfter(minDate)
	err := v.Validate()
	if err != nil {
		fmt.Println("Date is not after min date")
	} else {
		fmt.Println("Date is after min date")
	}
	// Output: Date is after min date
}

// Example_dateTimeValidation demonstrates comprehensive date/time validation
// Example_dateTimeValidation는 포괄적인 날짜/시간 검증을 보여줍니다
func Example_dateTimeValidation() {
	// Event scheduling validation
	minDate := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	maxDate := time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC)
	eventDateTime := time.Date(2025, 10, 17, 14, 30, 0, 0, time.UTC)

	mv := validation.NewValidator()
	mv.Field("2025-10-17", "event_date").Required().DateFormat("2006-01-02")
	mv.Field("14:30:00", "event_time").Required().TimeFormat("15:04:05")
	mv.Field(eventDateTime, "event_datetime").DateAfter(minDate).DateBefore(maxDate)

	err := mv.Validate()
	if err != nil {
		fmt.Println("Invalid event schedule")
	} else {
		fmt.Println("Valid event schedule")
	}
	// Output: Valid event schedule
}

// ExampleValidator_IntRange demonstrates integer range validation
// ExampleValidator_IntRange는 정수 범위 검증을 보여줍니다
func ExampleValidator_IntRange() {
	v := validation.New(25, "age")
	v.IntRange(18, 65)
	err := v.Validate()
	if err != nil {
		fmt.Println("Age is out of range")
	} else {
		fmt.Println("Valid age")
	}
	// Output: Valid age
}

// ExampleValidator_FloatRange demonstrates float range validation
// ExampleValidator_FloatRange는 실수 범위 검증을 보여줍니다
func ExampleValidator_FloatRange() {
	v := validation.New(98.6, "temperature")
	v.FloatRange(95.0, 105.0)
	err := v.Validate()
	if err != nil {
		fmt.Println("Temperature is out of range")
	} else {
		fmt.Println("Valid temperature")
	}
	// Output: Valid temperature
}

// ExampleValidator_DateRange demonstrates date range validation
// ExampleValidator_DateRange는 날짜 범위 검증을 보여줍니다
func ExampleValidator_DateRange() {
	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC)
	testDate := time.Date(2025, 6, 15, 12, 0, 0, 0, time.UTC)

	v := validation.New(testDate, "event_date")
	v.DateRange(start, end)
	err := v.Validate()
	if err != nil {
		fmt.Println("Date is out of range")
	} else {
		fmt.Println("Date is within range")
	}
	// Output: Date is within range
}

// Example_rangeValidationComprehensive demonstrates comprehensive range validation
// Example_rangeValidationComprehensive는 포괄적인 범위 검증을 보여줍니다
func Example_rangeValidationComprehensive() {
	// Event validation with multiple range checks
	start := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC)
	eventDate := time.Date(2025, 6, 15, 12, 0, 0, 0, time.UTC)

	mv := validation.NewValidator()
	mv.Field(25, "participant_age").IntRange(18, 65)
	mv.Field(25.50, "ticket_price").FloatRange(10.0, 100.0)
	mv.Field(eventDate, "event_date").DateRange(start, end)

	err := mv.Validate()
	if err != nil {
		fmt.Println("Invalid event data")
	} else {
		fmt.Println("Valid event data")
	}
	// Output: Valid event data
}

// Example_uuidv4Validation demonstrates UUIDv4 validation
// Example_uuidv4Validation은 UUIDv4 검증을 보여줍니다
func Example_uuidv4Validation() {
	v := validation.New("550e8400-e29b-41d4-a716-446655440000", "request_id")
	v.UUIDv4()

	err := v.Validate()
	if err != nil {
		fmt.Println("Invalid UUID v4")
	} else {
		fmt.Println("Valid UUID v4")
	}
	// Output: Valid UUID v4
}

// Example_xmlValidation demonstrates XML validation
// Example_xmlValidation은 XML 검증을 보여줍니다
func Example_xmlValidation() {
	xmlData := `<?xml version="1.0"?>
	<person>
		<name>John Doe</name>
		<age>30</age>
	</person>`

	v := validation.New(xmlData, "user_data")
	v.XML()

	err := v.Validate()
	if err != nil {
		fmt.Println("Invalid XML")
	} else {
		fmt.Println("Valid XML")
	}
	// Output: Valid XML
}

// Example_hexValidation demonstrates hexadecimal validation
// Example_hexValidation은 16진수 검증을 보여줍니다
func Example_hexValidation() {
	v := validation.New("0xdeadbeef", "color_code")
	v.Hex()

	err := v.Validate()
	if err != nil {
		fmt.Println("Invalid hex")
	} else {
		fmt.Println("Valid hex")
	}
	// Output: Valid hex
}

// Example_formatValidationComprehensive demonstrates comprehensive format validation
// Example_formatValidationComprehensive는 포괄적인 형식 검증을 보여줍니다
func Example_formatValidationComprehensive() {
	// API Request validation with multiple format checks
	requestID := "550e8400-e29b-41d4-a716-446655440000"
	configData := `{"timeout": 30, "retries": 3}`
	hexToken := "0xabcd1234"

	mv := validation.NewValidator()
	mv.Field(requestID, "request_id").UUIDv4()
	mv.Field(configData, "config").JSON()
	mv.Field(hexToken, "token").Hex()

	err := mv.Validate()
	if err != nil {
		fmt.Println("Invalid request format")
	} else {
		fmt.Println("Valid request format")
	}
	// Output: Valid request format
}

// Example_filePathValidation demonstrates file path validation
// Example_filePathValidation은 파일 경로 검증을 보여줍니다
func Example_filePathValidation() {
	v := validation.New("./config/app.json", "config_file")
	v.FilePath()

	err := v.Validate()
	if err != nil {
		fmt.Println("Invalid file path")
	} else {
		fmt.Println("Valid file path")
	}
	// Output: Valid file path
}

// Example_fileExistsValidation demonstrates file existence validation
// Example_fileExistsValidation은 파일 존재 검증을 보여줍니다
func Example_fileExistsValidation() {
	// Create a temporary file for demonstration
	tmpFile, _ := os.CreateTemp("", "example_*.txt")
	tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	v := validation.New(tmpFile.Name(), "log_file")
	v.FileExists()

	err := v.Validate()
	if err != nil {
		fmt.Println("File does not exist")
	} else {
		fmt.Println("File exists")
	}
	// Output: File exists
}

// Example_fileSizeValidation demonstrates file size validation
// Example_fileSizeValidation은 파일 크기 검증을 보여줍니다
func Example_fileSizeValidation() {
	// Create a temporary file with content
	tmpFile, _ := os.CreateTemp("", "example_*.txt")
	tmpFile.WriteString("Hello, World!")
	tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	v := validation.New(tmpFile.Name(), "upload_file")
	v.FileSize(0, 1024) // Max 1KB

	err := v.Validate()
	if err != nil {
		fmt.Println("File size out of range")
	} else {
		fmt.Println("File size OK")
	}
	// Output: File size OK
}

// Example_fileExtensionValidation demonstrates file extension validation
// Example_fileExtensionValidation은 파일 확장자 검증을 보여줍니다
func Example_fileExtensionValidation() {
	v := validation.New("document.pdf", "file_name")
	v.FileExtension(".pdf", ".doc", ".docx")

	err := v.Validate()
	if err != nil {
		fmt.Println("Invalid file extension")
	} else {
		fmt.Println("Valid file extension")
	}
	// Output: Valid file extension
}

// Example_fileValidationComprehensive demonstrates comprehensive file validation
// Example_fileValidationComprehensive는 포괄적인 파일 검증을 보여줍니다
func Example_fileValidationComprehensive() {
	// Create a temporary file for demonstration
	tmpFile, _ := os.CreateTemp("", "upload_*.txt")
	tmpFile.WriteString("Test content")
	tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	mv := validation.NewValidator()
	mv.Field(tmpFile.Name(), "upload_file").
		FileExists().
		FileReadable().
		FileSize(0, 10240). // Max 10KB
		FileExtension(".txt", ".log")

	err := mv.Validate()
	if err != nil {
		fmt.Println("Invalid file")
	} else {
		fmt.Println("Valid file")
	}
	// Output: Valid file
}
