// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

import (
	"github.com/GoLangsam/container/ccsafe/lsm"
	"github.com/GoLangsam/container/ccsafe/tag"
)

// Friendly - interface exposed for go doc only
// shows the composition as interface
type Friendly interface {
	tag.Friendly      // via "container/.../tag/..."
	lsm.Friendly      // via "container/.../lsm"
	StringFriendly    // dot.go: Set..., Assign..., Delete...
	ChildFriendly     // children.go: lookupDot getChild
	NavigatorFriendly // navigate.go: Up Root Path DownS
	PrinterFriendly   // print.go: PrintTree
	ErrorFriendly     // => dot!	error.go
	OutputFriendly    // output.go
	GoFriendly        // dot.go
}

// StringFriendly - interface exposed for go doc only
type StringFriendly interface {
	SetableFriendly // set.go: Set/replace Content: Set SetS SetM
	AssignFriendly  // assign.go: Add/overwrite Content: Assignss AssignSs AssignMs
	UserFriendly    // add.go: AddMap AddStrings AddStringS
	DeleteFriendly  // delete.go: Delete/remove vals from Content: Deletes, DeleteS, DeleteM
	PrivacyFriendly // content.go: add addM

}
