// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lsm

import (
	"github.com/golangsam/do/ats" // anything to string
)

// I love to be easy - thus: I give You a simple API!
//  Create me with New, if You like - Note: No need, I'm friendly :-)
//  and Init me to use me afresh,
//  Assign any name (as key) to any value,
//  and Fetch a value by it's name,
//  Lookup a value as string by it's name (as key),
//  Delete a key, if You don't need it any more
//  as You please :-)
// Note: this interface is exposed for godoc - only ;-)
type UserFriendly interface {
	// Following may be chained:
	Init() *LazyStringerMap                              // (re)start afresh: no names, no content
	Assign(key string, val interface{}) *LazyStringerMap // assign a string "val" to name "key"
	Delete(key string) *LazyStringerMap                  // forget name "key" (and related content, if any)
	// Following may also be used in templates
	Fetch(key string) (interface{}, bool) // obtain content named "key"
	Lookup(key string) string             // obtain content named "key" - as (eventually empty) string
	//
	Len() int // How many things do I contain right now?
}

var _ UserFriendly = New() // Interface satisfied? :-)

// Want my content reborn empty?
func (d *LazyStringerMap) Init() *LazyStringerMap {
	d.protectMe()       // protect me, and ...
	defer d.releaseMe() // release me, let me go ...
	d.init()
	return d
}

// You want to let my content named "key" to be the "val"-string?
func (d *LazyStringerMap) Assign(key string, val interface{}) *LazyStringerMap {
	d.protectMe()       // protect me, and ...
	defer d.releaseMe() // release me, let me go ...
	d.val[key] = val
	return d
}

// You want my content named "key" - as (eventually empty) string
func (d *LazyStringerMap) Lookup(key string) string {
	d.lazyInit()        // non-nil me ...
	d.l.RLock()         // protect me, and ...
	defer d.l.RUnlock() // release me, let me go ...
	if c, ok := d.val[key]; ok {
		return ats.GetString(c)
	} else {
		return ""
	}
}

// You want my content named "key"
func (d *LazyStringerMap) Fetch(key string) (interface{}, bool) {
	d.lazyInit()        // non-nil me ...
	d.l.RLock()         // protect me, and ...
	defer d.l.RUnlock() // release me, let me go ...
	if c, ok := d.val[key]; ok {
		return c, true
	} else {
		return nil, false
	}
}

// You want me to forget about name "key" (and it's related content)?
func (d *LazyStringerMap) Delete(key string) *LazyStringerMap {
	d.protectMe()       // protect me, and ...
	defer d.releaseMe() // release me, let me go ...
	delete(d.val, key)
	return d
}

// How many things do I contain right now?
func (d *LazyStringerMap) Len() int {
	d.lazyInit()        // non-nil me ...
	d.l.RLock()         // protect me, and ...
	defer d.l.RUnlock() // release me, let me go ...
	return len(d.val)
}
