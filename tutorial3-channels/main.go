package main

import (
	"fmt"
)

type Server struct {
	users map[string]string
}

func newServer() *Server {
	return &Server{
		users: make(map[string]string),
	}
}

func (s *Server) addUser(username, password string) {
	s.users[username] = password
}

func main() {
	s := newServer()
	for i := 0; i < 5; i++ {
		s.addUser(fmt.Sprintf("user%d", i), "password")
	}
	for username, password := range s.users {
		fmt.Printf("Username: %s, Password: %s\n", username, password)
	}
}
func sendMessage(msgch chan<- string, msg string) {

	msgch <- msg

}

func receiveMessage(msgch <-chan string) {
	msg := <-msgch
	fmt.Println(msg)
}
