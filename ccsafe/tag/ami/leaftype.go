// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tag

import (
	"github.com/golangsam/do/ami"
	"github.com/golangsam/do/nvp"
)

// I love to be informative - and even give metadata about my innermost content
//  use DeepValueTypeName to get the name of the type of my innermost content
//  use DeepValueTypePkgName to get the package name of the type of my innermost content
//  use DeepValueTypeString to get a 'nick-name' of the type of my innermost content
//  use DeepValueTypeKind to get the Kind of the type of my innermost content ( int, struct, func, ...)
//  use DeepValueTypeIsComparable ...
//  use DeepValueTypeIsVariadic (only useful, if my DeepValueTypeKind is a function)
// Note: I use the nvp.Leaf in order to go deep down to my innermost content
//  Thus: Use me if I have the recursive 'onion-skins' nature - my-kind as content of my-kind ... as my content.
type DeepInfoFriendly interface {
	DeepValueTypeName() string       // the type's name within its package
	DeepValueTypePkgPath() string    // the import path that uniquely identifies the package
	DeepValueTypeString() string     // may use shortened package names (e.g., base64 instead of "encoding/base64")
	DeepValueTypeKind() string       //
	DeepValueTypeIsComparable() bool //
	DeepValueTypeIsVariadic() bool   // (false, if Kind != Func)
}

var _ DeepInfoFriendly = New("Interface satisfied? :-)")

//
func (d *TagAny) DeepValueTypeName() string {
	return ami.TypeName(nvp.Leaf(d))
}

func (d *TagAny) DeepValueTypePkgPath() string {
	return ami.TypePkgPath(nvp.Leaf(d))
}

func (d *TagAny) DeepValueTypeString() string {
	return ami.TypeString(nvp.Leaf(d))
}

func (d *TagAny) DeepValueTypeKind() string {
	return ami.TypeKind(nvp.Leaf(d))
}

func (d *TagAny) DeepValueTypeIsComparable() bool {
	return ami.TypeIsComparable(nvp.Leaf(d))
}

func (d *TagAny) DeepValueTypeIsVariadic() bool {
	return ami.TypeIsVariadic(nvp.Leaf(d))
}
