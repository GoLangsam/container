// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
listmore.go extends list.go with:

Create (TODO: also as method!)
	- NewList( v, vals... )	returns a list, the Root() of which carries v as Value

	- l.Equals( *List )	bool
	- e.Equals( *Element )	bool

	- e.IsRoot()		bool
	- e.IsNode()		bool

	- l.IsEmpty()		bool
*/

package list

// ===========================================================================

// NewList returns a list of vals, the Root() of which carries v
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
