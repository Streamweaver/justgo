// Stuff for handling the server itself.
// Playing around with example from https://gist.github.com/775526
package server

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
)

type JsonObject struct {
	Server ServerInfo
}

type ServerInfo struct {
	Name string
	Protocol string
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
func LoadSettings() JsonObject {

	// Load json file with config information.
	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		fmt.Printf("%s\n", e)
		log.Fatal(e)
	}

	// Parse file info into data
	var data JsonObject
	e = json.Unmarshal(file, &data)
	if e != nil {
		fmt.Println("error parsing json: ", e)
		log.Fatal(e)
	}

	// Return it all
	fmt.Printf("%+v\n", data)
	return data
}

// func StartServer() {
// 	// Listen on TCP port 2000 on all interfaces.
// 	l, err := net.Listen("tcp", ":4201")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for {
// 		// Wait for a connection. 
// 		conn, err := l.Accept()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		// Handle the connection in a new goroutine.
// 		// The loop then returns to accepting, so that
// 		// multiple connections may be served concurrently.
// 		go func(c net.Conn) {
// 			// Echo all incoming data.
// 			io.Copy(c, c)
// 			// Shut down the connection.
// 			c.Close()
// 		}(conn)
// 	}
// }