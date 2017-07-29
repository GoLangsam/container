// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

// UserFriendly - interface exposed for go doc only
type UserFriendly interface {
	AddStrings(key string, val ...string) *Dot         // adds key to d, and adds variadic strings below key
	AddStringS(key string, val ...[]string) *Dot       // adds key to d, and adds slices below key
	AddMap(key string, vals ...map[string]string) *Dot // adds key to d, and adds map(s) below key
}

var _ UserFriendly = New("Interface satisfied? :-)")

// Creators - concurrency safe

// AddStrings adds key to current dot d
// and adds content below this key
// using given variadic strings
func (d *Dot) AddStrings(key string, val ...string) *Dot {
	d.l.Lock()           // protect me, and ...
	defer d.l.Unlock()   // release me, let me go ...
	c := d.getChild(key) // get key
	c.l.Lock()           // protect it, and ...
	defer c.l.Unlock()   // release it, let it go ...
	c.add(val...)        // fulfill the promise
	return d
}

// AddStringS adds key to current dot d
// and adds content below this key
// using given variadic string-slices
func (d *Dot) AddStringS(key string, val ...[]string) *Dot {
	d.l.Lock()           // protect me, and ...
	defer d.l.Unlock()   // release me, let me go ...
	c := d.getChild(key) // get key
	c.l.Lock()           // protect it, and ...
	defer c.l.Unlock()   // release it, let it go ...
	for _, vals := range val {
		c.add(vals...) // fulfill the promise
	}
	return d
}

// AddMap adds key to current dot d
// and adds content below this key
// using given variadic string-maps
func (d *Dot) AddMap(key string, val ...map[string]string) *Dot {
	d.l.Lock()           // protect me, and ...
	defer d.l.Unlock()   // release me, let me go ...
	c := d.getChild(key) // get key
	c.l.Lock()           // protect it, and ...
	defer c.l.Unlock()   // release it, let it go ...
	c.addM(val...)       // fulfill the promise
	return d
}
