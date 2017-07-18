// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package sync implements a no-op drop-in replacement for standard "sync"
//
// Let's say: You have a gr8 package. And it's concurrency safe due to import ("sync") and good use of "sync.Mutex" and friends.
//
// Problem: Now, You also want to use (some of) it's type in another thing, which is concurrency safe by itself, or does not need to be.
//
// Solution: Copy Your concurrency safe package (e.g. to some "oldway"/xyz path) and change the import ("sync") to here.
//
// Benefit: Easy and swiftly to achieve.
//
// Added benefit: Future maintenance will also be easy: just diff the "ccsafe" and the "oldway" packages - You should only see Your changed import!
//
// Further benefit: You reduce memory footprint and improve performance - given the compiler is clever, he may even optimise the no-op function calls away.
//
// Note: intentionally changes to the original code are kept to a minimum - to ease future maintenance.
package sync
