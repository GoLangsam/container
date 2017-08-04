// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package das_test // Dictionary by any for strings

import (
	"testing"

	"github.com/golangsam/container/ccsafe/das"
)

func TestDasBehavesFriendly(t *testing.T) {
	var i interface{} = new(das.Das)
	if _, ok := i.(das.Friendly); !ok {
		t.Fatalf("expected %t to behave Friendly", i)
	}
}
