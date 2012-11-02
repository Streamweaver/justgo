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
	"pogomud/user"
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
		conn, err := l.AcceptTCP()
		if err != nil {
			log.Fatal(err)
			return
		}
		u := user.NewUser(1, conn)
		u.Write("Welcome to " + server.Name + "!\n")
		fmt.Printf("Connection for user %d made from %s\n", u.Id, u.Conn.RemoteAddr())

		// Spawn a new connection and go back to listening.
		go func(u *user.User) {
			// Echo all incoming data.
			io.Copy(u.Conn, u.Conn)
			// Shut down the connection.
			u.Conn.Close()
		}(u)
	}
}