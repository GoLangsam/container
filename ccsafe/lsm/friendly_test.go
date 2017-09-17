// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lsm_test

import (
	"testing"

	"github.com/GoLangsam/container/ccsafe/lsm"
)

var _ lsm.Friendly = new(lsm.LazyStringerMap)

func TestLsmBehavesFriendly(t *testing.T) {
	var i interface{} = new(lsm.LazyStringerMap)
	if _, ok := i.(lsm.Friendly); !ok {
		t.Fatalf("expected %t to behave Friendly", i)
	}
}
