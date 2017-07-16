package fs

// FsBase represents a basename of the file system.
//  Note: FsBase is immutable, and as such safe for concurrent use.
type FsBase struct {
	fsPath // file system base name
}

// FsBaseS represents a collection (slice) of (pointers to) FsBase's.
type FsBaseS []*FsBase

// String returns the FsBaseS-slice as string.
func (f FsBaseS) String() string {
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

// ForceBase returns a fresh FsBase representing the Base(name) of the given a name.
func ForceBase(name string) *FsBase {
	return newBase(name)
}

// newBase returns a fresh FsBase representing the Base(name) of the given a name.
func newBase(name string) *FsBase {
	return &FsBase{*newPath(name)}
}

// AsBase returns a fresh FsBase for the given FsInfo,
// or panics, if path is not identical to it's own Base(name).
func (p *fsPath) AsBase() *FsBase {
	if fb, ok := p.TryBase(); !ok {
		panic("AsBase: path " + p.String() + " not equal to it's base " + p.Base().String())
	} else {
		return fb
	}
}

// TryBase returns a fresh FsBase for the given path,
// or false iff path is not identical to it's own Base(name).
func (p *fsPath) TryBase() (*FsBase, bool) {
	return &FsBase{*p}, (p.String() == p.Base().String())
}
