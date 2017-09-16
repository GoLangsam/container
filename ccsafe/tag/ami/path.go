// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tag

import (
	"github.com/GoLangsam/do/nvp" // NameValuePair - provides recursing functions
)

// PathFriendly - interface exposed for go doc only
//
// Note: You may like to use my own kind as stuff ;-)
//  Thus: build a path to where I hide - recursively!
//  just: Using Named() in a loop may not get You much :-(
//  just the same name over and over ... and over again ...
//  thus: better use me.Into(key string) to build Your path ;-)
// How to use? Easy:
//  Use Into to hide the treasure under another name
//  Use Leaf to retrieve the treasure
//  Use NameS or Path - they'll tell You where the treasure is hidden
//  ... just in case You forgot ;-)
//
type PathFriendly interface {
	Into(key string) *TagAny // put me into a new TagAny named key
	Leaf() interface{}       // return the treasure hidden deep inside *TagAny's
	NameS() []string         // return the names leading to the hidden treasure as a slice of strings
	NamePath() string        // return the names leading to the hidden treasure as a (cleaned!) "path"
}

// Into returns a new TagAny named "key" and containing d
func (d *TagAny) Into(key string) *TagAny {
	return Tag(key, d)
}

// Leaf returns a new TagAny named "key" and containing d
func (d *TagAny) Leaf() interface{} {
	return nvp.Leaf(d)
}

// NameS returns the names leading to the hidden treasure as a slice of strings
func (d *TagAny) NameS() []string {
	return nvp.NameS(d)
}

// NamePath returns the names leading to the hidden treasure as a "path"
//  Note: as the "path" is cleaned, it may not lead You back to the treasure!
func (d *TagAny) NamePath() string {
	return nvp.NamePath(d)
}
