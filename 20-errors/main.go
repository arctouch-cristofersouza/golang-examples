package main

import (
	"errors"
	"fmt"
)

func invalidNumber(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

func main() {

	//default way to deal with error in go
	for _, i := range []int{7, 42} {
		r, err := invalidNumber(i)
		if err != nil {
			fmt.Println("invalidNumber failed:", err)
		} else {
			fmt.Println("invalidNumber worked:", r)
		}
	}

	fmt.Println("========")
	//
	for i := range []int{1, 2, 3, 4, 5} {
		if err := makeTea(i); err != nil {

			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}

		fmt.Println("Tea is ready!")
	}
}

var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}
