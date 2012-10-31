// Basic dice rolling for an example
package main 

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 500; i++ {
		rnum := rand.Intn(6) + 1
		fmt.Printf("Dice are rolled and gives %v!\n", rnum)
	}
	
}