// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// MatchFiles matches pathName
// against the Disk (via MatchDisk/Glob) and then returns only those
// files
// the base name of which matches any of the given patternlists.
// Any eventual filesystem errors are ignored and skipped over.
func MatchFiles(pathName string, patterns ...*Pattern) (filS FsFileS) {
	dS, fS, _ := MatchDisk(pathName)
	_ = dS // Folds are ignored here
	for i := range fS {
		if ok, _ := fS[i].BaseMatches(patterns...); ok {
			filS = append(filS, fS[i])
		}
	}
	return filS
}
