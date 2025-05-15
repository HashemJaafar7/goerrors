// Package goerrors provides utilities for creating and handling structured errors in Go.
//
// The package allows creation of errors with associated names for better error identification
// and handling. It provides functions to create named errors and extract error names from
// existing errors  and extract error stack
//
// Example usage:
//
//	err := goerrors.Errorf("ValidationError", "invalid input: %s", "missing field")
//	name := goerrors.GetName(err) // Returns "ValidationError"
//	fmt.Println(err.Error())
package goerrors

import (
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/HashemJaafar7/testutils"
)

const (
	NIL                   string = "NIL"
	ErrorInvalidErrorType string = "ErrorInvalidErrorType"
)

type errorStruct struct {
	name    string
	message string
	stack   string
}

// Error returns a string representation of the error.
// It formats the error name and message into a string using the format "name : message\n".
func (e *errorStruct) Error() string {
	return fmt.Sprintf("%v : %v", e.name, e.message)
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

// Errorf creates a new error with a custom name, formatted message, and stack trace.
// It takes a name for the error, a format string, and optional arguments for formatting.
// The resulting error includes:
//   - A custom name identifier
//   - A formatted error message
//   - A colorized stack trace where .go file references are highlighted in red
//
// Parameters:
//   - name: String identifier for the error type
//   - format: Printf-style format string for the error message
//   - a: Optional variadic arguments used in format string
//
// Returns:
//   - error: A new error instance containing the name, formatted message and stack trace
func Errorf(name string, format string, a ...any) error {
	stack := string(debug.Stack())
	stackSlice := strings.Split(stack, "\n")

	var resultStack string
	for _, line := range stackSlice {
		if strings.Contains(line, ".go:") {
			resultStack += testutils.ColorRed + line + testutils.ColorReset + "\n"
		} else {
			resultStack += line + "\n"
		}
	}

	return &errorStruct{
		name:    name,
		message: fmt.Sprintf(format, a...),
		stack:   resultStack,
	}
}

// GetStack retrieves the stack trace from an error.
// Returns:
//   - The stack trace string if the error is of type *errorStruct
//   - "nil" if the error is nil
//   - "invalid error type" if the error is not of type *errorStruct
//
// This function is used internally to extract stack trace information from custom errors.
func GetStack(err error) string {
	if err == nil {
		return NIL
	}
	switch err := err.(type) {
	case *errorStruct:
		return err.stack
	}
	return ErrorInvalidErrorType
}
