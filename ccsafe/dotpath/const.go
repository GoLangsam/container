// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.package dotpath

package dotpath

// Combinations of Dots representing the period(s) appearing in path names
const (
	// Dot represents the period appearing in path names
	Dot = `.`
	// SingleDot - just another name for Dot
	SingleDot = Dot // just another name
	// DoubleDot represents two consecutive periods `..` appearing in path names
	DoubleDot = Dot + Dot
	// TripleDot represents three period `...` appearing in path names
	TripleDot = Dot + Dot + Dot
	// Empty represents any empty string
	Empty = ``
)

func init() { // some paranoid sanity checks ;-)
	if len(Empty) != 0 {
		panic("My empty '" + Empty + "' has non-zero length!")
	}

}
