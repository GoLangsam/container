// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

type DeleteFriendly interface {
	Deletess(vals ...string) *Dot            // Delete/remove vals from Content - given as strings
	DeleteSs(vals ...[]string) *Dot          // Delete/remove val from Content - given as slices of strings
	DeleteMs(vals ...map[string]string) *Dot // Delete/remove val from Content - given as maps of strings
}

var _ DeleteFriendly = New("Interface satisfied? :-)")

// Value modifiers - concurrency safe

func (d *Dot) Deletess(vals ...string) *Dot {
	d.l.Lock()         // protect me, and ...
	defer d.l.Unlock() // release me, let me go ...
	for _, val := range vals {
		d.Delete(val)
	}
	return d
}

func (d *Dot) DeleteSs(vals ...[]string) *Dot {
	d.l.Lock()         // protect me, and ...
	defer d.l.Unlock() // release me, let me go ...
	for _, val := range vals {
		for _, v := range val {
			d.Delete(v)
		}
	}
	return d
}

func (d *Dot) DeleteMs(vals ...map[string]string) *Dot {
	d.l.Lock()         // protect me, and ...
	defer d.l.Unlock() // release me, let me go ...
	for _, val := range vals {
		for k, v := range val {
			if c, ok := d.lookupDot(k); ok { // for valid child k: delete v
				c.Delete(v)
			}
		}
	}
	return d
}
