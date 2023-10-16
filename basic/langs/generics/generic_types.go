package generics

import "fmt"

type ObjectOne struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

func (a *ObjectOne) GetName() string {
	return a.Name
}

type ObjectTwo struct {
	Name   string      `json:"name"`
	Data   interface{} `json:"data"`
	Option bool        `json:"option"`
}

func (a *ObjectTwo) GetName() string {
	return a.Name
}

// Generics
type objectType interface {
	*ObjectOne | *ObjectTwo
	GetName() string
}

type Object[T objectType] struct {
	Body T
}

func New[T objectType](object T) *Object[T] {
	return &Object[T]{
		Body: object,
	}
}

func (o *Object[T]) PrintName() string {
	return o.Body.GetName()
}

func CreateObjects() {
	o := New(&ObjectOne{
		Name: "Object-1",
		Data: 123,
	})
	fmt.Printf("%+v\n", *o.Body)

	t := New(&ObjectTwo{
		Name:   "Object-2",
		Data:   "payload",
		Option: true,
	})
	fmt.Printf("%+v\n", *t.Body)

	// Usage
	fmt.Printf("%+v\n", o.PrintName())
	fmt.Printf("%+v\n", t.PrintName())
}
