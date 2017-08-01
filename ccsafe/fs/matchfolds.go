// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// MatchFolds matches pathName
// against the Disk (via MatchDisk/Glob) and then returns only those
// folders/directories
// the base name of which matches any of the given patternlists.
// Any eventual filesystem errors are ignored and skipped over.
func MatchFolds(pathName string, patterns ...*Pattern) (dirS FsFoldS) {
	dS, fS, _ := MatchDisk(pathName)
	_ = fS // Files are ignored here
	for i := range dS {
		if ok, _ := dS[i].BaseMatches(patterns...); ok {
			dirS = append(dirS, dS[i])
		}
	}
	return dirS
}
