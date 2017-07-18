// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

type AssignFriendly interface {
	Assignss(vals ...string) *Dot           // Assign/overwrite Content with val - given as strings
	AssignSs(val ...[]string) *Dot          // Assign/overwrite Content with val - given as slices of strings
	AssignMs(val ...map[string]string) *Dot // Assign/overwrite Content with val - given as maps of strings
}

var _ AssignFriendly = New("Interface satisfied? :-)")

// Value modifiers - concurrency safe

func (d *Dot) Assignss(vals ...string) *Dot {
	d.l.Lock()         // protect me, and ...
	defer d.l.Unlock() // release me, let me go ...
	for _, v := range vals {
		d = d.add(v) // fulfill the promise
	}
	return d
}

func (d *Dot) AssignSs(val ...[]string) *Dot {
	d.l.Lock()         // protect me, and ...
	defer d.l.Unlock() // release me, let me go ...
	for _, vals := range val {
		for _, v := range vals { // same as Assign()
			d = d.add(v) // fulfill the promise
		}
	}
	return d
}

func (d *Dot) AssignMs(val ...map[string]string) *Dot {
	d.l.Lock()         // protect me, and ...
	defer d.l.Unlock() // release me, let me go ...
	for _, v := range val {
		d = d.addM(v) // fulfill the promise
	}
	return d
}
