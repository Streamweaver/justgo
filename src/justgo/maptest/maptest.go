// Just a quick method to test map deletes across go routines.
// having some trouble conceptualizing maps that aren't passed as pointers
// but still seem to act like it.
package main 

import (
	"fmt"
	"time"
	"math/rand"
	)

func mapDancer(prefix string, m map[string]int, ch chan string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * time.Duration(rand.Intn(10)))
		key := fmt.Sprintf("%s%d", prefix, i)
		value := i * 10
		m[key] = value
		ch <- fmt.Sprintf("Wrote %s with value %d", key, value)
		fmt.Printf("--->thisMap has %d elements.\n", len(m))
	}
}

func printChan(ch chan string) {
	for {
		msg := <- ch
		fmt.Println(msg)
	}
}

func main() {
	thisMap := make(map[string]int)
	ch := make(chan string)

	go mapDancer("Bob", thisMap, ch)
	go mapDancer("John", thisMap, ch)
	go mapDancer("Tom", thisMap, ch)

	go printChan(ch)

	var name string
	fmt.Scanln(&input)

	fmt.Printf("thisMap has %d elements.\n", len(thisMap))
}