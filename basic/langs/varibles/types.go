package varibles

import (
	"fmt"
	"runtime"
)

var (
	speed int
	heat  float64

	off   bool
	brand string
)
var packageLevelVar, Name string

func VariableTypesMore() {
	fmt.Println(42, 8500, 344433, -2323)
	fmt.Println(3.14, 6.28, -42.)
	fmt.Println(true, false)
	fmt.Println("Hi! I'm String!")
	//HEX
	fmt.Println(0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9)
	fmt.Println(0xa, 0xb, 0xc, 0xd, 0xe, 0xf)
	fmt.Println(0x11) // 17
	fmt.Println(0x19) // 25
	fmt.Println(0x32) // 50
	fmt.Println(0x64) // 100
}

func DeclarationIt() {
	var speed int

	fmt.Println(speed)

	var nFiles int
	var counter int
	var nCPU int

	fmt.Println(
		nFiles,
		counter,
		nCPU,
	)

	var heat float64
	var ratio float64
	var degree float64

	fmt.Println(
		heat,
		ratio,
		degree,
	)

	var off bool
	var valid bool
	var closed bool

	fmt.Println(
		off,
		valid,
		closed,
	)

	var msg string
	var name string
	var text string

	fmt.Println(
		msg,
		name,
		text,
	)

	_ = speed
	name, lastname := "Nikola", "Tesla"
	fmt.Println(name, lastname)

	birth, death := 1856, 1943
	fmt.Println(birth, death)

	on, off := true, false
	fmt.Println(on, off)
}

func Statement() {
	fmt.Println("Hello!")
	fmt.Println("Hello")
	fmt.Println("World!")
	//	fmt.Println("Hello"); fmt.Println("World!")

	// Statements change the execution flow
	// Especially the control flow statements like `if`
	if 5 > 1 { //controller  flow
		fmt.Println("bigger")
	}
	safe := true
	fmt.Println(safe)
	//operator
	fmt.Println("Hello!" + "!")
	fmt.Println("Hello!" + "!" + "!" + "?")
	fmt.Println(runtime.NumCPU() + 1)

}

func AssignmentAndSwap() {
	var (
		speed     = 100
		prevSpeed = 50
	)

	speed, prevSpeed = prevSpeed, speed

	fmt.Println(speed, prevSpeed)
}

func VariableDemo() {

	// explicit
	var foo string = "foo"

	// type inferred
	var bar = "foo"

	// shorthand
	baz := "bar"

	// constant
	const qux = "qux"

	_ = foo
	_ = bar
	_ = baz
	_ = qux

}
