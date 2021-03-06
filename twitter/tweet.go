package twitter

import (
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Keys struct {
	Consumer struct {
		Key    string `yaml:"key"`
		Secret string `yaml:"secret"`
	} `yaml:"consumer"`
	Access struct {
		Key    string `yaml:"key"`
		Secret string `yaml:"secret"`
	} `yaml:"access"`
}

func PostTweet(keys Keys, title string, room_id int, imageUrl string) error {
	config := oauth1.NewConfig(keys.Consumer.Key, keys.Consumer.Secret)
	token := oauth1.NewToken(keys.Access.Key, keys.Access.Secret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Send a Tweet
	_, _, err := client.Statuses.Update("Starting live streaming: \n"+title+" https://live.bilibili.com/"+strconv.Itoa(room_id)+" "+imageUrl, nil)
	if err != nil {
		return err
	}
	return nil
}
