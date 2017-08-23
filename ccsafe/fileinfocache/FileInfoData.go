// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fic

import (
	"os"
)

// FiData represents the data (as read upon it's creation) of a related os.FileInfo
// and is intended to be used by higher packages such as FsCache.
//  Note: FsData is immutable, and as such safe for concurrent use.
type FiData struct {
	fileInfo os.FileInfo
	byteS    []byte
}
