// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
listmore.go extends the (stolen and extended) list.go
with stuff, which is considered useful and helpfull, such as:

Create (also as method!)
	- NewList( v, vals... )	returns a list, the Root() of which carries v as Value

	- l.Equals( *List )	bool
	- e.Equals( *Element )	bool

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
	for i := range vals {
		list.PushBack(vals[i])
	}
	return list
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
