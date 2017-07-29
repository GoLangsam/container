// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dotpath

/*
type result struct {
	string
	bool
}
*/

// Parse returns a slice of pathname, bool
func Parse(args ...string) (recurlist []struct {
	string
	bool
}) {

	recurlist = make([]struct {
		string
		bool
	}, 0, len(args)*4) // initial cap: 4 per arg

	dotS := FilePathS(args...)
	for _, dotPath := range dotS {
		recurlist = append(recurlist, dotPathS(dotPath)...)
	}

	return recurlist
}

// dotPathS returns Recurse / NotDown folds from DotPath
func dotPathS(dotPath *DotPath) (recurlist []struct {
	string
	bool
}) {

	recurlist = make([]struct {
		string
		bool
	}, 0, 4) // initial cap: 4 per arg

	var waydown = make(map[string]bool)
	for _, p := range PathS() {
		waydown[p.String()] = false
	}

	var r struct {
		string
		bool
	}

	for _, p := range RecursePathS() {
		if _, ok := waydown[p]; ok {
			waydown[p.String()] = true
		} else {
			r = struct {
				string
				bool
			}{{string(p.String()), bool(true)}}
			recurlist = append(recurlist, r)
		}
	}

	for _, p := range dotPath.PathS() {
		if waydown[p.String()] {
			r = struct {
				string
				bool
			}{{p, true}}
			recurlist = append(recurlist, r)
		} else {
			r = struct {
				string
				bool
			}{{p, false}}
			recurlist = append(recurlist, r)
		}
	}
	return recurlist
}
