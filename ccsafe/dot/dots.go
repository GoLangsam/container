// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

// DotS implements sort.Interface
type DotS []Friendly

// Len implements sort.Interface
func (k DotS) Len() int { return len(k) }

// Less implements sort.Interface
func (k DotS) Less(i, j int) bool { return k[i].String() < k[j].String() }

// Swap implements sort.Interface
func (k DotS) Swap(i, j int) { k[i], k[j] = k[j], k[i] }
