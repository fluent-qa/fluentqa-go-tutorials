package conversion

import "fmt"

func Convert() {
	speed := 100 // int
	force := 2.5 // float64

	// ERROR: invalid op
	// speed = speed * force

	// conversion can be a destructive operation
	// `force` loses its fractional part...

	speed = speed * int(force)

	fmt.Println(speed)
}

func CorrectIt() {
	speed := 100
	force := 2.5

	fmt.Printf("speed     : %T\n", speed)
	fmt.Printf("conversion: %T\n", float64(speed))
	fmt.Printf("expression: %T\n", float64(speed)*force)

	// TYPE MISMATCH:
	//   speed is int
	//   expression is float64
	// speed = float64(speed) * force

	// CORRECT: int * int
	speed = int(float64(speed) * force)

	fmt.Println(speed)
}

func StringInt() {
	var apple int
	var orange int32
	apple = int(orange)

	// you can't convert a numeric type to a bool:
	// isDelicious := bool(orange)

	// but you can convert an int to a string
	// this only works with int types
	orange = 65 // 65 is A
	color := string(orange)
	fmt.Println(color)

	// this doesn't work. 65.0 is a float.
	// fmt.Println(string(65.0))

	// this works: there are two byte values
	// byte is also an int
	fmt.Println(string([]byte{104, 105}))

	_ = apple
}
