// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

import (
	"errors"
)

// ErrorFriendly - interface exposed for go doc only
type ErrorFriendly interface {
	SeeError(sender, place string, err error) bool
	SeeNotOk(sender, place string, ok bool, complain string) bool
}

// ErrorName is the name of a node-type error
const ErrorName = "Error"

// ErrorID is the ID of of a node of type error
const ErrorID = ":" + ErrorName + ":"

var _ ErrorFriendly = New("Interface satisfied? :-)")

type notOkError error

func notOk(text string) notOkError {
	var n notOkError = errors.New(text)
	return n
}

func (d *Dot) dotError() *Dot {
	return d.getChild(ErrorID)
}

func (d *Dot) dotNameTag(text string, tag interface{}) *Dot {
	c := d.getChild(text)
	c.Tag(tag)
	return c
}

func (d *Dot) dotErrorErr(sender, place string, err error) *Dot {
	msg := sender + " encountered: " + err.Error() + " @ " + place
	d.PutOut(msg)
	_ = d.dotNameTag(msg, err)
	return d
}

func (d *Dot) dotErrorNok(sender, place string, err error) *Dot {
	msg := sender + " found: " + err.Error() + " @ " + place
	d.PutOut(msg)
	_ = d.dotNameTag(msg, err)
	return d
}

// Helpers to handle an Error resp. NotOk complaint consistently

// SeeError -
// If err!=nil, attach an error below "Error:" of d, and return true
func (d *Dot) SeeError(sender, place string, err error) bool {
	if err == nil {
		return false
	}

	_ = d.dotError().dotErrorErr(sender, place, err)
	d.Tag("") // invalidate value of d

	return true
}

// SeeNotOk -
// If ok!=true, attach a not-ok complain below "Error:" of d, and return true
func (d *Dot) SeeNotOk(sender, place string, ok bool, complain string) bool {
	if ok == true {
		return false
	}

	_ = d.dotError().dotErrorNok(sender, place, notOk(complain))
	// do *NOT* invalidate value of d

	return true
}
