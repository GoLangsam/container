// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

import (
	"sort"
)

// DotS implements sort.Interface
type DotS []Friendly

func (k DotS) Len() int           { return len(k) }
func (k DotS) Less(i, j int) bool { return k[i].String() < k[j].String() }
func (k DotS) Swap(i, j int)      { k[i], k[j] = k[j], k[i] }

var _ sort.Interface = new(DotS)
