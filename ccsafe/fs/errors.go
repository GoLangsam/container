// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

const (
	// Tab is the horizontal tabulation character/rune
	Tab = "\t"
)

// Errors is a slice of error
type Errors struct {
	errs []error
}

// Error returns a Tab-terminated string listing the accumulated errors
// (the string is suitable for text\tabwriter;
// if there are no errors, a single Tab is returned - an empty Tab-terminated string)
func (er Errors) Error() string {
	s := Tab
	for i := range er.errs {
		s = s + er.errs[i].Error() + Tab
	}
	return s
}

// err adds any non-nil error
func (er Errors) err(err error) {
	if err != nil {
		er.errs = append(er.errs, err)
	}
}

// got returns the error(s) got, or nil, if there are none
func (er Errors) got() error {
	switch {
	case len(er.errs) > 0:
		return er
	default:
		return nil
	}
}
