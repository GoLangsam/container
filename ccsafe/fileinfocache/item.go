// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fic

import (
	"os"
)

// Item represents the data (as read upon it's creation) of a related os.FileInfo
//
//  Note: Item is immutable, and as such safe for concurrent use.
type Item struct {
	fileInfo os.FileInfo
	byteS    []byte
}
