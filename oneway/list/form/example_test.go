// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package form

import (
	"fmt"

	"github.com/GoLangsam/container/oneway/list"
)

func ExampleForm() {
	e := list.NewList("TestList", "Element One").Front()
	fmt.Println(e.Value) // Element One

	undo := Form(e, Value(3))
	fmt.Println(e.Value) // 3 (temporarily)

	redo := UnDo(e, undo)
	fmt.Println(e.Value) // Element One (temporary setting undone)

	_ = Form(e, redo...)
	fmt.Println(e.Value) // 3 (undo undone)

	// Output:
	// Element One
	// 3
	// Element One
	// 3
}

func ExampleUnDo() {
	e := list.NewList("TestList", "Element One").Front()
	fmt.Println(e.Value) // Element One

	undo := Form(e, Value(3))
	fmt.Println(e.Value) // 3 (temporarily)

	redo := UnDo(e, undo)
	fmt.Println(e.Value) // Element One (temporary setting undone)

	_ = Form(e, redo...)
	fmt.Println(e.Value) // 3 (undo undone)

	// Output:
	// Element One
	// 3
	// Element One
	// 3
}

func ExampleValue() {
	e := list.NewList("TestList", "Element One").Front()

	setValue := func(e *list.Element, value interface{}) {
		fmt.Println(e.Value) // Original Value

		// upon exit restore Original while setting to new value now via
		// undo := Form(e, Value(value))
		defer UnDo(e, Form(e, Value(value)))

		// ... do some stuff with Elements Value being temporarily set to value
		fmt.Println(e.Value) // Changed Value
	}

	setValue(e, 5)
	fmt.Println(e.Value)

	// Output:
	// Element One
	// 5
	// Element One
}
