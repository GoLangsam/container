// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

import (
	"fmt"
	"io"
	"os"
)

type PrinterFriendly interface {
	PrintTree(prefix ...string) *Dot // prints the tree, one line per node, each prefixed by prefix and indented with tab "\t"
}

var _ PrinterFriendly = New("Interface satisfied?")

const tab = "\t"

func prefix(pfx ...string) string {
	var indent string
	for i, pre := range pfx {
		if i == 0 {
			indent = indent + pre
		} else {
			indent = indent + tab + pre
		}
	}
	return indent
}

// PrintTree prints the tree to os.Stdout, one line per node, each prefixed by prefix and indented with tab "\t"
func (d *Dot) PrintTree(prefix ...string) *Dot {
	return d.FprintTree(os.Stdout, prefix...)
}

// FprintTree prints the tree to io.Writer, one line per node, each prefixed by prefix and indented with tab "\t"
func (d *Dot) FprintTree(w io.Writer, pfx ...string) *Dot {
	if d != nil {
		fprintTree(w, d, tab, prefix(pfx...))
	}
	return d
}

func fprintTree(w io.Writer, d *Dot, delim, indent string) {
	id := indent + delim
	d.l.RLock() // protect me, and ...
	fmt.Fprintln(w, id, d.String())
	d.l.RUnlock() // release me, let me go ...
	for _, m := range d.DownS() {
		fprintTree(w, m, delim, id)
	}
}
