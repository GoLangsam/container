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
	Init() *list.List
	List() *list.List
	Front() *list.Element
	Back() *list.Element
	Next() *list.Element
	Prev() *list.Element
	Root() *list.Element
	Len() int
	// vLen() int
	CVs() *list.ComposedValue
	IsAtom() bool
	IsComposed() bool

	ForEachNext(f func(*list.Element))
	ForEachPrev(f func(*list.Element))

	AtomValues() list.Values
	Elements() []*list.Element

	PrintValue(args ...interface{})
	PrintAtomValues(args ...interface{})
}

// list only:
type listonly interface {
	node

	InsertAfter(*list.Element) *list.Element
	InsertBefore(*list.Element) *list.Element
	MoveAfter(*list.Element) *list.Element
	MoveBefore(*list.Element) *list.Element
	MoveToBack() *list.Element
	MoveToFront() *list.Element
	Print(args ...interface{})
	PushBack(*list.Element) *list.Element
	PushBackList(*list.List) *list.Element
	PushFront(*list.Element) *list.Element
	PushFrontList(*list.List) *list.Element
	Remove(*list.Element) *list.Element
}

// element only
type elementonly interface {
	node

	MoveToPrevOf(*list.Element) *list.Element
	MoveToNextOf(*list.Element) *list.Element
}

/* symmetric
	New(vals...) node
	Equals(x node) bool
	With(x node) *ComposedValue
*/

var l node = list.New()
var e node = l.Root()

var _listonly = l
var _listelement = e
