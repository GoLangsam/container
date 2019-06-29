// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

// Ring returns a new ring.
// Iff no values are given, a nil element is returned.
func Ring(vals ...interface{}) (ring *Element) {
	if len(vals) > 0 {
		ring = &Element{Value: vals[0]}
		ring.prev, ring.next = ring, ring
		for i := range vals[1:] {
			ring.PushBack(vals[1+i])
		}
	}

	return
}

func (e *Element) assertRing() {
	if e != nil && e.list != nil {
		panic("element does not belong to some ring but to some list!")
	}
}

// Len returns the number of elements in the list of e
// or 0 (zero), if e is root of it's (non-nil) list,
// or -1 iff e == nil.
//
// The complexity is O(1) iff e is element of a list
// and O(n) otherwise (e is element of some ring).
func (e *Element) Len() int { 
	switch{
	case e == nil:
		return -1
	case e.list != nil:
		if e == &e.list.root { // IsRoot()
			return 0
		}
		return e.list.len		

	default:
		var length int
		for at := e.next; at != e; at = at.next{
			length++
		}
		return length
	
	}
}

// insert inserts e after at, increments l.len iff l is not nil, and returns e.
func (e *Element) insert(at *Element) *Element {
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e
	e.list = at.list
	if e.list != nil {
		e.list.len++
	}
	return e
}

// insertValue is a convenience wrapper for e.insert(&Element{Value: v}).
func (e *Element) insertValue(v interface{}) *Element {
	return e.insert(&Element{Value: v})
}

// remove removes e from its list or ring, decrements l.len iff l is not nil, and returns e.
func (e *Element) remove() *Element {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	//away = nil // no touch: MoveXyz functions use 'removes' followed by 'insert'
	if e.list != nil {
		e.list.len--
	}
	e.list = nil
	return e
}

// move moves e to next to at and returns e.
func (e *Element) move(at *Element) *Element {
	if e == at {
		return e
	}
	e.prev.next = e.next
	e.next.prev = e.prev

	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e

	return e
}

// Remove removes e from l if e is an element of list l.
// It returns the element value e.Value.
// The element must not be nil.
func (e *Element) Remove() interface{} {
	switch{
	case e == nil:
		return nil
	case e.list != nil:
		e.list.remove(e)
	default:
		e.remove()
	}

	return e.Value
}

// PushFront inserts a new element e with value v at the front of ring e and returns e.
func (e *Element) PushFront(v interface{}) *Element {
	e.assertRing()
	return e.insertValue(v)
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (e *Element) PushBack(v interface{}) *Element {
	e.assertRing()
	return e.prev.insertValue(v)
}

// InsertBefore inserts a new element e with value v immediately before mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (mark *Element) InsertBefore(v interface{}) *Element {
	if mark.list != nil {
		return mark.list.InsertBefore(v, mark)
	}
	return mark.prev.insertValue(v)
}

// InsertAfter inserts a new element e with value v immediately after mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (mark *Element) InsertAfter(v interface{}) *Element {
	if mark.list != nil {
		return mark.list.InsertAfter(v, mark)
	}
	return mark.insertValue(v)
}

// MoveToFront moves element e to the front of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (mark *Element) MoveToFront(e *Element) {
	if mark.list != nil {
		mark.list.MoveToFront(e)
	}
	mark.move(e)
}

// MoveToBack moves element e to the back of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (mark *Element) MoveToBack(e *Element) {
	if mark.list != nil {
		mark.list.MoveToBack(e)
	}
	mark.prev.move(e)
}

// MoveBefore moves element e to its new position before mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (mark *Element) MoveBefore(e *Element) {
	if mark.list != nil {
		mark.list.MoveBefore(e, mark)
	}
	mark.prev.move(e)
}

// MoveAfter moves element e to its new position after mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (mark *Element) MoveAfter(e *Element) {
	if mark.list != nil {
		mark.list.MoveAfter(e, mark)
	}
	mark.move(e)
}

/*

// PushBackList inserts a copy of an other list at the back of list l.
// The lists l and other may be the same. They must not be nil.
func (l *List) PushBackList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.root.prev)
	}
}

// PushFrontList inserts a copy of an other list at the front of list l.
// The lists l and other may be the same. They must not be nil.
func (l *List) PushFrontList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.Value, &l.root)
	}
}

*/