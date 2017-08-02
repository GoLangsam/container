// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

// DeleteFriendly - interface exposed for go doc only
type DeleteFriendly interface {
	Deletess(vals ...string) *Dot            // Delete/remove vals from Content - given as strings
	DeleteSs(vals ...[]string) *Dot          // Delete/remove val from Content - given as slices of strings
	DeleteMs(vals ...map[string]string) *Dot // Delete/remove val from Content - given as maps of strings
}

var _ DeleteFriendly = New("Interface satisfied? :-)")

// Value modifiers - concurrency safe

// Deletess deletes / removes
// content below current dot d
// using given variadic strings
func (d *Dot) Deletess(vals ...string) *Dot {
	d.l.Lock()         // protect me, and ...
	defer d.l.Unlock() // release me, let me go ...
	for i := range vals {
		d.Delete(vals[i])
	}
	return d
}

// DeleteSs deletes / removes
// content below current dot d
// using given variadic string-slices
func (d *Dot) DeleteSs(vals ...[]string) *Dot {
	d.l.Lock()         // protect me, and ...
	defer d.l.Unlock() // release me, let me go ...
	for i := range vals {
		for j := range vals[i] {
			d.Delete(vals[i][j])
		}
	}
	return d
}

// DeleteMs deletes / removes
// content below current dot d
// using given variadic string-maps
func (d *Dot) DeleteMs(vals ...map[string]string) *Dot {
	d.l.Lock()         // protect me, and ...
	defer d.l.Unlock() // release me, let me go ...
	for i := range vals {
		for k, v := range vals[i] {
			if c, ok := d.lookupDot(k); ok { // for valid child k: delete v
				c.Delete(v)
			}
		}
	}
	return d
}
