// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

import (
	"github.com/GoLangsam/container/oneway/list"
)

// Times returns a new list: the cross product of l with some lists...
// ( recursively as [[[ l * l ] * l ] ... ] )
//
// Note: Times( l, nil ) returns a new empty list
// the root of which carries the CVs of the original l.Root()
// and the elements carry the CVs of the original elements
// Note: The Away's in the new list point to nil - thus, the new list is isolated.
func Times(l *list.List, lists ...*list.List) *list.List {
	n := len(lists)
	switch {
	case n == 0:
		return times(l, nil)
	case n == 1:
		return times(l, lists[0])
	default:
		return times(l, Times(lists[0], lists[1:]...))
	}
}

// ===========================================================================

// times returns a new list with len(X) * len(Y) Elements
// representing the cross-product of the list X * Y
//
// Note: l.times( nil ) returns a new list with no elements
// Note: The Away's in the new list point to nil - thus, the new list is isolated.
func times(X, Y *list.List) (l *list.List) {
	if X == nil {
		return X.New()
	}

	l = X.New(X.CVs())
	if Y != nil {
		for x := X.Front(); x != nil; x = x.Next() {
			for y := Y.Front(); y != nil; y = y.Next() {
				l.PushBack(x.With(y))
			}
		}
		l.Root().Value = X.With(Y)
	}
	return l
}
