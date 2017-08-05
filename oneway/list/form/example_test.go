// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package form_test

import (
	"fmt"

	"github.com/golangsam/container/oneway/list"
	"github.com/golangsam/container/oneway/list/form"
)

func ExampleForm() {
	e := form.Element(*(list.NewList("TestList", "Element One").Front()))
	fmt.Println(e.Value)
	undo := e.Form(form.Value(3))
	fmt.Println(e.Value)
	redo := e.UnDo(undo)
	fmt.Println(e.Value)
	_ = e.Form(redo...)
	fmt.Println(e.Value)

	// Output:
	// Element One
	// 3
	// Element One
	// 3
}

func ExampleUnDo() {
	e := form.Element(*(list.NewList("TestList", "Element One").Front()))
	fmt.Println(e.Value)
	undo := e.Form(form.Value(3))
	fmt.Println(e.Value)
	redo := e.UnDo(undo)
	fmt.Println(e.Value)
	_ = e.Form(redo...)
	fmt.Println(e.Value)

	// Output:
	// Element One
	// 3
	// Element One
	// 3
}

func setValue(e *list.Element, value interface{}) {
	fmt.Println(e.Value)
	f := form.Element(*e)
	prev := f.Form(form.Value(value))
	defer f.UnDo(prev)
	// ... do some stuff with Elements Value being temporarily set to value
	fmt.Println(f.Value)
}

func ExampleValue() {
	e := list.NewList("TestList", "Element One").Front()
	setValue(e, 5)
	fmt.Println(e.Value)

	// Output:
	// Element One
	// 5
	// Element One
}
