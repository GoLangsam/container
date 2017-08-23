// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package dotpath is all about Pathinking - Path..think.ing - Path.ink.ing - Pa.think.ing - Pa.thin.king ...
//
// Package dotpath is intended as a parser for user-provided strings
// separated by some single character delimiter (which is not a Dot)
// such as filepath information or other hierarchical identifiers
// and gives an extended meaning to multiple dots and trailing slashes:
//
//  ... => 'use Subtree also' - recurse
//  ..  => 'use Parent also', if .. is part of trailing dots and slashes
//  ./  => 'this MUST be a directory', if a trailing slash is given
//
// lexical.go: strictly lexical analysis of fullPath
//
// lexical.go uses some dedicated functions exported by pathsplitter.go, and the packages "strings" and "do/string".
//
// Typical use is upon e.g. flag.Args(), flags.StringVar(...)
package dotpath

// DotPath represents a lexically analysed filepath string.
//
// Special attention is given to slashes; especially to trailing slash(es)
// as indicator for "I mean: directory! Not file.",
// and to sequences of multiple dots; in particular tripledots "..."
// as indicators for subtree recursion, and trailing doubledots
// as indicators for "waydown" accumulation along parents/anchestors.
//
// Hint: Use New(path string) *DotPath for single entries, or
// Parse(pathNames ...string) for multiple entries.
//
// Note: A DotPath is immutable, and thus concurrency safe.
type DotPath struct {
	separator  string   // the one-char separator - must not be dot
	original   string   // the original input - never touched - String()
	volumeName string   // the VolumeName (if any)
	rootSlashs string   // leading slash(es) (if any)
	trailSlash string   // trailing slash(es) (if any)
	fullPath   string   // less any rooting                   - Name()
	lessDown   string   // less any triple ... dots           - Path
	lessTail   string   // lowest explicitly mentioned level  - Base
	goUpTail   []string // one dot plus the list of .. found in tail
	butDots    bool     // anything but dots?
}

// New returns a parsed *DotPath - separated by sep
//  Note: sep must have len == 1 and must not be dot
func New(path, sep string) *DotPath {
	if len(sep) != 1 {
		panic("DotPath: Separator must have length == 1!")
	}
	if sep == Dot {
		panic("DotPath: Separator must not be a Dot1!")
	}
	dp := new(DotPath)
	dp.separator = sep
	dp.original = noEmpty(path) // make sure it's not empty
	return dp.fill()
}

// NewPath returns a parsed *DotPath - separated by go's PathSeparator Slash `/`
func NewPath(path string) *DotPath {
	return New(path, GoPathSeparator)
}

// NewFilePath returns a parsed *DotPath - separated by the current os.PathSeparator
func NewFilePath(path string) *DotPath {
	return New(path, OsPathSeparator)
}

func (dp *DotPath) fill() *DotPath {
	dp = dp.getVolumeName()                            // VolumenName (if any)
	dp = dp.stripVolumeName()                          // fullPath = original less VolumeName (if any)
	dp = dp.stripPrefixSlashes()                       // fullPath = fullPath less leading slash(es) (if any);
	dp = dp.stripSuffixSlashes()                       // fullPath = fullPath less trailing slash(es) (if any);
	dp.butDots = false                                 // butDots  = initialised
	dp = dp.fullPathNoMultipleDots()                   // fullPath = fullPath with cleaned sequences of any multiple consecutive dots;
	dp.lessDown = dp.fullPath                          // lessDown = fullPath
	dp = dp.lessDownLessTripleDots()                   // lessDown = lessDown less any tripledots
	dp.lessTail, dp.goUpTail = dp.lessDown, []string{} // lessTail = lessDown
	dp = dp.lessTailLessDoubleDots()                   // lessTail = lessDown less trailing doubledots;

	return dp
}

func (dp *DotPath) verbosefill() *DotPath {
	dp = dp.getVolumeName()             // VolumenName (if any)
	dp = dp.stripVolumeName()           // fullPath = original less VolumeName (if any)
	println("dp.fullPath", dp.fullPath) //
	dp = dp.stripPrefixSlashes()        // fullPath = fullPath less leading slash(es) (if any);
	dp = dp.stripSuffixSlashes()        // fullPath = fullPath less trailing slash(es) (if any);
	dp.butDots = false                  // butDots  = initialised
	dp = dp.fullPathNoMultipleDots()    // fullPath = fullPath with cleaned sequences of any multiple consecutive dots;
	println("dp.fullPath", dp.fullPath, "\tbutDots", dp.butDots)
	dp.lessDown = dp.fullPath                          // lessDown = fullPath
	println("dp.lessDown", dp.lessDown)                //
	dp = dp.lessDownLessTripleDots()                   // lessDown = lessDown less any tripledots
	println("dp.lessDown", dp.lessDown, "\tless ...")  //
	dp.lessTail, dp.goUpTail = dp.lessDown, []string{} // lessTail = lessDown
	println("dp.lessTail", dp.lessTail)                //
	dp = dp.lessTailLessDoubleDots()                   // lessTail = lessDown less trailing doubledots;
	println("dp.lessTail", dp.lessTail, "\tless ..")   //

	return dp
}

// LooksLikeIsAbs - returns true if some leading Slash was found
func (dp *DotPath) LooksLikeIsAbs() bool {
	return (len(dp.rootSlashs) > 0)
}

// LooksLikeIsDir - returns true if some trailing Slash was found
func (dp *DotPath) LooksLikeIsDir() bool {
	return (len(dp.trailSlash) > 0)
}

// Separator - returns the single character separator string
func (dp *DotPath) Separator() string {
	return dp.separator
}

// String - returns the original input - skipping nothing
func (dp *DotPath) String() string {
	return dp.original
}

// PathName - returns the full analysed path - incl. any triple ... dots
//  Note: intentionally, path.Cleaning is is not applied as
//  some tripledot might disappear due to being followed by some doubledot.
func (dp *DotPath) PathName() string {
	return dp.volumeName + dp.rootSlashs + dp.fullPath + dp.trailSlash
}

// PathBase - returns the path - less any triple ... dots and trailing up .. double dots,
// (including volumename and trailing delimiter, if any).
func (dp *DotPath) PathBase() string {
	return dp.goodPath(dp.lessTail)
}

// PathText - returns the path - less any triple ... dots,
// (including volumename and trailing delimiter, if any).
func (dp *DotPath) PathText() string {
	return dp.goodPath(dp.lessDown)
}

// Path - returns the path - less any triple ... dots and any (eventual) trailing slash
// (including volumename, if any).
func (dp *DotPath) Path() string {
	return dp.clean(dp.volumeName + dp.rootSlashs + dp.lessDown)
}

// RecursePathS returns any (partial) path which ends in a Down ...
func (dp *DotPath) RecursePathS() (pathS []string) {

	for _, downPath := range dp.downPathS() {
		pathS = append(pathS, dp.goodPath(downPath))
	}
	return pathS
}

// WaydownPathS returns any (partial) path which is followed by some trailing GoUp ..
//  Note: such pathS should be considered before PathBase(), and are
//  only relevant if PathText() != PathBase(); Path() would be the first WaydownPath
func (dp *DotPath) WaydownPathS() (pathS []string) {

	for pos := len(dp.goUpTail); pos > 0; pos-- { // shrinking list of GoUp's
		pathS = append(pathS, dp.goodPath(dp.goUpPath(pos)))
	}
	return pathS
}

// PathS conveniently returns any WaydownPath and the Base.
//  Note: If there is no way down mentioned (and thus: Path and Base are same),
//  the returned list just holds this single element: the Path
func (dp *DotPath) PathS() (pathS []string) {
	pathS = append(pathS, dp.WaydownPathS()...)
	pathS = append(pathS, dp.PathBase())
	return pathS
}

// HasVolumeName - returns true if some filepath.VolumeName was found
//  Note: only relevant, if separator == OsPathSeparator, e.g. NewFilePath
func (dp *DotPath) HasVolumeName() bool {
	return (len(dp.volumeName) > 0)
}

// VolumeName - returns the VolumeName, if any
//  Note: only relevant, if separator == OsPathSeparator, e.g. NewFilePath
func (dp *DotPath) VolumeName() string {
	return dp.volumeName
}

// Intentionally we differ from Clean here and respect any trailSlash,
// as such indicates unambigously: a directory is intended to be named.
func (dp *DotPath) goodPath(pathName string) (fullName string) {
	return dp.clean(dp.volumeName+dp.rootSlashs+pathName) + dp.trailSlash
}

/*
Note: We focus on 'trailing' dots!

base/../foo/../bar/..	=> bar/.. & bar

Otherwise we get things like these, with duplicate visits:

base					=> base)
base/..					=> .)
base/../foo/			=> foo)
base/../foo/..			=> .	which we had before!)
base/../foo/../bar		=> bar
base/../foo/../bar/..	=> .	which we had before!)

root/base/../foo/../../bar/../..

root/base					 		=> root/base
root/base/..						=> root
root/base/../foo					=> root/foo
root/base/../foo/..					=> root	which we had before!)
root/base/../foo/../..				=> .
root/base/../foo/../../bar			=> bar
root/base/../foo/../../bar/..		=> .		which we had before!)
root/base/../foo/../../bar/../..	=> ..

*/
