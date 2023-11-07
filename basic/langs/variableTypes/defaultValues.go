package variableTypes

import "fmt"

func greet(name *string) string {
	n := "stranger"
	// use pointers and check for nil to know if explicitly left blank
	if name != nil {
		n = *name
	}

	return fmt.Sprintf("hello %s", n)
}

func DefaultValues() {
	message := greet(nil)
	fmt.Println(message)

	name := "bob"
	message = greet(&name)
	fmt.Println(message)
}
