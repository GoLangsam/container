// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tag_test

import (
	"testing"

	"github.com/GoLangsam/container/ccsafe/tag/ami"
)

func TestTagBehavesFriendly(t *testing.T) {
	var i interface{} = new(tag.TagAny)
	if _, ok := i.(tag.Friendly); !ok {
		t.Fatalf("expected %t to behave Friendly", i)
	}
}
