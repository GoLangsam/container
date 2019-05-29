// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package form allows to form something, and to undo such forming later on.
//
// Inspired by:
//   - http://commandcenter.blogspot.com.au/2014/01/self-referential-functions-and-design.html
//   - https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
//   - https://www.calhoun.io/using-functional-options-instead-of-method-chaining-in-go/
//
// These samples support undo only for the last of several formers passed to Form.
//
// This implementation provides full undo.
package form

// DoFn is the signature of a
// self referential function.
// It returns it's undo DoFn.
type DoFn func() DoFn

// Form applies the formers
// and returns its undo function
// which, when evaluated,
// restores the Formable a
// to it's previous state.
func Form(a Formable, doit ...func(Formable) DoFn) DoFn {
	prev := make([]DoFn, 0, len(doit))
	for i := range doit {
		prev = append(prev, doit[i](a))
	}
	return func() DoFn {
		return undo(prev...)
	}
}

// undo applies the given doit functions in reverse order
// and returns it's own undo..
func undo(doit ...DoFn) DoFn { // TODO: may optimise for len(doit) == 0 and == 1
	prev := make([]DoFn, 0, len(doit))
	for i := len(doit) - 1; i >= 0; i-- {
		prev = append(prev, doit[i]())
	}
	return func() DoFn {
		return undo(prev...)
	}
}
