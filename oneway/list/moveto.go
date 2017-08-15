// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
moveto.go extends list.go with:

	- e.MoveToPrevOf( *Element ) *Element
	- e.MoveToNextOf( *Element ) *Element	// TODO:

Note: For good performance, the functions are implemented directly
and do not use the "l.insert(l.remove(e), mark)" pattern
*/

package list

// ========================================================

// MoveToPrevOf moves e to before at (or to the back of e.list, if at is nil)
// If at is not an element of the list of e, the list of e is not modified.
func (e *Element) MoveToPrevOf (at *Element) *Element {
	if at == e { return e }

	if at == nil {
		e.prev.next		= e.next		// Unlink e from e.prev
		e.next.prev		= e.prev		// Unlink e from e.next
		e.prev			= e.list.root.prev	// Relink e.prev to e.list.root.prev
		e.next			= &e.list.root		// Relink e.next to e.list.root
		e.list.root.prev.next	= e			// Link e to root.prev.next
		e.list.root.prev	= e			// Link e to root.prev
	} else {
		if e.list != at.list { return e }
		e.prev.next		= e.next		// Unlink e from e.prev
		e.next.prev		= e.prev		// Unlink e from e.next
		e.prev			= at.prev		// Relink e.prev to at.prev
		e.next			= at			// Relink e.next to at
		at.prev.next		= e			// Link e to at.prev.next
		at.prev			= e			// Link e to at.prev
	}

	return e
}
