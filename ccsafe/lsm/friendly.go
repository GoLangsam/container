// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lsm

// Friendly - interface exposed for go doc only
//
// I love to contain strings, named strings, named things, things which can name themselves.
//	And I love to give them back to You, when You need 'em.
//	And also their names -as slice or map- as You need 'em.
//	And also sorted, or reversed, all for Your convenience.
//
type Friendly interface {
	AccessFriendly      // gets.go
	UserFriendly        // use.go
	PerformanceFriendly // lazy.go
}
