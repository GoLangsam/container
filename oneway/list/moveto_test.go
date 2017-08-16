// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

func ExampleElement_MoveToPrevOf() {

	var l = NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.
	var e = l.Front().Next().Next().Next()                        // D
	var f = e.Next().Next()                                       // F
	l.PrintAtomValues("l.before\t")                               // Show

	e.MoveToPrevOf(f) // move D before F (after E, that is)

	l.PrintAtomValues("l.after\t") // Show again

	// Output:
	// l.before	: List=Example	A | B | C | D | E | F | G | Total=7
	// l.after	: List=Example	A | B | C | E | D | F | G | Total=7
}
