// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package svp implements StringValuePair - named constants
//  Note: Being immutable implies concurrency saftey.
package svp // StringValuePair

// Note: this interface is exposed for godoc - only ;-)
//
// I love to be friendly - thus: I give You a simple API!
//  Create me with Svp(name, stuff), or Set(name, stuff string),
//  or New(stuff) with an fmt.Stringer, so stuff names himself.
//
type Friendly interface {
	UserFriendly     // use.go
	ConcurrencySafe  // sync.go
	InfoFriendly     // valuetype.go
	DeepInfoFriendly // leaftype.go
}

var _ Friendly = New("Interface satisfied? :-)", empty)

var empty interface{}

type StringValuePair struct {
	k string      // my key text
	v interface{} // my value
}

// New returns a new StringValuePair, named key, containg string val
func New(key string, val interface{}) *StringValuePair {
	return &StringValuePair{key, val}
}
