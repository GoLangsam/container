package fs

import (
	"path/filepath" // filepath.SplitList for NewS
	"strings"
)

// fsPath represents a file system path with os related functionalities
// and makes it's methods available to the derived types of this package.
//  Note: fsPath itself is intentionally not exported.
//  Note: fsPath is immutable, and as such safe for concurrent use.
type fsPath struct {
	name string // file system path
}

// FsPathS represents a collection (slice) of (pointers to) fsPathes
type FsPathS []*fsPath

// String returns the FsPathS-slice as string.
func (f FsPathS) String() string {
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

// forcePath returns a fresh fsPath representing the the given name.
func forcePath(name string) *fsPath {
	return newPath(name)
}

// newPath returns a fresh fsPath representing some file system element (directory/file)
// with the given path, which is also filepath.FromSlash-normalised.
// regardless whether it exists, or not, or is even a pattern.
func newPath(path string) *fsPath {
	return &fsPath{filepath.FromSlash(path)}
}

// AsPath returns (a pointer to) the underlying fsPath
// or panics, if TryPath detected some invlaid content.
//  Note: AsPath exists only for symmetry with respect to the other, higher types.
func (p *fsPath) AsPath() *fsPath {
	if _, ok := p.TryPath(); !ok {
		panic("AsPath: " + p.name + " contains an invalid character such as '" + string(filepath.ListSeparator) + "' or '" + MatchAny + "' or '" + MatchOne)
	} else {
		return p
	}
}

// TryPath returns (a pointer to) the underlying fsPath
// false, iff fsPath contains
//  - any filepath.ListSeparator (= os.PathListSeparator) or
//  - any of the Match-Metacharacters (MatchAny "*" or MatchOne "?")
//  Note: Match-Metacharacters "[" and "]" are intentionally permitted;
//  they may be used not only in patterns, but also as valid name of some file or folder/directory.
func (p *fsPath) TryPath() (*fsPath, bool) {
	switch {
	case strings.ContainsAny(p.name, string(filepath.ListSeparator)+MatchAny+MatchOne):
		return p, false
	default:
		return p, true
	}
}

// NewS returns a non-empty slice of fsPath obtained via filepath.SplitList
func NewS(names ...string) (pathS FsPathS) {
	if len(names) < 1 {
		pathS = append(pathS, newPath(""))
	} else {
		for _, nameList := range names {
			for _, pathName := range filepath.SplitList(nameList) {
				pathS = append(pathS, newPath(pathName))
			}
		}
	}
	return pathS
}

// String returns the pathtext repreented by fsPath
func (p *fsPath) String() string {
	return p.name
}

// MatchDisk is a convenience for MatchDisk(name).
func (p *fsPath) MatchDisk() (dirS FsFoldS, filS FsFileS, err error) {
	return MatchDisk(p.name)
}

// PathMatches reports whether fsPath matches any of the patterns.
func (p *fsPath) PathMatches(patterns ...*Pattern) (matched bool, err error) {
	return Match(p.name, patterns...)
}

// BaseMatches reports whether base name of fsPath matches any of the patterns.
func (p *fsPath) BaseMatches(patterns ...*Pattern) (matched bool, err error) {
	return Match(p.Base().String(), patterns...)
}

// BaseLessExt: name.Base() less name.Ext()
func (p *fsPath) BaseLessExt() *FsBase {
	return BaseLessExt(p.name)
}
