// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list_test

import (
	"github.com/golangsam/container/oneway/list"

	"fmt"
)

func ExampleElement_ForEachNext() {
	fmt.Println("Starting")

	// Create a new list and put some numbers in it.
	var l = list.NewList("ForEach", "A", "B", "C", "D", "E", "F", "G")
	var e = l.Front().Next().Next().Next() // D

	fmt.Print("l.ForEachNext\t"); l.ForEachNext( func(e *list.Element){e.PrintValue()} ); fmt.Println("<")
	fmt.Print("e.ForEachNext\t"); e.ForEachNext( func(e *list.Element){e.PrintValue()} ); fmt.Println("<")
	fmt.Print("l.ForEachPrev\t"); l.ForEachPrev( func(e *list.Element){e.PrintValue()} ); fmt.Println("<")
	fmt.Print("e.ForEachPrev\t"); e.ForEachPrev( func(e *list.Element){e.PrintValue()} ); fmt.Println("<")

	// Output:
	// Starting
	// l.ForEachNext	ABCDEFG<
	// e.ForEachNext	EFGABC<
	// l.ForEachPrev	GFEDCBA<
	// e.ForEachPrev	CBAGFE<
}
