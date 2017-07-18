// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package das // Dictionary by any for strings

import (
	"bytes"
	"testing"
	"text/template"
)

var keyBuff = bytes.NewBufferString("Test")
var keyTmpl = template.New("Test")
var keyStrg = "Test"
var keyInt8 = 4711
var keyBool = true

var newData = []string{"Foo", "Bar", "Buh", "Foo", "Bar"}
var addData = []string{"Foo", "Bar", "Buh", "Foo", "Bar"}

func Example_Assign(T *testing.T) {
	var das *Das // test also lazyInit

	das = das.Assign(keyBuff, newData...)
	das = das.Assign(keyTmpl, newData...)
	das = das.Assign(keyStrg, newData...)
	das = das.Assign(keyInt8, newData...)
	das = das.Assign(keyBool, newData...)
}

func Example_Append(T *testing.T) {
	var das *Das // test also lazyInit

	das = das.Assign(keyBuff, newData...)
	das = das.Assign(keyTmpl, newData...)
	das = das.Assign(keyStrg, newData...)
	das = das.Assign(keyInt8, newData...)
	das = das.Assign(keyBool, newData...)

	das = das.Append(keyBuff, addData...)
	das = das.Append(keyTmpl, addData...)
	das = das.Append(keyStrg, addData...)
	das = das.Append(keyInt8, addData...)
	das = das.Append(keyBool, addData...)

	das = das.Assign(keyBuff, newData...)
	das = das.Assign(keyTmpl, newData...)
	das = das.Assign(keyStrg, newData...)
	das = das.Assign(keyInt8, newData...)
	das = das.Assign(keyBool, newData...)
}

func Example_Delete(T *testing.T) {
	var das *Das // test also lazyInit

	das = das.Assign(keyBuff, newData...)
	das = das.Assign(keyTmpl, newData...)
	das = das.Assign(keyStrg, newData...)
	das = das.Assign(keyInt8, newData...)
	das = das.Assign(keyBool, newData...)

	das = das.Append(keyBuff, addData...)
	das = das.Append(keyTmpl, addData...)
	das = das.Append(keyStrg, addData...)
	das = das.Append(keyInt8, addData...)
	das = das.Append(keyBool, addData...)

	das = das.Delete(keyBuff)
	das = das.Delete(keyTmpl)
	das = das.Delete(keyStrg)
	das = das.Delete(keyInt8)
	das = das.Delete(keyBool)

	println("Len == 0 ?", das.Len())
}

func Example_Lookup(T *testing.T) {
	var das *Das // test also lazyInit

	das = das.Assign(keyBuff, newData...)
	das = das.Assign(keyTmpl, newData...)
	das = das.Assign(keyStrg, newData...)
	das = das.Assign(keyInt8, newData...)
	das = das.Assign(keyBool, newData...)

	das = das.Append(keyBuff, addData...)
	das = das.Append(keyTmpl, addData...)
	das = das.Append(keyStrg, addData...)
	das = das.Append(keyInt8, addData...)
	das = das.Append(keyBool, addData...)

	var res []string
	res = das.Lookup(keyBuff)
	res = das.Lookup(keyTmpl)
	res = das.Lookup(keyStrg)
	res = das.Lookup(keyInt8)
	res = das.Lookup(keyBool)

	println("Len == 5 ?", len(res))
}

func Example_KeyS(T *testing.T) {
	das := New()
	das = das.Assign(keyBuff, newData...)
	das = das.Assign(keyTmpl, newData...)
	das = das.Assign(keyStrg, newData...)
	das = das.Assign(keyInt8, newData...)
	das = das.Assign(keyBool, newData...)

	das = das.Append(keyBuff, addData...)
	das = das.Append(keyTmpl, addData...)
	das = das.Append(keyStrg, addData...)
	das = das.Append(keyInt8, addData...)
	das = das.Append(keyBool, addData...)

	var res []interface{}
	res = das.KeyS()
	println("Len == 3 ???", len(res))
	println("Is result sorted?", res)

}

func Example_Init(T *testing.T) {
	var das *Das // test also lazyInit

	das = das.Assign(keyBuff, newData...)
	das = das.Assign(keyTmpl, newData...)
	das = das.Assign(keyStrg, newData...)
	das = das.Assign(keyInt8, newData...)
	das = das.Assign(keyBool, newData...)

	das = das.Append(keyBuff, addData...)
	das = das.Append(keyTmpl, addData...)
	das = das.Append(keyStrg, addData...)
	das = das.Append(keyInt8, addData...)
	das = das.Append(keyBool, addData...)

	println("Len == 5 ?", das.Len())
	das = das.Init()
	println("Len == 0 ?", das.Len())

	das = das.Assign(keyBuff, newData...)
	das = das.Assign(keyTmpl, newData...)
	das = das.Assign(keyStrg, newData...)
	das = das.Assign(keyInt8, newData...)
	das = das.Assign(keyBool, newData...)

	das = das.Append(keyBuff, addData...)
	das = das.Append(keyTmpl, addData...)
	das = das.Append(keyStrg, addData...)
	das = das.Append(keyInt8, addData...)
	das = das.Append(keyBool, addData...)

	println("Len == 5 ?", das.Len())
}
