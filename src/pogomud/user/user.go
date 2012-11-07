// Represents a MUD user and needed function.
package user

import (
	"net"
	"log"
)

type User struct {
	Id int
	Name string
	Conn *net.TCPConn
	Incoming chan string
	Outgoing chan string
	Quit chan bool
	UserList map[int]User
}

func (u *User) Write(s string) {
	u.Conn.Write([]byte(s))
}

func (u *User) Read(buffer []byte) bool {
	_, err := u.Conn.Read(buffer)
	if err != nil {
		u.Close()
		log.Fatal(err)
		return false
	}
	return true
}

func (u *User) IsMe(other *User) bool {
	if u.Id == other.Id{
		return true
	}
	return false
}

func (u *User) Close() {
	u.Quit <- true
	u.Conn.Close()
	u.Destroy()
}

func (u *User) Destroy() {
	delete(u.UserList, u.Id)
}

// Handle User Communications.
func Broadcaster(user *User) {
	buffer := make([]byte, 2048)
	for user.Read(buffer) {
		user.Outgoing <- user.Name + ": " + string(buffer)
	}
}

func Reciever(user *User) {
    for {
        select {
            case buffer := <-user.Incoming:    
                user.Conn.Write([]byte(buffer))
            case <-user.Quit:
                user.Conn.Close()
                break
        }
    }
}

func nameSetter(conn *net.TCPConn) string {
	conn.Write([]byte("Enter a name to use: "))
	buffer := make([]byte, 2048)
	bRead, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	return string(buffer[0:bRead])
}

func UserHandler(conn *net.TCPConn, out chan string, userList map[int]User) {
	name := nameSetter(conn)
	newUser := User{
		len(userList),
		name,
		conn,
		make(chan string),
		out,
		make(chan bool),
		userList,
	}
	userList[newUser.Id] = newUser
	go Broadcaster(&newUser)
	go Reciever(&newUser)
	out <-string(name + " has connected.\n")
}