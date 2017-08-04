// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package das provides a Dictionary by Anything for Strings
package das

import (
	"sync"
)

// Das - a Dictionary by Anything for Strings
//
// I love to be easy - thus: I give You a simple API!
//	Create me with 'New', and then
//	'Assign' and/or 'Append' and/or 'Add',
//	'Fetch' and/or 'Lookup', and
//	'Delete' or 'Init' - as You please :-)
//
// I love to be responsive :-)
//	Get my 'KeyS', or 'Lookup' my (sorted!) content as slice,
//	or as map with sorted! strings: 'Das' - Dictionary by any for strings
//
// I love to be lazy - do not like to do things over and over again.
//	Thus: only when You ask the question, then, on Your demand, so to say
//	do I prepare the answer for such certain question about my content. ;-)
//
// I love to be concurrency-safe :-)
//	Thus: I always protect myself!
//	Alas: You may not copy myself after frist use!
//
type Das struct { // Dictionary by Anything for Strings
	val    map[interface{}][]string // content: M() or S()
	l      sync.RWMutex             // concurency included - we care for it - and hide it
	noCopy noCopy                   // Important: Do not copy me after first use!
}

// Note: I forget() m and s upon any eventual change to val,
// and recreate -on demand!- e.g. via lazyS

// New - my creator - for good orders sake ;-)
//
// Note: no need to call me - I use lazyInit to care for myself :-)
//
// Hint: just plug me into Your "type favourite structure{}" :-)
//
func New() *Das {
	d := new(Das)
	d = d.init()
	return d
}

// Accessors - internal

// Want my content reborn empty?
func (d *Das) init() *Das {
	d.val = make(map[interface{}][]string)
	d.forget() // destroy my being valuable, if need
	return d
}

func (d *Das) protectMe() {
	d.lazyInit() // non-nil me ...
	d.l.Lock()   // protect me, and ...
}

func (d *Das) releaseMe() {
	d.forget()   // destroy my being valuable, if need
	d.l.Unlock() // release me, let me go ...
}
