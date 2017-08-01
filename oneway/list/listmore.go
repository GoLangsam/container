// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
listmore.go extends the (stolen and extended) list.go
with stuff, which is considered useful and helpfull, such as:

	- NewList( v, vals... )	returns a list, the Root() of which carries v as Value

	- l.Clear()		*List

	- l.Equals( *List )	bool
	- e.Equals( *Element )	bool

	- e.List()		*List

	- l.Root()		*Element
	- e.Root()		*Element

	- l.Next()		*Element
	- l.Prev()		*Element

	- e.Front()		*Element
	- e.Back()		*Element

	- e.IsRoot()		bool
	- e.IsNode()		bool

	- l.IsEmpty()		bool

*/
package list

// ===========================================================================

// NewList( v, vals ) returns a list, the Root() of which carries v, (and is Away to nil)
func NewList(v interface{}, vals ...interface{}) *List {
	var list = New()
	list.root.Value = v
	for _, val := range vals {
		list.PushBack(val)
	}
	return list
}

// ===========================================================================
// func (l *List) ...

// Clear disconnects all list elements from the list l,
func (l *List) Clear() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.root.list = l
	l.len = 0
	return l

}

// ===========================================================================
// Binary => bool

// Equals reports whether the lists l and t have the same element-values.
func (l *List) Equals(t *List) bool {

	if l == t {
		return true
	}
	le := l.root.next
	te := t.root.next
	for {
		switch {
			case le == &l.root && te == &t.root:	return true
			case le == &l.root || te == &t.root:	return false
			case !le.Equals(te):			return false
		}
		le = le.next
		te = te.next
	}
}

// Equals reports whether the elements e and i have the same values.
func (e *Element) Equals(i *Element) bool {
	return e.Value == i.Value
}

// ===========================================================================

// List returns the list this element belongs to
func (e *Element) List() *List {
	return e.list
}

// ===========================================================================
// Move => *Element

// Note: A rather "compact" Element would return
// e.Root() == e (or is it not?)
// e.Len() == 0
// e.Front() == e
// e.Back() == e
// and this would be a very very boring world ;-)

// Root returns the root element of list l
func (l *List) Root() *Element {
	return &l.root
}

// Root returns the Root of this elements list
func (e *Element) Root() *Element {
	if &e.list == nil {	return nil	}
	return e.list.Root()
}

// Prev returns the root element of list l
func (l *List) Prev() *Element {
	return l.root.prev
}

// Next returns the root element of list l
func (l *List) Next() *Element {
	return l.root.next
}

// Front returns the Front of this elements list
func (e *Element) Front() *Element {
	if &e.list == nil {	return nil	}
	return e.list.Front()
}

// Back returns the Back of this elements list
func (e *Element) Back() *Element {
	if &e.list == nil {	return nil	}
	return e.list.Back()
}

// => int

// Len returns the number of elements in the list of e,
// or 0 (zero), if e.IsRoot or -1 if e.list == nil
// The complexity is O(1).
func (e *Element) Len() int {
	if &e.list == nil	{	return -1	}
	if e.IsRoot()		{	return 0	}
	return e.list.Len()
}

// => bool

// IsRoot reports whether the element e is Root() of it's list
func (e *Element) IsRoot() bool {
	return (e == &e.list.root)
}

// IsNode: an element which is not root can be seen as a node
func (e *Element) IsNode() bool {
	return (e != &e.list.root)
}

// IsEmpty reports whether the list l is empty.
// Note: Does not evaluate Len(), as this could be scrambled temporarily
func (l *List) IsEmpty() bool {
	return l.root.next == &l.root
}
