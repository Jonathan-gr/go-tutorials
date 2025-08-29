package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	now := time.Now()
	respch := make(chan string, 3)
	wg := &sync.WaitGroup{}
	userID := 10

	wg.Add(3)
	go fetchUserData(userID, respch, wg)
	go fetchUserRecomendations(userID, respch, wg)
	go fetchUserLikes(userID, respch, wg)

	wg.Wait()
	close(respch)
	for res := range respch {
		fmt.Println(res)
	}

	fmt.Println(time.Since(now))
}

func fetchUserData(userID int, respch chan<- string, wg *sync.WaitGroup) {
	time.Sleep(100 * time.Millisecond)
	respch <- "user data"
	defer wg.Done()
}

func fetchUserRecomendations(userID int, respch chan<- string, wg *sync.WaitGroup) {
	time.Sleep(100 * time.Millisecond)
	respch <- "user recomendations"
	defer wg.Done()
}

func fetchUserLikes(userID int, respch chan<- string, wg *sync.WaitGroup) {
	time.Sleep(100 * time.Millisecond)
	respch <- "user likes "
	defer wg.Done()
}
