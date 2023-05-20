package main

import "fmt"

func SquaresOfSumAndDiff(a int64, b int64) (s int64, d int64) {
	x, y := a+b, a-b
	s = x * x
	d = y * y
	return // <=> return s, d
}

func main() {
	fmt.Println(SquaresOfSumAndDiff(10, 45))
	x, y := func() (int, int) {
		println("This function has no parameters.")
		return 3, 4
	}()
	func() {
		// The following line prints: x*x + y*y = 25
		println("x*x + y*y =", x*x+y*y)
	}()
}
