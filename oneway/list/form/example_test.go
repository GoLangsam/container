// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package form_test

import (
	"fmt"

	"github.com/GoLangsam/container/oneway/list"
	"github.com/GoLangsam/container/oneway/list/form"
)

// Value returns a function which
// sets Formable's Value to v
// and returns it's undo DoFn.
func Value(v interface{}) func(form.Formable) form.DoFn {
	return func(a form.Formable) form.DoFn {
		prev := a.Value
		a.Value = v
		return func() form.DoFn {
			return Value(prev)(a)
		}
	}
}

func Example() {
	e := list.NewList("TestList", "Element One").Front()
	fmt.Println(e.Value) // Element One

	undo := Value(3)(e)
	fmt.Println(e.Value) // 3 (temporarily)

	redo := undo()
	fmt.Println(e.Value) // Element One (temporary setting undone)

	rere := redo()
	fmt.Println(e.Value) // 3 (undo undone)

	_ = rere()
	fmt.Println(e.Value) // Element One (temporary setting undone)

	// Output:
	// Element One
	// 3
	// Element One
	// 3
	// Element One
}

func ExampleForm() {
	e := list.NewList("TestList", "Element One").Front()
	fmt.Println(e.Value) // Element One

	undo := form.Form(e, Value(3), Value("5"), Value(7))
	fmt.Println(e.Value) // 7 (temporarily)

	redo := undo()
	fmt.Println(e.Value) // Element One (temporary setting undone)

	_ = redo()
	fmt.Println(e.Value) // 7 (undo undone)

	// Output:
	// Element One
	// 7
	// Element One
	// 7
}

func ExampleValue() {
	e := list.NewList("TestList", "Element One").Front()

	setValue := func(e *list.Element, v interface{}) {
		// upon exit apply undo to restore original value while setting to new value v now via
		defer Value(v)(e)() // Note the triple evaluation.

		// ... do some stuff with Value being temporarily set to v.
		fmt.Println(e.Value) // Changed Value
	}

	fmt.Println(e.Value) // Original Value
	setValue(e, 5)
	fmt.Println(e.Value)

	// Output:
	// Element One
	// 5
	// Element One
}
