// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

import (
	"os"
)

const (
	// ListSep is the os.PathListSeparator
	ListSep = os.PathListSeparator
)

// Perm defaults to '0644' as os.FileMode
var Perm os.FileMode = 0644 // default os.FileMode

func init() { // some paranoid sanity checks ;-)
}
