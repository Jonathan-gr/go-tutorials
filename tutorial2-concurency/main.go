package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	userID := 10

	data := fetchUserData(userID)
	rec := fetchUserRecomendations(userID)
	likes := fetchUserLikes(userID)

	fmt.Println(data)
	fmt.Println(rec)
	fmt.Println(likes)

	fmt.Println(time.Since(now))

}

func fetchUserData(userID int) string {
	time.Sleep(100 * time.Millisecond)
	return "user data"
}

func fetchUserRecomendations(userID int) string {
	time.Sleep(100 * time.Millisecond)
	return "user recomendations"
}

func fetchUserLikes(userID int) string {
	time.Sleep(100 * time.Millisecond)
	return "user likes "
}
