// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:pattern "github.com/GoLangsam/container/oneway/stack/stack.go"

/*
Package stack provides a normal (=non-concurrency-safe) stack
for anything (interface{})

Note: the very crisp implementation was found in cmd/go/pkg.go importStack
*/
package stack

// Stack implements a normal (=non-concurrency-safe) stack
// for anything (interface{})
// based on a slice, and never shrinking
type Stack []interface{}

// New returns a new empty stack with given initial capacity
func New(cap int) *Stack {
	return new(Stack).Init(cap)
}

// Init returns an empty stack with given initial capacity
func (s *Stack) Init(cap int) *Stack {
	return &Stack{make([]interface{}, 0, cap)}

}

// Push sth onto the current stack
func (s *Stack) Push(l interface{}) {
	//	s.Lock()
	//	defer s.Unlock()

	*s = append(*s, l)
}

// Pop sth off the current stack
func (s *Stack) Pop() interface{} {
	//	s.Lock()
	//	defer s.Unlock()

	p := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return p
}

// Drop pops sth silently off the current stack
func (s *Stack) Drop() {
	//	s.Lock()
	//	defer s.Unlock()

	*s = (*s)[0 : len(*s)-1]
}

// Top returns the top of the current stack
func (s *Stack) Top() interface{} {
	//	s.Lock()
	//	defer s.Unlock()

	return (*s)[len(*s)-1]
}

// Get returns a copy of the current stack
func (s *Stack) Get() []interface{} {
	//	s.Lock()
	//	defer s.Unlock()

	//	return append([]interface{}{}, *s...)
	var stack = make([]interface{}, len(*s))
	copy(stack, *s)
	return stack
}

// Len returns the length of the current stack
func (s *Stack) Len() int {
	//	s.Lock()
	//	defer s.Unlock()

	return len(*s)
}
