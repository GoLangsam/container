package fs

const (
	// Metacharacters for filepath.Glob pattern
	MatchAny = `*` // matches any sequence of non-Separator characters
	MatchOne = `?` // matches any single non-Separator character
	Dot      = `.` // extension separator
)
