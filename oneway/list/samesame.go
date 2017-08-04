// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
samesame.go extends the (stolen and extended) list.go
with stuff, which is considered useful and helpfull, such as:

Blurr the difference between list and element so element can represent both
	- e.Init()		*Element

	- e.Front()		*Element
	- e.Back()		*Element

	- l.Next()		*Element
	- l.Prev()		*Element

Going home
	- l.List()		*List
	- e.List()		*List

	- l.Root()		*Element
	- e.Root()		*Element

*/
// Note: A rather "compact" Element would return
// e.Front() == e
// e.Back() == e
// e.Root() == e (or is it not?)
// e.Len() == 0
// and this would be a very very boring world ;-)
package list

// ===========================================================================
// Blurr the difference between list and element so either can represent both

// Init initializes or clears list of e.
// A fresh and empty list is returned iff e or e.list is nil.
func (e *Element) Init() (list *List) {
	if &e == nil || &e.list == nil {
		list = New()
	} else {
		list = e.list.Init()
	}
	return list
}

// Front returns the Front of this elements list
func (e *Element) Front() *Element {
	if &e == nil || &e.list == nil {
		return nil
	}
	return e.list.Front()
}

// Back returns the Back of this elements list
func (e *Element) Back() *Element {
	if &e == nil || &e.list == nil {
		return nil
	}
	return e.list.Back()
}

// Next returns the Front of this list l
func (l *List) Next() *Element {
	return l.root.next
}

// Prev returns the Back of this list l
func (l *List) Prev() *Element {
	return l.root.prev
}

// Len returns the number of elements in the list of e,
// or 0 (zero), if e.IsRoot or -1 if e.list == nil
// The complexity is O(1).
func (e *Element) Len() int {
	if &e.list == nil {
		return -1
	}
	if e == &e.list.root { // IsRoot()
		return 0
	}
	return e.list.Len()
}

// ===========================================================================

// List returns this list
func (l *List) List() *List {
	return l
}

// List returns this list this element belongs to
func (e *Element) List() *List {
	return e.list
}

// Root returns the root element of list l
func (l *List) Root() *Element {
	return &l.root
}

// Root returns the Root of this elements list
func (e *Element) Root() *Element {
	if &e.list == nil {
		return nil
	}
	return e.list.Root()
}