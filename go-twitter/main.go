package main

import (
	"fmt"
	"go-twitter/twitter"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	handle := os.Args[1]

	user := twitter.GetTwitterUserByHandle(handle)

	if user.Id != "" {
		tweets := twitter.GetTweetsByUserId(user.Id, 10)

		// for each tweet
		for _, tweet := range tweets {
			text := twitter.FormatTweet(tweet, handle)

			fmt.Print("\n", text, "\n")
		}
	}
}
