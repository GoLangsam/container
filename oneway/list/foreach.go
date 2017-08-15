// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
foreach.go extends list.go with:

	- l.ForEachNext( f func(*Element) )
	- e.ForEachNext( f func(*Element) )

	- l.ForEachPrev( f func(*Element) )
	- e.ForEachPrev( f func(*Element) )

Important: f must not mutate l !!!

Thus, "ForEach" is not really safe to expose to clients.

Note: using e.g. some Walker from "dlx/list/walk" or "dlx/trix/walk"
allows more natural control flow with continue/break/return etc.
*/

package list

// ForEachNext applies function f to each element of the list l in natural order.
func (l *List) ForEachNext(f func(*Element)) {
	for e := l.root.next; e != &l.root; e = e.next {
		f(e)
	}
}

// ForEachPrev applies function f to each element of the list l in reverse order.
func (l *List) ForEachPrev(f func(*Element)) {
	for e := l.root.prev; e != &l.root; e = e.prev {
		f(e)
	}
}

// ForEachNext applies function f to each other element of e's list in natural order.
func (e *Element) ForEachNext(f func(*Element)) {
	for i := e.next; i != e; i = i.next {
		if i != &e.list.root {
			f(i)
		}
	}
}

// ForEachPrev applies function f to each other element of e's list in reverse order.
func (e *Element) ForEachPrev(f func(*Element)) {
	for i := e.prev; i != e; i = i.prev {
		if i != &e.list.root {
			f(i)
		}
	}
}
