package fs

import (
	"os"
)

// Stat - returns the actual os.Stat() and the error received from os.Stat (of type *PathError)
//  Note: Stat does not refer to the FileInfo originally embedded into fsPath and thus may return
//  different FileInfo, if content of file system representetd by fsPath has changed.
func (p *fsPath) Stat() (os.FileInfo, error) {
	fi, err := os.Stat(p.name)
	return fi, err
}

// Accessible - returns nil, or the error received from os.Open (or Close)
func (p *fsPath) Accessible() error {
	file, err := os.Open(p.name)
	if err != nil {
		return err
	}
	return file.Close()
}

// MkDir uses os.MkdirAll to create a directory named by fsPath, along with any necessary parents,
// and returns nil, or else returns an error.
// The permission bits Perm are used for all created directories.
// If fsPath is already a directory, MkDir does nothing and returns nil.
func (p *FsFold) MkDir() error {
	return os.MkdirAll(p.name, Perm)
}

// Create creates the file named by OsPath.Path() with mode 0666 (before umask),
// truncating it if it already exists.
// If successful, methods on the returned File can be used for I/O;
// the associated file descriptor has mode O_RDWR.
// If there is an error, it will be of type *PathError.
//  Note: do not forget to defer file.Close()
func (p *FsFile) Create() (*os.File, error) {
	file, err := os.Create(p.name)
	return file, err
}

// Open opens the file named by OsPath.Path() for reading.
// If successful, methods on the returned file can be used for reading;
// the associated file descriptor has mode O_RDONLY.
// If there is an error, it will be of type *PathError.
//  Note: do not forget to defer file.Close()
func (p *FsFile) Open() (*os.File, error) {
	file, err := os.Open(p.name)
	return file, err
}
