// Package newmath trival example
package newmath

import "fmt"

var TUNING int = 1000

func Sqrt(x float64) float64 {
	var z float64
	for i := 0; i < TUNING; i++ {
		z -= (z*z - x) / (2 * x)
	}
	return z
}

func SqrtLog(x float64) {
	var z float64
	for i := 0; i < TUNING; i++ {
		z -= (z*z - x) / (2 * x)
		fmt.Printf("Pass %v value is %v\n", i, z)
	}
}