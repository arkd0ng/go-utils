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

// ExampleValidator_CreditCard demonstrates credit card validation
// ExampleValidator_CreditCard는 신용카드 검증을 보여줍니다
func ExampleValidator_CreditCard() {
	v := validation.New("4532015112830366", "card_number")
	v.CreditCard()

	if len(v.GetErrors()) > 0 {
		fmt.Println("Invalid credit card")
	} else {
		fmt.Println("Valid credit card")
	}
	// Output: Valid credit card
}

// ExampleValidator_CreditCardType demonstrates credit card type validation
// ExampleValidator_CreditCardType는 신용카드 타입 검증을 보여줍니다
func ExampleValidator_CreditCardType() {
	v := validation.New("4532015112830366", "card_number")
	v.CreditCardType("visa")

	if len(v.GetErrors()) > 0 {
		fmt.Println("Invalid Visa card")
	} else {
		fmt.Println("Valid Visa card")
	}
	// Output: Valid Visa card
}

// ExampleValidator_Luhn demonstrates Luhn algorithm validation
// ExampleValidator_Luhn는 Luhn 알고리즘 검증을 보여줍니다
func ExampleValidator_Luhn() {
	v := validation.New("79927398713", "identifier")
	v.Luhn()

	if len(v.GetErrors()) > 0 {
		fmt.Println("Invalid Luhn number")
	} else {
		fmt.Println("Valid Luhn number")
	}
	// Output: Valid Luhn number
}

// Example_creditCardValidation demonstrates credit card validation with spaces and hyphens
// Example_creditCardValidation는 공백과 하이픈이 있는 신용카드 검증을 보여줍니다
func Example_creditCardValidation() {
	// Credit card with spaces
	v := validation.New("4532 0151 1283 0366", "card_number")
	v.CreditCard()

	if len(v.GetErrors()) > 0 {
		fmt.Println("Invalid")
	} else {
		fmt.Println("Valid with spaces")
	}
	// Output: Valid with spaces
}

// Example_creditCardTypeValidation demonstrates validation of different card types
// Example_creditCardTypeValidation는 다양한 카드 타입 검증을 보여줍니다
func Example_creditCardTypeValidation() {
	mv := validation.NewValidator()

	// Validate Visa
	mv.Field("4532015112830366", "visa_card").CreditCardType("visa")

	// Validate Mastercard
	mv.Field("5425233430109903", "mastercard").CreditCardType("mastercard")

	// Validate Amex
	mv.Field("374245455400126", "amex_card").CreditCardType("amex")

	err := mv.Validate()
	if err != nil {
		fmt.Println("Invalid cards")
	} else {
		fmt.Println("All cards valid")
	}
	// Output: All cards valid
}

// ExampleValidator_ISBN demonstrates ISBN validation
// ExampleValidator_ISBN는 ISBN 검증을 보여줍니다
func ExampleValidator_ISBN() {
	v := validation.New("978-0-596-52068-7", "isbn")
	v.ISBN()

	if len(v.GetErrors()) > 0 {
		fmt.Println("Invalid ISBN")
	} else {
		fmt.Println("Valid ISBN")
	}
	// Output: Valid ISBN
}

// ExampleValidator_ISSN demonstrates ISSN validation
// ExampleValidator_ISSN는 ISSN 검증을 보여줍니다
func ExampleValidator_ISSN() {
	v := validation.New("2049-3630", "issn")
	v.ISSN()

	if len(v.GetErrors()) > 0 {
		fmt.Println("Invalid ISSN")
	} else {
		fmt.Println("Valid ISSN")
	}
	// Output: Valid ISSN
}

// ExampleValidator_EAN demonstrates EAN validation
// ExampleValidator_EAN는 EAN 검증을 보여줍니다
func ExampleValidator_EAN() {
	v := validation.New("4006381333931", "ean")
	v.EAN()

	if len(v.GetErrors()) > 0 {
		fmt.Println("Invalid EAN")
	} else {
		fmt.Println("Valid EAN")
	}
	// Output: Valid EAN
}

// Example_businessIDValidation demonstrates business ID validation
// Example_businessIDValidation는 비즈니스 ID 검증을 보여줍니다
func Example_businessIDValidation() {
	mv := validation.NewValidator()

	// Validate book ISBN
	mv.Field("978-0-596-52068-7", "book_isbn").ISBN()

	// Validate journal ISSN
	mv.Field("2049-3630", "journal_issn").ISSN()

	// Validate product EAN
	mv.Field("4006381333931", "product_ean").EAN()

	err := mv.Validate()
	if err != nil {
		fmt.Println("Invalid IDs")
	} else {
		fmt.Println("All IDs valid")
	}
	// Output: All IDs valid
}

// ExampleValidator_Latitude demonstrates latitude validation
// ExampleValidator_Latitude는 위도 검증을 보여줍니다
func ExampleValidator_Latitude() {
	v := validation.New(37.5665, "latitude")
	v.Latitude()

	if len(v.GetErrors()) > 0 {
		fmt.Println("Invalid latitude")
	} else {
		fmt.Println("Valid latitude")
	}
	// Output: Valid latitude
}

// ExampleValidator_Longitude demonstrates longitude validation
// ExampleValidator_Longitude는 경도 검증을 보여줍니다
func ExampleValidator_Longitude() {
	v := validation.New(126.9780, "longitude")
	v.Longitude()

	if len(v.GetErrors()) > 0 {
		fmt.Println("Invalid longitude")
	} else {
		fmt.Println("Valid longitude")
	}
	// Output: Valid longitude
}

// ExampleValidator_Coordinate demonstrates coordinate string validation
// ExampleValidator_Coordinate는 좌표 문자열 검증을 보여줍니다
func ExampleValidator_Coordinate() {
	v := validation.New("37.5665,126.9780", "location")
	v.Coordinate()

	if len(v.GetErrors()) > 0 {
		fmt.Println("Invalid coordinate")
	} else {
		fmt.Println("Valid coordinate")
	}
	// Output: Valid coordinate
}

// Example_geographicValidation demonstrates geographic validation
// Example_geographicValidation는 지리적 검증을 보여줍니다
func Example_geographicValidation() {
	mv := validation.NewValidator()

	// Validate latitude
	mv.Field(37.5665, "latitude").Latitude()

	// Validate longitude
	mv.Field(126.9780, "longitude").Longitude()

	// Validate coordinate string
	mv.Field("40.7128,-74.0060", "nyc_location").Coordinate()

	err := mv.Validate()
	if err != nil {
		fmt.Println("Invalid location")
	} else {
		fmt.Println("All locations valid")
	}
	// Output: All locations valid
}

// ExampleValidator_JWT demonstrates JWT validation
// ExampleValidator_JWT는 JWT 검증을 보여줍니다
func ExampleValidator_JWT() {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIn0.dozjgNryP4J3jVmNHl0w5N_XgL0n3I9PlFUP0THsR8U"
	v := validation.New(token, "jwt_token")
	v.JWT()

	if len(v.GetErrors()) > 0 {
		fmt.Println("Invalid JWT")
	} else {
		fmt.Println("Valid JWT")
	}
	// Output: Valid JWT
}

// ExampleValidator_BCrypt demonstrates BCrypt hash validation
// ExampleValidator_BCrypt는 BCrypt 해시 검증을 보여줍니다
func ExampleValidator_BCrypt() {
	hash := "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy"
	v := validation.New(hash, "password_hash")
	v.BCrypt()

	if len(v.GetErrors()) > 0 {
		fmt.Println("Invalid BCrypt hash")
	} else {
		fmt.Println("Valid BCrypt hash")
	}
	// Output: Valid BCrypt hash
}

// ExampleValidator_MD5 demonstrates MD5 hash validation
// ExampleValidator_MD5는 MD5 해시 검증을 보여줍니다
func ExampleValidator_MD5() {
	hash := "5d41402abc4b2a76b9719d911017c592"
	v := validation.New(hash, "file_hash")
	v.MD5()

	if len(v.GetErrors()) > 0 {
		fmt.Println("Invalid MD5")
	} else {
		fmt.Println("Valid MD5")
	}
	// Output: Valid MD5
}

// ExampleValidator_SHA256 demonstrates SHA256 hash validation
// ExampleValidator_SHA256는 SHA256 해시 검증을 보여줍니다
func ExampleValidator_SHA256() {
	hash := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	v := validation.New(hash, "file_hash")
	v.SHA256()

	if len(v.GetErrors()) > 0 {
		fmt.Println("Invalid SHA256")
	} else {
		fmt.Println("Valid SHA256")
	}
	// Output: Valid SHA256
}

// Example_securityValidation demonstrates security validation
// Example_securityValidation는 보안 검증을 보여줍니다
func Example_securityValidation() {
	mv := validation.NewValidator()

	// Validate JWT token
	mv.Field("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIn0.dozjgNryP4J3jVmNHl0w5N_XgL0n3I9PlFUP0THsR8U", "token").JWT()

	// Validate BCrypt password hash
	mv.Field("$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy", "password").BCrypt()

	// Validate file hash
	mv.Field("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", "file_hash").SHA256()

	err := mv.Validate()
	if err != nil {
		fmt.Println("Invalid security data")
	} else {
		fmt.Println("All security data valid")
	}
	// Output: All security data valid
}

// ExampleValidator_HexColor demonstrates hex color validation
// ExampleValidator_HexColor는 16진수 색상 검증을 보여줍니다
func ExampleValidator_HexColor() {
	color := "#FF5733"
	v := validation.New(color, "brand_color")
	v.HexColor()

	if len(v.GetErrors()) > 0 {
		fmt.Println("Invalid hex color")
	} else {
		fmt.Println("Valid hex color")
	}
	// Output: Valid hex color
}

// ExampleValidator_RGB demonstrates RGB color validation
// ExampleValidator_RGB는 RGB 색상 검증을 보여줍니다
func ExampleValidator_RGB() {
	color := "rgb(255, 87, 51)"
	v := validation.New(color, "background_color")
	v.RGB()

	if len(v.GetErrors()) > 0 {
		fmt.Println("Invalid RGB color")
	} else {
		fmt.Println("Valid RGB color")
	}
	// Output: Valid RGB color
}

// ExampleValidator_RGBA demonstrates RGBA color validation
// ExampleValidator_RGBA는 RGBA 색상 검증을 보여줍니다
func ExampleValidator_RGBA() {
	color := "rgba(255, 87, 51, 0.8)"
	v := validation.New(color, "overlay_color")
	v.RGBA()

	if len(v.GetErrors()) > 0 {
		fmt.Println("Invalid RGBA color")
	} else {
		fmt.Println("Valid RGBA color")
	}
	// Output: Valid RGBA color
}

// ExampleValidator_HSL demonstrates HSL color validation
// ExampleValidator_HSL는 HSL 색상 검증을 보여줍니다
func ExampleValidator_HSL() {
	color := "hsl(9, 100%, 60%)"
	v := validation.New(color, "theme_color")
	v.HSL()

	if len(v.GetErrors()) > 0 {
		fmt.Println("Invalid HSL color")
	} else {
		fmt.Println("Valid HSL color")
	}
	// Output: Valid HSL color
}

// Example_colorValidation demonstrates color validation
// Example_colorValidation는 색상 검증을 보여줍니다
func Example_colorValidation() {
	mv := validation.NewValidator()

	// Validate hex color
	mv.Field("#FF5733", "primary_color").HexColor()

	// Validate RGB color
	mv.Field("rgb(255, 87, 51)", "secondary_color").RGB()

	// Validate RGBA color
	mv.Field("rgba(255, 87, 51, 0.8)", "overlay_color").RGBA()

	// Validate HSL color
	mv.Field("hsl(9, 100%, 60%)", "accent_color").HSL()

	err := mv.Validate()
	if err != nil {
		fmt.Println("Invalid colors")
	} else {
		fmt.Println("All colors valid")
	}
	// Output: All colors valid
}

// ExampleValidator_ASCII demonstrates ASCII character validation.
// ExampleValidator_ASCII는 ASCII 문자 검증을 보여줍니다.
func ExampleValidator_ASCII() {
	v := validation.New("Hello World 123", "text")
	v.ASCII()

	err := v.Validate()
	if err != nil {
		fmt.Println("Invalid ASCII")
	} else {
		fmt.Println("Valid ASCII")
	}
	// Output: Valid ASCII
}

// ExampleValidator_Printable demonstrates printable ASCII character validation.
// ExampleValidator_Printable는 인쇄 가능한 ASCII 문자 검증을 보여줍니다.
func ExampleValidator_Printable() {
	v := validation.New("Hello World! 123", "text")
	v.Printable()

	err := v.Validate()
	if err != nil {
		fmt.Println("Invalid printable")
	} else {
		fmt.Println("Valid printable")
	}
	// Output: Valid printable
}

// ExampleValidator_Whitespace demonstrates whitespace-only validation.
// ExampleValidator_Whitespace는 공백 문자만 있는지 검증을 보여줍니다.
func ExampleValidator_Whitespace() {
	v := validation.New(" \t\n  ", "text")
	v.Whitespace()

	err := v.Validate()
	if err != nil {
		fmt.Println("Not whitespace")
	} else {
		fmt.Println("Valid whitespace")
	}
	// Output: Valid whitespace
}

// ExampleValidator_AlphaSpace demonstrates alpha+space validation.
// ExampleValidator_AlphaSpace는 문자와 공백만 있는지 검증을 보여줍니다.
func ExampleValidator_AlphaSpace() {
	v := validation.New("John Doe", "name")
	v.AlphaSpace()

	err := v.Validate()
	if err != nil {
		fmt.Println("Invalid name")
	} else {
		fmt.Println("Valid name")
	}
	// Output: Valid name
}

// Example_dataFormatValidation demonstrates multiple data format validators.
// Example_dataFormatValidation는 여러 데이터 형식 검증기를 보여줍니다.
func Example_dataFormatValidation() {
	mv := validation.NewValidator()

	// Validate ASCII text
	mv.Field("Hello World", "ascii_text").ASCII()

	// Validate printable text (no control characters)
	mv.Field("Display Text!", "display_text").Printable()

	// Validate whitespace
	mv.Field("   ", "spacing").Whitespace()

	// Validate name with spaces
	mv.Field("John Doe", "full_name").AlphaSpace()

	err := mv.Validate()
	if err != nil {
		fmt.Println("Invalid data")
	} else {
		fmt.Println("All data valid")
	}
	// Output: All data valid
}

// ExampleValidator_OneOf demonstrates OneOf validation.
// ExampleValidator_OneOf는 OneOf 검증을 보여줍니다.
func ExampleValidator_OneOf() {
	v := validation.New("active", "status")
	v.OneOf("active", "inactive", "pending")

	err := v.Validate()
	if err != nil {
		fmt.Println("Invalid status")
	} else {
		fmt.Println("Valid status")
	}
	// Output: Valid status
}

// ExampleValidator_NotOneOf demonstrates NotOneOf validation.
// ExampleValidator_NotOneOf는 NotOneOf 검증을 보여줍니다.
func ExampleValidator_NotOneOf() {
	v := validation.New("user123", "username")
	v.NotOneOf("admin", "root", "administrator")

	err := v.Validate()
	if err != nil {
		fmt.Println("Forbidden username")
	} else {
		fmt.Println("Valid username")
	}
	// Output: Valid username
}

// ExampleValidator_When demonstrates conditional validation with When.
// ExampleValidator_When는 When을 사용한 조건부 검증을 보여줍니다.
func ExampleValidator_When() {
	email := "user@example.com"
	isRequired := true

	v := validation.New(email, "email")
	v.When(isRequired, func(val *validation.Validator) {
		val.Required().Email()
	})

	err := v.Validate()
	if err != nil {
		fmt.Println("Invalid email")
	} else {
		fmt.Println("Valid email")
	}
	// Output: Valid email
}

// ExampleValidator_Unless demonstrates conditional validation with Unless.
// ExampleValidator_Unless는 Unless를 사용한 조건부 검증을 보여줍니다.
func ExampleValidator_Unless() {
	email := "user@example.com"
	isGuest := false

	v := validation.New(email, "email")
	v.Unless(isGuest, func(val *validation.Validator) {
		val.Required().Email()
	})

	err := v.Validate()
	if err != nil {
		fmt.Println("Invalid email")
	} else {
		fmt.Println("Valid email")
	}
	// Output: Valid email
}

// Example_logicalValidation demonstrates multiple logical validators.
// Example_logicalValidation는 여러 논리 검증기를 보여줍니다.
func Example_logicalValidation() {
	mv := validation.NewValidator()

	// Validate status is one of allowed values
	mv.Field("active", "status").OneOf("active", "inactive", "pending")

	// Validate username is not forbidden
	mv.Field("user123", "username").NotOneOf("admin", "root")

	// Conditional validation
	isProduction := true
	mv.Field("prod-server", "server").When(isProduction, func(val *validation.Validator) {
		val.Required().MinLength(5)
	})

	// Inverse conditional validation
	isTestMode := false
	mv.Field("test@example.com", "email").Unless(isTestMode, func(val *validation.Validator) {
		val.Required().Email()
	})

	err := mv.Validate()
	if err != nil {
		fmt.Println("Invalid data")
	} else {
		fmt.Println("All validations passed")
	}
	// Output: All validations passed
}

// ExampleValidator_True demonstrates boolean true validation.
// ExampleValidator_True는 불리언 true 검증을 보여줍니다.
func ExampleValidator_True() {
	// Validate terms acceptance (must be true)
	accepted := true
	v := validation.New(accepted, "terms_accepted")
	v.True()

	err := v.Validate()
	if err != nil {
		fmt.Println("Terms must be accepted")
	} else {
		fmt.Println("Terms accepted")
	}
	// Output: Terms accepted
}

// ExampleValidator_False demonstrates boolean false validation.
// ExampleValidator_False는 불리언 false 검증을 보여줍니다.
func ExampleValidator_False() {
	// Validate newsletter opt-out (must be false)
	optIn := false
	v := validation.New(optIn, "newsletter_opt_in")
	v.False()

	err := v.Validate()
	if err != nil {
		fmt.Println("Must opt out of newsletter")
	} else {
		fmt.Println("Newsletter opt-out confirmed")
	}
	// Output: Newsletter opt-out confirmed
}

// ExampleValidator_Nil demonstrates nil value validation.
// ExampleValidator_Nil는 nil 값 검증을 보여줍니다.
func ExampleValidator_Nil() {
	// Validate optional field is nil
	var optionalField *string
	v := validation.New(optionalField, "optional_field")
	v.Nil()

	err := v.Validate()
	if err != nil {
		fmt.Println("Field should be nil")
	} else {
		fmt.Println("Field is correctly nil")
	}
	// Output: Field is correctly nil
}

// ExampleValidator_NotNil demonstrates non-nil value validation.
// ExampleValidator_NotNil는 non-nil 값 검증을 보여줍니다.
func ExampleValidator_NotNil() {
	// Validate required pointer is not nil
	value := "test"
	ptr := &value
	v := validation.New(ptr, "required_ptr")
	v.NotNil()

	err := v.Validate()
	if err != nil {
		fmt.Println("Pointer must not be nil")
	} else {
		fmt.Println("Pointer has a value")
	}
	// Output: Pointer has a value
}

// ExampleValidator_Type demonstrates type validation.
// ExampleValidator_Type는 타입 검증을 보여줍니다.
func ExampleValidator_Type() {
	// Validate value is a string
	text := "hello"
	v := validation.New(text, "text")
	v.Type("string")

	err := v.Validate()
	if err != nil {
		fmt.Println("Must be a string")
	} else {
		fmt.Println("Correct type")
	}
	// Output: Correct type
}

// ExampleValidator_Empty demonstrates empty value validation.
// ExampleValidator_Empty는 빈 값 검증을 보여줍니다.
func ExampleValidator_Empty() {
	// Validate optional field is empty
	optionalField := ""
	v := validation.New(optionalField, "optional_field")
	v.Empty()

	err := v.Validate()
	if err != nil {
		fmt.Println("Field should be empty")
	} else {
		fmt.Println("Field is empty as expected")
	}
	// Output: Field is empty as expected
}

// ExampleValidator_NotEmpty demonstrates non-empty value validation.
// ExampleValidator_NotEmpty는 비어있지 않은 값 검증을 보여줍니다.
func ExampleValidator_NotEmpty() {
	// Validate required field is not empty
	requiredField := "value"
	v := validation.New(requiredField, "required_field")
	v.NotEmpty()

	err := v.Validate()
	if err != nil {
		fmt.Println("Field must not be empty")
	} else {
		fmt.Println("Field has a value")
	}
	// Output: Field has a value
}

// ExampleValidator_BetweenTime demonstrates time range validation.
// ExampleValidator_BetweenTime는 시간 범위 검증을 보여줍니다.
func ExampleValidator_BetweenTime() {
	// Define the valid time range for 2024
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)

	// Validate a date within the range
	eventDate := time.Date(2024, 6, 15, 12, 0, 0, 0, time.UTC)
	v := validation.New(eventDate, "event_date")
	v.BetweenTime(start, end)

	err := v.Validate()
	if err != nil {
		fmt.Println("Date is out of range")
	} else {
		fmt.Println("Date is within 2024")
	}
	// Output: Date is within 2024
}

// ExampleValidator_WithCustomMessage demonstrates setting a single custom error message.
// ExampleValidator_WithCustomMessage는 단일 커스텀 에러 메시지 설정을 보여줍니다.
func ExampleValidator_WithCustomMessage() {
	v := validation.New("", "email")
	v.WithCustomMessage("required", "Please enter your email address")
	v.Required()

	err := v.Validate()
	if err != nil {
		fmt.Println(err.Error())
	}
	// Output: Please enter your email address
}

// ExampleValidator_WithCustomMessages demonstrates setting multiple custom error messages.
// ExampleValidator_WithCustomMessages는 여러 커스텀 에러 메시지 설정을 보여줍니다.
func ExampleValidator_WithCustomMessages() {
	v := validation.New("ab", "password")
	v.WithCustomMessages(map[string]string{
		"required":  "비밀번호를 입력해주세요",
		"minlength": "비밀번호는 8자 이상이어야 합니다",
		"maxlength": "비밀번호는 20자 이하여야 합니다",
	})
	v.Required().MinLength(8).MaxLength(20)

	err := v.Validate()
	if err != nil {
		errors := v.GetErrors()
		fmt.Println(errors[0].Message)
	}
	// Output: 비밀번호는 8자 이상이어야 합니다
}

// ExampleValidator_WithCustomMessages_multiValidator demonstrates custom messages with MultiValidator.
// ExampleValidator_WithCustomMessages_multiValidator는 MultiValidator에서 커스텀 메시지를 보여줍니다.
func ExampleValidator_WithCustomMessages_multiValidator() {
	mv := validation.NewValidator()

	mv.Field("", "email").WithCustomMessages(map[string]string{
		"required": "Email is required",
		"email":    "Invalid email format",
	}).Required()

	mv.Field("short", "password").WithCustomMessages(map[string]string{
		"required":  "Password is required",
		"minlength": "Password must be at least 8 characters",
	}).Required().MinLength(8)

	err := mv.Validate()
	if err != nil {
		errors := mv.GetErrors()
		for _, e := range errors {
			fmt.Printf("%s: %s\n", e.Field, e.Message)
		}
	}
	// Output:
	// email: Email is required
	// password: Password must be at least 8 characters
}
