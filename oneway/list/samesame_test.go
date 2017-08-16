// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

func ExampleList_Init() {

	var l = NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.

	l.PrintAtomValues("l.before\t")
	l.Init()
	l.PrintAtomValues("l.after\t")

	// Output:
	// l.before	: List=Example	A | B | C | D | E | F | G | Total=7
	// l.after	: List=Example	Total=0
}

func ExampleElement_Init() {

	var l = NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.
	var e = l.Front().Next().Next().Next()                        // D

	e.PrintValue("e.before\t")
	e.Init()
	e.PrintValue("e.after\t")

	// Output:
	// e.before	: De.after	: D
}

func ExampleList_Front() {

	var l = NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.
	var e = l.Front().Next().Next().Next()                        // D

	e.PrintValue("D = ")

	// Output:
	// D = : D
}

func ExampleElement_Front() {

	var l = NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.
	var e = l.Front().Next().Next().Next()                        // D

	e.PrintValue("D = ")

	e = e.Front().Next().Next().Next() // D
	e.PrintValue("D = ")
	// Output:
	// D = : DD = : D
}

func ExampleList_Back() {

	var l = NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.
	var e = l.Back().Prev().Prev().Prev()                         // D

	e.PrintValue("D = ")

	// Output:
	// D = : D
}

func ExampleElement_Back() {

	var l = NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.
	var e = l.Front().Next().Next().Next()                        // D

	e.PrintValue("D = ")

	e = e.Back().Prev().Prev().Prev() // D
	e.PrintValue("D = ")
	// Output:
	// D = : DD = : D
}

func ExampleList_Next() {

	var l = NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.

	if l.Next() != l.Front() {
		l.Print("Next != Front")
	}

	// Output:
}

func ExampleElement_Next() {

	var l = NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.
	var e = l.Front().Next().Next().Next()                        // D

	e.PrintValue("D = ")

	e = e.Back().Prev().Prev().Prev() // D
	e.PrintValue("D = ")
	// Output:
	// D = : DD = : D
}

func ExampleList_Prev() {

	var l = NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.

	if l.Prev() != l.Back() {
		l.Print("Prev != Back")
	}

	// Output:
}

func ExampleElement_Prev() {

	var l = NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.
	var e = l.Back().Prev().Prev().Prev()                         // D

	e.PrintValue("D = ")

	// Output:
	// D = : D
}

func ExampleList_List() {

	var l = NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.
	var e = l.Front().Next().Next().Next()                        // D

	if l.List() != e.List() {
		l.Print("List != List")
	}

	// Output:
}

func ExampleElement_List() {

	var l = NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.
	var e = l.Back().Prev().Prev().Prev()                         // D

	if e.List() != l.List() {
		l.Print("List != List")
	}

	// Output:
}

func ExampleList_Root() {

	var l = NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.
	var e = l.Front().Next().Next().Next()                        // D

	if l.Root() != e.Root() {
		l.Print("List != List")
	}

	l.Root().PrintValue("Root =")

	// Output:
	// Root =: Example
}

func ExampleElement_Root() {

	var l = NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.
	var e = l.Back().Prev().Prev().Prev()                         // D

	if e.List() != l.List() {
		l.Print("List != List")
	}

	e.Root().PrintValue("Root =")

	// Output:
	// Root =: Example
}
