// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

import (
	"path/filepath"
)

// JoinWith joins f with any number of path elements into a single path,
// adding a Separator if necessary.
// Join calls Clean on the result; in particular, all empty strings are ignored.
// On Windows, the result is a UNC path if and only if the first path element is a UNC path.
func (p *fsPath) JoinWith(elem ...string) string {
	e := make([]string, 0, len(elem)+1)
	e = append(e, p.name)
	e = append(e, elem...)
	return filepath.Join(e...)
}

// Split splits path immediately following the final Separator, separating it
// into a directory and file name component. If there is no Separator in path,
// Split returns an empty dir and file set to path.
// The returned values have the property that path = dir+file.
func (p *fsPath) Split() (dir, file string) {
	return filepath.Split(p.name)
}

// SplitList splits a list of paths joined by the OS-specific ListSeparator,
// usually found in PATH or GOPATH environment variables.
// Unlike strings.Split, SplitList returns an empty slice when passed an empty string.
func (p *fsPath) SplitList() []string {
	return filepath.SplitList(p.name)
}

// Base returns the last element of path.
// Trailing path separators are removed before extracting the last element.
// If the path is empty, Base returns ".".
// If the path consists entirely of separators, Base returns a single separator.
func (p *fsPath) Base() *FsBase {
	return newBase(filepath.Base(p.name))
}

// Ext returns the file name extension used by path.
// The extension is the suffix beginning at the final dot in the final element of path;
// it is empty if there is no dot.
func (p *fsPath) Ext() *FsBase {
	return newBase(filepath.Ext(p.name))
}

// Match reports whether name matches the shell file name pattern.
func (p *fsPath) Match(pattern *Pattern) (matched bool, err error) {
	return filepath.Match(pattern.String(), p.name)
}

// Glob returns the names of all files matching the pattern (=PathText()),
// or nil if there is no matching entry(file/directory).
//  Note. The pattern may describe hierarchical names such as
//  /usr/*/bin/ed (assuming the Separator is '/').
//
// Glob ignores file system errors such as I/O errors reading directories.
// The only possible returned error is ErrBadPattern, when pattern is malformed.
//
//  Note: If fsPath does not exist, Glob may return matches, if fsPath represents a pattern.
//  If fsPath exists, Glob should return exactly one match: itself (unless Ospwd has changed).
func (p *fsPath) Glob() (matches []string, err error) {
	matches, err = filepath.Glob(p.name)
	return matches, err
}
