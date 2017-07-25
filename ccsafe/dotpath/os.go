// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dotpath

import (
	"os"
	"path"
	"path/filepath"
)

const (
	// GoPathSeparator is the pathseparator used (but not exported) by standard package `path`
	GoPathSeparator = `/`
	// OsPathSeparator is the path-list separator `os.PathSeparator`
	OsPathSeparator = string(os.PathSeparator)
)

func init() { // some paranoid sanity checks ;-)
	if Dot != path.Clean(Empty) {
		panic("My dot '" + Dot + "' differs from '" + path.Clean(Empty) + "' = path.Clean(Empty)")
	}

	if Dot != filepath.Clean(Empty) {
		panic("My dot '" + Dot + "' differs from '" + filepath.Clean(Empty) + "' = filepath.Clean(Empty)")
	}

	if GoPathSeparator != path.Clean(GoPathSeparator+GoPathSeparator+GoPathSeparator+GoPathSeparator+GoPathSeparator+GoPathSeparator+GoPathSeparator+GoPathSeparator+GoPathSeparator+GoPathSeparator+GoPathSeparator+GoPathSeparator+GoPathSeparator) {
		panic("My slash '" + GoPathSeparator + "' differs from '" + path.Clean(GoPathSeparator+GoPathSeparator) + "' = path.Clean(GoPathSeparator+GoPathSeparator+...)")
	}

	if OsPathSeparator != filepath.Clean(OsPathSeparator+OsPathSeparator+OsPathSeparator+OsPathSeparator+OsPathSeparator+OsPathSeparator+OsPathSeparator+OsPathSeparator+OsPathSeparator+OsPathSeparator+OsPathSeparator+OsPathSeparator+OsPathSeparator) {
		panic("My slash '" + OsPathSeparator + "' differs from '" + filepath.Clean(OsPathSeparator+OsPathSeparator) + "' = filepath.Clean(OsPathSeparator+OsPathSeparator+...)")
	}

}

func (dp *DotPath) clean(name string) string {
	switch {
	case dp.separator == OsPathSeparator:
		return filepath.Clean(name)
	case dp.separator == GoPathSeparator:
		return path.Clean(name)
	default:
		return name
	}
}

// getVolumeName - returns the volumeName, if any
func (dp *DotPath) getVolumeName() *DotPath {
	switch {
	case dp.separator == OsPathSeparator:
		dp.volumeName = filepath.VolumeName(dp.original)
	default:
		dp.volumeName = Empty
	}
	return dp
}

// helper to avoid empty elements
func noEmpty(pathName string) string {
	switch {
	case len(pathName) == 0:
		return path.Clean(pathName)
	default:
		return pathName
	}
}

// PathS returns a non-empty slice of NewPath
//  Note: each pathName is not split any further -filepath.SplitList is not applied-, but
//  is normalised via filepath.ToSlash
func PathS(pathNames ...string) (pathS []DotPath) {
	if len(pathNames) < 1 {
		pathS = append(pathS, *NewPath(""))
	} else {
		for _, pathName := range pathNames {
			pathName = filepath.ToSlash(pathName)
			pathS = append(pathS, *NewPath(pathName))
		}
	}
	return pathS
}

// FilePathS returns a non-empty slice of NewFilePath
//  Note: each pathList is split via filepath.SplitList,
//  is expanded against the os environmant, and
//  is normalised via filepath.FromSlash
func FilePathS(pathLists ...string) (filePathS []*DotPath) {
	if len(pathLists) < 1 {
		filePathS = append(filePathS, NewFilePath(""))
	} else {
		for _, pathList := range pathLists {
			for _, pathName := range filepath.SplitList(pathList) {
				pathName = os.ExpandEnv(pathName)
				pathName = filepath.FromSlash(pathName)
				filePathS = append(filePathS, NewFilePath(pathName))
			}
		}
	}
	return filePathS
}

/*
// FilepathIsAbs - returns the answer of filepath.IsAbs
//  Note: only relevant, if separator == OsPathSeparator, e.g. NewFilePath
func (dp *DotPath) FilepathIsAbs() bool {
	return filepath.IsAbs(dp.Path())
}
*/
