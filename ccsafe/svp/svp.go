// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package svp implements StringValuePair - named constants
//  Note: Being immutable implies concurrency safetey.
package svp // StringValuePair

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

var _ Friendly = New("Interface satisfied? :-)", empty)

var empty interface{}

// StringValuePair is what package svp implements
type StringValuePair struct {
	k string      // my key text
	v interface{} // my value
}

// New returns a new StringValuePair, named key, containing string val
func New(key string, val interface{}) *StringValuePair {
	return &StringValuePair{key, val}
}
