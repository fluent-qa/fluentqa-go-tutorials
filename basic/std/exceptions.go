package std

import (
	"fmt"
)

func foo() {
	panic("my exception")
}

func ExceptionDemo() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("caught exception: %s", r)
		}
	}()

	foo()
}
