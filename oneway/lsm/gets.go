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
	M() map[string]string        // return Content as map[key]value-string
	S() []string                 // return my keys as sorted slice
	Range() []interface{}        // return my values as slice, sorted by name
	LSM() map[string]interface{} // return Content as is
}

var _ AccessFriendly = New() // Interface satisfied? :-)

// M returns the content as map[key]value-string
// Thus: {{ .M.key }} accesses the map for key="key"
// and returns its content as a slice
func (d *LazyStringerMap) M() map[string]string {
	d.lazyInit()        // non-nil me ...
	d.l.RLock()         // protect me, and ...
	defer d.l.RUnlock() // release me, let me go ...
	return d.lazyM()    // fulfill the promise
}

// S returns the my keys as sorted slice
// Thus: {{ range .S }}...{{end}} walks my keys
func (d *LazyStringerMap) S() []string {
	d.lazyInit()        // non-nil me ...
	d.l.RLock()         // protect me, and ...
	defer d.l.RUnlock() // release me, let me go ...
	return d.lazyS()    // fulfill the promise
}

// Range returns the my values as sorted slice
// Thus: {{ range .Range }}...{{end}} walks my (sorted) values
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

// LSM returns my complete content
func (d *LazyStringerMap) LSM() map[string]interface{} {
	d.lazyInit()        // non-nil me ...
	d.l.RLock()         // protect me, and ...
	defer d.l.RUnlock() // release me, let me go ...
	return d.val
}
