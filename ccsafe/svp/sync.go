// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package svp

// I love to be concurrency safe
//  I am concurrency safe by nature
//  Make me, use me, forget me ...
//  Don't worry, go happy :)
// Note: I am immutable - thus: I am concurrency safe ;-)
type ConcurrencySafe interface {
}

var _ ConcurrencySafe = New("Interface satisfied? :-)", empty)
