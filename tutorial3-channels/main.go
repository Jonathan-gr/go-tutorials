package main

func main() {
	// START OMIT
	c := make(chan int)

	go func() {
		c <- 42
	}()

	v := <-c
	println(v)
	// END OMIT
}
