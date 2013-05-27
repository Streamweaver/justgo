// A simple set of scripts to test filehashing and the general results.
package main

import (
"fmt"
"os"
"crypto/sha1"
)

// Test the output of some hashes to assess use in fixity checks.
func main() {
	fi, err := os.Open("/Users/swt8w/Desktop/PresentationDryRun.mp4")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	hsh := sha1.New()
	hsh.Write([]byte("His money is twice tainted: 'taint yours and 'taint mine."))
	byteSum := hsh.Sum(nil)
	fmt.Printf("%x", byteSum) // Convert to base16 on formatting.
}