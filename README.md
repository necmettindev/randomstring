# Random String Generator

This library assists in generating random strings based on specific characteristics. You can create random strings consisting of alphanumeric characters, letters only, digits only, special characters, or even a custom character set.

## Features

- Generate strings with alphanumeric, letters only, digits only, or special characters.
- Ability to disable uppercase, lowercase, or numeric characters.
- Option to use a custom character set.

## Installation

```bash
go get -u github.com/necmettindev/randomstring
```

## Usage

Below are different scenarios demonstrating how to use the `randomstring` package:

### Basic Usage

Generate a string of length 10 using the default alphanumeric characters:

```go
package main

import (
"fmt"
"github.com/necmettindev/randomstring"
)

func main() {
    opts := randomstring.GenerationOptions{
        Length: 10,
    }

    result, _ := randomstring.GenerateString(opts)

    fmt.Println(result)
}
```

### Disabling Features

For instance, to generate a string containing letters only:

```go
    opts := randomstring.GenerationOptions{
        Length: 10,
        DisableNumeric: true,
    }

    result, _ := randomstring.GenerateString(opts)
    fmt.Println(result)
```

### Using Special Characters

Generate a string with enabled special characters:

```go
    opts := randomstring.GenerationOptions{
        Length: 15,
        EnableSpecialCharacter: true,
    }

    result, _ := randomstring.GenerateString(opts)
    fmt.Println(result)
```

### Using a Custom Character Set

Generate a string with your own custom character set:

```go
    opts := randomstring.GenerationOptions{
        Length: 12,
        CustomCharset: "abcXYZ789",
    }

    result, _ := randomstring.GenerateString(opts)
    fmt.Println(result) // Produces a string consisting only of 'a', 'b', 'c', 'X', 'Y', 'Z', '7', '8', '9'.
```

## Contributing

If you wish to contribute to this project, please first open an issue or pick up an existing one. Then make your contribution and submit a pull request.

## License

This project is licensed under the MIT License.
