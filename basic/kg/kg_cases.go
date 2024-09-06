package kg

import (
	"maps"
	"slices"
	"sync"
)

func Merge[M ~map[K]V, K comparable, V any](ms ...M) M {
	result := M{}
	for _, m := range ms {
		maps.Copy[map[K]V](result, m)
	}
	return result
}

func ContainsFunc[E any](s []E, f func(E) bool) bool {
	return slices.IndexFunc(s, f) >= 0
}

type Channel[T any] struct {
	lock            *sync.RWMutex
	ch              chan T
	sends, receives int
}

func NewChannel[T any](length int) *Channel[T] {
	return &Channel[T]{
		lock: &sync.RWMutex{},
		ch:   make(chan T, length),
	}
}

func (c *Channel[T]) Send(v T) {
	c.ch <- v
	c.lock.Lock()
	defer c.lock.Unlock()
	c.sends++
}

func (c *Channel[T]) Receive() T {
	v := <-c.ch
	c.lock.Lock()
	defer c.lock.Unlock()
	c.receives++
	return v
}

func (c *Channel[T]) Sends() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.sends
}

func (c *Channel[T]) Receives() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.receives
}

type Stack[E any] struct {
	data []E
}

func (s Stack[E]) Len() int {
	return len(s.data)
}

func (s *Stack[E]) Push(vals ...E) {
	s.data = append(s.data, vals...)
}

func (s *Stack[E]) Pop() (v E, ok bool) {
	if len(s.data) == 0 {
		return v, false
	}
	v = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, true
}

type FuncMap[T, U any] map[string]func(T) U

func (fm FuncMap[T, U]) Apply(name string, val T) U {
	return fm[name](val)
}
