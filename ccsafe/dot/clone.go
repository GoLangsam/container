// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

// Clone returns a completely new Dot tree:
// a copy of d and it's entire subtree
func (d *Dot) Clone() *Dot {
	d.l.Lock()         // protect me, and ...
	defer d.l.Unlock() // release me, let me go ...
	new := New(d.K())
	new.Tag(d.V())
	d.clone(new)
	return new
}

// clone - depth first
func (d *Dot) clone(copy *Dot) {
	for _, key := range d.S() { // children
		c := d.getChild(key)
		c.l.Lock()         // protect me, and ...
		defer c.l.Unlock() // release me, let me go ...
		new := copy.G(c.K())
		new.Tag(c.V())
		c.clone(new)
	}
}
