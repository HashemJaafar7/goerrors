package goerrors_test

import (
	"fmt"
	"testing"

	"github.com/HashemJaafar7/goerrors"
	"github.com/HashemJaafar7/testutils"
)

func fTest[t any](actual t, expected t) {
	testutils.Test(true, false, true, 10, "v", actual, expected)
}

func ExampleErrorf() {
	// Create a new error with a name and message
	err := goerrors.Errorf("ValidationError", "invalid field %q", "email")
	fmt.Println(err)
	// Output: ValidationError : invalid field "email"
}

func ExampleGetName() {
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
	// goroutine 1 [running]:
	// runtime/debug.Stack()
	//         C:/Program Files/Go/src/runtime/debug/stack.go:26 +0x5e
	// github.com/HashemJaafar7/goerrors.Errorf({0x429b63, 0xd}, {0x42ae23, 0x11}, {0x0, 0x0, 0x0})
	//         c:/Users/hashem/Desktop/libraries/goerrors/lib.go:65 +0x57
	// github.com/HashemJaafar7/goerrors_test.ExampleGetName()
	//         c:/Users/hashem/Desktop/libraries/goerrors/lib_test.go:24 +0x3f
	// testing.runExample({{0x429fce, 0xe}, 0x436df0, {0x43658c, 0x358}, 0x0})
	//         C:/Program Files/Go/src/testing/run_example.go:63 +0x2b0
	// testing.runExamples(0x467240?, {0x567d40, 0x2, 0x2?})
	//         C:/Program Files/Go/src/testing/example.go:41 +0x125
	// testing.(*M).Run(0xc0000581e0)
	//         C:/Program Files/Go/src/testing/testing.go:2144 +0x71b
	// main.main()
	//         _testmain.go:53 +0x9b
}

func TestErrorf(t *testing.T) {
	{
		err := goerrors.Errorf("DatabaseError", "connection failed")

		errString := err.Error()
		fTest(errString, "DatabaseError : connection failed")

		name := goerrors.GetName(err)
		fTest(name, "DatabaseError")
	}

	{
		err := goerrors.Errorf("DatabaseError", "connection failed")
		err = nil

		name := goerrors.GetName(err)
		fTest(name, goerrors.NIL)
	}

	{
		err := fmt.Errorf("DatabaseError : connection failed")

		errString := err.Error()
		fTest(errString, "DatabaseError : connection failed")

		name := goerrors.GetName(err)
		fTest(name, goerrors.ErrorInvalidErrorType)
	}
}

func foo() error {
	return goerrors.Errorf("TestError", "test message")
}

func Test_GetStack(t *testing.T) {
	// Case 1: nil error
	{
		var err error
		actual := goerrors.GetStack(err)
		expected := goerrors.NIL
		fTest(actual, expected)
	}

	// Case 2: error created with goerrors.Errorf
	{
		err := goerrors.Errorf("TestError", "test message")
		actual := goerrors.GetStack(err)
		// Stack trace will be non-empty string
		fmt.Println(actual)
	}

	// Case 3: standard error
	{
		err := fmt.Errorf("standard error")
		actual := goerrors.GetStack(err)
		expected := goerrors.ErrorInvalidErrorType
		fTest(actual, expected)
	}

	// Case 3: standard error
	{
		err := foo()
		actual := goerrors.GetStack(err)
		fmt.Println(actual)
	}
}
