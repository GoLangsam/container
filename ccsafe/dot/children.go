// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

type ChildFriendly interface {
	lookupDot(key string) (*Dot, bool)
	getChild(key string) *Dot
}

var _ ChildFriendly = New("Interface satisfied? :-)")

// lookupDot
func (d *Dot) lookupDot(key string) (*Dot, bool) {
	var c *Dot

	if x, ok := d.Fetch(key); !ok { // || c == nil ???
		return nil, false
	} else if c, ok = x.(*Dot); !ok {
		panic("lookupDot: Did not get type Dot from lsm")
		return nil, false //
	} else if c.p != d { // sanity check
		panic("lookupDot: Ups: I am *not* parent of child?!?") // TODO: better formatting
	} else if c.home != d.home { // sanity check
		panic("lookupDot: Ups: We do *not* have same home?!?") // TODO: better formatting
	} else {
		return c, true
	}
}

// getChild returns the child named by key, and panics if wrong type was found
//
// Note:
// - the child is handled unsafe - use with locked value container only
// - a new child is created and linked, if need
// - if d:Get(key) returns wrong type, getChild panics
func (d *Dot) getChild(key string) *Dot {
	var c *Dot

	if x, ok := d.lookupDot(key); !ok || x == nil {
		c = New(key)     // new child
		c.home = d.home  // inherit home - never changes
		c.p = d          // set me as it's parent
		d.Assign(key, c) // assign it to key
	} else {
		c = x
	}
	return c
}
