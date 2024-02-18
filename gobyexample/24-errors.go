package main 

import (
	"fmt"
	"errors"
)

type argError struct {
	arg int
	reason string
}

// sending errors as last value is standard followed
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Told ya to not send 42")
	}

	return arg+3, nil
}


func (e *argError) Error() string {
	return fmt.Sprintf("%d %s", e.arg, e.reason)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg: arg, reason: "Told you to not pass 42."}
	} 
	return arg+3, nil
}

func main() {

	if f1Ret, err := f2(32); err != nil {
		fmt.Println("F1 failed.", f1Ret, err)
	} else {
		fmt.Println("F1 passed.", f1Ret, err)
		
	}

	
	if f2Ret, err1 := f2(42); err1 != nil {
		fmt.Println("F2 failed.", f2Ret, err1)
	} else {
		fmt.Println("F2 passed.", f2Ret, err1)
	}


	for _, num := range []int{49,44,43,42} {
		if r, e := f1(num); e != nil {
			fmt.Println("Failed.", num, e, r)
		} else {
			fmt.Println("Passed.", num, e, r)
		}
	}


}