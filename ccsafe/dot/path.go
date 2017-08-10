// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

import (
	"path"
)

// PathBase - func path.Base(path string) string
func (d *Dot) PathBase() string {
	return path.Base(d.String())
}

// PathClean - func path.Clean(path string) string
func (d *Dot) PathClean() string {
	return path.Clean(d.String())
}

// PathDir - func path.Dir(path string) string
func (d *Dot) PathDir() string {
	return path.Dir(d.String())
}

// PathExt - func path.Ext(path string) string
func (d *Dot) PathExt() string {
	return path.Ext(d.String())
}

// PathIsAbs - func path.IsAbs(path string) bool
func (d *Dot) PathIsAbs() bool {
	return path.IsAbs(d.String())
}

// PathMatch - func path.Match(pattern, name string) (matched bool, err error)
func (d *Dot) PathMatch(pattern string) (matched bool, err error) {
	return path.Match(pattern, d.String())
}

// PathJoin - func path.Join(elem ...string) string
//  Note: joins elem... to Dot
func (d *Dot) PathJoin(elem ...string) string {
	var es = make([]string, len(elem)+1)
	es = append(es, d.String())
	es = append(es, elem...)
	return path.Join(es...)
}

// PathJoinThese - func path.Join(elem ...string) string
//  Note: joins elem... - does not consider Dot
func (d *Dot) PathJoinThese(elem ...string) string {
	return path.Join(elem...)
}

// func Split(path string) (dir, file string)
//  Note: not useful for Dot

// PathDown returns the path from Root to here as a path.Join'ed string.
// Note: Intentionally, path.Clean is neither applied to components nor to the result, as this might make components such as ".." 'disappear';
// but path.Clean is applied to empty nodes in order to gain a "." being joined!
func (d *Dot) PathDown() string {
	var dwnpth string
	dpS := d.Path()
	for i := range dpS {
		curpth := dpS[i].String()
		if curpth == "" {
			curpth = path.Clean(curpth)
		}
		dwnpth = path.Join(dpS[i].String(), curpth)
	}
	return dwnpth
}
