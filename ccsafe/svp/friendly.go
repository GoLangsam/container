// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package svp

// Friendly - interface exposed for go doc only
//
// I love to be friendly - thus: I give You a simple API!
//  Create me with New(name, stuff)
//
type Friendly interface {
	UserFriendly     // use.go
	InfoFriendly     // valuetype.go
	DeepInfoFriendly // leaftype.go
}
