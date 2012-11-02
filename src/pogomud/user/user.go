// Represents a MUD user and needed function.
package user

import (
	"time"
	"net"
)

type User struct {
	Id int
	LastConnection time.Date
	conn *net.TCPConn
}

var UserPool map[int]*User