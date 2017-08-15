// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list_test

import (
	"github.com/golangsam/container/oneway/list"
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

/* list only:
InsertAfter
InsertBefore
MoveAfter
MoveBefore
MoveToBack
MoveToFront
Print
PushBack
PushBackList
PushFront
PushFrontList
Remove
*/

/* element only
MoveToPrevOf
*/

/* symmetric
Equals(x node) bool
With(x node) *list.ComposedValue
*/

var l node = list.New()
var _ node = l.Root()
