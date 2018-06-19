// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package das_test // Dictionary by any for strings

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/GoLangsam/container/ccsafe/das"
)

var keyBuff = bytes.NewBufferString("Test")
var keyTmpl = template.New("Test")
var keyStrg = "Test"
var keyInt8 = 4711
var keyBool = true

var newData = []string{"Foo", "Bar", "Buh", "Faa", "Bir"}
var addData = []string{"Fuu", "Bor", "Bah", "Fee", "Ber"}

func ExampleDas() *das.Das {
	var das *das.Das // test also lazyInit

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

	return das
}

func ExampleDas_Assign() {
	das := ExampleDas() // a populated *Das

	das = das.Assign(keyBuff, newData...)
	das = das.Assign(keyTmpl, newData...)
	das = das.Assign(keyStrg, newData...)
	das = das.Assign(keyInt8, newData...)
	das = das.Assign(keyBool, newData...)
}

func ExampleDas_Append() {
	das := ExampleDas() // a populated *Das

	das = das.Append(keyBuff, addData...)
	das = das.Append(keyTmpl, addData...)
	das = das.Append(keyStrg, addData...)
	das = das.Append(keyInt8, addData...)
	das = das.Append(keyBool, addData...)
}

func ExampleDas_Das() {
	das := ExampleDas() // a populated *Das

	for key, val := range das.Das() {
		fmt.Printf("%v:\t\t\n", key)
		for v := range val {
			fmt.Printf("\t%v\t\n", v)
		}
	}
}

func ExampleDas_Delete() {
	das := ExampleDas() // a populated *Das

	fmt.Println("Len == 5 ?", das.Len())

	das = das.Delete(keyBuff)
	das = das.Delete(keyTmpl)
	das = das.Delete(keyStrg)
	das = das.Delete(keyInt8)
	das = das.Delete(keyBool)

	fmt.Println("Len == 0 ?", das.Len())
}

func ExampleDas_Fetch() {
	das := ExampleDas() // a populated *Das

	key := keyBool
	fmt.Printf("%v:\t\t\n", key)
	if vS, ok := das.Fetch(key); ok {
		for i := range vS {
			fmt.Printf("\t%v\t\n", vS[i])
		}
	}
}

func ExampleDas_Len() {
	das := ExampleDas() // a populated *Das

	fmt.Println("Len == 5 ?", das.Len())
}

func ExampleDas_Lookup() {
	das := ExampleDas() // a populated *Das

	var res []string
	res = das.Lookup(keyBuff)
	res = das.Lookup(keyTmpl)
	res = das.Lookup(keyStrg)
	res = das.Lookup(keyInt8)
	res = das.Lookup(keyBool)

	fmt.Println("Len == 5 ?", len(res))
}

func ExampleDas_KeyS() {
	das := ExampleDas() // a populated *Das

	var res []interface{}
	res = das.KeyS()
	fmt.Println("Len == 3 ???", len(res))
	fmt.Println("Is result sorted?", res)

}

func ExampleDas_Init() {
	das := ExampleDas() // a populated *Das
	fmt.Println("Len == 5 ?", das.Len())

	das = das.Init()
	fmt.Println("Len == 0 ?", das.Len())

	das = ExampleDas() // a populated *Das - again
	fmt.Println("Len == 5 ?", das.Len())
}
