// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lsm

import (
	"sort"

	"github.com/golangsam/do/ats"
)

// PerformanceFriendly - interface exposed for go doc only
//
// I love to be fast :-)
//  Thus: I memoize answers about my content, and about when to forget my memos
//
// I love to be lazy - do not like to do things over and over again.
//	Thus: only when You ask the question, then, on Your demand, so to say
//	do I prepare the answer for such certain question about my content. ;-)
//
type PerformanceFriendly interface {
	forget()                  // helper to forget - "destroy my being valuable" :-)
	lazyInit()                // helper for init-on-demand
	lazyS() []string          // helper to rebuild and keep hold of val-map[keys] as sorted slice
	lazyM() map[string]string // helper to rebuild and keep hold of val-map[keys] as map of strings
}

// helper to forget - "destroy my being valuable" :-)
func (d *LazyStringerMap) forget() {
	if len(d.val) != len(d.s) {
		d.s = d.s[:0] // forget the slice
	} // else: no new key
	d.m = nil // destroy my being valuable
}

// helper for init-on-demand
func (d *LazyStringerMap) lazyInit() {
	if d == nil {
		d = New()
	}
	if d.val == nil {
		d = d.init()
	}
}

// helper to rebuild and keep hold of val-map[keys] as sorted slice
func (d *LazyStringerMap) lazyS() []string {
	if len(d.val) == len(d.s) {
		return d.s // no new keys
	}
	d.l.RUnlock()       // release my RLock, and ...
	defer d.l.RLock()   // restore my RLock ...
	d.protectMe()       // protect me, and ...
	defer d.releaseMe() // release me, let me go ...

	d.s = d.s[:0] // start afresh
	for k := range d.val {
		d.s = append(d.s, k) // collect the keys
	}
	sort.Strings(d.s) // and sort 'em
	return d.s
}

// helper to rebuild and keep hold of val-map[keys] as map of strings
func (d *LazyStringerMap) lazyM() map[string]string {
	if d.m != nil {
		return d.m // no new keys
	}
	d.l.RUnlock()       // release my RLock, and ...
	defer d.l.RLock()   // restore my RLock ...
	d.protectMe()       // protect me, and ...
	defer d.releaseMe() // release me, let me go ...

	d.m = make(map[string]string, d.Len()) // start afresh
	for key, val := range d.val {
		d.m[key] = ats.GetString(val) // build the map
	}
	return d.m
}

// helper to transform a map of strings into a sorted slice of strings
func sortM(m map[string]string) []string {
	slice := make([]string, 0, len(m))
	for key := range m {
		slice = append(slice, key)
	}
	sort.Strings(slice)
	return slice
}
