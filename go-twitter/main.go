package main

import (
	"fmt"
	"go-twitter/twitter"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	handle := os.Args[1]

	user := twitter.GetTwitterUserByHandle(handle)

	if user.Id != "" {
		tweets := twitter.GetTweetsByUserId(user.Id)

		// for each tweet
		for _, tweet := range tweets {

			var titleLine []string
			createdAt, _ := time.Parse(time.RFC3339, tweet.CreatedAt)
			formattedTime := createdAt.Format("Jan 02 2006")
			title := fmt.Sprintf("On %v, %v tweeted:", formattedTime, handle)
			titleLine = append(titleLine, title)

			lines := []string{strings.Join(titleLine, " ")}

			// split the tweet text into user entered "\n" chars
			tmpLines := strings.Split(tweet.Text, "\n")

			// consume each line
			for _, line := range tmpLines {
				// if line length under threshold, let through
				if len(line) < 60 {
					lines = append(lines, line)
					continue
				}

				// Otherwise, we need to shorten it to just under 60
				tmpLine := ""

				// split the line into words
				words := strings.Split(line, " ")

				// iterate each word and build up a line until the next word would be over 60
				for i, word := range words {
					// if adding the current word passes threshold, reset
					if len(tmpLine+" "+word) > 60 {
						lines = append(lines, tmpLine)
						tmpLine = word

					} else {
						tmpLine = fmt.Sprintf("%v %v", tmpLine, word)
					}

					if i == len(words)-1 {
						lines = append(lines, tmpLine)
					}
				}
			}

			longest := 0
			for _, line := range lines {
				if len(line) > longest {
					longest = len(line)
				}
			}

			for i, line := range lines {
				if i == 0 {
					fmt.Println(line)
					fmt.Println("# ------------------------------------------------------------ #")
					continue
				}
				fmt.Printf("| %-60v |\n", strings.Trim(line, " "))
			}
			fmt.Println("# ------------------------------------------------------------ #")
			fmt.Println("\n\n")
		}
	}
}
