# goerrors

A Go package for creating and handling structured errors with names.

## Installation

```bash
go get github.com/HashemJaafar7/goerrors
```

## Usage

```go
package main

import "github.com/HashemJaafar7/goerrors"
func main(){
	// Create an error and get its name
	err := goerrors.Errorf("DatabaseError", "connection failed")
	name := goerrors.GetName(err)
	stack := goerrors.GetStack(err)
	fmt.Println(err.Error())
	fmt.Println(name)
	fmt.Println(stack)
	// Output:
	// DatabaseError : connection failed
	// DatabaseError
	// c:/Users/hashem/Desktop/libraries/goerrors/lib_test.go:24 +0x3f
	// C:/Program Files/Go/src/testing/run_example.go:63 +0x2b0
	// C:/Program Files/Go/src/testing/example.go:41 +0x125
	// C:/Program Files/Go/src/testing/testing.go:2144 +0x71b
	// _testmain.go:53 +0x9b
}
```

## Features

- Create errors with associated names
- Extract error names from errors
- Extract error stack trace from errors
- Special handling for nil errors
- Type-safe error name extraction

## Documentation

For detailed documentation, visit [pkg.go.dev](https://pkg.go.dev/github.com/HashemJaafar7/goerrors)

## License

MIT License

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
