// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package form

import (
	"github.com/GoLangsam/container/oneway/list"
)

// Value sets Element's Value to v.
// and returns it's undo function.
func Value(v interface{}) DoFn {
	return func(e *list.Element) (undo DoFn) {
		undo = Value(e.Value) // refer to self
		e.Value = v           // do it
		return                // done
	}
}
