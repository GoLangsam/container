// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list_test

import (
	"testing"

	"github.com/GoLangsam/container/oneway/list"
)

func TestNil(t *testing.T) {
	var l *list.List
	var e *list.Element

	mark := e
	if e == mark || e == nil || mark == nil {
		//	skip
	} else {
		l.MoveAfter(e, mark)
		l.MoveBefore(e, mark)
	}

	if e == nil || l == nil {
		//	skip
	} else {
		l.MoveToBack(e)
		l.MoveToFront(e)
	}

	if l == nil {
		//	skip
	} else {
		_ = l.Front()
		_ = l.Back()

		_ = l.Elements()
		_ = l.Values()

		l.ForEachNext(func(e *list.Element) {})
		l.ForEachPrev(func(e *list.Element) {})
	}

}
