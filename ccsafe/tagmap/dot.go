// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

import (
	//	"io"
	"sync"

	"github.com/golangsam/container/ccsafe/lsm"
	"github.com/golangsam/container/ccsafe/tag/ami"
)

// Dot is the type provided by package dot
type Dot struct {
	*tag.TagAny                        // key - a Key Value Pair
	*lsm.LazyStringerMap               // value(s) - an 'Anything' StringMap
	l                    *sync.RWMutex // private lock - concurency included!
}

// New returns what You need in order to keep a hand on me :-)
func New(key string) *Dot {
	dot := &Dot{
		tag.New(key),      // init key
		lsm.New(),         // init val
		new(sync.RWMutex), // mutex
	}
	return dot
}

// GoFriendly - interface exposed for go doc only
type GoFriendly interface {
	// helper for templates:
	A(vals ...string) string // Add values, and return an empty string
	G(keys ...string) *Dot   // Go into key(s)
	KV(key, val string) *Dot // Assign a Key Value Pair
	// helper for "do/dot"
	UnlockedGet(key string) (interface{}, bool)
	UnlockedAdd(key string, val ...string) (interface{}, bool)
}

// A is a helper method for templates:
// Add value(s), and return an empty string
func (d *Dot) A(vals ...string) string {
	d.l.Lock()         // protect me, and ...
	defer d.l.Unlock() // release me, let me go ...
	for i := range vals {
		d = d.add(vals[i]) // fulfill the promise
	}
	return ""
}

// G is a helper method for templates:
// Go into (eventually new!) key(s) - returns the final child (key)
func (d *Dot) G(keys ...string) *Dot {
	c := d
	for i := range keys {
		c.l.Lock()         // protect me, and ...
		defer c.l.Unlock() // release me, let me go ...
		c = c.getChild(keys[i])
	}
	return c
}

// KV is a helper method for templates:
// Assign Key & Value to current dot.
func (d *Dot) KV(key, val string) *Dot {
	d.l.Lock()         // protect me, and ...
	defer d.l.Unlock() // release me, let me go ...
	c := d.getChild(key)
	c.Tag(val)
	return d
}

/*
Locking - handling with "do/dot"
All calls into "do/dot" have d.l.Lock()
All callbacks into Dot used (via interface "do/dot.Dot")
are named "UnlockedXyz" and assume d.Lock is held
*/

// UnlockedGet is a helper method and is exported only for use by the function library "do/dot".
// It returns the (eventually new!) child (key)
//  Note: It's 2nd arg (bool) intentionally avoids usage from templates!
// Other clients must behave as if this method is not exported!
func (d *Dot) UnlockedGet(key string) (interface{}, bool) {
	//	d.l.Lock()         // protect me, and ...
	//	defer d.l.Unlock() // release me, let me go ...
	c := d.getChild(key)
	return c, true // bool avoids usage from templates!
}

// UnlockedAdd is a helper method and is exported only for use by the function library "do/dot".
// It adds key to d, and adds variadic strings below key, and returns child "key"
//  Note: It's 2nd arg (bool) intentionally avoids usage from templates!
// Other clients must behave as if this method is not exported!
func (d *Dot) UnlockedAdd(key string, val ...string) (interface{}, bool) {
	//	d.l.Lock()             // protect me, and ...
	//	defer d.l.Unlock()     // release me, let me go ...
	c := d.getChild(key)
	c.l.Lock()         // protect it, and ...
	defer c.l.Unlock() // release it, let it go ...
	c.add(val...)      // fulfill the promise
	return c, true     // bool avoids usage from templates!
}
