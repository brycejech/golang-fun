package twitter

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

func GetTweetsByUserId(userId string) []Tweet {
	method := "GET"
	url := fmt.Sprintf("%v/2/users/%v/tweets?max_results=5&tweet.fields=id,text,created_at", BASE_URL, userId)

	userTweets := doRequest[GetTweets](method, url)

	return userTweets.Data
}

func FormatTweet(tweet Tweet, handle string) string {
	const length = 100

	var strings []string

	createdAt, _ := time.Parse(time.RFC3339, tweet.CreatedAt)
	formattedTime := createdAt.Format("Jan 02 2006")
	title := fmt.Sprintf("On %v, %v tweeted:\n\n", formattedTime, handle)
	strings = append(strings, title)

	var tmpStr string
	for _, char := range tweet.Text {
		str := string(char)
		if str == "\n" {
			strings = append(strings, tmpStr)
			tmpStr = ""
		}
		tmpStr += str
		// if()
	}
	return ""
}
