// Small script file that gives an example of channels in multiple
// go routines.
package main

import (
	"fmt"
	"time"
	)

// Broadcast to channel.
func chanBroadcast(ch chan string) {
	for {
		time.Sleep(time.Second * 3)
		ch <- "Broadcast!"
	}
}

// Listen on a channel for broadcast.
// Demonstrates blocking nature of channels.
func chanListen(ch chan string) {
	for {
		fmt.Println("... listening ...")
		c := <- ch
		fmt.Println(c)
	}
}

// Another technique uses range to assign.
// This is a great technique if you don't need to do something
// before recieving C.
func chanListenTwo(ch chan string) {
	for c := range ch {
		fmt.Println("......... HUH? ....")
		fmt.Println(c + " Two!")
	}
}

// Setup the go routines and run it all.  Pressing any key ends it.
func main() {
	fmt.Println("PRESS ANY KEY TO EXIT")

	ch := make(chan string)

	go chanBroadcast(ch)
	go chanListen(ch)
	go chanListenTwo(ch)

	var input string
	fmt.Scanln(&input)
	fmt.Println("closed...")

}