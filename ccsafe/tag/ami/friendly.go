// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tag

import (
	myTag "github.com/golangsam/container/ccsafe/tag"
)

// Friendly - interface exposed for go doc only
//
// I love to be friendly - thus: I give You a simple API!
//  Create me with New(name) or Tag(name, stuff),
//
type Friendly interface {
	myTag.Friendly   // inherited
	PathFriendly     // path.go
	InfoFriendly     // valuetype.go
	DeepInfoFriendly // leaftype.go
}
