package fs

import (
	"os"
)

const (
	ListSep = os.PathListSeparator
)

var Perm os.FileMode = 0644 // default os.FileMode

func init() { // some paranoid sanity checks ;-)
}
