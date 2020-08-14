package main

import (
	"fmt"
	"github.com/lolbinarycat/taskstack/ll"
)

type Task struct {
	Body string
}

type Stack struct {
	Top *LL[Task]
}

// LL is a simple linked list
type LL[type T] struct {
	C T      // contents
	N *LL[T] // next
}

func (s *Stack) Push(t Task) {
	ll.Push(&s.Top,t)
}

// Peek returns a slice of the top n items
func (s *Stack) Peek(n int) []Task {
	ll.Peek(&s.Top,)
}

func (s *Stack) Remove(index int) (ok bool) {

}

