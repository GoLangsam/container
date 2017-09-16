// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot_test

import (
	"testing"

	dot "github.com/GoLangsam/container/ccsafe/tagmap"
)

func TestSimple(t *testing.T) {

}

func Example() {
	var sms *dot.Dot = dot.New("<root>")
	//	sms.PrintTree("1 >>")
	sms.Assignss("foo", "bar")
	//	sms.PrintTree("1 >>")
}
