# go-utils

A collection of frequently used utility functions for Golang development.

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.16-blue)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Installation

```bash
go get github.com/arkd0ng/go-utils
```

## Package Structure

This library is organized into subpackages for better modularity and easier maintenance:

```
go-utils/
├── random/          # Random generation utilities
├── stringutil/      # String manipulation (coming soon)
├── sliceutil/       # Slice helpers (coming soon)
├── maputil/         # Map utilities (coming soon)
└── ...
```

## Features

### Random Package

Generate random strings with various character sets and customizable length ranges.

#### Available Methods

- **`random.GenString.Alpha(min, max int)`** - Alphabetic characters only (a-z, A-Z)
- **`random.GenString.AlphaNum(min, max int)`** - Alphanumeric characters (a-z, A-Z, 0-9)
- **`random.GenString.AlphaNumSpecial(min, max int)`** - Alphanumeric + all special characters
- **`random.GenString.AlphaNumSpecialLimited(min, max int)`** - Alphanumeric + limited special characters (!@#$%^&*-_)
- **`random.GenString.Custom(charset string, min, max int)`** - Custom character set

## Usage

### Basic Examples

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils/random"
)

func main() {
    // Generate alphabetic string (32-128 characters)
    str1 := random.GenString.Alpha(32, 128)
    fmt.Println(str1)

    // Generate alphanumeric string (32-128 characters)
    str2 := random.GenString.AlphaNum(32, 128)
    fmt.Println(str2)

    // Generate string with special characters (16-32 characters)
    str3 := random.GenString.AlphaNumSpecial(16, 32)
    fmt.Println(str3)

    // Generate string with limited special characters (20-40 characters)
    str4 := random.GenString.AlphaNumSpecialLimited(20, 40)
    fmt.Println(str4)

    // Generate string with custom character set (10-20 characters)
    str5 := random.GenString.Custom("ABC123xyz", 10, 20)
    fmt.Println(str5)
}
```

### Fixed Length String

To generate a string with a fixed length, set `min` and `max` to the same value:

```go
// Generate exactly 32 characters
password := random.GenString.AlphaNum(32, 32)
```

### Common Use Cases

```go
import "github.com/arkd0ng/go-utils/random"

// Generate a secure password
password := random.GenString.AlphaNumSpecial(16, 24)

// Generate a random API key
apiKey := random.GenString.AlphaNum(40, 40)

// Generate a random username
username := random.GenString.Alpha(8, 12)

// Generate a verification code with numbers only
code := random.GenString.Custom("0123456789", 6, 6)
```

## Character Sets

- **Alpha**: `A-Z`, `a-z`
- **AlphaNum**: `A-Z`, `a-z`, `0-9`
- **AlphaNumSpecial**: `A-Z`, `a-z`, `0-9`, `!@#$%^&*()-_=+[]{}|;:,.<>?/`
- **AlphaNumSpecialLimited**: `A-Z`, `a-z`, `0-9`, `!@#$%^&*-_`

## Security

The random string generator uses `crypto/rand` for cryptographically secure random generation, making it suitable for:

- Password generation
- API key generation
- Token generation
- Security-sensitive applications

## Testing

Run the test suite:

```bash
go test ./... -v
```

Run benchmarks:

```bash
go test ./... -bench=.
```

## Contributing

Contributions are welcome! This library will grow with frequently used utility functions. Feel free to:

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/new-utility`)
3. Commit your changes (`git commit -am 'Add new utility function'`)
4. Push to the branch (`git push origin feature/new-utility`)
5. Create a Pull Request

## Roadmap

Future utilities being considered:

- **stringutil/** - String manipulation utilities
- **sliceutil/** - Slice/Array helpers (filter, map, unique, etc.)
- **maputil/** - Map utilities (merge, keys, values, etc.)
- **fileutil/** - File/Path utilities
- **httputil/** - HTTP helpers
- **timeutil/** - Time/Date utilities
- **validation/** - Validation utilities (email, URL, etc.)
- **errorutil/** - Error handling helpers
- **concurrency/** - Concurrency helpers

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

**arkd0ng**

- GitHub: [@arkd0ng](https://github.com/arkd0ng)

## Changelog

### v0.2.0 (Current)

- **BREAKING CHANGE**: Refactored to subpackage structure
  - Moved `GenRandomString` to `random.GenString`
  - Import path changed from `github.com/arkd0ng/go-utils` to `github.com/arkd0ng/go-utils/random`
- Improved package organization for better scalability
- Prepared structure for future utility additions

### v0.1.0 (Initial Release)

- Added `GenRandomString` with support for:
  - Alpha (alphabetic only)
  - AlphaNum (alphanumeric)
  - AlphaNumSpecial (with all special characters)
  - AlphaNumSpecialLimited (with limited special characters)
  - Custom (user-defined character set)
