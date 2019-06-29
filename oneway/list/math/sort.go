// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

import (
	"sort"

	"github.com/GoLangsam/container/oneway/list"
)

// Arrange sorts t in place according to weight.
// The sort is stable.
func Arrange(l *list.List, weight func(*list.Element) int) *list.List {
	if l.Len() < 1 {
		return l
	}

	these := l.Elements()
	less := func(i, j int) bool { return weight(these[i]) < weight(these[j]) }
	return Stable(these, less)
}

// Stable sorts these in place according to less.
// The sort is stable.
func Stable(these []*list.Element, less func(i, j int) bool) *list.List {
	if len(these) < 1 {
		return nil
	}

	mark := these[0]
	sort.SliceStable(these, less)

	for _, e := range these {
		mark = e.MoveToPrevOf(mark)
	}
	return mark.List()
}

// Weight returns both:
//  a map of the elements with their respective weight and
//  a map of the weights with their elements
func Weight(l list.List, weight func(*list.Element) int) (map[*list.Element]int, map[int][]*list.Element) {
	var elems = make(map[*list.Element]int, l.Len())
	var sizes = make(map[int][]*list.Element, l.Len())

	var size int
	for e := l.Front(); e != nil; e = e.Next() {
		size = weight(e)
		elems[e] = size
		sizes[size] = append(sizes[size], e)
	}
	return elems, sizes
}
