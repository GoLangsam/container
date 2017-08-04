// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package svp

import (
	"github.com/golangsam/do/ami"
)

// InfoFriendly interface - exported for go doc
//
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

// ValueTypeName returns the Name of the Type of Value
func (p *StringValuePair) ValueTypeName() string {
	return ami.TypeName(p.V())
}

// ValueTypePkgPath returns the PkgPath of the Type of Value
func (p *StringValuePair) ValueTypePkgPath() string {
	return ami.TypePkgPath(p.V())
}

// ValueTypeString returns the String of the Type of Value
func (p *StringValuePair) ValueTypeString() string {
	return ami.TypeString(p.V())
}

// ValueTypeKind returns the Kind of the Type of Value
func (p *StringValuePair) ValueTypeKind() string {
	return ami.TypeKind(p.V())
}

// ValueTypeIsComparable returns the IsComparable of the Type of Value
func (p *StringValuePair) ValueTypeIsComparable() bool {
	return ami.TypeIsComparable(p.V())
}

// ValueTypeIsVariadic returns the IsVariadic of the Type of Value
func (p *StringValuePair) ValueTypeIsVariadic() bool {
	return ami.TypeIsVariadic(p.V())
}
