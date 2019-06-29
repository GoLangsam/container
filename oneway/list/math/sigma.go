// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

import (
	"github.com/GoLangsam/container/oneway/list"
)

// ===========================================================================

//  let rec sigma f = function
//     | [] -> 0
//     | x :: l -> f x + sigma f l;;
//
// val sigma : ('a -> int) -> 'a list -> int = <fun>

// Sigma returns the sum of the results of applying a given function f to each element of list l
func Sigma(
	f func(*list.Element) int,
	l *list.List,
) (result int) {

	if l != nil {
		for e := l.Front(); e != nil; e = e.Next() {
			result += f(e)
		}
	}

	return
}

// Note: Sigma may also be expressed using
//	FoldInt(func(e *list.Element, r int) int {return r + f(e)}, l, 0)

// SigmaInt returns the sum of the results of applying a given function f to each element of list l using
//	FoldInt(func(e *list.Element, r int) int {return r + f(e)}, l, 0)
func SigmaInt(
	f func(*list.Element) int,
	l *list.List,
) (result int) {
	return FoldInt(func(e *list.Element, r int) int { return r + f(e) }, l, 0)
}
