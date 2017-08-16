// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

import (
	"github.com/golangsam/container/oneway/list"
)

// ===========================================================================

// foldr _ z [] = z
// foldr f z (x:xs) = f x (foldr f z xs)

// FoldAny uses f(*Element, interface{}) to fold list l (from front to back) given initial i
func FoldAny(
	f func(*list.Element, interface{}) interface{},
	l *list.List,
	i interface{},
) (result interface{}) {

	result = i

	for e := l.Front(); e != nil; e = e.Next() {
		result = f(e, result)
	}

	return result
}

// FoldInt uses f(*Element, int) to fold list l (from front to back) given initial i
func FoldInt(
	f func(*list.Element, int) int,
	l *list.List,
	i int,
) (result int) {

	result = i

	for e := l.Front(); e != nil; e = e.Next() {
		result = f(e, result)
	}

	return result
}

// FoldString uses f(*Element, string) to fold list l (from front to back) given initial i
func FoldString(
	f func(*list.Element, string) string,
	l *list.List,
	i string,
) (result string) {

	result = i

	for e := l.Front(); e != nil; e = e.Next() {
		result = f(e, result)
	}

	return result
}
