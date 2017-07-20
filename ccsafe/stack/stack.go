// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package stack provides a concurrency-safe stack for interface{}
guarded by a sync.Mutex

*/
package stack

import (
	"sync"
)

// Stack implements a concurrency-safe stack for anything (interface{})
// guarded by a sync.Mutex
type Stack struct {
	stack []interface{}
	sync.Mutex
}

// New returns a new empty stack with given initial capacity
func New(cap int) *Stack {
	return new(Stack).Init(cap)
}

// Init returns an empty stack with given initial capacity
func (s *Stack) Init(cap int) *Stack {
	s.stack = make([]interface{}, 0, cap)
	return s
}

// Push sth onto the current stack
func (s *Stack) Push(l interface{}) {
	s.Lock()
	defer s.Unlock()

	s.stack = append(s.stack, l)
}

// Pop sth off the current stack
func (s *Stack) Pop() interface{} {
	s.Lock()
	defer s.Unlock()

	var p = s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return p
}

// Drop pops sth silently off the current stack
func (s *Stack) Drop() {
	s.Lock()
	defer s.Unlock()

	//	var p = s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return // p
}

// Top returns the top of the current stack
func (s *Stack) Top() interface{} {
	s.Lock()
	defer s.Unlock()

	return s.stack[len(s.stack)-1]
}

// Get returns a copy of the current stack
func (s *Stack) Get() []interface{} {
	s.Lock()
	defer s.Unlock()

	var stack = make([]interface{}, len(s.stack))
	copy(stack, s.stack)
	return stack
}

// Len returns the length of the current stack
func (s *Stack) Len() int {
	s.Lock()
	defer s.Unlock()

	return len(s.stack)
}
