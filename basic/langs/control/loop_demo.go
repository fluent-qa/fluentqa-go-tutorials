package control

import "fmt"

func LoopDemo() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
}

func WhileDemo() {
	i := 0

	for i <= 5 {
		fmt.Println(i)

		i++
	}
}
