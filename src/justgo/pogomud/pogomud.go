// file gomud/gomud.go - Entry point for GoMud
package main

import (
	"net"
	"log"
	"io"
)

func main() {
	// Listen on TCP port 2000 on all interfaces.
	l, err := net.Listen("tcp", ":4201")
	if err != nil {
		log.Fatal(err)
	}
	for {
		// Wait for a connection. 
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			// Echo all incoming data.
			io.Copy(c, c)
			// Shut down the connection.
			c.Close()
		}(conn)
	}
}