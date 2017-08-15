// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
fmt.go extends list.go with:

	- l.Print()

	- l.PrintValue()
	- e.PrintValue()

	- l.PrintAtomValues()
	- e.PrintAtomValues()
*/

package list

import (
	"fmt"
)

// ===========================================================================

// Print "List=" AtomValues | Total= & lf
func (l *List) Print(args ...interface{}) {
	if l.print(args...) {
		fmt.Print("List=")
		l.root.printAtomValues()
		fmt.Print(" | ")
		fmt.Print("Total=")
		fmt.Println(l.len)
	}
}

// ===========================================================================

// PrintValue of l.Root()
func (l *List) PrintValue(args ...interface{}) {
	if l.print(args...) {
		l.Root().PrintValue()
	}
}

// PrintValue (AtomValues)
func (e *Element) PrintValue(args ...interface{}) {
	if e.print(args...) {
		e.printAtomValues()
	}
}

// ===========================================================================

// PrintAtomValues is a convenience to print a list with all elements atom values
func (l *List) PrintAtomValues(args ...interface{}) {
	if l.print(args...) {
		fmt.Print("List=")
		l.root.printAtomValues()
		fmt.Print("\t")
		// Iterate through list and print its contents.
		for e := l.Front(); e != nil; e = e.Next() {
			e.printAtomValues()
			fmt.Print(" | ")
		}
		fmt.Print("Total=")
		fmt.Println(l.len)
	}
}

// PrintAtomValues is a convenience to print the atom values of e
func (e *Element) PrintAtomValues(args ...interface{}) {
	if e.print(args...) {
		fmt.Print("Element=")
		e.printAtomValues()
	}
	fmt.Println(".")
}

// printAtomValues:
func (e *Element) printAtomValues() {
	switch ev := e.Value.(type) {
	case *ComposedValue:
		for i, x := range *ev {
			x.printAtomValues()
			if i+1 < len(*ev) {
				fmt.Print("|")
			}
		}
	default:
		fmt.Print(e.Value)
	}
}

// ===========================================================================
func (l *List) print(args ...interface{}) bool {
	printArgs(args...)
	if l == nil {
		fmt.Println("List is nil!")
		return false
	}
	return true
}

func (e *Element) print(args ...interface{}) bool {
	printArgs(args...)
	if e == nil {
		fmt.Println("Element is nil!")
		return false
	}
	return true
}

// ===========================================================================
func printArgs(args ...interface{}) {
	if args != nil {
		fmt.Print(args...)
		fmt.Print(": ")
	}
}
