// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

import (
	"fmt"
)

func ExampleList_ForEachNext() {

	// Create a new list with some elements.
	var l = NewList("ForEach", "A", "B", "C", "D", "E", "F", "G")
	var e = l.Front().Next().Next().Next() // D

	fmt.Print("l.ForEachNext\t")                       // Prefix
	l.ForEachNext(func(e *Element) { e.PrintValue() }) // each Value
	fmt.Println("<")                                   // Suffix line

	fmt.Print("e.ForEachNext\t")                       // Prefix
	e.ForEachNext(func(e *Element) { e.PrintValue() }) // each Value
	fmt.Println("<")                                   // Suffix line

	// Notice the subtle difference:
	//  - for a list l all elements are iterated along the list l
	//  - for an element e all *other* elements are iterated around e

	// Output:
	// l.ForEachNext	ABCDEFG<
	// e.ForEachNext	EFGABC<
}

func ExampleElement_ForEachNext() {

	// Create a new list with some elements.
	var l = NewList("ForEach", "A", "B", "C", "D", "E", "F", "G")
	var e = l.Front().Next().Next().Next() // D

	fmt.Print("l.ForEachNext\t")                       // Prefix
	l.ForEachNext(func(e *Element) { e.PrintValue() }) // each Value
	fmt.Println("<")                                   // Suffix line

	fmt.Print("e.ForEachNext\t")                       // Prefix
	e.ForEachNext(func(e *Element) { e.PrintValue() }) // each Value
	fmt.Println("<")                                   // Suffix line

	// Notice the subtle difference:
	//  - for a list l all elements are iterated along the list l
	//  - for an element e all *other* elements are iterated around e

	// Output:
	// l.ForEachNext	ABCDEFG<
	// e.ForEachNext	EFGABC<
}

func ExampleList_ForEachPrev() {

	// Create a new list with some elements.
	var l = NewList("ForEach", "A", "B", "C", "D", "E", "F", "G")
	var e = l.Front().Next().Next().Next() // D

	fmt.Print("l.ForEachPrev\t")                       // Prefix
	l.ForEachPrev(func(e *Element) { e.PrintValue() }) // each Value
	fmt.Println("<")                                   // Suffix line

	fmt.Print("e.ForEachPrev\t")                       // Prefix
	e.ForEachPrev(func(e *Element) { e.PrintValue() }) // each Value
	fmt.Println("<")                                   // Suffix line

	// Notice the subtle difference:
	//  - for a list l all elements are iterated along the list l
	//  - for an element e all *other* elements are iterated around e

	// Output:
	// l.ForEachPrev	GFEDCBA<
	// e.ForEachPrev	CBAGFE<
}

func ExampleElement_ForEachPrev() {

	// Create a new list with some elements.
	var l = NewList("ForEach", "A", "B", "C", "D", "E", "F", "G")
	var e = l.Front().Next().Next().Next() // D

	fmt.Print("l.ForEachPrev\t")                       // Prefix
	l.ForEachPrev(func(e *Element) { e.PrintValue() }) // each Value
	fmt.Println("<")                                   // Suffix line

	fmt.Print("e.ForEachPrev\t")                       // Prefix
	e.ForEachPrev(func(e *Element) { e.PrintValue() }) // each Value
	fmt.Println("<")                                   // Suffix line

	// Notice the subtle difference:
	//  - for a list l all elements are iterated along the list l
	//  - for an element e all *other* elements are iterated around e

	// Output:
	// l.ForEachPrev	GFEDCBA<
	// e.ForEachPrev	CBAGFE<
}
