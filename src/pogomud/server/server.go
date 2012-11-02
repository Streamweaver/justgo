// Stuff for handling the server itself.
// Playing around with example from https://gist.github.com/775526
package server

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"io"
	"net"
)

type ServerObject struct {
	Name string
	Protocol string
	Host string
	Port int
	BufferLimit int
	Database DatabaseInfo
}

type DatabaseInfo struct {
	Host string // Hostname of Database
	Port string // Port number of Database
	Name string // Name of Database
	User string // Admin Username for Database
	Pass string // Password for User above.
	Engine string // Type of Database being connected to.
}

// Loads config information from JSON file.
func NewServer() ServerObject {

	// Load json file with config information.
	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		fmt.Printf("%s\n", e)
		log.Fatal(e)
	}

	// Parse file info into data
	var data ServerObject
	e = json.Unmarshal(file, &data)
	if e != nil {
		fmt.Println("error parsing json: ", e)
		log.Fatal(e)
	}

	// Return it all
	//fmt.Printf("%+v\n", data)
	return data
}

func (server *ServerObject)Start() {
	addr, err := net.ResolveTCPAddr("tcp", net.JoinHostPort(server.Host, fmt.Sprintf("%d", server.Port)))
	if err != nil {
		log.Fatal(err)
		return
	}

	l, err := net.ListenTCP(server.Protocol, addr)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("%s server started and listening on port %d.\n", server.Name, server.Port)

	// Accept connections here.
	for {
		// Wait for a connection. 
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
			return
		}
		conn.Write([]byte("Welcome to " + server.Name + "!\n"))
		fmt.Printf("Connection made from %s\n", conn.RemoteAddr())

		// Spawn a new connection and go back to listening.
		go func(c net.Conn) {
			// Echo all incoming data.
			io.Copy(c, c)
			// Shut down the connection.
			c.Close()
		}(conn)
	}
}