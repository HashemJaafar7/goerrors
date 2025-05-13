package atoErr_test

import (
	"fmt"
	"testing"

	"github.com/HashemJaafar7/atoErr"
	test "github.com/HashemJaafar7/go_test"
)

func ExampleErrorf() {
	// Create a new error with a name and message
	err := atoErr.Errorf("ValidationError", "invalid field %q", "email")
	fmt.Println(err)
	// Output: ValidationError : invalid field "email"
}

func ExampleGetName() {
	// Create an error and get its name
	err := atoErr.Errorf("DatabaseError", "connection failed")
	name := atoErr.GetName(err)
	fmt.Println(name)
	// Output: DatabaseError
}

func TestErrorf(t *testing.T) {
	{
		err := atoErr.Errorf("DatabaseError", "connection failed")

		errString := err.Error()
		test.Test(false, true, "#v", errString, "DatabaseError : connection failed\n")

		name := atoErr.GetName(err)
		test.Test(false, true, "v", name, "DatabaseError")
	}

	{
		err := atoErr.Errorf("DatabaseError", "connection failed")
		err = nil

		name := atoErr.GetName(err)
		test.Test(false, true, "v", name, atoErr.NIL)
	}

	{
		err := fmt.Errorf("DatabaseError : connection failed")

		errString := err.Error()
		test.Test(false, true, "v", errString, "DatabaseError : connection failed")

		name := atoErr.GetName(err)
		test.Test(false, true, "v", name, atoErr.ErrorInvalidErrorType)
	}
}
