// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package drum

import "fmt"

type VerboseType bool

var (
	Verbose VerboseType
)

func (v VerboseType) Printf(s string, a ...interface{}) {
	if v {
		fmt.Printf(s, a...)
	}
}
