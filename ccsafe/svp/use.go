﻿// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package svp

import (
	"github.com/golangsam/do/ats"
)

// Note: this interface is exposed for godoc - only ;-)
//
// I love to be easy - easy to use:
//  use K (or String) to get my name
//  use V to get my (named) stuff
//  use me, as You please :-)
// Hint: I behave like a named constant - just with other names
type UserFriendly interface {
	String() string    // returns my Key as string
	K() string         // returns my Key as string (same as String())
	V() string         // returns my Value as string via ats.GetString
	GetV() interface{} // returns my Value as is
}

var _ UserFriendly = New("Interface satisfied? :-)", empty)

// implement fmt.Stringer
func (p *StringValuePair) String() string {
	return p.k
}

// K returns my Key string, my name - so to say
func (p *StringValuePair) K() string {
	return p.k
}

// V returns my Value as string
func (p *StringValuePair) V() string {
	return ats.GetString(p.v)
}

// GetV returns my Value, my content - so to say
func (p *StringValuePair) GetV() interface{} {
	return p.v
}
