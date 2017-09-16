// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package svp_test

import (
	"testing"

	"github.com/GoLangsam/container/ccsafe/svp"
)

func TestSvpBehavesFriendly(t *testing.T) {
	var i interface{} = new(svp.StringValuePair)
	if _, ok := i.(svp.Friendly); !ok {
		t.Fatalf("expected %t to behave Friendly", i)
	}
}
