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
		return 0, errors.New("Cannot call with 0")
	}
}
