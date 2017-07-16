// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

type SetableFriendly interface {
	Set(vals ...string) *Dot            // Set/replace Content with val - given as strings
	SetM(val ...map[string]string) *Dot // Set/replace Content with val - given as maps of strings
	SetS(val ...[]string) *Dot          // Set/replace Content with val - given as slices of strings
}

var _ SetableFriendly = New("Interface satisfied? :-)")

// Value modifiers - concurrency safe

func (d *Dot) Set(vals ...string) *Dot {
	d.l.Lock()         // protect me, and ...
	d.Init()           // reset my being valueable, and ...
	defer d.l.Unlock() // release me, let me go ...
	for _, v := range vals {
		d = d.add(v) // fullfill the promise
	}
	return d
}

func (d *Dot) SetM(val ...map[string]string) *Dot {
	d.l.Lock()         // protect me, and ...
	d.Init()           // reset my being valueable, and ...
	defer d.l.Unlock() // release me, let me go ...
	for _, v := range val {
		d = d.addM(v) // fullfill the promise
	}
	return d
}

func (d *Dot) SetS(val ...[]string) *Dot {
	d.l.Lock()         // protect me, and ...
	d.Init()           // reset my being valueable, and ...
	defer d.l.Unlock() // release me, let me go ...
	for _, vals := range val {
		for _, v := range vals { // same as Add()
			d = d.add(v) // fullfill the promise
		}
	}
	return d
}
