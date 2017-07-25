// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

import (
	"os" // os.SameFile
)

// fsInfo represents a file system object including it's os.Info as found upon creation.
// and makes it's methods available to the derived types of this package.
//  Note: fsInfo itself is intentionally not exported.
//  Note: fsInfo is immutable, and as such safe for concurrent use.
type fsInfo struct {
	fsPath           // file system path - inherits many methods
	os.FileInfo      // file info
	exists      bool // true, if verified to exist on disk
}

// FsInfoS represents a collection (slice) of (pointers to) FsInfo's.
type FsInfoS []*fsInfo

// String returns the FsInfoS-slice as string.
func (f FsInfoS) String() string {
	var s string
	s = s + "{"
	first := true
	for _, e := range f {
		if first {
			first = false
		} else {
			s = s + ", "
		}
		s = s + e.String()
	}
	s = s + "}"
	return s
}

// forceInfo returns a fresh fsInfo representing the the given name.
func forceInfo(name string) *fsInfo {
	return newInfo(name)
}

// newInfo returns a fresh FsInfo,
// based on a fresh fsPath and a fresh os.FileInfo.
func newInfo(path string) *fsInfo {
	fi, _ := newPath(path).TryInfo()
	return fi
}

// AsInfo returns a fresh FsInfo, and the error received from os.Stat() (if any),
// or panics, if TryInfo returns an error (from os.Stat()).
func (p *fsPath) AsInfo() *fsInfo {
	anew, err := p.TryInfo()
	switch {
	case err == nil:
		panic("newInfo: Stat returned Error: " + err.Error())
	default:
		return anew
	}
}

// TryInfo returns a fresh fsInfo, and the error received from os.Stat() (if any)
func (p *fsPath) TryInfo() (*fsInfo, error) {
	finfo, err := p.Stat()
	switch {
	case err == nil:
		return &fsInfo{*p, finfo, true}, nil
	default:
		return &fsInfo{*p, *new(os.FileInfo), false}, err
	}
}

// newExists returns a new OsPath representing an existing file system element (directory/file)
func newExists(name string, fi os.FileInfo) *fsInfo {
	fp := newPath(name)
	return &fsInfo{*fp, fi, true}
}

// IsFold returns IsDir, if it is known to exist, and false otherwise.
func (p *fsInfo) IsFold() bool {
	switch {
	case p.Exists():
		return p.IsDir()
	default:
		return false
	}
}

// Exists returns true, if OsPath is known to represent a real disk element
// which already exists on disk
func (p *fsInfo) Exists() bool {
	return (p.exists)
}

// SameFile reports whether f and oi describe the same file.
// For example, on Unix this means that the device and inode fields
// of the two underlying structures are identical;
// on other systems the decision may be based on the path names.
// SameFile only applies to results returned by os.Stat.
// SameFile returns false in other cases.
func (p *fsInfo) SameFile(oi os.FileInfo) bool {
	return os.SameFile(p, oi)
}

// InfoEquals returns true, if all FileInfo data is same.
func (p *fsInfo) InfoEquals(oi os.FileInfo) bool {
	return (p.Name() == oi.Name() &&
		p.Size() == oi.Size() &&
		p.Mode() == oi.Mode() &&
		p.IsDir() == oi.IsDir() &&
		p.ModTime() == oi.ModTime() &&
		p.Sys() == oi.Sys())
}

// AllDirs returns true if and only if for all elements IsDir is true.
func (f FsInfoS) AllDirs() bool {
	for _, fi := range f {
		if !fi.IsDir() {
			return false
		}
	}
	return true
}
