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
	IsComposed() bool
	IsAtom() bool
	/*
		ForEachNext(f func(node))
		ForEachPrev(f func(node))
	*/
}

var l node = list.New()
var _ node = l.Root()
