// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

import (
	"github.com/GoLangsam/container/oneway/list"
)

// ===========================================================================

// Zipp returns an iterator function returning successive pairs.
func Zipp(X, Y *list.List) func() (*list.Element, *list.Element) {
	var x = X.Front()
	var y = Y.Front()

	return func() (*list.Element, *list.Element) {
		var currx = x
		var curry = y
		if x != nil {
			x = x.Next()
		}
		if y != nil {
			y = y.Next()
		}
		return currx, curry
	}
}
