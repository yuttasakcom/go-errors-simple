package main

import (
	"errors"
	"fmt"

	"github.com/yuttasakcom/go-errors-simple/validation"
)

func main() {
	if err := validation.PasswordValidation("#"); err != nil {
		fmt.Printf("%v\n", err)
		fmt.Println("------------")
		if errors.Is(err, validation.ErrInvalidLength) {
			fmt.Println("custom : Invalid Length")
			errInvalidLength := &validation.ErrInvalidLengthType{}
			if errors.As(err, errInvalidLength) {
				fmt.Println("\t--->", errInvalidLength.ActualLength)
			}

		}
		if errors.Is(err, validation.ErrMissingSmallLetter) {
			fmt.Println("custom : MissingSmallLetter")
		}
		if errors.Is(err, validation.ErrMissingCapitalLetter) {
			fmt.Println("custom : MissingCapitalLetter")
		}
		if errors.Is(err, validation.ErrMissingDigit) {
			fmt.Println("custom : MissingDigit")
			errMissingDigit := &validation.ErrMissingDigitType{}
			if errors.As(err, errMissingDigit) {
				fmt.Println("\t--->", errMissingDigit.Desc)
			}
		}
	}
	fmt.Println("Done")
}
