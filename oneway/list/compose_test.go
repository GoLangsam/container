// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

func ExampleList_With() {

	var l1 = NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.
	var l2 = NewList("List #2", 111, 222, 333, 444, 555, 666, 777) // And another list with some elements.

	var l3 = NewList(l1.With(l2))

	l3.PrintAtomValues() // Show

	// Output:
	// List=Example|List #2	Total=0
}

func ExampleElement_With() {

	var l = NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.
	var e = l.Front().Next().Next().Next()                        // D
	var f = e.Next().Next()                                       // F
	l.PrintAtomValues("l.before\t")                               // Show

	l.PushBack(e.With(f)) // move D before F (after E, that is)

	l.PrintAtomValues("l.after\t") // Show again

	// Output:
	// l.before	: List=Example	A | B | C | D | E | F | G | Total=7
	// l.after	: List=Example	A | B | C | D | E | F | G | D|F | Total=8
}
