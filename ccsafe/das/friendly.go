// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package das

// Friendly - interface exposed for go doc only
//
// I love to contain strings, identified by 'anything'.
//	And I love to give them back to You, when You need 'em.
//	And also as slice or map - as You need 'em.
//	And also sorted, or reversed, all for Your convenience.
//
type Friendly interface {
	UserFriendly        // use.go
	PerformanceFriendly // lazy.go
}
