package std

import "fmt"

type DemoObj struct {
	Key   string
	Value string
}

func (o *DemoObj) Read() (string, string) {
	return o.Key, o.Value
}

func main() {
	obj := DemoObj{
		Key:   "foo",
		Value: "bar",
	}

	// option 1: multiple variable assignment
	key, value := obj.Key, obj.Value
	fmt.Println(key, value)

	// option 2: return multiple values from a function
	key, value = obj.Read()
	fmt.Println(key, value)
}
