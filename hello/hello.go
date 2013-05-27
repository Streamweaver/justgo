// file hello.go - from the go example pages.
package main 

import (
	"example/newmath"
	"fmt"
)

func main() {
	fmt.Printf("Hello World. Sqrt(2) = %v\n", newmath.Sqrt(2))
	newmath.SqrtLog(2)
}