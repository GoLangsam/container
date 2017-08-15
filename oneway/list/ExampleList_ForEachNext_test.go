// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

import (
	"fmt"
)

func ExampleList_ForEachNext() {

	// Create a new list and put some numbers in it.
	var l = NewList("ForEach", "A", "B", "C", "D", "E", "F", "G")
	var e = l.Front().Next().Next().Next() // D

	fmt.Print("l.ForEachNext\t"); l.ForEachNext( func(e *Element){e.PrintValue()} ); fmt.Println("<")
	fmt.Print("e.ForEachNext\t"); e.ForEachNext( func(e *Element){e.PrintValue()} ); fmt.Println("<")

	// Output:
	// l.ForEachNext	ABCDEFG<
	// e.ForEachNext	EFGABC<
}

func ExampleList_ForEachPrev() {

	// Create a new list and put some numbers in it.
	var l = NewList("ForEach", "A", "B", "C", "D", "E", "F", "G")
	var e = l.Front().Next().Next().Next() // D

	fmt.Print("l.ForEachPrev\t"); l.ForEachPrev( func(e *Element){e.PrintValue()} ); fmt.Println("<")
	fmt.Print("e.ForEachPrev\t"); e.ForEachPrev( func(e *Element){e.PrintValue()} ); fmt.Println("<")

	// Output:
	// l.ForEachPrev	GFEDCBA<
	// e.ForEachPrev	CBAGFE<
}
