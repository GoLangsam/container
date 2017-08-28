// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lsm

import (
	"github.com/golangsam/do/ats" // anything to string
)

// UserFriendly - interface exposed for go doc only
//
// I love to be easy - thus: I give You a simple API!
//  Create me with New, if You like - Note: No need, I'm friendly :-)
//  and Init me to use me afresh,
//  Assign any name (as key) to any value,
//  and Fetch a value by it's name,
//  Lookup a value as string by it's name (as key),
//  Delete a key, if You don't need it any more
//  as You please :-)
//
type UserFriendly interface {
	// Following may be chained:
	Init() *LazyStringerMap                              // (re)start afresh: no names, no content
	Assign(key string, val interface{}) *LazyStringerMap // assign a string "val" to name "key"
	Delete(key string) *LazyStringerMap                  // forget name "key" (and related content, if any)
	// Following may also be used in templates
	Fetch(key string) (interface{}, bool) // obtain content named "key", iff any
	Lookup(key string) string             // obtain content named "key" - as (eventually empty) string
	//
	Len() int // How many things do I contain right now?
}

// Init - Want my content reborn empty?
func (d *LazyStringerMap) Init() *LazyStringerMap {
	d.protectMe()       // protect me, and ...
	defer d.releaseMe() // release me, let me go ...
	d.init()
	return d
}

// Assign - You want to let my content named "key" to be the "val"-string?
func (d *LazyStringerMap) Assign(key string, val interface{}) *LazyStringerMap {
	d.protectMe()       // protect me, and ...
	defer d.releaseMe() // release me, let me go ...
	d.val[key] = val
	return d
}

// Lookup - You want my content named "key" - as (eventually empty) string
func (d *LazyStringerMap) Lookup(key string) string {
	d.lazyInit()        // non-nil me ...
	d.l.RLock()         // protect me, and ...
	defer d.l.RUnlock() // release me, let me go ...
	if c, ok := d.val[key]; ok {
		return ats.GetString(c)
	}

	return ""
}

// Fetch - You want my content named "key", iff any
func (d *LazyStringerMap) Fetch(key string) (interface{}, bool) {
	d.lazyInit()        // non-nil me ...
	d.l.RLock()         // protect me, and ...
	defer d.l.RUnlock() // release me, let me go ...
	if c, ok := d.val[key]; ok {
		return c, true
	}

	return nil, false
}

// Delete - You want me to forget about name "key" (and it's related content)?
func (d *LazyStringerMap) Delete(key string) *LazyStringerMap {
	d.protectMe()       // protect me, and ...
	defer d.releaseMe() // release me, let me go ...
	delete(d.val, key)
	return d
}

// Len - How many things do I contain right now?
func (d *LazyStringerMap) Len() int {
	d.lazyInit()        // non-nil me ...
	d.l.RLock()         // protect me, and ...
	defer d.l.RUnlock() // release me, let me go ...
	return len(d.val)
}
