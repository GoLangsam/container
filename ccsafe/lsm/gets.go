// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lsm

// AccessFriendly - interface exposed for go doc only
//
// I love to be responsive - not only for templates :-)
//	Get my content as
//      - sorted slice of strings: S
//      - sorted slice of my kind: Range
//      - string map of strings: M
//      - string map of my kind: LSM
//
type AccessFriendly interface {
	M() map[string]string        // get my content as map[key]value-string
	S() []string                 // get my keys as sorted slice
	Range() []interface{}        // get my values as slice, sorted by name
	LSM() map[string]interface{} // get my complete content as is
}

// M returns my content as map[key]value-string:
//  {{ .M.key }}
// fetches key="key" from the map
// and returns its content as string.
func (d *LazyStringerMap) M() map[string]string {
	d.lazyInit()        // non-nil me ...
	d.l.RLock()         // protect me, and ...
	defer d.l.RUnlock() // release me, let me go ...
	return d.lazyM()    // fulfill the promise
}

// S returns my keys as sorted slice:
//  {{ range .S }}...{{end}}
// walks my keys.
func (d *LazyStringerMap) S() []string {
	d.lazyInit()        // non-nil me ...
	d.l.RLock()         // protect me, and ...
	defer d.l.RUnlock() // release me, let me go ...
	return d.lazyS()    // fulfill the promise
}

// Range returns my values as slice, sorted by name:
//  {{ range .Range }}...{{end}}
// walks my (sorted) values.
func (d *LazyStringerMap) Range() []interface{} {
	d.lazyInit()        // non-nil me ...
	d.l.RLock()         // protect me, and ...
	defer d.l.RUnlock() // release me, let me go ...
	var r = make([]interface{}, 0, len(d.val))
	for _, k := range d.lazyS() {
		r = append(r, d.val[k])
	}
	return r // fulfill the promise
}

// LSM returns my complete content as is.
func (d *LazyStringerMap) LSM() map[string]interface{} {
	d.lazyInit()        // non-nil me ...
	d.l.RLock()         // protect me, and ...
	defer d.l.RUnlock() // release me, let me go ...
	return d.val
}
