// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

// node abstracts the distinction between List and Element
// - the focus is on common behaviour.
type node interface {
	Init() *List
	List() *List
	Front() *Element
	Back() *Element
	Next() *Element
	Prev() *Element
	Root() *Element
	Len() int
	// vLen() int
	CVs() *ComposedValue
	IsAtom() bool
	IsComposed() bool

	ForEachNext(f func(*Element))
	ForEachPrev(f func(*Element))

	AtomValues() Values
	Elements() []*Element

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
With(x node) *ComposedValue
*/

var l node = New()
var _ node = l.Root()
