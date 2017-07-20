// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package drum provides a simple counter (with names based on musical methaphores)
package drum

import (
	"fmt"
	"sort"
	//	"sync"
)

type counter map[int]int64

// Drum is named counter
type Drum struct {
	Nam string
	Cnt int
	Map counter
	//	sync.Mutex
}

// NewDrum returns a new and initialized drum
func NewDrum(nam string, cap int) *Drum {
	return new(Drum).Init(nam, cap)
}

// Init returns an initialized drum
func (b *Drum) Init(nam string, cap int) *Drum {
	//	b.Lock()
	//	defer b.Unlock()
	b.Nam = nam
	b.Cnt = 0
	b.Map = make(counter, cap)
	return b
}

// Beat increments by one
func (b *Drum) Beat(cur int) {
	//	b.Lock()
	//	defer b.Unlock()
	b.Cnt++
	if Verbose {
		b.Map[cur]++
	}
}

// Sort returns the keys of b.Map in a sorted slice
func (b *Drum) Sort() []int {
	//	b.Lock()
	//	defer b.Unlock()
	var keys sort.IntSlice
	for key := range b.Map {
		keys = append(keys, key)
	}
	keys.Sort() // Note: see also sort.Ints( []int )
	return keys
}

// Print prints a counter, if it's not empty, as lines of tab-terminated cells
func (b *Drum) Print() {
	//	b.Lock()
	//	defer b.Unlock()
	if b.Cnt < 1 { // do not print empty counter
		return
	}
	fmt.Printf("%s\t% 9d\t"+"\n", b.Nam, b.Cnt)
	if Verbose {
		for _, key := range b.Sort() {
			fmt.Printf("%6d\t% 9d\t"+"\n", key, b.Map[key])
		}
	}
}
