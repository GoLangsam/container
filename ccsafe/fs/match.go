// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

import (
	"path/filepath"
)

// Match reports whether name matches any of the shell file name pattern lists.
//  Note: any name matches an empty patternlist and any empty pattern!
func Match(name string, patterns ...*Pattern) (matched bool, err error) {
	var match bool
	for i := range patterns {
		p := patterns[i].String()
		if len(p) > 0 {
			if ok, err := filepath.Match(p, name); err != nil {
				return false, err
			} else if !ok {
				// keep looking
			} else {
				match = true
				break
			}
		} else {
			match = true
			break
		}
	}
	if len(patterns) > 0 {
		return match, nil
	}

	return true, nil
}
