// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

import (
	"io"
	"sync"

	"github.com/GoLangsam/container/ccsafe/lsm"
	"github.com/GoLangsam/container/ccsafe/tag" // or "container/ccsafe/tag/ami"
)

// Dot - a tree of named anythings - useful for agnostic template-driven text generation
//
// "Templates are executed by applying them to a data structure.
// Annotations in the template refer to elements of the data structure
// (typically a field of a struct or a key in a map)
// to control execution and derive values to be displayed.
//
// Execution of the template walks the structure and sets the cursor,
// represented by a period '.' and called "dot",
// to the value at the current location in the structure as execution proceeds."
//
// quoted from `text/template/doc.go`
type Dot struct {
	*tag.TagAny                        // key - a Key Value Pair
	*lsm.LazyStringerMap               // value(s) - an 'Anything' StringMap
	*sync.RWMutex                      // public lock - concurency enabled!
	p                    *Dot          // parent: Up(), Root(), Path(), DownS()
	home                 *Dot          // home - never changes
	output               io.Writer     // for errors etc - nil means stderr; use out() accessor
	l                    *sync.RWMutex // private lock - concurency included internally!
}

// New returns what You need in order to keep a hand on me :-)
func New(key string) *Dot {
	dot := &Dot{
		tag.New(key),      // init key
		lsm.New(),         // init val
		new(sync.RWMutex), // public mutex
		nil,               // no parent
		nil,               // no home
		nil,               // no output - nil means stderr
		new(sync.RWMutex), // private mutex
	}
	dot.home = dot // home - never changes
	return dot
}

// GoFriendly - interface exposed for go doc only
type GoFriendly interface {
	// helper for templates:
	A(vals ...string) string     // Add values, and return an empty string
	G(keys ...string) *Dot       // Go into key(s)
	DotDot(key, val string) *Dot // Assign a Key Value Pair
	Dot() interface{}            // The value of the current dot
	DotDotDots() []*Dot          // The entire subtree, depth first
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
// Go down into (eventually new!) key(s) and return the final child dot (key).
func (d *Dot) G(keys ...string) *Dot {
	c := d
	for i := range keys {
		c.l.Lock()         // protect me, and ...
		defer c.l.Unlock() // release me, let me go ...
		c = c.getChild(keys[i])
	}
	return c
}

// DotDot is a helper method for templates:
// Assign Key & Value to current dot.
func (d *Dot) DotDot(key, val string) *Dot {
	d.l.Lock()         // protect me, and ...
	defer d.l.Unlock() // release me, let me go ...
	c := d.getChild(key)
	c.Tag(val)
	return d
}

// Dot is a helper method for templates:
// Dot returns the value of the current dot.
func (d *Dot) Dot() interface{} {
	d.l.Lock()         // protect me, and ...
	defer d.l.Unlock() // release me, let me go ...
	return d.GetV()
}

// DotDotDots is a helper method for templates:
// Returns the entire subtree, depth first.
func (d *Dot) DotDotDots() []*Dot {
	d.l.RLock()         // protect me, and ...
	defer d.l.RUnlock() // release me, let me go ...
	var dotdotdots = make([]*Dot, 0, d.Len())
	d.WalkDepth1st(func(d *Dot) { dotdotdots = append(dotdotdots, d) })
	return dotdotdots
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

// Try returns anything as a *Dot,
// or nil and false, iff no *Dot was given.
func Try(v interface{}) (*Dot, bool) {
	switch v := v.(type) {
	case *Dot:
		return v, true
	default:
		return nil, false
	}
}
