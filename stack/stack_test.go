package stack_test

import (
	"testing"

	"github.com/bumpsoo/golib/stack"
)

func TestStack(t *testing.T) {
	s := stack.Stack[int]{}
	for i := 5; i >= 0; i-- {
		s.Push(i)
	}
	for i := 0; i <= 5; i++ {
		v, ok := s.Pop()
		if !ok || v != i {
			t.Fatal()
		}
	}
	_, ok := s.Pop()
	if ok {
		t.Fatal()
	}
}
