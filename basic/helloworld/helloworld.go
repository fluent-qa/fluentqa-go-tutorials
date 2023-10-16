package helloworld

import (
	"fmt"
)

// Greet ...
func Greet(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
