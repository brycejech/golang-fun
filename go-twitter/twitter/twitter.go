package twitter

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type TokenResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

type TwitterUser struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type GetTwitterUserByUsername struct {
	Data TwitterUser `json:"data"`
}

type Tweet struct {
	Id        string `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

type GetTweets struct {
	Data []Tweet `json:"data"`
}

const BASE_URL = "https://api.twitter.com"

func getStringOfLen(str string, len int) string {
	var final string
	for i := 0; i < len; i++ {
		final += str
	}
	return final
}

func getLongestString(strings []string) (longest string, length int) {
	longestStr := ""
	longestLen := 0

	for _, str := range strings {
		if len(str) > longestLen {
			longestLen = len(str)
			longestStr = str
		}
	}

	return longestStr, longestLen
}

func getBasicAuthToken() string {
	var apiKey string = os.Getenv("API_KEY")
	var apiSecret string = os.Getenv("API_SECRET")

	authKey := fmt.Sprintf("%v:%v", apiKey, apiSecret)
	encodedKey := base64.StdEncoding.EncodeToString([]byte(authKey))

	return encodedKey
}

func doRequest[T any](method string, url string) T {
	token := GetBearerToken()
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var data T
	err = json.Unmarshal(body, &data)

	if err != nil {
		log.Fatal(err)
	}

	return data
}

func GetBearerToken() string {
	var bearerToken string = os.Getenv("BEARER_TOKEN")
	if bearerToken != "" {
		return bearerToken
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%v/oauth2/token?grant_type=client_credentials", BASE_URL), nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Basic %v", getBasicAuthToken()))

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var bodyJSON TokenResponse
	err = json.Unmarshal(body, &bodyJSON)

	if err != nil {
		log.Fatal(err)
	}

	return bodyJSON.AccessToken
}

func GetTwitterUserByHandle(handle string) TwitterUser {
	method := "GET"
	url := fmt.Sprintf("%v/2/users/by/username/%v", BASE_URL, handle)

	userData := doRequest[GetTwitterUserByUsername](method, url)

	return userData.Data
}

func GetTweetsByUserId(userId string, max int) []Tweet {
	method := "GET"
	url := fmt.Sprintf("%v/2/users/%v/tweets?max_results=%v&tweet.fields=id,text,created_at", BASE_URL, userId, max)

	userTweets := doRequest[GetTweets](method, url)

	return userTweets.Data
}

func FormatTweet(tweet Tweet, handle string) string {
	const MAX_LEN = 65

	var final []string

	var titleLine []string
	createdAt, _ := time.Parse(time.RFC3339, tweet.CreatedAt)
	formattedTime := createdAt.Format("Jan 02 2006")
	title := fmt.Sprintf("On %v, %v tweeted:", formattedTime, handle)
	titleLine = append(titleLine, title)

	lines := []string{strings.Join(titleLine, " ")}

	// split the tweet text into user entered "\n" chars
	tmpLines := strings.Split(tweet.Text, "\n")

	for _, line := range tmpLines {
		// if line length under threshold, let through
		if len(line) < MAX_LEN {
			lines = append(lines, line)
			continue
		}

		// Otherwise, we need to shorten it to just under 60
		tmpLine := ""
		words := strings.Split(line, " ")

		for i, word := range words {
			// if adding the current word passes threshold, reset
			if len(tmpLine+" "+word) > MAX_LEN {
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

	_, longest := getLongestString(lines)

	FINAL_LEN := MAX_LEN
	if longest < MAX_LEN {
		FINAL_LEN = longest
	}

	sep := fmt.Sprintf("# %v #", getStringOfLen("-", FINAL_LEN))
	for i, line := range lines {
		if i == 0 {
			final = append(final, line /* title */, sep)
			continue
		}
		tweetLine := fmt.Sprintf("| %-*v |", FINAL_LEN, strings.Trim(line, " "))
		final = append(final, tweetLine)
	}
	final = append(final, sep)

	return strings.Join(final, "\n")
}
