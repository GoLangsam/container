package fs

// FsFold represents a folder / directory of the file system.
// It may indicate to be (or have been) recursed into.
//  Note: FsFold is immutable, and as such safe for concurrent use.
type FsFold struct {
	fsInfo
	recurse *bool
}

// FsFoldS represents a collection (slice) of (pointers to) FsFold's.
type FsFoldS []*FsFold

// String returns the FsFoldS-slice as string.
func (f FsFoldS) String() string {
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

// ForceFold returns a fresh FsFold given a pathname.
func ForceFold(name string) *FsFold {
	return newFold(name)
}

// newFold returns a fresh unqualified FsFold
func newFold(name string) *FsFold {
	return &FsFold{*newInfo(name), nil}
}

// AsFold returns a fresh FsFold for the given FsInfo,
// or panics, if the FsInfo does not represent a Dir.
func (p *fsInfo) AsFold() *FsFold {
	if fd, ok := p.TryFold(); !ok {
		panic("AsFold: " + p.String() + " seems not to be a directory!")
	} else {
		return fd
	}
}

// TryFold returns a fresh FsFold,
// or nil and false iff not p.IsFold().
func (p *fsInfo) TryFold() (*FsFold, bool) {
	if !p.IsFold() {
		return nil, false
	} else {
		return &FsFold{*p, nil}, true
	}
}

// FileS returns all files in f matching any of the patterns in the patternlists
func (f *FsFold) FileS(patterns ...*Pattern) (FilS FsFileS) {
	return MatchFiles(f.JoinWith(MatchAny), patterns...)
}

// SubDirS returns fi and all it's (recursed) subdirectories,
// the directory tree rooted at fi, so to say.
func (f *FsFold) SubDirS() (DirS []*FsFold) {
	if !f.IsFold() {
		return DirS
	}
	dir := f.AsRecurse()
	DirS = append(DirS, dir)
	if dirInfoS, err := f.ReadDir(); err == nil {
		for _, dirInfo := range dirInfoS {
			if dirInfo.IsDir() {
				dir := Recurse(f.JoinWith(dirInfo.Name()))
				DirS = append(DirS, dir.SubDirS()...)
			}
		}
	}
	return DirS
}

// ReadDirS reads the directory named by OsPath.Path() and returns a FsInfo slice
// of directory entries sorted by filename.
func (f *FsFold) ReadDirS() (entrieS FsInfoS, err error) {
	if finfoS, err := f.ReadDir(); err == nil {
		for _, finfo := range finfoS {
			entrieS = append(entrieS, newExists(f.JoinWith(finfo.Name()), f))
		}
	}
	return entrieS, err
}

// AsNotDown returns a fresh FsFold representing a file system directory/file not to be recursed into.
func (p *fsInfo) AsNotDown() *FsFold {
	var recurse bool = false
	return &FsFold{*p, &recurse}
}

// NotDown returns a fresh FsFold given a pathname, which also indicates not to be recursed into.
func NotDown(name string) *FsFold {
	return newInfo(name).AsNotDown()
}

// AsRecurse returns a fresh FsFold representing a file system directory/file to be recursed into.
func (p *fsInfo) AsRecurse() *FsFold {
	var recurse bool = true
	return &FsFold{*p, &recurse}
}

// Recurse returns a fresh FsFold given a pathname, which also indicates to be recursed.
func Recurse(name string) *FsFold {
	return newInfo(name).AsRecurse()
}

// Recurse returns true if this folder indicates to be recursed into
func (f *FsFold) Recurse() bool {
	if f.recurse == nil {
		return false
	} else {
		return *f.recurse
	}
}

// HasRecurse returns true if this folder has a recurse indicator
func (f *FsFold) HasRecurse() bool {
	if f.recurse == nil {
		return false
	} else {
		return true
	}
}

// TabString returns the Name and the Recurse flag as a Tab terminated string
func (f *FsFold) TabString() string {
	if f.Recurse() {
		return f.String() + "\t" + "Recurse=true" + "\t"
	} else {
		return f.String() + "\t"
	}
}