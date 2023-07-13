package main

import (
	"fmt"

	"github.com/yuttasakcom/go-errors-simple/validation"
)

type Unwrappable struct {
	msg     string
	wrapped error
}

func (u Unwrappable) Error() string {
	return u.msg
}

func (u Unwrappable) Unwrap() error {
	return u.wrapped
}

func main() {
	if err := validation.PasswordValidation("#"); err != nil {
		fmt.Printf("%+v", *err.(*validation.PasswordError))
	}
}
