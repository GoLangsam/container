// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package drum

import (
	"fmt"
	"sort"
	//	"sync"
)

// Histogram represents a named integer histrogram (with total)
type Histogram struct {
	Nam string
	Cnt int
	Map counter
	Tot int64
	Sum counter
	//	sync.Mutex
}

// NewHistogram returns a new and initialized histogram
func NewHistogram(nam string, cap int) *Histogram {
	return new(Histogram).Init(nam, cap)
}

// Init returns an initialized histogram
func (b *Histogram) Init(nam string, cap int) *Histogram {
	//	b.Lock()
	//	defer b.Unlock()

	b.Nam = nam
	b.Cnt = 0
	b.Map = make(counter, cap)
	b.Tot = 0
	b.Sum = make(counter, cap)
	return b
}

// Beat increments
func (b *Histogram) Beat(cur int) {
	//	b.Lock()
	//	defer b.Unlock()

	b.Cnt++
	if Verbose {
		b.Map[cur]++
	}
	b.Tot += int64(cur)
	if Verbose {
		b.Sum[cur] += int64(cur)
	}
}

// Sort returns the keys of b.Map in a sorted slice
func (b *Histogram) Sort() []int {
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
func (b *Histogram) Print() {
	//	b.Lock()
	//	defer b.Unlock()

	if b.Cnt < 1 { // do not print empty counter
		return
	}
	fmt.Printf("%s\t"+"% 9d\t"+"% 9d\t"+"\n", b.Nam, b.Cnt, b.Tot)
	if Verbose {
		for _, key := range b.Sort() {
			fmt.Printf("%6d\t"+"% 9d\t"+"% 9d\t"+"\n", key, b.Map[key], b.Sum[key])
		}
	}
}
