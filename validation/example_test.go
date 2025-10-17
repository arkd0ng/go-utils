package validation_test

import (
	"fmt"
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
