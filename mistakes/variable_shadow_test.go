package main

import (
	"fmt"
	"testing"
)

func Test_listing1(t *testing.T) {
	fmt.Println("this is main entry")
	createClientWithTracing()
	createDefaultClient()
	fmt.Println(tracing)
}
