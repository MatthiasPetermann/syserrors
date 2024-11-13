# syserrors

syserrors is a Go package that provides structured error definitions with
standardized exit status codes, inspired by the traditional UNIX sysexits.h
exit codes. This package helps you manage and handle different types of errors
in your Go applications, providing clear and consistent exit codes to the
operating system upon failure.

## Features

- Predefined error types for common error scenarios (e.g., configuration
errors, I/O errors, remote errors).
- Integration with standard UNIX sysexits codes for consistent exit status
management.
- Easily extendable to cover additional error categories as needed.
- Simplifies error handling in CLI tools and system-level applications.

## Installation

To install the package, use:

```
go get github.com/MatthiasPetermann/syserrors
```

## Usage

### Importing the Package

```
import (
    "fmt"
    "os"
    "github.com/MatthiasPetermann/syserrors"
)
```

### Example: Handling Errors with Exit Codes

```
package main

import (
    "fmt"
    "os"
    "github.com/MatthiasPetermann/syserrors"
)

func performAction() error {
    // Example: Return a configuration error
    return &syserrors.ConfigurationError{Msg: "Invalid configuration detected"}
}

func main() {
    if err := performAction(); err != nil {
        if codedErr, ok := err.(syserrors.CodedError); ok {
            fmt.Printf("Error: %s\n", codedErr.Error())
            os.Exit(codedErr.Code())
        } else {
            fmt.Println("Unknown error:", err)
            os.Exit(1)
        }
    }
    fmt.Println("Action completed successfully")
}
```

### Defining and Handling Custom Errors

You can extend the package by creating your own error types based on the CodedError interface:

```
type MyCustomError struct {
    Msg string
}

func (e *MyCustomError) Error() string {
    return e.Msg
}

func (e *MyCustomError) Code() int {
    return sysexits.EX_SOFTWARE // Custom exit code
}
```
## License

This project is licensed under the MIT License - see the LICENSE file for
details.
