// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tag

import (
	"sync"
)

// I love to be friendly - thus: I give You a simple API!
//  Create me with New(name) or Tag(name, stuff),
//
// Note: this interface is exposed for godoc - only ;-)
type Friendly interface {
	UserFriendly // use.go
}

var _ Friendly = New("Interface satisfied? :-)")

// TagAny is a concurrency safe string-named any-value (string-any pair)
type TagAny struct {
	k string      // my tag text
	v interface{} // my value
	l sync.Mutex  // my private mutex
	noCopy
}

// New returns a new (empty) Tag
func New(tag string) *TagAny {
	t := new(TagAny)
	t.k = tag
	// t.v intentionally not set
	return t
}

// Tag returns a new Tag, initially set to val
func Tag(tag string, val interface{}) *TagAny {
	t := new(TagAny)
	t.k = tag
	t.v = val
	return t
}
