// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

const (
	// Metacharacters for filepath.Glob pattern

	// MatchAny matches any sequence of non-Separator characters
	MatchAny = `*`

	// MatchOne matches any single non-Separator character
	MatchOne = `?`

	// Dot is the extension separator - a period.
	Dot = `.`
)
