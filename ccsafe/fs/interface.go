// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

import (
	"os"
)

// PathFriendly summarises many methods inherited from internal type fsPath
// in related groups / sub interfaces
//  Note: this interface is exposed not only for godoc ;-)
type PathFriendly interface {
	FilePathFriendly
	MatchFriendly
	OsFileInfoFriendly
	PathTypeFriendly
	Accessible() error
}

// FilePathFriendly - interface exported for go doc only
type FilePathFriendly interface {
	String() string

	Base() *FsBase
	Ext() *FsBase
	BaseLessExt() *FsBase
	Split() (dir, file string)
	SplitList() []string
	JoinWith(elem ...string) string
}

// MatchFriendly - interface exported for go doc only
type MatchFriendly interface {
	DiskFriendly
	Match(pattern *Pattern) (matched bool, err error)
	PathMatches(patterns ...*Pattern) (matched bool, err error)
	BaseMatches(patterns ...*Pattern) (matched bool, err error)
}

// DiskFriendly - interface exported for go doc only
type DiskFriendly interface {
	Glob() (matches []string, err error)
	MatchDisk() (dirS FsFoldS, filS FsFileS, err error)
}

// OsFileInfoFriendly - interface exported for go doc only
type OsFileInfoFriendly interface {
	// obtain os.FileInfo
	Stat() (os.FileInfo, error)
}

// PathTypeFriendly - interface exported for go doc only
type PathTypeFriendly interface {
	AsPath() *fsPath
	TryPath() (*fsPath, bool)
	AsInfo() *fsInfo
	TryInfo() (*fsInfo, error)
	AsBase() *FsBase
	TryBase() (*FsBase, bool)
	AsPattern() *Pattern
	TryPattern() (*Pattern, bool)
}

// PatternFriendly summarises methods of type Pattern
//  Note: this interface is exposed not only for godoc ;-)
type PatternFriendly interface {
	PathFriendly
}

// BaseFriendly summarises methods of type FsBase
//  Note: this interface is exposed not only for godoc ;-)
type BaseFriendly interface {
	PathFriendly
}

// InfoFriendly summarises many methods inherited from internal type fsInfo
// in related groups / sub interfaces
//  Note: this interface is exposed not only for godoc ;-)
type InfoFriendly interface {
	PathFriendly
	AsFile() *FsFile
	TryFile() (*FsFile, bool)
	AsFold() *FsFold
	TryFold() (*FsFold, bool)
	AsNotDown() *FsFold
	AsRecurse() *FsFold

	Exists() bool
	IsFold() bool

	SameFile(oi os.FileInfo) bool
	InfoEquals(oi os.FileInfo) bool
}

// FileFriendly summarises methods of type FsFile
//  Note: this interface is exposed not only for godoc ;-)
type FileFriendly interface {
	InfoFriendly
	// File Create/Open
	Create() (*os.File, error)
	Open() (*os.File, error)
	// File Read/Write
	ReadFile() ([]byte, error)
	WriteFile(byteS []byte) error
}

// FoldFriendly summarises methods of type FsFold
//  Note: this interface is exposed not only for godoc ;-)
type FoldFriendly interface {
	InfoFriendly
	MkDir() error
	ReadDir() ([]os.FileInfo, error)

	// AsNotDown() *FsFold

	Recurse() bool
	HasRecurse() bool

	TabString() string

	FileS(patterns ...*Pattern) (FilS FsFileS)
	ReadDirS() (entrieS FsInfoS, err error)
	SubDirS() (DirS []*FsFold)
}

// DataFriendly summarises methods of type FsData
//  Note: this interface is exposed not only for godoc ;-)
type DataFriendly interface {
	InfoFriendly

	ByteS() []byte
	Data() string
}

var _ PathFriendly = newPath("Interface statisifed") // fsPath

var _ InfoFriendly = newInfo("Interface statisifed") // FsInfo

var _ FileFriendly = newFile("Interface statisifed") // FsFile
//r _ FileFriendly = newInfo("Interface statisifed").AsFile() // FsFile - TODO:panics

var _ FoldFriendly = newFold("Interface statisifed") // FsFold
//r _ FoldFriendly = newInfo("Interface statisifed").AsFold() // FsFold - TODO:panics

var _ BaseFriendly = newBase("Interface statisifed")          // FsBase
var _ BaseFriendly = newInfo("Interface statisifed").AsBase() // FsBase

var _ DataFriendly = newData("Interface statisifed", []byte("Are You Sure?")) // FsData
//r _ DataFriendly = newFile("Interface statisifed").AsData() // FsData - panics
