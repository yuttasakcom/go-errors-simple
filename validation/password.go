package validation

import (
	"fmt"
	"regexp"
	"unicode/utf8"
)

type PasswordError struct {
	msg string
}

func (p *PasswordError) Error() string {
	return p.msg
}

func PasswordValidation(pw string) error {
	pwError := &PasswordError{}
	if err := checkLength(pw); err != nil {
		pwError.msg = err.Error() + "\n"
	}

	if err := containSmallLetter(pw); err != nil {
		pwError.msg = pwError.msg + err.Error() + "\n"
	}

	if err := containCapitalLetter(pw); err != nil {
		pwError.msg = pwError.msg + err.Error() + "\n"
	}

	if err := containDigit(pw); err != nil {
		pwError.msg = pwError.msg + err.Error() + "\n"
	}

	if pwError.msg != "" {
		return pwError
	}
	return nil
}

func regexHelper(pw, pattern, msg string) error {
	re := regexp.MustCompile(pattern)
	result := re.FindString(pw)
	if result == "" {
		return fmt.Errorf(msg)
	}
	return nil
}

func containSmallLetter(pw string) error {
	return regexHelper(pw, `[a-z]`, "password must contain at least one small letter")
}

func containCapitalLetter(pw string) error {
	return regexHelper(pw, `[A-Z]`, "password must contain at least one capital letter")
}

func containDigit(pw string) error {
	return regexHelper(pw, `[0-9]`, "password must contain at least one digit")
}

func checkLength(pw string) error {
	pwLen := utf8.RuneCountInString(pw)
	if pwLen < 7 || pwLen > 16 {
		return fmt.Errorf("password length must be between 7 and 16")
	}
	return nil
}
