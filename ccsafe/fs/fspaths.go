// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// Validate returns any errors encountered when validating it's elements
func (f FsPathS) Validate() error {
	er := new(Errors)
	var err error
	for i := range f {
		_, err = f[i].Stat()
		er.err(err)
	}
	return er.got()
}

// Accessible returns any errors encountered when accessing it's elements
func (f FsPathS) Accessible() error {
	er := new(Errors)
	for i := range f {
		er.err(f[i].Accessible())
	}
	return er.got()
}
