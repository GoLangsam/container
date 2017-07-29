// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

// NavigatorFriendly - interface exposed for go doc only
type NavigatorFriendly interface {
	parent() *Dot  // returns the parent (or nil, if at root)
	Back() *Dot    // returns the parent (or nil, if at root)
	Up() *Dot      // returns the parent (or nil, if at root) (same as Back())
	Home() *Dot    // returns the original home - just in case You lost it :-)
	Root() *Dot    // goes all the way Up() and returns the root (of subtree)
	Path() []*Dot  // returns a slice from here up to the root
	DownS() []*Dot // returns the children as a Dot-slice, sorted ascending bykey
	Level() int    // returns the # of levels above - Root().Depth() == 0
}

var _ NavigatorFriendly = New("Interface satisfied? :-)")

// Navigators - concurrency safe

// parent returns the parent (or nil, if at root)
func (d *Dot) parent() *Dot {
	d.l.RLock()         // protect me, and ...
	defer d.l.RUnlock() // release me, let me go ...
	return d.p
}

// Up returns the parent (or nil, if at root)
func (d *Dot) Up() *Dot {
	// Note: our locking is delegated to up.parent()
	return d.parent()
}

// Back returns the parent (or nil, if at root)
func (d *Dot) Back() *Dot {
	// Note: our locking is delegated to up.parent()
	return d.parent()
}

// Home returns the original home - just in case You lost it :-)
func (d *Dot) Home() *Dot {
	d.l.RLock()         // protect me, and ...
	defer d.l.RUnlock() // release me, let me go ...
	return d.home
}

// Root goes all the way Up() and returns the root - just in case You lost it :-)
//
// Note: usually, Root == Home. Alas: operations (such as Remove subtree) may invalidate this.
func (d *Dot) Root() *Dot {
	// Note: our locking is delegated to up.parent()
	var root *Dot
	for up := d; up != nil; up = up.parent() {
		root = up
	}
	return root
}

// Path returns a slice from here up to the root
// Note: as the slice is returned bottom-up, may like to reverse it :-)
func (d *Dot) Path() []*Dot {
	// Note: our locking is delegated to up.parent()
	var ups []*Dot
	for up := d; up != nil; up = up.parent() {
		ups = append(ups, up)
	}
	return ups
}

// DownS returns the children as a Dot-slice, sorted ascending bykey
func (d *Dot) DownS() []*Dot {
	d.l.RLock()         // protect me, and ...
	defer d.l.RUnlock() // release me, let me go ...
	var dns []*Dot = make([]*Dot, 0, d.Len())
	for _, key := range d.S() { // children
		dns = append(dns, d.getChild(key))
	}
	return dns
}

// Level returns the # of levels above - Root().Depth() == 0
func (d *Dot) Level() int {
	// Note: our locking is delegated to up.parent()
	var ups int
	for up := d; up != nil; up = up.parent() {
		ups++
	}
	return ups
}

// WalkFunc is the signature of a Walk function.
type WalkFunc func(d *Dot)

// WalkBreadth1st applies a given WalkFunc wf
// in breadth-first order
func (d *Dot) WalkBreadth1st(wf WalkFunc) *Dot {
	d.l.RLock()         // protect me, and ...
	defer d.l.RUnlock() // release me, let me go ...
	wf(d)
	for _, key := range d.S() { // children
		d.getChild(key).WalkBreadth1st(wf)
	}
	return d
}

// WalkDepth1st applies a given WalkFunc wf
// in depth-first order
func (d *Dot) WalkDepth1st(wf WalkFunc) *Dot {
	d.l.RLock()                 // protect me, and ...
	defer d.l.RUnlock()         // release me, let me go ...
	for _, key := range d.S() { // children
		d.getChild(key).WalkDepth1st(wf)
	}
	wf(d)
	return d
}
