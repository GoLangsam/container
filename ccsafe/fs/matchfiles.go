package fs

// MatchFiles
// matches pathName against the Disk (via MatchDisk/Glob) and then returns only those
// files
// the base name of which matches any of the given patternlists.
// Any eventual filesystem errors are ignored and skipped over.
func MatchFiles(pathName string, patterns ...*Pattern) (filS FsFileS) {
	dS, fS, _ := MatchDisk(pathName)
	_ = dS // Folds are ignored here
	for _, f := range fS {
		if ok, _ := f.BaseMatches(patterns...); ok {
			filS = append(filS, f)
		}
	}
	return filS
}
