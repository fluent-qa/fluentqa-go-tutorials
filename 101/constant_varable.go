package main

import "fmt"

// Declare two individual constants. Yes,
// non-ASCII letters can be used in identifiers.
const π = 3.1416
const Pi = π // <=> const Pi = 3.1416

// Declare multiple constants in a group.
const (
	T          = iota
	No         = !Yes
	Yes        = true
	MaxDegrees = 360
	Unit       = "radian"
	q          = iota * 2
)

const (
	Failed    = iota - 1 // == -1
	Unknown              // == 0
	Succeeded            // == 1
)

func main() {
	// Declare multiple constants in one line.
	const TwoPi, HalfPi, Unit2 = π * 2, π * 0.5, "degree"
	fmt.Println(TwoPi, HalfPi, Unit2, A, B, X, Y)
	fmt.Println(T, q)
	fmt.Println(Unknown, Failed, Succeeded)
	var lang, website string = "Go", "https://golang.org"
	var compiled, dynamic bool = true, false
	var announceYear int = 2009
	fmt.Println(lang, website, compiled, dynamic, announceYear)
}

const X = float32(3.14)

const (
	A, B = int64(-3), int64(5)
	Y    = float32(2.718)
)
