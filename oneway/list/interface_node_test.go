// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list_test

import (
	"github.com/GoLangsam/container/oneway/list"
)

// node abstracts the distinction between List and Element
// - the focus is on common behaviour.
type node interface {
	Beam
	Dust

	Init() *list.List
	List() *list.List

	AtomValues() list.Values
	Elements() []*list.Element
}

// Beam abstracts the 'lengthy' behaviour common to *list.Element & *list.List
type Beam interface {
	CanGrow // Push Insert
	CanMove // Move MoveTo
	CanIter // Front Next
	CanReti // Back  Prev
	Len() int
	Root() *list.Element
}

// CanGrow abstracts the growing behaviour common to *list.Element & *list.List
type CanGrow interface {
	ValuesPushFront(values ...interface{})
	ValuesPushBack(values ...interface{})

	PushBack(v interface{}) *list.Element
	PushFront(v interface{}) *list.Element

	InsertAfter(v interface{}, mark *list.Element) *list.Element
	InsertBefore(v interface{}, mark *list.Element) *list.Element
}

// CanMove abstracts the moveing behaviour common to *list.Element & *list.List
type CanMove interface {
	MoveToBack(e *list.Element)
	MoveToFront(e *list.Element)

	MoveAfter(e, mark *list.Element)
	MoveBefore(e, mark *list.Element)
}

// CanIter allows to iterate forward by starting with Front() and, if non-nil, repeating Next() until Next() returns nil
type CanIter interface {
	Front() *list.Element
	Next() *list.Element
	ForEachNext(f func(*list.Element))
}

// CanReti allows to iterate backward by starting with Back() and, if non-nil, repeating Prev() until Prev() returns nil
//  Note: Reti is Iter spelled backwards.
type CanReti interface {
	Back() *list.Element
	Prev() *list.Element
	ForEachPrev(f func(*list.Element))
}

// Dust abstracts the 'pointy' behaviour common to *list.Element & *list.List
type Dust interface {
	CVs() *list.ComposedValue

	IsAtom() bool
	IsComposed() bool

	PrintAtomValues(args ...interface{})
	PrintValue(args ...interface{})
	//	Values() list.Values
}

// Coll combines all methods unique to any List, and not shared with Element
type Coll interface {
	//	Clear() *list.List

	IsEmpty() bool // TODO: e.IsEmpty() iff Value == nil?

	Print(args ...interface{}) // TODO: samesame for e, especially for Ring

	PushBackList(other *list.List) // TODO: samesame for e, especially for Ring
	PushFrontList(other *list.List)

	Remove(*list.Element) interface{}

	Values() list.Values // TODO: samesame for e, especially for Ring
}

// Atom combines all methods unique to any Element, and not shared with List
type Atom interface {
	Remove() interface{}

	IsNode() bool
	IsRoot() bool

	MoveToPrevOf(*list.Element) *list.Element
	//	MoveToNextOf(*list.Element) *list.Element
}

/* symmetric
New(vals...) node
Equals(x node) bool
With(x node) *ComposedValue
*/

func Example_interface() {

	var _ node = list.New()
	var _ node = list.New().Root()

	var _ Coll = list.New()
	var _ Atom = list.New().Root()
}
