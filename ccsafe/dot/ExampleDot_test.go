// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot_test

import (
	"testing"

	"github.com/golangsam/container/ccsafe/dot"
)

func TestSimple(t *testing.T) {

}

func TestDot(t *testing.T) {
	var sms = dot.New("<root>")
	sms.PrintTree("1 >>")
	sms.Assignss("foo", "bar")
	sms.PrintTree("2 >>")

	// Output:
	// <root>
	// <root>
	// 	foo
	// 	bar
}
