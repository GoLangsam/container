// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

import (
	"fmt"
)

func ExampleList_Elements() {

	var l = NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.

	for _, e := range l.Elements() {
		e.PrintValue()
	}

	// Output:
	// ABCDEFG
}

func ExampleElement_Elements() {

	var l = NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.
	var e = l.Front().Next().Next().Next()                        // D

	for _, ele := range e.Elements() {
		ele.PrintValue()
	}

	// Output:
	// ABCDEFG
}

func ExampleList_Values() {

	var l = NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.

	for _, v := range l.Values() {
		fmt.Print(v)
	}
	fmt.Println()

	// Output:
	// ABCDEFG
}

func ExampleElement_Values() {

	var l = NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.
	var e = l.Front().Next().Next().Next()                        // D

	for _, v := range e.Values() {
		fmt.Print(v)
	}
	fmt.Println()

	// Output:
	// ABCDEFG
}

func ExampleList_ValuesPushBack() {

	var l = NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.
	l.ValuesPushBack("H", "I", "J", "K")                          // Push some more at Back
	l.PrintAtomValues()                                           // Show

	// Output:
	// List=Example	A | B | C | D | E | F | G | H | I | J | K | Total=11
}

func ExampleList_ValuesPushFront() {

	var l = NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.
	l.ValuesPushFront("X", "Y", "Z")                              // Push some more at Front
	l.PrintAtomValues()                                           // Show

	// Output:
	// List=Example	Z | Y | X | A | B | C | D | E | F | G | Total=10
}
