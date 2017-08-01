// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

import (
	"os"            // os.SameFile
	"path/filepath" // filepath.Glob for Glob
)

// MatchDisk uses filepath.Glob and returns all entries matching the pattern,
// separated into directories and files, as slices of FsPath, and eventual
// encountered errors, which can only be ErrBadPattern or PathError
func MatchDisk(name string) (dirS FsFoldS, filS FsFileS, err error) {
	errS := new(Errors)

	mS, err := filepath.Glob(name)
	if err != nil {
		errS.err(err)
	} else {
		for i := range mS {
			fi, err := os.Stat(mS[i])
			if err == nil {
				fs := newExists(mS[i], fi)
				if fs.IsFold() {
					dirS = append(dirS, fs.AsFold())
				} else {
					filS = append(filS, fs.AsFile())
				}
			} else {
				errS.err(err) // panic - can only be *PathError
			}
		}
	}
	return dirS, filS, errS.got()
}

// SubDirS matches pathName
// against the Disk (via MatchDisk/Glob) and then returns
// all directories below any directory matching pathname
// as a breadth first sorted slice.
// Any eventual filesystem errors are ignored and skipped over.
func SubDirS(pathName string) (dirS FsFoldS) {
	dS, fS, _ := MatchDisk(pathName)
	_ = fS // Files are ignored here
	for i := range dS {
		dirS = append(dirS, dS[i])
		dirInfoS, _ := dS[i].ReadDir()
		for j := range dirInfoS {
			pathName := filepath.Join(dS[i].String(), dirInfoS[j].Name())
			dirS = append(dirS, SubDirS(pathName)...)
		}
	}
	return dirS
}
