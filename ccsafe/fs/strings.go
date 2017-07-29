// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

import (
	"path/filepath"
	"strings"
)

// Base is a typesafe convenience for filepath.Base()
func Base(name string) *FsBase {
	return newBase(filepath.Base(name))
}

// Ext is a typesafe convenience for filepath.Ext()
func Ext(name string) *FsBase {
	return newBase(filepath.Ext(name))
}

// BaseLessExt returns name.Base() less name.Ext() as *FsBase
func BaseLessExt(name string) *FsBase {
	return newBase(strings.TrimSuffix(filepath.Base(name), filepath.Ext(name)))
}
