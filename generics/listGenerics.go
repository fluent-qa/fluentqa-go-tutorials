package generics

import "fmt"

// return -1 if not found
// generic function
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

type ListG[T any] struct {
	next *ListG[T]
	val  T
}

func CreateListG() {
	g := &ListG[any]{
		next: &ListG[any]{},
		val:  "test",
	}
	g.next.val = "test1"
	g.next.next = g
	fmt.Println(g.val, g.next.val)
}
