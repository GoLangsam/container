// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package tag extends "container/.../tag" with the pathfriendly functions from "do/nvp"
package tag

import (
	myTag "github.com/golangsam/container/ccsafe/tag"
)

// I love to be friendly - thus: I give You a simple API!
//  Create me with New(name) or Tag(name, stuff),
//
// Note: this interface is exposed for godoc - only ;-)
type Friendly interface {
	myTag.Friendly   // inherited
	PathFriendly     // path.go
	InfoFriendly     // valuetype.go
	DeepInfoFriendly // leaftype.go
}

var _ Friendly = New("Interface satisfied? :-)")

type TagAny struct {
	myTag.TagAny // my Key Value Pair
}

// New returns a new (empty) Tag
func New(tag string) *TagAny {
	return &TagAny{*myTag.New(tag)}
}

// Tag returns a new Tag, initially set to val
func Tag(tag string, val interface{}) *TagAny {
	return &TagAny{*myTag.Tag(tag, val)}
}
