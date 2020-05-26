package main

import (
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	config := oauth1.NewConfig(os.Getenv("INPUT_CONSUMER_KEY"), os.Getenv("INPUT_CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("INPUT_ACCESS_TOKEN"), os.Getenv("INPUT_TOKEN_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)
	// Twitter client
	client := twitter.NewClient(httpClient)
	_, _, err := client.Statuses.Update(os.Getenv("INPUT_CONTENT"), nil)
	if err != nil {
		panic(err.Error())
	}
}
