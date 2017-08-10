// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

// AssignFriendly - interface exposed for go doc only
type AssignFriendly interface {
	Assignss(vals ...string) *Dot           // Assign/overwrite Content with val - given as strings
	AssignSs(val ...[]string) *Dot          // Assign/overwrite Content with val - given as slices of strings
	AssignMs(val ...map[string]string) *Dot // Assign/overwrite Content with val - given as maps of strings
}

// Value modifiers - concurrency safe

// Assignss adds to (or replaces with)
// content below current dot d
// using given variadic strings
func (d *Dot) Assignss(vals ...string) *Dot {
	d.l.Lock()         // protect me, and ...
	defer d.l.Unlock() // release me, let me go ...
	for i := range vals {
		d = d.add(vals[i]) // fulfill the promise
	}
	return d
}

// AssignSs adds to (or replaces with)
// content below current dot d
// using given variadic string-slices
func (d *Dot) AssignSs(val ...[]string) *Dot {
	d.l.Lock()         // protect me, and ...
	defer d.l.Unlock() // release me, let me go ...
	for i := range val {
		for j := range val[i] { // same as Assign()
			d = d.add(val[i][j]) // fulfill the promise
		}
	}
	return d
}

// AssignMs adds to (or replaces with)
// content below current dot d
// using given variadic string-maps
func (d *Dot) AssignMs(val ...map[string]string) *Dot {
	d.l.Lock()         // protect me, and ...
	defer d.l.Unlock() // release me, let me go ...
	for i := range val {
		d = d.addM(val[i]) // fulfill the promise
	}
	return d
}
