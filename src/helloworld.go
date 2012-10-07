// file quickreference.go
/* 
This file contains examples and notes about the Go language
for reference and study.

These notes and examples are from the book "The Way To Go"
by Ivo Balbaert.

Further examples can be found on the Go language website at 
http://golang.org/ 
*/
package main // file belongs to this package.

import (
	"fmt"
	"runtime"
)

func main() { 
	helloWorld() // appease the helloworld gods.
	printGoVersion()
}

func helloWorld() {
	println("Hello", "world!") // simple print with newline
}

func printGoVersion() {
	// print with formatted text %s, %d
	fmt.Printf("Using Go version %s\n", runtime.Version())
}
