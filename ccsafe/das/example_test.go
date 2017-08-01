// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package das // Dictionary by any for strings

import (
	"bytes"
	"fmt"
	"text/template"
)

var keyBuff = bytes.NewBufferString("Test")
var keyTmpl = template.New("Test")
var keyStrg = "Test"
var keyInt8 = 4711
var keyBool = true

var newData = []string{"Foo", "Bar", "Buh", "Foo", "Bar"}
var addData = []string{"Foo", "Bar", "Buh", "Foo", "Bar"}

func ExampleDas_Assign() {
	var das *Das // test also lazyInit

	das = das.Assign(keyBuff, newData...)
	das = das.Assign(keyTmpl, newData...)
	das = das.Assign(keyStrg, newData...)
	das = das.Assign(keyInt8, newData...)
	das = das.Assign(keyBool, newData...)
}

func ExampleDas_Append() {
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
}

func ExampleDas_Das() {
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

	for key, val := range das.Das() {
		fmt.Printf("%s:\t\t\n", key)
		for v := range val {
			fmt.Printf("\t%s\t\n", v)
		}
	}
}

func ExampleDas_Delete() {
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

	fmt.Println("Len == 0 ?", das.Len())
}

func ExampleDas_Fetch() {
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

	key := keyBool
	fmt.Printf("%s:\t\t\n", key)
	for v := range das.Fetch(key) {
		fmt.Printf("\t%s\t\n", v)
	}
}

func ExampleDas_Len() {
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

	fmt.Println("Len == 10 ?", das.Len())
}

func ExampleDas_Lookup() {
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

	fmt.Println("Len == 5 ?", len(res))
}

func ExampleDas_KeyS() {
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
	fmt.Println("Len == 3 ???", len(res))
	fmt.Println("Is result sorted?", res)

}

func ExampleDas_Init() {
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

	fmt.Println("Len == 5 ?", das.Len())
	das = das.Init()
	fmt.Println("Len == 0 ?", das.Len())

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

	fmt.Println("Len == 5 ?", das.Len())
}
