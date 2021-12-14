package interfaces

import (
	"errors"
	"fmt"
)

func TypeAssertionsExample() {
	fmt.Printf("%s\n", formatOneValue(2))                        // Non scrivera nulla
	fmt.Printf("%s\n", formatOneValue(errors.New("an error")))   // err: an error
	fmt.Printf("%s\n", formatOneValue(MyStringer{"some value"})) // some value
}

type Stringer interface {
	String() string
}

type MyStringer struct {
	text string
}

func (s MyStringer) String() string {
	return s.text
}

func formatOneValue(x interface{}) string {
	if err, ok := x.(error); ok {
		return "err: " + err.Error()
	}
	if str, ok := x.(Stringer); ok { // ok := x.(Stringer), verifica se il tipo (dynamic type) di x corrisponde a Stringer
		return str.String()
	}
	return ""
}
