package main

import (
	"fmt"
)

type Server struct {
	users  map[string]string
	userch chan string
}

func newServer() *Server {
	return &Server{
		users:  make(map[string]string),
		userch: make(chan string),
	}
}

func (s *Server) addUser(username, password string) {
	s.users[username] = password
}

func (s *Server) loop() {
	for {
		username := <-s.userch
		s.addUser(username, "default_password")
	}
}
func (s *Server) start() {
	go s.loop()
}

func main() {
	s := newServer()
	s.start()
	for i := 0; i < 5; i++ {
		go func(i int) {
			s.userch <- fmt.Sprintf("user%d", i)
		}(i)
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
