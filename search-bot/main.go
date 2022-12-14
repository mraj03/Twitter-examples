package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
	
}



func Twitter(_ *Credentials) (*twitter.Client, error){
	config := oauth1.NewConfig("idjt5RJxuupDcZ8DVoKoNEw5C","b5GhCqqJ8D0REbknW1qIkqSbR8QhctU6LhOZIwgbsWjLd8Rmv0")
	token := oauth1.NewToken("1399697326713688069-oA09uq6KZL2kmNA4iXdorgVC1a9Ofi","VnacAQiYlpwQolFTzMpw4vkrRlbSwLCO6NFxKYGg0xZu0")
	

	


	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		IncludeEntities: new(bool),
		SkipStatus:      twitter.Bool(true),
		IncludeEmail:    twitter.Bool(true),
	}

	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil ,err
	}
	log.Printf("User Account:\n%+v\n", user)

	return client,nil
}


func main() {
	fmt.Println("Twitter search bot v 0.01")


	creds := Credentials{
		AccessToken: os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey: os.Getenv("CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("CONSUMER_SECRET"),
	}

	fmt.Printf("%+v\n", creds)

	client, err := Twitter(&creds)
	if err != nil {
		log.Println("Error in twitter client")
		log.Println(err)
	}

	search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query:           "(from:austinkleon) new -is:retweet",
		Geocode:         "",
		Lang:            "",
		Locale:          "",
		ResultType:      "",
		Count:           10,
		SinceID:         0,
		MaxID:           0,
		Until:           "",
		Since:           "",
		Filter:          "",
		IncludeEntities: new(bool),
		TweetMode:       "",
	})
	
	if err != nil {
		log.Print(err)
	}
	
	log.Printf("%+v\n", resp)
	log.Printf("%+v\n", search)
}