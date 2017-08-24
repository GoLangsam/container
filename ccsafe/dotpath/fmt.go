// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dotpath

import (
	"fmt"
)

// Print prints a commented tab-terminated list
func (dp *DotPath) Print() {
	fmt.Println("Got:\t", dp.String(), "\t")
	rp := dp.RecursePathS()
	fmt.Println("=> Path:\t", dp.Path(), "\t")
	if len(rp) > 0 {
		fmt.Println("=> Name:\t", dp.PathName(), "\t")
		fmt.Println("Recurse:\t", rp, "\t")
	}
	wd := dp.WaydownPathS()
	if len(wd) > 0 || dp.Path() != dp.PathBase() {
		fmt.Println("=> Base:\t", dp.PathBase(), "\t")
		fmt.Println("WayDown:\t", dp.PathS(), "\t")
	}
}

// IfPrintFlagArgs prints the arguments, iff flag is true
func IfPrintFlagArgs(flag bool, args ...string) {
	if flag {
		ds := FilePathS(args...)
		fmt.Println("===============================================================================")
		for i := range ds {
			ds[i].Print()
			fmt.Println("-------------------------------------------------------------------------------")
		}
		as := ToFolds(args...)
		for i := range as {
			fmt.Println(as[i].String(), "\t", as[i].Recurse())
		}
		fmt.Println("===============================================================================")
	}
}
