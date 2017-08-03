// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
compose.go extends the (stolen and extended) list.go
with stuff, which is considered useful and helpfull, such as:

	- l.With( *List )		*ComposedValue
	- e.With( *Element )		*ComposedValue

	- l.CVs()			*ComposedValue
	- e.CVs()			*ComposedValue

	- l.IsComposed()		bool
	- e.IsComposed()		bool

	- l.IsAtom()			bool
	- e.IsAtom()			bool

*/
package list

// ComposedValue aliases the Element.Value of IsComposed() Elements
type ComposedValue []*Element

// With returns a new slice of *Element to be used as ComposedValue for new Elements
func (l *List) With(x *List) *ComposedValue {
	return l.root.With(&x.root)
}

// With returns a new slice of *Element to be used as ComposedValue for new Elements
func (e *Element) With(x *Element) *ComposedValue {
	var r ComposedValue = make([]*Element, 0, e.vLen()+x.vLen())
	r = append(r, *e.CVs()...)
	r = append(r, *x.CVs()...)
	return &r
}

func (l *List) vLen() int {
	return l.root.vLen()
}
func (e *Element) vLen() int {
	switch ev := e.Value.(type) {
		case *ComposedValue:	return len(*ev)
		default:		return 1
	}
}

// CVs - obtain slice of *Element to be used as new Composed Value
//  Convenience for l.root.CVs()
func (l *List) CVs() *ComposedValue {
	return l.root.CVs()
}

// CVs - obtain slice of *Element to be used as new Composed Value
//  If e.IsAtom() the slice has one element which points to e
//  else a pointer(!) to the existing slice of atoms is returned
func (e *Element) CVs() *ComposedValue {
	switch ev := e.Value.(type) {
		case *ComposedValue:	return ev
		default:		return &ComposedValue{e}
	}
}

// IsComposed <=> l.root is composed
// (and thus carries a Value.(type) []*Element)
//  Convenience for l.root.IsComposed()
func (l *List) IsComposed() bool {
	return l.root.IsComposed()
}

// IsComposed <=> element is composed
// (and thus carries a Value.(type) []*Element)
func (e *Element) IsComposed() bool {
	switch e.Value.(type) {
	case *ComposedValue:	return true
	default:		return false
	}
}

// IsAtom <=> l.root is not composed
//  Convenience for l.root.IsAtom()
func (l *List) IsAtom() bool {
	return l.root.IsAtom()
}

// IsAtom <=> element is not composed
func (e *Element) IsAtom() bool {
	switch e.Value.(type) {
	case *ComposedValue:	return false
	default:		return true
	}
}

// AtomValues - obtain slice of original values
//  Convenience for l.root.AtomValues()
func (l *List) AtomValues() Values {
	return l.root.AtomValues()
}

// AtomValues - obtain slice of original values
func (e *Element) AtomValues() Values {
	switch ev := e.Value.(type) {
	case *ComposedValue:
		r := make([]interface{}, 0, len(*ev))
		for _, x := range *ev {
			r = append(r, x.AtomValues()...)
		}
		return r

	default:
		return []interface{}{e.Value}
	}
}
