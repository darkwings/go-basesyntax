package errors

import (
	"errors"
	"fmt"
	"os"
)

// ====================================================
// os.PathError
// ====================================================

func OsErrors() {
	_, err := os.Open("/no/such/file")
	fmt.Println(os.IsNotExist(err)) // "true"
}

func BaseError(x int) (int, error) {
	if x != 0 {
		return x * 2, nil
	} else {
		return 0, errors.New("cannot call with 0")
	}
}

// Custom error

type MyCustomError struct {
	message string
}

func (m *MyCustomError) Error() string {
	return m.message
}

func doingStuffAndGenerateCustomError(x int) (int, error) {
	if x != 0 {
		return x * 2, nil
	} else {
		return 0, &MyCustomError{"this is a custom error for a wrong division"}
	}
}

func ManageCustomError(x int) int {
	res, err := doingStuffAndGenerateCustomError(x)
	if err != nil {
		me, ok := err.(*MyCustomError) // error is MyCustomError?
		if ok {
			// doing something smarter with the custom error
			fmt.Print(me.Error())
		} else {
			// manage the standard error
			fmt.Print(err.Error())
		}
		return -666
	} else {
		return res
	}
}
