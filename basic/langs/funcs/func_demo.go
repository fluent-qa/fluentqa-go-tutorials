package funcs

import "fmt"

func add(a int, b int) int {
	return a + b
}

func Walkthrough() {
	result := add(2, 3)
	fmt.Println(result)
}
