// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
listmany.go extends list.go with:

	- l.Elements()		[]*Element
	- e.Elements()		[]*Element

	- l.Values()		Values
	- e.Values()		Values

	- l.ValuesPushBack( v... )
	- l.ValuesPushFront( v... )
*/

package list

// ===========================================================================

// Values aliases a slice of Element.Value
type Values []interface{}

// ===========================================================================

// Elements returns the elements of list l as a slice
func (l *List) Elements() []*Element {
	var data = make([]*Element, 0, l.Len())
	for e := l.Front(); e != nil; e = e.Next() {
		data = append(data, e)
	}
	return data
}

// Elements returns the Elements as a slice
// of the list of the element (syntactic sugar)
func (e *Element) Elements() []*Element {
	return e.List().Elements()
}

// Values returns all Element.Values as Values-slice
func (l *List) Values() Values {
	var data = make([]interface{}, 0, l.Len())
	for e := l.Front(); e != nil; e = e.Next() {
		data = append(data, e.Value)
	}
	return data
}

// Values returns all Element.Values of e.List()
// as Values-slice (syntactic sugar)
func (e *Element) Values() Values {
	return e.List().Values()
}

// ===========================================================================
// func (l *List) ...

// ValuesPushBack appends a slice of Values
func (l *List) ValuesPushBack(values ...interface{}) {
	valuesDo(l.PushBack, values...)
}

// ValuesPushFront prepends a slice of Values
func (l *List) ValuesPushFront(values ...interface{}) {
	valuesDo(l.PushFront, values...)
}

// ValuesPushBack appends a slice of Values
func (e *Element) ValuesPushBack(values ...interface{}) {
	valuesDo(e.PushBack, values...)
}

// ValuesPushFront prepends a slice of Values
func (e *Element) ValuesPushFront(values ...interface{}) {
	valuesDo(e.PushFront, values...)
}

// valuesDo executes the given function on a slice of Values
func valuesDo(do func(v interface{}) *Element, values ...interface{}) {
	for i := range values {
		do(values[i])
	}
}
