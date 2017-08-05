// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package form allows to Form an element, and UnDo it.
// To "Form" an element currently means to change it's Value.
package form

// https://www.calhoun.io/using-functional-options-instead-of-method-chaining-in-go/
// https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis

import (
	"github.com/golangsam/container/oneway/list"
)

// DoFn is the signature of a
// self referential function.
// It returns it's undo DoFn.
type DoFn func(f *list.Element) DoFn

// undo is used to reverse a previous application of Form(...) via UnDo(undo)
// undo is the type returned by Form(...)
// undo is a slice of forms to be applied in reverse order
// in order to undo a previously applied Form(...)
type undo []DoFn

// Form applies the specified doit forms.
// It returns an undo, a slice of forms to restore f to it's previous situation.
func Form(e *list.Element, doit ...DoFn) (undo undo) {
	undo = make([]DoFn, 0, len(doit))
	for i := range doit {
		undo = append(undo, doit[i](e))
	}
	return undo
}

// UnDo applies the specified undo.
// It returns a redo, a slice of forms to restore e to it's previous situation.
func UnDo(e *list.Element, undo undo) (redo []DoFn) {
	redo = make([]DoFn, 0, len(undo))
	for i := len(undo) - 1; i >= 0; i-- {
		redo = append(redo, undo[i](e))
	}
	return redo
}

// Value sets Element's Value to v.
// and returns it's undo form
func Value(v interface{}) DoFn {
	return func(e *list.Element) DoFn {
		previous := e.Value
		e.Value = v
		return Value(previous)
	}
}
