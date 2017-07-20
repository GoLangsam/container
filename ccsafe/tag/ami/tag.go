// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tag extends "container/.../tag" with the pathfriendly functions from "do/nvp"
package tag

import (
	myTag "github.com/golangsam/container/ccsafe/tag"
)

// Friendly - interface exposed for go doc only
//
// I love to be friendly - thus: I give You a simple API!
//  Create me with New(name) or Tag(name, stuff),
//
type Friendly interface {
	myTag.Friendly   // inherited
	PathFriendly     // path.go
	InfoFriendly     // valuetype.go
	DeepInfoFriendly // leaftype.go
}

var _ Friendly = New("Interface satisfied? :-)")

// TagAny is the type provided by package tag
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
