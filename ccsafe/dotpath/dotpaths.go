// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dotpath

// DotPathS - get DirPathS (each with Recurse flag) from args.
// Uses dotpath functionalities; useful e.g. for commandline arguments.
func DotPathS(args ...string) (dirS []struct {
	DirPath string
	Recurse bool
}) {
	dirS = make([]struct {
		DirPath string
		Recurse bool
	}, 0, len(args))
	dirS = append(dirS, toDirPathS(FilePathS(args...)...)...)
	return dirS
}

// toDirPath converts pathS to pathS - each with recurse flag
func toDirPathS(pathS ...*DotPath) (dirS []struct {
	DirPath string
	Recurse bool
}) {
	for i := range pathS {
		dirS = append(dirS, dotPathS(pathS[i])...)
	}
	return dirS
}

// dotPathS - pathS - each with recurse flag from given DotPath
func dotPathS(dotPath *DotPath) (dirS []struct {
	DirPath string
	Recurse bool
}) {

	var waydown = make(map[string]bool)

	dpS := dotPath.PathS()
	for i := range dpS {
		waydown[dpS[i]] = false
	}

	rpS := dotPath.RecursePathS()
	for i := range rpS {
		if _, ok := waydown[rpS[i]]; ok {
			waydown[rpS[i]] = true
		} else {
			dirS = append(dirS, struct {
				DirPath string
				Recurse bool
			}{rpS[i], true})
		}
	}

	for i := range dpS {
		if waydown[dpS[i]] {
			dirS = append(dirS, struct {
				DirPath string
				Recurse bool
			}{dpS[i], true})
		} else {
			dirS = append(dirS, struct {
				DirPath string
				Recurse bool
			}{dpS[i], false})
		}
	}

	return dirS
}
