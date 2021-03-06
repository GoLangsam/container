// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dot

// PrivacyFriendly - interface exposed for go doc only
type PrivacyFriendly interface {
	add(vals ...string) *Dot            // add val as children to current content
	addM(val ...map[string]string) *Dot // add content named val: to any child named as key add content named as val[key]
}

// Value modifiers - internal - to be used with locked d

// add content named val
func (d *Dot) add(vals ...string) *Dot {
	for i := range vals {
		_ = d.getChild(vals[i]) // get key
	}
	return d
}

// addM adds children named as key and adds content named as val[key]
//
// Note: as many childs may be added, the common parent (which is me) is returned
func (d *Dot) addM(val ...map[string]string) *Dot {
	for i := range val {
		for key, v := range val[i] {
			c := d.getChild(key) // key
			_ = c.getChild(v)    // value
		}
	}
	return d
}
