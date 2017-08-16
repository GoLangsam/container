// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

func ExampleList_Equals() {

	var l1 = NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.
	var l2 = NewList("List #2", "A", "B", "C", 444, "E", "F", "G") // a different one
	var l3 = NewList("List #3", "A", "B", "C", "D", "E", "F", "G") // a similar one

	if !l1.Equals(l2) { // "D" != 444
		// fine
	} else {
		l1.Print("l1")
		l2.Print("l2")
	}

	if l1.Equals(l3) { // different list, same values in same sequence
		// fine
	} else {
		l1.Print("l1")
		l2.Print("l3")
	}

	// Output:
}

func ExampleElement_Equals() {

	var l = NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.
	var e1 = l.Front().Next().Next().Next()                       // D
	var e2 = l.Back().Prev().Prev().Prev()                        // D

	if e1.Equals(e2) { // "D" == "D"
		// fine
	} else {
		e1.PrintValue("Element != Element")
		e2.PrintValue("Element != Element")
	}

	var l3 = NewList("List #3", "A", "B", "C", "D", "E", "F", "G") // another list with same values
	var e3 = l3.Front().Next().Next().Next()                       // D

	if e1.Equals(e3) { // different list, same values: "D" == "D"
		// fine
	} else {
		e1.PrintValue("Element != Element")
		e3.PrintValue("Element != Element")
	}

	if e1.Equals(e1.Next()) { // same lists, different values: "D" != "E"
		e1.PrintValue("Element != Element")
		e3.PrintValue("Element != Element")
	} else {
		// fine
	}

	e3.Value = 4711
	if e1.Equals(e3) { // different lists, different values: "D" != 4711
		e1.PrintValue("Element != Element")
		e3.PrintValue("Element != Element")
	} else {
		// fine
	}

	// Output:
}

func ExampleElement_IsRoot() {

	var l = NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.
	var e1 = l.Front().Next().Next().Next()                       // D
	var e2 = l.Root()                                             // Root

	if e1.IsRoot() {
		e1.PrintValue("no root?")
	} else {
		// fine
	}

	if e2.IsRoot() {
		// fine
	} else {
		e2.PrintValue("is root?")
	}

	// Output:
}

func ExampleElement_IsNode() {

	var l = NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.
	var e1 = l.Front().Next().Next().Next()                       // D
	var e2 = l.Root()                                             // Root

	if e1.IsNode() {
		// fine
	} else {
		e1.PrintValue("is node?")
	}

	if e2.IsNode() {
		e2.PrintValue("is node?")
	} else {
		// fine
	}

	// Output:
}

func ExampleList_IsEmpty() {

	var l1 = New()                                                 // A new list
	var l2 = NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.
	var l3 = new(List)                                             // Root

	if l1.IsEmpty() {
		// fine
	} else {
		l1.Print("is empty?")
	}

	if l2.IsEmpty() {
		l2.PrintValue("not empty?")
	} else {
		// fine
	}

	if l3.IsEmpty() {
		// fine
	} else {
		l3.Print("is empty?")
	}

	// Output:
}
