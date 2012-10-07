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

import fm "fmt" // alias a package import like so

func main() {
	fm.Println("hello, world from private main!") // Uppercase so it's public
	Main()
}

func Main() {
	fm.Println("Hello, world from visible main!")
}