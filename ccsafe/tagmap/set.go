// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

// SetableFriendly - interface exposed for go doc only
type SetableFriendly interface {
	Set(vals ...string) *Dot            // Set/replace Content with val - given as strings
	SetM(val ...map[string]string) *Dot // Set/replace Content with val - given as maps of strings
	SetS(val ...[]string) *Dot          // Set/replace Content with val - given as slices of strings
}

var _ SetableFriendly = New("Interface satisfied? :-)")

// Value modifiers - concurrency safe

// Set - set strings
func (d *Dot) Set(vals ...string) *Dot {
	d.l.Lock()         // protect me, and ...
	d.Init()           // reset my being valuable, and ...
	defer d.l.Unlock() // release me, let me go ...
	for i := range vals {
		d = d.add(vals[i]) // fulfill the promise
	}
	return d
}

// SetS - set string-slices
func (d *Dot) SetS(val ...[]string) *Dot {
	d.l.Lock()         // protect me, and ...
	d.Init()           // reset my being valuable, and ...
	defer d.l.Unlock() // release me, let me go ...
	for i := range val {
		for j := range val[i] { // same as Add()
			d = d.add(val[i][j]) // fulfill the promise
		}
	}
	return d
}

// SetM - set string-maps
func (d *Dot) SetM(val ...map[string]string) *Dot {
	d.l.Lock()         // protect me, and ...
	d.Init()           // reset my being valuable, and ...
	defer d.l.Unlock() // release me, let me go ...
	for i := range val {
		d = d.addM(val[i]) // fulfill the promise
	}
	return d
}
