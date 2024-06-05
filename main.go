package main

import (
	"errors"
	"fmt"

	. "github.com/manwitha1000names/gofp/v2/MaybeResult"
)

type ErrorYes struct {
	msg string
}

func (e ErrorYes) Error() string {
	return e.msg
}

var (
	e1 = fmt.Errorf("What is this??")
	e2 = ErrorYes{"yes"}
)

func main() {
	err1 := Err[string](e1)
	err2 := Err[string](e2)
	if errors.Is(err1, e1) {
		fmt.Println("yup err1 == e1")
	}
	var y ErrorYes
	if errors.As(err2, &y) {
		fmt.Println("yup err2 == e2:", y.msg)
	}

	fmt.Println("nothing is alive")
}
