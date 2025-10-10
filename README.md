# go-utils

A collection of frequently used utility functions for Golang development.

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.16-blue)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Installation

```bash
go get github.com/arkd0ng/go-utils
```

## Features

### Random String Generation

Generate random strings with various character sets and customizable length ranges.

#### Available Methods

- **`Alpha(min, max int)`** - Alphabetic characters only (a-z, A-Z)
- **`AlphaNum(min, max int)`** - Alphanumeric characters (a-z, A-Z, 0-9)
- **`AlphaNumSpecial(min, max int)`** - Alphanumeric + all special characters
- **`AlphaNumSpecialLimited(min, max int)`** - Alphanumeric + limited special characters (!@#$%^&*-_)
- **`Custom(charset string, min, max int)`** - Custom character set

## Usage

### Basic Examples

```go
package main

import (
    "fmt"
    "github.com/arkd0ng/go-utils"
)

func main() {
    // Generate alphabetic string (32-128 characters)
    str1 := goutils.GenRandomString.Alpha(32, 128)
    fmt.Println(str1)

    // Generate alphanumeric string (32-128 characters)
    str2 := goutils.GenRandomString.AlphaNum(32, 128)
    fmt.Println(str2)

    // Generate string with special characters (16-32 characters)
    str3 := goutils.GenRandomString.AlphaNumSpecial(16, 32)
    fmt.Println(str3)

    // Generate string with limited special characters (20-40 characters)
    str4 := goutils.GenRandomString.AlphaNumSpecialLimited(20, 40)
    fmt.Println(str4)

    // Generate string with custom character set (10-20 characters)
    str5 := goutils.GenRandomString.Custom("ABC123xyz", 10, 20)
    fmt.Println(str5)
}
```

### Fixed Length String

To generate a string with a fixed length, set `min` and `max` to the same value:

```go
// Generate exactly 32 characters
password := goutils.GenRandomString.AlphaNum(32, 32)
```

### Common Use Cases

```go
// Generate a secure password
password := goutils.GenRandomString.AlphaNumSpecial(16, 24)

// Generate a random API key
apiKey := goutils.GenRandomString.AlphaNum(40, 40)

// Generate a random username
username := goutils.GenRandomString.Alpha(8, 12)

// Generate a verification code with numbers only
code := goutils.GenRandomString.Custom("0123456789", 6, 6)
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
go test -v
```

Run benchmarks:

```bash
go test -bench=.
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

- String manipulation utilities
- Slice/Array helpers
- Map utilities
- File/Path utilities
- Error handling helpers
- Time/Date utilities
- JSON/Struct conversion helpers
- HTTP utilities
- Concurrency helpers
- Validation utilities

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

**arkd0ng**

- GitHub: [@arkd0ng](https://github.com/arkd0ng)

## Changelog

### v0.1.0 (Initial Release)

- Added `GenRandomString` with support for:
  - Alpha (alphabetic only)
  - AlphaNum (alphanumeric)
  - AlphaNumSpecial (with all special characters)
  - AlphaNumSpecialLimited (with limited special characters)
  - Custom (user-defined character set)
