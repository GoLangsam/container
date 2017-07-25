// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package das // Dictionary by any for strings

// UserFriendly - interface exposed for go doc only
//
// I love to be easy - thus: I give You a simple API!
//  Create me with New, if You like - Note: No need, I'm friendly :-)
//  and Init me to use me afresh,
//  Assign anything (as key) to any slice of strings,
//  and Fetch a value by it's key,
//  Lookup a value (or nil) by it's key,
//  Delete a key, if You don't need it any more
//  as You please :-)
//
type UserFriendly interface {
	// Following may be chained:
	Init() *Das                                  // (re)start afresh: no names, no content
	Assign(key interface{}, vals ...string) *Das // assign strings "vals" to name "key" (replacing prev. content!)
	Append(key interface{}, vals ...string) *Das // append strings "vals" to name "key" (respecting prev. content!)
	Delete(key interface{}) *Das                 // forget name "key" (and related content, if any)
	// Following may also be used in templates
	Fetch(key interface{}) ([]string, bool) // obtain content named "key"
	Lookup(key interface{}) []string        // obtain content named "key" - as (eventually empty) string
	//
	KeyS() []interface{}           // return my keys as slice (in random order)
	Das() map[interface{}][]string // return Content with sorted duplicatefree stringslices
	//
	Len() int // How many things do I contain right now?
}

var _ UserFriendly = New() // Interface satisfied? :-)

// Init - Want my content reborn empty?
func (d *Das) Init() *Das {
	d.protectMe()       // protect me, and ...
	defer d.releaseMe() // release me, let me go ...
	d.init()
	return d
}

// Append - You want me to append "val"-strings? to my "key" content
func (d *Das) Append(key interface{}, vals ...string) *Das {
	d.protectMe()       // protect me, and ...
	defer d.releaseMe() // release me, let me go ...
	r := d.val[key]
	r = append(r, vals...)
	d.val[key] = r
	return d
}

// Assign - You want to reset my "key" content "val"-strings?
func (d *Das) Assign(key interface{}, vals ...string) *Das {
	d.protectMe()       // protect me, and ...
	defer d.releaseMe() // release me, let me go ...
	d.val[key] = vals
	return d
}

// Lookup - You want my content of "key" - as (eventually empty) string-slice
func (d *Das) Lookup(key interface{}) []string {
	d.lazyInit()        // non-nil me ...
	d.l.RLock()         // protect me, and ...
	defer d.l.RUnlock() // release me, let me go ...
	if _, ok := d.val[key]; ok {
		return d.lazyS(key)
	}

	return []string{}
}

// Fetch - You want my content of "key"
func (d *Das) Fetch(key interface{}) ([]string, bool) {
	d.lazyInit()        // non-nil me ...
	d.l.RLock()         // protect me, and ...
	defer d.l.RUnlock() // release me, let me go ...
	if _, ok := d.val[key]; ok {
		return d.lazyS(key), true
	}

	return []string{}, false
}

// Delete - You want me to forget abou "key" (and it's related content)?
func (d *Das) Delete(key interface{}) *Das {
	d.protectMe()       // protect me, and ...
	defer d.releaseMe() // release me, let me go ...
	delete(d.val, key)
	return d
}

// Len - How many keys do I contain right now?
func (d *Das) Len() int {
	d.lazyInit()        // non-nil me ...
	d.l.RLock()         // protect me, and ...
	defer d.l.RUnlock() // release me, let me go ...
	return len(d.val)
}

// KeyS returns the my keys as unsorted slice
// Thus: {{ range .Range }}...{{end}} walks my (sorted) values
func (d *Das) KeyS() []interface{} {
	d.lazyInit()        // non-nil me ...
	d.l.RLock()         // protect me, and ...
	defer d.l.RUnlock() // release me, let me go ...
	var res = make([]interface{}, 0, len(d.val))
	for k := range d.val {
		res = append(res, k)
	}
	return res // fulfill the promise
}

// Das returns my complete content
func (d *Das) Das() map[interface{}][]string {
	d.lazyInit()        // non-nil me ...
	d.l.RLock()         // protect me, and ...
	defer d.l.RUnlock() // release me, let me go ...
	var res = make(map[interface{}][]string, len(d.val))
	for k := range d.val {
		res[k] = d.lazyS(k)
	}
	return res // fulfill the promise
}
