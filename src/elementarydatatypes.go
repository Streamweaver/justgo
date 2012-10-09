// file elementarydatatypes.go
// PACKAGE DEC
package main // write package names in lowercase letters.
// Expectations about package naming from java is screwnig
// me up here.  don't think I can have multiple main decs

// IMPORTS
// Good form to list package in alphabetical order

// If imports start with '.' or '/' it looks for them in the local
// directory.  Without the it looks in the Go global packages.

/*
VISIBILITY RULE

Identifiers that start uppercase is considered publicly visible i.e.
outside the package.

Lowercase indicates visivility inside the package only.

NOTE THIS IS PACKAGE LEVEL VISIBILITY
*/

import (
	fm "fmt" // alias a package import like so
	"os"
)

/* 
TYPES

These are elementary (primative) or structured (struct, array, slice, map,
channel, etc)
*/
// As with import can be single line or multi-line function call.
type (
	IZ int 
	//FZ float
	STR string
)
type T struct{}

// Constants and vars get dec'd too.
var v int = 5
const c = "C"
const (
	aI = iota
	bI = iota
	cI = iota
)

/*
INIT
*/
func init() {
}

// main initiates all other functions in this file.
func main() {
	fm.Println("hello, world from private main!") // Uppercase so it's public
	Main()
	CastStuff()
	TestIota()
	ReadGoSystmeInfo()
}

/*
FUNCTIONS

These can have a return type as well so functions are said to be typed.
Functions with multiple return types are separated by commas.
*/

// Main just shows me that func are case senastive too.
func Main() { 
	fm.Println("Hello, world from visible main!")
}

func CastStuff() {
	a := 5.0
	fm.Printf("CastStuff says a is %f\n", a)	
	b := int(a)
	fm.Printf("CastStuff says a is %f and b is %d\n", a, b)
}

// TestIota is showing me an example of how the iota var works, which confuses me.
func TestIota() {
	// Sweet jesus I'm sloppy and repeating myself
	fm.Printf("TestIota says aI is %d\n", aI)
	fm.Printf("TestIota says bI is %d\n", bI)
	fm.Printf("TestIota says cI is %d\n", cI)
}

func ReadGoSystmeInfo() {
	var goos string = os.Getenv("GOOS")
	fm.Printf("The OS is %s\n", goos)
	path := os.Getenv("PATH") // shorthand var dec
	fm.Printf("Path is %s\n", path)
	fm.Printf("Go root is %s\n", os.Getenv("GOROOT"))
}