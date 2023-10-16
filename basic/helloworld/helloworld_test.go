package helloworld

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGreet(t *testing.T) {
	result := Greet("world")
	assert.Equal(t, result, "hello world")

	var result1 = Greet("world")
	fmt.Println(result1)
}
