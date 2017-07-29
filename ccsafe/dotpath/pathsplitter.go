// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dotpath

import (
	"strings" // Note: We use only strings here! - path/filepath provides higher abstractions
)

// OneTripleAnyPairs - a helper to get e.g. '........' as '... .. .. .'
// Invariant: len(text) == len(segmentS)
func OneTripleAnyPairs(text string) (segmentS []string) {
	dotsleft := text          // what's left
	lengleft := len(dotsleft) // how many

	if lengleft > 2 { // take Triple ... once
		segmentS = append(segmentS, dotsleft[:3])
		dotsleft = dotsleft[3:]
		lengleft = lengleft - 3 // 2+1
	}

	for lengleft > 1 { // take Double ..
		segmentS = append(segmentS, dotsleft[:2])
		dotsleft = dotsleft[2:]
		lengleft = lengleft - 2 // 1+1
	}

	for lengleft > 0 { // take Single . (should be only one, if any)
		segmentS = append(segmentS, dotsleft[:1])
		dotsleft = dotsleft[1:]
		lengleft = lengleft - 1 // 0+1
	}
	return segmentS
}

// PathDotTailor returns a slice of strings each representing a valid node / element / PathBase,
// or a triple-dot as indicator for a full tree of subdirectories.
func PathDotTailor(head, tail string) []string {
	lenhead := len(head)
	lentail := len(tail)
	if lentail < 1 {
		return []string{head}
	}

	if lentail != strings.Count(tail, Dot) {
		panic("PathDotTailor: Expect tail to consist only of dots - but found:" + tail)
	}

	switch {
	case lenhead < 1:
		return OneTripleAnyPairs(tail)
	default:
		/* old version: keep one dot with head
		l := []string{head + string(tail[:1])}
		return append(l, OneTripleAnyPairs(tail[1:])...)
		*/
		return append([]string{head}, OneTripleAnyPairs(tail)...)
	}
}
