package main

import (
	"fmt"
	"errors"
)
type argError struct {
	arg int
	msg string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s",e.arg, e.msg)
}

func f(arg int) (int,error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with 42"}
	}else {
		return arg,nil
	}
}
func main() {
	_, e := f(42)
	var ae *argError

	if errors.As(e, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.msg)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
