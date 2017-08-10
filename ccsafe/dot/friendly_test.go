// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot_test

import (
	"sort"
	"testing"

	"github.com/golangsam/container/ccsafe/dot"
)

func TestDotBehavesFriendly(t *testing.T) {
	var i interface{} = new(dot.Dot)
	if _, ok := i.(dot.Friendly); !ok {
		t.Fatalf("expected %t to behave Friendly", i)
	}
}

func TestDotS(t *testing.T) {
	var i interface{} = new(dot.DotS)
	if _, ok := i.(sort.Interface); !ok {
		t.Fatalf("expected %t to be sortable", i)
	}
}
