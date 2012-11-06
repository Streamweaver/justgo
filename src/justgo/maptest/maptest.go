// Just a quick method to test map deletes across go routines.
// having some trouble conceptualizing maps that aren't passed as pointers
// but still seem to act like it.
package main 

import "fmt"

func mapDancer(prefix string, m map[string]int, max int) {
	for i := 0; i < max; i++ {
		key := prefix + string(i)
		m[key] = 0
	}
}

func main() {
	thisMap := make(map[string]int)
	mapDancer("one", thisMap, 10)
	fmt.Printf("Map is %d long.\n", len(thisMap))
	mapDancer("two", thisMap, 10)
	fmt.Printf("Map is %d long.\n", len(thisMap))
	mapDancer("three", thisMap, 10)
	fmt.Printf("Map is %d long.\n", len(thisMap))
}