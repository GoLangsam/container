// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lsm

import (
	"sync"
)

// I love to contain strings, named strings, named things, things which can name themselfs.
//	And I love to give them back to You, when You need 'em.
//	And also their names -as slice or map- as You need 'em.
//	And also sorted, or reversed, all for Your convenience.
// Note: this interface is exposed for godoc - only ;-)
type Friendly interface {
	AccessFriendly      // get.go
	UserFriendly        // use.go
	PerformanceFriendly // lazy.go
}

var _ Friendly = New() // Interface satisfied? :-)

// I love to be easy - thus: I give You a simple API!
//	Create me with New, and Assign or Lookup, and Delete or Init as You please :-)
//
// I love to be responsive :-)
//	Get my content as sorted slice: S or as map: M (or even as is: LSM)
//
// I love to be fast :-)
//  Thus: I memoize answers about my content, and about when to forget my memos
//
// I love to be lazy - do not like to do things over and over again.
//	Thus: only when You ask the question, then, on Your demand, so to say
//	do I prepare the answer for such certain question about my content. ;-)
//
// I love to be concurrency-safe :-)
//  Thus: I always protect myself - Alas: You may not copy myself after frist use!
//
type LazyStringerMap struct {
	val    map[string]interface{} // content: M() or S()
	m      map[string]string      // on-demand buffer for content as map of strings
	s      []string               // on-demand buffer for content as ascending sorted slice of strings
	l      sync.RWMutex           // concurency included - we care for it - and hide it
	noCopy noCopy                 // Important: Do not copy me after first use!
}

// Note: I forget() m and s upon any eventual change to val,
// and recreate -on demand!- e.g. via lazyS

// New - my creator - for good orders sake ;-)
//
// Note: no need to call me - I use lazyInit to care for myself :-)
//
// Hint: just plug me into Your "type favourite structure{}" :-)
//
func New() *LazyStringerMap {
	d := new(LazyStringerMap)
	d = d.init()
	return d
}

// Accessors - internal

// Want my content reborn empty?
func (d *LazyStringerMap) init() *LazyStringerMap {
	d.val = make(map[string]interface{})
	d.s = make([]string, 0, 16)
	d.m = nil
	d.forget() // destroy my being valueable, if need
	return d
}

func (d *LazyStringerMap) protectMe() {
	d.lazyInit() // non-nil me ...
	d.l.Lock()   // protect me, and ...
}

func (d *LazyStringerMap) releaseMe() {
	d.forget()   // destroy my being valueable, if need
	d.l.Unlock() // release me, let me go ...
}
