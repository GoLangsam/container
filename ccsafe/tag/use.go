// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tag

import (
	"github.com/GoLangsam/do/ats"
)

// UserFriendly - interface exposed for go doc only
//
// I love to be easy - easy to use:
//  use me to Tag stuff,
//  use K (or String) to get my name
//  use V to get the name of my stuff
//  use GetV to get my stuff
//  use me, as You please :-)
//
// Hint: I behave like a named variable - just with other names
//  Thus: You may like to use me, where 'normal' methodnames are used otherwise ;-)
//
type UserFriendly interface {
	Tag(val interface{}) // Set/replace AnyValue/Payload
	String() string      // implement `fmt.Stringer`
	K() string           // returns my Key as string (shortcut for String())
	V() string           // returns my Value as string (via ats.GetString)
	GetV() interface{}   // returns my Value as is
}

// Tag attaches me to my AnyValue/Payload
func (d *TagAny) Tag(val interface{}) {
	d.l.Lock()         // proctect me, and ...
	defer d.l.Unlock() // release me, let me go ...
	d.v = val
}

// String implements `fmt.Stringer`
func (d *TagAny) String() string {
	d.l.Lock()         // proctect me, and ...
	defer d.l.Unlock() // release me, let me go ...
	return d.k
}

// K returns the tag string (= Key)
func (d *TagAny) K() string {
	d.l.Lock()         // proctect me, and ...
	defer d.l.Unlock() // release me, let me go ...
	return d.k
}

// V returns my AnyValue/Payload as string
func (d *TagAny) V() string {
	d.l.Lock()         // proctect me, and ...
	defer d.l.Unlock() // release me, let me go ...
	return ats.GetString(d.v)
}

// GetV returns my AnyValue/Payload
func (d *TagAny) GetV() interface{} {
	d.l.Lock()         // proctect me, and ...
	defer d.l.Unlock() // release me, let me go ...
	return d.v
}
