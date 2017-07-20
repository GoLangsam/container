// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package das // Dictionary by any for strings

import (
	"sort"
)

// PerformanceFriendly - interface exposed for godoc only
//
// I love to be fast :-)
//  Thus: I memoize answers about my content, and about when to forget my memos
//
// I love to be lazy - do not like to do things over and over again.
//	Thus: only when You ask the question, then, on Your demand, so to say
//	do I prepare the answer for such certain question about my content. ;-)
//
type PerformanceFriendly interface {
	forget()
	lazyInit()
	lazyS(key interface{}) []string
}

var _ PerformanceFriendly = New() // Interface satisfied? :-)

// helper to forget - "destroy my being valuable" :-)
func (d *Das) forget() {
	// currently: no-op
}

// helper for init-on-demand
func (d *Das) lazyInit() {
	if d == nil {
		d = New()
	}
	if d.val == nil {
		d = d.init()
	}
}

// helper to obtain the val-map: []string as sorted slice
func (d *Das) lazyS(key interface{}) []string {
	d.l.RUnlock()     // release my RLock, and ...
	defer d.l.RLock() // restore my RLock ...
	d.protectMe()
	defer d.releaseMe()

	slice := make([]string, 0, len(d.val[key]))

	for _, s := range d.val[key] {
		slice = append(slice, s) // collect the values
	}
	sort.Strings(slice) // and sort 'em
	// TODO: need to filter duplicates! And panic, if not well sorted!
	return slice
}
