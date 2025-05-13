// Package atoErr provides utilities for creating and handling structured errors in Go.
//
// The package allows creation of errors with associated names for better error identification
// and handling. It provides functions to create named errors and extract error names from
// existing errors.
//
// Example usage:
//
//	err := atoErr.Errorf("ValidationError", "invalid input: %s", "missing field")
//	name := atoErr.GetName(err) // Returns "ValidationError"
package atoErr

import "fmt"

const (
	NIL                   string = "NIL"
	ErrorInvalidErrorType string = "ErrorInvalidErrorType"
)

type errorStruct struct {
	name    string
	message string
}

// Error returns a string representation of the error.
// It formats the error name and message into a string using the format "name : message\n".
func (e *errorStruct) Error() string {
	return fmt.Sprintf("%v : %v\n", e.name, e.message)
}

// GetName extracts and returns the error name from the provided error.
// If the error is nil, it returns NIL.
// If the error is not of type *errorStruct, it returns ErrorInvalidErrorType.
func GetName(err error) string {
	if err == nil {
		return NIL
	}
	switch err := err.(type) {
	case *errorStruct:
		return err.name
	}
	return ErrorInvalidErrorType
}

// Errorf creates a new error with a specified name and formatted message.
// It takes a name string, a format string, and variadic arguments similar to fmt.Sprintf.
// The returned error contains both the name identifier and the formatted message.
//
// Parameters:
//   - name: A string identifying the type or category of the error
//   - format: A format string as used in fmt.Sprintf
//   - a: Variadic arguments to be formatted into the message
//
// Returns:
//   - error: A new error instance containing the name and formatted message
func Errorf(name string, format string, a ...any) error {
	return &errorStruct{
		name:    name,
		message: fmt.Sprintf(format, a...),
	}
}
