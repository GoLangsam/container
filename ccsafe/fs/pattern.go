// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

import (
	"path/filepath" // filepath.SplitList for NewPatternS
)

// Pattern represents a file system pattern
//  Note: Pattern is immutable, and as such safe for concurrent use.
type Pattern struct {
	fsPath // file system pattern
}

// PatternS represents a collection (slice) of (pointers to) Pattern's.
type PatternS []*Pattern

// String returns the PatternS-slice as string.
func (f PatternS) String() string {
	var s string
	s = s + "{"
	first := true
	for i := range f {
		if first {
			first = false
		} else {
			s = s + ", "
		}
		s = s + f[i].String()
	}
	s = s + "}"
	return s
}

// ForcePattern returns a fresh Pattern representing the given a name.
func ForcePattern(name string) *Pattern {
	return newPattern(name)
}

// newPattern returns a fresh Pattern representing some file system pattern,
// based on a fresh fsPath.
func newPattern(pattern string) *Pattern {
	return &Pattern{*newPath(pattern)}

}

// AsPattern returns a fresh Pattern for the given fsPath,
// or panics, if TryPattern returns an error (ErrBadPattern).
func (p *fsPath) AsPattern() *Pattern {
	if anew, ok := p.TryPattern(); !ok {
		panic("newPattern: Match returned an ErrBadPattern-error!")
	} else {
		return anew
	}
}

// TryPattern returns a fresh Pattern for the given path,
// and false iff ErrBadPattern is returned from Match.
func (p *fsPath) TryPattern() (*Pattern, bool) {
	pattern := &Pattern{*p}
	_, err := p.Match(pattern)
	return &Pattern{*p}, (err == nil)
}

// NewPatternS returns a non-empty slice of Patterns obtained via filepath.SplitList
func NewPatternS(names ...string) (patternS PatternS) {
	if len(names) < 1 {
		patternS = append(patternS, newPattern(""))
	} else {
		for i := range names {
			parts := filepath.SplitList(names[i])
			for i := range parts {
				patternS = append(patternS, newPattern(parts[i]))
			}
		}
	}
	return patternS
}
