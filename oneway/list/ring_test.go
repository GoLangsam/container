// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

import "testing"

func checkRingLen(t *testing.T, l *Element, len int) bool {
	if n := l.Len(); n != len {
		t.Errorf("l.Len() = %d, want %d", n, len)
		return false
	}
	return true
}

func checkRingPointers(t *testing.T, l *Element, es []*Element) {
	root := l

	es = append(es, l)
	if !checkRingLen(t, l, len(es)) {
		return
	}

	// zero length lists must be the zero value or properly initialized (sentinel circle)
	if len(es) == 0 {
		if l.Root().next != nil && l.Root().next != root || l.Root().prev != nil && l.Root().prev != root {
			t.Errorf("l.root.next = %p, l.root.prev = %p; both should both be nil or %p", l.Root().next, l.Root().prev, root)
		}
		return
	}
	// len(es) > 0

	// check internal and external prev/next connections
	for i, e := range es {
		prev := es[len(es)-1]
		Prev := prev
		if i > 0 {
			prev = es[i-1]
			Prev = prev
		}
		if p := e.prev; p != prev {
			t.Errorf("elt[%d](%p).prev = %p, want %p", i, e, p, prev)
		}
		if p := e.Prev(); p != Prev {
			t.Errorf("elt[%d](%p).Prev() = %p, want %p", i, e, p, Prev)
		}

		next := es[0]
		Next := next
		if i < len(es)-1 {
			next = es[i+1]
			Next = next
		}
		if n := e.next; n != next {
			t.Errorf("elt[%d](%p).next = %p, want %p", i, e, n, next)
		}
		if n := e.Next(); n != Next {
			t.Errorf("elt[%d](%p).Next() = %p, want %p", i, e, n, Next)
		}
	}
}

func TestRingX(t *testing.T) {
	l := NewRing("test")
	checkRingPointers(t, l, []*Element{})

	// Single element list
	e := l.PushFront("a")
	checkRingPointers(t, l, []*Element{e})
	l.MoveToFront(e)
	checkRingPointers(t, l, []*Element{e})
	l.MoveToBack(e)
	checkRingPointers(t, l, []*Element{e})
	e.Remove()
	checkRingPointers(t, l, []*Element{})

	// Bigger list
	e2 := l.PushFront(2)
	e1 := l.PushFront(1)
	e3 := l.PushBack(3)
	e4 := l.PushBack("banana")
	checkRingPointers(t, l, []*Element{e1, e2, e3, e4})

	e2.Remove()
	checkRingPointers(t, l, []*Element{e1, e3, e4})

	l.MoveToFront(e3) // move from middle
	checkRingPointers(t, l, []*Element{e3, e1, e4})

	l.MoveToFront(e1)
	l.MoveToBack(e3) // move from middle
	checkRingPointers(t, l, []*Element{e1, e4, e3})

	l.MoveToFront(e3) // move from back
	checkRingPointers(t, l, []*Element{e3, e1, e4})
	l.MoveToFront(e3) // should be no-op
	checkRingPointers(t, l, []*Element{e3, e1, e4})

	l.MoveToBack(e3) // move from front
	checkRingPointers(t, l, []*Element{e1, e4, e3})
	l.MoveToBack(e3) // should be no-op
	checkRingPointers(t, l, []*Element{e1, e4, e3})

	e2 = l.InsertBefore(2, e1) // insert before front
	checkRingPointers(t, l, []*Element{e2, e1, e4, e3})
	e2.Remove()
	e2 = l.InsertBefore(2, e4) // insert before middle
	checkRingPointers(t, l, []*Element{e1, e2, e4, e3})
	e2.Remove()
	e2 = l.InsertBefore(2, e3) // insert before back
	checkRingPointers(t, l, []*Element{e1, e4, e2, e3})
	e2.Remove()

	e2 = l.InsertAfter(2, e1) // insert after front
	checkRingPointers(t, l, []*Element{e1, e2, e4, e3})
	e2.Remove()
	e2 = l.InsertAfter(2, e4) // insert after middle
	checkRingPointers(t, l, []*Element{e1, e4, e2, e3})
	e2.Remove()
	e2 = l.InsertAfter(2, e3) // insert after back
	checkRingPointers(t, l, []*Element{e1, e4, e3, e2})
	e2.Remove()

	// Check standard iteration.
	sum := 0
	for e := l.Front(); e != l; e = e.Next() {
		if i, ok := e.Value.(int); ok {
			sum += i
		}
	}
	if sum != 4 {
		t.Errorf("sum over l = %d, want 4", sum)
	}

	// Clear all elements by iterating
	var next *Element
	for e := l.Front(); e != l; e = next {
		next = e.Next()
		e.Remove()
	}
	checkRingPointers(t, l, []*Element{})
}

func checkRing(t *testing.T, l *Element, es []interface{}) {
	length := len(es)
	if es == nil {
		length = -1
	}
	if !checkRingLen(t, l, length) {
		return
	}

	i := 0
	for e := l.Front(); e != l; e = e.Next() {
		le := e.Value.(int)
		if le != es[i] {
			t.Errorf("elt[%d].Value = %v, want %v", i, le, es[i])
		}
		i++
	}
}

/*
func TestRingExtending(t *testing.T) {
	l1 := New()
	l2 := New()

	l1.PushBack(1)
	l1.PushBack(2)
	l1.PushBack(3)

	l2.PushBack(4)
	l2.PushBack(5)

	l3 := NewRing("test")
	l3.PushBackList(l1)
	checkRing(t, l3, []interface{}{1, 2, 3})
	l3.PushBackList(l2)
	checkRing(t, l3, []interface{}{1, 2, 3, 4, 5})

	l3 = NewRing("test")
	l3.PushFrontList(l2)
	checkRing(t, l3, []interface{}{4, 5})
	l3.PushFrontList(l1)
	checkRing(t, l3, []interface{}{1, 2, 3, 4, 5})

	checkRing(t, l1, []interface{}{1, 2, 3})
	checkRing(t, l2, []interface{}{4, 5})

	l3 = NewRing("test")
	l3.PushBackList(l1)
	checkRing(t, l3, []interface{}{1, 2, 3})
	l3.PushBackList(l3)
	checkRing(t, l3, []interface{}{1, 2, 3, 1, 2, 3})

	l3 = NewRing("test")
	l3.PushFrontList(l1)
	checkRing(t, l3, []interface{}{1, 2, 3})
	l3.PushFrontList(l3)
	checkRing(t, l3, []interface{}{1, 2, 3, 1, 2, 3})

	l3 = NewRing("test")
	l1.PushBackList(l3)
	checkRing(t, l1, []interface{}{1, 2, 3})
	l1.PushFrontList(l3)
	checkRing(t, l1, []interface{}{1, 2, 3})
}
*/

func TestRingRemove(t *testing.T) {
	l := NewRing("test")
	e1 := l.PushBack(1)
	e2 := l.PushBack(2)
	checkRingPointers(t, l, []*Element{e1, e2})
	e := l.Next()
	e.Remove()
	checkRingPointers(t, l, []*Element{e2})
	e.Remove()
	checkRingPointers(t, l, []*Element{e2})
}

func TestRingIssue4103(t *testing.T) {
	l1 := NewRing("test")
	l1.PushBack(1)
	l1.PushBack(2)

	l2 := NewRing("test")
	l2.PushBack(3)
	l2.PushBack(4)

	e := l1.Front()
	/*
		l2.Remove(e) // l2 should not change because e is not an element of l2
		if n := l2.Len(); n != 2 {
			t.Errorf("l2.Len() = %d, want 2", n)
		}
	*/

	e.InsertBefore(8, e)
	if n := l1.Len(); n != 4 {
		t.Errorf("l1.Len() = %d, want 4", n)
	}
}

func TestRingIssue6349(t *testing.T) {
	l := NewRing("test")
	l.PushBack(1)
	l.PushBack(2)

	e := l.Front()
	e.Remove()
	if e.Value != 1 {
		t.Errorf("e.value = %d, want 1", e.Value)
	}
	if e.Next() != nil {
		t.Errorf("e.Next() != nil")
	}
	if e.Prev() != nil {
		t.Errorf("e.Prev() != nil")
	}
}

func TestRingMove(t *testing.T) {
	l := NewRing("test")
	e1 := l.PushBack(1)
	e2 := l.PushBack(2)
	e3 := l.PushBack(3)
	e4 := l.PushBack(4)

	l.MoveAfter(e3, e3)
	checkRingPointers(t, l, []*Element{e1, e2, e3, e4})
	l.MoveBefore(e2, e2)
	checkRingPointers(t, l, []*Element{e1, e2, e3, e4})

	l.MoveAfter(e3, e2)
	checkRingPointers(t, l, []*Element{e1, e2, e3, e4})
	l.MoveBefore(e2, e3)
	checkRingPointers(t, l, []*Element{e1, e2, e3, e4})

	l.MoveBefore(e2, e4)
	checkRingPointers(t, l, []*Element{e1, e3, e2, e4})
	e2, e3 = e3, e2

	e1.MoveBefore(e4, e1)
	checkRingPointers(t, l, []*Element{e4, e1, e2, e3})
	e1, e2, e3, e4 = e4, e1, e2, e3

	e1.MoveAfter(e4, e1)
	checkRingPointers(t, l, []*Element{e1, e4, e2, e3})
	e2, e3, e4 = e4, e2, e3

	e1.MoveAfter(e2, e3)
	checkRingPointers(t, l, []*Element{e1, e3, e2, e4})
	e2, e3 = e3, e2
}
