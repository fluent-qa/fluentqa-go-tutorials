package lessons

import "fmt"

func PrintName() {
	// this is a line comment

	/*
	   this is a block comment
	*/
	fmt.Println("Hello Name")
	name := "bob"
	age := 21
	message := fmt.Sprintf("%s is %d years old", name, age)

	fmt.Println(message)
}
