package main

import (
	"fmt"
	"time"

	"github.com/arkd0ng/go-utils/validation"
)

func main() {
	fmt.Println("=== Validation Package Examples ===")
	fmt.Println()

	// Example 1: Simple String Validation
	example1SimpleStringValidation()

	// Example 2: Numeric Validation
	example2NumericValidation()

	// Example 3: Collection Validation
	example3CollectionValidation()

	// Example 4: Comparison Validation
	example4ComparisonValidation()

	// Example 5: Multi-Field Validation
	example5MultiFieldValidation()

	// Example 6: User Registration
	example6UserRegistration()

	// Example 7: Custom Validators
	example7CustomValidators()

	// Example 8: Stop on First Error
	example8StopOnFirstError()
}

func example1SimpleStringValidation() {
	fmt.Println("=== Example 1: Simple String Validation ===")

	// Valid email
	v := validation.New("john@example.com", "email")
	v.Required().Email()
	if err := v.Validate(); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Println("✅ Valid email")
	}

	// Invalid email
	v = validation.New("invalid-email", "email")
	v.Required().Email()
	if err := v.Validate(); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	}

	// String length validation
	v = validation.New("Hello", "message")
	v.MinLength(3).MaxLength(10)
	if err := v.Validate(); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Println("✅ Valid message length")
	}

	fmt.Println()
}

func example2NumericValidation() {
	fmt.Println("=== Example 2: Numeric Validation ===")

	// Age validation
	v := validation.New(25, "age")
	v.Positive().Min(18).Max(120)
	if err := v.Validate(); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Println("✅ Valid age")
	}

	// Score validation
	v = validation.New(85, "score")
	v.Between(0, 100).Positive()
	if err := v.Validate(); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Println("✅ Valid score")
	}

	// Even number validation
	v = validation.New(10, "value")
	v.Even().MultipleOf(5)
	if err := v.Validate(); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Println("✅ Valid even number")
	}

	fmt.Println()
}

func example3CollectionValidation() {
	fmt.Println("=== Example 3: Collection Validation ===")

	// Array validation
	numbers := []int{1, 2, 3, 4, 5}
	v := validation.New(numbers, "numbers")
	v.ArrayNotEmpty().ArrayMinLength(3).ArrayUnique()
	if err := v.Validate(); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Println("✅ Valid array")
	}

	// Array with duplicates
	duplicates := []int{1, 2, 2, 3}
	v = validation.New(duplicates, "duplicates")
	v.ArrayUnique()
	if err := v.Validate(); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	}

	// Map validation
	data := map[string]int{"name": 1, "age": 25, "score": 85}
	v = validation.New(data, "data")
	v.MapNotEmpty().MapHasKeys("name", "age")
	if err := v.Validate(); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Println("✅ Valid map")
	}

	// In validation
	v = validation.New("admin", "role")
	v.In("admin", "moderator", "user")
	if err := v.Validate(); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Println("✅ Valid role")
	}

	fmt.Println()
}

func example4ComparisonValidation() {
	fmt.Println("=== Example 4: Comparison Validation ===")

	// Numeric comparison
	v := validation.New(50, "score")
	v.GreaterThan(0).LessThan(100)
	if err := v.Validate(); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Println("✅ Valid score range")
	}

	// Time comparison
	now := time.Now()
	yesterday := now.Add(-24 * time.Hour)
	tomorrow := now.Add(24 * time.Hour)

	v = validation.New(now, "date")
	v.After(yesterday).Before(tomorrow)
	if err := v.Validate(); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Println("✅ Valid date range")
	}

	// Equality check
	password := "secret123"
	confirmPassword := "secret123"
	v = validation.New(password, "password")
	v.Equals(confirmPassword)
	if err := v.Validate(); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Println("✅ Passwords match")
	}

	fmt.Println()
}

func example5MultiFieldValidation() {
	fmt.Println("=== Example 5: Multi-Field Validation ===")

	// Validate multiple fields at once
	username := "john_doe"
	email := "john@example.com"
	age := 25

	mv := validation.NewValidator()

	mv.Field(username, "username").
		Required().
		MinLength(3).
		MaxLength(20).
		Alphanumeric()

	mv.Field(email, "email").
		Required().
		Email().
		MaxLength(100)

	mv.Field(age, "age").
		Positive().
		Between(18, 120)

	if err := mv.Validate(); err != nil {
		fmt.Printf("❌ Validation errors:\n")
		errors := err.(validation.ValidationErrors)
		for _, e := range errors {
			fmt.Printf("  - %s: %s\n", e.Field, e.Message)
		}
	} else {
		fmt.Println("✅ All fields valid")
	}

	fmt.Println()
}

type User struct {
	Username string
	Email    string
	Password string
	Age      int
	Country  string
	Tags     []string
}

func example6UserRegistration() {
	fmt.Println("=== Example 6: User Registration Validation ===")

	user := User{
		Username: "john_doe",
		Email:    "john@example.com",
		Password: "Secure123!",
		Age:      25,
		Country:  "KR",
		Tags:     []string{"developer", "golang"},
	}

	if err := ValidateUser(user); err != nil {
		fmt.Printf("❌ Validation errors:\n")
		errors := err.(validation.ValidationErrors)
		for _, e := range errors {
			fmt.Printf("  - %s: %s\n", e.Field, e.Message)
		}
	} else {
		fmt.Println("✅ User registration valid")
	}

	// Invalid user
	invalidUser := User{
		Username: "ab",
		Email:    "invalid-email",
		Password: "weak",
		Age:      15,
		Country:  "XX",
		Tags:     []string{},
	}

	if err := ValidateUser(invalidUser); err != nil {
		fmt.Printf("❌ Invalid user - validation errors:\n")
		errors := err.(validation.ValidationErrors)
		for _, e := range errors {
			fmt.Printf("  - %s: %s\n", e.Field, e.Message)
		}
	}

	fmt.Println()
}

func ValidateUser(user User) error {
	mv := validation.NewValidator()

	mv.Field(user.Username, "username").
		Required().
		MinLength(3).
		MaxLength(20).
		Alphanumeric()

	mv.Field(user.Email, "email").
		Required().
		Email().
		MaxLength(100)

	mv.Field(user.Password, "password").
		Required().
		MinLength(8).
		MaxLength(100)

	mv.Field(user.Age, "age").
		Positive().
		Between(18, 120)

	mv.Field(user.Country, "country").
		Required().
		In("US", "KR", "JP", "CN", "UK", "FR", "DE")

	mv.Field(user.Tags, "tags").
		ArrayNotEmpty().
		ArrayMinLength(1).
		ArrayMaxLength(10)

	return mv.Validate()
}

func example7CustomValidators() {
	fmt.Println("=== Example 7: Custom Validators ===")

	// Custom password strength validator
	password := "Secure123!"
	v := validation.New(password, "password")
	v.Required().
		MinLength(8).
		Custom(func(val interface{}) bool {
			s := val.(string)
			// Check for at least one uppercase, one lowercase, one digit, and one special char
			hasUpper := false
			hasLower := false
			hasDigit := false
			hasSpecial := false
			for _, r := range s {
				if r >= 'A' && r <= 'Z' {
					hasUpper = true
				}
				if r >= 'a' && r <= 'z' {
					hasLower = true
				}
				if r >= '0' && r <= '9' {
					hasDigit = true
				}
				if string(r) == "!" || string(r) == "@" || string(r) == "#" || string(r) == "$" {
					hasSpecial = true
				}
			}
			return hasUpper && hasLower && hasDigit && hasSpecial
		}, "Password must contain uppercase, lowercase, digit, and special character")

	if err := v.Validate(); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Println("✅ Strong password")
	}

	// Weak password
	weakPassword := "simple"
	v = validation.New(weakPassword, "password")
	v.MinLength(8).Custom(func(val interface{}) bool {
		s := val.(string)
		hasUpper := false
		hasLower := false
		hasDigit := false
		for _, r := range s {
			if r >= 'A' && r <= 'Z' {
				hasUpper = true
			}
			if r >= 'a' && r <= 'z' {
				hasLower = true
			}
			if r >= '0' && r <= '9' {
				hasDigit = true
			}
		}
		return hasUpper && hasLower && hasDigit
	}, "Password must contain uppercase, lowercase, and digit")

	if err := v.Validate(); err != nil {
		fmt.Printf("❌ Weak password: %v\n", err)
	}

	fmt.Println()
}

func example8StopOnFirstError() {
	fmt.Println("=== Example 8: Stop on First Error ===")

	// Without StopOnError - collects all errors
	v := validation.New("", "email")
	v.Required().Email().MaxLength(100)
	if err := v.Validate(); err != nil {
		fmt.Printf("❌ Without StopOnError: %v\n", err)
	}

	// With StopOnError - stops at first error
	v = validation.New("", "email")
	v.StopOnError().Required().Email().MaxLength(100)
	if err := v.Validate(); err != nil {
		fmt.Printf("❌ With StopOnError: %v\n", err)
	}

	fmt.Println()
}
