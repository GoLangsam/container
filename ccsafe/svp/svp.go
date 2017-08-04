// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package svp - StringValuePair - named constants
//
//  Note: Being immutable implies concurrency safetey.
package svp // StringValuePair

// StringValuePair - aka svp - named constants
type StringValuePair struct {
	k string      // my key text
	v interface{} // my value
}

// New returns a new StringValuePair, named key, containing string val
func New(key string, val interface{}) *StringValuePair {
	return &StringValuePair{key, val}
}
