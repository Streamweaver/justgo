// Represents a MUD user and needed function.
package user

import (
	"net"
	"log"
	"bufio"
)

type User struct {
	Id int
	Name string
	Conn *net.TCPConn
	Outgoing chan string
}

// Closes the connection and preforms anything needed with it.
func (u *User) Close() {
	//
}

// Removes the user from the server userlist.
func (u *User) Destroy() {
	//
}

// Listends to the users Outgoing channel and sends
// new values to the connection.
func MessageSender(user *User) {
	for {
		msg := <- user.Outgoing
		user.Conn.Write([]bytes(msg))
	}
}

func MessageListener(user *User) {
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
	r := bufio.NewReader(conn)
	line, err := r.ReadString(byte('\n'))
	if err != nil {
		log.Fatal(err)
	}
	
	return line
}

func HandleUser(conn *net.TCPConn, out chan string, userList map[int]User) {
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