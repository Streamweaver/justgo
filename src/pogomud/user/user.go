// Represents a MUD user and needed function.
package user

import (
	"time"
	"net"
)

func NewUser(i int, c *net.TCPConn) *User {
	return &User{i, c, time.Now()}
}

type User struct {
	Id int
	Conn *net.TCPConn
	Spawned time.Time
}

func (u *User) Write(s string) {
	u.Conn.Write([]byte(s))
}

func (u *User) Read() {
	// Something to read text sent?
}

var UserPool map[int]*User