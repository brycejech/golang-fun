package main

import (
	"fmt"
	"go-twitter/twitter"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	handle := os.Args[1]

	user := twitter.GetTwitterUserByHandle(handle)

	if user.Id != "" {
		tweets := twitter.GetTweetsByUserId(user.Id)

		for _, tweet := range tweets {
			fmt.Println("\n\n-------------------------------------")
			createdAt, _ := time.Parse(time.RFC3339, tweet.CreatedAt)
			formattedTime := createdAt.Format("Jan 02 2006")
			fmt.Printf("On %v, %v tweeted:\n\n", formattedTime, handle)
			fmt.Println(tweet.Text)
			fmt.Println("-------------------------------------")
		}
	}
}
