// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list_test

import (
	"github.com/golangsam/container/oneway/list"

	"fmt"
)

func ExampleElement_Value() {
	fmt.Println("Starting")

	// Create a new list and put some numbers in it.
	l := list.New()

	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("Incrementing")
	for e := l.Front(); e != nil; e = e.Next() {
		var c = e.Value.(int)
		c = c + 1
		fmt.Println(c)
	}

	// Output:
	// Starting
	// 1
	// 2
	// 3
	// 4
	// Incrementing
	// 2
	// 3
	// 4
	// 5
}
