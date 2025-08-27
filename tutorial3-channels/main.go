package main

func main() {

	//unbuffered channel
	//causes a deadlock
	// userch := make(chan string)

	// userch <- "hello2"
	// userch <- "hello"
	// user := <-userch
	// user2 := <-userch
	// println(user)
	// println(user2)

	// buffered channel
	userch := make(chan string, 2)
	userch <- "hello2"
	userch <- "hello"
	user := <-userch
	user2 := <-userch
	println(user)
	println(user2)

}
