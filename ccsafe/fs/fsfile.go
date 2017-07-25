// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// FsFile represents a file of the file system.
//  Note: FsFile is immutable, and as such safe for concurrent use.
type FsFile struct {
	fsInfo
}

// FsFileS represents a collection (slice) of (pointers to) FsFile's.
type FsFileS []*FsFile

// String returns the FsFileS-slice as string.
func (f FsFileS) String() string {
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

// ForceFile returns a fresh FsFile having given name.
func ForceFile(name string) *FsFile {
	return newFile(name)
}

// newFile returns a fresh FsFile having given name.
func newFile(name string) *FsFile {
	return &FsFile{*newInfo(name)}
}

// AsFile returns a fresh FsFile for the given FsInfo,
// or panics, if the FsInfo represents a Dir.
func (p *fsInfo) AsFile() *FsFile {
	if fd, ok := p.TryFile(); !ok {
		panic("AsFile: " + p.String() + " seems to be a directory!")
	} else {
		return fd
	}
}

// TryFile returns a fresh FsFile,
// or nil and false iff fi.IsFold().
func (p *fsInfo) TryFile() (*FsFile, bool) {
	switch {
	case p.IsFold():
		return nil, false
	default:
		return &FsFile{*p}, true
	}
}
