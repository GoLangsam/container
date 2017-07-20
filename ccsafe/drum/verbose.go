// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package drum

import "fmt"

// VerboseType - usage pattern:
//  Verbose = VerboseType(true)
// or: let v be e.g. a boolean flag
//  Verbose = VerboseType(v)
type VerboseType bool

var (
	// Verbose determines, if verbose is on or off (=default)
	Verbose VerboseType
)

// Printf prints iff v is true/on
func (v VerboseType) Printf(s string, a ...interface{}) {
	if v {
		fmt.Printf(s, a...)
	}
}
