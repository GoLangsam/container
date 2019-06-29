// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

import (
	"github.com/GoLangsam/container/oneway/list"
)

// Append returns a new list: the union of l with some lists...
// ( recursively as [[[ l + l ] + l ] ... ] )
//
// Note: Append( l, nil ) returns a new copy of l with composedValues
// the root of which carries the CVs of the original l.Root()
// and the elements carry the CVs of the original elements
// Note: The Away's in the new list point to nil - thus, the new list is isolated.
func Append(l *list.List, lists ...*list.List) *list.List {
	n := len(lists)
	switch {
	case n == 0:
		return plus(l, nil)
	case n == 1:
		return plus(l, lists[0])
	default:
		return plus(l, Append(lists[0], lists[1:]...))
	}
}

// ===========================================================================

// plus returns a new list with len(X) + len(Y) Elements
// representing the union of the list X plus Y
//
// Note: plus(X, nil ) returns a new copy of X with composedValues
// Note: The Away's in the new list point to nil - thus, the new list is isolated.
func plus(X, Y *list.List) (l *list.List) {
	if X == nil {
		return X.New()
	}

	l = X.New(X.CVs())
	for x := X.Front(); x != nil; x = x.Next() {
		l.PushBack(x.CVs())
	}
	if Y != nil {
		for y := Y.Front(); y != nil; y = y.Next() {
			l.PushBack(y.CVs())
		}
		l.Root().Value = X.With(Y)
	}
	return l
}
