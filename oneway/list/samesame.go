// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
samesame.go extends list.go with:

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
// and this would be a very very boring world ;-)

package list

// ===========================================================================
// Blurr the difference between list and element so either can represent both

// Init returns fresh and empty list.
// e is left untouched.
func (e *Element) Init() (list *List) {
	return New()
}

// Front returns the Front of this elements list
// (its first one, if any)
// or e.Next(), iff e belongs to a Ring (and has no list).
func (e *Element) Front() *Element {
	switch {
	case e == nil:
		return nil
	case e.list != nil:
		return e.list.Front()
	default:
		return e.next
	}
}

// Back returns the Back of this elements list
// (its last one, if any)
// or e.Prev(), iff e belongs to a Ring (and has no list).
func (e *Element) Back() *Element {
	switch {
	case e == nil:
		return nil
	case e.list != nil:
		return e.list.Back()
	default:
		return e.prev
	}
}

// Next returns the Front element of list l
// (its first one, if any) or nil.
func (l *List) Next() *Element {
	if &l == nil {
		return nil
	}
	return l.root.next
}

// Prev returns the Back element of list l
// (its last one, if any) or nil.
func (l *List) Prev() *Element {
	if &l == nil {
		return nil
	}
	return l.root.prev
}

// ===========================================================================

// List returns this list (which may be nil).
func (l *List) List() *List {
	return l
}

// List returns the list of e
// or nil iff e == nil or its list is nil.
func (e *Element) List() *List {
	if e == nil || e.list == nil {
		return nil
	}
	return e.list
}

// Root returns the root element of list l
// or nil iff l == nil or its root is nil.
func (l *List) Root() *Element {
	if l == nil {
		return nil
	}
	return &l.root
}

// Root returns the Root of this elements list
// or nil iff e == nil or e iff its list is nil.
func (e *Element) Root() *Element {
	switch {
	case e == nil:
		return nil
	case e.list == nil:
		return e
	default:
		return e.list.Root()
	}
}
