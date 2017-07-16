﻿// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tag

import (
	"github.com/golangsam/do/ami"
)

// I love to be informative - and even give metadata about my content
//  use ValueTypeName to get the name of the type of my content
//  use ValueTypePkgName to get the package name of the type of my content
//  use ValueTypeString to get a 'nick-name' of the type of my content
//  use ValueTypeKind to get the Kind of the type of my content ( int, struct, func, ...)
//  use ValueTypeIsComparable ...
//  use ValueTypeIsVariadic (only useful, if my ValueTypeKind is a function)
// Note: I use the "reflect" package to obtain metadata about my content - as You may have guessed ;-)
type InfoFriendly interface {
	ValueTypeName() string       // the type's name within its package
	ValueTypePkgPath() string    // the import path that uniquely identifies the package
	ValueTypeString() string     // may use shortened package names (e.g., base64 instead of "encoding/base64")
	ValueTypeKind() string       //
	ValueTypeIsComparable() bool //
	ValueTypeIsVariadic() bool   // (false, if Kind != Func)
}

var _ InfoFriendly = New("Interface satisfied? :-)")

//
func (d *TagAny) ValueTypeName() string {
	return ami.TypeName(d.V())
}

func (d *TagAny) ValueTypePkgPath() string {
	return ami.TypePkgPath(d.V())
}

func (d *TagAny) ValueTypeString() string {
	return ami.TypeString(d.V())
}

func (d *TagAny) ValueTypeKind() string {
	return ami.TypeKind(d.V())
}

func (d *TagAny) ValueTypeIsComparable() bool {
	return ami.TypeIsComparable(d.V())
}

func (d *TagAny) ValueTypeIsVariadic() bool {
	return ami.TypeIsVariadic(d.V())
}
