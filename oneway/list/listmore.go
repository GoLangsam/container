// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
listmore.go extends list.go with:

	- NewList( vals... )	*List
	- l.New( vals... )		*List
	- e.New( vals... )		*Element // == Root()

	- l.Equals( *List )		bool
	- e.Equals( *Element )	bool

	- e.IsRoot()			bool
	- e.IsNode()			bool

	- l.IsEmpty()			bool
*/

package list

// ===========================================================================

// NewList returns a new list
// the Root() of which carries vals[0] (if any)
// and the list has elements vals[1:] (if any)
func NewList(vals ...interface{}) *List {
	var list = New()
	if len(vals) > 0 {
		list.root.Value = vals[0]
		for i := range vals[1:] {
			list.PushBack(vals[1+i])
		}
	}

	return list
}

// New returns a new list
// the Root() of which carries vals[0] (if any)
// and the list has elements vals[1:] (if any)
func (l *List) New(vals ...interface{}) *List {
	return NewList(vals...)
}

// New returns the root element of a new list
// the Root() of which carries vals[0] (if any)
// and the list has elements vals[1:] (if any)
func (e *Element) New(vals ...interface{}) *Element {
	return NewList(vals...).Root()
}

// ===========================================================================
// func (l *List) ...

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
		case le == &l.root && te == &t.root:
			return true
		case le == &l.root || te == &t.root:
			return false
		case !le.Equals(te):
			return false
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

// => bool

// IsRoot reports whether the element e is Root() of it's list
// unless it's nil, belongs to nil list or list's root is nil
func (e *Element) IsRoot() bool {
	if &e == nil || &e.list == nil || &e.list.root == nil {
		return false
	}
	return (e == &e.list.root)
}

// IsNode - any element which !IsRoot can be seen as a node
// unless it's nil, belongs to nil list or list's root is nil
func (e *Element) IsNode() bool {
	if &e == nil || &e.list == nil || &e.list.root == nil {
		return false
	}
	return (e != &e.list.root)
}

// IsEmpty reports whether the list l is empty.
// Note: Does not evaluate Len(), as this could be scrambled temporarily
func (l *List) IsEmpty() bool {
	l.lazyInit()
	return l.root.next == &l.root
}
