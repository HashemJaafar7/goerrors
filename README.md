# atoErr

A Go package for creating and handling structured errors with names.

## Installation

```bash
go get github.com/HashemJaafar7/atoErr
```

## Usage

```go
import "github.com/HashemJaafar7/atoErr"

// Create an error and get its name
err := atoErr.Errorf("DatabaseError", "connection failed")
name := atoErr.GetName(err)
fmt.Println(name)
// Output: DatabaseError
```

## Features

- Create errors with associated names
- Extract error names from errors
- Special handling for nil errors
- Type-safe error name extraction

## Documentation

For detailed documentation, visit [pkg.go.dev](https://pkg.go.dev/github.com/HashemJaafar7/atoErr)

## License

MIT License

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
