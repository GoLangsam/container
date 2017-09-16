// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dotpath

import (
	"strings" // Note: We use only strings here! - path/filepath provides higher abstractions

	ds "github.com/GoLangsam/do/strings" // do strings
)

// fullPath = original less strip VolumeName (if any)
func (dp *DotPath) stripVolumeName() *DotPath {
	if len(dp.volumeName) > 0 {
		dp.fullPath = strings.TrimPrefix(dp.original, dp.volumeName)
	} else {
		dp.fullPath = dp.original
	}
	return dp
}

// fullPath = fullPath less leading slash(es) (if any);
// dp.rootSlashs gets 'em
func (dp *DotPath) stripPrefixSlashes() *DotPath {
	dp.rootSlashs = Empty
	for strings.HasPrefix(dp.fullPath, dp.separator) {
		dp.rootSlashs = dp.rootSlashs + dp.separator
		dp.fullPath = strings.TrimPrefix(dp.fullPath, dp.separator)
	}
	return dp
}

// fullPath = fullPath less trailing slash(es) (if any);
// dp.tailSlashs gets 'em
func (dp *DotPath) stripSuffixSlashes() *DotPath {
	dp.trailSlash = Empty
	for strings.HasSuffix(dp.fullPath, dp.separator) {
		dp.trailSlash = dp.separator // we keep only one
		dp.fullPath = strings.TrimSuffix(dp.fullPath, dp.separator)
	}
	return dp
}

// fullPath = fullPath with cleaned sequences of any multiple consecutive dots;
// dp.butDots gets true eventually
func (dp *DotPath) fullPathNoMultipleDots() *DotPath {
	todo, rest := dp.fullPath, Empty

	buff := []string{}
	for len(todo) > 0 {
		todo, rest = ds.SplitAtFirst(todo, dp.separator)
		if len(todo) < 1 { // use rest
			todo, rest = rest, Empty
		}
		char, dots := ds.SplitAllSuffixe(todo, Dot)
		if len(char) > 0 {
			dp.butDots = true
		}
		buff = append(buff, PathDotTailor(char, dots)...)
		todo, rest = rest, Empty // continue eating what's left
	}
	if len(buff) > 0 {
		dp.fullPath = strings.Join(buff, dp.separator)
	}
	return dp
}

// lessDown = lessDown less any tripledots
func (dp *DotPath) lessDownLessTripleDots() *DotPath {
	dp.lessDown = strings.Replace(dp.lessDown, dp.separator+TripleDot+dp.separator, dp.separator, -1)
	dp.lessDown = strings.Replace(dp.lessDown, dp.separator+TripleDot, dp.separator+Dot, -1)
	dp.lessDown = strings.Replace(dp.lessDown, TripleDot+dp.separator, Dot+dp.separator, -1)
	if dp.lessDown == TripleDot {
		dp.lessDown = Dot
	}
	return dp
}

// lessTail = lessTail less trailing doubledots;
// dp.goUpTail gets them appended
func (dp *DotPath) lessTailLessDoubleDots() *DotPath {
	pathBaseS := dp.PathBaseS()
	dp.lessTail = Empty

	for i, base := range pathBaseS {
		switch {
		case (base == TripleDot): // skip	TODO: panic?
		case (base == DoubleDot): // UpTail
			dp.goUpTail = append(dp.goUpTail, base)
		case (base == SingleDot): // skip
		case (base == Empty): // skip	TODO: panic?
		default: // done - reassemble by prepending remaining bases
			for j := i; j < len(pathBaseS); j++ {
				dp.lessTail = ds.Join2(pathBaseS[j], dp.lessTail, dp.separator)
			}
			return dp
		}
	}
	return dp
}

// helpers for results

// lessTail + pos-times goUpTail
func (dp *DotPath) goUpPath(pos int) string {
	return ds.Join2(dp.lessTail, strings.Join(dp.goUpTail[:pos], dp.separator), dp.separator)
}

// fullPath => list of parts ending in tripledots (if any)
func (dp *DotPath) downPathS() (pathS []string) {
	todo, rest := dp.fullPath, Empty

	for strings.Index(todo, TripleDot) > -1 {
		todo, rest = ds.SplitAtFirst(todo, TripleDot)
		pathS = append(pathS, todo)
		todo = todo + rest // start again, with tripledots removed
	}
	return pathS
}

// PathBaseS returns a slice of strings each representing a node / element / PathBase.
// The order is last-first / bottom up / reversed. All slashes are gone. VolumeName is ignored.
func (dp *DotPath) PathBaseS() (pathBaseS []string) {
	todo, rest := dp.lessTail, Empty

	for len(todo) > 0 {
		rest, todo = ds.SplitAtLastChar(todo, dp.separator)
		pathBaseS = append(pathBaseS, noEmpty(todo))
		todo, rest = strings.TrimSuffix(rest, dp.separator), Empty
	}
	if len(rest) > 0 {
		pathBaseS = append(pathBaseS, rest)
	}
	return pathBaseS
}
