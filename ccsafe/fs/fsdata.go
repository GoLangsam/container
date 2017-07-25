// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// FsData represents the data (as read upon it's creation) of a related FsFile
// and is intended to be used by higher packages such as FsCache.
//  Note: FsData is immutable, and as such safe for concurrent use.
type FsData struct {
	FsFile
	byteS []byte
}

// FsDataS represents a collection (slice) of (pointers to) FsData's.
type FsDataS []*FsData

// String returns the FsDataS-slice as string.
func (f FsDataS) String() string {
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

// ForceData returns a fresh FsData having given name and data.
func ForceData(name string, data []byte) *FsData {
	return newData(name, data)
}

// newData returns a fresh FsData given a name and some data,
// based on a fresh FsFile.
func newData(name string, data []byte) *FsData {
	return &FsData{*newFile(name), data}
}

// AsData returns a fresh FsData for the given FsFile,
// or panics, if TryData returns an error (from ReadFile).
func (f *FsFile) AsData() *FsData {
	if anew, ok := f.TryData(); !ok {
		panic("newFsData: ReadFile returned an error!")
	} else {
		return anew
	}
}

// TryData returns a fresh FsData for the given FsFile,
// or false (and empty byteS/data) iff ReadFile() returned an error.
func (f *FsFile) TryData() (*FsData, bool) {
	fd := &FsData{*f, []byte{}}
	data, err := fd.ReadFile()
	switch {
	case err == nil:
		fd.byteS = data
		return fd, true
	default:
		return fd, false
	}
}

// ByteS returns the cached data
func (p *FsData) ByteS() []byte {
	return p.byteS
}

// Data returns the cached data as string
func (p *FsData) Data() string {
	return string(p.byteS)
}
