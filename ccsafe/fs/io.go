package fs

import (
	"io/ioutil"
	"os"
)

// ReadFile reads the file named by OsPath.Path() and returns the contents.
//  Note: A successful call returns err == nil, not err == EOF.
//  Because ReadFile reads the whole file, it does not treat an EOF from Read as an error to be reported.
//  Note: for convenience, the contents is returned as string, not as []byte.
func (p *FsFile) ReadFile() ([]byte, error) {
	byteS, err := ioutil.ReadFile(p.name)
	return byteS, err
}

// WriteFile writes data to a file named by fsPath.
// If the file does not exist, WriteFile creates it with permissions Perm;
// otherwise WriteFile truncates it before writing.
func (p *FsFile) WriteFile(byteS []byte) error {
	return ioutil.WriteFile(p.name, byteS, Perm)
}

// ReadDir reads the directory named by OsPath.Path() and returns a list of directory
// entries sorted by filename.
func (p *FsFold) ReadDir() ([]os.FileInfo, error) {
	fiS, err := ioutil.ReadDir(p.name)
	return fiS, err
}

/*
func TempDir(dir, prefix string) (name string, err error)
    TempDir creates a new temporary directory in the directory dir with a name
    beginning with prefix and returns the path of the new directory. If dir is
    the empty string, TempDir uses the default directory for temporary files
    (see os.TempDir). Multiple programs calling TempDir simultaneously will not
    choose the same directory. It is the caller's responsibility to remove the
    directory when no longer needed.

func TempFile(dir, prefix string) (f *os.File, err error)
    TempFile creates a new temporary file in the directory dir with a name
    beginning with prefix, opens the file for reading and writing, and returns
    the resulting *os.File. If dir is the empty string, TempFile uses the
    default directory for temporary files (see os.TempDir). Multiple programs
    calling TempFile simultaneously will not choose the same file. The caller
    can use p.Name() to find the pathname of the file. It is the caller's
    responsibility to remove the file when no longer needed.

*/
