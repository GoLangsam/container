// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

import (
	"fmt"
	"io"
	"os"
)

type OutputFriendly interface {
	out() io.Writer
	SetOutput(output io.Writer) *Dot // sets the destination for usage and error messages. If output is nil, os.Stderr is used.
	PutOut(msg ...interface{}) *Dot  // prints msg's on a line to out()
}

var _ OutputFriendly = New("Interface satisfied? :-)")

// as found in "flag/flag.go"

func (d *Dot) out() io.Writer {
	if d.output == nil {
		return os.Stderr
	}
	return d.output
}

// SetOutput sets the destination for usage and error messages.
// If output is nil, os.Stderr is used.
func (d *Dot) SetOutput(output io.Writer) *Dot {
	d.l.Lock()         // protect me, and ...
	defer d.l.Unlock() // release me, let me go ...
	d.output = output
	return d
}

// OutOut prints msg's on a line to out()
func (d *Dot) PutOut(msg ...interface{}) *Dot {
	d.l.RLock()         // protect me, and ...
	defer d.l.RUnlock() // release me, let me go ...
	fmt.Fprintln(d.out(), msg...)
	return d
}
