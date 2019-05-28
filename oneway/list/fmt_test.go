// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list_test

import (
	"github.com/GoLangsam/container/oneway/list"
)

func ExampleList_Print() {

	var l = list.NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.

	l.Print("Show", "Me")

	// Output:
	// ShowMe: List=Example | Total=7
}

func ExampleList_PrintValue() {

	var l = list.NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.

	l.PrintValue("Show", "Me")

	// Output:
	// ShowMe: Example
}

func ExampleElement_PrintValue() {

	var l = list.NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.
	var e = l.Front().Next().Next().Next()                             // D

	e.PrintValue("Show", "Me")

	// Output:
	// ShowMe: D
}

func ExampleList_PrintAtomValues() {

	var l = list.NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.

	l.PrintAtomValues("Show", "Me")

	// Output:
	// ShowMe: List=Example	A | B | C | D | E | F | G | Total=7
}

func ExampleElement_PrintAtomValues() {

	var l = list.NewList("Example", "A", "B", "C", "D", "E", "F", "G") // Create a new list with some elements.
	var e = l.Front().Next().Next().Next()                             // D

	e.PrintAtomValues("Show", "Me")

	// Output:
	// ShowMe: Element=D.
}
